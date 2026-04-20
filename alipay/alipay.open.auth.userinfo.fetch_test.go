package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayOpenAuthUserinfoFetch(t *testing.T) {
	mockResp := `{"alipay_open_auth_userinfo_fetch_response":{"code":"10000","msg":"Success","user_id":"2088xxx","name":"测试商家","alias":"测试","login_id":"test@example.com"},"sign":"fakesign"}`

	client, _ := newTestClient(t, mockResp)

	req := &types.OpenAuthUserinfoFetch{
		AppAuthToken: "test_app_auth_token",
	}
	resp, err := client.AlipayOpenAuthUserinfoFetch(req)
	if err != nil {
		t.Fatalf("AlipayOpenAuthUserinfoFetch error: %v", err)
	}
	if resp.AlipayOpenAuthUserinfoFetchResponse.UserId != "2088xxx" {
		t.Errorf("UserId = %q, want 2088xxx", resp.AlipayOpenAuthUserinfoFetchResponse.UserId)
	}
}
