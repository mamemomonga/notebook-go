package simple

import (
	"fmt"
	"log"
	"os"
	"time"
	"context"
	"io/ioutil"
	"encoding/json"
	"github.com/mattn/go-mastodon"
	"github.com/microcosm-cc/bluemonday"
	"github.com/davecgh/go-spew/spew"
)

const Debug=true

func logDebug(s string) {
	if ! Debug {
		return
	}
	log.Printf("debug: %s",s)
}
func spewDump(s interface{}) {
	if ! Debug {
		return
	}
	spew.Dump(s)
}

type Mastodon struct {
	c        *MastodonConfig
	client   *mastodon.Client
	Ready    bool
	lastToot string
	AccountCurrentUser *mastodon.Account
}

type MastodonConfig struct {
	Server     string
	Email      string
	Password   string
	ClientName string
	ClientFile string
}

type ClientConfigs struct {
	Tokens map[string]ClientTokens `json:"tokens"`
}

type ClientTokens struct {
	ClientID     string `json:"id"`
	ClientSecret string `json:"secret"`
}

func NewMastodon(c *MastodonConfig) *Mastodon {
	t := new(Mastodon)
	t.c = c
	t.Ready = false
	return t
}

func (t *Mastodon) Connect() (err error) {
	ctx := context.Background()

	ccs := &ClientConfigs{
		Tokens: make(map[string]ClientTokens),
	}

	if _, err := os.Stat(t.c.ClientFile); !os.IsNotExist(err) {
		if err := t.loadClientFile(ccs); err != nil {
			return err
		}
	}
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
		if err := t.saveClientFile(ccs); err != nil {
			return err
		}
	}
	t.client = mastodon.NewClient(&mastodon.Config{
		Server:       fmt.Sprintf("https://%s/", t.c.Server),
		ClientID:     ccs.Tokens[t.c.Server].ClientID,
		ClientSecret: ccs.Tokens[t.c.Server].ClientSecret,
	})
	if err := t.client.Authenticate(ctx, t.c.Email, t.c.Password); err != nil {
		return err
	}
	account, err := t.client.GetAccountCurrentUser(ctx)
	if err != nil {
		return err
	}
	t.AccountCurrentUser=account
	t.Ready = true
	return nil
}

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
	logDebug(fmt.Sprintf("Toot: %s",toot.Status))
	return nil
}

func (t *Mastodon) HomeTimeline(page int) error {
	ctx := context.Background()
	err := t.tlPages( page, func(pg *mastodon.Pagination) ([]*mastodon.Status, error){
		return t.client.GetTimelineHome(ctx,pg)
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *Mastodon) TailHomeTimeline() error {
	wsc := t.client.NewWSClient()
	ctx := context.Background()
	q,err := wsc.StreamingWSUser(ctx)
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


func (t *Mastodon) tlPages(max int, f func(*mastodon.Pagination) ([]*mastodon.Status, error) ) error {
	var statuses []*mastodon.Status
	var maxid mastodon.ID
	for i:=0; i < max; i++ {
		pg := mastodon.Pagination{
			MaxID: maxid,
			Limit: 40,
		}
		s, err := f(&pg)
		if err != nil {
			return err
		}
		// spewDump(pg)
		statuses=append(statuses,s...)
		if pg.MaxID == "" {
			break
		}
		maxid = pg.MaxID
		time.Sleep(time.Second)
	}
	statuses = t.reverseStatuses(statuses)
	for _,v := range statuses {
		t.displayTimeline(v)
	}
	return nil
}

func (t *Mastodon) displayTimeline(s *mastodon.Status) {
	createdAt := s.CreatedAt.In(time.Local).Format(time.RFC3339)
	content := bluemonday.StrictPolicy().Sanitize(s.Content)
	fmt.Printf("[%s] (%s) %s\n",createdAt, s.Account.Username, s.Account.DisplayName)
	fmt.Printf(" %s\n",content)
	fmt.Println("")
}

func (t *Mastodon) reverseStatuses(s []*mastodon.Status) []*mastodon.Status {
	r := []*mastodon.Status{}
	e := len(s)-1
	for i:=0; i<=e; i++ {
		r = append(r,s[e-i])
	}
	return r
}
