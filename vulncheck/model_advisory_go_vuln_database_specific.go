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

// checks if the AdvisoryGoVulnDatabaseSpecific type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryGoVulnDatabaseSpecific{}

// AdvisoryGoVulnDatabaseSpecific struct for AdvisoryGoVulnDatabaseSpecific
type AdvisoryGoVulnDatabaseSpecific struct {
	Url *string `json:"url,omitempty"`
}

// NewAdvisoryGoVulnDatabaseSpecific instantiates a new AdvisoryGoVulnDatabaseSpecific object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryGoVulnDatabaseSpecific() *AdvisoryGoVulnDatabaseSpecific {
	this := AdvisoryGoVulnDatabaseSpecific{}
	return &this
}

// NewAdvisoryGoVulnDatabaseSpecificWithDefaults instantiates a new AdvisoryGoVulnDatabaseSpecific object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryGoVulnDatabaseSpecificWithDefaults() *AdvisoryGoVulnDatabaseSpecific {
	this := AdvisoryGoVulnDatabaseSpecific{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryGoVulnDatabaseSpecific) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGoVulnDatabaseSpecific) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryGoVulnDatabaseSpecific) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryGoVulnDatabaseSpecific) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisoryGoVulnDatabaseSpecific) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryGoVulnDatabaseSpecific) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAdvisoryGoVulnDatabaseSpecific struct {
	value *AdvisoryGoVulnDatabaseSpecific
	isSet bool
}

func (v NullableAdvisoryGoVulnDatabaseSpecific) Get() *AdvisoryGoVulnDatabaseSpecific {
	return v.value
}

func (v *NullableAdvisoryGoVulnDatabaseSpecific) Set(val *AdvisoryGoVulnDatabaseSpecific) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryGoVulnDatabaseSpecific) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryGoVulnDatabaseSpecific) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryGoVulnDatabaseSpecific(val *AdvisoryGoVulnDatabaseSpecific) *NullableAdvisoryGoVulnDatabaseSpecific {
	return &NullableAdvisoryGoVulnDatabaseSpecific{value: val, isSet: true}
}

func (v NullableAdvisoryGoVulnDatabaseSpecific) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryGoVulnDatabaseSpecific) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


