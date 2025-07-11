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

// checks if the AdvisoryNginxAdvisory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryNginxAdvisory{}

// AdvisoryNginxAdvisory struct for AdvisoryNginxAdvisory
type AdvisoryNginxAdvisory struct {
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Description *string `json:"description,omitempty"`
	NotVulnVersions []string `json:"not_vuln_versions,omitempty"`
	PatchPgp *string `json:"patch_pgp,omitempty"`
	PatchUrl *string `json:"patch_url,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Url *string `json:"url,omitempty"`
	VulnVersions []string `json:"vuln_versions,omitempty"`
}

// NewAdvisoryNginxAdvisory instantiates a new AdvisoryNginxAdvisory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryNginxAdvisory() *AdvisoryNginxAdvisory {
	this := AdvisoryNginxAdvisory{}
	return &this
}

// NewAdvisoryNginxAdvisoryWithDefaults instantiates a new AdvisoryNginxAdvisory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryNginxAdvisoryWithDefaults() *AdvisoryNginxAdvisory {
	this := AdvisoryNginxAdvisory{}
	return &this
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryNginxAdvisory) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNginxAdvisory) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryNginxAdvisory) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryNginxAdvisory) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryNginxAdvisory) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNginxAdvisory) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryNginxAdvisory) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryNginxAdvisory) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AdvisoryNginxAdvisory) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNginxAdvisory) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AdvisoryNginxAdvisory) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AdvisoryNginxAdvisory) SetDescription(v string) {
	o.Description = &v
}

// GetNotVulnVersions returns the NotVulnVersions field value if set, zero value otherwise.
func (o *AdvisoryNginxAdvisory) GetNotVulnVersions() []string {
	if o == nil || IsNil(o.NotVulnVersions) {
		var ret []string
		return ret
	}
	return o.NotVulnVersions
}

// GetNotVulnVersionsOk returns a tuple with the NotVulnVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNginxAdvisory) GetNotVulnVersionsOk() ([]string, bool) {
	if o == nil || IsNil(o.NotVulnVersions) {
		return nil, false
	}
	return o.NotVulnVersions, true
}

// HasNotVulnVersions returns a boolean if a field has been set.
func (o *AdvisoryNginxAdvisory) HasNotVulnVersions() bool {
	if o != nil && !IsNil(o.NotVulnVersions) {
		return true
	}

	return false
}

// SetNotVulnVersions gets a reference to the given []string and assigns it to the NotVulnVersions field.
func (o *AdvisoryNginxAdvisory) SetNotVulnVersions(v []string) {
	o.NotVulnVersions = v
}

// GetPatchPgp returns the PatchPgp field value if set, zero value otherwise.
func (o *AdvisoryNginxAdvisory) GetPatchPgp() string {
	if o == nil || IsNil(o.PatchPgp) {
		var ret string
		return ret
	}
	return *o.PatchPgp
}

// GetPatchPgpOk returns a tuple with the PatchPgp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNginxAdvisory) GetPatchPgpOk() (*string, bool) {
	if o == nil || IsNil(o.PatchPgp) {
		return nil, false
	}
	return o.PatchPgp, true
}

// HasPatchPgp returns a boolean if a field has been set.
func (o *AdvisoryNginxAdvisory) HasPatchPgp() bool {
	if o != nil && !IsNil(o.PatchPgp) {
		return true
	}

	return false
}

// SetPatchPgp gets a reference to the given string and assigns it to the PatchPgp field.
func (o *AdvisoryNginxAdvisory) SetPatchPgp(v string) {
	o.PatchPgp = &v
}

// GetPatchUrl returns the PatchUrl field value if set, zero value otherwise.
func (o *AdvisoryNginxAdvisory) GetPatchUrl() string {
	if o == nil || IsNil(o.PatchUrl) {
		var ret string
		return ret
	}
	return *o.PatchUrl
}

// GetPatchUrlOk returns a tuple with the PatchUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNginxAdvisory) GetPatchUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PatchUrl) {
		return nil, false
	}
	return o.PatchUrl, true
}

// HasPatchUrl returns a boolean if a field has been set.
func (o *AdvisoryNginxAdvisory) HasPatchUrl() bool {
	if o != nil && !IsNil(o.PatchUrl) {
		return true
	}

	return false
}

// SetPatchUrl gets a reference to the given string and assigns it to the PatchUrl field.
func (o *AdvisoryNginxAdvisory) SetPatchUrl(v string) {
	o.PatchUrl = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *AdvisoryNginxAdvisory) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNginxAdvisory) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *AdvisoryNginxAdvisory) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *AdvisoryNginxAdvisory) SetSeverity(v string) {
	o.Severity = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryNginxAdvisory) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNginxAdvisory) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryNginxAdvisory) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryNginxAdvisory) SetUrl(v string) {
	o.Url = &v
}

// GetVulnVersions returns the VulnVersions field value if set, zero value otherwise.
func (o *AdvisoryNginxAdvisory) GetVulnVersions() []string {
	if o == nil || IsNil(o.VulnVersions) {
		var ret []string
		return ret
	}
	return o.VulnVersions
}

// GetVulnVersionsOk returns a tuple with the VulnVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNginxAdvisory) GetVulnVersionsOk() ([]string, bool) {
	if o == nil || IsNil(o.VulnVersions) {
		return nil, false
	}
	return o.VulnVersions, true
}

// HasVulnVersions returns a boolean if a field has been set.
func (o *AdvisoryNginxAdvisory) HasVulnVersions() bool {
	if o != nil && !IsNil(o.VulnVersions) {
		return true
	}

	return false
}

// SetVulnVersions gets a reference to the given []string and assigns it to the VulnVersions field.
func (o *AdvisoryNginxAdvisory) SetVulnVersions(v []string) {
	o.VulnVersions = v
}

func (o AdvisoryNginxAdvisory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryNginxAdvisory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.NotVulnVersions) {
		toSerialize["not_vuln_versions"] = o.NotVulnVersions
	}
	if !IsNil(o.PatchPgp) {
		toSerialize["patch_pgp"] = o.PatchPgp
	}
	if !IsNil(o.PatchUrl) {
		toSerialize["patch_url"] = o.PatchUrl
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.VulnVersions) {
		toSerialize["vuln_versions"] = o.VulnVersions
	}
	return toSerialize, nil
}

type NullableAdvisoryNginxAdvisory struct {
	value *AdvisoryNginxAdvisory
	isSet bool
}

func (v NullableAdvisoryNginxAdvisory) Get() *AdvisoryNginxAdvisory {
	return v.value
}

func (v *NullableAdvisoryNginxAdvisory) Set(val *AdvisoryNginxAdvisory) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryNginxAdvisory) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryNginxAdvisory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryNginxAdvisory(val *AdvisoryNginxAdvisory) *NullableAdvisoryNginxAdvisory {
	return &NullableAdvisoryNginxAdvisory{value: val, isSet: true}
}

func (v NullableAdvisoryNginxAdvisory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryNginxAdvisory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


