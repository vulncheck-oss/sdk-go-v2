# IndexIndexCountEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocCount** | Pointer to **int32** | DocCount is the number of matching documents in this index. | [optional] 
**Index** | Pointer to **string** | Index is the name of the index. | [optional] 

## Methods

### NewIndexIndexCountEntry

`func NewIndexIndexCountEntry() *IndexIndexCountEntry`

NewIndexIndexCountEntry instantiates a new IndexIndexCountEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexIndexCountEntryWithDefaults

`func NewIndexIndexCountEntryWithDefaults() *IndexIndexCountEntry`

NewIndexIndexCountEntryWithDefaults instantiates a new IndexIndexCountEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocCount

`func (o *IndexIndexCountEntry) GetDocCount() int32`

GetDocCount returns the DocCount field if non-nil, zero value otherwise.

### GetDocCountOk

`func (o *IndexIndexCountEntry) GetDocCountOk() (*int32, bool)`

GetDocCountOk returns a tuple with the DocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocCount

`func (o *IndexIndexCountEntry) SetDocCount(v int32)`

SetDocCount sets DocCount field to given value.

### HasDocCount

`func (o *IndexIndexCountEntry) HasDocCount() bool`

HasDocCount returns a boolean if a field has been set.

### GetIndex

`func (o *IndexIndexCountEntry) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *IndexIndexCountEntry) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *IndexIndexCountEntry) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *IndexIndexCountEntry) HasIndex() bool`

HasIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


