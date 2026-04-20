package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayTradeAdvanceConsult(t *testing.T) {
	mockResp := `{"alipay_trade_advance_consult_response":{"code":"10000","msg":"Success","consult_id":"consult20260420001","next_action":"ACCEPT","risk_level":"LOW"}}`
	client, _ := newTestClient(t, mockResp)
	req := &types.TradeAdvanceConsult{
		OutTradeNo:  "consult20260420001",
		TotalAmount: "88.00",
		Subject:     "测试预咨询",
	}
	resp, err := client.AlipayTradeAdvanceConsult(req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.Response.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.Response.Code)
	}
	if resp.Response.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.Response.Msg)
	}
	if resp.Response.ConsultId != "consult20260420001" {
		t.Errorf("ConsultId = %q unexpected", resp.Response.ConsultId)
	}
	if resp.Response.NextAction != "ACCEPT" {
		t.Errorf("NextAction = %q, want ACCEPT", resp.Response.NextAction)
	}
}
