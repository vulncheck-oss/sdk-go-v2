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

// checks if the AdvisoryChrome type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryChrome{}

// AdvisoryChrome struct for AdvisoryChrome
type AdvisoryChrome struct {
	Affected []AdvisoryAffectedChrome `json:"affected,omitempty"`
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Title *string `json:"title,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAdvisoryChrome instantiates a new AdvisoryChrome object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryChrome() *AdvisoryChrome {
	this := AdvisoryChrome{}
	return &this
}

// NewAdvisoryChromeWithDefaults instantiates a new AdvisoryChrome object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryChromeWithDefaults() *AdvisoryChrome {
	this := AdvisoryChrome{}
	return &this
}

// GetAffected returns the Affected field value if set, zero value otherwise.
func (o *AdvisoryChrome) GetAffected() []AdvisoryAffectedChrome {
	if o == nil || IsNil(o.Affected) {
		var ret []AdvisoryAffectedChrome
		return ret
	}
	return o.Affected
}

// GetAffectedOk returns a tuple with the Affected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryChrome) GetAffectedOk() ([]AdvisoryAffectedChrome, bool) {
	if o == nil || IsNil(o.Affected) {
		return nil, false
	}
	return o.Affected, true
}

// HasAffected returns a boolean if a field has been set.
func (o *AdvisoryChrome) HasAffected() bool {
	if o != nil && !IsNil(o.Affected) {
		return true
	}

	return false
}

// SetAffected gets a reference to the given []AdvisoryAffectedChrome and assigns it to the Affected field.
func (o *AdvisoryChrome) SetAffected(v []AdvisoryAffectedChrome) {
	o.Affected = v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryChrome) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryChrome) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryChrome) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryChrome) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryChrome) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryChrome) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryChrome) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryChrome) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryChrome) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryChrome) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryChrome) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryChrome) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryChrome) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryChrome) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryChrome) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryChrome) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisoryChrome) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryChrome) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Affected) {
		toSerialize["affected"] = o.Affected
	}
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

type NullableAdvisoryChrome struct {
	value *AdvisoryChrome
	isSet bool
}

func (v NullableAdvisoryChrome) Get() *AdvisoryChrome {
	return v.value
}

func (v *NullableAdvisoryChrome) Set(val *AdvisoryChrome) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryChrome) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryChrome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryChrome(val *AdvisoryChrome) *NullableAdvisoryChrome {
	return &NullableAdvisoryChrome{value: val, isSet: true}
}

func (v NullableAdvisoryChrome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryChrome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


