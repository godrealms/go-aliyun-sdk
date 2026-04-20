# 资金线 API 支持设计文档

**日期：** 2026-04-20
**作者：** WuJie

---

## 背景

当前 SDK 已实现交易支付、用户协议、ISV 授权共 21 个接口。资金转账（`alipay.fund.trans.*`）和资金冻结（`alipay.fund.joint.*`）是支付之后最高频的业务需求，尚未覆盖。

---

## 目标

新增 9 个资金线接口，挂在现有 `*Client` 上，与现有接口组织方式完全一致。开发在 `develop` 分支进行。

---

## 接口清单

### alipay.fund.trans.* （6 个）

| 方法名 | API | 说明 |
|--------|-----|------|
| `AlipayFundTransUniTransfer` | `alipay.fund.trans.uni.transfer` | 新版单笔转账（转支付宝/银行卡） |
| `AlipayFundTransOrderQuery` | `alipay.fund.trans.order.query` | 新版转账订单查询 |
| `AlipayFundTransRefund` | `alipay.fund.trans.refund` | 转账撤销（T+0 内） |
| `AlipayFundTransToalipayTransfer` | `alipay.fund.trans.toalipay.transfer` | 旧版转到支付宝账户 |
| `AlipayFundTransToaccountTransfer` | `alipay.fund.trans.toaccount.transfer` | 旧版转到账户 |
| `AlipayFundTransCommonQuery` | `alipay.fund.trans.common.query` | 旧版通用转账查询 |

### alipay.fund.joint.* （3 个）

| 方法名 | API | 说明 |
|--------|-----|------|
| `AlipayFundJointFrozen` | `alipay.fund.joint.frozen` | 资金冻结（押金/预授权场景） |
| `AlipayFundJointThaw` | `alipay.fund.joint.thaw` | 资金解冻 |
| `AlipayFundJointDeduct` | `alipay.fund.joint.deduct` | 从冻结金额扣款 |

---

## 架构设计

复用现有模式，不引入新包或新类型。

### 目录结构

```
alipay/
  alipay.fund.trans.uni.transfer.go
  alipay.fund.trans.order.query.go
  alipay.fund.trans.refund.go
  alipay.fund.trans.toalipay.transfer.go
  alipay.fund.trans.toaccount.transfer.go
  alipay.fund.trans.common.query.go
  alipay.fund.joint.frozen.go
  alipay.fund.joint.thaw.go
  alipay.fund.joint.deduct.go
  types/
    alipay.fund.trans.uni.transfer.go
    alipay.fund.trans.order.query.go
    alipay.fund.trans.refund.go
    alipay.fund.trans.toalipay.transfer.go
    alipay.fund.trans.toaccount.transfer.go
    alipay.fund.trans.common.query.go
    alipay.fund.joint.frozen.go
    alipay.fund.joint.thaw.go
    alipay.fund.joint.deduct.go
example/alipay/
  alipay.fund.trans.uni.transfer.go
  alipay.fund.trans.order.query.go
  alipay.fund.trans.refund.go
  alipay.fund.trans.toalipay.transfer.go
  alipay.fund.trans.toaccount.transfer.go
  alipay.fund.trans.common.query.go
  alipay.fund.joint.frozen.go
  alipay.fund.joint.thaw.go
  alipay.fund.joint.deduct.go
```

### 请求流程

与现有接口完全一致：`PublicRequestParameters` → `biz_content` JSON → RSA2 签名 → `PostForm` → 解析响应。

---

## 类型定义

### FundTransPayee（共享子类型）

```go
type FundTransPayee struct {
    Identity     string `json:"identity"`
    IdentityType string `json:"identity_type"` // ALIPAY_USER_ID | ALIPAY_LOGON_ID
    Name         string `json:"name,omitempty"`
}
```

### alipay.fund.trans.uni.transfer

```go
type FundTransUniTransfer struct {
    OutBizNo    string          `json:"out_biz_no"`
    TransAmount string          `json:"trans_amount"`
    ProductCode string          `json:"product_code"`         // TRANS_ACCOUNT_NO_PWD
    BizScene    string          `json:"biz_scene,omitempty"`
    OrderTitle  string          `json:"order_title,omitempty"`
    PayeeInfo   *FundTransPayee `json:"payee_info"`
    Remark      string          `json:"remark,omitempty"`
}

type AlipayFundTransUniTransferResponse struct {
    AlipayFundTransUniTransferResponse struct {
        PublicResponseParameters
        OutBizNo       string `json:"out_biz_no"`
        OrderId        string `json:"order_id"`
        PayFundOrderId string `json:"pay_fund_order_id"`
        Status         string `json:"status"`
        TransDate      string `json:"trans_date"`
    } `json:"alipay_fund_trans_uni_transfer_response"`
    Sign string `json:"sign"`
}
```

