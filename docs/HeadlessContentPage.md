# HeadlessContentPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageID** | Pointer to **int32** | ID of the page. | [optional] 
**Name** | Pointer to **NullableString** | Page name. | [optional] 
**Path** | Pointer to **NullableString** | Path of the page. | [optional] 
**Title** | Pointer to **NullableString** | Title of the page. | [optional] 
**MenuText** | Pointer to **NullableString** | Menu text for the page. | [optional] 
**Visible** | Pointer to [**HeadlessContentPageVisibility**](HeadlessContentPageVisibility.md) |  | [optional] 
**TemplateName** | Pointer to **NullableString** | Name of the template used to create the page. | [optional] 
**RedirectUrl** | Pointer to **NullableString** | Redirect url of the page. | [optional] 
**SecurePage** | Pointer to **bool** | Indicates if the page is secure. | [optional] 
**ExcludeFromOutputCache** | Pointer to **bool** | Indicates if the page is excluded from cache. | [optional] 
**Seo** | Pointer to [**HeadlessContentItemSeo**](HeadlessContentItemSeo.md) |  | [optional] 
**Scripts** | Pointer to [**HeadlessContentScripts**](HeadlessContentScripts.md) |  | [optional] 
**Properties** | Pointer to [**HeadlessContentItemProperties**](HeadlessContentItemProperties.md) |  | [optional] 
**Zones** | Pointer to [**map[string][]HeadlessContentZone**](array.md) | Zones on the page defined by the template | [optional] 

## Methods

### NewHeadlessContentPage

`func NewHeadlessContentPage() *HeadlessContentPage`

NewHeadlessContentPage instantiates a new HeadlessContentPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeadlessContentPageWithDefaults

`func NewHeadlessContentPageWithDefaults() *HeadlessContentPage`

NewHeadlessContentPageWithDefaults instantiates a new HeadlessContentPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageID

`func (o *HeadlessContentPage) GetPageID() int32`

GetPageID returns the PageID field if non-nil, zero value otherwise.

### GetPageIDOk

`func (o *HeadlessContentPage) GetPageIDOk() (*int32, bool)`

GetPageIDOk returns a tuple with the PageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageID

`func (o *HeadlessContentPage) SetPageID(v int32)`

SetPageID sets PageID field to given value.

### HasPageID

`func (o *HeadlessContentPage) HasPageID() bool`

HasPageID returns a boolean if a field has been set.

### GetName

`func (o *HeadlessContentPage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeadlessContentPage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeadlessContentPage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeadlessContentPage) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HeadlessContentPage) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HeadlessContentPage) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *HeadlessContentPage) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HeadlessContentPage) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HeadlessContentPage) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HeadlessContentPage) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *HeadlessContentPage) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *HeadlessContentPage) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetTitle

`func (o *HeadlessContentPage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HeadlessContentPage) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HeadlessContentPage) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HeadlessContentPage) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *HeadlessContentPage) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *HeadlessContentPage) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetMenuText

`func (o *HeadlessContentPage) GetMenuText() string`

GetMenuText returns the MenuText field if non-nil, zero value otherwise.

### GetMenuTextOk

`func (o *HeadlessContentPage) GetMenuTextOk() (*string, bool)`

GetMenuTextOk returns a tuple with the MenuText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuText

`func (o *HeadlessContentPage) SetMenuText(v string)`

SetMenuText sets MenuText field to given value.

### HasMenuText

`func (o *HeadlessContentPage) HasMenuText() bool`

HasMenuText returns a boolean if a field has been set.

### SetMenuTextNil

`func (o *HeadlessContentPage) SetMenuTextNil(b bool)`

 SetMenuTextNil sets the value for MenuText to be an explicit nil

### UnsetMenuText
`func (o *HeadlessContentPage) UnsetMenuText()`

UnsetMenuText ensures that no value is present for MenuText, not even an explicit nil
### GetVisible

`func (o *HeadlessContentPage) GetVisible() HeadlessContentPageVisibility`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *HeadlessContentPage) GetVisibleOk() (*HeadlessContentPageVisibility, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *HeadlessContentPage) SetVisible(v HeadlessContentPageVisibility)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *HeadlessContentPage) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetTemplateName

