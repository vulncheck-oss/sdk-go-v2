# AdvisoryCloudAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedVersions** | Pointer to [**[]AdvisoryCloudAdvisoryVersionRange**](AdvisoryCloudAdvisoryVersionRange.md) |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryCloudAdvisory

`func NewAdvisoryCloudAdvisory() *AdvisoryCloudAdvisory`

NewAdvisoryCloudAdvisory instantiates a new AdvisoryCloudAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCloudAdvisoryWithDefaults

`func NewAdvisoryCloudAdvisoryWithDefaults() *AdvisoryCloudAdvisory`

NewAdvisoryCloudAdvisoryWithDefaults instantiates a new AdvisoryCloudAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedVersions

`func (o *AdvisoryCloudAdvisory) GetAffectedVersions() []AdvisoryCloudAdvisoryVersionRange`

GetAffectedVersions returns the AffectedVersions field if non-nil, zero value otherwise.

### GetAffectedVersionsOk

`func (o *AdvisoryCloudAdvisory) GetAffectedVersionsOk() (*[]AdvisoryCloudAdvisoryVersionRange, bool)`

GetAffectedVersionsOk returns a tuple with the AffectedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVersions

`func (o *AdvisoryCloudAdvisory) SetAffectedVersions(v []AdvisoryCloudAdvisoryVersionRange)`

SetAffectedVersions sets AffectedVersions field to given value.

### HasAffectedVersions

`func (o *AdvisoryCloudAdvisory) HasAffectedVersions() bool`

HasAffectedVersions returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryCloudAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryCloudAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryCloudAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryCloudAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryCloudAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryCloudAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryCloudAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryCloudAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetProduct

`func (o *AdvisoryCloudAdvisory) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AdvisoryCloudAdvisory) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AdvisoryCloudAdvisory) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AdvisoryCloudAdvisory) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetService

`func (o *AdvisoryCloudAdvisory) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AdvisoryCloudAdvisory) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AdvisoryCloudAdvisory) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *AdvisoryCloudAdvisory) HasService() bool`

HasService returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryCloudAdvisory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryCloudAdvisory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryCloudAdvisory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryCloudAdvisory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryCloudAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryCloudAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryCloudAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryCloudAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVendor

`func (o *AdvisoryCloudAdvisory) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AdvisoryCloudAdvisory) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AdvisoryCloudAdvisory) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AdvisoryCloudAdvisory) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


