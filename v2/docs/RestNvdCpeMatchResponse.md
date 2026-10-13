# RestNvdCpeMatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** |  | 
**MatchStrings** | [**[]NvdcpematchNvdMatchStringEnvelopeEntry**](NvdcpematchNvdMatchStringEnvelopeEntry.md) |  | 
**NextCursor** | Pointer to **string** | Present only when the request sent &#x60;cursor&#x60; and a further page exists. | [optional] 
**ResultsPerPage** | **int32** |  | 
**StartIndex** | **int32** |  | 
**Timestamp** | **string** |  | 
**TotalResults** | **int32** |  | 
**Version** | **string** |  | 

## Methods

### NewRestNvdCpeMatchResponse

`func NewRestNvdCpeMatchResponse(format string, matchStrings []NvdcpematchNvdMatchStringEnvelopeEntry, resultsPerPage int32, startIndex int32, timestamp string, totalResults int32, version string, ) *RestNvdCpeMatchResponse`

NewRestNvdCpeMatchResponse instantiates a new RestNvdCpeMatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestNvdCpeMatchResponseWithDefaults

`func NewRestNvdCpeMatchResponseWithDefaults() *RestNvdCpeMatchResponse`

NewRestNvdCpeMatchResponseWithDefaults instantiates a new RestNvdCpeMatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *RestNvdCpeMatchResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RestNvdCpeMatchResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RestNvdCpeMatchResponse) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetMatchStrings

`func (o *RestNvdCpeMatchResponse) GetMatchStrings() []NvdcpematchNvdMatchStringEnvelopeEntry`

GetMatchStrings returns the MatchStrings field if non-nil, zero value otherwise.

### GetMatchStringsOk

`func (o *RestNvdCpeMatchResponse) GetMatchStringsOk() (*[]NvdcpematchNvdMatchStringEnvelopeEntry, bool)`

GetMatchStringsOk returns a tuple with the MatchStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchStrings

`func (o *RestNvdCpeMatchResponse) SetMatchStrings(v []NvdcpematchNvdMatchStringEnvelopeEntry)`

SetMatchStrings sets MatchStrings field to given value.


### GetNextCursor

`func (o *RestNvdCpeMatchResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *RestNvdCpeMatchResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *RestNvdCpeMatchResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *RestNvdCpeMatchResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetResultsPerPage

`func (o *RestNvdCpeMatchResponse) GetResultsPerPage() int32`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *RestNvdCpeMatchResponse) GetResultsPerPageOk() (*int32, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *RestNvdCpeMatchResponse) SetResultsPerPage(v int32)`

SetResultsPerPage sets ResultsPerPage field to given value.


### GetStartIndex

`func (o *RestNvdCpeMatchResponse) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *RestNvdCpeMatchResponse) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *RestNvdCpeMatchResponse) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.


### GetTimestamp

`func (o *RestNvdCpeMatchResponse) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RestNvdCpeMatchResponse) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RestNvdCpeMatchResponse) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetTotalResults

`func (o *RestNvdCpeMatchResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *RestNvdCpeMatchResponse) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *RestNvdCpeMatchResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.


### GetVersion

`func (o *RestNvdCpeMatchResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RestNvdCpeMatchResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RestNvdCpeMatchResponse) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


