package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayTradeCreate(t *testing.T) {
	mockResp := `{"alipay_trade_create_response":{"code":"10000","msg":"Success","trade_no":"2026042022001480551404403001"}}`
	client, _ := newTestClient(t, mockResp)
	req := &types.TradeCreate{
		OutTradeNo:  "create20260420001",
		Subject:     "测试建单",
		TotalAmount: "88.00",
		BuyerId:     "2088102146225135",
	}
	resp, err := client.AlipayTradeCreate(context.Background(), req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.Response.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.Response.Code)
	}
	if resp.Response.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.Response.Msg)
	}
	if resp.Response.TradeNo != "2026042022001480551404403001" {
		t.Errorf("TradeNo = %q unexpected", resp.Response.TradeNo)
	}
}
