package types

import "encoding/json"

type CampaignCashTrigger struct {
	//【描述】活动 ID，由 AlipayMarketingCampaignCashCreate 返回
	//【示例值】20260420001
	CrowdNo string `json:"crowd_no"`
	//【描述】用户支付宝 openId
	//【示例值】0680809090909090909090909090
	OpenId string `json:"open_id"`
	//【描述】外部业务号，幂等控制
	//【示例值】biz20260420001
	OutBizNo string `json:"out_biz_no"`
}

func (r *CampaignCashTrigger) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

type CampaignCashTriggerDetail struct {
	PublicResponseParameters
	//【描述】奖励记录 ID
	//【示例值】AWARD20260420001
	AwardId string `json:"award_id"`
}

type AlipayMarketingCampaignCashTriggerResponse struct {
	AlipayMarketingCampaignCashTriggerResponse CampaignCashTriggerDetail `json:"alipay_marketing_campaign_cash_trigger_response"`
	Sign                                        string                    `json:"sign"`
}
