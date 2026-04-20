package types

import "encoding/json"

// TradeAdvanceConsult 交易预咨询请求参数
type TradeAdvanceConsult struct {
	//【描述】商户网站唯一订单号
	//【示例值】20150320010101001
	OutTradeNo string `json:"out_trade_no"`
	//【描述】订单总金额，单位为元
	//【示例值】88.00
	TotalAmount string `json:"total_amount"`
	//【描述】订单标题
	//【示例值】大乐透
	Subject string `json:"subject"`
	//【描述】销售产品码
	//【示例值】QUICK_MSECURITY_PAY
	ProductCode string `json:"product_code,omitempty"`
}

func (r *TradeAdvanceConsult) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// TradeAdvanceConsultResponse 交易预咨询响应
type TradeAdvanceConsultResponse struct {
	Code       string `json:"code"`
	Msg        string `json:"msg"`
	ConsultId  string `json:"consult_id"`
	NextAction string `json:"next_action"`
	RiskLevel  string `json:"risk_level,omitempty"`
}

// AlipayTradeAdvanceConsultResponse 交易预咨询响应 wrapper
type AlipayTradeAdvanceConsultResponse struct {
	PublicResponseParameters
	Response TradeAdvanceConsultResponse `json:"alipay_trade_advance_consult_response"`
}
