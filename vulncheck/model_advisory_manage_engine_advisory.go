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

// checks if the AdvisoryManageEngineAdvisory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryManageEngineAdvisory{}

// AdvisoryManageEngineAdvisory struct for AdvisoryManageEngineAdvisory
type AdvisoryManageEngineAdvisory struct {
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	ManageEngine *AdvisoryManageEngine `json:"manage_engine,omitempty"`
	Title *string `json:"title,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAdvisoryManageEngineAdvisory instantiates a new AdvisoryManageEngineAdvisory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryManageEngineAdvisory() *AdvisoryManageEngineAdvisory {
	this := AdvisoryManageEngineAdvisory{}
	return &this
}

// NewAdvisoryManageEngineAdvisoryWithDefaults instantiates a new AdvisoryManageEngineAdvisory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryManageEngineAdvisoryWithDefaults() *AdvisoryManageEngineAdvisory {
	this := AdvisoryManageEngineAdvisory{}
	return &this
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryManageEngineAdvisory) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngineAdvisory) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryManageEngineAdvisory) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryManageEngineAdvisory) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryManageEngineAdvisory) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngineAdvisory) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryManageEngineAdvisory) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryManageEngineAdvisory) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetManageEngine returns the ManageEngine field value if set, zero value otherwise.
func (o *AdvisoryManageEngineAdvisory) GetManageEngine() AdvisoryManageEngine {
	if o == nil || IsNil(o.ManageEngine) {
		var ret AdvisoryManageEngine
		return ret
	}
	return *o.ManageEngine
}

// GetManageEngineOk returns a tuple with the ManageEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngineAdvisory) GetManageEngineOk() (*AdvisoryManageEngine, bool) {
	if o == nil || IsNil(o.ManageEngine) {
		return nil, false
	}
	return o.ManageEngine, true
}

// HasManageEngine returns a boolean if a field has been set.
func (o *AdvisoryManageEngineAdvisory) HasManageEngine() bool {
	if o != nil && !IsNil(o.ManageEngine) {
		return true
	}

	return false
}

// SetManageEngine gets a reference to the given AdvisoryManageEngine and assigns it to the ManageEngine field.
func (o *AdvisoryManageEngineAdvisory) SetManageEngine(v AdvisoryManageEngine) {
	o.ManageEngine = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryManageEngineAdvisory) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngineAdvisory) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryManageEngineAdvisory) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryManageEngineAdvisory) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryManageEngineAdvisory) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngineAdvisory) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryManageEngineAdvisory) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryManageEngineAdvisory) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisoryManageEngineAdvisory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryManageEngineAdvisory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.ManageEngine) {
		toSerialize["manage_engine"] = o.ManageEngine
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAdvisoryManageEngineAdvisory struct {
	value *AdvisoryManageEngineAdvisory
	isSet bool
}

func (v NullableAdvisoryManageEngineAdvisory) Get() *AdvisoryManageEngineAdvisory {
	return v.value
}

func (v *NullableAdvisoryManageEngineAdvisory) Set(val *AdvisoryManageEngineAdvisory) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryManageEngineAdvisory) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryManageEngineAdvisory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryManageEngineAdvisory(val *AdvisoryManageEngineAdvisory) *NullableAdvisoryManageEngineAdvisory {
	return &NullableAdvisoryManageEngineAdvisory{value: val, isSet: true}
}

func (v NullableAdvisoryManageEngineAdvisory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryManageEngineAdvisory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


