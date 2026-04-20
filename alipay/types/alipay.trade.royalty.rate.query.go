package types

import "encoding/json"

// TradeRoyaltyRateQuery 分账比例查询请求参数
type TradeRoyaltyRateQuery struct {
	//【描述】商户请求号，由商家自定义，需保证在商户端不重复
	//【示例值】20150320010101001
	OutRequestNo string `json:"out_request_no"`
	//【描述】分账类型：ROYALTY（分账）、TRANSFER（转账）等
	//【示例值】ROYALTY
	RoyaltyType string `json:"royalty_type,omitempty"`
}

func (r *TradeRoyaltyRateQuery) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// RoyaltyRateInfo 分账比例信息
type RoyaltyRateInfo struct {
	//【描述】分账类型
	//【示例值】ROYALTY
	RoyaltyType string `json:"royalty_type"`
	//【描述】分账比例，如 0.05 代表 5%
	//【示例值】0.05
	Rate string `json:"rate,omitempty"`
	//【描述】固定分账金额，单位为元
	//【示例值】10.00
	Amount string `json:"amount,omitempty"`
	//【描述】收款方账号
	//【示例值】receiver@example.com
	ReceiverAccount string `json:"receiver_account,omitempty"`
}

// TradeRoyaltyRateQueryResponse 分账比例查询响应
type TradeRoyaltyRateQueryResponse struct {
	Code         string            `json:"code"`
	Msg          string            `json:"msg"`
	RoyaltyInfos []RoyaltyRateInfo `json:"royalty_infos"`
}

// AlipayTradeRoyaltyRateQueryResponse 分账比例查询响应 wrapper
type AlipayTradeRoyaltyRateQueryResponse struct {
	PublicResponseParameters
	Response TradeRoyaltyRateQueryResponse `json:"alipay_trade_royalty_rate_query_response"`
}
