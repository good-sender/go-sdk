//go:build !conformance

// Mock-mode smoke test for the Go SDK.
package main

import (
	"context"
	"fmt"
	"os"

	goodsender "github.com/good-sender/go-sdk"
)

type result struct {
	method string
	ok     bool
	detail string
}

var results []result

func run(method string, fn func() (string, error)) {
	detail, err := fn()
	if err != nil {
		results = append(results, result{method, false, fmt.Sprintf("%T: %v", err, err)})
		return
	}
	results = append(results, result{method, true, detail})
}

func main() {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:4010"
	}
	apiKey := os.Getenv("GOODSENDER_API_KEY")
	if apiKey == "" {
		apiKey = "test-key"
	}

	cfg := goodsender.NewConfiguration()
	cfg.Servers = goodsender.ServerConfigurations{{URL: baseURL}}
	cfg.AddDefaultHeader("Authorization", "Bearer "+apiKey)
	client := goodsender.NewAPIClient(cfg)
	ctx := context.Background()

	textBody := "Body"

	run("sendEmail", func() (string, error) {
		req := goodsender.SendEmailRequest{
			Emails: []goodsender.SendEmail{
				{
					From:        goodsender.Address{Email: "sender@example.com"},
					To:          []goodsender.Address{{Email: "recipient@example.com"}},
					Subject:     "Hello",
					TextContent: &textBody,
				},
			},
		}
		res, _, err := client.EmailsAPI.SendEmail(ctx).SendEmailRequest(req).Execute()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("sent=%d declined=%d", res.Sent, res.Declined), nil
	})

	run("sendTemplateEmail", func() (string, error) {
		req := goodsender.TemplateEmailRequest{
			From:    goodsender.Address{Email: "sender@example.com"},
			To:      goodsender.Address{Email: "recipient@example.com"},
			Subject: "OTP",
			Template: goodsender.TemplateEmailRequestTemplate{
				TemplateId: "otp_code",
				Variables:  map[string]string{"code": "123456"},
			},
		}
		res, _, err := client.EmailsAPI.SendTemplateEmail(ctx).TemplateEmailRequest(req).Execute()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("status=%v", res.Status), nil
	})

	run("requestEmailConsent", func() (string, error) {
		probe := "smoke-go@example.com"
		req := goodsender.ConsentEmailRequest{
			Domain: "example.com",
			Emails: []goodsender.ConsentEmailEntry{
				goodsender.StringAsConsentEmailEntry(&probe),
			},
		}
		res, _, err := client.EmailsAPI.RequestEmailConsent(ctx).ConsentEmailRequest(req).Execute()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("emails=%d", len(res.Emails)), nil
	})

	run("getEmailConsentStatus", func() (string, error) {
		res, _, err := client.EmailsAPI.GetEmailConsentStatus(ctx, "user@example.com").Domain("example.com").Execute()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("entries=%d", len(res)), nil
	})

	run("listEmailConsents", func() (string, error) {
		res, _, err := client.EmailsAPI.ListEmailConsents(ctx).Domain("example.com").Limit(50).Execute()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("emails=%d", len(res.Emails)), nil
	})

	run("listDomains", func() (string, error) {
		res, _, err := client.DomainsAPI.ListDomains(ctx).Limit(50).Execute()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("domains=%d", len(res.Domains)), nil
	})

	for _, r := range results {
		tag := "PASS"
		if !r.ok {
			tag = "FAIL"
		}
		fmt.Printf("%-4s  go      %-22s  %s\n", tag, r.method, r.detail)
	}
	failed := 0
	for _, r := range results {
		if !r.ok {
			failed++
		}
	}
	passed := len(results) - failed
	fmt.Printf("\n%d passed, %d failed\n", passed, failed)
	if failed > 0 {
		os.Exit(1)
	}
}
