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

// checks if the AdvisoryCloudVulnDBAdvisory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryCloudVulnDBAdvisory{}

// AdvisoryCloudVulnDBAdvisory struct for AdvisoryCloudVulnDBAdvisory
type AdvisoryCloudVulnDBAdvisory struct {
	AffectedServices *string `json:"affectedServices,omitempty"`
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Link *string `json:"link,omitempty"`
	References []string `json:"references,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewAdvisoryCloudVulnDBAdvisory instantiates a new AdvisoryCloudVulnDBAdvisory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryCloudVulnDBAdvisory() *AdvisoryCloudVulnDBAdvisory {
	this := AdvisoryCloudVulnDBAdvisory{}
	return &this
}

// NewAdvisoryCloudVulnDBAdvisoryWithDefaults instantiates a new AdvisoryCloudVulnDBAdvisory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryCloudVulnDBAdvisoryWithDefaults() *AdvisoryCloudVulnDBAdvisory {
	this := AdvisoryCloudVulnDBAdvisory{}
	return &this
}

// GetAffectedServices returns the AffectedServices field value if set, zero value otherwise.
func (o *AdvisoryCloudVulnDBAdvisory) GetAffectedServices() string {
	if o == nil || IsNil(o.AffectedServices) {
		var ret string
		return ret
	}
	return *o.AffectedServices
}

// GetAffectedServicesOk returns a tuple with the AffectedServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCloudVulnDBAdvisory) GetAffectedServicesOk() (*string, bool) {
	if o == nil || IsNil(o.AffectedServices) {
		return nil, false
	}
	return o.AffectedServices, true
}

// HasAffectedServices returns a boolean if a field has been set.
func (o *AdvisoryCloudVulnDBAdvisory) HasAffectedServices() bool {
	if o != nil && !IsNil(o.AffectedServices) {
		return true
	}

	return false
}

// SetAffectedServices gets a reference to the given string and assigns it to the AffectedServices field.
func (o *AdvisoryCloudVulnDBAdvisory) SetAffectedServices(v string) {
	o.AffectedServices = &v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryCloudVulnDBAdvisory) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCloudVulnDBAdvisory) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryCloudVulnDBAdvisory) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryCloudVulnDBAdvisory) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryCloudVulnDBAdvisory) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCloudVulnDBAdvisory) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryCloudVulnDBAdvisory) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryCloudVulnDBAdvisory) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *AdvisoryCloudVulnDBAdvisory) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCloudVulnDBAdvisory) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *AdvisoryCloudVulnDBAdvisory) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *AdvisoryCloudVulnDBAdvisory) SetLink(v string) {
	o.Link = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *AdvisoryCloudVulnDBAdvisory) GetReferences() []string {
	if o == nil || IsNil(o.References) {
		var ret []string
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCloudVulnDBAdvisory) GetReferencesOk() ([]string, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *AdvisoryCloudVulnDBAdvisory) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []string and assigns it to the References field.
func (o *AdvisoryCloudVulnDBAdvisory) SetReferences(v []string) {
	o.References = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryCloudVulnDBAdvisory) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCloudVulnDBAdvisory) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryCloudVulnDBAdvisory) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryCloudVulnDBAdvisory) SetTitle(v string) {
	o.Title = &v
}

func (o AdvisoryCloudVulnDBAdvisory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryCloudVulnDBAdvisory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AffectedServices) {
		toSerialize["affectedServices"] = o.AffectedServices
	}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableAdvisoryCloudVulnDBAdvisory struct {
	value *AdvisoryCloudVulnDBAdvisory
	isSet bool
}

func (v NullableAdvisoryCloudVulnDBAdvisory) Get() *AdvisoryCloudVulnDBAdvisory {
	return v.value
}

func (v *NullableAdvisoryCloudVulnDBAdvisory) Set(val *AdvisoryCloudVulnDBAdvisory) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryCloudVulnDBAdvisory) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryCloudVulnDBAdvisory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryCloudVulnDBAdvisory(val *AdvisoryCloudVulnDBAdvisory) *NullableAdvisoryCloudVulnDBAdvisory {
	return &NullableAdvisoryCloudVulnDBAdvisory{value: val, isSet: true}
}

func (v NullableAdvisoryCloudVulnDBAdvisory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryCloudVulnDBAdvisory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


