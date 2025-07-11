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

// checks if the ApiTemporalMetricV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTemporalMetricV3{}

// ApiTemporalMetricV3 struct for ApiTemporalMetricV3
type ApiTemporalMetricV3 struct {
	CvssV3 *ApiTemporalCVSSV3 `json:"cvssV3,omitempty"`
}

// NewApiTemporalMetricV3 instantiates a new ApiTemporalMetricV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTemporalMetricV3() *ApiTemporalMetricV3 {
	this := ApiTemporalMetricV3{}
	return &this
}

// NewApiTemporalMetricV3WithDefaults instantiates a new ApiTemporalMetricV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTemporalMetricV3WithDefaults() *ApiTemporalMetricV3 {
	this := ApiTemporalMetricV3{}
	return &this
}

// GetCvssV3 returns the CvssV3 field value if set, zero value otherwise.
func (o *ApiTemporalMetricV3) GetCvssV3() ApiTemporalCVSSV3 {
	if o == nil || IsNil(o.CvssV3) {
		var ret ApiTemporalCVSSV3
		return ret
	}
	return *o.CvssV3
}

// GetCvssV3Ok returns a tuple with the CvssV3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTemporalMetricV3) GetCvssV3Ok() (*ApiTemporalCVSSV3, bool) {
	if o == nil || IsNil(o.CvssV3) {
		return nil, false
	}
	return o.CvssV3, true
}

// HasCvssV3 returns a boolean if a field has been set.
func (o *ApiTemporalMetricV3) HasCvssV3() bool {
	if o != nil && !IsNil(o.CvssV3) {
		return true
	}

	return false
}

// SetCvssV3 gets a reference to the given ApiTemporalCVSSV3 and assigns it to the CvssV3 field.
func (o *ApiTemporalMetricV3) SetCvssV3(v ApiTemporalCVSSV3) {
	o.CvssV3 = &v
}

func (o ApiTemporalMetricV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTemporalMetricV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CvssV3) {
		toSerialize["cvssV3"] = o.CvssV3
	}
	return toSerialize, nil
}

type NullableApiTemporalMetricV3 struct {
	value *ApiTemporalMetricV3
	isSet bool
}

func (v NullableApiTemporalMetricV3) Get() *ApiTemporalMetricV3 {
	return v.value
}

func (v *NullableApiTemporalMetricV3) Set(val *ApiTemporalMetricV3) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTemporalMetricV3) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTemporalMetricV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTemporalMetricV3(val *ApiTemporalMetricV3) *NullableApiTemporalMetricV3 {
	return &NullableApiTemporalMetricV3{value: val, isSet: true}
}

func (v NullableApiTemporalMetricV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTemporalMetricV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


