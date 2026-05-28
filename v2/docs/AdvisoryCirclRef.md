# AdvisoryCirclRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | Pointer to [**AdvisoryCirclContainers**](AdvisoryCirclContainers.md) |  | [optional] 
**CveMetadata** | Pointer to [**AdvisoryCirclCveMetadata**](AdvisoryCirclCveMetadata.md) |  | [optional] 
**DataType** | Pointer to **string** |  | [optional] 
**DataVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryCirclRef

`func NewAdvisoryCirclRef() *AdvisoryCirclRef`

NewAdvisoryCirclRef instantiates a new AdvisoryCirclRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCirclRefWithDefaults

`func NewAdvisoryCirclRefWithDefaults() *AdvisoryCirclRef`

NewAdvisoryCirclRefWithDefaults instantiates a new AdvisoryCirclRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *AdvisoryCirclRef) GetContainers() AdvisoryCirclContainers`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *AdvisoryCirclRef) GetContainersOk() (*AdvisoryCirclContainers, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *AdvisoryCirclRef) SetContainers(v AdvisoryCirclContainers)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *AdvisoryCirclRef) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetCveMetadata

`func (o *AdvisoryCirclRef) GetCveMetadata() AdvisoryCirclCveMetadata`

GetCveMetadata returns the CveMetadata field if non-nil, zero value otherwise.

### GetCveMetadataOk

`func (o *AdvisoryCirclRef) GetCveMetadataOk() (*AdvisoryCirclCveMetadata, bool)`

GetCveMetadataOk returns a tuple with the CveMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveMetadata

`func (o *AdvisoryCirclRef) SetCveMetadata(v AdvisoryCirclCveMetadata)`

SetCveMetadata sets CveMetadata field to given value.

### HasCveMetadata

`func (o *AdvisoryCirclRef) HasCveMetadata() bool`

HasCveMetadata returns a boolean if a field has been set.

### GetDataType

`func (o *AdvisoryCirclRef) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *AdvisoryCirclRef) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *AdvisoryCirclRef) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *AdvisoryCirclRef) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDataVersion

`func (o *AdvisoryCirclRef) GetDataVersion() string`

GetDataVersion returns the DataVersion field if non-nil, zero value otherwise.

### GetDataVersionOk

`func (o *AdvisoryCirclRef) GetDataVersionOk() (*string, bool)`

GetDataVersionOk returns a tuple with the DataVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVersion

`func (o *AdvisoryCirclRef) SetDataVersion(v string)`

SetDataVersion sets DataVersion field to given value.

### HasDataVersion

`func (o *AdvisoryCirclRef) HasDataVersion() bool`

HasDataVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


