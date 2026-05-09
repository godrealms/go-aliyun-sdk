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
	resp, err := client.AlipayMarketingVoucherCreate(context.Background(), &types.VoucherCreate{
		VoucherName: "满100减10", VoucherType: "MONEY_VOUCHER",
		DenominationMoney: "10.00", MerchantId: os.Getenv("ALIPAY_MERCHANT_ID"),
		Quantity: "100", ValidBeginTime: os.Getenv("ALIPAY_VOUCHER_BEGIN_TIME"), ValidEndTime: os.Getenv("ALIPAY_VOUCHER_END_TIME"),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("VoucherId: %s", resp.AlipayMarketingVoucherCreateResponse.VoucherId)
}
