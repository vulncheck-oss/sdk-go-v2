# AdvisoryMCna

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to [**[]AdvisoryMAffected**](AdvisoryMAffected.md) |  | [optional] 
**CpeApplicability** | Pointer to [**[]AdvisoryMCPEApplicability**](AdvisoryMCPEApplicability.md) |  | [optional] 
**Credits** | Pointer to [**[]AdvisoryCredit**](AdvisoryCredit.md) |  | [optional] 
**Descriptions** | Pointer to [**[]AdvisoryMDescriptions**](AdvisoryMDescriptions.md) |  | [optional] 
**Impacts** | Pointer to [**[]AdvisoryImpact**](AdvisoryImpact.md) |  | [optional] 
**Metrics** | Pointer to [**[]AdvisoryMetric**](AdvisoryMetric.md) |  | [optional] 
**ProblemTypes** | Pointer to [**[]AdvisoryMProblemTypes**](AdvisoryMProblemTypes.md) |  | [optional] 
**ProviderMetadata** | Pointer to [**AdvisoryMProviderMetadata**](AdvisoryMProviderMetadata.md) |  | [optional] 
**References** | Pointer to [**[]AdvisoryMReference**](AdvisoryMReference.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Timeline** | Pointer to [**[]AdvisoryTimeline**](AdvisoryTimeline.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

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

### GetCpeApplicability

`func (o *AdvisoryMCna) GetCpeApplicability() []AdvisoryMCPEApplicability`

GetCpeApplicability returns the CpeApplicability field if non-nil, zero value otherwise.

### GetCpeApplicabilityOk

`func (o *AdvisoryMCna) GetCpeApplicabilityOk() (*[]AdvisoryMCPEApplicability, bool)`

GetCpeApplicabilityOk returns a tuple with the CpeApplicability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeApplicability

`func (o *AdvisoryMCna) SetCpeApplicability(v []AdvisoryMCPEApplicability)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


