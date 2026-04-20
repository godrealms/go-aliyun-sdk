package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayOpenAuthTokenApp(t *testing.T) {
	mockResp := `{"alipay_open_auth_token_app_response":{"code":"10000","msg":"Success","user_id":"2088xxx","auth_app_id":"2021xxx","app_auth_token":"newtoken","app_refresh_token":"refreshtoken","expires_in":31536000,"re_expires_in":32140800,"token_begin_time":"2026-04-20 10:00:00"},"sign":"fakesign"}`

	client, _ := newTestClient(t, mockResp)

	req := &types.OpenAuthTokenApp{
		GrantType: "authorization_code",
		Code:      "testcode",
	}
	resp, err := client.AlipayOpenAuthTokenApp(req)
	if err != nil {
		t.Fatalf("AlipayOpenAuthTokenApp error: %v", err)
	}
	if resp.AlipayOpenAuthTokenAppResponse.AppAuthToken != "newtoken" {
		t.Errorf("AppAuthToken = %q, want %q", resp.AlipayOpenAuthTokenAppResponse.AppAuthToken, "newtoken")
	}
	if resp.AlipayOpenAuthTokenAppResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayOpenAuthTokenAppResponse.Code)
	}
}
