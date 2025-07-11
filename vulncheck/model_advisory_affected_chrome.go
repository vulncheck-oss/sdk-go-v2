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

// checks if the AdvisoryAffectedChrome type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryAffectedChrome{}

// AdvisoryAffectedChrome struct for AdvisoryAffectedChrome
type AdvisoryAffectedChrome struct {
	FixedVersion *string `json:"fixed_version,omitempty"`
	Product *string `json:"product,omitempty"`
}

// NewAdvisoryAffectedChrome instantiates a new AdvisoryAffectedChrome object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryAffectedChrome() *AdvisoryAffectedChrome {
	this := AdvisoryAffectedChrome{}
	return &this
}

// NewAdvisoryAffectedChromeWithDefaults instantiates a new AdvisoryAffectedChrome object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryAffectedChromeWithDefaults() *AdvisoryAffectedChrome {
	this := AdvisoryAffectedChrome{}
	return &this
}

// GetFixedVersion returns the FixedVersion field value if set, zero value otherwise.
func (o *AdvisoryAffectedChrome) GetFixedVersion() string {
	if o == nil || IsNil(o.FixedVersion) {
		var ret string
		return ret
	}
	return *o.FixedVersion
}

// GetFixedVersionOk returns a tuple with the FixedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAffectedChrome) GetFixedVersionOk() (*string, bool) {
	if o == nil || IsNil(o.FixedVersion) {
		return nil, false
	}
	return o.FixedVersion, true
}

// HasFixedVersion returns a boolean if a field has been set.
func (o *AdvisoryAffectedChrome) HasFixedVersion() bool {
	if o != nil && !IsNil(o.FixedVersion) {
		return true
	}

	return false
}

// SetFixedVersion gets a reference to the given string and assigns it to the FixedVersion field.
func (o *AdvisoryAffectedChrome) SetFixedVersion(v string) {
	o.FixedVersion = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *AdvisoryAffectedChrome) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAffectedChrome) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *AdvisoryAffectedChrome) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *AdvisoryAffectedChrome) SetProduct(v string) {
	o.Product = &v
}

func (o AdvisoryAffectedChrome) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryAffectedChrome) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FixedVersion) {
		toSerialize["fixed_version"] = o.FixedVersion
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	return toSerialize, nil
}

type NullableAdvisoryAffectedChrome struct {
	value *AdvisoryAffectedChrome
	isSet bool
}

func (v NullableAdvisoryAffectedChrome) Get() *AdvisoryAffectedChrome {
	return v.value
}

func (v *NullableAdvisoryAffectedChrome) Set(val *AdvisoryAffectedChrome) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryAffectedChrome) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryAffectedChrome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryAffectedChrome(val *AdvisoryAffectedChrome) *NullableAdvisoryAffectedChrome {
	return &NullableAdvisoryAffectedChrome{value: val, isSet: true}
}

func (v NullableAdvisoryAffectedChrome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryAffectedChrome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


