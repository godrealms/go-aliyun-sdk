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
	resp, err := client.AlipayMarketingVoucherQuery(&types.VoucherQuery{VoucherId: os.Getenv("ALIPAY_VOUCHER_ID")})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("VoucherName: %s, Status: %s", resp.AlipayMarketingVoucherQueryResponse.VoucherName, resp.AlipayMarketingVoucherQueryResponse.Status)
}
