# ApiNVD20MetricExtended

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CvssMetricV2** | Pointer to [**[]ApiNVD20CvssMetricV2**](ApiNVD20CvssMetricV2.md) |  | [optional] 
**CvssMetricV30** | Pointer to [**[]ApiNVD20CvssMetricV3**](ApiNVD20CvssMetricV3.md) |  | [optional] 
**CvssMetricV31** | Pointer to [**[]ApiNVD20CvssMetricV3**](ApiNVD20CvssMetricV3.md) |  | [optional] 
**CvssMetricV40** | Pointer to [**[]ApiNVD20CvssMetricV40**](ApiNVD20CvssMetricV40.md) |  | [optional] 
**Epss** | Pointer to [**ApiEPSS**](ApiEPSS.md) |  | [optional] 
**Ssvc** | Pointer to [**[]ApiSSVC**](ApiSSVC.md) |  | [optional] 
**TemporalCVSSV2** | Pointer to [**ApiNVD20TemporalCVSSV2**](ApiNVD20TemporalCVSSV2.md) |  | [optional] 
**TemporalCVSSV2Secondary** | Pointer to [**[]ApiNVD20TemporalCVSSV2**](ApiNVD20TemporalCVSSV2.md) |  | [optional] 
**TemporalCVSSV30** | Pointer to [**ApiNVD20TemporalCVSSV3**](ApiNVD20TemporalCVSSV3.md) |  | [optional] 
**TemporalCVSSV30Secondary** | Pointer to [**[]ApiNVD20TemporalCVSSV3**](ApiNVD20TemporalCVSSV3.md) |  | [optional] 
**TemporalCVSSV31** | Pointer to [**ApiNVD20TemporalCVSSV3**](ApiNVD20TemporalCVSSV3.md) |  | [optional] 
**TemporalCVSSV31Secondary** | Pointer to [**[]ApiNVD20TemporalCVSSV3**](ApiNVD20TemporalCVSSV3.md) |  | [optional] 
**ThreatCVSSV40** | Pointer to [**ApiNVD20ThreatCVSSV40**](ApiNVD20ThreatCVSSV40.md) |  | [optional] 
**ThreatCVSSV40Secondary** | Pointer to [**[]ApiNVD20ThreatCVSSV40**](ApiNVD20ThreatCVSSV40.md) |  | [optional] 

## Methods

### NewApiNVD20MetricExtended

`func NewApiNVD20MetricExtended() *ApiNVD20MetricExtended`

NewApiNVD20MetricExtended instantiates a new ApiNVD20MetricExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNVD20MetricExtendedWithDefaults

`func NewApiNVD20MetricExtendedWithDefaults() *ApiNVD20MetricExtended`

NewApiNVD20MetricExtendedWithDefaults instantiates a new ApiNVD20MetricExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvssMetricV2

`func (o *ApiNVD20MetricExtended) GetCvssMetricV2() []ApiNVD20CvssMetricV2`

GetCvssMetricV2 returns the CvssMetricV2 field if non-nil, zero value otherwise.

### GetCvssMetricV2Ok

`func (o *ApiNVD20MetricExtended) GetCvssMetricV2Ok() (*[]ApiNVD20CvssMetricV2, bool)`

GetCvssMetricV2Ok returns a tuple with the CvssMetricV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssMetricV2

`func (o *ApiNVD20MetricExtended) SetCvssMetricV2(v []ApiNVD20CvssMetricV2)`

SetCvssMetricV2 sets CvssMetricV2 field to given value.

### HasCvssMetricV2

`func (o *ApiNVD20MetricExtended) HasCvssMetricV2() bool`

HasCvssMetricV2 returns a boolean if a field has been set.

### GetCvssMetricV30

`func (o *ApiNVD20MetricExtended) GetCvssMetricV30() []ApiNVD20CvssMetricV3`

GetCvssMetricV30 returns the CvssMetricV30 field if non-nil, zero value otherwise.

### GetCvssMetricV30Ok

