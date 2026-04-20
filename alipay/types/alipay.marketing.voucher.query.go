package types

import "encoding/json"

type VoucherQuery struct {
	//【描述】优惠券模板 ID（与 DetailId 二选一）
	//【示例值】VCH20260420001
	VoucherId string `json:"voucher_id,omitempty"`
	//【描述】用户领券记录 ID（与 VoucherId 二选一）
	//【示例值】DTL20260420001
	DetailId string `json:"detail_id,omitempty"`
}

func (r *VoucherQuery) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

type VoucherQueryDetail struct {
	PublicResponseParameters
	VoucherId      string `json:"voucher_id"`
	VoucherName    string `json:"voucher_name"`
	VoucherType    string `json:"voucher_type"`
	Status         string `json:"status"`
	ValidBeginTime string `json:"valid_begin_time"`
	ValidEndTime   string `json:"valid_end_time"`
}

type AlipayMarketingVoucherQueryResponse struct {
	AlipayMarketingVoucherQueryResponse VoucherQueryDetail `json:"alipay_marketing_voucher_query_response"`
	Sign                                string             `json:"sign"`
}
