package awsutils

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "log"
)

type AWSUtils struct {
	session *session.Session
}

func NewAWSUtils(region string) *AWSUtils {
	t := new(AWSUtils)
    sess, err := session.NewSession(&aws.Config{ Region: aws.String(region)})
	if err != nil {
		log.Println(err)
		log.Fatal("alert: セッション開始に失敗しました。")
	}
	t.session = sess

	return t
}

