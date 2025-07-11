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

// checks if the AdvisoryGoVulnJSON type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryGoVulnJSON{}

// AdvisoryGoVulnJSON struct for AdvisoryGoVulnJSON
type AdvisoryGoVulnJSON struct {
	AdvisoryUrl *string `json:"advisory_url,omitempty"`
	Affected []AdvisoryGoVulnAffected `json:"affected,omitempty"`
	Aliases []string `json:"aliases,omitempty"`
	Credits []AdvisoryGoCredits `json:"credits,omitempty"`
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Details *string `json:"details,omitempty"`
	Ghsa []string `json:"ghsa,omitempty"`
	GoAdvisoryId *string `json:"go_advisory_id,omitempty"`
	Modified *string `json:"modified,omitempty"`
	Published *string `json:"published,omitempty"`
	References []AdvisoryGoVulnReference `json:"references,omitempty"`
}

// NewAdvisoryGoVulnJSON instantiates a new AdvisoryGoVulnJSON object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryGoVulnJSON() *AdvisoryGoVulnJSON {
	this := AdvisoryGoVulnJSON{}
	return &this
}

// NewAdvisoryGoVulnJSONWithDefaults instantiates a new AdvisoryGoVulnJSON object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryGoVulnJSONWithDefaults() *AdvisoryGoVulnJSON {
	this := AdvisoryGoVulnJSON{}
	return &this
}

// GetAdvisoryUrl returns the AdvisoryUrl field value if set, zero value otherwise.
func (o *AdvisoryGoVulnJSON) GetAdvisoryUrl() string {
	if o == nil || IsNil(o.AdvisoryUrl) {
		var ret string
		return ret
	}
	return *o.AdvisoryUrl
}

// GetAdvisoryUrlOk returns a tuple with the AdvisoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGoVulnJSON) GetAdvisoryUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AdvisoryUrl) {
		return nil, false
	}
	return o.AdvisoryUrl, true
}

// HasAdvisoryUrl returns a boolean if a field has been set.
func (o *AdvisoryGoVulnJSON) HasAdvisoryUrl() bool {
	if o != nil && !IsNil(o.AdvisoryUrl) {
		return true
	}

	return false
}

// SetAdvisoryUrl gets a reference to the given string and assigns it to the AdvisoryUrl field.
func (o *AdvisoryGoVulnJSON) SetAdvisoryUrl(v string) {
	o.AdvisoryUrl = &v
}

// GetAffected returns the Affected field value if set, zero value otherwise.
func (o *AdvisoryGoVulnJSON) GetAffected() []AdvisoryGoVulnAffected {
	if o == nil || IsNil(o.Affected) {
		var ret []AdvisoryGoVulnAffected
		return ret
	}
	return o.Affected
}

// GetAffectedOk returns a tuple with the Affected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGoVulnJSON) GetAffectedOk() ([]AdvisoryGoVulnAffected, bool) {
	if o == nil || IsNil(o.Affected) {
		return nil, false
	}
	return o.Affected, true
}

// HasAffected returns a boolean if a field has been set.
func (o *AdvisoryGoVulnJSON) HasAffected() bool {
	if o != nil && !IsNil(o.Affected) {
		return true
	}

	return false
}

// SetAffected gets a reference to the given []AdvisoryGoVulnAffected and assigns it to the Affected field.
func (o *AdvisoryGoVulnJSON) SetAffected(v []AdvisoryGoVulnAffected) {
	o.Affected = v
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *AdvisoryGoVulnJSON) GetAliases() []string {
	if o == nil || IsNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGoVulnJSON) GetAliasesOk() ([]string, bool) {
	if o == nil || IsNil(o.Aliases) {
		return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *AdvisoryGoVulnJSON) HasAliases() bool {
	if o != nil && !IsNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *AdvisoryGoVulnJSON) SetAliases(v []string) {
	o.Aliases = v
}

// GetCredits returns the Credits field value if set, zero value otherwise.
func (o *AdvisoryGoVulnJSON) GetCredits() []AdvisoryGoCredits {
	if o == nil || IsNil(o.Credits) {
		var ret []AdvisoryGoCredits
		return ret
	}
	return o.Credits
}

// GetCreditsOk returns a tuple with the Credits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGoVulnJSON) GetCreditsOk() ([]AdvisoryGoCredits, bool) {
	if o == nil || IsNil(o.Credits) {
		return nil, false
	}
	return o.Credits, true
}

// HasCredits returns a boolean if a field has been set.
func (o *AdvisoryGoVulnJSON) HasCredits() bool {
	if o != nil && !IsNil(o.Credits) {
		return true
	}

	return false
}

// SetCredits gets a reference to the given []AdvisoryGoCredits and assigns it to the Credits field.
func (o *AdvisoryGoVulnJSON) SetCredits(v []AdvisoryGoCredits) {
	o.Credits = v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryGoVulnJSON) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGoVulnJSON) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryGoVulnJSON) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryGoVulnJSON) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryGoVulnJSON) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGoVulnJSON) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryGoVulnJSON) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryGoVulnJSON) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *AdvisoryGoVulnJSON) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGoVulnJSON) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *AdvisoryGoVulnJSON) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *AdvisoryGoVulnJSON) SetDetails(v string) {
	o.Details = &v
}

