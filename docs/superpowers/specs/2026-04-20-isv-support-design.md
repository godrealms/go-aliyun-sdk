# 服务商模式支持设计文档

**日期：** 2026-04-20
**作者：** WuJie

---

## 背景

当前 SDK 已支持商户直连模式。服务商（ISV）模式下，服务商使用自己的 AppId 签名，通过 `app_auth_token` 代商户调用支付宝接口。`AppAuthToken` 字段已存在于 `Client` 和 `PublicRequestParameters`，但缺少：

1. 专用的 ISV 客户端构造函数
2. 授权令牌生命周期管理接口（换取、查询、解除）
3. 商家信息查询接口
4. 商家授权页 URL 生成工具

---

## 目标

- 新增 `NewISVClient()` 工厂函数，明确标识服务商客户端
- 新增 5 个 ISV 专属方法挂载在现有 `Client` 上
- 新增配套类型定义和示例文件
- 不改动任何现有接口逻辑

---

## 方案选择

采用**方案 A：直接在现有包扩展**。

理由：与现有 16 个接口的组织方式完全一致，`AppAuthToken` 字段已就位，无需引入新包，调用方零迁移成本。

---

## 架构设计

### 目录结构

```
alipay/
  client.go                               ← 新增 NewISVClient()
  alipay.open.auth.token.app.go           ← 换取授权令牌
  alipay.open.auth.token.app.query.go     ← 查询授权令牌
  alipay.open.auth.revoke.go              ← 解除授权
  alipay.open.auth.userinfo.fetch.go      ← 查询可代运营商家信息
  alipay.open.auth.page.go                ← 生成商家授权页 URL
  types/
    alipay.open.auth.token.app.go
    alipay.open.auth.token.app.query.go
    alipay.open.auth.revoke.go
    alipay.open.auth.userinfo.fetch.go
    alipay.open.auth.page.go
example/alipay/
    alipay.open.auth.token.app.go
    alipay.open.auth.token.app.query.go
    alipay.open.auth.revoke.go
    alipay.open.auth.userinfo.fetch.go
    alipay.open.auth.page.go
```

---

## 详细设计

### 1. NewISVClient 工厂函数

文件：`alipay/client.go`

```go
func NewISVClient(appId, privateKey, alipayPublicKey, appAuthToken string) *Client {
    client := NewClient()
    client.AppId = appId
    client.PrivateKey = privateKey
    client.AlipayPublicKey = alipayPublicKey
    client.AppAuthToken = appAuthToken
    return client
}
```

返回与 `NewClient()` 相同的 `*Client` 类型，所有现有方法（`AlipayTradeQuery` 等）自动在服务商模式下工作，因为它们已经读取 `c.AppAuthToken` 并写入 `PublicRequestParameters.AppAuthToken`。

---

### 2. 接口方法签名

所有方法内部流程与现有接口一致：构造 `PublicRequestParameters` → RSA2 签名 → PostForm → 解析响应。

```go
// 换取/刷新授权令牌
func (c *Client) AlipayOpenAuthTokenApp(req *types.OpenAuthTokenApp) (*types.AlipayOpenAuthTokenAppResponse, error)

// 查询授权令牌信息
func (c *Client) AlipayOpenAuthTokenAppQuery(req *types.OpenAuthTokenAppQuery) (*types.AlipayOpenAuthTokenAppQueryResponse, error)

// 解除商家授权
func (c *Client) AlipayOpenAuthRevoke(req *types.OpenAuthRevoke) (*types.AlipayOpenAuthRevokeResponse, error)

// 查询可代运营的商家信息
func (c *Client) AlipayOpenAuthUserinfoFetch(req *types.OpenAuthUserinfoFetch) (*types.AlipayOpenAuthUserinfoFetchResponse, error)

// 生成商家授权页跳转 URL（纯本地拼接，无网络请求，无需签名）
func (c *Client) GetOpenAuthPageURL(req *types.OpenAuthPage) (string, error)
```

---

### 3. 类型定义

#### alipay.open.auth.token.app

```go
type OpenAuthTokenApp struct {
    GrantType    string `json:"grant_type"`              // authorization_code | refresh_token
    Code         string `json:"code,omitempty"`
    RefreshToken string `json:"refresh_token,omitempty"`
}

type AlipayOpenAuthTokenAppResponse struct {
    AlipayOpenAuthTokenAppResponse struct {
        PublicResponseParameters
        UserId          string `json:"user_id"`
        AuthAppId       string `json:"auth_app_id"`
        AppAuthToken    string `json:"app_auth_token"`
        AppRefreshToken string `json:"app_refresh_token"`
        ExpiresIn       int64  `json:"expires_in"`
        ReExpiresIn     int64  `json:"re_expires_in"`
        TokenBeginTime  string `json:"token_begin_time"`
    } `json:"alipay_open_auth_token_app_response"`
}
```

