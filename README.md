# go-aliyun-sdk

[![Go Report Card](https://goreportcard.com/badge/github.com/godrealms/go-aliyun-sdk)](https://goreportcard.com/report/github.com/godrealms/go-aliyun-sdk)
[![GoDoc](https://godoc.org/github.com/godrealms/go-aliyun-sdk?status.svg)](https://godoc.org/github.com/godrealms/go-aliyun-sdk)

Go SDK for Aliyun services, including Alipay API integration.

## 功能特性

- 支付宝接口完整封装，包括支付、退款、查询、关闭交易等
- 支持用户协议管理（签约、解约、查询、修改等）
- 账单下载功能
- 异步通知处理与验证
- 支持沙箱环境测试
- 完整的签名与验签机制
- 易于使用的API设计

## 安装

```bash
go get github.com/godrealms/go-aliyun-sdk
```

## 快速开始

### 创建客户端

```go
import "github.com/godrealms/go-aliyun-sdk/alipay"

client := alipay.NewClient()
client.AppId = "your_app_id"
client.PrivateKey = "your_private_key"
client.AlipayPublicKey = "alipay_public_key"
```

### 执行查询交易

```go
import "github.com/godrealms/go-aliyun-sdk/alipay/types"

query := &types.TradeQuery{
    OutTradeNo: "out_trade_no",
}

response, err := client.AlipayTradeQuery(query)
if err != nil {
    // 处理错误
}

// 处理响应
fmt.Println(response)
```

## 支持的接口

### 支付相关
- `alipay.trade.pay` - 统一收单交易支付接口
- `alipay.trade.precreate` - 统一收单线下交易预创建
- `alipay.trade.app.pay` - App支付接口
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

## 异步通知处理

```go
// 解析并验证异步通知
notify, err := client.Notify(request.Form)
if err != nil {
    // 处理错误
}

// 处理通知数据
fmt.Println(notify.TradeStatus)
```

## 沙箱环境

```go
client.Sandbox = true
client.Http.SetBaseURL("https://openapi-sandbox.dl.alipaydev.com/gateway.do")
```

## 贡献

欢迎提交Issue或Pull Request来改进这个项目。

## 许可证

MIT