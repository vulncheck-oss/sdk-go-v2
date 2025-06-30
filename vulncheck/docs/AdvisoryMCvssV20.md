# AdvisoryMCvssV20

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessVector** | Pointer to **string** |  | [optional] 
**AttackComplexity** | Pointer to **string** |  | [optional] 
**Authentication** | Pointer to **string** |  | [optional] 
**AvailabilityImpact** | Pointer to **string** |  | [optional] 
**BaseScore** | Pointer to **float32** |  | [optional] 
**ConfidentialityImpact** | Pointer to **string** |  | [optional] 
**IntegrityImpact** | Pointer to **string** |  | [optional] 
**VectorString** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryMCvssV20

`func NewAdvisoryMCvssV20() *AdvisoryMCvssV20`

NewAdvisoryMCvssV20 instantiates a new AdvisoryMCvssV20 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMCvssV20WithDefaults

`func NewAdvisoryMCvssV20WithDefaults() *AdvisoryMCvssV20`

NewAdvisoryMCvssV20WithDefaults instantiates a new AdvisoryMCvssV20 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessVector

`func (o *AdvisoryMCvssV20) GetAccessVector() string`

GetAccessVector returns the AccessVector field if non-nil, zero value otherwise.

### GetAccessVectorOk

`func (o *AdvisoryMCvssV20) GetAccessVectorOk() (*string, bool)`

GetAccessVectorOk returns a tuple with the AccessVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessVector

`func (o *AdvisoryMCvssV20) SetAccessVector(v string)`

SetAccessVector sets AccessVector field to given value.

### HasAccessVector

`func (o *AdvisoryMCvssV20) HasAccessVector() bool`

HasAccessVector returns a boolean if a field has been set.

### GetAttackComplexity

`func (o *AdvisoryMCvssV20) GetAttackComplexity() string`

GetAttackComplexity returns the AttackComplexity field if non-nil, zero value otherwise.

### GetAttackComplexityOk

`func (o *AdvisoryMCvssV20) GetAttackComplexityOk() (*string, bool)`

GetAttackComplexityOk returns a tuple with the AttackComplexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackComplexity

`func (o *AdvisoryMCvssV20) SetAttackComplexity(v string)`

SetAttackComplexity sets AttackComplexity field to given value.

### HasAttackComplexity

`func (o *AdvisoryMCvssV20) HasAttackComplexity() bool`

HasAttackComplexity returns a boolean if a field has been set.

### GetAuthentication

`func (o *AdvisoryMCvssV20) GetAuthentication() string`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *AdvisoryMCvssV20) GetAuthenticationOk() (*string, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *AdvisoryMCvssV20) SetAuthentication(v string)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *AdvisoryMCvssV20) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetAvailabilityImpact

`func (o *AdvisoryMCvssV20) GetAvailabilityImpact() string`

GetAvailabilityImpact returns the AvailabilityImpact field if non-nil, zero value otherwise.

### GetAvailabilityImpactOk

`func (o *AdvisoryMCvssV20) GetAvailabilityImpactOk() (*string, bool)`

GetAvailabilityImpactOk returns a tuple with the AvailabilityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityImpact

`func (o *AdvisoryMCvssV20) SetAvailabilityImpact(v string)`

SetAvailabilityImpact sets AvailabilityImpact field to given value.

### HasAvailabilityImpact

`func (o *AdvisoryMCvssV20) HasAvailabilityImpact() bool`

HasAvailabilityImpact returns a boolean if a field has been set.

### GetBaseScore

`func (o *AdvisoryMCvssV20) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *AdvisoryMCvssV20) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *AdvisoryMCvssV20) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *AdvisoryMCvssV20) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetConfidentialityImpact

`func (o *AdvisoryMCvssV20) GetConfidentialityImpact() string`

GetConfidentialityImpact returns the ConfidentialityImpact field if non-nil, zero value otherwise.

### GetConfidentialityImpactOk

`func (o *AdvisoryMCvssV20) GetConfidentialityImpactOk() (*string, bool)`

GetConfidentialityImpactOk returns a tuple with the ConfidentialityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialityImpact

`func (o *AdvisoryMCvssV20) SetConfidentialityImpact(v string)`

SetConfidentialityImpact sets ConfidentialityImpact field to given value.

### HasConfidentialityImpact

`func (o *AdvisoryMCvssV20) HasConfidentialityImpact() bool`

HasConfidentialityImpact returns a boolean if a field has been set.

### GetIntegrityImpact

`func (o *AdvisoryMCvssV20) GetIntegrityImpact() string`

GetIntegrityImpact returns the IntegrityImpact field if non-nil, zero value otherwise.

### GetIntegrityImpactOk

`func (o *AdvisoryMCvssV20) GetIntegrityImpactOk() (*string, bool)`

GetIntegrityImpactOk returns a tuple with the IntegrityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityImpact

`func (o *AdvisoryMCvssV20) SetIntegrityImpact(v string)`

SetIntegrityImpact sets IntegrityImpact field to given value.

### HasIntegrityImpact

`func (o *AdvisoryMCvssV20) HasIntegrityImpact() bool`

HasIntegrityImpact returns a boolean if a field has been set.

### GetVectorString

`func (o *AdvisoryMCvssV20) GetVectorString() string`

GetVectorString returns the VectorString field if non-nil, zero value otherwise.

### GetVectorStringOk

`func (o *AdvisoryMCvssV20) GetVectorStringOk() (*string, bool)`

GetVectorStringOk returns a tuple with the VectorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorString

`func (o *AdvisoryMCvssV20) SetVectorString(v string)`

SetVectorString sets VectorString field to given value.

### HasVectorString

`func (o *AdvisoryMCvssV20) HasVectorString() bool`

HasVectorString returns a boolean if a field has been set.

### GetVersion

`func (o *AdvisoryMCvssV20) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AdvisoryMCvssV20) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AdvisoryMCvssV20) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AdvisoryMCvssV20) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


