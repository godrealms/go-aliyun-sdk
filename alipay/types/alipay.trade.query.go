package types

import "encoding/json"

// TradeQuery 统一收单交易查询
type TradeQuery struct {
	//【描述】订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no
	//【示例值】20150320010101001
	OutTradeNo string `json:"out_trade_no,omitempty"`
	//【描述】支付宝交易号，和商户订单号不能同时为空
	//【示例值】2014112611001004680 073956707
	TradeNo string `json:"trade_no,omitempty"`
	//【描述】银行间联模式下有用，其它场景请不要使用； 双联通过该参数指定需要查询的交易所属收单机构的pid;
	//【示例值】2088101117952222
	OrgPid string `json:"org_pid,omitempty"`
	//【描述】查询选项，商户传入该参数可定制本接口同步响应额外返回的信息字段，数组格式。
	//【枚举值】
	//	交易结算信息: trade_settle_info
	//	交易支付使用的资金渠道: fund_bill_list
	//	交易支付时使用的所有优惠券信息: voucher_detail_list
	//	交易支付使用单品券优惠的商品优惠信息: discount_goods_detail
	//	商家优惠金额: mdiscount_amount
	//	医保信息: medical_insurance_info
	//	碰一下支付信息: tap_pay_info
	//【示例值】["trade_settle_info"]
	QueryOptions []QueryOption `json:"query_options"`
}

func (r *TradeQuery) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// TradeSettleDetail
// 【描述】交易结算明细信息
type TradeSettleDetail struct {
	//【描述】结算操作类型。有以下几种类型： replenish(补差)、replenish_refund(退补差)、transfer(分账)、transfer_refund(退分账)、settle(结算)、settle_refund(退结算)、on_settle(待结算)。
	//【枚举值】
	//	分账: transfer
	//	退分账: transfer_refund
	//	补差: replenish
	//	退补差: replenish_refund
	//	结算: settle
	//	退结算: settle_refund
	//	待结算: on_settle
	//【示例值】replenish
	OperationType string `json:"operation_type"`
	//【描述】商户操作序列号。商户发起请求的外部请求号。
	//【示例值】2321232323232
	OperationSerialNo string `json:"operation_serial_no"`
	//【描述】操作日期
	//【示例值】2019-05-16 09:59:17
	OperationDt string `json:"operation_dt"`
	//【描述】转出账号
	//【示例值】2088111111111111
	TransOut string `json:"trans_out"`
	//【描述】转入账号
	//【示例值】2088111111111111
	TransIn string `json:"trans_in"`
	//【描述】实际操作金额，单位为元，两位小数。该参数的值为分账或补差或结算时传入
	//【示例值】10.00
	Amount string `json:"amount"`
	//【描述】商户请求的转出账号
	//【示例值】2088111111111111
	OriTransOut string `json:"ori_trans_out"`
	//【描述】商户请求的转入账号
	//【示例值】2088111111111111
	OriTransIn string `json:"ori_trans_in"`
}

// TradeSettleInfo
// 【描述】返回的交易结算信息，包含分账、补差等信息。
//
//	只有在query_options中指定时才返回该字段信息。
type TradeSettleInfo struct {
	//【描述】交易结算明细信息
	TradeSettleDetailList []TradeSettleDetail `json:"trade_settle_detail_list"`
	//【描述】直付通账期、直连账期下返回，其他场景为空，表示一笔订单剩余待结算金额
	//【示例值】1.1
	TradeUnsettledAmount string `json:"trade_unsettled_amount"`
}

// HbFqPayInfo
// 【描述】若用户使用花呗分期支付，且商家开通返回此通知参数，则会返回花呗分期信息。
//
//	json格式其它说明详见花呗分期信息说明。
//	注意：商家需与支付宝约定后才返回本参数。
type HbFqPayInfo struct {
	//【描述】用户使用花呗分期支付的分期数
	//【示例值】3
	UserInstallNum string `json:"user_install_num"`
}

