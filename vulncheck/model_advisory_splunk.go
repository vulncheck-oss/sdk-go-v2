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

// checks if the AdvisorySplunk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisorySplunk{}

// AdvisorySplunk struct for AdvisorySplunk
type AdvisorySplunk struct {
	AdvisoryId *string `json:"advisory_id,omitempty"`
	AffectedProducts []AdvisorySplunkProduct `json:"affected_products,omitempty"`
	BugId *string `json:"bug_id,omitempty"`
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Title *string `json:"title,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAdvisorySplunk instantiates a new AdvisorySplunk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisorySplunk() *AdvisorySplunk {
	this := AdvisorySplunk{}
	return &this
}

// NewAdvisorySplunkWithDefaults instantiates a new AdvisorySplunk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisorySplunkWithDefaults() *AdvisorySplunk {
	this := AdvisorySplunk{}
	return &this
}

// GetAdvisoryId returns the AdvisoryId field value if set, zero value otherwise.
func (o *AdvisorySplunk) GetAdvisoryId() string {
	if o == nil || IsNil(o.AdvisoryId) {
		var ret string
		return ret
	}
	return *o.AdvisoryId
}

// GetAdvisoryIdOk returns a tuple with the AdvisoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySplunk) GetAdvisoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdvisoryId) {
		return nil, false
	}
	return o.AdvisoryId, true
}

// HasAdvisoryId returns a boolean if a field has been set.
func (o *AdvisorySplunk) HasAdvisoryId() bool {
	if o != nil && !IsNil(o.AdvisoryId) {
		return true
	}

	return false
}

// SetAdvisoryId gets a reference to the given string and assigns it to the AdvisoryId field.
func (o *AdvisorySplunk) SetAdvisoryId(v string) {
	o.AdvisoryId = &v
}

// GetAffectedProducts returns the AffectedProducts field value if set, zero value otherwise.
func (o *AdvisorySplunk) GetAffectedProducts() []AdvisorySplunkProduct {
	if o == nil || IsNil(o.AffectedProducts) {
		var ret []AdvisorySplunkProduct
		return ret
	}
	return o.AffectedProducts
}

// GetAffectedProductsOk returns a tuple with the AffectedProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySplunk) GetAffectedProductsOk() ([]AdvisorySplunkProduct, bool) {
	if o == nil || IsNil(o.AffectedProducts) {
		return nil, false
	}
	return o.AffectedProducts, true
}

// HasAffectedProducts returns a boolean if a field has been set.
func (o *AdvisorySplunk) HasAffectedProducts() bool {
	if o != nil && !IsNil(o.AffectedProducts) {
		return true
	}

	return false
}

// SetAffectedProducts gets a reference to the given []AdvisorySplunkProduct and assigns it to the AffectedProducts field.
func (o *AdvisorySplunk) SetAffectedProducts(v []AdvisorySplunkProduct) {
	o.AffectedProducts = v
}

// GetBugId returns the BugId field value if set, zero value otherwise.
func (o *AdvisorySplunk) GetBugId() string {
	if o == nil || IsNil(o.BugId) {
		var ret string
		return ret
	}
	return *o.BugId
}

// GetBugIdOk returns a tuple with the BugId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySplunk) GetBugIdOk() (*string, bool) {
	if o == nil || IsNil(o.BugId) {
		return nil, false
	}
	return o.BugId, true
}

// HasBugId returns a boolean if a field has been set.
func (o *AdvisorySplunk) HasBugId() bool {
	if o != nil && !IsNil(o.BugId) {
		return true
	}

	return false
}

// SetBugId gets a reference to the given string and assigns it to the BugId field.
func (o *AdvisorySplunk) SetBugId(v string) {
	o.BugId = &v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisorySplunk) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySplunk) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisorySplunk) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisorySplunk) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisorySplunk) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySplunk) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisorySplunk) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisorySplunk) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AdvisorySplunk) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySplunk) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AdvisorySplunk) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AdvisorySplunk) SetSummary(v string) {
	o.Summary = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisorySplunk) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySplunk) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisorySplunk) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisorySplunk) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisorySplunk) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySplunk) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisorySplunk) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisorySplunk) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisorySplunk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisorySplunk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdvisoryId) {
		toSerialize["advisory_id"] = o.AdvisoryId
	}
	if !IsNil(o.AffectedProducts) {
		toSerialize["affected_products"] = o.AffectedProducts
	}
	if !IsNil(o.BugId) {
		toSerialize["bug_id"] = o.BugId
	}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
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

type NullableAdvisorySplunk struct {
	value *AdvisorySplunk
	isSet bool
}

func (v NullableAdvisorySplunk) Get() *AdvisorySplunk {
	return v.value
}

func (v *NullableAdvisorySplunk) Set(val *AdvisorySplunk) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisorySplunk) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisorySplunk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisorySplunk(val *AdvisorySplunk) *NullableAdvisorySplunk {
	return &NullableAdvisorySplunk{value: val, isSet: true}
}

func (v NullableAdvisorySplunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisorySplunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


