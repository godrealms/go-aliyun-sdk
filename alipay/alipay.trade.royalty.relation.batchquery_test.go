package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayTradeRoyaltyRelationBatchquery(t *testing.T) {
	mockResp := `{"alipay_trade_royalty_relation_batchquery_response":{"code":"10000","msg":"Success","count":2,"receiver_infos":[{"receiver_type":"loginName","receiver_account":"a@example.com","receiver_name":"甲方","relation_type":"PARTNER"},{"receiver_type":"loginName","receiver_account":"b@example.com","receiver_name":"乙方","relation_type":"PARTNER"}]}}`
	client, _ := newTestClient(t, mockResp)
	req := &types.TradeRoyaltyRelationBatchquery{
		OutRequestNo: "bq20260420001",
	}
	resp, err := client.AlipayTradeRoyaltyRelationBatchquery(req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.Response.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.Response.Code)
	}
	if resp.Response.Msg != "Success" {
		t.Errorf("Msg = %q, want Success", resp.Response.Msg)
	}
	if resp.Response.Count != 2 {
		t.Errorf("Count = %d, want 2", resp.Response.Count)
	}
	if len(resp.Response.ReceiverInfos) != 2 {
		t.Fatalf("ReceiverInfos len = %d, want 2", len(resp.Response.ReceiverInfos))
	}
	if resp.Response.ReceiverInfos[0].ReceiverAccount != "a@example.com" {
		t.Errorf("ReceiverInfos[0].ReceiverAccount = %q unexpected", resp.Response.ReceiverInfos[0].ReceiverAccount)
	}
}
