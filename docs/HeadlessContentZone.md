# HeadlessContentZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | Pointer to **NullableString** | Name of the module | [optional] 
**Item** | Pointer to [**HeadlessContentItem**](HeadlessContentItem.md) |  | [optional] 

## Methods

### NewHeadlessContentZone

`func NewHeadlessContentZone() *HeadlessContentZone`

NewHeadlessContentZone instantiates a new HeadlessContentZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeadlessContentZoneWithDefaults

`func NewHeadlessContentZoneWithDefaults() *HeadlessContentZone`

NewHeadlessContentZoneWithDefaults instantiates a new HeadlessContentZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *HeadlessContentZone) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *HeadlessContentZone) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *HeadlessContentZone) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *HeadlessContentZone) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *HeadlessContentZone) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *HeadlessContentZone) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetItem

`func (o *HeadlessContentZone) GetItem() HeadlessContentItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *HeadlessContentZone) GetItemOk() (*HeadlessContentItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *HeadlessContentZone) SetItem(v HeadlessContentItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *HeadlessContentZone) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


