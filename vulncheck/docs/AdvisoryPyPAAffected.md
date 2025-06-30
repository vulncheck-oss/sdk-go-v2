# AdvisoryPyPAAffected

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Package** | Pointer to [**AdvisoryPyPAPackage**](AdvisoryPyPAPackage.md) |  | [optional] 
**Ranges** | Pointer to [**[]AdvisoryPyPARange**](AdvisoryPyPARange.md) |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdvisoryPyPAAffected

`func NewAdvisoryPyPAAffected() *AdvisoryPyPAAffected`

NewAdvisoryPyPAAffected instantiates a new AdvisoryPyPAAffected object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryPyPAAffectedWithDefaults

`func NewAdvisoryPyPAAffectedWithDefaults() *AdvisoryPyPAAffected`

NewAdvisoryPyPAAffectedWithDefaults instantiates a new AdvisoryPyPAAffected object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackage

`func (o *AdvisoryPyPAAffected) GetPackage() AdvisoryPyPAPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *AdvisoryPyPAAffected) GetPackageOk() (*AdvisoryPyPAPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *AdvisoryPyPAAffected) SetPackage(v AdvisoryPyPAPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *AdvisoryPyPAAffected) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetRanges

`func (o *AdvisoryPyPAAffected) GetRanges() []AdvisoryPyPARange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *AdvisoryPyPAAffected) GetRangesOk() (*[]AdvisoryPyPARange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *AdvisoryPyPAAffected) SetRanges(v []AdvisoryPyPARange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *AdvisoryPyPAAffected) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetVersions

`func (o *AdvisoryPyPAAffected) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AdvisoryPyPAAffected) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AdvisoryPyPAAffected) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AdvisoryPyPAAffected) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


