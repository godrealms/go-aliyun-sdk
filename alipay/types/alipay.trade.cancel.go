package types

import "encoding/json"

// TradeCancel 统一收单交易撤销接口
type TradeCancel struct {
	//【描述】原支付请求的商户订单号,和支付宝交易号不能同时为空
	//【示例值】20150320010101001
	OutTradeNo string `json:"out_trade_no"`
	//【描述】支付宝交易号，和商户订单号不能同时为空
	//【示例值】2014112611001004680073956707
	TradeNo string `json:"trade_no"`
}

func (c *TradeCancel) ToString() string {
	marshal, _ := json.Marshal(c)
	return string(marshal)
}

type TradeCancelResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	//【描述】支付宝交易号; 当发生交易关闭或交易退款时返回；
	//【示例值】2013112011001004330000121536
	TradeNo string `json:"trade_no,omitempty"`
	//【描述】商户订单号
	//【示例值】6823789339978248
	OutTradeNo string `json:"out_trade_no"`
	//【描述】是否需要重试
	//【示例值】N
	RetryFlag string `json:"retry_flag"`
	//【描述】本次撤销触发的交易动作,接口调用成功且交易存在时返回。可能的返回值：
	// 	close：交易未支付，触发关闭交易动作，无退款；
	// 	refund：交易已支付，触发交易退款动作；
	// 	未返回：未查询到交易，或接口调用失败；
	//【示例值】close
	Action string `json:"action,omitempty"`
}

type AlipayTradeCancelResponse struct {
	PublicResponseParameters
	Response TradeCancelResponse `json:"alipay_trade_cancel_response"`
}
