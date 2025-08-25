package alipay

import (
	"context"
	"github.com/godrealms/go-aliyun-sdk/alipay/types"
	"github.com/godrealms/go-aliyun-sdk/community"
	"time"
)

// AlipayUserAgreementExecutionPlanModify 周期性扣款协议执行计划修改接口
func (c *Client) AlipayUserAgreementExecutionPlanModify(form *types.AgreementExecutionPlanModify) (*types.AgreementExecutionPlanModifyResponse, error) {
	data := types.PublicRequestParameters{
		AppId:        c.AppId,
		Method:       "alipay.user.agreement.executionplan.modify",
		Format:       "JSON",
		ReturnUrl:    c.ReturnUrl,
		Charset:      "UTF-8",
		SignType:     "RSA2",
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Version:      "1.0",
		NotifyUrl:    c.NotifyUrl,
		AppAuthToken: c.AppAuthToken,
		BizContent:   form.ToString(),
	}

	signature, err := community.NewSignatureHelper(c.PrivateKey)
	if err != nil {
		return nil, err
	}
	data.Sign, err = signature.GenerateSignature(data)
	if err != nil {
		return nil, err
	}
	value := data.ToUrlValue()
	result := &types.AgreementExecutionPlanModifyResponse{}
	err = c.Http.Get(context.Background(), "", value, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
