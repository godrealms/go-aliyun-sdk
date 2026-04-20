package alipay

import (
	"context"
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"github.com/godrealms/go-aliyun-sdk/community"
)

// AlipayTradeCreate 统一收单下单（alipay.trade.create）
func (c *Client) AlipayTradeCreate(req *types.TradeCreate) (*types.AlipayTradeCreateResponse, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.trade.create",
		Format:       "JSON",
		Charset:      "UTF-8",
		SignType:     "RSA2",
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Version:      "1.0",
		NotifyUrl:    c.NotifyUrl,
		AppAuthToken: c.AppAuthToken,
		BizContent:   req.ToString(),
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
	result := &types.AlipayTradeCreateResponse{}
	err = c.Http.Get(context.Background(), "", value, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
