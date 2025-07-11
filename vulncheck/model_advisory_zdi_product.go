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

// checks if the AdvisoryZDIProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryZDIProduct{}

// AdvisoryZDIProduct struct for AdvisoryZDIProduct
type AdvisoryZDIProduct struct {
	Name *string `json:"name,omitempty"`
	Uri *string `json:"uri,omitempty"`
	Vendor *AdvisoryZDIVendor `json:"vendor,omitempty"`
}

// NewAdvisoryZDIProduct instantiates a new AdvisoryZDIProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryZDIProduct() *AdvisoryZDIProduct {
	this := AdvisoryZDIProduct{}
	return &this
}

// NewAdvisoryZDIProductWithDefaults instantiates a new AdvisoryZDIProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryZDIProductWithDefaults() *AdvisoryZDIProduct {
	this := AdvisoryZDIProduct{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdvisoryZDIProduct) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryZDIProduct) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdvisoryZDIProduct) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdvisoryZDIProduct) SetName(v string) {
	o.Name = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *AdvisoryZDIProduct) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryZDIProduct) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *AdvisoryZDIProduct) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *AdvisoryZDIProduct) SetUri(v string) {
	o.Uri = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *AdvisoryZDIProduct) GetVendor() AdvisoryZDIVendor {
	if o == nil || IsNil(o.Vendor) {
		var ret AdvisoryZDIVendor
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryZDIProduct) GetVendorOk() (*AdvisoryZDIVendor, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *AdvisoryZDIProduct) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given AdvisoryZDIVendor and assigns it to the Vendor field.
func (o *AdvisoryZDIProduct) SetVendor(v AdvisoryZDIVendor) {
	o.Vendor = &v
}

func (o AdvisoryZDIProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryZDIProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	return toSerialize, nil
}

type NullableAdvisoryZDIProduct struct {
	value *AdvisoryZDIProduct
	isSet bool
}

func (v NullableAdvisoryZDIProduct) Get() *AdvisoryZDIProduct {
	return v.value
}

func (v *NullableAdvisoryZDIProduct) Set(val *AdvisoryZDIProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryZDIProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryZDIProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryZDIProduct(val *AdvisoryZDIProduct) *NullableAdvisoryZDIProduct {
	return &NullableAdvisoryZDIProduct{value: val, isSet: true}
}

func (v NullableAdvisoryZDIProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryZDIProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


