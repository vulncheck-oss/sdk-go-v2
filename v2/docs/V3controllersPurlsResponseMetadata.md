# V3controllersPurlsResponseMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** | time of the transaction | [optional] 
**TotalDocuments** | Pointer to **int32** | number of results found | [optional] 
**TotalSubmitted** | Pointer to **int32** | number of purls in the request | [optional] 
**Unprocessed** | Pointer to [**[]PurlUnprocessedPurl**](PurlUnprocessedPurl.md) | Unprocessed lists purls we could not look up. Not inferable from the counts above: purls with no vulnerabilities are omitted from data too. | [optional] 

## Methods

### NewV3controllersPurlsResponseMetadata

`func NewV3controllersPurlsResponseMetadata() *V3controllersPurlsResponseMetadata`

NewV3controllersPurlsResponseMetadata instantiates a new V3controllersPurlsResponseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3controllersPurlsResponseMetadataWithDefaults

`func NewV3controllersPurlsResponseMetadataWithDefaults() *V3controllersPurlsResponseMetadata`

NewV3controllersPurlsResponseMetadataWithDefaults instantiates a new V3controllersPurlsResponseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *V3controllersPurlsResponseMetadata) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *V3controllersPurlsResponseMetadata) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *V3controllersPurlsResponseMetadata) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *V3controllersPurlsResponseMetadata) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTotalDocuments

`func (o *V3controllersPurlsResponseMetadata) GetTotalDocuments() int32`

GetTotalDocuments returns the TotalDocuments field if non-nil, zero value otherwise.

### GetTotalDocumentsOk

`func (o *V3controllersPurlsResponseMetadata) GetTotalDocumentsOk() (*int32, bool)`

GetTotalDocumentsOk returns a tuple with the TotalDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDocuments

`func (o *V3controllersPurlsResponseMetadata) SetTotalDocuments(v int32)`

SetTotalDocuments sets TotalDocuments field to given value.

### HasTotalDocuments

`func (o *V3controllersPurlsResponseMetadata) HasTotalDocuments() bool`

HasTotalDocuments returns a boolean if a field has been set.

### GetTotalSubmitted

`func (o *V3controllersPurlsResponseMetadata) GetTotalSubmitted() int32`

GetTotalSubmitted returns the TotalSubmitted field if non-nil, zero value otherwise.

### GetTotalSubmittedOk

`func (o *V3controllersPurlsResponseMetadata) GetTotalSubmittedOk() (*int32, bool)`

GetTotalSubmittedOk returns a tuple with the TotalSubmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSubmitted

`func (o *V3controllersPurlsResponseMetadata) SetTotalSubmitted(v int32)`

SetTotalSubmitted sets TotalSubmitted field to given value.

### HasTotalSubmitted

`func (o *V3controllersPurlsResponseMetadata) HasTotalSubmitted() bool`

HasTotalSubmitted returns a boolean if a field has been set.

### GetUnprocessed

`func (o *V3controllersPurlsResponseMetadata) GetUnprocessed() []PurlUnprocessedPurl`

GetUnprocessed returns the Unprocessed field if non-nil, zero value otherwise.

### GetUnprocessedOk

`func (o *V3controllersPurlsResponseMetadata) GetUnprocessedOk() (*[]PurlUnprocessedPurl, bool)`

GetUnprocessedOk returns a tuple with the Unprocessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprocessed

`func (o *V3controllersPurlsResponseMetadata) SetUnprocessed(v []PurlUnprocessedPurl)`

SetUnprocessed sets Unprocessed field to given value.

### HasUnprocessed

`func (o *V3controllersPurlsResponseMetadata) HasUnprocessed() bool`

HasUnprocessed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


