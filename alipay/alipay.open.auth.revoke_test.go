package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayOpenAuthRevoke(t *testing.T) {
	mockResp := `{"alipay_open_auth_revoke_response":{"code":"10000","msg":"Success"},"sign":"fakesign"}`

	client, _ := newTestClient(t, mockResp)

	req := &types.OpenAuthRevoke{
		AppAuthToken: "tokentorevoke",
	}
	resp, err := client.AlipayOpenAuthRevoke(req)
	if err != nil {
		t.Fatalf("AlipayOpenAuthRevoke error: %v", err)
	}
	if resp.AlipayOpenAuthRevokeResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayOpenAuthRevokeResponse.Code)
	}
}
