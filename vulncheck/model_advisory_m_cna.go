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

// checks if the AdvisoryMCna type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryMCna{}

// AdvisoryMCna struct for AdvisoryMCna
type AdvisoryMCna struct {
	Affected []AdvisoryMAffected `json:"affected,omitempty"`
	CpeApplicability []AdvisoryMCPEApplicability `json:"cpeApplicability,omitempty"`
	Credits []AdvisoryCredit `json:"credits,omitempty"`
	Descriptions []AdvisoryMDescriptions `json:"descriptions,omitempty"`
	Impacts []AdvisoryImpact `json:"impacts,omitempty"`
	Metrics []AdvisoryMetric `json:"metrics,omitempty"`
	ProblemTypes []AdvisoryMProblemTypes `json:"problemTypes,omitempty"`
	ProviderMetadata *AdvisoryMProviderMetadata `json:"providerMetadata,omitempty"`
	References []AdvisoryMReference `json:"references,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Timeline []AdvisoryTimeline `json:"timeline,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewAdvisoryMCna instantiates a new AdvisoryMCna object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryMCna() *AdvisoryMCna {
	this := AdvisoryMCna{}
	return &this
}

// NewAdvisoryMCnaWithDefaults instantiates a new AdvisoryMCna object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryMCnaWithDefaults() *AdvisoryMCna {
	this := AdvisoryMCna{}
	return &this
}

// GetAffected returns the Affected field value if set, zero value otherwise.
func (o *AdvisoryMCna) GetAffected() []AdvisoryMAffected {
	if o == nil || IsNil(o.Affected) {
		var ret []AdvisoryMAffected
		return ret
	}
	return o.Affected
}

// GetAffectedOk returns a tuple with the Affected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMCna) GetAffectedOk() ([]AdvisoryMAffected, bool) {
	if o == nil || IsNil(o.Affected) {
		return nil, false
	}
	return o.Affected, true
}

// HasAffected returns a boolean if a field has been set.
func (o *AdvisoryMCna) HasAffected() bool {
	if o != nil && !IsNil(o.Affected) {
		return true
	}

	return false
}

// SetAffected gets a reference to the given []AdvisoryMAffected and assigns it to the Affected field.
func (o *AdvisoryMCna) SetAffected(v []AdvisoryMAffected) {
	o.Affected = v
}

// GetCpeApplicability returns the CpeApplicability field value if set, zero value otherwise.
func (o *AdvisoryMCna) GetCpeApplicability() []AdvisoryMCPEApplicability {
	if o == nil || IsNil(o.CpeApplicability) {
		var ret []AdvisoryMCPEApplicability
		return ret
	}
	return o.CpeApplicability
}

// GetCpeApplicabilityOk returns a tuple with the CpeApplicability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMCna) GetCpeApplicabilityOk() ([]AdvisoryMCPEApplicability, bool) {
	if o == nil || IsNil(o.CpeApplicability) {
		return nil, false
	}
	return o.CpeApplicability, true
}

// HasCpeApplicability returns a boolean if a field has been set.
func (o *AdvisoryMCna) HasCpeApplicability() bool {
	if o != nil && !IsNil(o.CpeApplicability) {
		return true
	}

	return false
}

// SetCpeApplicability gets a reference to the given []AdvisoryMCPEApplicability and assigns it to the CpeApplicability field.
func (o *AdvisoryMCna) SetCpeApplicability(v []AdvisoryMCPEApplicability) {
	o.CpeApplicability = v
}

// GetCredits returns the Credits field value if set, zero value otherwise.
func (o *AdvisoryMCna) GetCredits() []AdvisoryCredit {
	if o == nil || IsNil(o.Credits) {
		var ret []AdvisoryCredit
		return ret
	}
	return o.Credits
}

// GetCreditsOk returns a tuple with the Credits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMCna) GetCreditsOk() ([]AdvisoryCredit, bool) {
	if o == nil || IsNil(o.Credits) {
		return nil, false
	}
	return o.Credits, true
}

