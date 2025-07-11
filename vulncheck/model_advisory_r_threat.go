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

// checks if the AdvisoryRThreat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryRThreat{}

// AdvisoryRThreat struct for AdvisoryRThreat
type AdvisoryRThreat struct {
	Date *string `json:"Date,omitempty"`
	DateSpecified *bool `json:"DateSpecified,omitempty"`
	Description *AdvisoryIVal `json:"Description,omitempty"`
	ProductID []string `json:"ProductID,omitempty"`
	// diff
	Type *int32 `json:"Type,omitempty"`
}

// NewAdvisoryRThreat instantiates a new AdvisoryRThreat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryRThreat() *AdvisoryRThreat {
	this := AdvisoryRThreat{}
	return &this
}

// NewAdvisoryRThreatWithDefaults instantiates a new AdvisoryRThreat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryRThreatWithDefaults() *AdvisoryRThreat {
	this := AdvisoryRThreat{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *AdvisoryRThreat) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRThreat) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *AdvisoryRThreat) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *AdvisoryRThreat) SetDate(v string) {
	o.Date = &v
}

// GetDateSpecified returns the DateSpecified field value if set, zero value otherwise.
func (o *AdvisoryRThreat) GetDateSpecified() bool {
	if o == nil || IsNil(o.DateSpecified) {
		var ret bool
		return ret
	}
	return *o.DateSpecified
}

// GetDateSpecifiedOk returns a tuple with the DateSpecified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRThreat) GetDateSpecifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.DateSpecified) {
		return nil, false
	}
	return o.DateSpecified, true
}

// HasDateSpecified returns a boolean if a field has been set.
func (o *AdvisoryRThreat) HasDateSpecified() bool {
	if o != nil && !IsNil(o.DateSpecified) {
		return true
	}

	return false
}

// SetDateSpecified gets a reference to the given bool and assigns it to the DateSpecified field.
func (o *AdvisoryRThreat) SetDateSpecified(v bool) {
	o.DateSpecified = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AdvisoryRThreat) GetDescription() AdvisoryIVal {
	if o == nil || IsNil(o.Description) {
		var ret AdvisoryIVal
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRThreat) GetDescriptionOk() (*AdvisoryIVal, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AdvisoryRThreat) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given AdvisoryIVal and assigns it to the Description field.
func (o *AdvisoryRThreat) SetDescription(v AdvisoryIVal) {
	o.Description = &v
}

// GetProductID returns the ProductID field value if set, zero value otherwise.
func (o *AdvisoryRThreat) GetProductID() []string {
	if o == nil || IsNil(o.ProductID) {
		var ret []string
		return ret
	}
	return o.ProductID
}

// GetProductIDOk returns a tuple with the ProductID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRThreat) GetProductIDOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductID) {
		return nil, false
	}
	return o.ProductID, true
}

// HasProductID returns a boolean if a field has been set.
func (o *AdvisoryRThreat) HasProductID() bool {
	if o != nil && !IsNil(o.ProductID) {
		return true
	}

	return false
}

// SetProductID gets a reference to the given []string and assigns it to the ProductID field.
func (o *AdvisoryRThreat) SetProductID(v []string) {
	o.ProductID = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AdvisoryRThreat) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryRThreat) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AdvisoryRThreat) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *AdvisoryRThreat) SetType(v int32) {
	o.Type = &v
}

func (o AdvisoryRThreat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryRThreat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["Date"] = o.Date
	}
	if !IsNil(o.DateSpecified) {
		toSerialize["DateSpecified"] = o.DateSpecified
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.ProductID) {
		toSerialize["ProductID"] = o.ProductID
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAdvisoryRThreat struct {
	value *AdvisoryRThreat
	isSet bool
}

func (v NullableAdvisoryRThreat) Get() *AdvisoryRThreat {
	return v.value
}

func (v *NullableAdvisoryRThreat) Set(val *AdvisoryRThreat) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryRThreat) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryRThreat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryRThreat(val *AdvisoryRThreat) *NullableAdvisoryRThreat {
	return &NullableAdvisoryRThreat{value: val, isSet: true}
}

func (v NullableAdvisoryRThreat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryRThreat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


