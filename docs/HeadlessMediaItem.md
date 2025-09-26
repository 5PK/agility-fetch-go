# HeadlessMediaItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaID** | Pointer to **int32** | media Id. | [optional] 
**FileName** | Pointer to **NullableString** | Media item file name. | [optional] 
**Url** | Pointer to **NullableString** | Url that media item can be accessed at. | [optional] 
**Size** | Pointer to **int32** | Size of media item in bytes. | [optional] 
**ModifiedOn** | Pointer to **string** | Last modified date. | [optional] 
**MetaData** | Pointer to **map[string]string** | List of meta data related to media item in key value format. | [optional] 

## Methods

### NewHeadlessMediaItem

`func NewHeadlessMediaItem() *HeadlessMediaItem`

NewHeadlessMediaItem instantiates a new HeadlessMediaItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeadlessMediaItemWithDefaults

`func NewHeadlessMediaItemWithDefaults() *HeadlessMediaItem`

NewHeadlessMediaItemWithDefaults instantiates a new HeadlessMediaItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaID

`func (o *HeadlessMediaItem) GetMediaID() int32`

GetMediaID returns the MediaID field if non-nil, zero value otherwise.

### GetMediaIDOk

`func (o *HeadlessMediaItem) GetMediaIDOk() (*int32, bool)`

GetMediaIDOk returns a tuple with the MediaID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaID

`func (o *HeadlessMediaItem) SetMediaID(v int32)`

SetMediaID sets MediaID field to given value.

### HasMediaID

`func (o *HeadlessMediaItem) HasMediaID() bool`

HasMediaID returns a boolean if a field has been set.

### GetFileName

`func (o *HeadlessMediaItem) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *HeadlessMediaItem) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *HeadlessMediaItem) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *HeadlessMediaItem) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *HeadlessMediaItem) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *HeadlessMediaItem) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetUrl

`func (o *HeadlessMediaItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HeadlessMediaItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HeadlessMediaItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HeadlessMediaItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *HeadlessMediaItem) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *HeadlessMediaItem) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetSize

`func (o *HeadlessMediaItem) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *HeadlessMediaItem) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *HeadlessMediaItem) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *HeadlessMediaItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetModifiedOn

`func (o *HeadlessMediaItem) GetModifiedOn() string`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *HeadlessMediaItem) GetModifiedOnOk() (*string, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *HeadlessMediaItem) SetModifiedOn(v string)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *HeadlessMediaItem) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetMetaData

`func (o *HeadlessMediaItem) GetMetaData() map[string]string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *HeadlessMediaItem) GetMetaDataOk() (*map[string]string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *HeadlessMediaItem) SetMetaData(v map[string]string)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *HeadlessMediaItem) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### SetMetaDataNil

`func (o *HeadlessMediaItem) SetMetaDataNil(b bool)`

 SetMetaDataNil sets the value for MetaData to be an explicit nil

### UnsetMetaData
`func (o *HeadlessMediaItem) UnsetMetaData()`

UnsetMetaData ensures that no value is present for MetaData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


