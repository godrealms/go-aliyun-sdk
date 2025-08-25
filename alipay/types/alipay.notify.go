package types

type Notify struct {
	// 通知时间。通知的发送时间。格式为 yyyy-MM-dd HH:mm:ss。
	NotifyTime string `json:"notify_time" form:"notify_time"`
	// 通知类型。
	// 枚举值：trade_status_sync。
	NotifyType string `json:"notify_type" form:"notify_type"`
	// 通知校验 ID。
	NotifyId string `json:"notify_id" form:"notify_id"`
	// 签名类型。商家生成签名字符串所使用的签名算法类型，
	// 目前支持 RSA2 和 RSA，推荐使用 RSA2（如果开发者手动验签，不使用 SDK 验签，可以不传此参数）。
	SignType string `json:"sign_type" form:"sign_type"`
	// 签名。可查看异步返回结果的验签（如果开发者手动验签，不使用 SDK 验签，可以不传此参数）。
	Sign string `json:"sign" form:"sign"`
	// 支付宝交易号。支付宝交易凭证号。
	TradeNo string `json:"trade_no" form:"trade_no"`
	// 开发者的 app_id。支付宝分配给开发者的应用 APPID。
	AppId string `json:"app_id" form:"app_id"`
	// 开发者的 app_id，在服务商调用的场景下为授权方的 app_id。	2014072300007148
	AuthAppId string `json:"auth_app_id" form:"auth_app_id"`
	// 商户订单号。	6823789339978248
	OutTradeNo string `json:"out_trade_no" form:"out_trade_no"`
	// 商家业务号。商家业务 ID，主要是退款通知中返回退款申请的流水号。	HZRF001
	OutBizNo string `json:"out_biz_no" form:"out_biz_no"`
	// 买家支付宝用户号。买家支付宝账号对应的支付宝唯一用户号。新商户建议使用open_id替代该字段。对于新商户，user_id字段未来计划逐步回收，存量商户可继续使用。如使用open_id，请确认 应用-开发配置-openid配置管理 已启用。无该配置项，可查看openid配置申请。	-
	BuyerOpenId string `json:"buyer_open_id" form:"buyer_open_id"`
	// 买家支付宝账号。	180****0062
	BuyerLogonId string `json:"buyer_logon_id" form:"buyer_logon_id"`
	// 卖家支付宝用户号。	2088101106499364
	SellerId string `json:"seller_id" form:"seller_id"`
	// 卖家支付宝账号。	zhuzhanghu@alitest.com
	SellerEmail string `json:"seller_email" form:"seller_email"`
	//【描述】交易状态：
	//	WAIT_BUYER_PAY（交易创建，等待买家付款）、
	//	TRADE_CLOSED（未付款交易超时关闭，或支付完成后全额退款）、
	//	TRADE_SUCCESS（交易支付成功）、
	//	TRADE_FINISHED（交易结束，不可退款）
	//【示例值】TRADE_CLOSED
	TradeStatus string `json:"trade_status" form:"trade_status"`
	// 订单金额。本次交易支付的订单金额，单位为人民币（元）。支持小数点后两位。	20
	TotalAmount float64 `json:"total_amount" form:"total_amount"`
	// 实收金额。商家在交易中实际收到的款项，单位为人民币（元）。支持小数点后两位。	15
	ReceiptAmount float64 `json:"receipt_amount" form:"receipt_amount"`
	// 开票金额。用户在交易中支付的可开发票的金额。支持小数点后两位。	10.00
	InvoiceAmount float64 `json:"invoice_amount" form:"invoice_amount"`
	// 付款金额。用户在交易中支付的金额。支持小数点后两位。	13.88
	BuyerPayAmount float64 `json:"buyer_pay_amount" form:"buyer_pay_amount"`
	// 集分宝金额。使用集分宝支付的金额。支持小数点后两位。	12.00
	PointAmount float64 `json:"point_amount" form:"point_amount"`
	// 总退款金额。退款通知中，返回总退款金额，单位为元，支持小数点后两位。	2.58
	RefundFee float64 `json:"refund_fee" form:"refund_fee"`
	// 实际退款金额。商家实际退款给用户的金额，单位为元，支持小数点后两位。	2.08
	SendBackFee float64 `json:"send_back_fee" form:"send_back_fee"`
	// 订单标题。商品的标题/交易标题/订单标题/订单关键字等，是请求时对应的参数，原样通知回来。	XXX交易
	Subject string `json:"subject" form:"subject"`
	// 。该订单的备注、描述、明细等。对应请求时的 Body 参数，原样通知回来。	XXX交易内容
	Body string `json:"body" form:"body"`
	// 公共回传参数，如果请求时传递了该参数，则返回给商家时会在异步通知时将该参数原样返回。本参数必须进行 UrlEncode 之后才可以发送给支付宝。	-
	PassbackParams string `json:"passback_params" form:"passback_params"`
	// 交易创建时间。该笔交易创建的时间。格式 为 yyyy-MM-dd HH:mm:ss。	2015-04-27 15:45:57
	GmtCreate string `json:"gmt_create" form:"gmt_create"`
	// 交易 付款时间。该笔交易的买家付款时间。格式为 yyyy-MM-dd HH:mm:ss。	2015-04-27 15:45:57
	GmtPayment string `json:"gmt_payment" form:"gmt_payment"`
	// 交易退款时间。该笔交易的退款时间。格式 为 yyyy-MM-dd HH:mm:ss.SS。	2015-04-28 15:45:57.320
	GmtRefund string `json:"gmt_refund" form:"gmt_refund"`
	// 交易结束时间。该笔交易结束时间。格式为 yyyy-MM-dd HH:mm:ss。	2015-04-29 15:45:57
	GmtClose string `json:"gmt_close" form:"gmt_close"`
	// 支付金额信息。支付成功的各个渠道金额信息，详请可查看下表 资金明细信息说明 。	[{"amount":"15.00", "fundChannel":"ALIPAYACCOUNT"}]
	FundBillList string `json:"fund_bill_list" form:"fund_bill_list"`
	// 优惠券信息。本交易支付时所使用的所有优惠券信息，详请可查看下表 优惠券信息说明 。	[{"amount":"0.20","merchantContribute":"0.00","name":"一键创建券模板的券名称"，"otherContribute":"0.20","type":"ALIPAY_BIZ_VOUCHER","memo":"学生8折优惠"}]
	VoucherDetailList string `json:"voucher_detail_list" form:"voucher_detail_list"`
}
