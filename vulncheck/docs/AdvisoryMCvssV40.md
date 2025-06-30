# AdvisoryMCvssV40

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Automatable** | Pointer to **string** |  | [optional] 
**Recovery** | Pointer to **string** |  | [optional] 
**Safety** | Pointer to **string** |  | [optional] 
**AttackComplexity** | Pointer to **string** |  | [optional] 
**AttackRequirements** | Pointer to **string** |  | [optional] 
**AttackVector** | Pointer to **string** |  | [optional] 
**BaseScore** | Pointer to **float32** |  | [optional] 
**BaseSeverity** | Pointer to **string** |  | [optional] 
**PrivilegesRequired** | Pointer to **string** |  | [optional] 
**ProviderUrgency** | Pointer to **string** |  | [optional] 
**SubAvailabilityImpact** | Pointer to **string** |  | [optional] 
**SubConfidentialityImpact** | Pointer to **string** |  | [optional] 
**SubIntegrityImpact** | Pointer to **string** |  | [optional] 
**UserInteraction** | Pointer to **string** |  | [optional] 
**ValueDensity** | Pointer to **string** |  | [optional] 
**VectorString** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**VulnAvailabilityImpact** | Pointer to **string** |  | [optional] 
**VulnConfidentialityImpact** | Pointer to **string** |  | [optional] 
**VulnIntegrityImpact** | Pointer to **string** |  | [optional] 
**VulnerabilityResponseEffort** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryMCvssV40

`func NewAdvisoryMCvssV40() *AdvisoryMCvssV40`

NewAdvisoryMCvssV40 instantiates a new AdvisoryMCvssV40 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMCvssV40WithDefaults

`func NewAdvisoryMCvssV40WithDefaults() *AdvisoryMCvssV40`

NewAdvisoryMCvssV40WithDefaults instantiates a new AdvisoryMCvssV40 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomatable

`func (o *AdvisoryMCvssV40) GetAutomatable() string`

GetAutomatable returns the Automatable field if non-nil, zero value otherwise.

### GetAutomatableOk

`func (o *AdvisoryMCvssV40) GetAutomatableOk() (*string, bool)`

GetAutomatableOk returns a tuple with the Automatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatable

`func (o *AdvisoryMCvssV40) SetAutomatable(v string)`

SetAutomatable sets Automatable field to given value.

### HasAutomatable

`func (o *AdvisoryMCvssV40) HasAutomatable() bool`

HasAutomatable returns a boolean if a field has been set.

### GetRecovery

`func (o *AdvisoryMCvssV40) GetRecovery() string`

GetRecovery returns the Recovery field if non-nil, zero value otherwise.

### GetRecoveryOk

`func (o *AdvisoryMCvssV40) GetRecoveryOk() (*string, bool)`

GetRecoveryOk returns a tuple with the Recovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecovery

`func (o *AdvisoryMCvssV40) SetRecovery(v string)`

SetRecovery sets Recovery field to given value.

### HasRecovery

`func (o *AdvisoryMCvssV40) HasRecovery() bool`

HasRecovery returns a boolean if a field has been set.

### GetSafety

`func (o *AdvisoryMCvssV40) GetSafety() string`

GetSafety returns the Safety field if non-nil, zero value otherwise.

### GetSafetyOk

`func (o *AdvisoryMCvssV40) GetSafetyOk() (*string, bool)`

GetSafetyOk returns a tuple with the Safety field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafety

`func (o *AdvisoryMCvssV40) SetSafety(v string)`

SetSafety sets Safety field to given value.

### HasSafety

`func (o *AdvisoryMCvssV40) HasSafety() bool`

HasSafety returns a boolean if a field has been set.

### GetAttackComplexity

`func (o *AdvisoryMCvssV40) GetAttackComplexity() string`

GetAttackComplexity returns the AttackComplexity field if non-nil, zero value otherwise.

### GetAttackComplexityOk

`func (o *AdvisoryMCvssV40) GetAttackComplexityOk() (*string, bool)`

GetAttackComplexityOk returns a tuple with the AttackComplexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackComplexity

