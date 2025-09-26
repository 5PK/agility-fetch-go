# HeadlessContentScripts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedFromGlobal** | Pointer to **bool** | Indicates if scripts are excluded from global. | [optional] 
**Top** | Pointer to **NullableString** | Scripts to include on top of the page. | [optional] 
**Bottom** | Pointer to **NullableString** | Scripts to include on the bottom of the page. | [optional] 

## Methods

### NewHeadlessContentScripts

`func NewHeadlessContentScripts() *HeadlessContentScripts`

NewHeadlessContentScripts instantiates a new HeadlessContentScripts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeadlessContentScriptsWithDefaults

`func NewHeadlessContentScriptsWithDefaults() *HeadlessContentScripts`

NewHeadlessContentScriptsWithDefaults instantiates a new HeadlessContentScripts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludedFromGlobal

`func (o *HeadlessContentScripts) GetExcludedFromGlobal() bool`

GetExcludedFromGlobal returns the ExcludedFromGlobal field if non-nil, zero value otherwise.

### GetExcludedFromGlobalOk

`func (o *HeadlessContentScripts) GetExcludedFromGlobalOk() (*bool, bool)`

GetExcludedFromGlobalOk returns a tuple with the ExcludedFromGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedFromGlobal

`func (o *HeadlessContentScripts) SetExcludedFromGlobal(v bool)`

SetExcludedFromGlobal sets ExcludedFromGlobal field to given value.

### HasExcludedFromGlobal

`func (o *HeadlessContentScripts) HasExcludedFromGlobal() bool`

HasExcludedFromGlobal returns a boolean if a field has been set.

### GetTop

`func (o *HeadlessContentScripts) GetTop() string`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *HeadlessContentScripts) GetTopOk() (*string, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *HeadlessContentScripts) SetTop(v string)`

SetTop sets Top field to given value.

### HasTop

`func (o *HeadlessContentScripts) HasTop() bool`

HasTop returns a boolean if a field has been set.

### SetTopNil

`func (o *HeadlessContentScripts) SetTopNil(b bool)`

 SetTopNil sets the value for Top to be an explicit nil

### UnsetTop
`func (o *HeadlessContentScripts) UnsetTop()`

UnsetTop ensures that no value is present for Top, not even an explicit nil
### GetBottom

`func (o *HeadlessContentScripts) GetBottom() string`

GetBottom returns the Bottom field if non-nil, zero value otherwise.

### GetBottomOk

`func (o *HeadlessContentScripts) GetBottomOk() (*string, bool)`

GetBottomOk returns a tuple with the Bottom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottom

`func (o *HeadlessContentScripts) SetBottom(v string)`

SetBottom sets Bottom field to given value.

### HasBottom

`func (o *HeadlessContentScripts) HasBottom() bool`

HasBottom returns a boolean if a field has been set.

### SetBottomNil

`func (o *HeadlessContentScripts) SetBottomNil(b bool)`

 SetBottomNil sets the value for Bottom to be an explicit nil

### UnsetBottom
`func (o *HeadlessContentScripts) UnsetBottom()`

UnsetBottom ensures that no value is present for Bottom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


