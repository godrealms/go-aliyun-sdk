package types

import "encoding/json"

// TradeClose 统一收单交易关闭接口
type TradeClose struct {
	//【描述】该交易在支付宝系统中的交易流水号。
	//	最短 16 位，最长 64 位。和out_trade_no不能同时为空，
	//	如果同时传了 out_trade_no和 trade_no，则以 trade_no为准。
	//【示例值】2013112611001004680073956707
	TradeNo string `json:"trade_no,omitempty"`
	//【描述】订单支付时传入的商户订单号,和支付宝交易号不能同时为空。
	//	trade_no,out_trade_no如果同时存在优先取trade_no
	//【示例值】HZ0120131127001
	OutTradeNo string `json:"out_trade_no,omitempty"`
	//【描述】商家操作员编号 id，由商家自定义。
	//【示例值】YX01
	OperatorId string `json:"operator_id,omitempty"`
}

func (c *TradeClose) ToString() string {
	marshal, _ := json.Marshal(c)
	return string(marshal)
}

type TradeCloseResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	//【描述】该交易在支付宝系统中的交易流水号。
	//	最短 16 位，最长 64 位。和out_trade_no不能同时为空，
	//	如果同时传了 out_trade_no和 trade_no，则以 trade_no为准。
	//【示例值】2013112611001004680073956707
	TradeNo string `json:"trade_no"`
	//【描述】订单支付时传入的商户订单号,和支付宝交易号不能同时为空。
	//	trade_no,out_trade_no如果同时存在优先取trade_no
	//【示例值】HZ0120131127001
	OutTradeNo string `json:"out_trade_no"`
}

type AlipayTradeCloseResponse struct {
	PublicResponseParameters
	Response TradeCloseResponse `json:"alipay_trade_close_response"`
}
