package types

import "encoding/json"

type EffectType string

const (
	EffectTypeDIRECT          EffectType = "DIRECT"           // 短信签约
	EffectTypeNOTICE          EffectType = "NOTICE"           // 扫码签约
	EffectTypeALLOWINACTIVATE EffectType = "ALLOW_INACTIVATE" // 钱包签约
)

type DeviceType string

const (
	DeviceTypeVRMACHINE DeviceType = "VR_MACHINE" // VR一体机
	DeviceTypeTV        DeviceType = "TV"         // 电视
	DeviceTypeIDCARD    DeviceType = "ID_CARD"    // 身份证
	DeviceTypeWORKCARD  DeviceType = "WORK_CARD"  // 工牌
)

type Channel string

const (
	ChannelALIPAYAPP   Channel = "ALIPAYAPP"
	ChannelQRCODE      Channel = "QRCODE"
	ChannelQRCODEORSMS Channel = "QRCODEORSMS"
)

type PeriodType string

const (
	PeriodTypeDAY   PeriodType = "DAY"
	PeriodTypeMONTH PeriodType = "MONTH"
)

// AccessParams
// 【描述】请按当前接入的方式进行填充，且输入值必须为文档中的参数取值范围。
// 扫码或者短信页面签约需要拼装http的请求地址访问中间页面，钱包h5页面签约可直接拼接scheme的请求地址
type AccessParams struct {
	//【描述】目前支持以下值： 1. ALIPAYAPP （钱包h5页面签约） 2. QRCODE(扫码签约) 3. QRCODEORSMS(扫码签约或者短信签约)
	//【示例值】ALIPAYAPP
	Channel Channel `json:"channel"`
}

// SceneRuleParams
// 【描述】周期扣中场景化规则信息，例如影音会员续费、保险等场景特殊规则字段
type SceneRuleParams struct {
	//【描述】在周期扣场景化模板中配置优惠类型为优惠期玩法时需要该参数，表示后续有多少期扣款可享受优惠，值为自然数代表周期
	//【示例值】1
	DiscountPeriod string `json:"discount_period,omitempty"`
	//【描述】在周期扣场景化模板中配置优惠类型为低价玩法时需要该参数，表示代扣低价期持续的时间。单位是天，该值为自然数
	//【示例值】2
	LowPricePeriod string `json:"low_price_period,omitempty"`
}

// PeriodRuleParams
// 【描述】周期管控规则参数period_rule_params，在签约周期扣款产品（如CYCLE_PAY_AUTH_P）时必传，在签约其他产品时无需传入。
// 周期扣款产品，会按照这里传入的参数提示用户，并对发起扣款的时间、金额、次数等做相应限制。
type PeriodRuleParams struct {
	//【描述】单次扣款最大金额single_amount是周期扣款产品必填，即每次发起扣款时限制的最大金额，单位为元。商户每次发起扣款都不允许大于此金额。
	//【示例值】10.99
	SingleAmount float64 `json:"single_amount,omitempty"`
	//【描述】周期类型period_type是周期扣款产品必填，枚举值为DAY和MONTH。 DAY即扣款周期按天计，MONTH代表扣款周期按自然月。
	//	与另一参数period组合使用确定扣款周期，例如period_type为DAY，period=30，则扣款周期为30天；
	//	period_type为MONTH，period=3，则扣款周期为3个自然月。 自然月是指，不论这个月有多少天，周期都计算到月份中的同一日期。
	//	例如1月3日到2月3日为一个自然月，1月3日到4月3日为三个自然月。
	//	注意周期类型使用MONTH的时候，计划扣款时间execute_time不允许传28日之后的日期（可以传28日），
	//	以此避免有些月份可能不存在对应日期的情况。
	//【枚举值】
	//	自然日: DAY
	//	自然月: MONTH
	//【示例值】DAY
	PeriodType PeriodType `json:"period_type,omitempty"`
	//【描述】周期数period是周期扣款产品必填。与另一参数period_type组合使用确定扣款周期，
	//	例如period_type为DAY，period=90，则扣款周期为90天。
	//【示例值】3
	Period int64 `json:"period,omitempty"`
	//【描述】首次执行时间execute_time是周期扣款产品必填，即商户发起首次扣款的时间。
	//	精确到日，格式为yyyy-MM-dd 结合其他必填的扣款周期参数，会确定商户以后的扣款计划。
	//	发起扣款的时间需符合这里的扣款计划。
	//【示例值】2019-01-23
	ExecuteTime string `json:"execute_time,omitempty"`
	//【描述】总金额限制，单位为元。如果传入此参数，商户多次扣款的累计金额不允许超过此金额。
	//【示例值】600.00
	TotalAmount float64 `json:"total_amount,omitempty"`
	//【描述】总扣款次数。如果传入此参数，则商户成功扣款的次数不能超过此次数限制（扣款失败不计入）。
	//【示例值】12
	TotalPayments int64 `json:"total_payments,omitempty"`
	//【描述】周期扣中场景化规则信息，例如影音会员续费、保险等场景特殊规则字段
	SceneRuleParams *SceneRuleParams `json:"scene_rule_params,omitempty"`
}

