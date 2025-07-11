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

// checks if the AdvisoryNCSC type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryNCSC{}

// AdvisoryNCSC struct for AdvisoryNCSC
type AdvisoryNCSC struct {
	Csaf *AdvisoryCSAF `json:"csaf,omitempty"`
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Id *string `json:"id,omitempty"`
	References []string `json:"references,omitempty"`
	SummaryNl *string `json:"summary_nl,omitempty"`
	TitleNl *string `json:"title_nl,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAdvisoryNCSC instantiates a new AdvisoryNCSC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryNCSC() *AdvisoryNCSC {
	this := AdvisoryNCSC{}
	return &this
}

// NewAdvisoryNCSCWithDefaults instantiates a new AdvisoryNCSC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryNCSCWithDefaults() *AdvisoryNCSC {
	this := AdvisoryNCSC{}
	return &this
}

// GetCsaf returns the Csaf field value if set, zero value otherwise.
func (o *AdvisoryNCSC) GetCsaf() AdvisoryCSAF {
	if o == nil || IsNil(o.Csaf) {
		var ret AdvisoryCSAF
		return ret
	}
	return *o.Csaf
}

// GetCsafOk returns a tuple with the Csaf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNCSC) GetCsafOk() (*AdvisoryCSAF, bool) {
	if o == nil || IsNil(o.Csaf) {
		return nil, false
	}
	return o.Csaf, true
}

// HasCsaf returns a boolean if a field has been set.
func (o *AdvisoryNCSC) HasCsaf() bool {
	if o != nil && !IsNil(o.Csaf) {
		return true
	}

	return false
}

// SetCsaf gets a reference to the given AdvisoryCSAF and assigns it to the Csaf field.
func (o *AdvisoryNCSC) SetCsaf(v AdvisoryCSAF) {
	o.Csaf = &v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryNCSC) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNCSC) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryNCSC) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryNCSC) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryNCSC) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNCSC) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryNCSC) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryNCSC) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdvisoryNCSC) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNCSC) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdvisoryNCSC) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdvisoryNCSC) SetId(v string) {
	o.Id = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *AdvisoryNCSC) GetReferences() []string {
	if o == nil || IsNil(o.References) {
		var ret []string
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNCSC) GetReferencesOk() ([]string, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *AdvisoryNCSC) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []string and assigns it to the References field.
func (o *AdvisoryNCSC) SetReferences(v []string) {
	o.References = v
}

// GetSummaryNl returns the SummaryNl field value if set, zero value otherwise.
func (o *AdvisoryNCSC) GetSummaryNl() string {
	if o == nil || IsNil(o.SummaryNl) {
		var ret string
		return ret
	}
	return *o.SummaryNl
}

// GetSummaryNlOk returns a tuple with the SummaryNl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNCSC) GetSummaryNlOk() (*string, bool) {
	if o == nil || IsNil(o.SummaryNl) {
		return nil, false
	}
	return o.SummaryNl, true
}

// HasSummaryNl returns a boolean if a field has been set.
func (o *AdvisoryNCSC) HasSummaryNl() bool {
	if o != nil && !IsNil(o.SummaryNl) {
		return true
	}

	return false
}

// SetSummaryNl gets a reference to the given string and assigns it to the SummaryNl field.
func (o *AdvisoryNCSC) SetSummaryNl(v string) {
	o.SummaryNl = &v
}

// GetTitleNl returns the TitleNl field value if set, zero value otherwise.
func (o *AdvisoryNCSC) GetTitleNl() string {
	if o == nil || IsNil(o.TitleNl) {
		var ret string
		return ret
	}
	return *o.TitleNl
}

// GetTitleNlOk returns a tuple with the TitleNl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNCSC) GetTitleNlOk() (*string, bool) {
	if o == nil || IsNil(o.TitleNl) {
		return nil, false
	}
	return o.TitleNl, true
}

// HasTitleNl returns a boolean if a field has been set.
func (o *AdvisoryNCSC) HasTitleNl() bool {
	if o != nil && !IsNil(o.TitleNl) {
		return true
	}

	return false
}

// SetTitleNl gets a reference to the given string and assigns it to the TitleNl field.
func (o *AdvisoryNCSC) SetTitleNl(v string) {
	o.TitleNl = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryNCSC) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNCSC) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryNCSC) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryNCSC) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisoryNCSC) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryNCSC) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Csaf) {
		toSerialize["csaf"] = o.Csaf
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
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	if !IsNil(o.SummaryNl) {
		toSerialize["summary_nl"] = o.SummaryNl
	}
	if !IsNil(o.TitleNl) {
		toSerialize["title_nl"] = o.TitleNl
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAdvisoryNCSC struct {
	value *AdvisoryNCSC
	isSet bool
}

func (v NullableAdvisoryNCSC) Get() *AdvisoryNCSC {
	return v.value
}

func (v *NullableAdvisoryNCSC) Set(val *AdvisoryNCSC) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryNCSC) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryNCSC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryNCSC(val *AdvisoryNCSC) *NullableAdvisoryNCSC {
	return &NullableAdvisoryNCSC{value: val, isSet: true}
}

func (v NullableAdvisoryNCSC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryNCSC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


