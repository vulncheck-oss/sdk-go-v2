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

// checks if the AdvisoryGHAdvisoryJSONLean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryGHAdvisoryJSONLean{}

// AdvisoryGHAdvisoryJSONLean struct for AdvisoryGHAdvisoryJSONLean
type AdvisoryGHAdvisoryJSONLean struct {
	Classification *string `json:"classification,omitempty"`
	Cve []string `json:"cve,omitempty"`
	Cvss *AdvisoryGHCvss `json:"cvss,omitempty"`
	Cwes *AdvisoryCwes `json:"cwes,omitempty"`
	DatabaseId *int32 `json:"databaseId,omitempty"`
	Description *string `json:"description,omitempty"`
	GhsaId *string `json:"ghsaId,omitempty"`
	Id *string `json:"id,omitempty"`
	Identifiers []AdvisoryGHIdentifier `json:"identifiers,omitempty"`
	NotificationsPermalink *string `json:"notificationsPermalink,omitempty"`
	Origin *string `json:"origin,omitempty"`
	Permalink *string `json:"permalink,omitempty"`
	PublishedAt *string `json:"publishedAt,omitempty"`
	References []AdvisoryGHReference `json:"references,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Summary *string `json:"summary,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	Vulnerabilities *AdvisoryGHVulnerabilities `json:"vulnerabilities,omitempty"`
	WithdrawnAt *string `json:"withdrawnAt,omitempty"`
}

// NewAdvisoryGHAdvisoryJSONLean instantiates a new AdvisoryGHAdvisoryJSONLean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryGHAdvisoryJSONLean() *AdvisoryGHAdvisoryJSONLean {
	this := AdvisoryGHAdvisoryJSONLean{}
	return &this
}

// NewAdvisoryGHAdvisoryJSONLeanWithDefaults instantiates a new AdvisoryGHAdvisoryJSONLean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryGHAdvisoryJSONLeanWithDefaults() *AdvisoryGHAdvisoryJSONLean {
	this := AdvisoryGHAdvisoryJSONLean{}
	return &this
}

// GetClassification returns the Classification field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetClassification() string {
	if o == nil || IsNil(o.Classification) {
		var ret string
		return ret
	}
	return *o.Classification
}

// GetClassificationOk returns a tuple with the Classification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetClassificationOk() (*string, bool) {
	if o == nil || IsNil(o.Classification) {
		return nil, false
	}
	return o.Classification, true
}

// HasClassification returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasClassification() bool {
	if o != nil && !IsNil(o.Classification) {
		return true
	}

	return false
}

// SetClassification gets a reference to the given string and assigns it to the Classification field.
func (o *AdvisoryGHAdvisoryJSONLean) SetClassification(v string) {
	o.Classification = &v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryGHAdvisoryJSONLean) SetCve(v []string) {
	o.Cve = v
}

// GetCvss returns the Cvss field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetCvss() AdvisoryGHCvss {
	if o == nil || IsNil(o.Cvss) {
		var ret AdvisoryGHCvss
		return ret
	}
	return *o.Cvss
}

// GetCvssOk returns a tuple with the Cvss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetCvssOk() (*AdvisoryGHCvss, bool) {
	if o == nil || IsNil(o.Cvss) {
		return nil, false
	}
	return o.Cvss, true
}

// HasCvss returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasCvss() bool {
	if o != nil && !IsNil(o.Cvss) {
		return true
	}

	return false
}

// SetCvss gets a reference to the given AdvisoryGHCvss and assigns it to the Cvss field.
func (o *AdvisoryGHAdvisoryJSONLean) SetCvss(v AdvisoryGHCvss) {
	o.Cvss = &v
}

// GetCwes returns the Cwes field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetCwes() AdvisoryCwes {
	if o == nil || IsNil(o.Cwes) {
		var ret AdvisoryCwes
		return ret
	}
	return *o.Cwes
}

// GetCwesOk returns a tuple with the Cwes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetCwesOk() (*AdvisoryCwes, bool) {
	if o == nil || IsNil(o.Cwes) {
		return nil, false
	}
	return o.Cwes, true
}

// HasCwes returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasCwes() bool {
	if o != nil && !IsNil(o.Cwes) {
		return true
	}

	return false
}

// SetCwes gets a reference to the given AdvisoryCwes and assigns it to the Cwes field.
func (o *AdvisoryGHAdvisoryJSONLean) SetCwes(v AdvisoryCwes) {
	o.Cwes = &v
}

// GetDatabaseId returns the DatabaseId field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetDatabaseId() int32 {
	if o == nil || IsNil(o.DatabaseId) {
		var ret int32
		return ret
	}
	return *o.DatabaseId
}

