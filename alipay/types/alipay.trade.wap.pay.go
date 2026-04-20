package types

import "encoding/json"

// TradeWapPay H5手机网站支付请求参数
type TradeWapPay struct {
	//【描述】商户网站唯一订单号
	//【示例值】70501111111S001111119
	OutTradeNo string `json:"out_trade_no"`
	//【描述】订单标题
	//【示例值】大乐透
	Subject string `json:"subject"`
	//【描述】订单总金额，单位为元，精确到小数点后两位
	//【示例值】9.00
	TotalAmount string `json:"total_amount"`
	//【描述】销售产品码，H5场景固定值 QUICK_WAP_WAY
	//【示例值】QUICK_WAP_WAY
	ProductCode string `json:"product_code,omitempty"`
	//【描述】订单描述
	//【示例值】Iphone6 16G
	Body string `json:"body,omitempty"`
	//【描述】用户付款中途退出返回商户网站的地址
	//【示例值】https://m.alipay.com/Gk8NF23
	QuitUrl string `json:"quit_url,omitempty"`
	//【描述】绝对超时时间，格式为yyyy-MM-dd HH:mm:ss
	//【示例值】2016-12-31 10:05:00
	TimeExpire string `json:"time_expire,omitempty"`
	//【描述】公用回传参数
	PassbackParams string `json:"passback_params,omitempty"`
	//【描述】可用渠道，多个渠道以逗号分割
	EnablePayChannels string `json:"enable_pay_channels,omitempty"`
	//【描述】禁用渠道，多个渠道以逗号分割
	DisablePayChannels string `json:"disable_pay_channels,omitempty"`
}

func (r *TradeWapPay) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}
