# HeadlessContentSiteMapItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageID** | Pointer to **int32** | ID of the page. | [optional] 
**Name** | Pointer to **NullableString** | Page name. | [optional] 
**Path** | Pointer to **NullableString** | Path of the page. | [optional] 
**Title** | Pointer to **NullableString** | Title of the page. | [optional] 
**MenuText** | Pointer to **NullableString** | Menu text for the page. | [optional] 
**Visible** | Pointer to [**HeadlessContentPageVisibility**](HeadlessContentPageVisibility.md) |  | [optional] 
**Redirect** | Pointer to [**RedirectUrl**](RedirectUrl.md) |  | [optional] 
**IsFolder** | Pointer to **bool** | Indicates if the item is a folder. | [optional] 
**ContentID** | Pointer to **NullableInt32** | Content ID of the item. | [optional] 

## Methods

### NewHeadlessContentSiteMapItem

`func NewHeadlessContentSiteMapItem() *HeadlessContentSiteMapItem`

NewHeadlessContentSiteMapItem instantiates a new HeadlessContentSiteMapItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeadlessContentSiteMapItemWithDefaults

`func NewHeadlessContentSiteMapItemWithDefaults() *HeadlessContentSiteMapItem`

NewHeadlessContentSiteMapItemWithDefaults instantiates a new HeadlessContentSiteMapItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageID

`func (o *HeadlessContentSiteMapItem) GetPageID() int32`

GetPageID returns the PageID field if non-nil, zero value otherwise.

### GetPageIDOk

`func (o *HeadlessContentSiteMapItem) GetPageIDOk() (*int32, bool)`

GetPageIDOk returns a tuple with the PageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageID

`func (o *HeadlessContentSiteMapItem) SetPageID(v int32)`

SetPageID sets PageID field to given value.

### HasPageID

`func (o *HeadlessContentSiteMapItem) HasPageID() bool`

HasPageID returns a boolean if a field has been set.

### GetName

`func (o *HeadlessContentSiteMapItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeadlessContentSiteMapItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeadlessContentSiteMapItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeadlessContentSiteMapItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HeadlessContentSiteMapItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HeadlessContentSiteMapItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *HeadlessContentSiteMapItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HeadlessContentSiteMapItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HeadlessContentSiteMapItem) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HeadlessContentSiteMapItem) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *HeadlessContentSiteMapItem) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *HeadlessContentSiteMapItem) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetTitle

`func (o *HeadlessContentSiteMapItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HeadlessContentSiteMapItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HeadlessContentSiteMapItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HeadlessContentSiteMapItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *HeadlessContentSiteMapItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *HeadlessContentSiteMapItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetMenuText

`func (o *HeadlessContentSiteMapItem) GetMenuText() string`

GetMenuText returns the MenuText field if non-nil, zero value otherwise.

### GetMenuTextOk

`func (o *HeadlessContentSiteMapItem) GetMenuTextOk() (*string, bool)`

GetMenuTextOk returns a tuple with the MenuText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuText

`func (o *HeadlessContentSiteMapItem) SetMenuText(v string)`

SetMenuText sets MenuText field to given value.

### HasMenuText

`func (o *HeadlessContentSiteMapItem) HasMenuText() bool`

HasMenuText returns a boolean if a field has been set.

### SetMenuTextNil

`func (o *HeadlessContentSiteMapItem) SetMenuTextNil(b bool)`

 SetMenuTextNil sets the value for MenuText to be an explicit nil

### UnsetMenuText
`func (o *HeadlessContentSiteMapItem) UnsetMenuText()`

UnsetMenuText ensures that no value is present for MenuText, not even an explicit nil
### GetVisible

`func (o *HeadlessContentSiteMapItem) GetVisible() HeadlessContentPageVisibility`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *HeadlessContentSiteMapItem) GetVisibleOk() (*HeadlessContentPageVisibility, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *HeadlessContentSiteMapItem) SetVisible(v HeadlessContentPageVisibility)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *HeadlessContentSiteMapItem) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetRedirect

`func (o *HeadlessContentSiteMapItem) GetRedirect() RedirectUrl`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *HeadlessContentSiteMapItem) GetRedirectOk() (*RedirectUrl, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *HeadlessContentSiteMapItem) SetRedirect(v RedirectUrl)`

SetRedirect sets Redirect field to given value.

### HasRedirect

`func (o *HeadlessContentSiteMapItem) HasRedirect() bool`

HasRedirect returns a boolean if a field has been set.

### GetIsFolder

`func (o *HeadlessContentSiteMapItem) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *HeadlessContentSiteMapItem) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *HeadlessContentSiteMapItem) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *HeadlessContentSiteMapItem) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### GetContentID

`func (o *HeadlessContentSiteMapItem) GetContentID() int32`

GetContentID returns the ContentID field if non-nil, zero value otherwise.

### GetContentIDOk

`func (o *HeadlessContentSiteMapItem) GetContentIDOk() (*int32, bool)`

GetContentIDOk returns a tuple with the ContentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentID

`func (o *HeadlessContentSiteMapItem) SetContentID(v int32)`

SetContentID sets ContentID field to given value.

### HasContentID

`func (o *HeadlessContentSiteMapItem) HasContentID() bool`

HasContentID returns a boolean if a field has been set.

### SetContentIDNil

`func (o *HeadlessContentSiteMapItem) SetContentIDNil(b bool)`

 SetContentIDNil sets the value for ContentID to be an explicit nil

### UnsetContentID
`func (o *HeadlessContentSiteMapItem) UnsetContentID()`

UnsetContentID ensures that no value is present for ContentID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


