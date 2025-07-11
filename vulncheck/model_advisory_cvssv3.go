/*
VulnCheck API

Version 3 of the VulnCheck API

API version: 3.0
Contact: support@vulncheck.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vulncheck

import (
	"encoding/json"
)

// checks if the AdvisoryCVSSV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryCVSSV3{}

// AdvisoryCVSSV3 struct for AdvisoryCVSSV3
type AdvisoryCVSSV3 struct {
	AttackComplexity *string `json:"attackComplexity,omitempty"`
	AttackVector *string `json:"attackVector,omitempty"`
	AvailabilityImpact *string `json:"availabilityImpact,omitempty"`
	BaseScore *float32 `json:"baseScore,omitempty"`
	BaseSeverity *string `json:"baseSeverity,omitempty"`
	ConfidentialityImpact *string `json:"confidentialityImpact,omitempty"`
	IntegrityImpact *string `json:"integrityImpact,omitempty"`
	PrivilegesRequired *string `json:"privilegesRequired,omitempty"`
	Scope *string `json:"scope,omitempty"`
	UserInteraction *string `json:"userInteraction,omitempty"`
	VectorString *string `json:"vectorString,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewAdvisoryCVSSV3 instantiates a new AdvisoryCVSSV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryCVSSV3() *AdvisoryCVSSV3 {
	this := AdvisoryCVSSV3{}
	return &this
}

// NewAdvisoryCVSSV3WithDefaults instantiates a new AdvisoryCVSSV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryCVSSV3WithDefaults() *AdvisoryCVSSV3 {
	this := AdvisoryCVSSV3{}
	return &this
}

// GetAttackComplexity returns the AttackComplexity field value if set, zero value otherwise.
func (o *AdvisoryCVSSV3) GetAttackComplexity() string {
	if o == nil || IsNil(o.AttackComplexity) {
		var ret string
		return ret
	}
	return *o.AttackComplexity
}

// GetAttackComplexityOk returns a tuple with the AttackComplexity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCVSSV3) GetAttackComplexityOk() (*string, bool) {
	if o == nil || IsNil(o.AttackComplexity) {
		return nil, false
	}
	return o.AttackComplexity, true
}

// HasAttackComplexity returns a boolean if a field has been set.
func (o *AdvisoryCVSSV3) HasAttackComplexity() bool {
	if o != nil && !IsNil(o.AttackComplexity) {
		return true
	}

	return false
}

// SetAttackComplexity gets a reference to the given string and assigns it to the AttackComplexity field.
func (o *AdvisoryCVSSV3) SetAttackComplexity(v string) {
	o.AttackComplexity = &v
}

// GetAttackVector returns the AttackVector field value if set, zero value otherwise.
func (o *AdvisoryCVSSV3) GetAttackVector() string {
	if o == nil || IsNil(o.AttackVector) {
		var ret string
		return ret
	}
	return *o.AttackVector
}

// GetAttackVectorOk returns a tuple with the AttackVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCVSSV3) GetAttackVectorOk() (*string, bool) {
	if o == nil || IsNil(o.AttackVector) {
		return nil, false
	}
	return o.AttackVector, true
}

// HasAttackVector returns a boolean if a field has been set.
func (o *AdvisoryCVSSV3) HasAttackVector() bool {
	if o != nil && !IsNil(o.AttackVector) {
		return true
	}

	return false
}

// SetAttackVector gets a reference to the given string and assigns it to the AttackVector field.
func (o *AdvisoryCVSSV3) SetAttackVector(v string) {
	o.AttackVector = &v
}

// GetAvailabilityImpact returns the AvailabilityImpact field value if set, zero value otherwise.
func (o *AdvisoryCVSSV3) GetAvailabilityImpact() string {
	if o == nil || IsNil(o.AvailabilityImpact) {
		var ret string
		return ret
	}
	return *o.AvailabilityImpact
}

// GetAvailabilityImpactOk returns a tuple with the AvailabilityImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCVSSV3) GetAvailabilityImpactOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityImpact) {
		return nil, false
	}
	return o.AvailabilityImpact, true
}

// HasAvailabilityImpact returns a boolean if a field has been set.
func (o *AdvisoryCVSSV3) HasAvailabilityImpact() bool {
	if o != nil && !IsNil(o.AvailabilityImpact) {
		return true
	}

	return false
}

// SetAvailabilityImpact gets a reference to the given string and assigns it to the AvailabilityImpact field.
func (o *AdvisoryCVSSV3) SetAvailabilityImpact(v string) {
	o.AvailabilityImpact = &v
}

// GetBaseScore returns the BaseScore field value if set, zero value otherwise.
func (o *AdvisoryCVSSV3) GetBaseScore() float32 {
	if o == nil || IsNil(o.BaseScore) {
		var ret float32
		return ret
	}
	return *o.BaseScore
}

// GetBaseScoreOk returns a tuple with the BaseScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCVSSV3) GetBaseScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseScore) {
		return nil, false
	}
	return o.BaseScore, true
}

// HasBaseScore returns a boolean if a field has been set.
func (o *AdvisoryCVSSV3) HasBaseScore() bool {
	if o != nil && !IsNil(o.BaseScore) {
		return true
	}

	return false
}

// SetBaseScore gets a reference to the given float32 and assigns it to the BaseScore field.
func (o *AdvisoryCVSSV3) SetBaseScore(v float32) {
	o.BaseScore = &v
}

// GetBaseSeverity returns the BaseSeverity field value if set, zero value otherwise.
func (o *AdvisoryCVSSV3) GetBaseSeverity() string {
	if o == nil || IsNil(o.BaseSeverity) {
		var ret string
		return ret
	}
	return *o.BaseSeverity
}

// GetBaseSeverityOk returns a tuple with the BaseSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCVSSV3) GetBaseSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.BaseSeverity) {
		return nil, false
	}
	return o.BaseSeverity, true
}

// HasBaseSeverity returns a boolean if a field has been set.
func (o *AdvisoryCVSSV3) HasBaseSeverity() bool {
	if o != nil && !IsNil(o.BaseSeverity) {
		return true
	}

	return false
}

// SetBaseSeverity gets a reference to the given string and assigns it to the BaseSeverity field.
func (o *AdvisoryCVSSV3) SetBaseSeverity(v string) {
	o.BaseSeverity = &v
}

// GetConfidentialityImpact returns the ConfidentialityImpact field value if set, zero value otherwise.
func (o *AdvisoryCVSSV3) GetConfidentialityImpact() string {
	if o == nil || IsNil(o.ConfidentialityImpact) {
		var ret string
		return ret
	}
	return *o.ConfidentialityImpact
}

// GetConfidentialityImpactOk returns a tuple with the ConfidentialityImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCVSSV3) GetConfidentialityImpactOk() (*string, bool) {
	if o == nil || IsNil(o.ConfidentialityImpact) {
		return nil, false
	}
	return o.ConfidentialityImpact, true
}

// HasConfidentialityImpact returns a boolean if a field has been set.
func (o *AdvisoryCVSSV3) HasConfidentialityImpact() bool {
	if o != nil && !IsNil(o.ConfidentialityImpact) {
		return true
	}

	return false
}

// SetConfidentialityImpact gets a reference to the given string and assigns it to the ConfidentialityImpact field.
func (o *AdvisoryCVSSV3) SetConfidentialityImpact(v string) {
	o.ConfidentialityImpact = &v
}

// GetIntegrityImpact returns the IntegrityImpact field value if set, zero value otherwise.
func (o *AdvisoryCVSSV3) GetIntegrityImpact() string {
	if o == nil || IsNil(o.IntegrityImpact) {
		var ret string
		return ret
	}
	return *o.IntegrityImpact
}

// GetIntegrityImpactOk returns a tuple with the IntegrityImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCVSSV3) GetIntegrityImpactOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrityImpact) {
		return nil, false
	}
	return o.IntegrityImpact, true
}

// HasIntegrityImpact returns a boolean if a field has been set.
func (o *AdvisoryCVSSV3) HasIntegrityImpact() bool {
	if o != nil && !IsNil(o.IntegrityImpact) {
		return true
	}

	return false
}

// SetIntegrityImpact gets a reference to the given string and assigns it to the IntegrityImpact field.
func (o *AdvisoryCVSSV3) SetIntegrityImpact(v string) {
	o.IntegrityImpact = &v
}

// GetPrivilegesRequired returns the PrivilegesRequired field value if set, zero value otherwise.
func (o *AdvisoryCVSSV3) GetPrivilegesRequired() string {
	if o == nil || IsNil(o.PrivilegesRequired) {
		var ret string
		return ret
	}
	return *o.PrivilegesRequired
}

// GetPrivilegesRequiredOk returns a tuple with the PrivilegesRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCVSSV3) GetPrivilegesRequiredOk() (*string, bool) {
	if o == nil || IsNil(o.PrivilegesRequired) {
		return nil, false
	}
	return o.PrivilegesRequired, true
}

// HasPrivilegesRequired returns a boolean if a field has been set.
func (o *AdvisoryCVSSV3) HasPrivilegesRequired() bool {
	if o != nil && !IsNil(o.PrivilegesRequired) {
		return true
	}

	return false
}

// SetPrivilegesRequired gets a reference to the given string and assigns it to the PrivilegesRequired field.
func (o *AdvisoryCVSSV3) SetPrivilegesRequired(v string) {
	o.PrivilegesRequired = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AdvisoryCVSSV3) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCVSSV3) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AdvisoryCVSSV3) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *AdvisoryCVSSV3) SetScope(v string) {
	o.Scope = &v
}

// GetUserInteraction returns the UserInteraction field value if set, zero value otherwise.
func (o *AdvisoryCVSSV3) GetUserInteraction() string {
	if o == nil || IsNil(o.UserInteraction) {
		var ret string
		return ret
	}
	return *o.UserInteraction
}

// GetUserInteractionOk returns a tuple with the UserInteraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCVSSV3) GetUserInteractionOk() (*string, bool) {
	if o == nil || IsNil(o.UserInteraction) {
		return nil, false
	}
	return o.UserInteraction, true
}

// HasUserInteraction returns a boolean if a field has been set.
func (o *AdvisoryCVSSV3) HasUserInteraction() bool {
	if o != nil && !IsNil(o.UserInteraction) {
		return true
	}

	return false
}

// SetUserInteraction gets a reference to the given string and assigns it to the UserInteraction field.
func (o *AdvisoryCVSSV3) SetUserInteraction(v string) {
	o.UserInteraction = &v
}

// GetVectorString returns the VectorString field value if set, zero value otherwise.
func (o *AdvisoryCVSSV3) GetVectorString() string {
	if o == nil || IsNil(o.VectorString) {
		var ret string
		return ret
	}
	return *o.VectorString
}

// GetVectorStringOk returns a tuple with the VectorString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCVSSV3) GetVectorStringOk() (*string, bool) {
	if o == nil || IsNil(o.VectorString) {
		return nil, false
	}
	return o.VectorString, true
}

// HasVectorString returns a boolean if a field has been set.
func (o *AdvisoryCVSSV3) HasVectorString() bool {
	if o != nil && !IsNil(o.VectorString) {
		return true
	}

	return false
}

// SetVectorString gets a reference to the given string and assigns it to the VectorString field.
func (o *AdvisoryCVSSV3) SetVectorString(v string) {
	o.VectorString = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AdvisoryCVSSV3) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCVSSV3) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AdvisoryCVSSV3) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AdvisoryCVSSV3) SetVersion(v string) {
	o.Version = &v
}

func (o AdvisoryCVSSV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryCVSSV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttackComplexity) {
		toSerialize["attackComplexity"] = o.AttackComplexity
	}
	if !IsNil(o.AttackVector) {
		toSerialize["attackVector"] = o.AttackVector
	}
	if !IsNil(o.AvailabilityImpact) {
		toSerialize["availabilityImpact"] = o.AvailabilityImpact
	}
	if !IsNil(o.BaseScore) {
		toSerialize["baseScore"] = o.BaseScore
	}
	if !IsNil(o.BaseSeverity) {
		toSerialize["baseSeverity"] = o.BaseSeverity
	}
	if !IsNil(o.ConfidentialityImpact) {
		toSerialize["confidentialityImpact"] = o.ConfidentialityImpact
	}
	if !IsNil(o.IntegrityImpact) {
		toSerialize["integrityImpact"] = o.IntegrityImpact
	}
	if !IsNil(o.PrivilegesRequired) {
		toSerialize["privilegesRequired"] = o.PrivilegesRequired
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.UserInteraction) {
		toSerialize["userInteraction"] = o.UserInteraction
	}
	if !IsNil(o.VectorString) {
		toSerialize["vectorString"] = o.VectorString
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableAdvisoryCVSSV3 struct {
	value *AdvisoryCVSSV3
	isSet bool
}

func (v NullableAdvisoryCVSSV3) Get() *AdvisoryCVSSV3 {
	return v.value
}

func (v *NullableAdvisoryCVSSV3) Set(val *AdvisoryCVSSV3) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryCVSSV3) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryCVSSV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryCVSSV3(val *AdvisoryCVSSV3) *NullableAdvisoryCVSSV3 {
	return &NullableAdvisoryCVSSV3{value: val, isSet: true}
}

func (v NullableAdvisoryCVSSV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryCVSSV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


