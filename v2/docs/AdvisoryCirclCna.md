# AdvisoryCirclCna

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
**Source** | Pointer to [**AdvisoryCirclSource**](AdvisoryCirclSource.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TaxonomyMappings** | Pointer to [**[]AdvisoryTaxonomyMapping**](AdvisoryTaxonomyMapping.md) |  | [optional] 
**Timeline** | Pointer to [**[]AdvisoryTimeline**](AdvisoryTimeline.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Workarounds** | Pointer to [**[]AdvisoryMDescriptions**](AdvisoryMDescriptions.md) |  | [optional] 
**XGcve** | Pointer to [**[]AdvisoryCirclXGcve**](AdvisoryCirclXGcve.md) |  | [optional] 
**XGenerator** | Pointer to [**AdvisoryCirclXGenerator**](AdvisoryCirclXGenerator.md) |  | [optional] 

## Methods

### NewAdvisoryCirclCna

`func NewAdvisoryCirclCna() *AdvisoryCirclCna`

NewAdvisoryCirclCna instantiates a new AdvisoryCirclCna object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCirclCnaWithDefaults

`func NewAdvisoryCirclCnaWithDefaults() *AdvisoryCirclCna`

NewAdvisoryCirclCnaWithDefaults instantiates a new AdvisoryCirclCna object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *AdvisoryCirclCna) GetAffected() []AdvisoryMAffected`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryCirclCna) GetAffectedOk() (*[]AdvisoryMAffected, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryCirclCna) SetAffected(v []AdvisoryMAffected)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryCirclCna) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetConfigurations

`func (o *AdvisoryCirclCna) GetConfigurations() []AdvisoryMDescriptions`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *AdvisoryCirclCna) GetConfigurationsOk() (*[]AdvisoryMDescriptions, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *AdvisoryCirclCna) SetConfigurations(v []AdvisoryMDescriptions)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *AdvisoryCirclCna) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetCpeApplicability

`func (o *AdvisoryCirclCna) GetCpeApplicability() []AdvisoryCustomCPE`

GetCpeApplicability returns the CpeApplicability field if non-nil, zero value otherwise.

### GetCpeApplicabilityOk

`func (o *AdvisoryCirclCna) GetCpeApplicabilityOk() (*[]AdvisoryCustomCPE, bool)`

GetCpeApplicabilityOk returns a tuple with the CpeApplicability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeApplicability

`func (o *AdvisoryCirclCna) SetCpeApplicability(v []AdvisoryCustomCPE)`

SetCpeApplicability sets CpeApplicability field to given value.

### HasCpeApplicability

`func (o *AdvisoryCirclCna) HasCpeApplicability() bool`

HasCpeApplicability returns a boolean if a field has been set.

### GetCredits

`func (o *AdvisoryCirclCna) GetCredits() []AdvisoryCredit`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *AdvisoryCirclCna) GetCreditsOk() (*[]AdvisoryCredit, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *AdvisoryCirclCna) SetCredits(v []AdvisoryCredit)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *AdvisoryCirclCna) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetDateAssigned

`func (o *AdvisoryCirclCna) GetDateAssigned() string`

GetDateAssigned returns the DateAssigned field if non-nil, zero value otherwise.

### GetDateAssignedOk

`func (o *AdvisoryCirclCna) GetDateAssignedOk() (*string, bool)`

GetDateAssignedOk returns a tuple with the DateAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAssigned

`func (o *AdvisoryCirclCna) SetDateAssigned(v string)`

SetDateAssigned sets DateAssigned field to given value.

### HasDateAssigned

`func (o *AdvisoryCirclCna) HasDateAssigned() bool`

HasDateAssigned returns a boolean if a field has been set.

### GetDatePublic

`func (o *AdvisoryCirclCna) GetDatePublic() string`

GetDatePublic returns the DatePublic field if non-nil, zero value otherwise.

### GetDatePublicOk

`func (o *AdvisoryCirclCna) GetDatePublicOk() (*string, bool)`

GetDatePublicOk returns a tuple with the DatePublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePublic

`func (o *AdvisoryCirclCna) SetDatePublic(v string)`

SetDatePublic sets DatePublic field to given value.

### HasDatePublic

`func (o *AdvisoryCirclCna) HasDatePublic() bool`

HasDatePublic returns a boolean if a field has been set.

### GetDescriptions

`func (o *AdvisoryCirclCna) GetDescriptions() []AdvisoryMDescriptions`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *AdvisoryCirclCna) GetDescriptionsOk() (*[]AdvisoryMDescriptions, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *AdvisoryCirclCna) SetDescriptions(v []AdvisoryMDescriptions)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *AdvisoryCirclCna) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetExploits

`func (o *AdvisoryCirclCna) GetExploits() []AdvisoryMDescriptions`

GetExploits returns the Exploits field if non-nil, zero value otherwise.

### GetExploitsOk

`func (o *AdvisoryCirclCna) GetExploitsOk() (*[]AdvisoryMDescriptions, bool)`

GetExploitsOk returns a tuple with the Exploits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploits

`func (o *AdvisoryCirclCna) SetExploits(v []AdvisoryMDescriptions)`

SetExploits sets Exploits field to given value.

### HasExploits

`func (o *AdvisoryCirclCna) HasExploits() bool`

HasExploits returns a boolean if a field has been set.

### GetImpacts

`func (o *AdvisoryCirclCna) GetImpacts() []AdvisoryImpact`

GetImpacts returns the Impacts field if non-nil, zero value otherwise.

### GetImpactsOk

`func (o *AdvisoryCirclCna) GetImpactsOk() (*[]AdvisoryImpact, bool)`

GetImpactsOk returns a tuple with the Impacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpacts

`func (o *AdvisoryCirclCna) SetImpacts(v []AdvisoryImpact)`

SetImpacts sets Impacts field to given value.

### HasImpacts

`func (o *AdvisoryCirclCna) HasImpacts() bool`

HasImpacts returns a boolean if a field has been set.

### GetMetrics

`func (o *AdvisoryCirclCna) GetMetrics() []AdvisoryMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *AdvisoryCirclCna) GetMetricsOk() (*[]AdvisoryMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *AdvisoryCirclCna) SetMetrics(v []AdvisoryMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *AdvisoryCirclCna) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetProblemTypes

