package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay"
	"github.com/godrealms/go-aliyun-sdk/alipay/types"
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

	query := &types.TradePay{
		OutTradeNo:      fmt.Sprintf("%s%d", time.Now().Format("20060102150405"), time.Now().Unix()),
		TotalAmount:     "0.01",
		Subject:         "测试支付",
		ProductCode:     "QUICK_MSECURITY_PAY",
		GoodsDetail:     nil,
		TimeExpire:      "",
		ExtendParams:    nil,
		PassbackParams:  "",
		MerchantOrderNo: "",
		ExtUserInfo:     nil,
		QueryOptions:    nil,
	}

	tradeAppPay, err := client.AlipayTradeAppPay(query)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tradeAppPay)
}
