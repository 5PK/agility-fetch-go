# \ContentModelsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentmodelsGet**](ContentModelsAPI.md#ContentmodelsGet) | **Get** /{guid}/{apitype}/contentmodels | Returns content models for the Agility instance.



## ContentmodelsGet

> []AgilityDefinition ContentmodelsGet(ctx, guid, apitype).LastModifiedDate(lastModifiedDate).Execute()

Returns content models for the Agility instance.

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
	lastModifiedDate := "lastModifiedDate_example" // string | The last modified date/time that the Models list was updated, eg: 2020-09-24T10:00:00.00-04:00. (optional)

	configuration := agilityClient.NewConfiguration()
	apiClient := agilityClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentModelsAPI.ContentmodelsGet(context.Background(), guid, apitype).LastModifiedDate(lastModifiedDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentModelsAPI.ContentmodelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentmodelsGet`: []AgilityDefinition
	fmt.Fprintf(os.Stdout, "Response from `ContentModelsAPI.ContentmodelsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The instance GUID, available from the API Keys section. | 
**apitype** | [**APIType**](.md) | The API type (preview, fetch). | 

### Other Parameters

Other parameters are passed through a pointer to a ContentmodelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lastModifiedDate** | **string** | The last modified date/time that the Models list was updated, eg: 2020-09-24T10:00:00.00-04:00. | 

### Return type

[**[]AgilityDefinition**](AgilityDefinition.md)

### Authorization

[APIKeyAuthorization](../README.md#APIKeyAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

