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

// checks if the AdvisoryVendorNameForThreatActor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryVendorNameForThreatActor{}

// AdvisoryVendorNameForThreatActor struct for AdvisoryVendorNameForThreatActor
type AdvisoryVendorNameForThreatActor struct {
	ThreatActorName *string `json:"threat_actor_name,omitempty"`
	Url *string `json:"url,omitempty"`
	VendorName *string `json:"vendor_name,omitempty"`
}

// NewAdvisoryVendorNameForThreatActor instantiates a new AdvisoryVendorNameForThreatActor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryVendorNameForThreatActor() *AdvisoryVendorNameForThreatActor {
	this := AdvisoryVendorNameForThreatActor{}
	return &this
}

// NewAdvisoryVendorNameForThreatActorWithDefaults instantiates a new AdvisoryVendorNameForThreatActor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryVendorNameForThreatActorWithDefaults() *AdvisoryVendorNameForThreatActor {
	this := AdvisoryVendorNameForThreatActor{}
	return &this
}

// GetThreatActorName returns the ThreatActorName field value if set, zero value otherwise.
func (o *AdvisoryVendorNameForThreatActor) GetThreatActorName() string {
	if o == nil || IsNil(o.ThreatActorName) {
		var ret string
		return ret
	}
	return *o.ThreatActorName
}

// GetThreatActorNameOk returns a tuple with the ThreatActorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryVendorNameForThreatActor) GetThreatActorNameOk() (*string, bool) {
	if o == nil || IsNil(o.ThreatActorName) {
		return nil, false
	}
	return o.ThreatActorName, true
}

// HasThreatActorName returns a boolean if a field has been set.
func (o *AdvisoryVendorNameForThreatActor) HasThreatActorName() bool {
	if o != nil && !IsNil(o.ThreatActorName) {
		return true
	}

	return false
}

// SetThreatActorName gets a reference to the given string and assigns it to the ThreatActorName field.
func (o *AdvisoryVendorNameForThreatActor) SetThreatActorName(v string) {
	o.ThreatActorName = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryVendorNameForThreatActor) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryVendorNameForThreatActor) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryVendorNameForThreatActor) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryVendorNameForThreatActor) SetUrl(v string) {
	o.Url = &v
}

// GetVendorName returns the VendorName field value if set, zero value otherwise.
func (o *AdvisoryVendorNameForThreatActor) GetVendorName() string {
	if o == nil || IsNil(o.VendorName) {
		var ret string
		return ret
	}
	return *o.VendorName
}

// GetVendorNameOk returns a tuple with the VendorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryVendorNameForThreatActor) GetVendorNameOk() (*string, bool) {
	if o == nil || IsNil(o.VendorName) {
		return nil, false
	}
	return o.VendorName, true
}

// HasVendorName returns a boolean if a field has been set.
func (o *AdvisoryVendorNameForThreatActor) HasVendorName() bool {
	if o != nil && !IsNil(o.VendorName) {
		return true
	}

	return false
}

// SetVendorName gets a reference to the given string and assigns it to the VendorName field.
func (o *AdvisoryVendorNameForThreatActor) SetVendorName(v string) {
	o.VendorName = &v
}

func (o AdvisoryVendorNameForThreatActor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryVendorNameForThreatActor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ThreatActorName) {
		toSerialize["threat_actor_name"] = o.ThreatActorName
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.VendorName) {
		toSerialize["vendor_name"] = o.VendorName
	}
	return toSerialize, nil
}

type NullableAdvisoryVendorNameForThreatActor struct {
	value *AdvisoryVendorNameForThreatActor
	isSet bool
}

func (v NullableAdvisoryVendorNameForThreatActor) Get() *AdvisoryVendorNameForThreatActor {
	return v.value
}

func (v *NullableAdvisoryVendorNameForThreatActor) Set(val *AdvisoryVendorNameForThreatActor) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryVendorNameForThreatActor) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryVendorNameForThreatActor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryVendorNameForThreatActor(val *AdvisoryVendorNameForThreatActor) *NullableAdvisoryVendorNameForThreatActor {
	return &NullableAdvisoryVendorNameForThreatActor{value: val, isSet: true}
}

func (v NullableAdvisoryVendorNameForThreatActor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryVendorNameForThreatActor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


