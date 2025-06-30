# AdvisoryJVNAdvisoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpe** | Pointer to [**[]AdvisoryJVNCPE**](AdvisoryJVNCPE.md) |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**Cvss** | Pointer to [**[]AdvisoryCVSS**](AdvisoryCVSS.md) |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DescriptionEn** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Issued** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **string** |  | [optional] 
**References** | Pointer to [**[]AdvisoryJVNReference**](AdvisoryJVNReference.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TitleEn** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UrlEn** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryJVNAdvisoryItem

`func NewAdvisoryJVNAdvisoryItem() *AdvisoryJVNAdvisoryItem`

NewAdvisoryJVNAdvisoryItem instantiates a new AdvisoryJVNAdvisoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryJVNAdvisoryItemWithDefaults

`func NewAdvisoryJVNAdvisoryItemWithDefaults() *AdvisoryJVNAdvisoryItem`

NewAdvisoryJVNAdvisoryItemWithDefaults instantiates a new AdvisoryJVNAdvisoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpe

`func (o *AdvisoryJVNAdvisoryItem) GetCpe() []AdvisoryJVNCPE`

GetCpe returns the Cpe field if non-nil, zero value otherwise.

### GetCpeOk

`func (o *AdvisoryJVNAdvisoryItem) GetCpeOk() (*[]AdvisoryJVNCPE, bool)`

GetCpeOk returns a tuple with the Cpe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpe

`func (o *AdvisoryJVNAdvisoryItem) SetCpe(v []AdvisoryJVNCPE)`

SetCpe sets Cpe field to given value.

### HasCpe

`func (o *AdvisoryJVNAdvisoryItem) HasCpe() bool`

HasCpe returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryJVNAdvisoryItem) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryJVNAdvisoryItem) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryJVNAdvisoryItem) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryJVNAdvisoryItem) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvss

`func (o *AdvisoryJVNAdvisoryItem) GetCvss() []AdvisoryCVSS`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *AdvisoryJVNAdvisoryItem) GetCvssOk() (*[]AdvisoryCVSS, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *AdvisoryJVNAdvisoryItem) SetCvss(v []AdvisoryCVSS)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *AdvisoryJVNAdvisoryItem) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryJVNAdvisoryItem) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryJVNAdvisoryItem) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryJVNAdvisoryItem) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryJVNAdvisoryItem) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryJVNAdvisoryItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryJVNAdvisoryItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryJVNAdvisoryItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryJVNAdvisoryItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionEn

`func (o *AdvisoryJVNAdvisoryItem) GetDescriptionEn() string`

GetDescriptionEn returns the DescriptionEn field if non-nil, zero value otherwise.

### GetDescriptionEnOk

`func (o *AdvisoryJVNAdvisoryItem) GetDescriptionEnOk() (*string, bool)`

GetDescriptionEnOk returns a tuple with the DescriptionEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionEn

`func (o *AdvisoryJVNAdvisoryItem) SetDescriptionEn(v string)`

SetDescriptionEn sets DescriptionEn field to given value.

### HasDescriptionEn

`func (o *AdvisoryJVNAdvisoryItem) HasDescriptionEn() bool`

HasDescriptionEn returns a boolean if a field has been set.

### GetIdentifier

`func (o *AdvisoryJVNAdvisoryItem) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AdvisoryJVNAdvisoryItem) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AdvisoryJVNAdvisoryItem) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *AdvisoryJVNAdvisoryItem) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIssued

`func (o *AdvisoryJVNAdvisoryItem) GetIssued() string`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *AdvisoryJVNAdvisoryItem) GetIssuedOk() (*string, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *AdvisoryJVNAdvisoryItem) SetIssued(v string)`

SetIssued sets Issued field to given value.

### HasIssued

`func (o *AdvisoryJVNAdvisoryItem) HasIssued() bool`

HasIssued returns a boolean if a field has been set.

### GetModified

`func (o *AdvisoryJVNAdvisoryItem) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AdvisoryJVNAdvisoryItem) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AdvisoryJVNAdvisoryItem) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *AdvisoryJVNAdvisoryItem) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryJVNAdvisoryItem) GetReferences() []AdvisoryJVNReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryJVNAdvisoryItem) GetReferencesOk() (*[]AdvisoryJVNReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryJVNAdvisoryItem) SetReferences(v []AdvisoryJVNReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryJVNAdvisoryItem) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryJVNAdvisoryItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryJVNAdvisoryItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryJVNAdvisoryItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryJVNAdvisoryItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleEn

`func (o *AdvisoryJVNAdvisoryItem) GetTitleEn() string`

GetTitleEn returns the TitleEn field if non-nil, zero value otherwise.

### GetTitleEnOk

`func (o *AdvisoryJVNAdvisoryItem) GetTitleEnOk() (*string, bool)`

GetTitleEnOk returns a tuple with the TitleEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleEn

`func (o *AdvisoryJVNAdvisoryItem) SetTitleEn(v string)`

SetTitleEn sets TitleEn field to given value.

### HasTitleEn

`func (o *AdvisoryJVNAdvisoryItem) HasTitleEn() bool`

HasTitleEn returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryJVNAdvisoryItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryJVNAdvisoryItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryJVNAdvisoryItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryJVNAdvisoryItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryJVNAdvisoryItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryJVNAdvisoryItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryJVNAdvisoryItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryJVNAdvisoryItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUrlEn

`func (o *AdvisoryJVNAdvisoryItem) GetUrlEn() string`

GetUrlEn returns the UrlEn field if non-nil, zero value otherwise.

### GetUrlEnOk

`func (o *AdvisoryJVNAdvisoryItem) GetUrlEnOk() (*string, bool)`

GetUrlEnOk returns a tuple with the UrlEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlEn

`func (o *AdvisoryJVNAdvisoryItem) SetUrlEn(v string)`

SetUrlEn sets UrlEn field to given value.

### HasUrlEn

`func (o *AdvisoryJVNAdvisoryItem) HasUrlEn() bool`

HasUrlEn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


