package alipay

import (
	"context"
	"strings"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayUserCertifyOpenCertify(t *testing.T) {
	client, ts := newTestClient(t, "")
	req := &types.UserCertifyOpenCertify{CertifyId: "OcCp2413fkv09diXXXXX"}
	u, err := client.AlipayUserCertifyOpenCertify(context.Background(), req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if !strings.HasPrefix(u, ts.URL) {
		t.Errorf("URL should start with %s, got: %s", ts.URL, u)
	}
	if !strings.Contains(u, "method=alipay.user.certify.open.certify") {
		t.Errorf("URL missing method param, got: %s", u)
	}
	if !strings.Contains(u, "sign=") {
		t.Errorf("URL missing sign param, got: %s", u)
	}
}
