# GoodSender SDK for Go

Official client library for the GoodSender email API. Package: `github.com/good-sender/go-sdk`

## Installation

```bash
go get github.com/good-sender/go-sdk
```

## Quick start

```go
package main

import (
    "context"
    "fmt"

    goodsender "github.com/good-sender/go-sdk"
)

func main() {
    cfg := goodsender.NewConfiguration()
    cfg.Servers = goodsender.ServerConfigurations{{URL: "https://api.goodsender.com"}}
    cfg.AddDefaultHeader("Authorization", "Bearer YOUR_API_KEY")
    client := goodsender.NewAPIClient(cfg)
    ctx := context.Background()

    body := "Body"
    req := goodsender.SendEmailRequest{
        Emails: []goodsender.SendEmail{
            {
                From:        goodsender.Address{Email: "sender@example.com"},
                To:          []goodsender.Address{{Email: "recipient@example.com"}},
                Subject:     "Hello",
                TextContent: &body,
            },
        },
    }
    res, _, err := client.EmailsAPI.SendEmail(ctx).SendEmailRequest(req).Execute()
    if err != nil {
        panic(err)
    }
    fmt.Printf("sent=%d declined=%d\n", res.Sent, res.Declined)
}
```

## Examples

### Send via a template

```go
req := goodsender.TemplateEmailRequest{
    From:    goodsender.Address{Email: "sender@example.com"},
    To:      goodsender.Address{Email: "recipient@example.com"},
    Subject: "Your OTP",
    Template: goodsender.TemplateEmailRequestTemplate{
        TemplateId: "otp_code",
        Variables:  map[string]string{"code": "123456"},
    },
}
res, _, err := client.EmailsAPI.SendTemplateEmail(ctx).TemplateEmailRequest(req).Execute()
```

### List domains

```go
res, _, err := client.DomainsAPI.ListDomains(ctx).Limit(50).Execute()
fmt.Printf("domains=%d\n", len(res.Domains))
```

### Check consent status

```go
res, _, err := client.EmailsAPI.GetEmailConsentStatus(ctx, "user@example.com").
    Domain("example.com").Execute()
fmt.Printf("entries=%d\n", len(res))

// List all consents for a domain
list, _, err := client.EmailsAPI.ListEmailConsents(ctx).Domain("example.com").Limit(50).Execute()
fmt.Printf("emails=%d\n", len(list.Emails))
```

## Documentation

- API reference: <https://goodsender.com/docs>
- OpenAPI spec: `openapi/goodsender.yaml` in this repo
- Conformance tests: `tests/`

## Development

- Regenerate from spec: `scripts/regen.sh` (preserves `tests/`, `.github/`, and hand-curated files per `.regen-ignore`)
- Run conformance tests against local mock: `tests/run.sh mock`
- Run conformance against real dev API: `tests/run.sh dev` (requires `tests/.env.dev`)

## License

MIT — see [LICENSE](LICENSE).