`func (o *HeadlessContentPage) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *HeadlessContentPage) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *HeadlessContentPage) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *HeadlessContentPage) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### SetTemplateNameNil

`func (o *HeadlessContentPage) SetTemplateNameNil(b bool)`

 SetTemplateNameNil sets the value for TemplateName to be an explicit nil

### UnsetTemplateName
`func (o *HeadlessContentPage) UnsetTemplateName()`

UnsetTemplateName ensures that no value is present for TemplateName, not even an explicit nil
### GetRedirectUrl

`func (o *HeadlessContentPage) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *HeadlessContentPage) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *HeadlessContentPage) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *HeadlessContentPage) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *HeadlessContentPage) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *HeadlessContentPage) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetSecurePage

`func (o *HeadlessContentPage) GetSecurePage() bool`

GetSecurePage returns the SecurePage field if non-nil, zero value otherwise.

### GetSecurePageOk

`func (o *HeadlessContentPage) GetSecurePageOk() (*bool, bool)`

GetSecurePageOk returns a tuple with the SecurePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePage

`func (o *HeadlessContentPage) SetSecurePage(v bool)`

SetSecurePage sets SecurePage field to given value.

### HasSecurePage

`func (o *HeadlessContentPage) HasSecurePage() bool`

HasSecurePage returns a boolean if a field has been set.

### GetExcludeFromOutputCache

`func (o *HeadlessContentPage) GetExcludeFromOutputCache() bool`

GetExcludeFromOutputCache returns the ExcludeFromOutputCache field if non-nil, zero value otherwise.

### GetExcludeFromOutputCacheOk

`func (o *HeadlessContentPage) GetExcludeFromOutputCacheOk() (*bool, bool)`

GetExcludeFromOutputCacheOk returns a tuple with the ExcludeFromOutputCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromOutputCache

`func (o *HeadlessContentPage) SetExcludeFromOutputCache(v bool)`

SetExcludeFromOutputCache sets ExcludeFromOutputCache field to given value.

### HasExcludeFromOutputCache

`func (o *HeadlessContentPage) HasExcludeFromOutputCache() bool`

HasExcludeFromOutputCache returns a boolean if a field has been set.

### GetSeo

`func (o *HeadlessContentPage) GetSeo() HeadlessContentItemSeo`

GetSeo returns the Seo field if non-nil, zero value otherwise.

### GetSeoOk

`func (o *HeadlessContentPage) GetSeoOk() (*HeadlessContentItemSeo, bool)`

GetSeoOk returns a tuple with the Seo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeo

`func (o *HeadlessContentPage) SetSeo(v HeadlessContentItemSeo)`

SetSeo sets Seo field to given value.

### HasSeo

`func (o *HeadlessContentPage) HasSeo() bool`

HasSeo returns a boolean if a field has been set.

### GetScripts

`func (o *HeadlessContentPage) GetScripts() HeadlessContentScripts`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *HeadlessContentPage) GetScriptsOk() (*HeadlessContentScripts, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *HeadlessContentPage) SetScripts(v HeadlessContentScripts)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *HeadlessContentPage) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetProperties

`func (o *HeadlessContentPage) GetProperties() HeadlessContentItemProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *HeadlessContentPage) GetPropertiesOk() (*HeadlessContentItemProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *HeadlessContentPage) SetProperties(v HeadlessContentItemProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *HeadlessContentPage) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetZones

`func (o *HeadlessContentPage) GetZones() map[string][]HeadlessContentZone`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *HeadlessContentPage) GetZonesOk() (*map[string][]HeadlessContentZone, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *HeadlessContentPage) SetZones(v map[string][]HeadlessContentZone)`

SetZones sets Zones field to given value.

### HasZones

`func (o *HeadlessContentPage) HasZones() bool`

HasZones returns a boolean if a field has been set.

### SetZonesNil

`func (o *HeadlessContentPage) SetZonesNil(b bool)`

 SetZonesNil sets the value for Zones to be an explicit nil

### UnsetZones
`func (o *HeadlessContentPage) UnsetZones()`

UnsetZones ensures that no value is present for Zones, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


