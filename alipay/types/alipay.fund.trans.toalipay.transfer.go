package types

import "encoding/json"

// FundTransToalipayTransfer 转账到支付宝账户请求（旧版）
type FundTransToalipayTransfer struct {
	//【描述】商家流水号，唯一标识转账请求
	OutBizNo string `json:"out_biz_no"`
	//【描述】收款方账户类型。ALIPAY_USERID | ALIPAY_LOGONID
	PayeeType string `json:"payee_type"`
	//【描述】收款方账户
	PayeeAccount string `json:"payee_account"`
	//【描述】转账金额，单位元，精确到小数点后两位
	Amount string `json:"amount"`
	//【描述】收款方真实姓名
	PayeeRealName string `json:"payee_real_name,omitempty"`
	//【描述】转账备注
	Remark string `json:"remark,omitempty"`
}

func (r *FundTransToalipayTransfer) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// FundTransToalipayTransferDetail 转账到支付宝账户响应详情
type FundTransToalipayTransferDetail struct {
	PublicResponseParameters
	//【描述】商家流水号
	OutBizNo string `json:"out_biz_no"`
	//【描述】支付宝转账单据号
	OrderId string `json:"order_id"`
	//【描述】支付时间，格式 yyyy-MM-dd HH:mm:ss
	PayDate string `json:"pay_date"`
}

// AlipayFundTransToalipayTransferResponse 转账到支付宝账户响应
type AlipayFundTransToalipayTransferResponse struct {
	AlipayFundTransToalipayTransferResponse FundTransToalipayTransferDetail `json:"alipay_fund_trans_toalipay_transfer_response"`
	Sign                                    string                          `json:"sign"`
}
