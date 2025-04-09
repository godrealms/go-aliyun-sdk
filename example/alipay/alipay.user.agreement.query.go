package main

import (
	"github.com/godrealms/go-aliyun-sdk/alipay"
	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"log"
)

func main() {
	client := alipay.NewClient()
	client.Sandbox = true
	client.AppId = "2021000147675551"
	client.PrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAxcGXfIfa9ni7rQhyv9PI4C8gxRWJINcGkpEhG7dta3/cxnH2s0Sq2k04Kvfyi3oJYaV4Pu5XrSodocIp0joK+kB6WzppeiKtl0zMqAOAxAiH4vhUoMWspPq0bYdCIcu0oePXupGWly4fpQ8fv4xK4Ns/1a/JZge9pj3PxE7BbzNdo+OmjVDCXDwg3AAYB8uq1R08QETnfmuDBcBmQlr6UllgaNSbT7mjWNYfZN66oIry5nvANWbIjGboQNTBrF2k8R6BUTw/EFNx2SSPMcc9n+fEBe/0sIQKOYWmhU24cIrP89FJl2AcN0MA36/E9pOYwDUVWsNoLW7bUBLDRvUFCQIDAQABAoIBAQCaKOU/teRyuQGMzQYvFhkE8pg4KyagDB/ah/poUeChqaZgqJypeUE7NnEvxeQ06KiBj22imUiN/EPnQqY5dPfzx+naTjnFdV0fkkYEP97UctYDkJTgU9jXL30R2RwrnlIs1aHple7v4YvXr2ePef0xNpmXqhZ1TFmCS7M4lDBkzyQ30YkN87e0WnWe8mDlotCG1mLtKkUa78BGdRAHwUVRWW9oBY6qt8zqotwNd567Vg/sa05MAwQB4fTvvzGAeKq0HM2EpbCYwH/+l5fEexgSR/PGTG7U3yYZYJET8PZveF9kq3Kkch1ECmmcn5b9ANRKiK+ig/J6nmQ2bb6kfwmFAoGBAPq1f5R1pTyxILsux1jZRjpqbRDr934qe3eeb874aFWt/y226oaFLX1oPcEHQ2Pj/fqXqgdLQSUyrpxFHaZWPQHqr98UbUCOuk+uZuXop2VbynUXgcVYd+2bQ6blbrtDBl7V8356f2f6gsCR6dmNQ8NwoI+eILOxqTuCPFl1n9NrAoGBAMnuAZzkFE0PcE+3jdeRzERRYPQlnDS2eNXFcTMAmjDl4XFEsqHvkgSLlRKx6URKsJnUpfGe+PTEef4draFHS2N9DUJtdT+Vyp30cfmzGY7zQ5snxhCXOrkVtBQdyzln/XW8iUIyTjAXDQT0d+/9A1ImUiH0W8Qo3i2lwDhKBBpbAoGBAL4wtBE20lxhbxgxG49+fLVSCV1d9QnkRnfvXJihf90Hu169tdrI60KIsthdnzUHP6Q8LAOOhmQmt6nbEOwf3fY3SKYuA9eGrzZFctAVF/Bfmw4LACpqu8goEkFpKgCrjwigvDCF02NY+poF12ZvsSlxpoxtBBftvhlj6k5fWhjzAoGBAIx0WXjYiGplb7hzM10bU5q9hBOuSW288mW3FRrls1qJu3r+zsWmjslMkZ9UUq2myhnl+JsM+Zu/Qh+IYypIHd8Qr3ViD2jv9uRRkCmf8wokmQTF8JW/qx/sQbqwUpgWhg0r28lIlmwKzIqCxR4PvgyZDQ4DGwVha9ESdpZwnVP7AoGAMteLwlw++UjeajFmPE/QO+LJRpzekiDauwIsCkiWmYaTnXEHEAHBh2UUJqmh4crMa8V9XmxWGT27LNaxDJADw5kGMWl7JDu96mSPqCa1njjWGtsSVN6TgLZnmjP4STHaZh+kc9FlAG0BiVbT/XEEsBSIKdLJg3aeZlNQ3JvrQ9U=
-----END RSA PRIVATE KEY-----`
	client.PublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxcGXfIfa9ni7rQhyv9PI4C8gxRWJINcGkpEhG7dta3/cxnH2s0Sq2k04Kvfyi3oJYaV4Pu5XrSodocIp0joK+kB6WzppeiKtl0zMqAOAxAiH4vhUoMWspPq0bYdCIcu0oePXupGWly4fpQ8fv4xK4Ns/1a/JZge9pj3PxE7BbzNdo+OmjVDCXDwg3AAYB8uq1R08QETnfmuDBcBmQlr6UllgaNSbT7mjWNYfZN66oIry5nvANWbIjGboQNTBrF2k8R6BUTw/EFNx2SSPMcc9n+fEBe/0sIQKOYWmhU24cIrP89FJl2AcN0MA36/E9pOYwDUVWsNoLW7bUBLDRvUFCQIDAQAB
-----END PUBLIC KEY-----`
	client.AlipayPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAniS+LEHNKHSdjQtpwuLzIQFLcMgIWcgdbkIuB4KLYRioqjh5MG5wNuOFDtqRZG6YUUOl7ikhZ/s0XTivkFE1e8eyDhc4fZgFStL3SIYJB1CI0OOKYUDS1/qfrWW/y+PUBJazgIQY8C+6IZ6oC1BeUp+e2k27h4hc9fv/rBL7JYGG7dxpgNQba6WIGB4TSm04e1R/om5cyl6NZpgOFhHWiFbOmUxnn92r137XMvCl5HX0pMENi6S0kueIH2BlOKDuscgE3YYV7TFghJxAVWesSBci/lyGo/MMbRg7jiGXjjmkt9FuVYjfGO8lxb3DHowNpeCXdHQltWc2CHpd0QYcIwIDAQAB
-----END PUBLIC KEY-----`
	client.Http.SetBaseURL("https://openapi-sandbox.dl.alipaydev.com/gateway.do")

	query := &types.UserAgreementQuery{
		AlipayUserId: "2088101122675263",
	}
	response, err := client.AlipayUserAgreementQuery(query)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Query response: %+v\n", response)
}
