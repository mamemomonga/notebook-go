package main

import (
	"log"
	"errors"
	"fmt"
	"net/http"
	"io/ioutil"
)

func fetch(url string) (string,error) {
	log.Printf("Fetch: %s", url)
	res, err := http.Get(url)
	if err != nil {
		return "",err
	}
	if res.StatusCode != 200 {
		return "",errors.New(fmt.Sprintf("Status Code: %d", res.StatusCode))
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "",err
	}
	return string(b),nil
}
