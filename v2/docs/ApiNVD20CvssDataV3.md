# ApiNVD20CvssDataV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttackComplexity** | Pointer to **string** |  | [optional] 
**AttackVector** | Pointer to **string** |  | [optional] 
**AvailabilityImpact** | Pointer to **string** |  | [optional] 
**AvailabilityRequirement** | Pointer to **string** |  | [optional] 
**BaseScore** | Pointer to **float32** |  | [optional] 
**BaseSeverity** | Pointer to **string** |  | [optional] 
**ConfidentialityImpact** | Pointer to **string** |  | [optional] 
**ConfidentialityRequirement** | Pointer to **string** |  | [optional] 
**EnvironmentalScore** | Pointer to **float32** |  | [optional] 
**EnvironmentalSeverity** | Pointer to **string** |  | [optional] 
**ExploitCodeMaturity** | Pointer to **string** |  | [optional] 
**IntegrityImpact** | Pointer to **string** |  | [optional] 
**IntegrityRequirement** | Pointer to **string** |  | [optional] 
**ModifiedAttackComplexity** | Pointer to **string** |  | [optional] 
**ModifiedAttackVector** | Pointer to **string** |  | [optional] 
**ModifiedAvailabilityImpact** | Pointer to **string** |  | [optional] 
**ModifiedConfidentialityImpact** | Pointer to **string** |  | [optional] 
**ModifiedIntegrityImpact** | Pointer to **string** |  | [optional] 
**ModifiedPrivilegesRequired** | Pointer to **string** |  | [optional] 
**ModifiedScope** | Pointer to **string** |  | [optional] 
**ModifiedUserInteraction** | Pointer to **string** |  | [optional] 
**PrivilegesRequired** | Pointer to **string** |  | [optional] 
**RemediationLevel** | Pointer to **string** |  | [optional] 
**ReportConfidence** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**TemporalScore** | Pointer to **float32** |  | [optional] 
**TemporalSeverity** | Pointer to **string** |  | [optional] 
**UserInteraction** | Pointer to **string** |  | [optional] 
**VectorString** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewApiNVD20CvssDataV3

`func NewApiNVD20CvssDataV3() *ApiNVD20CvssDataV3`

NewApiNVD20CvssDataV3 instantiates a new ApiNVD20CvssDataV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNVD20CvssDataV3WithDefaults

`func NewApiNVD20CvssDataV3WithDefaults() *ApiNVD20CvssDataV3`

NewApiNVD20CvssDataV3WithDefaults instantiates a new ApiNVD20CvssDataV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttackComplexity

`func (o *ApiNVD20CvssDataV3) GetAttackComplexity() string`

GetAttackComplexity returns the AttackComplexity field if non-nil, zero value otherwise.

### GetAttackComplexityOk

`func (o *ApiNVD20CvssDataV3) GetAttackComplexityOk() (*string, bool)`

GetAttackComplexityOk returns a tuple with the AttackComplexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackComplexity

`func (o *ApiNVD20CvssDataV3) SetAttackComplexity(v string)`

SetAttackComplexity sets AttackComplexity field to given value.

### HasAttackComplexity

`func (o *ApiNVD20CvssDataV3) HasAttackComplexity() bool`

HasAttackComplexity returns a boolean if a field has been set.

### GetAttackVector

`func (o *ApiNVD20CvssDataV3) GetAttackVector() string`

GetAttackVector returns the AttackVector field if non-nil, zero value otherwise.

### GetAttackVectorOk

`func (o *ApiNVD20CvssDataV3) GetAttackVectorOk() (*string, bool)`

GetAttackVectorOk returns a tuple with the AttackVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackVector

`func (o *ApiNVD20CvssDataV3) SetAttackVector(v string)`

SetAttackVector sets AttackVector field to given value.

### HasAttackVector

`func (o *ApiNVD20CvssDataV3) HasAttackVector() bool`

HasAttackVector returns a boolean if a field has been set.

### GetAvailabilityImpact

`func (o *ApiNVD20CvssDataV3) GetAvailabilityImpact() string`

GetAvailabilityImpact returns the AvailabilityImpact field if non-nil, zero value otherwise.

