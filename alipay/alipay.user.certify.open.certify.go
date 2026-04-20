package alipay

import (
	"time"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"github.com/godrealms/go-aliyun-sdk/community"
)

// AlipayUserCertifyOpenCertify 开始实名认证，返回认证页跳转 URL（alipay.user.certify.open.certify）
// 调用方将返回的 URL 作为跳转地址，在浏览器中打开即进入认证流程。
func (c *Client) AlipayUserCertifyOpenCertify(request *types.UserCertifyOpenCertify) (string, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.user.certify.open.certify",
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
		return "", err
	}
	data.Sign, err = signature.GenerateSignature(data)
	if err != nil {
		return "", err
	}

	return c.Http.BaseURL + "?" + data.ToUrlValue().Encode(), nil
}
