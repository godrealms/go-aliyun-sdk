package main

import (
	"log"
	"os"

	"github.com/godrealms/go-aliyun-sdk/alipay"
	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func main() {
	// Required env vars: ALIPAY_APP_ID
	client := alipay.NewISVClient(
		os.Getenv("ALIPAY_APP_ID"),
		"",
		"",
		"",
	)
	client.Sandbox = true

	u, err := client.GetOpenAuthPageURL(&types.OpenAuthPage{
		RedirectUri: "https://example.com/alipay/callback",
		State:       "order_123",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("授权页 URL: %s", u)
}
