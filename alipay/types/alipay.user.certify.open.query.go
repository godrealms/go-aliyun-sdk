package types

import "encoding/json"

// UserCertifyOpenQuery 查询认证结果请求
type UserCertifyOpenQuery struct {
	//【描述】认证单号，由 AlipayUserCertifyOpenInitialize 返回
	//【示例值】OcCp2413fkv09diXXXXX
	CertifyId string `json:"certify_id"`
}

func (r *UserCertifyOpenQuery) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// UserCertifyOpenQueryDetail 查询认证结果响应详情
type UserCertifyOpenQueryDetail struct {
	PublicResponseParameters
	//【描述】认证是否通过：T（通过）| F（未通过）
	Passed string `json:"passed"`
	//【描述】身份信息（passed=T 时返回），JSON 字符串
	IdentityInfo string `json:"identity_info,omitempty"`
	//【描述】认证材料信息，JSON 字符串
	MaterialInfo string `json:"material_info,omitempty"`
}

// AlipayUserCertifyOpenQueryResponse 查询认证结果响应
type AlipayUserCertifyOpenQueryResponse struct {
	AlipayUserCertifyOpenQueryResponse UserCertifyOpenQueryDetail `json:"alipay_user_certify_open_query_response"`
	Sign                               string                     `json:"sign"`
}