`func (o *AdvisoryCirclCna) GetProblemTypes() []AdvisoryMProblemTypes`

GetProblemTypes returns the ProblemTypes field if non-nil, zero value otherwise.

### GetProblemTypesOk

`func (o *AdvisoryCirclCna) GetProblemTypesOk() (*[]AdvisoryMProblemTypes, bool)`

GetProblemTypesOk returns a tuple with the ProblemTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemTypes

`func (o *AdvisoryCirclCna) SetProblemTypes(v []AdvisoryMProblemTypes)`

SetProblemTypes sets ProblemTypes field to given value.

### HasProblemTypes

`func (o *AdvisoryCirclCna) HasProblemTypes() bool`

HasProblemTypes returns a boolean if a field has been set.

### GetProviderMetadata

`func (o *AdvisoryCirclCna) GetProviderMetadata() AdvisoryMProviderMetadata`

GetProviderMetadata returns the ProviderMetadata field if non-nil, zero value otherwise.

### GetProviderMetadataOk

`func (o *AdvisoryCirclCna) GetProviderMetadataOk() (*AdvisoryMProviderMetadata, bool)`

GetProviderMetadataOk returns a tuple with the ProviderMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadata

`func (o *AdvisoryCirclCna) SetProviderMetadata(v AdvisoryMProviderMetadata)`

SetProviderMetadata sets ProviderMetadata field to given value.

### HasProviderMetadata

`func (o *AdvisoryCirclCna) HasProviderMetadata() bool`

HasProviderMetadata returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryCirclCna) GetReferences() []AdvisoryMReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryCirclCna) GetReferencesOk() (*[]AdvisoryMReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryCirclCna) SetReferences(v []AdvisoryMReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryCirclCna) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRejectedReasons

`func (o *AdvisoryCirclCna) GetRejectedReasons() []AdvisoryMDescriptions`

GetRejectedReasons returns the RejectedReasons field if non-nil, zero value otherwise.

### GetRejectedReasonsOk

`func (o *AdvisoryCirclCna) GetRejectedReasonsOk() (*[]AdvisoryMDescriptions, bool)`

GetRejectedReasonsOk returns a tuple with the RejectedReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedReasons

`func (o *AdvisoryCirclCna) SetRejectedReasons(v []AdvisoryMDescriptions)`

SetRejectedReasons sets RejectedReasons field to given value.

### HasRejectedReasons

`func (o *AdvisoryCirclCna) HasRejectedReasons() bool`

HasRejectedReasons returns a boolean if a field has been set.

### GetReplacedBy

`func (o *AdvisoryCirclCna) GetReplacedBy() []string`

GetReplacedBy returns the ReplacedBy field if non-nil, zero value otherwise.

### GetReplacedByOk

`func (o *AdvisoryCirclCna) GetReplacedByOk() (*[]string, bool)`

GetReplacedByOk returns a tuple with the ReplacedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedBy

`func (o *AdvisoryCirclCna) SetReplacedBy(v []string)`

SetReplacedBy sets ReplacedBy field to given value.

### HasReplacedBy

`func (o *AdvisoryCirclCna) HasReplacedBy() bool`

