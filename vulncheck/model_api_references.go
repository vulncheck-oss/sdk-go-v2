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

// checks if the ApiReferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiReferences{}

// ApiReferences struct for ApiReferences
type ApiReferences struct {
	ReferenceData []ApiReferenceData `json:"reference_data,omitempty"`
}

// NewApiReferences instantiates a new ApiReferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiReferences() *ApiReferences {
	this := ApiReferences{}
	return &this
}

// NewApiReferencesWithDefaults instantiates a new ApiReferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiReferencesWithDefaults() *ApiReferences {
	this := ApiReferences{}
	return &this
}

// GetReferenceData returns the ReferenceData field value if set, zero value otherwise.
func (o *ApiReferences) GetReferenceData() []ApiReferenceData {
	if o == nil || IsNil(o.ReferenceData) {
		var ret []ApiReferenceData
		return ret
	}
	return o.ReferenceData
}

// GetReferenceDataOk returns a tuple with the ReferenceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReferences) GetReferenceDataOk() ([]ApiReferenceData, bool) {
	if o == nil || IsNil(o.ReferenceData) {
		return nil, false
	}
	return o.ReferenceData, true
}

// HasReferenceData returns a boolean if a field has been set.
func (o *ApiReferences) HasReferenceData() bool {
	if o != nil && !IsNil(o.ReferenceData) {
		return true
	}

	return false
}

// SetReferenceData gets a reference to the given []ApiReferenceData and assigns it to the ReferenceData field.
func (o *ApiReferences) SetReferenceData(v []ApiReferenceData) {
	o.ReferenceData = v
}

func (o ApiReferences) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiReferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReferenceData) {
		toSerialize["reference_data"] = o.ReferenceData
	}
	return toSerialize, nil
}

type NullableApiReferences struct {
	value *ApiReferences
	isSet bool
}

func (v NullableApiReferences) Get() *ApiReferences {
	return v.value
}

func (v *NullableApiReferences) Set(val *ApiReferences) {
	v.value = val
	v.isSet = true
}

func (v NullableApiReferences) IsSet() bool {
	return v.isSet
}

func (v *NullableApiReferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiReferences(val *ApiReferences) *NullableApiReferences {
	return &NullableApiReferences{value: val, isSet: true}
}

func (v NullableApiReferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiReferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


