package alipay

import (
	"context"
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

// AlipayTradeRoyaltyRelationBatchquery 分账关系批量查询（alipay.trade.royalty.relation.batchquery）
func (c *Client) AlipayTradeRoyaltyRelationBatchquery(ctx context.Context, req *types.TradeRoyaltyRelationBatchquery) (*types.AlipayTradeRoyaltyRelationBatchqueryResponse, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.trade.royalty.relation.batchquery",
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
	result := &types.AlipayTradeRoyaltyRelationBatchqueryResponse{}
	err = c.Http.Get(ctx, "", value, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
