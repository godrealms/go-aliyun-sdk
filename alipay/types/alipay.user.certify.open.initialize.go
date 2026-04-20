package types

import "encoding/json"

// UserCertifyOpenInitialize 实名认证初始化请求
type UserCertifyOpenInitialize struct {
	//【描述】商户请求的唯一标识，64 个字符以内
	//【示例值】order_certify_001
	OuterOrderNo string `json:"outer_order_no"`
	//【描述】认证方案，固定填 FACE
	//【示例值】FACE
	BizCode string `json:"biz_code"`
	//【描述】需要验证的身份信息，JSON 字符串，例：{"identity_type":"CERT_INFO","cert_type":"IDENTITY_CARD","cert_name":"张三","cert_no":"310000199001011234"}
	//【示例值】{"identity_type":"CERT_INFO","cert_type":"IDENTITY_CARD","cert_name":"张三","cert_no":"310000199001011234"}
	IdentityParam string `json:"identity_param"`
	//【描述】商户个性化配置，JSON 字符串，例：{"return_url":"https://example.com/certify/callback"}
	//【示例值】{"return_url":"https://example.com/certify/callback"}
	MerchantConfig string `json:"merchant_config,omitempty"`
}

func (r *UserCertifyOpenInitialize) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// UserCertifyOpenInitializeDetail 实名认证初始化响应详情
type UserCertifyOpenInitializeDetail struct {
	PublicResponseParameters
	//【描述】认证单号，后续认证和查询使用
	CertifyId string `json:"certify_id"`
}

// AlipayUserCertifyOpenInitializeResponse 实名认证初始化响应
type AlipayUserCertifyOpenInitializeResponse struct {
	AlipayUserCertifyOpenInitializeResponse UserCertifyOpenInitializeDetail `json:"alipay_user_certify_open_initialize_response"`
	Sign                                    string                          `json:"sign"`
}
