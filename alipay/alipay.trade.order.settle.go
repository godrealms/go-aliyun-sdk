package alipay

import (
	"context"
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

// AlipayTradeOrderSettle 统一收单交易结算（alipay.trade.order.settle）
func (c *Client) AlipayTradeOrderSettle(ctx context.Context, req *types.TradeOrderSettle) (*types.AlipayTradeOrderSettleResponse, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.trade.order.settle",
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
		return nil, err
	}
	data.Sign, err = signer.GenerateSignature(data)
	if err != nil {
		return nil, err
	}
	value := data.ToUrlValue()
	result := &types.AlipayTradeOrderSettleResponse{}
	err = c.Http.PostForm(ctx, "", value, nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
