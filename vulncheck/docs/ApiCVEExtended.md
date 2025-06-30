# ApiCVEExtended

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CVEDataMeta** | Pointer to [**ApiCVEDataMetaExtended**](ApiCVEDataMetaExtended.md) |  | [optional] 
**Categorization** | Pointer to [**ApiCategorizationExtended**](ApiCategorizationExtended.md) |  | [optional] 
**DataFormat** | Pointer to **string** |  | [optional] 
**DataType** | Pointer to **string** |  | [optional] 
**DataVersion** | Pointer to **string** |  | [optional] 
**Description** | Pointer to [**ApiDescription**](ApiDescription.md) |  | [optional] 
**Problemtype** | Pointer to [**ApiProblemTypeExtended**](ApiProblemTypeExtended.md) |  | [optional] 
**References** | Pointer to [**ApiReferencesExtended**](ApiReferencesExtended.md) |  | [optional] 

## Methods

### NewApiCVEExtended

`func NewApiCVEExtended() *ApiCVEExtended`

NewApiCVEExtended instantiates a new ApiCVEExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCVEExtendedWithDefaults

`func NewApiCVEExtendedWithDefaults() *ApiCVEExtended`

NewApiCVEExtendedWithDefaults instantiates a new ApiCVEExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCVEDataMeta

`func (o *ApiCVEExtended) GetCVEDataMeta() ApiCVEDataMetaExtended`

GetCVEDataMeta returns the CVEDataMeta field if non-nil, zero value otherwise.

### GetCVEDataMetaOk

`func (o *ApiCVEExtended) GetCVEDataMetaOk() (*ApiCVEDataMetaExtended, bool)`

GetCVEDataMetaOk returns a tuple with the CVEDataMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCVEDataMeta

`func (o *ApiCVEExtended) SetCVEDataMeta(v ApiCVEDataMetaExtended)`

SetCVEDataMeta sets CVEDataMeta field to given value.

### HasCVEDataMeta

`func (o *ApiCVEExtended) HasCVEDataMeta() bool`

HasCVEDataMeta returns a boolean if a field has been set.

### GetCategorization

`func (o *ApiCVEExtended) GetCategorization() ApiCategorizationExtended`

GetCategorization returns the Categorization field if non-nil, zero value otherwise.

### GetCategorizationOk

`func (o *ApiCVEExtended) GetCategorizationOk() (*ApiCategorizationExtended, bool)`

GetCategorizationOk returns a tuple with the Categorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorization

`func (o *ApiCVEExtended) SetCategorization(v ApiCategorizationExtended)`

SetCategorization sets Categorization field to given value.

### HasCategorization

`func (o *ApiCVEExtended) HasCategorization() bool`

HasCategorization returns a boolean if a field has been set.

### GetDataFormat

`func (o *ApiCVEExtended) GetDataFormat() string`

GetDataFormat returns the DataFormat field if non-nil, zero value otherwise.

### GetDataFormatOk

`func (o *ApiCVEExtended) GetDataFormatOk() (*string, bool)`

GetDataFormatOk returns a tuple with the DataFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFormat

`func (o *ApiCVEExtended) SetDataFormat(v string)`

SetDataFormat sets DataFormat field to given value.

### HasDataFormat

`func (o *ApiCVEExtended) HasDataFormat() bool`

HasDataFormat returns a boolean if a field has been set.

### GetDataType

`func (o *ApiCVEExtended) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ApiCVEExtended) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ApiCVEExtended) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *ApiCVEExtended) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDataVersion

`func (o *ApiCVEExtended) GetDataVersion() string`

GetDataVersion returns the DataVersion field if non-nil, zero value otherwise.

### GetDataVersionOk

`func (o *ApiCVEExtended) GetDataVersionOk() (*string, bool)`

GetDataVersionOk returns a tuple with the DataVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVersion

`func (o *ApiCVEExtended) SetDataVersion(v string)`

SetDataVersion sets DataVersion field to given value.

### HasDataVersion

`func (o *ApiCVEExtended) HasDataVersion() bool`

HasDataVersion returns a boolean if a field has been set.

### GetDescription

`func (o *ApiCVEExtended) GetDescription() ApiDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiCVEExtended) GetDescriptionOk() (*ApiDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiCVEExtended) SetDescription(v ApiDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiCVEExtended) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProblemtype

`func (o *ApiCVEExtended) GetProblemtype() ApiProblemTypeExtended`

GetProblemtype returns the Problemtype field if non-nil, zero value otherwise.

### GetProblemtypeOk

`func (o *ApiCVEExtended) GetProblemtypeOk() (*ApiProblemTypeExtended, bool)`

GetProblemtypeOk returns a tuple with the Problemtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemtype

`func (o *ApiCVEExtended) SetProblemtype(v ApiProblemTypeExtended)`

SetProblemtype sets Problemtype field to given value.

### HasProblemtype

`func (o *ApiCVEExtended) HasProblemtype() bool`

HasProblemtype returns a boolean if a field has been set.

### GetReferences

`func (o *ApiCVEExtended) GetReferences() ApiReferencesExtended`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ApiCVEExtended) GetReferencesOk() (*ApiReferencesExtended, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ApiCVEExtended) SetReferences(v ApiReferencesExtended)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ApiCVEExtended) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


