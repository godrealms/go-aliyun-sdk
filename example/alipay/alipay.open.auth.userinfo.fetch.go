package main

import (
	"log"
	"os"

	"github.com/godrealms/go-aliyun-sdk/alipay"
	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func main() {
	// Required env vars: ALIPAY_APP_ID, ALIPAY_PRIVATE_KEY, ALIPAY_PUBLIC_KEY_FROM_ALIPAY, ALIPAY_APP_AUTH_TOKEN
	client := alipay.NewISVClient(
		os.Getenv("ALIPAY_APP_ID"),
		os.Getenv("ALIPAY_PRIVATE_KEY"),
		os.Getenv("ALIPAY_PUBLIC_KEY_FROM_ALIPAY"),
		os.Getenv("ALIPAY_APP_AUTH_TOKEN"),
	)
	client.Sandbox = true
	client.Http.SetBaseURL("https://openapi-sandbox.dl.alipaydev.com/gateway.do")

	resp, err := client.AlipayOpenAuthUserinfoFetch(&types.OpenAuthUserinfoFetch{
		AppAuthToken: os.Getenv("ALIPAY_APP_AUTH_TOKEN"),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("UserId: %s, Name: %s", resp.AlipayOpenAuthUserinfoFetchResponse.UserId, resp.AlipayOpenAuthUserinfoFetchResponse.Name)
}
