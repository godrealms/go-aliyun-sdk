package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayFundTransToalipayTransfer(t *testing.T) {
	mockResp := `{"alipay_fund_trans_toalipay_transfer_response":{"code":"10000","msg":"Success","out_biz_no":"biz002","order_id":"20260420110070000006880000000003","pay_date":"2026-04-20 10:00:00"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	req := &types.FundTransToalipayTransfer{
		OutBizNo:     "biz002",
		PayeeType:    "ALIPAY_LOGONID",
		PayeeAccount: "test@example.com",
		Amount:       "5.00",
	}
	resp, err := client.AlipayFundTransToalipayTransfer(context.Background(), req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayFundTransToalipayTransferResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayFundTransToalipayTransferResponse.Code)
	}
	if resp.AlipayFundTransToalipayTransferResponse.OrderId != "20260420110070000006880000000003" {
		t.Errorf("OrderId = %q unexpected", resp.AlipayFundTransToalipayTransferResponse.OrderId)
	}
	if resp.AlipayFundTransToalipayTransferResponse.OutBizNo != "biz002" {
		t.Errorf("OutBizNo = %q, want biz002", resp.AlipayFundTransToalipayTransferResponse.OutBizNo)
	}
}
