//go:build conformance

// State-aware conformance test for the Go SDK against the real dev API.
// Mirrors tests/runners/node/conformance.ts (same scenario IDs).
package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	goodsender "github.com/good-sender/go-sdk"
)

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		fmt.Fprintf(os.Stderr, "FATAL: %s is not set in .env.dev\n", key)
		os.Exit(2)
	}
	return v
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func randomSuffix(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

type result struct {
	id, name, status, detail string
}

var results []result

func record(id, name, status, detail string) {
	results = append(results, result{id, name, status, detail})
}

func scenario(id, name string, fn func() (bool, string)) {
	defer func() {
		if r := recover(); r != nil {
			record(id, name, "FAIL", fmt.Sprintf("panic: %v", r))
		}
	}()
	ok, detail := fn()
	if ok {
		record(id, name, "PASS", detail)
	} else {
		record(id, name, "FAIL", detail)
	}
}

func skipScenario(id, name string) {
	record(id, name, "SKIP", "destructive — set ALLOW_DESTRUCTIVE=1")
}

func errBodySnippet(err error) string {
	if err == nil {
		return ""
	}
	s := err.Error()
	if len(s) > 160 {
		return s[:160]
	}
	return s
}

func main() {
	baseURL := mustEnv("BASE_URL")
	apiKey := mustEnv("GOODSENDER_API_KEY")
	allowDestructive := os.Getenv("ALLOW_DESTRUCTIVE") == "1"

	verifiedDomain := mustEnv("VERIFIED_SENDER_DOMAIN")
	verifiedEmail := mustEnv("VERIFIED_SENDER_EMAIL")
	verifiedName := envOr("VERIFIED_SENDER_NAME", "GoodSender SDK Tests")
	unverifiedDomain := mustEnv("UNVERIFIED_SENDER_DOMAIN")
	unverifiedEmail := mustEnv("UNVERIFIED_SENDER_EMAIL")
	granted1 := mustEnv("RECIPIENT_GRANTED_1")
	granted2 := mustEnv("RECIPIENT_GRANTED_2")
	denied1 := mustEnv("RECIPIENT_DENIED_1")
	denied2 := mustEnv("RECIPIENT_DENIED_2")
	templateID := mustEnv("TEMPLATE_ID")

	runTag := fmt.Sprintf("sdk-%x-%s", time.Now().Unix(), randomSuffix(3))
	fresh1 := fmt.Sprintf("%s-1@%s", runTag, verifiedDomain)
	fresh2 := fmt.Sprintf("%s-2@%s", runTag, verifiedDomain)

	cfg := goodsender.NewConfiguration()
	cfg.Servers = goodsender.ServerConfigurations{{URL: baseURL}}
	cfg.AddDefaultHeader("Authorization", "Bearer "+apiKey)
	client := goodsender.NewAPIClient(cfg)
	ctx := newCtx()

	// ─── Read-only (R1–R6) ──────────────────────────────────────

	scenario("R1", "listDomains returns both fixtures with correct verification flags", func() (bool, string) {
		res, _, err := client.DomainsAPI.ListDomains(ctx).Limit(100).Execute()
		if err != nil {
			return false, errBodySnippet(err)
		}
		byName := map[string]goodsender.Domain{}
		for _, d := range res.Domains {
			byName[d.Domain] = d
		}
		v, vOk := byName[verifiedDomain]
		u, uOk := byName[unverifiedDomain]
		if !vOk {
			return false, fmt.Sprintf("%s not in listDomains response", verifiedDomain)
		}
		if !uOk {
			return false, fmt.Sprintf("%s not in listDomains response", unverifiedDomain)
		}
		if !v.Verification.Verified {
			return false, fmt.Sprintf("%s has verification.verified=false; should be true", verifiedDomain)
		}
		if u.Verification.Verified {
			return false, fmt.Sprintf("%s has verification.verified=true; should be false", unverifiedDomain)
		}
		return true, fmt.Sprintf("domains=%d, verified=true, unverified=false", len(res.Domains))
	})

	scenario("R2", "getEmailConsentStatus returns granted for approved recipient", func() (bool, string) {
		res, _, err := client.EmailsAPI.GetEmailConsentStatus(ctx, granted1).Domain(verifiedDomain).Execute()
		if err != nil {
			return false, errBodySnippet(err)
		}
		var found *goodsender.EmailAccount
		for i := range res {
			if res[i].Domain == verifiedDomain {
				found = &res[i]
				break
			}
		}
		if found == nil {
			return false, fmt.Sprintf("no entry for domain=%s", verifiedDomain)
		}
		if found.ConsentStatus != "granted" {
			return false, fmt.Sprintf("consentStatus=%s, expected granted", found.ConsentStatus)
		}
		return true, fmt.Sprintf("consentStatus=%s", found.ConsentStatus)
	})

	scenario("R3", "getEmailConsentStatus returns denied for rejected recipient", func() (bool, string) {
		res, _, err := client.EmailsAPI.GetEmailConsentStatus(ctx, denied1).Domain(verifiedDomain).Execute()
		if err != nil {
			return false, errBodySnippet(err)
		}
		var found *goodsender.EmailAccount
		for i := range res {
			if res[i].Domain == verifiedDomain {
				found = &res[i]
				break
			}
		}
		if found == nil {
			return false, fmt.Sprintf("no entry for domain=%s", verifiedDomain)
		}
		if found.ConsentStatus != "denied" {
			return false, fmt.Sprintf("consentStatus=%s, expected denied", found.ConsentStatus)
		}
		return true, fmt.Sprintf("consentStatus=%s", found.ConsentStatus)
	})

	scenario("R4", "getEmailConsentStatus returns 404 for unknown recipient", func() (bool, string) {
		probe := fmt.Sprintf("%s-r4-probe@%s", runTag, verifiedDomain)
		_, httpResp, err := client.EmailsAPI.GetEmailConsentStatus(ctx, probe).Domain(verifiedDomain).Execute()
		if err == nil {
			return false, fmt.Sprintf("expected 404, got %d", httpResp.StatusCode)
		}
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			return true, fmt.Sprintf("404 (probe=%s)", probe)
		}
		return false, fmt.Sprintf("expected 404, got status=%d %s", statusOf(httpResp), errBodySnippet(err))
	})

	scenario("R5", "listEmailConsents for verified domain includes all 4 fixtures", func() (bool, string) {
		collected := map[string]bool{}
		var cursor string
		var total int
		for p := 0; p < 20; p++ {
			req := client.EmailsAPI.ListEmailConsents(ctx).Domain(verifiedDomain).Limit(100)
			if cursor != "" {
				req = req.Cursor(cursor)
			}
			res, _, err := req.Execute()
			if err != nil {
				return false, errBodySnippet(err)
			}
			for _, e := range res.Emails {
				collected[e.Email] = true
				total++
			}
			if res.NextCursor == nil || *res.NextCursor == "" {
				break
			}
			cursor = *res.NextCursor
		}
		var missing []string
		for _, e := range []string{granted1, granted2, denied1, denied2} {
			if !collected[e] {
				missing = append(missing, e)
			}
		}
		if len(missing) > 0 {
			return false, fmt.Sprintf("missing from listEmailConsents: %s", strings.Join(missing, ", "))
		}
		return true, fmt.Sprintf("%d entries scanned; all 4 fixtures present", total)
	})

	scenario("R6", "listEmailConsents with consentStatus=granted filter excludes denied", func() (bool, string) {
		collected := map[string]bool{}
		statuses := map[string]bool{}
		var cursor string
		var pages int
		for p := 0; p < 20; p++ {
			req := client.EmailsAPI.ListEmailConsents(ctx).Domain(verifiedDomain).Limit(100).ConsentStatus("granted")
			if cursor != "" {
				req = req.Cursor(cursor)
			}
			res, _, err := req.Execute()
			if err != nil {
				return false, errBodySnippet(err)
			}
			pages++
			for _, e := range res.Emails {
				collected[e.Email] = true
				statuses[e.ConsentStatus] = true
			}
			if res.NextCursor == nil || *res.NextCursor == "" {
				break
			}
			cursor = *res.NextCursor
		}
		var nonGranted []string
		for s := range statuses {
			if s != "granted" {
				nonGranted = append(nonGranted, s)
			}
		}
		if len(nonGranted) > 0 {
			sort.Strings(nonGranted)
			return false, "filter leaked non-granted statuses: " + strings.Join(nonGranted, ",")
		}
		var missing []string
		for _, e := range []string{granted1, granted2} {
			if !collected[e] {
				missing = append(missing, e)
			}
		}
		if len(missing) > 0 {
			var sample []string
			for k := range collected {
				sample = append(sample, k)
				if len(sample) >= 5 {
					break
				}
			}
			sampleStr := strings.Join(sample, ", ")
			if sampleStr == "" {
				sampleStr = "(none)"
			}
			return false, fmt.Sprintf("filter returned %d entries across %d page(s); missing=%s; sample=[%s]", len(collected), pages, strings.Join(missing, ","), sampleStr)
		}
		if collected[denied1] || collected[denied2] {
			return false, "denied fixtures leaked into granted filter"
		}
		return true, fmt.Sprintf("%d granted entries; denied fixtures absent", len(collected))
	})

	// ─── Destructive (D1–D6, E1–E5) ────────────────────────────

	if allowDestructive {
		scenario("D1", "sendEmail to 2 granted recipients delivers both", func() (bool, string) {
			textBody := "Conformance D1 — granted recipients."
			req := goodsender.SendEmailRequest{Emails: []goodsender.SendEmail{{
				From:        goodsender.Address{Email: verifiedEmail, Name: &verifiedName},
				To:          []goodsender.Address{{Email: granted1}, {Email: granted2}},
				Subject:     fmt.Sprintf("SDK conformance D1 %s", runTag),
				TextContent: &textBody,
			}}}
			res, _, err := client.EmailsAPI.SendEmail(ctx).SendEmailRequest(req).Execute()
			if err != nil {
				return false, errBodySnippet(err)
			}
			if res.Sent == 2 && res.Declined == 0 {
				return true, fmt.Sprintf("sent=%d declined=%d", res.Sent, res.Declined)
			}
			return false, fmt.Sprintf("sent=%d declined=%d, expected 2/0", res.Sent, res.Declined)
		})

		scenario("D2", "sendEmail to 2 denied recipients declines both", func() (bool, string) {
			textBody := "D2"
			req := goodsender.SendEmailRequest{Emails: []goodsender.SendEmail{{
				From:        goodsender.Address{Email: verifiedEmail, Name: &verifiedName},
				To:          []goodsender.Address{{Email: denied1}, {Email: denied2}},
				Subject:     fmt.Sprintf("SDK conformance D2 %s", runTag),
				TextContent: &textBody,
			}}}
			res, _, err := client.EmailsAPI.SendEmail(ctx).SendEmailRequest(req).Execute()
			if err != nil {
				return false, errBodySnippet(err)
			}
			if res.Sent == 0 && res.Declined == 2 {
				return true, fmt.Sprintf("sent=%d declined=%d", res.Sent, res.Declined)
			}
			return false, fmt.Sprintf("sent=%d declined=%d, expected 0/2", res.Sent, res.Declined)
		})

		scenario("D3", "sendEmail granted+denied mix splits correctly", func() (bool, string) {
			textBody := "D3"
			req := goodsender.SendEmailRequest{Emails: []goodsender.SendEmail{{
				From:        goodsender.Address{Email: verifiedEmail, Name: &verifiedName},
				To:          []goodsender.Address{{Email: granted1}, {Email: denied1}},
				Subject:     fmt.Sprintf("SDK conformance D3 %s", runTag),
				TextContent: &textBody,
			}}}
			res, _, err := client.EmailsAPI.SendEmail(ctx).SendEmailRequest(req).Execute()
			if err != nil {
				return false, errBodySnippet(err)
			}
			if res.Sent == 1 && res.Declined == 1 {
				return true, fmt.Sprintf("sent=%d declined=%d", res.Sent, res.Declined)
			}
			return false, fmt.Sprintf("sent=%d declined=%d, expected 1/1", res.Sent, res.Declined)
		})

		scenario("D4", "sendTemplateEmail to granted recipient returns status=sent", func() (bool, string) {
			req := goodsender.TemplateEmailRequest{
				From:     goodsender.Address{Email: verifiedEmail, Name: &verifiedName},
				To:       goodsender.Address{Email: granted1},
				Subject:  fmt.Sprintf("SDK conformance D4 %s", runTag),
				Template: goodsender.TemplateEmailRequestTemplate{TemplateId: templateID, Variables: map[string]string{}},
			}
			res, _, err := client.EmailsAPI.SendTemplateEmail(ctx).TemplateEmailRequest(req).Execute()
			if err != nil {
				return false, errBodySnippet(err)
			}
			s := fmt.Sprintf("%v", res.Status)
			if s == "sent" {
				return true, "status=sent"
			}
			return false, fmt.Sprintf("status=%s, expected sent", s)
		})

		scenario("D5", "sendTemplateEmail to denied recipient returns status=declined", func() (bool, string) {
			req := goodsender.TemplateEmailRequest{
				From:     goodsender.Address{Email: verifiedEmail, Name: &verifiedName},
				To:       goodsender.Address{Email: denied1},
				Subject:  fmt.Sprintf("SDK conformance D5 %s", runTag),
				Template: goodsender.TemplateEmailRequestTemplate{TemplateId: templateID, Variables: map[string]string{}},
			}
			res, _, err := client.EmailsAPI.SendTemplateEmail(ctx).TemplateEmailRequest(req).Execute()
			if err != nil {
				return false, errBodySnippet(err)
			}
			s := fmt.Sprintf("%v", res.Status)
			if s == "declined" {
				return true, "status=declined"
			}
			return false, fmt.Sprintf("status=%s, expected declined", s)
		})

		scenario("D6", "requestEmailConsent registers 2 fresh addresses", func() (bool, string) {
			name1, name2 := "Fresh 1", "Fresh 2"
			req := goodsender.ConsentEmailRequest{
				Domain: verifiedDomain,
				Emails: []goodsender.ConsentEmailEntry{
					goodsender.ConsentEmailRecipientAsConsentEmailEntry(&goodsender.ConsentEmailRecipient{Email: fresh1, Name: *goodsender.NewNullableString(&name1)}),
					goodsender.ConsentEmailRecipientAsConsentEmailEntry(&goodsender.ConsentEmailRecipient{Email: fresh2, Name: *goodsender.NewNullableString(&name2)}),
				},
			}
			res, _, err := client.EmailsAPI.RequestEmailConsent(ctx).ConsentEmailRequest(req).Execute()
			if err != nil {
				return false, errBodySnippet(err)
			}
			if len(res.Emails) != 2 {
				return false, fmt.Sprintf("expected 2 entries, got %d", len(res.Emails))
			}
			statuses := []string{}
			for _, e := range res.Emails {
				statuses = append(statuses, e.ConsentStatus)
			}
			sort.Strings(statuses)
			return true, fmt.Sprintf("2 fresh addresses; statuses=[%s] %s %s", strings.Join(statuses, ","), fresh1, fresh2)
		})

		scenario("E1", "sendEmail from unverified domain is rejected", func() (bool, string) {
			textBody := "should be rejected"
			req := goodsender.SendEmailRequest{Emails: []goodsender.SendEmail{{
				From:        goodsender.Address{Email: unverifiedEmail},
				To:          []goodsender.Address{{Email: granted1}},
				Subject:     fmt.Sprintf("SDK conformance E1 %s", runTag),
				TextContent: &textBody,
			}}}
			res, httpResp, err := client.EmailsAPI.SendEmail(ctx).SendEmailRequest(req).Execute()
			if err == nil {
				return false, fmt.Sprintf("expected 4xx, got %d sent=%d", httpResp.StatusCode, res.Sent)
			}
			code := statusOf(httpResp)
			if code >= 400 && code < 500 {
				return true, fmt.Sprintf("%d %s", code, errBodySnippet(err))
			}
			return false, fmt.Sprintf("expected 4xx, got %d %s", code, errBodySnippet(err))
		})

		scenario("E2", "sendTemplateEmail from unverified domain is rejected", func() (bool, string) {
			req := goodsender.TemplateEmailRequest{
				From:     goodsender.Address{Email: unverifiedEmail},
				To:       goodsender.Address{Email: granted1},
				Subject:  fmt.Sprintf("SDK conformance E2 %s", runTag),
				Template: goodsender.TemplateEmailRequestTemplate{TemplateId: templateID, Variables: map[string]string{}},
			}
			res, httpResp, err := client.EmailsAPI.SendTemplateEmail(ctx).TemplateEmailRequest(req).Execute()
			if err == nil {
				return false, fmt.Sprintf("expected 4xx, got %d status=%v", httpResp.StatusCode, res.Status)
			}
			code := statusOf(httpResp)
			if code >= 400 && code < 500 {
				return true, fmt.Sprintf("%d %s", code, errBodySnippet(err))
			}
			return false, fmt.Sprintf("expected 4xx, got %d %s", code, errBodySnippet(err))
		})

		scenario("E3", "sendTemplateEmail with bogus template_id returns 404", func() (bool, string) {
			bad := fmt.Sprintf("%s-does-not-exist", runTag)
			req := goodsender.TemplateEmailRequest{
				From:     goodsender.Address{Email: verifiedEmail},
				To:       goodsender.Address{Email: granted1},
				Subject:  fmt.Sprintf("SDK conformance E3 %s", runTag),
				Template: goodsender.TemplateEmailRequestTemplate{TemplateId: bad, Variables: map[string]string{}},
			}
			res, httpResp, err := client.EmailsAPI.SendTemplateEmail(ctx).TemplateEmailRequest(req).Execute()
			if err == nil {
				return false, fmt.Sprintf("expected 404, got %d status=%v", httpResp.StatusCode, res.Status)
			}
			code := statusOf(httpResp)
			if code == http.StatusNotFound {
				return true, fmt.Sprintf("404 %s", errBodySnippet(err))
			}
			return false, fmt.Sprintf("expected 404, got %d %s", code, errBodySnippet(err))
		})

		scenario("E4", "requestEmailConsent for unverified domain is rejected", func() (bool, string) {
			fresh := fmt.Sprintf("%s-e4-target@example.com", runTag)
			req := goodsender.ConsentEmailRequest{
				Domain: unverifiedDomain,
				Emails: []goodsender.ConsentEmailEntry{goodsender.StringAsConsentEmailEntry(&fresh)},
			}
			res, httpResp, err := client.EmailsAPI.RequestEmailConsent(ctx).ConsentEmailRequest(req).Execute()
			if err == nil {
				return false, fmt.Sprintf("expected 4xx, got %d emails=%d", httpResp.StatusCode, len(res.Emails))
			}
			code := statusOf(httpResp)
			if code >= 400 && code < 500 {
				return true, fmt.Sprintf("%d %s", code, errBodySnippet(err))
			}
			return false, fmt.Sprintf("expected 4xx, got %d %s", code, errBodySnippet(err))
		})

		scenario("E5", "listEmailConsents for non-existent domain", func() (bool, string) {
			bogus := fmt.Sprintf("not-a-real-domain-%s.invalid", runTag)
			res, httpResp, err := client.EmailsAPI.ListEmailConsents(ctx).Domain(bogus).Limit(1).Execute()
			if err == nil {
				return true, fmt.Sprintf("200 emails=%d (no error path for unknown domain)", len(res.Emails))
			}
			code := statusOf(httpResp)
			if code >= 400 && code < 500 {
				return true, fmt.Sprintf("%d %s", code, errBodySnippet(err))
			}
			return false, fmt.Sprintf("unexpected: %d %s", code, errBodySnippet(err))
		})
	} else {
		for _, s := range [][2]string{
			{"D1", "sendEmail to 2 granted recipients"},
			{"D2", "sendEmail to 2 denied recipients"},
			{"D3", "sendEmail granted+denied mix"},
			{"D4", "sendTemplateEmail to granted"},
			{"D5", "sendTemplateEmail to denied"},
			{"D6", "requestEmailConsent for 2 fresh addresses"},
			{"E1", "sendEmail from unverified domain rejected"},
			{"E2", "sendTemplateEmail from unverified domain rejected"},
			{"E3", "sendTemplateEmail with bogus template_id"},
			{"E4", "requestEmailConsent for unverified domain rejected"},
			{"E5", "listEmailConsents for non-existent domain"},
		} {
			skipScenario(s[0], s[1])
		}
	}

	for _, r := range results {
		nameTruncated := r.name
		if len(nameTruncated) > 58 {
			nameTruncated = nameTruncated[:58]
		}
		fmt.Printf("%-4s  go      %s  %-58s  %s\n", r.status, r.id, nameTruncated, r.detail)
	}
	passed, failed, skipped := 0, 0, 0
	for _, r := range results {
		switch r.status {
		case "PASS":
			passed++
		case "FAIL":
			failed++
		case "SKIP":
			skipped++
		}
	}
	fmt.Printf("\n%d passed, %d failed, %d skipped\n", passed, failed, skipped)
	if allowDestructive {
		fmt.Printf("\nDestructive run created consent records for cleanup:\n  %s\n  %s\n", fresh1, fresh2)
	}
	if failed > 0 {
		os.Exit(1)
	}
}

func newCtx() context.Context { return context.Background() }

func statusOf(resp *http.Response) int {
	if resp == nil {
		return 0
	}
	return resp.StatusCode
}

