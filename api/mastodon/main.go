package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"flag"
	"time"
	"github.com/comail/colog"
	don "github.com/mamemomonga/notebook-go/api/mastodon/mastodon/simple"
)


// Locationの設定
const Location = "Asia/Tokyo"

func init() {

	// cologの設定
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
	        Colors: true,
	        Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()

	// Localtimeの設定
	loc, err := time.LoadLocation(Location)

	// Locale情報がなければ、JSTにする
	if err != nil {
		loc = time.FixedZone(Location, 9*60*60)
	}
	time.Local = loc
}

// 設定ファイルの構造
type Config struct {
	Server     string `yaml:"server"`
	Email      string `yaml:"email"`
	Password   string `yaml:"password"`
	ClientName string `yaml:"client_name"`
	ClientFile string `yaml:"client_file"`
}

// メイン
func main() {

	// フラグの読込
	fgc := flag.String("c", "./etc/config.yaml", "設定ファイル")
	fgt := flag.String("t", "","トゥート")
	fgh := flag.Bool("h", false,"ホームタイムライン")
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
		if err := d.Toot(*fgt); err != nil {
			log.Fatal(err)
		}
	}

	// ホームタイムライン
	if *fgh {
		if err := d.HomeTimeline(2); err != nil {
			log.Fatal(err)
		}
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

