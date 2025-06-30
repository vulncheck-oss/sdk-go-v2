# AdvisoryRockyErrata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advisory** | Pointer to [**AdvisoryRockyAdvisory**](AdvisoryRockyAdvisory.md) |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Packages** | Pointer to [**[]AdvisoryRockyPackage**](AdvisoryRockyPackage.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryRockyErrata

`func NewAdvisoryRockyErrata() *AdvisoryRockyErrata`

NewAdvisoryRockyErrata instantiates a new AdvisoryRockyErrata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryRockyErrataWithDefaults

`func NewAdvisoryRockyErrataWithDefaults() *AdvisoryRockyErrata`

NewAdvisoryRockyErrataWithDefaults instantiates a new AdvisoryRockyErrata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisory

`func (o *AdvisoryRockyErrata) GetAdvisory() AdvisoryRockyAdvisory`

GetAdvisory returns the Advisory field if non-nil, zero value otherwise.

### GetAdvisoryOk

`func (o *AdvisoryRockyErrata) GetAdvisoryOk() (*AdvisoryRockyAdvisory, bool)`

GetAdvisoryOk returns a tuple with the Advisory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisory

`func (o *AdvisoryRockyErrata) SetAdvisory(v AdvisoryRockyAdvisory)`

SetAdvisory sets Advisory field to given value.

### HasAdvisory

`func (o *AdvisoryRockyErrata) HasAdvisory() bool`

HasAdvisory returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryRockyErrata) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryRockyErrata) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryRockyErrata) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryRockyErrata) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryRockyErrata) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryRockyErrata) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryRockyErrata) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryRockyErrata) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetPackages

`func (o *AdvisoryRockyErrata) GetPackages() []AdvisoryRockyPackage`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *AdvisoryRockyErrata) GetPackagesOk() (*[]AdvisoryRockyPackage, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *AdvisoryRockyErrata) SetPackages(v []AdvisoryRockyPackage)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *AdvisoryRockyErrata) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryRockyErrata) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryRockyErrata) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryRockyErrata) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryRockyErrata) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


