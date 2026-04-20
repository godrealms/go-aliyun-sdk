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

	resp, err := client.AlipayTradeRoyaltyRelationBatchquery(&types.TradeRoyaltyRelationBatchquery{
		OutRequestNo: os.Getenv("ALIPAY_OUT_REQUEST_NO"),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Count: %d", resp.Response.Count)
	for _, r := range resp.Response.ReceiverInfos {
		log.Printf("  Receiver: %s (%s)", r.ReceiverAccount, r.ReceiverName)
	}
}
