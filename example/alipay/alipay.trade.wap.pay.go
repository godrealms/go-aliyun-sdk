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
	client.ReturnUrl = os.Getenv("ALIPAY_RETURN_URL")
	client.NotifyUrl = os.Getenv("ALIPAY_NOTIFY_URL")

	resp, err := client.AlipayTradeWapPay(context.Background(), &types.TradeWapPay{
		OutTradeNo:  os.Getenv("ALIPAY_OUT_TRADE_NO"),
		Subject:     os.Getenv("ALIPAY_SUBJECT"),
		TotalAmount: os.Getenv("ALIPAY_TOTAL_AMOUNT"),
		ProductCode: "QUICK_WAP_WAY",
		QuitUrl:     os.Getenv("ALIPAY_QUIT_URL"),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
