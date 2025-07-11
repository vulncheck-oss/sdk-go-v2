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

// checks if the AdvisoryDocumentNote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryDocumentNote{}

// AdvisoryDocumentNote struct for AdvisoryDocumentNote
type AdvisoryDocumentNote struct {
	Text *string `json:"text,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewAdvisoryDocumentNote instantiates a new AdvisoryDocumentNote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryDocumentNote() *AdvisoryDocumentNote {
	this := AdvisoryDocumentNote{}
	return &this
}

// NewAdvisoryDocumentNoteWithDefaults instantiates a new AdvisoryDocumentNote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryDocumentNoteWithDefaults() *AdvisoryDocumentNote {
	this := AdvisoryDocumentNote{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *AdvisoryDocumentNote) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryDocumentNote) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *AdvisoryDocumentNote) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *AdvisoryDocumentNote) SetText(v string) {
	o.Text = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryDocumentNote) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryDocumentNote) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryDocumentNote) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryDocumentNote) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AdvisoryDocumentNote) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryDocumentNote) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AdvisoryDocumentNote) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AdvisoryDocumentNote) SetType(v string) {
	o.Type = &v
}

func (o AdvisoryDocumentNote) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryDocumentNote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAdvisoryDocumentNote struct {
	value *AdvisoryDocumentNote
	isSet bool
}

func (v NullableAdvisoryDocumentNote) Get() *AdvisoryDocumentNote {
	return v.value
}

func (v *NullableAdvisoryDocumentNote) Set(val *AdvisoryDocumentNote) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryDocumentNote) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryDocumentNote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryDocumentNote(val *AdvisoryDocumentNote) *NullableAdvisoryDocumentNote {
	return &NullableAdvisoryDocumentNote{value: val, isSet: true}
}

func (v NullableAdvisoryDocumentNote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryDocumentNote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


