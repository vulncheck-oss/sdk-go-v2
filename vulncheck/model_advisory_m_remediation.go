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

// checks if the AdvisoryMRemediation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryMRemediation{}

// AdvisoryMRemediation struct for AdvisoryMRemediation
type AdvisoryMRemediation struct {
	AffectedFiles []AdvisoryAffectedFile `json:"AffectedFiles,omitempty"`
	Date *string `json:"Date,omitempty"`
	DateSpecified *bool `json:"DateSpecified,omitempty"`
	Description *AdvisoryIVal `json:"Description,omitempty"`
	FixedBuild *string `json:"FixedBuild,omitempty"`
	ProductID []string `json:"ProductID,omitempty"`
	RestartRequired *AdvisoryIVal `json:"RestartRequired,omitempty"`
	SubType *string `json:"SubType,omitempty"`
	// diff
	Type *int32 `json:"Type,omitempty"`
	Url *string `json:"Url,omitempty"`
	Supercedence *string `json:"supercedence,omitempty"`
}

// NewAdvisoryMRemediation instantiates a new AdvisoryMRemediation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryMRemediation() *AdvisoryMRemediation {
	this := AdvisoryMRemediation{}
	return &this
}

// NewAdvisoryMRemediationWithDefaults instantiates a new AdvisoryMRemediation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryMRemediationWithDefaults() *AdvisoryMRemediation {
	this := AdvisoryMRemediation{}
	return &this
}

// GetAffectedFiles returns the AffectedFiles field value if set, zero value otherwise.
func (o *AdvisoryMRemediation) GetAffectedFiles() []AdvisoryAffectedFile {
	if o == nil || IsNil(o.AffectedFiles) {
		var ret []AdvisoryAffectedFile
		return ret
	}
	return o.AffectedFiles
}

// GetAffectedFilesOk returns a tuple with the AffectedFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMRemediation) GetAffectedFilesOk() ([]AdvisoryAffectedFile, bool) {
	if o == nil || IsNil(o.AffectedFiles) {
		return nil, false
	}
	return o.AffectedFiles, true
}

// HasAffectedFiles returns a boolean if a field has been set.
func (o *AdvisoryMRemediation) HasAffectedFiles() bool {
	if o != nil && !IsNil(o.AffectedFiles) {
		return true
	}

	return false
}

// SetAffectedFiles gets a reference to the given []AdvisoryAffectedFile and assigns it to the AffectedFiles field.
func (o *AdvisoryMRemediation) SetAffectedFiles(v []AdvisoryAffectedFile) {
	o.AffectedFiles = v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *AdvisoryMRemediation) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMRemediation) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *AdvisoryMRemediation) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *AdvisoryMRemediation) SetDate(v string) {
	o.Date = &v
}

// GetDateSpecified returns the DateSpecified field value if set, zero value otherwise.
func (o *AdvisoryMRemediation) GetDateSpecified() bool {
	if o == nil || IsNil(o.DateSpecified) {
		var ret bool
		return ret
	}
	return *o.DateSpecified
}

// GetDateSpecifiedOk returns a tuple with the DateSpecified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMRemediation) GetDateSpecifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.DateSpecified) {
		return nil, false
	}
	return o.DateSpecified, true
}

// HasDateSpecified returns a boolean if a field has been set.
func (o *AdvisoryMRemediation) HasDateSpecified() bool {
	if o != nil && !IsNil(o.DateSpecified) {
		return true
	}

	return false
}

// SetDateSpecified gets a reference to the given bool and assigns it to the DateSpecified field.
func (o *AdvisoryMRemediation) SetDateSpecified(v bool) {
	o.DateSpecified = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AdvisoryMRemediation) GetDescription() AdvisoryIVal {
	if o == nil || IsNil(o.Description) {
		var ret AdvisoryIVal
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMRemediation) GetDescriptionOk() (*AdvisoryIVal, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AdvisoryMRemediation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given AdvisoryIVal and assigns it to the Description field.
func (o *AdvisoryMRemediation) SetDescription(v AdvisoryIVal) {
	o.Description = &v
}

// GetFixedBuild returns the FixedBuild field value if set, zero value otherwise.
func (o *AdvisoryMRemediation) GetFixedBuild() string {
	if o == nil || IsNil(o.FixedBuild) {
		var ret string
		return ret
	}
	return *o.FixedBuild
}

// GetFixedBuildOk returns a tuple with the FixedBuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMRemediation) GetFixedBuildOk() (*string, bool) {
	if o == nil || IsNil(o.FixedBuild) {
		return nil, false
	}
	return o.FixedBuild, true
}

// HasFixedBuild returns a boolean if a field has been set.
func (o *AdvisoryMRemediation) HasFixedBuild() bool {
	if o != nil && !IsNil(o.FixedBuild) {
		return true
	}

	return false
}

// SetFixedBuild gets a reference to the given string and assigns it to the FixedBuild field.
func (o *AdvisoryMRemediation) SetFixedBuild(v string) {
	o.FixedBuild = &v
}

// GetProductID returns the ProductID field value if set, zero value otherwise.
func (o *AdvisoryMRemediation) GetProductID() []string {
	if o == nil || IsNil(o.ProductID) {
		var ret []string
		return ret
	}
	return o.ProductID
}

// GetProductIDOk returns a tuple with the ProductID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMRemediation) GetProductIDOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductID) {
		return nil, false
	}
	return o.ProductID, true
}

