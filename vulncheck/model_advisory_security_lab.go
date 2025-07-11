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

// checks if the AdvisorySecurityLab type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisorySecurityLab{}

// AdvisorySecurityLab struct for AdvisorySecurityLab
type AdvisorySecurityLab struct {
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Id *string `json:"id,omitempty"`
	TitleRu *string `json:"title_ru,omitempty"`
	Url *string `json:"url,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
}

// NewAdvisorySecurityLab instantiates a new AdvisorySecurityLab object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisorySecurityLab() *AdvisorySecurityLab {
	this := AdvisorySecurityLab{}
	return &this
}

// NewAdvisorySecurityLabWithDefaults instantiates a new AdvisorySecurityLab object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisorySecurityLabWithDefaults() *AdvisorySecurityLab {
	this := AdvisorySecurityLab{}
	return &this
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisorySecurityLab) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySecurityLab) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisorySecurityLab) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisorySecurityLab) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisorySecurityLab) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySecurityLab) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisorySecurityLab) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisorySecurityLab) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdvisorySecurityLab) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySecurityLab) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdvisorySecurityLab) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdvisorySecurityLab) SetId(v string) {
	o.Id = &v
}

// GetTitleRu returns the TitleRu field value if set, zero value otherwise.
func (o *AdvisorySecurityLab) GetTitleRu() string {
	if o == nil || IsNil(o.TitleRu) {
		var ret string
		return ret
	}
	return *o.TitleRu
}

// GetTitleRuOk returns a tuple with the TitleRu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySecurityLab) GetTitleRuOk() (*string, bool) {
	if o == nil || IsNil(o.TitleRu) {
		return nil, false
	}
	return o.TitleRu, true
}

// HasTitleRu returns a boolean if a field has been set.
func (o *AdvisorySecurityLab) HasTitleRu() bool {
	if o != nil && !IsNil(o.TitleRu) {
		return true
	}

	return false
}

// SetTitleRu gets a reference to the given string and assigns it to the TitleRu field.
func (o *AdvisorySecurityLab) SetTitleRu(v string) {
	o.TitleRu = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisorySecurityLab) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySecurityLab) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisorySecurityLab) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisorySecurityLab) SetUrl(v string) {
	o.Url = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *AdvisorySecurityLab) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySecurityLab) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *AdvisorySecurityLab) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *AdvisorySecurityLab) SetVendor(v string) {
	o.Vendor = &v
}

func (o AdvisorySecurityLab) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisorySecurityLab) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TitleRu) {
		toSerialize["title_ru"] = o.TitleRu
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	return toSerialize, nil
}

type NullableAdvisorySecurityLab struct {
	value *AdvisorySecurityLab
	isSet bool
}

func (v NullableAdvisorySecurityLab) Get() *AdvisorySecurityLab {
	return v.value
}

func (v *NullableAdvisorySecurityLab) Set(val *AdvisorySecurityLab) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisorySecurityLab) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisorySecurityLab) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisorySecurityLab(val *AdvisorySecurityLab) *NullableAdvisorySecurityLab {
	return &NullableAdvisorySecurityLab{value: val, isSet: true}
}

func (v NullableAdvisorySecurityLab) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisorySecurityLab) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


