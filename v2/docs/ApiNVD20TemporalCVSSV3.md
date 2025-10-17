# ApiNVD20TemporalCVSSV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedBaseMetricV3** | Pointer to [**ApiNVD20TemporalAssociatedBaseMetric**](ApiNVD20TemporalAssociatedBaseMetric.md) |  | [optional] 
**ExploitCodeMaturity** | Pointer to **string** |  | [optional] 
**RemediationLevel** | Pointer to **string** |  | [optional] 
**ReportConfidence** | Pointer to **string** |  | [optional] 
**TemporalScore** | Pointer to **float32** |  | [optional] 
**VectorString** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewApiNVD20TemporalCVSSV3

`func NewApiNVD20TemporalCVSSV3() *ApiNVD20TemporalCVSSV3`

NewApiNVD20TemporalCVSSV3 instantiates a new ApiNVD20TemporalCVSSV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNVD20TemporalCVSSV3WithDefaults

`func NewApiNVD20TemporalCVSSV3WithDefaults() *ApiNVD20TemporalCVSSV3`

NewApiNVD20TemporalCVSSV3WithDefaults instantiates a new ApiNVD20TemporalCVSSV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedBaseMetricV3

`func (o *ApiNVD20TemporalCVSSV3) GetAssociatedBaseMetricV3() ApiNVD20TemporalAssociatedBaseMetric`

GetAssociatedBaseMetricV3 returns the AssociatedBaseMetricV3 field if non-nil, zero value otherwise.

### GetAssociatedBaseMetricV3Ok

`func (o *ApiNVD20TemporalCVSSV3) GetAssociatedBaseMetricV3Ok() (*ApiNVD20TemporalAssociatedBaseMetric, bool)`

GetAssociatedBaseMetricV3Ok returns a tuple with the AssociatedBaseMetricV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedBaseMetricV3

`func (o *ApiNVD20TemporalCVSSV3) SetAssociatedBaseMetricV3(v ApiNVD20TemporalAssociatedBaseMetric)`

SetAssociatedBaseMetricV3 sets AssociatedBaseMetricV3 field to given value.

### HasAssociatedBaseMetricV3

`func (o *ApiNVD20TemporalCVSSV3) HasAssociatedBaseMetricV3() bool`

HasAssociatedBaseMetricV3 returns a boolean if a field has been set.

### GetExploitCodeMaturity

`func (o *ApiNVD20TemporalCVSSV3) GetExploitCodeMaturity() string`

GetExploitCodeMaturity returns the ExploitCodeMaturity field if non-nil, zero value otherwise.

### GetExploitCodeMaturityOk

`func (o *ApiNVD20TemporalCVSSV3) GetExploitCodeMaturityOk() (*string, bool)`

GetExploitCodeMaturityOk returns a tuple with the ExploitCodeMaturity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitCodeMaturity

`func (o *ApiNVD20TemporalCVSSV3) SetExploitCodeMaturity(v string)`

SetExploitCodeMaturity sets ExploitCodeMaturity field to given value.

### HasExploitCodeMaturity

`func (o *ApiNVD20TemporalCVSSV3) HasExploitCodeMaturity() bool`

HasExploitCodeMaturity returns a boolean if a field has been set.

### GetRemediationLevel

`func (o *ApiNVD20TemporalCVSSV3) GetRemediationLevel() string`

GetRemediationLevel returns the RemediationLevel field if non-nil, zero value otherwise.

### GetRemediationLevelOk

`func (o *ApiNVD20TemporalCVSSV3) GetRemediationLevelOk() (*string, bool)`

GetRemediationLevelOk returns a tuple with the RemediationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationLevel

`func (o *ApiNVD20TemporalCVSSV3) SetRemediationLevel(v string)`

SetRemediationLevel sets RemediationLevel field to given value.

### HasRemediationLevel

`func (o *ApiNVD20TemporalCVSSV3) HasRemediationLevel() bool`

HasRemediationLevel returns a boolean if a field has been set.

### GetReportConfidence

`func (o *ApiNVD20TemporalCVSSV3) GetReportConfidence() string`

GetReportConfidence returns the ReportConfidence field if non-nil, zero value otherwise.

### GetReportConfidenceOk

`func (o *ApiNVD20TemporalCVSSV3) GetReportConfidenceOk() (*string, bool)`

GetReportConfidenceOk returns a tuple with the ReportConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportConfidence

`func (o *ApiNVD20TemporalCVSSV3) SetReportConfidence(v string)`

SetReportConfidence sets ReportConfidence field to given value.

### HasReportConfidence

`func (o *ApiNVD20TemporalCVSSV3) HasReportConfidence() bool`

HasReportConfidence returns a boolean if a field has been set.

### GetTemporalScore

`func (o *ApiNVD20TemporalCVSSV3) GetTemporalScore() float32`

GetTemporalScore returns the TemporalScore field if non-nil, zero value otherwise.

### GetTemporalScoreOk

`func (o *ApiNVD20TemporalCVSSV3) GetTemporalScoreOk() (*float32, bool)`

GetTemporalScoreOk returns a tuple with the TemporalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalScore

`func (o *ApiNVD20TemporalCVSSV3) SetTemporalScore(v float32)`

SetTemporalScore sets TemporalScore field to given value.

### HasTemporalScore

`func (o *ApiNVD20TemporalCVSSV3) HasTemporalScore() bool`

HasTemporalScore returns a boolean if a field has been set.

### GetVectorString

`func (o *ApiNVD20TemporalCVSSV3) GetVectorString() string`

GetVectorString returns the VectorString field if non-nil, zero value otherwise.

### GetVectorStringOk

`func (o *ApiNVD20TemporalCVSSV3) GetVectorStringOk() (*string, bool)`

GetVectorStringOk returns a tuple with the VectorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorString

`func (o *ApiNVD20TemporalCVSSV3) SetVectorString(v string)`

SetVectorString sets VectorString field to given value.

### HasVectorString

`func (o *ApiNVD20TemporalCVSSV3) HasVectorString() bool`

HasVectorString returns a boolean if a field has been set.

### GetVersion

`func (o *ApiNVD20TemporalCVSSV3) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiNVD20TemporalCVSSV3) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiNVD20TemporalCVSSV3) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApiNVD20TemporalCVSSV3) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