// SubFeeDetail
// 【描述】组合支付收费明细
type SubFeeDetail struct {
	//【描述】实收费用。单位：元。
	//【示例值】0.10
	ChargeFee string `json:"charge_fee"`
	//【描述】原始费用。单位：元。
	//【示例值】0.20
	OriginalChargeFee string `json:"original_charge_fee"`
	//【描述】签约费率
	//【示例值】0.03
	SwitchFeeRate string `json:"switch_fee_rate"`
}

// ChargeInfo
// 【描述】计费信息列表
type ChargeInfo struct {
	//【描述】实收费用。单位：元。
	//【示例值】0.01
	ChargeFee string `json:"charge_fee"`
	//【描述】原始费用。单位：元。
	//【示例值】0.01
	OriginalChargeFee string `json:"original_charge_fee"`
	//【描述】签约费率
	//【示例值】0.03
	SwitchFeeRate string `json:"switch_fee_rate"`
	//【描述】是否收款账号出资，值为"Y"或"N"
	//【示例值】Y
	IsRatingOnTradeReceiver string `json:"is_rating_on_trade_receiver"`
	//【描述】是否合约指定收费账号出资，值为"Y"或"N"
	//【示例值】Y
	IsRatingOnSwitch string `json:"is_rating_on_switch"`
	//【描述】收单手续费trade，花呗分期手续hbfq，其他手续费charge
	//【示例值】trade
	ChargeType string `json:"charge_type"`
	//【描述】组合支付收费明细
	SubFeeDetailList []*SubFeeDetail `json:"sub_fee_detail_list"`
}

// ReqGoodsDetail
// 【描述】支付请求的商品明细列表
type ReqGoodsDetail struct {
	//【描述】商品的编号，该参数传入支付券上绑定商品goods_id, 倘若无支付券需要消费，该字段传入商品最小粒度的商品ID（如：若商品有sku粒度，则传商户sku粒度的ID）
	//【示例值】apple-01
	GoodsId string `json:"goods_id"`
	//【描述】支付宝定义的统一商品编号
	//【示例值】20010001
	AlipayGoodsId string `json:"alipay_goods_id"`
	//【描述】商品名称
	//【示例值】ipad
	GoodsName string `json:"goods_name"`
	//【描述】商品数量
	//【示例值】1
	Quantity int `json:"quantity"`
	//【描述】商品单价，单位为元
	//【示例值】2000
	Price string `json:"price"`
	//【描述】商品类目
	//【示例值】34543238
	GoodsCategory string `json:"goods_category"`
	//【描述】商品类目树，从商品类目根节点到叶子节点的类目id组成，类目id值使用|分割
	//【示例值】124868003|126232002|126252004
	CategoriesTree string `json:"categories_tree"`
	//【描述】商品描述信息
	//【示例值】特价手机
	Body string `json:"body"`
	//【描述】商品的展示地址
	//【示例值】http://www.alipay.com/xxx.jpg
	ShowUrl string `json:"show_url"`
	//【描述】商家侧小程序商品ID，指商家提报给小程序商品库的商品。
	//	当前接口的extend_params.trade_component_order_id字段不为空时该字段必填，
	//	且与交易组件订单参数保持一致。
	//	了解小程序商品请参考：https://opendocs.alipay.com/mini/06uila?pathHash=63b6fba7
	//【示例值】outItem_01
	OutItemId string `json:"out_item_id"`
	//【描述】商家侧小程序商品ID，指商家提报给小程序商品库的商品。
	//	当前接口的extend_params.trade_component_order_id字段不为空时该字段必填，
	//	且与交易组件订单参数保持一致。
	//	了解小程序商品请参考：https://opendocs.alipay.com/mini/06uila?pathHash=63b6fba7
	//【示例值】outSku_01
	OutSkuId string `json:"out_sku_id"`
}

