# AdvisoryAlmaPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** |  | [optional] 
**Epoch** | Pointer to **string** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RebootSuggested** | Pointer to **int32** |  | [optional] 
**Release** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Sum** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryAlmaPackage

`func NewAdvisoryAlmaPackage() *AdvisoryAlmaPackage`

NewAdvisoryAlmaPackage instantiates a new AdvisoryAlmaPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAlmaPackageWithDefaults

`func NewAdvisoryAlmaPackageWithDefaults() *AdvisoryAlmaPackage`

NewAdvisoryAlmaPackageWithDefaults instantiates a new AdvisoryAlmaPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *AdvisoryAlmaPackage) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *AdvisoryAlmaPackage) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *AdvisoryAlmaPackage) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *AdvisoryAlmaPackage) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetEpoch

`func (o *AdvisoryAlmaPackage) GetEpoch() string`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *AdvisoryAlmaPackage) GetEpochOk() (*string, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *AdvisoryAlmaPackage) SetEpoch(v string)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *AdvisoryAlmaPackage) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### GetFilename

`func (o *AdvisoryAlmaPackage) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *AdvisoryAlmaPackage) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *AdvisoryAlmaPackage) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *AdvisoryAlmaPackage) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetName

`func (o *AdvisoryAlmaPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvisoryAlmaPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvisoryAlmaPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvisoryAlmaPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRebootSuggested

`func (o *AdvisoryAlmaPackage) GetRebootSuggested() int32`

GetRebootSuggested returns the RebootSuggested field if non-nil, zero value otherwise.

### GetRebootSuggestedOk

`func (o *AdvisoryAlmaPackage) GetRebootSuggestedOk() (*int32, bool)`

GetRebootSuggestedOk returns a tuple with the RebootSuggested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootSuggested

`func (o *AdvisoryAlmaPackage) SetRebootSuggested(v int32)`

SetRebootSuggested sets RebootSuggested field to given value.

### HasRebootSuggested

`func (o *AdvisoryAlmaPackage) HasRebootSuggested() bool`

HasRebootSuggested returns a boolean if a field has been set.

### GetRelease

`func (o *AdvisoryAlmaPackage) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *AdvisoryAlmaPackage) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *AdvisoryAlmaPackage) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *AdvisoryAlmaPackage) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetSource

`func (o *AdvisoryAlmaPackage) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AdvisoryAlmaPackage) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AdvisoryAlmaPackage) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AdvisoryAlmaPackage) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSum

`func (o *AdvisoryAlmaPackage) GetSum() string`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *AdvisoryAlmaPackage) GetSumOk() (*string, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *AdvisoryAlmaPackage) SetSum(v string)`

SetSum sets Sum field to given value.

### HasSum

`func (o *AdvisoryAlmaPackage) HasSum() bool`

HasSum returns a boolean if a field has been set.

### GetVersion

`func (o *AdvisoryAlmaPackage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AdvisoryAlmaPackage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AdvisoryAlmaPackage) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AdvisoryAlmaPackage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


