package alipay

import (
	"testing"
)

func TestNewISVClient(t *testing.T) {
	client := NewISVClient("app123", "privkey", "pubkey", "token456")
	if client.AppId != "app123" {
		t.Errorf("AppId = %q, want %q", client.AppId, "app123")
	}
	if client.PrivateKey != "privkey" {
		t.Errorf("PrivateKey = %q, want %q", client.PrivateKey, "privkey")
	}
	if client.AlipayPublicKey != "pubkey" {
		t.Errorf("AlipayPublicKey = %q, want %q", client.AlipayPublicKey, "pubkey")
	}
	if client.AppAuthToken != "token456" {
		t.Errorf("AppAuthToken = %q, want %q", client.AppAuthToken, "token456")
	}
	if client.Http == nil {
		t.Error("Http client should not be nil")
	}
}
