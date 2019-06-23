package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetGCEInstanceList インスタンスリストを得る
func GetGCEInstanceList(client *http.Client, project string, region string) (data GCEInstanceList, err error) {
	resp, err := client.Get(fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/zones/%s/instances", project, region))
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Status: %d", resp.StatusCode)
		return data, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return data, err
	}

	//	log.Println(string(bodyBytes))

	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
