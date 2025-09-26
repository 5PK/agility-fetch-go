# \UrlRedirectionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UrlredirectionGet**](UrlRedirectionAPI.md#UrlredirectionGet) | **Get** /{guid}/{apitype}/urlredirection | Gets the list of all URL Redirections since a particular date/time. Persist the lastAccessDate that is returned and use that value to maintain state on subsequent requests.



## UrlredirectionGet

> []UrlRedirection UrlredirectionGet(ctx, guid, apitype).LastAccessDate(lastAccessDate).Execute()

Gets the list of all URL Redirections since a particular date/time. Persist the lastAccessDate that is returned and use that value to maintain state on subsequent requests.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
     
	agilityClient "agility-fetch-go"
)

func main() {
	guid := "guid_example" // string | The instance GUID, available from the API Keys section.
	apitype := openapiclient.APIType("preview") // APIType | The Type of API - fetch or preview.
	lastAccessDate := time.Now() // string | The last access date/time that the URL Redirections list was previously accessed, eg: 2020-09-24T10:00:00.00-04:00. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UrlRedirectionAPI.UrlredirectionGet(context.Background(), guid, apitype).LastAccessDate(lastAccessDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UrlRedirectionAPI.UrlredirectionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UrlredirectionGet`: []UrlRedirection
	fmt.Fprintf(os.Stdout, "Response from `UrlRedirectionAPI.UrlredirectionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The instance GUID, available from the API Keys section. | 
**apitype** | [**APIType**](.md) | The Type of API - fetch or preview. | 

### Other Parameters

Other parameters are passed through a pointer to a UrlredirectionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lastAccessDate** | **string** | The last access date/time that the URL Redirections list was previously accessed, eg: 2020-09-24T10:00:00.00-04:00. | 

### Return type

[**[]UrlRedirection**](UrlRedirection.md)

### Authorization

[APIKeyAuthorization](../README.md#APIKeyAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

