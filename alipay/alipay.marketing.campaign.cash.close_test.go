package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayMarketingCampaignCashClose(t *testing.T) {
	mockResp := `{"alipay_marketing_campaign_cash_close_response":{"code":"10000","msg":"Success"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	resp, err := client.AlipayMarketingCampaignCashClose(context.Background(), &types.CampaignCashClose{CrowdNo: "20260420001"})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayMarketingCampaignCashCloseResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayMarketingCampaignCashCloseResponse.Code)
	}
	if resp.AlipayMarketingCampaignCashCloseResponse.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.AlipayMarketingCampaignCashCloseResponse.Msg)
	}
}
