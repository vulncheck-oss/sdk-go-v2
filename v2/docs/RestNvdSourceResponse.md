# RestNvdSourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** |  | 
**ResultsPerPage** | **int32** |  | 
**Sources** | [**[]NvdsourceNvdSource**](NvdsourceNvdSource.md) |  | 
**StartIndex** | **int32** |  | 
**Timestamp** | **string** |  | 
**TotalResults** | **int32** |  | 
**Version** | **string** |  | 

## Methods

### NewRestNvdSourceResponse

`func NewRestNvdSourceResponse(format string, resultsPerPage int32, sources []NvdsourceNvdSource, startIndex int32, timestamp string, totalResults int32, version string, ) *RestNvdSourceResponse`

NewRestNvdSourceResponse instantiates a new RestNvdSourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestNvdSourceResponseWithDefaults

`func NewRestNvdSourceResponseWithDefaults() *RestNvdSourceResponse`

NewRestNvdSourceResponseWithDefaults instantiates a new RestNvdSourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *RestNvdSourceResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RestNvdSourceResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RestNvdSourceResponse) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetResultsPerPage

`func (o *RestNvdSourceResponse) GetResultsPerPage() int32`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *RestNvdSourceResponse) GetResultsPerPageOk() (*int32, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *RestNvdSourceResponse) SetResultsPerPage(v int32)`

SetResultsPerPage sets ResultsPerPage field to given value.


### GetSources

`func (o *RestNvdSourceResponse) GetSources() []NvdsourceNvdSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *RestNvdSourceResponse) GetSourcesOk() (*[]NvdsourceNvdSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *RestNvdSourceResponse) SetSources(v []NvdsourceNvdSource)`

SetSources sets Sources field to given value.


### GetStartIndex

`func (o *RestNvdSourceResponse) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *RestNvdSourceResponse) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *RestNvdSourceResponse) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.


### GetTimestamp

`func (o *RestNvdSourceResponse) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RestNvdSourceResponse) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RestNvdSourceResponse) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetTotalResults

`func (o *RestNvdSourceResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *RestNvdSourceResponse) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *RestNvdSourceResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.


### GetVersion

`func (o *RestNvdSourceResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RestNvdSourceResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RestNvdSourceResponse) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


