# AdvisoryMispValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**AdvisoryMispMeta**](AdvisoryMispMeta.md) |  | [optional] 
**Related** | Pointer to [**[]AdvisoryMispRelatedItem**](AdvisoryMispRelatedItem.md) |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryMispValue

`func NewAdvisoryMispValue() *AdvisoryMispValue`

NewAdvisoryMispValue instantiates a new AdvisoryMispValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMispValueWithDefaults

`func NewAdvisoryMispValueWithDefaults() *AdvisoryMispValue`

NewAdvisoryMispValueWithDefaults instantiates a new AdvisoryMispValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AdvisoryMispValue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryMispValue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryMispValue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryMispValue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *AdvisoryMispValue) GetMeta() AdvisoryMispMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AdvisoryMispValue) GetMetaOk() (*AdvisoryMispMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AdvisoryMispValue) SetMeta(v AdvisoryMispMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AdvisoryMispValue) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRelated

`func (o *AdvisoryMispValue) GetRelated() []AdvisoryMispRelatedItem`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *AdvisoryMispValue) GetRelatedOk() (*[]AdvisoryMispRelatedItem, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *AdvisoryMispValue) SetRelated(v []AdvisoryMispRelatedItem)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *AdvisoryMispValue) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetUuid

`func (o *AdvisoryMispValue) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AdvisoryMispValue) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AdvisoryMispValue) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AdvisoryMispValue) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetValue

`func (o *AdvisoryMispValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AdvisoryMispValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AdvisoryMispValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AdvisoryMispValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


