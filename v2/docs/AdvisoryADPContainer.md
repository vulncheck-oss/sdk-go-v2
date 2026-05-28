# AdvisoryADPContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to [**[]AdvisoryMAffected**](AdvisoryMAffected.md) |  | [optional] 
**Configurations** | Pointer to [**[]AdvisoryMDescriptions**](AdvisoryMDescriptions.md) |  | [optional] 
**CpeApplicability** | Pointer to [**[]AdvisoryCustomCPE**](AdvisoryCustomCPE.md) |  | [optional] 
**Credits** | Pointer to [**[]AdvisoryCredit**](AdvisoryCredit.md) |  | [optional] 
**DatePublic** | Pointer to **string** |  | [optional] 
**Descriptions** | Pointer to [**[]AdvisoryMDescriptions**](AdvisoryMDescriptions.md) |  | [optional] 
**Exploits** | Pointer to [**[]AdvisoryMDescriptions**](AdvisoryMDescriptions.md) |  | [optional] 
**Impacts** | Pointer to [**[]AdvisoryImpact**](AdvisoryImpact.md) |  | [optional] 
**Metrics** | Pointer to [**[]AdvisoryMetric**](AdvisoryMetric.md) |  | [optional] 
**ProblemTypes** | Pointer to [**[]AdvisoryMProblemTypes**](AdvisoryMProblemTypes.md) |  | [optional] 
**ProviderMetadata** | Pointer to [**AdvisoryMProviderMetadata**](AdvisoryMProviderMetadata.md) |  | [optional] 
**References** | Pointer to [**[]AdvisoryMReference**](AdvisoryMReference.md) |  | [optional] 
**Solutions** | Pointer to [**[]AdvisoryMDescriptions**](AdvisoryMDescriptions.md) |  | [optional] 
**Source** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TaxonomyMappings** | Pointer to [**[]AdvisoryTaxonomyMapping**](AdvisoryTaxonomyMapping.md) |  | [optional] 
**Timeline** | Pointer to [**[]AdvisoryTimeline**](AdvisoryTimeline.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Workarounds** | Pointer to [**[]AdvisoryMDescriptions**](AdvisoryMDescriptions.md) |  | [optional] 

## Methods

### NewAdvisoryADPContainer

`func NewAdvisoryADPContainer() *AdvisoryADPContainer`

NewAdvisoryADPContainer instantiates a new AdvisoryADPContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryADPContainerWithDefaults

`func NewAdvisoryADPContainerWithDefaults() *AdvisoryADPContainer`

NewAdvisoryADPContainerWithDefaults instantiates a new AdvisoryADPContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *AdvisoryADPContainer) GetAffected() []AdvisoryMAffected`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryADPContainer) GetAffectedOk() (*[]AdvisoryMAffected, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryADPContainer) SetAffected(v []AdvisoryMAffected)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryADPContainer) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetConfigurations

`func (o *AdvisoryADPContainer) GetConfigurations() []AdvisoryMDescriptions`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *AdvisoryADPContainer) GetConfigurationsOk() (*[]AdvisoryMDescriptions, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *AdvisoryADPContainer) SetConfigurations(v []AdvisoryMDescriptions)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *AdvisoryADPContainer) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetCpeApplicability

`func (o *AdvisoryADPContainer) GetCpeApplicability() []AdvisoryCustomCPE`

GetCpeApplicability returns the CpeApplicability field if non-nil, zero value otherwise.

### GetCpeApplicabilityOk

`func (o *AdvisoryADPContainer) GetCpeApplicabilityOk() (*[]AdvisoryCustomCPE, bool)`

GetCpeApplicabilityOk returns a tuple with the CpeApplicability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeApplicability

`func (o *AdvisoryADPContainer) SetCpeApplicability(v []AdvisoryCustomCPE)`

SetCpeApplicability sets CpeApplicability field to given value.

### HasCpeApplicability

`func (o *AdvisoryADPContainer) HasCpeApplicability() bool`

HasCpeApplicability returns a boolean if a field has been set.

### GetCredits

`func (o *AdvisoryADPContainer) GetCredits() []AdvisoryCredit`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *AdvisoryADPContainer) GetCreditsOk() (*[]AdvisoryCredit, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *AdvisoryADPContainer) SetCredits(v []AdvisoryCredit)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *AdvisoryADPContainer) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetDatePublic

`func (o *AdvisoryADPContainer) GetDatePublic() string`

GetDatePublic returns the DatePublic field if non-nil, zero value otherwise.

### GetDatePublicOk

`func (o *AdvisoryADPContainer) GetDatePublicOk() (*string, bool)`

GetDatePublicOk returns a tuple with the DatePublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePublic

`func (o *AdvisoryADPContainer) SetDatePublic(v string)`

SetDatePublic sets DatePublic field to given value.

### HasDatePublic

`func (o *AdvisoryADPContainer) HasDatePublic() bool`

HasDatePublic returns a boolean if a field has been set.

### GetDescriptions

`func (o *AdvisoryADPContainer) GetDescriptions() []AdvisoryMDescriptions`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *AdvisoryADPContainer) GetDescriptionsOk() (*[]AdvisoryMDescriptions, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *AdvisoryADPContainer) SetDescriptions(v []AdvisoryMDescriptions)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *AdvisoryADPContainer) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetExploits

`func (o *AdvisoryADPContainer) GetExploits() []AdvisoryMDescriptions`

GetExploits returns the Exploits field if non-nil, zero value otherwise.

### GetExploitsOk

`func (o *AdvisoryADPContainer) GetExploitsOk() (*[]AdvisoryMDescriptions, bool)`

GetExploitsOk returns a tuple with the Exploits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploits

`func (o *AdvisoryADPContainer) SetExploits(v []AdvisoryMDescriptions)`

SetExploits sets Exploits field to given value.

### HasExploits

`func (o *AdvisoryADPContainer) HasExploits() bool`

HasExploits returns a boolean if a field has been set.

### GetImpacts

`func (o *AdvisoryADPContainer) GetImpacts() []AdvisoryImpact`

GetImpacts returns the Impacts field if non-nil, zero value otherwise.

### GetImpactsOk

`func (o *AdvisoryADPContainer) GetImpactsOk() (*[]AdvisoryImpact, bool)`

GetImpactsOk returns a tuple with the Impacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpacts

`func (o *AdvisoryADPContainer) SetImpacts(v []AdvisoryImpact)`

SetImpacts sets Impacts field to given value.

### HasImpacts

`func (o *AdvisoryADPContainer) HasImpacts() bool`

HasImpacts returns a boolean if a field has been set.

### GetMetrics

`func (o *AdvisoryADPContainer) GetMetrics() []AdvisoryMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *AdvisoryADPContainer) GetMetricsOk() (*[]AdvisoryMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *AdvisoryADPContainer) SetMetrics(v []AdvisoryMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *AdvisoryADPContainer) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetProblemTypes

