# AdvisoryMISPValueNoID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**AdvisoryMispMeta**](AdvisoryMispMeta.md) |  | [optional] 
**Related** | Pointer to [**[]AdvisoryMispRelatedItem**](AdvisoryMispRelatedItem.md) |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryMISPValueNoID

`func NewAdvisoryMISPValueNoID() *AdvisoryMISPValueNoID`

NewAdvisoryMISPValueNoID instantiates a new AdvisoryMISPValueNoID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMISPValueNoIDWithDefaults

`func NewAdvisoryMISPValueNoIDWithDefaults() *AdvisoryMISPValueNoID`

NewAdvisoryMISPValueNoIDWithDefaults instantiates a new AdvisoryMISPValueNoID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AdvisoryMISPValueNoID) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryMISPValueNoID) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryMISPValueNoID) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryMISPValueNoID) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *AdvisoryMISPValueNoID) GetMeta() AdvisoryMispMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AdvisoryMISPValueNoID) GetMetaOk() (*AdvisoryMispMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AdvisoryMISPValueNoID) SetMeta(v AdvisoryMispMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AdvisoryMISPValueNoID) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRelated

`func (o *AdvisoryMISPValueNoID) GetRelated() []AdvisoryMispRelatedItem`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *AdvisoryMISPValueNoID) GetRelatedOk() (*[]AdvisoryMispRelatedItem, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *AdvisoryMISPValueNoID) SetRelated(v []AdvisoryMispRelatedItem)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *AdvisoryMISPValueNoID) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetValue

`func (o *AdvisoryMISPValueNoID) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AdvisoryMISPValueNoID) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AdvisoryMISPValueNoID) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AdvisoryMISPValueNoID) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


