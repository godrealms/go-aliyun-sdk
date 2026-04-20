package main

import (
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

	u, err := client.GetUserAuthPageURL(&types.UserInfoAuthPage{
		Scope:       "auth_user",
		RedirectUri: "https://example.com/alipay/callback",
		State:       "order_123",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("用户授权页 URL: %s", u)
}
