package types

import "encoding/json"

// RefundGoodsDetail
// 【描述】退款包含的商品列表信息
type RefundGoodsDetail struct {
	//【描述】商家侧小程序商品sku ID，对应支付时传入的out_sku_id
	//【示例值】outSku_01
	OutSkuId string `json:"out_sku_id"`
	//【描述】商家侧小程序商品ID，对应支付时传入的out_item_id
	//【示例值】outItem_01
	OutItemId string `json:"out_item_id"`
	//【描述】商品编号。 对应支付时传入的goods_id
	//【示例值】apple-01
	GoodsId string `json:"goods_id"`
	//【描述】该商品的退款总金额，单位为元
	//【示例值】19.50
	RefundAmount string `json:"refund_amount"`
	//【描述】外部商品凭证编号列表
	//【示例值】["202407013232143241231243243423"]
	OutCertificateNoList []string `json:"out_certificate_no_list"`
}

// RefundRoyaltyParameter
// 【描述】退分账明细信息。
//
//	注：
//	1.当面付且非直付通模式无需传入退分账明细，系统自动按退款金额与订单金额的比率，从收款方和分账收入方退款，不支持指定退款金额与退款方。
//	2.直付通模式，电脑网站支付，手机 APP 支付，手机网站支付产品，须在退款请求中明确是否退分账，从哪个分账收入方退，退多少分账金额；
//	如不明确，默认从收款方退款，收款方余额不足退款失败。不支持系统按比率退款。
type RefundRoyaltyParameter struct {
	//【描述】分账的金额，单位为元
	//【示例值】0.1
	Amount string `json:"amount,omitempty"`
	//【描述】分账类型.
	//【枚举值】
	//	分账: transfer
	//	营销补差: replenish
	//【注意事项】为空默认为分账transfer;
	//【示例值】transfer
	RoyaltyType string `json:"royalty_type,omitempty"`
	//【描述】可选值：达人佣金、平台服务费、技术服务费、其他
	//【示例值】达人佣金
	RoyaltyScene string `json:"royalty_scene,omitempty"`
	//【描述】支出方账户。
	//	如果支出方账户类型为userId，本参数为支出方的支付宝账号对应的支付宝唯一用户号，
	//	以2088开头的纯16位数字；如果支出方类型为loginName，本参数为支出方的支付宝登录号。
	//	泛金融类商户分账时，该字段不要上送。
	//【示例值】2088101126765726
	TransOut string `json:"trans_out,omitempty"`
	//【描述】支出方账户类型。
	//【枚举值】
	//	支付宝账号对应的支付宝唯一用户号: userId
	//	支付宝登录号: loginName
	//【注意事项】泛金融类商户分账时，该字段不要上送。
	//【示例值】userId
	TransOutType string `json:"trans_out_type,omitempty"`
	//【描述】收入方账户。
	//	如果收入方账户类型为userId，本参数为收入方的支付宝账号对应的支付宝唯一用户号，以2088开头的纯16位数字；
	//	如果收入方类型为cardAliasNo，本参数为收入方在支付宝绑定的卡编号；
	//	如果收入方类型为loginName，本参数为收入方的支付宝登录号；
	//【示例值】2088101126708402
	TransIn string `json:"trans_in,omitempty"`
	//【描述】收入方账户类型。
	//【枚举值】
	//	支付宝账号对应的支付宝唯一用户号: userId
	//	支付宝登录号: loginName
	//卡编号: cardAliasNo
	//【示例值】userId
	TransInType string `json:"trans_in_type,omitempty"`
	//【描述】分账收款方姓名，上送则进行姓名与支付宝账号的一致性校验，校验不一致则分账失败。
	//	不上送则不进行姓名校验
	//【示例值】张三
	TransInName string `json:"trans_in_name,omitempty"`
	//【描述】分账描述
	//【示例值】分账给2088101126708402
	Desc string `json:"desc,omitempty"`
}

