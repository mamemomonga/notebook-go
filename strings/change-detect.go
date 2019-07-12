package main

import (
	"github.com/thoas/go-funk"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	p := NewPlayers(&PlayersConfig{
		CBJoin: func(users []string) {
			log.Printf("  ユーザが加わりました: %s", strings.Join(users, " と "))
		},
		CBLeave: func(users []string) {
			log.Printf("  ユーザが去りました: %s", strings.Join(users, " と "))
		},
		CBCurrent: func(num int) {
			log.Printf("  現在のプレイヤー数: %d",num)
		},
	})
	p.Check("There are 2 of a max 10 players online: user1, user2")
	p.Check("There are 2 of a max 10 players online: user1, user2")
	p.Check("There are 2 of a max 10 players online: user1, user3")
	p.Check("There are 1 of a max 10 players online: user3")
	p.Check("There are 0 of a max 10 players online: ")
}

type Players struct {
	c           *PlayersConfig
	prevPlayers []string       // 前回のプレイヤー
	rex1        *regexp.Regexp // 正規表現
}

type PlayersConfig struct {
	CBJoin  func([]string) // プレイヤーが加わった時のコールバック
	CBLeave func([]string) // プレイヤーが抜けたときのコールバック
	CBCurrent func(int)    // 現在のプレイヤー数のコールバック
}

func NewPlayers(c *PlayersConfig) (t *Players) {
	t = new(Players)
	t.c = c
	t.rex1 = regexp.MustCompile(`^There are (\d+) of a max (\d+) players online:(.+)`)
	return t
}

func (t *Players) Check(buf string) {

	match := t.rex1.FindStringSubmatch(buf) // マッチ実行
	current, _ := strconv.Atoi(match[1])  // 整数に変換

	// 取得したプレイヤーをカンマで分割し、文字列のスペースを除去
	names := funk.Map(strings.Split(match[3], ","), func(x string) string {
		return strings.TrimSpace(x)
	}).([]string)

	// 空欄は除外する
	names = funk.FilterString(names, func(x string) bool {
		if x == "" {
			return false
		}
		return true
	})

	// 参加プレイヤー: 現在のプレイヤーに前回のプレイヤーが存在している
	pjoin := funk.FilterString(names, func(x string) bool { return !funk.ContainsString(t.prevPlayers, x) })
	// 退出プレイヤー: 前回のプレイヤーに現在のプレイヤーが存在していない
	pleave := funk.FilterString(t.prevPlayers, func(x string) bool { return !funk.ContainsString(names, x) })

	// 前回のプレイヤーの状態を保存する
	t.prevPlayers = names

	// コールバックの実行
	t.c.CBCurrent(current)
	if len(pjoin) > 0 {
		t.c.CBJoin(pjoin)
	}
	if len(pleave) > 0 {
		t.c.CBLeave(pleave)
	}
}

