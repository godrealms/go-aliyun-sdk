package types

import "encoding/json"

type CardOpen struct {
	//【描述】请求 ID，幂等控制
	//【示例值】req20260420001
	RequestId string `json:"request_id"`
	//【描述】卡模板 ID，由 AlipayMarketingCardTemplateCreate 返回
	//【示例值】TPL20260420001
	TemplateId string `json:"template_id"`
	//【描述】扩展信息，JSON 字符串（可选）
	ExtInfo string `json:"ext_info,omitempty"`
}

func (r *CardOpen) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

type CardOpenDetail struct {
	PublicResponseParameters
	//【描述】开卡跳转 URL，引导用户在支付宝完成开卡
	//【示例值】https://render.alipay.com/p/s/xxx
	OpenUrl string `json:"open_url"`
}

type AlipayMarketingCardOpenResponse struct {
	AlipayMarketingCardOpenResponse CardOpenDetail `json:"alipay_marketing_card_open_response"`
	Sign                             string         `json:"sign"`
}