// HasCredits returns a boolean if a field has been set.
func (o *AdvisoryMCna) HasCredits() bool {
	if o != nil && !IsNil(o.Credits) {
		return true
	}

	return false
}

// SetCredits gets a reference to the given []AdvisoryCredit and assigns it to the Credits field.
func (o *AdvisoryMCna) SetCredits(v []AdvisoryCredit) {
	o.Credits = v
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *AdvisoryMCna) GetDescriptions() []AdvisoryMDescriptions {
	if o == nil || IsNil(o.Descriptions) {
		var ret []AdvisoryMDescriptions
		return ret
	}
	return o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMCna) GetDescriptionsOk() ([]AdvisoryMDescriptions, bool) {
	if o == nil || IsNil(o.Descriptions) {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *AdvisoryMCna) HasDescriptions() bool {
	if o != nil && !IsNil(o.Descriptions) {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []AdvisoryMDescriptions and assigns it to the Descriptions field.
func (o *AdvisoryMCna) SetDescriptions(v []AdvisoryMDescriptions) {
	o.Descriptions = v
}

// GetImpacts returns the Impacts field value if set, zero value otherwise.
func (o *AdvisoryMCna) GetImpacts() []AdvisoryImpact {
	if o == nil || IsNil(o.Impacts) {
		var ret []AdvisoryImpact
		return ret
	}
	return o.Impacts
}

// GetImpactsOk returns a tuple with the Impacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMCna) GetImpactsOk() ([]AdvisoryImpact, bool) {
	if o == nil || IsNil(o.Impacts) {
		return nil, false
	}
	return o.Impacts, true
}

// HasImpacts returns a boolean if a field has been set.
func (o *AdvisoryMCna) HasImpacts() bool {
	if o != nil && !IsNil(o.Impacts) {
		return true
	}

	return false
}

// SetImpacts gets a reference to the given []AdvisoryImpact and assigns it to the Impacts field.
func (o *AdvisoryMCna) SetImpacts(v []AdvisoryImpact) {
	o.Impacts = v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *AdvisoryMCna) GetMetrics() []AdvisoryMetric {
	if o == nil || IsNil(o.Metrics) {
		var ret []AdvisoryMetric
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMCna) GetMetricsOk() ([]AdvisoryMetric, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *AdvisoryMCna) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given []AdvisoryMetric and assigns it to the Metrics field.
func (o *AdvisoryMCna) SetMetrics(v []AdvisoryMetric) {
	o.Metrics = v
}

// GetProblemTypes returns the ProblemTypes field value if set, zero value otherwise.
func (o *AdvisoryMCna) GetProblemTypes() []AdvisoryMProblemTypes {
	if o == nil || IsNil(o.ProblemTypes) {
		var ret []AdvisoryMProblemTypes
		return ret
	}
	return o.ProblemTypes
}

// GetProblemTypesOk returns a tuple with the ProblemTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMCna) GetProblemTypesOk() ([]AdvisoryMProblemTypes, bool) {
	if o == nil || IsNil(o.ProblemTypes) {
		return nil, false
	}
	return o.ProblemTypes, true
}

// HasProblemTypes returns a boolean if a field has been set.
func (o *AdvisoryMCna) HasProblemTypes() bool {
	if o != nil && !IsNil(o.ProblemTypes) {
		return true
	}

	return false
}

// SetProblemTypes gets a reference to the given []AdvisoryMProblemTypes and assigns it to the ProblemTypes field.
func (o *AdvisoryMCna) SetProblemTypes(v []AdvisoryMProblemTypes) {
	o.ProblemTypes = v
}

// GetProviderMetadata returns the ProviderMetadata field value if set, zero value otherwise.
func (o *AdvisoryMCna) GetProviderMetadata() AdvisoryMProviderMetadata {
	if o == nil || IsNil(o.ProviderMetadata) {
		var ret AdvisoryMProviderMetadata
		return ret
	}
	return *o.ProviderMetadata
}

// GetProviderMetadataOk returns a tuple with the ProviderMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMCna) GetProviderMetadataOk() (*AdvisoryMProviderMetadata, bool) {
	if o == nil || IsNil(o.ProviderMetadata) {
		return nil, false
	}
	return o.ProviderMetadata, true
}

// HasProviderMetadata returns a boolean if a field has been set.
func (o *AdvisoryMCna) HasProviderMetadata() bool {
	if o != nil && !IsNil(o.ProviderMetadata) {
		return true
	}

	return false
}

// SetProviderMetadata gets a reference to the given AdvisoryMProviderMetadata and assigns it to the ProviderMetadata field.
func (o *AdvisoryMCna) SetProviderMetadata(v AdvisoryMProviderMetadata) {
	o.ProviderMetadata = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *AdvisoryMCna) GetReferences() []AdvisoryMReference {
	if o == nil || IsNil(o.References) {
		var ret []AdvisoryMReference
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMCna) GetReferencesOk() ([]AdvisoryMReference, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *AdvisoryMCna) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []AdvisoryMReference and assigns it to the References field.
func (o *AdvisoryMCna) SetReferences(v []AdvisoryMReference) {
	o.References = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AdvisoryMCna) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMCna) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AdvisoryMCna) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AdvisoryMCna) SetTags(v []string) {
	o.Tags = v
}

