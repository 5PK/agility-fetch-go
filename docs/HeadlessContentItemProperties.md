# HeadlessContentItemProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **int32** | The state of this content item. 1 &#x3D; Staging, 2 &#x3D; Published, 3 &#x3D; Deleted, 4 &#x3D; Approved, 5 &#x3D; AwaitingApproval, 6 &#x3D; Declined, 7 &#x3D; Unpublished | [optional] 
**Modified** | Pointer to **string** | The date/time the item was last modified | [optional] 
**Created** | Pointer to **string** | The date/time this item was last modified. | [optional] 
**VersionID** | Pointer to **int32** | The unique versionID of this content item. | [optional] 
**ReferenceName** | Pointer to **NullableString** | The reference name of this item is part of. This may be the reference name of a content list, item, or module. | [optional] 
**DefinitionName** | Pointer to **NullableString** | The reference name of the content/module definition this item is based on. | [optional] 
**SyncToken** | Pointer to **int64** | Used for sync tracking. | [optional] 

## Methods

### NewHeadlessContentItemProperties

`func NewHeadlessContentItemProperties() *HeadlessContentItemProperties`

NewHeadlessContentItemProperties instantiates a new HeadlessContentItemProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeadlessContentItemPropertiesWithDefaults

`func NewHeadlessContentItemPropertiesWithDefaults() *HeadlessContentItemProperties`

NewHeadlessContentItemPropertiesWithDefaults instantiates a new HeadlessContentItemProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *HeadlessContentItemProperties) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HeadlessContentItemProperties) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HeadlessContentItemProperties) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *HeadlessContentItemProperties) HasState() bool`

HasState returns a boolean if a field has been set.

### GetModified

`func (o *HeadlessContentItemProperties) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *HeadlessContentItemProperties) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *HeadlessContentItemProperties) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *HeadlessContentItemProperties) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetCreated

`func (o *HeadlessContentItemProperties) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *HeadlessContentItemProperties) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *HeadlessContentItemProperties) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *HeadlessContentItemProperties) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetVersionID

`func (o *HeadlessContentItemProperties) GetVersionID() int32`

GetVersionID returns the VersionID field if non-nil, zero value otherwise.

### GetVersionIDOk

`func (o *HeadlessContentItemProperties) GetVersionIDOk() (*int32, bool)`

GetVersionIDOk returns a tuple with the VersionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionID

`func (o *HeadlessContentItemProperties) SetVersionID(v int32)`

SetVersionID sets VersionID field to given value.

### HasVersionID

`func (o *HeadlessContentItemProperties) HasVersionID() bool`

HasVersionID returns a boolean if a field has been set.

### GetReferenceName

`func (o *HeadlessContentItemProperties) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *HeadlessContentItemProperties) GetReferenceNameOk() (*string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceName

`func (o *HeadlessContentItemProperties) SetReferenceName(v string)`

SetReferenceName sets ReferenceName field to given value.

### HasReferenceName

`func (o *HeadlessContentItemProperties) HasReferenceName() bool`

HasReferenceName returns a boolean if a field has been set.

### SetReferenceNameNil

`func (o *HeadlessContentItemProperties) SetReferenceNameNil(b bool)`

 SetReferenceNameNil sets the value for ReferenceName to be an explicit nil

### UnsetReferenceName
`func (o *HeadlessContentItemProperties) UnsetReferenceName()`

UnsetReferenceName ensures that no value is present for ReferenceName, not even an explicit nil
### GetDefinitionName

`func (o *HeadlessContentItemProperties) GetDefinitionName() string`

GetDefinitionName returns the DefinitionName field if non-nil, zero value otherwise.

### GetDefinitionNameOk

`func (o *HeadlessContentItemProperties) GetDefinitionNameOk() (*string, bool)`

GetDefinitionNameOk returns a tuple with the DefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionName

`func (o *HeadlessContentItemProperties) SetDefinitionName(v string)`

SetDefinitionName sets DefinitionName field to given value.

### HasDefinitionName

`func (o *HeadlessContentItemProperties) HasDefinitionName() bool`

HasDefinitionName returns a boolean if a field has been set.

### SetDefinitionNameNil

`func (o *HeadlessContentItemProperties) SetDefinitionNameNil(b bool)`

 SetDefinitionNameNil sets the value for DefinitionName to be an explicit nil

### UnsetDefinitionName
`func (o *HeadlessContentItemProperties) UnsetDefinitionName()`

UnsetDefinitionName ensures that no value is present for DefinitionName, not even an explicit nil
### GetSyncToken

`func (o *HeadlessContentItemProperties) GetSyncToken() int64`

GetSyncToken returns the SyncToken field if non-nil, zero value otherwise.

### GetSyncTokenOk

`func (o *HeadlessContentItemProperties) GetSyncTokenOk() (*int64, bool)`

GetSyncTokenOk returns a tuple with the SyncToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncToken

`func (o *HeadlessContentItemProperties) SetSyncToken(v int64)`

SetSyncToken sets SyncToken field to given value.

### HasSyncToken

`func (o *HeadlessContentItemProperties) HasSyncToken() bool`

HasSyncToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


