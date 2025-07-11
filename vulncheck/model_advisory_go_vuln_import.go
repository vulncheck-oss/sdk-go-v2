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

// checks if the AdvisoryGoVulnImport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryGoVulnImport{}

// AdvisoryGoVulnImport struct for AdvisoryGoVulnImport
type AdvisoryGoVulnImport struct {
	Path *string `json:"path,omitempty"`
	Symbols []string `json:"symbols,omitempty"`
}

// NewAdvisoryGoVulnImport instantiates a new AdvisoryGoVulnImport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryGoVulnImport() *AdvisoryGoVulnImport {
	this := AdvisoryGoVulnImport{}
	return &this
}

// NewAdvisoryGoVulnImportWithDefaults instantiates a new AdvisoryGoVulnImport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryGoVulnImportWithDefaults() *AdvisoryGoVulnImport {
	this := AdvisoryGoVulnImport{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *AdvisoryGoVulnImport) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGoVulnImport) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *AdvisoryGoVulnImport) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *AdvisoryGoVulnImport) SetPath(v string) {
	o.Path = &v
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *AdvisoryGoVulnImport) GetSymbols() []string {
	if o == nil || IsNil(o.Symbols) {
		var ret []string
		return ret
	}
	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGoVulnImport) GetSymbolsOk() ([]string, bool) {
	if o == nil || IsNil(o.Symbols) {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *AdvisoryGoVulnImport) HasSymbols() bool {
	if o != nil && !IsNil(o.Symbols) {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given []string and assigns it to the Symbols field.
func (o *AdvisoryGoVulnImport) SetSymbols(v []string) {
	o.Symbols = v
}

func (o AdvisoryGoVulnImport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryGoVulnImport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Symbols) {
		toSerialize["symbols"] = o.Symbols
	}
	return toSerialize, nil
}

type NullableAdvisoryGoVulnImport struct {
	value *AdvisoryGoVulnImport
	isSet bool
}

func (v NullableAdvisoryGoVulnImport) Get() *AdvisoryGoVulnImport {
	return v.value
}

func (v *NullableAdvisoryGoVulnImport) Set(val *AdvisoryGoVulnImport) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryGoVulnImport) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryGoVulnImport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryGoVulnImport(val *AdvisoryGoVulnImport) *NullableAdvisoryGoVulnImport {
	return &NullableAdvisoryGoVulnImport{value: val, isSet: true}
}

func (v NullableAdvisoryGoVulnImport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryGoVulnImport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


