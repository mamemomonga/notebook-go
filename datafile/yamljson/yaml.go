package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// YAML書出
func writeYAML(filename string, data *MyData) error {
	buf, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, buf, 0644)
	if err != nil {
		return err
	}
	log.Printf("Write: %s", filename)
	return nil
}

// YAML読込
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
