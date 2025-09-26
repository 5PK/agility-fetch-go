# \GalleryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GalleryIdGet**](GalleryAPI.md#GalleryIdGet) | **Get** /{guid}/{apitype}/gallery/{id} | Gets the gallery by ID.



## GalleryIdGet

> HeadlessGallery GalleryIdGet(ctx, guid, apitype, id).Execute()

Gets the gallery by ID.

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
	id := int32(56) // int32 | The galleryID of the requested gallery.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GalleryAPI.GalleryIdGet(context.Background(), guid, apitype, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GalleryAPI.GalleryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GalleryIdGet`: HeadlessGallery
	fmt.Fprintf(os.Stdout, "Response from `GalleryAPI.GalleryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The instance GUID, available from the API Keys section. | 
**apitype** | [**APIType**](.md) | The Type of API - fetch or preview. | 
**id** | **int32** | The galleryID of the requested gallery. | 

### Other Parameters

Other parameters are passed through a pointer to a GalleryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**HeadlessGallery**](HeadlessGallery.md)

### Authorization

[APIKeyAuthorization](../README.md#APIKeyAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

