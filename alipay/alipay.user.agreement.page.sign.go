package alipay

import (
	"fmt"
	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"github.com/godrealms/go-aliyun-sdk/community"
	"time"
)

// GetAlipayUserAgreementPageSign 支付宝个人协议页面签约接口(跳转地址)
func (c *Client) GetAlipayUserAgreementPageSign(page *types.AgreementPageSign) (string, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.user.agreement.page.sign",
		Format:       "JSON",
		ReturnUrl:    c.ReturnUrl,
		Charset:      "UTF-8",
		SignType:     "RSA2",
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Version:      "1.0",
		NotifyUrl:    c.NotifyUrl,
		AppAuthToken: c.AppAuthToken,
		BizContent:   page.ToString(),
	}
	signature, err := community.NewSignatureHelper(c.PrivateKey)
	if err != nil {
		return "", err
	}

	data.Sign, err = signature.GenerateSignature(data)
	if err != nil {
		return "", err
	}

	urlValue := data.ToUrlValue()
	return fmt.Sprintf("%s?%s", c.Http.BaseURL, urlValue.Encode()), nil
}

// PostAlipayUserAgreementPageSign 支付宝个人协议页面签约接口(拼接参数)
func (c *Client) PostAlipayUserAgreementPageSign(page *types.AgreementPageSign) (*types.PublicRequestParameters, error) {
	data := &types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.user.agreement.page.sign",
		Format:       "JSON",
		ReturnUrl:    c.ReturnUrl,
		Charset:      "UTF-8",
		SignType:     "RSA2",
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Version:      "1.0",
		NotifyUrl:    c.NotifyUrl,
		AppAuthToken: c.AppAuthToken,
		BizContent:   page.ToString(),
	}
	signature, err := community.NewSignatureHelper(c.PrivateKey)
	if err != nil {
		return nil, err
	}

	data.Sign, err = signature.GenerateSignature(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
