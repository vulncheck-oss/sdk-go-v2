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

// checks if the AdvisoryPyPAAdvisory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryPyPAAdvisory{}

// AdvisoryPyPAAdvisory struct for AdvisoryPyPAAdvisory
type AdvisoryPyPAAdvisory struct {
	// ID is the PYSEC- identifier
	AdvisoryId *string `json:"advisory_id,omitempty"`
	// Affected will list out the vulnerable versions.
	Affected []AdvisoryPyPAAffected `json:"affected,omitempty"`
	// Aliases are other identifiers that refer to this, such as a CVE
	Aliases []string `json:"aliases,omitempty"`
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	// Details discuss the vulnerability information
	Details *string `json:"details,omitempty"`
	// Modified is non-zero Time if entry was updated
	Modified *string `json:"modified,omitempty"`
	// Published is the datetime when this was released
	Published *string `json:"published,omitempty"`
	// References are links to more detailed advisories, fixes, etc.
	References []AdvisoryPyPAReference `json:"references,omitempty"`
	// WasWD indicates if the advisory was withdrawn or not.
	WasWithdrawn *bool `json:"was_withdrawn,omitempty"`
	// Withdrawn is non-zero if this advisory has been withdrawn
	Withdrawn *string `json:"withdrawn,omitempty"`
}

// NewAdvisoryPyPAAdvisory instantiates a new AdvisoryPyPAAdvisory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryPyPAAdvisory() *AdvisoryPyPAAdvisory {
	this := AdvisoryPyPAAdvisory{}
	return &this
}

// NewAdvisoryPyPAAdvisoryWithDefaults instantiates a new AdvisoryPyPAAdvisory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryPyPAAdvisoryWithDefaults() *AdvisoryPyPAAdvisory {
	this := AdvisoryPyPAAdvisory{}
	return &this
}

// GetAdvisoryId returns the AdvisoryId field value if set, zero value otherwise.
func (o *AdvisoryPyPAAdvisory) GetAdvisoryId() string {
	if o == nil || IsNil(o.AdvisoryId) {
		var ret string
		return ret
	}
	return *o.AdvisoryId
}

// GetAdvisoryIdOk returns a tuple with the AdvisoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPyPAAdvisory) GetAdvisoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdvisoryId) {
		return nil, false
	}
	return o.AdvisoryId, true
}

// HasAdvisoryId returns a boolean if a field has been set.
func (o *AdvisoryPyPAAdvisory) HasAdvisoryId() bool {
	if o != nil && !IsNil(o.AdvisoryId) {
		return true
	}

	return false
}

// SetAdvisoryId gets a reference to the given string and assigns it to the AdvisoryId field.
func (o *AdvisoryPyPAAdvisory) SetAdvisoryId(v string) {
	o.AdvisoryId = &v
}

// GetAffected returns the Affected field value if set, zero value otherwise.
func (o *AdvisoryPyPAAdvisory) GetAffected() []AdvisoryPyPAAffected {
	if o == nil || IsNil(o.Affected) {
		var ret []AdvisoryPyPAAffected
		return ret
	}
	return o.Affected
}

// GetAffectedOk returns a tuple with the Affected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPyPAAdvisory) GetAffectedOk() ([]AdvisoryPyPAAffected, bool) {
	if o == nil || IsNil(o.Affected) {
		return nil, false
	}
	return o.Affected, true
}

// HasAffected returns a boolean if a field has been set.
func (o *AdvisoryPyPAAdvisory) HasAffected() bool {
	if o != nil && !IsNil(o.Affected) {
		return true
	}

	return false
}

