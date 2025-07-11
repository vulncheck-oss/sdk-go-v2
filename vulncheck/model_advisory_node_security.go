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

// checks if the AdvisoryNodeSecurity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryNodeSecurity{}

// AdvisoryNodeSecurity struct for AdvisoryNodeSecurity
type AdvisoryNodeSecurity struct {
	AffectedEnvironments []string `json:"affected_environments,omitempty"`
	Author *AdvisoryNodeAuthor `json:"author,omitempty"`
	CoordinatingVendor *string `json:"coordinating_vendor,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Cve []string `json:"cve,omitempty"`
	CvssScore *float32 `json:"cvss_score,omitempty"`
	CvssVector *string `json:"cvss_vector,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Id *int32 `json:"id,omitempty"`
	ModuleName *string `json:"module_name,omitempty"`
	Overview *string `json:"overview,omitempty"`
	PatchedVersions *string `json:"patched_versions,omitempty"`
	PublishDate *string `json:"publish_date,omitempty"`
	Recommendation *string `json:"recommendation,omitempty"`
	References []string `json:"references,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	Url *string `json:"url,omitempty"`
	VulnerableVersions *string `json:"vulnerable_versions,omitempty"`
}

// NewAdvisoryNodeSecurity instantiates a new AdvisoryNodeSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryNodeSecurity() *AdvisoryNodeSecurity {
	this := AdvisoryNodeSecurity{}
	return &this
}

// NewAdvisoryNodeSecurityWithDefaults instantiates a new AdvisoryNodeSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryNodeSecurityWithDefaults() *AdvisoryNodeSecurity {
	this := AdvisoryNodeSecurity{}
	return &this
}

// GetAffectedEnvironments returns the AffectedEnvironments field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetAffectedEnvironments() []string {
	if o == nil || IsNil(o.AffectedEnvironments) {
		var ret []string
		return ret
	}
	return o.AffectedEnvironments
}

// GetAffectedEnvironmentsOk returns a tuple with the AffectedEnvironments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetAffectedEnvironmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.AffectedEnvironments) {
		return nil, false
	}
	return o.AffectedEnvironments, true
}

// HasAffectedEnvironments returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasAffectedEnvironments() bool {
	if o != nil && !IsNil(o.AffectedEnvironments) {
		return true
	}

	return false
}

// SetAffectedEnvironments gets a reference to the given []string and assigns it to the AffectedEnvironments field.
func (o *AdvisoryNodeSecurity) SetAffectedEnvironments(v []string) {
	o.AffectedEnvironments = v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetAuthor() AdvisoryNodeAuthor {
	if o == nil || IsNil(o.Author) {
		var ret AdvisoryNodeAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetAuthorOk() (*AdvisoryNodeAuthor, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given AdvisoryNodeAuthor and assigns it to the Author field.
func (o *AdvisoryNodeSecurity) SetAuthor(v AdvisoryNodeAuthor) {
	o.Author = &v
}

// GetCoordinatingVendor returns the CoordinatingVendor field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetCoordinatingVendor() string {
	if o == nil || IsNil(o.CoordinatingVendor) {
		var ret string
		return ret
	}
	return *o.CoordinatingVendor
}

// GetCoordinatingVendorOk returns a tuple with the CoordinatingVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetCoordinatingVendorOk() (*string, bool) {
	if o == nil || IsNil(o.CoordinatingVendor) {
		return nil, false
	}
	return o.CoordinatingVendor, true
}

// HasCoordinatingVendor returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasCoordinatingVendor() bool {
	if o != nil && !IsNil(o.CoordinatingVendor) {
		return true
	}

	return false
}

// SetCoordinatingVendor gets a reference to the given string and assigns it to the CoordinatingVendor field.
func (o *AdvisoryNodeSecurity) SetCoordinatingVendor(v string) {
	o.CoordinatingVendor = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AdvisoryNodeSecurity) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryNodeSecurity) SetCve(v []string) {
	o.Cve = v
}

// GetCvssScore returns the CvssScore field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetCvssScore() float32 {
	if o == nil || IsNil(o.CvssScore) {
		var ret float32
		return ret
	}
	return *o.CvssScore
}

// GetCvssScoreOk returns a tuple with the CvssScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetCvssScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.CvssScore) {
		return nil, false
	}
	return o.CvssScore, true
}

// HasCvssScore returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasCvssScore() bool {
	if o != nil && !IsNil(o.CvssScore) {
		return true
	}

	return false
}

// SetCvssScore gets a reference to the given float32 and assigns it to the CvssScore field.
func (o *AdvisoryNodeSecurity) SetCvssScore(v float32) {
	o.CvssScore = &v
}

// GetCvssVector returns the CvssVector field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetCvssVector() string {
	if o == nil || IsNil(o.CvssVector) {
		var ret string
		return ret
	}
	return *o.CvssVector
}

// GetCvssVectorOk returns a tuple with the CvssVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetCvssVectorOk() (*string, bool) {
	if o == nil || IsNil(o.CvssVector) {
		return nil, false
	}
	return o.CvssVector, true
}

// HasCvssVector returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasCvssVector() bool {
	if o != nil && !IsNil(o.CvssVector) {
		return true
	}

	return false
}

// SetCvssVector gets a reference to the given string and assigns it to the CvssVector field.
func (o *AdvisoryNodeSecurity) SetCvssVector(v string) {
	o.CvssVector = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryNodeSecurity) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AdvisoryNodeSecurity) SetId(v int32) {
	o.Id = &v
}

// GetModuleName returns the ModuleName field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetModuleName() string {
	if o == nil || IsNil(o.ModuleName) {
		var ret string
		return ret
	}
	return *o.ModuleName
}

// GetModuleNameOk returns a tuple with the ModuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetModuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.ModuleName) {
		return nil, false
	}
	return o.ModuleName, true
}

// HasModuleName returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasModuleName() bool {
	if o != nil && !IsNil(o.ModuleName) {
		return true
	}

	return false
}

// SetModuleName gets a reference to the given string and assigns it to the ModuleName field.
func (o *AdvisoryNodeSecurity) SetModuleName(v string) {
	o.ModuleName = &v
}

// GetOverview returns the Overview field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetOverview() string {
	if o == nil || IsNil(o.Overview) {
		var ret string
		return ret
	}
	return *o.Overview
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetOverviewOk() (*string, bool) {
	if o == nil || IsNil(o.Overview) {
		return nil, false
	}
	return o.Overview, true
}

// HasOverview returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasOverview() bool {
	if o != nil && !IsNil(o.Overview) {
		return true
	}

	return false
}

// SetOverview gets a reference to the given string and assigns it to the Overview field.
func (o *AdvisoryNodeSecurity) SetOverview(v string) {
	o.Overview = &v
}

// GetPatchedVersions returns the PatchedVersions field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetPatchedVersions() string {
	if o == nil || IsNil(o.PatchedVersions) {
		var ret string
		return ret
	}
	return *o.PatchedVersions
}

// GetPatchedVersionsOk returns a tuple with the PatchedVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetPatchedVersionsOk() (*string, bool) {
	if o == nil || IsNil(o.PatchedVersions) {
		return nil, false
	}
	return o.PatchedVersions, true
}

// HasPatchedVersions returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasPatchedVersions() bool {
	if o != nil && !IsNil(o.PatchedVersions) {
		return true
	}

	return false
}

// SetPatchedVersions gets a reference to the given string and assigns it to the PatchedVersions field.
func (o *AdvisoryNodeSecurity) SetPatchedVersions(v string) {
	o.PatchedVersions = &v
}

// GetPublishDate returns the PublishDate field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetPublishDate() string {
	if o == nil || IsNil(o.PublishDate) {
		var ret string
		return ret
	}
	return *o.PublishDate
}

// GetPublishDateOk returns a tuple with the PublishDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetPublishDateOk() (*string, bool) {
	if o == nil || IsNil(o.PublishDate) {
		return nil, false
	}
	return o.PublishDate, true
}

// HasPublishDate returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasPublishDate() bool {
	if o != nil && !IsNil(o.PublishDate) {
		return true
	}

	return false
}

// SetPublishDate gets a reference to the given string and assigns it to the PublishDate field.
func (o *AdvisoryNodeSecurity) SetPublishDate(v string) {
	o.PublishDate = &v
}

// GetRecommendation returns the Recommendation field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetRecommendation() string {
	if o == nil || IsNil(o.Recommendation) {
		var ret string
		return ret
	}
	return *o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetRecommendationOk() (*string, bool) {
	if o == nil || IsNil(o.Recommendation) {
		return nil, false
	}
	return o.Recommendation, true
}

// HasRecommendation returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasRecommendation() bool {
	if o != nil && !IsNil(o.Recommendation) {
		return true
	}

	return false
}

// SetRecommendation gets a reference to the given string and assigns it to the Recommendation field.
func (o *AdvisoryNodeSecurity) SetRecommendation(v string) {
	o.Recommendation = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetReferences() []string {
	if o == nil || IsNil(o.References) {
		var ret []string
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetReferencesOk() ([]string, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []string and assigns it to the References field.
func (o *AdvisoryNodeSecurity) SetReferences(v []string) {
	o.References = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryNodeSecurity) SetTitle(v string) {
	o.Title = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AdvisoryNodeSecurity) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryNodeSecurity) SetUrl(v string) {
	o.Url = &v
}

// GetVulnerableVersions returns the VulnerableVersions field value if set, zero value otherwise.
func (o *AdvisoryNodeSecurity) GetVulnerableVersions() string {
	if o == nil || IsNil(o.VulnerableVersions) {
		var ret string
		return ret
	}
	return *o.VulnerableVersions
}

// GetVulnerableVersionsOk returns a tuple with the VulnerableVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNodeSecurity) GetVulnerableVersionsOk() (*string, bool) {
	if o == nil || IsNil(o.VulnerableVersions) {
		return nil, false
	}
	return o.VulnerableVersions, true
}

// HasVulnerableVersions returns a boolean if a field has been set.
func (o *AdvisoryNodeSecurity) HasVulnerableVersions() bool {
	if o != nil && !IsNil(o.VulnerableVersions) {
		return true
	}

	return false
}

// SetVulnerableVersions gets a reference to the given string and assigns it to the VulnerableVersions field.
func (o *AdvisoryNodeSecurity) SetVulnerableVersions(v string) {
	o.VulnerableVersions = &v
}

func (o AdvisoryNodeSecurity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryNodeSecurity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AffectedEnvironments) {
		toSerialize["affected_environments"] = o.AffectedEnvironments
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.CoordinatingVendor) {
		toSerialize["coordinating_vendor"] = o.CoordinatingVendor
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.CvssScore) {
		toSerialize["cvss_score"] = o.CvssScore
	}
	if !IsNil(o.CvssVector) {
		toSerialize["cvss_vector"] = o.CvssVector
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ModuleName) {
		toSerialize["module_name"] = o.ModuleName
	}
	if !IsNil(o.Overview) {
		toSerialize["overview"] = o.Overview
	}
	if !IsNil(o.PatchedVersions) {
		toSerialize["patched_versions"] = o.PatchedVersions
	}
	if !IsNil(o.PublishDate) {
		toSerialize["publish_date"] = o.PublishDate
	}
	if !IsNil(o.Recommendation) {
		toSerialize["recommendation"] = o.Recommendation
	}
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.VulnerableVersions) {
		toSerialize["vulnerable_versions"] = o.VulnerableVersions
	}
	return toSerialize, nil
}

type NullableAdvisoryNodeSecurity struct {
	value *AdvisoryNodeSecurity
	isSet bool
}

func (v NullableAdvisoryNodeSecurity) Get() *AdvisoryNodeSecurity {
	return v.value
}

func (v *NullableAdvisoryNodeSecurity) Set(val *AdvisoryNodeSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryNodeSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryNodeSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryNodeSecurity(val *AdvisoryNodeSecurity) *NullableAdvisoryNodeSecurity {
	return &NullableAdvisoryNodeSecurity{value: val, isSet: true}
}

func (v NullableAdvisoryNodeSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryNodeSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


