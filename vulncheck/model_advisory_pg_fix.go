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

// checks if the AdvisoryPGFix type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryPGFix{}

// AdvisoryPGFix struct for AdvisoryPGFix
type AdvisoryPGFix struct {
	Affected *string `json:"affected,omitempty"`
	Fixed *string `json:"fixed,omitempty"`
}

// NewAdvisoryPGFix instantiates a new AdvisoryPGFix object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryPGFix() *AdvisoryPGFix {
	this := AdvisoryPGFix{}
	return &this
}

// NewAdvisoryPGFixWithDefaults instantiates a new AdvisoryPGFix object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryPGFixWithDefaults() *AdvisoryPGFix {
	this := AdvisoryPGFix{}
	return &this
}

// GetAffected returns the Affected field value if set, zero value otherwise.
func (o *AdvisoryPGFix) GetAffected() string {
	if o == nil || IsNil(o.Affected) {
		var ret string
		return ret
	}
	return *o.Affected
}

// GetAffectedOk returns a tuple with the Affected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPGFix) GetAffectedOk() (*string, bool) {
	if o == nil || IsNil(o.Affected) {
		return nil, false
	}
	return o.Affected, true
}

// HasAffected returns a boolean if a field has been set.
func (o *AdvisoryPGFix) HasAffected() bool {
	if o != nil && !IsNil(o.Affected) {
		return true
	}

	return false
}

// SetAffected gets a reference to the given string and assigns it to the Affected field.
func (o *AdvisoryPGFix) SetAffected(v string) {
	o.Affected = &v
}

// GetFixed returns the Fixed field value if set, zero value otherwise.
func (o *AdvisoryPGFix) GetFixed() string {
	if o == nil || IsNil(o.Fixed) {
		var ret string
		return ret
	}
	return *o.Fixed
}

// GetFixedOk returns a tuple with the Fixed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPGFix) GetFixedOk() (*string, bool) {
	if o == nil || IsNil(o.Fixed) {
		return nil, false
	}
	return o.Fixed, true
}

// HasFixed returns a boolean if a field has been set.
func (o *AdvisoryPGFix) HasFixed() bool {
	if o != nil && !IsNil(o.Fixed) {
		return true
	}

	return false
}

// SetFixed gets a reference to the given string and assigns it to the Fixed field.
func (o *AdvisoryPGFix) SetFixed(v string) {
	o.Fixed = &v
}

func (o AdvisoryPGFix) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryPGFix) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Affected) {
		toSerialize["affected"] = o.Affected
	}
	if !IsNil(o.Fixed) {
		toSerialize["fixed"] = o.Fixed
	}
	return toSerialize, nil
}

type NullableAdvisoryPGFix struct {
	value *AdvisoryPGFix
	isSet bool
}

func (v NullableAdvisoryPGFix) Get() *AdvisoryPGFix {
	return v.value
}

func (v *NullableAdvisoryPGFix) Set(val *AdvisoryPGFix) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryPGFix) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryPGFix) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryPGFix(val *AdvisoryPGFix) *NullableAdvisoryPGFix {
	return &NullableAdvisoryPGFix{value: val, isSet: true}
}

func (v NullableAdvisoryPGFix) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryPGFix) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


