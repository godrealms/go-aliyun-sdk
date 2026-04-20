package alipay

import (
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayUserCertifyOpenInitialize(t *testing.T) {
	mockResp := `{"alipay_user_certify_open_initialize_response":{"code":"10000","msg":"Success","certify_id":"OcCp2413fkv09diXXXXX"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	req := &types.UserCertifyOpenInitialize{
		OuterOrderNo:  "order_certify_001",
		BizCode:       "FACE",
		IdentityParam: `{"identity_type":"CERT_INFO","cert_type":"IDENTITY_CARD","cert_name":"张三","cert_no":"310000199001011234"}`,
	}
	resp, err := client.AlipayUserCertifyOpenInitialize(req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayUserCertifyOpenInitializeResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayUserCertifyOpenInitializeResponse.Code)
	}
	if resp.AlipayUserCertifyOpenInitializeResponse.CertifyId != "OcCp2413fkv09diXXXXX" {
		t.Errorf("CertifyId = %q, want OcCp2413fkv09diXXXXX", resp.AlipayUserCertifyOpenInitializeResponse.CertifyId)
	}
}
