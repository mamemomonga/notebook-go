package main

import (
	"fmt"
	"log"

	"example.local/errors/src/myapp"
)

var (
	version  string
	revision string
)

func main() {
	verString := fmt.Sprintf("%s-%s", version, revision)
	log.Printf("Hello World Version: %s\n", verString)

	myap := myapp.New()
	log.Printf("MYAP: %s", myap.Hello())

	err := myap.Error0()
	if err != nil {
		log.Println(err)
	}

	err = myap.Error1()
	if err != nil {
		ew, _ := err.(*myapp.AppError)
		log.Printf("Code: %d", ew.Code)
	}

}
