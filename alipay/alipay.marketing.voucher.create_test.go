package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayMarketingVoucherCreate(t *testing.T) {
	mockResp := `{"alipay_marketing_voucher_create_response":{"code":"10000","msg":"Success","voucher_id":"VCH20260420001"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	resp, err := client.AlipayMarketingVoucherCreate(context.Background(), &types.VoucherCreate{
		VoucherName: "满100减10", VoucherType: "MONEY_VOUCHER",
		DenominationMoney: "10.00", MerchantId: "merchant001",
		Quantity: "100", ValidBeginTime: "2026-04-20 00:00:00", ValidEndTime: "2026-05-20 23:59:59",
	})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayMarketingVoucherCreateResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayMarketingVoucherCreateResponse.Code)
	}
	if resp.AlipayMarketingVoucherCreateResponse.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.AlipayMarketingVoucherCreateResponse.Msg)
	}
	if resp.AlipayMarketingVoucherCreateResponse.VoucherId != "VCH20260420001" {
		t.Errorf("VoucherId = %q, want VCH20260420001", resp.AlipayMarketingVoucherCreateResponse.VoucherId)
	}
}
