# HeadlessGallery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GalleryId** | Pointer to **int32** | Gallery ID. | [optional] 
**Name** | Pointer to **NullableString** | Gallery name. | [optional] 
**Description** | Pointer to **NullableString** | Gallery description. | [optional] 
**Media** | Pointer to [**[]HeadlessMediaItem**](HeadlessMediaItem.md) | List of Media items in gallery. | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewHeadlessGallery

`func NewHeadlessGallery() *HeadlessGallery`

NewHeadlessGallery instantiates a new HeadlessGallery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeadlessGalleryWithDefaults

`func NewHeadlessGalleryWithDefaults() *HeadlessGallery`

NewHeadlessGalleryWithDefaults instantiates a new HeadlessGallery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGalleryId

`func (o *HeadlessGallery) GetGalleryId() int32`

GetGalleryId returns the GalleryId field if non-nil, zero value otherwise.

### GetGalleryIdOk

`func (o *HeadlessGallery) GetGalleryIdOk() (*int32, bool)`

GetGalleryIdOk returns a tuple with the GalleryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGalleryId

`func (o *HeadlessGallery) SetGalleryId(v int32)`

SetGalleryId sets GalleryId field to given value.

### HasGalleryId

`func (o *HeadlessGallery) HasGalleryId() bool`

HasGalleryId returns a boolean if a field has been set.

### GetName

`func (o *HeadlessGallery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeadlessGallery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeadlessGallery) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeadlessGallery) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HeadlessGallery) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HeadlessGallery) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *HeadlessGallery) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HeadlessGallery) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HeadlessGallery) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HeadlessGallery) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HeadlessGallery) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HeadlessGallery) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMedia

`func (o *HeadlessGallery) GetMedia() []HeadlessMediaItem`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *HeadlessGallery) GetMediaOk() (*[]HeadlessMediaItem, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *HeadlessGallery) SetMedia(v []HeadlessMediaItem)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *HeadlessGallery) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### SetMediaNil

`func (o *HeadlessGallery) SetMediaNil(b bool)`

 SetMediaNil sets the value for Media to be an explicit nil

### UnsetMedia
`func (o *HeadlessGallery) UnsetMedia()`

UnsetMedia ensures that no value is present for Media, not even an explicit nil
### GetCount

`func (o *HeadlessGallery) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HeadlessGallery) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HeadlessGallery) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *HeadlessGallery) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


