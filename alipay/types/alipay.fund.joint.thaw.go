package types

import "encoding/json"

// FundJointThaw 资金解冻请求
type FundJointThaw struct {
	//【描述】商户解冻流水号，唯一标识本次解冻请求
	OutRequestNo string `json:"out_request_no"`
	//【描述】支付宝冻结单号，由冻结接口返回
	FreezeId string `json:"freeze_id"`
	//【描述】解冻金额，单位元，不超过冻结金额
	Amount string `json:"amount"`
	//【描述】解冻说明
	Remark string `json:"remark,omitempty"`
}

func (r *FundJointThaw) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// FundJointThawDetail 资金解冻响应详情
type FundJointThawDetail struct {
	PublicResponseParameters
	//【描述】商户解冻流水号
	OutRequestNo string `json:"out_request_no"`
	//【描述】实际解冻金额，单位元
	Amount string `json:"amount"`
	//【描述】解冻时间，格式 yyyy-MM-dd HH:mm:ss
	ThawDate string `json:"thaw_date"`
}

// AlipayFundJointThawResponse 资金解冻响应
type AlipayFundJointThawResponse struct {
	AlipayFundJointThawResponse FundJointThawDetail `json:"alipay_fund_joint_thaw_response"`
	Sign                        string              `json:"sign"`
}
