package main

import (
	"log"
	"fmt"
)

func main() {
	// 京都の天気
	weather,err := fetch("http://weather.livedoor.com/forecast/webservice/json/v1?city=260010")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println( weather )
}

