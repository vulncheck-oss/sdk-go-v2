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

// checks if the AdvisoryStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryStatus{}

// AdvisoryStatus struct for AdvisoryStatus
type AdvisoryStatus struct {
	ProductID []string `json:"productID,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewAdvisoryStatus instantiates a new AdvisoryStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryStatus() *AdvisoryStatus {
	this := AdvisoryStatus{}
	return &this
}

// NewAdvisoryStatusWithDefaults instantiates a new AdvisoryStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryStatusWithDefaults() *AdvisoryStatus {
	this := AdvisoryStatus{}
	return &this
}

// GetProductID returns the ProductID field value if set, zero value otherwise.
func (o *AdvisoryStatus) GetProductID() []string {
	if o == nil || IsNil(o.ProductID) {
		var ret []string
		return ret
	}
	return o.ProductID
}

// GetProductIDOk returns a tuple with the ProductID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryStatus) GetProductIDOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductID) {
		return nil, false
	}
	return o.ProductID, true
}

// HasProductID returns a boolean if a field has been set.
func (o *AdvisoryStatus) HasProductID() bool {
	if o != nil && !IsNil(o.ProductID) {
		return true
	}

	return false
}

// SetProductID gets a reference to the given []string and assigns it to the ProductID field.
func (o *AdvisoryStatus) SetProductID(v []string) {
	o.ProductID = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AdvisoryStatus) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryStatus) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AdvisoryStatus) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AdvisoryStatus) SetType(v string) {
	o.Type = &v
}

func (o AdvisoryStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductID) {
		toSerialize["productID"] = o.ProductID
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAdvisoryStatus struct {
	value *AdvisoryStatus
	isSet bool
}

func (v NullableAdvisoryStatus) Get() *AdvisoryStatus {
	return v.value
}

func (v *NullableAdvisoryStatus) Set(val *AdvisoryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryStatus(val *AdvisoryStatus) *NullableAdvisoryStatus {
	return &NullableAdvisoryStatus{value: val, isSet: true}
}

func (v NullableAdvisoryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


