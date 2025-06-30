# V3controllersPurlResponseMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PurlStruct** | Pointer to [**PurlPackageURLJSON**](PurlPackageURLJSON.md) | meta-data about the purl | [optional] 
**Timestamp** | Pointer to **string** | time of the transaction | [optional] 
**TotalDocuments** | Pointer to **int32** | number of results found | [optional] 

## Methods

### NewV3controllersPurlResponseMetadata

`func NewV3controllersPurlResponseMetadata() *V3controllersPurlResponseMetadata`

NewV3controllersPurlResponseMetadata instantiates a new V3controllersPurlResponseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3controllersPurlResponseMetadataWithDefaults

`func NewV3controllersPurlResponseMetadataWithDefaults() *V3controllersPurlResponseMetadata`

NewV3controllersPurlResponseMetadataWithDefaults instantiates a new V3controllersPurlResponseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurlStruct

`func (o *V3controllersPurlResponseMetadata) GetPurlStruct() PurlPackageURLJSON`

GetPurlStruct returns the PurlStruct field if non-nil, zero value otherwise.

### GetPurlStructOk

`func (o *V3controllersPurlResponseMetadata) GetPurlStructOk() (*PurlPackageURLJSON, bool)`

GetPurlStructOk returns a tuple with the PurlStruct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurlStruct

`func (o *V3controllersPurlResponseMetadata) SetPurlStruct(v PurlPackageURLJSON)`

SetPurlStruct sets PurlStruct field to given value.

### HasPurlStruct

`func (o *V3controllersPurlResponseMetadata) HasPurlStruct() bool`

HasPurlStruct returns a boolean if a field has been set.

### GetTimestamp

`func (o *V3controllersPurlResponseMetadata) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *V3controllersPurlResponseMetadata) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *V3controllersPurlResponseMetadata) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *V3controllersPurlResponseMetadata) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTotalDocuments

`func (o *V3controllersPurlResponseMetadata) GetTotalDocuments() int32`

GetTotalDocuments returns the TotalDocuments field if non-nil, zero value otherwise.

### GetTotalDocumentsOk

`func (o *V3controllersPurlResponseMetadata) GetTotalDocumentsOk() (*int32, bool)`

GetTotalDocumentsOk returns a tuple with the TotalDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDocuments

`func (o *V3controllersPurlResponseMetadata) SetTotalDocuments(v int32)`

SetTotalDocuments sets TotalDocuments field to given value.

### HasTotalDocuments

`func (o *V3controllersPurlResponseMetadata) HasTotalDocuments() bool`

HasTotalDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


