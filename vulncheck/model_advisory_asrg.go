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

// checks if the AdvisoryASRG type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryASRG{}

// AdvisoryASRG struct for AdvisoryASRG
type AdvisoryASRG struct {
	AffectedProducts *string `json:"affected_products,omitempty"`
	Capec *string `json:"capec,omitempty"`
	Cve []string `json:"cve,omitempty"`
	Cvss *string `json:"cvss,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Description *string `json:"description,omitempty"`
	ProblemType *string `json:"problem_type,omitempty"`
	References []string `json:"references,omitempty"`
	Title *string `json:"title,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAdvisoryASRG instantiates a new AdvisoryASRG object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryASRG() *AdvisoryASRG {
	this := AdvisoryASRG{}
	return &this
}

// NewAdvisoryASRGWithDefaults instantiates a new AdvisoryASRG object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryASRGWithDefaults() *AdvisoryASRG {
	this := AdvisoryASRG{}
	return &this
}

// GetAffectedProducts returns the AffectedProducts field value if set, zero value otherwise.
func (o *AdvisoryASRG) GetAffectedProducts() string {
	if o == nil || IsNil(o.AffectedProducts) {
		var ret string
		return ret
	}
	return *o.AffectedProducts
}

// GetAffectedProductsOk returns a tuple with the AffectedProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryASRG) GetAffectedProductsOk() (*string, bool) {
	if o == nil || IsNil(o.AffectedProducts) {
		return nil, false
	}
	return o.AffectedProducts, true
}

// HasAffectedProducts returns a boolean if a field has been set.
func (o *AdvisoryASRG) HasAffectedProducts() bool {
	if o != nil && !IsNil(o.AffectedProducts) {
		return true
	}

	return false
}

// SetAffectedProducts gets a reference to the given string and assigns it to the AffectedProducts field.
func (o *AdvisoryASRG) SetAffectedProducts(v string) {
	o.AffectedProducts = &v
}

// GetCapec returns the Capec field value if set, zero value otherwise.
func (o *AdvisoryASRG) GetCapec() string {
	if o == nil || IsNil(o.Capec) {
		var ret string
		return ret
	}
	return *o.Capec
}

// GetCapecOk returns a tuple with the Capec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryASRG) GetCapecOk() (*string, bool) {
	if o == nil || IsNil(o.Capec) {
		return nil, false
	}
	return o.Capec, true
}

// HasCapec returns a boolean if a field has been set.
func (o *AdvisoryASRG) HasCapec() bool {
	if o != nil && !IsNil(o.Capec) {
		return true
	}

	return false
}

// SetCapec gets a reference to the given string and assigns it to the Capec field.
func (o *AdvisoryASRG) SetCapec(v string) {
	o.Capec = &v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryASRG) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryASRG) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryASRG) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryASRG) SetCve(v []string) {
	o.Cve = v
}

// GetCvss returns the Cvss field value if set, zero value otherwise.
func (o *AdvisoryASRG) GetCvss() string {
	if o == nil || IsNil(o.Cvss) {
		var ret string
		return ret
	}
	return *o.Cvss
}

// GetCvssOk returns a tuple with the Cvss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryASRG) GetCvssOk() (*string, bool) {
	if o == nil || IsNil(o.Cvss) {
		return nil, false
	}
	return o.Cvss, true
}

// HasCvss returns a boolean if a field has been set.
func (o *AdvisoryASRG) HasCvss() bool {
	if o != nil && !IsNil(o.Cvss) {
		return true
	}

	return false
}

// SetCvss gets a reference to the given string and assigns it to the Cvss field.
func (o *AdvisoryASRG) SetCvss(v string) {
	o.Cvss = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryASRG) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryASRG) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryASRG) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryASRG) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AdvisoryASRG) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryASRG) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AdvisoryASRG) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AdvisoryASRG) SetDescription(v string) {
	o.Description = &v
}

// GetProblemType returns the ProblemType field value if set, zero value otherwise.
func (o *AdvisoryASRG) GetProblemType() string {
	if o == nil || IsNil(o.ProblemType) {
		var ret string
		return ret
	}
	return *o.ProblemType
}

// GetProblemTypeOk returns a tuple with the ProblemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryASRG) GetProblemTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProblemType) {
		return nil, false
	}
	return o.ProblemType, true
}

// HasProblemType returns a boolean if a field has been set.
func (o *AdvisoryASRG) HasProblemType() bool {
	if o != nil && !IsNil(o.ProblemType) {
		return true
	}

	return false
}

// SetProblemType gets a reference to the given string and assigns it to the ProblemType field.
func (o *AdvisoryASRG) SetProblemType(v string) {
	o.ProblemType = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *AdvisoryASRG) GetReferences() []string {
	if o == nil || IsNil(o.References) {
		var ret []string
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryASRG) GetReferencesOk() ([]string, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *AdvisoryASRG) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []string and assigns it to the References field.
func (o *AdvisoryASRG) SetReferences(v []string) {
	o.References = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryASRG) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryASRG) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryASRG) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryASRG) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryASRG) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryASRG) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryASRG) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryASRG) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisoryASRG) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryASRG) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AffectedProducts) {
		toSerialize["affected_products"] = o.AffectedProducts
	}
	if !IsNil(o.Capec) {
		toSerialize["capec"] = o.Capec
	}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.Cvss) {
		toSerialize["cvss"] = o.Cvss
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ProblemType) {
		toSerialize["problem_type"] = o.ProblemType
	}
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAdvisoryASRG struct {
	value *AdvisoryASRG
	isSet bool
}

func (v NullableAdvisoryASRG) Get() *AdvisoryASRG {
	return v.value
}

func (v *NullableAdvisoryASRG) Set(val *AdvisoryASRG) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryASRG) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryASRG) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryASRG(val *AdvisoryASRG) *NullableAdvisoryASRG {
	return &NullableAdvisoryASRG{value: val, isSet: true}
}

func (v NullableAdvisoryASRG) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryASRG) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


