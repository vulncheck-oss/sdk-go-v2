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

// checks if the RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination{}

// RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination struct for RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination
type RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination struct {
	Benchmark *float32 `json:"_benchmark,omitempty"`
	Meta *PaginatePagination `json:"_meta,omitempty"`
	Data []AdvisoryApacheArchiva `json:"data,omitempty"`
}

// NewRenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination instantiates a new RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination() *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination {
	this := RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination{}
	return &this
}

// NewRenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePaginationWithDefaults instantiates a new RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePaginationWithDefaults() *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination {
	this := RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination{}
	return &this
}

// GetBenchmark returns the Benchmark field value if set, zero value otherwise.
func (o *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) GetBenchmark() float32 {
	if o == nil || IsNil(o.Benchmark) {
		var ret float32
		return ret
	}
	return *o.Benchmark
}

// GetBenchmarkOk returns a tuple with the Benchmark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) GetBenchmarkOk() (*float32, bool) {
	if o == nil || IsNil(o.Benchmark) {
		return nil, false
	}
	return o.Benchmark, true
}

// HasBenchmark returns a boolean if a field has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) HasBenchmark() bool {
	if o != nil && !IsNil(o.Benchmark) {
		return true
	}

	return false
}

// SetBenchmark gets a reference to the given float32 and assigns it to the Benchmark field.
func (o *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) SetBenchmark(v float32) {
	o.Benchmark = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) GetMeta() PaginatePagination {
	if o == nil || IsNil(o.Meta) {
		var ret PaginatePagination
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) GetMetaOk() (*PaginatePagination, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginatePagination and assigns it to the Meta field.
func (o *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) SetMeta(v PaginatePagination) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) GetData() []AdvisoryApacheArchiva {
	if o == nil || IsNil(o.Data) {
		var ret []AdvisoryApacheArchiva
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) GetDataOk() ([]AdvisoryApacheArchiva, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AdvisoryApacheArchiva and assigns it to the Data field.
func (o *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) SetData(v []AdvisoryApacheArchiva) {
	o.Data = v
}

func (o RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) ToMap() (map[string]interface{}, error) {
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

type NullableRenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination struct {
	value *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination
	isSet bool
}

func (v NullableRenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) Get() *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination {
	return v.value
}

func (v *NullableRenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) Set(val *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) {
	v.value = val
	v.isSet = true
}

func (v NullableRenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) IsSet() bool {
	return v.isSet
}

func (v *NullableRenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination(val *RenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) *NullableRenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination {
	return &NullableRenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination{value: val, isSet: true}
}

func (v NullableRenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenderResponseWithMetadataArrayAdvisoryApacheArchivaPaginatePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