`func (o *ApiNVD20MetricExtended) GetCvssMetricV30Ok() (*[]ApiNVD20CvssMetricV3, bool)`

GetCvssMetricV30Ok returns a tuple with the CvssMetricV30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssMetricV30

`func (o *ApiNVD20MetricExtended) SetCvssMetricV30(v []ApiNVD20CvssMetricV3)`

SetCvssMetricV30 sets CvssMetricV30 field to given value.

### HasCvssMetricV30

`func (o *ApiNVD20MetricExtended) HasCvssMetricV30() bool`

HasCvssMetricV30 returns a boolean if a field has been set.

### GetCvssMetricV31

`func (o *ApiNVD20MetricExtended) GetCvssMetricV31() []ApiNVD20CvssMetricV3`

GetCvssMetricV31 returns the CvssMetricV31 field if non-nil, zero value otherwise.

### GetCvssMetricV31Ok

`func (o *ApiNVD20MetricExtended) GetCvssMetricV31Ok() (*[]ApiNVD20CvssMetricV3, bool)`

GetCvssMetricV31Ok returns a tuple with the CvssMetricV31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssMetricV31

`func (o *ApiNVD20MetricExtended) SetCvssMetricV31(v []ApiNVD20CvssMetricV3)`

SetCvssMetricV31 sets CvssMetricV31 field to given value.

### HasCvssMetricV31

`func (o *ApiNVD20MetricExtended) HasCvssMetricV31() bool`

HasCvssMetricV31 returns a boolean if a field has been set.

### GetCvssMetricV40

`func (o *ApiNVD20MetricExtended) GetCvssMetricV40() []ApiNVD20CvssMetricV40`

GetCvssMetricV40 returns the CvssMetricV40 field if non-nil, zero value otherwise.

### GetCvssMetricV40Ok

`func (o *ApiNVD20MetricExtended) GetCvssMetricV40Ok() (*[]ApiNVD20CvssMetricV40, bool)`

GetCvssMetricV40Ok returns a tuple with the CvssMetricV40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssMetricV40

`func (o *ApiNVD20MetricExtended) SetCvssMetricV40(v []ApiNVD20CvssMetricV40)`

SetCvssMetricV40 sets CvssMetricV40 field to given value.

### HasCvssMetricV40

`func (o *ApiNVD20MetricExtended) HasCvssMetricV40() bool`

HasCvssMetricV40 returns a boolean if a field has been set.

### GetEpss

`func (o *ApiNVD20MetricExtended) GetEpss() ApiEPSS`

GetEpss returns the Epss field if non-nil, zero value otherwise.

### GetEpssOk

`func (o *ApiNVD20MetricExtended) GetEpssOk() (*ApiEPSS, bool)`

GetEpssOk returns a tuple with the Epss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpss

`func (o *ApiNVD20MetricExtended) SetEpss(v ApiEPSS)`

SetEpss sets Epss field to given value.

### HasEpss

`func (o *ApiNVD20MetricExtended) HasEpss() bool`

HasEpss returns a boolean if a field has been set.

### GetSsvc

`func (o *ApiNVD20MetricExtended) GetSsvc() []ApiSSVC`

GetSsvc returns the Ssvc field if non-nil, zero value otherwise.

### GetSsvcOk

`func (o *ApiNVD20MetricExtended) GetSsvcOk() (*[]ApiSSVC, bool)`

GetSsvcOk returns a tuple with the Ssvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsvc

`func (o *ApiNVD20MetricExtended) SetSsvc(v []ApiSSVC)`

SetSsvc sets Ssvc field to given value.

### HasSsvc

`func (o *ApiNVD20MetricExtended) HasSsvc() bool`

HasSsvc returns a boolean if a field has been set.

### GetTemporalCVSSV2

`func (o *ApiNVD20MetricExtended) GetTemporalCVSSV2() ApiNVD20TemporalCVSSV2`

GetTemporalCVSSV2 returns the TemporalCVSSV2 field if non-nil, zero value otherwise.

### GetTemporalCVSSV2Ok

