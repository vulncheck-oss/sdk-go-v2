# ApiNVD20CvssDataV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessComplexity** | Pointer to **string** |  | [optional] 
**AccessVector** | Pointer to **string** |  | [optional] 
**Authentication** | Pointer to **string** |  | [optional] 
**AvailabilityImpact** | Pointer to **string** |  | [optional] 
**AvailabilityRequirement** | Pointer to **string** |  | [optional] 
**BaseScore** | Pointer to **float32** |  | [optional] 
**CollateralDamagePotential** | Pointer to **string** |  | [optional] 
**ConfidentialityImpact** | Pointer to **string** |  | [optional] 
**ConfidentialityRequirement** | Pointer to **string** |  | [optional] 
**EnvironmentalScore** | Pointer to **float32** |  | [optional] 
**Exploitability** | Pointer to **string** |  | [optional] 
**IntegrityImpact** | Pointer to **string** |  | [optional] 
**IntegrityRequirement** | Pointer to **string** |  | [optional] 
**RemediationLevel** | Pointer to **string** |  | [optional] 
**ReportConfidence** | Pointer to **string** |  | [optional] 
**TargetDistribution** | Pointer to **string** |  | [optional] 
**TemporalScore** | Pointer to **float32** |  | [optional] 
**VectorString** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewApiNVD20CvssDataV2

`func NewApiNVD20CvssDataV2() *ApiNVD20CvssDataV2`

NewApiNVD20CvssDataV2 instantiates a new ApiNVD20CvssDataV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNVD20CvssDataV2WithDefaults

`func NewApiNVD20CvssDataV2WithDefaults() *ApiNVD20CvssDataV2`

NewApiNVD20CvssDataV2WithDefaults instantiates a new ApiNVD20CvssDataV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessComplexity

`func (o *ApiNVD20CvssDataV2) GetAccessComplexity() string`

GetAccessComplexity returns the AccessComplexity field if non-nil, zero value otherwise.

### GetAccessComplexityOk

`func (o *ApiNVD20CvssDataV2) GetAccessComplexityOk() (*string, bool)`

GetAccessComplexityOk returns a tuple with the AccessComplexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessComplexity

`func (o *ApiNVD20CvssDataV2) SetAccessComplexity(v string)`

SetAccessComplexity sets AccessComplexity field to given value.

### HasAccessComplexity

`func (o *ApiNVD20CvssDataV2) HasAccessComplexity() bool`

HasAccessComplexity returns a boolean if a field has been set.

### GetAccessVector

`func (o *ApiNVD20CvssDataV2) GetAccessVector() string`

GetAccessVector returns the AccessVector field if non-nil, zero value otherwise.

### GetAccessVectorOk

`func (o *ApiNVD20CvssDataV2) GetAccessVectorOk() (*string, bool)`

GetAccessVectorOk returns a tuple with the AccessVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessVector

`func (o *ApiNVD20CvssDataV2) SetAccessVector(v string)`

SetAccessVector sets AccessVector field to given value.

### HasAccessVector

`func (o *ApiNVD20CvssDataV2) HasAccessVector() bool`

HasAccessVector returns a boolean if a field has been set.

### GetAuthentication

`func (o *ApiNVD20CvssDataV2) GetAuthentication() string`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *ApiNVD20CvssDataV2) GetAuthenticationOk() (*string, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *ApiNVD20CvssDataV2) SetAuthentication(v string)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *ApiNVD20CvssDataV2) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetAvailabilityImpact

`func (o *ApiNVD20CvssDataV2) GetAvailabilityImpact() string`

GetAvailabilityImpact returns the AvailabilityImpact field if non-nil, zero value otherwise.

### GetAvailabilityImpactOk

`func (o *ApiNVD20CvssDataV2) GetAvailabilityImpactOk() (*string, bool)`

GetAvailabilityImpactOk returns a tuple with the AvailabilityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityImpact

`func (o *ApiNVD20CvssDataV2) SetAvailabilityImpact(v string)`

SetAvailabilityImpact sets AvailabilityImpact field to given value.

### HasAvailabilityImpact

`func (o *ApiNVD20CvssDataV2) HasAvailabilityImpact() bool`

HasAvailabilityImpact returns a boolean if a field has been set.

### GetAvailabilityRequirement

`func (o *ApiNVD20CvssDataV2) GetAvailabilityRequirement() string`

GetAvailabilityRequirement returns the AvailabilityRequirement field if non-nil, zero value otherwise.

### GetAvailabilityRequirementOk

`func (o *ApiNVD20CvssDataV2) GetAvailabilityRequirementOk() (*string, bool)`

GetAvailabilityRequirementOk returns a tuple with the AvailabilityRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityRequirement

`func (o *ApiNVD20CvssDataV2) SetAvailabilityRequirement(v string)`

