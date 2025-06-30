# ApiCveItemsExtended

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** |  | [optional] 
**Configurations** | Pointer to [**ApiConfigurations**](ApiConfigurations.md) |  | [optional] 
**Cve** | Pointer to [**ApiCVEExtended**](ApiCVEExtended.md) |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**DocumentGenerationDate** | Pointer to **string** | the deep tag instructs deep.Equal to ignore this field (used during OpenSearch loading) | [optional] 
**Impact** | Pointer to [**ApiImpactExtended**](ApiImpactExtended.md) |  | [optional] 
**LastModifiedDate** | Pointer to **string** |  | [optional] 
**MitreAttackTechniques** | Pointer to [**[]ApiMitreAttackTech**](ApiMitreAttackTech.md) |  | [optional] 
**PublishedDate** | Pointer to **string** |  | [optional] 
**RelatedAttackPatterns** | Pointer to [**[]ApiRelatedAttackPattern**](ApiRelatedAttackPattern.md) |  | [optional] 
**VcConfigurations** | Pointer to [**ApiConfigurations**](ApiConfigurations.md) |  | [optional] 
**VcVulnerableCPEs** | Pointer to **[]string** |  | [optional] 
**VulnerableCpes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApiCveItemsExtended

`func NewApiCveItemsExtended() *ApiCveItemsExtended`

NewApiCveItemsExtended instantiates a new ApiCveItemsExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCveItemsExtendedWithDefaults

`func NewApiCveItemsExtendedWithDefaults() *ApiCveItemsExtended`

NewApiCveItemsExtendedWithDefaults instantiates a new ApiCveItemsExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ApiCveItemsExtended) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApiCveItemsExtended) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApiCveItemsExtended) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ApiCveItemsExtended) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetConfigurations

`func (o *ApiCveItemsExtended) GetConfigurations() ApiConfigurations`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *ApiCveItemsExtended) GetConfigurationsOk() (*ApiConfigurations, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *ApiCveItemsExtended) SetConfigurations(v ApiConfigurations)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *ApiCveItemsExtended) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetCve

`func (o *ApiCveItemsExtended) GetCve() ApiCVEExtended`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *ApiCveItemsExtended) GetCveOk() (*ApiCVEExtended, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *ApiCveItemsExtended) SetCve(v ApiCVEExtended)`

SetCve sets Cve field to given value.

### HasCve

`func (o *ApiCveItemsExtended) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *ApiCveItemsExtended) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *ApiCveItemsExtended) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *ApiCveItemsExtended) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *ApiCveItemsExtended) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDocumentGenerationDate

`func (o *ApiCveItemsExtended) GetDocumentGenerationDate() string`

GetDocumentGenerationDate returns the DocumentGenerationDate field if non-nil, zero value otherwise.

### GetDocumentGenerationDateOk

`func (o *ApiCveItemsExtended) GetDocumentGenerationDateOk() (*string, bool)`

GetDocumentGenerationDateOk returns a tuple with the DocumentGenerationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentGenerationDate

`func (o *ApiCveItemsExtended) SetDocumentGenerationDate(v string)`

SetDocumentGenerationDate sets DocumentGenerationDate field to given value.

### HasDocumentGenerationDate

`func (o *ApiCveItemsExtended) HasDocumentGenerationDate() bool`

HasDocumentGenerationDate returns a boolean if a field has been set.

### GetImpact

