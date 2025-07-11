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

// checks if the AdvisoryCweAcceptanceLevel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryCweAcceptanceLevel{}

// AdvisoryCweAcceptanceLevel struct for AdvisoryCweAcceptanceLevel
type AdvisoryCweAcceptanceLevel struct {
	Description *string `json:"description,omitempty"`
	LastModified *string `json:"lastModified,omitempty"`
}

// NewAdvisoryCweAcceptanceLevel instantiates a new AdvisoryCweAcceptanceLevel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryCweAcceptanceLevel() *AdvisoryCweAcceptanceLevel {
	this := AdvisoryCweAcceptanceLevel{}
	return &this
}

// NewAdvisoryCweAcceptanceLevelWithDefaults instantiates a new AdvisoryCweAcceptanceLevel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryCweAcceptanceLevelWithDefaults() *AdvisoryCweAcceptanceLevel {
	this := AdvisoryCweAcceptanceLevel{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AdvisoryCweAcceptanceLevel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCweAcceptanceLevel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AdvisoryCweAcceptanceLevel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AdvisoryCweAcceptanceLevel) SetDescription(v string) {
	o.Description = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *AdvisoryCweAcceptanceLevel) GetLastModified() string {
	if o == nil || IsNil(o.LastModified) {
		var ret string
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCweAcceptanceLevel) GetLastModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *AdvisoryCweAcceptanceLevel) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given string and assigns it to the LastModified field.
func (o *AdvisoryCweAcceptanceLevel) SetLastModified(v string) {
	o.LastModified = &v
}

func (o AdvisoryCweAcceptanceLevel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryCweAcceptanceLevel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableAdvisoryCweAcceptanceLevel struct {
	value *AdvisoryCweAcceptanceLevel
	isSet bool
}

func (v NullableAdvisoryCweAcceptanceLevel) Get() *AdvisoryCweAcceptanceLevel {
	return v.value
}

func (v *NullableAdvisoryCweAcceptanceLevel) Set(val *AdvisoryCweAcceptanceLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryCweAcceptanceLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryCweAcceptanceLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryCweAcceptanceLevel(val *AdvisoryCweAcceptanceLevel) *NullableAdvisoryCweAcceptanceLevel {
	return &NullableAdvisoryCweAcceptanceLevel{value: val, isSet: true}
}

func (v NullableAdvisoryCweAcceptanceLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryCweAcceptanceLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


