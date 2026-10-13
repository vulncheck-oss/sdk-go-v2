# RestNvdCvesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** |  | 
**NextCursor** | Pointer to **string** | Present only when the request sent &#x60;cursor&#x60; and a further page exists. | [optional] 
**ResultsPerPage** | **int32** |  | 
**StartIndex** | **int32** |  | 
**Timestamp** | **string** |  | 
**TotalResults** | **int32** |  | 
**Version** | **string** |  | 
**Vulnerabilities** | [**[]NvdcvesNvdCveEnvelopeEntry**](NvdcvesNvdCveEnvelopeEntry.md) |  | 

## Methods

### NewRestNvdCvesResponse

`func NewRestNvdCvesResponse(format string, resultsPerPage int32, startIndex int32, timestamp string, totalResults int32, version string, vulnerabilities []NvdcvesNvdCveEnvelopeEntry, ) *RestNvdCvesResponse`

NewRestNvdCvesResponse instantiates a new RestNvdCvesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestNvdCvesResponseWithDefaults

`func NewRestNvdCvesResponseWithDefaults() *RestNvdCvesResponse`

NewRestNvdCvesResponseWithDefaults instantiates a new RestNvdCvesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *RestNvdCvesResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RestNvdCvesResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RestNvdCvesResponse) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetNextCursor

`func (o *RestNvdCvesResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *RestNvdCvesResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *RestNvdCvesResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *RestNvdCvesResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetResultsPerPage

`func (o *RestNvdCvesResponse) GetResultsPerPage() int32`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *RestNvdCvesResponse) GetResultsPerPageOk() (*int32, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *RestNvdCvesResponse) SetResultsPerPage(v int32)`

SetResultsPerPage sets ResultsPerPage field to given value.


### GetStartIndex

`func (o *RestNvdCvesResponse) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *RestNvdCvesResponse) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *RestNvdCvesResponse) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.


### GetTimestamp

`func (o *RestNvdCvesResponse) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RestNvdCvesResponse) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RestNvdCvesResponse) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetTotalResults

`func (o *RestNvdCvesResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *RestNvdCvesResponse) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *RestNvdCvesResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.


### GetVersion

`func (o *RestNvdCvesResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RestNvdCvesResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RestNvdCvesResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVulnerabilities

`func (o *RestNvdCvesResponse) GetVulnerabilities() []NvdcvesNvdCveEnvelopeEntry`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *RestNvdCvesResponse) GetVulnerabilitiesOk() (*[]NvdcvesNvdCveEnvelopeEntry, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *RestNvdCvesResponse) SetVulnerabilities(v []NvdcvesNvdCveEnvelopeEntry)`

SetVulnerabilities sets Vulnerabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


