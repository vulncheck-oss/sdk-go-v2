# AdvisoryCISAAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedProducts** | Pointer to **string** |  | [optional] 
**AlertID** | Pointer to **string** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**CveexploitedITW** | Pointer to **bool** |  | [optional] 
**Cvss** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Icsa** | Pointer to **bool** |  | [optional] 
**Icsma** | Pointer to **bool** |  | [optional] 
**Mitigations** | Pointer to **string** |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryCISAAlert

`func NewAdvisoryCISAAlert() *AdvisoryCISAAlert`

NewAdvisoryCISAAlert instantiates a new AdvisoryCISAAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCISAAlertWithDefaults

`func NewAdvisoryCISAAlertWithDefaults() *AdvisoryCISAAlert`

NewAdvisoryCISAAlertWithDefaults instantiates a new AdvisoryCISAAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedProducts

`func (o *AdvisoryCISAAlert) GetAffectedProducts() string`

GetAffectedProducts returns the AffectedProducts field if non-nil, zero value otherwise.

### GetAffectedProductsOk

`func (o *AdvisoryCISAAlert) GetAffectedProductsOk() (*string, bool)`

GetAffectedProductsOk returns a tuple with the AffectedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProducts

`func (o *AdvisoryCISAAlert) SetAffectedProducts(v string)`

SetAffectedProducts sets AffectedProducts field to given value.

### HasAffectedProducts

`func (o *AdvisoryCISAAlert) HasAffectedProducts() bool`

HasAffectedProducts returns a boolean if a field has been set.

### GetAlertID

`func (o *AdvisoryCISAAlert) GetAlertID() string`

GetAlertID returns the AlertID field if non-nil, zero value otherwise.

### GetAlertIDOk

`func (o *AdvisoryCISAAlert) GetAlertIDOk() (*string, bool)`

GetAlertIDOk returns a tuple with the AlertID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertID

`func (o *AdvisoryCISAAlert) SetAlertID(v string)`

SetAlertID sets AlertID field to given value.

### HasAlertID

`func (o *AdvisoryCISAAlert) HasAlertID() bool`

HasAlertID returns a boolean if a field has been set.

### GetArchived

`func (o *AdvisoryCISAAlert) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *AdvisoryCISAAlert) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *AdvisoryCISAAlert) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *AdvisoryCISAAlert) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryCISAAlert) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryCISAAlert) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryCISAAlert) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryCISAAlert) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCveexploitedITW

`func (o *AdvisoryCISAAlert) GetCveexploitedITW() bool`

GetCveexploitedITW returns the CveexploitedITW field if non-nil, zero value otherwise.

### GetCveexploitedITWOk

`func (o *AdvisoryCISAAlert) GetCveexploitedITWOk() (*bool, bool)`

GetCveexploitedITWOk returns a tuple with the CveexploitedITW field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveexploitedITW

`func (o *AdvisoryCISAAlert) SetCveexploitedITW(v bool)`

SetCveexploitedITW sets CveexploitedITW field to given value.

### HasCveexploitedITW

`func (o *AdvisoryCISAAlert) HasCveexploitedITW() bool`

HasCveexploitedITW returns a boolean if a field has been set.

### GetCvss

`func (o *AdvisoryCISAAlert) GetCvss() string`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *AdvisoryCISAAlert) GetCvssOk() (*string, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *AdvisoryCISAAlert) SetCvss(v string)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *AdvisoryCISAAlert) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryCISAAlert) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryCISAAlert) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryCISAAlert) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryCISAAlert) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetIcsa

`func (o *AdvisoryCISAAlert) GetIcsa() bool`

GetIcsa returns the Icsa field if non-nil, zero value otherwise.

### GetIcsaOk

`func (o *AdvisoryCISAAlert) GetIcsaOk() (*bool, bool)`

GetIcsaOk returns a tuple with the Icsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcsa

`func (o *AdvisoryCISAAlert) SetIcsa(v bool)`

SetIcsa sets Icsa field to given value.

### HasIcsa

`func (o *AdvisoryCISAAlert) HasIcsa() bool`

HasIcsa returns a boolean if a field has been set.

### GetIcsma

`func (o *AdvisoryCISAAlert) GetIcsma() bool`

GetIcsma returns the Icsma field if non-nil, zero value otherwise.

### GetIcsmaOk

`func (o *AdvisoryCISAAlert) GetIcsmaOk() (*bool, bool)`

GetIcsmaOk returns a tuple with the Icsma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcsma

`func (o *AdvisoryCISAAlert) SetIcsma(v bool)`

SetIcsma sets Icsma field to given value.

### HasIcsma

`func (o *AdvisoryCISAAlert) HasIcsma() bool`

HasIcsma returns a boolean if a field has been set.

### GetMitigations

`func (o *AdvisoryCISAAlert) GetMitigations() string`

GetMitigations returns the Mitigations field if non-nil, zero value otherwise.

### GetMitigationsOk

`func (o *AdvisoryCISAAlert) GetMitigationsOk() (*string, bool)`

GetMitigationsOk returns a tuple with the Mitigations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigations

`func (o *AdvisoryCISAAlert) SetMitigations(v string)`

SetMitigations sets Mitigations field to given value.

### HasMitigations

`func (o *AdvisoryCISAAlert) HasMitigations() bool`

HasMitigations returns a boolean if a field has been set.

### GetReleaseDate

`func (o *AdvisoryCISAAlert) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *AdvisoryCISAAlert) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *AdvisoryCISAAlert) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *AdvisoryCISAAlert) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryCISAAlert) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryCISAAlert) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryCISAAlert) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryCISAAlert) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryCISAAlert) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryCISAAlert) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryCISAAlert) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryCISAAlert) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryCISAAlert) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryCISAAlert) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryCISAAlert) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryCISAAlert) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVendor

`func (o *AdvisoryCISAAlert) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AdvisoryCISAAlert) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AdvisoryCISAAlert) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AdvisoryCISAAlert) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


