package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayMarketingCardDelete(t *testing.T) {
	mockResp := `{"alipay_marketing_card_delete_response":{"code":"10000","msg":"Success"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	resp, err := client.AlipayMarketingCardDelete(&types.CardDelete{CardId: "CARD20260420001", TemplateId: "TPL20260420001"})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayMarketingCardDeleteResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayMarketingCardDeleteResponse.Code)
	}
	if resp.AlipayMarketingCardDeleteResponse.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.AlipayMarketingCardDeleteResponse.Msg)
	}
}
