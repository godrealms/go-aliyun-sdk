package types

import "encoding/json"

type VoucherSend struct {
	//【描述】优惠券模板 ID，由 AlipayMarketingVoucherCreate 返回
	//【示例值】VCH20260420001
	VoucherId string `json:"voucher_id"`
	//【描述】用户支付宝 openId
	//【示例值】0680809090909090909090909090
	OpenId string `json:"open_id"`
	//【描述】外部业务号，幂等控制
	//【示例值】send20260420001
	OutBizNo string `json:"out_biz_no"`
}

func (r *VoucherSend) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

type VoucherSendDetail struct {
	PublicResponseParameters
	//【描述】用户领券记录 ID
	//【示例值】DTL20260420001
	DetailId string `json:"detail_id"`
}

type AlipayMarketingVoucherSendResponse struct {
	AlipayMarketingVoucherSendResponse VoucherSendDetail `json:"alipay_marketing_voucher_send_response"`
	Sign                               string            `json:"sign"`
}
