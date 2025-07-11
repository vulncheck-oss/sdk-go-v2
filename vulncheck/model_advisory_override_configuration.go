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

// checks if the AdvisoryOverrideConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryOverrideConfiguration{}

// AdvisoryOverrideConfiguration struct for AdvisoryOverrideConfiguration
type AdvisoryOverrideConfiguration struct {
	Nodes []AdvisoryCPENode `json:"nodes,omitempty"`
}

// NewAdvisoryOverrideConfiguration instantiates a new AdvisoryOverrideConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryOverrideConfiguration() *AdvisoryOverrideConfiguration {
	this := AdvisoryOverrideConfiguration{}
	return &this
}

// NewAdvisoryOverrideConfigurationWithDefaults instantiates a new AdvisoryOverrideConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryOverrideConfigurationWithDefaults() *AdvisoryOverrideConfiguration {
	this := AdvisoryOverrideConfiguration{}
	return &this
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *AdvisoryOverrideConfiguration) GetNodes() []AdvisoryCPENode {
	if o == nil || IsNil(o.Nodes) {
		var ret []AdvisoryCPENode
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOverrideConfiguration) GetNodesOk() ([]AdvisoryCPENode, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *AdvisoryOverrideConfiguration) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []AdvisoryCPENode and assigns it to the Nodes field.
func (o *AdvisoryOverrideConfiguration) SetNodes(v []AdvisoryCPENode) {
	o.Nodes = v
}

func (o AdvisoryOverrideConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryOverrideConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	return toSerialize, nil
}

type NullableAdvisoryOverrideConfiguration struct {
	value *AdvisoryOverrideConfiguration
	isSet bool
}

func (v NullableAdvisoryOverrideConfiguration) Get() *AdvisoryOverrideConfiguration {
	return v.value
}

func (v *NullableAdvisoryOverrideConfiguration) Set(val *AdvisoryOverrideConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryOverrideConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryOverrideConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryOverrideConfiguration(val *AdvisoryOverrideConfiguration) *NullableAdvisoryOverrideConfiguration {
	return &NullableAdvisoryOverrideConfiguration{value: val, isSet: true}
}

func (v NullableAdvisoryOverrideConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryOverrideConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