// GetDatabaseIdOk returns a tuple with the DatabaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetDatabaseIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DatabaseId) {
		return nil, false
	}
	return o.DatabaseId, true
}

// HasDatabaseId returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasDatabaseId() bool {
	if o != nil && !IsNil(o.DatabaseId) {
		return true
	}

	return false
}

// SetDatabaseId gets a reference to the given int32 and assigns it to the DatabaseId field.
func (o *AdvisoryGHAdvisoryJSONLean) SetDatabaseId(v int32) {
	o.DatabaseId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AdvisoryGHAdvisoryJSONLean) SetDescription(v string) {
	o.Description = &v
}

// GetGhsaId returns the GhsaId field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetGhsaId() string {
	if o == nil || IsNil(o.GhsaId) {
		var ret string
		return ret
	}
	return *o.GhsaId
}

// GetGhsaIdOk returns a tuple with the GhsaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetGhsaIdOk() (*string, bool) {
	if o == nil || IsNil(o.GhsaId) {
		return nil, false
	}
	return o.GhsaId, true
}

// HasGhsaId returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasGhsaId() bool {
	if o != nil && !IsNil(o.GhsaId) {
		return true
	}

	return false
}

// SetGhsaId gets a reference to the given string and assigns it to the GhsaId field.
func (o *AdvisoryGHAdvisoryJSONLean) SetGhsaId(v string) {
	o.GhsaId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdvisoryGHAdvisoryJSONLean) SetId(v string) {
	o.Id = &v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetIdentifiers() []AdvisoryGHIdentifier {
	if o == nil || IsNil(o.Identifiers) {
		var ret []AdvisoryGHIdentifier
		return ret
	}
	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetIdentifiersOk() ([]AdvisoryGHIdentifier, bool) {
	if o == nil || IsNil(o.Identifiers) {
		return nil, false
	}
	return o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasIdentifiers() bool {
	if o != nil && !IsNil(o.Identifiers) {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given []AdvisoryGHIdentifier and assigns it to the Identifiers field.
func (o *AdvisoryGHAdvisoryJSONLean) SetIdentifiers(v []AdvisoryGHIdentifier) {
	o.Identifiers = v
}

// GetNotificationsPermalink returns the NotificationsPermalink field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetNotificationsPermalink() string {
	if o == nil || IsNil(o.NotificationsPermalink) {
		var ret string
		return ret
	}
	return *o.NotificationsPermalink
}

// GetNotificationsPermalinkOk returns a tuple with the NotificationsPermalink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetNotificationsPermalinkOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationsPermalink) {
		return nil, false
	}
	return o.NotificationsPermalink, true
}

// HasNotificationsPermalink returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasNotificationsPermalink() bool {
	if o != nil && !IsNil(o.NotificationsPermalink) {
		return true
	}

	return false
}

// SetNotificationsPermalink gets a reference to the given string and assigns it to the NotificationsPermalink field.
func (o *AdvisoryGHAdvisoryJSONLean) SetNotificationsPermalink(v string) {
	o.NotificationsPermalink = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *AdvisoryGHAdvisoryJSONLean) SetOrigin(v string) {
	o.Origin = &v
}

// GetPermalink returns the Permalink field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetPermalink() string {
	if o == nil || IsNil(o.Permalink) {
		var ret string
		return ret
	}
	return *o.Permalink
}

// GetPermalinkOk returns a tuple with the Permalink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetPermalinkOk() (*string, bool) {
	if o == nil || IsNil(o.Permalink) {
		return nil, false
	}
	return o.Permalink, true
}

// HasPermalink returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasPermalink() bool {
	if o != nil && !IsNil(o.Permalink) {
		return true
	}

	return false
}

