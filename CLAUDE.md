# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
go build ./...    # Build all packages
go test ./...     # Run tests
go vet ./...      # Static analysis
go fmt ./...      # Format code
go mod tidy       # Sync dependencies
```

No Makefile, CI, or test files exist yet. No external dependencies beyond the Go standard library.

## Architecture

This is a Go SDK for Aliyun services, currently implementing Alipay APIs.

**Package responsibilities:**

- **`alipay/`** — Core client. `client.go` defines the `Client` struct (AppId, PrivateKey, AlipayPublicKey, NotifyUrl, ReturnUrl, sandbox flag, embedded `*community.HTTP`). Each API endpoint has its own file (e.g., `alipay.trade.query.go`) as a receiver method on `Client`.
- **`alipay/types/`** — Typed request and response structs for each endpoint, plus shared `PublicRequestParameters` / `PublicResponseParameters`. Fields use JSON tags and Chinese comments documenting Alipay's API specs.
- **`community/`** — Shared infrastructure: `http.go` is a generic HTTP client; `signature.go` handles RSA2/RSA/ECDSA signing & verification, supporting PKCS#8, PKCS#1, and EC key formats.
- **`utils/`** — `StructToValues()` converts any struct to `url.Values` via reflection, respecting JSON tags and handling nested structs, slices, and maps.
- **`example/alipay/`** — One example file per API endpoint showing complete client setup.

## Request Flow

Every API call follows the same pattern:
1. Caller passes a typed request struct (e.g., `*types.TradeQuery`)
2. Client wraps it in `PublicRequestParameters` (adds `app_id`, `method`, `timestamp`, `version`, `charset`, `sign_type`)
3. Request struct is serialized to `biz_content` JSON
4. `SignatureHelper` generates an RSA2 signature over the sorted, filtered parameters
5. Parameters are URL-form-encoded and sent via GET request
6. JSON response is unmarshaled into a typed response struct

## Adding a New Endpoint

1. Add request/response types in `alipay/types/alipay.<method>.go`
2. Add the method file `alipay/alipay.<method>.go` implementing a `Client` receiver
3. Follow the existing pattern: build `PublicRequestParameters`, marshal `biz_content`, sign, POST/GET, unmarshal response
4. Add an example under `example/alipay/`
