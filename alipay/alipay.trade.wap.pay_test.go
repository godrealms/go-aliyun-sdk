package alipay

import (
	"strings"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestAlipayTradeWapPay(t *testing.T) {
	client, _ := newTestClient(t, "")
	req := &types.TradeWapPay{
		OutTradeNo:  "wap20260420001",
		Subject:     "测试H5支付",
		TotalAmount: "88.00",
		ProductCode: "QUICK_WAP_WAY",
		QuitUrl:     "https://example.com/quit",
	}
	result, err := client.AlipayTradeWapPay(req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if result == "" {
		t.Error("result URL is empty")
	}
	if !strings.Contains(result, "alipay.trade.wap.pay") {
		t.Errorf("result does not contain method name: %q", result)
	}
}
