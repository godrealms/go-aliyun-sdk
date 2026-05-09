package alipay

import (
	"context"
	"fmt"
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func (c *Client) AlipayTradeRefund(ctx context.Context, request *types.TradeRefund) (*types.AlipayTradeRefundResponse, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.trade.refund",
		Format:       "JSON",
		ReturnUrl:    c.ReturnUrl,
		Charset:      "UTF-8",
		SignType:     "RSA2",
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Version:      "1.0",
		NotifyUrl:    c.NotifyUrl,
		AppAuthToken: c.AppAuthToken,
		BizContent:   request.ToString(),
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
	result := &types.AlipayTradeRefundResponse{}
	err = c.Http.PostForm(ctx, "", value, nil, result)
	if err != nil {
		return nil, err
	}
	if result.Response.Code != "10000" {
		return nil, fmt.Errorf("alipay trade refund failed: %s", result.Response.Msg)
	}
	return result, nil
}
