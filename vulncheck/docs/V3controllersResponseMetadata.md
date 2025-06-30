# V3controllersResponseMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpe** | Pointer to **string** |  | [optional] 
**CpeStruct** | Pointer to [**ApiCPE**](ApiCPE.md) |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**TotalDocuments** | Pointer to **int32** |  | [optional] 

## Methods

### NewV3controllersResponseMetadata

`func NewV3controllersResponseMetadata() *V3controllersResponseMetadata`

NewV3controllersResponseMetadata instantiates a new V3controllersResponseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3controllersResponseMetadataWithDefaults

`func NewV3controllersResponseMetadataWithDefaults() *V3controllersResponseMetadata`

NewV3controllersResponseMetadataWithDefaults instantiates a new V3controllersResponseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpe

`func (o *V3controllersResponseMetadata) GetCpe() string`

GetCpe returns the Cpe field if non-nil, zero value otherwise.

### GetCpeOk

`func (o *V3controllersResponseMetadata) GetCpeOk() (*string, bool)`

GetCpeOk returns a tuple with the Cpe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpe

`func (o *V3controllersResponseMetadata) SetCpe(v string)`

SetCpe sets Cpe field to given value.

### HasCpe

`func (o *V3controllersResponseMetadata) HasCpe() bool`

HasCpe returns a boolean if a field has been set.

### GetCpeStruct

`func (o *V3controllersResponseMetadata) GetCpeStruct() ApiCPE`

GetCpeStruct returns the CpeStruct field if non-nil, zero value otherwise.

### GetCpeStructOk

`func (o *V3controllersResponseMetadata) GetCpeStructOk() (*ApiCPE, bool)`

GetCpeStructOk returns a tuple with the CpeStruct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeStruct

`func (o *V3controllersResponseMetadata) SetCpeStruct(v ApiCPE)`

SetCpeStruct sets CpeStruct field to given value.

### HasCpeStruct

`func (o *V3controllersResponseMetadata) HasCpeStruct() bool`

HasCpeStruct returns a boolean if a field has been set.

### GetTimestamp

`func (o *V3controllersResponseMetadata) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *V3controllersResponseMetadata) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *V3controllersResponseMetadata) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *V3controllersResponseMetadata) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTotalDocuments

`func (o *V3controllersResponseMetadata) GetTotalDocuments() int32`

GetTotalDocuments returns the TotalDocuments field if non-nil, zero value otherwise.

### GetTotalDocumentsOk

`func (o *V3controllersResponseMetadata) GetTotalDocumentsOk() (*int32, bool)`

GetTotalDocumentsOk returns a tuple with the TotalDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDocuments

`func (o *V3controllersResponseMetadata) SetTotalDocuments(v int32)`

SetTotalDocuments sets TotalDocuments field to given value.

### HasTotalDocuments

`func (o *V3controllersResponseMetadata) HasTotalDocuments() bool`

HasTotalDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


