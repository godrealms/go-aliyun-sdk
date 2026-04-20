package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayMarketingCardTemplateQuery(t *testing.T) {
	mockResp := `{"alipay_marketing_card_template_query_response":{"code":"10000","msg":"Success","template_id":"TPL20260420001","template_name":"黄金会员卡","status":"NORMAL"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	resp, err := client.AlipayMarketingCardTemplateQuery(&types.CardTemplateQuery{TemplateId: "TPL20260420001"})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayMarketingCardTemplateQueryResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayMarketingCardTemplateQueryResponse.Code)
	}
	if resp.AlipayMarketingCardTemplateQueryResponse.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.AlipayMarketingCardTemplateQueryResponse.Msg)
	}
	if resp.AlipayMarketingCardTemplateQueryResponse.TemplateId != "TPL20260420001" {
		t.Errorf("TemplateId = %q, want TPL20260420001", resp.AlipayMarketingCardTemplateQueryResponse.TemplateId)
	}
	if resp.AlipayMarketingCardTemplateQueryResponse.Status != "NORMAL" {
		t.Errorf("Status = %q, want NORMAL", resp.AlipayMarketingCardTemplateQueryResponse.Status)
	}
}
