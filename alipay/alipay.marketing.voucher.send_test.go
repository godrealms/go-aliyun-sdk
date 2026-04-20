package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayMarketingVoucherSend(t *testing.T) {
	mockResp := `{"alipay_marketing_voucher_send_response":{"code":"10000","msg":"Success","detail_id":"DTL20260420001"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	resp, err := client.AlipayMarketingVoucherSend(&types.VoucherSend{
		VoucherId: "VCH20260420001", OpenId: "0680809090909090909090909090", OutBizNo: "send20260420001",
	})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayMarketingVoucherSendResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayMarketingVoucherSendResponse.Code)
	}
	if resp.AlipayMarketingVoucherSendResponse.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.AlipayMarketingVoucherSendResponse.Msg)
	}
	if resp.AlipayMarketingVoucherSendResponse.DetailId != "DTL20260420001" {
		t.Errorf("DetailId = %q, want DTL20260420001", resp.AlipayMarketingVoucherSendResponse.DetailId)
	}
}
