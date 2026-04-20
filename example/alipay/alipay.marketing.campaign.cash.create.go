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

	resp, err := client.AlipayMarketingCampaignCashCreate(&types.CampaignCashCreate{
		CrowdName:   "春季促销活动",
		PrizeType:   "CASH",
		BudgetInfo:  `{"total_budget":"1000.00"}`,
		PrizeInfo:   `{"prize_amount":"10.00"}`,
		SendChannel: "ALISEND",
		StartTime:   "2026-04-20 00:00:00",
		EndTime:     "2026-05-20 23:59:59",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CrowdNo: %s", resp.AlipayMarketingCampaignCashCreateResponse.CrowdNo)
}
