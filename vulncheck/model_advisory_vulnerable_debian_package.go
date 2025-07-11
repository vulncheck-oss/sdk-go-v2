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

// checks if the AdvisoryVulnerableDebianPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryVulnerableDebianPackage{}

// AdvisoryVulnerableDebianPackage struct for AdvisoryVulnerableDebianPackage
type AdvisoryVulnerableDebianPackage struct {
	AssociatedCves []AdvisoryDebianCVE `json:"associated_cves,omitempty"`
	Cve []string `json:"cve,omitempty"`
	PackageName *string `json:"package_name,omitempty"`
}

// NewAdvisoryVulnerableDebianPackage instantiates a new AdvisoryVulnerableDebianPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryVulnerableDebianPackage() *AdvisoryVulnerableDebianPackage {
	this := AdvisoryVulnerableDebianPackage{}
	return &this
}

// NewAdvisoryVulnerableDebianPackageWithDefaults instantiates a new AdvisoryVulnerableDebianPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryVulnerableDebianPackageWithDefaults() *AdvisoryVulnerableDebianPackage {
	this := AdvisoryVulnerableDebianPackage{}
	return &this
}

// GetAssociatedCves returns the AssociatedCves field value if set, zero value otherwise.
func (o *AdvisoryVulnerableDebianPackage) GetAssociatedCves() []AdvisoryDebianCVE {
	if o == nil || IsNil(o.AssociatedCves) {
		var ret []AdvisoryDebianCVE
		return ret
	}
	return o.AssociatedCves
}

// GetAssociatedCvesOk returns a tuple with the AssociatedCves field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryVulnerableDebianPackage) GetAssociatedCvesOk() ([]AdvisoryDebianCVE, bool) {
	if o == nil || IsNil(o.AssociatedCves) {
		return nil, false
	}
	return o.AssociatedCves, true
}

// HasAssociatedCves returns a boolean if a field has been set.
func (o *AdvisoryVulnerableDebianPackage) HasAssociatedCves() bool {
	if o != nil && !IsNil(o.AssociatedCves) {
		return true
	}

	return false
}

// SetAssociatedCves gets a reference to the given []AdvisoryDebianCVE and assigns it to the AssociatedCves field.
func (o *AdvisoryVulnerableDebianPackage) SetAssociatedCves(v []AdvisoryDebianCVE) {
	o.AssociatedCves = v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryVulnerableDebianPackage) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryVulnerableDebianPackage) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryVulnerableDebianPackage) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryVulnerableDebianPackage) SetCve(v []string) {
	o.Cve = v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *AdvisoryVulnerableDebianPackage) GetPackageName() string {
	if o == nil || IsNil(o.PackageName) {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryVulnerableDebianPackage) GetPackageNameOk() (*string, bool) {
	if o == nil || IsNil(o.PackageName) {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *AdvisoryVulnerableDebianPackage) HasPackageName() bool {
	if o != nil && !IsNil(o.PackageName) {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *AdvisoryVulnerableDebianPackage) SetPackageName(v string) {
	o.PackageName = &v
}

func (o AdvisoryVulnerableDebianPackage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryVulnerableDebianPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociatedCves) {
		toSerialize["associated_cves"] = o.AssociatedCves
	}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.PackageName) {
		toSerialize["package_name"] = o.PackageName
	}
	return toSerialize, nil
}

type NullableAdvisoryVulnerableDebianPackage struct {
	value *AdvisoryVulnerableDebianPackage
	isSet bool
}

func (v NullableAdvisoryVulnerableDebianPackage) Get() *AdvisoryVulnerableDebianPackage {
	return v.value
}

func (v *NullableAdvisoryVulnerableDebianPackage) Set(val *AdvisoryVulnerableDebianPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryVulnerableDebianPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryVulnerableDebianPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryVulnerableDebianPackage(val *AdvisoryVulnerableDebianPackage) *NullableAdvisoryVulnerableDebianPackage {
	return &NullableAdvisoryVulnerableDebianPackage{value: val, isSet: true}
}

func (v NullableAdvisoryVulnerableDebianPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryVulnerableDebianPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


