# SearchV4AdvisoryReturnValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**SearchV4AdvisoryMeta**](SearchV4AdvisoryMeta.md) |  | [optional] 
**Data** | Pointer to [**[]AdvisoryMitreCVEListV5Ref**](AdvisoryMitreCVEListV5Ref.md) |  | [optional] 

## Methods

### NewSearchV4AdvisoryReturnValue

`func NewSearchV4AdvisoryReturnValue() *SearchV4AdvisoryReturnValue`

NewSearchV4AdvisoryReturnValue instantiates a new SearchV4AdvisoryReturnValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchV4AdvisoryReturnValueWithDefaults

`func NewSearchV4AdvisoryReturnValueWithDefaults() *SearchV4AdvisoryReturnValue`

NewSearchV4AdvisoryReturnValueWithDefaults instantiates a new SearchV4AdvisoryReturnValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *SearchV4AdvisoryReturnValue) GetMeta() SearchV4AdvisoryMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SearchV4AdvisoryReturnValue) GetMetaOk() (*SearchV4AdvisoryMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SearchV4AdvisoryReturnValue) SetMeta(v SearchV4AdvisoryMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SearchV4AdvisoryReturnValue) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *SearchV4AdvisoryReturnValue) GetData() []AdvisoryMitreCVEListV5Ref`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SearchV4AdvisoryReturnValue) GetDataOk() (*[]AdvisoryMitreCVEListV5Ref, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SearchV4AdvisoryReturnValue) SetData(v []AdvisoryMitreCVEListV5Ref)`

SetData sets Data field to given value.

### HasData

`func (o *SearchV4AdvisoryReturnValue) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


