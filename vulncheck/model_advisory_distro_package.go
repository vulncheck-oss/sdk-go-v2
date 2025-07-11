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

// checks if the AdvisoryDistroPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryDistroPackage{}

// AdvisoryDistroPackage struct for AdvisoryDistroPackage
type AdvisoryDistroPackage struct {
	Binary *bool `json:"binary,omitempty"`
	Cve []string `json:"cve,omitempty"`
	License []string `json:"license,omitempty"`
	Name *string `json:"name,omitempty"`
	SecFixes []AdvisorySecFix `json:"secFixes,omitempty"`
	Versions []AdvisoryDistroVersion `json:"versions,omitempty"`
}

// NewAdvisoryDistroPackage instantiates a new AdvisoryDistroPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryDistroPackage() *AdvisoryDistroPackage {
	this := AdvisoryDistroPackage{}
	return &this
}

// NewAdvisoryDistroPackageWithDefaults instantiates a new AdvisoryDistroPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryDistroPackageWithDefaults() *AdvisoryDistroPackage {
	this := AdvisoryDistroPackage{}
	return &this
}

// GetBinary returns the Binary field value if set, zero value otherwise.
func (o *AdvisoryDistroPackage) GetBinary() bool {
	if o == nil || IsNil(o.Binary) {
		var ret bool
		return ret
	}
	return *o.Binary
}

// GetBinaryOk returns a tuple with the Binary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryDistroPackage) GetBinaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Binary) {
		return nil, false
	}
	return o.Binary, true
}

// HasBinary returns a boolean if a field has been set.
func (o *AdvisoryDistroPackage) HasBinary() bool {
	if o != nil && !IsNil(o.Binary) {
		return true
	}

	return false
}

// SetBinary gets a reference to the given bool and assigns it to the Binary field.
func (o *AdvisoryDistroPackage) SetBinary(v bool) {
	o.Binary = &v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryDistroPackage) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryDistroPackage) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryDistroPackage) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryDistroPackage) SetCve(v []string) {
	o.Cve = v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *AdvisoryDistroPackage) GetLicense() []string {
	if o == nil || IsNil(o.License) {
		var ret []string
		return ret
	}
	return o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryDistroPackage) GetLicenseOk() ([]string, bool) {
	if o == nil || IsNil(o.License) {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *AdvisoryDistroPackage) HasLicense() bool {
	if o != nil && !IsNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given []string and assigns it to the License field.
func (o *AdvisoryDistroPackage) SetLicense(v []string) {
	o.License = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdvisoryDistroPackage) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryDistroPackage) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdvisoryDistroPackage) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdvisoryDistroPackage) SetName(v string) {
	o.Name = &v
}

// GetSecFixes returns the SecFixes field value if set, zero value otherwise.
func (o *AdvisoryDistroPackage) GetSecFixes() []AdvisorySecFix {
	if o == nil || IsNil(o.SecFixes) {
		var ret []AdvisorySecFix
		return ret
	}
	return o.SecFixes
}

// GetSecFixesOk returns a tuple with the SecFixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryDistroPackage) GetSecFixesOk() ([]AdvisorySecFix, bool) {
	if o == nil || IsNil(o.SecFixes) {
		return nil, false
	}
	return o.SecFixes, true
}

// HasSecFixes returns a boolean if a field has been set.
func (o *AdvisoryDistroPackage) HasSecFixes() bool {
	if o != nil && !IsNil(o.SecFixes) {
		return true
	}

	return false
}

// SetSecFixes gets a reference to the given []AdvisorySecFix and assigns it to the SecFixes field.
func (o *AdvisoryDistroPackage) SetSecFixes(v []AdvisorySecFix) {
	o.SecFixes = v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *AdvisoryDistroPackage) GetVersions() []AdvisoryDistroVersion {
	if o == nil || IsNil(o.Versions) {
		var ret []AdvisoryDistroVersion
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryDistroPackage) GetVersionsOk() ([]AdvisoryDistroVersion, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *AdvisoryDistroPackage) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []AdvisoryDistroVersion and assigns it to the Versions field.
func (o *AdvisoryDistroPackage) SetVersions(v []AdvisoryDistroVersion) {
	o.Versions = v
}

func (o AdvisoryDistroPackage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryDistroPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Binary) {
		toSerialize["binary"] = o.Binary
	}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.License) {
		toSerialize["license"] = o.License
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SecFixes) {
		toSerialize["secFixes"] = o.SecFixes
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableAdvisoryDistroPackage struct {
	value *AdvisoryDistroPackage
	isSet bool
}

func (v NullableAdvisoryDistroPackage) Get() *AdvisoryDistroPackage {
	return v.value
}

func (v *NullableAdvisoryDistroPackage) Set(val *AdvisoryDistroPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryDistroPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryDistroPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryDistroPackage(val *AdvisoryDistroPackage) *NullableAdvisoryDistroPackage {
	return &NullableAdvisoryDistroPackage{value: val, isSet: true}
}

func (v NullableAdvisoryDistroPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryDistroPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


