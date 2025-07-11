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

// checks if the AdvisoryVulnrichmentCVERef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryVulnrichmentCVERef{}

// AdvisoryVulnrichmentCVERef struct for AdvisoryVulnrichmentCVERef
type AdvisoryVulnrichmentCVERef struct {
	Containers *AdvisoryVulnrichmentContainers `json:"containers,omitempty"`
	CveMetadata *AdvisoryMCveMetadata `json:"cveMetadata,omitempty"`
	DataType *string `json:"dataType,omitempty"`
	DataVersion *string `json:"dataVersion,omitempty"`
}

// NewAdvisoryVulnrichmentCVERef instantiates a new AdvisoryVulnrichmentCVERef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryVulnrichmentCVERef() *AdvisoryVulnrichmentCVERef {
	this := AdvisoryVulnrichmentCVERef{}
	return &this
}

// NewAdvisoryVulnrichmentCVERefWithDefaults instantiates a new AdvisoryVulnrichmentCVERef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryVulnrichmentCVERefWithDefaults() *AdvisoryVulnrichmentCVERef {
	this := AdvisoryVulnrichmentCVERef{}
	return &this
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *AdvisoryVulnrichmentCVERef) GetContainers() AdvisoryVulnrichmentContainers {
	if o == nil || IsNil(o.Containers) {
		var ret AdvisoryVulnrichmentContainers
		return ret
	}
	return *o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryVulnrichmentCVERef) GetContainersOk() (*AdvisoryVulnrichmentContainers, bool) {
	if o == nil || IsNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *AdvisoryVulnrichmentCVERef) HasContainers() bool {
	if o != nil && !IsNil(o.Containers) {
		return true
	}

	return false
}

// SetContainers gets a reference to the given AdvisoryVulnrichmentContainers and assigns it to the Containers field.
func (o *AdvisoryVulnrichmentCVERef) SetContainers(v AdvisoryVulnrichmentContainers) {
	o.Containers = &v
}

// GetCveMetadata returns the CveMetadata field value if set, zero value otherwise.
func (o *AdvisoryVulnrichmentCVERef) GetCveMetadata() AdvisoryMCveMetadata {
	if o == nil || IsNil(o.CveMetadata) {
		var ret AdvisoryMCveMetadata
		return ret
	}
	return *o.CveMetadata
}

// GetCveMetadataOk returns a tuple with the CveMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryVulnrichmentCVERef) GetCveMetadataOk() (*AdvisoryMCveMetadata, bool) {
	if o == nil || IsNil(o.CveMetadata) {
		return nil, false
	}
	return o.CveMetadata, true
}

// HasCveMetadata returns a boolean if a field has been set.
func (o *AdvisoryVulnrichmentCVERef) HasCveMetadata() bool {
	if o != nil && !IsNil(o.CveMetadata) {
		return true
	}

	return false
}

// SetCveMetadata gets a reference to the given AdvisoryMCveMetadata and assigns it to the CveMetadata field.
func (o *AdvisoryVulnrichmentCVERef) SetCveMetadata(v AdvisoryMCveMetadata) {
	o.CveMetadata = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *AdvisoryVulnrichmentCVERef) GetDataType() string {
	if o == nil || IsNil(o.DataType) {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryVulnrichmentCVERef) GetDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *AdvisoryVulnrichmentCVERef) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *AdvisoryVulnrichmentCVERef) SetDataType(v string) {
	o.DataType = &v
}

// GetDataVersion returns the DataVersion field value if set, zero value otherwise.
func (o *AdvisoryVulnrichmentCVERef) GetDataVersion() string {
	if o == nil || IsNil(o.DataVersion) {
		var ret string
		return ret
	}
	return *o.DataVersion
}

// GetDataVersionOk returns a tuple with the DataVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryVulnrichmentCVERef) GetDataVersionOk() (*string, bool) {
	if o == nil || IsNil(o.DataVersion) {
		return nil, false
	}
	return o.DataVersion, true
}

// HasDataVersion returns a boolean if a field has been set.
func (o *AdvisoryVulnrichmentCVERef) HasDataVersion() bool {
	if o != nil && !IsNil(o.DataVersion) {
		return true
	}

	return false
}

// SetDataVersion gets a reference to the given string and assigns it to the DataVersion field.
func (o *AdvisoryVulnrichmentCVERef) SetDataVersion(v string) {
	o.DataVersion = &v
}

func (o AdvisoryVulnrichmentCVERef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryVulnrichmentCVERef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Containers) {
		toSerialize["containers"] = o.Containers
	}
	if !IsNil(o.CveMetadata) {
		toSerialize["cveMetadata"] = o.CveMetadata
	}
	if !IsNil(o.DataType) {
		toSerialize["dataType"] = o.DataType
	}
	if !IsNil(o.DataVersion) {
		toSerialize["dataVersion"] = o.DataVersion
	}
	return toSerialize, nil
}

type NullableAdvisoryVulnrichmentCVERef struct {
	value *AdvisoryVulnrichmentCVERef
	isSet bool
}

func (v NullableAdvisoryVulnrichmentCVERef) Get() *AdvisoryVulnrichmentCVERef {
	return v.value
}

func (v *NullableAdvisoryVulnrichmentCVERef) Set(val *AdvisoryVulnrichmentCVERef) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryVulnrichmentCVERef) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryVulnrichmentCVERef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryVulnrichmentCVERef(val *AdvisoryVulnrichmentCVERef) *NullableAdvisoryVulnrichmentCVERef {
	return &NullableAdvisoryVulnrichmentCVERef{value: val, isSet: true}
}

func (v NullableAdvisoryVulnrichmentCVERef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryVulnrichmentCVERef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


