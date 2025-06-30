# AdvisoryImpact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapecId** | Pointer to **string** |  | [optional] 
**Descriptions** | Pointer to [**[]AdvisoryMDescriptions**](AdvisoryMDescriptions.md) |  | [optional] 

## Methods

### NewAdvisoryImpact

`func NewAdvisoryImpact() *AdvisoryImpact`

NewAdvisoryImpact instantiates a new AdvisoryImpact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryImpactWithDefaults

`func NewAdvisoryImpactWithDefaults() *AdvisoryImpact`

NewAdvisoryImpactWithDefaults instantiates a new AdvisoryImpact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapecId

`func (o *AdvisoryImpact) GetCapecId() string`

GetCapecId returns the CapecId field if non-nil, zero value otherwise.

### GetCapecIdOk

`func (o *AdvisoryImpact) GetCapecIdOk() (*string, bool)`

GetCapecIdOk returns a tuple with the CapecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapecId

`func (o *AdvisoryImpact) SetCapecId(v string)`

SetCapecId sets CapecId field to given value.

### HasCapecId

`func (o *AdvisoryImpact) HasCapecId() bool`

HasCapecId returns a boolean if a field has been set.

### GetDescriptions

`func (o *AdvisoryImpact) GetDescriptions() []AdvisoryMDescriptions`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *AdvisoryImpact) GetDescriptionsOk() (*[]AdvisoryMDescriptions, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *AdvisoryImpact) SetDescriptions(v []AdvisoryMDescriptions)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *AdvisoryImpact) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


