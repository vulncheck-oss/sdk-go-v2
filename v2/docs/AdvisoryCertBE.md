# AdvisoryCertBE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedSoftware** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Mitigation** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Risk** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**VulnerabilityType** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdvisoryCertBE

`func NewAdvisoryCertBE() *AdvisoryCertBE`

NewAdvisoryCertBE instantiates a new AdvisoryCertBE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCertBEWithDefaults

`func NewAdvisoryCertBEWithDefaults() *AdvisoryCertBE`

NewAdvisoryCertBEWithDefaults instantiates a new AdvisoryCertBE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedSoftware

`func (o *AdvisoryCertBE) GetAffectedSoftware() []string`

GetAffectedSoftware returns the AffectedSoftware field if non-nil, zero value otherwise.

### GetAffectedSoftwareOk

`func (o *AdvisoryCertBE) GetAffectedSoftwareOk() (*[]string, bool)`

GetAffectedSoftwareOk returns a tuple with the AffectedSoftware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedSoftware

`func (o *AdvisoryCertBE) SetAffectedSoftware(v []string)`

SetAffectedSoftware sets AffectedSoftware field to given value.

### HasAffectedSoftware

`func (o *AdvisoryCertBE) HasAffectedSoftware() bool`

HasAffectedSoftware returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryCertBE) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryCertBE) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryCertBE) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryCertBE) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryCertBE) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryCertBE) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryCertBE) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryCertBE) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryCertBE) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryCertBE) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryCertBE) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryCertBE) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMitigation

`func (o *AdvisoryCertBE) GetMitigation() string`

GetMitigation returns the Mitigation field if non-nil, zero value otherwise.

### GetMitigationOk

`func (o *AdvisoryCertBE) GetMitigationOk() (*string, bool)`

GetMitigationOk returns a tuple with the Mitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigation

`func (o *AdvisoryCertBE) SetMitigation(v string)`

SetMitigation sets Mitigation field to given value.

### HasMitigation

`func (o *AdvisoryCertBE) HasMitigation() bool`

HasMitigation returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryCertBE) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryCertBE) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryCertBE) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryCertBE) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRisk

`func (o *AdvisoryCertBE) GetRisk() string`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *AdvisoryCertBE) GetRiskOk() (*string, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *AdvisoryCertBE) SetRisk(v string)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *AdvisoryCertBE) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryCertBE) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryCertBE) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryCertBE) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryCertBE) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryCertBE) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryCertBE) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryCertBE) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryCertBE) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryCertBE) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryCertBE) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryCertBE) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryCertBE) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryCertBE) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryCertBE) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryCertBE) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryCertBE) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVulnerabilityType

`func (o *AdvisoryCertBE) GetVulnerabilityType() []string`

GetVulnerabilityType returns the VulnerabilityType field if non-nil, zero value otherwise.

### GetVulnerabilityTypeOk

`func (o *AdvisoryCertBE) GetVulnerabilityTypeOk() (*[]string, bool)`

GetVulnerabilityTypeOk returns a tuple with the VulnerabilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityType

`func (o *AdvisoryCertBE) SetVulnerabilityType(v []string)`

SetVulnerabilityType sets VulnerabilityType field to given value.

### HasVulnerabilityType

`func (o *AdvisoryCertBE) HasVulnerabilityType() bool`

HasVulnerabilityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


