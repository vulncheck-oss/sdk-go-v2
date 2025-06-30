# AdvisoryCVSSV40

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Automatable** | Pointer to **string** |  | [optional] 
**Recovery** | Pointer to **string** |  | [optional] 
**Safety** | Pointer to **string** |  | [optional] 
**AttackComplexity** | Pointer to **string** |  | [optional] 
**AttackRequirements** | Pointer to **string** |  | [optional] 
**AttackVector** | Pointer to **string** |  | [optional] 
**AvailabilityRequirement** | Pointer to **string** |  | [optional] 
**BaseScore** | Pointer to **float32** |  | [optional] 
**BaseSeverity** | Pointer to **string** |  | [optional] 
**ConfidentialityRequirement** | Pointer to **string** |  | [optional] 
**ExploitMaturity** | Pointer to **string** |  | [optional] 
**IntegrityRequirement** | Pointer to **string** |  | [optional] 
**ModifiedAttackComplexity** | Pointer to **string** |  | [optional] 
**ModifiedAttackRequirements** | Pointer to **string** |  | [optional] 
**ModifiedAttackVector** | Pointer to **string** |  | [optional] 
**ModifiedPrivilegesRequired** | Pointer to **string** |  | [optional] 
**ModifiedSubAvailabilityImpact** | Pointer to **string** |  | [optional] 
**ModifiedSubConfidentialityImpact** | Pointer to **string** |  | [optional] 
**ModifiedSubIntegrityImpact** | Pointer to **string** |  | [optional] 
**ModifiedUserInteraction** | Pointer to **string** |  | [optional] 
**ModifiedVulnAvailabilityImpact** | Pointer to **string** |  | [optional] 
**ModifiedVulnConfidentialityImpact** | Pointer to **string** |  | [optional] 
**ModifiedVulnIntegrityImpact** | Pointer to **string** |  | [optional] 
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

### NewAdvisoryCVSSV40

`func NewAdvisoryCVSSV40() *AdvisoryCVSSV40`

NewAdvisoryCVSSV40 instantiates a new AdvisoryCVSSV40 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCVSSV40WithDefaults

`func NewAdvisoryCVSSV40WithDefaults() *AdvisoryCVSSV40`

NewAdvisoryCVSSV40WithDefaults instantiates a new AdvisoryCVSSV40 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomatable

`func (o *AdvisoryCVSSV40) GetAutomatable() string`

GetAutomatable returns the Automatable field if non-nil, zero value otherwise.

### GetAutomatableOk

`func (o *AdvisoryCVSSV40) GetAutomatableOk() (*string, bool)`

GetAutomatableOk returns a tuple with the Automatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatable

`func (o *AdvisoryCVSSV40) SetAutomatable(v string)`

SetAutomatable sets Automatable field to given value.

### HasAutomatable

`func (o *AdvisoryCVSSV40) HasAutomatable() bool`

HasAutomatable returns a boolean if a field has been set.

### GetRecovery

`func (o *AdvisoryCVSSV40) GetRecovery() string`

GetRecovery returns the Recovery field if non-nil, zero value otherwise.

### GetRecoveryOk

`func (o *AdvisoryCVSSV40) GetRecoveryOk() (*string, bool)`

GetRecoveryOk returns a tuple with the Recovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecovery

`func (o *AdvisoryCVSSV40) SetRecovery(v string)`

SetRecovery sets Recovery field to given value.

### HasRecovery

`func (o *AdvisoryCVSSV40) HasRecovery() bool`

HasRecovery returns a boolean if a field has been set.

### GetSafety

`func (o *AdvisoryCVSSV40) GetSafety() string`

GetSafety returns the Safety field if non-nil, zero value otherwise.

### GetSafetyOk

`func (o *AdvisoryCVSSV40) GetSafetyOk() (*string, bool)`

GetSafetyOk returns a tuple with the Safety field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafety

`func (o *AdvisoryCVSSV40) SetSafety(v string)`

SetSafety sets Safety field to given value.

### HasSafety

`func (o *AdvisoryCVSSV40) HasSafety() bool`

HasSafety returns a boolean if a field has been set.

### GetAttackComplexity

`func (o *AdvisoryCVSSV40) GetAttackComplexity() string`

GetAttackComplexity returns the AttackComplexity field if non-nil, zero value otherwise.

### GetAttackComplexityOk

