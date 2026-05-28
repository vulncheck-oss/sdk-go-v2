# RenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Benchmark** | Pointer to **float32** | Benchmark is the server-side processing time for the request in seconds. Example: 0.122322 &#x3D; approximately 122 milliseconds | [optional] 
**Meta** | Pointer to [**PaginatePagination**](PaginatePagination.md) |  | [optional] 
**Data** | Pointer to [**[]AdvisoryApacheCouchDB**](AdvisoryApacheCouchDB.md) | Data is the data returned by the endpoint | [optional] 

## Methods

### NewRenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination

`func NewRenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination() *RenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination`

NewRenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination instantiates a new RenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePaginationWithDefaults

`func NewRenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePaginationWithDefaults() *RenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination`

NewRenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePaginationWithDefaults instantiates a new RenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmark

`func (o *RenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination) GetBenchmark() float32`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *RenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination) GetBenchmarkOk() (*float32, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *RenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination) SetBenchmark(v float32)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *RenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetMeta

`func (o *RenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination) GetMeta() PaginatePagination`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination) GetMetaOk() (*PaginatePagination, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination) SetMeta(v PaginatePagination)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *RenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination) GetData() []AdvisoryApacheCouchDB`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination) GetDataOk() (*[]AdvisoryApacheCouchDB, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination) SetData(v []AdvisoryApacheCouchDB)`

SetData sets Data field to given value.

### HasData

`func (o *RenderResponseWithMetadataArrayAdvisoryApacheCouchDBPaginatePagination) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


