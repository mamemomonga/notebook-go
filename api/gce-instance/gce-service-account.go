package main

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
)

// GetGCEClient API接続用のhttp.Clientを得る
func GetGCEClient(filename string) (client *http.Client, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return client, err
	}
	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/compute")
	if err != nil {
		return client, err
	}
	return conf.Client(oauth2.NoContext), nil
}
