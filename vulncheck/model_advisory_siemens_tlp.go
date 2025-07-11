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

// checks if the AdvisorySiemensTLP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisorySiemensTLP{}

// AdvisorySiemensTLP struct for AdvisorySiemensTLP
type AdvisorySiemensTLP struct {
	Label *string `json:"label,omitempty"`
}

// NewAdvisorySiemensTLP instantiates a new AdvisorySiemensTLP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisorySiemensTLP() *AdvisorySiemensTLP {
	this := AdvisorySiemensTLP{}
	return &this
}

// NewAdvisorySiemensTLPWithDefaults instantiates a new AdvisorySiemensTLP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisorySiemensTLPWithDefaults() *AdvisorySiemensTLP {
	this := AdvisorySiemensTLP{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *AdvisorySiemensTLP) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySiemensTLP) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *AdvisorySiemensTLP) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *AdvisorySiemensTLP) SetLabel(v string) {
	o.Label = &v
}

func (o AdvisorySiemensTLP) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisorySiemensTLP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableAdvisorySiemensTLP struct {
	value *AdvisorySiemensTLP
	isSet bool
}

func (v NullableAdvisorySiemensTLP) Get() *AdvisorySiemensTLP {
	return v.value
}

func (v *NullableAdvisorySiemensTLP) Set(val *AdvisorySiemensTLP) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisorySiemensTLP) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisorySiemensTLP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisorySiemensTLP(val *AdvisorySiemensTLP) *NullableAdvisorySiemensTLP {
	return &NullableAdvisorySiemensTLP{value: val, isSet: true}
}

func (v NullableAdvisorySiemensTLP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisorySiemensTLP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


