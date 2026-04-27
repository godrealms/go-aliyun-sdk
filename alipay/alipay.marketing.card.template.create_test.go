package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayMarketingCardTemplateCreate(t *testing.T) {
	mockResp := `{"alipay_marketing_card_template_create_response":{"code":"10000","msg":"Success","template_id":"TPL20260420001"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	resp, err := client.AlipayMarketingCardTemplateCreate(context.Background(), &types.CardTemplateCreate{
		TemplateName: "黄金会员卡", LogoUrl: "https://example.com/logo.png",
	})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayMarketingCardTemplateCreateResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayMarketingCardTemplateCreateResponse.Code)
	}
	if resp.AlipayMarketingCardTemplateCreateResponse.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.AlipayMarketingCardTemplateCreateResponse.Msg)
	}
	if resp.AlipayMarketingCardTemplateCreateResponse.TemplateId != "TPL20260420001" {
		t.Errorf("TemplateId = %q, want TPL20260420001", resp.AlipayMarketingCardTemplateCreateResponse.TemplateId)
	}
}
