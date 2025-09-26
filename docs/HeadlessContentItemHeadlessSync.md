# HeadlessContentItemHeadlessSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncToken** | Pointer to **int64** |  | [optional] 
**Items** | Pointer to [**[]HeadlessContentItem**](HeadlessContentItem.md) |  | [optional] 

## Methods

### NewHeadlessContentItemHeadlessSync

`func NewHeadlessContentItemHeadlessSync() *HeadlessContentItemHeadlessSync`

NewHeadlessContentItemHeadlessSync instantiates a new HeadlessContentItemHeadlessSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeadlessContentItemHeadlessSyncWithDefaults

`func NewHeadlessContentItemHeadlessSyncWithDefaults() *HeadlessContentItemHeadlessSync`

NewHeadlessContentItemHeadlessSyncWithDefaults instantiates a new HeadlessContentItemHeadlessSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyncToken

`func (o *HeadlessContentItemHeadlessSync) GetSyncToken() int64`

GetSyncToken returns the SyncToken field if non-nil, zero value otherwise.

### GetSyncTokenOk

`func (o *HeadlessContentItemHeadlessSync) GetSyncTokenOk() (*int64, bool)`

GetSyncTokenOk returns a tuple with the SyncToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncToken

`func (o *HeadlessContentItemHeadlessSync) SetSyncToken(v int64)`

SetSyncToken sets SyncToken field to given value.

### HasSyncToken

`func (o *HeadlessContentItemHeadlessSync) HasSyncToken() bool`

HasSyncToken returns a boolean if a field has been set.

### GetItems

`func (o *HeadlessContentItemHeadlessSync) GetItems() []HeadlessContentItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *HeadlessContentItemHeadlessSync) GetItemsOk() (*[]HeadlessContentItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *HeadlessContentItemHeadlessSync) SetItems(v []HeadlessContentItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *HeadlessContentItemHeadlessSync) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *HeadlessContentItemHeadlessSync) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *HeadlessContentItemHeadlessSync) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


