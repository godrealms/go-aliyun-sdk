package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayMarketingCardUpdate(t *testing.T) {
	mockResp := `{"alipay_marketing_card_update_response":{"code":"10000","msg":"Success"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	resp, err := client.AlipayMarketingCardUpdate(context.Background(), &types.CardUpdate{
		CardId: "CARD20260420001", TemplateId: "TPL20260420001", BalanceInfo: `{"balance":"200.00"}`,
	})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayMarketingCardUpdateResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayMarketingCardUpdateResponse.Code)
	}
	if resp.AlipayMarketingCardUpdateResponse.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.AlipayMarketingCardUpdateResponse.Msg)
	}
}
