package alipay

import (
	"context"
	"fmt"
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func (c *Client) AlipayTradePagePay(ctx context.Context, form *types.TradePay) (string, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.trade.page.pay",
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

	signer, err := c.getSigner()
	if err != nil {
		return "", err
	}
	data.Sign, err = signer.GenerateSignature(data)
	if err != nil {
		return "", err
	}
	value := data.ToUrlValue()

	fullURL := fmt.Sprintf("%s?%s", c.Http.GetBaseURL(), value.Encode())
	return fullURL, nil
}
