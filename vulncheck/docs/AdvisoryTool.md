# AdvisoryTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**References** | Pointer to [**[]AdvisoryToolRef**](AdvisoryToolRef.md) |  | [optional] 

## Methods

### NewAdvisoryTool

`func NewAdvisoryTool() *AdvisoryTool`

NewAdvisoryTool instantiates a new AdvisoryTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryToolWithDefaults

`func NewAdvisoryToolWithDefaults() *AdvisoryTool`

NewAdvisoryToolWithDefaults instantiates a new AdvisoryTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdvisoryTool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvisoryTool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvisoryTool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvisoryTool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryTool) GetReferences() []AdvisoryToolRef`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryTool) GetReferencesOk() (*[]AdvisoryToolRef, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryTool) SetReferences(v []AdvisoryToolRef)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryTool) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


