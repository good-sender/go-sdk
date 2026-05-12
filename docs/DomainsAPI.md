# \DomainsAPI

All URIs are relative to *https://api.goodsender.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDomains**](DomainsAPI.md#ListDomains) | **Get** /v1/domains | List domains



## ListDomains

> DomainListResponse ListDomains(ctx).Limit(limit).Cursor(cursor).Execute()

List domains



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
	limit := int32(56) // int32 | Maximum number of records to return. (optional) (default to 50)
	cursor := "cursor_example" // string | Cursor for pagination, returned as `nextCursor` from a previous response. (optional)

	configuration := goodsender.NewConfiguration()
	apiClient := goodsender.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.ListDomains(context.Background()).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.ListDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDomains`: DomainListResponse
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.ListDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of records to return. | [default to 50]
 **cursor** | **string** | Cursor for pagination, returned as &#x60;nextCursor&#x60; from a previous response. | 

### Return type

[**DomainListResponse**](DomainListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

