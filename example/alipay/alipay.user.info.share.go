package main

import (
	"log"
	"os"

	"github.com/godrealms/go-aliyun-sdk/alipay"
)

func main() {
	client := alipay.NewClient()
	client.AppId = os.Getenv("ALIPAY_APP_ID")
	client.PrivateKey = os.Getenv("ALIPAY_PRIVATE_KEY")
	client.AlipayPublicKey = os.Getenv("ALIPAY_PUBLIC_KEY_FROM_ALIPAY")

	resp, err := client.AlipayUserInfoShare(os.Getenv("ALIPAY_USER_ACCESS_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("UserId: %s, NickName: %s", resp.AlipayUserInfoShareResponse.UserId, resp.AlipayUserInfoShareResponse.NickName)
}
