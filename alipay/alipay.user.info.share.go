package alipay

import (
	"context"
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"github.com/godrealms/go-aliyun-sdk/community"
)

// AlipayUserInfoShare 获取用户基本信息（alipay.user.info.share）
// req.AuthToken 为用户级 OAuth access_token（通过 AlipaySystemOauthToken 换取），
// 该接口以用户身份调用，不附带 app_auth_token。
func (c *Client) AlipayUserInfoShare(req *types.UserInfoShare) (*types.AlipayUserInfoShareResponse, error) {
	data := types.PublicRequestParameters{
		AppId:      c.AppId,
		Method:     "alipay.user.info.share",
		Format:     "JSON",
		Charset:    "UTF-8",
		SignType:   "RSA2",
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Version:    "1.0",
		AuthToken:  req.AuthToken,
		BizContent: req.ToString(),
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
	result := &types.AlipayUserInfoShareResponse{}
	err = c.Http.PostForm(context.Background(), "", value, nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
