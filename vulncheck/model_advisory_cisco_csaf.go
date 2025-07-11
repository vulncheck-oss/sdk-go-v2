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

// checks if the AdvisoryCiscoCSAF type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryCiscoCSAF{}

// AdvisoryCiscoCSAF struct for AdvisoryCiscoCSAF
type AdvisoryCiscoCSAF struct {
	Csaf map[string]interface{} `json:"csaf,omitempty"`
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAdvisoryCiscoCSAF instantiates a new AdvisoryCiscoCSAF object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryCiscoCSAF() *AdvisoryCiscoCSAF {
	this := AdvisoryCiscoCSAF{}
	return &this
}

// NewAdvisoryCiscoCSAFWithDefaults instantiates a new AdvisoryCiscoCSAF object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryCiscoCSAFWithDefaults() *AdvisoryCiscoCSAF {
	this := AdvisoryCiscoCSAF{}
	return &this
}

// GetCsaf returns the Csaf field value if set, zero value otherwise.
func (o *AdvisoryCiscoCSAF) GetCsaf() map[string]interface{} {
	if o == nil || IsNil(o.Csaf) {
		var ret map[string]interface{}
		return ret
	}
	return o.Csaf
}

// GetCsafOk returns a tuple with the Csaf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCiscoCSAF) GetCsafOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Csaf) {
		return map[string]interface{}{}, false
	}
	return o.Csaf, true
}

// HasCsaf returns a boolean if a field has been set.
func (o *AdvisoryCiscoCSAF) HasCsaf() bool {
	if o != nil && !IsNil(o.Csaf) {
		return true
	}

	return false
}

// SetCsaf gets a reference to the given map[string]interface{} and assigns it to the Csaf field.
func (o *AdvisoryCiscoCSAF) SetCsaf(v map[string]interface{}) {
	o.Csaf = v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryCiscoCSAF) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCiscoCSAF) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryCiscoCSAF) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryCiscoCSAF) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryCiscoCSAF) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCiscoCSAF) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryCiscoCSAF) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryCiscoCSAF) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *AdvisoryCiscoCSAF) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCiscoCSAF) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *AdvisoryCiscoCSAF) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *AdvisoryCiscoCSAF) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryCiscoCSAF) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCiscoCSAF) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryCiscoCSAF) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryCiscoCSAF) SetTitle(v string) {
	o.Title = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AdvisoryCiscoCSAF) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCiscoCSAF) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AdvisoryCiscoCSAF) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AdvisoryCiscoCSAF) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryCiscoCSAF) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCiscoCSAF) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryCiscoCSAF) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryCiscoCSAF) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisoryCiscoCSAF) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryCiscoCSAF) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Csaf) {
		toSerialize["csaf"] = o.Csaf
	}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
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

type NullableAdvisoryCiscoCSAF struct {
	value *AdvisoryCiscoCSAF
	isSet bool
}

func (v NullableAdvisoryCiscoCSAF) Get() *AdvisoryCiscoCSAF {
	return v.value
}

func (v *NullableAdvisoryCiscoCSAF) Set(val *AdvisoryCiscoCSAF) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryCiscoCSAF) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryCiscoCSAF) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryCiscoCSAF(val *AdvisoryCiscoCSAF) *NullableAdvisoryCiscoCSAF {
	return &NullableAdvisoryCiscoCSAF{value: val, isSet: true}
}

func (v NullableAdvisoryCiscoCSAF) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryCiscoCSAF) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


