# HeadlessContentListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]HeadlessContentItem**](HeadlessContentItem.md) | The paginated items returned from the query. | [optional] 
**TotalCount** | Pointer to **int32** | The total number of items in the list (excluding any pagination). | [optional] 

## Methods

### NewHeadlessContentListResponse

`func NewHeadlessContentListResponse() *HeadlessContentListResponse`

NewHeadlessContentListResponse instantiates a new HeadlessContentListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeadlessContentListResponseWithDefaults

`func NewHeadlessContentListResponseWithDefaults() *HeadlessContentListResponse`

NewHeadlessContentListResponseWithDefaults instantiates a new HeadlessContentListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *HeadlessContentListResponse) GetItems() []HeadlessContentItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *HeadlessContentListResponse) GetItemsOk() (*[]HeadlessContentItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *HeadlessContentListResponse) SetItems(v []HeadlessContentItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *HeadlessContentListResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *HeadlessContentListResponse) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *HeadlessContentListResponse) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetTotalCount

`func (o *HeadlessContentListResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *HeadlessContentListResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *HeadlessContentListResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *HeadlessContentListResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


