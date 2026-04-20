package types

import "encoding/json"

// TradeRoyaltyRelationUnbind 分账关系解绑请求参数
type TradeRoyaltyRelationUnbind struct {
	//【描述】商户请求号，由商家自定义，需保证在商户端不重复
	//【示例值】20150320010101001
	OutRequestNo string `json:"out_request_no"`
	//【描述】分账收款方信息列表，JSON 字符串
	//【示例值】[{"type":"loginName","account":"xxx@example.com","name":"收款方"}]
	ReceiverList string `json:"receiver_list"`
}

func (r *TradeRoyaltyRelationUnbind) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// TradeRoyaltyRelationUnbindResponse 分账关系解绑响应
type TradeRoyaltyRelationUnbindResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// AlipayTradeRoyaltyRelationUnbindResponse 分账关系解绑响应 wrapper
type AlipayTradeRoyaltyRelationUnbindResponse struct {
	PublicResponseParameters
	Response TradeRoyaltyRelationUnbindResponse `json:"alipay_trade_royalty_relation_unbind_response"`
}
