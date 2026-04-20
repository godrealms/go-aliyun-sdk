package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayTradeOrderSettle(t *testing.T) {
	mockResp := `{"alipay_trade_order_settle_response":{"code":"10000","msg":"Success","trade_no":"2026042022001480551404403002"}}`
	client, _ := newTestClient(t, mockResp)
	req := &types.TradeOrderSettle{
		OutRequestNo:      "settle20260420001",
		TradeNo:           "2026042022001480551404403002",
		RoyaltyParameters: `[{"trans_in":"2088102146225136","amount":"10.00","desc":"分账给供应商"}]`,
	}
	resp, err := client.AlipayTradeOrderSettle(req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.Response.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.Response.Code)
	}
	if resp.Response.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.Response.Msg)
	}
	if resp.Response.TradeNo != "2026042022001480551404403002" {
		t.Errorf("TradeNo = %q unexpected", resp.Response.TradeNo)
	}
}