`func (o *AdvisoryCVSSV40) GetAttackComplexityOk() (*string, bool)`

GetAttackComplexityOk returns a tuple with the AttackComplexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackComplexity

`func (o *AdvisoryCVSSV40) SetAttackComplexity(v string)`

SetAttackComplexity sets AttackComplexity field to given value.

### HasAttackComplexity

`func (o *AdvisoryCVSSV40) HasAttackComplexity() bool`

HasAttackComplexity returns a boolean if a field has been set.

### GetAttackRequirements

`func (o *AdvisoryCVSSV40) GetAttackRequirements() string`

GetAttackRequirements returns the AttackRequirements field if non-nil, zero value otherwise.

### GetAttackRequirementsOk

`func (o *AdvisoryCVSSV40) GetAttackRequirementsOk() (*string, bool)`

GetAttackRequirementsOk returns a tuple with the AttackRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackRequirements

`func (o *AdvisoryCVSSV40) SetAttackRequirements(v string)`

SetAttackRequirements sets AttackRequirements field to given value.

### HasAttackRequirements

`func (o *AdvisoryCVSSV40) HasAttackRequirements() bool`

HasAttackRequirements returns a boolean if a field has been set.

### GetAttackVector

`func (o *AdvisoryCVSSV40) GetAttackVector() string`

GetAttackVector returns the AttackVector field if non-nil, zero value otherwise.

### GetAttackVectorOk

`func (o *AdvisoryCVSSV40) GetAttackVectorOk() (*string, bool)`

GetAttackVectorOk returns a tuple with the AttackVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackVector

`func (o *AdvisoryCVSSV40) SetAttackVector(v string)`

SetAttackVector sets AttackVector field to given value.

### HasAttackVector

`func (o *AdvisoryCVSSV40) HasAttackVector() bool`

HasAttackVector returns a boolean if a field has been set.

### GetAvailabilityRequirement

`func (o *AdvisoryCVSSV40) GetAvailabilityRequirement() string`

GetAvailabilityRequirement returns the AvailabilityRequirement field if non-nil, zero value otherwise.

### GetAvailabilityRequirementOk

`func (o *AdvisoryCVSSV40) GetAvailabilityRequirementOk() (*string, bool)`

GetAvailabilityRequirementOk returns a tuple with the AvailabilityRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityRequirement

`func (o *AdvisoryCVSSV40) SetAvailabilityRequirement(v string)`

SetAvailabilityRequirement sets AvailabilityRequirement field to given value.

### HasAvailabilityRequirement

`func (o *AdvisoryCVSSV40) HasAvailabilityRequirement() bool`

HasAvailabilityRequirement returns a boolean if a field has been set.

### GetBaseScore

