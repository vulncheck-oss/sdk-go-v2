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

// checks if the RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination{}

// RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination struct for RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination
type RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination struct {
	Benchmark *float32 `json:"_benchmark,omitempty"`
	Meta *PaginatePagination `json:"_meta,omitempty"`
	Data []AdvisoryNetskope `json:"data,omitempty"`
}

// NewRenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination instantiates a new RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination() *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination {
	this := RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination{}
	return &this
}

// NewRenderResponseWithMetadataArrayAdvisoryNetskopePaginatePaginationWithDefaults instantiates a new RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenderResponseWithMetadataArrayAdvisoryNetskopePaginatePaginationWithDefaults() *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination {
	this := RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination{}
	return &this
}

// GetBenchmark returns the Benchmark field value if set, zero value otherwise.
func (o *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) GetBenchmark() float32 {
	if o == nil || IsNil(o.Benchmark) {
		var ret float32
		return ret
	}
	return *o.Benchmark
}

// GetBenchmarkOk returns a tuple with the Benchmark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) GetBenchmarkOk() (*float32, bool) {
	if o == nil || IsNil(o.Benchmark) {
		return nil, false
	}
	return o.Benchmark, true
}

// HasBenchmark returns a boolean if a field has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) HasBenchmark() bool {
	if o != nil && !IsNil(o.Benchmark) {
		return true
	}

	return false
}

// SetBenchmark gets a reference to the given float32 and assigns it to the Benchmark field.
func (o *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) SetBenchmark(v float32) {
	o.Benchmark = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) GetMeta() PaginatePagination {
	if o == nil || IsNil(o.Meta) {
		var ret PaginatePagination
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) GetMetaOk() (*PaginatePagination, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginatePagination and assigns it to the Meta field.
func (o *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) SetMeta(v PaginatePagination) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) GetData() []AdvisoryNetskope {
	if o == nil || IsNil(o.Data) {
		var ret []AdvisoryNetskope
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) GetDataOk() ([]AdvisoryNetskope, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AdvisoryNetskope and assigns it to the Data field.
func (o *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) SetData(v []AdvisoryNetskope) {
	o.Data = v
}

func (o RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) ToMap() (map[string]interface{}, error) {
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

type NullableRenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination struct {
	value *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination
	isSet bool
}

func (v NullableRenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) Get() *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination {
	return v.value
}

func (v *NullableRenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) Set(val *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) {
	v.value = val
	v.isSet = true
}

func (v NullableRenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) IsSet() bool {
	return v.isSet
}

func (v *NullableRenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination(val *RenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) *NullableRenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination {
	return &NullableRenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination{value: val, isSet: true}
}

func (v NullableRenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenderResponseWithMetadataArrayAdvisoryNetskopePaginatePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


