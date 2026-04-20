package types

// UserInfoAuthPage 用户 OAuth 授权页参数
type UserInfoAuthPage struct {
	//【描述】授权范围：auth_base（静默授权，仅获取 user_id）| auth_user（用户信息）| auth_user_mobile（含手机号）
	Scope string
	//【描述】授权回调地址，需与应用配置一致
	RedirectUri string
	//【描述】商户自定义参数，授权完成后原样返回
	State string
}