// ZmAuthParams
// 【描述】芝麻授权信息，针对于信用代扣签约。json格式。
type ZmAuthParams struct {
	//【描述】商户在芝麻端申请的merchantId
	//【示例值】268820000000414397785
	BuckleMerchantId string `json:"buckle_merchant_id"`
	//【描述】商户在芝麻端申请的appId
	//【示例值】1001164
	BuckleAppId string `json:"buckle_app_id"`
}

// ProdParams
// 【描述】签约产品属性，json格式
type ProdParams struct {
	//【描述】预授权业务信息
	//【示例值】{"platform":"taobao"}
	AuthBizParams string `json:"personal_product_code"`
	//【描述】前置收银id，商户接入前置收银台咨询时生成
	//【示例值】23954234125612
	PreConsultId string `json:"pre_consult_id"`
	//【描述】前置营销信息，由商户接入前置收银台后生成，在拉起独立签约时传递
	//【示例值】{\\\"payOperationId\\\":\\\"20240717192716a02065e6000YNN4482\\\",\\\"head_node_sceneCode\\\":\\\"shopMiniAppPreAlipayCommon\\\",\\\"head_node_itemId\\\":\\\"24071212104015\\\"}
	PayOperationInfo string `json:"pay_operation_info"`
}

// SubMerchantParams
// 【描述】此参数用于传递子商户信息，无特殊需求时不用关注。
//
//	目前商户代扣、海外代扣、淘旅行信用住产品支持传入该参数（在销售方案中“是否允许自定义子商户信息”需要选是）。
type SubMerchantParams struct {
	//【描述】子商户的商户id
	//【示例值】2088123412341234
	SubMerchantId string `json:"sub_merchant_id,omitempty"`
	//【描述】子商户的商户名称
	//【示例值】滴滴出行
	SubMerchantName string `json:"sub_merchant_name,omitempty"`
	//【描述】子商户的服务名称
	//【示例值】滴滴出行免密支付
	SubMerchantServiceName string `json:"sub_merchant_service_name,omitempty"`
	//【描述】子商户的服务描述
	//【示例值】免密付车费，单次最高500
	SubMerchantServiceDescription string `json:"sub_merchant_service_description,omitempty"`
}

// DeviceParams
// 【描述】设备信息参数，在使用设备维度签约代扣协议时，可以传这些信息
type DeviceParams struct {
	//【描述】设备Id
	//【示例值】device12345
	DeviceId string `json:"device_id"`
	//【描述】设备名称
	//【示例值】电视
	DeviceName string `json:"device_name"`
	//【描述】设备类型
	//【枚举值】
	//	VR一体机: VR_MACHINE
	//	电视: TV
	//	身份证: ID_CARD
	//	工牌: WORK_CARD
	//【示例值】TV
	DeviceType DeviceType `json:"device_type"`
}

// IdentityParams
// 【描述】用户实名信息参数，包含：姓名、身份证号、签约指定uid。商户传入用户实名信息参数，支付宝会对比用户在支付宝端的实名信息。
type IdentityParams struct {
	//【描述】签约指定用户的uid，如用户登录的uid和指定的用户uid不一致则报错新商户建议使用sign_open_id替代该字段。
	//	对于新商户，sign_user_id字段未来计划逐步回收，存量商户可继续使用。如使用sign_open_id，请确认 应用-开发配置-openid配置管理 已启用。
	//	无该配置项，可查看openid配置申请。
	//【示例值】2088202888530893
	SignUserId string `json:"sign_user_id,omitempty"`
	//【描述】签约指定用户的openid  详情可查看 openid简介
	//【示例值】031_DfFvT0Ufzk1852BLPnhuSWiztu4NqbkO35ylXPow-Y6
	SignOpenId string `json:"sign_open_id,omitempty"`
	//【描述】用户姓名
	//【示例值】张三
	UserName string `json:"user_name,omitempty"`
	//【描述】用户身份证号
	//【示例值】61102619921108888
	CertNo string `json:"cert_no,omitempty"`
	//【描述】用户实名信息hash值
	//【示例值】8D969EEF6ECAD3C29A3A629280E686CF0C3F5D5A86AFF3CA12020C923ADC6C92
	IdentityHash string `json:"identity_hash,omitempty"`
}

