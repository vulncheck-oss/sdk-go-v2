# AdvisoryChainGuard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apkurl** | Pointer to **string** |  | [optional] 
**Archs** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** | un-used | [optional] 
**Packages** | Pointer to [**[]AdvisoryChainGuardPackage**](AdvisoryChainGuardPackage.md) |  | [optional] 
**Reponame** | Pointer to **string** |  | [optional] 
**Urlprefix** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryChainGuard

`func NewAdvisoryChainGuard() *AdvisoryChainGuard`

NewAdvisoryChainGuard instantiates a new AdvisoryChainGuard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryChainGuardWithDefaults

`func NewAdvisoryChainGuardWithDefaults() *AdvisoryChainGuard`

NewAdvisoryChainGuardWithDefaults instantiates a new AdvisoryChainGuard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApkurl

`func (o *AdvisoryChainGuard) GetApkurl() string`

GetApkurl returns the Apkurl field if non-nil, zero value otherwise.

### GetApkurlOk

`func (o *AdvisoryChainGuard) GetApkurlOk() (*string, bool)`

GetApkurlOk returns a tuple with the Apkurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApkurl

`func (o *AdvisoryChainGuard) SetApkurl(v string)`

SetApkurl sets Apkurl field to given value.

### HasApkurl

`func (o *AdvisoryChainGuard) HasApkurl() bool`

HasApkurl returns a boolean if a field has been set.

### GetArchs

`func (o *AdvisoryChainGuard) GetArchs() []string`

GetArchs returns the Archs field if non-nil, zero value otherwise.

### GetArchsOk

`func (o *AdvisoryChainGuard) GetArchsOk() (*[]string, bool)`

GetArchsOk returns a tuple with the Archs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchs

`func (o *AdvisoryChainGuard) SetArchs(v []string)`

SetArchs sets Archs field to given value.

### HasArchs

`func (o *AdvisoryChainGuard) HasArchs() bool`

HasArchs returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryChainGuard) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryChainGuard) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryChainGuard) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryChainGuard) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetPackages

`func (o *AdvisoryChainGuard) GetPackages() []AdvisoryChainGuardPackage`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *AdvisoryChainGuard) GetPackagesOk() (*[]AdvisoryChainGuardPackage, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *AdvisoryChainGuard) SetPackages(v []AdvisoryChainGuardPackage)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *AdvisoryChainGuard) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetReponame

`func (o *AdvisoryChainGuard) GetReponame() string`

GetReponame returns the Reponame field if non-nil, zero value otherwise.

### GetReponameOk

`func (o *AdvisoryChainGuard) GetReponameOk() (*string, bool)`

GetReponameOk returns a tuple with the Reponame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReponame

`func (o *AdvisoryChainGuard) SetReponame(v string)`

SetReponame sets Reponame field to given value.

### HasReponame

`func (o *AdvisoryChainGuard) HasReponame() bool`

HasReponame returns a boolean if a field has been set.

### GetUrlprefix

`func (o *AdvisoryChainGuard) GetUrlprefix() string`

GetUrlprefix returns the Urlprefix field if non-nil, zero value otherwise.

### GetUrlprefixOk

`func (o *AdvisoryChainGuard) GetUrlprefixOk() (*string, bool)`

GetUrlprefixOk returns a tuple with the Urlprefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlprefix

`func (o *AdvisoryChainGuard) SetUrlprefix(v string)`

SetUrlprefix sets Urlprefix field to given value.

### HasUrlprefix

`func (o *AdvisoryChainGuard) HasUrlprefix() bool`

HasUrlprefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


