# AgilityField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**AgilityFieldType**](AgilityFieldType.md) |  | [optional] 
**Settings** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAgilityField

`func NewAgilityField() *AgilityField`

NewAgilityField instantiates a new AgilityField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgilityFieldWithDefaults

`func NewAgilityFieldWithDefaults() *AgilityField`

NewAgilityFieldWithDefaults instantiates a new AgilityField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AgilityField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgilityField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgilityField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgilityField) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AgilityField) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AgilityField) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLabel

`func (o *AgilityField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AgilityField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AgilityField) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AgilityField) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *AgilityField) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *AgilityField) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetType

`func (o *AgilityField) GetType() AgilityFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgilityField) GetTypeOk() (*AgilityFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgilityField) SetType(v AgilityFieldType)`

SetType sets Type field to given value.

### HasType

`func (o *AgilityField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSettings

`func (o *AgilityField) GetSettings() map[string]string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AgilityField) GetSettingsOk() (*map[string]string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AgilityField) SetSettings(v map[string]string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *AgilityField) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *AgilityField) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *AgilityField) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


