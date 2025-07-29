package alipay

import (
	"context"
	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"github.com/godrealms/go-aliyun-sdk/community"
	"net/url"
	"time"
)

func (c *Client) AlipayTradeRefund(request *types.TradeRefund) (*types.AlipayTradeRefundResponse, error) {
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

	signature, err := community.NewSignatureHelper(c.PrivateKey)
	if err != nil {
		return nil, err
	}
	data.Sign, err = signature.GenerateSignature(data)
	if err != nil {
		return nil, err
	}
	value := data.ToUrlValue()
	result := &types.AlipayTradeRefundResponse{}
	query := url.Values{
		"charset":   []string{"UTF-8"},
		"method":    []string{"alipay.trade.refund"},
		"format":    []string{"JSON"},
		"sign":      []string{data.Sign},
		"app_id":    []string{c.AppId},
		"version":   []string{"1.0"},
		"sign_type": []string{"RSA2"},
		"timestamp": []string{time.Now().Format("2006-01-02 15:04:05")},
	}
	err = c.Http.PostForm(context.Background(), "", value, query, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
