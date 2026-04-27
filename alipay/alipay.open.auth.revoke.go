package alipay

import (
	"context"
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

// AlipayOpenAuthRevoke 解除应用授权（alipay.open.auth.revoke）
func (c *Client) AlipayOpenAuthRevoke(ctx context.Context, request *types.OpenAuthRevoke) (*types.AlipayOpenAuthRevokeResponse, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.open.auth.revoke",
		Format:       "JSON",
		Charset:      "UTF-8",
		SignType:     "RSA2",
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Version:      "1.0",
		AppAuthToken: c.AppAuthToken,
		BizContent:   request.ToString(),
	}

	signer, err := c.getSigner()
	if err != nil {
		return nil, err
	}
	data.Sign, err = signer.GenerateSignature(data)
	if err != nil {
		return nil, err
	}

	value := data.ToUrlValue()
	result := &types.AlipayOpenAuthRevokeResponse{}
	err = c.Http.PostForm(ctx, "", value, nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
