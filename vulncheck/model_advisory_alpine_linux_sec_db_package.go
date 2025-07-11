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

// checks if the AdvisoryAlpineLinuxSecDBPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryAlpineLinuxSecDBPackage{}

// AdvisoryAlpineLinuxSecDBPackage struct for AdvisoryAlpineLinuxSecDBPackage
type AdvisoryAlpineLinuxSecDBPackage struct {
	PackageName *string `json:"package_name,omitempty"`
	Secfixes []AdvisoryAlpineLinuxSecurityFix `json:"secfixes,omitempty"`
}

// NewAdvisoryAlpineLinuxSecDBPackage instantiates a new AdvisoryAlpineLinuxSecDBPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryAlpineLinuxSecDBPackage() *AdvisoryAlpineLinuxSecDBPackage {
	this := AdvisoryAlpineLinuxSecDBPackage{}
	return &this
}

// NewAdvisoryAlpineLinuxSecDBPackageWithDefaults instantiates a new AdvisoryAlpineLinuxSecDBPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryAlpineLinuxSecDBPackageWithDefaults() *AdvisoryAlpineLinuxSecDBPackage {
	this := AdvisoryAlpineLinuxSecDBPackage{}
	return &this
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *AdvisoryAlpineLinuxSecDBPackage) GetPackageName() string {
	if o == nil || IsNil(o.PackageName) {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAlpineLinuxSecDBPackage) GetPackageNameOk() (*string, bool) {
	if o == nil || IsNil(o.PackageName) {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *AdvisoryAlpineLinuxSecDBPackage) HasPackageName() bool {
	if o != nil && !IsNil(o.PackageName) {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *AdvisoryAlpineLinuxSecDBPackage) SetPackageName(v string) {
	o.PackageName = &v
}

// GetSecfixes returns the Secfixes field value if set, zero value otherwise.
func (o *AdvisoryAlpineLinuxSecDBPackage) GetSecfixes() []AdvisoryAlpineLinuxSecurityFix {
	if o == nil || IsNil(o.Secfixes) {
		var ret []AdvisoryAlpineLinuxSecurityFix
		return ret
	}
	return o.Secfixes
}

// GetSecfixesOk returns a tuple with the Secfixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAlpineLinuxSecDBPackage) GetSecfixesOk() ([]AdvisoryAlpineLinuxSecurityFix, bool) {
	if o == nil || IsNil(o.Secfixes) {
		return nil, false
	}
	return o.Secfixes, true
}

// HasSecfixes returns a boolean if a field has been set.
func (o *AdvisoryAlpineLinuxSecDBPackage) HasSecfixes() bool {
	if o != nil && !IsNil(o.Secfixes) {
		return true
	}

	return false
}

// SetSecfixes gets a reference to the given []AdvisoryAlpineLinuxSecurityFix and assigns it to the Secfixes field.
func (o *AdvisoryAlpineLinuxSecDBPackage) SetSecfixes(v []AdvisoryAlpineLinuxSecurityFix) {
	o.Secfixes = v
}

func (o AdvisoryAlpineLinuxSecDBPackage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryAlpineLinuxSecDBPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PackageName) {
		toSerialize["package_name"] = o.PackageName
	}
	if !IsNil(o.Secfixes) {
		toSerialize["secfixes"] = o.Secfixes
	}
	return toSerialize, nil
}

type NullableAdvisoryAlpineLinuxSecDBPackage struct {
	value *AdvisoryAlpineLinuxSecDBPackage
	isSet bool
}

func (v NullableAdvisoryAlpineLinuxSecDBPackage) Get() *AdvisoryAlpineLinuxSecDBPackage {
	return v.value
}

func (v *NullableAdvisoryAlpineLinuxSecDBPackage) Set(val *AdvisoryAlpineLinuxSecDBPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryAlpineLinuxSecDBPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryAlpineLinuxSecDBPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryAlpineLinuxSecDBPackage(val *AdvisoryAlpineLinuxSecDBPackage) *NullableAdvisoryAlpineLinuxSecDBPackage {
	return &NullableAdvisoryAlpineLinuxSecDBPackage{value: val, isSet: true}
}

func (v NullableAdvisoryAlpineLinuxSecDBPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryAlpineLinuxSecDBPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


