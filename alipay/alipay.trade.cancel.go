package alipay

import (
	"context"
	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"time"
)

func (c *Client) AlipayTradeCancel(ctx context.Context, cancel *types.TradeCancel) (*types.AlipayTradeCancelResponse, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.trade.cancel",
		Format:       "JSON",
		ReturnUrl:    c.ReturnUrl,
		Charset:      "UTF-8",
		SignType:     "RSA2",
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Version:      "1.0",
		NotifyUrl:    c.NotifyUrl,
		AppAuthToken: c.AppAuthToken,
		BizContent:   cancel.ToString(),
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
	result := &types.AlipayTradeCancelResponse{}
	err = c.Http.PostForm(ctx, "", value, nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil

}
