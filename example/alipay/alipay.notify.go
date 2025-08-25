package main

import (
	"github.com/godrealms/go-aliyun-sdk/alipay"
	"log"
	"net/url"
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

	values, err := url.ParseQuery("notify_type=trade_status_sync&notify_id=91722adff935e8cfa58b3aabf4dead6ibe&notify_time=2017-02-16 21:46:15&sign_type=RSA2&sign=WcO+t3D8Kg71dTlKwN7r9PzUOXeaBJwp8/FOuSxcuSkXsoVYxBpsAidprySCjHCjmaglNcjoKJQLJ28/Asl93joTW39FX6i07lXhnbPknezAlwmvPdnQuI01HZsZF9V1i6ggZjBiAd5lG8bZtTxZOJ87ub2i9GuJ3Nr/NUc9VeY=&receipt_amount=null&invoice_amount=null&buyer_pay_amount=null&point_amount=null&voucher_detail_list=null&buyer_logon_id=null&seller_email=null&out_biz_no=null&passback_params=null&out_channel_type=null&trade_no=null&app_id=null&out_trade_no=null&seller_id=null&trade_status=null&total_amount=null&refund_fee=null&subject=null&body=null&gmt_create=null&gmt_payment=null&gmt_refund=null&gmt_close=null&buyer_id=null&fund_bill_list=[{\"amount\":\"50.00\",\"fundChannel\":\"ALIPAYACCOUNT\"}]&notify_action_type=payByAccountAction/closeTradeAction/reverseAction/finishFPAction/confirmDisburseAction/financingReceiptAction&current_seller_received_amount=88.88&seller_received_total_amount=88.88&total_from_seller_fee=88.88&ff_current_period=1&discount_amount=88.88&charge_amount=8.88&charge_flags=bluesea_1&settlement_id=2018101610032004620239146945&industry_sepc_detail={\"registration_order_pay\":{\"brlx\":\"1\",\"cblx\":\"1\"}}&industry_sepc_detail_acc={\"registration_order_pay\":{\"brlx\":\"1\",\"cblx\":\"1\"}}&industry_sepc_detail_gov={\"registration_order_pay\":{\"brlx\":\"1\",\"cblx\":\"1\"}}&discount_goods_detail=\"[{\\\"goodsId\\\":\\\"STANDARD1026181538\\\",\\\"goodsName\\\":\\\"雪碧\\\",\\\"discountAmount\\\":\\\"10.00\\\"}]\"&mdiscount_amount=88.88&hb_fq_pay_info={\"USER_INSTALL_NUM\":\"3\"}&receipt_currency_type=DC&enterprise_pay_info={\"invoice_amount\":\"28.00\",\"is_use_enterprise_pay\":\"true\"}&hyb_amount=10.24&charge_info_list=[{\"charge_fee\":\"0.01\",\"original_charge_fee\":\"0.02\",\"switch_fee_rate\":\"0.03\",\"is_rating_on_trade_receiver\":\"Y\",\"is_rating_on_switch\":\"Y\"}]&medical_insurance_info={\"medicareCardType\":\"1\",\"medicareCardHolderHiddenName\":\"**专\"}&credit_pay_mode=creditAdvanceV2&buyer_charge_amt=1.00&fulfillment_amount=10.24&cashier_type=APP&refund_voucher_detail_list=[{\"amount\":\"0.10\",\"merchantContribute\":\"0.00\",\"name\":\"工具立减_固定金额\",\"otherContribute\":\"0.10\",\"type\":\"DISCOUNT\",\"voucherId\":\"20241207000730027708000CRGZP\"}]")
	if err != nil {
		log.Fatal(err)
	}
	notify, err := client.Notify(values)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(notify)
}
