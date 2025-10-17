# AdvisoryHaskellAffected

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedConstraint** | Pointer to **string** |  | [optional] 
**AffectedVersions** | Pointer to [**[]AdvisoryHaskellVersion**](AdvisoryHaskellVersion.md) |  | [optional] 
**Arch** | Pointer to **[]string** |  | [optional] 
**Cvss** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **[]string** |  | [optional] 
**Package** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryHaskellAffected

`func NewAdvisoryHaskellAffected() *AdvisoryHaskellAffected`

NewAdvisoryHaskellAffected instantiates a new AdvisoryHaskellAffected object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryHaskellAffectedWithDefaults

`func NewAdvisoryHaskellAffectedWithDefaults() *AdvisoryHaskellAffected`

NewAdvisoryHaskellAffectedWithDefaults instantiates a new AdvisoryHaskellAffected object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedConstraint

`func (o *AdvisoryHaskellAffected) GetAffectedConstraint() string`

GetAffectedConstraint returns the AffectedConstraint field if non-nil, zero value otherwise.

### GetAffectedConstraintOk

`func (o *AdvisoryHaskellAffected) GetAffectedConstraintOk() (*string, bool)`

GetAffectedConstraintOk returns a tuple with the AffectedConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedConstraint

`func (o *AdvisoryHaskellAffected) SetAffectedConstraint(v string)`

SetAffectedConstraint sets AffectedConstraint field to given value.

### HasAffectedConstraint

`func (o *AdvisoryHaskellAffected) HasAffectedConstraint() bool`

HasAffectedConstraint returns a boolean if a field has been set.

### GetAffectedVersions

`func (o *AdvisoryHaskellAffected) GetAffectedVersions() []AdvisoryHaskellVersion`

GetAffectedVersions returns the AffectedVersions field if non-nil, zero value otherwise.

### GetAffectedVersionsOk

`func (o *AdvisoryHaskellAffected) GetAffectedVersionsOk() (*[]AdvisoryHaskellVersion, bool)`

GetAffectedVersionsOk returns a tuple with the AffectedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVersions

`func (o *AdvisoryHaskellAffected) SetAffectedVersions(v []AdvisoryHaskellVersion)`

SetAffectedVersions sets AffectedVersions field to given value.

### HasAffectedVersions

`func (o *AdvisoryHaskellAffected) HasAffectedVersions() bool`

HasAffectedVersions returns a boolean if a field has been set.

### GetArch

`func (o *AdvisoryHaskellAffected) GetArch() []string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *AdvisoryHaskellAffected) GetArchOk() (*[]string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *AdvisoryHaskellAffected) SetArch(v []string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *AdvisoryHaskellAffected) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetCvss

`func (o *AdvisoryHaskellAffected) GetCvss() string`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *AdvisoryHaskellAffected) GetCvssOk() (*string, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *AdvisoryHaskellAffected) SetCvss(v string)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *AdvisoryHaskellAffected) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetOs

`func (o *AdvisoryHaskellAffected) GetOs() []string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *AdvisoryHaskellAffected) GetOsOk() (*[]string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *AdvisoryHaskellAffected) SetOs(v []string)`

SetOs sets Os field to given value.

### HasOs

`func (o *AdvisoryHaskellAffected) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPackage

`func (o *AdvisoryHaskellAffected) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *AdvisoryHaskellAffected) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *AdvisoryHaskellAffected) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *AdvisoryHaskellAffected) HasPackage() bool`

HasPackage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


