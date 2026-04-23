package types

import "encoding/json"

// TradeCreate 统一收单下单请求参数
// 本结构覆盖常用字段；未列出的长尾字段可在使用方按需自行拼接 biz_content JSON。
type TradeCreate struct {
	//【描述】商户网站唯一订单号
	//【示例值】20150320010101001
	OutTradeNo string `json:"out_trade_no"`
	//【描述】订单标题
	//【示例值】大乐透
	Subject string `json:"subject"`
	//【描述】订单总金额，单位为元，精确到小数点后两位
	//【示例值】88.00
	TotalAmount string `json:"total_amount"`
	//【描述】买家的支付宝用户id；buyer_id 与 buyer_logon_id 二选一必填
	//【示例值】2088102146225135
	BuyerId string `json:"buyer_id,omitempty"`
	//【描述】买家支付宝账号；buyer_id 与 buyer_logon_id 二选一必填
	//【示例值】15900000001
	BuyerLogonId string `json:"buyer_logon_id,omitempty"`
	//【描述】销售产品码
	//【示例值】FACE_TO_FACE_PAYMENT
	ProductCode string `json:"product_code,omitempty"`
	//【描述】订单可打折金额
	//【示例值】8.00
	DiscountableAmount string `json:"discountable_amount,omitempty"`
	//【描述】订单不可打折金额
	//【示例值】80.00
	UndiscountableAmount string `json:"undiscountable_amount,omitempty"`
	//【描述】卖家支付宝用户ID，为空则取商户签约账号对应的用户ID
	//【示例值】2088102146225135
	SellerId string `json:"seller_id,omitempty"`
	//【描述】订单描述
	//【示例值】Iphone6 16G
	Body string `json:"body,omitempty"`
	//【描述】商户操作员编号
	//【示例值】yx_001
	OperatorId string `json:"operator_id,omitempty"`
	//【描述】商户门店编号
	//【示例值】NJ_S_001
	StoreId string `json:"store_id,omitempty"`
	//【描述】商户机具终端编号
	//【示例值】NJ_T_001
	TerminalId string `json:"terminal_id,omitempty"`
	//【描述】订单包含的商品列表信息
	GoodsDetail []*GoodsDetail `json:"goods_detail,omitempty"`
	//【描述】业务扩展参数
	ExtendParams *ExtendParams `json:"extend_params,omitempty"`
	//【描述】二级商户信息；直付通/机构间连模式下必传
	SubMerchant *SubMerchant `json:"sub_merchant,omitempty"`
	//【描述】外部指定买家信息
	ExtUserInfo *ExtUserInfo `json:"ext_user_info,omitempty"`
	//【描述】商户传入业务信息，json 格式，应用于安全、营销等参数直传场景
	//【示例值】{"mc_create_trade_ip":"127.0.0.1"}
	BusinessParams string `json:"business_params,omitempty"`
	//【描述】绝对超时时间，格式为 yyyy-MM-dd HH:mm:ss
	//【示例值】2016-12-31 10:05:00
	TimeExpire string `json:"time_expire,omitempty"`
	//【描述】相对超时时间；从交易创建开始计时，超过该时间未支付则关闭交易，如 "90m"、"1h"
	//【示例值】90m
	TimeoutExpress string `json:"timeout_express,omitempty"`
	//【描述】公用回传参数；如请求时传递，支付宝异步通知及同步返回时将原样回传，
	//	发送前需进行 URL encode。
	//【示例值】merchantBizType%3d3C%26merchantBizNo%3d2016010101111
	PassbackParams string `json:"passback_params,omitempty"`
	//【描述】商户原始订单号，最大长度限制32位
	//【示例值】20161008001
	MerchantOrderNo string `json:"merchant_order_no,omitempty"`
}

func (r *TradeCreate) ToString() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}

// TradeCreateResponse 统一收单下单响应
type TradeCreateResponse struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	TradeNo string `json:"trade_no"`
}

// AlipayTradeCreateResponse 统一收单下单响应 wrapper
type AlipayTradeCreateResponse struct {
	PublicResponseParameters
	Response TradeCreateResponse `json:"alipay_trade_create_response"`
}
