package alipay

import (
	"context"
	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"time"
)

func (c *Client) AlipayTradeFastPayRefundQuery(ctx context.Context, query *types.TradeFastPayRefundQuery) (*types.AlipayTradeFastPayRefundQueryResponse, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.trade.fastpay.refund.query",
		Format:       "JSON",
		ReturnUrl:    c.ReturnUrl,
		Charset:      "UTF-8",
		SignType:     "RSA2",
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Version:      "1.0",
		NotifyUrl:    c.NotifyUrl,
		AppAuthToken: c.AppAuthToken,
		BizContent:   query.ToString(),
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
	result := &types.AlipayTradeFastPayRefundQueryResponse{}
	err = c.Http.Get(ctx, "", value, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
