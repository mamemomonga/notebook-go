package main

import (
	"log"
)

func main() {
	g := NewGCE()
	err := g.LoadCredentialsFile("service-account-key.json")
	if err != nil {
		log.Fatal(err)
	}
	g.Project = "the-project"
	g.Zone = "asia-northeast1-b"
	g.Instance = "the-instance-name"

	{
		r, err := g.Get()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(r.Status)
	}
	{
		r, err := g.Stop()
		if err != nil {
			log.Fatal(err)
		}
		spewDump(r)
	}

}
