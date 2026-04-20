package types

import "encoding/json"

type CardQuery struct {
	//【描述】会员卡 ID（与 TemplateId 至少填一个）
	//【示例值】CARD20260420001
	CardId string `json:"card_id,omitempty"`
	//【描述】卡模板 ID（与 CardId 至少填一个）
	//【示例值】TPL20260420001
	TemplateId string `json:"template_id,omitempty"`
}

func (r *CardQuery) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

type CardQueryDetail struct {
	PublicResponseParameters
	CardId     string `json:"card_id"`
	TemplateId string `json:"template_id"`
	//【描述】卡状态：OPENED | CLOSED
	Status  string `json:"status"`
	Balance string `json:"balance,omitempty"`
	Point   string `json:"point,omitempty"`
}

type AlipayMarketingCardQueryResponse struct {
	AlipayMarketingCardQueryResponse CardQueryDetail `json:"alipay_marketing_card_query_response"`
	Sign                              string          `json:"sign"`
}
