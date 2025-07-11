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

// checks if the AdvisoryIssued type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryIssued{}

// AdvisoryIssued struct for AdvisoryIssued
type AdvisoryIssued struct {
	Date *string `json:"date,omitempty"`
}

// NewAdvisoryIssued instantiates a new AdvisoryIssued object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryIssued() *AdvisoryIssued {
	this := AdvisoryIssued{}
	return &this
}

// NewAdvisoryIssuedWithDefaults instantiates a new AdvisoryIssued object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryIssuedWithDefaults() *AdvisoryIssued {
	this := AdvisoryIssued{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *AdvisoryIssued) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryIssued) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *AdvisoryIssued) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *AdvisoryIssued) SetDate(v string) {
	o.Date = &v
}

func (o AdvisoryIssued) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryIssued) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	return toSerialize, nil
}

type NullableAdvisoryIssued struct {
	value *AdvisoryIssued
	isSet bool
}

func (v NullableAdvisoryIssued) Get() *AdvisoryIssued {
	return v.value
}

func (v *NullableAdvisoryIssued) Set(val *AdvisoryIssued) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryIssued) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryIssued) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryIssued(val *AdvisoryIssued) *NullableAdvisoryIssued {
	return &NullableAdvisoryIssued{value: val, isSet: true}
}

func (v NullableAdvisoryIssued) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryIssued) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