// SetAffected gets a reference to the given []AdvisoryPyPAAffected and assigns it to the Affected field.
func (o *AdvisoryPyPAAdvisory) SetAffected(v []AdvisoryPyPAAffected) {
	o.Affected = v
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *AdvisoryPyPAAdvisory) GetAliases() []string {
	if o == nil || IsNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPyPAAdvisory) GetAliasesOk() ([]string, bool) {
	if o == nil || IsNil(o.Aliases) {
		return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *AdvisoryPyPAAdvisory) HasAliases() bool {
	if o != nil && !IsNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *AdvisoryPyPAAdvisory) SetAliases(v []string) {
	o.Aliases = v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryPyPAAdvisory) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPyPAAdvisory) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryPyPAAdvisory) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryPyPAAdvisory) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryPyPAAdvisory) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPyPAAdvisory) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryPyPAAdvisory) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryPyPAAdvisory) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *AdvisoryPyPAAdvisory) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPyPAAdvisory) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *AdvisoryPyPAAdvisory) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *AdvisoryPyPAAdvisory) SetDetails(v string) {
	o.Details = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *AdvisoryPyPAAdvisory) GetModified() string {
	if o == nil || IsNil(o.Modified) {
		var ret string
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPyPAAdvisory) GetModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *AdvisoryPyPAAdvisory) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given string and assigns it to the Modified field.
func (o *AdvisoryPyPAAdvisory) SetModified(v string) {
	o.Modified = &v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *AdvisoryPyPAAdvisory) GetPublished() string {
	if o == nil || IsNil(o.Published) {
		var ret string
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPyPAAdvisory) GetPublishedOk() (*string, bool) {
	if o == nil || IsNil(o.Published) {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *AdvisoryPyPAAdvisory) HasPublished() bool {
	if o != nil && !IsNil(o.Published) {
		return true
	}

	return false
}

// SetPublished gets a reference to the given string and assigns it to the Published field.
func (o *AdvisoryPyPAAdvisory) SetPublished(v string) {
	o.Published = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *AdvisoryPyPAAdvisory) GetReferences() []AdvisoryPyPAReference {
	if o == nil || IsNil(o.References) {
		var ret []AdvisoryPyPAReference
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPyPAAdvisory) GetReferencesOk() ([]AdvisoryPyPAReference, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *AdvisoryPyPAAdvisory) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []AdvisoryPyPAReference and assigns it to the References field.
func (o *AdvisoryPyPAAdvisory) SetReferences(v []AdvisoryPyPAReference) {
	o.References = v
}

// GetWasWithdrawn returns the WasWithdrawn field value if set, zero value otherwise.
func (o *AdvisoryPyPAAdvisory) GetWasWithdrawn() bool {
	if o == nil || IsNil(o.WasWithdrawn) {
		var ret bool
		return ret
	}
	return *o.WasWithdrawn
}

// GetWasWithdrawnOk returns a tuple with the WasWithdrawn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPyPAAdvisory) GetWasWithdrawnOk() (*bool, bool) {
	if o == nil || IsNil(o.WasWithdrawn) {
		return nil, false
	}
	return o.WasWithdrawn, true
}

// HasWasWithdrawn returns a boolean if a field has been set.
func (o *AdvisoryPyPAAdvisory) HasWasWithdrawn() bool {
	if o != nil && !IsNil(o.WasWithdrawn) {
		return true
	}

	return false
}

// SetWasWithdrawn gets a reference to the given bool and assigns it to the WasWithdrawn field.
func (o *AdvisoryPyPAAdvisory) SetWasWithdrawn(v bool) {
	o.WasWithdrawn = &v
}

// GetWithdrawn returns the Withdrawn field value if set, zero value otherwise.
func (o *AdvisoryPyPAAdvisory) GetWithdrawn() string {
	if o == nil || IsNil(o.Withdrawn) {
		var ret string
		return ret
	}
	return *o.Withdrawn
}

// GetWithdrawnOk returns a tuple with the Withdrawn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPyPAAdvisory) GetWithdrawnOk() (*string, bool) {
	if o == nil || IsNil(o.Withdrawn) {
		return nil, false
	}
	return o.Withdrawn, true
}

// HasWithdrawn returns a boolean if a field has been set.
func (o *AdvisoryPyPAAdvisory) HasWithdrawn() bool {
	if o != nil && !IsNil(o.Withdrawn) {
		return true
	}

	return false
}

// SetWithdrawn gets a reference to the given string and assigns it to the Withdrawn field.
func (o *AdvisoryPyPAAdvisory) SetWithdrawn(v string) {
	o.Withdrawn = &v
}

func (o AdvisoryPyPAAdvisory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryPyPAAdvisory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdvisoryId) {
		toSerialize["advisory_id"] = o.AdvisoryId
	}
	if !IsNil(o.Affected) {
		toSerialize["affected"] = o.Affected
	}
	if !IsNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
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
	if !IsNil(o.WasWithdrawn) {
		toSerialize["was_withdrawn"] = o.WasWithdrawn
	}
	if !IsNil(o.Withdrawn) {
		toSerialize["withdrawn"] = o.Withdrawn
	}
	return toSerialize, nil
}

type NullableAdvisoryPyPAAdvisory struct {
	value *AdvisoryPyPAAdvisory
	isSet bool
}

func (v NullableAdvisoryPyPAAdvisory) Get() *AdvisoryPyPAAdvisory {
	return v.value
}

func (v *NullableAdvisoryPyPAAdvisory) Set(val *AdvisoryPyPAAdvisory) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryPyPAAdvisory) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryPyPAAdvisory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryPyPAAdvisory(val *AdvisoryPyPAAdvisory) *NullableAdvisoryPyPAAdvisory {
	return &NullableAdvisoryPyPAAdvisory{value: val, isSet: true}
}

func (v NullableAdvisoryPyPAAdvisory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryPyPAAdvisory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


