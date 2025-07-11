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

// checks if the AdvisoryCBLMariner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryCBLMariner{}

// AdvisoryCBLMariner struct for AdvisoryCBLMariner
type AdvisoryCBLMariner struct {
	AdvisoryId *string `json:"advisory_id,omitempty"`
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Package *string `json:"package,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Title *string `json:"title,omitempty"`
	Url *string `json:"url,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewAdvisoryCBLMariner instantiates a new AdvisoryCBLMariner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryCBLMariner() *AdvisoryCBLMariner {
	this := AdvisoryCBLMariner{}
	return &this
}

// NewAdvisoryCBLMarinerWithDefaults instantiates a new AdvisoryCBLMariner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryCBLMarinerWithDefaults() *AdvisoryCBLMariner {
	this := AdvisoryCBLMariner{}
	return &this
}

// GetAdvisoryId returns the AdvisoryId field value if set, zero value otherwise.
func (o *AdvisoryCBLMariner) GetAdvisoryId() string {
	if o == nil || IsNil(o.AdvisoryId) {
		var ret string
		return ret
	}
	return *o.AdvisoryId
}

// GetAdvisoryIdOk returns a tuple with the AdvisoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCBLMariner) GetAdvisoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdvisoryId) {
		return nil, false
	}
	return o.AdvisoryId, true
}

// HasAdvisoryId returns a boolean if a field has been set.
func (o *AdvisoryCBLMariner) HasAdvisoryId() bool {
	if o != nil && !IsNil(o.AdvisoryId) {
		return true
	}

	return false
}

// SetAdvisoryId gets a reference to the given string and assigns it to the AdvisoryId field.
func (o *AdvisoryCBLMariner) SetAdvisoryId(v string) {
	o.AdvisoryId = &v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryCBLMariner) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCBLMariner) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryCBLMariner) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryCBLMariner) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryCBLMariner) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCBLMariner) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryCBLMariner) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryCBLMariner) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *AdvisoryCBLMariner) GetPackage() string {
	if o == nil || IsNil(o.Package) {
		var ret string
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCBLMariner) GetPackageOk() (*string, bool) {
	if o == nil || IsNil(o.Package) {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *AdvisoryCBLMariner) HasPackage() bool {
	if o != nil && !IsNil(o.Package) {
		return true
	}

	return false
}

// SetPackage gets a reference to the given string and assigns it to the Package field.
func (o *AdvisoryCBLMariner) SetPackage(v string) {
	o.Package = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *AdvisoryCBLMariner) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCBLMariner) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *AdvisoryCBLMariner) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *AdvisoryCBLMariner) SetSeverity(v string) {
	o.Severity = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AdvisoryCBLMariner) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCBLMariner) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AdvisoryCBLMariner) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AdvisoryCBLMariner) SetSummary(v string) {
	o.Summary = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryCBLMariner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCBLMariner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryCBLMariner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryCBLMariner) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryCBLMariner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCBLMariner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryCBLMariner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryCBLMariner) SetUrl(v string) {
	o.Url = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AdvisoryCBLMariner) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCBLMariner) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AdvisoryCBLMariner) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AdvisoryCBLMariner) SetVersion(v string) {
	o.Version = &v
}

func (o AdvisoryCBLMariner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryCBLMariner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdvisoryId) {
		toSerialize["advisory_id"] = o.AdvisoryId
	}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Package) {
		toSerialize["package"] = o.Package
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
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
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableAdvisoryCBLMariner struct {
	value *AdvisoryCBLMariner
	isSet bool
}

func (v NullableAdvisoryCBLMariner) Get() *AdvisoryCBLMariner {
	return v.value
}

func (v *NullableAdvisoryCBLMariner) Set(val *AdvisoryCBLMariner) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryCBLMariner) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryCBLMariner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryCBLMariner(val *AdvisoryCBLMariner) *NullableAdvisoryCBLMariner {
	return &NullableAdvisoryCBLMariner{value: val, isSet: true}
}

func (v NullableAdvisoryCBLMariner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryCBLMariner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


