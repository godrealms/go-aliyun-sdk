package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayMarketingCampaignCashCreate(t *testing.T) {
	mockResp := `{"alipay_marketing_campaign_cash_create_response":{"code":"10000","msg":"Success","crowd_no":"20260420001"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	req := &types.CampaignCashCreate{
		CrowdName:   "春季促销活动",
		PrizeType:   "CASH",
		BudgetInfo:  `{"total_budget":"1000.00"}`,
		PrizeInfo:   `{"prize_amount":"10.00"}`,
		SendChannel: "ALISEND",
		StartTime:   "2026-04-20 00:00:00",
		EndTime:     "2026-05-20 23:59:59",
	}
	resp, err := client.AlipayMarketingCampaignCashCreate(req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayMarketingCampaignCashCreateResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayMarketingCampaignCashCreateResponse.Code)
	}
	if resp.AlipayMarketingCampaignCashCreateResponse.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.AlipayMarketingCampaignCashCreateResponse.Msg)
	}
	if resp.AlipayMarketingCampaignCashCreateResponse.CrowdNo != "20260420001" {
		t.Errorf("CrowdNo = %q, want 20260420001", resp.AlipayMarketingCampaignCashCreateResponse.CrowdNo)
	}
}
