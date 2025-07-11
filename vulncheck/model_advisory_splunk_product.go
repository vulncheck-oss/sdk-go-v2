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

// checks if the AdvisorySplunkProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisorySplunkProduct{}

// AdvisorySplunkProduct struct for AdvisorySplunkProduct
type AdvisorySplunkProduct struct {
	AffectedVersion *string `json:"affected_version,omitempty"`
	Component *string `json:"component,omitempty"`
	FixedVersion *string `json:"fixed_version,omitempty"`
	Product *string `json:"product,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewAdvisorySplunkProduct instantiates a new AdvisorySplunkProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisorySplunkProduct() *AdvisorySplunkProduct {
	this := AdvisorySplunkProduct{}
	return &this
}

// NewAdvisorySplunkProductWithDefaults instantiates a new AdvisorySplunkProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisorySplunkProductWithDefaults() *AdvisorySplunkProduct {
	this := AdvisorySplunkProduct{}
	return &this
}

// GetAffectedVersion returns the AffectedVersion field value if set, zero value otherwise.
func (o *AdvisorySplunkProduct) GetAffectedVersion() string {
	if o == nil || IsNil(o.AffectedVersion) {
		var ret string
		return ret
	}
	return *o.AffectedVersion
}

// GetAffectedVersionOk returns a tuple with the AffectedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySplunkProduct) GetAffectedVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AffectedVersion) {
		return nil, false
	}
	return o.AffectedVersion, true
}

// HasAffectedVersion returns a boolean if a field has been set.
func (o *AdvisorySplunkProduct) HasAffectedVersion() bool {
	if o != nil && !IsNil(o.AffectedVersion) {
		return true
	}

	return false
}

// SetAffectedVersion gets a reference to the given string and assigns it to the AffectedVersion field.
func (o *AdvisorySplunkProduct) SetAffectedVersion(v string) {
	o.AffectedVersion = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AdvisorySplunkProduct) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySplunkProduct) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AdvisorySplunkProduct) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AdvisorySplunkProduct) SetComponent(v string) {
	o.Component = &v
}

// GetFixedVersion returns the FixedVersion field value if set, zero value otherwise.
func (o *AdvisorySplunkProduct) GetFixedVersion() string {
	if o == nil || IsNil(o.FixedVersion) {
		var ret string
		return ret
	}
	return *o.FixedVersion
}

// GetFixedVersionOk returns a tuple with the FixedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySplunkProduct) GetFixedVersionOk() (*string, bool) {
	if o == nil || IsNil(o.FixedVersion) {
		return nil, false
	}
	return o.FixedVersion, true
}

// HasFixedVersion returns a boolean if a field has been set.
func (o *AdvisorySplunkProduct) HasFixedVersion() bool {
	if o != nil && !IsNil(o.FixedVersion) {
		return true
	}

	return false
}

// SetFixedVersion gets a reference to the given string and assigns it to the FixedVersion field.
func (o *AdvisorySplunkProduct) SetFixedVersion(v string) {
	o.FixedVersion = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *AdvisorySplunkProduct) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySplunkProduct) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *AdvisorySplunkProduct) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *AdvisorySplunkProduct) SetProduct(v string) {
	o.Product = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AdvisorySplunkProduct) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySplunkProduct) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AdvisorySplunkProduct) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AdvisorySplunkProduct) SetVersion(v string) {
	o.Version = &v
}

func (o AdvisorySplunkProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisorySplunkProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AffectedVersion) {
		toSerialize["affected_version"] = o.AffectedVersion
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.FixedVersion) {
		toSerialize["fixed_version"] = o.FixedVersion
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableAdvisorySplunkProduct struct {
	value *AdvisorySplunkProduct
	isSet bool
}

func (v NullableAdvisorySplunkProduct) Get() *AdvisorySplunkProduct {
	return v.value
}

func (v *NullableAdvisorySplunkProduct) Set(val *AdvisorySplunkProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisorySplunkProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisorySplunkProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisorySplunkProduct(val *AdvisorySplunkProduct) *NullableAdvisorySplunkProduct {
	return &NullableAdvisorySplunkProduct{value: val, isSet: true}
}

func (v NullableAdvisorySplunkProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisorySplunkProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