// FulfillmentDetail
// 【描述】履约详情列表。
//
//	只有入参的query_options中指定fulfillment_detail_list并且所查询的交易存在履约明细时才返回该字段信息。
//
// 【必选条件】履约详情列表。
//
//	只有入参的query_options中指定fulfillment_detail_list并且所查询的交易存在履约明细时才返回该字段信息。
type FulfillmentDetail struct {
	//【描述】履约金额
	//【示例值】80.00
	FulfillmentAmount string `json:"fulfillment_amount"`
	//【描述】商户发起履约请求时，传入的out_request_no，标识一次请求的唯一id
	//【示例值】20200320010101001
	OutRequestNo string `json:"out_request_no"`
	//【描述】履约支付时间
	//【示例值】2021-03-17 09:45:57
	GmtPayment string `json:"gmt_payment"`
}

// TapPayInfo
// 【描述】碰一下支付信息
type TapPayInfo struct {
	//【描述】碰一下支付的支付介质类型，标识当前支付为碰一下支付
	//【示例值】TAP_PAY
	PaymentMediumType string `json:"payment_medium_type"`
}

// BkAgentRespInfo
// 【描述】间联交易下，返回给机构的信 息
type BkAgentRespInfo struct {
	//【描述】原快捷交易流水号
	//【示例值】123412341234
	BindtrxId string `json:"bindtrx_id"`
	//【描述】枚举值，01 银联；02 网联；03 连通等
	//【示例值】01
	BindclrissrId string `json:"bindclrissr_id"`
	//【描述】付款机构在清算组织登记或分配的机构代码
	//【示例值】123123123123
	BindpyeracctbkId string `json:"bindpyeracctbk_id"`
	//【描述】用户在银行付款账号的标记化处理编号
	//【示例值】123451234512345
	BkpyeruserCode string `json:"bkpyeruser_code"`
	//【描述】设备推测位置
	//【示例值】+37.28/-121.268
	EstterLocation string `json:"estter_location"`
}

type TradeQueryResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	//【描述】支付宝交易号，和商户订单号不能同时为空
	//【示例值】2014112611001004680 073956707
	TradeNo string `json:"trade_no"`
	//【描述】商家订单号
	//【示例值】6823789339978248
	OutTradeNo string `json:"out_trade_no"`
	//【描述】买家支付宝账号
	//【注意事项】在未生成真实交易时，不返回，需要商户多次调用该接口或支付通知，获取最终的用户信息
	//【示例值】159****5620
	BuyerLogonId string `json:"buyer_logon_id"`
	//【描述】交易状态：
	//	WAIT_BUYER_PAY（交易创建，等待买家付款）、
	//	TRADE_CLOSED（未付款交易超时关闭，或支付完成后全额退款）、
	//	TRADE_SUCCESS（交易支付成功）、
	//	TRADE_FINISHED（交易结束，不可退款）
	//【示例值】TRADE_CLOSED
	TradeStatus string `json:"trade_status"`
	//【描述】交易附加状态： SELLER_NOT_RECEIVED（买家已付款，卖家未收款）；
	//【示例值】SELLER_NOT_RECEIVED
	AdditionalStatus string `json:"additional_status"`
	//【描述】交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount
	//【示例值】88.88
	TotalAmount string `json:"total_amount"`
	//【描述】标价币种，该参数的值为支付时传入的trans_currency，
	//	英镑：GBP、
	//	港币：HKD、
	//	美元：USD、
	//	新加坡元：SGD、
	//	日元：JPY、
	//	加拿大元：CAD、
	//	澳元：AUD、
	//	欧元：EUR、
	//	新西兰元：NZD、
	//	韩元：KRW、
	//	泰铢：THB、
	//	瑞士法郎：CHF、
	//	瑞典克朗：SEK、
	//	丹麦克朗：DKK、
	//	挪威克朗：NOK、
	//	马来西亚林吉特：MYR、
	//	印尼卢比：IDR、
	//	菲律宾比索：PHP、
	//	毛里求斯卢比：MUR、
	//	以色列新谢克尔：ILS、
	//	斯里兰卡卢比：LKR、
	//	俄罗斯卢布：RUB、
	//	阿联酋迪拉姆：AED、
	//	捷克克朗：CZK、
	//	南非兰特：ZAR、
	//	人民币：CNY、
	//	新台币：TWD。当trans_currency 和 settle_currency 不一致时，trans_currency支持人民币：CNY、新台币：TWD
	//【示例值】TWD
	TransCurrency string `json:"trans_currency"`
	//【描述】订单结算币种，对应支付接口传入的settle_currency，
	//	英镑：GBP、
	//	港币：HKD、
	//	美元：USD、
	//	新加坡元：SGD、
	//	日元：JPY、
	//	加拿大元：CAD、
	//	澳元：AUD、
	//	欧元：EUR、
	//	新西兰元：NZD、
	//	韩元：KRW、
	//	泰铢：THB、
	//	瑞士法郎：CHF、
	//	瑞典克朗：SEK、
	//	丹麦克朗：DKK、
	//	挪威克朗：NOK、
	//	马来西亚林吉特：MYR、
	//	印尼卢比：IDR、
	//	菲律宾比索：PHP、
	//	毛里求斯卢比：MUR、
	//	以色列新谢克尔：ILS、
	//	斯里兰卡卢比：LKR、
	//	俄罗斯卢布：RUB、
	//	阿联酋迪拉姆：AED、
	//	捷克克朗：CZK、
	//	南非兰特：ZAR
	//【示例值】USD
	SettleCurrency string `json:"settle_currency"`
	//【描述】结算币种订单金额
	//【示例值】2.96
	SettleAmount string `json:"settle_amount"`
	//【描述】订单支付币种
	//【示例值】CNY
	PayCurrency string `json:"pay_currency"`
	//【描述】支付币种订单金额
	//【示例值】8.88
	PayAmount string `json:"pay_amount"`
	//【描述】结算币种兑换标价币种汇率
	//【示例值】30.025
	SettleTransRate string `json:"settle_trans_rate"`
	//【描述】标价币种兑换支付币种汇率
	//【示例值】0.264
	TransPayRate string `json:"trans_pay_rate"`
	//【描述】买家实付金额，单位为元，两位小数。
	//	该金额代表该笔交易买家实际支付的金额，不包含商户折扣等金额
	//【示例值】8.88
	BuyerPayAmount string `json:"buyer_pay_amount"`
	//【描述】积分支付的金额，单位为元，两位小数。
	//	该金额代表该笔交易中用户使用积分支付的金额，比如集分宝或者支付宝实时优惠等
	//【示例值】10
	PointAmount string `json:"point_amount"`
	//【描述】交易中用户支付的可开具发票的金额，单位为元，两位小数。
	//	该金额代表该笔交易中可以给用户开具发票的金额
	//【示例值】12.11
	InvoiceAmount string `json:"invoice_amount"`
	//【描述】本次交易打款给卖家的时间
	//【示例值】2014-11-27 15:45:57
	SendPayDate string `json:"send_pay_date"`
	//【描述】实收金额，单位为元，两位小数。该金额为本笔交易，商户账户能够实际收到的金额
	//【示例值】15.25
	ReceiptAmount string `json:"receipt_amount"`
	//【描述】商户门店编号
	//【示例值】NJ_S_001
	StoreId string `json:"store_id"`
	//【描述】商户机具终端编号
	//【示例值】NJ_T_001
	TerminalId string `json:"terminal_id"`
	//【描述】交易支付使用的资金渠道。 只有在签约中指定需要返回资金明细，或者入参的query_options中指定时才返回该字段信息。
	FundBillList []*FundBill `json:"fund_bill_list"`
	//【描述】请求交易支付中的商户店铺的名称
	//【示例值】证大五道口店
	StoreName string `json:"store_name"`
	//【描述】买家在支付宝的用户id新商户建议使用buyer_open_id替代该字段。
	//	对于新商户，buyer_user_id字段未来计划逐步回收，存量商户可继续使用。
	//	如使用buyer_open_id，请确认 应用-开发配置-openid配置管理 已启用。
	//	无该配置项，可查看openid配置申请：https://opendocs.alipay.com/mini/0ai9ok?pathHash=de631c06。
	//【示例值】2088101117955611
	BuyerUserId string `json:"buyer_user_id"`
	//【描述】买家支付宝用户唯一标识
	//	详情可查看 openid简介: https://opendocs.alipay.com/mini/0ai2i6?pathHash=13dd5946
	//【示例值】01501o8f93I3nJAGB1jG4ONxtxV25DCN3Gec3uggnC4CJU0
	BuyerOpenId string `json:"buyer_open_id"`
	//【描述】行业特殊信息-统筹相关
	//【示例值】{"registration_order_pay":{"brlx":"1","cblx":"1"}}
	IndustrySepcDetailGov string `json:"industry_sepc_detail_gov"`
	//【描述】行业特殊信息-个账相关
	//【示例值】{"registration_order_pay":{"brlx":"1","cblx":"1"}}
	IndustrySepcDetailAcc string `json:"industry_sepc_detail_acc"`
	//【描述】该笔交易针对收款方的收费金额；单位：元。 只在银行间联交易场景下返回该信息；
	//【示例值】8.88
	ChargeAmount string `json:"charge_amount"`
	//【描述】费率活动标识。
	//	当交易享受特殊行业或活动费率时，返回该场景的标识。
	//	具体场景如下：
	//		trade_special_00：订单优惠费率；
	//		industry_special_on_00：线上行业特殊费率0；
	//		industry_special_on_01：线上行业特殊费率1；
	//		industry_special_00：线下行业特殊费率0；
	//		industry_special_01：线下行业特殊费率1；
	//		bluesea_1：蓝海活动优惠费率标签；
	//		注：只在机构间联模式下返回，其它场景下不返回该字段；
	//【示例值】bluesea_1
	ChargeFlags string `json:"charge_flags"`
	//【描述】支付清算编号，用于清算对账使用； 只在银行间联交易场景下返回该信息；
	//【示例值】2018101610032004620239146945
	SettlementId string `json:"settlement_id"`
	//【描述】返回的交易结算信息，包含分账、补差等信息。 只有在query_options中指定时才返回该字段信息。
	TradeSettleInfo *TradeSettleInfo `json:"trade_settle_info"`
	//【描述】预授权支付模式，该参数仅在信用预授权支付场景下返回。信用预授权支付：CREDIT_PREAUTH_PAY
	//【枚举值】
	//	信用预授权支付: CREDIT_PREAUTH_PAY
	//【示例值】CREDIT_PREAUTH_PAY
	AuthTradePayMode string `json:"auth_trade_pay_mode"`
	//【描述】平台优惠金额。单位：元。
	//【示例值】88.88
	DiscountAmount string `json:"discount_amount"`
	//【描述】订单标题； 只在银行间联交易场景下返回该信息；
	//【示例值】Iphone6 16G
	Subject string `json:"subject"`
	//【描述】订单描述； 只在银行间联交易场景下返回该信息；
	//【示例值】Iphone6 16G
	Body string `json:"body"`
	//【描述】间连商户在支付宝端的商户编号； 只在银行间联交易场景下返回该信息；
	//【示例值】2088301372182171
	AlipaySubMerchantId string `json:"alipay_sub_merchant_id"`
	//【描述】交易额外信息，特殊场景下与支付宝约定返回。 json格式。
	//【示例值】{"action":"cancel"}
	ExtInfos string `json:"ext_infos"`
	//【描述】公用回传参数。 返回支付时传入的passback_params参数信息
	//【示例值】merchantBizType%3d3C%26merchantBizNo%3d2016010101111
	PassbackParams string `json:"passback_params"`
	//【描述】买家用户类型。CORPORATE:企业用户；PRIVATE:个人用户。
	//【枚举值】
	//	企业用户: CORPORATE
	//	个人用户: PRIVATE
	//【示例值】PRIVATE
	BuyerUserType string `json:"buyer_user_type"`
	//【描述】商家优惠金额。单位：元。
	//【示例值】88.88
	MdiscountAmount string `json:"mdiscount_amount"`
	//【描述】若用户使用花呗分期支付，且商家开通返回此通知参数，则会返回花呗分期信息。
	//	json格式其它说明详见花呗分期信息说明。
	//	注意：商家需与支付宝约定后才返回本参数。
	HbFqPayInfo *HbFqPayInfo `json:"hb_fq_pay_info"`
	//【描述】信用支付模式。表示订单是采用信用支付方式（支付时买家没有出资，需要后续履约）。
	//	"creditAdvanceV2"表示芝麻先用后付模式，用户后续需要履约扣款。
	//	此字段只有信用支付场景才有值，商户需要根据字段值单独处理。
	//	此字段以后可能扩展其他值，建议商户使用白名单方式识别，对于未识别的值做失败处理，并联系支付宝技术支持人员。
	//【枚举值】
	//	芝麻先用后付模式: creditAdvanceV2
	//【示例值】creditAdvanceV2
	CreditPayMode string `json:"credit_pay_mode"`
	//【描述】信用业务单号。信用支付场景才有值，先用后付产品里是芝麻订单号。
	//【示例值】ZMCB99202103310000450000041833
	CreditBizOrderId string `json:"credit_biz_order_id"`
	//【描述】惠营宝回票金额。单位：元。
	//【示例值】10.24
	HybAmount string `json:"hyb_amount"`
	//【描述】间联交易下，返回给机构的信 息
	BkAgentRespInfo *BkAgentRespInfo `json:"bkagent_resp_info"`
	//【描述】计费信息列表
	ChargeInfoList []*ChargeInfo `json:"charge_info_list"`
	//【描述】账期结算标识，指已完成支付的订单会进行账期管控，不会实时结算。
	//	该参数目前会在使用小程序交易组件场景下返回
	//【枚举值】
	//	账期模式: PERIOD
	//【示例值】PERIOD
	BizSettleMode string `json:"biz_settle_mode"`
	//【描述】支付请求的商品明细列表
	ReqGoodsDetail []*ReqGoodsDetail `json:"req_goods_detail"`
	//【描述】履约详情列表。 只有入参的query_options中指定fulfillment_detail_list并且所查询的交易存在履约明细时才返回该字段信息。
	//【必选条件】履约详情列表。 只有入参的query_options中指定fulfillment_detail_list并且所查询的交易存在履约明细时才返回该字段信息。
	FulfillmentDetailList []*FulfillmentDetail `json:"fulfillment_detail_list"`
	//【描述】该字段用于描述当前账期交易的场景。
	//【示例值】账期交易的场景。
	PeriodScene string `json:"period_scene"`
	//【描述】异步支付受理状态，仅异步支付模式且query_options指定async_pay_info时返回。
	//	S：受理成功，支付宝内部会在一定期限内捞起任务推进支付，直到支付成功或超出可重试期限；
	//	其它：受理结果未知，可重试查询。
	//【枚举值】
	//	异步支付受理成功: S
	//【示例值】S
	AsyncPayApplyStatus string `json:"async_pay_apply_status"`
	//【描述】收银台类型。 用户支付的收银台类型，取值如下：
	//	APP：支付宝APP收银台支付；
	//	WAP：支付H5收银台支付；
	//	注：只有在无线产品支付接口请求中query_options指定cashier_type才返回该字段。
	//【示例值】APP
	CashierType string `json:"cashier_type"`
	//【描述】碰一下支付信息
	TapPayInfo *TapPayInfo `json:"tap_pay_info"`
	//【描述】当用户使用芝麻信用先享后付时，会返回该字段，代表整笔交易的原始待履约金额，单位元。
	//【必选条件】当用户使用芝麻信用先享后付时，会返回该字段。
	//【示例值】12.46
	PreAuthPayAmount string `json:"pre_auth_pay_amount"`
}

type AlipayTradeQueryResponse struct {
	PublicResponseParameters
	Response *TradeQueryResponse `json:"alipay_trade_query_response"`
}
