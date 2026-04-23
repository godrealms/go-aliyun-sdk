package alipay

import (
	"fmt"
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"github.com/godrealms/go-aliyun-sdk/community"
)

// AlipayTradeWapPay H5手机网站支付（alipay.trade.wap.pay）
func (c *Client) AlipayTradeWapPay(req *types.TradeWapPay) (string, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.trade.wap.pay",
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

	signature, err := community.NewSignatureHelper(c.PrivateKey)
	if err != nil {
		return "", err
	}
	data.Sign, err = signature.GenerateSignature(data)
	if err != nil {
		return "", err
	}
	value := data.ToUrlValue()
	return fmt.Sprintf("%s?%s", c.Http.GetBaseURL(), value.Encode()), nil
}
