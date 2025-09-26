# HeadlessContentPageBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageID** | Pointer to **int32** | ID of the page. | [optional] 
**Name** | Pointer to **NullableString** | Page name. | [optional] 
**Path** | Pointer to **NullableString** | Path of the page. | [optional] 
**Title** | Pointer to **NullableString** | Title of the page. | [optional] 
**MenuText** | Pointer to **NullableString** | Menu text for the page. | [optional] 
**Visible** | Pointer to [**HeadlessContentPageVisibility**](HeadlessContentPageVisibility.md) |  | [optional] 

## Methods

### NewHeadlessContentPageBase

`func NewHeadlessContentPageBase() *HeadlessContentPageBase`

NewHeadlessContentPageBase instantiates a new HeadlessContentPageBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeadlessContentPageBaseWithDefaults

`func NewHeadlessContentPageBaseWithDefaults() *HeadlessContentPageBase`

NewHeadlessContentPageBaseWithDefaults instantiates a new HeadlessContentPageBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageID

`func (o *HeadlessContentPageBase) GetPageID() int32`

GetPageID returns the PageID field if non-nil, zero value otherwise.

### GetPageIDOk

`func (o *HeadlessContentPageBase) GetPageIDOk() (*int32, bool)`

GetPageIDOk returns a tuple with the PageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageID

`func (o *HeadlessContentPageBase) SetPageID(v int32)`

SetPageID sets PageID field to given value.

### HasPageID

`func (o *HeadlessContentPageBase) HasPageID() bool`

HasPageID returns a boolean if a field has been set.

### GetName

`func (o *HeadlessContentPageBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeadlessContentPageBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeadlessContentPageBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeadlessContentPageBase) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HeadlessContentPageBase) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HeadlessContentPageBase) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *HeadlessContentPageBase) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HeadlessContentPageBase) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HeadlessContentPageBase) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HeadlessContentPageBase) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *HeadlessContentPageBase) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *HeadlessContentPageBase) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetTitle

`func (o *HeadlessContentPageBase) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HeadlessContentPageBase) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HeadlessContentPageBase) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HeadlessContentPageBase) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *HeadlessContentPageBase) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *HeadlessContentPageBase) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetMenuText

`func (o *HeadlessContentPageBase) GetMenuText() string`

GetMenuText returns the MenuText field if non-nil, zero value otherwise.

### GetMenuTextOk

`func (o *HeadlessContentPageBase) GetMenuTextOk() (*string, bool)`

GetMenuTextOk returns a tuple with the MenuText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuText

`func (o *HeadlessContentPageBase) SetMenuText(v string)`

SetMenuText sets MenuText field to given value.

### HasMenuText

`func (o *HeadlessContentPageBase) HasMenuText() bool`

HasMenuText returns a boolean if a field has been set.

### SetMenuTextNil

`func (o *HeadlessContentPageBase) SetMenuTextNil(b bool)`

 SetMenuTextNil sets the value for MenuText to be an explicit nil

### UnsetMenuText
`func (o *HeadlessContentPageBase) UnsetMenuText()`

UnsetMenuText ensures that no value is present for MenuText, not even an explicit nil
### GetVisible

`func (o *HeadlessContentPageBase) GetVisible() HeadlessContentPageVisibility`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *HeadlessContentPageBase) GetVisibleOk() (*HeadlessContentPageVisibility, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *HeadlessContentPageBase) SetVisible(v HeadlessContentPageVisibility)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *HeadlessContentPageBase) HasVisible() bool`

HasVisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


