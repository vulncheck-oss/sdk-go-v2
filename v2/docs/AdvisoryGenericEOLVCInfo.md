# AdvisoryGenericEOLVCInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**References** | Pointer to [**[]AdvisoryGenericEOLReference**](AdvisoryGenericEOLReference.md) |  | [optional] 
**Replacement** | Pointer to [**AdvisoryGenericEOLProduct**](AdvisoryGenericEOLProduct.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryGenericEOLVCInfo

`func NewAdvisoryGenericEOLVCInfo() *AdvisoryGenericEOLVCInfo`

NewAdvisoryGenericEOLVCInfo instantiates a new AdvisoryGenericEOLVCInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryGenericEOLVCInfoWithDefaults

`func NewAdvisoryGenericEOLVCInfoWithDefaults() *AdvisoryGenericEOLVCInfo`

NewAdvisoryGenericEOLVCInfoWithDefaults instantiates a new AdvisoryGenericEOLVCInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryGenericEOLVCInfo) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryGenericEOLVCInfo) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryGenericEOLVCInfo) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryGenericEOLVCInfo) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryGenericEOLVCInfo) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryGenericEOLVCInfo) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryGenericEOLVCInfo) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryGenericEOLVCInfo) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryGenericEOLVCInfo) GetReferences() []AdvisoryGenericEOLReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryGenericEOLVCInfo) GetReferencesOk() (*[]AdvisoryGenericEOLReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryGenericEOLVCInfo) SetReferences(v []AdvisoryGenericEOLReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryGenericEOLVCInfo) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetReplacement

`func (o *AdvisoryGenericEOLVCInfo) GetReplacement() AdvisoryGenericEOLProduct`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *AdvisoryGenericEOLVCInfo) GetReplacementOk() (*AdvisoryGenericEOLProduct, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *AdvisoryGenericEOLVCInfo) SetReplacement(v AdvisoryGenericEOLProduct)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *AdvisoryGenericEOLVCInfo) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryGenericEOLVCInfo) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryGenericEOLVCInfo) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryGenericEOLVCInfo) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryGenericEOLVCInfo) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *AdvisoryGenericEOLVCInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AdvisoryGenericEOLVCInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AdvisoryGenericEOLVCInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AdvisoryGenericEOLVCInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


