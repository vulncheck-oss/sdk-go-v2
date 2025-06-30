# AdvisoryDebianSecurityAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedPackages** | Pointer to [**[]AdvisoryAffectedDebianPackage**](AdvisoryAffectedDebianPackage.md) |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Dsa** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryDebianSecurityAdvisory

`func NewAdvisoryDebianSecurityAdvisory() *AdvisoryDebianSecurityAdvisory`

NewAdvisoryDebianSecurityAdvisory instantiates a new AdvisoryDebianSecurityAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryDebianSecurityAdvisoryWithDefaults

`func NewAdvisoryDebianSecurityAdvisoryWithDefaults() *AdvisoryDebianSecurityAdvisory`

NewAdvisoryDebianSecurityAdvisoryWithDefaults instantiates a new AdvisoryDebianSecurityAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedPackages

`func (o *AdvisoryDebianSecurityAdvisory) GetAffectedPackages() []AdvisoryAffectedDebianPackage`

GetAffectedPackages returns the AffectedPackages field if non-nil, zero value otherwise.

### GetAffectedPackagesOk

`func (o *AdvisoryDebianSecurityAdvisory) GetAffectedPackagesOk() (*[]AdvisoryAffectedDebianPackage, bool)`

GetAffectedPackagesOk returns a tuple with the AffectedPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPackages

`func (o *AdvisoryDebianSecurityAdvisory) SetAffectedPackages(v []AdvisoryAffectedDebianPackage)`

SetAffectedPackages sets AffectedPackages field to given value.

### HasAffectedPackages

`func (o *AdvisoryDebianSecurityAdvisory) HasAffectedPackages() bool`

HasAffectedPackages returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryDebianSecurityAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryDebianSecurityAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryDebianSecurityAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryDebianSecurityAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryDebianSecurityAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryDebianSecurityAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryDebianSecurityAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryDebianSecurityAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDsa

`func (o *AdvisoryDebianSecurityAdvisory) GetDsa() string`

GetDsa returns the Dsa field if non-nil, zero value otherwise.

### GetDsaOk

`func (o *AdvisoryDebianSecurityAdvisory) GetDsaOk() (*string, bool)`

GetDsaOk returns a tuple with the Dsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsa

`func (o *AdvisoryDebianSecurityAdvisory) SetDsa(v string)`

SetDsa sets Dsa field to given value.

### HasDsa

`func (o *AdvisoryDebianSecurityAdvisory) HasDsa() bool`

HasDsa returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryDebianSecurityAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryDebianSecurityAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryDebianSecurityAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryDebianSecurityAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryDebianSecurityAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryDebianSecurityAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryDebianSecurityAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryDebianSecurityAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


