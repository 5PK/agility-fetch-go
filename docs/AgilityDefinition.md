# AgilityDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**ModelId** | Pointer to **int32** |  | [optional] 
**Guid** | Pointer to **NullableString** |  | [optional] 
**LastModifiedDate** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**ReferenceName** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ContentReferenceNames** | Pointer to **[]string** |  | [optional] 
**Fields** | Pointer to [**[]AgilityField**](AgilityField.md) |  | [optional] 

## Methods

### NewAgilityDefinition

`func NewAgilityDefinition() *AgilityDefinition`

NewAgilityDefinition instantiates a new AgilityDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgilityDefinitionWithDefaults

`func NewAgilityDefinitionWithDefaults() *AgilityDefinition`

NewAgilityDefinitionWithDefaults instantiates a new AgilityDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgilityDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgilityDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgilityDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgilityDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AgilityDefinition) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AgilityDefinition) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetModelId

`func (o *AgilityDefinition) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AgilityDefinition) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AgilityDefinition) SetModelId(v int32)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *AgilityDefinition) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetGuid

`func (o *AgilityDefinition) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *AgilityDefinition) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *AgilityDefinition) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *AgilityDefinition) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *AgilityDefinition) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *AgilityDefinition) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetLastModifiedDate

`func (o *AgilityDefinition) GetLastModifiedDate() string`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *AgilityDefinition) GetLastModifiedDateOk() (*string, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *AgilityDefinition) SetLastModifiedDate(v string)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *AgilityDefinition) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetDisplayName

`func (o *AgilityDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AgilityDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AgilityDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AgilityDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AgilityDefinition) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AgilityDefinition) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetReferenceName

`func (o *AgilityDefinition) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *AgilityDefinition) GetReferenceNameOk() (*string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceName

`func (o *AgilityDefinition) SetReferenceName(v string)`

SetReferenceName sets ReferenceName field to given value.

### HasReferenceName

`func (o *AgilityDefinition) HasReferenceName() bool`

HasReferenceName returns a boolean if a field has been set.

### SetReferenceNameNil

`func (o *AgilityDefinition) SetReferenceNameNil(b bool)`

 SetReferenceNameNil sets the value for ReferenceName to be an explicit nil

### UnsetReferenceName
`func (o *AgilityDefinition) UnsetReferenceName()`

UnsetReferenceName ensures that no value is present for ReferenceName, not even an explicit nil
### GetName

`func (o *AgilityDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgilityDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgilityDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgilityDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AgilityDefinition) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AgilityDefinition) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetContentReferenceNames

`func (o *AgilityDefinition) GetContentReferenceNames() []string`

GetContentReferenceNames returns the ContentReferenceNames field if non-nil, zero value otherwise.

### GetContentReferenceNamesOk

`func (o *AgilityDefinition) GetContentReferenceNamesOk() (*[]string, bool)`

GetContentReferenceNamesOk returns a tuple with the ContentReferenceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentReferenceNames

`func (o *AgilityDefinition) SetContentReferenceNames(v []string)`

SetContentReferenceNames sets ContentReferenceNames field to given value.

### HasContentReferenceNames

`func (o *AgilityDefinition) HasContentReferenceNames() bool`

HasContentReferenceNames returns a boolean if a field has been set.

### SetContentReferenceNamesNil

`func (o *AgilityDefinition) SetContentReferenceNamesNil(b bool)`

 SetContentReferenceNamesNil sets the value for ContentReferenceNames to be an explicit nil

### UnsetContentReferenceNames
`func (o *AgilityDefinition) UnsetContentReferenceNames()`

UnsetContentReferenceNames ensures that no value is present for ContentReferenceNames, not even an explicit nil
### GetFields

`func (o *AgilityDefinition) GetFields() []AgilityField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AgilityDefinition) GetFieldsOk() (*[]AgilityField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AgilityDefinition) SetFields(v []AgilityField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AgilityDefinition) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *AgilityDefinition) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *AgilityDefinition) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