`func (o *ApiNVD20MetricExtended) GetTemporalCVSSV2Ok() (*ApiNVD20TemporalCVSSV2, bool)`

GetTemporalCVSSV2Ok returns a tuple with the TemporalCVSSV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalCVSSV2

`func (o *ApiNVD20MetricExtended) SetTemporalCVSSV2(v ApiNVD20TemporalCVSSV2)`

SetTemporalCVSSV2 sets TemporalCVSSV2 field to given value.

### HasTemporalCVSSV2

`func (o *ApiNVD20MetricExtended) HasTemporalCVSSV2() bool`

HasTemporalCVSSV2 returns a boolean if a field has been set.

### GetTemporalCVSSV2Secondary

`func (o *ApiNVD20MetricExtended) GetTemporalCVSSV2Secondary() []ApiNVD20TemporalCVSSV2`

GetTemporalCVSSV2Secondary returns the TemporalCVSSV2Secondary field if non-nil, zero value otherwise.

### GetTemporalCVSSV2SecondaryOk

`func (o *ApiNVD20MetricExtended) GetTemporalCVSSV2SecondaryOk() (*[]ApiNVD20TemporalCVSSV2, bool)`

GetTemporalCVSSV2SecondaryOk returns a tuple with the TemporalCVSSV2Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalCVSSV2Secondary

`func (o *ApiNVD20MetricExtended) SetTemporalCVSSV2Secondary(v []ApiNVD20TemporalCVSSV2)`

SetTemporalCVSSV2Secondary sets TemporalCVSSV2Secondary field to given value.

### HasTemporalCVSSV2Secondary

`func (o *ApiNVD20MetricExtended) HasTemporalCVSSV2Secondary() bool`

HasTemporalCVSSV2Secondary returns a boolean if a field has been set.

### GetTemporalCVSSV30

`func (o *ApiNVD20MetricExtended) GetTemporalCVSSV30() ApiNVD20TemporalCVSSV3`

GetTemporalCVSSV30 returns the TemporalCVSSV30 field if non-nil, zero value otherwise.

### GetTemporalCVSSV30Ok

`func (o *ApiNVD20MetricExtended) GetTemporalCVSSV30Ok() (*ApiNVD20TemporalCVSSV3, bool)`

GetTemporalCVSSV30Ok returns a tuple with the TemporalCVSSV30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalCVSSV30

`func (o *ApiNVD20MetricExtended) SetTemporalCVSSV30(v ApiNVD20TemporalCVSSV3)`

SetTemporalCVSSV30 sets TemporalCVSSV30 field to given value.

### HasTemporalCVSSV30

`func (o *ApiNVD20MetricExtended) HasTemporalCVSSV30() bool`

HasTemporalCVSSV30 returns a boolean if a field has been set.

### GetTemporalCVSSV30Secondary

`func (o *ApiNVD20MetricExtended) GetTemporalCVSSV30Secondary() []ApiNVD20TemporalCVSSV3`

GetTemporalCVSSV30Secondary returns the TemporalCVSSV30Secondary field if non-nil, zero value otherwise.

### GetTemporalCVSSV30SecondaryOk

`func (o *ApiNVD20MetricExtended) GetTemporalCVSSV30SecondaryOk() (*[]ApiNVD20TemporalCVSSV3, bool)`

GetTemporalCVSSV30SecondaryOk returns a tuple with the TemporalCVSSV30Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalCVSSV30Secondary

`func (o *ApiNVD20MetricExtended) SetTemporalCVSSV30Secondary(v []ApiNVD20TemporalCVSSV3)`

SetTemporalCVSSV30Secondary sets TemporalCVSSV30Secondary field to given value.

### HasTemporalCVSSV30Secondary

`func (o *ApiNVD20MetricExtended) HasTemporalCVSSV30Secondary() bool`

HasTemporalCVSSV30Secondary returns a boolean if a field has been set.

### GetTemporalCVSSV31

`func (o *ApiNVD20MetricExtended) GetTemporalCVSSV31() ApiNVD20TemporalCVSSV3`

