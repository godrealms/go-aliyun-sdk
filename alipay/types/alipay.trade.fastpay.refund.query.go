package types

import "encoding/json"

// TradeFastPayRefundQuery 统一收单交易退款查询
type TradeFastPayRefundQuery struct {
	//【描述】支付宝交易号。 和商户订单号不能同时为空
	//【示例值】2021081722001419121412730660
	TradeNo string `json:"trade_no"`
	//【描述】商户订单号。 订单支付时传入的商户订单号,和支付宝交易号不能同时为空。
	//	trade_no,out_trade_no如果同时存在优先取trade_no
	//【示例值】2014112611001004680073956707
	OutTradeNo string `json:"out_trade_no"`
	//【描述】退款请求号。 请求退款接口时，传入的退款请求号，如果在退款请求时未传入，则该值为创建交易时的商户订单号。
	//【示例值】HZ01RF001
	OutRequestNo string `json:"out_request_no"`
	//【描述】查询选项，商户通过上送该参数来定制同步需要额外返回的信息字段，数组格式。枚举支持： refund_detail_item_list：本次退款使用的资金渠道； gmt_refund_pay：退款执行成功的时间； deposit_back_info：银行卡冲退信息；
	//【枚举值】
	//	本次退款使用的资金渠道: refund_detail_item_list
	//	退款执行成功的时间: gmt_refund_pay
	//	银行卡冲退信息: deposit_back_info
	//	本次交易使用的券信息: refund_voucher_detail_list
	//【示例值】["refund_detail_item_list"]
	QueryOptions []QueryOption `json:"query_options"`
}

func (q *TradeFastPayRefundQuery) ToString() string {
	marshal, _ := json.Marshal(q)
	return string(marshal)
}

// RefundRoyalty 【描述】退分账明细信息，当前仅在直付通产品中返回。
type RefundRoyalty struct {
	//【描述】退分账金额。单位：元。
	//【示例值】10
	RefundAmount string `json:"refund_amount"`
	//【描述】分账类型. 字段为空默认为普通分账类型transfer
	//【枚举值】
	//	普通分账类型: transfer
	//	补差分账类型: replenish
	//【示例值】transfer
	RoyaltyType string `json:"royalty_type"`
	//【描述】退分账结果码
	//【示例值】SUCCESS
	ResultCode string `json:"result_code"`
	//【描述】转出人支付宝账号对应用户ID
	//【示例值】2088102210397302
	TransOut string `json:"trans_out"`
	//【描述】转出人支付宝账号
	//【示例值】alipay-test03@alipay.com
	TransOutEmail string `json:"trans_out_email"`
	//【描述】转入人支付宝账号对应用户ID
	//【示例值】2088102210397302
	TransIn string `json:"trans_in"`
	//【描述】转入人支付宝账号
	//【示例值】zen_gwen@hotmail.com
	TransInEmail string `json:"trans_in_email"`
	//【描述】商户请求的转出账号
	//【示例值】2088111111111111
	OriTransOut string `json:"ori_trans_out"`
	//【描述】商户请求的转入账号
	//【示例值】2088111111111111
	OriTransIn string `json:"ori_trans_in"`
}

// DepositBackInfo
// 【描述】银行卡冲退信息；
//
//	默认不返回该信息，需要在入参的query_options中指定"deposit_back_info"值时才返回该字段信息。
type DepositBackInfo struct {
	//【描述】是否存在银行卡冲退信息。
	//【示例值】true
	HasDepositBack string `json:"has_deposit_back"`
	//【描述】银行卡冲退状态。S-成功，F-失败，P-处理中。银行卡冲退失败，资金自动转入用户支付宝余额。
	//【示例值】S
	DbackStatus string `json:"dback_status"`
	//【描述】银行卡冲退金额。单位：元。
	//【示例值】1.01
	DbackAmount string `json:"dback_amount"`
	//【描述】银行响应时间，格式为yyyy-MM-dd HH:mm:ss
	//【示例值】2020-06-02 14:03:48
	BankAckTime string `json:"bank_ack_time"`
	//【描述】预估银行到账时间，格式为yyyy-MM-dd HH:mm:ss
	//【示例值】2020-06-02 14:03:48
	EstBankReceiptTime string `json:"est_bank_receipt_time"`
}

type TradeFastPayRefundQueryResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	//【描述】支付宝交易号
	//【示例值】2014112611001004680073956707
	TradeNo string `json:"trade_no"`
	//【描述】创建交易传入的商户订单号
	//【示例值】20150320010101001
	OutTradeNo string `json:"out_trade_no"`
	//【描述】本笔退款对应的退款请求号
	//【示例值】20150320010101001
	OutRequestNo string `json:"out_request_no"`
	//【描述】该笔退款所对应的交易的订单金额。单位：元。
	//【示例值】100.20
	TotalAmount string `json:"total_amount"`
	//【描述】本次退款请求，对应的退款金额。单位：元。
	//【示例值】12.33
	RefundAmount string `json:"refund_amount"`
	//【描述】退款状态。枚举值： REFUND_SUCCESS 退款处理成功；未返回该字段表示退款请求未收到或者退款失败；
	//	注：如果退款查询发起时间早于退款时间，或者间隔退款发起时间太短，
	//	可能出现退款查询时还没处理成功，后面又处理成功的情况，
	//	建议商户在退款发起后间隔10秒以上再发起退款查询请求。
	//【枚举值】
	//	退款处理成功: REFUND_SUCCESS
	//【示例值】REFUND_SUCCESS
	RefundStatus string `json:"refund_status"`
	//【描述】退分账明细信息，当前仅在直付通产品中返回。
	RefundRoyaltys []*RefundRoyalty `json:"refund_royaltys"`
	//【描述】退款时间。默认不返回该信息，需要在入参的query_options中指定"gmt_refund_pay"值时才返回该字段信息。
	//【示例值】2014-11-27 15:45:57
	GmtRefundPay string `json:"gmt_refund_pay"`
	//【描述】本次退款使用的资金渠道； 默认不返回该信息，需要在入参的query_options中指定"refund_detail_item_list"值时才返回该字段信息。
	RefundDetailItemList []*RefundDetailItem `json:"refund_detail_item_list"`
	//【描述】本次商户实际退回金额；单位：元。
	//	默认不返回该信息，需要在入参的query_options中指定"refund_detail_item_list"值时才返回该字段信息。
	//【示例值】88
	SendBackFee string `json:"send_back_fee"`
	//【描述】银行卡冲退信息； 默认不返回该信息，需要在入参的query_options中指定"deposit_back_info"值时才返回该字段信息。
	DepositBackInfo *DepositBackInfo `json:"deposit_back_info"`
	//【描述】本次退款金额中退惠营宝的金额。单位：元。
	//【示例值】10.24
	RefundHybAmount string `json:"refund_hyb_amount"`
	//【描述】退费信息
	RefundChargeInfoList []*RefundChargeInfo `json:"refund_charge_info_list"`
	//【描述】银行卡冲退信息列表。 默认不返回该信息，需要在入参的query_options中指定"deposit_back_info_list"值时才返回该字段信息。
	DepositBackInfoList []*DepositBackInfo `json:"deposit_back_info_list"`
	//【描述】本交易支付时使用的所有优惠券信息。 只有在query_options中指定refund_voucher_detail_list时才返回该字段信息。
	//【必选条件】query_options中包含refund_voucher_detail_list时，才会返回券信息列表
	RefundVoucherDetailList []*RefundVoucherDetail `json:"refund_voucher_detail_list"`
	//【描述】当用户使用芝麻信用先享后付时，且当前的操作为预授权撤销动作时，会返回该字段，代表当前撤销的预授权金额，单位元。
	//【必选条件】当用户使用芝麻信用先享后付时，且当前的操作为预授权撤销动作时，会返回该字段。
	//【示例值】12.45
	PreAuthCancelFee string `json:"pre_auth_cancel_fee"`
}

type AlipayTradeFastPayRefundQueryResponse struct {
	PublicResponseParameters
	Response TradeFastPayRefundQueryResponse `json:"alipay_trade_fastpay_refund_query_response"`
}
