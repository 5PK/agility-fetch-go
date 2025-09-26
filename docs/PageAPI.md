# \PageAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PageChannelGet**](PageAPI.md#PageChannelGet) | **Get** /{guid}/{apitype}/{locale}/page/{channel} | Gets the details of a page by its Page path, in a specific channel.
[**PageIdGet**](PageAPI.md#PageIdGet) | **Get** /{guid}/{apitype}/{locale}/page/{id} | Gets the details of a page by its Page ID.



## PageChannelGet

> HeadlessContentPageByPath PageChannelGet(ctx, guid, apitype, locale, channel).Path(path).ContentLinkDepth(contentLinkDepth).ExpandAllContentLinks(expandAllContentLinks).Execute()

Gets the details of a page by its Page path, in a specific channel.

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
	channel := "channel_example" // string | The reference name of the digital channel of the sitemap to return. If you only have one channel, your channel reference name is likely 'website'.
	path := "path_example" // string | The path of the page (optional)
	contentLinkDepth := int32(56) // int32 | The maximum level to expand linked content from this item (optional) (default to 2)
	expandAllContentLinks := true // bool | Whether or not to expand entire linked content references, includings lists and items that are rendered in the CMS as Grid or Link and linked to modules or linked items on this page. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PageAPI.PageChannelGet(context.Background(), guid, apitype, locale, channel).Path(path).ContentLinkDepth(contentLinkDepth).ExpandAllContentLinks(expandAllContentLinks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.PageChannelGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageChannelGet`: HeadlessContentPageByPath
	fmt.Fprintf(os.Stdout, "Response from `PageAPI.PageChannelGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The instance GUID, available from the API Keys section. | 
**apitype** | [**APIType**](.md) | The API type (preview, fetch). | 
**locale** | **string** | The locale code you want to retrieve content for. | 
**channel** | **string** | The reference name of the digital channel of the sitemap to return. If you only have one channel, your channel reference name is likely &#39;website&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a LocalePageChannelGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **path** | **string** | The path of the page | 
 **contentLinkDepth** | **int32** | The maximum level to expand linked content from this item | [default to 2]
 **expandAllContentLinks** | **bool** | Whether or not to expand entire linked content references, includings lists and items that are rendered in the CMS as Grid or Link and linked to modules or linked items on this page. | [default to false]

### Return type

[**HeadlessContentPageByPath**](HeadlessContentPageByPath.md)

### Authorization

[APIKeyAuthorization](../README.md#APIKeyAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageIdGet

> HeadlessContentPage PageIdGet(ctx, guid, apitype, locale, id).ContentLinkDepth(contentLinkDepth).ExpandAllContentLinks(expandAllContentLinks).Execute()

Gets the details of a page by its Page ID.

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
	id := int32(56) // int32 | The unique page ID of the page you wish to retrieve in the current locale.
	contentLinkDepth := int32(56) // int32 | The maximum level to expand linked content from this item (optional) (default to 2)
	expandAllContentLinks := true // bool | Whether or not to expand entire linked content references, includings lists and items that are rendered in the CMS as Grid or Link and linked to modules or linked items on this page. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PageAPI.PageIdGet(context.Background(), guid, apitype, locale, id).ContentLinkDepth(contentLinkDepth).ExpandAllContentLinks(expandAllContentLinks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.PageIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageIdGet`: HeadlessContentPage
	fmt.Fprintf(os.Stdout, "Response from `PageAPI.PageIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The instance GUID, available from the API Keys section. | 
**apitype** | [**APIType**](.md) | The API type (preview, fetch). | 
**locale** | **string** | The locale code you want to retrieve content for. | 
**id** | **int32** | The unique page ID of the page you wish to retrieve in the current locale. | 

### Other Parameters

Other parameters are passed through a pointer to a PageIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **contentLinkDepth** | **int32** | The maximum level to expand linked content from this item | [default to 2]
 **expandAllContentLinks** | **bool** | Whether or not to expand entire linked content references, includings lists and items that are rendered in the CMS as Grid or Link and linked to modules or linked items on this page. | [default to false]

### Return type

[**HeadlessContentPage**](HeadlessContentPage.md)

### Authorization

[APIKeyAuthorization](../README.md#APIKeyAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

