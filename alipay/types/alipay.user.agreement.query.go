package types

import "encoding/json"

// UserAgreementQuery 支付宝个人代扣协议查询接口参数
type UserAgreementQuery struct {
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
	//【描述】协议产品码，商户和支付宝签约时确定，商户可咨询技术支持。
	//【示例值】GENERAL_WITHHOLDING_P
	PersonalProductCode string `json:"personal_product_code,omitempty"`
	//【描述】用户的支付宝登录账号，支持邮箱或手机号码格式。本参数与alipay_open_id 或 alipay_user_id 同时填写，
	//	优先按照 alipay_open_id 或 alipay_user_id 处理。
	//【示例值】abx@alitest.com
	AlipayLogonId string `json:"alipay_logon_id,omitempty"`
	//【描述】签约场景码，该值需要与系统/页面签约接口调用时传入的值保持一 致。
	//	如：周期扣款场景与调用 alipay.user.agreement.page.sign(支付宝个人协议页面签约接口) 签约时的 sign_scene 相同。
	//	注意：当传入商户签约号 external_agreement_no 时，该值不能为空或默认值 DEFAULT|DEFAULT。
	//【示例值】INDUSTRY|MEDICAL
	SignScene string `json:"sign_scene,omitempty"`
	//【描述】代扣协议中标示用户的唯一签约号(确保在商户系统中 唯一)。 格式规则:支持大写小写字母和数字，最长 32 位。
	//【示例值】test
	ExternalAgreementNo string `json:"external_agreement_no,omitempty"`
	//【描述】签约第三方主体类型。对于三方协议，表示当前用户和哪一类的第三方主体进行签约。 默认为PARTNER。
	//【枚举值】
	//	平台商户: PARTNER
	//【示例值】PARTNER
	ThirdPartyType string `json:"third_party_type,omitempty"`
	//【描述】支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ） ，如果传了该参数，其他参数会被忽略
	//【示例值】20170322450983769228
	AgreementNo string `json:"agreement_no"`
}

func (q *UserAgreementQuery) ToString() string {
	marshal, _ := json.Marshal(q)
	return string(marshal)
}

type ExecutionPlan struct {
	//【描述】周期扣中单笔金额，单位是元
	//【示例值】100.00
	SingleAmount string `json:"single_amount"`
	//【描述】该值为自然数，表示周期扣期数。
	//【示例值】1
	PeriodId string `json:"period_id"`
	//【描述】周期扣预期执行时间，格式为YYYY-MM-DD
	//【示例值】2024-04-29
	ExecuteTime string `json:"execute_time"`
	//【描述】周期扣执行计划最晚执行时间，格式为YYYY-MM-DD
	//【示例值】2024-05-30
	LatestExecuteTime string `json:"latest_execute_time"`
}

type AgreementQueryResponseInfo struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	//【描述】签约主体标识。 当principal_type为CARD 时，该字段为支付宝用户号;
	//	新商户建议使用principal_open_id替代该字段。对于新商户，principal_id字段未来计划逐步回收，存量商户可继续使用。
	//	如使用principal_open_id，请确认 应用-开发配置-openid配置管理 已启用。
	//	无该配置项，可查看openid配置申请: https://opendocs.alipay.com/mini/0ai9ok?pathHash=de631c06。
	//【示例值】2088101122675263
	PrincipalId string `json:"principal_id"`
	//【描述】签约主体标识。 当principal_type为CARD 时，该字段为支付宝用户号;
	//	详情可查看: https://opendocs.alipay.com/mini/0ai2i6?pathHash=13dd5946
	//【示例值】074a1CcTG1LelxKe4xQC0zgNdId0nxi95b5lsNpazWYoCo5
	PrincipalOpenId string `json:"principal_open_id"`
	ValidTime       string `json:"valid_time"`
	AlipayLogonId   string `json:"alipay_logon_id"`
	InvalidTime     string `json:"invalid_time"`
	PricipalType    string `json:"pricipal_type"`
	//【描述】设备Id
	//【示例值】RSED235F875932
	DeviceId string `json:"device_id"`
	//【描述】签约场景码。
	//	如：周期扣款场景与调用 alipay.user.agreement.page.sign(支付宝个人协议页面签约接口) 签约时的 sign_scene 相同。
	//	注意：当传入商户签约号 external_agreement_no 时，该值不能为空或默认值 DEFAULT|DEFAULT。
	//【示例值】INDUSTRY|MEDICAL
	SignScene string `json:"sign_scene"`
	//【描述】支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ） ，如果传了该参数，其他参数会被忽略
	//【示例值】20170322450983769228
	AgreementNo string `json:"agreement_no"`
	//【描述】签约第三方主体类型。对于三方协议，表示当前用户和哪一类的第三方主体进行签约。 默认为PARTNER。
	//【枚举值】
	//	平台商户: PARTNER
	//【示例值】PARTNER
	ThirdPartyType      string `json:"third_party_type"`
	Status              string `json:"status"`
	SignTime            string `json:"sign_time"`
	PersonalProductCode string `json:"personal_product_code"`
	//【描述】代扣协议中标示用户的唯一签约号(确保在商户系统中唯一)。当入参中传了此参数时返回。
	//【示例值】test
	ExternalAgreementNo string `json:"external_agreement_no"`
	//【描述】用户的芝麻信用 openId，供商 户查询用户芝麻信用使用。
	//【示例值】268816057852461313538942792
	ZmOpenId string `json:"zm_open_id"`
	//【描述】外部登录Id。当入参中传了此参数时返回。
	//【示例值】2088101118392209
	ExternalLogonId string `json:"external_logon_id"`
	//【描述】授信模式，目前只在花芝代扣（即花芝go）协议时才会返回
	//【枚举值】
	//	花芝GO: DEDUCT_HUAZHI
	//【示例值】DEDUCT_HUAZHI
	CreditAuthMode string `json:"credit_auth_mode"`
	//【描述】单笔代扣额度
	//【示例值】100.00
	SingleQuota string `json:"single_quota"`
	//【描述】周期扣协议，上次扣款成功时间
	//【示例值】2022-05-15
	LastDeductTime string `json:"last_deduct_time"`
	//【描述】周期扣协议，预计下次扣款时间
	//【示例值】2022-06-15
	NextDeductTime string `json:"next_deduct_time"`
	//【描述】还款计划列表
	ExecutionPlans []ExecutionPlan `json:"execution_plans"`
}

// AgreementQueryResponse 支付宝个人代扣协议查询接口
type AgreementQueryResponse struct {
	PublicResponseParameters
	Response AgreementQueryResponseInfo `json:"alipay_user_agreement_query_response"`
}