`func (o *AdvisoryCVSSV40) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *AdvisoryCVSSV40) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *AdvisoryCVSSV40) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *AdvisoryCVSSV40) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetBaseSeverity

`func (o *AdvisoryCVSSV40) GetBaseSeverity() string`

GetBaseSeverity returns the BaseSeverity field if non-nil, zero value otherwise.

### GetBaseSeverityOk

`func (o *AdvisoryCVSSV40) GetBaseSeverityOk() (*string, bool)`

GetBaseSeverityOk returns a tuple with the BaseSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSeverity

`func (o *AdvisoryCVSSV40) SetBaseSeverity(v string)`

SetBaseSeverity sets BaseSeverity field to given value.

### HasBaseSeverity

`func (o *AdvisoryCVSSV40) HasBaseSeverity() bool`

HasBaseSeverity returns a boolean if a field has been set.

### GetConfidentialityRequirement

`func (o *AdvisoryCVSSV40) GetConfidentialityRequirement() string`

GetConfidentialityRequirement returns the ConfidentialityRequirement field if non-nil, zero value otherwise.

### GetConfidentialityRequirementOk

`func (o *AdvisoryCVSSV40) GetConfidentialityRequirementOk() (*string, bool)`

GetConfidentialityRequirementOk returns a tuple with the ConfidentialityRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialityRequirement

`func (o *AdvisoryCVSSV40) SetConfidentialityRequirement(v string)`

SetConfidentialityRequirement sets ConfidentialityRequirement field to given value.

### HasConfidentialityRequirement

`func (o *AdvisoryCVSSV40) HasConfidentialityRequirement() bool`

HasConfidentialityRequirement returns a boolean if a field has been set.

### GetExploitMaturity

`func (o *AdvisoryCVSSV40) GetExploitMaturity() string`

GetExploitMaturity returns the ExploitMaturity field if non-nil, zero value otherwise.

### GetExploitMaturityOk

`func (o *AdvisoryCVSSV40) GetExploitMaturityOk() (*string, bool)`

GetExploitMaturityOk returns a tuple with the ExploitMaturity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitMaturity

`func (o *AdvisoryCVSSV40) SetExploitMaturity(v string)`

SetExploitMaturity sets ExploitMaturity field to given value.

### HasExploitMaturity

`func (o *AdvisoryCVSSV40) HasExploitMaturity() bool`

HasExploitMaturity returns a boolean if a field has been set.

### GetIntegrityRequirement

`func (o *AdvisoryCVSSV40) GetIntegrityRequirement() string`

GetIntegrityRequirement returns the IntegrityRequirement field if non-nil, zero value otherwise.

### GetIntegrityRequirementOk

`func (o *AdvisoryCVSSV40) GetIntegrityRequirementOk() (*string, bool)`

GetIntegrityRequirementOk returns a tuple with the IntegrityRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityRequirement

`func (o *AdvisoryCVSSV40) SetIntegrityRequirement(v string)`

SetIntegrityRequirement sets IntegrityRequirement field to given value.

### HasIntegrityRequirement

`func (o *AdvisoryCVSSV40) HasIntegrityRequirement() bool`

HasIntegrityRequirement returns a boolean if a field has been set.

### GetModifiedAttackComplexity

`func (o *AdvisoryCVSSV40) GetModifiedAttackComplexity() string`

GetModifiedAttackComplexity returns the ModifiedAttackComplexity field if non-nil, zero value otherwise.

### GetModifiedAttackComplexityOk

`func (o *AdvisoryCVSSV40) GetModifiedAttackComplexityOk() (*string, bool)`

GetModifiedAttackComplexityOk returns a tuple with the ModifiedAttackComplexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAttackComplexity

`func (o *AdvisoryCVSSV40) SetModifiedAttackComplexity(v string)`

SetModifiedAttackComplexity sets ModifiedAttackComplexity field to given value.

### HasModifiedAttackComplexity

`func (o *AdvisoryCVSSV40) HasModifiedAttackComplexity() bool`

HasModifiedAttackComplexity returns a boolean if a field has been set.

### GetModifiedAttackRequirements

`func (o *AdvisoryCVSSV40) GetModifiedAttackRequirements() string`

GetModifiedAttackRequirements returns the ModifiedAttackRequirements field if non-nil, zero value otherwise.

### GetModifiedAttackRequirementsOk

`func (o *AdvisoryCVSSV40) GetModifiedAttackRequirementsOk() (*string, bool)`

GetModifiedAttackRequirementsOk returns a tuple with the ModifiedAttackRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAttackRequirements

`func (o *AdvisoryCVSSV40) SetModifiedAttackRequirements(v string)`

SetModifiedAttackRequirements sets ModifiedAttackRequirements field to given value.

### HasModifiedAttackRequirements

`func (o *AdvisoryCVSSV40) HasModifiedAttackRequirements() bool`

HasModifiedAttackRequirements returns a boolean if a field has been set.

### GetModifiedAttackVector

`func (o *AdvisoryCVSSV40) GetModifiedAttackVector() string`

GetModifiedAttackVector returns the ModifiedAttackVector field if non-nil, zero value otherwise.

### GetModifiedAttackVectorOk

`func (o *AdvisoryCVSSV40) GetModifiedAttackVectorOk() (*string, bool)`

GetModifiedAttackVectorOk returns a tuple with the ModifiedAttackVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAttackVector

`func (o *AdvisoryCVSSV40) SetModifiedAttackVector(v string)`

SetModifiedAttackVector sets ModifiedAttackVector field to given value.

### HasModifiedAttackVector

`func (o *AdvisoryCVSSV40) HasModifiedAttackVector() bool`

HasModifiedAttackVector returns a boolean if a field has been set.

### GetModifiedPrivilegesRequired

`func (o *AdvisoryCVSSV40) GetModifiedPrivilegesRequired() string`

GetModifiedPrivilegesRequired returns the ModifiedPrivilegesRequired field if non-nil, zero value otherwise.

### GetModifiedPrivilegesRequiredOk

`func (o *AdvisoryCVSSV40) GetModifiedPrivilegesRequiredOk() (*string, bool)`

GetModifiedPrivilegesRequiredOk returns a tuple with the ModifiedPrivilegesRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedPrivilegesRequired

`func (o *AdvisoryCVSSV40) SetModifiedPrivilegesRequired(v string)`

SetModifiedPrivilegesRequired sets ModifiedPrivilegesRequired field to given value.

### HasModifiedPrivilegesRequired

`func (o *AdvisoryCVSSV40) HasModifiedPrivilegesRequired() bool`

HasModifiedPrivilegesRequired returns a boolean if a field has been set.

### GetModifiedSubAvailabilityImpact

`func (o *AdvisoryCVSSV40) GetModifiedSubAvailabilityImpact() string`

GetModifiedSubAvailabilityImpact returns the ModifiedSubAvailabilityImpact field if non-nil, zero value otherwise.

### GetModifiedSubAvailabilityImpactOk

`func (o *AdvisoryCVSSV40) GetModifiedSubAvailabilityImpactOk() (*string, bool)`

GetModifiedSubAvailabilityImpactOk returns a tuple with the ModifiedSubAvailabilityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedSubAvailabilityImpact

`func (o *AdvisoryCVSSV40) SetModifiedSubAvailabilityImpact(v string)`

SetModifiedSubAvailabilityImpact sets ModifiedSubAvailabilityImpact field to given value.

### HasModifiedSubAvailabilityImpact

`func (o *AdvisoryCVSSV40) HasModifiedSubAvailabilityImpact() bool`

HasModifiedSubAvailabilityImpact returns a boolean if a field has been set.

### GetModifiedSubConfidentialityImpact

`func (o *AdvisoryCVSSV40) GetModifiedSubConfidentialityImpact() string`

GetModifiedSubConfidentialityImpact returns the ModifiedSubConfidentialityImpact field if non-nil, zero value otherwise.

### GetModifiedSubConfidentialityImpactOk

`func (o *AdvisoryCVSSV40) GetModifiedSubConfidentialityImpactOk() (*string, bool)`

GetModifiedSubConfidentialityImpactOk returns a tuple with the ModifiedSubConfidentialityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedSubConfidentialityImpact

`func (o *AdvisoryCVSSV40) SetModifiedSubConfidentialityImpact(v string)`

SetModifiedSubConfidentialityImpact sets ModifiedSubConfidentialityImpact field to given value.

### HasModifiedSubConfidentialityImpact

`func (o *AdvisoryCVSSV40) HasModifiedSubConfidentialityImpact() bool`

HasModifiedSubConfidentialityImpact returns a boolean if a field has been set.

### GetModifiedSubIntegrityImpact

`func (o *AdvisoryCVSSV40) GetModifiedSubIntegrityImpact() string`

GetModifiedSubIntegrityImpact returns the ModifiedSubIntegrityImpact field if non-nil, zero value otherwise.

### GetModifiedSubIntegrityImpactOk

`func (o *AdvisoryCVSSV40) GetModifiedSubIntegrityImpactOk() (*string, bool)`

GetModifiedSubIntegrityImpactOk returns a tuple with the ModifiedSubIntegrityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedSubIntegrityImpact

`func (o *AdvisoryCVSSV40) SetModifiedSubIntegrityImpact(v string)`

SetModifiedSubIntegrityImpact sets ModifiedSubIntegrityImpact field to given value.

### HasModifiedSubIntegrityImpact

`func (o *AdvisoryCVSSV40) HasModifiedSubIntegrityImpact() bool`

HasModifiedSubIntegrityImpact returns a boolean if a field has been set.

### GetModifiedUserInteraction

`func (o *AdvisoryCVSSV40) GetModifiedUserInteraction() string`

GetModifiedUserInteraction returns the ModifiedUserInteraction field if non-nil, zero value otherwise.

### GetModifiedUserInteractionOk

`func (o *AdvisoryCVSSV40) GetModifiedUserInteractionOk() (*string, bool)`

GetModifiedUserInteractionOk returns a tuple with the ModifiedUserInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedUserInteraction

`func (o *AdvisoryCVSSV40) SetModifiedUserInteraction(v string)`

SetModifiedUserInteraction sets ModifiedUserInteraction field to given value.

### HasModifiedUserInteraction

`func (o *AdvisoryCVSSV40) HasModifiedUserInteraction() bool`

HasModifiedUserInteraction returns a boolean if a field has been set.

### GetModifiedVulnAvailabilityImpact

`func (o *AdvisoryCVSSV40) GetModifiedVulnAvailabilityImpact() string`

GetModifiedVulnAvailabilityImpact returns the ModifiedVulnAvailabilityImpact field if non-nil, zero value otherwise.

### GetModifiedVulnAvailabilityImpactOk

`func (o *AdvisoryCVSSV40) GetModifiedVulnAvailabilityImpactOk() (*string, bool)`

GetModifiedVulnAvailabilityImpactOk returns a tuple with the ModifiedVulnAvailabilityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedVulnAvailabilityImpact

`func (o *AdvisoryCVSSV40) SetModifiedVulnAvailabilityImpact(v string)`

SetModifiedVulnAvailabilityImpact sets ModifiedVulnAvailabilityImpact field to given value.

### HasModifiedVulnAvailabilityImpact

`func (o *AdvisoryCVSSV40) HasModifiedVulnAvailabilityImpact() bool`

HasModifiedVulnAvailabilityImpact returns a boolean if a field has been set.

### GetModifiedVulnConfidentialityImpact

`func (o *AdvisoryCVSSV40) GetModifiedVulnConfidentialityImpact() string`

GetModifiedVulnConfidentialityImpact returns the ModifiedVulnConfidentialityImpact field if non-nil, zero value otherwise.

### GetModifiedVulnConfidentialityImpactOk

`func (o *AdvisoryCVSSV40) GetModifiedVulnConfidentialityImpactOk() (*string, bool)`

GetModifiedVulnConfidentialityImpactOk returns a tuple with the ModifiedVulnConfidentialityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedVulnConfidentialityImpact

`func (o *AdvisoryCVSSV40) SetModifiedVulnConfidentialityImpact(v string)`

SetModifiedVulnConfidentialityImpact sets ModifiedVulnConfidentialityImpact field to given value.

### HasModifiedVulnConfidentialityImpact

`func (o *AdvisoryCVSSV40) HasModifiedVulnConfidentialityImpact() bool`

HasModifiedVulnConfidentialityImpact returns a boolean if a field has been set.

### GetModifiedVulnIntegrityImpact

`func (o *AdvisoryCVSSV40) GetModifiedVulnIntegrityImpact() string`

GetModifiedVulnIntegrityImpact returns the ModifiedVulnIntegrityImpact field if non-nil, zero value otherwise.

### GetModifiedVulnIntegrityImpactOk

`func (o *AdvisoryCVSSV40) GetModifiedVulnIntegrityImpactOk() (*string, bool)`

GetModifiedVulnIntegrityImpactOk returns a tuple with the ModifiedVulnIntegrityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedVulnIntegrityImpact

`func (o *AdvisoryCVSSV40) SetModifiedVulnIntegrityImpact(v string)`

SetModifiedVulnIntegrityImpact sets ModifiedVulnIntegrityImpact field to given value.

### HasModifiedVulnIntegrityImpact

`func (o *AdvisoryCVSSV40) HasModifiedVulnIntegrityImpact() bool`

HasModifiedVulnIntegrityImpact returns a boolean if a field has been set.

### GetPrivilegesRequired

`func (o *AdvisoryCVSSV40) GetPrivilegesRequired() string`

GetPrivilegesRequired returns the PrivilegesRequired field if non-nil, zero value otherwise.

### GetPrivilegesRequiredOk

`func (o *AdvisoryCVSSV40) GetPrivilegesRequiredOk() (*string, bool)`

GetPrivilegesRequiredOk returns a tuple with the PrivilegesRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegesRequired

`func (o *AdvisoryCVSSV40) SetPrivilegesRequired(v string)`

SetPrivilegesRequired sets PrivilegesRequired field to given value.

### HasPrivilegesRequired

`func (o *AdvisoryCVSSV40) HasPrivilegesRequired() bool`

HasPrivilegesRequired returns a boolean if a field has been set.

### GetProviderUrgency

`func (o *AdvisoryCVSSV40) GetProviderUrgency() string`

GetProviderUrgency returns the ProviderUrgency field if non-nil, zero value otherwise.

### GetProviderUrgencyOk

`func (o *AdvisoryCVSSV40) GetProviderUrgencyOk() (*string, bool)`

GetProviderUrgencyOk returns a tuple with the ProviderUrgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrgency

`func (o *AdvisoryCVSSV40) SetProviderUrgency(v string)`

SetProviderUrgency sets ProviderUrgency field to given value.

### HasProviderUrgency

`func (o *AdvisoryCVSSV40) HasProviderUrgency() bool`

HasProviderUrgency returns a boolean if a field has been set.

### GetSubAvailabilityImpact

`func (o *AdvisoryCVSSV40) GetSubAvailabilityImpact() string`

GetSubAvailabilityImpact returns the SubAvailabilityImpact field if non-nil, zero value otherwise.

### GetSubAvailabilityImpactOk

`func (o *AdvisoryCVSSV40) GetSubAvailabilityImpactOk() (*string, bool)`

GetSubAvailabilityImpactOk returns a tuple with the SubAvailabilityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAvailabilityImpact

`func (o *AdvisoryCVSSV40) SetSubAvailabilityImpact(v string)`

SetSubAvailabilityImpact sets SubAvailabilityImpact field to given value.

### HasSubAvailabilityImpact

`func (o *AdvisoryCVSSV40) HasSubAvailabilityImpact() bool`

HasSubAvailabilityImpact returns a boolean if a field has been set.

### GetSubConfidentialityImpact

`func (o *AdvisoryCVSSV40) GetSubConfidentialityImpact() string`

GetSubConfidentialityImpact returns the SubConfidentialityImpact field if non-nil, zero value otherwise.

### GetSubConfidentialityImpactOk

`func (o *AdvisoryCVSSV40) GetSubConfidentialityImpactOk() (*string, bool)`

GetSubConfidentialityImpactOk returns a tuple with the SubConfidentialityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubConfidentialityImpact

`func (o *AdvisoryCVSSV40) SetSubConfidentialityImpact(v string)`

SetSubConfidentialityImpact sets SubConfidentialityImpact field to given value.

### HasSubConfidentialityImpact

`func (o *AdvisoryCVSSV40) HasSubConfidentialityImpact() bool`

HasSubConfidentialityImpact returns a boolean if a field has been set.

### GetSubIntegrityImpact

`func (o *AdvisoryCVSSV40) GetSubIntegrityImpact() string`

GetSubIntegrityImpact returns the SubIntegrityImpact field if non-nil, zero value otherwise.

### GetSubIntegrityImpactOk

`func (o *AdvisoryCVSSV40) GetSubIntegrityImpactOk() (*string, bool)`

GetSubIntegrityImpactOk returns a tuple with the SubIntegrityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubIntegrityImpact

`func (o *AdvisoryCVSSV40) SetSubIntegrityImpact(v string)`

SetSubIntegrityImpact sets SubIntegrityImpact field to given value.

### HasSubIntegrityImpact

`func (o *AdvisoryCVSSV40) HasSubIntegrityImpact() bool`

HasSubIntegrityImpact returns a boolean if a field has been set.

### GetUserInteraction

`func (o *AdvisoryCVSSV40) GetUserInteraction() string`

GetUserInteraction returns the UserInteraction field if non-nil, zero value otherwise.

### GetUserInteractionOk

`func (o *AdvisoryCVSSV40) GetUserInteractionOk() (*string, bool)`

GetUserInteractionOk returns a tuple with the UserInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInteraction

`func (o *AdvisoryCVSSV40) SetUserInteraction(v string)`

SetUserInteraction sets UserInteraction field to given value.

### HasUserInteraction

`func (o *AdvisoryCVSSV40) HasUserInteraction() bool`

HasUserInteraction returns a boolean if a field has been set.

### GetValueDensity

`func (o *AdvisoryCVSSV40) GetValueDensity() string`

GetValueDensity returns the ValueDensity field if non-nil, zero value otherwise.

### GetValueDensityOk

`func (o *AdvisoryCVSSV40) GetValueDensityOk() (*string, bool)`

GetValueDensityOk returns a tuple with the ValueDensity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDensity

`func (o *AdvisoryCVSSV40) SetValueDensity(v string)`

SetValueDensity sets ValueDensity field to given value.

### HasValueDensity

`func (o *AdvisoryCVSSV40) HasValueDensity() bool`

HasValueDensity returns a boolean if a field has been set.

### GetVectorString

`func (o *AdvisoryCVSSV40) GetVectorString() string`

GetVectorString returns the VectorString field if non-nil, zero value otherwise.

### GetVectorStringOk

`func (o *AdvisoryCVSSV40) GetVectorStringOk() (*string, bool)`

GetVectorStringOk returns a tuple with the VectorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorString

`func (o *AdvisoryCVSSV40) SetVectorString(v string)`

SetVectorString sets VectorString field to given value.

### HasVectorString

`func (o *AdvisoryCVSSV40) HasVectorString() bool`

HasVectorString returns a boolean if a field has been set.

### GetVersion

`func (o *AdvisoryCVSSV40) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AdvisoryCVSSV40) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AdvisoryCVSSV40) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AdvisoryCVSSV40) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVulnAvailabilityImpact

