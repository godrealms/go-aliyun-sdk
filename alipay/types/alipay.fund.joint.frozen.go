package types

import "encoding/json"

// FundJointFrozen 资金冻结请求
type FundJointFrozen struct {
	//【描述】商户冻结流水号，唯一标识本次冻结请求
	OutRequestNo string `json:"out_request_no"`
	//【描述】付款方支付宝 UID
	PayerUserId string `json:"payer_user_id"`
	//【描述】冻结金额，单位元，精确到小数点后两位
	Amount string `json:"amount"`
	//【描述】冻结说明，用于收款方账单展示
	Remark string `json:"remark,omitempty"`
	//【描述】扩展参数，JSON 字符串格式
	ExtraParam string `json:"extra_param,omitempty"`
}

func (r *FundJointFrozen) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// FundJointFrozenDetail 资金冻结响应详情
type FundJointFrozenDetail struct {
	PublicResponseParameters
	//【描述】支付宝冻结单号
	FreezeId string `json:"freeze_id"`
	//【描述】冻结时间，格式 yyyy-MM-dd HH:mm:ss
	FreezeDate string `json:"freeze_date"`
	//【描述】实际冻结金额，单位元
	Amount string `json:"amount"`
}

// AlipayFundJointFrozenResponse 资金冻结响应
type AlipayFundJointFrozenResponse struct {
	AlipayFundJointFrozenResponse FundJointFrozenDetail `json:"alipay_fund_joint_frozen_response"`
	Sign                          string                `json:"sign"`
}
