package awsutils

import (
	"github.com/aws/aws-sdk-go/service/sts"
	"log"
)

// GetAccountID アカウントIDを返す
func (t *AWSUtils) GetAccountID() string {
	svc := sts.New(t.session)
	input := sts.GetCallerIdentityInput{}
	ret,err := svc.GetCallerIdentity(&input)
	if err != nil {
		log.Println(err)
		log.Fatal("alert: アカウントID取得に失敗しました。認証設定を確認してください")
	}
	return *ret.Account
}

