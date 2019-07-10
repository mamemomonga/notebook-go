package main

import (
	"flag"
	"path/filepath"
	don "github.com/mamemomonga/notebook-go/api/mastodon/mastodon/simple"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// Config 設定ファイルの構造
type Config struct {
	Server     string `yaml:"server"`
	Email      string `yaml:"email"`
	Password   string `yaml:"password"`
	ClientName string `yaml:"client_name"`
	ClientFile string `yaml:"client_file"`
}

// メイン
func main() {

	home, _:= homedir.Dir()

	// フラグの読込
	fgc  := flag.String("c", filepath.Join(home,".easydon/config.yaml"), "設定ファイル")
	fgt  := flag.String("t", "", "トゥート内容")
	fga  := flag.String("a", "", "指定したファイルを添付してトゥート")
	fgh  := flag.Int("h", 0,  "ホームタイムラインを指定ページ数表示する")
	fghf := flag.Bool("f", false, "ストリーミング表示する")

	flag.Parse()

	// 設定ファイルの読み込み
	config, err := readConfigYAML(*fgc)
	if err != nil {
		log.Fatal(err)
	}

	// マストドンAPIの初期化
	d := don.NewMastodon(&don.MastodonConfig{
		Server:     config.Server,
		Email:      config.Email,
		Password:   config.Password,
		ClientName: config.ClientName,
		ClientFile: config.ClientFile,
	})
	// マストドンに接続
	if err := d.Connect(); err != nil {
		log.Fatal(err)
	}

	// トゥートする
	if *fgt != "" {
		if *fga == "" {
			// 添付なしでトゥートする
			if err := d.Toot(*fgt); err != nil {
				log.Fatal(err)
			}
		} else {
			// ファイルを添付してトゥートする
			if err := d.TootWithAttachment(*fgt,*fga); err != nil {
			log.Fatal(err)
			}
		}
	}

	// ホームタイムライン
	if *fgh > 0 {
		if err := d.HomeTimeline(*fgh); err != nil {
			log.Fatal(err)
		}
	}
	// ホームタイムラインストリーミング
	if *fghf {
		if err := d.TailHomeTimeline(); err != nil {
			log.Fatal(err)
		}
	}
}

// 設定ファイル読み込み
func readConfigYAML(filename string) (data *Config, err error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		return
	}
	log.Printf("Read: %s", filename)
	return data, nil
}
