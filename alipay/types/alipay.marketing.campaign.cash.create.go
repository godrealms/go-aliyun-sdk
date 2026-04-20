package types

import "encoding/json"

type CampaignCashCreate struct {
	//【描述】活动名称
	//【示例值】春季促销活动
	CrowdName string `json:"crowd_name"`
	//【描述】奖励类型：CASH（现金）
	//【示例值】CASH
	PrizeType string `json:"prize_type"`
	//【描述】预算信息，JSON 字符串，如 {"total_budget":"1000.00"}
	BudgetInfo string `json:"budget_info"`
	//【描述】奖励信息，JSON 字符串，如 {"prize_amount":"10.00"}
	PrizeInfo string `json:"prize_info"`
	//【描述】发放渠道：ALISEND（支付宝）| SMSSEND（短信）
	//【示例值】ALISEND
	SendChannel string `json:"send_channel"`
	//【描述】活动开始时间，格式 yyyy-MM-dd HH:mm:ss
	//【示例值】2026-04-20 00:00:00
	StartTime string `json:"start_time"`
	//【描述】活动结束时间，格式 yyyy-MM-dd HH:mm:ss
	//【示例值】2026-05-20 23:59:59
	EndTime string `json:"end_time"`
	//【描述】商户推广链接
	MerchantLink string `json:"merchant_link,omitempty"`
}

func (r *CampaignCashCreate) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

type CampaignCashCreateDetail struct {
	PublicResponseParameters
	//【描述】活动 ID
	//【示例值】20260420001
	CrowdNo string `json:"crowd_no"`
}

type AlipayMarketingCampaignCashCreateResponse struct {
	AlipayMarketingCampaignCashCreateResponse CampaignCashCreateDetail `json:"alipay_marketing_campaign_cash_create_response"`
	Sign                                       string                   `json:"sign"`
}
