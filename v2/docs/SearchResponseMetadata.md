# SearchResponseMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** |  | [optional] 
**TotalDocuments** | Pointer to **int32** |  | [optional] 

## Methods

### NewSearchResponseMetadata

`func NewSearchResponseMetadata() *SearchResponseMetadata`

NewSearchResponseMetadata instantiates a new SearchResponseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseMetadataWithDefaults

`func NewSearchResponseMetadataWithDefaults() *SearchResponseMetadata`

NewSearchResponseMetadataWithDefaults instantiates a new SearchResponseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *SearchResponseMetadata) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SearchResponseMetadata) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SearchResponseMetadata) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SearchResponseMetadata) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTotalDocuments

`func (o *SearchResponseMetadata) GetTotalDocuments() int32`

GetTotalDocuments returns the TotalDocuments field if non-nil, zero value otherwise.

### GetTotalDocumentsOk

`func (o *SearchResponseMetadata) GetTotalDocumentsOk() (*int32, bool)`

GetTotalDocumentsOk returns a tuple with the TotalDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDocuments

`func (o *SearchResponseMetadata) SetTotalDocuments(v int32)`

SetTotalDocuments sets TotalDocuments field to given value.

### HasTotalDocuments

`func (o *SearchResponseMetadata) HasTotalDocuments() bool`

HasTotalDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


