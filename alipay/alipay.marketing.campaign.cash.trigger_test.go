package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayMarketingCampaignCashTrigger(t *testing.T) {
	mockResp := `{"alipay_marketing_campaign_cash_trigger_response":{"code":"10000","msg":"Success","award_id":"AWARD20260420001"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	resp, err := client.AlipayMarketingCampaignCashTrigger(context.Background(), &types.CampaignCashTrigger{
		CrowdNo: "20260420001", OpenId: "0680809090909090909090909090", OutBizNo: "biz20260420001",
	})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayMarketingCampaignCashTriggerResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayMarketingCampaignCashTriggerResponse.Code)
	}
	if resp.AlipayMarketingCampaignCashTriggerResponse.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.AlipayMarketingCampaignCashTriggerResponse.Msg)
	}
	if resp.AlipayMarketingCampaignCashTriggerResponse.AwardId != "AWARD20260420001" {
		t.Errorf("AwardId = %q, want AWARD20260420001", resp.AlipayMarketingCampaignCashTriggerResponse.AwardId)
	}
}
