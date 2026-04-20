# 用户信息线 API 支持设计文档

**日期：** 2026-04-20
**作者：** WuJie

---

## 背景

当前 SDK 已实现交易支付、用户协议、ISV 授权、资金转账共 30 个接口。用户 OAuth 信息获取（`alipay.user.info.share`）和实名认证（`alipay.user.certify.open.*`）是接入支付宝登录与身份核验的核心需求，尚未覆盖。

---

## 目标

新增 6 个用户信息线方法，挂在现有 `*Client` 上，与现有接口组织方式完全一致。开发在 `develop` 分支进行。

---

## 接口清单

| 方法名 | API / 说明 |
|--------|-----------|
| `GetUserAuthPageURL` | 本地构建用户 OAuth 授权页 URL（不发网络请求） |
| `AlipaySystemOauthToken` | `alipay.system.oauth.token` — 换取/刷新用户 access_token |
| `AlipayUserInfoShare` | `alipay.user.info.share` — 获取用户基本信息 |
| `AlipayUserCertifyOpenInitialize` | `alipay.user.certify.open.initialize` — 实名认证初始化 |
| `AlipayUserCertifyOpenCertify` | `alipay.user.certify.open.certify` — 开始认证（返回认证页 URL） |
| `AlipayUserCertifyOpenQuery` | `alipay.user.certify.open.query` — 查询认证结果 |

---

## 架构设计

复用现有模式，不引入新包或新类型。

### 目录结构

```
alipay/
  alipay.user.info.auth.page.go
  alipay.system.oauth.token.go
  alipay.user.info.share.go
  alipay.user.certify.open.initialize.go
  alipay.user.certify.open.certify.go
  alipay.user.certify.open.query.go
  types/
    alipay.user.info.auth.page.go
    alipay.system.oauth.token.go
    alipay.user.info.share.go
    alipay.user.certify.open.initialize.go
    alipay.user.certify.open.certify.go
    alipay.user.certify.open.query.go
example/alipay/
  alipay.user.info.auth.page.go
  alipay.system.oauth.token.go
  alipay.user.info.share.go
  alipay.user.certify.open.initialize.go
  alipay.user.certify.open.certify.go
  alipay.user.certify.open.query.go
```

### PublicRequestParameters 修改

在现有 `types/public.request.parameters.go` 中新增 `AuthToken` 字段：

```go
AuthToken string `json:"auth_token,omitempty"`
```

该字段用于传递用户级 OAuth token（`user_access_token`），参与签名，与现有 `AppAuthToken` 平行。

### 特殊接口说明

- **`GetUserAuthPageURL`**：本地拼接，不签名，不发网络请求。URL 格式：
  `https://openauth.alipay.com/oauth2/publicAppAuthorize.htm?app_id=...&scope=...&redirect_uri=...&state=...`
  沙箱地址：`https://openauth-sandbox.dl.alipaydev.com/oauth2/publicAppAuthorize.htm`

- **`AlipayUserInfoShare`**：biz_content 为空 `{}`，`authToken` 作为方法参数注入 `PublicRequestParameters.AuthToken`，其余流程与其它接口完全一致。

- **`AlipayUserCertifyOpenCertify`**：该接口不返回 JSON，浏览器直接跳转到认证页。实现方式：将签名后参数拼成 GET URL 字符串返回，调用方负责跳转。URL 格式与 `GetUserAuthPageURL` 类似（gateway + 参数）。

---

## 类型定义

### PublicRequestParameters 新增字段

```go
// 在 types/public.request.parameters.go 中新增
AuthToken string `json:"auth_token,omitempty"`
```

同时在 `ToUrlValue()` 方法中确保 `auth_token` 被正确序列化（若非空）。

### alipay.user.info.auth.page（本地 URL 构建）

```go
type UserInfoAuthPage struct {
    Scope       string `json:"scope"`              // auth_base | auth_user | auth_user_mobile
    RedirectUri string `json:"redirect_uri"`
    State       string `json:"state,omitempty"`
}
```

### alipay.system.oauth.token

```go
type SystemOauthToken struct {
    GrantType    string `json:"grant_type"`              // authorization_code | refresh_token
    Code         string `json:"code,omitempty"`
    RefreshToken string `json:"refresh_token,omitempty"`
}

func (r *SystemOauthToken) ToString() string

type SystemOauthTokenDetail struct {
    PublicResponseParameters
    UserId       string `json:"user_id"`
    AccessToken  string `json:"access_token"`
    ExpiresIn    string `json:"expires_in"`
    RefreshToken string `json:"refresh_token"`
    ReExpiresIn  string `json:"re_expires_in"`
}

type AlipaySystemOauthTokenResponse struct {
    AlipaySystemOauthTokenResponse SystemOauthTokenDetail `json:"alipay_system_oauth_token_response"`
    Sign string `json:"sign"`
}
```

