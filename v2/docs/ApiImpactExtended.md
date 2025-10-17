# ApiImpactExtended

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseMetricV2** | Pointer to [**ApiBaseMetricV2**](ApiBaseMetricV2.md) |  | [optional] 
**BaseMetricV3** | Pointer to [**ApiBaseMetricV3**](ApiBaseMetricV3.md) |  | [optional] 
**CorrectedBaseMetricV3** | Pointer to [**ApiBaseMetricV3**](ApiBaseMetricV3.md) |  | [optional] 
**Epss** | Pointer to [**ApiEPSS**](ApiEPSS.md) |  | [optional] 
**MetricV40** | Pointer to [**AdvisoryCVSSV40**](AdvisoryCVSSV40.md) |  | [optional] 
**Ssvc** | Pointer to [**[]ApiSSVC**](ApiSSVC.md) |  | [optional] 
**TemporalMetricV2** | Pointer to [**ApiTemporalMetricV2**](ApiTemporalMetricV2.md) |  | [optional] 
**TemporalMetricV3** | Pointer to [**ApiTemporalMetricV3**](ApiTemporalMetricV3.md) |  | [optional] 
**TemporalV3Corrected** | Pointer to [**ApiTemporalMetricV3**](ApiTemporalMetricV3.md) |  | [optional] 
**ThreatMetricV40** | Pointer to [**AdvisoryCVSSV40Threat**](AdvisoryCVSSV40Threat.md) |  | [optional] 

## Methods

### NewApiImpactExtended

`func NewApiImpactExtended() *ApiImpactExtended`

NewApiImpactExtended instantiates a new ApiImpactExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiImpactExtendedWithDefaults

`func NewApiImpactExtendedWithDefaults() *ApiImpactExtended`

NewApiImpactExtendedWithDefaults instantiates a new ApiImpactExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseMetricV2

`func (o *ApiImpactExtended) GetBaseMetricV2() ApiBaseMetricV2`

GetBaseMetricV2 returns the BaseMetricV2 field if non-nil, zero value otherwise.

### GetBaseMetricV2Ok

`func (o *ApiImpactExtended) GetBaseMetricV2Ok() (*ApiBaseMetricV2, bool)`

GetBaseMetricV2Ok returns a tuple with the BaseMetricV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseMetricV2

`func (o *ApiImpactExtended) SetBaseMetricV2(v ApiBaseMetricV2)`

SetBaseMetricV2 sets BaseMetricV2 field to given value.

### HasBaseMetricV2

`func (o *ApiImpactExtended) HasBaseMetricV2() bool`

HasBaseMetricV2 returns a boolean if a field has been set.

### GetBaseMetricV3

`func (o *ApiImpactExtended) GetBaseMetricV3() ApiBaseMetricV3`

GetBaseMetricV3 returns the BaseMetricV3 field if non-nil, zero value otherwise.

### GetBaseMetricV3Ok

`func (o *ApiImpactExtended) GetBaseMetricV3Ok() (*ApiBaseMetricV3, bool)`

GetBaseMetricV3Ok returns a tuple with the BaseMetricV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseMetricV3

`func (o *ApiImpactExtended) SetBaseMetricV3(v ApiBaseMetricV3)`

SetBaseMetricV3 sets BaseMetricV3 field to given value.

### HasBaseMetricV3

`func (o *ApiImpactExtended) HasBaseMetricV3() bool`

HasBaseMetricV3 returns a boolean if a field has been set.

### GetCorrectedBaseMetricV3

`func (o *ApiImpactExtended) GetCorrectedBaseMetricV3() ApiBaseMetricV3`

GetCorrectedBaseMetricV3 returns the CorrectedBaseMetricV3 field if non-nil, zero value otherwise.

### GetCorrectedBaseMetricV3Ok

`func (o *ApiImpactExtended) GetCorrectedBaseMetricV3Ok() (*ApiBaseMetricV3, bool)`

GetCorrectedBaseMetricV3Ok returns a tuple with the CorrectedBaseMetricV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectedBaseMetricV3

`func (o *ApiImpactExtended) SetCorrectedBaseMetricV3(v ApiBaseMetricV3)`

SetCorrectedBaseMetricV3 sets CorrectedBaseMetricV3 field to given value.

### HasCorrectedBaseMetricV3

`func (o *ApiImpactExtended) HasCorrectedBaseMetricV3() bool`

HasCorrectedBaseMetricV3 returns a boolean if a field has been set.

### GetEpss

`func (o *ApiImpactExtended) GetEpss() ApiEPSS`

GetEpss returns the Epss field if non-nil, zero value otherwise.

### GetEpssOk

`func (o *ApiImpactExtended) GetEpssOk() (*ApiEPSS, bool)`

GetEpssOk returns a tuple with the Epss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpss

`func (o *ApiImpactExtended) SetEpss(v ApiEPSS)`

SetEpss sets Epss field to given value.

