package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayFundTransRefund(t *testing.T) {
	mockResp := `{"alipay_fund_trans_refund_response":{"code":"10000","msg":"Success","order_id":"20260420110070000006880000000000","refund_order_id":"20260420110070000006880000000002","refund_amount":"10.00","refund_date":"2026-04-20 10:05:00"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	req := &types.FundTransRefund{
		OutRequestNo: "refund001",
		OrderId:      "20260420110070000006880000000000",
		RefundAmount: "10.00",
	}
	resp, err := client.AlipayFundTransRefund(context.Background(), req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayFundTransRefundResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayFundTransRefundResponse.Code)
	}
	if resp.AlipayFundTransRefundResponse.RefundAmount != "10.00" {
		t.Errorf("RefundAmount = %q, want 10.00", resp.AlipayFundTransRefundResponse.RefundAmount)
	}
	if resp.AlipayFundTransRefundResponse.RefundOrderId != "20260420110070000006880000000002" {
		t.Errorf("RefundOrderId = %q unexpected", resp.AlipayFundTransRefundResponse.RefundOrderId)
	}
}
