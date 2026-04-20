package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayTradeOrderinfoSync(t *testing.T) {
	mockResp := `{"alipay_trade_orderinfo_sync_response":{"code":"10000","msg":"Success"}}`
	client, _ := newTestClient(t, mockResp)
	req := &types.TradeOrderinfoSync{
		TradeNo:      "2026042022001480551404403003",
		OutRequestNo: "sync20260420001",
		OrderType:    "CREDIT_ADVANCE",
		OrderScene:   "CREDIT_ADVANCE_SETTLE",
	}
	resp, err := client.AlipayTradeOrderinfoSync(req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.Response.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.Response.Code)
	}
	if resp.Response.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.Response.Msg)
	}
}