`func (o *AdvisoryMCvssV40) SetAttackComplexity(v string)`

SetAttackComplexity sets AttackComplexity field to given value.

### HasAttackComplexity

`func (o *AdvisoryMCvssV40) HasAttackComplexity() bool`

HasAttackComplexity returns a boolean if a field has been set.

### GetAttackRequirements

`func (o *AdvisoryMCvssV40) GetAttackRequirements() string`

GetAttackRequirements returns the AttackRequirements field if non-nil, zero value otherwise.

### GetAttackRequirementsOk

`func (o *AdvisoryMCvssV40) GetAttackRequirementsOk() (*string, bool)`

GetAttackRequirementsOk returns a tuple with the AttackRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackRequirements

`func (o *AdvisoryMCvssV40) SetAttackRequirements(v string)`

SetAttackRequirements sets AttackRequirements field to given value.

### HasAttackRequirements

`func (o *AdvisoryMCvssV40) HasAttackRequirements() bool`

HasAttackRequirements returns a boolean if a field has been set.

### GetAttackVector

`func (o *AdvisoryMCvssV40) GetAttackVector() string`

GetAttackVector returns the AttackVector field if non-nil, zero value otherwise.

### GetAttackVectorOk

`func (o *AdvisoryMCvssV40) GetAttackVectorOk() (*string, bool)`

GetAttackVectorOk returns a tuple with the AttackVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackVector

`func (o *AdvisoryMCvssV40) SetAttackVector(v string)`

SetAttackVector sets AttackVector field to given value.

### HasAttackVector

`func (o *AdvisoryMCvssV40) HasAttackVector() bool`

HasAttackVector returns a boolean if a field has been set.

### GetBaseScore

