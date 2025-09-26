# \ItemAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ItemIdGet**](ItemAPI.md#ItemIdGet) | **Get** /{guid}/{apitype}/{locale}/item/{id} | Gets the details of a content item by ther Content ID.



## ItemIdGet

> HeadlessContentItem ItemIdGet(ctx, guid, apitype, locale, id).ContentLinkDepth(contentLinkDepth).ExpandAllContentLinks(expandAllContentLinks).Execute()

Gets the details of a content item by ther Content ID.

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
	apitype := openapiclient.APIType("preview") // APIType | The API type (preview, fetch).
	locale := "locale_example" // string | The locale code you want to retrieve content for.
	id := int32(56) // int32 | The contentID of the requested item in this locale.
	contentLinkDepth := int32(56) // int32 | The maximum level to expand linked content from this item (optional) (default to 1)
	expandAllContentLinks := true // bool | Whether or not to expand entire linked content references, includings lists and items that are rendered in the CMS as Grid or Link. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemAPI.ItemIdGet(context.Background(), guid, apitype, locale, id).ContentLinkDepth(contentLinkDepth).ExpandAllContentLinks(expandAllContentLinks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAPI.ItemIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ItemIdGet`: HeadlessContentItem
	fmt.Fprintf(os.Stdout, "Response from `ItemAPI.ItemIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The instance GUID, available from the API Keys section. | 
**apitype** | [**APIType**](.md) | The API type (preview, fetch). | 
**locale** | **string** | The locale code you want to retrieve content for. | 
**id** | **int32** | The contentID of the requested item in this locale. | 

### Other Parameters

Other parameters are passed through a pointer to a LocaleItemIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **contentLinkDepth** | **int32** | The maximum level to expand linked content from this item | [default to 1]
 **expandAllContentLinks** | **bool** | Whether or not to expand entire linked content references, includings lists and items that are rendered in the CMS as Grid or Link. | [default to false]

### Return type

[**HeadlessContentItem**](HeadlessContentItem.md)

### Authorization

[APIKeyAuthorization](../README.md#APIKeyAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

