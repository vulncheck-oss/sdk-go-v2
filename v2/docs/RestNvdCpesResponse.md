# RestNvdCpesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** |  | 
**NextCursor** | Pointer to **string** | Present only when the request sent &#x60;cursor&#x60; and a further page exists. | [optional] 
**Products** | [**[]NvdcpesNvdCpeEnvelopeEntry**](NvdcpesNvdCpeEnvelopeEntry.md) |  | 
**ResultsPerPage** | **int32** |  | 
**StartIndex** | **int32** |  | 
**Timestamp** | **string** |  | 
**TotalResults** | **int32** |  | 
**Version** | **string** |  | 

## Methods

### NewRestNvdCpesResponse

`func NewRestNvdCpesResponse(format string, products []NvdcpesNvdCpeEnvelopeEntry, resultsPerPage int32, startIndex int32, timestamp string, totalResults int32, version string, ) *RestNvdCpesResponse`

NewRestNvdCpesResponse instantiates a new RestNvdCpesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestNvdCpesResponseWithDefaults

`func NewRestNvdCpesResponseWithDefaults() *RestNvdCpesResponse`

NewRestNvdCpesResponseWithDefaults instantiates a new RestNvdCpesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *RestNvdCpesResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RestNvdCpesResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RestNvdCpesResponse) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetNextCursor

`func (o *RestNvdCpesResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *RestNvdCpesResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *RestNvdCpesResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *RestNvdCpesResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetProducts

`func (o *RestNvdCpesResponse) GetProducts() []NvdcpesNvdCpeEnvelopeEntry`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *RestNvdCpesResponse) GetProductsOk() (*[]NvdcpesNvdCpeEnvelopeEntry, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *RestNvdCpesResponse) SetProducts(v []NvdcpesNvdCpeEnvelopeEntry)`

SetProducts sets Products field to given value.


### GetResultsPerPage

`func (o *RestNvdCpesResponse) GetResultsPerPage() int32`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *RestNvdCpesResponse) GetResultsPerPageOk() (*int32, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *RestNvdCpesResponse) SetResultsPerPage(v int32)`

SetResultsPerPage sets ResultsPerPage field to given value.


### GetStartIndex

`func (o *RestNvdCpesResponse) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *RestNvdCpesResponse) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *RestNvdCpesResponse) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.


### GetTimestamp

`func (o *RestNvdCpesResponse) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RestNvdCpesResponse) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RestNvdCpesResponse) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetTotalResults

`func (o *RestNvdCpesResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *RestNvdCpesResponse) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *RestNvdCpesResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.


### GetVersion

`func (o *RestNvdCpesResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RestNvdCpesResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RestNvdCpesResponse) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


