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

	resp, err := client.AlipayTradeOrderinfoSync(context.Background(), &types.TradeOrderinfoSync{
		TradeNo:      os.Getenv("ALIPAY_TRADE_NO"),
		OutRequestNo: os.Getenv("ALIPAY_OUT_REQUEST_NO"),
		OrderType:    os.Getenv("ALIPAY_ORDER_TYPE"),
		OrderScene:   os.Getenv("ALIPAY_ORDER_SCENE"),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Code: %s", resp.Response.Code)
}
