# ApiBaseMetricV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CvssV3** | Pointer to [**ApiCVSSV3**](ApiCVSSV3.md) |  | [optional] 
**ExploitabilityScore** | Pointer to **float32** |  | [optional] 
**ImpactScore** | Pointer to **float32** |  | [optional] 

## Methods

### NewApiBaseMetricV3

`func NewApiBaseMetricV3() *ApiBaseMetricV3`

NewApiBaseMetricV3 instantiates a new ApiBaseMetricV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiBaseMetricV3WithDefaults

`func NewApiBaseMetricV3WithDefaults() *ApiBaseMetricV3`

NewApiBaseMetricV3WithDefaults instantiates a new ApiBaseMetricV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvssV3

`func (o *ApiBaseMetricV3) GetCvssV3() ApiCVSSV3`

GetCvssV3 returns the CvssV3 field if non-nil, zero value otherwise.

### GetCvssV3Ok

`func (o *ApiBaseMetricV3) GetCvssV3Ok() (*ApiCVSSV3, bool)`

GetCvssV3Ok returns a tuple with the CvssV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV3

`func (o *ApiBaseMetricV3) SetCvssV3(v ApiCVSSV3)`

SetCvssV3 sets CvssV3 field to given value.

### HasCvssV3

`func (o *ApiBaseMetricV3) HasCvssV3() bool`

HasCvssV3 returns a boolean if a field has been set.

### GetExploitabilityScore

`func (o *ApiBaseMetricV3) GetExploitabilityScore() float32`

GetExploitabilityScore returns the ExploitabilityScore field if non-nil, zero value otherwise.

### GetExploitabilityScoreOk

`func (o *ApiBaseMetricV3) GetExploitabilityScoreOk() (*float32, bool)`

GetExploitabilityScoreOk returns a tuple with the ExploitabilityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitabilityScore

`func (o *ApiBaseMetricV3) SetExploitabilityScore(v float32)`

SetExploitabilityScore sets ExploitabilityScore field to given value.

### HasExploitabilityScore

`func (o *ApiBaseMetricV3) HasExploitabilityScore() bool`

HasExploitabilityScore returns a boolean if a field has been set.

### GetImpactScore

`func (o *ApiBaseMetricV3) GetImpactScore() float32`

GetImpactScore returns the ImpactScore field if non-nil, zero value otherwise.

### GetImpactScoreOk

`func (o *ApiBaseMetricV3) GetImpactScoreOk() (*float32, bool)`

GetImpactScoreOk returns a tuple with the ImpactScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactScore

`func (o *ApiBaseMetricV3) SetImpactScore(v float32)`

SetImpactScore sets ImpactScore field to given value.

### HasImpactScore

`func (o *ApiBaseMetricV3) HasImpactScore() bool`

HasImpactScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


