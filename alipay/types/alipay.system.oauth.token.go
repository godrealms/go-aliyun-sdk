package types

import "encoding/json"

// SystemOauthToken 换取/刷新用户授权访问令牌请求
type SystemOauthToken struct {
	//【描述】授权方式：authorization_code（换取 token）| refresh_token（刷新 token）
	//【示例值】authorization_code
	GrantType string `json:"grant_type"`
	//【描述】授权码，grant_type=authorization_code 时必填
	//【示例值】4b203fe6c11548bcabd8da5bb087a83b
	Code string `json:"code,omitempty"`
	//【描述】刷新令牌，grant_type=refresh_token 时必填
	//【示例值】201208134b203fe6c11548bcabd8da5bb087a83b
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (r *SystemOauthToken) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// SystemOauthTokenDetail 换取用户授权令牌响应详情
// 注意：该接口的 code/msg 位于 alipay_system_oauth_token_response 对象内部，
// 因此 PublicResponseParameters 嵌入在 Detail 中而非顶层（有别于 trade 系列接口）。
type SystemOauthTokenDetail struct {
	PublicResponseParameters
	//【描述】支付宝用户 ID
	UserId string `json:"user_id"`
	//【描述】访问令牌，有效期见 expires_in
	AccessToken string `json:"access_token"`
	//【描述】访问令牌有效期，单位秒
	ExpiresIn int64 `json:"expires_in"`
	//【描述】刷新令牌，有效期见 re_expires_in
	RefreshToken string `json:"refresh_token"`
	//【描述】刷新令牌有效期，单位秒
	ReExpiresIn int64 `json:"re_expires_in"`
}

// AlipaySystemOauthTokenResponse 换取用户授权令牌响应
type AlipaySystemOauthTokenResponse struct {
	AlipaySystemOauthTokenResponse SystemOauthTokenDetail `json:"alipay_system_oauth_token_response"`
	Sign                           string                 `json:"sign"`
}
