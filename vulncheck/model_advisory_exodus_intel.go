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

// checks if the AdvisoryExodusIntel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryExodusIntel{}

// AdvisoryExodusIntel struct for AdvisoryExodusIntel
type AdvisoryExodusIntel struct {
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	DisclosedPublic *string `json:"disclosed_public,omitempty"`
	DisclosedVendor *string `json:"disclosed_vendor,omitempty"`
	Id *string `json:"id,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Title *string `json:"title,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAdvisoryExodusIntel instantiates a new AdvisoryExodusIntel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryExodusIntel() *AdvisoryExodusIntel {
	this := AdvisoryExodusIntel{}
	return &this
}

// NewAdvisoryExodusIntelWithDefaults instantiates a new AdvisoryExodusIntel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryExodusIntelWithDefaults() *AdvisoryExodusIntel {
	this := AdvisoryExodusIntel{}
	return &this
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryExodusIntel) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryExodusIntel) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryExodusIntel) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryExodusIntel) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryExodusIntel) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryExodusIntel) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryExodusIntel) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryExodusIntel) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetDisclosedPublic returns the DisclosedPublic field value if set, zero value otherwise.
func (o *AdvisoryExodusIntel) GetDisclosedPublic() string {
	if o == nil || IsNil(o.DisclosedPublic) {
		var ret string
		return ret
	}
	return *o.DisclosedPublic
}

// GetDisclosedPublicOk returns a tuple with the DisclosedPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryExodusIntel) GetDisclosedPublicOk() (*string, bool) {
	if o == nil || IsNil(o.DisclosedPublic) {
		return nil, false
	}
	return o.DisclosedPublic, true
}

// HasDisclosedPublic returns a boolean if a field has been set.
func (o *AdvisoryExodusIntel) HasDisclosedPublic() bool {
	if o != nil && !IsNil(o.DisclosedPublic) {
		return true
	}

	return false
}

// SetDisclosedPublic gets a reference to the given string and assigns it to the DisclosedPublic field.
func (o *AdvisoryExodusIntel) SetDisclosedPublic(v string) {
	o.DisclosedPublic = &v
}

// GetDisclosedVendor returns the DisclosedVendor field value if set, zero value otherwise.
func (o *AdvisoryExodusIntel) GetDisclosedVendor() string {
	if o == nil || IsNil(o.DisclosedVendor) {
		var ret string
		return ret
	}
	return *o.DisclosedVendor
}

// GetDisclosedVendorOk returns a tuple with the DisclosedVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryExodusIntel) GetDisclosedVendorOk() (*string, bool) {
	if o == nil || IsNil(o.DisclosedVendor) {
		return nil, false
	}
	return o.DisclosedVendor, true
}

// HasDisclosedVendor returns a boolean if a field has been set.
func (o *AdvisoryExodusIntel) HasDisclosedVendor() bool {
	if o != nil && !IsNil(o.DisclosedVendor) {
		return true
	}

	return false
}

// SetDisclosedVendor gets a reference to the given string and assigns it to the DisclosedVendor field.
func (o *AdvisoryExodusIntel) SetDisclosedVendor(v string) {
	o.DisclosedVendor = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdvisoryExodusIntel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryExodusIntel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdvisoryExodusIntel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdvisoryExodusIntel) SetId(v string) {
	o.Id = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AdvisoryExodusIntel) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryExodusIntel) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AdvisoryExodusIntel) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AdvisoryExodusIntel) SetSummary(v string) {
	o.Summary = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryExodusIntel) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryExodusIntel) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryExodusIntel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryExodusIntel) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryExodusIntel) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryExodusIntel) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryExodusIntel) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryExodusIntel) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisoryExodusIntel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryExodusIntel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.DisclosedPublic) {
		toSerialize["disclosed_public"] = o.DisclosedPublic
	}
	if !IsNil(o.DisclosedVendor) {
		toSerialize["disclosed_vendor"] = o.DisclosedVendor
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
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

type NullableAdvisoryExodusIntel struct {
	value *AdvisoryExodusIntel
	isSet bool
}

func (v NullableAdvisoryExodusIntel) Get() *AdvisoryExodusIntel {
	return v.value
}

func (v *NullableAdvisoryExodusIntel) Set(val *AdvisoryExodusIntel) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryExodusIntel) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryExodusIntel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryExodusIntel(val *AdvisoryExodusIntel) *NullableAdvisoryExodusIntel {
	return &NullableAdvisoryExodusIntel{value: val, isSet: true}
}

func (v NullableAdvisoryExodusIntel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryExodusIntel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


