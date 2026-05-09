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

	resp, err := client.AlipayFundJointThaw(context.Background(), &types.FundJointThaw{
		OutRequestNo: "THAW_REQUEST_001",
		FreezeId:     os.Getenv("ALIPAY_FREEZE_ID"),
		Amount:       "0.10",
		Remark:       "联合解冻测试",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("OutRequestNo: %s", resp.AlipayFundJointThawResponse.OutRequestNo)
}
