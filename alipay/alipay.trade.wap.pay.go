package alipay

import (
	"context"
	"fmt"
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

// AlipayTradeWapPay H5手机网站支付（alipay.trade.wap.pay）
func (c *Client) AlipayTradeWapPay(ctx context.Context, req *types.TradeWapPay) (string, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.trade.wap.pay",
		Format:       "JSON",
		ReturnUrl:    c.ReturnUrl,
		Charset:      "UTF-8",
		SignType:     "RSA2",
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Version:      "1.0",
		NotifyUrl:    c.NotifyUrl,
		AppAuthToken: c.AppAuthToken,
		BizContent:   req.ToString(),
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
	return fmt.Sprintf("%s?%s", c.Http.GetBaseURL(), value.Encode()), nil
}
