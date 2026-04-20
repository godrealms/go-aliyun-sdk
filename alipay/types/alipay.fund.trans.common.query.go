package types

import "encoding/json"

// FundTransCommonQuery 通用转账查询请求（旧版）
type FundTransCommonQuery struct {
	//【描述】产品码，与转账时一致
	ProductCode string `json:"product_code,omitempty"`
	//【描述】业务场景
	BizScene string `json:"biz_scene,omitempty"`
	//【描述】商家侧唯一订单号，与 order_id 二选一
	OutBizNo string `json:"out_biz_no,omitempty"`
	//【描述】支付宝转账单据号，与 out_biz_no 二选一
	OrderId string `json:"order_id,omitempty"`
	//【描述】支付宝支付资金流水号
	PayFundOrderId string `json:"pay_fund_order_id,omitempty"`
}

func (r *FundTransCommonQuery) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// FundTransCommonQueryDetail 通用转账查询响应详情
type FundTransCommonQueryDetail struct {
	PublicResponseParameters
	//【描述】支付宝转账单据号
	OrderId string `json:"order_id"`
	//【描述】支付宝支付资金流水号
	PayFundOrderId string `json:"pay_fund_order_id"`
	//【描述】转账单据状态。SUCCESS | FAIL | DEALING
	Status string `json:"status"`
	//【描述】付款完成时间，格式 yyyy-MM-dd HH:mm:ss
	PayDate string `json:"pay_date"`
	//【描述】转账金额，单位元
	OrderFee string `json:"order_fee"`
	//【描述】失败原因
	FailReason string `json:"fail_reason"`
}

// AlipayFundTransCommonQueryResponse 通用转账查询响应
type AlipayFundTransCommonQueryResponse struct {
	AlipayFundTransCommonQueryResponse FundTransCommonQueryDetail `json:"alipay_fund_trans_common_query_response"`
	Sign                               string                     `json:"sign"`
}
