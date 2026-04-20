package types

import "encoding/json"

type CardTemplateQuery struct {
	//【描述】卡模板 ID，由 AlipayMarketingCardTemplateCreate 返回
	//【示例值】TPL20260420001
	TemplateId string `json:"template_id"`
}

func (r *CardTemplateQuery) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

type CardTemplateQueryDetail struct {
	PublicResponseParameters
	TemplateId   string `json:"template_id"`
	TemplateName string `json:"template_name"`
	//【描述】模板状态：NORMAL | STOP
	Status string `json:"status"`
}

type AlipayMarketingCardTemplateQueryResponse struct {
	AlipayMarketingCardTemplateQueryResponse CardTemplateQueryDetail `json:"alipay_marketing_card_template_query_response"`
	Sign                                     string                  `json:"sign"`
}
