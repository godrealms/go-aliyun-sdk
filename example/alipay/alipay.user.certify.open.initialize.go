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

	resp, err := client.AlipayUserCertifyOpenInitialize(context.Background(), &types.UserCertifyOpenInitialize{
		OuterOrderNo:   "order_certify_001",
		BizCode:        "FACE",
		IdentityParam:  `{"identity_type":"CERT_INFO","cert_type":"IDENTITY_CARD","cert_name":"张三","cert_no":"310000199001011234"}`,
		MerchantConfig: `{"return_url":"https://example.com/certify/callback"}`,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CertifyId: %s", resp.AlipayUserCertifyOpenInitializeResponse.CertifyId)
}
