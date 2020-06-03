package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func fetchFile(url string, filename string) error {
	log.Printf("Fetch: %s", url)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatal(fmt.Sprintf("Status Code: %d", res.StatusCode))
	}

	defer res.Body.Close()
	errChan := make(chan error)
	defer close(errChan)

	go func() {
		log.Printf("Write: %s\n", filename)
		outfile, err := os.Create(filename)
		if err != nil {
			errChan <- err
			return
		}
		defer outfile.Close()
		if _, err := io.Copy(outfile, res.Body); err != nil {
			errChan <- err
			return
		}
		errChan <- nil
		return
	}()

	err = <-errChan
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := fetchFile("https://code.jquery.com/jquery-3.4.1.min.js", "jquery.js")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Finish")
}
