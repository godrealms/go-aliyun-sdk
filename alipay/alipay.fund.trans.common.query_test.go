package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayFundTransCommonQuery(t *testing.T) {
	mockResp := `{"alipay_fund_trans_common_query_response":{"code":"10000","msg":"Success","order_id":"20260420110070000006880000000000","pay_fund_order_id":"20260420110070000006880000000001","status":"SUCCESS","pay_date":"2026-04-20 10:00:00","order_fee":"10.00"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	req := &types.FundTransCommonQuery{OutBizNo: "biz001"}
	resp, err := client.AlipayFundTransCommonQuery(req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayFundTransCommonQueryResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayFundTransCommonQueryResponse.Code)
	}
	if resp.AlipayFundTransCommonQueryResponse.Status != "SUCCESS" {
		t.Errorf("Status = %q, want SUCCESS", resp.AlipayFundTransCommonQueryResponse.Status)
	}
	if resp.AlipayFundTransCommonQueryResponse.OrderId != "20260420110070000006880000000000" {
		t.Errorf("OrderId = %q unexpected", resp.AlipayFundTransCommonQueryResponse.OrderId)
	}
}
