package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayFundJointFrozen(t *testing.T) {
	mockResp := `{"alipay_fund_joint_frozen_response":{"code":"10000","msg":"Success","freeze_id":"2026042022001480551404403022","freeze_date":"2026-04-20 10:00:00","amount":"50.00"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	req := &types.FundJointFrozen{
		OutRequestNo: "freeze001",
		PayerUserId:  "2088000000000000",
		Amount:       "50.00",
		Remark:       "押金冻结",
	}
	resp, err := client.AlipayFundJointFrozen(context.Background(), req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayFundJointFrozenResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayFundJointFrozenResponse.Code)
	}
	if resp.AlipayFundJointFrozenResponse.FreezeId != "2026042022001480551404403022" {
		t.Errorf("FreezeId = %q unexpected", resp.AlipayFundJointFrozenResponse.FreezeId)
	}
	if resp.AlipayFundJointFrozenResponse.Amount != "50.00" {
		t.Errorf("Amount = %q, want 50.00", resp.AlipayFundJointFrozenResponse.Amount)
	}
}
