# RenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Benchmark** | Pointer to **float32** | Benchmark is the server-side processing time for the request in seconds. Example: 0.122322 &#x3D; approximately 122 milliseconds | [optional] 
**Meta** | Pointer to [**PaginatePagination**](PaginatePagination.md) |  | [optional] 
**Data** | Pointer to [**[]AdvisoryDraegerAdvisory**](AdvisoryDraegerAdvisory.md) | Data is the data returned by the endpoint | [optional] 

## Methods

### NewRenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination

`func NewRenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination() *RenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination`

NewRenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination instantiates a new RenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePaginationWithDefaults

`func NewRenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePaginationWithDefaults() *RenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination`

NewRenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePaginationWithDefaults instantiates a new RenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmark

`func (o *RenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination) GetBenchmark() float32`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *RenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination) GetBenchmarkOk() (*float32, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *RenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination) SetBenchmark(v float32)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *RenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetMeta

`func (o *RenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination) GetMeta() PaginatePagination`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination) GetMetaOk() (*PaginatePagination, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination) SetMeta(v PaginatePagination)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *RenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination) GetData() []AdvisoryDraegerAdvisory`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination) GetDataOk() (*[]AdvisoryDraegerAdvisory, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination) SetData(v []AdvisoryDraegerAdvisory)`

SetData sets Data field to given value.

### HasData

`func (o *RenderResponseWithMetadataArrayAdvisoryDraegerAdvisoryPaginatePagination) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


