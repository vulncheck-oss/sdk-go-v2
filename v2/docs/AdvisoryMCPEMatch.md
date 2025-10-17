# AdvisoryMCPEMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | Pointer to **string** |  | [optional] 
**MatchCriteriaId** | Pointer to **string** |  | [optional] 
**VersionEndExcluding** | Pointer to **string** |  | [optional] 
**VersionEndIncluding** | Pointer to **string** |  | [optional] 
**VersionStartExcluding** | Pointer to **string** |  | [optional] 
**VersionStartIncluding** | Pointer to **string** |  | [optional] 
**Vulnerable** | Pointer to **bool** |  | [optional] 

## Methods

### NewAdvisoryMCPEMatch

`func NewAdvisoryMCPEMatch() *AdvisoryMCPEMatch`

NewAdvisoryMCPEMatch instantiates a new AdvisoryMCPEMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMCPEMatchWithDefaults

`func NewAdvisoryMCPEMatchWithDefaults() *AdvisoryMCPEMatch`

NewAdvisoryMCPEMatchWithDefaults instantiates a new AdvisoryMCPEMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *AdvisoryMCPEMatch) GetCriteria() string`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *AdvisoryMCPEMatch) GetCriteriaOk() (*string, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *AdvisoryMCPEMatch) SetCriteria(v string)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *AdvisoryMCPEMatch) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetMatchCriteriaId

`func (o *AdvisoryMCPEMatch) GetMatchCriteriaId() string`

GetMatchCriteriaId returns the MatchCriteriaId field if non-nil, zero value otherwise.

### GetMatchCriteriaIdOk

`func (o *AdvisoryMCPEMatch) GetMatchCriteriaIdOk() (*string, bool)`

GetMatchCriteriaIdOk returns a tuple with the MatchCriteriaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCriteriaId

`func (o *AdvisoryMCPEMatch) SetMatchCriteriaId(v string)`

SetMatchCriteriaId sets MatchCriteriaId field to given value.

### HasMatchCriteriaId

`func (o *AdvisoryMCPEMatch) HasMatchCriteriaId() bool`

HasMatchCriteriaId returns a boolean if a field has been set.

### GetVersionEndExcluding

`func (o *AdvisoryMCPEMatch) GetVersionEndExcluding() string`

GetVersionEndExcluding returns the VersionEndExcluding field if non-nil, zero value otherwise.

### GetVersionEndExcludingOk

`func (o *AdvisoryMCPEMatch) GetVersionEndExcludingOk() (*string, bool)`

GetVersionEndExcludingOk returns a tuple with the VersionEndExcluding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionEndExcluding

`func (o *AdvisoryMCPEMatch) SetVersionEndExcluding(v string)`

SetVersionEndExcluding sets VersionEndExcluding field to given value.

### HasVersionEndExcluding

`func (o *AdvisoryMCPEMatch) HasVersionEndExcluding() bool`

HasVersionEndExcluding returns a boolean if a field has been set.

### GetVersionEndIncluding

`func (o *AdvisoryMCPEMatch) GetVersionEndIncluding() string`

GetVersionEndIncluding returns the VersionEndIncluding field if non-nil, zero value otherwise.

### GetVersionEndIncludingOk

`func (o *AdvisoryMCPEMatch) GetVersionEndIncludingOk() (*string, bool)`

GetVersionEndIncludingOk returns a tuple with the VersionEndIncluding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionEndIncluding

`func (o *AdvisoryMCPEMatch) SetVersionEndIncluding(v string)`

SetVersionEndIncluding sets VersionEndIncluding field to given value.

### HasVersionEndIncluding

`func (o *AdvisoryMCPEMatch) HasVersionEndIncluding() bool`

HasVersionEndIncluding returns a boolean if a field has been set.

### GetVersionStartExcluding

`func (o *AdvisoryMCPEMatch) GetVersionStartExcluding() string`

GetVersionStartExcluding returns the VersionStartExcluding field if non-nil, zero value otherwise.

### GetVersionStartExcludingOk

`func (o *AdvisoryMCPEMatch) GetVersionStartExcludingOk() (*string, bool)`

GetVersionStartExcludingOk returns a tuple with the VersionStartExcluding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionStartExcluding

`func (o *AdvisoryMCPEMatch) SetVersionStartExcluding(v string)`

SetVersionStartExcluding sets VersionStartExcluding field to given value.

### HasVersionStartExcluding

`func (o *AdvisoryMCPEMatch) HasVersionStartExcluding() bool`

HasVersionStartExcluding returns a boolean if a field has been set.

### GetVersionStartIncluding

`func (o *AdvisoryMCPEMatch) GetVersionStartIncluding() string`

GetVersionStartIncluding returns the VersionStartIncluding field if non-nil, zero value otherwise.

### GetVersionStartIncludingOk

`func (o *AdvisoryMCPEMatch) GetVersionStartIncludingOk() (*string, bool)`

GetVersionStartIncludingOk returns a tuple with the VersionStartIncluding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionStartIncluding

`func (o *AdvisoryMCPEMatch) SetVersionStartIncluding(v string)`

SetVersionStartIncluding sets VersionStartIncluding field to given value.

### HasVersionStartIncluding

`func (o *AdvisoryMCPEMatch) HasVersionStartIncluding() bool`

HasVersionStartIncluding returns a boolean if a field has been set.

### GetVulnerable

`func (o *AdvisoryMCPEMatch) GetVulnerable() bool`

GetVulnerable returns the Vulnerable field if non-nil, zero value otherwise.

### GetVulnerableOk

`func (o *AdvisoryMCPEMatch) GetVulnerableOk() (*bool, bool)`

GetVulnerableOk returns a tuple with the Vulnerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerable

`func (o *AdvisoryMCPEMatch) SetVulnerable(v bool)`

SetVulnerable sets Vulnerable field to given value.

### HasVulnerable

`func (o *AdvisoryMCPEMatch) HasVulnerable() bool`

HasVulnerable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