`func (o *AdvisoryMCvssV40) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *AdvisoryMCvssV40) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *AdvisoryMCvssV40) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *AdvisoryMCvssV40) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetBaseSeverity

`func (o *AdvisoryMCvssV40) GetBaseSeverity() string`

GetBaseSeverity returns the BaseSeverity field if non-nil, zero value otherwise.

### GetBaseSeverityOk

`func (o *AdvisoryMCvssV40) GetBaseSeverityOk() (*string, bool)`

GetBaseSeverityOk returns a tuple with the BaseSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSeverity

`func (o *AdvisoryMCvssV40) SetBaseSeverity(v string)`

SetBaseSeverity sets BaseSeverity field to given value.

### HasBaseSeverity

`func (o *AdvisoryMCvssV40) HasBaseSeverity() bool`

HasBaseSeverity returns a boolean if a field has been set.

### GetPrivilegesRequired

`func (o *AdvisoryMCvssV40) GetPrivilegesRequired() string`

GetPrivilegesRequired returns the PrivilegesRequired field if non-nil, zero value otherwise.

### GetPrivilegesRequiredOk

`func (o *AdvisoryMCvssV40) GetPrivilegesRequiredOk() (*string, bool)`

GetPrivilegesRequiredOk returns a tuple with the PrivilegesRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegesRequired

`func (o *AdvisoryMCvssV40) SetPrivilegesRequired(v string)`

SetPrivilegesRequired sets PrivilegesRequired field to given value.

### HasPrivilegesRequired

`func (o *AdvisoryMCvssV40) HasPrivilegesRequired() bool`

HasPrivilegesRequired returns a boolean if a field has been set.

### GetProviderUrgency

`func (o *AdvisoryMCvssV40) GetProviderUrgency() string`

GetProviderUrgency returns the ProviderUrgency field if non-nil, zero value otherwise.

### GetProviderUrgencyOk

`func (o *AdvisoryMCvssV40) GetProviderUrgencyOk() (*string, bool)`

GetProviderUrgencyOk returns a tuple with the ProviderUrgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrgency

`func (o *AdvisoryMCvssV40) SetProviderUrgency(v string)`

SetProviderUrgency sets ProviderUrgency field to given value.

### HasProviderUrgency

`func (o *AdvisoryMCvssV40) HasProviderUrgency() bool`

HasProviderUrgency returns a boolean if a field has been set.

### GetSubAvailabilityImpact

`func (o *AdvisoryMCvssV40) GetSubAvailabilityImpact() string`

GetSubAvailabilityImpact returns the SubAvailabilityImpact field if non-nil, zero value otherwise.

### GetSubAvailabilityImpactOk

`func (o *AdvisoryMCvssV40) GetSubAvailabilityImpactOk() (*string, bool)`

GetSubAvailabilityImpactOk returns a tuple with the SubAvailabilityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAvailabilityImpact

`func (o *AdvisoryMCvssV40) SetSubAvailabilityImpact(v string)`

SetSubAvailabilityImpact sets SubAvailabilityImpact field to given value.

### HasSubAvailabilityImpact

`func (o *AdvisoryMCvssV40) HasSubAvailabilityImpact() bool`

HasSubAvailabilityImpact returns a boolean if a field has been set.

### GetSubConfidentialityImpact

`func (o *AdvisoryMCvssV40) GetSubConfidentialityImpact() string`

GetSubConfidentialityImpact returns the SubConfidentialityImpact field if non-nil, zero value otherwise.

### GetSubConfidentialityImpactOk

`func (o *AdvisoryMCvssV40) GetSubConfidentialityImpactOk() (*string, bool)`

GetSubConfidentialityImpactOk returns a tuple with the SubConfidentialityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubConfidentialityImpact

`func (o *AdvisoryMCvssV40) SetSubConfidentialityImpact(v string)`

SetSubConfidentialityImpact sets SubConfidentialityImpact field to given value.

### HasSubConfidentialityImpact

`func (o *AdvisoryMCvssV40) HasSubConfidentialityImpact() bool`

HasSubConfidentialityImpact returns a boolean if a field has been set.

### GetSubIntegrityImpact

`func (o *AdvisoryMCvssV40) GetSubIntegrityImpact() string`

GetSubIntegrityImpact returns the SubIntegrityImpact field if non-nil, zero value otherwise.

### GetSubIntegrityImpactOk

`func (o *AdvisoryMCvssV40) GetSubIntegrityImpactOk() (*string, bool)`

GetSubIntegrityImpactOk returns a tuple with the SubIntegrityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubIntegrityImpact

`func (o *AdvisoryMCvssV40) SetSubIntegrityImpact(v string)`

SetSubIntegrityImpact sets SubIntegrityImpact field to given value.

### HasSubIntegrityImpact

`func (o *AdvisoryMCvssV40) HasSubIntegrityImpact() bool`

HasSubIntegrityImpact returns a boolean if a field has been set.

### GetUserInteraction

`func (o *AdvisoryMCvssV40) GetUserInteraction() string`

GetUserInteraction returns the UserInteraction field if non-nil, zero value otherwise.

### GetUserInteractionOk

`func (o *AdvisoryMCvssV40) GetUserInteractionOk() (*string, bool)`

GetUserInteractionOk returns a tuple with the UserInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInteraction

`func (o *AdvisoryMCvssV40) SetUserInteraction(v string)`

SetUserInteraction sets UserInteraction field to given value.

### HasUserInteraction

`func (o *AdvisoryMCvssV40) HasUserInteraction() bool`

HasUserInteraction returns a boolean if a field has been set.

### GetValueDensity

`func (o *AdvisoryMCvssV40) GetValueDensity() string`

GetValueDensity returns the ValueDensity field if non-nil, zero value otherwise.

### GetValueDensityOk

`func (o *AdvisoryMCvssV40) GetValueDensityOk() (*string, bool)`

GetValueDensityOk returns a tuple with the ValueDensity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDensity

`func (o *AdvisoryMCvssV40) SetValueDensity(v string)`

SetValueDensity sets ValueDensity field to given value.

### HasValueDensity

`func (o *AdvisoryMCvssV40) HasValueDensity() bool`

HasValueDensity returns a boolean if a field has been set.

### GetVectorString

`func (o *AdvisoryMCvssV40) GetVectorString() string`

GetVectorString returns the VectorString field if non-nil, zero value otherwise.

### GetVectorStringOk

`func (o *AdvisoryMCvssV40) GetVectorStringOk() (*string, bool)`

GetVectorStringOk returns a tuple with the VectorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorString

`func (o *AdvisoryMCvssV40) SetVectorString(v string)`

SetVectorString sets VectorString field to given value.

### HasVectorString

`func (o *AdvisoryMCvssV40) HasVectorString() bool`

HasVectorString returns a boolean if a field has been set.

### GetVersion

`func (o *AdvisoryMCvssV40) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AdvisoryMCvssV40) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AdvisoryMCvssV40) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AdvisoryMCvssV40) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVulnAvailabilityImpact

