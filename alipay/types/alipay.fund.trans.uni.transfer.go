package types

import "encoding/json"

// FundTransPayee 收款方信息
type FundTransPayee struct {
	//【描述】收款方账户，与 identity_type 配合使用
	Identity string `json:"identity"`
	//【描述】账户类型。ALIPAY_USER_ID：支付宝 UID；ALIPAY_LOGON_ID：登录账号（手机/邮箱）
	IdentityType string `json:"identity_type"`
	//【描述】收款方真实姓名，部分场景必填
	Name string `json:"name,omitempty"`
}

// FundTransUniTransfer 单笔转账请求（新版）
type FundTransUniTransfer struct {
	//【描述】商家侧唯一订单号，由商家自定义，不可重复
	//【示例值】201806300001
	OutBizNo string `json:"out_biz_no"`
	//【描述】转账金额，单位：元，精确到小数点后两位
	//【示例值】10.00
	TransAmount string `json:"trans_amount"`
	//【描述】产品码，固定值 TRANS_ACCOUNT_NO_PWD
	//【示例值】TRANS_ACCOUNT_NO_PWD
	ProductCode string `json:"product_code"`
	//【描述】业务场景，默认 DIRECT_TRANSFER
	//【示例值】DIRECT_TRANSFER
	BizScene string `json:"biz_scene,omitempty"`
	//【描述】转账业务的标题，用于在收款方账单中展示
	//【示例值】提现
	OrderTitle string `json:"order_title,omitempty"`
	//【描述】收款方信息
	PayeeInfo *FundTransPayee `json:"payee_info"`
	//【描述】业务备注
	Remark string `json:"remark,omitempty"`
}

func (r *FundTransUniTransfer) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// FundTransUniTransferDetail 单笔转账响应详情
type FundTransUniTransferDetail struct {
	PublicResponseParameters
	//【描述】商家侧唯一订单号
	OutBizNo string `json:"out_biz_no"`
	//【描述】支付宝转账单据号
	OrderId string `json:"order_id"`
	//【描述】支付宝支付资金流水号
	PayFundOrderId string `json:"pay_fund_order_id"`
	//【描述】转账单据状态。SUCCESS：成功；FAIL：失败；DEALING：处理中
	Status string `json:"status"`
	//【描述】订单支付时间，格式 yyyy-MM-dd HH:mm:ss
	TransDate string `json:"trans_date"`
}

// AlipayFundTransUniTransferResponse 单笔转账响应
type AlipayFundTransUniTransferResponse struct {
	AlipayFundTransUniTransferResponse FundTransUniTransferDetail `json:"alipay_fund_trans_uni_transfer_response"`
	Sign                               string                     `json:"sign"`
}
