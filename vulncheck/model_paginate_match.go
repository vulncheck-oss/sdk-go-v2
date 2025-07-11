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

// checks if the PaginateMatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginateMatch{}

// PaginateMatch struct for PaginateMatch
type PaginateMatch struct {
	Field *string `json:"field,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewPaginateMatch instantiates a new PaginateMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginateMatch() *PaginateMatch {
	this := PaginateMatch{}
	return &this
}

// NewPaginateMatchWithDefaults instantiates a new PaginateMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginateMatchWithDefaults() *PaginateMatch {
	this := PaginateMatch{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *PaginateMatch) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginateMatch) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *PaginateMatch) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *PaginateMatch) SetField(v string) {
	o.Field = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PaginateMatch) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginateMatch) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PaginateMatch) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PaginateMatch) SetValue(v string) {
	o.Value = &v
}

func (o PaginateMatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginateMatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullablePaginateMatch struct {
	value *PaginateMatch
	isSet bool
}

func (v NullablePaginateMatch) Get() *PaginateMatch {
	return v.value
}

func (v *NullablePaginateMatch) Set(val *PaginateMatch) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginateMatch) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginateMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginateMatch(val *PaginateMatch) *NullablePaginateMatch {
	return &NullablePaginateMatch{value: val, isSet: true}
}

func (v NullablePaginateMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginateMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


