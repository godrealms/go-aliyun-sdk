package types

import "encoding/json"

// TradeCreate 统一收单下单请求参数
type TradeCreate struct {
	//【描述】商户网站唯一订单号
	//【示例值】20150320010101001
	OutTradeNo string `json:"out_trade_no"`
	//【描述】订单标题
	//【示例值】大乐透
	Subject string `json:"subject"`
	//【描述】订单总金额，单位为元，精确到小数点后两位
	//【示例值】88.00
	TotalAmount string `json:"total_amount"`
	//【描述】买家的支付宝用户id
	//【示例值】2088102146225135
	BuyerId string `json:"buyer_id,omitempty"`
	//【描述】买家支付宝账号
	//【示例值】15900000001
	BuyerLogonId string `json:"buyer_logon_id,omitempty"`
	//【描述】销售产品码
	//【示例值】FACE_TO_FACE_PAYMENT
	ProductCode string `json:"product_code,omitempty"`
	//【描述】订单描述
	//【示例值】Iphone6 16G
	Body string `json:"body,omitempty"`
	//【描述】商户操作员编号
	//【示例值】yx_001
	OperatorId string `json:"operator_id,omitempty"`
	//【描述】商户门店编号
	//【示例值】NJ_S_001
	StoreId string `json:"store_id,omitempty"`
}

func (r *TradeCreate) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// TradeCreateResponse 统一收单下单响应
type TradeCreateResponse struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	TradeNo string `json:"trade_no"`
}

// AlipayTradeCreateResponse 统一收单下单响应 wrapper
type AlipayTradeCreateResponse struct {
	PublicResponseParameters
	Response TradeCreateResponse `json:"alipay_trade_create_response"`
}
