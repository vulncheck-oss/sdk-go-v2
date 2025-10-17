# ApiNodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]ApiNodes**](ApiNodes.md) |  | [optional] 
**CpeMatch** | Pointer to [**[]ApiCPEMatch**](ApiCPEMatch.md) |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 

## Methods

### NewApiNodes

`func NewApiNodes() *ApiNodes`

NewApiNodes instantiates a new ApiNodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNodesWithDefaults

`func NewApiNodesWithDefaults() *ApiNodes`

NewApiNodesWithDefaults instantiates a new ApiNodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *ApiNodes) GetChildren() []ApiNodes`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ApiNodes) GetChildrenOk() (*[]ApiNodes, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ApiNodes) SetChildren(v []ApiNodes)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ApiNodes) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCpeMatch

`func (o *ApiNodes) GetCpeMatch() []ApiCPEMatch`

GetCpeMatch returns the CpeMatch field if non-nil, zero value otherwise.

### GetCpeMatchOk

`func (o *ApiNodes) GetCpeMatchOk() (*[]ApiCPEMatch, bool)`

GetCpeMatchOk returns a tuple with the CpeMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeMatch

`func (o *ApiNodes) SetCpeMatch(v []ApiCPEMatch)`

SetCpeMatch sets CpeMatch field to given value.

### HasCpeMatch

`func (o *ApiNodes) HasCpeMatch() bool`

HasCpeMatch returns a boolean if a field has been set.

### GetOperator

`func (o *ApiNodes) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ApiNodes) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ApiNodes) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ApiNodes) HasOperator() bool`

HasOperator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


