package main

import (
	"context"
	"github.com/godrealms/go-aliyun-sdk/alipay"
	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"log"
	"os"
	"time"
)

// GetAlipayUserAgreementPageSign 支付宝个人协议页面签约接口(URL)
func GetAlipayUserAgreementPageSign(client *alipay.Client, page *types.AgreementPageSign) {
	url, err := client.GetAlipayUserAgreementPageSign(context.Background(), page)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("url: %+v\n", url)
}

// PostAlipayUserAgreementPageSign 支付宝个人协议页面签约接口(form表单)
func PostAlipayUserAgreementPageSign(client *alipay.Client, page *types.AgreementPageSign) {
	data, err := client.PostAlipayUserAgreementPageSign(context.Background(), page)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("data: %+v\n", data)
}

func main() {
	client := alipay.NewClient()
	client.Sandbox = true
	client.AppId = "2021000147675551"
	// Required env vars: ALIPAY_PRIVATE_KEY, ALIPAY_PUBLIC_KEY, ALIPAY_PUBLIC_KEY_FROM_ALIPAY
	client.PrivateKey = os.Getenv("ALIPAY_PRIVATE_KEY")
	client.PublicKey = os.Getenv("ALIPAY_PUBLIC_KEY")
	client.AlipayPublicKey = os.Getenv("ALIPAY_PUBLIC_KEY_FROM_ALIPAY")
	client.Http.SetBaseURL("https://openapi-sandbox.dl.alipaydev.com/gateway.do")
	page := &types.AgreementPageSign{
		AccessParams: &types.AccessParams{
			Channel: types.ChannelQRCODE,
		},
		PeriodRuleParams: &types.PeriodRuleParams{
			SingleAmount: 0.1,
			PeriodType:   types.PeriodTypeDAY,
			Period:       1,
			ExecuteTime:  time.Now().Format("2006-01-02"),
		},
		ExternalAgreementNo: time.Now().String(), // 商户签约号，代扣协议中标示用户的唯一签约号（确保在商户系统中唯一）
		PersonalProductCode: "CYCLE_PAY_AUTH_P",
		ProductCode:         "CYCLE_PAY_AUTH",
		SignScene:           "INDUSTRY|DEFAULT_SCENE",
	}

	GetAlipayUserAgreementPageSign(client, page)
	//PostAlipayUserAgreementPageSign(client, page)
}

//Py9nLIRWE2dS/dWjDkvp7WQOrL8CjGoq8XiiCdwGt/mT5d7k9gXY+mci32RpSuvvUUSjFjTJAW2s0atzBnonqHo6ZOLq5fuUk56Bar2sHTmrXvygrWDLSQmYGbuQ9k0vC4d6EKDCUHAmN+zz62XNskJ34Q0zRJkyE3ErUKMxRAnl6j+9p+rj31ihIp6XgoXmrhW4skG2qUYCwci0wCY6wiUF8ajvZUxX512N+u07LzCqC5QNB8ImGSOWCAw1b77+O4JDg3dj+yRNWLp+E7+BFPBTb7hOTaPOXLv1l0ZrfZayl/11pftJTnj6ua4QghiEC5DyORmmCEgLiGUkllyFQw==
//Py9nLIRWE2dS/dWjDkvp7WQOrL8CjGoq8XiiCdwGt/mT5d7k9gXY+mci32RpSuvvUUSjFjTJAW2s0atzBnonqHo6ZOLq5fuUk56Bar2sHTmrXvygrWDLSQmYGbuQ9k0vC4d6EKDCUHAmN+zz62XNskJ34Q0zRJkyE3ErUKMxRAnl6j+9p+rj31ihIp6XgoXmrhW4skG2qUYCwci0wCY6wiUF8ajvZUxX512N+u07LzCqC5QNB8ImGSOWCAw1b77+O4JDg3dj+yRNWLp+E7+BFPBTb7hOTaPOXLv1l0ZrfZayl/11pftJTnj6ua4QghiEC5DyORmmCEgLiGUkllyFQw==
