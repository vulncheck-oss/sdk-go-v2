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

// checks if the AdvisoryOSVPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryOSVPackage{}

// AdvisoryOSVPackage struct for AdvisoryOSVPackage
type AdvisoryOSVPackage struct {
	Ecosystem *string `json:"ecosystem,omitempty"`
	Name *string `json:"name,omitempty"`
	Purl *string `json:"purl,omitempty"`
}

// NewAdvisoryOSVPackage instantiates a new AdvisoryOSVPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryOSVPackage() *AdvisoryOSVPackage {
	this := AdvisoryOSVPackage{}
	return &this
}

// NewAdvisoryOSVPackageWithDefaults instantiates a new AdvisoryOSVPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryOSVPackageWithDefaults() *AdvisoryOSVPackage {
	this := AdvisoryOSVPackage{}
	return &this
}

// GetEcosystem returns the Ecosystem field value if set, zero value otherwise.
func (o *AdvisoryOSVPackage) GetEcosystem() string {
	if o == nil || IsNil(o.Ecosystem) {
		var ret string
		return ret
	}
	return *o.Ecosystem
}

// GetEcosystemOk returns a tuple with the Ecosystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOSVPackage) GetEcosystemOk() (*string, bool) {
	if o == nil || IsNil(o.Ecosystem) {
		return nil, false
	}
	return o.Ecosystem, true
}

// HasEcosystem returns a boolean if a field has been set.
func (o *AdvisoryOSVPackage) HasEcosystem() bool {
	if o != nil && !IsNil(o.Ecosystem) {
		return true
	}

	return false
}

// SetEcosystem gets a reference to the given string and assigns it to the Ecosystem field.
func (o *AdvisoryOSVPackage) SetEcosystem(v string) {
	o.Ecosystem = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdvisoryOSVPackage) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOSVPackage) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdvisoryOSVPackage) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdvisoryOSVPackage) SetName(v string) {
	o.Name = &v
}

// GetPurl returns the Purl field value if set, zero value otherwise.
func (o *AdvisoryOSVPackage) GetPurl() string {
	if o == nil || IsNil(o.Purl) {
		var ret string
		return ret
	}
	return *o.Purl
}

// GetPurlOk returns a tuple with the Purl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOSVPackage) GetPurlOk() (*string, bool) {
	if o == nil || IsNil(o.Purl) {
		return nil, false
	}
	return o.Purl, true
}

// HasPurl returns a boolean if a field has been set.
func (o *AdvisoryOSVPackage) HasPurl() bool {
	if o != nil && !IsNil(o.Purl) {
		return true
	}

	return false
}

// SetPurl gets a reference to the given string and assigns it to the Purl field.
func (o *AdvisoryOSVPackage) SetPurl(v string) {
	o.Purl = &v
}

func (o AdvisoryOSVPackage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryOSVPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ecosystem) {
		toSerialize["ecosystem"] = o.Ecosystem
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Purl) {
		toSerialize["purl"] = o.Purl
	}
	return toSerialize, nil
}

type NullableAdvisoryOSVPackage struct {
	value *AdvisoryOSVPackage
	isSet bool
}

func (v NullableAdvisoryOSVPackage) Get() *AdvisoryOSVPackage {
	return v.value
}

func (v *NullableAdvisoryOSVPackage) Set(val *AdvisoryOSVPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryOSVPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryOSVPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryOSVPackage(val *AdvisoryOSVPackage) *NullableAdvisoryOSVPackage {
	return &NullableAdvisoryOSVPackage{value: val, isSet: true}
}

func (v NullableAdvisoryOSVPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryOSVPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


