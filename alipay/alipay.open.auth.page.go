package alipay

import (
	"fmt"
	"net/url"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

const (
	openAuthPageURLProduction = "https://openauth.alipay.com/oauth2/appToAppAuth.htm"
	openAuthPageURLSandbox    = "https://openauth-sandbox.dl.alipaydev.com/oauth2/appToAppAuth.htm"
)

// GetOpenAuthPageURL 生成商家授权页跳转 URL（本地拼接，无网络请求）
func (c *Client) GetOpenAuthPageURL(req *types.OpenAuthPage) (string, error) {
	if req.RedirectUri == "" {
		return "", fmt.Errorf("redirect_uri is required")
	}

	baseURL := openAuthPageURLProduction
	if c.Sandbox {
		baseURL = openAuthPageURLSandbox
	}

	params := url.Values{}
	params.Set("app_id", c.AppId)
	params.Set("redirect_uri", req.RedirectUri)
	if req.State != "" {
		params.Set("state", req.State)
	}

	return baseURL + "?" + params.Encode(), nil
}