`func (o *AdvisoryADPContainer) GetProblemTypes() []AdvisoryMProblemTypes`

GetProblemTypes returns the ProblemTypes field if non-nil, zero value otherwise.

### GetProblemTypesOk

`func (o *AdvisoryADPContainer) GetProblemTypesOk() (*[]AdvisoryMProblemTypes, bool)`

GetProblemTypesOk returns a tuple with the ProblemTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemTypes

`func (o *AdvisoryADPContainer) SetProblemTypes(v []AdvisoryMProblemTypes)`

SetProblemTypes sets ProblemTypes field to given value.

### HasProblemTypes

`func (o *AdvisoryADPContainer) HasProblemTypes() bool`

HasProblemTypes returns a boolean if a field has been set.

### GetProviderMetadata

`func (o *AdvisoryADPContainer) GetProviderMetadata() AdvisoryMProviderMetadata`

GetProviderMetadata returns the ProviderMetadata field if non-nil, zero value otherwise.

### GetProviderMetadataOk

`func (o *AdvisoryADPContainer) GetProviderMetadataOk() (*AdvisoryMProviderMetadata, bool)`

GetProviderMetadataOk returns a tuple with the ProviderMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadata

`func (o *AdvisoryADPContainer) SetProviderMetadata(v AdvisoryMProviderMetadata)`

