# \SyncAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyncItemsGet**](SyncAPI.md#SyncItemsGet) | **Get** /{guid}/{apitype}/{locale}/sync/items | Retrieves all content items in a paged format.  Each call returns a sync token that should be persisted and passed into subsequent requests to maintain sync state.
[**SyncPagesGet**](SyncAPI.md#SyncPagesGet) | **Get** /{guid}/{apitype}/{locale}/sync/pages | Retrieves all pages items in a paged format.  Each call returns a sync token that should be persisted and passed into subsequent requests to maintain sync state.



## SyncItemsGet

> HeadlessContentItemHeadlessSync SyncItemsGet(ctx, guid, apitype, locale).SyncToken(syncToken).PageSize(pageSize).Execute()

Retrieves all content items in a paged format.  Each call returns a sync token that should be persisted and passed into subsequent requests to maintain sync state.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	agilityClient "github.com/5PK/agility-fetch-go"
)

func main() {
	guid := "guid_example" // string | The instance GUID, available from the API Keys section.
	apitype := agilityClient.APIType("preview") // APIType | The Type of API - fetch or preview.
	locale := "locale_example" // string | The locale code you want to sync content items for.
	syncToken := int64(789) // int64 | The token from the most recently synced value. (optional)
	pageSize := int32(56) // int32 | The number of items to return per set. (optional) (default to 500)

	configuration := agilityClient.NewConfiguration()
	apiClient := agilityClient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncAPI.SyncItemsGet(context.Background(), guid, apitype, locale).SyncToken(syncToken).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.SyncItemsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncItemsGet`: HeadlessContentItemHeadlessSync
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.SyncItemsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The instance GUID, available from the API Keys section. | 
**apitype** | [**APIType**](.md) | The Type of API - fetch or preview. | 
**locale** | **string** | The locale code you want to sync content items for. | 

### Other Parameters

Other parameters are passed through a pointer to a LocaleSyncItemsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **syncToken** | **int64** | The token from the most recently synced value. | 
 **pageSize** | **int32** | The number of items to return per set. | [default to 500]

### Return type

[**HeadlessContentItemHeadlessSync**](HeadlessContentItemHeadlessSync.md)

### Authorization

[APIKeyAuthorization](../README.md#APIKeyAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncPagesGet

> HeadlessContentPageHeadlessSync SyncPagesGet(ctx, guid, apitype, locale).SyncToken(syncToken).PageSize(pageSize).Execute()

Retrieves all pages items in a paged format.  Each call returns a sync token that should be persisted and passed into subsequent requests to maintain sync state.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	agilityClient "github.com/5PK/agility-fetch-go"
)

func main() {
	guid := "guid_example" // string | The instance GUID, available from the API Keys section.
	apitype := agilityClient.APIType("preview") // APIType | The Type of API - fetch or preview.
	locale := "locale_example" // string | The locale code you want to sync pages for.
	syncToken := int64(789) // int64 | The token from the most recently synced value. (optional)
	pageSize := int32(56) // int32 | The number of items to return per set. (optional) (default to 500)

	configuration := agilityClient.NewConfiguration()
	apiClient := agilityClient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncAPI.SyncPagesGet(context.Background(), guid, apitype, locale).SyncToken(syncToken).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.SyncPagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncPagesGet`: HeadlessContentPageHeadlessSync
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.SyncPagesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The instance GUID, available from the API Keys section. | 
**apitype** | [**APIType**](.md) | The Type of API - fetch or preview. | 
**locale** | **string** | The locale code you want to sync pages for. | 

### Other Parameters

Other parameters are passed through a pointer to a LocaleSyncPagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **syncToken** | **int64** | The token from the most recently synced value. | 
 **pageSize** | **int32** | The number of items to return per set. | [default to 500]

### Return type

[**HeadlessContentPageHeadlessSync**](HeadlessContentPageHeadlessSync.md)

### Authorization

[APIKeyAuthorization](../README.md#APIKeyAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

