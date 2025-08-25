package types

import "encoding/json"

type CertType string

const (
	CertTypeIDENTITYCARD                CertType = "IDENTITY_CARD"                 // 身份证
	CertTypePASSPORT                    CertType = "PASSPORT"                      // 护照
	CertTypeOFFICERCARD                 CertType = "OFFICER_CARD"                  // 军官证
	CertTypeSOLDIERCARD                 CertType = "SOLDIER_CARD"                  // 士兵证
	CertTypeHOKOU                       CertType = "HOKOU"                         // 户口本
	CertTypePERMANENTRESIDENCEFOREIGNER CertType = "PERMANENT_RESIDENCE_FOREIGNER" // 外国人永久居留身份证
)

type QueryOption string

const (
	QueryOptionMedicalInsuranceInfo QueryOption = "medical_insurance_info" // 医保信息

	QueryOptionHybAmount         QueryOption = "hyb_amount"          // 惠营宝回票金额
	QueryOptionEnterprisePayInfo QueryOption = "enterprise_pay_info" // 因公付支付信息

	QueryOptionFundBillList        QueryOption = "fund_bill_list"        // 资金明细信息
	QueryOptionVoucherDetailList   QueryOption = "voucher_detail_list"   // 优惠券信息
	QueryOptionDiscountGoodsDetail QueryOption = "discount_goods_detail" // 商品优惠信息
	QueryOptionDiscountAmount      QueryOption = "discount_amount"       // 平台优惠金额
	QueryOptionMdiscountAmount     QueryOption = "mdiscount_amount"      // 商家优惠金额

	QueryOptionTradeSettleInfo QueryOption = "trade_settle_info" // 交易结算信息
	QueryOptionTapPayInfo      QueryOption = "tap_pay_info"      // 碰一下支付信息
)

type GoodsDetail struct {
	//【描述】商品的编号
	//【示例值】apple-01
	GoodsId string `json:"goods_id"`
	//【描述】商品名称
	//【示例值】ipad
	GoodsName string `json:"goods_name"`
	//【描述】商品数量
	//【示例值】1
	Quantity int64 `json:"quantity"`
	//【描述】商品单价，单位为元
	//【示例值】2000
	Price string `json:"price"`
	//【描述】支付宝定义的统一商品编号
	//【示例值】20010001
	AlipayGoodsId string `json:"alipay_goods_id"`
	//【描述】商品类目
	//【示例值】34543238
	GoodsCategory string `json:"goods_category"`
	//【描述】商品类目树，从商品类目根节点到叶子节点的类目id组成，类目id值使用|分割
	//【示例值】124868003|126232002|126252004
	CategoriesTree string `json:"categories_tree"`
	//【描述】商品的展示地址
	//【示例值】http://www.alipay.com/xxx.jpg
	ShowUrl string `json:"show_url"`
}

type ExtendParams struct {
	//【描述】系统商编号
	//	该参数作为系统商返佣数据提取的依据，请填写系统商签约协议的PID
	//【示例值】2088511833207846
	SysServiceProviderId string `json:"sys_service_provider_id"`
	//【描述】使用花呗分期要进行的分期数
	//【示例值】3
	HbFqNum string `json:"hb_fq_num"`
	//【描述】使用花呗分期需要卖家承担的手续费比例的百分值，传入100代表100%
	//【示例值】100
	HbFqSellerPercent string `json:"hb_fq_seller_percent"`
	//【描述】行业数据回流信息, 详见：地铁支付接口参数补充说明
	//【示例值】{\"scene_code\":\"metro_tradeorder\",\"channel\":\"xxxx\",\"scene_data\":{\"asset_name\":\"ALIPAY\"}}
	IndustryRefluxInfo string `json:"industry_reflux_info"`
	//【描述】卡类型
	//【枚举值】
	//	S0JP0000: S0JP0000
	//【示例值】S0JP0000
	CardType string `json:"card_type"`
	//【描述】是否进行资金冻结，用于后续分账，true表示资金冻结，false或不传表示资金不冻结
	//【示例值】true
	RoyaltyFreeze string `json:"royalty_freeze"`
}

