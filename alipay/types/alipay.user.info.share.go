package types

// UserInfoShare 获取用户基本信息请求
// biz_content 为空；AuthToken 填充到公共参数 auth_token
type UserInfoShare struct {
	//【描述】用户级 OAuth access_token（通过 AlipaySystemOauthToken 换取）
	//【示例值】composeB0b18c4*******a1ae91cb026F80
	AuthToken string `json:"-"`
}

func (r *UserInfoShare) ToString() string {
	return "{}"
}

// UserInfoShareDetail 获取用户基本信息响应详情
// 注意：该接口的 code/msg 位于 alipay_user_info_share_response 对象内部，
// 因此 PublicResponseParameters 嵌入在 Detail 中而非顶层（有别于 trade 系列接口）。
type UserInfoShareDetail struct {
	PublicResponseParameters
	//【描述】支付宝用户 ID
	UserId string `json:"user_id"`
	//【描述】用户头像地址
	Avatar string `json:"avatar"`
	//【描述】省份
	Province string `json:"province"`
	//【描述】城市
	City string `json:"city"`
	//【描述】用户昵称
	NickName string `json:"nick_name"`
	//【描述】性别：F（女）| M（男）
	Gender string `json:"gender"`
	//【描述】手机号（scope=auth_user_mobile 时返回）
	Mobile string `json:"mobile"`
	//【描述】用户状态：T（正常）| F（异常）
	UserStatus string `json:"user_status"`
	//【描述】是否通过实名认证：T | F
	IsCertified string `json:"is_certified"`
}

// AlipayUserInfoShareResponse 获取用户基本信息响应
type AlipayUserInfoShareResponse struct {
	AlipayUserInfoShareResponse UserInfoShareDetail `json:"alipay_user_info_share_response"`
	Sign                        string              `json:"sign"`
}
