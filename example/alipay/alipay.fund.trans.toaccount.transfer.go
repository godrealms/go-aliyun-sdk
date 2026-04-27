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

	resp, err := client.AlipayFundTransToaccountTransfer(context.Background(), &types.FundTransToaccountTransfer{
		OutBizNo:      "OUT_BIZ_NO_001",
		PayeeType:     "ALIPAY_LOGONID",
		PayeeAccount:  os.Getenv("ALIPAY_PAYEE_ACCOUNT"),
		Amount:        "0.10",
		PayeeRealName: os.Getenv("ALIPAY_PAYEE_NAME"),
		Remark:        "转账到账户",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("OrderId: %s", resp.AlipayFundTransToaccountTransferResponse.OrderId)
}