`func (o *AdvisoryCVSSV40) GetVulnAvailabilityImpact() string`

GetVulnAvailabilityImpact returns the VulnAvailabilityImpact field if non-nil, zero value otherwise.

### GetVulnAvailabilityImpactOk

`func (o *AdvisoryCVSSV40) GetVulnAvailabilityImpactOk() (*string, bool)`

GetVulnAvailabilityImpactOk returns a tuple with the VulnAvailabilityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnAvailabilityImpact

`func (o *AdvisoryCVSSV40) SetVulnAvailabilityImpact(v string)`

SetVulnAvailabilityImpact sets VulnAvailabilityImpact field to given value.

### HasVulnAvailabilityImpact

`func (o *AdvisoryCVSSV40) HasVulnAvailabilityImpact() bool`

HasVulnAvailabilityImpact returns a boolean if a field has been set.

### GetVulnConfidentialityImpact

`func (o *AdvisoryCVSSV40) GetVulnConfidentialityImpact() string`

GetVulnConfidentialityImpact returns the VulnConfidentialityImpact field if non-nil, zero value otherwise.

### GetVulnConfidentialityImpactOk

`func (o *AdvisoryCVSSV40) GetVulnConfidentialityImpactOk() (*string, bool)`

GetVulnConfidentialityImpactOk returns a tuple with the VulnConfidentialityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnConfidentialityImpact

