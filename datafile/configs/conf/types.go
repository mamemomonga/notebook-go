package conf

// Configs 設定データの定義
type Configs struct {
	StatesFile string   `yaml:"states_file"`
	Key1       string   `yaml:"key_1"`
	Key2       string   `yaml:"key_2"`
	Key3       string   `yaml:"key_3"`
	Users      []string `yaml:"users"`
}

// States 状態データの出意義
type States struct {
	Passwords map[string]string `json:"passwords"`
}

// ConfigsInit 設定データの初期化
func (t *Conf) ConfigsInit() {
	t.Configs = &Configs{
		Key1:  "Value1",
		Key2:  "Value2",
		Key3:  "Value3",
		Users: []string{"user1", "user2", "user3"},
	}
}

// StatesInit 状態データの初期化
func (t *Conf) StatesInit() {
	t.States = &States{
		Passwords: make(map[string]string),
	}
}