`func (o *ApiCveItemsExtended) GetImpact() ApiImpactExtended`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *ApiCveItemsExtended) GetImpactOk() (*ApiImpactExtended, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *ApiCveItemsExtended) SetImpact(v ApiImpactExtended)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *ApiCveItemsExtended) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *ApiCveItemsExtended) GetLastModifiedDate() string`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *ApiCveItemsExtended) GetLastModifiedDateOk() (*string, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *ApiCveItemsExtended) SetLastModifiedDate(v string)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *ApiCveItemsExtended) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetMitreAttackTechniques

`func (o *ApiCveItemsExtended) GetMitreAttackTechniques() []ApiMitreAttackTech`

GetMitreAttackTechniques returns the MitreAttackTechniques field if non-nil, zero value otherwise.

### GetMitreAttackTechniquesOk

`func (o *ApiCveItemsExtended) GetMitreAttackTechniquesOk() (*[]ApiMitreAttackTech, bool)`

GetMitreAttackTechniquesOk returns a tuple with the MitreAttackTechniques field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitreAttackTechniques

`func (o *ApiCveItemsExtended) SetMitreAttackTechniques(v []ApiMitreAttackTech)`

SetMitreAttackTechniques sets MitreAttackTechniques field to given value.

### HasMitreAttackTechniques

`func (o *ApiCveItemsExtended) HasMitreAttackTechniques() bool`

HasMitreAttackTechniques returns a boolean if a field has been set.

### GetPublishedDate

`func (o *ApiCveItemsExtended) GetPublishedDate() string`

GetPublishedDate returns the PublishedDate field if non-nil, zero value otherwise.

### GetPublishedDateOk

`func (o *ApiCveItemsExtended) GetPublishedDateOk() (*string, bool)`

GetPublishedDateOk returns a tuple with the PublishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDate

`func (o *ApiCveItemsExtended) SetPublishedDate(v string)`

SetPublishedDate sets PublishedDate field to given value.

### HasPublishedDate

`func (o *ApiCveItemsExtended) HasPublishedDate() bool`

HasPublishedDate returns a boolean if a field has been set.

### GetRelatedAttackPatterns

`func (o *ApiCveItemsExtended) GetRelatedAttackPatterns() []ApiRelatedAttackPattern`

GetRelatedAttackPatterns returns the RelatedAttackPatterns field if non-nil, zero value otherwise.

### GetRelatedAttackPatternsOk

`func (o *ApiCveItemsExtended) GetRelatedAttackPatternsOk() (*[]ApiRelatedAttackPattern, bool)`

GetRelatedAttackPatternsOk returns a tuple with the RelatedAttackPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedAttackPatterns

`func (o *ApiCveItemsExtended) SetRelatedAttackPatterns(v []ApiRelatedAttackPattern)`

SetRelatedAttackPatterns sets RelatedAttackPatterns field to given value.

### HasRelatedAttackPatterns

`func (o *ApiCveItemsExtended) HasRelatedAttackPatterns() bool`

HasRelatedAttackPatterns returns a boolean if a field has been set.

### GetVcConfigurations

`func (o *ApiCveItemsExtended) GetVcConfigurations() ApiConfigurations`

GetVcConfigurations returns the VcConfigurations field if non-nil, zero value otherwise.

### GetVcConfigurationsOk

`func (o *ApiCveItemsExtended) GetVcConfigurationsOk() (*ApiConfigurations, bool)`

GetVcConfigurationsOk returns a tuple with the VcConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcConfigurations

`func (o *ApiCveItemsExtended) SetVcConfigurations(v ApiConfigurations)`

SetVcConfigurations sets VcConfigurations field to given value.

### HasVcConfigurations

`func (o *ApiCveItemsExtended) HasVcConfigurations() bool`

HasVcConfigurations returns a boolean if a field has been set.

### GetVcVulnerableCPEs

`func (o *ApiCveItemsExtended) GetVcVulnerableCPEs() []string`

GetVcVulnerableCPEs returns the VcVulnerableCPEs field if non-nil, zero value otherwise.

### GetVcVulnerableCPEsOk

`func (o *ApiCveItemsExtended) GetVcVulnerableCPEsOk() (*[]string, bool)`

GetVcVulnerableCPEsOk returns a tuple with the VcVulnerableCPEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcVulnerableCPEs

`func (o *ApiCveItemsExtended) SetVcVulnerableCPEs(v []string)`

SetVcVulnerableCPEs sets VcVulnerableCPEs field to given value.

### HasVcVulnerableCPEs

`func (o *ApiCveItemsExtended) HasVcVulnerableCPEs() bool`

HasVcVulnerableCPEs returns a boolean if a field has been set.

### GetVulnerableCpes

`func (o *ApiCveItemsExtended) GetVulnerableCpes() []string`

GetVulnerableCpes returns the VulnerableCpes field if non-nil, zero value otherwise.

### GetVulnerableCpesOk

`func (o *ApiCveItemsExtended) GetVulnerableCpesOk() (*[]string, bool)`

GetVulnerableCpesOk returns a tuple with the VulnerableCpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableCpes

`func (o *ApiCveItemsExtended) SetVulnerableCpes(v []string)`

SetVulnerableCpes sets VulnerableCpes field to given value.

### HasVulnerableCpes

`func (o *ApiCveItemsExtended) HasVulnerableCpes() bool`

HasVulnerableCpes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


