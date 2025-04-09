package alipay

import "github.com/godrealms/go-aliyun-sdk/community"

type Client struct {
	AppId           string // 应用ID
	PrivateKey      string // 应用私钥
	PublicKey       string // 应用公钥
	AlipayPublicKey string // 支付宝公钥
	ReturnUrl       string // 支付宝同步通知地址
	NotifyUrl       string // 支付宝异步通知地址
	AppAuthToken    string // 应用授权令牌
	Sandbox         bool   // 是否沙箱环境
	Http            *community.HTTP
}

func NewClient() *Client {
	client := &Client{
		Http: community.NewHTTP("https://api.community.alipay.com/gateway.do"),
	}
	return client
}
