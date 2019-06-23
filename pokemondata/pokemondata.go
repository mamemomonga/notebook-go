// Package pokemondata データ構造のサンプル
// ポケモン図鑑を参考にしています
package pokemondata

// 性別を定義します 不明, オス, メス
const (
	SexUnknown = 0               // 0 性別不明
	SexMale    = int8(1 << iota) // 2 オス
	SexFemale                    // 4 メス
)

// Pokemon ポケモン構造体
type Pokemon struct {
	Name      string   `yaml:"name"      json:"name"`
	Number    int      `yaml:"number"    json:"number"`
	Type      string   `yaml:"type"      json:"type"`
	Category  string   `yaml:"category"  json:"category"`
	Weak      []string `yaml:"weak"      json:"waak"`
	Height    float32  `yaml:"height"    json:"height"`
	Weight    float32  `yaml:"weight"    json:"weight"`
	Character []string `yaml:"character" json:"character"`
	Sex       int8     `yaml:"sex"       json:"sex"`
}

// InitData 初期データ
func InitData() []Pokemon {
	return []Pokemon{
		Pokemon{
			Name:      "イーブイ",
			Number:    133,
			Type:      "ノーマル",
			Category:  "しんかポケモン",
			Weak:      []string{"かくとう"},
			Height:    0.3,
			Weight:    6.5,
			Sex:       SexMale | SexFemale,
			Character: []string{"にげあし", "てきおうりょく"},
		},
		Pokemon{
			Name:      "シャワーズ",
			Number:    134,
			Type:      "みず",
			Category:  "あわはきポケモン",
			Weak:      []string{"くさ", "でんき"},
			Height:    1.0,
			Weight:    29.0,
			Sex:       SexMale | SexFemale,
			Character: []string{"ちょすい"},
		},
		Pokemon{
			Name:      "ブースター",
			Number:    136,
			Type:      "ほのお",
			Category:  "ほのおポケモン",
			Weak:      []string{"みず", "じめん", "いわ"},
			Height:    0.9,
			Weight:    25.0,
			Sex:       SexMale | SexFemale,
			Character: []string{"もらいび"},
		},
	}
}

