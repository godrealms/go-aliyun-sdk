package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayFundTransOrderQuery(t *testing.T) {
	mockResp := `{"alipay_fund_trans_order_query_response":{"code":"10000","msg":"Success","order_id":"20260420110070000006880000000000","status":"SUCCESS","pay_date":"2026-04-20 10:00:00","order_fee":"10.00","fail_reason":""},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	req := &types.FundTransOrderQuery{OutBizNo: "biz001"}
	resp, err := client.AlipayFundTransOrderQuery(req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayFundTransOrderQueryResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayFundTransOrderQueryResponse.Code)
	}
	if resp.AlipayFundTransOrderQueryResponse.Status != "SUCCESS" {
		t.Errorf("Status = %q, want SUCCESS", resp.AlipayFundTransOrderQueryResponse.Status)
	}
	if resp.AlipayFundTransOrderQueryResponse.OrderId != "20260420110070000006880000000000" {
		t.Errorf("OrderId = %q unexpected", resp.AlipayFundTransOrderQueryResponse.OrderId)
	}
}
