package types

type AlipayTradePreCreate struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	//【描述】商户的订单号
	//【示例值】6823789339978248
	OutTradeNo string `json:"out_trade_no"`
	//【描述】当前预下单请求生成的二维码码串，有效时间2小时，可以用二维码生成工具根据该码串值生成对应的二维码
	//【示例值】https://qr.alipay.com/bavh4wjlxf12tper3a
	QrCode string `json:"qr_code"`
}

type AlipayTradePreCreateResponse struct {
	PublicResponseParameters
	AlipayTradePreCreateResponse AlipayTradePreCreate `json:"alipay_trade_precreate_response"`
}
