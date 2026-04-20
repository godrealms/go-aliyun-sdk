package types

import "encoding/json"

type CampaignCashList struct {
	//【描述】页码，从 1 开始
	//【示例值】1
	PageIndex string `json:"page_index"`
	//【描述】每页数量
	//【示例值】10
	PageSize string `json:"page_size"`
}

func (r *CampaignCashList) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

type CampaignCashItem struct {
	//【描述】活动 ID
	CrowdNo string `json:"crowd_no"`
	//【描述】活动名称
	CrowdName string `json:"crowd_name"`
	//【描述】活动状态：ONGOING | ENDED | CLOSED
	Status string `json:"status"`
	//【描述】活动开始时间
	StartTime string `json:"start_time"`
	//【描述】活动结束时间
	EndTime string `json:"end_time"`
}

type CampaignCashListDetail struct {
	PublicResponseParameters
	//【描述】总条数
	TotalCount int `json:"total_count"`
	//【描述】活动列表
	ResultList []CampaignCashItem `json:"result_list"`
}

type AlipayMarketingCampaignCashListResponse struct {
	AlipayMarketingCampaignCashListResponse CampaignCashListDetail `json:"alipay_marketing_campaign_cash_list_response"`
	Sign                                    string                 `json:"sign"`
}
