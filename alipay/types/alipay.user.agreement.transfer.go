package types

import "encoding/json"

type AgreementTransfer struct {
	//【描述】支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ），如果传了该参数，其他参数会被忽略
	//【示例值】20170322450983769228
	AgreementNo string `json:"agreement_no,omitempty"`
	//【描述】协议产品码，商户和支付宝签约时确定，不同业务场景对应不同的签约产品码。这里指的是需要修改目标产品码的值
	//【示例值】CYCLE_PAY_AUTH_P
	TargetProductCode string `json:"target_product_code"`
	//【描述】周期管控规则参数period_rule_params，在签约周期扣款产品（如CYCLE_PAY_AUTH_P）时必传。
	//	周期扣款产品，会按照这里传入的参数提示用户，并对发起扣款的时间、金额、次数等做相应限制。
	PeriodRuleParams *PeriodRuleParams `json:"period_rule_params"`
}

func (r *AgreementTransfer) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

type AlipayUserAgreementTransferResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	//【描述】执行时间
	//【示例值】2019-01-23
	ExecuteTime string `json:"execute_time"`
	//【描述】周期类型
	//【示例值】DAY
	PeriodType string `json:"period_type"`
	//【描述】单次金额限制，单位为元
	//【示例值】100
	Amount string `json:"amount"`
	//【描述】总金额限制，单位为元
	//【示例值】600
	TotalAmount string `json:"total_amount"`
	//【描述】总支付次数
	//【示例值】12
	TotalPayments string `json:"total_payments"`
	//【描述】周期
	//【示例值】7
	Period string `json:"period"`
}

type AgreementTransferResponse struct {
	PublicResponseParameters
	Response AlipayUserAgreementTransferResponse `json:"alipay_user_agreement_transfer_response"`
}
