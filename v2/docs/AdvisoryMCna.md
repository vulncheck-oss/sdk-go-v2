# AdvisoryMCna

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to [**[]AdvisoryMAffected**](AdvisoryMAffected.md) |  | [optional] 
**Configurations** | Pointer to [**[]AdvisoryMDescriptions**](AdvisoryMDescriptions.md) |  | [optional] 
**CpeApplicability** | Pointer to [**[]AdvisoryCustomCPE**](AdvisoryCustomCPE.md) |  | [optional] 
**Credits** | Pointer to [**[]AdvisoryCredit**](AdvisoryCredit.md) |  | [optional] 
**DateAssigned** | Pointer to **string** |  | [optional] 
**DatePublic** | Pointer to **string** |  | [optional] 
**Descriptions** | Pointer to [**[]AdvisoryMDescriptions**](AdvisoryMDescriptions.md) |  | [optional] 
**Exploits** | Pointer to [**[]AdvisoryMDescriptions**](AdvisoryMDescriptions.md) |  | [optional] 
**Impacts** | Pointer to [**[]AdvisoryImpact**](AdvisoryImpact.md) |  | [optional] 
**Metrics** | Pointer to [**[]AdvisoryMetric**](AdvisoryMetric.md) |  | [optional] 
**ProblemTypes** | Pointer to [**[]AdvisoryMProblemTypes**](AdvisoryMProblemTypes.md) |  | [optional] 
**ProviderMetadata** | Pointer to [**AdvisoryMProviderMetadata**](AdvisoryMProviderMetadata.md) |  | [optional] 
**References** | Pointer to [**[]AdvisoryMReference**](AdvisoryMReference.md) |  | [optional] 
**RejectedReasons** | Pointer to [**[]AdvisoryMDescriptions**](AdvisoryMDescriptions.md) | Fields below appear only on rejected records (cveMetadata.state &#x3D;&#x3D; \&quot;REJECTED\&quot;). | [optional] 
**ReplacedBy** | Pointer to **[]string** |  | [optional] 
**Solutions** | Pointer to [**[]AdvisoryMDescriptions**](AdvisoryMDescriptions.md) |  | [optional] 
**Source** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TaxonomyMappings** | Pointer to [**[]AdvisoryTaxonomyMapping**](AdvisoryTaxonomyMapping.md) |  | [optional] 
**Timeline** | Pointer to [**[]AdvisoryTimeline**](AdvisoryTimeline.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Workarounds** | Pointer to [**[]AdvisoryMDescriptions**](AdvisoryMDescriptions.md) |  | [optional] 

## Methods

### NewAdvisoryMCna

`func NewAdvisoryMCna() *AdvisoryMCna`

NewAdvisoryMCna instantiates a new AdvisoryMCna object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMCnaWithDefaults

`func NewAdvisoryMCnaWithDefaults() *AdvisoryMCna`

NewAdvisoryMCnaWithDefaults instantiates a new AdvisoryMCna object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *AdvisoryMCna) GetAffected() []AdvisoryMAffected`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryMCna) GetAffectedOk() (*[]AdvisoryMAffected, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryMCna) SetAffected(v []AdvisoryMAffected)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryMCna) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetConfigurations

`func (o *AdvisoryMCna) GetConfigurations() []AdvisoryMDescriptions`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *AdvisoryMCna) GetConfigurationsOk() (*[]AdvisoryMDescriptions, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *AdvisoryMCna) SetConfigurations(v []AdvisoryMDescriptions)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *AdvisoryMCna) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetCpeApplicability

`func (o *AdvisoryMCna) GetCpeApplicability() []AdvisoryCustomCPE`

GetCpeApplicability returns the CpeApplicability field if non-nil, zero value otherwise.

### GetCpeApplicabilityOk

`func (o *AdvisoryMCna) GetCpeApplicabilityOk() (*[]AdvisoryCustomCPE, bool)`

GetCpeApplicabilityOk returns a tuple with the CpeApplicability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeApplicability

`func (o *AdvisoryMCna) SetCpeApplicability(v []AdvisoryCustomCPE)`

SetCpeApplicability sets CpeApplicability field to given value.

### HasCpeApplicability

`func (o *AdvisoryMCna) HasCpeApplicability() bool`

HasCpeApplicability returns a boolean if a field has been set.

### GetCredits

`func (o *AdvisoryMCna) GetCredits() []AdvisoryCredit`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *AdvisoryMCna) GetCreditsOk() (*[]AdvisoryCredit, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *AdvisoryMCna) SetCredits(v []AdvisoryCredit)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *AdvisoryMCna) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetDateAssigned

`func (o *AdvisoryMCna) GetDateAssigned() string`

GetDateAssigned returns the DateAssigned field if non-nil, zero value otherwise.

### GetDateAssignedOk

`func (o *AdvisoryMCna) GetDateAssignedOk() (*string, bool)`

GetDateAssignedOk returns a tuple with the DateAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAssigned

`func (o *AdvisoryMCna) SetDateAssigned(v string)`

SetDateAssigned sets DateAssigned field to given value.

### HasDateAssigned

`func (o *AdvisoryMCna) HasDateAssigned() bool`

HasDateAssigned returns a boolean if a field has been set.

### GetDatePublic

`func (o *AdvisoryMCna) GetDatePublic() string`

GetDatePublic returns the DatePublic field if non-nil, zero value otherwise.

### GetDatePublicOk

`func (o *AdvisoryMCna) GetDatePublicOk() (*string, bool)`

GetDatePublicOk returns a tuple with the DatePublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePublic

`func (o *AdvisoryMCna) SetDatePublic(v string)`

SetDatePublic sets DatePublic field to given value.

### HasDatePublic

`func (o *AdvisoryMCna) HasDatePublic() bool`

HasDatePublic returns a boolean if a field has been set.

### GetDescriptions

`func (o *AdvisoryMCna) GetDescriptions() []AdvisoryMDescriptions`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *AdvisoryMCna) GetDescriptionsOk() (*[]AdvisoryMDescriptions, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *AdvisoryMCna) SetDescriptions(v []AdvisoryMDescriptions)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *AdvisoryMCna) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetExploits

`func (o *AdvisoryMCna) GetExploits() []AdvisoryMDescriptions`

GetExploits returns the Exploits field if non-nil, zero value otherwise.

### GetExploitsOk

`func (o *AdvisoryMCna) GetExploitsOk() (*[]AdvisoryMDescriptions, bool)`

GetExploitsOk returns a tuple with the Exploits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploits

`func (o *AdvisoryMCna) SetExploits(v []AdvisoryMDescriptions)`

SetExploits sets Exploits field to given value.

### HasExploits

`func (o *AdvisoryMCna) HasExploits() bool`

HasExploits returns a boolean if a field has been set.

### GetImpacts

`func (o *AdvisoryMCna) GetImpacts() []AdvisoryImpact`

GetImpacts returns the Impacts field if non-nil, zero value otherwise.

### GetImpactsOk

`func (o *AdvisoryMCna) GetImpactsOk() (*[]AdvisoryImpact, bool)`

GetImpactsOk returns a tuple with the Impacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpacts

`func (o *AdvisoryMCna) SetImpacts(v []AdvisoryImpact)`

SetImpacts sets Impacts field to given value.

### HasImpacts

`func (o *AdvisoryMCna) HasImpacts() bool`

HasImpacts returns a boolean if a field has been set.

### GetMetrics

`func (o *AdvisoryMCna) GetMetrics() []AdvisoryMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *AdvisoryMCna) GetMetricsOk() (*[]AdvisoryMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *AdvisoryMCna) SetMetrics(v []AdvisoryMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *AdvisoryMCna) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetProblemTypes

