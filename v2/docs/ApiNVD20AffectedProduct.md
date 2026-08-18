# ApiNVD20AffectedProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionURL** | Pointer to **string** |  | [optional] 
**Cpes** | Pointer to **[]string** |  | [optional] 
**DefaultStatus** | Pointer to **string** | DefaultStatus is the status for versions not otherwise listed: one of \&quot;affected\&quot;, \&quot;unaffected\&quot;, or \&quot;unknown\&quot;. | [optional] 
**Modules** | Pointer to **[]string** |  | [optional] 
**PackageName** | Pointer to **string** |  | [optional] 
**PackageURL** | Pointer to **string** |  | [optional] 
**Platforms** | Pointer to **[]string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**ProgramFiles** | Pointer to **[]string** |  | [optional] 
**ProgramRoutines** | Pointer to [**[]ApiNVD20AffectedProgramRoutine**](ApiNVD20AffectedProgramRoutine.md) |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Versions** | Pointer to [**[]ApiNVD20AffectedVersion**](ApiNVD20AffectedVersion.md) |  | [optional] 

## Methods

### NewApiNVD20AffectedProduct

`func NewApiNVD20AffectedProduct() *ApiNVD20AffectedProduct`

NewApiNVD20AffectedProduct instantiates a new ApiNVD20AffectedProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNVD20AffectedProductWithDefaults

`func NewApiNVD20AffectedProductWithDefaults() *ApiNVD20AffectedProduct`

NewApiNVD20AffectedProductWithDefaults instantiates a new ApiNVD20AffectedProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionURL

`func (o *ApiNVD20AffectedProduct) GetCollectionURL() string`

GetCollectionURL returns the CollectionURL field if non-nil, zero value otherwise.

### GetCollectionURLOk

`func (o *ApiNVD20AffectedProduct) GetCollectionURLOk() (*string, bool)`

GetCollectionURLOk returns a tuple with the CollectionURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionURL

`func (o *ApiNVD20AffectedProduct) SetCollectionURL(v string)`

SetCollectionURL sets CollectionURL field to given value.

### HasCollectionURL

`func (o *ApiNVD20AffectedProduct) HasCollectionURL() bool`

HasCollectionURL returns a boolean if a field has been set.

### GetCpes

`func (o *ApiNVD20AffectedProduct) GetCpes() []string`

GetCpes returns the Cpes field if non-nil, zero value otherwise.

### GetCpesOk

`func (o *ApiNVD20AffectedProduct) GetCpesOk() (*[]string, bool)`

GetCpesOk returns a tuple with the Cpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpes

`func (o *ApiNVD20AffectedProduct) SetCpes(v []string)`

SetCpes sets Cpes field to given value.

### HasCpes

`func (o *ApiNVD20AffectedProduct) HasCpes() bool`

HasCpes returns a boolean if a field has been set.

### GetDefaultStatus

`func (o *ApiNVD20AffectedProduct) GetDefaultStatus() string`

GetDefaultStatus returns the DefaultStatus field if non-nil, zero value otherwise.

### GetDefaultStatusOk

`func (o *ApiNVD20AffectedProduct) GetDefaultStatusOk() (*string, bool)`

GetDefaultStatusOk returns a tuple with the DefaultStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStatus

`func (o *ApiNVD20AffectedProduct) SetDefaultStatus(v string)`

SetDefaultStatus sets DefaultStatus field to given value.

### HasDefaultStatus

`func (o *ApiNVD20AffectedProduct) HasDefaultStatus() bool`

HasDefaultStatus returns a boolean if a field has been set.

### GetModules

`func (o *ApiNVD20AffectedProduct) GetModules() []string`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ApiNVD20AffectedProduct) GetModulesOk() (*[]string, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ApiNVD20AffectedProduct) SetModules(v []string)`

SetModules sets Modules field to given value.

### HasModules

`func (o *ApiNVD20AffectedProduct) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetPackageName

`func (o *ApiNVD20AffectedProduct) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *ApiNVD20AffectedProduct) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *ApiNVD20AffectedProduct) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *ApiNVD20AffectedProduct) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetPackageURL

`func (o *ApiNVD20AffectedProduct) GetPackageURL() string`

GetPackageURL returns the PackageURL field if non-nil, zero value otherwise.

### GetPackageURLOk

`func (o *ApiNVD20AffectedProduct) GetPackageURLOk() (*string, bool)`

GetPackageURLOk returns a tuple with the PackageURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageURL

`func (o *ApiNVD20AffectedProduct) SetPackageURL(v string)`

SetPackageURL sets PackageURL field to given value.

### HasPackageURL

`func (o *ApiNVD20AffectedProduct) HasPackageURL() bool`

HasPackageURL returns a boolean if a field has been set.

### GetPlatforms

`func (o *ApiNVD20AffectedProduct) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *ApiNVD20AffectedProduct) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *ApiNVD20AffectedProduct) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *ApiNVD20AffectedProduct) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetProduct

`func (o *ApiNVD20AffectedProduct) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ApiNVD20AffectedProduct) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ApiNVD20AffectedProduct) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ApiNVD20AffectedProduct) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProgramFiles

`func (o *ApiNVD20AffectedProduct) GetProgramFiles() []string`

GetProgramFiles returns the ProgramFiles field if non-nil, zero value otherwise.

### GetProgramFilesOk

`func (o *ApiNVD20AffectedProduct) GetProgramFilesOk() (*[]string, bool)`

GetProgramFilesOk returns a tuple with the ProgramFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramFiles

`func (o *ApiNVD20AffectedProduct) SetProgramFiles(v []string)`

SetProgramFiles sets ProgramFiles field to given value.

### HasProgramFiles

`func (o *ApiNVD20AffectedProduct) HasProgramFiles() bool`

HasProgramFiles returns a boolean if a field has been set.

### GetProgramRoutines

`func (o *ApiNVD20AffectedProduct) GetProgramRoutines() []ApiNVD20AffectedProgramRoutine`

GetProgramRoutines returns the ProgramRoutines field if non-nil, zero value otherwise.

### GetProgramRoutinesOk

`func (o *ApiNVD20AffectedProduct) GetProgramRoutinesOk() (*[]ApiNVD20AffectedProgramRoutine, bool)`

GetProgramRoutinesOk returns a tuple with the ProgramRoutines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramRoutines

`func (o *ApiNVD20AffectedProduct) SetProgramRoutines(v []ApiNVD20AffectedProgramRoutine)`

SetProgramRoutines sets ProgramRoutines field to given value.

### HasProgramRoutines

`func (o *ApiNVD20AffectedProduct) HasProgramRoutines() bool`

HasProgramRoutines returns a boolean if a field has been set.

### GetRepo

`func (o *ApiNVD20AffectedProduct) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ApiNVD20AffectedProduct) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ApiNVD20AffectedProduct) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ApiNVD20AffectedProduct) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetVendor

`func (o *ApiNVD20AffectedProduct) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ApiNVD20AffectedProduct) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ApiNVD20AffectedProduct) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ApiNVD20AffectedProduct) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersions

`func (o *ApiNVD20AffectedProduct) GetVersions() []ApiNVD20AffectedVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ApiNVD20AffectedProduct) GetVersionsOk() (*[]ApiNVD20AffectedVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ApiNVD20AffectedProduct) SetVersions(v []ApiNVD20AffectedVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ApiNVD20AffectedProduct) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


