package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayFundTransUniTransfer(t *testing.T) {
	mockResp := `{"alipay_fund_trans_uni_transfer_response":{"code":"10000","msg":"Success","out_biz_no":"biz001","order_id":"20260420110070000006880000000000","pay_fund_order_id":"20260420110070000006880000000001","status":"SUCCESS","trans_date":"2026-04-20 10:00:00"},"sign":"fakesign"}`

	client, _ := newTestClient(t, mockResp)

	req := &types.FundTransUniTransfer{
		OutBizNo:    "biz001",
		TransAmount: "10.00",
		ProductCode: "TRANS_ACCOUNT_NO_PWD",
		PayeeInfo: &types.FundTransPayee{
			Identity:     "2088000000000000",
			IdentityType: "ALIPAY_USER_ID",
			Name:         "测试用户",
		},
	}
	resp, err := client.AlipayFundTransUniTransfer(req)
	if err != nil {
		t.Fatalf("AlipayFundTransUniTransfer error: %v", err)
	}
	if resp.AlipayFundTransUniTransferResponse.Status != "SUCCESS" {
		t.Errorf("Status = %q, want SUCCESS", resp.AlipayFundTransUniTransferResponse.Status)
	}
	if resp.AlipayFundTransUniTransferResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayFundTransUniTransferResponse.Code)
	}
	if resp.AlipayFundTransUniTransferResponse.OutBizNo != "biz001" {
		t.Errorf("OutBizNo = %q, want biz001", resp.AlipayFundTransUniTransferResponse.OutBizNo)
	}
	if resp.AlipayFundTransUniTransferResponse.OrderId != "20260420110070000006880000000000" {
		t.Errorf("OrderId = %q, want 20260420110070000006880000000000", resp.AlipayFundTransUniTransferResponse.OrderId)
	}
}