`func (o *AdvisoryMCna) GetProblemTypes() []AdvisoryMProblemTypes`

GetProblemTypes returns the ProblemTypes field if non-nil, zero value otherwise.

### GetProblemTypesOk

`func (o *AdvisoryMCna) GetProblemTypesOk() (*[]AdvisoryMProblemTypes, bool)`

GetProblemTypesOk returns a tuple with the ProblemTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemTypes

`func (o *AdvisoryMCna) SetProblemTypes(v []AdvisoryMProblemTypes)`

SetProblemTypes sets ProblemTypes field to given value.

### HasProblemTypes

`func (o *AdvisoryMCna) HasProblemTypes() bool`

HasProblemTypes returns a boolean if a field has been set.

### GetProviderMetadata

`func (o *AdvisoryMCna) GetProviderMetadata() AdvisoryMProviderMetadata`

GetProviderMetadata returns the ProviderMetadata field if non-nil, zero value otherwise.

### GetProviderMetadataOk

`func (o *AdvisoryMCna) GetProviderMetadataOk() (*AdvisoryMProviderMetadata, bool)`

GetProviderMetadataOk returns a tuple with the ProviderMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadata

`func (o *AdvisoryMCna) SetProviderMetadata(v AdvisoryMProviderMetadata)`

SetProviderMetadata sets ProviderMetadata field to given value.

### HasProviderMetadata

`func (o *AdvisoryMCna) HasProviderMetadata() bool`

