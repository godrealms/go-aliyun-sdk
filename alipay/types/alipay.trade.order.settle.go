package types

import "encoding/json"

// TradeOrderSettle 统一收单交易结算请求参数
type TradeOrderSettle struct {
	//【描述】结算请求流水号，由商家自定义，需保证在商户端不重复
	//【示例值】20150320010101001
	OutRequestNo string `json:"out_request_no"`
	//【描述】支付宝交易号
	//【示例值】2014112611001004680073956707
	TradeNo string `json:"trade_no"`
	//【描述】分账明细信息，JSON 字符串，包含收款方账号、金额、描述等
	//【示例值】[{"trans_in":"2088xxx","amount":"10.00","desc":"分账"}]
	RoyaltyParameters string `json:"royalty_parameters"`
	//【描述】商户操作员编号
	//【示例值】yx_001
	OperatorId string `json:"operator_id,omitempty"`
}

func (r *TradeOrderSettle) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// TradeOrderSettleResponse 统一收单交易结算响应
type TradeOrderSettleResponse struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	TradeNo string `json:"trade_no"`
}

// AlipayTradeOrderSettleResponse 统一收单交易结算响应 wrapper
type AlipayTradeOrderSettleResponse struct {
	PublicResponseParameters
	Response TradeOrderSettleResponse `json:"alipay_trade_order_settle_response"`
}
