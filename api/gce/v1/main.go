package main

import (
	"flag"
	"log"
)

func main() {

	var (
		projectID string
		region    string
	)
	{
		p := flag.String("i", "mamemo-190623", "プロジェクトID")
		r := flag.String("r", "asia-northeast1-b", "リージョン")
		flag.Parse()

		projectID = *p
		region = *r
	}

	client, err := GetGCEClient("service-account-key.json")
	if err != nil {
		log.Fatal(err)
	}

	data, err := GetGCEInstanceList(client, projectID, region)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range data.Items {
		log.Printf("Name: %20s State: %10s\n", v.Name, v.Status)
	}
}
