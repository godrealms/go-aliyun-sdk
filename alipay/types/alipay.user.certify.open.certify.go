package types

import "encoding/json"

// UserCertifyOpenCertify 开始认证请求
type UserCertifyOpenCertify struct {
	//【描述】认证单号，由 AlipayUserCertifyOpenInitialize 返回
	//【示例值】OcCp2413fkv09diXXXXX
	CertifyId string `json:"certify_id"`
}

func (r *UserCertifyOpenCertify) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}
