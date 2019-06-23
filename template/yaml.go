package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"text/template"

	"github.com/davecgh/go-spew/spew"
	yaml "gopkg.in/yaml.v2"
)

// MyData struct of template
type MyData struct {
	Template interface{} `yaml:"template"`
}

func runYaml() {
	tpl := template.Must(template.ParseFiles("template.txt"))
	data, err := readYAML("data.yaml")
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(data)

	var buf bytes.Buffer

	if err := tpl.Execute(&buf, data.Template); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())

}

func readYAML(filename string) (data *MyData, err error) {
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
