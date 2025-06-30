# AdvisoryNVD20Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpeMatch** | Pointer to [**[]AdvisoryNVD20CVECPEMatch**](AdvisoryNVD20CVECPEMatch.md) |  | [optional] 
**Negate** | Pointer to **bool** |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryNVD20Node

`func NewAdvisoryNVD20Node() *AdvisoryNVD20Node`

NewAdvisoryNVD20Node instantiates a new AdvisoryNVD20Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryNVD20NodeWithDefaults

`func NewAdvisoryNVD20NodeWithDefaults() *AdvisoryNVD20Node`

NewAdvisoryNVD20NodeWithDefaults instantiates a new AdvisoryNVD20Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpeMatch

`func (o *AdvisoryNVD20Node) GetCpeMatch() []AdvisoryNVD20CVECPEMatch`

GetCpeMatch returns the CpeMatch field if non-nil, zero value otherwise.

### GetCpeMatchOk

`func (o *AdvisoryNVD20Node) GetCpeMatchOk() (*[]AdvisoryNVD20CVECPEMatch, bool)`

GetCpeMatchOk returns a tuple with the CpeMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeMatch

`func (o *AdvisoryNVD20Node) SetCpeMatch(v []AdvisoryNVD20CVECPEMatch)`

SetCpeMatch sets CpeMatch field to given value.

### HasCpeMatch

`func (o *AdvisoryNVD20Node) HasCpeMatch() bool`

HasCpeMatch returns a boolean if a field has been set.

### GetNegate

`func (o *AdvisoryNVD20Node) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *AdvisoryNVD20Node) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *AdvisoryNVD20Node) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *AdvisoryNVD20Node) HasNegate() bool`

HasNegate returns a boolean if a field has been set.

### GetOperator

`func (o *AdvisoryNVD20Node) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AdvisoryNVD20Node) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AdvisoryNVD20Node) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AdvisoryNVD20Node) HasOperator() bool`

HasOperator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


