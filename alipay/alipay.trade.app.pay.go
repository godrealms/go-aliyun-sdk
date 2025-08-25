package alipay

import (
	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"github.com/godrealms/go-aliyun-sdk/community"
	"time"
)

func (c *Client) AlipayTradeAppPay(form *types.TradePay) (string, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.trade.app.pay",
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
		return "", err
	}

	data.Sign, err = signature.GenerateSignature(data)
	if err != nil {
		return "", err
	}

	return data.ToUrlValue().Encode(), nil
}
