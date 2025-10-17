# ApiEPSSData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **string** |  | [optional] 
**EpssPercentile** | Pointer to **float32** |  | [optional] 
**EpssScore** | Pointer to **float32** |  | [optional] 

## Methods

### NewApiEPSSData

`func NewApiEPSSData() *ApiEPSSData`

NewApiEPSSData instantiates a new ApiEPSSData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiEPSSDataWithDefaults

`func NewApiEPSSDataWithDefaults() *ApiEPSSData`

NewApiEPSSDataWithDefaults instantiates a new ApiEPSSData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ApiEPSSData) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApiEPSSData) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApiEPSSData) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ApiEPSSData) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCve

`func (o *ApiEPSSData) GetCve() string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *ApiEPSSData) GetCveOk() (*string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *ApiEPSSData) SetCve(v string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *ApiEPSSData) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetEpssPercentile

`func (o *ApiEPSSData) GetEpssPercentile() float32`

GetEpssPercentile returns the EpssPercentile field if non-nil, zero value otherwise.

### GetEpssPercentileOk

`func (o *ApiEPSSData) GetEpssPercentileOk() (*float32, bool)`

GetEpssPercentileOk returns a tuple with the EpssPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpssPercentile

`func (o *ApiEPSSData) SetEpssPercentile(v float32)`

SetEpssPercentile sets EpssPercentile field to given value.

### HasEpssPercentile

`func (o *ApiEPSSData) HasEpssPercentile() bool`

HasEpssPercentile returns a boolean if a field has been set.

### GetEpssScore

`func (o *ApiEPSSData) GetEpssScore() float32`

GetEpssScore returns the EpssScore field if non-nil, zero value otherwise.

### GetEpssScoreOk

`func (o *ApiEPSSData) GetEpssScoreOk() (*float32, bool)`

GetEpssScoreOk returns a tuple with the EpssScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpssScore

`func (o *ApiEPSSData) SetEpssScore(v float32)`

SetEpssScore sets EpssScore field to given value.

### HasEpssScore

`func (o *ApiEPSSData) HasEpssScore() bool`

HasEpssScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


