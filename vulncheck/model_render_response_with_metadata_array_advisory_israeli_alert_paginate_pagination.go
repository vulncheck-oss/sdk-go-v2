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

// checks if the RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination{}

// RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination struct for RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination
type RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination struct {
	Benchmark *float32 `json:"_benchmark,omitempty"`
	Meta *PaginatePagination `json:"_meta,omitempty"`
	Data []AdvisoryIsraeliAlert `json:"data,omitempty"`
}

// NewRenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination instantiates a new RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination() *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination {
	this := RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination{}
	return &this
}

// NewRenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePaginationWithDefaults instantiates a new RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePaginationWithDefaults() *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination {
	this := RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination{}
	return &this
}

// GetBenchmark returns the Benchmark field value if set, zero value otherwise.
func (o *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) GetBenchmark() float32 {
	if o == nil || IsNil(o.Benchmark) {
		var ret float32
		return ret
	}
	return *o.Benchmark
}

// GetBenchmarkOk returns a tuple with the Benchmark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) GetBenchmarkOk() (*float32, bool) {
	if o == nil || IsNil(o.Benchmark) {
		return nil, false
	}
	return o.Benchmark, true
}

// HasBenchmark returns a boolean if a field has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) HasBenchmark() bool {
	if o != nil && !IsNil(o.Benchmark) {
		return true
	}

	return false
}

// SetBenchmark gets a reference to the given float32 and assigns it to the Benchmark field.
func (o *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) SetBenchmark(v float32) {
	o.Benchmark = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) GetMeta() PaginatePagination {
	if o == nil || IsNil(o.Meta) {
		var ret PaginatePagination
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) GetMetaOk() (*PaginatePagination, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginatePagination and assigns it to the Meta field.
func (o *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) SetMeta(v PaginatePagination) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) GetData() []AdvisoryIsraeliAlert {
	if o == nil || IsNil(o.Data) {
		var ret []AdvisoryIsraeliAlert
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) GetDataOk() ([]AdvisoryIsraeliAlert, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AdvisoryIsraeliAlert and assigns it to the Data field.
func (o *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) SetData(v []AdvisoryIsraeliAlert) {
	o.Data = v
}

func (o RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) ToMap() (map[string]interface{}, error) {
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

type NullableRenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination struct {
	value *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination
	isSet bool
}

func (v NullableRenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) Get() *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination {
	return v.value
}

func (v *NullableRenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) Set(val *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) {
	v.value = val
	v.isSet = true
}

func (v NullableRenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) IsSet() bool {
	return v.isSet
}

func (v *NullableRenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination(val *RenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) *NullableRenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination {
	return &NullableRenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination{value: val, isSet: true}
}

func (v NullableRenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenderResponseWithMetadataArrayAdvisoryIsraeliAlertPaginatePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


