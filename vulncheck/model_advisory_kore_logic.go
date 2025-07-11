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

// checks if the AdvisoryKoreLogic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryKoreLogic{}

// AdvisoryKoreLogic struct for AdvisoryKoreLogic
type AdvisoryKoreLogic struct {
	AffectedProduct *string `json:"affected_product,omitempty"`
	AffectedVendor *string `json:"affected_vendor,omitempty"`
	AffectedVersion *string `json:"affected_version,omitempty"`
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Id *string `json:"id,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Title *string `json:"title,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAdvisoryKoreLogic instantiates a new AdvisoryKoreLogic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryKoreLogic() *AdvisoryKoreLogic {
	this := AdvisoryKoreLogic{}
	return &this
}

// NewAdvisoryKoreLogicWithDefaults instantiates a new AdvisoryKoreLogic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryKoreLogicWithDefaults() *AdvisoryKoreLogic {
	this := AdvisoryKoreLogic{}
	return &this
}

// GetAffectedProduct returns the AffectedProduct field value if set, zero value otherwise.
func (o *AdvisoryKoreLogic) GetAffectedProduct() string {
	if o == nil || IsNil(o.AffectedProduct) {
		var ret string
		return ret
	}
	return *o.AffectedProduct
}

// GetAffectedProductOk returns a tuple with the AffectedProduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryKoreLogic) GetAffectedProductOk() (*string, bool) {
	if o == nil || IsNil(o.AffectedProduct) {
		return nil, false
	}
	return o.AffectedProduct, true
}

// HasAffectedProduct returns a boolean if a field has been set.
func (o *AdvisoryKoreLogic) HasAffectedProduct() bool {
	if o != nil && !IsNil(o.AffectedProduct) {
		return true
	}

	return false
}

// SetAffectedProduct gets a reference to the given string and assigns it to the AffectedProduct field.
func (o *AdvisoryKoreLogic) SetAffectedProduct(v string) {
	o.AffectedProduct = &v
}

// GetAffectedVendor returns the AffectedVendor field value if set, zero value otherwise.
func (o *AdvisoryKoreLogic) GetAffectedVendor() string {
	if o == nil || IsNil(o.AffectedVendor) {
		var ret string
		return ret
	}
	return *o.AffectedVendor
}

// GetAffectedVendorOk returns a tuple with the AffectedVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryKoreLogic) GetAffectedVendorOk() (*string, bool) {
	if o == nil || IsNil(o.AffectedVendor) {
		return nil, false
	}
	return o.AffectedVendor, true
}

// HasAffectedVendor returns a boolean if a field has been set.
func (o *AdvisoryKoreLogic) HasAffectedVendor() bool {
	if o != nil && !IsNil(o.AffectedVendor) {
		return true
	}

	return false
}

// SetAffectedVendor gets a reference to the given string and assigns it to the AffectedVendor field.
func (o *AdvisoryKoreLogic) SetAffectedVendor(v string) {
	o.AffectedVendor = &v
}

// GetAffectedVersion returns the AffectedVersion field value if set, zero value otherwise.
func (o *AdvisoryKoreLogic) GetAffectedVersion() string {
	if o == nil || IsNil(o.AffectedVersion) {
		var ret string
		return ret
	}
	return *o.AffectedVersion
}

// GetAffectedVersionOk returns a tuple with the AffectedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryKoreLogic) GetAffectedVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AffectedVersion) {
		return nil, false
	}
	return o.AffectedVersion, true
}

// HasAffectedVersion returns a boolean if a field has been set.
func (o *AdvisoryKoreLogic) HasAffectedVersion() bool {
	if o != nil && !IsNil(o.AffectedVersion) {
		return true
	}

	return false
}

// SetAffectedVersion gets a reference to the given string and assigns it to the AffectedVersion field.
func (o *AdvisoryKoreLogic) SetAffectedVersion(v string) {
	o.AffectedVersion = &v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryKoreLogic) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryKoreLogic) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryKoreLogic) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryKoreLogic) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryKoreLogic) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryKoreLogic) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryKoreLogic) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryKoreLogic) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdvisoryKoreLogic) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryKoreLogic) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdvisoryKoreLogic) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdvisoryKoreLogic) SetId(v string) {
	o.Id = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AdvisoryKoreLogic) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryKoreLogic) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AdvisoryKoreLogic) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AdvisoryKoreLogic) SetSummary(v string) {
	o.Summary = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryKoreLogic) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryKoreLogic) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryKoreLogic) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryKoreLogic) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryKoreLogic) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryKoreLogic) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryKoreLogic) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryKoreLogic) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisoryKoreLogic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryKoreLogic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AffectedProduct) {
		toSerialize["affected_product"] = o.AffectedProduct
	}
	if !IsNil(o.AffectedVendor) {
		toSerialize["affected_vendor"] = o.AffectedVendor
	}
	if !IsNil(o.AffectedVersion) {
		toSerialize["affected_version"] = o.AffectedVersion
	}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAdvisoryKoreLogic struct {
	value *AdvisoryKoreLogic
	isSet bool
}

func (v NullableAdvisoryKoreLogic) Get() *AdvisoryKoreLogic {
	return v.value
}

func (v *NullableAdvisoryKoreLogic) Set(val *AdvisoryKoreLogic) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryKoreLogic) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryKoreLogic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryKoreLogic(val *AdvisoryKoreLogic) *NullableAdvisoryKoreLogic {
	return &NullableAdvisoryKoreLogic{value: val, isSet: true}
}

func (v NullableAdvisoryKoreLogic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryKoreLogic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


