# AdvisoryGallagher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveExploitation** | Pointer to **bool** |  | [optional] 
**Affected** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Fixes** | Pointer to **string** |  | [optional] 
**ReportedBy** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryGallagher

`func NewAdvisoryGallagher() *AdvisoryGallagher`

NewAdvisoryGallagher instantiates a new AdvisoryGallagher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryGallagherWithDefaults

`func NewAdvisoryGallagherWithDefaults() *AdvisoryGallagher`

NewAdvisoryGallagherWithDefaults instantiates a new AdvisoryGallagher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveExploitation

`func (o *AdvisoryGallagher) GetActiveExploitation() bool`

GetActiveExploitation returns the ActiveExploitation field if non-nil, zero value otherwise.

### GetActiveExploitationOk

`func (o *AdvisoryGallagher) GetActiveExploitationOk() (*bool, bool)`

GetActiveExploitationOk returns a tuple with the ActiveExploitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveExploitation

`func (o *AdvisoryGallagher) SetActiveExploitation(v bool)`

SetActiveExploitation sets ActiveExploitation field to given value.

### HasActiveExploitation

`func (o *AdvisoryGallagher) HasActiveExploitation() bool`

HasActiveExploitation returns a boolean if a field has been set.

### GetAffected

`func (o *AdvisoryGallagher) GetAffected() string`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryGallagher) GetAffectedOk() (*string, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryGallagher) SetAffected(v string)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryGallagher) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryGallagher) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryGallagher) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryGallagher) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryGallagher) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryGallagher) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryGallagher) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryGallagher) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryGallagher) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryGallagher) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryGallagher) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryGallagher) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryGallagher) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFixes

`func (o *AdvisoryGallagher) GetFixes() string`

GetFixes returns the Fixes field if non-nil, zero value otherwise.

### GetFixesOk

`func (o *AdvisoryGallagher) GetFixesOk() (*string, bool)`

GetFixesOk returns a tuple with the Fixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixes

`func (o *AdvisoryGallagher) SetFixes(v string)`

SetFixes sets Fixes field to given value.

### HasFixes

`func (o *AdvisoryGallagher) HasFixes() bool`

HasFixes returns a boolean if a field has been set.

### GetReportedBy

`func (o *AdvisoryGallagher) GetReportedBy() string`

GetReportedBy returns the ReportedBy field if non-nil, zero value otherwise.

### GetReportedByOk

`func (o *AdvisoryGallagher) GetReportedByOk() (*string, bool)`

GetReportedByOk returns a tuple with the ReportedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedBy

`func (o *AdvisoryGallagher) SetReportedBy(v string)`

SetReportedBy sets ReportedBy field to given value.

### HasReportedBy

`func (o *AdvisoryGallagher) HasReportedBy() bool`

HasReportedBy returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryGallagher) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryGallagher) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryGallagher) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryGallagher) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryGallagher) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryGallagher) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryGallagher) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryGallagher) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryGallagher) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryGallagher) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryGallagher) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryGallagher) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


