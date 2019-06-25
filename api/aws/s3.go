package main

import (
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
    "log"
)

func S3Download() {
    sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	file := ""
	bucket := ""
	item := ""

	downloader := s3manager.NewDownloader(sess)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(item),
	})
	if err != nil {
		log.Fatal("Unable to download item %q, %v", item, err)
	}
	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

