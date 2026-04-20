package types

import "encoding/json"

// TradeRoyaltyRelationBind 分账关系绑定请求参数
type TradeRoyaltyRelationBind struct {
	//【描述】商户请求号，由商家自定义，需保证在商户端不重复
	//【示例值】20150320010101001
	OutRequestNo string `json:"out_request_no"`
	//【描述】分账收款方信息列表，JSON 字符串
	//【示例值】[{"type":"loginName","account":"xxx@example.com","name":"收款方"}]
	ReceiverList string `json:"receiver_list"`
}

func (r *TradeRoyaltyRelationBind) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// TradeRoyaltyRelationBindResponse 分账关系绑定响应
type TradeRoyaltyRelationBindResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// AlipayTradeRoyaltyRelationBindResponse 分账关系绑定响应 wrapper
type AlipayTradeRoyaltyRelationBindResponse struct {
	PublicResponseParameters
	Response TradeRoyaltyRelationBindResponse `json:"alipay_trade_royalty_relation_bind_response"`
}
