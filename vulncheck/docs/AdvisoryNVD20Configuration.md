# AdvisoryNVD20Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Negate** | Pointer to **bool** |  | [optional] 
**Nodes** | Pointer to [**[]AdvisoryNVD20Node**](AdvisoryNVD20Node.md) |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryNVD20Configuration

`func NewAdvisoryNVD20Configuration() *AdvisoryNVD20Configuration`

NewAdvisoryNVD20Configuration instantiates a new AdvisoryNVD20Configuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryNVD20ConfigurationWithDefaults

`func NewAdvisoryNVD20ConfigurationWithDefaults() *AdvisoryNVD20Configuration`

NewAdvisoryNVD20ConfigurationWithDefaults instantiates a new AdvisoryNVD20Configuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNegate

`func (o *AdvisoryNVD20Configuration) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *AdvisoryNVD20Configuration) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *AdvisoryNVD20Configuration) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *AdvisoryNVD20Configuration) HasNegate() bool`

HasNegate returns a boolean if a field has been set.

### GetNodes

`func (o *AdvisoryNVD20Configuration) GetNodes() []AdvisoryNVD20Node`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *AdvisoryNVD20Configuration) GetNodesOk() (*[]AdvisoryNVD20Node, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *AdvisoryNVD20Configuration) SetNodes(v []AdvisoryNVD20Node)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *AdvisoryNVD20Configuration) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetOperator

`func (o *AdvisoryNVD20Configuration) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AdvisoryNVD20Configuration) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AdvisoryNVD20Configuration) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AdvisoryNVD20Configuration) HasOperator() bool`

HasOperator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


