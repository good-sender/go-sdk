# \EmailsAPI

All URIs are relative to *https://api.goodsender.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEmailConsentStatus**](EmailsAPI.md#GetEmailConsentStatus) | **Get** /v1/emails/{email} | Get recipient consent status
[**ListEmailConsents**](EmailsAPI.md#ListEmailConsents) | **Get** /v1/emails | List email consent statuses
[**RequestEmailConsent**](EmailsAPI.md#RequestEmailConsent) | **Post** /v1/emails/consent | Request recipients&#39; consent to receive emails from your domain
[**SendEmail**](EmailsAPI.md#SendEmail) | **Post** /v1/emails/send | Send an email or a batch of emails
[**SendTemplateEmail**](EmailsAPI.md#SendTemplateEmail) | **Post** /v1/emails/template | Send a transactional email using a template



## GetEmailConsentStatus

> []EmailAccount GetEmailConsentStatus(ctx, email).Domain(domain).Execute()

Get recipient consent status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	goodsender "github.com/good-sender/go-sdk"
)

func main() {
	email := "user@example.com" // string | Email address to look up.
	domain := "example.com" // string | Optional sender domain to filter consent records by. When omitted, returns consent across all domains. (optional)

	configuration := goodsender.NewConfiguration()
	apiClient := goodsender.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailsAPI.GetEmailConsentStatus(context.Background(), email).Domain(domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.GetEmailConsentStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailConsentStatus`: []EmailAccount
	fmt.Fprintf(os.Stdout, "Response from `EmailsAPI.GetEmailConsentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Email address to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailConsentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domain** | **string** | Optional sender domain to filter consent records by. When omitted, returns consent across all domains. | 

### Return type

[**[]EmailAccount**](EmailAccount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailConsents

> EmailListResponse ListEmailConsents(ctx).Domain(domain).Limit(limit).Cursor(cursor).ConsentStatus(consentStatus).EngagementStatus(engagementStatus).Execute()

List email consent statuses



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	goodsender "github.com/good-sender/go-sdk"
)

func main() {
	domain := "example.com" // string | Sender domain to filter consent records by.
	limit := int32(56) // int32 | Maximum number of records to return. (optional) (default to 50)
	cursor := "cursor_example" // string | Cursor for pagination. (optional)
	consentStatus := "consentStatus_example" // string | Status of the recipient's consent for receiving emails. 'pending' = awaiting consent email send, 'requested' = consent email dispatched, 'failed' = consent email delivery failed, 'granted' = recipient consented to receive emails, 'denied' = recipient declined to receive emails. (optional)
	engagementStatus := "engagementStatus_example" // string | Status of the recipient's engagement with the emails. (optional)

	configuration := goodsender.NewConfiguration()
	apiClient := goodsender.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailsAPI.ListEmailConsents(context.Background()).Domain(domain).Limit(limit).Cursor(cursor).ConsentStatus(consentStatus).EngagementStatus(engagementStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.ListEmailConsents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEmailConsents`: EmailListResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailsAPI.ListEmailConsents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEmailConsentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | Sender domain to filter consent records by. | 
 **limit** | **int32** | Maximum number of records to return. | [default to 50]
 **cursor** | **string** | Cursor for pagination. | 
 **consentStatus** | **string** | Status of the recipient&#39;s consent for receiving emails. &#39;pending&#39; &#x3D; awaiting consent email send, &#39;requested&#39; &#x3D; consent email dispatched, &#39;failed&#39; &#x3D; consent email delivery failed, &#39;granted&#39; &#x3D; recipient consented to receive emails, &#39;denied&#39; &#x3D; recipient declined to receive emails. | 
 **engagementStatus** | **string** | Status of the recipient&#39;s engagement with the emails. | 

### Return type

[**EmailListResponse**](EmailListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestEmailConsent

> ConsentEmailResult RequestEmailConsent(ctx).ConsentEmailRequest(consentEmailRequest).Execute()

Request recipients' consent to receive emails from your domain



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	goodsender "github.com/good-sender/go-sdk"
)

func main() {
	consentEmailRequest := *goodsender.NewConsentEmailRequest("example.com", []goodsender.ConsentEmailEntry{goodsender.ConsentEmailEntry{ConsentEmailRecipient: goodsender.NewConsentEmailRecipient("user@example.com")}}) // ConsentEmailRequest | 

	configuration := goodsender.NewConfiguration()
	apiClient := goodsender.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailsAPI.RequestEmailConsent(context.Background()).ConsentEmailRequest(consentEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.RequestEmailConsent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestEmailConsent`: ConsentEmailResult
	fmt.Fprintf(os.Stdout, "Response from `EmailsAPI.RequestEmailConsent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestEmailConsentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consentEmailRequest** | [**ConsentEmailRequest**](ConsentEmailRequest.md) |  | 

### Return type

[**ConsentEmailResult**](ConsentEmailResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEmail

> SendEmailResponse SendEmail(ctx).SendEmailRequest(sendEmailRequest).Execute()

Send an email or a batch of emails



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	goodsender "github.com/good-sender/go-sdk"
)

func main() {
	sendEmailRequest := *goodsender.NewSendEmailRequest([]goodsender.SendEmail{*goodsender.NewSendEmail(*goodsender.NewAddress("Email_example"), []goodsender.Address{*goodsender.NewAddress("Email_example")}, "Subject_example")}) // SendEmailRequest | List of emails to send

	configuration := goodsender.NewConfiguration()
	apiClient := goodsender.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailsAPI.SendEmail(context.Background()).SendEmailRequest(sendEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.SendEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendEmail`: SendEmailResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailsAPI.SendEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendEmailRequest** | [**SendEmailRequest**](SendEmailRequest.md) | List of emails to send | 

### Return type

[**SendEmailResponse**](SendEmailResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTemplateEmail

> TemplateEmailResponse SendTemplateEmail(ctx).TemplateEmailRequest(templateEmailRequest).Execute()

Send a transactional email using a template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	goodsender "github.com/good-sender/go-sdk"
)

func main() {
	templateEmailRequest := *goodsender.NewTemplateEmailRequest(*goodsender.NewAddress("Email_example"), *goodsender.NewAddress("Email_example"), "Subject_example", *goodsender.NewTemplateEmailRequestTemplate("TemplateId_example")) // TemplateEmailRequest | Template email to send

	configuration := goodsender.NewConfiguration()
	apiClient := goodsender.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailsAPI.SendTemplateEmail(context.Background()).TemplateEmailRequest(templateEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.SendTemplateEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendTemplateEmail`: TemplateEmailResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailsAPI.SendTemplateEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendTemplateEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateEmailRequest** | [**TemplateEmailRequest**](TemplateEmailRequest.md) | Template email to send | 

### Return type

[**TemplateEmailResponse**](TemplateEmailResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

