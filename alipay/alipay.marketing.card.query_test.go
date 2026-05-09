package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayMarketingCardQuery(t *testing.T) {
	mockResp := `{"alipay_marketing_card_query_response":{"code":"10000","msg":"Success","card_id":"CARD20260420001","template_id":"TPL20260420001","status":"OPENED","balance":"100.00","point":"500"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	resp, err := client.AlipayMarketingCardQuery(context.Background(), &types.CardQuery{CardId: "CARD20260420001"})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayMarketingCardQueryResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayMarketingCardQueryResponse.Code)
	}
	if resp.AlipayMarketingCardQueryResponse.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.AlipayMarketingCardQueryResponse.Msg)
	}
	if resp.AlipayMarketingCardQueryResponse.CardId != "CARD20260420001" {
		t.Errorf("CardId = %q, want CARD20260420001", resp.AlipayMarketingCardQueryResponse.CardId)
	}
	if resp.AlipayMarketingCardQueryResponse.Status != "OPENED" {
		t.Errorf("Status = %q, want OPENED", resp.AlipayMarketingCardQueryResponse.Status)
	}
}
