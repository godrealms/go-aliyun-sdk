package types

import "encoding/json"

type BillType string

const (
	BillTypeTrade            BillType = "trade"
	BillTypeSignCustomer     BillType = "signcustomer"
	BillTypeMerchantAct      BillType = "merchant_act"
	BillTypeTradeZftMerchant BillType = "trade_zft_merchant"
	BillTypeZftAcc           BillType = "zft_acc"
	BillTypeSettlementMerge  BillType = "settlementMerge"
)

type BillDownloadUrlQuery struct {
	//【描述】账单类型，商户通过接口或商户经开放平台授权后其所属服务商通过接口可以获取以下账单类型。
	//【枚举值】
	//	商户基于支付宝交易收单的业务账单: trade
	//	基于商户支付宝余额收入及支出等资金变动的账务账单: signcustomer
	//	营销活动账单，包含营销活动的发放，核销记录: merchant_act
	//	直付通二级商户查询交易的业务账单: trade_zft_merchant
	//	直付通平台商查询二级商户流水使用，返回所有二级商户流水。: zft_acc
	//	每日结算到卡的资金对应的明细，下载内容包含批次结算到卡明细文件（示例）和批次结算到卡汇总文件（示例）；若查询时间范围内有多个批次，会将多个批次的明细和汇总文件打包到一份压缩包中；: settlementMerge
	//【示例值】trade
	BillType BillType `json:"bill_type"`
	//【描述】账单时间：
	//	* 日账单格式为yyyy-MM-dd，最早可下载2016年1月1日开始的日账单。不支持下载当日账单，只能下载前一日24点前的账单数据（T+1），当日数据一般于次日 9 点前生成，特殊情况可能延迟。
	//	* 月账单格式为yyyy-MM，最早可下载2016年1月开始的月账单。不支持下载当月账单，只能下载上一月账单数据，当月账单一般在次月 3 日生成，特殊情况可能延迟。
	//	* 当biz_type为settlementMerge时候，时间为汇总批次结算资金到账的日期，日期格式为yyyy-MM-dd，最早可下载2023年4月17日及以后的账单。
	//【示例值】2016-04-05
	BillDate string `json:"bill_date"`
	//【描述】二级商户smid，这个参数只在bill_type是trade_zft_merchant时才能使用
	//【示例值】2088123412341234
	Smid string `json:"smid,omitempty"`
}

func (q *BillDownloadUrlQuery) ToString() string {
	marshal, _ := json.Marshal(q)
	return string(marshal)
}

type DownloadUrl struct {
	Code            string `json:"code"`
	Msg             string `json:"msg"`
	BillDownloadUrl string `json:"bill_download_url"`
	BillFileCode    string `json:"bill_file_code"`
}

type BillDownloadUrlResponse struct {
	PublicResponseParameters
	Response DownloadUrl `json:"alipay_data_dataservice_bill_downloadurl_query_response"`
}
