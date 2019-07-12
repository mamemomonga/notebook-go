package main

/*
goでJSON/YAML読み書き

*/

import (
	"log"
	"os"

	"github.com/mamemomonga/notebook-go/pokemondata"
)

// バージョンとリビジョン
var (
	Version  string
	Revision string
)

// MyData データの定義
type MyData struct {
	Name     string `yaml:"my_name" json:"my_name"`
	Pokemons []pokemondata.Pokemon
}

// メイン
func main() {

	// ログの出力先をstdoutに変更
	log.SetOutput(os.Stdout)

	// バージョンとリビジョンの表示
	log.Printf("yamljson version %s revision %s\n", Version, Revision)

	{
		msgYellow("YAMLとJSONの書込")
		d := MyData{
			Name:     "まめも",
			Pokemons: pokemondata.InitData(),
		}
		err := writeYAML("pokemon.yaml", &d)
		if err != nil {
			log.Fatal(err)
		}
		err = writeJSON("pokemon.json", &d)
		if err != nil {
			log.Fatal(err)
		}
	}
	{
		msgYellow("YAMLの読込")
		d, err := readYAML("pokemon.yaml")
		if err != nil {
			log.Fatal(err)
		}
		spewDump(d)
	}
	{
		msgYellow("JSONの読込")
		d, err := readJSON("pokemon.json")
		if err != nil {
			log.Fatal(err)
		}
		spewDump(d)
	}
}
