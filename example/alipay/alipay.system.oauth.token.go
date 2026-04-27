package main

import (
	"context"
	"log"
	"os"

	"github.com/godrealms/go-aliyun-sdk/alipay"
	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func main() {
	client := alipay.NewClient()
	client.AppId = os.Getenv("ALIPAY_APP_ID")
	client.PrivateKey = os.Getenv("ALIPAY_PRIVATE_KEY")
	client.AlipayPublicKey = os.Getenv("ALIPAY_PUBLIC_KEY_FROM_ALIPAY")

	resp, err := client.AlipaySystemOauthToken(context.Background(), &types.SystemOauthToken{
		GrantType: "authorization_code",
		Code:      os.Getenv("ALIPAY_AUTH_CODE"),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("AccessToken: %s", resp.AlipaySystemOauthTokenResponse.AccessToken)
}
