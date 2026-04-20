package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayFundJointThaw(t *testing.T) {
	mockResp := `{"alipay_fund_joint_thaw_response":{"code":"10000","msg":"Success","out_request_no":"thaw001","amount":"50.00","thaw_date":"2026-04-20 11:00:00"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	req := &types.FundJointThaw{
		OutRequestNo: "thaw001",
		FreezeId:     "2026042022001480551404403022",
		Amount:       "50.00",
	}
	resp, err := client.AlipayFundJointThaw(req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayFundJointThawResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayFundJointThawResponse.Code)
	}
	if resp.AlipayFundJointThawResponse.Amount != "50.00" {
		t.Errorf("Amount = %q, want 50.00", resp.AlipayFundJointThawResponse.Amount)
	}
	if resp.AlipayFundJointThawResponse.OutRequestNo != "thaw001" {
		t.Errorf("OutRequestNo = %q, want thaw001", resp.AlipayFundJointThawResponse.OutRequestNo)
	}
}