SetProviderMetadata sets ProviderMetadata field to given value.

### HasProviderMetadata

`func (o *AdvisoryADPContainer) HasProviderMetadata() bool`

HasProviderMetadata returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryADPContainer) GetReferences() []AdvisoryMReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryADPContainer) GetReferencesOk() (*[]AdvisoryMReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryADPContainer) SetReferences(v []AdvisoryMReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryADPContainer) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSolutions

`func (o *AdvisoryADPContainer) GetSolutions() []AdvisoryMDescriptions`

GetSolutions returns the Solutions field if non-nil, zero value otherwise.

### GetSolutionsOk

`func (o *AdvisoryADPContainer) GetSolutionsOk() (*[]AdvisoryMDescriptions, bool)`

GetSolutionsOk returns a tuple with the Solutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutions

`func (o *AdvisoryADPContainer) SetSolutions(v []AdvisoryMDescriptions)`

SetSolutions sets Solutions field to given value.

### HasSolutions

`func (o *AdvisoryADPContainer) HasSolutions() bool`

HasSolutions returns a boolean if a field has been set.

### GetSource

`func (o *AdvisoryADPContainer) GetSource() []int32`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AdvisoryADPContainer) GetSourceOk() (*[]int32, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AdvisoryADPContainer) SetSource(v []int32)`

SetSource sets Source field to given value.

### HasSource

`func (o *AdvisoryADPContainer) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTags

`func (o *AdvisoryADPContainer) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AdvisoryADPContainer) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AdvisoryADPContainer) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AdvisoryADPContainer) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTaxonomyMappings

`func (o *AdvisoryADPContainer) GetTaxonomyMappings() []AdvisoryTaxonomyMapping`

GetTaxonomyMappings returns the TaxonomyMappings field if non-nil, zero value otherwise.

### GetTaxonomyMappingsOk

`func (o *AdvisoryADPContainer) GetTaxonomyMappingsOk() (*[]AdvisoryTaxonomyMapping, bool)`

GetTaxonomyMappingsOk returns a tuple with the TaxonomyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomyMappings

`func (o *AdvisoryADPContainer) SetTaxonomyMappings(v []AdvisoryTaxonomyMapping)`

SetTaxonomyMappings sets TaxonomyMappings field to given value.

### HasTaxonomyMappings

`func (o *AdvisoryADPContainer) HasTaxonomyMappings() bool`

HasTaxonomyMappings returns a boolean if a field has been set.

### GetTimeline

`func (o *AdvisoryADPContainer) GetTimeline() []AdvisoryTimeline`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *AdvisoryADPContainer) GetTimelineOk() (*[]AdvisoryTimeline, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *AdvisoryADPContainer) SetTimeline(v []AdvisoryTimeline)`

SetTimeline sets Timeline field to given value.

### HasTimeline

`func (o *AdvisoryADPContainer) HasTimeline() bool`

HasTimeline returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryADPContainer) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryADPContainer) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryADPContainer) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryADPContainer) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWorkarounds

`func (o *AdvisoryADPContainer) GetWorkarounds() []AdvisoryMDescriptions`

GetWorkarounds returns the Workarounds field if non-nil, zero value otherwise.

### GetWorkaroundsOk

`func (o *AdvisoryADPContainer) GetWorkaroundsOk() (*[]AdvisoryMDescriptions, bool)`

GetWorkaroundsOk returns a tuple with the Workarounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkarounds

`func (o *AdvisoryADPContainer) SetWorkarounds(v []AdvisoryMDescriptions)`

SetWorkarounds sets Workarounds field to given value.

### HasWorkarounds

`func (o *AdvisoryADPContainer) HasWorkarounds() bool`

HasWorkarounds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


