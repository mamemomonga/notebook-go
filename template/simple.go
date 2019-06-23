package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

// MyTMPL is struct of template
type MyTMPL struct {
	Name      string
	HomePage  string
	Favorites []string
}

func runSimple() {
	tpl := template.Must(template.ParseFiles("template.txt"))
	m := MyTMPL{
		Name:      "まめも",
		HomePage:  "https://github.com/mamemomonga",
		Favorites: []string{"なし", "りんご", "バナナ"},
	}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, m); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}
