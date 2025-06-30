# AdvisoryVulnrichmentContainers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adp** | Pointer to [**[]AdvisoryADP**](AdvisoryADP.md) |  | [optional] 
**Cna** | Pointer to [**AdvisoryMCna**](AdvisoryMCna.md) |  | [optional] 

## Methods

### NewAdvisoryVulnrichmentContainers

`func NewAdvisoryVulnrichmentContainers() *AdvisoryVulnrichmentContainers`

NewAdvisoryVulnrichmentContainers instantiates a new AdvisoryVulnrichmentContainers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryVulnrichmentContainersWithDefaults

`func NewAdvisoryVulnrichmentContainersWithDefaults() *AdvisoryVulnrichmentContainers`

NewAdvisoryVulnrichmentContainersWithDefaults instantiates a new AdvisoryVulnrichmentContainers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdp

`func (o *AdvisoryVulnrichmentContainers) GetAdp() []AdvisoryADP`

GetAdp returns the Adp field if non-nil, zero value otherwise.

### GetAdpOk

`func (o *AdvisoryVulnrichmentContainers) GetAdpOk() (*[]AdvisoryADP, bool)`

GetAdpOk returns a tuple with the Adp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdp

`func (o *AdvisoryVulnrichmentContainers) SetAdp(v []AdvisoryADP)`

SetAdp sets Adp field to given value.

### HasAdp

`func (o *AdvisoryVulnrichmentContainers) HasAdp() bool`

HasAdp returns a boolean if a field has been set.

### GetCna

`func (o *AdvisoryVulnrichmentContainers) GetCna() AdvisoryMCna`

GetCna returns the Cna field if non-nil, zero value otherwise.

### GetCnaOk

`func (o *AdvisoryVulnrichmentContainers) GetCnaOk() (*AdvisoryMCna, bool)`

GetCnaOk returns a tuple with the Cna field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCna

`func (o *AdvisoryVulnrichmentContainers) SetCna(v AdvisoryMCna)`

SetCna sets Cna field to given value.

### HasCna

`func (o *AdvisoryVulnrichmentContainers) HasCna() bool`

HasCna returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


