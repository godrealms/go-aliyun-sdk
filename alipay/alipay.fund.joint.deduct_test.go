package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayFundJointDeduct(t *testing.T) {
	mockResp := `{"alipay_fund_joint_deduct_response":{"code":"10000","msg":"Success","out_request_no":"deduct001","order_id":"20260420110070000006880000000005","amount":"30.00","deduct_date":"2026-04-20 12:00:00"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	req := &types.FundJointDeduct{
		OutRequestNo: "deduct001",
		FreezeId:     "2026042022001480551404403022",
		Amount:       "30.00",
		Remark:       "扣除押金",
	}
	resp, err := client.AlipayFundJointDeduct(context.Background(), req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayFundJointDeductResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayFundJointDeductResponse.Code)
	}
	if resp.AlipayFundJointDeductResponse.Amount != "30.00" {
		t.Errorf("Amount = %q, want 30.00", resp.AlipayFundJointDeductResponse.Amount)
	}
	if resp.AlipayFundJointDeductResponse.OrderId != "20260420110070000006880000000005" {
		t.Errorf("OrderId = %q unexpected", resp.AlipayFundJointDeductResponse.OrderId)
	}
}
