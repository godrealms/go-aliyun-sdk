package alipay

import (
	"context"
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"github.com/godrealms/go-aliyun-sdk/community"
)

// AlipayMarketingVoucherCreate 创建优惠券（alipay.marketing.voucher.create）
func (c *Client) AlipayMarketingVoucherCreate(request *types.VoucherCreate) (*types.AlipayMarketingVoucherCreateResponse, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.marketing.voucher.create",
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
	result := &types.AlipayMarketingVoucherCreateResponse{}
	err = c.Http.PostForm(context.Background(), "", data.ToUrlValue(), nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
