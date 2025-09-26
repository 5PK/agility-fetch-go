# \ListAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListReferenceNameGet**](ListAPI.md#ListReferenceNameGet) | **Get** /{guid}/{apitype}/{locale}/list/{referenceName} | Retrieves a list of content items by reference name.



## ListReferenceNameGet

> HeadlessContentListResponse ListReferenceNameGet(ctx, guid, apitype, locale, referenceName).ContentLinkDepth(contentLinkDepth).ExpandAllContentLinks(expandAllContentLinks).Fields(fields).Take(take).Skip(skip).Filter(filter).Sort(sort).Direction(direction).Execute()

Retrieves a list of content items by reference name.

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
	apitype := agilityClient.APIType("preview") // APIType | The API type (preview, fetch).
	locale := "locale_example" // string | The locale code you want to retrieve content for.
	referenceName := "referenceName_example" // string | The unique reference name of the content list you wish to retrieve in the current locale. Reference names must be ALL lowercase.
	contentLinkDepth := int32(56) // int32 | [Optional] The depth of list items. Maximum allowed is 5. (optional)
	expandAllContentLinks := true // bool | [Optional] Whether or not to expand entire linked content references, includings lists and items that are rendered in the CMS as Grid or Link. (optional)
	fields := "fields_example" // string | [Optional] A comma separated list of the fields to return. (optional)
	take := int32(56) // int32 | [Optional] The maximum number of items to retrieve in this request. Default is 10. Maximum allowed is 250. (optional)
	skip := int32(56) // int32 | [Optional] The number of items to skip from the list. Default is 0. Used for implementing pagination. (optional)
	filter := "filter_example" // string | [Optional] The filter you wish to apply to the list query. Supports [eq (Equal To), ne (Not Equal), lt (Less Than), lte (Less Than or Equal), gt (Greater Than), gte (Greater Than or Equal)]. Example value: `fields.title[eq]\"some title\" or fields.details[like]\"some text\"` (optional)
	sort := "sort_example" // string | [Optional] The field to sort the results by. Example fields.title or properties.created, seo.metaDescription (optional)
	direction := "direction_example" // string | [Optional] The direction to sort the results by. Default is asc. Valid values are asc, desc. (optional)

	configuration := agilityClient.NewConfiguration()
	apiClient := agilityClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListAPI.ListReferenceNameGet(context.Background(), guid, apitype, locale, referenceName).ContentLinkDepth(contentLinkDepth).ExpandAllContentLinks(expandAllContentLinks).Fields(fields).Take(take).Skip(skip).Filter(filter).Sort(sort).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAPI.ListReferenceNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReferenceNameGet`: HeadlessContentListResponse
	fmt.Fprintf(os.Stdout, "Response from `ListAPI.ListReferenceNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The instance GUID, available from the API Keys section. | 
**apitype** | [**APIType**](.md) | The API type (preview, fetch). | 
**locale** | **string** | The locale code you want to retrieve content for. | 
**referenceName** | **string** | The unique reference name of the content list you wish to retrieve in the current locale. Reference names must be ALL lowercase. | 

### Other Parameters

Other parameters are passed through a pointer to a LocaleListReferenceNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **contentLinkDepth** | **int32** | [Optional] The depth of list items. Maximum allowed is 5. | 
 **expandAllContentLinks** | **bool** | [Optional] Whether or not to expand entire linked content references, includings lists and items that are rendered in the CMS as Grid or Link. | 
 **fields** | **string** | [Optional] A comma separated list of the fields to return. | 
 **take** | **int32** | [Optional] The maximum number of items to retrieve in this request. Default is 10. Maximum allowed is 250. | 
 **skip** | **int32** | [Optional] The number of items to skip from the list. Default is 0. Used for implementing pagination. | 
 **filter** | **string** | [Optional] The filter you wish to apply to the list query. Supports [eq (Equal To), ne (Not Equal), lt (Less Than), lte (Less Than or Equal), gt (Greater Than), gte (Greater Than or Equal)]. Example value: &#x60;fields.title[eq]\&quot;some title\&quot; or fields.details[like]\&quot;some text\&quot;&#x60; | 
 **sort** | **string** | [Optional] The field to sort the results by. Example fields.title or properties.created, seo.metaDescription | 
 **direction** | **string** | [Optional] The direction to sort the results by. Default is asc. Valid values are asc, desc. | 

### Return type

[**HeadlessContentListResponse**](HeadlessContentListResponse.md)

### Authorization

[APIKeyAuthorization](../README.md#APIKeyAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

