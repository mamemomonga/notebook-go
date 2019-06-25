package main

import (
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
    "log"
)

// GetCurrentAccountID アカウントIDを得る
func GetCurrentAccountID() string {
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
	return *ret.Account
}