GetTemporalCVSSV31 returns the TemporalCVSSV31 field if non-nil, zero value otherwise.

### GetTemporalCVSSV31Ok

`func (o *ApiNVD20MetricExtended) GetTemporalCVSSV31Ok() (*ApiNVD20TemporalCVSSV3, bool)`

GetTemporalCVSSV31Ok returns a tuple with the TemporalCVSSV31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalCVSSV31

`func (o *ApiNVD20MetricExtended) SetTemporalCVSSV31(v ApiNVD20TemporalCVSSV3)`

SetTemporalCVSSV31 sets TemporalCVSSV31 field to given value.

### HasTemporalCVSSV31

`func (o *ApiNVD20MetricExtended) HasTemporalCVSSV31() bool`

HasTemporalCVSSV31 returns a boolean if a field has been set.

### GetTemporalCVSSV31Secondary

`func (o *ApiNVD20MetricExtended) GetTemporalCVSSV31Secondary() []ApiNVD20TemporalCVSSV3`

GetTemporalCVSSV31Secondary returns the TemporalCVSSV31Secondary field if non-nil, zero value otherwise.

### GetTemporalCVSSV31SecondaryOk

`func (o *ApiNVD20MetricExtended) GetTemporalCVSSV31SecondaryOk() (*[]ApiNVD20TemporalCVSSV3, bool)`

GetTemporalCVSSV31SecondaryOk returns a tuple with the TemporalCVSSV31Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalCVSSV31Secondary

`func (o *ApiNVD20MetricExtended) SetTemporalCVSSV31Secondary(v []ApiNVD20TemporalCVSSV3)`

SetTemporalCVSSV31Secondary sets TemporalCVSSV31Secondary field to given value.

### HasTemporalCVSSV31Secondary

`func (o *ApiNVD20MetricExtended) HasTemporalCVSSV31Secondary() bool`

HasTemporalCVSSV31Secondary returns a boolean if a field has been set.

### GetThreatCVSSV40

`func (o *ApiNVD20MetricExtended) GetThreatCVSSV40() ApiNVD20ThreatCVSSV40`

GetThreatCVSSV40 returns the ThreatCVSSV40 field if non-nil, zero value otherwise.

### GetThreatCVSSV40Ok

`func (o *ApiNVD20MetricExtended) GetThreatCVSSV40Ok() (*ApiNVD20ThreatCVSSV40, bool)`

GetThreatCVSSV40Ok returns a tuple with the ThreatCVSSV40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatCVSSV40

`func (o *ApiNVD20MetricExtended) SetThreatCVSSV40(v ApiNVD20ThreatCVSSV40)`

SetThreatCVSSV40 sets ThreatCVSSV40 field to given value.

### HasThreatCVSSV40

`func (o *ApiNVD20MetricExtended) HasThreatCVSSV40() bool`

HasThreatCVSSV40 returns a boolean if a field has been set.

### GetThreatCVSSV40Secondary

`func (o *ApiNVD20MetricExtended) GetThreatCVSSV40Secondary() []ApiNVD20ThreatCVSSV40`

GetThreatCVSSV40Secondary returns the ThreatCVSSV40Secondary field if non-nil, zero value otherwise.

### GetThreatCVSSV40SecondaryOk

`func (o *ApiNVD20MetricExtended) GetThreatCVSSV40SecondaryOk() (*[]ApiNVD20ThreatCVSSV40, bool)`

GetThreatCVSSV40SecondaryOk returns a tuple with the ThreatCVSSV40Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatCVSSV40Secondary

`func (o *ApiNVD20MetricExtended) SetThreatCVSSV40Secondary(v []ApiNVD20ThreatCVSSV40)`

SetThreatCVSSV40Secondary sets ThreatCVSSV40Secondary field to given value.

### HasThreatCVSSV40Secondary

`func (o *ApiNVD20MetricExtended) HasThreatCVSSV40Secondary() bool`

HasThreatCVSSV40Secondary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