`func (o *AdvisoryCVSSV40) SetVulnConfidentialityImpact(v string)`

SetVulnConfidentialityImpact sets VulnConfidentialityImpact field to given value.

### HasVulnConfidentialityImpact

`func (o *AdvisoryCVSSV40) HasVulnConfidentialityImpact() bool`

HasVulnConfidentialityImpact returns a boolean if a field has been set.

### GetVulnIntegrityImpact

`func (o *AdvisoryCVSSV40) GetVulnIntegrityImpact() string`

GetVulnIntegrityImpact returns the VulnIntegrityImpact field if non-nil, zero value otherwise.

### GetVulnIntegrityImpactOk

`func (o *AdvisoryCVSSV40) GetVulnIntegrityImpactOk() (*string, bool)`

GetVulnIntegrityImpactOk returns a tuple with the VulnIntegrityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnIntegrityImpact

`func (o *AdvisoryCVSSV40) SetVulnIntegrityImpact(v string)`

SetVulnIntegrityImpact sets VulnIntegrityImpact field to given value.

### HasVulnIntegrityImpact

`func (o *AdvisoryCVSSV40) HasVulnIntegrityImpact() bool`

HasVulnIntegrityImpact returns a boolean if a field has been set.

### GetVulnerabilityResponseEffort

`func (o *AdvisoryCVSSV40) GetVulnerabilityResponseEffort() string`

GetVulnerabilityResponseEffort returns the VulnerabilityResponseEffort field if non-nil, zero value otherwise.

### GetVulnerabilityResponseEffortOk

`func (o *AdvisoryCVSSV40) GetVulnerabilityResponseEffortOk() (*string, bool)`

GetVulnerabilityResponseEffortOk returns a tuple with the VulnerabilityResponseEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityResponseEffort

`func (o *AdvisoryCVSSV40) SetVulnerabilityResponseEffort(v string)`

SetVulnerabilityResponseEffort sets VulnerabilityResponseEffort field to given value.

### HasVulnerabilityResponseEffort

`func (o *AdvisoryCVSSV40) HasVulnerabilityResponseEffort() bool`

HasVulnerabilityResponseEffort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


