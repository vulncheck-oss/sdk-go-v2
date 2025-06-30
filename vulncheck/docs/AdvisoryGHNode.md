# AdvisoryGHNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Package** | Pointer to [**AdvisoryGHPackage**](AdvisoryGHPackage.md) |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**VulnerableVersionRange** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryGHNode

`func NewAdvisoryGHNode() *AdvisoryGHNode`

NewAdvisoryGHNode instantiates a new AdvisoryGHNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryGHNodeWithDefaults

`func NewAdvisoryGHNodeWithDefaults() *AdvisoryGHNode`

NewAdvisoryGHNodeWithDefaults instantiates a new AdvisoryGHNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackage

`func (o *AdvisoryGHNode) GetPackage() AdvisoryGHPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *AdvisoryGHNode) GetPackageOk() (*AdvisoryGHPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *AdvisoryGHNode) SetPackage(v AdvisoryGHPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *AdvisoryGHNode) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryGHNode) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryGHNode) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryGHNode) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryGHNode) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryGHNode) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryGHNode) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryGHNode) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryGHNode) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVulnerableVersionRange

`func (o *AdvisoryGHNode) GetVulnerableVersionRange() string`

GetVulnerableVersionRange returns the VulnerableVersionRange field if non-nil, zero value otherwise.

### GetVulnerableVersionRangeOk

`func (o *AdvisoryGHNode) GetVulnerableVersionRangeOk() (*string, bool)`

GetVulnerableVersionRangeOk returns a tuple with the VulnerableVersionRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableVersionRange

`func (o *AdvisoryGHNode) SetVulnerableVersionRange(v string)`

SetVulnerableVersionRange sets VulnerableVersionRange field to given value.

### HasVulnerableVersionRange

`func (o *AdvisoryGHNode) HasVulnerableVersionRange() bool`

HasVulnerableVersionRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


