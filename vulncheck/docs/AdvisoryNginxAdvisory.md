# AdvisoryNginxAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**NotVulnVersions** | Pointer to **[]string** |  | [optional] 
**PatchPgp** | Pointer to **string** |  | [optional] 
**PatchUrl** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**VulnVersions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdvisoryNginxAdvisory

`func NewAdvisoryNginxAdvisory() *AdvisoryNginxAdvisory`

NewAdvisoryNginxAdvisory instantiates a new AdvisoryNginxAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryNginxAdvisoryWithDefaults

`func NewAdvisoryNginxAdvisoryWithDefaults() *AdvisoryNginxAdvisory`

NewAdvisoryNginxAdvisoryWithDefaults instantiates a new AdvisoryNginxAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryNginxAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryNginxAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryNginxAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryNginxAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryNginxAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryNginxAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryNginxAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryNginxAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryNginxAdvisory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryNginxAdvisory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryNginxAdvisory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryNginxAdvisory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotVulnVersions

`func (o *AdvisoryNginxAdvisory) GetNotVulnVersions() []string`

GetNotVulnVersions returns the NotVulnVersions field if non-nil, zero value otherwise.

### GetNotVulnVersionsOk

`func (o *AdvisoryNginxAdvisory) GetNotVulnVersionsOk() (*[]string, bool)`

GetNotVulnVersionsOk returns a tuple with the NotVulnVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotVulnVersions

`func (o *AdvisoryNginxAdvisory) SetNotVulnVersions(v []string)`

SetNotVulnVersions sets NotVulnVersions field to given value.

### HasNotVulnVersions

`func (o *AdvisoryNginxAdvisory) HasNotVulnVersions() bool`

HasNotVulnVersions returns a boolean if a field has been set.

### GetPatchPgp

`func (o *AdvisoryNginxAdvisory) GetPatchPgp() string`

GetPatchPgp returns the PatchPgp field if non-nil, zero value otherwise.

### GetPatchPgpOk

`func (o *AdvisoryNginxAdvisory) GetPatchPgpOk() (*string, bool)`

GetPatchPgpOk returns a tuple with the PatchPgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPgp

`func (o *AdvisoryNginxAdvisory) SetPatchPgp(v string)`

SetPatchPgp sets PatchPgp field to given value.

### HasPatchPgp

`func (o *AdvisoryNginxAdvisory) HasPatchPgp() bool`

HasPatchPgp returns a boolean if a field has been set.

### GetPatchUrl

`func (o *AdvisoryNginxAdvisory) GetPatchUrl() string`

GetPatchUrl returns the PatchUrl field if non-nil, zero value otherwise.

### GetPatchUrlOk

`func (o *AdvisoryNginxAdvisory) GetPatchUrlOk() (*string, bool)`

GetPatchUrlOk returns a tuple with the PatchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchUrl

`func (o *AdvisoryNginxAdvisory) SetPatchUrl(v string)`

SetPatchUrl sets PatchUrl field to given value.

### HasPatchUrl

`func (o *AdvisoryNginxAdvisory) HasPatchUrl() bool`

HasPatchUrl returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryNginxAdvisory) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryNginxAdvisory) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryNginxAdvisory) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryNginxAdvisory) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryNginxAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryNginxAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryNginxAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryNginxAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVulnVersions

`func (o *AdvisoryNginxAdvisory) GetVulnVersions() []string`

GetVulnVersions returns the VulnVersions field if non-nil, zero value otherwise.

### GetVulnVersionsOk

`func (o *AdvisoryNginxAdvisory) GetVulnVersionsOk() (*[]string, bool)`

GetVulnVersionsOk returns a tuple with the VulnVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnVersions

`func (o *AdvisoryNginxAdvisory) SetVulnVersions(v []string)`

SetVulnVersions sets VulnVersions field to given value.

### HasVulnVersions

`func (o *AdvisoryNginxAdvisory) HasVulnVersions() bool`

HasVulnVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


