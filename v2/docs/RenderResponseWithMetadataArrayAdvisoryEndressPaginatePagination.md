# RenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Benchmark** | Pointer to **float32** | Benchmark is the server-side processing time for the request in seconds. Example: 0.122322 &#x3D; approximately 122 milliseconds | [optional] 
**Meta** | Pointer to [**PaginatePagination**](PaginatePagination.md) |  | [optional] 
**Data** | Pointer to [**[]AdvisoryEndress**](AdvisoryEndress.md) | Data is the data returned by the endpoint | [optional] 

## Methods

### NewRenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination

`func NewRenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination() *RenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination`

NewRenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination instantiates a new RenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderResponseWithMetadataArrayAdvisoryEndressPaginatePaginationWithDefaults

`func NewRenderResponseWithMetadataArrayAdvisoryEndressPaginatePaginationWithDefaults() *RenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination`

NewRenderResponseWithMetadataArrayAdvisoryEndressPaginatePaginationWithDefaults instantiates a new RenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmark

`func (o *RenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination) GetBenchmark() float32`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *RenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination) GetBenchmarkOk() (*float32, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *RenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination) SetBenchmark(v float32)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *RenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetMeta

`func (o *RenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination) GetMeta() PaginatePagination`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination) GetMetaOk() (*PaginatePagination, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination) SetMeta(v PaginatePagination)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *RenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination) GetData() []AdvisoryEndress`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination) GetDataOk() (*[]AdvisoryEndress, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination) SetData(v []AdvisoryEndress)`

SetData sets Data field to given value.

### HasData

`func (o *RenderResponseWithMetadataArrayAdvisoryEndressPaginatePagination) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


