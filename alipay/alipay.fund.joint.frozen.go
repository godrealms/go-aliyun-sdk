package alipay

import (
	"context"
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"github.com/godrealms/go-aliyun-sdk/community"
)

// AlipayFundJointFrozen 资金冻结（alipay.fund.joint.frozen）
func (c *Client) AlipayFundJointFrozen(request *types.FundJointFrozen) (*types.AlipayFundJointFrozenResponse, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.fund.joint.frozen",
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
	result := &types.AlipayFundJointFrozenResponse{}
	err = c.Http.PostForm(context.Background(), "", value, nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
