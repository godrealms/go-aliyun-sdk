package types

import "encoding/json"

// TradeOrderinfoSync 商户订单信息同步请求参数
type TradeOrderinfoSync struct {
	//【描述】支付宝交易号，与 out_trade_no 不能同时为空
	//【示例值】2014112611001004680073956707
	TradeNo string `json:"trade_no,omitempty"`
	//【描述】商户订单号，与 trade_no 不能同时为空
	//【示例值】20150320010101001
	OutTradeNo string `json:"out_trade_no,omitempty"`
	//【描述】商户请求号，由商家自定义，需保证在商户端不重复
	//【示例值】20150320010101001
	OutRequestNo string `json:"out_request_no"`
	//【描述】订单同步类型，如 CREDIT_ADVANCE（先享后付）
	//【示例值】CREDIT_ADVANCE
	OrderType string `json:"order_type"`
	//【描述】订单业务场景
	//【示例值】CREDIT_ADVANCE_SETTLE
	OrderScene string `json:"order_scene"`
	//【描述】芝麻信用业务单号，先享后付场景下必填
	//【示例值】ZMCB99202103310000450000041833
	CreditBizOrderId string `json:"credit_biz_order_id,omitempty"`
	//【描述】扩展信息，JSON 字符串
	ExtInfo string `json:"ext_info,omitempty"`
}

func (r *TradeOrderinfoSync) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// TradeOrderinfoSyncResponse 商户订单信息同步响应
type TradeOrderinfoSyncResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// AlipayTradeOrderinfoSyncResponse 商户订单信息同步响应 wrapper
type AlipayTradeOrderinfoSyncResponse struct {
	PublicResponseParameters
	Response TradeOrderinfoSyncResponse `json:"alipay_trade_orderinfo_sync_response"`
}
