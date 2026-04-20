package types

import "encoding/json"

type VoucherCreate struct {
	//【描述】优惠券名称
	//【示例值】满100减10
	VoucherName string `json:"voucher_name"`
	//【描述】优惠券类型：MONEY_VOUCHER（代金券）| DISCOUNT_VOUCHER（折扣券）
	//【示例值】MONEY_VOUCHER
	VoucherType string `json:"voucher_type"`
	//【描述】优惠金额（元），VoucherType=MONEY_VOUCHER 时必填
	//【示例值】10.00
	DenominationMoney string `json:"denomination_money"`
	//【描述】商户 ID
	//【示例值】merchant001
	MerchantId string `json:"merchant_id"`
	//【描述】发行数量
	//【示例值】100
	Quantity string `json:"quantity"`
	//【描述】有效期开始时间，格式 yyyy-MM-dd HH:mm:ss
	//【示例值】2026-04-20 00:00:00
	ValidBeginTime string `json:"valid_begin_time"`
	//【描述】有效期结束时间，格式 yyyy-MM-dd HH:mm:ss
	//【示例值】2026-05-20 23:59:59
	ValidEndTime string `json:"valid_end_time"`
}

func (r *VoucherCreate) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

type VoucherCreateDetail struct {
	PublicResponseParameters
	//【描述】优惠券模板 ID
	//【示例值】VCH20260420001
	VoucherId string `json:"voucher_id"`
}

type AlipayMarketingVoucherCreateResponse struct {
	AlipayMarketingVoucherCreateResponse VoucherCreateDetail `json:"alipay_marketing_voucher_create_response"`
	Sign                                 string              `json:"sign"`
}
