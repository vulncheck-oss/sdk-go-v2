# AdvisoryUnisoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessVector** | Pointer to **string** |  | [optional] 
**AffectedChipsets** | Pointer to **string** |  | [optional] 
**AffectedSoftware** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Rating** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Technology** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Vulnerability** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryUnisoc

`func NewAdvisoryUnisoc() *AdvisoryUnisoc`

NewAdvisoryUnisoc instantiates a new AdvisoryUnisoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryUnisocWithDefaults

`func NewAdvisoryUnisocWithDefaults() *AdvisoryUnisoc`

NewAdvisoryUnisocWithDefaults instantiates a new AdvisoryUnisoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessVector

`func (o *AdvisoryUnisoc) GetAccessVector() string`

GetAccessVector returns the AccessVector field if non-nil, zero value otherwise.

### GetAccessVectorOk

`func (o *AdvisoryUnisoc) GetAccessVectorOk() (*string, bool)`

GetAccessVectorOk returns a tuple with the AccessVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessVector

`func (o *AdvisoryUnisoc) SetAccessVector(v string)`

SetAccessVector sets AccessVector field to given value.

### HasAccessVector

`func (o *AdvisoryUnisoc) HasAccessVector() bool`

HasAccessVector returns a boolean if a field has been set.

### GetAffectedChipsets

`func (o *AdvisoryUnisoc) GetAffectedChipsets() string`

GetAffectedChipsets returns the AffectedChipsets field if non-nil, zero value otherwise.

### GetAffectedChipsetsOk

`func (o *AdvisoryUnisoc) GetAffectedChipsetsOk() (*string, bool)`

GetAffectedChipsetsOk returns a tuple with the AffectedChipsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedChipsets

`func (o *AdvisoryUnisoc) SetAffectedChipsets(v string)`

SetAffectedChipsets sets AffectedChipsets field to given value.

### HasAffectedChipsets

`func (o *AdvisoryUnisoc) HasAffectedChipsets() bool`

HasAffectedChipsets returns a boolean if a field has been set.

### GetAffectedSoftware

`func (o *AdvisoryUnisoc) GetAffectedSoftware() string`

GetAffectedSoftware returns the AffectedSoftware field if non-nil, zero value otherwise.

### GetAffectedSoftwareOk

`func (o *AdvisoryUnisoc) GetAffectedSoftwareOk() (*string, bool)`

GetAffectedSoftwareOk returns a tuple with the AffectedSoftware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedSoftware

`func (o *AdvisoryUnisoc) SetAffectedSoftware(v string)`

SetAffectedSoftware sets AffectedSoftware field to given value.

### HasAffectedSoftware

`func (o *AdvisoryUnisoc) HasAffectedSoftware() bool`

HasAffectedSoftware returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryUnisoc) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryUnisoc) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryUnisoc) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryUnisoc) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryUnisoc) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryUnisoc) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryUnisoc) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryUnisoc) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryUnisoc) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryUnisoc) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryUnisoc) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryUnisoc) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRating

`func (o *AdvisoryUnisoc) GetRating() string`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *AdvisoryUnisoc) GetRatingOk() (*string, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *AdvisoryUnisoc) SetRating(v string)`

SetRating sets Rating field to given value.

### HasRating

`func (o *AdvisoryUnisoc) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetScore

`func (o *AdvisoryUnisoc) GetScore() string`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *AdvisoryUnisoc) GetScoreOk() (*string, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *AdvisoryUnisoc) SetScore(v string)`

SetScore sets Score field to given value.

### HasScore

`func (o *AdvisoryUnisoc) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryUnisoc) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryUnisoc) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryUnisoc) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryUnisoc) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTechnology

`func (o *AdvisoryUnisoc) GetTechnology() string`

GetTechnology returns the Technology field if non-nil, zero value otherwise.

### GetTechnologyOk

`func (o *AdvisoryUnisoc) GetTechnologyOk() (*string, bool)`

GetTechnologyOk returns a tuple with the Technology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnology

`func (o *AdvisoryUnisoc) SetTechnology(v string)`

SetTechnology sets Technology field to given value.

### HasTechnology

`func (o *AdvisoryUnisoc) HasTechnology() bool`

HasTechnology returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryUnisoc) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryUnisoc) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryUnisoc) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryUnisoc) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryUnisoc) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryUnisoc) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryUnisoc) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryUnisoc) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVulnerability

`func (o *AdvisoryUnisoc) GetVulnerability() string`

GetVulnerability returns the Vulnerability field if non-nil, zero value otherwise.

### GetVulnerabilityOk

`func (o *AdvisoryUnisoc) GetVulnerabilityOk() (*string, bool)`

GetVulnerabilityOk returns a tuple with the Vulnerability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerability

`func (o *AdvisoryUnisoc) SetVulnerability(v string)`

SetVulnerability sets Vulnerability field to given value.

### HasVulnerability

`func (o *AdvisoryUnisoc) HasVulnerability() bool`

HasVulnerability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


