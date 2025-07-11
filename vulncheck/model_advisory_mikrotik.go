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

// checks if the AdvisoryMikrotik type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryMikrotik{}

// AdvisoryMikrotik struct for AdvisoryMikrotik
type AdvisoryMikrotik struct {
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Title *string `json:"title,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAdvisoryMikrotik instantiates a new AdvisoryMikrotik object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryMikrotik() *AdvisoryMikrotik {
	this := AdvisoryMikrotik{}
	return &this
}

// NewAdvisoryMikrotikWithDefaults instantiates a new AdvisoryMikrotik object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryMikrotikWithDefaults() *AdvisoryMikrotik {
	this := AdvisoryMikrotik{}
	return &this
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryMikrotik) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMikrotik) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryMikrotik) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryMikrotik) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryMikrotik) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMikrotik) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryMikrotik) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryMikrotik) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryMikrotik) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMikrotik) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryMikrotik) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryMikrotik) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryMikrotik) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMikrotik) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryMikrotik) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryMikrotik) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisoryMikrotik) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryMikrotik) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAdvisoryMikrotik struct {
	value *AdvisoryMikrotik
	isSet bool
}

func (v NullableAdvisoryMikrotik) Get() *AdvisoryMikrotik {
	return v.value
}

func (v *NullableAdvisoryMikrotik) Set(val *AdvisoryMikrotik) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryMikrotik) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryMikrotik) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryMikrotik(val *AdvisoryMikrotik) *NullableAdvisoryMikrotik {
	return &NullableAdvisoryMikrotik{value: val, isSet: true}
}

func (v NullableAdvisoryMikrotik) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryMikrotik) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