### GetAvailabilityImpactOk

`func (o *ApiNVD20CvssDataV3) GetAvailabilityImpactOk() (*string, bool)`

GetAvailabilityImpactOk returns a tuple with the AvailabilityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityImpact

`func (o *ApiNVD20CvssDataV3) SetAvailabilityImpact(v string)`

SetAvailabilityImpact sets AvailabilityImpact field to given value.

### HasAvailabilityImpact

`func (o *ApiNVD20CvssDataV3) HasAvailabilityImpact() bool`

HasAvailabilityImpact returns a boolean if a field has been set.

### GetAvailabilityRequirement

`func (o *ApiNVD20CvssDataV3) GetAvailabilityRequirement() string`

GetAvailabilityRequirement returns the AvailabilityRequirement field if non-nil, zero value otherwise.

### GetAvailabilityRequirementOk

`func (o *ApiNVD20CvssDataV3) GetAvailabilityRequirementOk() (*string, bool)`

GetAvailabilityRequirementOk returns a tuple with the AvailabilityRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityRequirement

`func (o *ApiNVD20CvssDataV3) SetAvailabilityRequirement(v string)`

SetAvailabilityRequirement sets AvailabilityRequirement field to given value.

### HasAvailabilityRequirement

`func (o *ApiNVD20CvssDataV3) HasAvailabilityRequirement() bool`

HasAvailabilityRequirement returns a boolean if a field has been set.

### GetBaseScore

