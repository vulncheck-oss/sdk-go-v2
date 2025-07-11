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

// checks if the ApiNVD20ReferenceExtended type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiNVD20ReferenceExtended{}

// ApiNVD20ReferenceExtended struct for ApiNVD20ReferenceExtended
type ApiNVD20ReferenceExtended struct {
	DateAdded *string `json:"date_added,omitempty"`
	ExternalId *string `json:"external_id,omitempty"`
	Lang *string `json:"lang,omitempty"`
	Name *string `json:"name,omitempty"`
	PreviousUrl *string `json:"previous_url,omitempty"`
	Refsource *string `json:"refsource,omitempty"`
	Source *string `json:"source,omitempty"`
	Status *string `json:"status,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewApiNVD20ReferenceExtended instantiates a new ApiNVD20ReferenceExtended object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiNVD20ReferenceExtended() *ApiNVD20ReferenceExtended {
	this := ApiNVD20ReferenceExtended{}
	return &this
}

// NewApiNVD20ReferenceExtendedWithDefaults instantiates a new ApiNVD20ReferenceExtended object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiNVD20ReferenceExtendedWithDefaults() *ApiNVD20ReferenceExtended {
	this := ApiNVD20ReferenceExtended{}
	return &this
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *ApiNVD20ReferenceExtended) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNVD20ReferenceExtended) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *ApiNVD20ReferenceExtended) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *ApiNVD20ReferenceExtended) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ApiNVD20ReferenceExtended) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNVD20ReferenceExtended) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ApiNVD20ReferenceExtended) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ApiNVD20ReferenceExtended) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *ApiNVD20ReferenceExtended) GetLang() string {
	if o == nil || IsNil(o.Lang) {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNVD20ReferenceExtended) GetLangOk() (*string, bool) {
	if o == nil || IsNil(o.Lang) {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *ApiNVD20ReferenceExtended) HasLang() bool {
	if o != nil && !IsNil(o.Lang) {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *ApiNVD20ReferenceExtended) SetLang(v string) {
	o.Lang = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiNVD20ReferenceExtended) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNVD20ReferenceExtended) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiNVD20ReferenceExtended) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiNVD20ReferenceExtended) SetName(v string) {
	o.Name = &v
}

// GetPreviousUrl returns the PreviousUrl field value if set, zero value otherwise.
func (o *ApiNVD20ReferenceExtended) GetPreviousUrl() string {
	if o == nil || IsNil(o.PreviousUrl) {
		var ret string
		return ret
	}
	return *o.PreviousUrl
}

// GetPreviousUrlOk returns a tuple with the PreviousUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNVD20ReferenceExtended) GetPreviousUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousUrl) {
		return nil, false
	}
	return o.PreviousUrl, true
}

// HasPreviousUrl returns a boolean if a field has been set.
func (o *ApiNVD20ReferenceExtended) HasPreviousUrl() bool {
	if o != nil && !IsNil(o.PreviousUrl) {
		return true
	}

	return false
}

// SetPreviousUrl gets a reference to the given string and assigns it to the PreviousUrl field.
func (o *ApiNVD20ReferenceExtended) SetPreviousUrl(v string) {
	o.PreviousUrl = &v
}

// GetRefsource returns the Refsource field value if set, zero value otherwise.
func (o *ApiNVD20ReferenceExtended) GetRefsource() string {
	if o == nil || IsNil(o.Refsource) {
		var ret string
		return ret
	}
	return *o.Refsource
}

// GetRefsourceOk returns a tuple with the Refsource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNVD20ReferenceExtended) GetRefsourceOk() (*string, bool) {
	if o == nil || IsNil(o.Refsource) {
		return nil, false
	}
	return o.Refsource, true
}

// HasRefsource returns a boolean if a field has been set.
func (o *ApiNVD20ReferenceExtended) HasRefsource() bool {
	if o != nil && !IsNil(o.Refsource) {
		return true
	}

	return false
}

// SetRefsource gets a reference to the given string and assigns it to the Refsource field.
func (o *ApiNVD20ReferenceExtended) SetRefsource(v string) {
	o.Refsource = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ApiNVD20ReferenceExtended) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNVD20ReferenceExtended) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ApiNVD20ReferenceExtended) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ApiNVD20ReferenceExtended) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiNVD20ReferenceExtended) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNVD20ReferenceExtended) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiNVD20ReferenceExtended) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiNVD20ReferenceExtended) SetStatus(v string) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ApiNVD20ReferenceExtended) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNVD20ReferenceExtended) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApiNVD20ReferenceExtended) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ApiNVD20ReferenceExtended) SetTags(v []string) {
	o.Tags = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ApiNVD20ReferenceExtended) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNVD20ReferenceExtended) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ApiNVD20ReferenceExtended) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ApiNVD20ReferenceExtended) SetUrl(v string) {
	o.Url = &v
}

func (o ApiNVD20ReferenceExtended) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiNVD20ReferenceExtended) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.PreviousUrl) {
		toSerialize["previous_url"] = o.PreviousUrl
	}
	if !IsNil(o.Refsource) {
		toSerialize["refsource"] = o.Refsource
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableApiNVD20ReferenceExtended struct {
	value *ApiNVD20ReferenceExtended
	isSet bool
}

func (v NullableApiNVD20ReferenceExtended) Get() *ApiNVD20ReferenceExtended {
	return v.value
}

func (v *NullableApiNVD20ReferenceExtended) Set(val *ApiNVD20ReferenceExtended) {
	v.value = val
	v.isSet = true
}

func (v NullableApiNVD20ReferenceExtended) IsSet() bool {
	return v.isSet
}

func (v *NullableApiNVD20ReferenceExtended) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiNVD20ReferenceExtended(val *ApiNVD20ReferenceExtended) *NullableApiNVD20ReferenceExtended {
	return &NullableApiNVD20ReferenceExtended{value: val, isSet: true}
}

func (v NullableApiNVD20ReferenceExtended) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiNVD20ReferenceExtended) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


