# AdvisoryAlpineLinuxSecDBPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageName** | Pointer to **string** |  | [optional] 
**Secfixes** | Pointer to [**[]AdvisoryAlpineLinuxSecurityFix**](AdvisoryAlpineLinuxSecurityFix.md) |  | [optional] 

## Methods

### NewAdvisoryAlpineLinuxSecDBPackage

`func NewAdvisoryAlpineLinuxSecDBPackage() *AdvisoryAlpineLinuxSecDBPackage`

NewAdvisoryAlpineLinuxSecDBPackage instantiates a new AdvisoryAlpineLinuxSecDBPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAlpineLinuxSecDBPackageWithDefaults

`func NewAdvisoryAlpineLinuxSecDBPackageWithDefaults() *AdvisoryAlpineLinuxSecDBPackage`

NewAdvisoryAlpineLinuxSecDBPackageWithDefaults instantiates a new AdvisoryAlpineLinuxSecDBPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageName

`func (o *AdvisoryAlpineLinuxSecDBPackage) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *AdvisoryAlpineLinuxSecDBPackage) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *AdvisoryAlpineLinuxSecDBPackage) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *AdvisoryAlpineLinuxSecDBPackage) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetSecfixes

`func (o *AdvisoryAlpineLinuxSecDBPackage) GetSecfixes() []AdvisoryAlpineLinuxSecurityFix`

GetSecfixes returns the Secfixes field if non-nil, zero value otherwise.

### GetSecfixesOk

`func (o *AdvisoryAlpineLinuxSecDBPackage) GetSecfixesOk() (*[]AdvisoryAlpineLinuxSecurityFix, bool)`

GetSecfixesOk returns a tuple with the Secfixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecfixes

`func (o *AdvisoryAlpineLinuxSecDBPackage) SetSecfixes(v []AdvisoryAlpineLinuxSecurityFix)`

SetSecfixes sets Secfixes field to given value.

### HasSecfixes

`func (o *AdvisoryAlpineLinuxSecDBPackage) HasSecfixes() bool`

HasSecfixes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


