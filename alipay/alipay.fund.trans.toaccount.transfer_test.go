package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayFundTransToaccountTransfer(t *testing.T) {
	mockResp := `{"alipay_fund_trans_toaccount_transfer_response":{"code":"10000","msg":"Success","out_biz_no":"biz003","order_id":"20260420110070000006880000000004","pay_date":"2026-04-20 10:00:00"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	req := &types.FundTransToaccountTransfer{
		OutBizNo:     "biz003",
		PayeeType:    "ALIPAY_LOGONID",
		PayeeAccount: "test@example.com",
		Amount:       "5.00",
	}
	resp, err := client.AlipayFundTransToaccountTransfer(context.Background(), req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayFundTransToaccountTransferResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayFundTransToaccountTransferResponse.Code)
	}
	if resp.AlipayFundTransToaccountTransferResponse.OrderId != "20260420110070000006880000000004" {
		t.Errorf("OrderId = %q unexpected", resp.AlipayFundTransToaccountTransferResponse.OrderId)
	}
	if resp.AlipayFundTransToaccountTransferResponse.OutBizNo != "biz003" {
		t.Errorf("OutBizNo = %q, want biz003", resp.AlipayFundTransToaccountTransferResponse.OutBizNo)
	}
}
