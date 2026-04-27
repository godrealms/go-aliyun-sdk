package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayTradeRoyaltyRelationUnbind(t *testing.T) {
	mockResp := `{"alipay_trade_royalty_relation_unbind_response":{"code":"10000","msg":"Success"}}`
	client, _ := newTestClient(t, mockResp)
	req := &types.TradeRoyaltyRelationUnbind{
		OutRequestNo: "unbind20260420001",
		ReceiverList: `[{"type":"loginName","account":"receiver@example.com","name":"收款方"}]`,
	}
	resp, err := client.AlipayTradeRoyaltyRelationUnbind(context.Background(), req)
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
