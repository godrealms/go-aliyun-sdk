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

	resp, err := client.AlipayMarketingCampaignCashList(&types.CampaignCashList{
		PageIndex: "1",
		PageSize:  "10",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("TotalCount: %d, Items: %d", resp.AlipayMarketingCampaignCashListResponse.TotalCount, len(resp.AlipayMarketingCampaignCashListResponse.ResultList))
}
