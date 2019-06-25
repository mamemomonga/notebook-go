package main

import (
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
    "log"
)

// GetAccountID アカウントIDを返す
func GetAccountID() {
    sess, err := session.NewSession()
	if err != nil {
		log.Println(err)
	}
	svc := sts.New(sess)

	input := sts.GetCallerIdentityInput{}
	ret,err := svc.GetCallerIdentity(&input)
	if err != nil {
		log.Println(err)
	}

	spewDump(ret)
}


