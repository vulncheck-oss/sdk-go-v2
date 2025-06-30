# ApiNVD20TemporalCVSSV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedBaseMetricV2** | Pointer to [**ApiNVD20TemporalAssociatedBaseMetric**](ApiNVD20TemporalAssociatedBaseMetric.md) |  | [optional] 
**Exploitability** | Pointer to **string** |  | [optional] 
**RemediationLevel** | Pointer to **string** |  | [optional] 
**ReportConfidence** | Pointer to **string** |  | [optional] 
**TemporalScore** | Pointer to **float32** |  | [optional] 
**VectorString** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewApiNVD20TemporalCVSSV2

`func NewApiNVD20TemporalCVSSV2() *ApiNVD20TemporalCVSSV2`

NewApiNVD20TemporalCVSSV2 instantiates a new ApiNVD20TemporalCVSSV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNVD20TemporalCVSSV2WithDefaults

`func NewApiNVD20TemporalCVSSV2WithDefaults() *ApiNVD20TemporalCVSSV2`

NewApiNVD20TemporalCVSSV2WithDefaults instantiates a new ApiNVD20TemporalCVSSV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedBaseMetricV2

`func (o *ApiNVD20TemporalCVSSV2) GetAssociatedBaseMetricV2() ApiNVD20TemporalAssociatedBaseMetric`

GetAssociatedBaseMetricV2 returns the AssociatedBaseMetricV2 field if non-nil, zero value otherwise.

### GetAssociatedBaseMetricV2Ok

`func (o *ApiNVD20TemporalCVSSV2) GetAssociatedBaseMetricV2Ok() (*ApiNVD20TemporalAssociatedBaseMetric, bool)`

GetAssociatedBaseMetricV2Ok returns a tuple with the AssociatedBaseMetricV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedBaseMetricV2

`func (o *ApiNVD20TemporalCVSSV2) SetAssociatedBaseMetricV2(v ApiNVD20TemporalAssociatedBaseMetric)`

SetAssociatedBaseMetricV2 sets AssociatedBaseMetricV2 field to given value.

### HasAssociatedBaseMetricV2

`func (o *ApiNVD20TemporalCVSSV2) HasAssociatedBaseMetricV2() bool`

HasAssociatedBaseMetricV2 returns a boolean if a field has been set.

### GetExploitability

`func (o *ApiNVD20TemporalCVSSV2) GetExploitability() string`

GetExploitability returns the Exploitability field if non-nil, zero value otherwise.

### GetExploitabilityOk

`func (o *ApiNVD20TemporalCVSSV2) GetExploitabilityOk() (*string, bool)`

GetExploitabilityOk returns a tuple with the Exploitability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitability

`func (o *ApiNVD20TemporalCVSSV2) SetExploitability(v string)`

SetExploitability sets Exploitability field to given value.

### HasExploitability

`func (o *ApiNVD20TemporalCVSSV2) HasExploitability() bool`

HasExploitability returns a boolean if a field has been set.

### GetRemediationLevel

`func (o *ApiNVD20TemporalCVSSV2) GetRemediationLevel() string`

GetRemediationLevel returns the RemediationLevel field if non-nil, zero value otherwise.

### GetRemediationLevelOk

`func (o *ApiNVD20TemporalCVSSV2) GetRemediationLevelOk() (*string, bool)`

GetRemediationLevelOk returns a tuple with the RemediationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationLevel

`func (o *ApiNVD20TemporalCVSSV2) SetRemediationLevel(v string)`

SetRemediationLevel sets RemediationLevel field to given value.

### HasRemediationLevel

`func (o *ApiNVD20TemporalCVSSV2) HasRemediationLevel() bool`

HasRemediationLevel returns a boolean if a field has been set.

### GetReportConfidence

`func (o *ApiNVD20TemporalCVSSV2) GetReportConfidence() string`

GetReportConfidence returns the ReportConfidence field if non-nil, zero value otherwise.

### GetReportConfidenceOk

`func (o *ApiNVD20TemporalCVSSV2) GetReportConfidenceOk() (*string, bool)`

GetReportConfidenceOk returns a tuple with the ReportConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportConfidence

`func (o *ApiNVD20TemporalCVSSV2) SetReportConfidence(v string)`

SetReportConfidence sets ReportConfidence field to given value.

### HasReportConfidence

`func (o *ApiNVD20TemporalCVSSV2) HasReportConfidence() bool`

HasReportConfidence returns a boolean if a field has been set.

### GetTemporalScore

`func (o *ApiNVD20TemporalCVSSV2) GetTemporalScore() float32`

GetTemporalScore returns the TemporalScore field if non-nil, zero value otherwise.

### GetTemporalScoreOk

`func (o *ApiNVD20TemporalCVSSV2) GetTemporalScoreOk() (*float32, bool)`

GetTemporalScoreOk returns a tuple with the TemporalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalScore

`func (o *ApiNVD20TemporalCVSSV2) SetTemporalScore(v float32)`

SetTemporalScore sets TemporalScore field to given value.

### HasTemporalScore

`func (o *ApiNVD20TemporalCVSSV2) HasTemporalScore() bool`

HasTemporalScore returns a boolean if a field has been set.

### GetVectorString

`func (o *ApiNVD20TemporalCVSSV2) GetVectorString() string`

GetVectorString returns the VectorString field if non-nil, zero value otherwise.

### GetVectorStringOk

`func (o *ApiNVD20TemporalCVSSV2) GetVectorStringOk() (*string, bool)`

GetVectorStringOk returns a tuple with the VectorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorString

`func (o *ApiNVD20TemporalCVSSV2) SetVectorString(v string)`

SetVectorString sets VectorString field to given value.

### HasVectorString

`func (o *ApiNVD20TemporalCVSSV2) HasVectorString() bool`

HasVectorString returns a boolean if a field has been set.

### GetVersion

`func (o *ApiNVD20TemporalCVSSV2) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiNVD20TemporalCVSSV2) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiNVD20TemporalCVSSV2) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApiNVD20TemporalCVSSV2) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


