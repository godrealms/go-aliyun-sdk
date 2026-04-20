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
	resp, err := client.AlipayMarketingCardOpen(&types.CardOpen{
		RequestId: os.Getenv("ALIPAY_REQUEST_ID"), TemplateId: os.Getenv("ALIPAY_TEMPLATE_ID"),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("OpenUrl: %s", resp.AlipayMarketingCardOpenResponse.OpenUrl)
}
