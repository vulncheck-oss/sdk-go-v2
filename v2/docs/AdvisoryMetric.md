# AdvisoryMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CvssV20** | Pointer to [**AdvisoryMCvssV20**](AdvisoryMCvssV20.md) |  | [optional] 
**CvssV30** | Pointer to [**AdvisoryMCvssV30**](AdvisoryMCvssV30.md) |  | [optional] 
**CvssV31** | Pointer to [**AdvisoryMCvssV31**](AdvisoryMCvssV31.md) |  | [optional] 
**CvssV40** | Pointer to [**AdvisoryMCvssV40**](AdvisoryMCvssV40.md) |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Other** | Pointer to [**AdvisoryMetricsOther**](AdvisoryMetricsOther.md) |  | [optional] 

## Methods

### NewAdvisoryMetric

`func NewAdvisoryMetric() *AdvisoryMetric`

NewAdvisoryMetric instantiates a new AdvisoryMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMetricWithDefaults

`func NewAdvisoryMetricWithDefaults() *AdvisoryMetric`

NewAdvisoryMetricWithDefaults instantiates a new AdvisoryMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvssV20

`func (o *AdvisoryMetric) GetCvssV20() AdvisoryMCvssV20`

GetCvssV20 returns the CvssV20 field if non-nil, zero value otherwise.

### GetCvssV20Ok

`func (o *AdvisoryMetric) GetCvssV20Ok() (*AdvisoryMCvssV20, bool)`

GetCvssV20Ok returns a tuple with the CvssV20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV20

`func (o *AdvisoryMetric) SetCvssV20(v AdvisoryMCvssV20)`

SetCvssV20 sets CvssV20 field to given value.

### HasCvssV20

`func (o *AdvisoryMetric) HasCvssV20() bool`

HasCvssV20 returns a boolean if a field has been set.

### GetCvssV30

`func (o *AdvisoryMetric) GetCvssV30() AdvisoryMCvssV30`

GetCvssV30 returns the CvssV30 field if non-nil, zero value otherwise.

### GetCvssV30Ok

`func (o *AdvisoryMetric) GetCvssV30Ok() (*AdvisoryMCvssV30, bool)`

GetCvssV30Ok returns a tuple with the CvssV30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV30

`func (o *AdvisoryMetric) SetCvssV30(v AdvisoryMCvssV30)`

SetCvssV30 sets CvssV30 field to given value.

### HasCvssV30

`func (o *AdvisoryMetric) HasCvssV30() bool`

HasCvssV30 returns a boolean if a field has been set.

### GetCvssV31

`func (o *AdvisoryMetric) GetCvssV31() AdvisoryMCvssV31`

GetCvssV31 returns the CvssV31 field if non-nil, zero value otherwise.

### GetCvssV31Ok

`func (o *AdvisoryMetric) GetCvssV31Ok() (*AdvisoryMCvssV31, bool)`

GetCvssV31Ok returns a tuple with the CvssV31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV31

`func (o *AdvisoryMetric) SetCvssV31(v AdvisoryMCvssV31)`

SetCvssV31 sets CvssV31 field to given value.

### HasCvssV31

`func (o *AdvisoryMetric) HasCvssV31() bool`

HasCvssV31 returns a boolean if a field has been set.

### GetCvssV40

`func (o *AdvisoryMetric) GetCvssV40() AdvisoryMCvssV40`

GetCvssV40 returns the CvssV40 field if non-nil, zero value otherwise.

### GetCvssV40Ok

`func (o *AdvisoryMetric) GetCvssV40Ok() (*AdvisoryMCvssV40, bool)`

GetCvssV40Ok returns a tuple with the CvssV40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV40

`func (o *AdvisoryMetric) SetCvssV40(v AdvisoryMCvssV40)`

SetCvssV40 sets CvssV40 field to given value.

### HasCvssV40

`func (o *AdvisoryMetric) HasCvssV40() bool`

HasCvssV40 returns a boolean if a field has been set.

### GetFormat

`func (o *AdvisoryMetric) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AdvisoryMetric) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AdvisoryMetric) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *AdvisoryMetric) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetOther

`func (o *AdvisoryMetric) GetOther() AdvisoryMetricsOther`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *AdvisoryMetric) GetOtherOk() (*AdvisoryMetricsOther, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *AdvisoryMetric) SetOther(v AdvisoryMetricsOther)`

SetOther sets Other field to given value.

### HasOther

`func (o *AdvisoryMetric) HasOther() bool`

HasOther returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