type ExtUserInfo struct {
	//【描述】买家证件号。
	//	注：need_check_info=T时该参数才有效，支付宝会比较买家在支付宝留存的证件号码与该参数传入的值是否匹配。
	//【示例值】362334768769238881
	CertNo string `json:"cert_no"`
	//【描述】允许的最小买家年龄。
	//	买家年龄必须大于等于所传数值
	//	注：
	//	1. need_check_info=T时该参数才有效
	//	2. min_age为整数，必须大于等于0
	//【示例值】18
	MinAge string `json:"min_age"`
	//【描述】指定买家姓名。
	//	注： need_check_info=T时该参数才有效
	//【示例值】李明
	Name string `json:"name"`
	//【描述】指定买家手机号。
	//	注：该参数暂不校验
	//【示例值】16587658765
	Mobile string `json:"mobile"`
	//【描述】指定买家证件类型。
	//	枚举值：
	//		IDENTITY_CARD：身份证；
	//		PASSPORT：护照；
	//		OFFICER_CARD：军官证；
	//		SOLDIER_CARD：士兵证；
	//		HOKOU：户口本；
	//		PERMANENT_RESIDENCE_FOREIGNER：外国人永久居留身份证。 如有其它类型需要支持，请与支付宝工作人员联系。
	//	注： need_check_info=T时该参数才有效，支付宝会比较买家在支付宝留存的证件类型与该参数传入的值是否匹配。
	//【示例值】IDENTITY_CARD
	CertType CertType `json:"cert_type"`
	//【描述】是否强制校验买家信息；
	//	需要强制校验传：T;
	//	不需要强制校验传：F或者不传；
	//	当传T时，支付宝会校验支付买家的信息与接口上传递的cert_type、cert_no、name或age是否匹配，只有接口传递了信息才会进行对应项的校验；
	//	只要有任何一项信息校验不匹配交易都会失败。如果传递了need_check_info，但是没有传任何校验项，则不进行任何校验。
	//	默认为不校验。
	//【示例值】F
	NeedCheckInfo string `json:"need_check_info"`
	//【描述】买家加密身份信息。当指定了此参数且指定need_check_info=T时，
	//	支付宝会对买家身份进行校验，校验逻辑为买家姓名、买家证件号拼接后的字符串，以sha256算法utf-8编码计算hash，若与传入的值不匹配则会拦截本次支付。
	//	注意：如果同时指定了用户明文身份信息（name，cert_type，cert_no中任意一个），则忽略identity_hash以明文参数校验。
	//【示例值】27bfcd1dee4f22c8fe8a2374af9b660419d1361b1c207e9b41a754a113f38fcc
	IdentityHash string `json:"identity_hash"`
}

// SignParams
// 【描述】签约参数。如果希望在sdk中支付并签约，需要在这里传入签约信息。
//
//	周期扣款场景 product_code 为 CYCLE_PAY_AUTH 时必填。
type SignParams struct {
	//【描述】商家和支付宝签约的产品码。 商家扣款产品传入固定值：GENERAL_WITHHOLDING
	//【示例值】GENERAL_WITHHOLDING
	ProductCode string `json:"product_code"`
	//【描述】个人签约产品码，商户和支付宝签约时确定。
	//【示例值】CYCLE_PAY_AUTH_P
	PersonalProductCode string `json:"personal_product_code"`
	//【描述】协议签约场景，商户和支付宝签约时确定，商户可咨询技术支持。
	//【示例值】INDUSTRY|DIGITAL_MEDIA
	SignScene string `json:"sign_scene"`
	//【描述】请按当前接入的方式进行填充，且输入值必须为文档中的参数取值范围。
	AccessParams *AccessParams `json:"access_params,omitempty"`
	//【描述】周期管控规则参数period_rule_params，
	//	商家扣款产品，会按照这里传入的参数提示用户，并对发起扣款的时间、金额、次数等做相应限制。
	PeriodRuleParams *PeriodRuleParams `json:"period_rule_params,omitempty"`
	//【描述】设置签约请求的有效时间，单位为秒。
	//	如传入600，商户发起签约请求到用户进入支付宝签约页面的时间差不能超过10分钟。
	//【示例值】600
	EffectTime string `json:"effect_time,omitempty"`
	//【描述】商户签约号，代扣协议中标示用户的唯一签约号（确保在商户系统中唯一）。
	//	格式规则：支持大写小写字母和数字，最长32位。
	//	商户系统按需传入，如果同一用户在同一产品码、同一签约场景下，签订了多份代扣协议，那么需要指定并传入该值。
	//【示例值】test20190701
	ExternalAgreementNo string `json:"external_agreement_no,omitempty"`
	//【描述】用户在商户网站的登录账号，用于在签约页面展示，如果为空，则不展示
	//【示例值】13888888888
	ExternalLogonId string `json:"external_logon_id"`
	//【描述】签约成功后商户用于接收异步通知的地址。
	//	如果不传入，签约与支付的异步通知都会发到外层notify_url参数传入的地址；
	//	如果外层也未传入，签约与支付的异步通知都会发到商户appid配置的网关地址。
	//【示例值】http://www.merchant.com/receiveSignNotify
	SignNotifyUrl string `json:"sign_notify_url"`
}

