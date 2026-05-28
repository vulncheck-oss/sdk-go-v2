# PaginatePagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** | Cursor is an opaque string representing the current page position for cursor-based pagination. Contains encoded information about sort values and position markers. Use this value with the &#39;cursor&#39; query parameter to resume pagination from this point. | [optional] 
**FirstItem** | Pointer to **int32** | FirstItem is the 1-indexed position of the first item on the current page. Example: Page 1 with limit 100 → FirstItem &#x3D; 1          Page 2 with limit 100 → FirstItem &#x3D; 101 | [optional] 
**Index** | Pointer to **string** | Index name being queried, corresponding to the VulnCheck data feed. Examples: \&quot;exploits\&quot;, \&quot;vulncheck-nvd2\&quot;, \&quot;ipintel-3d\&quot;, \&quot;malicious-vscode-exts\&quot; | [optional] 
**LastItem** | Pointer to **int32** | LastItem is the 1-indexed position of the last item on the current page. Accounts for partial pages. Example: Page with 50 items → LastItem &#x3D; FirstItem + 49 On the final page, LastItem equals TotalRows. | [optional] 
**Limit** | Pointer to **int32** | Limit is the maximum number of items returned in this response. Range: 1-100 for standard indices, 1-10 for large indices. Default: 100 (or 10 for indices marked as \&quot;large\&quot; in metadata) | [optional] 
**Matches** | Pointer to [**[]PaginateMatch**](PaginateMatch.md) | Matches shows the active filter values applied to the current query. Each match includes the field name and the value being filtered on. Helps users understand which filters are currently active. | [optional] 
**MaxPages** | Pointer to **int32** | MaxPages is the maximum number of pages allowed for offset-based pagination. Performance limit to prevent expensive deep pagination. Default: 6 pages. When TotalPages &gt; MaxPages, cursor-based pagination should be used instead. | [optional] 
**NextCursor** | Pointer to **string** | NextCursor is an opaque string for fetching the next page in cursor-based pagination. When empty or null, indicates no more pages are available. More efficient than offset pagination for large datasets and real-time data. | [optional] 
**OpensearchQuery** | Pointer to **map[string]interface{}** | OSQuery contains the raw OpenSearch query JSON used internally. Only included when show_query&#x3D;true parameter is set. Useful for debugging complex queries and understanding performance. | [optional] 
**Order** | Pointer to **string** | Order specifies the sort direction: \&quot;asc\&quot; (ascending) or \&quot;desc\&quot; (descending). Default: \&quot;desc\&quot; for most time-based indices to show newest items first. | [optional] 
**Page** | Pointer to **int32** | Page is the current page number for offset-based pagination (1-indexed). First page &#x3D; 1, second page &#x3D; 2, etc. Used with Limit to calculate database offset. For cursor-based pagination, this field may be omitted or estimated. | [optional] 
**Pages** | Pointer to **[]string** | Pages contains page numbers to display in pagination UI (e.g., [1, 2, \&quot;...\&quot;, 45]). Only populated when ShowPages&#x3D;true. Includes ellipsis (\&quot;...\&quot;) for gaps. Limited by MaxPages to prevent overwhelming UI with too many page links. | [optional] 
**Parameters** | Pointer to [**[]PaginateParam**](PaginateParam.md) | Params describes available query parameters for this index. Each parameter includes name, format constraints, and filtering capabilities. Used for building dynamic query UIs and input validation. | [optional] 
**ShowPages** | Pointer to **bool** | ShowPages controls whether the Pages array should be calculated and included. Set to true for UI pagination controls, false for API-only usage to save processing. | [optional] 
**ShowQuery** | Pointer to **bool** | ShowQuery indicates whether the raw OpenSearch query should be included. When true, OSQuery field will be populated with the internal query structure. | [optional] 
**Sort** | Pointer to **string** | Sort field used for ordering results. Available fields vary by index. Common values: \&quot;_timestamp\&quot;, \&quot;date_added\&quot;, \&quot;_id\&quot;, \&quot;lastSeen\&quot; Affects cursor generation and ensures pagination consistency. | [optional] 
**Timestamp** | Pointer to **string** | Timestamp when the query was executed in UTC (ISO 8601 format). Used for debugging, cache management, and determining data freshness. Example: \&quot;2024-02-23T20:35:43.732591251Z\&quot; | [optional] 
**TotalDocuments** | Pointer to **int32** | TotalRows is the total number of documents matching the query across all pages. Used for calculating total_pages and building pagination UI elements. Note: May be capped for performance on very large result sets. | [optional] 
**TotalPages** | Pointer to **int32** | TotalPages is the total number of pages based on TotalRows and Limit. Calculated as: ceil(TotalRows / Limit). Used for pagination UI and validation. Example: 1000 items with limit 100 &#x3D; 10 total pages. | [optional] 
**Warnings** | Pointer to **[]string** | Warnings contains any non-fatal issues encountered during query processing. Examples: deprecated parameters, performance concerns, or data quality notes. | [optional] 

