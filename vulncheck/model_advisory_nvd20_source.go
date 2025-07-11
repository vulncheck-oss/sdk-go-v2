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

// checks if the AdvisoryNVD20Source type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryNVD20Source{}

// AdvisoryNVD20Source struct for AdvisoryNVD20Source
type AdvisoryNVD20Source struct {
	ContactEmail *string `json:"contactEmail,omitempty"`
	Created *string `json:"created,omitempty"`
	CweAcceptanceLevel *AdvisoryCweAcceptanceLevel `json:"cweAcceptanceLevel,omitempty"`
	LastModified *string `json:"lastModified,omitempty"`
	Name *string `json:"name,omitempty"`
	SourceIdentifiers []string `json:"sourceIdentifiers,omitempty"`
	V3AcceptanceLevel *AdvisoryV3AcceptanceLevel `json:"v3AcceptanceLevel,omitempty"`
}

// NewAdvisoryNVD20Source instantiates a new AdvisoryNVD20Source object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryNVD20Source() *AdvisoryNVD20Source {
	this := AdvisoryNVD20Source{}
	return &this
}

// NewAdvisoryNVD20SourceWithDefaults instantiates a new AdvisoryNVD20Source object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryNVD20SourceWithDefaults() *AdvisoryNVD20Source {
	this := AdvisoryNVD20Source{}
	return &this
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *AdvisoryNVD20Source) GetContactEmail() string {
	if o == nil || IsNil(o.ContactEmail) {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNVD20Source) GetContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ContactEmail) {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *AdvisoryNVD20Source) HasContactEmail() bool {
	if o != nil && !IsNil(o.ContactEmail) {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *AdvisoryNVD20Source) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AdvisoryNVD20Source) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNVD20Source) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AdvisoryNVD20Source) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *AdvisoryNVD20Source) SetCreated(v string) {
	o.Created = &v
}

// GetCweAcceptanceLevel returns the CweAcceptanceLevel field value if set, zero value otherwise.
func (o *AdvisoryNVD20Source) GetCweAcceptanceLevel() AdvisoryCweAcceptanceLevel {
	if o == nil || IsNil(o.CweAcceptanceLevel) {
		var ret AdvisoryCweAcceptanceLevel
		return ret
	}
	return *o.CweAcceptanceLevel
}

// GetCweAcceptanceLevelOk returns a tuple with the CweAcceptanceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNVD20Source) GetCweAcceptanceLevelOk() (*AdvisoryCweAcceptanceLevel, bool) {
	if o == nil || IsNil(o.CweAcceptanceLevel) {
		return nil, false
	}
	return o.CweAcceptanceLevel, true
}

// HasCweAcceptanceLevel returns a boolean if a field has been set.
func (o *AdvisoryNVD20Source) HasCweAcceptanceLevel() bool {
	if o != nil && !IsNil(o.CweAcceptanceLevel) {
		return true
	}

	return false
}

// SetCweAcceptanceLevel gets a reference to the given AdvisoryCweAcceptanceLevel and assigns it to the CweAcceptanceLevel field.
func (o *AdvisoryNVD20Source) SetCweAcceptanceLevel(v AdvisoryCweAcceptanceLevel) {
	o.CweAcceptanceLevel = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *AdvisoryNVD20Source) GetLastModified() string {
	if o == nil || IsNil(o.LastModified) {
		var ret string
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNVD20Source) GetLastModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *AdvisoryNVD20Source) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given string and assigns it to the LastModified field.
func (o *AdvisoryNVD20Source) SetLastModified(v string) {
	o.LastModified = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdvisoryNVD20Source) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNVD20Source) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdvisoryNVD20Source) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdvisoryNVD20Source) SetName(v string) {
	o.Name = &v
}

// GetSourceIdentifiers returns the SourceIdentifiers field value if set, zero value otherwise.
func (o *AdvisoryNVD20Source) GetSourceIdentifiers() []string {
	if o == nil || IsNil(o.SourceIdentifiers) {
		var ret []string
		return ret
	}
	return o.SourceIdentifiers
}

// GetSourceIdentifiersOk returns a tuple with the SourceIdentifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNVD20Source) GetSourceIdentifiersOk() ([]string, bool) {
	if o == nil || IsNil(o.SourceIdentifiers) {
		return nil, false
	}
	return o.SourceIdentifiers, true
}

// HasSourceIdentifiers returns a boolean if a field has been set.
func (o *AdvisoryNVD20Source) HasSourceIdentifiers() bool {
	if o != nil && !IsNil(o.SourceIdentifiers) {
		return true
	}

	return false
}

// SetSourceIdentifiers gets a reference to the given []string and assigns it to the SourceIdentifiers field.
func (o *AdvisoryNVD20Source) SetSourceIdentifiers(v []string) {
	o.SourceIdentifiers = v
}

// GetV3AcceptanceLevel returns the V3AcceptanceLevel field value if set, zero value otherwise.
func (o *AdvisoryNVD20Source) GetV3AcceptanceLevel() AdvisoryV3AcceptanceLevel {
	if o == nil || IsNil(o.V3AcceptanceLevel) {
		var ret AdvisoryV3AcceptanceLevel
		return ret
	}
	return *o.V3AcceptanceLevel
}

// GetV3AcceptanceLevelOk returns a tuple with the V3AcceptanceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryNVD20Source) GetV3AcceptanceLevelOk() (*AdvisoryV3AcceptanceLevel, bool) {
	if o == nil || IsNil(o.V3AcceptanceLevel) {
		return nil, false
	}
	return o.V3AcceptanceLevel, true
}

// HasV3AcceptanceLevel returns a boolean if a field has been set.
func (o *AdvisoryNVD20Source) HasV3AcceptanceLevel() bool {
	if o != nil && !IsNil(o.V3AcceptanceLevel) {
		return true
	}

	return false
}

// SetV3AcceptanceLevel gets a reference to the given AdvisoryV3AcceptanceLevel and assigns it to the V3AcceptanceLevel field.
func (o *AdvisoryNVD20Source) SetV3AcceptanceLevel(v AdvisoryV3AcceptanceLevel) {
	o.V3AcceptanceLevel = &v
}

func (o AdvisoryNVD20Source) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryNVD20Source) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContactEmail) {
		toSerialize["contactEmail"] = o.ContactEmail
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.CweAcceptanceLevel) {
		toSerialize["cweAcceptanceLevel"] = o.CweAcceptanceLevel
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SourceIdentifiers) {
		toSerialize["sourceIdentifiers"] = o.SourceIdentifiers
	}
	if !IsNil(o.V3AcceptanceLevel) {
		toSerialize["v3AcceptanceLevel"] = o.V3AcceptanceLevel
	}
	return toSerialize, nil
}

type NullableAdvisoryNVD20Source struct {
	value *AdvisoryNVD20Source
	isSet bool
}

func (v NullableAdvisoryNVD20Source) Get() *AdvisoryNVD20Source {
	return v.value
}

func (v *NullableAdvisoryNVD20Source) Set(val *AdvisoryNVD20Source) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryNVD20Source) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryNVD20Source) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryNVD20Source(val *AdvisoryNVD20Source) *NullableAdvisoryNVD20Source {
	return &NullableAdvisoryNVD20Source{value: val, isSet: true}
}

func (v NullableAdvisoryNVD20Source) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryNVD20Source) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


