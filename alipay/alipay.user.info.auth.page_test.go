package alipay

import (
	"strings"
	"testing"

	"github.com/godrealms/go-aliyun-sdk/alipay/types"
)

func TestGetUserAuthPageURL_Production(t *testing.T) {
	client := NewISVClient("app_test_001", "", "", "")
	req := &types.UserInfoAuthPage{
		Scope:       "auth_user",
		RedirectUri: "https://example.com/callback",
		State:       "state_xyz",
	}
	u, err := client.GetUserAuthPageURL(req)
	if err != nil {
		t.Fatalf("GetUserAuthPageURL error: %v", err)
	}
	if !strings.HasPrefix(u, "https://openauth.alipay.com") {
		t.Errorf("URL should start with production host, got: %s", u)
	}
	if !strings.Contains(u, "app_id=app_test_001") {
		t.Errorf("URL missing app_id, got: %s", u)
	}
	if !strings.Contains(u, "scope=auth_user") {
		t.Errorf("URL missing scope, got: %s", u)
	}
	if !strings.Contains(u, "redirect_uri=") {
		t.Errorf("URL missing redirect_uri, got: %s", u)
	}
	if !strings.Contains(u, "state=state_xyz") {
		t.Errorf("URL missing state, got: %s", u)
	}
}

func TestGetUserAuthPageURL_Sandbox(t *testing.T) {
	client := NewISVClient("app_test_001", "", "", "")
	client.Sandbox = true
	req := &types.UserInfoAuthPage{
		Scope:       "auth_base",
		RedirectUri: "https://example.com/callback",
	}
	u, err := client.GetUserAuthPageURL(req)
	if err != nil {
		t.Fatalf("GetUserAuthPageURL error: %v", err)
	}
	if !strings.HasPrefix(u, "https://openauth-sandbox.dl.alipaydev.com") {
		t.Errorf("URL should start with sandbox host, got: %s", u)
	}
}

func TestGetUserAuthPageURL_MissingRedirectUri(t *testing.T) {
	client := NewISVClient("app_test_001", "", "", "")
	req := &types.UserInfoAuthPage{Scope: "auth_user"}
	_, err := client.GetUserAuthPageURL(req)
	if err == nil {
		t.Error("expected error for missing redirect_uri, got nil")
	}
}

func TestGetUserAuthPageURL_MissingScope(t *testing.T) {
	client := NewISVClient("app_test_001", "", "", "")
	req := &types.UserInfoAuthPage{RedirectUri: "https://example.com/callback"}
	_, err := client.GetUserAuthPageURL(req)
	if err == nil {
		t.Error("expected error for missing scope, got nil")
	}
}