### alipay.user.info.share

```go
type UserInfoShare struct{} // biz_content 为空，无字段

func (r *UserInfoShare) ToString() string // 返回 "{}"

type UserInfoShareDetail struct {
    PublicResponseParameters
    UserId      string `json:"user_id"`
    Avatar      string `json:"avatar"`
    Province    string `json:"province"`
    City        string `json:"city"`
    NickName    string `json:"nick_name"`
    Gender      string `json:"gender"`       // F | M
    Mobile      string `json:"mobile"`       // scope=auth_user_mobile 时返回
    UserStatus  string `json:"user_status"`  // T | F
    IsCertified string `json:"is_certified"` // T | F
}

type AlipayUserInfoShareResponse struct {
    AlipayUserInfoShareResponse UserInfoShareDetail `json:"alipay_user_info_share_response"`
    Sign string `json:"sign"`
}
```

### alipay.user.certify.open.initialize

```go
type UserCertifyOpenInitialize struct {
    OuterOrderNo   string `json:"outer_order_no"`
    BizCode        string `json:"biz_code"`               // FACE
    IdentityParam  string `json:"identity_param"`         // JSON 字符串
    MerchantConfig string `json:"merchant_config,omitempty"` // JSON 字符串，含 return_url
}

func (r *UserCertifyOpenInitialize) ToString() string

type UserCertifyOpenInitializeDetail struct {
    PublicResponseParameters
    CertifyId string `json:"certify_id"`
}

type AlipayUserCertifyOpenInitializeResponse struct {
    AlipayUserCertifyOpenInitializeResponse UserCertifyOpenInitializeDetail `json:"alipay_user_certify_open_initialize_response"`
    Sign string `json:"sign"`
}
```

### alipay.user.certify.open.certify

```go
type UserCertifyOpenCertify struct {
    CertifyId string `json:"certify_id"`
}

func (r *UserCertifyOpenCertify) ToString() string

// 无 JSON 响应，方法返回认证页 URL string
```

### alipay.user.certify.open.query

```go
type UserCertifyOpenQuery struct {
    CertifyId string `json:"certify_id"`
}

func (r *UserCertifyOpenQuery) ToString() string

type UserCertifyOpenQueryDetail struct {
    PublicResponseParameters
    Passed       string `json:"passed"`                  // T | F
    IdentityInfo string `json:"identity_info,omitempty"`
    MaterialInfo string `json:"material_info,omitempty"`
}

type AlipayUserCertifyOpenQueryResponse struct {
    AlipayUserCertifyOpenQueryResponse UserCertifyOpenQueryDetail `json:"alipay_user_certify_open_query_response"`
    Sign string `json:"sign"`
}
```

---

## 方法签名

```go
func (c *Client) GetUserAuthPageURL(req *types.UserInfoAuthPage) (string, error)
func (c *Client) AlipaySystemOauthToken(req *types.SystemOauthToken) (*types.AlipaySystemOauthTokenResponse, error)
func (c *Client) AlipayUserInfoShare(authToken string) (*types.AlipayUserInfoShareResponse, error)
func (c *Client) AlipayUserCertifyOpenInitialize(req *types.UserCertifyOpenInitialize) (*types.AlipayUserCertifyOpenInitializeResponse, error)
func (c *Client) AlipayUserCertifyOpenCertify(req *types.UserCertifyOpenCertify) (string, error)
func (c *Client) AlipayUserCertifyOpenQuery(req *types.UserCertifyOpenQuery) (*types.AlipayUserCertifyOpenQueryResponse, error)
```

---

## 测试策略

复用 `alipay/testhelper_test.go` 中的 `newTestClient()`，每个接口一个测试函数，mock JSON 响应，验证关键字段解析正确。

`AlipayUserCertifyOpenCertify` 测试验证返回的 URL 包含正确的 certify_id 参数。
`GetUserAuthPageURL` 测试验证返回的 URL 包含 app_id、scope、redirect_uri。

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

- `alipay.user.agreement.*`（已实现）
- `alipay.user.auth.zhimaauth`（芝麻信用，独立产品线）
- 下一条产品线：营销（`alipay.marketing.*`）
