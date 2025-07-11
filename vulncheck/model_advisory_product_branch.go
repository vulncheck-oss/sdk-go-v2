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

// checks if the AdvisoryProductBranch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryProductBranch{}

// AdvisoryProductBranch struct for AdvisoryProductBranch
type AdvisoryProductBranch struct {
	Branches []AdvisoryProductBranch `json:"branches,omitempty"`
	Category *string `json:"category,omitempty"`
	Name *string `json:"name,omitempty"`
	Product *AdvisoryProduct `json:"product,omitempty"`
	Relationships []AdvisoryCSAFRelationship `json:"relationships,omitempty"`
}

// NewAdvisoryProductBranch instantiates a new AdvisoryProductBranch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryProductBranch() *AdvisoryProductBranch {
	this := AdvisoryProductBranch{}
	return &this
}

// NewAdvisoryProductBranchWithDefaults instantiates a new AdvisoryProductBranch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryProductBranchWithDefaults() *AdvisoryProductBranch {
	this := AdvisoryProductBranch{}
	return &this
}

// GetBranches returns the Branches field value if set, zero value otherwise.
func (o *AdvisoryProductBranch) GetBranches() []AdvisoryProductBranch {
	if o == nil || IsNil(o.Branches) {
		var ret []AdvisoryProductBranch
		return ret
	}
	return o.Branches
}

// GetBranchesOk returns a tuple with the Branches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryProductBranch) GetBranchesOk() ([]AdvisoryProductBranch, bool) {
	if o == nil || IsNil(o.Branches) {
		return nil, false
	}
	return o.Branches, true
}

// HasBranches returns a boolean if a field has been set.
func (o *AdvisoryProductBranch) HasBranches() bool {
	if o != nil && !IsNil(o.Branches) {
		return true
	}

	return false
}

// SetBranches gets a reference to the given []AdvisoryProductBranch and assigns it to the Branches field.
func (o *AdvisoryProductBranch) SetBranches(v []AdvisoryProductBranch) {
	o.Branches = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AdvisoryProductBranch) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryProductBranch) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AdvisoryProductBranch) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *AdvisoryProductBranch) SetCategory(v string) {
	o.Category = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdvisoryProductBranch) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryProductBranch) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdvisoryProductBranch) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdvisoryProductBranch) SetName(v string) {
	o.Name = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *AdvisoryProductBranch) GetProduct() AdvisoryProduct {
	if o == nil || IsNil(o.Product) {
		var ret AdvisoryProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryProductBranch) GetProductOk() (*AdvisoryProduct, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *AdvisoryProductBranch) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given AdvisoryProduct and assigns it to the Product field.
func (o *AdvisoryProductBranch) SetProduct(v AdvisoryProduct) {
	o.Product = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AdvisoryProductBranch) GetRelationships() []AdvisoryCSAFRelationship {
	if o == nil || IsNil(o.Relationships) {
		var ret []AdvisoryCSAFRelationship
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryProductBranch) GetRelationshipsOk() ([]AdvisoryCSAFRelationship, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AdvisoryProductBranch) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given []AdvisoryCSAFRelationship and assigns it to the Relationships field.
func (o *AdvisoryProductBranch) SetRelationships(v []AdvisoryCSAFRelationship) {
	o.Relationships = v
}

func (o AdvisoryProductBranch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryProductBranch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Branches) {
		toSerialize["branches"] = o.Branches
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableAdvisoryProductBranch struct {
	value *AdvisoryProductBranch
	isSet bool
}

func (v NullableAdvisoryProductBranch) Get() *AdvisoryProductBranch {
	return v.value
}

func (v *NullableAdvisoryProductBranch) Set(val *AdvisoryProductBranch) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryProductBranch) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryProductBranch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryProductBranch(val *AdvisoryProductBranch) *NullableAdvisoryProductBranch {
	return &NullableAdvisoryProductBranch{value: val, isSet: true}
}

func (v NullableAdvisoryProductBranch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryProductBranch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


