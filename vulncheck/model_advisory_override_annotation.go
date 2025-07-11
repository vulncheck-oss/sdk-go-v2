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

// checks if the AdvisoryOverrideAnnotation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryOverrideAnnotation{}

// AdvisoryOverrideAnnotation struct for AdvisoryOverrideAnnotation
type AdvisoryOverrideAnnotation struct {
	CveId *string `json:"cve_id,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Snapshot *string `json:"snapshot,omitempty"`
	TriageNotes *AdvisoryTriageNotes `json:"triage_notes,omitempty"`
}

// NewAdvisoryOverrideAnnotation instantiates a new AdvisoryOverrideAnnotation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryOverrideAnnotation() *AdvisoryOverrideAnnotation {
	this := AdvisoryOverrideAnnotation{}
	return &this
}

// NewAdvisoryOverrideAnnotationWithDefaults instantiates a new AdvisoryOverrideAnnotation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryOverrideAnnotationWithDefaults() *AdvisoryOverrideAnnotation {
	this := AdvisoryOverrideAnnotation{}
	return &this
}

// GetCveId returns the CveId field value if set, zero value otherwise.
func (o *AdvisoryOverrideAnnotation) GetCveId() string {
	if o == nil || IsNil(o.CveId) {
		var ret string
		return ret
	}
	return *o.CveId
}

// GetCveIdOk returns a tuple with the CveId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOverrideAnnotation) GetCveIdOk() (*string, bool) {
	if o == nil || IsNil(o.CveId) {
		return nil, false
	}
	return o.CveId, true
}

// HasCveId returns a boolean if a field has been set.
func (o *AdvisoryOverrideAnnotation) HasCveId() bool {
	if o != nil && !IsNil(o.CveId) {
		return true
	}

	return false
}

// SetCveId gets a reference to the given string and assigns it to the CveId field.
func (o *AdvisoryOverrideAnnotation) SetCveId(v string) {
	o.CveId = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AdvisoryOverrideAnnotation) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOverrideAnnotation) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AdvisoryOverrideAnnotation) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AdvisoryOverrideAnnotation) SetReason(v string) {
	o.Reason = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *AdvisoryOverrideAnnotation) GetSnapshot() string {
	if o == nil || IsNil(o.Snapshot) {
		var ret string
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOverrideAnnotation) GetSnapshotOk() (*string, bool) {
	if o == nil || IsNil(o.Snapshot) {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *AdvisoryOverrideAnnotation) HasSnapshot() bool {
	if o != nil && !IsNil(o.Snapshot) {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given string and assigns it to the Snapshot field.
func (o *AdvisoryOverrideAnnotation) SetSnapshot(v string) {
	o.Snapshot = &v
}

// GetTriageNotes returns the TriageNotes field value if set, zero value otherwise.
func (o *AdvisoryOverrideAnnotation) GetTriageNotes() AdvisoryTriageNotes {
	if o == nil || IsNil(o.TriageNotes) {
		var ret AdvisoryTriageNotes
		return ret
	}
	return *o.TriageNotes
}

// GetTriageNotesOk returns a tuple with the TriageNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOverrideAnnotation) GetTriageNotesOk() (*AdvisoryTriageNotes, bool) {
	if o == nil || IsNil(o.TriageNotes) {
		return nil, false
	}
	return o.TriageNotes, true
}

// HasTriageNotes returns a boolean if a field has been set.
func (o *AdvisoryOverrideAnnotation) HasTriageNotes() bool {
	if o != nil && !IsNil(o.TriageNotes) {
		return true
	}

	return false
}

// SetTriageNotes gets a reference to the given AdvisoryTriageNotes and assigns it to the TriageNotes field.
func (o *AdvisoryOverrideAnnotation) SetTriageNotes(v AdvisoryTriageNotes) {
	o.TriageNotes = &v
}

func (o AdvisoryOverrideAnnotation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryOverrideAnnotation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CveId) {
		toSerialize["cve_id"] = o.CveId
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Snapshot) {
		toSerialize["snapshot"] = o.Snapshot
	}
	if !IsNil(o.TriageNotes) {
		toSerialize["triage_notes"] = o.TriageNotes
	}
	return toSerialize, nil
}

type NullableAdvisoryOverrideAnnotation struct {
	value *AdvisoryOverrideAnnotation
	isSet bool
}

func (v NullableAdvisoryOverrideAnnotation) Get() *AdvisoryOverrideAnnotation {
	return v.value
}

func (v *NullableAdvisoryOverrideAnnotation) Set(val *AdvisoryOverrideAnnotation) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryOverrideAnnotation) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryOverrideAnnotation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryOverrideAnnotation(val *AdvisoryOverrideAnnotation) *NullableAdvisoryOverrideAnnotation {
	return &NullableAdvisoryOverrideAnnotation{value: val, isSet: true}
}

func (v NullableAdvisoryOverrideAnnotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryOverrideAnnotation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