// GetGhsa returns the Ghsa field value if set, zero value otherwise.
func (o *AdvisoryGoVulnJSON) GetGhsa() []string {
	if o == nil || IsNil(o.Ghsa) {
		var ret []string
		return ret
	}
	return o.Ghsa
}

// GetGhsaOk returns a tuple with the Ghsa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGoVulnJSON) GetGhsaOk() ([]string, bool) {
	if o == nil || IsNil(o.Ghsa) {
		return nil, false
	}
	return o.Ghsa, true
}

// HasGhsa returns a boolean if a field has been set.
func (o *AdvisoryGoVulnJSON) HasGhsa() bool {
	if o != nil && !IsNil(o.Ghsa) {
		return true
	}

	return false
}

// SetGhsa gets a reference to the given []string and assigns it to the Ghsa field.
func (o *AdvisoryGoVulnJSON) SetGhsa(v []string) {
	o.Ghsa = v
}

// GetGoAdvisoryId returns the GoAdvisoryId field value if set, zero value otherwise.
func (o *AdvisoryGoVulnJSON) GetGoAdvisoryId() string {
	if o == nil || IsNil(o.GoAdvisoryId) {
		var ret string
		return ret
	}
	return *o.GoAdvisoryId
}

// GetGoAdvisoryIdOk returns a tuple with the GoAdvisoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGoVulnJSON) GetGoAdvisoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.GoAdvisoryId) {
		return nil, false
	}
	return o.GoAdvisoryId, true
}

// HasGoAdvisoryId returns a boolean if a field has been set.
func (o *AdvisoryGoVulnJSON) HasGoAdvisoryId() bool {
	if o != nil && !IsNil(o.GoAdvisoryId) {
		return true
	}

	return false
}

// SetGoAdvisoryId gets a reference to the given string and assigns it to the GoAdvisoryId field.
func (o *AdvisoryGoVulnJSON) SetGoAdvisoryId(v string) {
	o.GoAdvisoryId = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *AdvisoryGoVulnJSON) GetModified() string {
	if o == nil || IsNil(o.Modified) {
		var ret string
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGoVulnJSON) GetModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *AdvisoryGoVulnJSON) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given string and assigns it to the Modified field.
func (o *AdvisoryGoVulnJSON) SetModified(v string) {
	o.Modified = &v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *AdvisoryGoVulnJSON) GetPublished() string {
	if o == nil || IsNil(o.Published) {
		var ret string
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGoVulnJSON) GetPublishedOk() (*string, bool) {
	if o == nil || IsNil(o.Published) {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *AdvisoryGoVulnJSON) HasPublished() bool {
	if o != nil && !IsNil(o.Published) {
		return true
	}

	return false
}

// SetPublished gets a reference to the given string and assigns it to the Published field.
func (o *AdvisoryGoVulnJSON) SetPublished(v string) {
	o.Published = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *AdvisoryGoVulnJSON) GetReferences() []AdvisoryGoVulnReference {
	if o == nil || IsNil(o.References) {
		var ret []AdvisoryGoVulnReference
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGoVulnJSON) GetReferencesOk() ([]AdvisoryGoVulnReference, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *AdvisoryGoVulnJSON) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []AdvisoryGoVulnReference and assigns it to the References field.
func (o *AdvisoryGoVulnJSON) SetReferences(v []AdvisoryGoVulnReference) {
	o.References = v
}

func (o AdvisoryGoVulnJSON) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryGoVulnJSON) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdvisoryUrl) {
		toSerialize["advisory_url"] = o.AdvisoryUrl
	}
	if !IsNil(o.Affected) {
		toSerialize["affected"] = o.Affected
	}
	if !IsNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	if !IsNil(o.Credits) {
		toSerialize["credits"] = o.Credits
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
	if !IsNil(o.Ghsa) {
		toSerialize["ghsa"] = o.Ghsa
	}
	if !IsNil(o.GoAdvisoryId) {
		toSerialize["go_advisory_id"] = o.GoAdvisoryId
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
	return toSerialize, nil
}

type NullableAdvisoryGoVulnJSON struct {
	value *AdvisoryGoVulnJSON
	isSet bool
}

func (v NullableAdvisoryGoVulnJSON) Get() *AdvisoryGoVulnJSON {
	return v.value
}

func (v *NullableAdvisoryGoVulnJSON) Set(val *AdvisoryGoVulnJSON) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryGoVulnJSON) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryGoVulnJSON) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryGoVulnJSON(val *AdvisoryGoVulnJSON) *NullableAdvisoryGoVulnJSON {
	return &NullableAdvisoryGoVulnJSON{value: val, isSet: true}
}

func (v NullableAdvisoryGoVulnJSON) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryGoVulnJSON) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


