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

// checks if the AdvisoryAcronis type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryAcronis{}

// AdvisoryAcronis struct for AdvisoryAcronis
type AdvisoryAcronis struct {
	Cve []string `json:"cve,omitempty"`
	Cvss *string `json:"cvss,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAdvisoryAcronis instantiates a new AdvisoryAcronis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryAcronis() *AdvisoryAcronis {
	this := AdvisoryAcronis{}
	return &this
}

// NewAdvisoryAcronisWithDefaults instantiates a new AdvisoryAcronis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryAcronisWithDefaults() *AdvisoryAcronis {
	this := AdvisoryAcronis{}
	return &this
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryAcronis) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAcronis) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryAcronis) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryAcronis) SetCve(v []string) {
	o.Cve = v
}

// GetCvss returns the Cvss field value if set, zero value otherwise.
func (o *AdvisoryAcronis) GetCvss() string {
	if o == nil || IsNil(o.Cvss) {
		var ret string
		return ret
	}
	return *o.Cvss
}

// GetCvssOk returns a tuple with the Cvss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAcronis) GetCvssOk() (*string, bool) {
	if o == nil || IsNil(o.Cvss) {
		return nil, false
	}
	return o.Cvss, true
}

// HasCvss returns a boolean if a field has been set.
func (o *AdvisoryAcronis) HasCvss() bool {
	if o != nil && !IsNil(o.Cvss) {
		return true
	}

	return false
}

// SetCvss gets a reference to the given string and assigns it to the Cvss field.
func (o *AdvisoryAcronis) SetCvss(v string) {
	o.Cvss = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryAcronis) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAcronis) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryAcronis) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryAcronis) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AdvisoryAcronis) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAcronis) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AdvisoryAcronis) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AdvisoryAcronis) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdvisoryAcronis) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAcronis) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdvisoryAcronis) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdvisoryAcronis) SetId(v string) {
	o.Id = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AdvisoryAcronis) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAcronis) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AdvisoryAcronis) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AdvisoryAcronis) SetSummary(v string) {
	o.Summary = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryAcronis) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAcronis) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryAcronis) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryAcronis) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisoryAcronis) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryAcronis) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.Cvss) {
		toSerialize["cvss"] = o.Cvss
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAdvisoryAcronis struct {
	value *AdvisoryAcronis
	isSet bool
}

func (v NullableAdvisoryAcronis) Get() *AdvisoryAcronis {
	return v.value
}

func (v *NullableAdvisoryAcronis) Set(val *AdvisoryAcronis) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryAcronis) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryAcronis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryAcronis(val *AdvisoryAcronis) *NullableAdvisoryAcronis {
	return &NullableAdvisoryAcronis{value: val, isSet: true}
}

func (v NullableAdvisoryAcronis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryAcronis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


