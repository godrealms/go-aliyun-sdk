package types

import "encoding/json"

type OperateType string

const (
	OperateTypeConfirm OperateType = "confirm" // 解约确认
	OperateTypeInvalid OperateType = "invalid" // 解约作废
)

// AgreementUnSign 支付宝个人代扣协议解约接口
type AgreementUnSign struct {
	//【描述】用户的支付宝账号对应 的支付宝唯一用户号，以 2088 开头的 16 位纯数字 组成。
	//	本参数与alipay_logon_id若都填写，则以本参数为准，优先级高于 alipay_logon_id。
	//	新商户建议使用alipay_open_id替代该字段。对于新商户，alipay_user_id字段未来计划逐步回收，存量商户可继续使用。
	//	如使用alipay_open_id，请确认 应用-开发配置-openid配置管理 已启用。无该配置项，可查看openid配置申请。
	//【示例值】2088101122675263
	AlipayUserId string `json:"alipay_user_id,omitempty"`
	//【描述】用户的支付宝账号对应 的支付宝唯一用户号， 本参数与alipay_logon_id若都填写，则以本参数为准，优先级高于 alipay_logon_id。
	//	详情可查看: https://opendocs.alipay.com/mini/0ai2i6?pathHash=13dd5946
	//【示例值】074a1CcTG1LelxKe4xQC0zgNdId0nxi95b5lsNpazWYoCo5
	AlipayOpenId string `json:"alipay_open_id,omitempty"`
	//【描述】用户的支付宝登录账号，支持邮箱或手机号码格式。本参数与alipay_user_id 不可同时为空，若都填写，则以alipay_user_id 为准。
	//【示例值】abx@alitest.com
	AlipayLogonId string `json:"alipay_logon_id,omitempty"`
	//【描述】代扣协议中标示用户的唯一签约号(确保在商户系统中 唯一)。
	//	格式规则:支持大写小写字母和数字，最长 32 位。
	//	注意：若调用 alipay.user.agreement.page.sign(支付宝个人协议页面签约接口) 签约时传入 external_agreement_no 则该值必填且需与签约接口传入值相同。
	//【示例值】test
	ExternalAgreementNo string `json:"external_agreement_no,omitempty"`
	//【描述】支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ），如果传了该参数，其他参数会被忽略 。
	//	本参数与 external_agreement_no 不可同时为空。
	//【示例值】20170322450983769228
	AgreementNo string `json:"agreement_no,omitempty"`
	//【描述】协议产品码，商户和支付宝签约时确定，不同业务场景对应不同的签约产品码。
	//【示例值】GENERAL_WITHHOLDING_P
	PersonalProductCode string `json:"personal_product_code,omitempty"`
	//【描述】签约协议场景，该值需要与系统/页面签约接口调用时传入的值保持一 致。
	//	如：周期扣款场景，需与调用 alipay.user.agreement.page.sign(支付宝个人协议页面签约接口) 签约时的 sign_scene 相同。
	//	当传入商户签约号 external_agreement_no时，场景不能为空或默认值 DEFAULT|DEFAULT。
	//【示例值】INDUSTRY|MEDICA
	SignScene string `json:"sign_scene,omitempty"`
	//【描述】签约第三方主体类型。对于三方协议，表示当前用户和哪一类的第三方主体进行签约。
	//【枚举值】
	//	平台商户: PARTNER
	//【注意事项】默认为PARTNER
	//【示例值】PARTNER
	ThirdPartyType string `json:"third_party_type,omitempty"`
	//【描述】扩展参数
	//【示例值】{"UNSIGN_ERROR_CODE": "USER_OWE_MONEY","UNSIGN_ERROR_INFO":"10.00"}
	ExtendParams string `json:"extend_params"`
	//【描述】注意：仅异步解约需传入，其余情况无需传递本参数。
	//【枚举值】
	//	解约确认: confirm
	//	解约作废: invalid
	//【示例值】confirm
	OperateType OperateType `json:"operate_type"`
}

func (r *AgreementUnSign) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

type UnsignResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// AgreementUnsignResponse 业务响应参数
type AgreementUnsignResponse struct {
	PublicResponseParameters
	Response UnsignResponse `json:"alipay_user_agreement_unsign_response"`
}
