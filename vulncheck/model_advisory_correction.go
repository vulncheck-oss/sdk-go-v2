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

// checks if the AdvisoryCorrection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryCorrection{}

// AdvisoryCorrection struct for AdvisoryCorrection
type AdvisoryCorrection struct {
	CorrectedAt *string `json:"correctedAt,omitempty"`
	Orelease *string `json:"orelease,omitempty"`
	Release *string `json:"release,omitempty"`
}

// NewAdvisoryCorrection instantiates a new AdvisoryCorrection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryCorrection() *AdvisoryCorrection {
	this := AdvisoryCorrection{}
	return &this
}

// NewAdvisoryCorrectionWithDefaults instantiates a new AdvisoryCorrection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryCorrectionWithDefaults() *AdvisoryCorrection {
	this := AdvisoryCorrection{}
	return &this
}

// GetCorrectedAt returns the CorrectedAt field value if set, zero value otherwise.
func (o *AdvisoryCorrection) GetCorrectedAt() string {
	if o == nil || IsNil(o.CorrectedAt) {
		var ret string
		return ret
	}
	return *o.CorrectedAt
}

// GetCorrectedAtOk returns a tuple with the CorrectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCorrection) GetCorrectedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CorrectedAt) {
		return nil, false
	}
	return o.CorrectedAt, true
}

// HasCorrectedAt returns a boolean if a field has been set.
func (o *AdvisoryCorrection) HasCorrectedAt() bool {
	if o != nil && !IsNil(o.CorrectedAt) {
		return true
	}

	return false
}

// SetCorrectedAt gets a reference to the given string and assigns it to the CorrectedAt field.
func (o *AdvisoryCorrection) SetCorrectedAt(v string) {
	o.CorrectedAt = &v
}

// GetOrelease returns the Orelease field value if set, zero value otherwise.
func (o *AdvisoryCorrection) GetOrelease() string {
	if o == nil || IsNil(o.Orelease) {
		var ret string
		return ret
	}
	return *o.Orelease
}

// GetOreleaseOk returns a tuple with the Orelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCorrection) GetOreleaseOk() (*string, bool) {
	if o == nil || IsNil(o.Orelease) {
		return nil, false
	}
	return o.Orelease, true
}

// HasOrelease returns a boolean if a field has been set.
func (o *AdvisoryCorrection) HasOrelease() bool {
	if o != nil && !IsNil(o.Orelease) {
		return true
	}

	return false
}

// SetOrelease gets a reference to the given string and assigns it to the Orelease field.
func (o *AdvisoryCorrection) SetOrelease(v string) {
	o.Orelease = &v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *AdvisoryCorrection) GetRelease() string {
	if o == nil || IsNil(o.Release) {
		var ret string
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCorrection) GetReleaseOk() (*string, bool) {
	if o == nil || IsNil(o.Release) {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *AdvisoryCorrection) HasRelease() bool {
	if o != nil && !IsNil(o.Release) {
		return true
	}

	return false
}

// SetRelease gets a reference to the given string and assigns it to the Release field.
func (o *AdvisoryCorrection) SetRelease(v string) {
	o.Release = &v
}

func (o AdvisoryCorrection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryCorrection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CorrectedAt) {
		toSerialize["correctedAt"] = o.CorrectedAt
	}
	if !IsNil(o.Orelease) {
		toSerialize["orelease"] = o.Orelease
	}
	if !IsNil(o.Release) {
		toSerialize["release"] = o.Release
	}
	return toSerialize, nil
}

type NullableAdvisoryCorrection struct {
	value *AdvisoryCorrection
	isSet bool
}

func (v NullableAdvisoryCorrection) Get() *AdvisoryCorrection {
	return v.value
}

func (v *NullableAdvisoryCorrection) Set(val *AdvisoryCorrection) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryCorrection) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryCorrection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryCorrection(val *AdvisoryCorrection) *NullableAdvisoryCorrection {
	return &NullableAdvisoryCorrection{value: val, isSet: true}
}

func (v NullableAdvisoryCorrection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryCorrection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