### alipay.fund.trans.order.query

```go
type FundTransOrderQuery struct {
    OutBizNo string `json:"out_biz_no,omitempty"`
    OrderId  string `json:"order_id,omitempty"`
}

type AlipayFundTransOrderQueryResponse struct {
    AlipayFundTransOrderQueryResponse struct {
        PublicResponseParameters
        OrderId    string `json:"order_id"`
        Status     string `json:"status"`      // SUCCESS | FAIL | DEALING
        PayDate    string `json:"pay_date"`
        OrderFee   string `json:"order_fee"`
        FailReason string `json:"fail_reason"`
    } `json:"alipay_fund_trans_order_query_response"`
    Sign string `json:"sign"`
}
```

### alipay.fund.trans.refund

```go
type FundTransRefund struct {
    OutRequestNo string `json:"out_request_no"`
    OrderId      string `json:"order_id,omitempty"`
    OutBizNo     string `json:"out_biz_no,omitempty"`
    RefundAmount string `json:"refund_amount"`
    Remark       string `json:"remark,omitempty"`
}

type AlipayFundTransRefundResponse struct {
    AlipayFundTransRefundResponse struct {
        PublicResponseParameters
        OrderId       string `json:"order_id"`
        RefundOrderId string `json:"refund_order_id"`
        RefundAmount  string `json:"refund_amount"`
        RefundDate    string `json:"refund_date"`
    } `json:"alipay_fund_trans_refund_response"`
    Sign string `json:"sign"`
}
```

### alipay.fund.trans.toalipay.transfer

```go
type FundTransToalipayTransfer struct {
    OutBizNo      string `json:"out_biz_no"`
    PayeeType     string `json:"payee_type"`   // ALIPAY_USERID | ALIPAY_LOGONID
    PayeeAccount  string `json:"payee_account"`
    Amount        string `json:"amount"`
    PayeeRealName string `json:"payee_real_name,omitempty"`
    Remark        string `json:"remark,omitempty"`
}

type AlipayFundTransToalipayTransferResponse struct {
    AlipayFundTransToalipayTransferResponse struct {
        PublicResponseParameters
        OutBizNo string `json:"out_biz_no"`
        OrderId  string `json:"order_id"`
        PayDate  string `json:"pay_date"`
    } `json:"alipay_fund_trans_toalipay_transfer_response"`
    Sign string `json:"sign"`
}
```

### alipay.fund.trans.toaccount.transfer

```go
type FundTransToaccountTransfer struct {
    OutBizNo      string `json:"out_biz_no"`
    PayeeType     string `json:"payee_type"`   // ALIPAY_USERID | ALIPAY_LOGONID | ALIPAY_OPENID
    PayeeAccount  string `json:"payee_account"`
    Amount        string `json:"amount"`
    PayerShowName string `json:"payer_show_name,omitempty"`
    PayeeRealName string `json:"payee_real_name,omitempty"`
    Remark        string `json:"remark,omitempty"`
}

type AlipayFundTransToaccountTransferResponse struct {
    AlipayFundTransToaccountTransferResponse struct {
        PublicResponseParameters
        OutBizNo string `json:"out_biz_no"`
        OrderId  string `json:"order_id"`
        PayDate  string `json:"pay_date"`
    } `json:"alipay_fund_trans_toaccount_transfer_response"`
    Sign string `json:"sign"`
}
```

### alipay.fund.trans.common.query

```go
type FundTransCommonQuery struct {
    ProductCode    string `json:"product_code,omitempty"`
    BizScene       string `json:"biz_scene,omitempty"`
    OutBizNo       string `json:"out_biz_no,omitempty"`
    OrderId        string `json:"order_id,omitempty"`
    PayFundOrderId string `json:"pay_fund_order_id,omitempty"`
}

type AlipayFundTransCommonQueryResponse struct {
    AlipayFundTransCommonQueryResponse struct {
        PublicResponseParameters
        OrderId        string `json:"order_id"`
        PayFundOrderId string `json:"pay_fund_order_id"`
        Status         string `json:"status"`
        PayDate        string `json:"pay_date"`
        OrderFee       string `json:"order_fee"`
        FailReason     string `json:"fail_reason"`
    } `json:"alipay_fund_trans_common_query_response"`
    Sign string `json:"sign"`
}
```

