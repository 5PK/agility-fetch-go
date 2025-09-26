# \SitemapAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitemapFlatChannelNameGet**](SitemapAPI.md#SitemapFlatChannelNameGet) | **Get** /{guid}/{apitype}/{locale}/sitemap/flat/{channelName} | Gets the sitemap, returned in a flat list, where the dictionary key is the page path. This method is ideal for page routing.
[**SitemapNestedChannelNameGet**](SitemapAPI.md#SitemapNestedChannelNameGet) | **Get** /{guid}/{apitype}/{locale}/sitemap/nested/{channelName} | Gets the sitemap as an array in a nested format, ideal for generating menus.



## SitemapFlatChannelNameGet

> map[string]HeadlessContentSiteMapItem SitemapFlatChannelNameGet(ctx, guid, apitype, locale, channelName).Execute()

Gets the sitemap, returned in a flat list, where the dictionary key is the page path. This method is ideal for page routing.

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
	locale := "locale_example" // string | The locale code you want to retreive content for
	channelName := "channelName_example" // string | The reference name of the digital channel of the sitemap to return. If you only have one channel, your channel reference name is likely 'website'.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitemapAPI.SitemapFlatChannelNameGet(context.Background(), guid, apitype, locale, channelName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitemapAPI.SitemapFlatChannelNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SitemapFlatChannelNameGet`: map[string]HeadlessContentSiteMapItem
	fmt.Fprintf(os.Stdout, "Response from `SitemapAPI.SitemapFlatChannelNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The instance GUID, available from the API Keys section. | 
**apitype** | [**APIType**](.md) | The Type of API - fetch or preview. | 
**locale** | **string** | The locale code you want to retreive content for | 
**channelName** | **string** | The reference name of the digital channel of the sitemap to return. If you only have one channel, your channel reference name is likely &#39;website&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a LocaleSitemapFlatChannelNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**map[string]HeadlessContentSiteMapItem**](HeadlessContentSiteMapItem.md)

### Authorization

[APIKeyAuthorization](../README.md#APIKeyAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitemapNestedChannelNameGet

> []HeadlessContentSiteMapNestedItem SitemapNestedChannelNameGet(ctx, guid, apitype, locale, channelName).Execute()

Gets the sitemap as an array in a nested format, ideal for generating menus.

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
	locale := "locale_example" // string | The locale code you want to retreive content for
	channelName := "channelName_example" // string | The reference name of the digital channel of the sitemap to return. If you only have one channel, your channel reference name is likely 'website'.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitemapAPI.SitemapNestedChannelNameGet(context.Background(), guid, apitype, locale, channelName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitemapAPI.SitemapNestedChannelNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SitemapNestedChannelNameGet`: []HeadlessContentSiteMapNestedItem
	fmt.Fprintf(os.Stdout, "Response from `SitemapAPI.SitemapNestedChannelNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The instance GUID, available from the API Keys section. | 
**apitype** | [**APIType**](.md) | The Type of API - fetch or preview. | 
**locale** | **string** | The locale code you want to retreive content for | 
**channelName** | **string** | The reference name of the digital channel of the sitemap to return. If you only have one channel, your channel reference name is likely &#39;website&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a LocaleSitemapNestedChannelNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**[]HeadlessContentSiteMapNestedItem**](HeadlessContentSiteMapNestedItem.md)

### Authorization

[APIKeyAuthorization](../README.md#APIKeyAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

