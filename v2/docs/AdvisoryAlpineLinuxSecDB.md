# AdvisoryAlpineLinuxSecDB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apkurl** | Pointer to **string** |  | [optional] 
**Archs** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Distroversion** | Pointer to **string** |  | [optional] 
**Packages** | Pointer to [**[]AdvisoryAlpineLinuxSecDBPackage**](AdvisoryAlpineLinuxSecDBPackage.md) |  | [optional] 
**Reponame** | Pointer to **string** |  | [optional] 
**Urlprefix** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryAlpineLinuxSecDB

`func NewAdvisoryAlpineLinuxSecDB() *AdvisoryAlpineLinuxSecDB`

NewAdvisoryAlpineLinuxSecDB instantiates a new AdvisoryAlpineLinuxSecDB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAlpineLinuxSecDBWithDefaults

`func NewAdvisoryAlpineLinuxSecDBWithDefaults() *AdvisoryAlpineLinuxSecDB`

NewAdvisoryAlpineLinuxSecDBWithDefaults instantiates a new AdvisoryAlpineLinuxSecDB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApkurl

`func (o *AdvisoryAlpineLinuxSecDB) GetApkurl() string`

GetApkurl returns the Apkurl field if non-nil, zero value otherwise.

### GetApkurlOk

`func (o *AdvisoryAlpineLinuxSecDB) GetApkurlOk() (*string, bool)`

GetApkurlOk returns a tuple with the Apkurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApkurl

`func (o *AdvisoryAlpineLinuxSecDB) SetApkurl(v string)`

SetApkurl sets Apkurl field to given value.

### HasApkurl

`func (o *AdvisoryAlpineLinuxSecDB) HasApkurl() bool`

HasApkurl returns a boolean if a field has been set.

### GetArchs

`func (o *AdvisoryAlpineLinuxSecDB) GetArchs() []string`

GetArchs returns the Archs field if non-nil, zero value otherwise.

### GetArchsOk

`func (o *AdvisoryAlpineLinuxSecDB) GetArchsOk() (*[]string, bool)`

GetArchsOk returns a tuple with the Archs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchs

`func (o *AdvisoryAlpineLinuxSecDB) SetArchs(v []string)`

SetArchs sets Archs field to given value.

### HasArchs

`func (o *AdvisoryAlpineLinuxSecDB) HasArchs() bool`

HasArchs returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryAlpineLinuxSecDB) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryAlpineLinuxSecDB) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryAlpineLinuxSecDB) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryAlpineLinuxSecDB) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryAlpineLinuxSecDB) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryAlpineLinuxSecDB) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryAlpineLinuxSecDB) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryAlpineLinuxSecDB) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDistroversion

`func (o *AdvisoryAlpineLinuxSecDB) GetDistroversion() string`

GetDistroversion returns the Distroversion field if non-nil, zero value otherwise.

### GetDistroversionOk

`func (o *AdvisoryAlpineLinuxSecDB) GetDistroversionOk() (*string, bool)`

GetDistroversionOk returns a tuple with the Distroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroversion

`func (o *AdvisoryAlpineLinuxSecDB) SetDistroversion(v string)`

SetDistroversion sets Distroversion field to given value.

### HasDistroversion

`func (o *AdvisoryAlpineLinuxSecDB) HasDistroversion() bool`

HasDistroversion returns a boolean if a field has been set.

### GetPackages

`func (o *AdvisoryAlpineLinuxSecDB) GetPackages() []AdvisoryAlpineLinuxSecDBPackage`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *AdvisoryAlpineLinuxSecDB) GetPackagesOk() (*[]AdvisoryAlpineLinuxSecDBPackage, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *AdvisoryAlpineLinuxSecDB) SetPackages(v []AdvisoryAlpineLinuxSecDBPackage)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *AdvisoryAlpineLinuxSecDB) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetReponame

`func (o *AdvisoryAlpineLinuxSecDB) GetReponame() string`

GetReponame returns the Reponame field if non-nil, zero value otherwise.

### GetReponameOk

`func (o *AdvisoryAlpineLinuxSecDB) GetReponameOk() (*string, bool)`

GetReponameOk returns a tuple with the Reponame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReponame

`func (o *AdvisoryAlpineLinuxSecDB) SetReponame(v string)`

SetReponame sets Reponame field to given value.

### HasReponame

`func (o *AdvisoryAlpineLinuxSecDB) HasReponame() bool`

HasReponame returns a boolean if a field has been set.

### GetUrlprefix

`func (o *AdvisoryAlpineLinuxSecDB) GetUrlprefix() string`

GetUrlprefix returns the Urlprefix field if non-nil, zero value otherwise.

### GetUrlprefixOk

`func (o *AdvisoryAlpineLinuxSecDB) GetUrlprefixOk() (*string, bool)`

GetUrlprefixOk returns a tuple with the Urlprefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlprefix

`func (o *AdvisoryAlpineLinuxSecDB) SetUrlprefix(v string)`

SetUrlprefix sets Urlprefix field to given value.

### HasUrlprefix

`func (o *AdvisoryAlpineLinuxSecDB) HasUrlprefix() bool`

HasUrlprefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