### alipay.fund.joint.frozen

```go
type FundJointFrozen struct {
    OutRequestNo string `json:"out_request_no"`
    PayerUserId  string `json:"payer_user_id"`
    Amount       string `json:"amount"`
    Remark       string `json:"remark,omitempty"`
    ExtraParam   string `json:"extra_param,omitempty"`
}

type AlipayFundJointFrozenResponse struct {
    AlipayFundJointFrozenResponse struct {
        PublicResponseParameters
        FreezeId   string `json:"freeze_id"`
        FreezeDate string `json:"freeze_date"`
        Amount     string `json:"amount"`
    } `json:"alipay_fund_joint_frozen_response"`
    Sign string `json:"sign"`
}
```

### alipay.fund.joint.thaw

```go
type FundJointThaw struct {
    OutRequestNo string `json:"out_request_no"`
    FreezeId     string `json:"freeze_id"`
    Amount       string `json:"amount"`
    Remark       string `json:"remark,omitempty"`
}

type AlipayFundJointThawResponse struct {
    AlipayFundJointThawResponse struct {
        PublicResponseParameters
        OutRequestNo string `json:"out_request_no"`
        Amount       string `json:"amount"`
        ThawDate     string `json:"thaw_date"`
    } `json:"alipay_fund_joint_thaw_response"`
    Sign string `json:"sign"`
}
```

### alipay.fund.joint.deduct

```go
type FundJointDeduct struct {
    OutRequestNo string `json:"out_request_no"`
    FreezeId     string `json:"freeze_id"`
    Amount       string `json:"amount"`
    TransIn      string `json:"trans_in,omitempty"`
    Remark       string `json:"remark,omitempty"`
}

type AlipayFundJointDeductResponse struct {
    AlipayFundJointDeductResponse struct {
        PublicResponseParameters
        OutRequestNo string `json:"out_request_no"`
        OrderId      string `json:"order_id"`
        Amount       string `json:"amount"`
        DeductDate   string `json:"deduct_date"`
    } `json:"alipay_fund_joint_deduct_response"`
    Sign string `json:"sign"`
}
```

---

## 方法签名

```go
func (c *Client) AlipayFundTransUniTransfer(req *types.FundTransUniTransfer) (*types.AlipayFundTransUniTransferResponse, error)
func (c *Client) AlipayFundTransOrderQuery(req *types.FundTransOrderQuery) (*types.AlipayFundTransOrderQueryResponse, error)
func (c *Client) AlipayFundTransRefund(req *types.FundTransRefund) (*types.AlipayFundTransRefundResponse, error)
func (c *Client) AlipayFundTransToalipayTransfer(req *types.FundTransToalipayTransfer) (*types.AlipayFundTransToalipayTransferResponse, error)
func (c *Client) AlipayFundTransToaccountTransfer(req *types.FundTransToaccountTransfer) (*types.AlipayFundTransToaccountTransferResponse, error)
func (c *Client) AlipayFundTransCommonQuery(req *types.FundTransCommonQuery) (*types.AlipayFundTransCommonQueryResponse, error)
func (c *Client) AlipayFundJointFrozen(req *types.FundJointFrozen) (*types.AlipayFundJointFrozenResponse, error)
func (c *Client) AlipayFundJointThaw(req *types.FundJointThaw) (*types.AlipayFundJointThawResponse, error)
func (c *Client) AlipayFundJointDeduct(req *types.FundJointDeduct) (*types.AlipayFundJointDeductResponse, error)
```

---

## 测试策略

复用 `alipay/testhelper_test.go` 中的 `newTestClient()`，每个接口一个测试函数，mock JSON 响应，验证关键字段解析正确。

---

## 示例文件规范

所有示例通过环境变量读取凭证：

```go
client := alipay.NewClient()
client.AppId = os.Getenv("ALIPAY_APP_ID")
client.PrivateKey = os.Getenv("ALIPAY_PRIVATE_KEY")
client.AlipayPublicKey = os.Getenv("ALIPAY_PUBLIC_KEY_FROM_ALIPAY")
```

---

## 不在本次范围内

- `alipay.fund.auth.*`（预授权收款，独立产品线）
- 自动重试、令牌刷新等高级封装
- 下一条产品线：用户信息（`alipay.user.info.*`）