type AgreementPageSign struct {
	//【描述】个人签约产品码，商户和支付宝签约时确定，商户可咨询技术支持。
	//【示例值】GENERAL_WITHHOLDING_P
	PersonalProductCode string `json:"personal_product_code"`
	//【描述】请按当前接入的方式进行填充，且输入值必须为文档中的参数取值范围。
	// 扫码或者短信页面签约需要拼装http的请求地址访问中间页面，钱包h5页面签约可直接拼接scheme的请求地址
	AccessParams *AccessParams `json:"access_params"`
	//【描述】周期管控规则参数period_rule_params，在签约周期扣款产品（如CYCLE_PAY_AUTH_P）时必传，在签约其他产品时无需传入。
	// 周期扣款产品，会按照这里传入的参数提示用户，并对发起扣款的时间、金额、次数等做相应限制。
	PeriodRuleParams *PeriodRuleParams `json:"period_rule_params,omitempty"`
	//【描述】销售产品码，商户签约的支付宝合同所对应的产品码。
	//【示例值】GENERAL_WITHHOLDING
	ProductCode string `json:"product_code,omitempty"`
	//【描述】用户在商户网站的登录账号，用于在签约页面展示，如果为空，则不展示
	//【示例值】13852852877
	ExternalLogonId string `json:"external_logon_id,omitempty"`
	//【描述】协议签约场景，商户可根据 代扣产品常见场景值 选择符合自身的行业场景。
	//	说明：当传入商户签约号 external_agreement_no 时，本参数必填，不能为默认值 DEFAULT|DEFAULT。
	//【示例值】INDUSTRY|CARRENTAL
	SignScene string `json:"sign_scene,omitempty"`
	//【描述】商户签约号，代扣协议中标示用户的唯一签约号（确保在商户系统中唯一）。
	//	格式规则：支持大写小写字母和数字，最长32位。
	//	商户系统按需自定义传入，如果同一用户在同一产品码、同一签约场景下，签订了多份代扣协议，那么需要指定并传入该值。
	//【示例值】test
	ExternalAgreementNo string `json:"external_agreement_no,omitempty"`
	//【描述】签约第三方主体类型。对于三方协议，表示当前用户和哪一类的第三方主体进行签约。 默认为PARTNER。
	//【枚举值】
	//	平台商户: PARTNER
	//【示例值】PARTNER
	ThirdPartyType string `json:"third_party_type,omitempty"`
	//【描述】当前用户签约请求的协议有效周期。
	//	整形数字加上时间单位的协议有效期，从发起签约请求的时间开始算起。
	//	目前支持的时间单位： 1. d：天 2. m：月 如果未传入，默认为长期有效。
	//【示例值】2m
	SignValidityPeriod string `json:"sign_validity_period,omitempty"`
	//【描述】芝麻授权信息，针对于信用代扣签约。json格式。
	ZmAuthParams *ZmAuthParams `json:"zm_auth_params,omitempty"`
	//【描述】签约产品属性，json格式
	ProdParams *ProdParams `json:"prod_params,omitempty"`
	//【描述】签约营销参数，此值为json格式；具体的key需与营销约定
	//【示例值】{"key":"value"}
	PromoParams string `json:"promo_params,omitempty"`
	//【描述】此参数用于传递子商户信息，无特殊需求时不用关注。
	//	目前商户代扣、海外代扣、淘旅行信用住产品支持传入该参数（在销售方案中“是否允许自定义子商户信息”需要选是）。
	SubMerchant *SubMerchantParams `json:"sub_merchant,omitempty"`
	//【描述】设备信息参数，在使用设备维度签约代扣协议时，可以传这些信息
	DeviceParams *DeviceParams `json:"device_params,omitempty"`
	//【描述】用户实名信息参数，包含：姓名、身份证号、签约指定uid。商户传入用户实名信息参数，支付宝会对比用户在支付宝端的实名信息。
	IdentityParams *IdentityParams `json:"identity_params,omitempty"`
	//【描述】协议生效类型, 用于指定协议是立即生效还是等待商户通知再生效. 可空, 不填默认为立即生效.
	//【枚举值】
	//	立即生效: DIRECT
	//	商户通知生效, 需要再次调用alipay.user.agreement.sign.effect （支付宝个人协议签约生效接口）接口推进协议生效.: NOTICE
	//	允许变更状态: ALLOW_INACTIVATE
	//【示例值】DIRECT
	AgreementEffectType EffectType `json:"agreement_effect_type,omitempty"`
	//【描述】商户希望限制的签约用户的年龄范围，min表示可签该协议的用户年龄下限，
	//	max表示年龄上限。如{"min": "18","max": "30"}表示18=<年龄<=30的用户可以签约该协议。
	//【示例值】{"min":"18","max":"30"}
	UserAgeRange string `json:"user_age_range,omitempty"`
	//【描述】签约有效时间限制，单位是秒，有效范围是0-86400，商户传入此字段会用商户传入的值否则使用支付宝侧默认值，
	//	在有效时间外进行签约，会进行安全拦截；（备注：此字段适用于需要开通安全防控的商户，且依赖商户传入生成签约时的时间戳字段timestamp）
	//【示例值】300
	EffectTime int64 `json:"effect_time,omitempty"`
}

func (p *AgreementPageSign) ToString() string {
	marshal, _ := json.Marshal(p)
	return string(marshal)
}
