package types

// OpenAuthPage 商家授权页跳转参数
type OpenAuthPage struct {
	//【描述】授权回调地址，需与应用配置中的授权回调地址一致
	//【示例值】https://example.com/callback
	RedirectUri string
	//【描述】商户自定义参数，授权完成后原样返回
	//【示例值】init
	State string
}