SetAvailabilityRequirement sets AvailabilityRequirement field to given value.

### HasAvailabilityRequirement

`func (o *ApiNVD20CvssDataV2) HasAvailabilityRequirement() bool`

HasAvailabilityRequirement returns a boolean if a field has been set.

### GetBaseScore

`func (o *ApiNVD20CvssDataV2) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *ApiNVD20CvssDataV2) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *ApiNVD20CvssDataV2) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *ApiNVD20CvssDataV2) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetCollateralDamagePotential

`func (o *ApiNVD20CvssDataV2) GetCollateralDamagePotential() string`

GetCollateralDamagePotential returns the CollateralDamagePotential field if non-nil, zero value otherwise.

### GetCollateralDamagePotentialOk

`func (o *ApiNVD20CvssDataV2) GetCollateralDamagePotentialOk() (*string, bool)`

GetCollateralDamagePotentialOk returns a tuple with the CollateralDamagePotential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralDamagePotential

`func (o *ApiNVD20CvssDataV2) SetCollateralDamagePotential(v string)`

SetCollateralDamagePotential sets CollateralDamagePotential field to given value.

### HasCollateralDamagePotential

`func (o *ApiNVD20CvssDataV2) HasCollateralDamagePotential() bool`

HasCollateralDamagePotential returns a boolean if a field has been set.

### GetConfidentialityImpact

`func (o *ApiNVD20CvssDataV2) GetConfidentialityImpact() string`

GetConfidentialityImpact returns the ConfidentialityImpact field if non-nil, zero value otherwise.

### GetConfidentialityImpactOk

`func (o *ApiNVD20CvssDataV2) GetConfidentialityImpactOk() (*string, bool)`

GetConfidentialityImpactOk returns a tuple with the ConfidentialityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialityImpact

`func (o *ApiNVD20CvssDataV2) SetConfidentialityImpact(v string)`

SetConfidentialityImpact sets ConfidentialityImpact field to given value.

### HasConfidentialityImpact

`func (o *ApiNVD20CvssDataV2) HasConfidentialityImpact() bool`

HasConfidentialityImpact returns a boolean if a field has been set.

### GetConfidentialityRequirement

`func (o *ApiNVD20CvssDataV2) GetConfidentialityRequirement() string`

GetConfidentialityRequirement returns the ConfidentialityRequirement field if non-nil, zero value otherwise.

### GetConfidentialityRequirementOk

`func (o *ApiNVD20CvssDataV2) GetConfidentialityRequirementOk() (*string, bool)`

GetConfidentialityRequirementOk returns a tuple with the ConfidentialityRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialityRequirement

`func (o *ApiNVD20CvssDataV2) SetConfidentialityRequirement(v string)`

SetConfidentialityRequirement sets ConfidentialityRequirement field to given value.

### HasConfidentialityRequirement

`func (o *ApiNVD20CvssDataV2) HasConfidentialityRequirement() bool`

HasConfidentialityRequirement returns a boolean if a field has been set.

### GetEnvironmentalScore

`func (o *ApiNVD20CvssDataV2) GetEnvironmentalScore() float32`

GetEnvironmentalScore returns the EnvironmentalScore field if non-nil, zero value otherwise.

### GetEnvironmentalScoreOk

`func (o *ApiNVD20CvssDataV2) GetEnvironmentalScoreOk() (*float32, bool)`

GetEnvironmentalScoreOk returns a tuple with the EnvironmentalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentalScore

`func (o *ApiNVD20CvssDataV2) SetEnvironmentalScore(v float32)`

SetEnvironmentalScore sets EnvironmentalScore field to given value.

### HasEnvironmentalScore

`func (o *ApiNVD20CvssDataV2) HasEnvironmentalScore() bool`

HasEnvironmentalScore returns a boolean if a field has been set.

### GetExploitability

`func (o *ApiNVD20CvssDataV2) GetExploitability() string`

GetExploitability returns the Exploitability field if non-nil, zero value otherwise.

### GetExploitabilityOk

`func (o *ApiNVD20CvssDataV2) GetExploitabilityOk() (*string, bool)`

GetExploitabilityOk returns a tuple with the Exploitability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitability

`func (o *ApiNVD20CvssDataV2) SetExploitability(v string)`

SetExploitability sets Exploitability field to given value.

### HasExploitability

`func (o *ApiNVD20CvssDataV2) HasExploitability() bool`

HasExploitability returns a boolean if a field has been set.

### GetIntegrityImpact

`func (o *ApiNVD20CvssDataV2) GetIntegrityImpact() string`

GetIntegrityImpact returns the IntegrityImpact field if non-nil, zero value otherwise.

### GetIntegrityImpactOk

`func (o *ApiNVD20CvssDataV2) GetIntegrityImpactOk() (*string, bool)`

GetIntegrityImpactOk returns a tuple with the IntegrityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityImpact

`func (o *ApiNVD20CvssDataV2) SetIntegrityImpact(v string)`

SetIntegrityImpact sets IntegrityImpact field to given value.

### HasIntegrityImpact

`func (o *ApiNVD20CvssDataV2) HasIntegrityImpact() bool`

HasIntegrityImpact returns a boolean if a field has been set.

### GetIntegrityRequirement

`func (o *ApiNVD20CvssDataV2) GetIntegrityRequirement() string`

GetIntegrityRequirement returns the IntegrityRequirement field if non-nil, zero value otherwise.

### GetIntegrityRequirementOk

`func (o *ApiNVD20CvssDataV2) GetIntegrityRequirementOk() (*string, bool)`

GetIntegrityRequirementOk returns a tuple with the IntegrityRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityRequirement

`func (o *ApiNVD20CvssDataV2) SetIntegrityRequirement(v string)`

SetIntegrityRequirement sets IntegrityRequirement field to given value.

### HasIntegrityRequirement

`func (o *ApiNVD20CvssDataV2) HasIntegrityRequirement() bool`

HasIntegrityRequirement returns a boolean if a field has been set.

### GetRemediationLevel

`func (o *ApiNVD20CvssDataV2) GetRemediationLevel() string`

GetRemediationLevel returns the RemediationLevel field if non-nil, zero value otherwise.

### GetRemediationLevelOk

`func (o *ApiNVD20CvssDataV2) GetRemediationLevelOk() (*string, bool)`

GetRemediationLevelOk returns a tuple with the RemediationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationLevel

`func (o *ApiNVD20CvssDataV2) SetRemediationLevel(v string)`

SetRemediationLevel sets RemediationLevel field to given value.

### HasRemediationLevel

`func (o *ApiNVD20CvssDataV2) HasRemediationLevel() bool`

HasRemediationLevel returns a boolean if a field has been set.

### GetReportConfidence

`func (o *ApiNVD20CvssDataV2) GetReportConfidence() string`

GetReportConfidence returns the ReportConfidence field if non-nil, zero value otherwise.

### GetReportConfidenceOk

`func (o *ApiNVD20CvssDataV2) GetReportConfidenceOk() (*string, bool)`

GetReportConfidenceOk returns a tuple with the ReportConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportConfidence

`func (o *ApiNVD20CvssDataV2) SetReportConfidence(v string)`

SetReportConfidence sets ReportConfidence field to given value.

### HasReportConfidence

`func (o *ApiNVD20CvssDataV2) HasReportConfidence() bool`

HasReportConfidence returns a boolean if a field has been set.

### GetTargetDistribution

`func (o *ApiNVD20CvssDataV2) GetTargetDistribution() string`

GetTargetDistribution returns the TargetDistribution field if non-nil, zero value otherwise.

### GetTargetDistributionOk

`func (o *ApiNVD20CvssDataV2) GetTargetDistributionOk() (*string, bool)`

GetTargetDistributionOk returns a tuple with the TargetDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDistribution

`func (o *ApiNVD20CvssDataV2) SetTargetDistribution(v string)`

SetTargetDistribution sets TargetDistribution field to given value.

### HasTargetDistribution

`func (o *ApiNVD20CvssDataV2) HasTargetDistribution() bool`

HasTargetDistribution returns a boolean if a field has been set.

### GetTemporalScore

`func (o *ApiNVD20CvssDataV2) GetTemporalScore() float32`

GetTemporalScore returns the TemporalScore field if non-nil, zero value otherwise.

### GetTemporalScoreOk

`func (o *ApiNVD20CvssDataV2) GetTemporalScoreOk() (*float32, bool)`

GetTemporalScoreOk returns a tuple with the TemporalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalScore

`func (o *ApiNVD20CvssDataV2) SetTemporalScore(v float32)`

SetTemporalScore sets TemporalScore field to given value.

### HasTemporalScore

`func (o *ApiNVD20CvssDataV2) HasTemporalScore() bool`

HasTemporalScore returns a boolean if a field has been set.

### GetVectorString

`func (o *ApiNVD20CvssDataV2) GetVectorString() string`

GetVectorString returns the VectorString field if non-nil, zero value otherwise.

### GetVectorStringOk

`func (o *ApiNVD20CvssDataV2) GetVectorStringOk() (*string, bool)`

GetVectorStringOk returns a tuple with the VectorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorString

`func (o *ApiNVD20CvssDataV2) SetVectorString(v string)`

SetVectorString sets VectorString field to given value.

### HasVectorString

`func (o *ApiNVD20CvssDataV2) HasVectorString() bool`

HasVectorString returns a boolean if a field has been set.

### GetVersion

`func (o *ApiNVD20CvssDataV2) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiNVD20CvssDataV2) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiNVD20CvssDataV2) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApiNVD20CvssDataV2) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