// HasProductID returns a boolean if a field has been set.
func (o *AdvisoryMRemediation) HasProductID() bool {
	if o != nil && !IsNil(o.ProductID) {
		return true
	}

	return false
}

// SetProductID gets a reference to the given []string and assigns it to the ProductID field.
func (o *AdvisoryMRemediation) SetProductID(v []string) {
	o.ProductID = v
}

// GetRestartRequired returns the RestartRequired field value if set, zero value otherwise.
func (o *AdvisoryMRemediation) GetRestartRequired() AdvisoryIVal {
	if o == nil || IsNil(o.RestartRequired) {
		var ret AdvisoryIVal
		return ret
	}
	return *o.RestartRequired
}

// GetRestartRequiredOk returns a tuple with the RestartRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMRemediation) GetRestartRequiredOk() (*AdvisoryIVal, bool) {
	if o == nil || IsNil(o.RestartRequired) {
		return nil, false
	}
	return o.RestartRequired, true
}

// HasRestartRequired returns a boolean if a field has been set.
func (o *AdvisoryMRemediation) HasRestartRequired() bool {
	if o != nil && !IsNil(o.RestartRequired) {
		return true
	}

	return false
}

// SetRestartRequired gets a reference to the given AdvisoryIVal and assigns it to the RestartRequired field.
func (o *AdvisoryMRemediation) SetRestartRequired(v AdvisoryIVal) {
	o.RestartRequired = &v
}

// GetSubType returns the SubType field value if set, zero value otherwise.
func (o *AdvisoryMRemediation) GetSubType() string {
	if o == nil || IsNil(o.SubType) {
		var ret string
		return ret
	}
	return *o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMRemediation) GetSubTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubType) {
		return nil, false
	}
	return o.SubType, true
}

// HasSubType returns a boolean if a field has been set.
func (o *AdvisoryMRemediation) HasSubType() bool {
	if o != nil && !IsNil(o.SubType) {
		return true
	}

	return false
}

// SetSubType gets a reference to the given string and assigns it to the SubType field.
func (o *AdvisoryMRemediation) SetSubType(v string) {
	o.SubType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AdvisoryMRemediation) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMRemediation) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AdvisoryMRemediation) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *AdvisoryMRemediation) SetType(v int32) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryMRemediation) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMRemediation) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryMRemediation) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryMRemediation) SetUrl(v string) {
	o.Url = &v
}

// GetSupercedence returns the Supercedence field value if set, zero value otherwise.
func (o *AdvisoryMRemediation) GetSupercedence() string {
	if o == nil || IsNil(o.Supercedence) {
		var ret string
		return ret
	}
	return *o.Supercedence
}

// GetSupercedenceOk returns a tuple with the Supercedence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMRemediation) GetSupercedenceOk() (*string, bool) {
	if o == nil || IsNil(o.Supercedence) {
		return nil, false
	}
	return o.Supercedence, true
}

// HasSupercedence returns a boolean if a field has been set.
func (o *AdvisoryMRemediation) HasSupercedence() bool {
	if o != nil && !IsNil(o.Supercedence) {
		return true
	}

	return false
}

// SetSupercedence gets a reference to the given string and assigns it to the Supercedence field.
func (o *AdvisoryMRemediation) SetSupercedence(v string) {
	o.Supercedence = &v
}

func (o AdvisoryMRemediation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryMRemediation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AffectedFiles) {
		toSerialize["AffectedFiles"] = o.AffectedFiles
	}
	if !IsNil(o.Date) {
		toSerialize["Date"] = o.Date
	}
	if !IsNil(o.DateSpecified) {
		toSerialize["DateSpecified"] = o.DateSpecified
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.FixedBuild) {
		toSerialize["FixedBuild"] = o.FixedBuild
	}
	if !IsNil(o.ProductID) {
		toSerialize["ProductID"] = o.ProductID
	}
	if !IsNil(o.RestartRequired) {
		toSerialize["RestartRequired"] = o.RestartRequired
	}
	if !IsNil(o.SubType) {
		toSerialize["SubType"] = o.SubType
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["Url"] = o.Url
	}
	if !IsNil(o.Supercedence) {
		toSerialize["supercedence"] = o.Supercedence
	}
	return toSerialize, nil
}

type NullableAdvisoryMRemediation struct {
	value *AdvisoryMRemediation
	isSet bool
}

func (v NullableAdvisoryMRemediation) Get() *AdvisoryMRemediation {
	return v.value
}

func (v *NullableAdvisoryMRemediation) Set(val *AdvisoryMRemediation) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryMRemediation) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryMRemediation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryMRemediation(val *AdvisoryMRemediation) *NullableAdvisoryMRemediation {
	return &NullableAdvisoryMRemediation{value: val, isSet: true}
}

func (v NullableAdvisoryMRemediation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryMRemediation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


