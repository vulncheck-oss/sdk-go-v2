# AdvisoryVulnerableDebianPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedCves** | Pointer to [**[]AdvisoryDebianCVE**](AdvisoryDebianCVE.md) |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**PackageName** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryVulnerableDebianPackage

`func NewAdvisoryVulnerableDebianPackage() *AdvisoryVulnerableDebianPackage`

NewAdvisoryVulnerableDebianPackage instantiates a new AdvisoryVulnerableDebianPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryVulnerableDebianPackageWithDefaults

`func NewAdvisoryVulnerableDebianPackageWithDefaults() *AdvisoryVulnerableDebianPackage`

NewAdvisoryVulnerableDebianPackageWithDefaults instantiates a new AdvisoryVulnerableDebianPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedCves

`func (o *AdvisoryVulnerableDebianPackage) GetAssociatedCves() []AdvisoryDebianCVE`

GetAssociatedCves returns the AssociatedCves field if non-nil, zero value otherwise.

### GetAssociatedCvesOk

`func (o *AdvisoryVulnerableDebianPackage) GetAssociatedCvesOk() (*[]AdvisoryDebianCVE, bool)`

GetAssociatedCvesOk returns a tuple with the AssociatedCves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedCves

`func (o *AdvisoryVulnerableDebianPackage) SetAssociatedCves(v []AdvisoryDebianCVE)`

SetAssociatedCves sets AssociatedCves field to given value.

### HasAssociatedCves

`func (o *AdvisoryVulnerableDebianPackage) HasAssociatedCves() bool`

HasAssociatedCves returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryVulnerableDebianPackage) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryVulnerableDebianPackage) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryVulnerableDebianPackage) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryVulnerableDebianPackage) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetPackageName

`func (o *AdvisoryVulnerableDebianPackage) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *AdvisoryVulnerableDebianPackage) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *AdvisoryVulnerableDebianPackage) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *AdvisoryVulnerableDebianPackage) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