// TradeRefund 统一收单交易退款
type TradeRefund struct {
	//【描述】商户订单号。 订单支付时传入的商户订单号，商家自定义且保证商家系统中唯一。
	//	与支付宝交易号 trade_no 不能同时为空。
	//【示例值】20150320010101001
	OutTradeNo string `json:"out_trade_no,omitempty"`
	//【描述】支付宝交易号。 和商户订单号 out_trade_no 不能同时为空，两者同时存在时，优先取值trade_no
	//【示例值】2014112611001004680073956707
	TradeNo string `json:"trade_no,omitempty"`
	//【描述】退款金额。 需要退款的金额，该金额不能大于订单金额，单位为元，支持两位小数。
	//	注：如果正向交易使用了营销，该退款金额包含营销金额，支付宝会按业务规则分配营销和买家自有资金分别退多少，
	//	默认优先退买家的自有资金。如交易总金额100元，用户支付时使用了80元自有资金和20元无资金流的营销券，商家实际收款80元。
	//	如果首次请求退款60元，则60元全部从商家收款资金扣除退回给用户自有资产；
	//	如果再请求退款40元，则从商家收款资金扣除20元退回用户资产以及把20元的营销券退回给用户（券是否可再使用取决于券的规则配置）。
	//【示例值】200.12
	RefundAmount string `json:"refund_amount"`
	//【描述】退款原因说明。 商家自定义，将在会在商户和用户的pc退款账单详情中展示
	//【示例值】正常退款
	RefundReason string `json:"refund_reason"`
	//【描述】本笔退款对应的退款请求号
	//【示例值】20150320010101001
	OutRequestNo string `json:"out_request_no"`
	//【描述】退款包含的商品列表信息
	RefundGoodsDetail []*RefundGoodsDetail `json:"refund_goods_detail"`
	//【描述】退分账明细信息。
	//	注：
	//	1.当面付且非直付通模式无需传入退分账明细，系统自动按退款金额与订单金额的比率，从收款方和分账收入方退款，不支持指定退款金额与退款方。
	//	2.直付通模式，电脑网站支付，手机 APP 支付，手机网站支付产品，须在退款请求中明确是否退分账，从哪个分账收入方退，退多少分账金额；
	//	如不明确，默认从收款方退款，收款方余额不足退款失败。不支持系统按比率退款。
	RefundRoyaltyParameters []*RefundRoyaltyParameter `json:"refund_royalty_parameters"`
	//【描述】查询选项。 商户通过上送该参数来定制同步需要额外返回的信息字段，数组格式。
	//【枚举值】
	//	本次退款使用的资金渠道: refund_detail_item_list
	//	银行卡冲退信息: deposit_back_info
	//	本次退款退的券信息: refund_voucher_detail_list
	//【示例值】["refund_detail_item_list"]
	QueryOptions []QueryOption `json:"query_options"`
	//【描述】针对账期交易，在确认结算后退款的话，需要指定确认结算时的结算单号。
	//【示例值】2024041122001495000530302869
	RelatedSettleConfirmNo string `json:"related_settle_confirm_no"`
}

func (r *TradeRefund) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// RefundDetailItem
// 【描述】本次退款使用的资金渠道；
//
//	默认不返回该信息，需要在入参的query_options中指定"refund_detail_item_list"值时才返回该字段信息。
type RefundDetailItem struct {
	//【描述】交易使用的资金渠道，
	//	详见 支付渠道列表：https://doc.open.alipay.com/doc2/detail?treeId=26&articleId=103259&docType=1
	//【示例值】ALIPAYACCOUNT
	FundChannel string `json:"fund_channel"`
	//【描述】该支付工具类型所使用的金额。单位：元。
	//【示例值】10
	Amount string `json:"amount"`
	//【描述】渠道实际付款金额。单位：元。
	//【示例值】11.21
	RealAmount string `json:"real_amount"`
	//【描述】渠道所使用的资金类型,目前只在资金渠道(fund_channel)是银行卡渠道(BANKCARD)的情况下才返回该信息
	//【枚举值】
	//	借记卡: DEBIT_CARD
	//	信用卡: CREDIT_CARD
	//	借贷合一卡: MIXED_CARD
	//【示例值】DEBIT_CARD
	FundType string `json:"fund_type"`
}

// RefundSubFeeDetail
// 【描述】组合支付退费明细
type RefundSubFeeDetail struct {
	//【描述】实退费用。单位：元。
	//【示例值】0.10
	RefundChargeFee string `json:"refund_charge_fee"`
	//【描述】签约费率
	//【示例值】0.01
	SwitchFeeRate string `json:"switch_fee_rate"`
}

// RefundChargeInfo
// 【描述】退费信息
type RefundChargeInfo struct {
	//【描述】实退费用。单位：元。
	//【示例值】0.01
	RefundChargeFee string `json:"refund_charge_fee"`
	//【描述】签约费率
	//【示例值】0.01
	SwitchFeeRate string `json:"switch_fee_rate"`
	//【描述】收单手续费trade，花呗分期手续hbfq，其他手续费charge
	//【示例值】trade
	ChargeType string `json:"charge_type"`
	//【描述】组合支付退费明细
	RefundSubFeeDetailList []*RefundSubFeeDetail `json:"refund_sub_fee_detail_list"`
}

