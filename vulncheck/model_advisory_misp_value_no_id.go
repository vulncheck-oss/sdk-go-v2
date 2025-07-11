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

// checks if the AdvisoryMISPValueNoID type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryMISPValueNoID{}

// AdvisoryMISPValueNoID struct for AdvisoryMISPValueNoID
type AdvisoryMISPValueNoID struct {
	Description *string `json:"description,omitempty"`
	Meta *AdvisoryMispMeta `json:"meta,omitempty"`
	Related []AdvisoryMispRelatedItem `json:"related,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewAdvisoryMISPValueNoID instantiates a new AdvisoryMISPValueNoID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryMISPValueNoID() *AdvisoryMISPValueNoID {
	this := AdvisoryMISPValueNoID{}
	return &this
}

// NewAdvisoryMISPValueNoIDWithDefaults instantiates a new AdvisoryMISPValueNoID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryMISPValueNoIDWithDefaults() *AdvisoryMISPValueNoID {
	this := AdvisoryMISPValueNoID{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AdvisoryMISPValueNoID) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMISPValueNoID) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AdvisoryMISPValueNoID) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AdvisoryMISPValueNoID) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AdvisoryMISPValueNoID) GetMeta() AdvisoryMispMeta {
	if o == nil || IsNil(o.Meta) {
		var ret AdvisoryMispMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMISPValueNoID) GetMetaOk() (*AdvisoryMispMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AdvisoryMISPValueNoID) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given AdvisoryMispMeta and assigns it to the Meta field.
func (o *AdvisoryMISPValueNoID) SetMeta(v AdvisoryMispMeta) {
	o.Meta = &v
}

// GetRelated returns the Related field value if set, zero value otherwise.
func (o *AdvisoryMISPValueNoID) GetRelated() []AdvisoryMispRelatedItem {
	if o == nil || IsNil(o.Related) {
		var ret []AdvisoryMispRelatedItem
		return ret
	}
	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMISPValueNoID) GetRelatedOk() ([]AdvisoryMispRelatedItem, bool) {
	if o == nil || IsNil(o.Related) {
		return nil, false
	}
	return o.Related, true
}

// HasRelated returns a boolean if a field has been set.
func (o *AdvisoryMISPValueNoID) HasRelated() bool {
	if o != nil && !IsNil(o.Related) {
		return true
	}

	return false
}

// SetRelated gets a reference to the given []AdvisoryMispRelatedItem and assigns it to the Related field.
func (o *AdvisoryMISPValueNoID) SetRelated(v []AdvisoryMispRelatedItem) {
	o.Related = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AdvisoryMISPValueNoID) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMISPValueNoID) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AdvisoryMISPValueNoID) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AdvisoryMISPValueNoID) SetValue(v string) {
	o.Value = &v
}

func (o AdvisoryMISPValueNoID) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryMISPValueNoID) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Related) {
		toSerialize["related"] = o.Related
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableAdvisoryMISPValueNoID struct {
	value *AdvisoryMISPValueNoID
	isSet bool
}

func (v NullableAdvisoryMISPValueNoID) Get() *AdvisoryMISPValueNoID {
	return v.value
}

func (v *NullableAdvisoryMISPValueNoID) Set(val *AdvisoryMISPValueNoID) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryMISPValueNoID) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryMISPValueNoID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryMISPValueNoID(val *AdvisoryMISPValueNoID) *NullableAdvisoryMISPValueNoID {
	return &NullableAdvisoryMISPValueNoID{value: val, isSet: true}
}

func (v NullableAdvisoryMISPValueNoID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryMISPValueNoID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


