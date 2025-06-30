# AdvisoryAlmaPackageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Packages** | Pointer to [**[]AdvisoryAlmaPackage**](AdvisoryAlmaPackage.md) |  | [optional] 
**Shortname** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryAlmaPackageList

`func NewAdvisoryAlmaPackageList() *AdvisoryAlmaPackageList`

NewAdvisoryAlmaPackageList instantiates a new AdvisoryAlmaPackageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAlmaPackageListWithDefaults

`func NewAdvisoryAlmaPackageListWithDefaults() *AdvisoryAlmaPackageList`

NewAdvisoryAlmaPackageListWithDefaults instantiates a new AdvisoryAlmaPackageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdvisoryAlmaPackageList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvisoryAlmaPackageList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvisoryAlmaPackageList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvisoryAlmaPackageList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackages

`func (o *AdvisoryAlmaPackageList) GetPackages() []AdvisoryAlmaPackage`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *AdvisoryAlmaPackageList) GetPackagesOk() (*[]AdvisoryAlmaPackage, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *AdvisoryAlmaPackageList) SetPackages(v []AdvisoryAlmaPackage)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *AdvisoryAlmaPackageList) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetShortname

`func (o *AdvisoryAlmaPackageList) GetShortname() string`

GetShortname returns the Shortname field if non-nil, zero value otherwise.

### GetShortnameOk

`func (o *AdvisoryAlmaPackageList) GetShortnameOk() (*string, bool)`

GetShortnameOk returns a tuple with the Shortname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortname

`func (o *AdvisoryAlmaPackageList) SetShortname(v string)`

SetShortname sets Shortname field to given value.

### HasShortname

`func (o *AdvisoryAlmaPackageList) HasShortname() bool`

HasShortname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