// OtherContributeDetail
// 【描述】优惠券的其他出资方明细
type OtherContributeDetail struct {
	//【描述】出资方类型
	//【枚举值】
	//	平台出资: PLATFORM
	//	品牌商出资: BRAND
	//	商圈出资 : MALL
	//【注意事项】不排除将来新增其他类型的可能，商家接入时注意兼容性避免硬编码
	//【示例值】BRAND
	ContributeType string `json:"contribute_type"`
	//【描述】出资方金额
	//【示例值】8.00
	ContributeAmount string `json:"contribute_amount"`
}
type RefundVoucherDetail struct {
	//【描述】券id
	//【示例值】2015102600073002039000002D5O
	Id string `json:"id"`
	//【描述】券名称
	//【示例值】XX超市5折优惠
	Name string `json:"name"`
	//【描述】券类型
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
	//【描述】优惠券的其他出资方明细
	OtherContributeDetail []*OtherContributeDetail `json:"other_contribute_detail"`
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

type TradeRefundResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	//【描述】支付宝交易号
	//【示例值】2013112011001004330000121536
	TradeNo string `json:"trade_no"`
	//【描述】商户订单号
	//【示例值】6823789339978248
	OutTradeNo string `json:"out_trade_no"`
	//【描述】用户的登录id
	//【示例值】159****5620
	BuyerLogonId string `json:"buyer_logon_id"`
	//【描述】本次退款是否发生了资金变化
	//【示例值】Y
	FundChange string `json:"fund_change"`
	//【描述】退款总金额。
	//	单位：元。 指该笔交易累计已经退款成功的金额。
	//【示例值】88.88
	RefundFee string `json:"refund_fee"`
	//【描述】退款使用的资金渠道。
	//	只有在签约中指定需要返回资金明细，或者入参的query_options中指定时才返回该字段信息。
	RefundDetailItemList []*RefundDetailItem `json:"refund_detail_item_list"`
	//【描述】交易在支付时候的门店名称
	//【必选条件】交易在支付时候的门店名称
	//【示例值】望湘园联洋店
	StoreName string `json:"store_name"`
	//【描述】买家在支付宝的用户id新商户建议使用buyer_open_id替代该字段。
	//	对于新商户，buyer_user_id字段未来计划逐步回收，存量商户可继续使用。
	//	如使用buyer_open_id，请确认 应用-开发配置-openid配置管理 已启用。
	//	无该配置项，可查看openid配置申请:https://opendocs.alipay.com/mini/0ai9ok?pathHash=de631c06
	//【示例值】2088101117955611
	BuyerUserId string `json:"buyer_user_id"`
	//【描述】买家支付宝用户唯一标识
	//	详情可查看 openid简介: https://opendocs.alipay.com/mini/0ai2i6?pathHash=13dd5946
	//【示例值】074a1CcTG1LelxKe4xQC0zgNdId0nxi95b5lsNpazWYoCo5
	BuyerOpenId string `json:"buyer_open_id"`
	//【描述】本次商户实际退回金额。
	//	单位：元。
	//	说明：如需获取该值，需在入参query_options中传入 refund_detail_item_list。
	//【示例值】1.8
	SendBackFee string `json:"send_back_fee"`
	//【描述】本次请求退惠营宝金额。单位：元。
	//【示例值】10.24
	RefundHybAmount string `json:"refund_hyb_amount"`
	//【描述】退费信息
	RefundChargeInfoList []*RefundChargeInfo `json:"refund_charge_info_list"`
	//【描述】本交易支付时使用的所有优惠券信息。
	//	只有在query_options中指定了refund_voucher_detail_list时才返回该字段信息。
	RefundVoucherDetailList []*RefundVoucherDetail `json:"refund_voucher_detail_list"`
	//【描述】当用户使用芝麻信用先享后付时，且当前的操作为预授权撤销动作时，会返回该字段，代表当前撤销的预授权金额，单位元。
	//【必选条件】当用户使用芝麻信用先享后付时，且当前的操作为预授权撤销动作时，会返回该字段。
	//【示例值】12.45
	PreAuthCancelFee string `json:"pre_auth_cancel_fee"`
}

type AlipayTradeRefundResponse struct {
	PublicResponseParameters
	Response TradeRefundResponse `json:"alipay_trade_refund_response"`
}
