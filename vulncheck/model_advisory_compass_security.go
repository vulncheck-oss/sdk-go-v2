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

// checks if the AdvisoryCompassSecurity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryCompassSecurity{}

// AdvisoryCompassSecurity struct for AdvisoryCompassSecurity
type AdvisoryCompassSecurity struct {
	CsncId *string `json:"csnc_id,omitempty"`
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Effect *string `json:"effect,omitempty"`
	Introduction *string `json:"introduction,omitempty"`
	Product *string `json:"product,omitempty"`
	References []string `json:"references,omitempty"`
	Risk *string `json:"risk,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Title *string `json:"title,omitempty"`
	Url *string `json:"url,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
}

// NewAdvisoryCompassSecurity instantiates a new AdvisoryCompassSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryCompassSecurity() *AdvisoryCompassSecurity {
	this := AdvisoryCompassSecurity{}
	return &this
}

// NewAdvisoryCompassSecurityWithDefaults instantiates a new AdvisoryCompassSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryCompassSecurityWithDefaults() *AdvisoryCompassSecurity {
	this := AdvisoryCompassSecurity{}
	return &this
}

// GetCsncId returns the CsncId field value if set, zero value otherwise.
func (o *AdvisoryCompassSecurity) GetCsncId() string {
	if o == nil || IsNil(o.CsncId) {
		var ret string
		return ret
	}
	return *o.CsncId
}

// GetCsncIdOk returns a tuple with the CsncId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCompassSecurity) GetCsncIdOk() (*string, bool) {
	if o == nil || IsNil(o.CsncId) {
		return nil, false
	}
	return o.CsncId, true
}

// HasCsncId returns a boolean if a field has been set.
func (o *AdvisoryCompassSecurity) HasCsncId() bool {
	if o != nil && !IsNil(o.CsncId) {
		return true
	}

	return false
}

// SetCsncId gets a reference to the given string and assigns it to the CsncId field.
func (o *AdvisoryCompassSecurity) SetCsncId(v string) {
	o.CsncId = &v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryCompassSecurity) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCompassSecurity) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryCompassSecurity) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryCompassSecurity) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryCompassSecurity) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCompassSecurity) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryCompassSecurity) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryCompassSecurity) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetEffect returns the Effect field value if set, zero value otherwise.
func (o *AdvisoryCompassSecurity) GetEffect() string {
	if o == nil || IsNil(o.Effect) {
		var ret string
		return ret
	}
	return *o.Effect
}

// GetEffectOk returns a tuple with the Effect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCompassSecurity) GetEffectOk() (*string, bool) {
	if o == nil || IsNil(o.Effect) {
		return nil, false
	}
	return o.Effect, true
}

// HasEffect returns a boolean if a field has been set.
func (o *AdvisoryCompassSecurity) HasEffect() bool {
	if o != nil && !IsNil(o.Effect) {
		return true
	}

	return false
}

// SetEffect gets a reference to the given string and assigns it to the Effect field.
func (o *AdvisoryCompassSecurity) SetEffect(v string) {
	o.Effect = &v
}

// GetIntroduction returns the Introduction field value if set, zero value otherwise.
func (o *AdvisoryCompassSecurity) GetIntroduction() string {
	if o == nil || IsNil(o.Introduction) {
		var ret string
		return ret
	}
	return *o.Introduction
}

// GetIntroductionOk returns a tuple with the Introduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCompassSecurity) GetIntroductionOk() (*string, bool) {
	if o == nil || IsNil(o.Introduction) {
		return nil, false
	}
	return o.Introduction, true
}

// HasIntroduction returns a boolean if a field has been set.
func (o *AdvisoryCompassSecurity) HasIntroduction() bool {
	if o != nil && !IsNil(o.Introduction) {
		return true
	}

	return false
}

// SetIntroduction gets a reference to the given string and assigns it to the Introduction field.
func (o *AdvisoryCompassSecurity) SetIntroduction(v string) {
	o.Introduction = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *AdvisoryCompassSecurity) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCompassSecurity) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *AdvisoryCompassSecurity) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *AdvisoryCompassSecurity) SetProduct(v string) {
	o.Product = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *AdvisoryCompassSecurity) GetReferences() []string {
	if o == nil || IsNil(o.References) {
		var ret []string
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCompassSecurity) GetReferencesOk() ([]string, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *AdvisoryCompassSecurity) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []string and assigns it to the References field.
func (o *AdvisoryCompassSecurity) SetReferences(v []string) {
	o.References = v
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *AdvisoryCompassSecurity) GetRisk() string {
	if o == nil || IsNil(o.Risk) {
		var ret string
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCompassSecurity) GetRiskOk() (*string, bool) {
	if o == nil || IsNil(o.Risk) {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *AdvisoryCompassSecurity) HasRisk() bool {
	if o != nil && !IsNil(o.Risk) {
		return true
	}

	return false
}

// SetRisk gets a reference to the given string and assigns it to the Risk field.
func (o *AdvisoryCompassSecurity) SetRisk(v string) {
	o.Risk = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *AdvisoryCompassSecurity) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCompassSecurity) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *AdvisoryCompassSecurity) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *AdvisoryCompassSecurity) SetSeverity(v string) {
	o.Severity = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryCompassSecurity) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCompassSecurity) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryCompassSecurity) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryCompassSecurity) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryCompassSecurity) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCompassSecurity) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryCompassSecurity) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryCompassSecurity) SetUrl(v string) {
	o.Url = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *AdvisoryCompassSecurity) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryCompassSecurity) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *AdvisoryCompassSecurity) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *AdvisoryCompassSecurity) SetVendor(v string) {
	o.Vendor = &v
}

func (o AdvisoryCompassSecurity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryCompassSecurity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CsncId) {
		toSerialize["csnc_id"] = o.CsncId
	}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Effect) {
		toSerialize["effect"] = o.Effect
	}
	if !IsNil(o.Introduction) {
		toSerialize["introduction"] = o.Introduction
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	if !IsNil(o.Risk) {
		toSerialize["risk"] = o.Risk
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	return toSerialize, nil
}

type NullableAdvisoryCompassSecurity struct {
	value *AdvisoryCompassSecurity
	isSet bool
}

func (v NullableAdvisoryCompassSecurity) Get() *AdvisoryCompassSecurity {
	return v.value
}

func (v *NullableAdvisoryCompassSecurity) Set(val *AdvisoryCompassSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryCompassSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryCompassSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryCompassSecurity(val *AdvisoryCompassSecurity) *NullableAdvisoryCompassSecurity {
	return &NullableAdvisoryCompassSecurity{value: val, isSet: true}
}

func (v NullableAdvisoryCompassSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryCompassSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


