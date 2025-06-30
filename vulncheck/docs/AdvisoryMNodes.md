# AdvisoryMNodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpeMatch** | Pointer to [**[]AdvisoryMCPEMatch**](AdvisoryMCPEMatch.md) |  | [optional] 
**Negate** | Pointer to **bool** |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryMNodes

`func NewAdvisoryMNodes() *AdvisoryMNodes`

NewAdvisoryMNodes instantiates a new AdvisoryMNodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMNodesWithDefaults

`func NewAdvisoryMNodesWithDefaults() *AdvisoryMNodes`

NewAdvisoryMNodesWithDefaults instantiates a new AdvisoryMNodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpeMatch

`func (o *AdvisoryMNodes) GetCpeMatch() []AdvisoryMCPEMatch`

GetCpeMatch returns the CpeMatch field if non-nil, zero value otherwise.

### GetCpeMatchOk

`func (o *AdvisoryMNodes) GetCpeMatchOk() (*[]AdvisoryMCPEMatch, bool)`

GetCpeMatchOk returns a tuple with the CpeMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeMatch

`func (o *AdvisoryMNodes) SetCpeMatch(v []AdvisoryMCPEMatch)`

SetCpeMatch sets CpeMatch field to given value.

### HasCpeMatch

`func (o *AdvisoryMNodes) HasCpeMatch() bool`

HasCpeMatch returns a boolean if a field has been set.

### GetNegate

`func (o *AdvisoryMNodes) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *AdvisoryMNodes) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *AdvisoryMNodes) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *AdvisoryMNodes) HasNegate() bool`

HasNegate returns a boolean if a field has been set.

### GetOperator

`func (o *AdvisoryMNodes) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AdvisoryMNodes) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AdvisoryMNodes) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AdvisoryMNodes) HasOperator() bool`

HasOperator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


