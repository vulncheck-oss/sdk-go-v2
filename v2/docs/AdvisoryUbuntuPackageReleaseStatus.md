# AdvisoryUbuntuPackageReleaseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to **bool** |  | [optional] 
**Fixed** | Pointer to **bool** |  | [optional] 
**FixedVersion** | Pointer to **string** |  | [optional] 
**Lts** | Pointer to **bool** |  | [optional] 
**Release** | Pointer to **string** |  | [optional] 
**ReleaseLong** | Pointer to **string** |  | [optional] 
**ReleaseVersion** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryUbuntuPackageReleaseStatus

`func NewAdvisoryUbuntuPackageReleaseStatus() *AdvisoryUbuntuPackageReleaseStatus`

NewAdvisoryUbuntuPackageReleaseStatus instantiates a new AdvisoryUbuntuPackageReleaseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryUbuntuPackageReleaseStatusWithDefaults

`func NewAdvisoryUbuntuPackageReleaseStatusWithDefaults() *AdvisoryUbuntuPackageReleaseStatus`

NewAdvisoryUbuntuPackageReleaseStatusWithDefaults instantiates a new AdvisoryUbuntuPackageReleaseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *AdvisoryUbuntuPackageReleaseStatus) GetAffected() bool`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryUbuntuPackageReleaseStatus) GetAffectedOk() (*bool, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryUbuntuPackageReleaseStatus) SetAffected(v bool)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryUbuntuPackageReleaseStatus) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetFixed

`func (o *AdvisoryUbuntuPackageReleaseStatus) GetFixed() bool`

GetFixed returns the Fixed field if non-nil, zero value otherwise.

### GetFixedOk

`func (o *AdvisoryUbuntuPackageReleaseStatus) GetFixedOk() (*bool, bool)`

GetFixedOk returns a tuple with the Fixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixed

`func (o *AdvisoryUbuntuPackageReleaseStatus) SetFixed(v bool)`

SetFixed sets Fixed field to given value.

### HasFixed

`func (o *AdvisoryUbuntuPackageReleaseStatus) HasFixed() bool`

HasFixed returns a boolean if a field has been set.

### GetFixedVersion

`func (o *AdvisoryUbuntuPackageReleaseStatus) GetFixedVersion() string`

GetFixedVersion returns the FixedVersion field if non-nil, zero value otherwise.

### GetFixedVersionOk

`func (o *AdvisoryUbuntuPackageReleaseStatus) GetFixedVersionOk() (*string, bool)`

GetFixedVersionOk returns a tuple with the FixedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedVersion

`func (o *AdvisoryUbuntuPackageReleaseStatus) SetFixedVersion(v string)`

SetFixedVersion sets FixedVersion field to given value.

### HasFixedVersion

`func (o *AdvisoryUbuntuPackageReleaseStatus) HasFixedVersion() bool`

HasFixedVersion returns a boolean if a field has been set.

### GetLts

`func (o *AdvisoryUbuntuPackageReleaseStatus) GetLts() bool`

GetLts returns the Lts field if non-nil, zero value otherwise.

### GetLtsOk

`func (o *AdvisoryUbuntuPackageReleaseStatus) GetLtsOk() (*bool, bool)`

GetLtsOk returns a tuple with the Lts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLts

`func (o *AdvisoryUbuntuPackageReleaseStatus) SetLts(v bool)`

SetLts sets Lts field to given value.

### HasLts

`func (o *AdvisoryUbuntuPackageReleaseStatus) HasLts() bool`

HasLts returns a boolean if a field has been set.

### GetRelease

`func (o *AdvisoryUbuntuPackageReleaseStatus) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *AdvisoryUbuntuPackageReleaseStatus) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *AdvisoryUbuntuPackageReleaseStatus) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *AdvisoryUbuntuPackageReleaseStatus) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetReleaseLong

`func (o *AdvisoryUbuntuPackageReleaseStatus) GetReleaseLong() string`

GetReleaseLong returns the ReleaseLong field if non-nil, zero value otherwise.

### GetReleaseLongOk

`func (o *AdvisoryUbuntuPackageReleaseStatus) GetReleaseLongOk() (*string, bool)`

GetReleaseLongOk returns a tuple with the ReleaseLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseLong

`func (o *AdvisoryUbuntuPackageReleaseStatus) SetReleaseLong(v string)`

SetReleaseLong sets ReleaseLong field to given value.

### HasReleaseLong

`func (o *AdvisoryUbuntuPackageReleaseStatus) HasReleaseLong() bool`

HasReleaseLong returns a boolean if a field has been set.

### GetReleaseVersion

`func (o *AdvisoryUbuntuPackageReleaseStatus) GetReleaseVersion() string`

GetReleaseVersion returns the ReleaseVersion field if non-nil, zero value otherwise.

### GetReleaseVersionOk

`func (o *AdvisoryUbuntuPackageReleaseStatus) GetReleaseVersionOk() (*string, bool)`

GetReleaseVersionOk returns a tuple with the ReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersion

`func (o *AdvisoryUbuntuPackageReleaseStatus) SetReleaseVersion(v string)`

SetReleaseVersion sets ReleaseVersion field to given value.

### HasReleaseVersion

`func (o *AdvisoryUbuntuPackageReleaseStatus) HasReleaseVersion() bool`

HasReleaseVersion returns a boolean if a field has been set.

### GetStatus

`func (o *AdvisoryUbuntuPackageReleaseStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdvisoryUbuntuPackageReleaseStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdvisoryUbuntuPackageReleaseStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdvisoryUbuntuPackageReleaseStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


