package types

import "encoding/json"

type CardTemplateCreate struct {
	//【描述】卡模板名称
	//【示例值】黄金会员卡
	TemplateName string `json:"template_name"`
	//【描述】卡片背景色或背景图 URL
	Background string `json:"background,omitempty"`
	//【描述】商户 Logo URL
	//【示例值】https://example.com/logo.png
	LogoUrl string `json:"logo_url,omitempty"`
	//【描述】通知跳转 URI
	NoticeUri string `json:"notice_uri,omitempty"`
}

func (r *CardTemplateCreate) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

type CardTemplateCreateDetail struct {
	PublicResponseParameters
	//【描述】卡模板 ID
	//【示例值】TPL20260420001
	TemplateId string `json:"template_id"`
}

type AlipayMarketingCardTemplateCreateResponse struct {
	AlipayMarketingCardTemplateCreateResponse CardTemplateCreateDetail `json:"alipay_marketing_card_template_create_response"`
	Sign                                       string                   `json:"sign"`
}
