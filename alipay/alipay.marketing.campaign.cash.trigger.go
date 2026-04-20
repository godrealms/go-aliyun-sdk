package alipay

import (
	"context"
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"github.com/godrealms/go-aliyun-sdk/community"
)

// AlipayMarketingCampaignCashTrigger 触发活动奖励发放（alipay.marketing.campaign.cash.trigger）
func (c *Client) AlipayMarketingCampaignCashTrigger(request *types.CampaignCashTrigger) (*types.AlipayMarketingCampaignCashTriggerResponse, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.marketing.campaign.cash.trigger",
		Format:       "JSON",
		Charset:      "UTF-8",
		SignType:     "RSA2",
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Version:      "1.0",
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
	result := &types.AlipayMarketingCampaignCashTriggerResponse{}
	err = c.Http.PostForm(context.Background(), "", data.ToUrlValue(), nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
