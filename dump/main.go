package main

import (
	"fmt"
	"log"

	"github.com/mamemomonga/notebook-go/pokemondata"
)

func main() {
	pokemons := pokemondata.InitData()

	title("Printf %v デフォルトフォーマット")
	fmt.Printf("%v\n", pokemons)

	title("Printf %+v デフォルトフォーマット+構造体名")
	fmt.Printf("%+v\n", pokemons)

	title("Printf %#v Go構文で値を表現")
	fmt.Printf("%#v\n", pokemons)

	title("Printf %T Go構文で型を表現")
	fmt.Printf("%T\n", pokemons)

	title("spewを使う")
	spewDump(pokemons)
}

func title(s string) {
	log.Printf("\033[44;1m   %s   \033[0m", s)
}