## Methods

### NewPaginatePagination

`func NewPaginatePagination() *PaginatePagination`

NewPaginatePagination instantiates a new PaginatePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatePaginationWithDefaults

`func NewPaginatePaginationWithDefaults() *PaginatePagination`

NewPaginatePaginationWithDefaults instantiates a new PaginatePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *PaginatePagination) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *PaginatePagination) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *PaginatePagination) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *PaginatePagination) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetFirstItem

`func (o *PaginatePagination) GetFirstItem() int32`

GetFirstItem returns the FirstItem field if non-nil, zero value otherwise.

### GetFirstItemOk

`func (o *PaginatePagination) GetFirstItemOk() (*int32, bool)`

GetFirstItemOk returns a tuple with the FirstItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstItem

`func (o *PaginatePagination) SetFirstItem(v int32)`

SetFirstItem sets FirstItem field to given value.

### HasFirstItem

`func (o *PaginatePagination) HasFirstItem() bool`

HasFirstItem returns a boolean if a field has been set.

### GetIndex

`func (o *PaginatePagination) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *PaginatePagination) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *PaginatePagination) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *PaginatePagination) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetLastItem

`func (o *PaginatePagination) GetLastItem() int32`

GetLastItem returns the LastItem field if non-nil, zero value otherwise.

### GetLastItemOk

`func (o *PaginatePagination) GetLastItemOk() (*int32, bool)`

GetLastItemOk returns a tuple with the LastItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastItem

`func (o *PaginatePagination) SetLastItem(v int32)`

SetLastItem sets LastItem field to given value.

### HasLastItem

`func (o *PaginatePagination) HasLastItem() bool`

HasLastItem returns a boolean if a field has been set.

### GetLimit

`func (o *PaginatePagination) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaginatePagination) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaginatePagination) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PaginatePagination) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMatches

`func (o *PaginatePagination) GetMatches() []PaginateMatch`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *PaginatePagination) GetMatchesOk() (*[]PaginateMatch, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *PaginatePagination) SetMatches(v []PaginateMatch)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *PaginatePagination) HasMatches() bool`

HasMatches returns a boolean if a field has been set.

### GetMaxPages

`func (o *PaginatePagination) GetMaxPages() int32`

GetMaxPages returns the MaxPages field if non-nil, zero value otherwise.

### GetMaxPagesOk

`func (o *PaginatePagination) GetMaxPagesOk() (*int32, bool)`

GetMaxPagesOk returns a tuple with the MaxPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPages

`func (o *PaginatePagination) SetMaxPages(v int32)`

SetMaxPages sets MaxPages field to given value.

### HasMaxPages

`func (o *PaginatePagination) HasMaxPages() bool`

HasMaxPages returns a boolean if a field has been set.

### GetNextCursor

`func (o *PaginatePagination) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *PaginatePagination) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *PaginatePagination) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *PaginatePagination) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetOpensearchQuery

`func (o *PaginatePagination) GetOpensearchQuery() map[string]interface{}`

GetOpensearchQuery returns the OpensearchQuery field if non-nil, zero value otherwise.

### GetOpensearchQueryOk

`func (o *PaginatePagination) GetOpensearchQueryOk() (*map[string]interface{}, bool)`

GetOpensearchQueryOk returns a tuple with the OpensearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpensearchQuery

`func (o *PaginatePagination) SetOpensearchQuery(v map[string]interface{})`

SetOpensearchQuery sets OpensearchQuery field to given value.

### HasOpensearchQuery

`func (o *PaginatePagination) HasOpensearchQuery() bool`

HasOpensearchQuery returns a boolean if a field has been set.

### GetOrder

`func (o *PaginatePagination) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PaginatePagination) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PaginatePagination) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PaginatePagination) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPage

`func (o *PaginatePagination) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PaginatePagination) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PaginatePagination) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PaginatePagination) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *PaginatePagination) GetPages() []string`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *PaginatePagination) GetPagesOk() (*[]string, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *PaginatePagination) SetPages(v []string)`

SetPages sets Pages field to given value.

### HasPages

`func (o *PaginatePagination) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetParameters

`func (o *PaginatePagination) GetParameters() []PaginateParam`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PaginatePagination) GetParametersOk() (*[]PaginateParam, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PaginatePagination) SetParameters(v []PaginateParam)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *PaginatePagination) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetShowPages

`func (o *PaginatePagination) GetShowPages() bool`

GetShowPages returns the ShowPages field if non-nil, zero value otherwise.

### GetShowPagesOk

`func (o *PaginatePagination) GetShowPagesOk() (*bool, bool)`

GetShowPagesOk returns a tuple with the ShowPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPages

`func (o *PaginatePagination) SetShowPages(v bool)`

SetShowPages sets ShowPages field to given value.

### HasShowPages

`func (o *PaginatePagination) HasShowPages() bool`

HasShowPages returns a boolean if a field has been set.

### GetShowQuery

`func (o *PaginatePagination) GetShowQuery() bool`

GetShowQuery returns the ShowQuery field if non-nil, zero value otherwise.

### GetShowQueryOk

`func (o *PaginatePagination) GetShowQueryOk() (*bool, bool)`

GetShowQueryOk returns a tuple with the ShowQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowQuery

`func (o *PaginatePagination) SetShowQuery(v bool)`

SetShowQuery sets ShowQuery field to given value.

### HasShowQuery

`func (o *PaginatePagination) HasShowQuery() bool`

HasShowQuery returns a boolean if a field has been set.

### GetSort

`func (o *PaginatePagination) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PaginatePagination) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PaginatePagination) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PaginatePagination) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTimestamp

`func (o *PaginatePagination) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PaginatePagination) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PaginatePagination) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PaginatePagination) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTotalDocuments

`func (o *PaginatePagination) GetTotalDocuments() int32`

GetTotalDocuments returns the TotalDocuments field if non-nil, zero value otherwise.

### GetTotalDocumentsOk

`func (o *PaginatePagination) GetTotalDocumentsOk() (*int32, bool)`

GetTotalDocumentsOk returns a tuple with the TotalDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDocuments

`func (o *PaginatePagination) SetTotalDocuments(v int32)`

SetTotalDocuments sets TotalDocuments field to given value.

### HasTotalDocuments

`func (o *PaginatePagination) HasTotalDocuments() bool`

HasTotalDocuments returns a boolean if a field has been set.

### GetTotalPages

`func (o *PaginatePagination) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PaginatePagination) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PaginatePagination) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PaginatePagination) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetWarnings

`func (o *PaginatePagination) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *PaginatePagination) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *PaginatePagination) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *PaginatePagination) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