`func (o *ApiNVD20CvssDataV3) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *ApiNVD20CvssDataV3) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *ApiNVD20CvssDataV3) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *ApiNVD20CvssDataV3) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetBaseSeverity

`func (o *ApiNVD20CvssDataV3) GetBaseSeverity() string`

GetBaseSeverity returns the BaseSeverity field if non-nil, zero value otherwise.

### GetBaseSeverityOk

`func (o *ApiNVD20CvssDataV3) GetBaseSeverityOk() (*string, bool)`

GetBaseSeverityOk returns a tuple with the BaseSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSeverity

`func (o *ApiNVD20CvssDataV3) SetBaseSeverity(v string)`

SetBaseSeverity sets BaseSeverity field to given value.

### HasBaseSeverity

`func (o *ApiNVD20CvssDataV3) HasBaseSeverity() bool`

HasBaseSeverity returns a boolean if a field has been set.

### GetConfidentialityImpact

`func (o *ApiNVD20CvssDataV3) GetConfidentialityImpact() string`

GetConfidentialityImpact returns the ConfidentialityImpact field if non-nil, zero value otherwise.

### GetConfidentialityImpactOk

`func (o *ApiNVD20CvssDataV3) GetConfidentialityImpactOk() (*string, bool)`

GetConfidentialityImpactOk returns a tuple with the ConfidentialityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialityImpact

`func (o *ApiNVD20CvssDataV3) SetConfidentialityImpact(v string)`

SetConfidentialityImpact sets ConfidentialityImpact field to given value.

### HasConfidentialityImpact

`func (o *ApiNVD20CvssDataV3) HasConfidentialityImpact() bool`

HasConfidentialityImpact returns a boolean if a field has been set.

### GetConfidentialityRequirement

`func (o *ApiNVD20CvssDataV3) GetConfidentialityRequirement() string`

GetConfidentialityRequirement returns the ConfidentialityRequirement field if non-nil, zero value otherwise.

### GetConfidentialityRequirementOk

`func (o *ApiNVD20CvssDataV3) GetConfidentialityRequirementOk() (*string, bool)`

GetConfidentialityRequirementOk returns a tuple with the ConfidentialityRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialityRequirement

`func (o *ApiNVD20CvssDataV3) SetConfidentialityRequirement(v string)`

SetConfidentialityRequirement sets ConfidentialityRequirement field to given value.

### HasConfidentialityRequirement

`func (o *ApiNVD20CvssDataV3) HasConfidentialityRequirement() bool`

HasConfidentialityRequirement returns a boolean if a field has been set.

### GetEnvironmentalScore

`func (o *ApiNVD20CvssDataV3) GetEnvironmentalScore() float32`

GetEnvironmentalScore returns the EnvironmentalScore field if non-nil, zero value otherwise.

### GetEnvironmentalScoreOk

`func (o *ApiNVD20CvssDataV3) GetEnvironmentalScoreOk() (*float32, bool)`

GetEnvironmentalScoreOk returns a tuple with the EnvironmentalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentalScore

`func (o *ApiNVD20CvssDataV3) SetEnvironmentalScore(v float32)`

SetEnvironmentalScore sets EnvironmentalScore field to given value.

### HasEnvironmentalScore

`func (o *ApiNVD20CvssDataV3) HasEnvironmentalScore() bool`

HasEnvironmentalScore returns a boolean if a field has been set.

### GetEnvironmentalSeverity

`func (o *ApiNVD20CvssDataV3) GetEnvironmentalSeverity() string`

GetEnvironmentalSeverity returns the EnvironmentalSeverity field if non-nil, zero value otherwise.

### GetEnvironmentalSeverityOk

`func (o *ApiNVD20CvssDataV3) GetEnvironmentalSeverityOk() (*string, bool)`

GetEnvironmentalSeverityOk returns a tuple with the EnvironmentalSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentalSeverity

`func (o *ApiNVD20CvssDataV3) SetEnvironmentalSeverity(v string)`

SetEnvironmentalSeverity sets EnvironmentalSeverity field to given value.

### HasEnvironmentalSeverity

`func (o *ApiNVD20CvssDataV3) HasEnvironmentalSeverity() bool`

HasEnvironmentalSeverity returns a boolean if a field has been set.

### GetExploitCodeMaturity

`func (o *ApiNVD20CvssDataV3) GetExploitCodeMaturity() string`

GetExploitCodeMaturity returns the ExploitCodeMaturity field if non-nil, zero value otherwise.

### GetExploitCodeMaturityOk

`func (o *ApiNVD20CvssDataV3) GetExploitCodeMaturityOk() (*string, bool)`

GetExploitCodeMaturityOk returns a tuple with the ExploitCodeMaturity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitCodeMaturity

`func (o *ApiNVD20CvssDataV3) SetExploitCodeMaturity(v string)`

SetExploitCodeMaturity sets ExploitCodeMaturity field to given value.

### HasExploitCodeMaturity

`func (o *ApiNVD20CvssDataV3) HasExploitCodeMaturity() bool`

HasExploitCodeMaturity returns a boolean if a field has been set.

### GetIntegrityImpact

`func (o *ApiNVD20CvssDataV3) GetIntegrityImpact() string`

GetIntegrityImpact returns the IntegrityImpact field if non-nil, zero value otherwise.

### GetIntegrityImpactOk

`func (o *ApiNVD20CvssDataV3) GetIntegrityImpactOk() (*string, bool)`

GetIntegrityImpactOk returns a tuple with the IntegrityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityImpact

`func (o *ApiNVD20CvssDataV3) SetIntegrityImpact(v string)`

SetIntegrityImpact sets IntegrityImpact field to given value.

### HasIntegrityImpact

`func (o *ApiNVD20CvssDataV3) HasIntegrityImpact() bool`

HasIntegrityImpact returns a boolean if a field has been set.

### GetIntegrityRequirement

`func (o *ApiNVD20CvssDataV3) GetIntegrityRequirement() string`

GetIntegrityRequirement returns the IntegrityRequirement field if non-nil, zero value otherwise.

### GetIntegrityRequirementOk

`func (o *ApiNVD20CvssDataV3) GetIntegrityRequirementOk() (*string, bool)`

GetIntegrityRequirementOk returns a tuple with the IntegrityRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityRequirement

`func (o *ApiNVD20CvssDataV3) SetIntegrityRequirement(v string)`

SetIntegrityRequirement sets IntegrityRequirement field to given value.

### HasIntegrityRequirement

`func (o *ApiNVD20CvssDataV3) HasIntegrityRequirement() bool`

HasIntegrityRequirement returns a boolean if a field has been set.

### GetModifiedAttackComplexity

`func (o *ApiNVD20CvssDataV3) GetModifiedAttackComplexity() string`

GetModifiedAttackComplexity returns the ModifiedAttackComplexity field if non-nil, zero value otherwise.

### GetModifiedAttackComplexityOk

`func (o *ApiNVD20CvssDataV3) GetModifiedAttackComplexityOk() (*string, bool)`

GetModifiedAttackComplexityOk returns a tuple with the ModifiedAttackComplexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAttackComplexity

`func (o *ApiNVD20CvssDataV3) SetModifiedAttackComplexity(v string)`

SetModifiedAttackComplexity sets ModifiedAttackComplexity field to given value.

### HasModifiedAttackComplexity

`func (o *ApiNVD20CvssDataV3) HasModifiedAttackComplexity() bool`

HasModifiedAttackComplexity returns a boolean if a field has been set.

### GetModifiedAttackVector

`func (o *ApiNVD20CvssDataV3) GetModifiedAttackVector() string`

GetModifiedAttackVector returns the ModifiedAttackVector field if non-nil, zero value otherwise.

### GetModifiedAttackVectorOk

`func (o *ApiNVD20CvssDataV3) GetModifiedAttackVectorOk() (*string, bool)`

GetModifiedAttackVectorOk returns a tuple with the ModifiedAttackVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAttackVector

`func (o *ApiNVD20CvssDataV3) SetModifiedAttackVector(v string)`

SetModifiedAttackVector sets ModifiedAttackVector field to given value.

### HasModifiedAttackVector

`func (o *ApiNVD20CvssDataV3) HasModifiedAttackVector() bool`

HasModifiedAttackVector returns a boolean if a field has been set.

### GetModifiedAvailabilityImpact

`func (o *ApiNVD20CvssDataV3) GetModifiedAvailabilityImpact() string`

GetModifiedAvailabilityImpact returns the ModifiedAvailabilityImpact field if non-nil, zero value otherwise.

### GetModifiedAvailabilityImpactOk

`func (o *ApiNVD20CvssDataV3) GetModifiedAvailabilityImpactOk() (*string, bool)`

GetModifiedAvailabilityImpactOk returns a tuple with the ModifiedAvailabilityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAvailabilityImpact

`func (o *ApiNVD20CvssDataV3) SetModifiedAvailabilityImpact(v string)`

SetModifiedAvailabilityImpact sets ModifiedAvailabilityImpact field to given value.

### HasModifiedAvailabilityImpact

`func (o *ApiNVD20CvssDataV3) HasModifiedAvailabilityImpact() bool`

HasModifiedAvailabilityImpact returns a boolean if a field has been set.

### GetModifiedConfidentialityImpact

`func (o *ApiNVD20CvssDataV3) GetModifiedConfidentialityImpact() string`

GetModifiedConfidentialityImpact returns the ModifiedConfidentialityImpact field if non-nil, zero value otherwise.

### GetModifiedConfidentialityImpactOk

`func (o *ApiNVD20CvssDataV3) GetModifiedConfidentialityImpactOk() (*string, bool)`

GetModifiedConfidentialityImpactOk returns a tuple with the ModifiedConfidentialityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedConfidentialityImpact

`func (o *ApiNVD20CvssDataV3) SetModifiedConfidentialityImpact(v string)`

SetModifiedConfidentialityImpact sets ModifiedConfidentialityImpact field to given value.

### HasModifiedConfidentialityImpact

`func (o *ApiNVD20CvssDataV3) HasModifiedConfidentialityImpact() bool`

HasModifiedConfidentialityImpact returns a boolean if a field has been set.

### GetModifiedIntegrityImpact

`func (o *ApiNVD20CvssDataV3) GetModifiedIntegrityImpact() string`

GetModifiedIntegrityImpact returns the ModifiedIntegrityImpact field if non-nil, zero value otherwise.

### GetModifiedIntegrityImpactOk

`func (o *ApiNVD20CvssDataV3) GetModifiedIntegrityImpactOk() (*string, bool)`

GetModifiedIntegrityImpactOk returns a tuple with the ModifiedIntegrityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedIntegrityImpact

`func (o *ApiNVD20CvssDataV3) SetModifiedIntegrityImpact(v string)`

SetModifiedIntegrityImpact sets ModifiedIntegrityImpact field to given value.

### HasModifiedIntegrityImpact

`func (o *ApiNVD20CvssDataV3) HasModifiedIntegrityImpact() bool`

HasModifiedIntegrityImpact returns a boolean if a field has been set.

### GetModifiedPrivilegesRequired

`func (o *ApiNVD20CvssDataV3) GetModifiedPrivilegesRequired() string`

GetModifiedPrivilegesRequired returns the ModifiedPrivilegesRequired field if non-nil, zero value otherwise.

### GetModifiedPrivilegesRequiredOk

`func (o *ApiNVD20CvssDataV3) GetModifiedPrivilegesRequiredOk() (*string, bool)`

GetModifiedPrivilegesRequiredOk returns a tuple with the ModifiedPrivilegesRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedPrivilegesRequired

`func (o *ApiNVD20CvssDataV3) SetModifiedPrivilegesRequired(v string)`

SetModifiedPrivilegesRequired sets ModifiedPrivilegesRequired field to given value.

### HasModifiedPrivilegesRequired

`func (o *ApiNVD20CvssDataV3) HasModifiedPrivilegesRequired() bool`

HasModifiedPrivilegesRequired returns a boolean if a field has been set.

### GetModifiedScope

`func (o *ApiNVD20CvssDataV3) GetModifiedScope() string`

GetModifiedScope returns the ModifiedScope field if non-nil, zero value otherwise.

### GetModifiedScopeOk

`func (o *ApiNVD20CvssDataV3) GetModifiedScopeOk() (*string, bool)`

GetModifiedScopeOk returns a tuple with the ModifiedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedScope

`func (o *ApiNVD20CvssDataV3) SetModifiedScope(v string)`

SetModifiedScope sets ModifiedScope field to given value.

### HasModifiedScope

`func (o *ApiNVD20CvssDataV3) HasModifiedScope() bool`

HasModifiedScope returns a boolean if a field has been set.

### GetModifiedUserInteraction

`func (o *ApiNVD20CvssDataV3) GetModifiedUserInteraction() string`

GetModifiedUserInteraction returns the ModifiedUserInteraction field if non-nil, zero value otherwise.

### GetModifiedUserInteractionOk

`func (o *ApiNVD20CvssDataV3) GetModifiedUserInteractionOk() (*string, bool)`

GetModifiedUserInteractionOk returns a tuple with the ModifiedUserInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedUserInteraction

`func (o *ApiNVD20CvssDataV3) SetModifiedUserInteraction(v string)`

SetModifiedUserInteraction sets ModifiedUserInteraction field to given value.

### HasModifiedUserInteraction

`func (o *ApiNVD20CvssDataV3) HasModifiedUserInteraction() bool`

HasModifiedUserInteraction returns a boolean if a field has been set.

### GetPrivilegesRequired

`func (o *ApiNVD20CvssDataV3) GetPrivilegesRequired() string`

GetPrivilegesRequired returns the PrivilegesRequired field if non-nil, zero value otherwise.

### GetPrivilegesRequiredOk

`func (o *ApiNVD20CvssDataV3) GetPrivilegesRequiredOk() (*string, bool)`

GetPrivilegesRequiredOk returns a tuple with the PrivilegesRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegesRequired

`func (o *ApiNVD20CvssDataV3) SetPrivilegesRequired(v string)`

SetPrivilegesRequired sets PrivilegesRequired field to given value.

### HasPrivilegesRequired

`func (o *ApiNVD20CvssDataV3) HasPrivilegesRequired() bool`

HasPrivilegesRequired returns a boolean if a field has been set.

### GetRemediationLevel

`func (o *ApiNVD20CvssDataV3) GetRemediationLevel() string`

GetRemediationLevel returns the RemediationLevel field if non-nil, zero value otherwise.

### GetRemediationLevelOk

`func (o *ApiNVD20CvssDataV3) GetRemediationLevelOk() (*string, bool)`

GetRemediationLevelOk returns a tuple with the RemediationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationLevel

`func (o *ApiNVD20CvssDataV3) SetRemediationLevel(v string)`

SetRemediationLevel sets RemediationLevel field to given value.

### HasRemediationLevel

`func (o *ApiNVD20CvssDataV3) HasRemediationLevel() bool`

HasRemediationLevel returns a boolean if a field has been set.

### GetReportConfidence

`func (o *ApiNVD20CvssDataV3) GetReportConfidence() string`

GetReportConfidence returns the ReportConfidence field if non-nil, zero value otherwise.

### GetReportConfidenceOk

`func (o *ApiNVD20CvssDataV3) GetReportConfidenceOk() (*string, bool)`

GetReportConfidenceOk returns a tuple with the ReportConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportConfidence

`func (o *ApiNVD20CvssDataV3) SetReportConfidence(v string)`

SetReportConfidence sets ReportConfidence field to given value.

### HasReportConfidence

`func (o *ApiNVD20CvssDataV3) HasReportConfidence() bool`

HasReportConfidence returns a boolean if a field has been set.

### GetScope

`func (o *ApiNVD20CvssDataV3) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ApiNVD20CvssDataV3) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ApiNVD20CvssDataV3) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ApiNVD20CvssDataV3) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTemporalScore

