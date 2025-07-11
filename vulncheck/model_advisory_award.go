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

// checks if the AdvisoryAward type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryAward{}

// AdvisoryAward struct for AdvisoryAward
type AdvisoryAward struct {
	Amount *string `json:"amount,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

// NewAdvisoryAward instantiates a new AdvisoryAward object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryAward() *AdvisoryAward {
	this := AdvisoryAward{}
	return &this
}

// NewAdvisoryAwardWithDefaults instantiates a new AdvisoryAward object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryAwardWithDefaults() *AdvisoryAward {
	this := AdvisoryAward{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AdvisoryAward) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAward) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AdvisoryAward) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AdvisoryAward) SetAmount(v string) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AdvisoryAward) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAward) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AdvisoryAward) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AdvisoryAward) SetCurrency(v string) {
	o.Currency = &v
}

func (o AdvisoryAward) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryAward) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableAdvisoryAward struct {
	value *AdvisoryAward
	isSet bool
}

func (v NullableAdvisoryAward) Get() *AdvisoryAward {
	return v.value
}

func (v *NullableAdvisoryAward) Set(val *AdvisoryAward) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryAward) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryAward) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryAward(val *AdvisoryAward) *NullableAdvisoryAward {
	return &NullableAdvisoryAward{value: val, isSet: true}
}

func (v NullableAdvisoryAward) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryAward) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


