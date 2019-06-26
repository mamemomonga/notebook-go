package main

import (
    "log"
	"github.com/mamemomonga/notebook-go/api/aws/awsutils"
)

func main() {

	// 開始
	au := awsutils.NewAWSUtils("ap-northeast-1")

	// 現在のアカウントIDを得る
	log.Printf("AccountID: %s",au.GetAccountID())

//	au.S3Download(
//		"bucketname",
//		"/s3file",
//		"localfile",
//	)

}

