package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayUserCertifyOpenQuery(t *testing.T) {
	mockResp := `{"alipay_user_certify_open_query_response":{"code":"10000","msg":"Success","passed":"T","identity_info":"{\"name\":\"张三\"}"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)
	req := &types.UserCertifyOpenQuery{CertifyId: "OcCp2413fkv09diXXXXX"}
	resp, err := client.AlipayUserCertifyOpenQuery(context.Background(), req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayUserCertifyOpenQueryResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayUserCertifyOpenQueryResponse.Code)
	}
	if resp.AlipayUserCertifyOpenQueryResponse.Passed != "T" {
		t.Errorf("Passed = %q, want T", resp.AlipayUserCertifyOpenQueryResponse.Passed)
	}
}
