package alipay

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestClient(t *testing.T, responseJSON string) (*Client, *httptest.Server) {
	t.Helper()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, responseJSON)
	}))
	t.Cleanup(ts.Close)

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generate rsa key: %v", err)
	}
	privDER, _ := x509.MarshalPKCS8PrivateKey(privateKey)
	privPEM := string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privDER}))

	pubDER, _ := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	pubPEM := string(pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubDER}))

	client := NewISVClient("test_app_id", privPEM, pubPEM, "test_app_auth_token")
	client.Http.SetBaseURL(ts.URL)
	return client, ts
}
