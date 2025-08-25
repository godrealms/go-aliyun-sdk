package alipay

import (
	"context"
	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"github.com/godrealms/go-aliyun-sdk/community"
	"time"
)

// AlipayUserAgreementUnSign 支付宝个人代扣协议解约接口
func (c *Client) AlipayUserAgreementUnSign(form *types.AgreementUnSign) (*types.AgreementUnsignResponse, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.user.agreement.unsign",
		Format:       "JSON",
		ReturnUrl:    c.ReturnUrl,
		Charset:      "UTF-8",
		SignType:     "RSA2",
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Version:      "1.0",
		NotifyUrl:    c.NotifyUrl,
		AppAuthToken: c.AppAuthToken,
		BizContent:   form.ToString(),
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
	result := &types.AgreementUnsignResponse{}
	err = c.Http.Get(context.Background(), "", value, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
