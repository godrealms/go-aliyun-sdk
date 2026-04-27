package alipay

import (
	"fmt"
	"sync"

	"github.com/godrealms/go-aliyun-sdk/community"
)

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

	signerOnce sync.Once
	signer     *community.SignatureHelper
	signerErr  error
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

// getSigner 返回缓存的签名助手，首次调用时从 PrivateKey 初始化。
// PrivateKey 必须在第一次 API 调用前设置；之后修改不生效。
func (c *Client) getSigner() (*community.SignatureHelper, error) {
	c.signerOnce.Do(func() {
		c.signer, c.signerErr = community.NewSignatureHelper(c.PrivateKey)
	})
	if c.signerErr != nil {
		return nil, fmt.Errorf("signature helper init: %w", c.signerErr)
	}
	if c.signer == nil {
		return nil, fmt.Errorf("signature helper not initialized: PrivateKey may be empty")
	}
	return c.signer, nil
}