// SetPermalink gets a reference to the given string and assigns it to the Permalink field.
func (o *AdvisoryGHAdvisoryJSONLean) SetPermalink(v string) {
	o.Permalink = &v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetPublishedAt() string {
	if o == nil || IsNil(o.PublishedAt) {
		var ret string
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetPublishedAtOk() (*string, bool) {
	if o == nil || IsNil(o.PublishedAt) {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasPublishedAt() bool {
	if o != nil && !IsNil(o.PublishedAt) {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given string and assigns it to the PublishedAt field.
func (o *AdvisoryGHAdvisoryJSONLean) SetPublishedAt(v string) {
	o.PublishedAt = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetReferences() []AdvisoryGHReference {
	if o == nil || IsNil(o.References) {
		var ret []AdvisoryGHReference
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetReferencesOk() ([]AdvisoryGHReference, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []AdvisoryGHReference and assigns it to the References field.
func (o *AdvisoryGHAdvisoryJSONLean) SetReferences(v []AdvisoryGHReference) {
	o.References = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *AdvisoryGHAdvisoryJSONLean) SetSeverity(v string) {
	o.Severity = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AdvisoryGHAdvisoryJSONLean) SetSummary(v string) {
	o.Summary = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AdvisoryGHAdvisoryJSONLean) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetVulnerabilities returns the Vulnerabilities field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetVulnerabilities() AdvisoryGHVulnerabilities {
	if o == nil || IsNil(o.Vulnerabilities) {
		var ret AdvisoryGHVulnerabilities
		return ret
	}
	return *o.Vulnerabilities
}

// GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetVulnerabilitiesOk() (*AdvisoryGHVulnerabilities, bool) {
	if o == nil || IsNil(o.Vulnerabilities) {
		return nil, false
	}
	return o.Vulnerabilities, true
}

// HasVulnerabilities returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasVulnerabilities() bool {
	if o != nil && !IsNil(o.Vulnerabilities) {
		return true
	}

	return false
}

// SetVulnerabilities gets a reference to the given AdvisoryGHVulnerabilities and assigns it to the Vulnerabilities field.
func (o *AdvisoryGHAdvisoryJSONLean) SetVulnerabilities(v AdvisoryGHVulnerabilities) {
	o.Vulnerabilities = &v
}

// GetWithdrawnAt returns the WithdrawnAt field value if set, zero value otherwise.
func (o *AdvisoryGHAdvisoryJSONLean) GetWithdrawnAt() string {
	if o == nil || IsNil(o.WithdrawnAt) {
		var ret string
		return ret
	}
	return *o.WithdrawnAt
}

// GetWithdrawnAtOk returns a tuple with the WithdrawnAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryGHAdvisoryJSONLean) GetWithdrawnAtOk() (*string, bool) {
	if o == nil || IsNil(o.WithdrawnAt) {
		return nil, false
	}
	return o.WithdrawnAt, true
}

// HasWithdrawnAt returns a boolean if a field has been set.
func (o *AdvisoryGHAdvisoryJSONLean) HasWithdrawnAt() bool {
	if o != nil && !IsNil(o.WithdrawnAt) {
		return true
	}

	return false
}

// SetWithdrawnAt gets a reference to the given string and assigns it to the WithdrawnAt field.
func (o *AdvisoryGHAdvisoryJSONLean) SetWithdrawnAt(v string) {
	o.WithdrawnAt = &v
}

func (o AdvisoryGHAdvisoryJSONLean) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryGHAdvisoryJSONLean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Classification) {
		toSerialize["classification"] = o.Classification
	}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.Cvss) {
		toSerialize["cvss"] = o.Cvss
	}
	if !IsNil(o.Cwes) {
		toSerialize["cwes"] = o.Cwes
	}
	if !IsNil(o.DatabaseId) {
		toSerialize["databaseId"] = o.DatabaseId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.GhsaId) {
		toSerialize["ghsaId"] = o.GhsaId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Identifiers) {
		toSerialize["identifiers"] = o.Identifiers
	}
	if !IsNil(o.NotificationsPermalink) {
		toSerialize["notificationsPermalink"] = o.NotificationsPermalink
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.Permalink) {
		toSerialize["permalink"] = o.Permalink
	}
	if !IsNil(o.PublishedAt) {
		toSerialize["publishedAt"] = o.PublishedAt
	}
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Vulnerabilities) {
		toSerialize["vulnerabilities"] = o.Vulnerabilities
	}
	if !IsNil(o.WithdrawnAt) {
		toSerialize["withdrawnAt"] = o.WithdrawnAt
	}
	return toSerialize, nil
}

type NullableAdvisoryGHAdvisoryJSONLean struct {
	value *AdvisoryGHAdvisoryJSONLean
	isSet bool
}

func (v NullableAdvisoryGHAdvisoryJSONLean) Get() *AdvisoryGHAdvisoryJSONLean {
	return v.value
}

func (v *NullableAdvisoryGHAdvisoryJSONLean) Set(val *AdvisoryGHAdvisoryJSONLean) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryGHAdvisoryJSONLean) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryGHAdvisoryJSONLean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryGHAdvisoryJSONLean(val *AdvisoryGHAdvisoryJSONLean) *NullableAdvisoryGHAdvisoryJSONLean {
	return &NullableAdvisoryGHAdvisoryJSONLean{value: val, isSet: true}
}

func (v NullableAdvisoryGHAdvisoryJSONLean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryGHAdvisoryJSONLean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