#### alipay.open.auth.token.app.query

```go
type OpenAuthTokenAppQuery struct {
    AppAuthToken string `json:"app_auth_token"`
}

type AlipayOpenAuthTokenAppQueryResponse struct {
    AlipayOpenAuthTokenAppQueryResponse struct {
        PublicResponseParameters
        UserId          string `json:"user_id"`
        AuthAppId       string `json:"auth_app_id"`
        AppAuthToken    string `json:"app_auth_token"`
        AppRefreshToken string `json:"app_refresh_token"`
        ExpiresIn       int64  `json:"expires_in"`
        ReExpiresIn     int64  `json:"re_expires_in"`
        TokenBeginTime  string `json:"token_begin_time"`
        Status          string `json:"status"`
    } `json:"alipay_open_auth_token_app_query_response"`
}
```

#### alipay.open.auth.revoke

```go
type OpenAuthRevoke struct {
    AppAuthToken string `json:"app_auth_token"`
}

type AlipayOpenAuthRevokeResponse struct {
    AlipayOpenAuthRevokeResponse struct {
        PublicResponseParameters
    } `json:"alipay_open_auth_revoke_response"`
}
```

#### alipay.open.auth.userinfo.fetch

```go
type OpenAuthUserinfoFetch struct {
    AppAuthToken string `json:"app_auth_token"`
}

type AlipayOpenAuthUserinfoFetchResponse struct {
    AlipayOpenAuthUserinfoFetchResponse struct {
        PublicResponseParameters
        UserId   string `json:"user_id"`
        Name     string `json:"name"`
        Alias    string `json:"alias"`
        LoginId  string `json:"login_id"`
    } `json:"alipay_open_auth_userinfo_fetch_response"`
}
```

#### 商家授权页 URL（本地生成）

```go
type OpenAuthPage struct {
    RedirectUri  string // 授权回调地址
    State        string // 自定义状态值，原样返回
}
```

`GetOpenAuthPageURL` 使用 `c.AppId` 拼接 `https://openauth.alipay.com/oauth2/appToAppAuth.htm` 的跳转 URL，沙箱时切换为 `https://openauth-sandbox.dl.alipaydev.com/oauth2/appToAppAuth.htm`。

---

### 4. 示例文件规范

5 个示例文件均通过环境变量读取凭证：

```go
client := alipay.NewISVClient(
    os.Getenv("ALIPAY_APP_ID"),
    os.Getenv("ALIPAY_PRIVATE_KEY"),
    os.Getenv("ALIPAY_PUBLIC_KEY_FROM_ALIPAY"),
    os.Getenv("ALIPAY_APP_AUTH_TOKEN"),
)
```

---

## 不在本次范围内

- `alipay.open.auth.token.app.revoke`（单独吊销 token，与 revoke 不同）
- 令牌自动刷新机制
- 批量商家管理

---

## 实现清单

| 文件 | 类型 | 描述 |
|------|------|------|
| `alipay/client.go` | 修改 | 新增 `NewISVClient()` |
| `alipay/alipay.open.auth.token.app.go` | 新增 | 接口实现 |
| `alipay/alipay.open.auth.token.app.query.go` | 新增 | 接口实现 |
| `alipay/alipay.open.auth.revoke.go` | 新增 | 接口实现 |
| `alipay/alipay.open.auth.userinfo.fetch.go` | 新增 | 接口实现 |
| `alipay/alipay.open.auth.page.go` | 新增 | URL 生成 |
| `alipay/types/alipay.open.auth.token.app.go` | 新增 | 类型定义 |
| `alipay/types/alipay.open.auth.token.app.query.go` | 新增 | 类型定义 |
| `alipay/types/alipay.open.auth.revoke.go` | 新增 | 类型定义 |
| `alipay/types/alipay.open.auth.userinfo.fetch.go` | 新增 | 类型定义 |
| `alipay/types/alipay.open.auth.page.go` | 新增 | 类型定义 |
| `example/alipay/alipay.open.auth.token.app.go` | 新增 | 示例 |
| `example/alipay/alipay.open.auth.token.app.query.go` | 新增 | 示例 |
| `example/alipay/alipay.open.auth.revoke.go` | 新增 | 示例 |
| `example/alipay/alipay.open.auth.userinfo.fetch.go` | 新增 | 示例 |
| `example/alipay/alipay.open.auth.page.go` | 新增 | 示例 |
