package types

import "encoding/json"

// AgreementExecutionPlanModify 周期性扣款协议执行计划修改接口
type AgreementExecutionPlanModify struct {
	//【描述】周期性扣款产品，授权免密支付协议号
	//【示例值】20185909000458725113
	AgreementNo string `json:"agreement_no"`
	//【描述】商户下一次扣款时间
	//【示例值】2019-05-12
	DeductTime string `json:"deduct_time"`
	//【描述】具体修改原因
	//【示例值】用户已购买半年包，需延期扣款时间
	Memo string `json:"memo,omitempty"`
}

func (r *AgreementExecutionPlanModify) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

type AlipayUserAgreementExecutionPlanModifyResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	//【描述】周期性扣款产品，授权免密支付协议号
	//【示例值】20185909000458725113
	AgreementNo string `json:"agreement_no"`
	//【描述】商户下一次扣款时间，格式 "yyyy-MM-dd"。
	//	例如：用户在1月1日开通了连续包月，使用了10天又另行购买了“季度包”，
	//	如果此时商户希望“季度包”立即优先生效，在季度包结束后能继续使用连续包月，那么原定的周期就被延后了。
	//	此时可以通过本接口将预计扣款时间推后“季度包”的时长。
	//【示例值】2020-05-12
	DeductTime string `json:"deduct_time"`
}

type AgreementExecutionPlanModifyResponse struct {
	PublicResponseParameters
	Response AlipayUserAgreementExecutionPlanModifyResponse `json:"alipay_user_agreement_executionplan_modify_response"`
}