### HasEpss

`func (o *ApiImpactExtended) HasEpss() bool`

HasEpss returns a boolean if a field has been set.

### GetMetricV40

`func (o *ApiImpactExtended) GetMetricV40() AdvisoryCVSSV40`

GetMetricV40 returns the MetricV40 field if non-nil, zero value otherwise.

### GetMetricV40Ok

`func (o *ApiImpactExtended) GetMetricV40Ok() (*AdvisoryCVSSV40, bool)`

GetMetricV40Ok returns a tuple with the MetricV40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricV40

`func (o *ApiImpactExtended) SetMetricV40(v AdvisoryCVSSV40)`

SetMetricV40 sets MetricV40 field to given value.

### HasMetricV40

`func (o *ApiImpactExtended) HasMetricV40() bool`

HasMetricV40 returns a boolean if a field has been set.

### GetSsvc

`func (o *ApiImpactExtended) GetSsvc() []ApiSSVC`

GetSsvc returns the Ssvc field if non-nil, zero value otherwise.

### GetSsvcOk

`func (o *ApiImpactExtended) GetSsvcOk() (*[]ApiSSVC, bool)`

GetSsvcOk returns a tuple with the Ssvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsvc

`func (o *ApiImpactExtended) SetSsvc(v []ApiSSVC)`

SetSsvc sets Ssvc field to given value.

### HasSsvc

`func (o *ApiImpactExtended) HasSsvc() bool`

HasSsvc returns a boolean if a field has been set.

### GetTemporalMetricV2

`func (o *ApiImpactExtended) GetTemporalMetricV2() ApiTemporalMetricV2`

GetTemporalMetricV2 returns the TemporalMetricV2 field if non-nil, zero value otherwise.

### GetTemporalMetricV2Ok

`func (o *ApiImpactExtended) GetTemporalMetricV2Ok() (*ApiTemporalMetricV2, bool)`

GetTemporalMetricV2Ok returns a tuple with the TemporalMetricV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalMetricV2

`func (o *ApiImpactExtended) SetTemporalMetricV2(v ApiTemporalMetricV2)`

SetTemporalMetricV2 sets TemporalMetricV2 field to given value.

### HasTemporalMetricV2

`func (o *ApiImpactExtended) HasTemporalMetricV2() bool`

HasTemporalMetricV2 returns a boolean if a field has been set.

### GetTemporalMetricV3

`func (o *ApiImpactExtended) GetTemporalMetricV3() ApiTemporalMetricV3`

GetTemporalMetricV3 returns the TemporalMetricV3 field if non-nil, zero value otherwise.

### GetTemporalMetricV3Ok

`func (o *ApiImpactExtended) GetTemporalMetricV3Ok() (*ApiTemporalMetricV3, bool)`

GetTemporalMetricV3Ok returns a tuple with the TemporalMetricV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalMetricV3

`func (o *ApiImpactExtended) SetTemporalMetricV3(v ApiTemporalMetricV3)`

SetTemporalMetricV3 sets TemporalMetricV3 field to given value.

### HasTemporalMetricV3

`func (o *ApiImpactExtended) HasTemporalMetricV3() bool`

HasTemporalMetricV3 returns a boolean if a field has been set.

### GetTemporalV3Corrected

`func (o *ApiImpactExtended) GetTemporalV3Corrected() ApiTemporalMetricV3`

GetTemporalV3Corrected returns the TemporalV3Corrected field if non-nil, zero value otherwise.

### GetTemporalV3CorrectedOk

`func (o *ApiImpactExtended) GetTemporalV3CorrectedOk() (*ApiTemporalMetricV3, bool)`

GetTemporalV3CorrectedOk returns a tuple with the TemporalV3Corrected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalV3Corrected

`func (o *ApiImpactExtended) SetTemporalV3Corrected(v ApiTemporalMetricV3)`

SetTemporalV3Corrected sets TemporalV3Corrected field to given value.

### HasTemporalV3Corrected

`func (o *ApiImpactExtended) HasTemporalV3Corrected() bool`

HasTemporalV3Corrected returns a boolean if a field has been set.

### GetThreatMetricV40

`func (o *ApiImpactExtended) GetThreatMetricV40() AdvisoryCVSSV40Threat`

GetThreatMetricV40 returns the ThreatMetricV40 field if non-nil, zero value otherwise.

### GetThreatMetricV40Ok

`func (o *ApiImpactExtended) GetThreatMetricV40Ok() (*AdvisoryCVSSV40Threat, bool)`

GetThreatMetricV40Ok returns a tuple with the ThreatMetricV40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatMetricV40

`func (o *ApiImpactExtended) SetThreatMetricV40(v AdvisoryCVSSV40Threat)`

SetThreatMetricV40 sets ThreatMetricV40 field to given value.

### HasThreatMetricV40

`func (o *ApiImpactExtended) HasThreatMetricV40() bool`

HasThreatMetricV40 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


