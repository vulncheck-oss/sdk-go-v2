# AdvisoryADP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to [**[]AdvisoryMAffected**](AdvisoryMAffected.md) |  | [optional] 
**Metrics** | Pointer to [**[]AdvisoryVulnrichmentMetric**](AdvisoryVulnrichmentMetric.md) |  | [optional] 
**ProviderMetadata** | Pointer to [**AdvisoryMProviderMetadata**](AdvisoryMProviderMetadata.md) |  | [optional] 

## Methods

### NewAdvisoryADP

`func NewAdvisoryADP() *AdvisoryADP`

NewAdvisoryADP instantiates a new AdvisoryADP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryADPWithDefaults

`func NewAdvisoryADPWithDefaults() *AdvisoryADP`

NewAdvisoryADPWithDefaults instantiates a new AdvisoryADP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *AdvisoryADP) GetAffected() []AdvisoryMAffected`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryADP) GetAffectedOk() (*[]AdvisoryMAffected, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryADP) SetAffected(v []AdvisoryMAffected)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryADP) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetMetrics

`func (o *AdvisoryADP) GetMetrics() []AdvisoryVulnrichmentMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *AdvisoryADP) GetMetricsOk() (*[]AdvisoryVulnrichmentMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *AdvisoryADP) SetMetrics(v []AdvisoryVulnrichmentMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *AdvisoryADP) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetProviderMetadata

`func (o *AdvisoryADP) GetProviderMetadata() AdvisoryMProviderMetadata`

GetProviderMetadata returns the ProviderMetadata field if non-nil, zero value otherwise.

### GetProviderMetadataOk

`func (o *AdvisoryADP) GetProviderMetadataOk() (*AdvisoryMProviderMetadata, bool)`

GetProviderMetadataOk returns a tuple with the ProviderMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadata

`func (o *AdvisoryADP) SetProviderMetadata(v AdvisoryMProviderMetadata)`

SetProviderMetadata sets ProviderMetadata field to given value.

### HasProviderMetadata

`func (o *AdvisoryADP) HasProviderMetadata() bool`

HasProviderMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


