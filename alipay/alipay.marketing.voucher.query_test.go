package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayMarketingVoucherQuery(t *testing.T) {
	mockResp := `{"alipay_marketing_voucher_query_response":{"code":"10000","msg":"Success","voucher_id":"VCH20260420001","voucher_name":"满100减10","voucher_type":"MONEY_VOUCHER","status":"VALID","valid_begin_time":"2026-04-20 00:00:00","valid_end_time":"2026-05-20 23:59:59"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	resp, err := client.AlipayMarketingVoucherQuery(&types.VoucherQuery{VoucherId: "VCH20260420001"})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayMarketingVoucherQueryResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayMarketingVoucherQueryResponse.Code)
	}
	if resp.AlipayMarketingVoucherQueryResponse.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.AlipayMarketingVoucherQueryResponse.Msg)
	}
	if resp.AlipayMarketingVoucherQueryResponse.VoucherId != "VCH20260420001" {
		t.Errorf("VoucherId = %q, want VCH20260420001", resp.AlipayMarketingVoucherQueryResponse.VoucherId)
	}
	if resp.AlipayMarketingVoucherQueryResponse.Status != "VALID" {
		t.Errorf("Status = %q, want VALID", resp.AlipayMarketingVoucherQueryResponse.Status)
	}
}
