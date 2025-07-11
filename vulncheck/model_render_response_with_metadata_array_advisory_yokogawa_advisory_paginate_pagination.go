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

// checks if the RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination{}

// RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination struct for RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination
type RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination struct {
	Benchmark *float32 `json:"_benchmark,omitempty"`
	Meta *PaginatePagination `json:"_meta,omitempty"`
	Data []AdvisoryYokogawaAdvisory `json:"data,omitempty"`
}

// NewRenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination instantiates a new RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination() *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination {
	this := RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination{}
	return &this
}

// NewRenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePaginationWithDefaults instantiates a new RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePaginationWithDefaults() *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination {
	this := RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination{}
	return &this
}

// GetBenchmark returns the Benchmark field value if set, zero value otherwise.
func (o *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) GetBenchmark() float32 {
	if o == nil || IsNil(o.Benchmark) {
		var ret float32
		return ret
	}
	return *o.Benchmark
}

// GetBenchmarkOk returns a tuple with the Benchmark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) GetBenchmarkOk() (*float32, bool) {
	if o == nil || IsNil(o.Benchmark) {
		return nil, false
	}
	return o.Benchmark, true
}

// HasBenchmark returns a boolean if a field has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) HasBenchmark() bool {
	if o != nil && !IsNil(o.Benchmark) {
		return true
	}

	return false
}

// SetBenchmark gets a reference to the given float32 and assigns it to the Benchmark field.
func (o *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) SetBenchmark(v float32) {
	o.Benchmark = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) GetMeta() PaginatePagination {
	if o == nil || IsNil(o.Meta) {
		var ret PaginatePagination
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) GetMetaOk() (*PaginatePagination, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginatePagination and assigns it to the Meta field.
func (o *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) SetMeta(v PaginatePagination) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) GetData() []AdvisoryYokogawaAdvisory {
	if o == nil || IsNil(o.Data) {
		var ret []AdvisoryYokogawaAdvisory
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) GetDataOk() ([]AdvisoryYokogawaAdvisory, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AdvisoryYokogawaAdvisory and assigns it to the Data field.
func (o *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) SetData(v []AdvisoryYokogawaAdvisory) {
	o.Data = v
}

func (o RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Benchmark) {
		toSerialize["_benchmark"] = o.Benchmark
	}
	if !IsNil(o.Meta) {
		toSerialize["_meta"] = o.Meta
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableRenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination struct {
	value *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination
	isSet bool
}

func (v NullableRenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) Get() *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination {
	return v.value
}

func (v *NullableRenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) Set(val *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) {
	v.value = val
	v.isSet = true
}

func (v NullableRenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) IsSet() bool {
	return v.isSet
}

func (v *NullableRenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination(val *RenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) *NullableRenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination {
	return &NullableRenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination{value: val, isSet: true}
}

func (v NullableRenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenderResponseWithMetadataArrayAdvisoryYokogawaAdvisoryPaginatePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


