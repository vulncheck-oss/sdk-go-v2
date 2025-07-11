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

// checks if the AdvisoryJuniper type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryJuniper{}

// AdvisoryJuniper struct for AdvisoryJuniper
type AdvisoryJuniper struct {
	Affected *string `json:"affected,omitempty"`
	Cve []string `json:"cve,omitempty"`
	Cvss3Score *string `json:"cvss3_score,omitempty"`
	Cvss3Vector *string `json:"cvss3_vector,omitempty"`
	Cvss4Score *string `json:"cvss4_score,omitempty"`
	Cvss4Vector *string `json:"cvss4_vector,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Fixed *string `json:"fixed,omitempty"`
	Id *string `json:"id,omitempty"`
	References []string `json:"references,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAdvisoryJuniper instantiates a new AdvisoryJuniper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryJuniper() *AdvisoryJuniper {
	this := AdvisoryJuniper{}
	return &this
}

// NewAdvisoryJuniperWithDefaults instantiates a new AdvisoryJuniper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryJuniperWithDefaults() *AdvisoryJuniper {
	this := AdvisoryJuniper{}
	return &this
}

// GetAffected returns the Affected field value if set, zero value otherwise.
func (o *AdvisoryJuniper) GetAffected() string {
	if o == nil || IsNil(o.Affected) {
		var ret string
		return ret
	}
	return *o.Affected
}

// GetAffectedOk returns a tuple with the Affected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryJuniper) GetAffectedOk() (*string, bool) {
	if o == nil || IsNil(o.Affected) {
		return nil, false
	}
	return o.Affected, true
}

// HasAffected returns a boolean if a field has been set.
func (o *AdvisoryJuniper) HasAffected() bool {
	if o != nil && !IsNil(o.Affected) {
		return true
	}

	return false
}

// SetAffected gets a reference to the given string and assigns it to the Affected field.
func (o *AdvisoryJuniper) SetAffected(v string) {
	o.Affected = &v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryJuniper) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryJuniper) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryJuniper) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryJuniper) SetCve(v []string) {
	o.Cve = v
}

// GetCvss3Score returns the Cvss3Score field value if set, zero value otherwise.
func (o *AdvisoryJuniper) GetCvss3Score() string {
	if o == nil || IsNil(o.Cvss3Score) {
		var ret string
		return ret
	}
	return *o.Cvss3Score
}

// GetCvss3ScoreOk returns a tuple with the Cvss3Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryJuniper) GetCvss3ScoreOk() (*string, bool) {
	if o == nil || IsNil(o.Cvss3Score) {
		return nil, false
	}
	return o.Cvss3Score, true
}

// HasCvss3Score returns a boolean if a field has been set.
func (o *AdvisoryJuniper) HasCvss3Score() bool {
	if o != nil && !IsNil(o.Cvss3Score) {
		return true
	}

	return false
}

// SetCvss3Score gets a reference to the given string and assigns it to the Cvss3Score field.
func (o *AdvisoryJuniper) SetCvss3Score(v string) {
	o.Cvss3Score = &v
}

// GetCvss3Vector returns the Cvss3Vector field value if set, zero value otherwise.
func (o *AdvisoryJuniper) GetCvss3Vector() string {
	if o == nil || IsNil(o.Cvss3Vector) {
		var ret string
		return ret
	}
	return *o.Cvss3Vector
}

// GetCvss3VectorOk returns a tuple with the Cvss3Vector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryJuniper) GetCvss3VectorOk() (*string, bool) {
	if o == nil || IsNil(o.Cvss3Vector) {
		return nil, false
	}
	return o.Cvss3Vector, true
}

// HasCvss3Vector returns a boolean if a field has been set.
func (o *AdvisoryJuniper) HasCvss3Vector() bool {
	if o != nil && !IsNil(o.Cvss3Vector) {
		return true
	}

	return false
}

// SetCvss3Vector gets a reference to the given string and assigns it to the Cvss3Vector field.
func (o *AdvisoryJuniper) SetCvss3Vector(v string) {
	o.Cvss3Vector = &v
}

// GetCvss4Score returns the Cvss4Score field value if set, zero value otherwise.
func (o *AdvisoryJuniper) GetCvss4Score() string {
	if o == nil || IsNil(o.Cvss4Score) {
		var ret string
		return ret
	}
	return *o.Cvss4Score
}

// GetCvss4ScoreOk returns a tuple with the Cvss4Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryJuniper) GetCvss4ScoreOk() (*string, bool) {
	if o == nil || IsNil(o.Cvss4Score) {
		return nil, false
	}
	return o.Cvss4Score, true
}

