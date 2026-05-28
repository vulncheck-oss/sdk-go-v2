# AdvisoryCirclContainers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adp** | Pointer to [**[]AdvisoryADPContainer**](AdvisoryADPContainer.md) |  | [optional] 
**Cna** | Pointer to [**AdvisoryCirclCna**](AdvisoryCirclCna.md) |  | [optional] 

## Methods

### NewAdvisoryCirclContainers

`func NewAdvisoryCirclContainers() *AdvisoryCirclContainers`

NewAdvisoryCirclContainers instantiates a new AdvisoryCirclContainers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCirclContainersWithDefaults

`func NewAdvisoryCirclContainersWithDefaults() *AdvisoryCirclContainers`

NewAdvisoryCirclContainersWithDefaults instantiates a new AdvisoryCirclContainers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdp

`func (o *AdvisoryCirclContainers) GetAdp() []AdvisoryADPContainer`

GetAdp returns the Adp field if non-nil, zero value otherwise.

### GetAdpOk

`func (o *AdvisoryCirclContainers) GetAdpOk() (*[]AdvisoryADPContainer, bool)`

GetAdpOk returns a tuple with the Adp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdp

`func (o *AdvisoryCirclContainers) SetAdp(v []AdvisoryADPContainer)`

SetAdp sets Adp field to given value.

### HasAdp

`func (o *AdvisoryCirclContainers) HasAdp() bool`

HasAdp returns a boolean if a field has been set.

### GetCna

`func (o *AdvisoryCirclContainers) GetCna() AdvisoryCirclCna`

GetCna returns the Cna field if non-nil, zero value otherwise.

### GetCnaOk

`func (o *AdvisoryCirclContainers) GetCnaOk() (*AdvisoryCirclCna, bool)`

GetCnaOk returns a tuple with the Cna field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCna

`func (o *AdvisoryCirclContainers) SetCna(v AdvisoryCirclCna)`

SetCna sets Cna field to given value.

### HasCna

`func (o *AdvisoryCirclContainers) HasCna() bool`

HasCna returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


