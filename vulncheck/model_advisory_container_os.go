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

// checks if the AdvisoryContainerOS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryContainerOS{}

// AdvisoryContainerOS struct for AdvisoryContainerOS
type AdvisoryContainerOS struct {
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Title *string `json:"title,omitempty"`
	Updates []AdvisoryCOSUpdate `json:"updates,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAdvisoryContainerOS instantiates a new AdvisoryContainerOS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryContainerOS() *AdvisoryContainerOS {
	this := AdvisoryContainerOS{}
	return &this
}

// NewAdvisoryContainerOSWithDefaults instantiates a new AdvisoryContainerOS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryContainerOSWithDefaults() *AdvisoryContainerOS {
	this := AdvisoryContainerOS{}
	return &this
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryContainerOS) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryContainerOS) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryContainerOS) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryContainerOS) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryContainerOS) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryContainerOS) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryContainerOS) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryContainerOS) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryContainerOS) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryContainerOS) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryContainerOS) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryContainerOS) SetTitle(v string) {
	o.Title = &v
}

// GetUpdates returns the Updates field value if set, zero value otherwise.
func (o *AdvisoryContainerOS) GetUpdates() []AdvisoryCOSUpdate {
	if o == nil || IsNil(o.Updates) {
		var ret []AdvisoryCOSUpdate
		return ret
	}
	return o.Updates
}

// GetUpdatesOk returns a tuple with the Updates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryContainerOS) GetUpdatesOk() ([]AdvisoryCOSUpdate, bool) {
	if o == nil || IsNil(o.Updates) {
		return nil, false
	}
	return o.Updates, true
}

// HasUpdates returns a boolean if a field has been set.
func (o *AdvisoryContainerOS) HasUpdates() bool {
	if o != nil && !IsNil(o.Updates) {
		return true
	}

	return false
}

// SetUpdates gets a reference to the given []AdvisoryCOSUpdate and assigns it to the Updates field.
func (o *AdvisoryContainerOS) SetUpdates(v []AdvisoryCOSUpdate) {
	o.Updates = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryContainerOS) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryContainerOS) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryContainerOS) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryContainerOS) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisoryContainerOS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryContainerOS) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Updates) {
		toSerialize["updates"] = o.Updates
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAdvisoryContainerOS struct {
	value *AdvisoryContainerOS
	isSet bool
}

func (v NullableAdvisoryContainerOS) Get() *AdvisoryContainerOS {
	return v.value
}

func (v *NullableAdvisoryContainerOS) Set(val *AdvisoryContainerOS) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryContainerOS) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryContainerOS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryContainerOS(val *AdvisoryContainerOS) *NullableAdvisoryContainerOS {
	return &NullableAdvisoryContainerOS{value: val, isSet: true}
}

func (v NullableAdvisoryContainerOS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryContainerOS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