`func (o *ApiNVD20CvssDataV3) GetTemporalScore() float32`

GetTemporalScore returns the TemporalScore field if non-nil, zero value otherwise.

### GetTemporalScoreOk

`func (o *ApiNVD20CvssDataV3) GetTemporalScoreOk() (*float32, bool)`

GetTemporalScoreOk returns a tuple with the TemporalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalScore

`func (o *ApiNVD20CvssDataV3) SetTemporalScore(v float32)`

SetTemporalScore sets TemporalScore field to given value.

### HasTemporalScore

`func (o *ApiNVD20CvssDataV3) HasTemporalScore() bool`

HasTemporalScore returns a boolean if a field has been set.

### GetTemporalSeverity

`func (o *ApiNVD20CvssDataV3) GetTemporalSeverity() string`

GetTemporalSeverity returns the TemporalSeverity field if non-nil, zero value otherwise.

### GetTemporalSeverityOk

`func (o *ApiNVD20CvssDataV3) GetTemporalSeverityOk() (*string, bool)`

GetTemporalSeverityOk returns a tuple with the TemporalSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalSeverity

`func (o *ApiNVD20CvssDataV3) SetTemporalSeverity(v string)`

SetTemporalSeverity sets TemporalSeverity field to given value.

### HasTemporalSeverity