HasProviderMetadata returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryMCna) GetReferences() []AdvisoryMReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryMCna) GetReferencesOk() (*[]AdvisoryMReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryMCna) SetReferences(v []AdvisoryMReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryMCna) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRejectedReasons

`func (o *AdvisoryMCna) GetRejectedReasons() []AdvisoryMDescriptions`

GetRejectedReasons returns the RejectedReasons field if non-nil, zero value otherwise.

### GetRejectedReasonsOk

`func (o *AdvisoryMCna) GetRejectedReasonsOk() (*[]AdvisoryMDescriptions, bool)`

GetRejectedReasonsOk returns a tuple with the RejectedReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedReasons

`func (o *AdvisoryMCna) SetRejectedReasons(v []AdvisoryMDescriptions)`

SetRejectedReasons sets RejectedReasons field to given value.

### HasRejectedReasons

`func (o *AdvisoryMCna) HasRejectedReasons() bool`

HasRejectedReasons returns a boolean if a field has been set.

### GetReplacedBy

`func (o *AdvisoryMCna) GetReplacedBy() []string`

GetReplacedBy returns the ReplacedBy field if non-nil, zero value otherwise.

### GetReplacedByOk

`func (o *AdvisoryMCna) GetReplacedByOk() (*[]string, bool)`

GetReplacedByOk returns a tuple with the ReplacedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedBy

`func (o *AdvisoryMCna) SetReplacedBy(v []string)`

SetReplacedBy sets ReplacedBy field to given value.

### HasReplacedBy

`func (o *AdvisoryMCna) HasReplacedBy() bool`

HasReplacedBy returns a boolean if a field has been set.

### GetSolutions

`func (o *AdvisoryMCna) GetSolutions() []AdvisoryMDescriptions`

GetSolutions returns the Solutions field if non-nil, zero value otherwise.

### GetSolutionsOk

`func (o *AdvisoryMCna) GetSolutionsOk() (*[]AdvisoryMDescriptions, bool)`

GetSolutionsOk returns a tuple with the Solutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutions

`func (o *AdvisoryMCna) SetSolutions(v []AdvisoryMDescriptions)`

SetSolutions sets Solutions field to given value.

### HasSolutions

`func (o *AdvisoryMCna) HasSolutions() bool`

HasSolutions returns a boolean if a field has been set.

### GetSource

`func (o *AdvisoryMCna) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AdvisoryMCna) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AdvisoryMCna) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *AdvisoryMCna) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTags

`func (o *AdvisoryMCna) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AdvisoryMCna) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AdvisoryMCna) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AdvisoryMCna) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTaxonomyMappings

`func (o *AdvisoryMCna) GetTaxonomyMappings() []AdvisoryTaxonomyMapping`

GetTaxonomyMappings returns the TaxonomyMappings field if non-nil, zero value otherwise.

### GetTaxonomyMappingsOk

`func (o *AdvisoryMCna) GetTaxonomyMappingsOk() (*[]AdvisoryTaxonomyMapping, bool)`

GetTaxonomyMappingsOk returns a tuple with the TaxonomyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomyMappings

`func (o *AdvisoryMCna) SetTaxonomyMappings(v []AdvisoryTaxonomyMapping)`

SetTaxonomyMappings sets TaxonomyMappings field to given value.

### HasTaxonomyMappings

`func (o *AdvisoryMCna) HasTaxonomyMappings() bool`

HasTaxonomyMappings returns a boolean if a field has been set.

### GetTimeline

`func (o *AdvisoryMCna) GetTimeline() []AdvisoryTimeline`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *AdvisoryMCna) GetTimelineOk() (*[]AdvisoryTimeline, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *AdvisoryMCna) SetTimeline(v []AdvisoryTimeline)`

SetTimeline sets Timeline field to given value.

### HasTimeline

`func (o *AdvisoryMCna) HasTimeline() bool`

HasTimeline returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryMCna) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryMCna) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryMCna) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryMCna) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWorkarounds

`func (o *AdvisoryMCna) GetWorkarounds() []AdvisoryMDescriptions`

GetWorkarounds returns the Workarounds field if non-nil, zero value otherwise.

### GetWorkaroundsOk

`func (o *AdvisoryMCna) GetWorkaroundsOk() (*[]AdvisoryMDescriptions, bool)`

GetWorkaroundsOk returns a tuple with the Workarounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkarounds

`func (o *AdvisoryMCna) SetWorkarounds(v []AdvisoryMDescriptions)`

SetWorkarounds sets Workarounds field to given value.

### HasWorkarounds

`func (o *AdvisoryMCna) HasWorkarounds() bool`

HasWorkarounds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


