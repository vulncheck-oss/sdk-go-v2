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

// checks if the AdvisoryWordfence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryWordfence{}

// AdvisoryWordfence struct for AdvisoryWordfence
type AdvisoryWordfence struct {
	Affected []string `json:"affected,omitempty"`
	Cve []string `json:"cve,omitempty"`
	CvssScore *string `json:"cvss_score,omitempty"`
	CvssVector *string `json:"cvss_vector,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Fixed []string `json:"fixed,omitempty"`
	References []string `json:"references,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAdvisoryWordfence instantiates a new AdvisoryWordfence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryWordfence() *AdvisoryWordfence {
	this := AdvisoryWordfence{}
	return &this
}

// NewAdvisoryWordfenceWithDefaults instantiates a new AdvisoryWordfence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryWordfenceWithDefaults() *AdvisoryWordfence {
	this := AdvisoryWordfence{}
	return &this
}

// GetAffected returns the Affected field value if set, zero value otherwise.
func (o *AdvisoryWordfence) GetAffected() []string {
	if o == nil || IsNil(o.Affected) {
		var ret []string
		return ret
	}
	return o.Affected
}

// GetAffectedOk returns a tuple with the Affected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryWordfence) GetAffectedOk() ([]string, bool) {
	if o == nil || IsNil(o.Affected) {
		return nil, false
	}
	return o.Affected, true
}

// HasAffected returns a boolean if a field has been set.
func (o *AdvisoryWordfence) HasAffected() bool {
	if o != nil && !IsNil(o.Affected) {
		return true
	}

	return false
}

// SetAffected gets a reference to the given []string and assigns it to the Affected field.
func (o *AdvisoryWordfence) SetAffected(v []string) {
	o.Affected = v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryWordfence) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryWordfence) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryWordfence) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryWordfence) SetCve(v []string) {
	o.Cve = v
}

// GetCvssScore returns the CvssScore field value if set, zero value otherwise.
func (o *AdvisoryWordfence) GetCvssScore() string {
	if o == nil || IsNil(o.CvssScore) {
		var ret string
		return ret
	}
	return *o.CvssScore
}

// GetCvssScoreOk returns a tuple with the CvssScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryWordfence) GetCvssScoreOk() (*string, bool) {
	if o == nil || IsNil(o.CvssScore) {
		return nil, false
	}
	return o.CvssScore, true
}

// HasCvssScore returns a boolean if a field has been set.
func (o *AdvisoryWordfence) HasCvssScore() bool {
	if o != nil && !IsNil(o.CvssScore) {
		return true
	}

	return false
}

// SetCvssScore gets a reference to the given string and assigns it to the CvssScore field.
func (o *AdvisoryWordfence) SetCvssScore(v string) {
	o.CvssScore = &v
}

// GetCvssVector returns the CvssVector field value if set, zero value otherwise.
func (o *AdvisoryWordfence) GetCvssVector() string {
	if o == nil || IsNil(o.CvssVector) {
		var ret string
		return ret
	}
	return *o.CvssVector
}

// GetCvssVectorOk returns a tuple with the CvssVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryWordfence) GetCvssVectorOk() (*string, bool) {
	if o == nil || IsNil(o.CvssVector) {
		return nil, false
	}
	return o.CvssVector, true
}

// HasCvssVector returns a boolean if a field has been set.
func (o *AdvisoryWordfence) HasCvssVector() bool {
	if o != nil && !IsNil(o.CvssVector) {
		return true
	}

	return false
}

// SetCvssVector gets a reference to the given string and assigns it to the CvssVector field.
func (o *AdvisoryWordfence) SetCvssVector(v string) {
	o.CvssVector = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryWordfence) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryWordfence) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryWordfence) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryWordfence) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetFixed returns the Fixed field value if set, zero value otherwise.
func (o *AdvisoryWordfence) GetFixed() []string {
	if o == nil || IsNil(o.Fixed) {
		var ret []string
		return ret
	}
	return o.Fixed
}

// GetFixedOk returns a tuple with the Fixed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryWordfence) GetFixedOk() ([]string, bool) {
	if o == nil || IsNil(o.Fixed) {
		return nil, false
	}
	return o.Fixed, true
}

// HasFixed returns a boolean if a field has been set.
func (o *AdvisoryWordfence) HasFixed() bool {
	if o != nil && !IsNil(o.Fixed) {
		return true
	}

	return false
}

// SetFixed gets a reference to the given []string and assigns it to the Fixed field.
func (o *AdvisoryWordfence) SetFixed(v []string) {
	o.Fixed = v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *AdvisoryWordfence) GetReferences() []string {
	if o == nil || IsNil(o.References) {
		var ret []string
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryWordfence) GetReferencesOk() ([]string, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *AdvisoryWordfence) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []string and assigns it to the References field.
func (o *AdvisoryWordfence) SetReferences(v []string) {
	o.References = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AdvisoryWordfence) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryWordfence) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AdvisoryWordfence) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AdvisoryWordfence) SetSummary(v string) {
	o.Summary = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryWordfence) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryWordfence) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryWordfence) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryWordfence) SetTitle(v string) {
	o.Title = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AdvisoryWordfence) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryWordfence) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AdvisoryWordfence) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AdvisoryWordfence) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryWordfence) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryWordfence) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryWordfence) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryWordfence) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisoryWordfence) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryWordfence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Affected) {
		toSerialize["affected"] = o.Affected
	}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.CvssScore) {
		toSerialize["cvss_score"] = o.CvssScore
	}
	if !IsNil(o.CvssVector) {
		toSerialize["cvss_vector"] = o.CvssVector
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Fixed) {
		toSerialize["fixed"] = o.Fixed
	}
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAdvisoryWordfence struct {
	value *AdvisoryWordfence
	isSet bool
}

func (v NullableAdvisoryWordfence) Get() *AdvisoryWordfence {
	return v.value
}

func (v *NullableAdvisoryWordfence) Set(val *AdvisoryWordfence) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryWordfence) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryWordfence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryWordfence(val *AdvisoryWordfence) *NullableAdvisoryWordfence {
	return &NullableAdvisoryWordfence{value: val, isSet: true}
}

func (v NullableAdvisoryWordfence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryWordfence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


