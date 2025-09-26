# RedirectUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **NullableString** | Redirect Url. | [optional] 
**Target** | Pointer to **NullableString** | Target to open the Url in. | [optional] 

## Methods

### NewRedirectUrl

`func NewRedirectUrl() *RedirectUrl`

NewRedirectUrl instantiates a new RedirectUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedirectUrlWithDefaults

`func NewRedirectUrlWithDefaults() *RedirectUrl`

NewRedirectUrlWithDefaults instantiates a new RedirectUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *RedirectUrl) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RedirectUrl) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RedirectUrl) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RedirectUrl) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *RedirectUrl) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *RedirectUrl) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetTarget

`func (o *RedirectUrl) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RedirectUrl) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RedirectUrl) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *RedirectUrl) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *RedirectUrl) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *RedirectUrl) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


