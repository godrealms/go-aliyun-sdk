package types

import "encoding/json"

// FundTransOrderQuery 查询转账订单请求（新版）
type FundTransOrderQuery struct {
	//【描述】商家侧唯一订单号，与 order_id 二选一
	OutBizNo string `json:"out_biz_no,omitempty"`
	//【描述】支付宝转账单据号，与 out_biz_no 二选一
	OrderId string `json:"order_id,omitempty"`
}

func (r *FundTransOrderQuery) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// FundTransOrderQueryDetail 查询转账订单响应详情
type FundTransOrderQueryDetail struct {
	PublicResponseParameters
	//【描述】支付宝转账单据号
	OrderId string `json:"order_id"`
	//【描述】转账单据状态。SUCCESS | FAIL | DEALING
	Status string `json:"status"`
	//【描述】付款完成时间，格式 yyyy-MM-dd HH:mm:ss
	PayDate string `json:"pay_date"`
	//【描述】预计到账时间
	ArrivalTimeEnd string `json:"arrival_time_end"`
	//【描述】转账金额，单位元
	OrderFee string `json:"order_fee"`
	//【描述】失败原因
	FailReason string `json:"fail_reason"`
}

// AlipayFundTransOrderQueryResponse 查询转账订单响应
type AlipayFundTransOrderQueryResponse struct {
	AlipayFundTransOrderQueryResponse FundTransOrderQueryDetail `json:"alipay_fund_trans_order_query_response"`
	Sign                              string                    `json:"sign"`
}
