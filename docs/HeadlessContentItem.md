# HeadlessContentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentID** | Pointer to **int32** | The unique ID of the content item in this locale. | [optional] 
**Properties** | Pointer to **map[string]interface{}** | The system properties of the content item. | [optional] 
**Fields** | Pointer to **map[string]interface{}** | A dictionary of the fields and the values of the content item. | [optional] 
**Seo** | Pointer to **map[string]interface{}** | SEO related fields for the content item. This is only returned for Dynamic Page Items. | [optional] 

## Methods

### NewHeadlessContentItem

`func NewHeadlessContentItem() *HeadlessContentItem`

NewHeadlessContentItem instantiates a new HeadlessContentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeadlessContentItemWithDefaults

`func NewHeadlessContentItemWithDefaults() *HeadlessContentItem`

NewHeadlessContentItemWithDefaults instantiates a new HeadlessContentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentID

`func (o *HeadlessContentItem) GetContentID() int32`

GetContentID returns the ContentID field if non-nil, zero value otherwise.

### GetContentIDOk

`func (o *HeadlessContentItem) GetContentIDOk() (*int32, bool)`

GetContentIDOk returns a tuple with the ContentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentID

`func (o *HeadlessContentItem) SetContentID(v int32)`

SetContentID sets ContentID field to given value.

### HasContentID

`func (o *HeadlessContentItem) HasContentID() bool`

HasContentID returns a boolean if a field has been set.

### GetProperties

`func (o *HeadlessContentItem) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *HeadlessContentItem) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *HeadlessContentItem) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *HeadlessContentItem) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *HeadlessContentItem) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *HeadlessContentItem) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetFields

`func (o *HeadlessContentItem) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *HeadlessContentItem) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *HeadlessContentItem) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *HeadlessContentItem) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *HeadlessContentItem) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *HeadlessContentItem) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetSeo

`func (o *HeadlessContentItem) GetSeo() map[string]interface{}`

GetSeo returns the Seo field if non-nil, zero value otherwise.

### GetSeoOk

`func (o *HeadlessContentItem) GetSeoOk() (*map[string]interface{}, bool)`

GetSeoOk returns a tuple with the Seo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeo

`func (o *HeadlessContentItem) SetSeo(v map[string]interface{})`

SetSeo sets Seo field to given value.

### HasSeo

`func (o *HeadlessContentItem) HasSeo() bool`

HasSeo returns a boolean if a field has been set.

### SetSeoNil

`func (o *HeadlessContentItem) SetSeoNil(b bool)`

 SetSeoNil sets the value for Seo to be an explicit nil

### UnsetSeo
`func (o *HeadlessContentItem) UnsetSeo()`

UnsetSeo ensures that no value is present for Seo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


