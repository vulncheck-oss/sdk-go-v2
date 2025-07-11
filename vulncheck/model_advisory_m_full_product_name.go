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

// checks if the AdvisoryMFullProductName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryMFullProductName{}

// AdvisoryMFullProductName struct for AdvisoryMFullProductName
type AdvisoryMFullProductName struct {
	CPE *string `json:"CPE,omitempty"`
	ProductID *string `json:"ProductID,omitempty"`
	Value *string `json:"Value,omitempty"`
}

// NewAdvisoryMFullProductName instantiates a new AdvisoryMFullProductName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryMFullProductName() *AdvisoryMFullProductName {
	this := AdvisoryMFullProductName{}
	return &this
}

// NewAdvisoryMFullProductNameWithDefaults instantiates a new AdvisoryMFullProductName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryMFullProductNameWithDefaults() *AdvisoryMFullProductName {
	this := AdvisoryMFullProductName{}
	return &this
}

// GetCPE returns the CPE field value if set, zero value otherwise.
func (o *AdvisoryMFullProductName) GetCPE() string {
	if o == nil || IsNil(o.CPE) {
		var ret string
		return ret
	}
	return *o.CPE
}

// GetCPEOk returns a tuple with the CPE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMFullProductName) GetCPEOk() (*string, bool) {
	if o == nil || IsNil(o.CPE) {
		return nil, false
	}
	return o.CPE, true
}

// HasCPE returns a boolean if a field has been set.
func (o *AdvisoryMFullProductName) HasCPE() bool {
	if o != nil && !IsNil(o.CPE) {
		return true
	}

	return false
}

// SetCPE gets a reference to the given string and assigns it to the CPE field.
func (o *AdvisoryMFullProductName) SetCPE(v string) {
	o.CPE = &v
}

// GetProductID returns the ProductID field value if set, zero value otherwise.
func (o *AdvisoryMFullProductName) GetProductID() string {
	if o == nil || IsNil(o.ProductID) {
		var ret string
		return ret
	}
	return *o.ProductID
}

// GetProductIDOk returns a tuple with the ProductID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMFullProductName) GetProductIDOk() (*string, bool) {
	if o == nil || IsNil(o.ProductID) {
		return nil, false
	}
	return o.ProductID, true
}

// HasProductID returns a boolean if a field has been set.
func (o *AdvisoryMFullProductName) HasProductID() bool {
	if o != nil && !IsNil(o.ProductID) {
		return true
	}

	return false
}

// SetProductID gets a reference to the given string and assigns it to the ProductID field.
func (o *AdvisoryMFullProductName) SetProductID(v string) {
	o.ProductID = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AdvisoryMFullProductName) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMFullProductName) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AdvisoryMFullProductName) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AdvisoryMFullProductName) SetValue(v string) {
	o.Value = &v
}

func (o AdvisoryMFullProductName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryMFullProductName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CPE) {
		toSerialize["CPE"] = o.CPE
	}
	if !IsNil(o.ProductID) {
		toSerialize["ProductID"] = o.ProductID
	}
	if !IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}
	return toSerialize, nil
}

type NullableAdvisoryMFullProductName struct {
	value *AdvisoryMFullProductName
	isSet bool
}

func (v NullableAdvisoryMFullProductName) Get() *AdvisoryMFullProductName {
	return v.value
}

func (v *NullableAdvisoryMFullProductName) Set(val *AdvisoryMFullProductName) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryMFullProductName) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryMFullProductName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryMFullProductName(val *AdvisoryMFullProductName) *NullableAdvisoryMFullProductName {
	return &NullableAdvisoryMFullProductName{value: val, isSet: true}
}

func (v NullableAdvisoryMFullProductName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryMFullProductName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


