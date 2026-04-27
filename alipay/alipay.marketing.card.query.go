package alipay

import (
	"context"
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

// AlipayMarketingCardQuery 查询会员卡（alipay.marketing.card.query）
func (c *Client) AlipayMarketingCardQuery(ctx context.Context, request *types.CardQuery) (*types.AlipayMarketingCardQueryResponse, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.marketing.card.query",
		Format:       "JSON",
		Charset:      "UTF-8",
		SignType:     "RSA2",
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Version:      "1.0",
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
	result := &types.AlipayMarketingCardQueryResponse{}
	err = c.Http.PostForm(ctx, "", value, nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
