// テストのときのみ同一ディレクトリで違うパッケージ名が許容される
package pokemondata_test

import (
	"github.com/mamemomonga/notebook-go/pokemondata"
	// テストに使用するパッケージ
	"testing" // https://golang.org/pkg/testing/
)

// ポケモンデータ読込先
var pokemons []pokemondata.Pokemon

// TestXXXXX という形でfunctionを定義が上から順に実行される。

// TestPokemondata01 ポケモンデータ読込
func TestPokemondata01(t *testing.T) {
	t.Log("InitData")
	pokemons = pokemondata.InitData()
}

// TestPokemodata02 最初のポケモンはイーブイである必要がある
func TestPokemodata02(t *testing.T) {
	t.Log("Check Eievui")
	t.Logf("%#v", pokemons[0])
	if pokemons[0].Name != "イーブイ" {
		t.Error("The first Pokemon needs Eievui.")
	}
}
