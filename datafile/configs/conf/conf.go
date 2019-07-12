package conf

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Conf 設定ストア
type Conf struct {
	Configs    *Configs
	States     *States
	c          *NewConfConfig
	BaseDir    string
	statesFile string
}

// NewConfConfig NewConf初期化設定
type NewConfConfig struct {
	ConfigsFile        string // 設定ファイル
	OffsetFromBin      string // 実行バイナリからの相対位置
	DefaultConfigsFile string // デフォルトの設定ファイル
	DefaultStatesFile  string // デフォルトの状態ファイル
}

// NewConf 設定ストア
func NewConf(c *NewConfConfig) *Conf {
	t := new(Conf)
	t.c = c
	return t
}

// GetDir ベースディレクトリからの絶対パス
func (t *Conf) GetDir(p string) (string, error) {
	return filepath.Abs(filepath.Join(t.BaseDir, p))
}

// Load 設定の読込
func (t *Conf) Load() error {
	// 実行バイナリからの相対位置を得る
	{
		exe, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		t.BaseDir, err = filepath.Abs(filepath.Join(filepath.Dir(exe), t.c.OffsetFromBin))
		if err != nil {
			log.Fatal(err)
		}
	}

	// 設定データの初期化
	t.ConfigsInit()
	// 状態データの初期化
	t.StatesInit()

	// ConfigsFileが空ならばデフォルト値を設定する
	if t.c.ConfigsFile == "" {
		if c, err := t.GetDir(t.c.DefaultConfigsFile); err != nil {
			log.Fatal(err)
		} else {
			t.c.ConfigsFile = c
		}
	}
	if _, err := os.Stat(t.c.ConfigsFile); !os.IsNotExist(err) {
		// 設定ファイルの読込
		buf, err := ioutil.ReadFile(t.c.ConfigsFile)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(buf, t.Configs)
		if err != nil {
			return err
		}
		log.Printf("Load: %s", t.c.ConfigsFile)
	} else {
		// デフォルト値でファイルを作成する
		buf, err := yaml.Marshal(t.Configs)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(t.c.ConfigsFile, buf, 0644)
		if err != nil {
			return err
		}
		log.Printf("Save: %s", t.c.ConfigsFile)
	}

	// StatesFileが空ならばデフォルト値を設定する
	if t.Configs.StatesFile == "" {
		if c, err := t.GetDir(t.c.DefaultStatesFile); err != nil {
			log.Fatal(err)
		} else {
			t.statesFile = c
		}
	} else {
		t.statesFile = t.Configs.StatesFile
	}

	if _, err := os.Stat(t.statesFile); !os.IsNotExist(err) {
		// 状態ファイルがあれば読み込む
		if err := t.LoadStates(); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

// SaveStates 状態の保存
func (t *Conf) SaveStates() error {
	buf, err := json.MarshalIndent(t.States, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(t.statesFile, buf, 0644)
	if err != nil {
		return err
	}
	log.Printf("Save: %s", t.statesFile)
	return nil
}

// LoadStates 状態の読込
func (t *Conf) LoadStates() error {
	buf, err := ioutil.ReadFile(t.statesFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, t.States)
	if err != nil {
		return err
	}
	log.Printf("Load: %s", t.statesFile)
	return nil
}
