package simple

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/mattn/go-mastodon"
	"github.com/microcosm-cc/bluemonday"
)

// Debug デバッグモード
const Debug = true

// デバッグ用の表示
func logDebug(s string) {
	if !Debug {
		return
	}
	log.Printf("debug: %s", s)
}

// デバッグ用の値ダンプ
func spewDump(s interface{}) {
	if !Debug {
		return
	}
	spew.Dump(s)
}

// Mastodon Mastodon API サンプル
type Mastodon struct {
	c                  *MastodonConfig
	client             *mastodon.Client
	Ready              bool // 準備ができている
	lastToot           string
	AccountCurrentUser *mastodon.Account // 本人の情報
}

// MastodonConfig NewMastodon設定
type MastodonConfig struct {
	Server     string
	Email      string
	Password   string
	ClientName string
	ClientFile string
}

// ClientConfigs クライアント情報
type ClientConfigs struct {
	Tokens map[string]ClientTokens `json:"tokens"`
}

// ClientTokens クライアントキーペア
type ClientTokens struct {
	ClientID     string `json:"id"`
	ClientSecret string `json:"secret"`
}

// NewMastodon Mastodon API サンプル
func NewMastodon(c *MastodonConfig) *Mastodon {
	t := new(Mastodon)
	t.c = c
	t.Ready = false
	return t
}

// Connect マストドンへ接続
func (t *Mastodon) Connect() (err error) {
	ctx := context.Background()

	// クライアント情報の初期値
	ccs := &ClientConfigs{
		Tokens: make(map[string]ClientTokens),
	}

	// クライアント設定ファイルがなければロード
	if _, err := os.Stat(t.c.ClientFile); !os.IsNotExist(err) {
		if err := t.loadClientFile(ccs); err != nil {
			return err
		}
	}
	// 該当サーバのクライアント情報がなければ取得
	if _, ok := ccs.Tokens[t.c.Server]; !ok {
		app, err := mastodon.RegisterApp(ctx, &mastodon.AppConfig{
			Server:     fmt.Sprintf("https://%s/", t.c.Server),
			ClientName: t.c.ClientName,
			Scopes:     "read write follow",
		})
		if err != nil {
			return err
		}
		ccs.Tokens[t.c.Server] = ClientTokens{
			ClientID:     app.ClientID,
			ClientSecret: app.ClientSecret,
		}
		// クライアント設定ファイルに保存
		if err := t.saveClientFile(ccs); err != nil {
			return err
		}
	}
	// マストドンクライアント
	t.client = mastodon.NewClient(&mastodon.Config{
		Server:       fmt.Sprintf("https://%s/", t.c.Server),
		ClientID:     ccs.Tokens[t.c.Server].ClientID,
		ClientSecret: ccs.Tokens[t.c.Server].ClientSecret,
	})
	// 認証
	if err := t.client.Authenticate(ctx, t.c.Email, t.c.Password); err != nil {
		return err
	}
	// 自分のアカウント情報を取得
	account, err := t.client.GetAccountCurrentUser(ctx)
	if err != nil {
		return err
	}
	t.AccountCurrentUser = account
	// 準備万端
	t.Ready = true
	return nil
}

// クライアント情報ファイルを保存
func (t *Mastodon) saveClientFile(cc *ClientConfigs) (err error) {
	buf, err := json.Marshal(cc)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(t.c.ClientFile, buf, 0644)
	if err != nil {
		return
	}
	logDebug("Save ClientFile")
	return nil
}

// クライアント情報ファイルを呼出
func (t *Mastodon) loadClientFile(cc *ClientConfigs) (err error) {
	buf, err := ioutil.ReadFile(t.c.ClientFile)
	if err != nil {
		return
	}
	err = json.Unmarshal(buf, cc)
	if err != nil {
		return
	}
	logDebug("Load ClientFile")
	return nil
}

// Toot トゥートする
func (t *Mastodon) Toot(s string) error {
	if !t.Ready {
		return nil
	}
	ctx := context.Background()
	toot := mastodon.Toot{Status: s}
	_, err := t.client.PostStatus(ctx, &toot)
	if err != nil {
		return err
	}
	logDebug(fmt.Sprintf("Toot: %s", toot.Status))
	return nil
}

// TootWithAttachment メディアをアップロードしてトゥートする
func (t *Mastodon) TootWithAttachment(s string, fn string) error {
	if !t.Ready {
		return nil
	}
	ctx := context.Background()
	log.Println("info: Uploading...")

	at,err := t.client.UploadMedia(ctx,fn)
	if err != nil {
		return err
	}

	spewDump(at)

	toot := mastodon.Toot{
		Status: s,
		MediaIDs: []mastodon.ID{ at.ID },
	}

	_, err = t.client.PostStatus(ctx, &toot)
	if err != nil {
		return err
	}
	logDebug(fmt.Sprintf("Toot: %s", toot.Status))
	return nil

}


// HomeTimeline HTLの取得
func (t *Mastodon) HomeTimeline(page int) error {
	ctx := context.Background()
	err := t.tlPages(page, func(pg *mastodon.Pagination) ([]*mastodon.Status, error) {
		return t.client.GetTimelineHome(ctx, pg)
	})
	if err != nil {
		return err
	}
	return nil
}

// TailHomeTimeline HTLのストリーミング取得
func (t *Mastodon) TailHomeTimeline() error {
	wsc := t.client.NewWSClient()
	ctx := context.Background()
	q, err := wsc.StreamingWSUser(ctx)
	if err != nil {
		return err
	}
	for e := range q {
		if u, ok := e.(*mastodon.UpdateEvent); ok {
			t.displayTimeline(u.Status)
		}
	}
	return nil
}

// ページングしながら取得する
func (t *Mastodon) tlPages(max int, f func(*mastodon.Pagination) ([]*mastodon.Status, error)) error {
	var statuses []*mastodon.Status
	var maxid mastodon.ID
	for i := 0; i < max; i++ {
		pg := mastodon.Pagination{
			MaxID: maxid,
			Limit: 40,
		}
		s, err := f(&pg)
		if err != nil {
			return err
		}
		// spewDump(pg)
		statuses = append(statuses, s...)
		if pg.MaxID == "" {
			break
		}
		maxid = pg.MaxID
		time.Sleep(time.Second)
	}
	statuses = t.reverseStatuses(statuses)
	for _, v := range statuses {
		t.displayTimeline(v)
	}
	return nil
}

// タイムラインの表示
func (t *Mastodon) displayTimeline(s *mastodon.Status) {
	createdAt := s.CreatedAt.In(time.Local).Format(time.RFC3339)
	content := bluemonday.StrictPolicy().Sanitize(s.Content)
	fmt.Printf("[%s] (%s) %s\n", createdAt, s.Account.Username, s.Account.DisplayName)
	fmt.Printf(" %s\n", content)
	fmt.Println("")
}

// ステータスの順番を逆転させる
func (t *Mastodon) reverseStatuses(s []*mastodon.Status) []*mastodon.Status {
	r := []*mastodon.Status{}
	e := len(s) - 1
	for i := 0; i <= e; i++ {
		r = append(r, s[e-i])
	}
	return r
}
