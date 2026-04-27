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

	resp, err := client.AlipayTradeAdvanceConsult(context.Background(), &types.TradeAdvanceConsult{
		OutTradeNo:  os.Getenv("ALIPAY_OUT_TRADE_NO"),
		TotalAmount: os.Getenv("ALIPAY_TOTAL_AMOUNT"),
		Subject:     os.Getenv("ALIPAY_SUBJECT"),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ConsultId: %s, NextAction: %s", resp.Response.ConsultId, resp.Response.NextAction)
}
