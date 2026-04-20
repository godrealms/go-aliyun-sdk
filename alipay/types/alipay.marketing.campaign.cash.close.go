package types

import "encoding/json"

type CampaignCashClose struct {
	//【描述】活动 ID，由 AlipayMarketingCampaignCashCreate 返回
	//【示例值】20260420001
	CrowdNo string `json:"crowd_no"`
}

func (r *CampaignCashClose) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

type CampaignCashCloseDetail struct {
	PublicResponseParameters
}

type AlipayMarketingCampaignCashCloseResponse struct {
	AlipayMarketingCampaignCashCloseResponse CampaignCashCloseDetail `json:"alipay_marketing_campaign_cash_close_response"`
	Sign                                      string                  `json:"sign"`
}
