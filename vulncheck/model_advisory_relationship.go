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

// checks if the AdvisoryRelationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryRelationship{}

// AdvisoryRelationship struct for AdvisoryRelationship
type AdvisoryRelationship struct {
	ProductReference *string `json:"productReference,omitempty"`
	RelatesToProductReference *string `json:"relatesToProductReference,omitempty"`
	RelationType *string `json:"relationType,omitempty"`
}

// NewAdvisoryRelationship instantiates a new AdvisoryRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryRelationship() *AdvisoryRelationship {
	this := AdvisoryRelationship{}
	return &this
}

// NewAdvisoryRelationshipWithDefaults instantiates a new AdvisoryRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryRelationshipWithDefaults() *AdvisoryRelationship {
	this := AdvisoryRelationship{}
	return &this
}

// GetProductReference returns the ProductReference field value if set, zero value otherwise.
func (o *AdvisoryRelationship) GetProductReference() string {
	if o == nil || IsNil(o.ProductReference) {
		var ret string
		return ret
	}
	return *o.ProductReference
}

// GetProductReferenceOk returns a tuple with the ProductReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRelationship) GetProductReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.ProductReference) {
		return nil, false
	}
	return o.ProductReference, true
}

// HasProductReference returns a boolean if a field has been set.
func (o *AdvisoryRelationship) HasProductReference() bool {
	if o != nil && !IsNil(o.ProductReference) {
		return true
	}

	return false
}

// SetProductReference gets a reference to the given string and assigns it to the ProductReference field.
func (o *AdvisoryRelationship) SetProductReference(v string) {
	o.ProductReference = &v
}

// GetRelatesToProductReference returns the RelatesToProductReference field value if set, zero value otherwise.
func (o *AdvisoryRelationship) GetRelatesToProductReference() string {
	if o == nil || IsNil(o.RelatesToProductReference) {
		var ret string
		return ret
	}
	return *o.RelatesToProductReference
}

// GetRelatesToProductReferenceOk returns a tuple with the RelatesToProductReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRelationship) GetRelatesToProductReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.RelatesToProductReference) {
		return nil, false
	}
	return o.RelatesToProductReference, true
}

// HasRelatesToProductReference returns a boolean if a field has been set.
func (o *AdvisoryRelationship) HasRelatesToProductReference() bool {
	if o != nil && !IsNil(o.RelatesToProductReference) {
		return true
	}

	return false
}

// SetRelatesToProductReference gets a reference to the given string and assigns it to the RelatesToProductReference field.
func (o *AdvisoryRelationship) SetRelatesToProductReference(v string) {
	o.RelatesToProductReference = &v
}

// GetRelationType returns the RelationType field value if set, zero value otherwise.
func (o *AdvisoryRelationship) GetRelationType() string {
	if o == nil || IsNil(o.RelationType) {
		var ret string
		return ret
	}
	return *o.RelationType
}

// GetRelationTypeOk returns a tuple with the RelationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRelationship) GetRelationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RelationType) {
		return nil, false
	}
	return o.RelationType, true
}

// HasRelationType returns a boolean if a field has been set.
func (o *AdvisoryRelationship) HasRelationType() bool {
	if o != nil && !IsNil(o.RelationType) {
		return true
	}

	return false
}

// SetRelationType gets a reference to the given string and assigns it to the RelationType field.
func (o *AdvisoryRelationship) SetRelationType(v string) {
	o.RelationType = &v
}

func (o AdvisoryRelationship) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryRelationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductReference) {
		toSerialize["productReference"] = o.ProductReference
	}
	if !IsNil(o.RelatesToProductReference) {
		toSerialize["relatesToProductReference"] = o.RelatesToProductReference
	}
	if !IsNil(o.RelationType) {
		toSerialize["relationType"] = o.RelationType
	}
	return toSerialize, nil
}

type NullableAdvisoryRelationship struct {
	value *AdvisoryRelationship
	isSet bool
}

func (v NullableAdvisoryRelationship) Get() *AdvisoryRelationship {
	return v.value
}

func (v *NullableAdvisoryRelationship) Set(val *AdvisoryRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryRelationship(val *AdvisoryRelationship) *NullableAdvisoryRelationship {
	return &NullableAdvisoryRelationship{value: val, isSet: true}
}

func (v NullableAdvisoryRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


