package types

// SubMerchant 【描述】二级商户信息。
// 直付通模式和机构间连模式下必传，其它场景下不需要传入。
type SubMerchant struct {
	//【描述】间连受理商户的支付宝商户编号，通过间连商户入驻后得到。间连业务下必传，并且需要按规范传递受理商户编号。
	//【示例值】2088000603999128
	MerchantId string `json:"merchant_id"`
	//【描述】二级商户编号类型。
	//	枚举值：
	//	alipay:支付宝分配的间联商户编号；
	//	目前仅支持alipay，默认可以不传。
	//【枚举值】
	//	alipay: alipay
	//【示例值】alipay
	MerchantType string `json:"merchant_type,omitempty"`
}

// KeyInfo 【描述】开票关键信息
type KeyInfo struct {
	//【描述】该交易是否支持开票
	//【示例值】true
	IsSupportInvoice bool `json:"is_support_invoice"`
	//【描述】开票商户名称：商户品牌简称|商户门店简称
	//【示例值】ABC|003
	InvoiceMerchantName string `json:"invoice_merchant_name"`
	//【描述】税号
	//【示例值】1464888883494
	TaxNum string `json:"tax_num"`
}

// InvoiceInfo 【描述】开票信息
type InvoiceInfo struct {
	//【描述】开票关键信息
	KeyInfo *KeyInfo `json:"key_info"`
	//【描述】开票内容
	//	注：json数组格式
	//【示例值】[{"code":"100294400","name":"服饰","num":"2","sumPrice":"200.00","taxRate":"6%"}]
	Details string `json:"details"`
}

type TradePayPayResponse struct {
	//【描述】用于跳转支付宝页面的信息，POST和GET方法生成内容不同：
	//	使用POST方法执行，结果为html form表单，在浏览器渲染即可；
	//	使用GET方法会得到支付宝URL，需要打开或重定向到该URL。
	//	建议使用POST方式。
	//	具体使用方法请参考: https://opendocs.alipay.com/open/270/105899?pathHash=d57664bf
	//【示例值】请参考响应示例
	PageRedirectionData string `json:"page_redirection_data"`
}
