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

// checks if the ApiUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiUpdate{}

// ApiUpdate struct for ApiUpdate
type ApiUpdate struct {
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Description *string `json:"description,omitempty"`
	// sort // key
	Id *string `json:"id,omitempty"`
	Issued *ApiDateTime `json:"issued,omitempty"`
	OsArch *string `json:"os_arch,omitempty"`
	OsVersion *string `json:"os_version,omitempty"`
	Packages []ApiPackage `json:"packages,omitempty"`
	References []ApiReference `json:"references,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
	Updated *ApiDateTime `json:"updated,omitempty"`
}

// NewApiUpdate instantiates a new ApiUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiUpdate() *ApiUpdate {
	this := ApiUpdate{}
	return &this
}

// NewApiUpdateWithDefaults instantiates a new ApiUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiUpdateWithDefaults() *ApiUpdate {
	this := ApiUpdate{}
	return &this
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *ApiUpdate) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUpdate) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *ApiUpdate) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *ApiUpdate) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *ApiUpdate) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUpdate) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *ApiUpdate) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *ApiUpdate) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiUpdate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUpdate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiUpdate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiUpdate) SetId(v string) {
	o.Id = &v
}

// GetIssued returns the Issued field value if set, zero value otherwise.
func (o *ApiUpdate) GetIssued() ApiDateTime {
	if o == nil || IsNil(o.Issued) {
		var ret ApiDateTime
		return ret
	}
	return *o.Issued
}

// GetIssuedOk returns a tuple with the Issued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUpdate) GetIssuedOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Issued) {
		return nil, false
	}
	return o.Issued, true
}

// HasIssued returns a boolean if a field has been set.
func (o *ApiUpdate) HasIssued() bool {
	if o != nil && !IsNil(o.Issued) {
		return true
	}

	return false
}

// SetIssued gets a reference to the given ApiDateTime and assigns it to the Issued field.
func (o *ApiUpdate) SetIssued(v ApiDateTime) {
	o.Issued = &v
}

// GetOsArch returns the OsArch field value if set, zero value otherwise.
func (o *ApiUpdate) GetOsArch() string {
	if o == nil || IsNil(o.OsArch) {
		var ret string
		return ret
	}
	return *o.OsArch
}

// GetOsArchOk returns a tuple with the OsArch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUpdate) GetOsArchOk() (*string, bool) {
	if o == nil || IsNil(o.OsArch) {
		return nil, false
	}
	return o.OsArch, true
}

// HasOsArch returns a boolean if a field has been set.
func (o *ApiUpdate) HasOsArch() bool {
	if o != nil && !IsNil(o.OsArch) {
		return true
	}

	return false
}

// SetOsArch gets a reference to the given string and assigns it to the OsArch field.
func (o *ApiUpdate) SetOsArch(v string) {
	o.OsArch = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *ApiUpdate) GetOsVersion() string {
	if o == nil || IsNil(o.OsVersion) {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUpdate) GetOsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *ApiUpdate) HasOsVersion() bool {
	if o != nil && !IsNil(o.OsVersion) {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *ApiUpdate) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *ApiUpdate) GetPackages() []ApiPackage {
	if o == nil || IsNil(o.Packages) {
		var ret []ApiPackage
		return ret
	}
	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUpdate) GetPackagesOk() ([]ApiPackage, bool) {
	if o == nil || IsNil(o.Packages) {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *ApiUpdate) HasPackages() bool {
	if o != nil && !IsNil(o.Packages) {
		return true
	}

	return false
}

// SetPackages gets a reference to the given []ApiPackage and assigns it to the Packages field.
func (o *ApiUpdate) SetPackages(v []ApiPackage) {
	o.Packages = v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *ApiUpdate) GetReferences() []ApiReference {
	if o == nil || IsNil(o.References) {
		var ret []ApiReference
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUpdate) GetReferencesOk() ([]ApiReference, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *ApiUpdate) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []ApiReference and assigns it to the References field.
func (o *ApiUpdate) SetReferences(v []ApiReference) {
	o.References = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ApiUpdate) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUpdate) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ApiUpdate) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ApiUpdate) SetSeverity(v string) {
	o.Severity = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ApiUpdate) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUpdate) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApiUpdate) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ApiUpdate) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiUpdate) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUpdate) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiUpdate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiUpdate) SetType(v string) {
	o.Type = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ApiUpdate) GetUpdated() ApiDateTime {
	if o == nil || IsNil(o.Updated) {
		var ret ApiDateTime
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUpdate) GetUpdatedOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ApiUpdate) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given ApiDateTime and assigns it to the Updated field.
func (o *ApiUpdate) SetUpdated(v ApiDateTime) {
	o.Updated = &v
}

func (o ApiUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Issued) {
		toSerialize["issued"] = o.Issued
	}
	if !IsNil(o.OsArch) {
		toSerialize["os_arch"] = o.OsArch
	}
	if !IsNil(o.OsVersion) {
		toSerialize["os_version"] = o.OsVersion
	}
	if !IsNil(o.Packages) {
		toSerialize["packages"] = o.Packages
	}
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	return toSerialize, nil
}

type NullableApiUpdate struct {
	value *ApiUpdate
	isSet bool
}

func (v NullableApiUpdate) Get() *ApiUpdate {
	return v.value
}

func (v *NullableApiUpdate) Set(val *ApiUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableApiUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableApiUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiUpdate(val *ApiUpdate) *NullableApiUpdate {
	return &NullableApiUpdate{value: val, isSet: true}
}

func (v NullableApiUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


