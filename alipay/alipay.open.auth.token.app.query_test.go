package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayOpenAuthTokenAppQuery(t *testing.T) {
	mockResp := `{"alipay_open_auth_token_app_query_response":{"code":"10000","msg":"Success","user_id":"2088xxx","auth_app_id":"2021xxx","app_auth_token":"existtoken","app_refresh_token":"refreshtoken","expires_in":31536000,"re_expires_in":32140800,"token_begin_time":"2026-04-20 10:00:00","status":"NORMAL"},"sign":"fakesign"}`

	client, _ := newTestClient(t, mockResp)

	req := &types.OpenAuthTokenAppQuery{
		AppAuthToken: "existtoken",
	}
	resp, err := client.AlipayOpenAuthTokenAppQuery(context.Background(), req)
	if err != nil {
		t.Fatalf("AlipayOpenAuthTokenAppQuery error: %v", err)
	}
	if resp.AlipayOpenAuthTokenAppQueryResponse.Status != "NORMAL" {
		t.Errorf("Status = %q, want NORMAL", resp.AlipayOpenAuthTokenAppQueryResponse.Status)
	}
}
