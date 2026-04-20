package main

import (
	"github.com/godrealms/go-aliyun-sdk/alipay"
	"log"
	"net/url"
	"os"
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
