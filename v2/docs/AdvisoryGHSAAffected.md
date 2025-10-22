# AdvisoryGHSAAffected

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EcosystemSpecific** | Pointer to [**AdvisoryGHSAEcoSystemSpecific**](AdvisoryGHSAEcoSystemSpecific.md) |  | [optional] 
**Package** | Pointer to [**AdvisoryGHSAPackage**](AdvisoryGHSAPackage.md) |  | [optional] 
**Ranges** | Pointer to [**[]AdvisoryGHSARange**](AdvisoryGHSARange.md) |  | [optional] 

## Methods

### NewAdvisoryGHSAAffected

`func NewAdvisoryGHSAAffected() *AdvisoryGHSAAffected`

NewAdvisoryGHSAAffected instantiates a new AdvisoryGHSAAffected object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryGHSAAffectedWithDefaults

`func NewAdvisoryGHSAAffectedWithDefaults() *AdvisoryGHSAAffected`

NewAdvisoryGHSAAffectedWithDefaults instantiates a new AdvisoryGHSAAffected object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEcosystemSpecific

`func (o *AdvisoryGHSAAffected) GetEcosystemSpecific() AdvisoryGHSAEcoSystemSpecific`

GetEcosystemSpecific returns the EcosystemSpecific field if non-nil, zero value otherwise.

### GetEcosystemSpecificOk

`func (o *AdvisoryGHSAAffected) GetEcosystemSpecificOk() (*AdvisoryGHSAEcoSystemSpecific, bool)`

GetEcosystemSpecificOk returns a tuple with the EcosystemSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcosystemSpecific

`func (o *AdvisoryGHSAAffected) SetEcosystemSpecific(v AdvisoryGHSAEcoSystemSpecific)`

SetEcosystemSpecific sets EcosystemSpecific field to given value.

### HasEcosystemSpecific

`func (o *AdvisoryGHSAAffected) HasEcosystemSpecific() bool`

HasEcosystemSpecific returns a boolean if a field has been set.

### GetPackage

`func (o *AdvisoryGHSAAffected) GetPackage() AdvisoryGHSAPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *AdvisoryGHSAAffected) GetPackageOk() (*AdvisoryGHSAPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *AdvisoryGHSAAffected) SetPackage(v AdvisoryGHSAPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *AdvisoryGHSAAffected) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetRanges

`func (o *AdvisoryGHSAAffected) GetRanges() []AdvisoryGHSARange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *AdvisoryGHSAAffected) GetRangesOk() (*[]AdvisoryGHSARange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *AdvisoryGHSAAffected) SetRanges(v []AdvisoryGHSARange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *AdvisoryGHSAAffected) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


