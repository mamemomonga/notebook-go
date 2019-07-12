package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// JSON書出
func writeJSON(filename string, data *MyData) (err error) {
	// インデントしない
	// buf, err := json.Marshal(data)

	// インデントする
	buf, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		return
	}
	err = ioutil.WriteFile(filename, buf, 0644)
	if err != nil {
		return
	}

	log.Printf("Write: %s", filename)
	return nil
}

// JSON読込
func readJSON(filename string) (data *MyData, err error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(buf, &data)
	if err != nil {
		return
	}
	log.Printf("Read: %s", filename)
	return data, nil
}
