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

// checks if the AdvisoryGoVulnEcosystemSpecific type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryGoVulnEcosystemSpecific{}

// AdvisoryGoVulnEcosystemSpecific struct for AdvisoryGoVulnEcosystemSpecific
type AdvisoryGoVulnEcosystemSpecific struct {
	Imports []AdvisoryGoVulnImport `json:"imports,omitempty"`
}

// NewAdvisoryGoVulnEcosystemSpecific instantiates a new AdvisoryGoVulnEcosystemSpecific object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryGoVulnEcosystemSpecific() *AdvisoryGoVulnEcosystemSpecific {
	this := AdvisoryGoVulnEcosystemSpecific{}
	return &this
}

// NewAdvisoryGoVulnEcosystemSpecificWithDefaults instantiates a new AdvisoryGoVulnEcosystemSpecific object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryGoVulnEcosystemSpecificWithDefaults() *AdvisoryGoVulnEcosystemSpecific {
	this := AdvisoryGoVulnEcosystemSpecific{}
	return &this
}

// GetImports returns the Imports field value if set, zero value otherwise.
func (o *AdvisoryGoVulnEcosystemSpecific) GetImports() []AdvisoryGoVulnImport {
	if o == nil || IsNil(o.Imports) {
		var ret []AdvisoryGoVulnImport
		return ret
	}
	return o.Imports
}

// GetImportsOk returns a tuple with the Imports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGoVulnEcosystemSpecific) GetImportsOk() ([]AdvisoryGoVulnImport, bool) {
	if o == nil || IsNil(o.Imports) {
		return nil, false
	}
	return o.Imports, true
}

// HasImports returns a boolean if a field has been set.
func (o *AdvisoryGoVulnEcosystemSpecific) HasImports() bool {
	if o != nil && !IsNil(o.Imports) {
		return true
	}

	return false
}

// SetImports gets a reference to the given []AdvisoryGoVulnImport and assigns it to the Imports field.
func (o *AdvisoryGoVulnEcosystemSpecific) SetImports(v []AdvisoryGoVulnImport) {
	o.Imports = v
}

func (o AdvisoryGoVulnEcosystemSpecific) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryGoVulnEcosystemSpecific) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Imports) {
		toSerialize["imports"] = o.Imports
	}
	return toSerialize, nil
}

type NullableAdvisoryGoVulnEcosystemSpecific struct {
	value *AdvisoryGoVulnEcosystemSpecific
	isSet bool
}

func (v NullableAdvisoryGoVulnEcosystemSpecific) Get() *AdvisoryGoVulnEcosystemSpecific {
	return v.value
}

func (v *NullableAdvisoryGoVulnEcosystemSpecific) Set(val *AdvisoryGoVulnEcosystemSpecific) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryGoVulnEcosystemSpecific) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryGoVulnEcosystemSpecific) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryGoVulnEcosystemSpecific(val *AdvisoryGoVulnEcosystemSpecific) *NullableAdvisoryGoVulnEcosystemSpecific {
	return &NullableAdvisoryGoVulnEcosystemSpecific{value: val, isSet: true}
}

func (v NullableAdvisoryGoVulnEcosystemSpecific) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryGoVulnEcosystemSpecific) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


