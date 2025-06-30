# ApiNVD20CvssMetricV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CvssData** | Pointer to [**ApiNVD20CvssDataV3**](ApiNVD20CvssDataV3.md) |  | [optional] 
**ExploitabilityScore** | Pointer to **float32** |  | [optional] 
**ImpactScore** | Pointer to **float32** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewApiNVD20CvssMetricV3

`func NewApiNVD20CvssMetricV3() *ApiNVD20CvssMetricV3`

NewApiNVD20CvssMetricV3 instantiates a new ApiNVD20CvssMetricV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNVD20CvssMetricV3WithDefaults

`func NewApiNVD20CvssMetricV3WithDefaults() *ApiNVD20CvssMetricV3`

NewApiNVD20CvssMetricV3WithDefaults instantiates a new ApiNVD20CvssMetricV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvssData

`func (o *ApiNVD20CvssMetricV3) GetCvssData() ApiNVD20CvssDataV3`

GetCvssData returns the CvssData field if non-nil, zero value otherwise.

### GetCvssDataOk

`func (o *ApiNVD20CvssMetricV3) GetCvssDataOk() (*ApiNVD20CvssDataV3, bool)`

GetCvssDataOk returns a tuple with the CvssData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssData

`func (o *ApiNVD20CvssMetricV3) SetCvssData(v ApiNVD20CvssDataV3)`

SetCvssData sets CvssData field to given value.

### HasCvssData

`func (o *ApiNVD20CvssMetricV3) HasCvssData() bool`

HasCvssData returns a boolean if a field has been set.

### GetExploitabilityScore

`func (o *ApiNVD20CvssMetricV3) GetExploitabilityScore() float32`

GetExploitabilityScore returns the ExploitabilityScore field if non-nil, zero value otherwise.

### GetExploitabilityScoreOk

`func (o *ApiNVD20CvssMetricV3) GetExploitabilityScoreOk() (*float32, bool)`

GetExploitabilityScoreOk returns a tuple with the ExploitabilityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitabilityScore

`func (o *ApiNVD20CvssMetricV3) SetExploitabilityScore(v float32)`

SetExploitabilityScore sets ExploitabilityScore field to given value.

### HasExploitabilityScore

`func (o *ApiNVD20CvssMetricV3) HasExploitabilityScore() bool`

HasExploitabilityScore returns a boolean if a field has been set.

### GetImpactScore

`func (o *ApiNVD20CvssMetricV3) GetImpactScore() float32`

GetImpactScore returns the ImpactScore field if non-nil, zero value otherwise.

### GetImpactScoreOk

`func (o *ApiNVD20CvssMetricV3) GetImpactScoreOk() (*float32, bool)`

GetImpactScoreOk returns a tuple with the ImpactScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactScore

`func (o *ApiNVD20CvssMetricV3) SetImpactScore(v float32)`

SetImpactScore sets ImpactScore field to given value.

### HasImpactScore

`func (o *ApiNVD20CvssMetricV3) HasImpactScore() bool`

HasImpactScore returns a boolean if a field has been set.

### GetSource

`func (o *ApiNVD20CvssMetricV3) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiNVD20CvssMetricV3) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiNVD20CvssMetricV3) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApiNVD20CvssMetricV3) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *ApiNVD20CvssMetricV3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiNVD20CvssMetricV3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiNVD20CvssMetricV3) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiNVD20CvssMetricV3) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


