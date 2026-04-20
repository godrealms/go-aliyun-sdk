package alipay

import (
	"strings"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestGetOpenAuthPageURL_Production(t *testing.T) {
	client := NewISVClient("app123", "", "", "")
	req := &types.OpenAuthPage{
		RedirectUri: "https://example.com/callback",
		State:       "mystate",
	}
	u, err := client.GetOpenAuthPageURL(req)
	if err != nil {
		t.Fatalf("GetOpenAuthPageURL error: %v", err)
	}
	if !strings.HasPrefix(u, "https://openauth.alipay.com") {
		t.Errorf("URL should start with production host, got: %s", u)
	}
	if !strings.Contains(u, "app_id=app123") {
		t.Errorf("URL missing app_id, got: %s", u)
	}
	if !strings.Contains(u, "redirect_uri=") {
		t.Errorf("URL missing redirect_uri, got: %s", u)
	}
	if !strings.Contains(u, "state=mystate") {
		t.Errorf("URL missing state, got: %s", u)
	}
}

func TestGetOpenAuthPageURL_Sandbox(t *testing.T) {
	client := NewISVClient("app123", "", "", "")
	client.Sandbox = true
	req := &types.OpenAuthPage{
		RedirectUri: "https://example.com/callback",
	}
	u, err := client.GetOpenAuthPageURL(req)
	if err != nil {
		t.Fatalf("GetOpenAuthPageURL error: %v", err)
	}
	if !strings.HasPrefix(u, "https://openauth-sandbox.dl.alipaydev.com") {
		t.Errorf("URL should start with sandbox host, got: %s", u)
	}
}

func TestGetOpenAuthPageURL_MissingRedirectUri(t *testing.T) {
	client := NewISVClient("app123", "", "", "")
	req := &types.OpenAuthPage{}
	_, err := client.GetOpenAuthPageURL(req)
	if err == nil {
		t.Error("expected error for missing redirect_uri, got nil")
	}
}