// AgreementParams
// 【描述】代扣信息。
//
//	代扣业务需要传入的协议相关信息，使用本参数传入协议号后scene和auth_code不需要再传值。
type AgreementParams struct {
	//【描述】支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ）
	//【示例值】20170322450983769228
	AgreementNo string `json:"agreement_no,omitempty"`
	//【描述】鉴权确认码，在需要做支付鉴权校验时，该参数不能为空
	//【示例值】423979
	AuthConfirmNo string `json:"auth_confirm_no,omitempty"`
	//【描述】鉴权申请token，其格式和内容，由支付宝定义。在需要做支付鉴权校验时，该参数不能为空。
	//【示例值】MDEDUCT0068292ca377d1d44b65fa24ec9cd89132f
	ApplyToken string `json:"apply_token,omitempty"`
}

// PayParams
// 【描述】支付相关参数
type PayParams struct {
	//【描述】普通异步支付, 传入该参数时，如果满足受理条件，会先同步受理支付，然后在异步调度推进支付
	//	NORMAL_ASYNC: 普通异步，受理成功之后，会在交易关单之前通过一定的策略重试
	//	NEAR_REAL_TIME_ASYNC: 准实时异步，受理成功之后，会准实时发起1次调度
	//【示例值】NORMAL_ASYNC
	AsyncType string `json:"async_type,omitempty"`
}

// PromoParams
// 【描述】优惠明细参数，通过此属性补充营销参数
type PromoParams struct {
	//【描述】存在延迟扣款这一类的场景，用这个时间表明用户发生交易的时间，
	//	比如说，在公交地铁场景，用户刷码出站的时间，和商户上送交易的时间是不一样的。
	//【示例值】2018-09-25 22:47:33
	ActualOrderTime string `json:"actual_order_time"`
}

