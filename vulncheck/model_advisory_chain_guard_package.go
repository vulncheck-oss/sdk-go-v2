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

// checks if the AdvisoryChainGuardPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryChainGuardPackage{}

// AdvisoryChainGuardPackage struct for AdvisoryChainGuardPackage
type AdvisoryChainGuardPackage struct {
	Name *string `json:"name,omitempty"`
	Secfixes []AdvisoryChainGuardSecFix `json:"secfixes,omitempty"`
}

// NewAdvisoryChainGuardPackage instantiates a new AdvisoryChainGuardPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryChainGuardPackage() *AdvisoryChainGuardPackage {
	this := AdvisoryChainGuardPackage{}
	return &this
}

// NewAdvisoryChainGuardPackageWithDefaults instantiates a new AdvisoryChainGuardPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryChainGuardPackageWithDefaults() *AdvisoryChainGuardPackage {
	this := AdvisoryChainGuardPackage{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdvisoryChainGuardPackage) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryChainGuardPackage) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdvisoryChainGuardPackage) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdvisoryChainGuardPackage) SetName(v string) {
	o.Name = &v
}

// GetSecfixes returns the Secfixes field value if set, zero value otherwise.
func (o *AdvisoryChainGuardPackage) GetSecfixes() []AdvisoryChainGuardSecFix {
	if o == nil || IsNil(o.Secfixes) {
		var ret []AdvisoryChainGuardSecFix
		return ret
	}
	return o.Secfixes
}

// GetSecfixesOk returns a tuple with the Secfixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryChainGuardPackage) GetSecfixesOk() ([]AdvisoryChainGuardSecFix, bool) {
	if o == nil || IsNil(o.Secfixes) {
		return nil, false
	}
	return o.Secfixes, true
}

// HasSecfixes returns a boolean if a field has been set.
func (o *AdvisoryChainGuardPackage) HasSecfixes() bool {
	if o != nil && !IsNil(o.Secfixes) {
		return true
	}

	return false
}

// SetSecfixes gets a reference to the given []AdvisoryChainGuardSecFix and assigns it to the Secfixes field.
func (o *AdvisoryChainGuardPackage) SetSecfixes(v []AdvisoryChainGuardSecFix) {
	o.Secfixes = v
}

func (o AdvisoryChainGuardPackage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryChainGuardPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Secfixes) {
		toSerialize["secfixes"] = o.Secfixes
	}
	return toSerialize, nil
}

type NullableAdvisoryChainGuardPackage struct {
	value *AdvisoryChainGuardPackage
	isSet bool
}

func (v NullableAdvisoryChainGuardPackage) Get() *AdvisoryChainGuardPackage {
	return v.value
}

func (v *NullableAdvisoryChainGuardPackage) Set(val *AdvisoryChainGuardPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryChainGuardPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryChainGuardPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryChainGuardPackage(val *AdvisoryChainGuardPackage) *NullableAdvisoryChainGuardPackage {
	return &NullableAdvisoryChainGuardPackage{value: val, isSet: true}
}

func (v NullableAdvisoryChainGuardPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryChainGuardPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


