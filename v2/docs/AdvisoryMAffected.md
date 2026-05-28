# AdvisoryMAffected

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionURL** | Pointer to **string** |  | [optional] 
**Cpes** | Pointer to **[]string** |  | [optional] 
**DefaultStatus** | Pointer to **string** |  | [optional] 
**Modules** | Pointer to **[]string** |  | [optional] 
**PackageName** | Pointer to **string** |  | [optional] 
**PackageURL** | Pointer to **string** |  | [optional] 
**Platforms** | Pointer to **[]string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**ProgramFiles** | Pointer to **[]string** |  | [optional] 
**ProgramRoutines** | Pointer to [**[]AdvisoryProgramRoutine**](AdvisoryProgramRoutine.md) |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Versions** | Pointer to [**[]AdvisoryMVersion**](AdvisoryMVersion.md) |  | [optional] 

## Methods

### NewAdvisoryMAffected

`func NewAdvisoryMAffected() *AdvisoryMAffected`

NewAdvisoryMAffected instantiates a new AdvisoryMAffected object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMAffectedWithDefaults

`func NewAdvisoryMAffectedWithDefaults() *AdvisoryMAffected`

NewAdvisoryMAffectedWithDefaults instantiates a new AdvisoryMAffected object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionURL

`func (o *AdvisoryMAffected) GetCollectionURL() string`

GetCollectionURL returns the CollectionURL field if non-nil, zero value otherwise.

### GetCollectionURLOk

`func (o *AdvisoryMAffected) GetCollectionURLOk() (*string, bool)`

GetCollectionURLOk returns a tuple with the CollectionURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionURL

`func (o *AdvisoryMAffected) SetCollectionURL(v string)`

SetCollectionURL sets CollectionURL field to given value.

### HasCollectionURL

`func (o *AdvisoryMAffected) HasCollectionURL() bool`

HasCollectionURL returns a boolean if a field has been set.

### GetCpes

`func (o *AdvisoryMAffected) GetCpes() []string`

GetCpes returns the Cpes field if non-nil, zero value otherwise.

### GetCpesOk

`func (o *AdvisoryMAffected) GetCpesOk() (*[]string, bool)`

GetCpesOk returns a tuple with the Cpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpes

`func (o *AdvisoryMAffected) SetCpes(v []string)`

SetCpes sets Cpes field to given value.

### HasCpes

`func (o *AdvisoryMAffected) HasCpes() bool`

HasCpes returns a boolean if a field has been set.

### GetDefaultStatus

`func (o *AdvisoryMAffected) GetDefaultStatus() string`

GetDefaultStatus returns the DefaultStatus field if non-nil, zero value otherwise.

### GetDefaultStatusOk

`func (o *AdvisoryMAffected) GetDefaultStatusOk() (*string, bool)`

GetDefaultStatusOk returns a tuple with the DefaultStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStatus

`func (o *AdvisoryMAffected) SetDefaultStatus(v string)`

SetDefaultStatus sets DefaultStatus field to given value.

### HasDefaultStatus

`func (o *AdvisoryMAffected) HasDefaultStatus() bool`

HasDefaultStatus returns a boolean if a field has been set.

### GetModules

`func (o *AdvisoryMAffected) GetModules() []string`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *AdvisoryMAffected) GetModulesOk() (*[]string, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *AdvisoryMAffected) SetModules(v []string)`

SetModules sets Modules field to given value.

### HasModules

`func (o *AdvisoryMAffected) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetPackageName

`func (o *AdvisoryMAffected) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *AdvisoryMAffected) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *AdvisoryMAffected) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *AdvisoryMAffected) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetPackageURL

`func (o *AdvisoryMAffected) GetPackageURL() string`

GetPackageURL returns the PackageURL field if non-nil, zero value otherwise.

### GetPackageURLOk

`func (o *AdvisoryMAffected) GetPackageURLOk() (*string, bool)`

GetPackageURLOk returns a tuple with the PackageURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageURL

`func (o *AdvisoryMAffected) SetPackageURL(v string)`

SetPackageURL sets PackageURL field to given value.

### HasPackageURL

`func (o *AdvisoryMAffected) HasPackageURL() bool`

HasPackageURL returns a boolean if a field has been set.

### GetPlatforms

`func (o *AdvisoryMAffected) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *AdvisoryMAffected) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *AdvisoryMAffected) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *AdvisoryMAffected) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetProduct

`func (o *AdvisoryMAffected) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AdvisoryMAffected) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AdvisoryMAffected) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AdvisoryMAffected) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProgramFiles

`func (o *AdvisoryMAffected) GetProgramFiles() []string`

GetProgramFiles returns the ProgramFiles field if non-nil, zero value otherwise.

### GetProgramFilesOk

`func (o *AdvisoryMAffected) GetProgramFilesOk() (*[]string, bool)`

GetProgramFilesOk returns a tuple with the ProgramFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramFiles

`func (o *AdvisoryMAffected) SetProgramFiles(v []string)`

SetProgramFiles sets ProgramFiles field to given value.

### HasProgramFiles

`func (o *AdvisoryMAffected) HasProgramFiles() bool`

HasProgramFiles returns a boolean if a field has been set.

### GetProgramRoutines

`func (o *AdvisoryMAffected) GetProgramRoutines() []AdvisoryProgramRoutine`

GetProgramRoutines returns the ProgramRoutines field if non-nil, zero value otherwise.

### GetProgramRoutinesOk

`func (o *AdvisoryMAffected) GetProgramRoutinesOk() (*[]AdvisoryProgramRoutine, bool)`

GetProgramRoutinesOk returns a tuple with the ProgramRoutines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramRoutines

`func (o *AdvisoryMAffected) SetProgramRoutines(v []AdvisoryProgramRoutine)`

SetProgramRoutines sets ProgramRoutines field to given value.

### HasProgramRoutines

`func (o *AdvisoryMAffected) HasProgramRoutines() bool`

HasProgramRoutines returns a boolean if a field has been set.

### GetRepo

`func (o *AdvisoryMAffected) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *AdvisoryMAffected) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *AdvisoryMAffected) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *AdvisoryMAffected) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetVendor

`func (o *AdvisoryMAffected) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AdvisoryMAffected) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AdvisoryMAffected) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AdvisoryMAffected) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersions

`func (o *AdvisoryMAffected) GetVersions() []AdvisoryMVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AdvisoryMAffected) GetVersionsOk() (*[]AdvisoryMVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AdvisoryMAffected) SetVersions(v []AdvisoryMVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AdvisoryMAffected) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


