# UrlRedirection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**OriginUrl** | Pointer to **NullableString** |  | [optional] 
**DestinationUrl** | Pointer to **NullableString** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewUrlRedirection

`func NewUrlRedirection() *UrlRedirection`

NewUrlRedirection instantiates a new UrlRedirection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlRedirectionWithDefaults

`func NewUrlRedirectionWithDefaults() *UrlRedirection`

NewUrlRedirectionWithDefaults instantiates a new UrlRedirection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UrlRedirection) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UrlRedirection) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UrlRedirection) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UrlRedirection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOriginUrl

`func (o *UrlRedirection) GetOriginUrl() string`

GetOriginUrl returns the OriginUrl field if non-nil, zero value otherwise.

### GetOriginUrlOk

`func (o *UrlRedirection) GetOriginUrlOk() (*string, bool)`

GetOriginUrlOk returns a tuple with the OriginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginUrl

`func (o *UrlRedirection) SetOriginUrl(v string)`

SetOriginUrl sets OriginUrl field to given value.

### HasOriginUrl

`func (o *UrlRedirection) HasOriginUrl() bool`

HasOriginUrl returns a boolean if a field has been set.

### SetOriginUrlNil

`func (o *UrlRedirection) SetOriginUrlNil(b bool)`

 SetOriginUrlNil sets the value for OriginUrl to be an explicit nil

### UnsetOriginUrl
`func (o *UrlRedirection) UnsetOriginUrl()`

UnsetOriginUrl ensures that no value is present for OriginUrl, not even an explicit nil
### GetDestinationUrl

`func (o *UrlRedirection) GetDestinationUrl() string`

GetDestinationUrl returns the DestinationUrl field if non-nil, zero value otherwise.

### GetDestinationUrlOk

`func (o *UrlRedirection) GetDestinationUrlOk() (*string, bool)`

GetDestinationUrlOk returns a tuple with the DestinationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUrl

`func (o *UrlRedirection) SetDestinationUrl(v string)`

SetDestinationUrl sets DestinationUrl field to given value.

### HasDestinationUrl

`func (o *UrlRedirection) HasDestinationUrl() bool`

HasDestinationUrl returns a boolean if a field has been set.

### SetDestinationUrlNil

`func (o *UrlRedirection) SetDestinationUrlNil(b bool)`

 SetDestinationUrlNil sets the value for DestinationUrl to be an explicit nil

### UnsetDestinationUrl
`func (o *UrlRedirection) UnsetDestinationUrl()`

UnsetDestinationUrl ensures that no value is present for DestinationUrl, not even an explicit nil
### GetStatusCode

`func (o *UrlRedirection) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *UrlRedirection) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *UrlRedirection) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *UrlRedirection) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


