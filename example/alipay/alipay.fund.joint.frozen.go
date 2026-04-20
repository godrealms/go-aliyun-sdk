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

	resp, err := client.AlipayFundJointFrozen(&types.FundJointFrozen{
		OutRequestNo: "FREEZE_REQUEST_001",
		PayerUserId:  os.Getenv("ALIPAY_PAYER_USER_ID"),
		Amount:       "0.10",
		Remark:       "联合冻结测试",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("FreezeId: %s", resp.AlipayFundJointFrozenResponse.FreezeId)
}
