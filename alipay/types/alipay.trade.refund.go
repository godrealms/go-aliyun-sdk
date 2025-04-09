package types

import "encoding/json"

type RefundGoodsDetail struct {
	OutSkuId             string   `json:"out_sku_id"`
	OutItemId            string   `json:"out_item_id"`
	GoodsId              string   `json:"goods_id"`
	RefundAmount         string   `json:"refund_amount"`
	OutCertificateNoList []string `json:"out_certificate_no_list"`
}

type RefundRoyaltyParameter struct {
	Amount       string `json:"amount"`
	TransIn      string `json:"trans_in"`
	RoyaltyType  string `json:"royalty_type"`
	TransOut     string `json:"trans_out"`
	TransOutType string `json:"trans_out_type"`
	RoyaltyScene string `json:"royalty_scene"`
	TransInType  string `json:"trans_in_type"`
	TransInName  string `json:"trans_in_name"`
	Desc         string `json:"desc"`
}

// TradeRefund 统一收单交易退款
type TradeRefund struct {
	OutTradeNo   string `json:"out_trade_no"`
	TradeNo      string `json:"trade_no"`
	RefundAmount string `json:"refund_amount"`
	RefundReason string `json:"refund_reason"`
	//【描述】本笔退款对应的退款请求号
	//【示例值】20150320010101001
	OutRequestNo            string                    `json:"out_request_no"`
	RefundGoodsDetail       []*RefundGoodsDetail      `json:"refund_goods_detail"`
	RefundRoyaltyParameters []*RefundRoyaltyParameter `json:"refund_royalty_parameters"`
	QueryOptions            []QueryOption             `json:"query_options"`
	RelatedSettleConfirmNo  string                    `json:"related_settle_confirm_no"`
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
	//【示例值】2014112611001004680073956707
	TradeNo string `json:"trade_no"`
	//【描述】创建交易传入的商户订单号
	//【示例值】20150320010101001
	OutTradeNo              string                 `json:"out_trade_no"`
	BuyerLogonId            string                 `json:"buyer_logon_id"`
	FundChange              string                 `json:"fund_change"`
	RefundFee               string                 `json:"refund_fee"`
	RefundDetailItemList    []*RefundDetailItem    `json:"refund_detail_item_list"`
	StoreName               string                 `json:"store_name"`
	BuyerUserId             string                 `json:"buyer_user_id"`
	BuyerOpenId             string                 `json:"buyer_open_id"`
	SendBackFee             string                 `json:"send_back_fee"`
	RefundHybAmount         string                 `json:"refund_hyb_amount"`
	RefundChargeInfoList    []*RefundChargeInfo    `json:"refund_charge_info_list"`
	RefundVoucherDetailList []*RefundVoucherDetail `json:"refund_voucher_detail_list"`
	PreAuthCancelFee        string                 `json:"pre_auth_cancel_fee"`
}

type AlipayTradeRefundResponse struct {
	PublicResponseParameters
	Response TradeRefundResponse `json:"alipay_trade_refund_response"`
}