type TradePay struct {
	//【描述】商户网站唯一订单号。
	//	由商家自定义，64个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复。
	//【示例值】70501111111S001111119
	OutTradeNo string `json:"out_trade_no"`
	//【描述】订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]，金额不能为0
	//【示例值】9.00
	TotalAmount string `json:"total_amount"`
	//【描述】订单标题。
	//	注意：不可使用特殊字符，如 /，=，& 等。
	//【示例值】大乐透
	Subject string `json:"subject"`
	//【描述】销售产品码，商家和支付宝签约的产品码
	//【示例值】QUICK_MSECURITY_PAY
	ProductCode string `json:"product_code,omitempty"`
	//【描述】PC扫码支付的方式。
	//	支持前置模式和跳转模式。
	//	前置模式是将二维码前置到商户的订单确认页的模式。需要商户在自己的页面中以 iframe 方式请求支付宝页面。
	//	具体支持的枚举值有以下几种：
	//		0：订单码-简约前置模式，对应 iframe 宽度不能小于600px，高度不能小于300px；
	//		1：订单码-前置模式，对应iframe 宽度不能小于 300px，高度不能小于600px；
	//		3：订单码-迷你前置模式，对应 iframe 宽度不能小于 75px，高度不能小于75px；
	//		4：订单码-可定义宽度的嵌入式二维码，商户可根据需要设定二维码的大小。
	//	跳转模式下，用户的扫码界面是由支付宝生成的，不在商户的域名下。
	//	支持传入的枚举值有：
	//		2：订单码-跳转模式
	//【枚举值】
	//	订单码-简约前置模式: 0
	//	订单码-前置模式: 1
	//	订单码-迷你前置模式: 3
	//	订单码-可定义宽度的嵌入式二维码: 4
	//【示例值】1
	QrPayMode string `json:"qr_pay_mode,omitempty"`
	//【描述】商户自定义二维码宽度。
	//	注：qr_pay_mode=4时该参数有效
	//【示例值】100
	QrcodeWidth int64 `json:"qrcode_width,omitempty"`
	//【描述】订单包含的商品列表信息，json格式，其它说明详见商品明细说明
	GoodsDetail []*GoodsDetail `json:"goods_detail,omitempty"`
	//【描述】绝对超时时间，格式为yyyy-MM-dd HH:mm:ss
	//【示例值】2016-12-31 10:05:00
	TimeExpire string `json:"time_expire,omitempty"`
	//【描述】二级商户信息。
	//	直付通模式和机构间连模式下必传，其它场景下不需要传入。
	SubMerchant *SubMerchant `json:"sub_merchant,omitempty"`
	//【描述】业务扩展参数
	ExtendParams *ExtendParams `json:"extend_params,omitempty"`
	//【描述】商户传入业务信息，具体值要和支付宝约定，应用于安全，营销等参数直传场景，格式为json格式
	//【示例值】{"mc_create_trade_ip":"127.0.0.1"}
	BusinessParams string `json:"business_params,omitempty"`
	//【描述】请求后页面的集成方式。
	//	枚举值：ALIAPP：支付宝钱包内;PCWEB：PC端访问;默认值为PCWEB。
	//【枚举值】
	//	支付宝钱包内: ALIAPP
	//	PC端访问: PCWEB
	//【示例值】PCWEB
	IntegrationType string `json:"integration_type,omitempty"`
	//【描述】请求来源地址。如果使用ALIAPP的集成方式，用户中途取消支付会返回该地址。
	//【示例值】https://
	RequestFromUrl string `json:"request_from_url"`
	//【描述】公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。支付宝只会在同步返回（包括跳转回商户网站）和异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝。
	//【示例值】merchantBizType%3d3C%26merchantBizNo%3d2016010101111
	PassbackParams string `json:"passback_params,omitempty"`
	//【描述】商户原始订单号，最大长度限制32位
	//【示例值】20161008001
	MerchantOrderNo string `json:"merchant_order_no,omitempty"`
	//【描述】外部指定买家
	ExtUserInfo *ExtUserInfo `json:"ext_user_info,omitempty"`
	//【描述】返回参数选项。 商户通过传递该参数来定制同步需要额外返回的信息字段，数组格式。包括但不限于：["hyb_amount","enterprise_pay_info"]
	//【枚举值】
	//	惠营宝回票金额: hyb_amount
	//	因公付支付信息: enterprise_pay_info
	//	医保信息: medical_insurance_info
	//【示例值】["hyb_amount","enterprise_pay_info"]
	QueryOptions []QueryOption `json:"query_options,omitempty"`
	//【描述】卖家支付宝用户ID。
	//	当需要指定收款账号时，通过该参数传入，如果该值为空，则默认为商户签约账号对应的支付宝用户ID。
	//	收款账号优先级规则：门店绑定的收款账户>请求传入的seller_id>商户签约账号对应的支付宝用户ID；
	//	注：直付通和机构间联场景下seller_id无需传入或者保持跟pid一致；
	//	如果传入的seller_id与pid不一致，需要联系支付宝小二配置收款关系；
	//	支付宝预授权和新当面资金授权场景下必填。
	//【示例值】2088102146225135
	SellerId string `json:"seller_id,omitempty"`
	//【描述】代扣信息。
	//	代扣业务需要传入的协议相关信息，使用本参数传入协议号后scene和auth_code不需要再传值。
	AgreementParams *AgreementParams `json:"agreement_params,omitempty"`
	//【描述】签约参数。
	//	如果希望在sdk中支付并签约，需要在这里传入签约信息。
	//	周期扣款场景 product_code 为 CYCLE_PAY_AUTH 时必填。
	AgreementSignParams *SignParams `json:"agreement_sign_params,omitempty"`
	//【描述】支付相关参数
	PayParams *PayParams `json:"pay_params,omitempty"`
	//【描述】优惠明细参数，通过此属性补充营销参数
	//【描述】优惠参数。为 JSON 格式。注：仅与支付宝协商后可用
	//【示例值】{"storeIdType":"1"}
	PromoParams *PromoParams `json:"promo_params,omitempty"`
	//【描述】开票信息
	InvoiceInfo *InvoiceInfo `json:"invoice_info,omitempty"`
}

func (p *TradePay) ToString() string {
	marshal, _ := json.Marshal(p)
	return string(marshal)
}

