# AdvisoryAndroidAffected

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EcosystemSpecific** | Pointer to [**AdvisoryEcoSystem**](AdvisoryEcoSystem.md) |  | [optional] 
**Package** | Pointer to [**AdvisoryAndroidPackage**](AdvisoryAndroidPackage.md) |  | [optional] 
**Ranges** | Pointer to [**[]AdvisoryAndroidRange**](AdvisoryAndroidRange.md) |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdvisoryAndroidAffected

`func NewAdvisoryAndroidAffected() *AdvisoryAndroidAffected`

NewAdvisoryAndroidAffected instantiates a new AdvisoryAndroidAffected object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAndroidAffectedWithDefaults

`func NewAdvisoryAndroidAffectedWithDefaults() *AdvisoryAndroidAffected`

NewAdvisoryAndroidAffectedWithDefaults instantiates a new AdvisoryAndroidAffected object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEcosystemSpecific

`func (o *AdvisoryAndroidAffected) GetEcosystemSpecific() AdvisoryEcoSystem`

GetEcosystemSpecific returns the EcosystemSpecific field if non-nil, zero value otherwise.

### GetEcosystemSpecificOk

`func (o *AdvisoryAndroidAffected) GetEcosystemSpecificOk() (*AdvisoryEcoSystem, bool)`

GetEcosystemSpecificOk returns a tuple with the EcosystemSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcosystemSpecific

`func (o *AdvisoryAndroidAffected) SetEcosystemSpecific(v AdvisoryEcoSystem)`

SetEcosystemSpecific sets EcosystemSpecific field to given value.

### HasEcosystemSpecific

`func (o *AdvisoryAndroidAffected) HasEcosystemSpecific() bool`

HasEcosystemSpecific returns a boolean if a field has been set.

### GetPackage

`func (o *AdvisoryAndroidAffected) GetPackage() AdvisoryAndroidPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *AdvisoryAndroidAffected) GetPackageOk() (*AdvisoryAndroidPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *AdvisoryAndroidAffected) SetPackage(v AdvisoryAndroidPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *AdvisoryAndroidAffected) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetRanges

`func (o *AdvisoryAndroidAffected) GetRanges() []AdvisoryAndroidRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *AdvisoryAndroidAffected) GetRangesOk() (*[]AdvisoryAndroidRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *AdvisoryAndroidAffected) SetRanges(v []AdvisoryAndroidRange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *AdvisoryAndroidAffected) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetVersions

`func (o *AdvisoryAndroidAffected) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AdvisoryAndroidAffected) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AdvisoryAndroidAffected) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AdvisoryAndroidAffected) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


