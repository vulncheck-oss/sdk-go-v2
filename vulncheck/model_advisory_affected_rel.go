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

// checks if the AdvisoryAffectedRel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryAffectedRel{}

// AdvisoryAffectedRel struct for AdvisoryAffectedRel
type AdvisoryAffectedRel struct {
	Advisory *string `json:"advisory,omitempty"`
	Cpe *string `json:"cpe,omitempty"`
	Package *string `json:"package,omitempty"`
	ProductName *string `json:"product_name,omitempty"`
	ReleaseDate *string `json:"release_date,omitempty"`
}

// NewAdvisoryAffectedRel instantiates a new AdvisoryAffectedRel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryAffectedRel() *AdvisoryAffectedRel {
	this := AdvisoryAffectedRel{}
	return &this
}

// NewAdvisoryAffectedRelWithDefaults instantiates a new AdvisoryAffectedRel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryAffectedRelWithDefaults() *AdvisoryAffectedRel {
	this := AdvisoryAffectedRel{}
	return &this
}

// GetAdvisory returns the Advisory field value if set, zero value otherwise.
func (o *AdvisoryAffectedRel) GetAdvisory() string {
	if o == nil || IsNil(o.Advisory) {
		var ret string
		return ret
	}
	return *o.Advisory
}

// GetAdvisoryOk returns a tuple with the Advisory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAffectedRel) GetAdvisoryOk() (*string, bool) {
	if o == nil || IsNil(o.Advisory) {
		return nil, false
	}
	return o.Advisory, true
}

// HasAdvisory returns a boolean if a field has been set.
func (o *AdvisoryAffectedRel) HasAdvisory() bool {
	if o != nil && !IsNil(o.Advisory) {
		return true
	}

	return false
}

// SetAdvisory gets a reference to the given string and assigns it to the Advisory field.
func (o *AdvisoryAffectedRel) SetAdvisory(v string) {
	o.Advisory = &v
}

// GetCpe returns the Cpe field value if set, zero value otherwise.
func (o *AdvisoryAffectedRel) GetCpe() string {
	if o == nil || IsNil(o.Cpe) {
		var ret string
		return ret
	}
	return *o.Cpe
}

// GetCpeOk returns a tuple with the Cpe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAffectedRel) GetCpeOk() (*string, bool) {
	if o == nil || IsNil(o.Cpe) {
		return nil, false
	}
	return o.Cpe, true
}

// HasCpe returns a boolean if a field has been set.
func (o *AdvisoryAffectedRel) HasCpe() bool {
	if o != nil && !IsNil(o.Cpe) {
		return true
	}

	return false
}

// SetCpe gets a reference to the given string and assigns it to the Cpe field.
func (o *AdvisoryAffectedRel) SetCpe(v string) {
	o.Cpe = &v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *AdvisoryAffectedRel) GetPackage() string {
	if o == nil || IsNil(o.Package) {
		var ret string
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAffectedRel) GetPackageOk() (*string, bool) {
	if o == nil || IsNil(o.Package) {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *AdvisoryAffectedRel) HasPackage() bool {
	if o != nil && !IsNil(o.Package) {
		return true
	}

	return false
}

// SetPackage gets a reference to the given string and assigns it to the Package field.
func (o *AdvisoryAffectedRel) SetPackage(v string) {
	o.Package = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *AdvisoryAffectedRel) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAffectedRel) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *AdvisoryAffectedRel) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *AdvisoryAffectedRel) SetProductName(v string) {
	o.ProductName = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *AdvisoryAffectedRel) GetReleaseDate() string {
	if o == nil || IsNil(o.ReleaseDate) {
		var ret string
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAffectedRel) GetReleaseDateOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseDate) {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *AdvisoryAffectedRel) HasReleaseDate() bool {
	if o != nil && !IsNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given string and assigns it to the ReleaseDate field.
func (o *AdvisoryAffectedRel) SetReleaseDate(v string) {
	o.ReleaseDate = &v
}

func (o AdvisoryAffectedRel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryAffectedRel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Advisory) {
		toSerialize["advisory"] = o.Advisory
	}
	if !IsNil(o.Cpe) {
		toSerialize["cpe"] = o.Cpe
	}
	if !IsNil(o.Package) {
		toSerialize["package"] = o.Package
	}
	if !IsNil(o.ProductName) {
		toSerialize["product_name"] = o.ProductName
	}
	if !IsNil(o.ReleaseDate) {
		toSerialize["release_date"] = o.ReleaseDate
	}
	return toSerialize, nil
}

type NullableAdvisoryAffectedRel struct {
	value *AdvisoryAffectedRel
	isSet bool
}

func (v NullableAdvisoryAffectedRel) Get() *AdvisoryAffectedRel {
	return v.value
}

func (v *NullableAdvisoryAffectedRel) Set(val *AdvisoryAffectedRel) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryAffectedRel) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryAffectedRel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryAffectedRel(val *AdvisoryAffectedRel) *NullableAdvisoryAffectedRel {
	return &NullableAdvisoryAffectedRel{value: val, isSet: true}
}

func (v NullableAdvisoryAffectedRel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryAffectedRel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


