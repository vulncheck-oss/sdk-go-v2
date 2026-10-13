# RestNvdCveHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CveChanges** | [**[]RestNvdChangeEnvelopeEntry**](RestNvdChangeEnvelopeEntry.md) |  | 
**Format** | **string** |  | 
**NextCursor** | Pointer to **string** | Present only when the request sent &#x60;cursor&#x60; and a further page exists. | [optional] 
**ResultsPerPage** | **int32** |  | 
**StartIndex** | **int32** |  | 
**Timestamp** | **string** |  | 
**TotalResults** | **int32** |  | 
**Version** | **string** |  | 

## Methods

### NewRestNvdCveHistoryResponse

`func NewRestNvdCveHistoryResponse(cveChanges []RestNvdChangeEnvelopeEntry, format string, resultsPerPage int32, startIndex int32, timestamp string, totalResults int32, version string, ) *RestNvdCveHistoryResponse`

NewRestNvdCveHistoryResponse instantiates a new RestNvdCveHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestNvdCveHistoryResponseWithDefaults

`func NewRestNvdCveHistoryResponseWithDefaults() *RestNvdCveHistoryResponse`

NewRestNvdCveHistoryResponseWithDefaults instantiates a new RestNvdCveHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCveChanges

`func (o *RestNvdCveHistoryResponse) GetCveChanges() []RestNvdChangeEnvelopeEntry`

GetCveChanges returns the CveChanges field if non-nil, zero value otherwise.

### GetCveChangesOk

`func (o *RestNvdCveHistoryResponse) GetCveChangesOk() (*[]RestNvdChangeEnvelopeEntry, bool)`

GetCveChangesOk returns a tuple with the CveChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveChanges

`func (o *RestNvdCveHistoryResponse) SetCveChanges(v []RestNvdChangeEnvelopeEntry)`

SetCveChanges sets CveChanges field to given value.


### GetFormat

`func (o *RestNvdCveHistoryResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RestNvdCveHistoryResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RestNvdCveHistoryResponse) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetNextCursor

`func (o *RestNvdCveHistoryResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *RestNvdCveHistoryResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *RestNvdCveHistoryResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *RestNvdCveHistoryResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetResultsPerPage

`func (o *RestNvdCveHistoryResponse) GetResultsPerPage() int32`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *RestNvdCveHistoryResponse) GetResultsPerPageOk() (*int32, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *RestNvdCveHistoryResponse) SetResultsPerPage(v int32)`

SetResultsPerPage sets ResultsPerPage field to given value.


### GetStartIndex

`func (o *RestNvdCveHistoryResponse) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *RestNvdCveHistoryResponse) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *RestNvdCveHistoryResponse) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.


### GetTimestamp

`func (o *RestNvdCveHistoryResponse) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RestNvdCveHistoryResponse) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RestNvdCveHistoryResponse) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetTotalResults

`func (o *RestNvdCveHistoryResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *RestNvdCveHistoryResponse) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *RestNvdCveHistoryResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.


### GetVersion

`func (o *RestNvdCveHistoryResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RestNvdCveHistoryResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RestNvdCveHistoryResponse) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


