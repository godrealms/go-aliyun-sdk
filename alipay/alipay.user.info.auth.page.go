package alipay

import (
	"fmt"
	"net/url"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

const (
	userAuthPageURLProduction = "https://openauth.alipay.com/oauth2/publicAppAuthorize.htm"
	userAuthPageURLSandbox    = "https://openauth-sandbox.dl.alipaydev.com/oauth2/publicAppAuthorize.htm"
)

// GetUserAuthPageURL 生成用户 OAuth 授权页跳转 URL（本地拼接，无网络请求）
func (c *Client) GetUserAuthPageURL(req *types.UserInfoAuthPage) (string, error) {
	if req.Scope == "" {
		return "", fmt.Errorf("scope is required")
	}
	if req.RedirectUri == "" {
		return "", fmt.Errorf("redirect_uri is required")
	}

	baseURL := userAuthPageURLProduction
	if c.Sandbox {
		baseURL = userAuthPageURLSandbox
	}

	params := url.Values{}
	params.Set("app_id", c.AppId)
	params.Set("scope", req.Scope)
	params.Set("redirect_uri", req.RedirectUri)
	if req.State != "" {
		params.Set("state", req.State)
	}

	return baseURL + "?" + params.Encode(), nil
}
