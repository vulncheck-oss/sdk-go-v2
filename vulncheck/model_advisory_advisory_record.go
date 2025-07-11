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

// checks if the AdvisoryAdvisoryRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryAdvisoryRecord{}

// AdvisoryAdvisoryRecord struct for AdvisoryAdvisoryRecord
type AdvisoryAdvisoryRecord struct {
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	ExternalId []string `json:"external_id,omitempty"`
	Lang *string `json:"lang,omitempty"`
	Name *string `json:"name,omitempty"`
	Refsource *string `json:"refsource,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAdvisoryAdvisoryRecord instantiates a new AdvisoryAdvisoryRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryAdvisoryRecord() *AdvisoryAdvisoryRecord {
	this := AdvisoryAdvisoryRecord{}
	return &this
}

// NewAdvisoryAdvisoryRecordWithDefaults instantiates a new AdvisoryAdvisoryRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryAdvisoryRecordWithDefaults() *AdvisoryAdvisoryRecord {
	this := AdvisoryAdvisoryRecord{}
	return &this
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryAdvisoryRecord) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAdvisoryRecord) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryAdvisoryRecord) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryAdvisoryRecord) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryAdvisoryRecord) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAdvisoryRecord) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryAdvisoryRecord) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryAdvisoryRecord) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *AdvisoryAdvisoryRecord) GetExternalId() []string {
	if o == nil || IsNil(o.ExternalId) {
		var ret []string
		return ret
	}
	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAdvisoryRecord) GetExternalIdOk() ([]string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *AdvisoryAdvisoryRecord) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given []string and assigns it to the ExternalId field.
func (o *AdvisoryAdvisoryRecord) SetExternalId(v []string) {
	o.ExternalId = v
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *AdvisoryAdvisoryRecord) GetLang() string {
	if o == nil || IsNil(o.Lang) {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAdvisoryRecord) GetLangOk() (*string, bool) {
	if o == nil || IsNil(o.Lang) {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *AdvisoryAdvisoryRecord) HasLang() bool {
	if o != nil && !IsNil(o.Lang) {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *AdvisoryAdvisoryRecord) SetLang(v string) {
	o.Lang = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdvisoryAdvisoryRecord) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAdvisoryRecord) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdvisoryAdvisoryRecord) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdvisoryAdvisoryRecord) SetName(v string) {
	o.Name = &v
}

// GetRefsource returns the Refsource field value if set, zero value otherwise.
func (o *AdvisoryAdvisoryRecord) GetRefsource() string {
	if o == nil || IsNil(o.Refsource) {
		var ret string
		return ret
	}
	return *o.Refsource
}

// GetRefsourceOk returns a tuple with the Refsource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAdvisoryRecord) GetRefsourceOk() (*string, bool) {
	if o == nil || IsNil(o.Refsource) {
		return nil, false
	}
	return o.Refsource, true
}

// HasRefsource returns a boolean if a field has been set.
func (o *AdvisoryAdvisoryRecord) HasRefsource() bool {
	if o != nil && !IsNil(o.Refsource) {
		return true
	}

	return false
}

// SetRefsource gets a reference to the given string and assigns it to the Refsource field.
func (o *AdvisoryAdvisoryRecord) SetRefsource(v string) {
	o.Refsource = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AdvisoryAdvisoryRecord) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAdvisoryRecord) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AdvisoryAdvisoryRecord) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AdvisoryAdvisoryRecord) SetTags(v []string) {
	o.Tags = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryAdvisoryRecord) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAdvisoryRecord) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryAdvisoryRecord) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryAdvisoryRecord) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisoryAdvisoryRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryAdvisoryRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	if !IsNil(o.Lang) {
		toSerialize["lang"] = o.Lang
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Refsource) {
		toSerialize["refsource"] = o.Refsource
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAdvisoryAdvisoryRecord struct {
	value *AdvisoryAdvisoryRecord
	isSet bool
}

func (v NullableAdvisoryAdvisoryRecord) Get() *AdvisoryAdvisoryRecord {
	return v.value
}

func (v *NullableAdvisoryAdvisoryRecord) Set(val *AdvisoryAdvisoryRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryAdvisoryRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryAdvisoryRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryAdvisoryRecord(val *AdvisoryAdvisoryRecord) *NullableAdvisoryAdvisoryRecord {
	return &NullableAdvisoryAdvisoryRecord{value: val, isSet: true}
}

func (v NullableAdvisoryAdvisoryRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryAdvisoryRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