`func (o *AdvisoryMCvssV40) GetVulnAvailabilityImpact() string`

GetVulnAvailabilityImpact returns the VulnAvailabilityImpact field if non-nil, zero value otherwise.

### GetVulnAvailabilityImpactOk

`func (o *AdvisoryMCvssV40) GetVulnAvailabilityImpactOk() (*string, bool)`

GetVulnAvailabilityImpactOk returns a tuple with the VulnAvailabilityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnAvailabilityImpact

`func (o *AdvisoryMCvssV40) SetVulnAvailabilityImpact(v string)`

SetVulnAvailabilityImpact sets VulnAvailabilityImpact field to given value.

### HasVulnAvailabilityImpact

`func (o *AdvisoryMCvssV40) HasVulnAvailabilityImpact() bool`

HasVulnAvailabilityImpact returns a boolean if a field has been set.

### GetVulnConfidentialityImpact

`func (o *AdvisoryMCvssV40) GetVulnConfidentialityImpact() string`

GetVulnConfidentialityImpact returns the VulnConfidentialityImpact field if non-nil, zero value otherwise.

### GetVulnConfidentialityImpactOk

`func (o *AdvisoryMCvssV40) GetVulnConfidentialityImpactOk() (*string, bool)`

GetVulnConfidentialityImpactOk returns a tuple with the VulnConfidentialityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnConfidentialityImpact

`func (o *AdvisoryMCvssV40) SetVulnConfidentialityImpact(v string)`

SetVulnConfidentialityImpact sets VulnConfidentialityImpact field to given value.

### HasVulnConfidentialityImpact

`func (o *AdvisoryMCvssV40) HasVulnConfidentialityImpact() bool`

HasVulnConfidentialityImpact returns a boolean if a field has been set.

### GetVulnIntegrityImpact

`func (o *AdvisoryMCvssV40) GetVulnIntegrityImpact() string`

GetVulnIntegrityImpact returns the VulnIntegrityImpact field if non-nil, zero value otherwise.

### GetVulnIntegrityImpactOk

`func (o *AdvisoryMCvssV40) GetVulnIntegrityImpactOk() (*string, bool)`

GetVulnIntegrityImpactOk returns a tuple with the VulnIntegrityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnIntegrityImpact

`func (o *AdvisoryMCvssV40) SetVulnIntegrityImpact(v string)`

SetVulnIntegrityImpact sets VulnIntegrityImpact field to given value.

### HasVulnIntegrityImpact

`func (o *AdvisoryMCvssV40) HasVulnIntegrityImpact() bool`

HasVulnIntegrityImpact returns a boolean if a field has been set.

### GetVulnerabilityResponseEffort

`func (o *AdvisoryMCvssV40) GetVulnerabilityResponseEffort() string`

GetVulnerabilityResponseEffort returns the VulnerabilityResponseEffort field if non-nil, zero value otherwise.

### GetVulnerabilityResponseEffortOk

`func (o *AdvisoryMCvssV40) GetVulnerabilityResponseEffortOk() (*string, bool)`

GetVulnerabilityResponseEffortOk returns a tuple with the VulnerabilityResponseEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityResponseEffort

`func (o *AdvisoryMCvssV40) SetVulnerabilityResponseEffort(v string)`

SetVulnerabilityResponseEffort sets VulnerabilityResponseEffort field to given value.

### HasVulnerabilityResponseEffort

`func (o *AdvisoryMCvssV40) HasVulnerabilityResponseEffort() bool`

HasVulnerabilityResponseEffort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


