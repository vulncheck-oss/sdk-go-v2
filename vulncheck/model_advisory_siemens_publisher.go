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

// checks if the AdvisorySiemensPublisher type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisorySiemensPublisher{}

// AdvisorySiemensPublisher struct for AdvisorySiemensPublisher
type AdvisorySiemensPublisher struct {
	Category *string `json:"category,omitempty"`
	ContactDetails *string `json:"contact_details,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// NewAdvisorySiemensPublisher instantiates a new AdvisorySiemensPublisher object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisorySiemensPublisher() *AdvisorySiemensPublisher {
	this := AdvisorySiemensPublisher{}
	return &this
}

// NewAdvisorySiemensPublisherWithDefaults instantiates a new AdvisorySiemensPublisher object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisorySiemensPublisherWithDefaults() *AdvisorySiemensPublisher {
	this := AdvisorySiemensPublisher{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AdvisorySiemensPublisher) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySiemensPublisher) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AdvisorySiemensPublisher) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *AdvisorySiemensPublisher) SetCategory(v string) {
	o.Category = &v
}

// GetContactDetails returns the ContactDetails field value if set, zero value otherwise.
func (o *AdvisorySiemensPublisher) GetContactDetails() string {
	if o == nil || IsNil(o.ContactDetails) {
		var ret string
		return ret
	}
	return *o.ContactDetails
}

// GetContactDetailsOk returns a tuple with the ContactDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySiemensPublisher) GetContactDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.ContactDetails) {
		return nil, false
	}
	return o.ContactDetails, true
}

// HasContactDetails returns a boolean if a field has been set.
func (o *AdvisorySiemensPublisher) HasContactDetails() bool {
	if o != nil && !IsNil(o.ContactDetails) {
		return true
	}

	return false
}

// SetContactDetails gets a reference to the given string and assigns it to the ContactDetails field.
func (o *AdvisorySiemensPublisher) SetContactDetails(v string) {
	o.ContactDetails = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdvisorySiemensPublisher) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySiemensPublisher) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdvisorySiemensPublisher) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdvisorySiemensPublisher) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *AdvisorySiemensPublisher) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisorySiemensPublisher) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *AdvisorySiemensPublisher) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *AdvisorySiemensPublisher) SetNamespace(v string) {
	o.Namespace = &v
}

func (o AdvisorySiemensPublisher) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisorySiemensPublisher) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.ContactDetails) {
		toSerialize["contact_details"] = o.ContactDetails
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	return toSerialize, nil
}

type NullableAdvisorySiemensPublisher struct {
	value *AdvisorySiemensPublisher
	isSet bool
}

func (v NullableAdvisorySiemensPublisher) Get() *AdvisorySiemensPublisher {
	return v.value
}

func (v *NullableAdvisorySiemensPublisher) Set(val *AdvisorySiemensPublisher) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisorySiemensPublisher) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisorySiemensPublisher) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisorySiemensPublisher(val *AdvisorySiemensPublisher) *NullableAdvisorySiemensPublisher {
	return &NullableAdvisorySiemensPublisher{value: val, isSet: true}
}

func (v NullableAdvisorySiemensPublisher) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisorySiemensPublisher) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


