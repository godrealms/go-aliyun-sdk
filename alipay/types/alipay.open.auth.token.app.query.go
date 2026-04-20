package types

import "encoding/json"

// OpenAuthTokenAppQuery 查询应用授权令牌请求
type OpenAuthTokenAppQuery struct {
	//【描述】要查询的应用授权令牌
	//【示例值】201208134b203fe6c11548bcabd8da5bb087a83b
	AppAuthToken string `json:"app_auth_token"`
}

func (r *OpenAuthTokenAppQuery) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// OpenAuthTokenAppQueryDetail 查询令牌详情
type OpenAuthTokenAppQueryDetail struct {
	PublicResponseParameters
	//【描述】授权商户的 user_id
	UserId string `json:"user_id"`
	//【描述】授权商户的 AppId
	AuthAppId string `json:"auth_app_id"`
	//【描述】接口调用凭证
	AppAuthToken string `json:"app_auth_token"`
	//【描述】刷新令牌
	AppRefreshToken string `json:"app_refresh_token"`
	//【描述】app_auth_token 有效时长，单位秒
	ExpiresIn int64 `json:"expires_in"`
	//【描述】app_refresh_token 有效时长，单位秒
	ReExpiresIn int64 `json:"re_expires_in"`
	//【描述】令牌起始时间
	TokenBeginTime string `json:"token_begin_time"`
	//【描述】令牌状态。NORMAL：正常；FREEZE：冻结
	//【枚举值】NORMAL | FREEZE
	Status string `json:"status"`
}

// AlipayOpenAuthTokenAppQueryResponse 查询应用授权令牌响应
type AlipayOpenAuthTokenAppQueryResponse struct {
	AlipayOpenAuthTokenAppQueryResponse OpenAuthTokenAppQueryDetail `json:"alipay_open_auth_token_app_query_response"`
	Sign                                string                      `json:"sign"`
}
