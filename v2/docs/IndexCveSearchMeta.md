# IndexCveSearchMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentCount** | Pointer to **int32** | CurrentCount is the number of documents returned in this response. | [optional] 
**Cve** | Pointer to **string** | CVE is the CVE identifier that was searched for. | [optional] 
**IndicesWithHits** | Pointer to **int32** | IndicesWithHits is the number of indices that contained at least one match. | [optional] 
**Limit** | Pointer to **int32** | Limit is the maximum number of results requested per page. | [optional] 
**NextCursor** | Pointer to **string** | NextCursor is the cursor for fetching the next page of results. Empty if there are no more results (current_count &lt; limit). | [optional] 
**Page** | Pointer to **int32** | Page is the current page number (only set in page-based pagination mode). | [optional] 
**QueriedIndexCount** | Pointer to **int32** | QueriedIndexCount is the number of indices that were searched. | [optional] 
**Timestamp** | Pointer to **string** | Timestamp is the UTC time when the search was executed. | [optional] 
**TopIndexCounts** | Pointer to [**[]IndexIndexCountEntry**](IndexIndexCountEntry.md) | TopIndexCounts contains the top indices by document count (up to 10). | [optional] 
**TotalCount** | Pointer to **int32** | TotalCount is the total number of matching documents across all queried indices. | [optional] 
**TotalPages** | Pointer to **int32** | TotalPages is the total number of pages (only set in page-based pagination mode). | [optional] 

## Methods

### NewIndexCveSearchMeta

`func NewIndexCveSearchMeta() *IndexCveSearchMeta`

NewIndexCveSearchMeta instantiates a new IndexCveSearchMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexCveSearchMetaWithDefaults

`func NewIndexCveSearchMetaWithDefaults() *IndexCveSearchMeta`

NewIndexCveSearchMetaWithDefaults instantiates a new IndexCveSearchMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentCount

`func (o *IndexCveSearchMeta) GetCurrentCount() int32`

GetCurrentCount returns the CurrentCount field if non-nil, zero value otherwise.

### GetCurrentCountOk

`func (o *IndexCveSearchMeta) GetCurrentCountOk() (*int32, bool)`

GetCurrentCountOk returns a tuple with the CurrentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCount

`func (o *IndexCveSearchMeta) SetCurrentCount(v int32)`

SetCurrentCount sets CurrentCount field to given value.

### HasCurrentCount

`func (o *IndexCveSearchMeta) HasCurrentCount() bool`

HasCurrentCount returns a boolean if a field has been set.

### GetCve

`func (o *IndexCveSearchMeta) GetCve() string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *IndexCveSearchMeta) GetCveOk() (*string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *IndexCveSearchMeta) SetCve(v string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *IndexCveSearchMeta) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetIndicesWithHits

`func (o *IndexCveSearchMeta) GetIndicesWithHits() int32`

GetIndicesWithHits returns the IndicesWithHits field if non-nil, zero value otherwise.

### GetIndicesWithHitsOk

`func (o *IndexCveSearchMeta) GetIndicesWithHitsOk() (*int32, bool)`

GetIndicesWithHitsOk returns a tuple with the IndicesWithHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicesWithHits

`func (o *IndexCveSearchMeta) SetIndicesWithHits(v int32)`

SetIndicesWithHits sets IndicesWithHits field to given value.

### HasIndicesWithHits

`func (o *IndexCveSearchMeta) HasIndicesWithHits() bool`

HasIndicesWithHits returns a boolean if a field has been set.

### GetLimit

`func (o *IndexCveSearchMeta) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *IndexCveSearchMeta) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *IndexCveSearchMeta) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *IndexCveSearchMeta) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNextCursor

`func (o *IndexCveSearchMeta) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *IndexCveSearchMeta) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *IndexCveSearchMeta) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *IndexCveSearchMeta) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetPage

`func (o *IndexCveSearchMeta) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *IndexCveSearchMeta) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *IndexCveSearchMeta) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *IndexCveSearchMeta) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetQueriedIndexCount

`func (o *IndexCveSearchMeta) GetQueriedIndexCount() int32`

GetQueriedIndexCount returns the QueriedIndexCount field if non-nil, zero value otherwise.

### GetQueriedIndexCountOk

`func (o *IndexCveSearchMeta) GetQueriedIndexCountOk() (*int32, bool)`

GetQueriedIndexCountOk returns a tuple with the QueriedIndexCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueriedIndexCount

`func (o *IndexCveSearchMeta) SetQueriedIndexCount(v int32)`

SetQueriedIndexCount sets QueriedIndexCount field to given value.

### HasQueriedIndexCount

`func (o *IndexCveSearchMeta) HasQueriedIndexCount() bool`

HasQueriedIndexCount returns a boolean if a field has been set.

### GetTimestamp

`func (o *IndexCveSearchMeta) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *IndexCveSearchMeta) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *IndexCveSearchMeta) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *IndexCveSearchMeta) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTopIndexCounts

`func (o *IndexCveSearchMeta) GetTopIndexCounts() []IndexIndexCountEntry`

GetTopIndexCounts returns the TopIndexCounts field if non-nil, zero value otherwise.

### GetTopIndexCountsOk

`func (o *IndexCveSearchMeta) GetTopIndexCountsOk() (*[]IndexIndexCountEntry, bool)`

GetTopIndexCountsOk returns a tuple with the TopIndexCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopIndexCounts

`func (o *IndexCveSearchMeta) SetTopIndexCounts(v []IndexIndexCountEntry)`

SetTopIndexCounts sets TopIndexCounts field to given value.

### HasTopIndexCounts

`func (o *IndexCveSearchMeta) HasTopIndexCounts() bool`

HasTopIndexCounts returns a boolean if a field has been set.

### GetTotalCount

`func (o *IndexCveSearchMeta) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *IndexCveSearchMeta) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *IndexCveSearchMeta) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *IndexCveSearchMeta) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *IndexCveSearchMeta) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *IndexCveSearchMeta) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *IndexCveSearchMeta) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *IndexCveSearchMeta) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


