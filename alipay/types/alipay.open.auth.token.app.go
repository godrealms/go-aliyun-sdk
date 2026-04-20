package types

import "encoding/json"

// OpenAuthTokenApp 换取应用授权令牌请求
type OpenAuthTokenApp struct {
	//【描述】授权方式。authorization_code：使用授权码换取令牌；refresh_token：使用 refresh_token 刷新令牌
	//【示例值】authorization_code
	GrantType string `json:"grant_type"`
	//【描述】授权码，grant_type=authorization_code 时必填
	//【示例值】4b203fe6c11548bcabd8da5bb087a83b
	Code string `json:"code,omitempty"`
	//【描述】刷新令牌，grant_type=refresh_token 时必填
	//【示例值】201208134b203fe6c11548bcabd8da5bb087a83b
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (r *OpenAuthTokenApp) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// OpenAuthTokenAppDetail 授权令牌详情
type OpenAuthTokenAppDetail struct {
	PublicResponseParameters
	//【描述】授权商户的 user_id
	UserId string `json:"user_id"`
	//【描述】授权商户的 AppId
	AuthAppId string `json:"auth_app_id"`
	//【描述】接口调用凭证，有效期限内多次调用返回同一令牌
	AppAuthToken string `json:"app_auth_token"`
	//【描述】刷新令牌，有效期更长，用于换取新的 app_auth_token
	AppRefreshToken string `json:"app_refresh_token"`
	//【描述】app_auth_token 有效时长，单位秒
	ExpiresIn int64 `json:"expires_in"`
	//【描述】app_refresh_token 有效时长，单位秒
	ReExpiresIn int64 `json:"re_expires_in"`
	//【描述】令牌起始时间，格式 yyyy-MM-dd HH:mm:ss
	TokenBeginTime string `json:"token_begin_time"`
}

// AlipayOpenAuthTokenAppResponse 换取应用授权令牌响应
type AlipayOpenAuthTokenAppResponse struct {
	AlipayOpenAuthTokenAppResponse OpenAuthTokenAppDetail `json:"alipay_open_auth_token_app_response"`
	Sign                           string                 `json:"sign"`
}
