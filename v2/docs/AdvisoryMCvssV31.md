# AdvisoryMCvssV31

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttackComplexity** | Pointer to **string** |  | [optional] 
**AttackVector** | Pointer to **string** |  | [optional] 
**AvailabilityImpact** | Pointer to **string** |  | [optional] 
**BaseScore** | Pointer to **float32** |  | [optional] 
**BaseSeverity** | Pointer to **string** |  | [optional] 
**ConfidentialityImpact** | Pointer to **string** |  | [optional] 
**IntegrityImpact** | Pointer to **string** |  | [optional] 
**PrivilegesRequired** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**UserInteraction** | Pointer to **string** |  | [optional] 
**VectorString** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryMCvssV31

`func NewAdvisoryMCvssV31() *AdvisoryMCvssV31`

NewAdvisoryMCvssV31 instantiates a new AdvisoryMCvssV31 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMCvssV31WithDefaults

`func NewAdvisoryMCvssV31WithDefaults() *AdvisoryMCvssV31`

NewAdvisoryMCvssV31WithDefaults instantiates a new AdvisoryMCvssV31 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttackComplexity

`func (o *AdvisoryMCvssV31) GetAttackComplexity() string`

GetAttackComplexity returns the AttackComplexity field if non-nil, zero value otherwise.

### GetAttackComplexityOk

`func (o *AdvisoryMCvssV31) GetAttackComplexityOk() (*string, bool)`

GetAttackComplexityOk returns a tuple with the AttackComplexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackComplexity

`func (o *AdvisoryMCvssV31) SetAttackComplexity(v string)`

SetAttackComplexity sets AttackComplexity field to given value.

### HasAttackComplexity

`func (o *AdvisoryMCvssV31) HasAttackComplexity() bool`

HasAttackComplexity returns a boolean if a field has been set.

### GetAttackVector

`func (o *AdvisoryMCvssV31) GetAttackVector() string`

GetAttackVector returns the AttackVector field if non-nil, zero value otherwise.

### GetAttackVectorOk

`func (o *AdvisoryMCvssV31) GetAttackVectorOk() (*string, bool)`

GetAttackVectorOk returns a tuple with the AttackVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackVector

`func (o *AdvisoryMCvssV31) SetAttackVector(v string)`

SetAttackVector sets AttackVector field to given value.

### HasAttackVector

`func (o *AdvisoryMCvssV31) HasAttackVector() bool`

HasAttackVector returns a boolean if a field has been set.

### GetAvailabilityImpact

`func (o *AdvisoryMCvssV31) GetAvailabilityImpact() string`

GetAvailabilityImpact returns the AvailabilityImpact field if non-nil, zero value otherwise.

### GetAvailabilityImpactOk

`func (o *AdvisoryMCvssV31) GetAvailabilityImpactOk() (*string, bool)`

GetAvailabilityImpactOk returns a tuple with the AvailabilityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityImpact

`func (o *AdvisoryMCvssV31) SetAvailabilityImpact(v string)`

SetAvailabilityImpact sets AvailabilityImpact field to given value.

### HasAvailabilityImpact

`func (o *AdvisoryMCvssV31) HasAvailabilityImpact() bool`

HasAvailabilityImpact returns a boolean if a field has been set.

### GetBaseScore

