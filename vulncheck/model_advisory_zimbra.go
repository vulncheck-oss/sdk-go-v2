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

// checks if the AdvisoryZimbra type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryZimbra{}

// AdvisoryZimbra struct for AdvisoryZimbra
type AdvisoryZimbra struct {
	Bugs []int32 `json:"bugs,omitempty"`
	Cve []string `json:"cve,omitempty"`
	Cvss *string `json:"cvss,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Fix *string `json:"fix,omitempty"`
	Rating *string `json:"rating,omitempty"`
	Reporter *string `json:"reporter,omitempty"`
	Summary *string `json:"summary,omitempty"`
}

// NewAdvisoryZimbra instantiates a new AdvisoryZimbra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryZimbra() *AdvisoryZimbra {
	this := AdvisoryZimbra{}
	return &this
}

// NewAdvisoryZimbraWithDefaults instantiates a new AdvisoryZimbra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryZimbraWithDefaults() *AdvisoryZimbra {
	this := AdvisoryZimbra{}
	return &this
}

// GetBugs returns the Bugs field value if set, zero value otherwise.
func (o *AdvisoryZimbra) GetBugs() []int32 {
	if o == nil || IsNil(o.Bugs) {
		var ret []int32
		return ret
	}
	return o.Bugs
}

// GetBugsOk returns a tuple with the Bugs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryZimbra) GetBugsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Bugs) {
		return nil, false
	}
	return o.Bugs, true
}

// HasBugs returns a boolean if a field has been set.
func (o *AdvisoryZimbra) HasBugs() bool {
	if o != nil && !IsNil(o.Bugs) {
		return true
	}

	return false
}

// SetBugs gets a reference to the given []int32 and assigns it to the Bugs field.
func (o *AdvisoryZimbra) SetBugs(v []int32) {
	o.Bugs = v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryZimbra) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryZimbra) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryZimbra) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryZimbra) SetCve(v []string) {
	o.Cve = v
}

// GetCvss returns the Cvss field value if set, zero value otherwise.
func (o *AdvisoryZimbra) GetCvss() string {
	if o == nil || IsNil(o.Cvss) {
		var ret string
		return ret
	}
	return *o.Cvss
}

// GetCvssOk returns a tuple with the Cvss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryZimbra) GetCvssOk() (*string, bool) {
	if o == nil || IsNil(o.Cvss) {
		return nil, false
	}
	return o.Cvss, true
}

// HasCvss returns a boolean if a field has been set.
func (o *AdvisoryZimbra) HasCvss() bool {
	if o != nil && !IsNil(o.Cvss) {
		return true
	}

	return false
}

// SetCvss gets a reference to the given string and assigns it to the Cvss field.
func (o *AdvisoryZimbra) SetCvss(v string) {
	o.Cvss = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryZimbra) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryZimbra) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryZimbra) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryZimbra) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetFix returns the Fix field value if set, zero value otherwise.
func (o *AdvisoryZimbra) GetFix() string {
	if o == nil || IsNil(o.Fix) {
		var ret string
		return ret
	}
	return *o.Fix
}

// GetFixOk returns a tuple with the Fix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryZimbra) GetFixOk() (*string, bool) {
	if o == nil || IsNil(o.Fix) {
		return nil, false
	}
	return o.Fix, true
}

// HasFix returns a boolean if a field has been set.
func (o *AdvisoryZimbra) HasFix() bool {
	if o != nil && !IsNil(o.Fix) {
		return true
	}

	return false
}

// SetFix gets a reference to the given string and assigns it to the Fix field.
func (o *AdvisoryZimbra) SetFix(v string) {
	o.Fix = &v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *AdvisoryZimbra) GetRating() string {
	if o == nil || IsNil(o.Rating) {
		var ret string
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryZimbra) GetRatingOk() (*string, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *AdvisoryZimbra) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given string and assigns it to the Rating field.
func (o *AdvisoryZimbra) SetRating(v string) {
	o.Rating = &v
}

// GetReporter returns the Reporter field value if set, zero value otherwise.
func (o *AdvisoryZimbra) GetReporter() string {
	if o == nil || IsNil(o.Reporter) {
		var ret string
		return ret
	}
	return *o.Reporter
}

// GetReporterOk returns a tuple with the Reporter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryZimbra) GetReporterOk() (*string, bool) {
	if o == nil || IsNil(o.Reporter) {
		return nil, false
	}
	return o.Reporter, true
}

// HasReporter returns a boolean if a field has been set.
func (o *AdvisoryZimbra) HasReporter() bool {
	if o != nil && !IsNil(o.Reporter) {
		return true
	}

	return false
}

// SetReporter gets a reference to the given string and assigns it to the Reporter field.
func (o *AdvisoryZimbra) SetReporter(v string) {
	o.Reporter = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AdvisoryZimbra) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryZimbra) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AdvisoryZimbra) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AdvisoryZimbra) SetSummary(v string) {
	o.Summary = &v
}

func (o AdvisoryZimbra) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryZimbra) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bugs) {
		toSerialize["bugs"] = o.Bugs
	}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.Cvss) {
		toSerialize["cvss"] = o.Cvss
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Fix) {
		toSerialize["fix"] = o.Fix
	}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	if !IsNil(o.Reporter) {
		toSerialize["reporter"] = o.Reporter
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	return toSerialize, nil
}

type NullableAdvisoryZimbra struct {
	value *AdvisoryZimbra
	isSet bool
}

func (v NullableAdvisoryZimbra) Get() *AdvisoryZimbra {
	return v.value
}

func (v *NullableAdvisoryZimbra) Set(val *AdvisoryZimbra) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryZimbra) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryZimbra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryZimbra(val *AdvisoryZimbra) *NullableAdvisoryZimbra {
	return &NullableAdvisoryZimbra{value: val, isSet: true}
}

func (v NullableAdvisoryZimbra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryZimbra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


