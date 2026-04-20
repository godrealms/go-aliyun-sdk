package main

import (
	"github.com/godrealms/go-aliyun-sdk/alipay"
	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"log"
	"os"
)

func main() {
	client := alipay.NewClient()
	client.Sandbox = true
	client.AppId = "2021000147675551"
	// Required env vars: ALIPAY_PRIVATE_KEY, ALIPAY_PUBLIC_KEY, ALIPAY_PUBLIC_KEY_FROM_ALIPAY
	client.PrivateKey = os.Getenv("ALIPAY_PRIVATE_KEY")
	client.PublicKey = os.Getenv("ALIPAY_PUBLIC_KEY")
	client.AlipayPublicKey = os.Getenv("ALIPAY_PUBLIC_KEY_FROM_ALIPAY")
	client.Http.SetBaseURL("https://openapi-sandbox.dl.alipaydev.com/gateway.do")

	query := &types.AgreementTransfer{
		AgreementNo: "2021000147675551",
		PeriodRuleParams: &types.PeriodRuleParams{
			ExecuteTime:   "2021-08-01 00:00:00",
			Period:        1,
			PeriodType:    types.PeriodTypeDAY,
			SingleAmount:  100,
			TotalAmount:   100,
			TotalPayments: 1,
		},
		TargetProductCode: "CYCLE_PAY_AUTH_P",
	}

	transfer, err := client.AgreementTransfer(query)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(transfer)
}
