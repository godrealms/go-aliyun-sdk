package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayMarketingCampaignCashList(t *testing.T) {
	mockResp := `{"alipay_marketing_campaign_cash_list_response":{"code":"10000","msg":"Success","total_count":1,"result_list":[{"crowd_no":"20260420001","crowd_name":"春季促销活动","status":"ONGOING","start_time":"2026-04-20 00:00:00","end_time":"2026-05-20 23:59:59"}]},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	req := &types.CampaignCashList{PageIndex: "1", PageSize: "10"}
	resp, err := client.AlipayMarketingCampaignCashList(context.Background(), req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayMarketingCampaignCashListResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayMarketingCampaignCashListResponse.Code)
	}
	if resp.AlipayMarketingCampaignCashListResponse.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.AlipayMarketingCampaignCashListResponse.Msg)
	}
	if resp.AlipayMarketingCampaignCashListResponse.TotalCount != 1 {
		t.Errorf("TotalCount = %d, want 1", resp.AlipayMarketingCampaignCashListResponse.TotalCount)
	}
	if len(resp.AlipayMarketingCampaignCashListResponse.ResultList) != 1 {
		t.Fatalf("ResultList len = %d, want 1", len(resp.AlipayMarketingCampaignCashListResponse.ResultList))
	}
	if resp.AlipayMarketingCampaignCashListResponse.ResultList[0].CrowdNo != "20260420001" {
		t.Errorf("CrowdNo = %q, want 20260420001", resp.AlipayMarketingCampaignCashListResponse.ResultList[0].CrowdNo)
	}
}
