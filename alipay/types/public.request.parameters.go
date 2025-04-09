package types

import (
	"github.com/godrealms/go-aliyun-sdk/utils"
	"net/url"
)

type PublicRequestParameters struct {
	//【描述】支付宝分配给开发者的应用ID
	//【示例值】2014072300007148
	AppId string `json:"app_id"`
	//【描述】接口名称
	//【示例值】alipay.user.agreement.page.sign
	Method string `json:"method"`
	//【描述】仅支持JSON
	//【示例值】JSON
	Format string `json:"format,omitempty"`
	//【描述】HTTP/HTTPS开头字符串
	//【示例值】https://m.alipay.com/Gk8NF23
	ReturnUrl string `json:"return_url,omitempty"`
	//【描述】请求使用的编码格式，如utf-8,gbk,gb2312等
	//【示例值】utf-8
	Charset string `json:"charset"`
	//【描述】商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2
	// 【示例值】RSA2
	SignType string `json:"sign_type"`
	//【描述】商户请求参数的签名串，详见: https://opendocs.alipay.com/common/02khjm
	Sign string `json:"sign"`
	//【描述】发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
	//【示例值】2016-07-29 16:00:00
	Timestamp string `json:"timestamp"`
	//【描述】调用的接口版本，固定为：1.0
	//【示例值】1.0
	Version string `json:"version"`
	//【描述】支付宝服务器主动通知商户服务器里指定的页面http/https路径。
	//【示例值】https://m.alipay.com/Gk8NF23
	NotifyUrl string `json:"notify_url,omitempty"`
	//【描述】详见应用授权概述: https://opendocs.alipay.com/isv/10467/xldcyq
	AppAuthToken string `json:"app_auth_token,omitempty"`
	//【描述】请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
	BizContent string `json:"biz_content"`
}

func (p *PublicRequestParameters) ToUrlValue() url.Values {
	return utils.StructToValues(p)
}

type PublicResponseParameters struct {
	Code    string `json:"code"`     // 必选	- 网关返回码,详见文档: https://opendoc.alipay.com/common/02km9f
	Msg     string `json:"msg"`      // 必选	- 网关返回码描述,详见文档: https://opendoc.alipay.com/common/02km9f
	SubCode string `json:"sub_code"` // 可选	- 业务返回码，参见具体的API接口文档
	SubMsg  string `json:"sub_msg"`  // 可选	- 业务返回码描述，参见具体的API接口文档
	Sign    string `json:"sign"`     // 必选	- 签名,详见文档: https://opendoc.alipay.com/common/02kf5q
}
