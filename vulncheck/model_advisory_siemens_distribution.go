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

// checks if the AdvisorySiemensDistribution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisorySiemensDistribution{}

// AdvisorySiemensDistribution struct for AdvisorySiemensDistribution
type AdvisorySiemensDistribution struct {
	Text *string `json:"text,omitempty"`
	Tlp *AdvisorySiemensTLP `json:"tlp,omitempty"`
}

// NewAdvisorySiemensDistribution instantiates a new AdvisorySiemensDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisorySiemensDistribution() *AdvisorySiemensDistribution {
	this := AdvisorySiemensDistribution{}
	return &this
}

// NewAdvisorySiemensDistributionWithDefaults instantiates a new AdvisorySiemensDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisorySiemensDistributionWithDefaults() *AdvisorySiemensDistribution {
	this := AdvisorySiemensDistribution{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *AdvisorySiemensDistribution) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySiemensDistribution) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *AdvisorySiemensDistribution) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *AdvisorySiemensDistribution) SetText(v string) {
	o.Text = &v
}

// GetTlp returns the Tlp field value if set, zero value otherwise.
func (o *AdvisorySiemensDistribution) GetTlp() AdvisorySiemensTLP {
	if o == nil || IsNil(o.Tlp) {
		var ret AdvisorySiemensTLP
		return ret
	}
	return *o.Tlp
}

// GetTlpOk returns a tuple with the Tlp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySiemensDistribution) GetTlpOk() (*AdvisorySiemensTLP, bool) {
	if o == nil || IsNil(o.Tlp) {
		return nil, false
	}
	return o.Tlp, true
}

// HasTlp returns a boolean if a field has been set.
func (o *AdvisorySiemensDistribution) HasTlp() bool {
	if o != nil && !IsNil(o.Tlp) {
		return true
	}

	return false
}

// SetTlp gets a reference to the given AdvisorySiemensTLP and assigns it to the Tlp field.
func (o *AdvisorySiemensDistribution) SetTlp(v AdvisorySiemensTLP) {
	o.Tlp = &v
}

func (o AdvisorySiemensDistribution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisorySiemensDistribution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Tlp) {
		toSerialize["tlp"] = o.Tlp
	}
	return toSerialize, nil
}

type NullableAdvisorySiemensDistribution struct {
	value *AdvisorySiemensDistribution
	isSet bool
}

func (v NullableAdvisorySiemensDistribution) Get() *AdvisorySiemensDistribution {
	return v.value
}

func (v *NullableAdvisorySiemensDistribution) Set(val *AdvisorySiemensDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisorySiemensDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisorySiemensDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisorySiemensDistribution(val *AdvisorySiemensDistribution) *NullableAdvisorySiemensDistribution {
	return &NullableAdvisorySiemensDistribution{value: val, isSet: true}
}

func (v NullableAdvisorySiemensDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisorySiemensDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