`func (o *AdvisoryMCvssV31) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *AdvisoryMCvssV31) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *AdvisoryMCvssV31) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *AdvisoryMCvssV31) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetBaseSeverity

`func (o *AdvisoryMCvssV31) GetBaseSeverity() string`

GetBaseSeverity returns the BaseSeverity field if non-nil, zero value otherwise.

### GetBaseSeverityOk

`func (o *AdvisoryMCvssV31) GetBaseSeverityOk() (*string, bool)`

GetBaseSeverityOk returns a tuple with the BaseSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSeverity

`func (o *AdvisoryMCvssV31) SetBaseSeverity(v string)`

SetBaseSeverity sets BaseSeverity field to given value.

### HasBaseSeverity

`func (o *AdvisoryMCvssV31) HasBaseSeverity() bool`

HasBaseSeverity returns a boolean if a field has been set.

### GetConfidentialityImpact

`func (o *AdvisoryMCvssV31) GetConfidentialityImpact() string`

GetConfidentialityImpact returns the ConfidentialityImpact field if non-nil, zero value otherwise.

### GetConfidentialityImpactOk

`func (o *AdvisoryMCvssV31) GetConfidentialityImpactOk() (*string, bool)`

GetConfidentialityImpactOk returns a tuple with the ConfidentialityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialityImpact

`func (o *AdvisoryMCvssV31) SetConfidentialityImpact(v string)`

SetConfidentialityImpact sets ConfidentialityImpact field to given value.

### HasConfidentialityImpact

`func (o *AdvisoryMCvssV31) HasConfidentialityImpact() bool`

HasConfidentialityImpact returns a boolean if a field has been set.

### GetIntegrityImpact

`func (o *AdvisoryMCvssV31) GetIntegrityImpact() string`

GetIntegrityImpact returns the IntegrityImpact field if non-nil, zero value otherwise.

### GetIntegrityImpactOk

`func (o *AdvisoryMCvssV31) GetIntegrityImpactOk() (*string, bool)`

GetIntegrityImpactOk returns a tuple with the IntegrityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityImpact

`func (o *AdvisoryMCvssV31) SetIntegrityImpact(v string)`

SetIntegrityImpact sets IntegrityImpact field to given value.

### HasIntegrityImpact

`func (o *AdvisoryMCvssV31) HasIntegrityImpact() bool`

HasIntegrityImpact returns a boolean if a field has been set.

### GetPrivilegesRequired

`func (o *AdvisoryMCvssV31) GetPrivilegesRequired() string`

GetPrivilegesRequired returns the PrivilegesRequired field if non-nil, zero value otherwise.

### GetPrivilegesRequiredOk

`func (o *AdvisoryMCvssV31) GetPrivilegesRequiredOk() (*string, bool)`

GetPrivilegesRequiredOk returns a tuple with the PrivilegesRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegesRequired

`func (o *AdvisoryMCvssV31) SetPrivilegesRequired(v string)`

SetPrivilegesRequired sets PrivilegesRequired field to given value.

### HasPrivilegesRequired

`func (o *AdvisoryMCvssV31) HasPrivilegesRequired() bool`

HasPrivilegesRequired returns a boolean if a field has been set.

### GetScope

`func (o *AdvisoryMCvssV31) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AdvisoryMCvssV31) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AdvisoryMCvssV31) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AdvisoryMCvssV31) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUserInteraction

`func (o *AdvisoryMCvssV31) GetUserInteraction() string`

GetUserInteraction returns the UserInteraction field if non-nil, zero value otherwise.

### GetUserInteractionOk

`func (o *AdvisoryMCvssV31) GetUserInteractionOk() (*string, bool)`

GetUserInteractionOk returns a tuple with the UserInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInteraction

`func (o *AdvisoryMCvssV31) SetUserInteraction(v string)`

SetUserInteraction sets UserInteraction field to given value.

### HasUserInteraction

`func (o *AdvisoryMCvssV31) HasUserInteraction() bool`

HasUserInteraction returns a boolean if a field has been set.

### GetVectorString

`func (o *AdvisoryMCvssV31) GetVectorString() string`

GetVectorString returns the VectorString field if non-nil, zero value otherwise.

### GetVectorStringOk

`func (o *AdvisoryMCvssV31) GetVectorStringOk() (*string, bool)`

GetVectorStringOk returns a tuple with the VectorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorString

`func (o *AdvisoryMCvssV31) SetVectorString(v string)`

SetVectorString sets VectorString field to given value.

### HasVectorString

`func (o *AdvisoryMCvssV31) HasVectorString() bool`

HasVectorString returns a boolean if a field has been set.

### GetVersion

`func (o *AdvisoryMCvssV31) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AdvisoryMCvssV31) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AdvisoryMCvssV31) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AdvisoryMCvssV31) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


