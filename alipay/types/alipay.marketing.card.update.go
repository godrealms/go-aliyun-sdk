package types

import "encoding/json"

type CardUpdate struct {
	//【描述】会员卡 ID，由开卡后获取
	//【示例值】CARD20260420001
	CardId string `json:"card_id"`
	//【描述】卡模板 ID
	//【示例值】TPL20260420001
	TemplateId string `json:"template_id"`
	//【描述】余额/积分更新信息，JSON 字符串，如 {"balance":"200.00"}
	BalanceInfo string `json:"balance_info,omitempty"`
}

func (r *CardUpdate) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

type CardUpdateDetail struct {
	PublicResponseParameters
}

type AlipayMarketingCardUpdateResponse struct {
	AlipayMarketingCardUpdateResponse CardUpdateDetail `json:"alipay_marketing_card_update_response"`
	Sign                              string           `json:"sign"`
}
