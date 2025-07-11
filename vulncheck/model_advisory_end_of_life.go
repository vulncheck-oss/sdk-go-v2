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

// checks if the AdvisoryEndOfLife type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryEndOfLife{}

// AdvisoryEndOfLife struct for AdvisoryEndOfLife
type AdvisoryEndOfLife struct {
	Cve []string `json:"cve,omitempty"`
	Cycles []AdvisoryCycle `json:"cycles,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Name *string `json:"name,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAdvisoryEndOfLife instantiates a new AdvisoryEndOfLife object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryEndOfLife() *AdvisoryEndOfLife {
	this := AdvisoryEndOfLife{}
	return &this
}

// NewAdvisoryEndOfLifeWithDefaults instantiates a new AdvisoryEndOfLife object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryEndOfLifeWithDefaults() *AdvisoryEndOfLife {
	this := AdvisoryEndOfLife{}
	return &this
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryEndOfLife) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryEndOfLife) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryEndOfLife) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryEndOfLife) SetCve(v []string) {
	o.Cve = v
}

// GetCycles returns the Cycles field value if set, zero value otherwise.
func (o *AdvisoryEndOfLife) GetCycles() []AdvisoryCycle {
	if o == nil || IsNil(o.Cycles) {
		var ret []AdvisoryCycle
		return ret
	}
	return o.Cycles
}

// GetCyclesOk returns a tuple with the Cycles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryEndOfLife) GetCyclesOk() ([]AdvisoryCycle, bool) {
	if o == nil || IsNil(o.Cycles) {
		return nil, false
	}
	return o.Cycles, true
}

// HasCycles returns a boolean if a field has been set.
func (o *AdvisoryEndOfLife) HasCycles() bool {
	if o != nil && !IsNil(o.Cycles) {
		return true
	}

	return false
}

// SetCycles gets a reference to the given []AdvisoryCycle and assigns it to the Cycles field.
func (o *AdvisoryEndOfLife) SetCycles(v []AdvisoryCycle) {
	o.Cycles = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryEndOfLife) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryEndOfLife) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryEndOfLife) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryEndOfLife) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdvisoryEndOfLife) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryEndOfLife) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdvisoryEndOfLife) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdvisoryEndOfLife) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryEndOfLife) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryEndOfLife) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryEndOfLife) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryEndOfLife) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisoryEndOfLife) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryEndOfLife) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.Cycles) {
		toSerialize["cycles"] = o.Cycles
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAdvisoryEndOfLife struct {
	value *AdvisoryEndOfLife
	isSet bool
}

func (v NullableAdvisoryEndOfLife) Get() *AdvisoryEndOfLife {
	return v.value
}

func (v *NullableAdvisoryEndOfLife) Set(val *AdvisoryEndOfLife) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryEndOfLife) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryEndOfLife) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryEndOfLife(val *AdvisoryEndOfLife) *NullableAdvisoryEndOfLife {
	return &NullableAdvisoryEndOfLife{value: val, isSet: true}
}

func (v NullableAdvisoryEndOfLife) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryEndOfLife) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


