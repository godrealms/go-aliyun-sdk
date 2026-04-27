package alipay

import (
	"context"
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

// AlipayMarketingCardUpdate 更新会员卡（alipay.marketing.card.update）
func (c *Client) AlipayMarketingCardUpdate(ctx context.Context, request *types.CardUpdate) (*types.AlipayMarketingCardUpdateResponse, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.marketing.card.update",
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
	result := &types.AlipayMarketingCardUpdateResponse{}
	err = c.Http.PostForm(ctx, "", value, nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
