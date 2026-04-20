package types

import "encoding/json"

// OpenAuthUserinfoFetch 查询可代运营商家信息请求
type OpenAuthUserinfoFetch struct {
	//【描述】商家授予服务商的 app_auth_token
	//【示例值】201208134b203fe6c11548bcabd8da5bb087a83b
	AppAuthToken string `json:"app_auth_token"`
}

func (r *OpenAuthUserinfoFetch) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// OpenAuthUserinfoFetchDetail 商家信息详情
type OpenAuthUserinfoFetchDetail struct {
	PublicResponseParameters
	//【描述】商家在支付宝的 user_id
	//【示例值】2088101117955611
	UserId string `json:"user_id"`
	//【描述】商家名称
	//【示例值】XXX 有限公司
	Name string `json:"name"`
	//【描述】商家别名
	//【示例值】XX 公司
	Alias string `json:"alias"`
	//【描述】商家账号（手机号或邮箱）
	//【示例值】152****1234
	LoginId string `json:"login_id"`
}

// AlipayOpenAuthUserinfoFetchResponse 查询可代运营商家信息响应
type AlipayOpenAuthUserinfoFetchResponse struct {
	AlipayOpenAuthUserinfoFetchResponse OpenAuthUserinfoFetchDetail `json:"alipay_open_auth_userinfo_fetch_response"`
	Sign                                string                      `json:"sign"`
}