HasReplacedBy returns a boolean if a field has been set.

### GetSolutions

`func (o *AdvisoryCirclCna) GetSolutions() []AdvisoryMDescriptions`

GetSolutions returns the Solutions field if non-nil, zero value otherwise.

### GetSolutionsOk

`func (o *AdvisoryCirclCna) GetSolutionsOk() (*[]AdvisoryMDescriptions, bool)`

GetSolutionsOk returns a tuple with the Solutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutions

`func (o *AdvisoryCirclCna) SetSolutions(v []AdvisoryMDescriptions)`

SetSolutions sets Solutions field to given value.

### HasSolutions

`func (o *AdvisoryCirclCna) HasSolutions() bool`

HasSolutions returns a boolean if a field has been set.

### GetSource

`func (o *AdvisoryCirclCna) GetSource() AdvisoryCirclSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AdvisoryCirclCna) GetSourceOk() (*AdvisoryCirclSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AdvisoryCirclCna) SetSource(v AdvisoryCirclSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *AdvisoryCirclCna) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTags

`func (o *AdvisoryCirclCna) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AdvisoryCirclCna) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AdvisoryCirclCna) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AdvisoryCirclCna) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTaxonomyMappings

`func (o *AdvisoryCirclCna) GetTaxonomyMappings() []AdvisoryTaxonomyMapping`

GetTaxonomyMappings returns the TaxonomyMappings field if non-nil, zero value otherwise.

### GetTaxonomyMappingsOk

`func (o *AdvisoryCirclCna) GetTaxonomyMappingsOk() (*[]AdvisoryTaxonomyMapping, bool)`

GetTaxonomyMappingsOk returns a tuple with the TaxonomyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomyMappings

`func (o *AdvisoryCirclCna) SetTaxonomyMappings(v []AdvisoryTaxonomyMapping)`

SetTaxonomyMappings sets TaxonomyMappings field to given value.

### HasTaxonomyMappings

`func (o *AdvisoryCirclCna) HasTaxonomyMappings() bool`

HasTaxonomyMappings returns a boolean if a field has been set.

### GetTimeline

`func (o *AdvisoryCirclCna) GetTimeline() []AdvisoryTimeline`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *AdvisoryCirclCna) GetTimelineOk() (*[]AdvisoryTimeline, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *AdvisoryCirclCna) SetTimeline(v []AdvisoryTimeline)`

SetTimeline sets Timeline field to given value.

### HasTimeline

`func (o *AdvisoryCirclCna) HasTimeline() bool`

HasTimeline returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryCirclCna) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryCirclCna) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryCirclCna) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryCirclCna) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWorkarounds

`func (o *AdvisoryCirclCna) GetWorkarounds() []AdvisoryMDescriptions`

GetWorkarounds returns the Workarounds field if non-nil, zero value otherwise.

### GetWorkaroundsOk

`func (o *AdvisoryCirclCna) GetWorkaroundsOk() (*[]AdvisoryMDescriptions, bool)`

GetWorkaroundsOk returns a tuple with the Workarounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkarounds

`func (o *AdvisoryCirclCna) SetWorkarounds(v []AdvisoryMDescriptions)`

SetWorkarounds sets Workarounds field to given value.

### HasWorkarounds

`func (o *AdvisoryCirclCna) HasWorkarounds() bool`

HasWorkarounds returns a boolean if a field has been set.

### GetXGcve

`func (o *AdvisoryCirclCna) GetXGcve() []AdvisoryCirclXGcve`

GetXGcve returns the XGcve field if non-nil, zero value otherwise.

### GetXGcveOk

`func (o *AdvisoryCirclCna) GetXGcveOk() (*[]AdvisoryCirclXGcve, bool)`

GetXGcveOk returns a tuple with the XGcve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXGcve

`func (o *AdvisoryCirclCna) SetXGcve(v []AdvisoryCirclXGcve)`

SetXGcve sets XGcve field to given value.

### HasXGcve

`func (o *AdvisoryCirclCna) HasXGcve() bool`

HasXGcve returns a boolean if a field has been set.

### GetXGenerator

`func (o *AdvisoryCirclCna) GetXGenerator() AdvisoryCirclXGenerator`

GetXGenerator returns the XGenerator field if non-nil, zero value otherwise.

### GetXGeneratorOk

`func (o *AdvisoryCirclCna) GetXGeneratorOk() (*AdvisoryCirclXGenerator, bool)`

GetXGeneratorOk returns a tuple with the XGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXGenerator

`func (o *AdvisoryCirclCna) SetXGenerator(v AdvisoryCirclXGenerator)`

SetXGenerator sets XGenerator field to given value.

### HasXGenerator

`func (o *AdvisoryCirclCna) HasXGenerator() bool`

HasXGenerator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


