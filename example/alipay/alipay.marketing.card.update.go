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
	resp, err := client.AlipayMarketingCardUpdate(context.Background(), &types.CardUpdate{
		CardId: os.Getenv("ALIPAY_CARD_ID"), TemplateId: os.Getenv("ALIPAY_TEMPLATE_ID"), BalanceInfo: `{"balance":"200.00"}`,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Code: %s, Msg: %s", resp.AlipayMarketingCardUpdateResponse.Code, resp.AlipayMarketingCardUpdateResponse.Msg)
}