// FundBill
// 【描述】交易支付使用的资金渠道。
//
//	只有在签约中指定需要返回资金明细，或者入参的query_options中指定时才返回该字段信息。
type FundBill struct {
	//【描述】交易使用的资金渠道，
	//	详见 支付渠道列表:https://doc.open.alipay.com/doc2/detail?treeId=26&articleId=103259&docType=1
	//【示例值】ALIPAYACCOUNT
	FundChannel string `json:"fund_channel"`
	//【描述】该支付工具类型所使用的金额
	//【示例值】10
	Amount string `json:"amount"`
	//【描述】渠道实际付款金额
	//【示例值】11.21
	RealAmount string `json:"real_amount"`
}

// VoucherDetail
// 【描述】本交易支付时使用的所有优惠券信息。
// 只有在query_options中指定时才返回该字段信息。
type VoucherDetail struct {
	//【描述】券id
	//【示例值】2015102600073002039000002D5O
	Id string `json:"id"`
	//【描述】券名称
	//【示例值】XX超市5折优惠
	Name string `json:"name"`
	//【描述】券类型，如：
	//	ALIPAY_FIX_VOUCHER - 全场代金券
	//	ALIPAY_DISCOUNT_VOUCHER - 折扣券
	//	ALIPAY_ITEM_VOUCHER - 单品优惠券
	//	ALIPAY_CASH_VOUCHER - 现金抵价券
	//	ALIPAY_BIZ_VOUCHER - 商家全场券
	//	注：不排除将来新增其他类型的可能，商家接入时注意兼容性避免硬编码
	//【枚举值】
	//	全场代金券: ALIPAY_FIX_VOUCHER
	//	折扣券: ALIPAY_DISCOUNT_VOUCHER
	//	单品优惠券: ALIPAY_ITEM_VOUCHER
	//	现金抵价券: ALIPAY_CASH_VOUCHER
	//	商家全场券: ALIPAY_BIZ_VOUCHER
	//【注意事项】不排除将来新增其他类型的可能，商家接入时注意兼容性避免硬编码
	//【示例值】ALIPAY_FIX_VOUCHER
	Type string `json:"type"`
	//【描述】优惠券面额，它应该会等于商家出资加上其他出资方出资
	//【示例值】10.00
	Amount string `json:"amount"`
	//【描述】商家出资（特指发起交易的商家出资金额）
	//【示例值】9.00
	MerchantContribute string `json:"merchant_contribute"`
	//【描述】其他出资方出资金额，可能是支付宝，可能是品牌商，或者其他方，也可能是他们的一起出资
	//【示例值】1.00
	OtherContribute string `json:"other_contribute"`
	//【描述】优惠券备注信息
	//【示例值】学生专用优惠
	Memo string `json:"memo"`
	//【描述】券模板id
	//【示例值】20171030000730015359000EMZP0
	TemplateId string `json:"template_id"`
	//【描述】如果使用的这张券是用户购买的，则该字段代表用户在购买这张券时用户实际付款的金额
	//【示例值】2.01
	PurchaseBuyerContribute string `json:"purchase_buyer_contribute"`
	//【描述】如果使用的这张券是用户购买的，则该字段代表用户在购买这张券时商户优惠的金额
	//【示例值】1.03
	PurchaseMerchantContribute string `json:"purchase_merchant_contribute"`
	//【描述】如果使用的这张券是用户购买的，则该字段代表用户在购买这张券时平台优惠的金额
	//【示例值】0.82
	PurchaseAntContribute string `json:"purchase_ant_contribute"`
}

type AlipayTradePayResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	//【描述】支付宝交易号
	//【注意事项】在未生成真实交易时，不返回，需要商户调用交易查询接口或接入支付通知，获取最终的交易号
	//【示例值】2013112011001004330000121536
	TradeNo string `json:"trade_no"`
	//【描述】商户订单号
	//【示例值】6823789339978248
	OutTradeNo string `json:"out_trade_no"`
	//【描述】买家支付宝账号
	//【注意事项】在未生成真实交易时，不返回，需要商户调用交易查询接口或接入支付通知，获取最终的用户信息
	//【示例值】159****5620
	BuyerLogonId string `json:"buyer_logon_id"`
	//【描述】交易金额
	//【示例值】120.88
	TotalAmount string `json:"total_amount"`
	//【描述】实收金额
	//【示例值】88.88
	ReceiptAmount string `json:"receipt_amount"`
	//【描述】买家付款的金额
	//【示例值】8.88
	BuyerPayAmount string `json:"buyer_pay_amount"`
	//【描述】使用集分宝付款的金额
	//【示例值】8.12
	PointAmount string `json:"point_amount"`
	//【描述】交易中可给用户开具发票的金额
	//【示例值】12.50
	InvoiceAmount string `json:"invoice_amount"`
	//【描述】交易支付时间
	//【示例值】2014-11-27 15:45:57
	GmtPayment string `json:"gmt_payment"`
	//【描述】交易支付使用的资金渠道。
	//	只有在签约中指定需要返回资金明细，或者入参的query_options中指定时才返回该字段信息。
	FundBillList []*FundBill `json:"fund_bill_list"`
	//【描述】发生支付交易的商户门店名称
	//【示例值】证大五道口店
	StoreName string `json:"store_name"`
	//【描述】本次交易支付所使用的单品券优惠的商品优惠信息。
	//	只有在query_options中指定时才返回该字段信息。
	//【示例值】[{"goods_id":"STANDARD1026181538","goods_name":"雪碧","discount_amount":"100.00","voucher_id":"2015102600073002039000002D5O"}]
	DiscountGoodsDetail string `json:"discount_goods_detail"`
	//【描述】买家支付宝用户唯一标识
	//	详情可查看 openid简介:https://opendocs.alipay.com/mini/0ai2i6?pathHash=13dd5946
	//【示例值】074a1CcTG1LelxKe4xQC0zgNdId0nxi95b5lsNpazWYoCo5
	BuyerOpenId string `json:"buyer_open_id"`
	//【描述】买家在支付宝的用户id新商户建议使用buyer_open_id替代该字段。
	//	对于新商户，buyer_user_id字段未来计划逐步回收，存量商户可继续使用。
	//	如使用buyer_open_id，请确认 应用-开发配置-openid配置管理 已启用。
	//	无该配置项，可查看openid配置申请:https://opendocs.alipay.com/mini/0ai9ok?pathHash=de631c06
	//【注意事项】在未生成真实交易时，不返回，需要商户调用交易查询接口或接入支付通知，获取最终的用户信息
	//【示例值】2088101117955611
	BuyerUserId string `json:"buyer_user_id"`
	//【描述】异步支付模式，目前有五种值：
	//	ASYNC_DELAY_PAY(异步延时付款);
	//	ASYNC_REALTIME_PAY(异步准实时付款);
	//	SYNC_DIRECT_PAY(同步直接扣款);
	//	NORMAL_ASYNC_PAY(纯异步付款);
	//	QUOTA_OCCUPYIED_ASYNC_PAY(异步支付并且预占了先享后付额度);
	//【枚举值】
	//	异步延时付款: ASYNC_DELAY_PAY
	//	异步准实时付款: ASYNC_REALTIME_PAY
	//	同步直接扣款: SYNC_DIRECT_PAY
	//	纯异步付款: NORMAL_ASYNC_PAY
	//	异步支付并且预占了先享后付额度: QUOTA_OCCUPYIED_ASYNC_PAY
	//【示例值】SYNC_DIRECT_PAY
	AsyncPaymentMode string `json:"async_payment_mode"`
	//【描述】本交易支付时使用的所有优惠券信息。
	//	只有在query_options中指定时才返回该字段信息。
	VoucherDetailList []*VoucherDetail `json:"voucher_detail_list"`
	//【描述】先享后付2.0垫资金额,不返回表示没有走垫资，非空表示垫资支付的金额
	//【示例值】88.8
	AdvanceAmount string `json:"advance_amount"`
	//【描述】费率活动标识，当交易享受活动优惠费率时，返回该活动的标识；
	//	只在机构间联模式下返回，其它场景下不返回该字段；
	//	可能的返回值列表：
	//	bluesea_1：蓝海活动标识;
	//	industry_special_00：行业特殊费率0；
	//	industry_special_01：行业特殊费率1；
	//【示例值】industry_special_00
	ChargeFlags string `json:"charge_flags"`
	//【描述】商家优惠金额
	//【示例值】88.88
	MdiscountAmount string `json:"mdiscount_amount"`
	//【描述】平台优惠金额
	//【示例值】88.88
	DiscountAmount string `json:"discount_amount"`
}

type TradePayResponse struct {
	PublicResponseParameters
	Response *AlipayTradePayResponse `json:"alipay_trade_pay_response"`
}
