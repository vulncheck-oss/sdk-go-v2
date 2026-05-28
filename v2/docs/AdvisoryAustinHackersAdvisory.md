# AdvisoryAustinHackersAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credit** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**CvssScore** | Pointer to **float32** | From detail page (omitted for external-linked CVEs) | [optional] 
**CvssVector** | Pointer to **string** |  | [optional] 
**Cwe** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Gcve** | Pointer to **string** |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**Mitigation** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** | From index table | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryAustinHackersAdvisory

`func NewAdvisoryAustinHackersAdvisory() *AdvisoryAustinHackersAdvisory`

NewAdvisoryAustinHackersAdvisory instantiates a new AdvisoryAustinHackersAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAustinHackersAdvisoryWithDefaults

`func NewAdvisoryAustinHackersAdvisoryWithDefaults() *AdvisoryAustinHackersAdvisory`

NewAdvisoryAustinHackersAdvisoryWithDefaults instantiates a new AdvisoryAustinHackersAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredit

`func (o *AdvisoryAustinHackersAdvisory) GetCredit() string`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *AdvisoryAustinHackersAdvisory) GetCreditOk() (*string, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *AdvisoryAustinHackersAdvisory) SetCredit(v string)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *AdvisoryAustinHackersAdvisory) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryAustinHackersAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryAustinHackersAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryAustinHackersAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryAustinHackersAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvssScore

`func (o *AdvisoryAustinHackersAdvisory) GetCvssScore() float32`

GetCvssScore returns the CvssScore field if non-nil, zero value otherwise.

### GetCvssScoreOk

`func (o *AdvisoryAustinHackersAdvisory) GetCvssScoreOk() (*float32, bool)`

GetCvssScoreOk returns a tuple with the CvssScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssScore

`func (o *AdvisoryAustinHackersAdvisory) SetCvssScore(v float32)`

SetCvssScore sets CvssScore field to given value.

### HasCvssScore

`func (o *AdvisoryAustinHackersAdvisory) HasCvssScore() bool`

HasCvssScore returns a boolean if a field has been set.

### GetCvssVector

`func (o *AdvisoryAustinHackersAdvisory) GetCvssVector() string`

GetCvssVector returns the CvssVector field if non-nil, zero value otherwise.

### GetCvssVectorOk

`func (o *AdvisoryAustinHackersAdvisory) GetCvssVectorOk() (*string, bool)`

GetCvssVectorOk returns a tuple with the CvssVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVector

`func (o *AdvisoryAustinHackersAdvisory) SetCvssVector(v string)`

SetCvssVector sets CvssVector field to given value.

### HasCvssVector

`func (o *AdvisoryAustinHackersAdvisory) HasCvssVector() bool`

HasCvssVector returns a boolean if a field has been set.

### GetCwe

`func (o *AdvisoryAustinHackersAdvisory) GetCwe() string`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *AdvisoryAustinHackersAdvisory) GetCweOk() (*string, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *AdvisoryAustinHackersAdvisory) SetCwe(v string)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *AdvisoryAustinHackersAdvisory) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryAustinHackersAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryAustinHackersAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryAustinHackersAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryAustinHackersAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryAustinHackersAdvisory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryAustinHackersAdvisory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryAustinHackersAdvisory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryAustinHackersAdvisory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGcve

`func (o *AdvisoryAustinHackersAdvisory) GetGcve() string`

GetGcve returns the Gcve field if non-nil, zero value otherwise.

### GetGcveOk

`func (o *AdvisoryAustinHackersAdvisory) GetGcveOk() (*string, bool)`

GetGcveOk returns a tuple with the Gcve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcve

`func (o *AdvisoryAustinHackersAdvisory) SetGcve(v string)`

SetGcve sets Gcve field to given value.

### HasGcve

`func (o *AdvisoryAustinHackersAdvisory) HasGcve() bool`

HasGcve returns a boolean if a field has been set.

### GetImpact

`func (o *AdvisoryAustinHackersAdvisory) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *AdvisoryAustinHackersAdvisory) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *AdvisoryAustinHackersAdvisory) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *AdvisoryAustinHackersAdvisory) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetMitigation

`func (o *AdvisoryAustinHackersAdvisory) GetMitigation() string`

GetMitigation returns the Mitigation field if non-nil, zero value otherwise.

### GetMitigationOk

`func (o *AdvisoryAustinHackersAdvisory) GetMitigationOk() (*string, bool)`

GetMitigationOk returns a tuple with the Mitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigation

`func (o *AdvisoryAustinHackersAdvisory) SetMitigation(v string)`

SetMitigation sets Mitigation field to given value.

### HasMitigation

`func (o *AdvisoryAustinHackersAdvisory) HasMitigation() bool`

HasMitigation returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryAustinHackersAdvisory) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryAustinHackersAdvisory) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryAustinHackersAdvisory) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryAustinHackersAdvisory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryAustinHackersAdvisory) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryAustinHackersAdvisory) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryAustinHackersAdvisory) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryAustinHackersAdvisory) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryAustinHackersAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryAustinHackersAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryAustinHackersAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryAustinHackersAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryAustinHackersAdvisory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryAustinHackersAdvisory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryAustinHackersAdvisory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryAustinHackersAdvisory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryAustinHackersAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryAustinHackersAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryAustinHackersAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryAustinHackersAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


