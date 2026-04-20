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

	resp, err := client.AlipayTradeRoyaltyRateQuery(&types.TradeRoyaltyRateQuery{
		OutRequestNo: os.Getenv("ALIPAY_OUT_REQUEST_NO"),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range resp.Response.RoyaltyInfos {
		log.Printf("RoyaltyType: %s, Rate: %s", info.RoyaltyType, info.Rate)
	}
}
