package configs

import (
	"errors"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type Configs struct {
	Configs C
}

func New() (t *Configs) {
	t = new(Configs)
	t.Configs = C{}
	return t
}

func (t *Configs) Load() error {

	// 設定ファイルの検索順
	sl := []string{}

	// バイナリと同じ場所か一階層上
	{
		exe, err := os.Executable()
		if err != nil {
			return err
		}
		b, err := filepath.Abs(filepath.Join(filepath.Dir(exe), "."))
		if err != nil {
			return err
		}

		suffix := []string{"yaml", "yml"}
		prefix := []string{"./", "../"}
		names := []string{defaultFilename, "." + defaultFilename}

		for _, s := range suffix {
			for _, p := range prefix {
				for _, n := range names {
					sl = append(sl, filepath.Join(b, p, n+"."+s))
				}
			}
		}
	}

	// ホームディレクトリ
	{
		h, err := homedir.Dir()
		if err != nil {
			return err
		}

		suffix := []string{"yaml", "yml"}
		names := []string{defaultFilename, "." + defaultFilename}

		for _, s := range suffix {
			for _, n := range names {
				sl = append(sl, filepath.Join(h, n+"."+s))
			}
		}

	}

	// カレントディレクトリ
	{
		p, err := os.Getwd()
		if err != nil {
			return err
		}
		suffix := []string{"yaml", "yml"}
		names := []string{defaultFilename, "." + defaultFilename}

		for _, s := range suffix {
			for _, n := range names {
				sl = append(sl, filepath.Join(p, n+"."+s))
			}
		}
	}

	// 検索
	configFile := ""
	for _, i := range sl {
		if _, err := os.Stat(i); !os.IsNotExist(err) {
			log.Printf("debug: FOUND %s", i)
			configFile = i
			break
		} else {
			log.Printf("debug: NOT FOUND %s", i)
		}
	}
	if configFile == "" {
		log.Println()
		return errors.New("alert: 設定ファイルがありません")
	}

	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	s := regexp.MustCompile(`\r\n|\r|\n`).ReplaceAllString(string(buf), "\n")
	err = yaml.Unmarshal([]byte(s), &t.Configs)

	if err != nil {
		return err
	}

	return nil
}
