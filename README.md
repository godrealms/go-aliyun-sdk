# go-aliyun-sdk

[![Go Report Card](https://goreportcard.com/badge/github.com/godrealms/go-aliyun-sdk)](https://goreportcard.com/report/github.com/godrealms/go-aliyun-sdk)
[![GoDoc](https://godoc.org/github.com/godrealms/go-aliyun-sdk?status.svg)](https://godoc.org/github.com/godrealms/go-aliyun-sdk)

Go SDK for Aliyun services, including Alipay API integration.

## 功能特性

- 支付宝接口完整封装，包括支付、退款、查询、关闭交易等
- 支持用户协议管理（签约、解约、查询、修改等）
- 账单下载功能
- 异步通知处理与验证
- **服务商（ISV）模式**：通过 `app_auth_token` 代商户调用接口，支持授权令牌全生命周期管理
- 支持沙箱环境测试
- 完整的签名与验签机制（RSA2 / RSA / ECDSA）
- 易于使用的 API 设计

## 安装

```bash
go get github.com/godrealms/go-aliyun-sdk
```

## 快速开始

### 商户直连模式

```go
import "github.com/godrealms/go-aliyun-sdk/alipay"

client := alipay.NewClient()
client.AppId = "your_app_id"
client.PrivateKey = "your_private_key"       // PKCS#8 PEM
client.AlipayPublicKey = "alipay_public_key" // 支付宝公钥 PEM
```

### 服务商（ISV）模式

服务商使用自己的 AppId 签名，通过 `app_auth_token` 代商户调用接口：

```go
client := alipay.NewISVClient(
    "your_isv_app_id",
    "your_isv_private_key",
    "alipay_public_key",
    "merchant_app_auth_token", // 商户授予服务商的授权令牌
)
```

`NewISVClient` 返回与 `NewClient` 相同的 `*Client` 类型，所有现有接口自动在服务商模式下工作。

### 生成商家授权页 URL

引导商户完成授权，获取 `app_auth_token`：

```go
u, err := client.GetOpenAuthPageURL(&types.OpenAuthPage{
    RedirectUri: "https://example.com/alipay/callback",
    State:       "order_123",
})
// https://openauth.alipay.com/oauth2/appToAppAuth.htm?app_id=...
```

### 换取授权令牌

```go
resp, err := client.AlipayOpenAuthTokenApp(&types.OpenAuthTokenApp{
    GrantType: "authorization_code",
    Code:      "授权码",
})
appAuthToken := resp.AlipayOpenAuthTokenAppResponse.AppAuthToken
```

### 查询交易

```go
import "github.com/godrealms/go-aliyun-sdk/alipay/types"

resp, err := client.AlipayTradeQuery(&types.TradeQuery{
    OutTradeNo: "out_trade_no",
})
```

## 支持的接口

### 支付相关
- `alipay.trade.pay` - 统一收单交易支付接口
- `alipay.trade.precreate` - 统一收单线下交易预创建
- `alipay.trade.app.pay` - App 支付接口
- `alipay.trade.page.pay` - 电脑网站支付接口
- `alipay.trade.wap.pay` - 手机网站支付接口

### 交易管理
- `alipay.trade.query` - 统一收单交易查询
- `alipay.trade.refund` - 统一收单交易退款接口
- `alipay.trade.fastpay.refund.query` - 统一收单交易退款查询
- `alipay.trade.close` - 统一收单交易关闭接口
- `alipay.trade.cancel` - 统一收单交易撤销接口

### 用户协议管理
- `alipay.user.agreement.page.sign` - 支付宝个人协议页面签约接口
- `alipay.user.agreement.query` - 支付宝个人代扣协议查询接口
- `alipay.user.agreement.unsign` - 个人代扣协议解约接口
- `alipay.user.agreement.transfer` - 代扣协议迁移接口
- `alipay.user.agreement.executionplan.modify` - 周期性扣款协议执行计划修改接口

### 账单服务
- `alipay.data.dataservice.bill.downloadurl.query` - 查询对账单下载地址接口

### 服务商授权管理
- `alipay.open.auth.token.app` - 换取/刷新应用授权令牌
- `alipay.open.auth.token.app.query` - 查询应用授权令牌信息
- `alipay.open.auth.revoke` - 解除应用授权
- `alipay.open.auth.userinfo.fetch` - 查询可代运营的商家信息
- `GetOpenAuthPageURL` - 生成商家授权页跳转 URL（本地拼接）

## 异步通知处理

```go
notify, err := client.Notify(request.Form)
if err != nil {
    // 处理错误
}
fmt.Println(notify.TradeStatus)
```

## 沙箱环境

```go
client.Sandbox = true
client.Http.SetBaseURL("https://openapi-sandbox.dl.alipaydev.com/gateway.do")
```

## 贡献

欢迎提交 Issue 或 Pull Request 来改进这个项目。

## 许可证

Apache License 2.0
