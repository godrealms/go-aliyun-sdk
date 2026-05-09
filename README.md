# go-aliyun-sdk

[![Go Report Card](https://goreportcard.com/badge/github.com/godrealms/go-aliyun-sdk)](https://goreportcard.com/report/github.com/godrealms/go-aliyun-sdk)
[![GoDoc](https://godoc.org/github.com/godrealms/go-aliyun-sdk?status.svg)](https://godoc.org/github.com/godrealms/go-aliyun-sdk)

Go SDK for Aliyun services, including Alipay API integration.

## 功能特性

- 支付宝接口完整封装，包括支付、退款、查询、关闭交易等
- 服务端建单（`trade.create`）、交易结算（`trade.order.settle`）、订单信息同步
- 分账管理：绑定/解绑分账关系、批量查询、比例查询
- 风控预咨询（`trade.advance.consult`）
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

## 升级说明（Breaking Changes）

本次更新引入了几处不兼容变更，从旧版本升级请同步调整调用方代码：

1. **所有 client API 方法新增 `ctx context.Context` 作为第一个参数。** 例如 `client.AlipayTradeQuery(req)` → `client.AlipayTradeQuery(ctx, req)`。`GetOpenAuthPageURL` / `GetUserAuthPageURL` / `Notify` 三个方法不发起网络请求，签名不变。
2. **`AlipayTradeRefund` 返回类型由 `*types.TradeRefundResponse` 调整为 `*types.AlipayTradeRefundResponse`**，与其他接口的外层包装类型保持一致。读取业务字段需经由 `resp.Response.XXX`。
3. **`Client.PrivateKey` 必须在第一次 API 调用前设置；**首次调用后，签名助手会被缓存，后续修改 `PrivateKey` 不再生效。

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
resp, err := client.AlipayOpenAuthTokenApp(context.Background(), &types.OpenAuthTokenApp{
    GrantType: "authorization_code",
    Code:      "授权码",
})
appAuthToken := resp.AlipayOpenAuthTokenAppResponse.AppAuthToken
```

### 查询交易

```go
import (
    "context"

    "github.com/godrealms/go-aliyun-sdk/alipay/types"
)

resp, err := client.AlipayTradeQuery(context.Background(), &types.TradeQuery{
    OutTradeNo: "out_trade_no",
})
```

### H5 手机网站支付

```go
client.ReturnUrl = "https://example.com/return"
client.NotifyUrl = "https://example.com/notify"

payURL, err := client.AlipayTradeWapPay(context.Background(), &types.TradeWapPay{
    OutTradeNo:  "20240101001",
    Subject:     "商品名称",
    TotalAmount: "88.00",
    ProductCode: "QUICK_WAP_WAY",
    QuitUrl:     "https://example.com/quit",
})
// 重定向到 payURL 完成支付
```

### 服务端建单

通过服务端预建订单，适合需要先在支付宝系统创建订单再唤起支付的场景：

```go
resp, err := client.AlipayTradeCreate(context.Background(), &types.TradeCreate{
    OutTradeNo:  "20240101001",
    Subject:     "商品名称",
    TotalAmount: "88.00",
    BuyerId:     "2088xxxxxxxx",  // 买家支付宝 UID
})
if err != nil {
    log.Fatal(err)
}
log.Printf("TradeNo: %s", resp.Response.TradeNo)
```

### 交易结算（分账）

对已完成的交易发起分账结算，`RoyaltyParameters` 为 JSON 字符串，描述分账比例/金额及收款方：

```go
royaltyParams := `[{"trans_in":"2088xxxxxxxx","trans_in_type":"userId","amount":"10.00","desc":"分账说明"}]`

resp, err := client.AlipayTradeOrderSettle(context.Background(), &types.TradeOrderSettle{
    OutRequestNo:      "settle_20240101001",
    TradeNo:           "2024xxxxxxxxxxxxxxxx",
    RoyaltyParameters: royaltyParams,
})
if err != nil {
    log.Fatal(err)
}
log.Printf("TradeNo: %s", resp.Response.TradeNo)
```

### 风控预咨询

在发起支付前查询风控建议，根据 `NextAction` 决定是否继续：

```go
resp, err := client.AlipayTradeAdvanceConsult(context.Background(), &types.TradeAdvanceConsult{
    OutTradeNo:  "20240101001",
    TotalAmount: "88.00",
    Subject:     "商品名称",
})
if err != nil {
    log.Fatal(err)
}
log.Printf("ConsultId: %s, NextAction: %s, RiskLevel: %s",
    resp.Response.ConsultId,
    resp.Response.NextAction,
    resp.Response.RiskLevel,
)
```

### 分账关系管理

**绑定分账关系**（`ReceiverList` 为 JSON 字符串，描述收款方信息）：

```go
receivers := `[{"type":"userId","account":"2088xxxxxxxx","name":"张三","memo":"合作商家"}]`

resp, err := client.AlipayTradeRoyaltyRelationBind(context.Background(), &types.TradeRoyaltyRelationBind{
    OutRequestNo: "bind_20240101001",
    ReceiverList: receivers,
})
```

**解绑分账关系**：

```go
resp, err := client.AlipayTradeRoyaltyRelationUnbind(context.Background(), &types.TradeRoyaltyRelationUnbind{
    OutRequestNo: "unbind_20240101001",
    ReceiverList: receivers,
})
```

**批量查询分账关系**：

```go
resp, err := client.AlipayTradeRoyaltyRelationBatchquery(context.Background(), &types.TradeRoyaltyRelationBatchquery{
    OutRequestNo: "query_20240101001",
    PageNum:      1,
    PageSize:     10,
})
if err != nil {
    log.Fatal(err)
}
log.Printf("共 %d 条", resp.Response.Count)
for _, r := range resp.Response.ReceiverInfos {
    log.Printf("  账号: %s (%s)", r.ReceiverAccount, r.ReceiverName)
}
```

**查询分账比例**：

```go
resp, err := client.AlipayTradeRoyaltyRateQuery(context.Background(), &types.TradeRoyaltyRateQuery{
    OutRequestNo: "ratequery_20240101001",
})
if err != nil {
    log.Fatal(err)
}
for _, info := range resp.Response.RoyaltyInfos {
    log.Printf("类型: %s, 比例: %s", info.RoyaltyType, info.Rate)
}
```

### 订单信息同步

将商户侧订单状态同步至支付宝，适用于先享后付、信用场景等：

```go
resp, err := client.AlipayTradeOrderinfoSync(context.Background(), &types.TradeOrderinfoSync{
    TradeNo:      "2024xxxxxxxxxxxxxxxx",
    OutRequestNo: "sync_20240101001",
    OrderType:    "CREDITCASHADVANCE",
    OrderScene:   "CONFIRM",
})
if err != nil {
    log.Fatal(err)
}
log.Printf("Code: %s", resp.Response.Code)
```

## 支持的接口

### 支付相关
- `alipay.trade.pay` - 统一收单交易支付接口（当面付条码）
- `alipay.trade.precreate` - 统一收单线下交易预创建（扫码付）
- `alipay.trade.app.pay` - App 支付接口
- `alipay.trade.page.pay` - 电脑网站支付接口
- `alipay.trade.wap.pay` - 手机网站支付接口（H5）

### 交易管理
- `alipay.trade.query` - 统一收单交易查询
- `alipay.trade.refund` - 统一收单交易退款接口
- `alipay.trade.fastpay.refund.query` - 统一收单交易退款查询
- `alipay.trade.close` - 统一收单交易关闭接口
- `alipay.trade.cancel` - 统一收单交易撤销接口
- `alipay.trade.create` - 统一收单下单（服务端建单）
- `alipay.trade.order.settle` - 统一收单交易结算（分账）
- `alipay.trade.orderinfo.sync` - 商户订单信息同步

### 风控预咨询
- `alipay.trade.advance.consult` - 交易风控预咨询

### 分账管理
- `alipay.trade.royalty.relation.bind` - 分账关系绑定
- `alipay.trade.royalty.relation.unbind` - 分账关系解绑
- `alipay.trade.royalty.relation.batchquery` - 分账关系批量查询
- `alipay.trade.royalty.rate.query` - 分账比例查询

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
```

设置 `Sandbox = true` 即可，客户端会自动将网关切换为沙箱地址 `https://openapi-sandbox.dl.alipaydev.com/gateway.do`。无需手动调用 `Http.SetBaseURL`。

## 贡献

欢迎提交 Issue 或 Pull Request 来改进这个项目。

## 许可证

Apache License 2.0
