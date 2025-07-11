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

// checks if the AdvisoryBectonDickinsonAdvisory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryBectonDickinsonAdvisory{}

// AdvisoryBectonDickinsonAdvisory struct for AdvisoryBectonDickinsonAdvisory
type AdvisoryBectonDickinsonAdvisory struct {
	Cve []string `json:"cve,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	ProductsAffected []AdvisoryProductsAffected `json:"products_affected,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAdvisoryBectonDickinsonAdvisory instantiates a new AdvisoryBectonDickinsonAdvisory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryBectonDickinsonAdvisory() *AdvisoryBectonDickinsonAdvisory {
	this := AdvisoryBectonDickinsonAdvisory{}
	return &this
}

// NewAdvisoryBectonDickinsonAdvisoryWithDefaults instantiates a new AdvisoryBectonDickinsonAdvisory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryBectonDickinsonAdvisoryWithDefaults() *AdvisoryBectonDickinsonAdvisory {
	this := AdvisoryBectonDickinsonAdvisory{}
	return &this
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryBectonDickinsonAdvisory) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryBectonDickinsonAdvisory) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryBectonDickinsonAdvisory) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryBectonDickinsonAdvisory) SetCve(v []string) {
	o.Cve = v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryBectonDickinsonAdvisory) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryBectonDickinsonAdvisory) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryBectonDickinsonAdvisory) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryBectonDickinsonAdvisory) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetProductsAffected returns the ProductsAffected field value if set, zero value otherwise.
func (o *AdvisoryBectonDickinsonAdvisory) GetProductsAffected() []AdvisoryProductsAffected {
	if o == nil || IsNil(o.ProductsAffected) {
		var ret []AdvisoryProductsAffected
		return ret
	}
	return o.ProductsAffected
}

// GetProductsAffectedOk returns a tuple with the ProductsAffected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryBectonDickinsonAdvisory) GetProductsAffectedOk() ([]AdvisoryProductsAffected, bool) {
	if o == nil || IsNil(o.ProductsAffected) {
		return nil, false
	}
	return o.ProductsAffected, true
}

// HasProductsAffected returns a boolean if a field has been set.
func (o *AdvisoryBectonDickinsonAdvisory) HasProductsAffected() bool {
	if o != nil && !IsNil(o.ProductsAffected) {
		return true
	}

	return false
}

// SetProductsAffected gets a reference to the given []AdvisoryProductsAffected and assigns it to the ProductsAffected field.
func (o *AdvisoryBectonDickinsonAdvisory) SetProductsAffected(v []AdvisoryProductsAffected) {
	o.ProductsAffected = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryBectonDickinsonAdvisory) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryBectonDickinsonAdvisory) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryBectonDickinsonAdvisory) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryBectonDickinsonAdvisory) SetTitle(v string) {
	o.Title = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AdvisoryBectonDickinsonAdvisory) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryBectonDickinsonAdvisory) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AdvisoryBectonDickinsonAdvisory) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AdvisoryBectonDickinsonAdvisory) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryBectonDickinsonAdvisory) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryBectonDickinsonAdvisory) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryBectonDickinsonAdvisory) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryBectonDickinsonAdvisory) SetUrl(v string) {
	o.Url = &v
}

func (o AdvisoryBectonDickinsonAdvisory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryBectonDickinsonAdvisory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.ProductsAffected) {
		toSerialize["products_affected"] = o.ProductsAffected
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
	return toSerialize, nil
}

type NullableAdvisoryBectonDickinsonAdvisory struct {
	value *AdvisoryBectonDickinsonAdvisory
	isSet bool
}

func (v NullableAdvisoryBectonDickinsonAdvisory) Get() *AdvisoryBectonDickinsonAdvisory {
	return v.value
}

func (v *NullableAdvisoryBectonDickinsonAdvisory) Set(val *AdvisoryBectonDickinsonAdvisory) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryBectonDickinsonAdvisory) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryBectonDickinsonAdvisory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryBectonDickinsonAdvisory(val *AdvisoryBectonDickinsonAdvisory) *NullableAdvisoryBectonDickinsonAdvisory {
	return &NullableAdvisoryBectonDickinsonAdvisory{value: val, isSet: true}
}

func (v NullableAdvisoryBectonDickinsonAdvisory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryBectonDickinsonAdvisory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


