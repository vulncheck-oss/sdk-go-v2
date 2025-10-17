# AdvisoryADPContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to [**[]AdvisoryMAffected**](AdvisoryMAffected.md) |  | [optional] 
**DatePublic** | Pointer to **string** | OK | [optional] 
**Descriptions** | Pointer to [**[]AdvisoryMDescriptions**](AdvisoryMDescriptions.md) | OK | [optional] 
**Impacts** | Pointer to [**[]AdvisoryImpact**](AdvisoryImpact.md) | OK | [optional] 
**Metrics** | Pointer to [**[]AdvisoryMetric**](AdvisoryMetric.md) | OK | [optional] 
**ProblemTypes** | Pointer to [**[]AdvisoryMProblemTypes**](AdvisoryMProblemTypes.md) | OK | [optional] 
**ProviderMetadata** | Pointer to [**AdvisoryMProviderMetadata**](AdvisoryMProviderMetadata.md) | OK | [optional] 
**References** | Pointer to [**[]AdvisoryMReference**](AdvisoryMReference.md) |  | [optional] 
**Tags** | Pointer to **[]string** | OK | [optional] 
**Title** | Pointer to **string** | OK | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