// HasCvss4Score returns a boolean if a field has been set.
func (o *AdvisoryJuniper) HasCvss4Score() bool {
	if o != nil && !IsNil(o.Cvss4Score) {
		return true
	}

	return false
}

// SetCvss4Score gets a reference to the given string and assigns it to the Cvss4Score field.
func (o *AdvisoryJuniper) SetCvss4Score(v string) {
	o.Cvss4Score = &v
}

// GetCvss4Vector returns the Cvss4Vector field value if set, zero value otherwise.
func (o *AdvisoryJuniper) GetCvss4Vector() string {
	if o == nil || IsNil(o.Cvss4Vector) {
		var ret string
		return ret
	}
	return *o.Cvss4Vector
}

// GetCvss4VectorOk returns a tuple with the Cvss4Vector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryJuniper) GetCvss4VectorOk() (*string, bool) {
	if o == nil || IsNil(o.Cvss4Vector) {
		return nil, false
	}
	return o.Cvss4Vector, true
}

// HasCvss4Vector returns a boolean if a field has been set.
func (o *AdvisoryJuniper) HasCvss4Vector() bool {
	if o != nil && !IsNil(o.Cvss4Vector) {
		return true
	}

	return false
}

// SetCvss4Vector gets a reference to the given string and assigns it to the Cvss4Vector field.
func (o *AdvisoryJuniper) SetCvss4Vector(v string) {
	o.Cvss4Vector = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryJuniper) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryJuniper) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryJuniper) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryJuniper) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetFixed returns the Fixed field value if set, zero value otherwise.
func (o *AdvisoryJuniper) GetFixed() string {
	if o == nil || IsNil(o.Fixed) {
		var ret string
		return ret
	}
	return *o.Fixed
}

// GetFixedOk returns a tuple with the Fixed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryJuniper) GetFixedOk() (*string, bool) {
	if o == nil || IsNil(o.Fixed) {
		return nil, false
	}
	return o.Fixed, true
}

// HasFixed returns a boolean if a field has been set.
func (o *AdvisoryJuniper) HasFixed() bool {
	if o != nil && !IsNil(o.Fixed) {
		return true
	}

	return false
}

// SetFixed gets a reference to the given string and assigns it to the Fixed field.
func (o *AdvisoryJuniper) SetFixed(v string) {
	o.Fixed = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdvisoryJuniper) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryJuniper) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdvisoryJuniper) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdvisoryJuniper) SetId(v string) {
	o.Id = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *AdvisoryJuniper) GetReferences() []string {
	if o == nil || IsNil(o.References) {
		var ret []string
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryJuniper) GetReferencesOk() ([]string, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *AdvisoryJuniper) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []string and assigns it to the References field.
func (o *AdvisoryJuniper) SetReferences(v []string) {
	o.References = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AdvisoryJuniper) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryJuniper) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AdvisoryJuniper) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AdvisoryJuniper) SetSummary(v string) {
	o.Summary = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryJuniper) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryJuniper) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryJuniper) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryJuniper) SetTitle(v string) {
	o.Title = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AdvisoryJuniper) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryJuniper) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AdvisoryJuniper) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AdvisoryJuniper) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryJuniper) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryJuniper) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryJuniper) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryJuniper) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisoryJuniper) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryJuniper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Affected) {
		toSerialize["affected"] = o.Affected
	}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.Cvss3Score) {
		toSerialize["cvss3_score"] = o.Cvss3Score
	}
	if !IsNil(o.Cvss3Vector) {
		toSerialize["cvss3_vector"] = o.Cvss3Vector
	}
	if !IsNil(o.Cvss4Score) {
		toSerialize["cvss4_score"] = o.Cvss4Score
	}
	if !IsNil(o.Cvss4Vector) {
		toSerialize["cvss4_vector"] = o.Cvss4Vector
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Fixed) {
		toSerialize["fixed"] = o.Fixed
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
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

type NullableAdvisoryJuniper struct {
	value *AdvisoryJuniper
	isSet bool
}

func (v NullableAdvisoryJuniper) Get() *AdvisoryJuniper {
	return v.value
}

func (v *NullableAdvisoryJuniper) Set(val *AdvisoryJuniper) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryJuniper) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryJuniper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryJuniper(val *AdvisoryJuniper) *NullableAdvisoryJuniper {
	return &NullableAdvisoryJuniper{value: val, isSet: true}
}

func (v NullableAdvisoryJuniper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryJuniper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


