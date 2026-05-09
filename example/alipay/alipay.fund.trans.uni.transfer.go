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

	resp, err := client.AlipayFundTransUniTransfer(context.Background(), &types.FundTransUniTransfer{
		OutBizNo:    "OUT_BIZ_NO_001",
		TransAmount: "0.10",
		ProductCode: "TRANS_ACCOUNT_NO_PWD",
		OrderTitle:  "转账测试",
		PayeeInfo: &types.FundTransPayee{
			Identity:     os.Getenv("ALIPAY_PAYEE_USER_ID"),
			IdentityType: "ALIPAY_USER_ID",
			Name:         os.Getenv("ALIPAY_PAYEE_NAME"),
		},
		Remark: "单笔转账",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("OrderId: %s", resp.AlipayFundTransUniTransferResponse.OrderId)
}
