# ApiImpact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseMetricV2** | Pointer to [**ApiBaseMetricV2**](ApiBaseMetricV2.md) |  | [optional] 
**BaseMetricV3** | Pointer to [**ApiBaseMetricV3**](ApiBaseMetricV3.md) |  | [optional] 
**MetricV40** | Pointer to [**AdvisoryCVSSV40**](AdvisoryCVSSV40.md) | this isn&#39;t called baseMetric, because it can contain other metrics -- typically supplemental metrics | [optional] 

## Methods

### NewApiImpact

`func NewApiImpact() *ApiImpact`

NewApiImpact instantiates a new ApiImpact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiImpactWithDefaults

`func NewApiImpactWithDefaults() *ApiImpact`

NewApiImpactWithDefaults instantiates a new ApiImpact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseMetricV2

`func (o *ApiImpact) GetBaseMetricV2() ApiBaseMetricV2`

GetBaseMetricV2 returns the BaseMetricV2 field if non-nil, zero value otherwise.

### GetBaseMetricV2Ok

`func (o *ApiImpact) GetBaseMetricV2Ok() (*ApiBaseMetricV2, bool)`

GetBaseMetricV2Ok returns a tuple with the BaseMetricV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseMetricV2

`func (o *ApiImpact) SetBaseMetricV2(v ApiBaseMetricV2)`

SetBaseMetricV2 sets BaseMetricV2 field to given value.

### HasBaseMetricV2

`func (o *ApiImpact) HasBaseMetricV2() bool`

HasBaseMetricV2 returns a boolean if a field has been set.

### GetBaseMetricV3

`func (o *ApiImpact) GetBaseMetricV3() ApiBaseMetricV3`

GetBaseMetricV3 returns the BaseMetricV3 field if non-nil, zero value otherwise.

### GetBaseMetricV3Ok

`func (o *ApiImpact) GetBaseMetricV3Ok() (*ApiBaseMetricV3, bool)`

GetBaseMetricV3Ok returns a tuple with the BaseMetricV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseMetricV3

`func (o *ApiImpact) SetBaseMetricV3(v ApiBaseMetricV3)`

SetBaseMetricV3 sets BaseMetricV3 field to given value.

### HasBaseMetricV3

`func (o *ApiImpact) HasBaseMetricV3() bool`

HasBaseMetricV3 returns a boolean if a field has been set.

### GetMetricV40

`func (o *ApiImpact) GetMetricV40() AdvisoryCVSSV40`

GetMetricV40 returns the MetricV40 field if non-nil, zero value otherwise.

### GetMetricV40Ok

`func (o *ApiImpact) GetMetricV40Ok() (*AdvisoryCVSSV40, bool)`

GetMetricV40Ok returns a tuple with the MetricV40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricV40

`func (o *ApiImpact) SetMetricV40(v AdvisoryCVSSV40)`

SetMetricV40 sets MetricV40 field to given value.

### HasMetricV40

`func (o *ApiImpact) HasMetricV40() bool`

HasMetricV40 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


