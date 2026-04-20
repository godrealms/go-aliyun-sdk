package types

import "encoding/json"

// FundJointDeduct 从冻结金额扣款请求
type FundJointDeduct struct {
	//【描述】商户扣款流水号，唯一标识本次扣款请求
	OutRequestNo string `json:"out_request_no"`
	//【描述】支付宝冻结单号，由冻结接口返回
	FreezeId string `json:"freeze_id"`
	//【描述】扣款金额，单位元，不超过冻结金额
	Amount string `json:"amount"`
	//【描述】转入方支付宝账号，不填则转入商户自身账号
	TransIn string `json:"trans_in,omitempty"`
	//【描述】扣款说明
	Remark string `json:"remark,omitempty"`
}

func (r *FundJointDeduct) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// FundJointDeductDetail 从冻结金额扣款响应详情
type FundJointDeductDetail struct {
	PublicResponseParameters
	//【描述】商户扣款流水号
	OutRequestNo string `json:"out_request_no"`
	//【描述】支付宝扣款单号
	OrderId string `json:"order_id"`
	//【描述】实际扣款金额，单位元
	Amount string `json:"amount"`
	//【描述】扣款时间，格式 yyyy-MM-dd HH:mm:ss
	DeductDate string `json:"deduct_date"`
}

// AlipayFundJointDeductResponse 从冻结金额扣款响应
type AlipayFundJointDeductResponse struct {
	AlipayFundJointDeductResponse FundJointDeductDetail `json:"alipay_fund_joint_deduct_response"`
	Sign                          string                `json:"sign"`
}
