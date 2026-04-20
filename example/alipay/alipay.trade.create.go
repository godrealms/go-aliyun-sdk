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
	client.NotifyUrl = os.Getenv("ALIPAY_NOTIFY_URL")

	resp, err := client.AlipayTradeCreate(&types.TradeCreate{
		OutTradeNo:  os.Getenv("ALIPAY_OUT_TRADE_NO"),
		Subject:     os.Getenv("ALIPAY_SUBJECT"),
		TotalAmount: os.Getenv("ALIPAY_TOTAL_AMOUNT"),
		BuyerId:     os.Getenv("ALIPAY_BUYER_ID"),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("TradeNo: %s", resp.Response.TradeNo)
}
