package alipay

import (
	"context"
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"github.com/godrealms/go-aliyun-sdk/community"
)

// AlipayUserCertifyOpenQuery 查询实名认证结果（alipay.user.certify.open.query）
func (c *Client) AlipayUserCertifyOpenQuery(request *types.UserCertifyOpenQuery) (*types.AlipayUserCertifyOpenQueryResponse, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.user.certify.open.query",
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
	result := &types.AlipayUserCertifyOpenQueryResponse{}
	err = c.Http.PostForm(context.Background(), "", value, nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
