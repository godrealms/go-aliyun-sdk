package alipay

import (
	"context"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayUserInfoShare(t *testing.T) {
	mockResp := `{"alipay_user_info_share_response":{"code":"10000","msg":"Success","user_id":"2088000000000001","nick_name":"测试用户","avatar":"https://tfs.alipayobjects.com/images/partner/avatar.jpg","province":"浙江省","city":"杭州市","gender":"M","user_status":"T","is_certified":"T"},"sign":"fakesign"}`
	client, _ := newTestClient(t, mockResp)

	resp, err := client.AlipayUserInfoShare(context.Background(), &types.UserInfoShare{AuthToken: "test_user_access_token"})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.AlipayUserInfoShareResponse.Code != "10000" {
		t.Errorf("Code = %q, want 10000", resp.AlipayUserInfoShareResponse.Code)
	}
	if resp.AlipayUserInfoShareResponse.UserId != "2088000000000001" {
		t.Errorf("UserId = %q, want 2088000000000001", resp.AlipayUserInfoShareResponse.UserId)
	}
	if resp.AlipayUserInfoShareResponse.NickName != "测试用户" {
		t.Errorf("NickName = %q, want 测试用户", resp.AlipayUserInfoShareResponse.NickName)
	}
}
