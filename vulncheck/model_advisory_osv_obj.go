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

// checks if the AdvisoryOSVObj type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryOSVObj{}

// AdvisoryOSVObj struct for AdvisoryOSVObj
type AdvisoryOSVObj struct {
	// collection based on https://ossf.github.io/osv-schema/
	Affected []AdvisoryAffected `json:"affected,omitempty"`
	Aliases []string `json:"aliases,omitempty"`
	Details *string `json:"details,omitempty"`
	Id *string `json:"id,omitempty"`
	Modified *string `json:"modified,omitempty"`
	Published *string `json:"published,omitempty"`
	References []AdvisoryOSVReference `json:"references,omitempty"`
	Related []string `json:"related,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Withdrawn *string `json:"withdrawn,omitempty"`
}

// NewAdvisoryOSVObj instantiates a new AdvisoryOSVObj object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryOSVObj() *AdvisoryOSVObj {
	this := AdvisoryOSVObj{}
	return &this
}

// NewAdvisoryOSVObjWithDefaults instantiates a new AdvisoryOSVObj object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryOSVObjWithDefaults() *AdvisoryOSVObj {
	this := AdvisoryOSVObj{}
	return &this
}

// GetAffected returns the Affected field value if set, zero value otherwise.
func (o *AdvisoryOSVObj) GetAffected() []AdvisoryAffected {
	if o == nil || IsNil(o.Affected) {
		var ret []AdvisoryAffected
		return ret
	}
	return o.Affected
}

// GetAffectedOk returns a tuple with the Affected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOSVObj) GetAffectedOk() ([]AdvisoryAffected, bool) {
	if o == nil || IsNil(o.Affected) {
		return nil, false
	}
	return o.Affected, true
}

// HasAffected returns a boolean if a field has been set.
func (o *AdvisoryOSVObj) HasAffected() bool {
	if o != nil && !IsNil(o.Affected) {
		return true
	}

	return false
}

// SetAffected gets a reference to the given []AdvisoryAffected and assigns it to the Affected field.
func (o *AdvisoryOSVObj) SetAffected(v []AdvisoryAffected) {
	o.Affected = v
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *AdvisoryOSVObj) GetAliases() []string {
	if o == nil || IsNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOSVObj) GetAliasesOk() ([]string, bool) {
	if o == nil || IsNil(o.Aliases) {
		return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *AdvisoryOSVObj) HasAliases() bool {
	if o != nil && !IsNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *AdvisoryOSVObj) SetAliases(v []string) {
	o.Aliases = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *AdvisoryOSVObj) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOSVObj) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *AdvisoryOSVObj) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *AdvisoryOSVObj) SetDetails(v string) {
	o.Details = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdvisoryOSVObj) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOSVObj) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdvisoryOSVObj) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdvisoryOSVObj) SetId(v string) {
	o.Id = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *AdvisoryOSVObj) GetModified() string {
	if o == nil || IsNil(o.Modified) {
		var ret string
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOSVObj) GetModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *AdvisoryOSVObj) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given string and assigns it to the Modified field.
func (o *AdvisoryOSVObj) SetModified(v string) {
	o.Modified = &v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *AdvisoryOSVObj) GetPublished() string {
	if o == nil || IsNil(o.Published) {
		var ret string
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOSVObj) GetPublishedOk() (*string, bool) {
	if o == nil || IsNil(o.Published) {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *AdvisoryOSVObj) HasPublished() bool {
	if o != nil && !IsNil(o.Published) {
		return true
	}

	return false
}

// SetPublished gets a reference to the given string and assigns it to the Published field.
func (o *AdvisoryOSVObj) SetPublished(v string) {
	o.Published = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *AdvisoryOSVObj) GetReferences() []AdvisoryOSVReference {
	if o == nil || IsNil(o.References) {
		var ret []AdvisoryOSVReference
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOSVObj) GetReferencesOk() ([]AdvisoryOSVReference, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *AdvisoryOSVObj) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []AdvisoryOSVReference and assigns it to the References field.
func (o *AdvisoryOSVObj) SetReferences(v []AdvisoryOSVReference) {
	o.References = v
}

// GetRelated returns the Related field value if set, zero value otherwise.
func (o *AdvisoryOSVObj) GetRelated() []string {
	if o == nil || IsNil(o.Related) {
		var ret []string
		return ret
	}
	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOSVObj) GetRelatedOk() ([]string, bool) {
	if o == nil || IsNil(o.Related) {
		return nil, false
	}
	return o.Related, true
}

// HasRelated returns a boolean if a field has been set.
func (o *AdvisoryOSVObj) HasRelated() bool {
	if o != nil && !IsNil(o.Related) {
		return true
	}

	return false
}

// SetRelated gets a reference to the given []string and assigns it to the Related field.
func (o *AdvisoryOSVObj) SetRelated(v []string) {
	o.Related = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AdvisoryOSVObj) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOSVObj) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AdvisoryOSVObj) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AdvisoryOSVObj) SetSummary(v string) {
	o.Summary = &v
}

// GetWithdrawn returns the Withdrawn field value if set, zero value otherwise.
func (o *AdvisoryOSVObj) GetWithdrawn() string {
	if o == nil || IsNil(o.Withdrawn) {
		var ret string
		return ret
	}
	return *o.Withdrawn
}

// GetWithdrawnOk returns a tuple with the Withdrawn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryOSVObj) GetWithdrawnOk() (*string, bool) {
	if o == nil || IsNil(o.Withdrawn) {
		return nil, false
	}
	return o.Withdrawn, true
}

// HasWithdrawn returns a boolean if a field has been set.
func (o *AdvisoryOSVObj) HasWithdrawn() bool {
	if o != nil && !IsNil(o.Withdrawn) {
		return true
	}

	return false
}

// SetWithdrawn gets a reference to the given string and assigns it to the Withdrawn field.
func (o *AdvisoryOSVObj) SetWithdrawn(v string) {
	o.Withdrawn = &v
}

func (o AdvisoryOSVObj) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryOSVObj) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Affected) {
		toSerialize["affected"] = o.Affected
	}
	if !IsNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.Published) {
		toSerialize["published"] = o.Published
	}
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	if !IsNil(o.Related) {
		toSerialize["related"] = o.Related
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Withdrawn) {
		toSerialize["withdrawn"] = o.Withdrawn
	}
	return toSerialize, nil
}

type NullableAdvisoryOSVObj struct {
	value *AdvisoryOSVObj
	isSet bool
}

func (v NullableAdvisoryOSVObj) Get() *AdvisoryOSVObj {
	return v.value
}

func (v *NullableAdvisoryOSVObj) Set(val *AdvisoryOSVObj) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryOSVObj) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryOSVObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryOSVObj(val *AdvisoryOSVObj) *NullableAdvisoryOSVObj {
	return &NullableAdvisoryOSVObj{value: val, isSet: true}
}

func (v NullableAdvisoryOSVObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryOSVObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


