package alipay

import "github.com/godrealms/go-aliyun-sdk/community"

const (
	// GatewayURL 生产环境网关
	GatewayURL = "https://openapi.alipay.com/gateway.do"
	// SandboxGatewayURL 沙箱环境网关
	SandboxGatewayURL = "https://openapi-sandbox.dl.alipaydev.com/gateway.do"
)

type Client struct {
	AppId           string // 应用ID
	PrivateKey      string // 应用私钥
	PublicKey       string // 应用公钥
	AlipayPublicKey string // 支付宝公钥
	ReturnUrl       string // 支付宝同步通知地址
	NotifyUrl       string // 支付宝异步通知地址
	AppAuthToken    string // 应用授权令牌
	Sandbox         bool   // 是否沙箱环境（影响网关及授权页URL）
	Http            *community.HTTP
}

func NewClient() *Client {
	client := &Client{
		Http: community.NewHTTP(""),
	}
	// 通过 BaseURLFunc 按 Sandbox 状态动态解析网关；
	// 若调用方通过 Http.SetBaseURL 显式指定（如测试或私有网关），则以该值为准。
	client.Http.BaseURLFunc = client.Gateway
	return client
}

func NewISVClient(appId, privateKey, alipayPublicKey, appAuthToken string) *Client {
	client := NewClient()
	client.AppId = appId
	client.PrivateKey = privateKey
	client.AlipayPublicKey = alipayPublicKey
	client.AppAuthToken = appAuthToken
	return client
}

// Gateway 根据 Sandbox 标志返回应使用的网关 URL。
func (c *Client) Gateway() string {
	if c.Sandbox {
		return SandboxGatewayURL
	}
	return GatewayURL
}
