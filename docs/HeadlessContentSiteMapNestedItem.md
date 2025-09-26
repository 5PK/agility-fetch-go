# HeadlessContentSiteMapNestedItem

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
**Children** | Pointer to [**[]HeadlessContentSiteMapNestedItem**](HeadlessContentSiteMapNestedItem.md) | Children of this item | [optional] 

## Methods

### NewHeadlessContentSiteMapNestedItem

`func NewHeadlessContentSiteMapNestedItem() *HeadlessContentSiteMapNestedItem`

NewHeadlessContentSiteMapNestedItem instantiates a new HeadlessContentSiteMapNestedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeadlessContentSiteMapNestedItemWithDefaults

`func NewHeadlessContentSiteMapNestedItemWithDefaults() *HeadlessContentSiteMapNestedItem`

NewHeadlessContentSiteMapNestedItemWithDefaults instantiates a new HeadlessContentSiteMapNestedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageID

`func (o *HeadlessContentSiteMapNestedItem) GetPageID() int32`

GetPageID returns the PageID field if non-nil, zero value otherwise.

### GetPageIDOk

`func (o *HeadlessContentSiteMapNestedItem) GetPageIDOk() (*int32, bool)`

GetPageIDOk returns a tuple with the PageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageID

`func (o *HeadlessContentSiteMapNestedItem) SetPageID(v int32)`

SetPageID sets PageID field to given value.

### HasPageID

`func (o *HeadlessContentSiteMapNestedItem) HasPageID() bool`

HasPageID returns a boolean if a field has been set.

### GetName

`func (o *HeadlessContentSiteMapNestedItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeadlessContentSiteMapNestedItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeadlessContentSiteMapNestedItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeadlessContentSiteMapNestedItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HeadlessContentSiteMapNestedItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HeadlessContentSiteMapNestedItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *HeadlessContentSiteMapNestedItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HeadlessContentSiteMapNestedItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HeadlessContentSiteMapNestedItem) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HeadlessContentSiteMapNestedItem) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *HeadlessContentSiteMapNestedItem) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *HeadlessContentSiteMapNestedItem) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetTitle

`func (o *HeadlessContentSiteMapNestedItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HeadlessContentSiteMapNestedItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HeadlessContentSiteMapNestedItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HeadlessContentSiteMapNestedItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *HeadlessContentSiteMapNestedItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *HeadlessContentSiteMapNestedItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetMenuText

`func (o *HeadlessContentSiteMapNestedItem) GetMenuText() string`

GetMenuText returns the MenuText field if non-nil, zero value otherwise.

### GetMenuTextOk

`func (o *HeadlessContentSiteMapNestedItem) GetMenuTextOk() (*string, bool)`

GetMenuTextOk returns a tuple with the MenuText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuText

`func (o *HeadlessContentSiteMapNestedItem) SetMenuText(v string)`

SetMenuText sets MenuText field to given value.

### HasMenuText

`func (o *HeadlessContentSiteMapNestedItem) HasMenuText() bool`

HasMenuText returns a boolean if a field has been set.

### SetMenuTextNil

`func (o *HeadlessContentSiteMapNestedItem) SetMenuTextNil(b bool)`

 SetMenuTextNil sets the value for MenuText to be an explicit nil

### UnsetMenuText
`func (o *HeadlessContentSiteMapNestedItem) UnsetMenuText()`

UnsetMenuText ensures that no value is present for MenuText, not even an explicit nil
### GetVisible

`func (o *HeadlessContentSiteMapNestedItem) GetVisible() HeadlessContentPageVisibility`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *HeadlessContentSiteMapNestedItem) GetVisibleOk() (*HeadlessContentPageVisibility, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *HeadlessContentSiteMapNestedItem) SetVisible(v HeadlessContentPageVisibility)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *HeadlessContentSiteMapNestedItem) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetRedirect

`func (o *HeadlessContentSiteMapNestedItem) GetRedirect() RedirectUrl`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *HeadlessContentSiteMapNestedItem) GetRedirectOk() (*RedirectUrl, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *HeadlessContentSiteMapNestedItem) SetRedirect(v RedirectUrl)`

SetRedirect sets Redirect field to given value.

### HasRedirect

`func (o *HeadlessContentSiteMapNestedItem) HasRedirect() bool`

HasRedirect returns a boolean if a field has been set.

### GetIsFolder

`func (o *HeadlessContentSiteMapNestedItem) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *HeadlessContentSiteMapNestedItem) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *HeadlessContentSiteMapNestedItem) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *HeadlessContentSiteMapNestedItem) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### GetContentID

`func (o *HeadlessContentSiteMapNestedItem) GetContentID() int32`

GetContentID returns the ContentID field if non-nil, zero value otherwise.

### GetContentIDOk

`func (o *HeadlessContentSiteMapNestedItem) GetContentIDOk() (*int32, bool)`

GetContentIDOk returns a tuple with the ContentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentID

`func (o *HeadlessContentSiteMapNestedItem) SetContentID(v int32)`

SetContentID sets ContentID field to given value.

### HasContentID

`func (o *HeadlessContentSiteMapNestedItem) HasContentID() bool`

HasContentID returns a boolean if a field has been set.

### SetContentIDNil

`func (o *HeadlessContentSiteMapNestedItem) SetContentIDNil(b bool)`

 SetContentIDNil sets the value for ContentID to be an explicit nil

### UnsetContentID
`func (o *HeadlessContentSiteMapNestedItem) UnsetContentID()`

UnsetContentID ensures that no value is present for ContentID, not even an explicit nil
### GetChildren

`func (o *HeadlessContentSiteMapNestedItem) GetChildren() []HeadlessContentSiteMapNestedItem`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *HeadlessContentSiteMapNestedItem) GetChildrenOk() (*[]HeadlessContentSiteMapNestedItem, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *HeadlessContentSiteMapNestedItem) SetChildren(v []HeadlessContentSiteMapNestedItem)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *HeadlessContentSiteMapNestedItem) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildrenNil

`func (o *HeadlessContentSiteMapNestedItem) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *HeadlessContentSiteMapNestedItem) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


