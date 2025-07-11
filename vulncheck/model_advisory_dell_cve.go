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

// checks if the AdvisoryDellCVE type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryDellCVE{}

// AdvisoryDellCVE struct for AdvisoryDellCVE
type AdvisoryDellCVE struct {
	Cve *string `json:"cve,omitempty"`
	CvssScore *string `json:"cvss_score,omitempty"`
	CvssVector *string `json:"cvss_vector,omitempty"`
}

// NewAdvisoryDellCVE instantiates a new AdvisoryDellCVE object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryDellCVE() *AdvisoryDellCVE {
	this := AdvisoryDellCVE{}
	return &this
}

// NewAdvisoryDellCVEWithDefaults instantiates a new AdvisoryDellCVE object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryDellCVEWithDefaults() *AdvisoryDellCVE {
	this := AdvisoryDellCVE{}
	return &this
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryDellCVE) GetCve() string {
	if o == nil || IsNil(o.Cve) {
		var ret string
		return ret
	}
	return *o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryDellCVE) GetCveOk() (*string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryDellCVE) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given string and assigns it to the Cve field.
func (o *AdvisoryDellCVE) SetCve(v string) {
	o.Cve = &v
}

// GetCvssScore returns the CvssScore field value if set, zero value otherwise.
func (o *AdvisoryDellCVE) GetCvssScore() string {
	if o == nil || IsNil(o.CvssScore) {
		var ret string
		return ret
	}
	return *o.CvssScore
}

// GetCvssScoreOk returns a tuple with the CvssScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryDellCVE) GetCvssScoreOk() (*string, bool) {
	if o == nil || IsNil(o.CvssScore) {
		return nil, false
	}
	return o.CvssScore, true
}

// HasCvssScore returns a boolean if a field has been set.
func (o *AdvisoryDellCVE) HasCvssScore() bool {
	if o != nil && !IsNil(o.CvssScore) {
		return true
	}

	return false
}

// SetCvssScore gets a reference to the given string and assigns it to the CvssScore field.
func (o *AdvisoryDellCVE) SetCvssScore(v string) {
	o.CvssScore = &v
}

// GetCvssVector returns the CvssVector field value if set, zero value otherwise.
func (o *AdvisoryDellCVE) GetCvssVector() string {
	if o == nil || IsNil(o.CvssVector) {
		var ret string
		return ret
	}
	return *o.CvssVector
}

// GetCvssVectorOk returns a tuple with the CvssVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryDellCVE) GetCvssVectorOk() (*string, bool) {
	if o == nil || IsNil(o.CvssVector) {
		return nil, false
	}
	return o.CvssVector, true
}

// HasCvssVector returns a boolean if a field has been set.
func (o *AdvisoryDellCVE) HasCvssVector() bool {
	if o != nil && !IsNil(o.CvssVector) {
		return true
	}

	return false
}

// SetCvssVector gets a reference to the given string and assigns it to the CvssVector field.
func (o *AdvisoryDellCVE) SetCvssVector(v string) {
	o.CvssVector = &v
}

func (o AdvisoryDellCVE) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryDellCVE) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.CvssScore) {
		toSerialize["cvss_score"] = o.CvssScore
	}
	if !IsNil(o.CvssVector) {
		toSerialize["cvss_vector"] = o.CvssVector
	}
	return toSerialize, nil
}

type NullableAdvisoryDellCVE struct {
	value *AdvisoryDellCVE
	isSet bool
}

func (v NullableAdvisoryDellCVE) Get() *AdvisoryDellCVE {
	return v.value
}

func (v *NullableAdvisoryDellCVE) Set(val *AdvisoryDellCVE) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryDellCVE) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryDellCVE) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryDellCVE(val *AdvisoryDellCVE) *NullableAdvisoryDellCVE {
	return &NullableAdvisoryDellCVE{value: val, isSet: true}
}

func (v NullableAdvisoryDellCVE) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryDellCVE) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