`func (o *ApiNVD20CvssDataV3) HasTemporalSeverity() bool`

HasTemporalSeverity returns a boolean if a field has been set.

### GetUserInteraction

`func (o *ApiNVD20CvssDataV3) GetUserInteraction() string`

GetUserInteraction returns the UserInteraction field if non-nil, zero value otherwise.

### GetUserInteractionOk

`func (o *ApiNVD20CvssDataV3) GetUserInteractionOk() (*string, bool)`

GetUserInteractionOk returns a tuple with the UserInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInteraction

`func (o *ApiNVD20CvssDataV3) SetUserInteraction(v string)`

SetUserInteraction sets UserInteraction field to given value.

### HasUserInteraction

`func (o *ApiNVD20CvssDataV3) HasUserInteraction() bool`

HasUserInteraction returns a boolean if a field has been set.

### GetVectorString

`func (o *ApiNVD20CvssDataV3) GetVectorString() string`

GetVectorString returns the VectorString field if non-nil, zero value otherwise.

### GetVectorStringOk

`func (o *ApiNVD20CvssDataV3) GetVectorStringOk() (*string, bool)`

GetVectorStringOk returns a tuple with the VectorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorString

`func (o *ApiNVD20CvssDataV3) SetVectorString(v string)`

SetVectorString sets VectorString field to given value.

### HasVectorString

`func (o *ApiNVD20CvssDataV3) HasVectorString() bool`

HasVectorString returns a boolean if a field has been set.

### GetVersion

`func (o *ApiNVD20CvssDataV3) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiNVD20CvssDataV3) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiNVD20CvssDataV3) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApiNVD20CvssDataV3) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


