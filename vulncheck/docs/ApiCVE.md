# ApiCVE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CVEDataMeta** | Pointer to [**ApiCVEDataMeta**](ApiCVEDataMeta.md) |  | [optional] 
**DataFormat** | Pointer to **string** |  | [optional] 
**DataType** | Pointer to **string** |  | [optional] 
**DataVersion** | Pointer to **string** |  | [optional] 
**Description** | Pointer to [**ApiDescription**](ApiDescription.md) |  | [optional] 
**Problemtype** | Pointer to [**ApiProblemType**](ApiProblemType.md) |  | [optional] 
**References** | Pointer to [**ApiReferences**](ApiReferences.md) |  | [optional] 

## Methods

### NewApiCVE

`func NewApiCVE() *ApiCVE`

NewApiCVE instantiates a new ApiCVE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCVEWithDefaults

`func NewApiCVEWithDefaults() *ApiCVE`

NewApiCVEWithDefaults instantiates a new ApiCVE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCVEDataMeta

`func (o *ApiCVE) GetCVEDataMeta() ApiCVEDataMeta`

GetCVEDataMeta returns the CVEDataMeta field if non-nil, zero value otherwise.

### GetCVEDataMetaOk

`func (o *ApiCVE) GetCVEDataMetaOk() (*ApiCVEDataMeta, bool)`

GetCVEDataMetaOk returns a tuple with the CVEDataMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCVEDataMeta

`func (o *ApiCVE) SetCVEDataMeta(v ApiCVEDataMeta)`

SetCVEDataMeta sets CVEDataMeta field to given value.

### HasCVEDataMeta

`func (o *ApiCVE) HasCVEDataMeta() bool`

HasCVEDataMeta returns a boolean if a field has been set.

### GetDataFormat

`func (o *ApiCVE) GetDataFormat() string`

GetDataFormat returns the DataFormat field if non-nil, zero value otherwise.

### GetDataFormatOk

`func (o *ApiCVE) GetDataFormatOk() (*string, bool)`

GetDataFormatOk returns a tuple with the DataFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFormat

`func (o *ApiCVE) SetDataFormat(v string)`

SetDataFormat sets DataFormat field to given value.

### HasDataFormat

`func (o *ApiCVE) HasDataFormat() bool`

HasDataFormat returns a boolean if a field has been set.

### GetDataType

`func (o *ApiCVE) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ApiCVE) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ApiCVE) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *ApiCVE) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDataVersion

`func (o *ApiCVE) GetDataVersion() string`

GetDataVersion returns the DataVersion field if non-nil, zero value otherwise.

### GetDataVersionOk

`func (o *ApiCVE) GetDataVersionOk() (*string, bool)`

GetDataVersionOk returns a tuple with the DataVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVersion

`func (o *ApiCVE) SetDataVersion(v string)`

SetDataVersion sets DataVersion field to given value.

### HasDataVersion

`func (o *ApiCVE) HasDataVersion() bool`

HasDataVersion returns a boolean if a field has been set.

### GetDescription

`func (o *ApiCVE) GetDescription() ApiDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiCVE) GetDescriptionOk() (*ApiDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiCVE) SetDescription(v ApiDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiCVE) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProblemtype

`func (o *ApiCVE) GetProblemtype() ApiProblemType`

GetProblemtype returns the Problemtype field if non-nil, zero value otherwise.

### GetProblemtypeOk

`func (o *ApiCVE) GetProblemtypeOk() (*ApiProblemType, bool)`

GetProblemtypeOk returns a tuple with the Problemtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemtype

`func (o *ApiCVE) SetProblemtype(v ApiProblemType)`

SetProblemtype sets Problemtype field to given value.

### HasProblemtype

`func (o *ApiCVE) HasProblemtype() bool`

HasProblemtype returns a boolean if a field has been set.

### GetReferences

`func (o *ApiCVE) GetReferences() ApiReferences`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ApiCVE) GetReferencesOk() (*ApiReferences, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ApiCVE) SetReferences(v ApiReferences)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ApiCVE) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


