package types

import "encoding/json"

// FundTransToaccountTransfer 转账到账户请求（旧版）
type FundTransToaccountTransfer struct {
	//【描述】商家流水号，唯一标识转账请求
	OutBizNo string `json:"out_biz_no"`
	//【描述】收款方账户类型。ALIPAY_USERID | ALIPAY_LOGONID | ALIPAY_OPENID
	PayeeType string `json:"payee_type"`
	//【描述】收款方账户
	PayeeAccount string `json:"payee_account"`
	//【描述】转账金额，单位元
	Amount string `json:"amount"`
	//【描述】付款方显示姓名
	PayerShowName string `json:"payer_show_name,omitempty"`
	//【描述】收款方真实姓名
	PayeeRealName string `json:"payee_real_name,omitempty"`
	//【描述】转账备注
	Remark string `json:"remark,omitempty"`
}

func (r *FundTransToaccountTransfer) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// FundTransToaccountTransferDetail 转账到账户响应详情
type FundTransToaccountTransferDetail struct {
	PublicResponseParameters
	//【描述】商家流水号
	OutBizNo string `json:"out_biz_no"`
	//【描述】支付宝转账单据号
	OrderId string `json:"order_id"`
	//【描述】支付时间，格式 yyyy-MM-dd HH:mm:ss
	PayDate string `json:"pay_date"`
}

// AlipayFundTransToaccountTransferResponse 转账到账户响应
type AlipayFundTransToaccountTransferResponse struct {
	AlipayFundTransToaccountTransferResponse FundTransToaccountTransferDetail `json:"alipay_fund_trans_toaccount_transfer_response"`
	Sign                                     string                           `json:"sign"`
}