// GetTimeline returns the Timeline field value if set, zero value otherwise.
func (o *AdvisoryMCna) GetTimeline() []AdvisoryTimeline {
	if o == nil || IsNil(o.Timeline) {
		var ret []AdvisoryTimeline
		return ret
	}
	return o.Timeline
}

// GetTimelineOk returns a tuple with the Timeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMCna) GetTimelineOk() ([]AdvisoryTimeline, bool) {
	if o == nil || IsNil(o.Timeline) {
		return nil, false
	}
	return o.Timeline, true
}

// HasTimeline returns a boolean if a field has been set.
func (o *AdvisoryMCna) HasTimeline() bool {
	if o != nil && !IsNil(o.Timeline) {
		return true
	}

	return false
}

// SetTimeline gets a reference to the given []AdvisoryTimeline and assigns it to the Timeline field.
func (o *AdvisoryMCna) SetTimeline(v []AdvisoryTimeline) {
	o.Timeline = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryMCna) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryMCna) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryMCna) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryMCna) SetTitle(v string) {
	o.Title = &v
}

func (o AdvisoryMCna) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryMCna) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Affected) {
		toSerialize["affected"] = o.Affected
	}
	if !IsNil(o.CpeApplicability) {
		toSerialize["cpeApplicability"] = o.CpeApplicability
	}
	if !IsNil(o.Credits) {
		toSerialize["credits"] = o.Credits
	}
	if !IsNil(o.Descriptions) {
		toSerialize["descriptions"] = o.Descriptions
	}
	if !IsNil(o.Impacts) {
		toSerialize["impacts"] = o.Impacts
	}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	if !IsNil(o.ProblemTypes) {
		toSerialize["problemTypes"] = o.ProblemTypes
	}
	if !IsNil(o.ProviderMetadata) {
		toSerialize["providerMetadata"] = o.ProviderMetadata
	}
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Timeline) {
		toSerialize["timeline"] = o.Timeline
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableAdvisoryMCna struct {
	value *AdvisoryMCna
	isSet bool
}

func (v NullableAdvisoryMCna) Get() *AdvisoryMCna {
	return v.value
}

func (v *NullableAdvisoryMCna) Set(val *AdvisoryMCna) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryMCna) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryMCna) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryMCna(val *AdvisoryMCna) *NullableAdvisoryMCna {
	return &NullableAdvisoryMCna{value: val, isSet: true}
}

func (v NullableAdvisoryMCna) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryMCna) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


