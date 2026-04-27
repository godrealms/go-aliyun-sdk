package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayMarketingCardOpen(t *testing.T) {
	mockResp := `{"alipay_marketing_card_open_response":{"code":"10000","msg":"Success","open_url":"https://render.alipay.com/p/s/xxx"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	resp, err := client.AlipayMarketingCardOpen(context.Background(), &types.CardOpen{RequestId: "req20260420001", TemplateId: "TPL20260420001"})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayMarketingCardOpenResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayMarketingCardOpenResponse.Code)
	}
	if resp.AlipayMarketingCardOpenResponse.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.AlipayMarketingCardOpenResponse.Msg)
	}
	if resp.AlipayMarketingCardOpenResponse.OpenUrl == "" {
		t.Error("OpenUrl is empty")
	}
}
