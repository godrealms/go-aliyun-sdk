package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayTradeRoyaltyRateQuery(t *testing.T) {
	mockResp := `{"alipay_trade_royalty_rate_query_response":{"code":"10000","msg":"Success","royalty_infos":[{"royalty_type":"ROYALTY","rate":"0.05","amount":"","receiver_account":"receiver@example.com"}]}}`
	client, _ := newTestClient(t, mockResp)
	req := &types.TradeRoyaltyRateQuery{
		OutRequestNo: "rq20260420001",
	}
	resp, err := client.AlipayTradeRoyaltyRateQuery(context.Background(), req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.Response.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.Response.Code)
	}
	if resp.Response.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.Response.Msg)
	}
	if len(resp.Response.RoyaltyInfos) != 1 {
		t.Fatalf("RoyaltyInfos len = %d, want 1", len(resp.Response.RoyaltyInfos))
	}
	if resp.Response.RoyaltyInfos[0].Rate != "0.05" {
		t.Errorf("RoyaltyInfos[0].Rate = %q, want 0.05", resp.Response.RoyaltyInfos[0].Rate)
	}
}
