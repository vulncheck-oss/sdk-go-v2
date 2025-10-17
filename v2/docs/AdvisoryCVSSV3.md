# AdvisoryCVSSV3

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

### NewAdvisoryCVSSV3

`func NewAdvisoryCVSSV3() *AdvisoryCVSSV3`

NewAdvisoryCVSSV3 instantiates a new AdvisoryCVSSV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCVSSV3WithDefaults

`func NewAdvisoryCVSSV3WithDefaults() *AdvisoryCVSSV3`

NewAdvisoryCVSSV3WithDefaults instantiates a new AdvisoryCVSSV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttackComplexity

`func (o *AdvisoryCVSSV3) GetAttackComplexity() string`

GetAttackComplexity returns the AttackComplexity field if non-nil, zero value otherwise.

### GetAttackComplexityOk

`func (o *AdvisoryCVSSV3) GetAttackComplexityOk() (*string, bool)`

GetAttackComplexityOk returns a tuple with the AttackComplexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackComplexity

`func (o *AdvisoryCVSSV3) SetAttackComplexity(v string)`

SetAttackComplexity sets AttackComplexity field to given value.

### HasAttackComplexity

`func (o *AdvisoryCVSSV3) HasAttackComplexity() bool`

HasAttackComplexity returns a boolean if a field has been set.

### GetAttackVector

`func (o *AdvisoryCVSSV3) GetAttackVector() string`

GetAttackVector returns the AttackVector field if non-nil, zero value otherwise.

### GetAttackVectorOk

`func (o *AdvisoryCVSSV3) GetAttackVectorOk() (*string, bool)`

GetAttackVectorOk returns a tuple with the AttackVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackVector

`func (o *AdvisoryCVSSV3) SetAttackVector(v string)`

SetAttackVector sets AttackVector field to given value.

### HasAttackVector

`func (o *AdvisoryCVSSV3) HasAttackVector() bool`

HasAttackVector returns a boolean if a field has been set.

### GetAvailabilityImpact

`func (o *AdvisoryCVSSV3) GetAvailabilityImpact() string`

GetAvailabilityImpact returns the AvailabilityImpact field if non-nil, zero value otherwise.

### GetAvailabilityImpactOk

`func (o *AdvisoryCVSSV3) GetAvailabilityImpactOk() (*string, bool)`

GetAvailabilityImpactOk returns a tuple with the AvailabilityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityImpact

`func (o *AdvisoryCVSSV3) SetAvailabilityImpact(v string)`

SetAvailabilityImpact sets AvailabilityImpact field to given value.

### HasAvailabilityImpact

`func (o *AdvisoryCVSSV3) HasAvailabilityImpact() bool`

HasAvailabilityImpact returns a boolean if a field has been set.

### GetBaseScore

`func (o *AdvisoryCVSSV3) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *AdvisoryCVSSV3) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *AdvisoryCVSSV3) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *AdvisoryCVSSV3) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetBaseSeverity

`func (o *AdvisoryCVSSV3) GetBaseSeverity() string`

GetBaseSeverity returns the BaseSeverity field if non-nil, zero value otherwise.

### GetBaseSeverityOk

`func (o *AdvisoryCVSSV3) GetBaseSeverityOk() (*string, bool)`

GetBaseSeverityOk returns a tuple with the BaseSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSeverity

`func (o *AdvisoryCVSSV3) SetBaseSeverity(v string)`

SetBaseSeverity sets BaseSeverity field to given value.

### HasBaseSeverity

`func (o *AdvisoryCVSSV3) HasBaseSeverity() bool`

HasBaseSeverity returns a boolean if a field has been set.

### GetConfidentialityImpact

`func (o *AdvisoryCVSSV3) GetConfidentialityImpact() string`

GetConfidentialityImpact returns the ConfidentialityImpact field if non-nil, zero value otherwise.

### GetConfidentialityImpactOk

`func (o *AdvisoryCVSSV3) GetConfidentialityImpactOk() (*string, bool)`

GetConfidentialityImpactOk returns a tuple with the ConfidentialityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialityImpact

`func (o *AdvisoryCVSSV3) SetConfidentialityImpact(v string)`

SetConfidentialityImpact sets ConfidentialityImpact field to given value.

### HasConfidentialityImpact

`func (o *AdvisoryCVSSV3) HasConfidentialityImpact() bool`

HasConfidentialityImpact returns a boolean if a field has been set.

### GetIntegrityImpact

`func (o *AdvisoryCVSSV3) GetIntegrityImpact() string`

GetIntegrityImpact returns the IntegrityImpact field if non-nil, zero value otherwise.

### GetIntegrityImpactOk

`func (o *AdvisoryCVSSV3) GetIntegrityImpactOk() (*string, bool)`

GetIntegrityImpactOk returns a tuple with the IntegrityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityImpact

`func (o *AdvisoryCVSSV3) SetIntegrityImpact(v string)`

SetIntegrityImpact sets IntegrityImpact field to given value.

### HasIntegrityImpact

`func (o *AdvisoryCVSSV3) HasIntegrityImpact() bool`

HasIntegrityImpact returns a boolean if a field has been set.

### GetPrivilegesRequired

`func (o *AdvisoryCVSSV3) GetPrivilegesRequired() string`

GetPrivilegesRequired returns the PrivilegesRequired field if non-nil, zero value otherwise.

### GetPrivilegesRequiredOk

`func (o *AdvisoryCVSSV3) GetPrivilegesRequiredOk() (*string, bool)`

GetPrivilegesRequiredOk returns a tuple with the PrivilegesRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegesRequired

`func (o *AdvisoryCVSSV3) SetPrivilegesRequired(v string)`

SetPrivilegesRequired sets PrivilegesRequired field to given value.

### HasPrivilegesRequired

`func (o *AdvisoryCVSSV3) HasPrivilegesRequired() bool`

HasPrivilegesRequired returns a boolean if a field has been set.

### GetScope

`func (o *AdvisoryCVSSV3) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AdvisoryCVSSV3) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AdvisoryCVSSV3) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AdvisoryCVSSV3) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUserInteraction

`func (o *AdvisoryCVSSV3) GetUserInteraction() string`

GetUserInteraction returns the UserInteraction field if non-nil, zero value otherwise.

### GetUserInteractionOk

`func (o *AdvisoryCVSSV3) GetUserInteractionOk() (*string, bool)`

GetUserInteractionOk returns a tuple with the UserInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInteraction

`func (o *AdvisoryCVSSV3) SetUserInteraction(v string)`

SetUserInteraction sets UserInteraction field to given value.

### HasUserInteraction

`func (o *AdvisoryCVSSV3) HasUserInteraction() bool`

HasUserInteraction returns a boolean if a field has been set.

### GetVectorString

`func (o *AdvisoryCVSSV3) GetVectorString() string`

GetVectorString returns the VectorString field if non-nil, zero value otherwise.

### GetVectorStringOk

`func (o *AdvisoryCVSSV3) GetVectorStringOk() (*string, bool)`

GetVectorStringOk returns a tuple with the VectorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorString

`func (o *AdvisoryCVSSV3) SetVectorString(v string)`

SetVectorString sets VectorString field to given value.

### HasVectorString

`func (o *AdvisoryCVSSV3) HasVectorString() bool`

HasVectorString returns a boolean if a field has been set.

### GetVersion

`func (o *AdvisoryCVSSV3) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AdvisoryCVSSV3) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AdvisoryCVSSV3) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AdvisoryCVSSV3) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


