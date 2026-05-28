# RenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Benchmark** | Pointer to **float32** | Benchmark is the server-side processing time for the request in seconds. Example: 0.122322 &#x3D; approximately 122 milliseconds | [optional] 
**Meta** | Pointer to [**PaginatePagination**](PaginatePagination.md) |  | [optional] 
**Data** | Pointer to [**[]ApiNVD20CVE**](ApiNVD20CVE.md) | Data is the data returned by the endpoint | [optional] 

## Methods

### NewRenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination

`func NewRenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination() *RenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination`

NewRenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination instantiates a new RenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderResponseWithMetadataArrayApiNVD20CVEPaginatePaginationWithDefaults

`func NewRenderResponseWithMetadataArrayApiNVD20CVEPaginatePaginationWithDefaults() *RenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination`

NewRenderResponseWithMetadataArrayApiNVD20CVEPaginatePaginationWithDefaults instantiates a new RenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmark

`func (o *RenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination) GetBenchmark() float32`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *RenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination) GetBenchmarkOk() (*float32, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *RenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination) SetBenchmark(v float32)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *RenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetMeta

`func (o *RenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination) GetMeta() PaginatePagination`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination) GetMetaOk() (*PaginatePagination, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination) SetMeta(v PaginatePagination)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *RenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination) GetData() []ApiNVD20CVE`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination) GetDataOk() (*[]ApiNVD20CVE, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination) SetData(v []ApiNVD20CVE)`

SetData sets Data field to given value.

### HasData

`func (o *RenderResponseWithMetadataArrayApiNVD20CVEPaginatePagination) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


