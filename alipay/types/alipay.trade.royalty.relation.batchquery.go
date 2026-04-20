package types

import "encoding/json"

// TradeRoyaltyRelationBatchquery 分账关系批量查询请求参数
type TradeRoyaltyRelationBatchquery struct {
	//【描述】商户请求号，由商家自定义，需保证在商户端不重复
	//【示例值】20150320010101001
	OutRequestNo string `json:"out_request_no"`
	//【描述】待查询的分账收款方信息列表，JSON 字符串，不传则返回全部绑定关系
	ReceiverList string `json:"receiver_list,omitempty"`
	//【描述】页码，从1开始
	//【示例值】1
	PageNum int `json:"page_num,omitempty"`
	//【描述】每页条数
	//【示例值】20
	PageSize int `json:"page_size,omitempty"`
}

func (r *TradeRoyaltyRelationBatchquery) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// RoyaltyReceiverInfo 分账收款方信息
type RoyaltyReceiverInfo struct {
	//【描述】账号类型：userId、loginName 等
	ReceiverType string `json:"receiver_type"`
	//【描述】收款方账号
	ReceiverAccount string `json:"receiver_account"`
	//【描述】收款方姓名
	ReceiverName string `json:"receiver_name,omitempty"`
	//【描述】关系类型
	RelationType string `json:"relation_type,omitempty"`
}

// TradeRoyaltyRelationBatchqueryResponse 分账关系批量查询响应
type TradeRoyaltyRelationBatchqueryResponse struct {
	Code          string                `json:"code"`
	Msg           string                `json:"msg"`
	Count         int                   `json:"count"`
	ReceiverInfos []RoyaltyReceiverInfo `json:"receiver_infos"`
}

// AlipayTradeRoyaltyRelationBatchqueryResponse 分账关系批量查询响应 wrapper
type AlipayTradeRoyaltyRelationBatchqueryResponse struct {
	PublicResponseParameters
	Response TradeRoyaltyRelationBatchqueryResponse `json:"alipay_trade_royalty_relation_batchquery_response"`
}
