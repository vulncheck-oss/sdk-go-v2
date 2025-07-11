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

// checks if the AdvisoryRockyAdvisory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryRockyAdvisory{}

// AdvisoryRockyAdvisory struct for AdvisoryRockyAdvisory
type AdvisoryRockyAdvisory struct {
	AffectedProducts []string `json:"affectedProducts,omitempty"`
	BuildReferences []string `json:"buildReferences,omitempty"`
	Cves []AdvisoryRockyCve `json:"cves,omitempty"`
	Description *string `json:"description,omitempty"`
	Fixes []AdvisoryRockyFix `json:"fixes,omitempty"`
	Name *string `json:"name,omitempty"`
	PublishedAt *string `json:"publishedAt,omitempty"`
	RebootSuggested *bool `json:"rebootSuggested,omitempty"`
	References []string `json:"references,omitempty"`
	Rpms *map[string]AdvisoryRockyVersion `json:"rpms,omitempty"`
	Severity *string `json:"severity,omitempty"`
	ShortCode *string `json:"shortCode,omitempty"`
	Solution *string `json:"solution,omitempty"`
	Synopsis *string `json:"synopsis,omitempty"`
	Topic *string `json:"topic,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewAdvisoryRockyAdvisory instantiates a new AdvisoryRockyAdvisory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryRockyAdvisory() *AdvisoryRockyAdvisory {
	this := AdvisoryRockyAdvisory{}
	return &this
}

// NewAdvisoryRockyAdvisoryWithDefaults instantiates a new AdvisoryRockyAdvisory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryRockyAdvisoryWithDefaults() *AdvisoryRockyAdvisory {
	this := AdvisoryRockyAdvisory{}
	return &this
}

// GetAffectedProducts returns the AffectedProducts field value if set, zero value otherwise.
func (o *AdvisoryRockyAdvisory) GetAffectedProducts() []string {
	if o == nil || IsNil(o.AffectedProducts) {
		var ret []string
		return ret
	}
	return o.AffectedProducts
}

// GetAffectedProductsOk returns a tuple with the AffectedProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRockyAdvisory) GetAffectedProductsOk() ([]string, bool) {
	if o == nil || IsNil(o.AffectedProducts) {
		return nil, false
	}
	return o.AffectedProducts, true
}

// HasAffectedProducts returns a boolean if a field has been set.
func (o *AdvisoryRockyAdvisory) HasAffectedProducts() bool {
	if o != nil && !IsNil(o.AffectedProducts) {
		return true
	}

	return false
}

// SetAffectedProducts gets a reference to the given []string and assigns it to the AffectedProducts field.
func (o *AdvisoryRockyAdvisory) SetAffectedProducts(v []string) {
	o.AffectedProducts = v
}

// GetBuildReferences returns the BuildReferences field value if set, zero value otherwise.
func (o *AdvisoryRockyAdvisory) GetBuildReferences() []string {
	if o == nil || IsNil(o.BuildReferences) {
		var ret []string
		return ret
	}
	return o.BuildReferences
}

// GetBuildReferencesOk returns a tuple with the BuildReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRockyAdvisory) GetBuildReferencesOk() ([]string, bool) {
	if o == nil || IsNil(o.BuildReferences) {
		return nil, false
	}
	return o.BuildReferences, true
}

// HasBuildReferences returns a boolean if a field has been set.
func (o *AdvisoryRockyAdvisory) HasBuildReferences() bool {
	if o != nil && !IsNil(o.BuildReferences) {
		return true
	}

	return false
}

// SetBuildReferences gets a reference to the given []string and assigns it to the BuildReferences field.
func (o *AdvisoryRockyAdvisory) SetBuildReferences(v []string) {
	o.BuildReferences = v
}

// GetCves returns the Cves field value if set, zero value otherwise.
func (o *AdvisoryRockyAdvisory) GetCves() []AdvisoryRockyCve {
	if o == nil || IsNil(o.Cves) {
		var ret []AdvisoryRockyCve
		return ret
	}
	return o.Cves
}

// GetCvesOk returns a tuple with the Cves field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRockyAdvisory) GetCvesOk() ([]AdvisoryRockyCve, bool) {
	if o == nil || IsNil(o.Cves) {
		return nil, false
	}
	return o.Cves, true
}

// HasCves returns a boolean if a field has been set.
func (o *AdvisoryRockyAdvisory) HasCves() bool {
	if o != nil && !IsNil(o.Cves) {
		return true
	}

	return false
}

// SetCves gets a reference to the given []AdvisoryRockyCve and assigns it to the Cves field.
func (o *AdvisoryRockyAdvisory) SetCves(v []AdvisoryRockyCve) {
	o.Cves = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AdvisoryRockyAdvisory) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRockyAdvisory) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AdvisoryRockyAdvisory) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AdvisoryRockyAdvisory) SetDescription(v string) {
	o.Description = &v
}

// GetFixes returns the Fixes field value if set, zero value otherwise.
func (o *AdvisoryRockyAdvisory) GetFixes() []AdvisoryRockyFix {
	if o == nil || IsNil(o.Fixes) {
		var ret []AdvisoryRockyFix
		return ret
	}
	return o.Fixes
}

// GetFixesOk returns a tuple with the Fixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRockyAdvisory) GetFixesOk() ([]AdvisoryRockyFix, bool) {
	if o == nil || IsNil(o.Fixes) {
		return nil, false
	}
	return o.Fixes, true
}

// HasFixes returns a boolean if a field has been set.
func (o *AdvisoryRockyAdvisory) HasFixes() bool {
	if o != nil && !IsNil(o.Fixes) {
		return true
	}

	return false
}

// SetFixes gets a reference to the given []AdvisoryRockyFix and assigns it to the Fixes field.
func (o *AdvisoryRockyAdvisory) SetFixes(v []AdvisoryRockyFix) {
	o.Fixes = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdvisoryRockyAdvisory) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRockyAdvisory) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdvisoryRockyAdvisory) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdvisoryRockyAdvisory) SetName(v string) {
	o.Name = &v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *AdvisoryRockyAdvisory) GetPublishedAt() string {
	if o == nil || IsNil(o.PublishedAt) {
		var ret string
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRockyAdvisory) GetPublishedAtOk() (*string, bool) {
	if o == nil || IsNil(o.PublishedAt) {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *AdvisoryRockyAdvisory) HasPublishedAt() bool {
	if o != nil && !IsNil(o.PublishedAt) {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given string and assigns it to the PublishedAt field.
func (o *AdvisoryRockyAdvisory) SetPublishedAt(v string) {
	o.PublishedAt = &v
}

// GetRebootSuggested returns the RebootSuggested field value if set, zero value otherwise.
func (o *AdvisoryRockyAdvisory) GetRebootSuggested() bool {
	if o == nil || IsNil(o.RebootSuggested) {
		var ret bool
		return ret
	}
	return *o.RebootSuggested
}

// GetRebootSuggestedOk returns a tuple with the RebootSuggested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRockyAdvisory) GetRebootSuggestedOk() (*bool, bool) {
	if o == nil || IsNil(o.RebootSuggested) {
		return nil, false
	}
	return o.RebootSuggested, true
}

// HasRebootSuggested returns a boolean if a field has been set.
func (o *AdvisoryRockyAdvisory) HasRebootSuggested() bool {
	if o != nil && !IsNil(o.RebootSuggested) {
		return true
	}

	return false
}

// SetRebootSuggested gets a reference to the given bool and assigns it to the RebootSuggested field.
func (o *AdvisoryRockyAdvisory) SetRebootSuggested(v bool) {
	o.RebootSuggested = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *AdvisoryRockyAdvisory) GetReferences() []string {
	if o == nil || IsNil(o.References) {
		var ret []string
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRockyAdvisory) GetReferencesOk() ([]string, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *AdvisoryRockyAdvisory) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []string and assigns it to the References field.
func (o *AdvisoryRockyAdvisory) SetReferences(v []string) {
	o.References = v
}

// GetRpms returns the Rpms field value if set, zero value otherwise.
func (o *AdvisoryRockyAdvisory) GetRpms() map[string]AdvisoryRockyVersion {
	if o == nil || IsNil(o.Rpms) {
		var ret map[string]AdvisoryRockyVersion
		return ret
	}
	return *o.Rpms
}

// GetRpmsOk returns a tuple with the Rpms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRockyAdvisory) GetRpmsOk() (*map[string]AdvisoryRockyVersion, bool) {
	if o == nil || IsNil(o.Rpms) {
		return nil, false
	}
	return o.Rpms, true
}

// HasRpms returns a boolean if a field has been set.
func (o *AdvisoryRockyAdvisory) HasRpms() bool {
	if o != nil && !IsNil(o.Rpms) {
		return true
	}

	return false
}

// SetRpms gets a reference to the given map[string]AdvisoryRockyVersion and assigns it to the Rpms field.
func (o *AdvisoryRockyAdvisory) SetRpms(v map[string]AdvisoryRockyVersion) {
	o.Rpms = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *AdvisoryRockyAdvisory) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRockyAdvisory) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *AdvisoryRockyAdvisory) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *AdvisoryRockyAdvisory) SetSeverity(v string) {
	o.Severity = &v
}

// GetShortCode returns the ShortCode field value if set, zero value otherwise.
func (o *AdvisoryRockyAdvisory) GetShortCode() string {
	if o == nil || IsNil(o.ShortCode) {
		var ret string
		return ret
	}
	return *o.ShortCode
}

// GetShortCodeOk returns a tuple with the ShortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRockyAdvisory) GetShortCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ShortCode) {
		return nil, false
	}
	return o.ShortCode, true
}

// HasShortCode returns a boolean if a field has been set.
func (o *AdvisoryRockyAdvisory) HasShortCode() bool {
	if o != nil && !IsNil(o.ShortCode) {
		return true
	}

	return false
}

// SetShortCode gets a reference to the given string and assigns it to the ShortCode field.
func (o *AdvisoryRockyAdvisory) SetShortCode(v string) {
	o.ShortCode = &v
}

// GetSolution returns the Solution field value if set, zero value otherwise.
func (o *AdvisoryRockyAdvisory) GetSolution() string {
	if o == nil || IsNil(o.Solution) {
		var ret string
		return ret
	}
	return *o.Solution
}

// GetSolutionOk returns a tuple with the Solution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRockyAdvisory) GetSolutionOk() (*string, bool) {
	if o == nil || IsNil(o.Solution) {
		return nil, false
	}
	return o.Solution, true
}

// HasSolution returns a boolean if a field has been set.
func (o *AdvisoryRockyAdvisory) HasSolution() bool {
	if o != nil && !IsNil(o.Solution) {
		return true
	}

	return false
}

// SetSolution gets a reference to the given string and assigns it to the Solution field.
func (o *AdvisoryRockyAdvisory) SetSolution(v string) {
	o.Solution = &v
}

// GetSynopsis returns the Synopsis field value if set, zero value otherwise.
func (o *AdvisoryRockyAdvisory) GetSynopsis() string {
	if o == nil || IsNil(o.Synopsis) {
		var ret string
		return ret
	}
	return *o.Synopsis
}

// GetSynopsisOk returns a tuple with the Synopsis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRockyAdvisory) GetSynopsisOk() (*string, bool) {
	if o == nil || IsNil(o.Synopsis) {
		return nil, false
	}
	return o.Synopsis, true
}

// HasSynopsis returns a boolean if a field has been set.
func (o *AdvisoryRockyAdvisory) HasSynopsis() bool {
	if o != nil && !IsNil(o.Synopsis) {
		return true
	}

	return false
}

// SetSynopsis gets a reference to the given string and assigns it to the Synopsis field.
func (o *AdvisoryRockyAdvisory) SetSynopsis(v string) {
	o.Synopsis = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *AdvisoryRockyAdvisory) GetTopic() string {
	if o == nil || IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRockyAdvisory) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *AdvisoryRockyAdvisory) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *AdvisoryRockyAdvisory) SetTopic(v string) {
	o.Topic = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AdvisoryRockyAdvisory) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRockyAdvisory) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AdvisoryRockyAdvisory) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AdvisoryRockyAdvisory) SetType(v string) {
	o.Type = &v
}

func (o AdvisoryRockyAdvisory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryRockyAdvisory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AffectedProducts) {
		toSerialize["affectedProducts"] = o.AffectedProducts
	}
	if !IsNil(o.BuildReferences) {
		toSerialize["buildReferences"] = o.BuildReferences
	}
	if !IsNil(o.Cves) {
		toSerialize["cves"] = o.Cves
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Fixes) {
		toSerialize["fixes"] = o.Fixes
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PublishedAt) {
		toSerialize["publishedAt"] = o.PublishedAt
	}
	if !IsNil(o.RebootSuggested) {
		toSerialize["rebootSuggested"] = o.RebootSuggested
	}
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	if !IsNil(o.Rpms) {
		toSerialize["rpms"] = o.Rpms
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.ShortCode) {
		toSerialize["shortCode"] = o.ShortCode
	}
	if !IsNil(o.Solution) {
		toSerialize["solution"] = o.Solution
	}
	if !IsNil(o.Synopsis) {
		toSerialize["synopsis"] = o.Synopsis
	}
	if !IsNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAdvisoryRockyAdvisory struct {
	value *AdvisoryRockyAdvisory
	isSet bool
}

func (v NullableAdvisoryRockyAdvisory) Get() *AdvisoryRockyAdvisory {
	return v.value
}

func (v *NullableAdvisoryRockyAdvisory) Set(val *AdvisoryRockyAdvisory) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryRockyAdvisory) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryRockyAdvisory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryRockyAdvisory(val *AdvisoryRockyAdvisory) *NullableAdvisoryRockyAdvisory {
	return &NullableAdvisoryRockyAdvisory{value: val, isSet: true}
}

func (v NullableAdvisoryRockyAdvisory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryRockyAdvisory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


