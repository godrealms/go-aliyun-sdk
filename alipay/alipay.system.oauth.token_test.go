package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipaySystemOauthToken(t *testing.T) {
	mockResp := `{"alipay_system_oauth_token_response":{"code":"10000","msg":"Success","user_id":"2088000000000000","access_token":"test_access_token_001","expires_in":"3600","refresh_token":"test_refresh_token_001","re_expires_in":"2592000"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	req := &types.SystemOauthToken{
		GrantType: "authorization_code",
		Code:      "test_auth_code",
	}
	resp, err := client.AlipaySystemOauthToken(req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipaySystemOauthTokenResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipaySystemOauthTokenResponse.Code)
	}
	if resp.AlipaySystemOauthTokenResponse.UserId != "2088000000000000" {
		t.Errorf("UserId = %q, want 2088000000000000", resp.AlipaySystemOauthTokenResponse.UserId)
	}
	if resp.AlipaySystemOauthTokenResponse.AccessToken != "test_access_token_001" {
		t.Errorf("AccessToken = %q, want test_access_token_001", resp.AlipaySystemOauthTokenResponse.AccessToken)
	}
}
