package types

import "encoding/json"

// FundTransRefund 转账撤销请求
type FundTransRefund struct {
	//【描述】退款请求流水号，由商家自定义，不可重复
	OutRequestNo string `json:"out_request_no"`
	//【描述】支付宝转账单据号，与 out_biz_no 二选一
	OrderId string `json:"order_id,omitempty"`
	//【描述】商家侧唯一订单号，与 order_id 二选一
	OutBizNo string `json:"out_biz_no,omitempty"`
	//【描述】退款金额，单位元，不超过原转账金额
	RefundAmount string `json:"refund_amount"`
	//【描述】退款原因说明
	Remark string `json:"remark,omitempty"`
}

func (r *FundTransRefund) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// FundTransRefundDetail 转账撤销响应详情
type FundTransRefundDetail struct {
	PublicResponseParameters
	//【描述】支付宝转账单据号
	OrderId string `json:"order_id"`
	//【描述】退款单据号
	RefundOrderId string `json:"refund_order_id"`
	//【描述】实际退款金额，单位元
	RefundAmount string `json:"refund_amount"`
	//【描述】退款时间，格式 yyyy-MM-dd HH:mm:ss
	RefundDate string `json:"refund_date"`
}

// AlipayFundTransRefundResponse 转账撤销响应
type AlipayFundTransRefundResponse struct {
	AlipayFundTransRefundResponse FundTransRefundDetail `json:"alipay_fund_trans_refund_response"`
	Sign                          string                `json:"sign"`
}
