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

// checks if the ApiMitreAttackTech type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMitreAttackTech{}

// ApiMitreAttackTech struct for ApiMitreAttackTech
type ApiMitreAttackTech struct {
	Domain *string `json:"domain,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Subtechnique *bool `json:"subtechnique,omitempty"`
	Tactics []string `json:"tactics,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewApiMitreAttackTech instantiates a new ApiMitreAttackTech object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMitreAttackTech() *ApiMitreAttackTech {
	this := ApiMitreAttackTech{}
	return &this
}

// NewApiMitreAttackTechWithDefaults instantiates a new ApiMitreAttackTech object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMitreAttackTechWithDefaults() *ApiMitreAttackTech {
	this := ApiMitreAttackTech{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ApiMitreAttackTech) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMitreAttackTech) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ApiMitreAttackTech) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ApiMitreAttackTech) SetDomain(v string) {
	o.Domain = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiMitreAttackTech) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMitreAttackTech) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiMitreAttackTech) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiMitreAttackTech) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiMitreAttackTech) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMitreAttackTech) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiMitreAttackTech) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiMitreAttackTech) SetName(v string) {
	o.Name = &v
}

// GetSubtechnique returns the Subtechnique field value if set, zero value otherwise.
func (o *ApiMitreAttackTech) GetSubtechnique() bool {
	if o == nil || IsNil(o.Subtechnique) {
		var ret bool
		return ret
	}
	return *o.Subtechnique
}

// GetSubtechniqueOk returns a tuple with the Subtechnique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMitreAttackTech) GetSubtechniqueOk() (*bool, bool) {
	if o == nil || IsNil(o.Subtechnique) {
		return nil, false
	}
	return o.Subtechnique, true
}

// HasSubtechnique returns a boolean if a field has been set.
func (o *ApiMitreAttackTech) HasSubtechnique() bool {
	if o != nil && !IsNil(o.Subtechnique) {
		return true
	}

	return false
}

// SetSubtechnique gets a reference to the given bool and assigns it to the Subtechnique field.
func (o *ApiMitreAttackTech) SetSubtechnique(v bool) {
	o.Subtechnique = &v
}

// GetTactics returns the Tactics field value if set, zero value otherwise.
func (o *ApiMitreAttackTech) GetTactics() []string {
	if o == nil || IsNil(o.Tactics) {
		var ret []string
		return ret
	}
	return o.Tactics
}

// GetTacticsOk returns a tuple with the Tactics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMitreAttackTech) GetTacticsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tactics) {
		return nil, false
	}
	return o.Tactics, true
}

// HasTactics returns a boolean if a field has been set.
func (o *ApiMitreAttackTech) HasTactics() bool {
	if o != nil && !IsNil(o.Tactics) {
		return true
	}

	return false
}

// SetTactics gets a reference to the given []string and assigns it to the Tactics field.
func (o *ApiMitreAttackTech) SetTactics(v []string) {
	o.Tactics = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ApiMitreAttackTech) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMitreAttackTech) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ApiMitreAttackTech) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ApiMitreAttackTech) SetUrl(v string) {
	o.Url = &v
}

func (o ApiMitreAttackTech) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMitreAttackTech) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Subtechnique) {
		toSerialize["subtechnique"] = o.Subtechnique
	}
	if !IsNil(o.Tactics) {
		toSerialize["tactics"] = o.Tactics
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableApiMitreAttackTech struct {
	value *ApiMitreAttackTech
	isSet bool
}

func (v NullableApiMitreAttackTech) Get() *ApiMitreAttackTech {
	return v.value
}

func (v *NullableApiMitreAttackTech) Set(val *ApiMitreAttackTech) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMitreAttackTech) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMitreAttackTech) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMitreAttackTech(val *ApiMitreAttackTech) *NullableApiMitreAttackTech {
	return &NullableApiMitreAttackTech{value: val, isSet: true}
}

func (v NullableApiMitreAttackTech) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMitreAttackTech) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


