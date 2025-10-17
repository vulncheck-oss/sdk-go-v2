# AdvisoryDistroPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binary** | Pointer to **bool** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**License** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SecFixes** | Pointer to [**[]AdvisorySecFix**](AdvisorySecFix.md) |  | [optional] 
**Versions** | Pointer to [**[]AdvisoryDistroVersion**](AdvisoryDistroVersion.md) |  | [optional] 

## Methods

### NewAdvisoryDistroPackage

`func NewAdvisoryDistroPackage() *AdvisoryDistroPackage`

NewAdvisoryDistroPackage instantiates a new AdvisoryDistroPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryDistroPackageWithDefaults

`func NewAdvisoryDistroPackageWithDefaults() *AdvisoryDistroPackage`

NewAdvisoryDistroPackageWithDefaults instantiates a new AdvisoryDistroPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinary

`func (o *AdvisoryDistroPackage) GetBinary() bool`

GetBinary returns the Binary field if non-nil, zero value otherwise.

### GetBinaryOk

`func (o *AdvisoryDistroPackage) GetBinaryOk() (*bool, bool)`

GetBinaryOk returns a tuple with the Binary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinary

`func (o *AdvisoryDistroPackage) SetBinary(v bool)`

SetBinary sets Binary field to given value.

### HasBinary

`func (o *AdvisoryDistroPackage) HasBinary() bool`

HasBinary returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryDistroPackage) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryDistroPackage) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryDistroPackage) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryDistroPackage) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetLicense

`func (o *AdvisoryDistroPackage) GetLicense() []string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *AdvisoryDistroPackage) GetLicenseOk() (*[]string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *AdvisoryDistroPackage) SetLicense(v []string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *AdvisoryDistroPackage) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetName

`func (o *AdvisoryDistroPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvisoryDistroPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvisoryDistroPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvisoryDistroPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecFixes

`func (o *AdvisoryDistroPackage) GetSecFixes() []AdvisorySecFix`

GetSecFixes returns the SecFixes field if non-nil, zero value otherwise.

### GetSecFixesOk

`func (o *AdvisoryDistroPackage) GetSecFixesOk() (*[]AdvisorySecFix, bool)`

GetSecFixesOk returns a tuple with the SecFixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecFixes

`func (o *AdvisoryDistroPackage) SetSecFixes(v []AdvisorySecFix)`

SetSecFixes sets SecFixes field to given value.

### HasSecFixes

`func (o *AdvisoryDistroPackage) HasSecFixes() bool`

HasSecFixes returns a boolean if a field has been set.

### GetVersions

`func (o *AdvisoryDistroPackage) GetVersions() []AdvisoryDistroVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AdvisoryDistroPackage) GetVersionsOk() (*[]AdvisoryDistroVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AdvisoryDistroPackage) SetVersions(v []AdvisoryDistroVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AdvisoryDistroPackage) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


