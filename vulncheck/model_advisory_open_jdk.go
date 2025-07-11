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

// checks if the AdvisoryOpenJDK type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryOpenJDK{}

// AdvisoryOpenJDK struct for AdvisoryOpenJDK
type AdvisoryOpenJDK struct {
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	OpenjdkCves []AdvisoryOpenJDKCVE `json:"openjdk_cves,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Title *string `json:"title,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAdvisoryOpenJDK instantiates a new AdvisoryOpenJDK object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryOpenJDK() *AdvisoryOpenJDK {
	this := AdvisoryOpenJDK{}
	return &this
}

// NewAdvisoryOpenJDKWithDefaults instantiates a new AdvisoryOpenJDK object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryOpenJDKWithDefaults() *AdvisoryOpenJDK {
	this := AdvisoryOpenJDK{}
	return &this
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryOpenJDK) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOpenJDK) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryOpenJDK) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryOpenJDK) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryOpenJDK) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOpenJDK) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryOpenJDK) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryOpenJDK) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetOpenjdkCves returns the OpenjdkCves field value if set, zero value otherwise.
func (o *AdvisoryOpenJDK) GetOpenjdkCves() []AdvisoryOpenJDKCVE {
	if o == nil || IsNil(o.OpenjdkCves) {
		var ret []AdvisoryOpenJDKCVE
		return ret
	}
	return o.OpenjdkCves
}

// GetOpenjdkCvesOk returns a tuple with the OpenjdkCves field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOpenJDK) GetOpenjdkCvesOk() ([]AdvisoryOpenJDKCVE, bool) {
	if o == nil || IsNil(o.OpenjdkCves) {
		return nil, false
	}
	return o.OpenjdkCves, true
}

// HasOpenjdkCves returns a boolean if a field has been set.
func (o *AdvisoryOpenJDK) HasOpenjdkCves() bool {
	if o != nil && !IsNil(o.OpenjdkCves) {
		return true
	}

	return false
}

// SetOpenjdkCves gets a reference to the given []AdvisoryOpenJDKCVE and assigns it to the OpenjdkCves field.
func (o *AdvisoryOpenJDK) SetOpenjdkCves(v []AdvisoryOpenJDKCVE) {
	o.OpenjdkCves = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AdvisoryOpenJDK) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOpenJDK) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AdvisoryOpenJDK) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AdvisoryOpenJDK) SetSummary(v string) {
	o.Summary = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryOpenJDK) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOpenJDK) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryOpenJDK) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryOpenJDK) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryOpenJDK) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOpenJDK) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryOpenJDK) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryOpenJDK) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisoryOpenJDK) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryOpenJDK) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.OpenjdkCves) {
		toSerialize["openjdk_cves"] = o.OpenjdkCves
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAdvisoryOpenJDK struct {
	value *AdvisoryOpenJDK
	isSet bool
}

func (v NullableAdvisoryOpenJDK) Get() *AdvisoryOpenJDK {
	return v.value
}

func (v *NullableAdvisoryOpenJDK) Set(val *AdvisoryOpenJDK) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryOpenJDK) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryOpenJDK) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryOpenJDK(val *AdvisoryOpenJDK) *NullableAdvisoryOpenJDK {
	return &NullableAdvisoryOpenJDK{value: val, isSet: true}
}

func (v NullableAdvisoryOpenJDK) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryOpenJDK) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


