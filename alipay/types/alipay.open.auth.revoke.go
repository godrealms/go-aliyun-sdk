package types

import "encoding/json"

// OpenAuthRevoke 解除应用授权请求
type OpenAuthRevoke struct {
	//【描述】要解除的应用授权令牌
	//【示例值】201208134b203fe6c11548bcabd8da5bb087a83b
	AppAuthToken string `json:"app_auth_token"`
}

func (r *OpenAuthRevoke) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// OpenAuthRevokeDetail 解除授权响应详情
type OpenAuthRevokeDetail struct {
	PublicResponseParameters
}

// AlipayOpenAuthRevokeResponse 解除应用授权响应
type AlipayOpenAuthRevokeResponse struct {
	AlipayOpenAuthRevokeResponse OpenAuthRevokeDetail `json:"alipay_open_auth_revoke_response"`
	Sign                         string               `json:"sign"`
}
