package main

import (
	"log"
)

func main() {

	client, err := GetGCEClient("service-account-key.json")
	if err != nil {
		log.Fatal(err)
	}

	data, err := GetGCEInstanceList(client, "PROJECT_ID", "REGION")
	// data,err := GetGCEInstanceList(client,"mamemo-190623","asia-northeast1-b")

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range data.Items {
		log.Printf("Name: %20s State: %10s\n", v.Name, v.Status)
	}
}
