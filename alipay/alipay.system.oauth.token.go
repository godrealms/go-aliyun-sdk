package alipay

import (
	"context"
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"github.com/godrealms/go-aliyun-sdk/community"
)

// AlipaySystemOauthToken 换取/刷新用户授权访问令牌（alipay.system.oauth.token）
func (c *Client) AlipaySystemOauthToken(request *types.SystemOauthToken) (*types.AlipaySystemOauthTokenResponse, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.system.oauth.token",
		Format:       "JSON",
		Charset:      "UTF-8",
		SignType:     "RSA2",
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Version:      "1.0",
		AppAuthToken: c.AppAuthToken,
		BizContent:   request.ToString(),
	}

	signature, err := community.NewSignatureHelper(c.PrivateKey)
	if err != nil {
		return nil, err
	}
	data.Sign, err = signature.GenerateSignature(data)
	if err != nil {
		return nil, err
	}

	value := data.ToUrlValue()
	result := &types.AlipaySystemOauthTokenResponse{}
	err = c.Http.PostForm(context.Background(), "", value, nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
