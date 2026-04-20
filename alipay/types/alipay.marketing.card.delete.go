package types

import "encoding/json"

type CardDelete struct {
	//【描述】会员卡 ID，由开卡后获取
	//【示例值】CARD20260420001
	CardId string `json:"card_id"`
	//【描述】卡模板 ID
	//【示例值】TPL20260420001
	TemplateId string `json:"template_id"`
}

func (r *CardDelete) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

type CardDeleteDetail struct {
	PublicResponseParameters
}

type AlipayMarketingCardDeleteResponse struct {
	AlipayMarketingCardDeleteResponse CardDeleteDetail `json:"alipay_marketing_card_delete_response"`
	Sign                               string           `json:"sign"`
}
