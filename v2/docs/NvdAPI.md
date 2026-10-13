# \NvdAPI

All URIs are relative to *https://api.vulncheck.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NvdCpeMatchGet**](NvdAPI.md#NvdCpeMatchGet) | **Get** /rest/json/cpematch/2.0 | Search CPE match criteria (NVD API 2.0 compatible)
[**NvdCpesGet**](NvdAPI.md#NvdCpesGet) | **Get** /rest/json/cpes/2.0 | Search the CPE dictionary (NVD API 2.0 compatible)
[**NvdCveHistoryGet**](NvdAPI.md#NvdCveHistoryGet) | **Get** /rest/json/cvehistory/2.0 | Search CVE change history (NVD API 2.0 compatible)
[**NvdCvesGet**](NvdAPI.md#NvdCvesGet) | **Get** /rest/json/cves/2.0 | Search CVEs (NVD API 2.0 compatible)
[**NvdSourceGet**](NvdAPI.md#NvdSourceGet) | **Get** /rest/json/source/2.0 | List data sources (NVD API 2.0 compatible)



## NvdCpeMatchGet

> RestNvdCpeMatchResponse NvdCpeMatchGet(ctx).CveId(cveId).MatchCriteriaId(matchCriteriaId).LastModStartDate(lastModStartDate).LastModEndDate(lastModEndDate).ResultsPerPage(resultsPerPage).StartIndex(startIndex).Cursor(cursor).Execute()

Search CPE match criteria (NVD API 2.0 compatible)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vulncheck-oss/sdk-go-v2/v2"
)

func main() {
	cveId := "cveId_example" // string | Only match criteria referenced by this CVE's configurations, e.g. CVE-2021-44228 (optional)
	matchCriteriaId := "matchCriteriaId_example" // string | A single match criteria by its NVD UUID (optional)
	lastModStartDate := "lastModStartDate_example" // string | Last modified on or after, ISO-8601 (e.g. 2026-01-01T00:00:00.000). Requires lastModEndDate; range at most 120 days (optional)
	lastModEndDate := "lastModEndDate_example" // string | Last modified on or before, ISO-8601. Requires lastModStartDate (optional)
	resultsPerPage := int32(56) // int32 | Page size (optional) (default to 500)
	startIndex := int32(56) // int32 | Zero-based offset of the first result. Must not be combined with a non-empty cursor (optional) (default to 0)
	cursor := "cursor_example" // string | VulnCheck extension: opt in to cursor paging. Send it empty for the first page, then the previous response's nextCursor. Responses carry nextCursor only when this is sent (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NvdAPI.NvdCpeMatchGet(context.Background()).CveId(cveId).MatchCriteriaId(matchCriteriaId).LastModStartDate(lastModStartDate).LastModEndDate(lastModEndDate).ResultsPerPage(resultsPerPage).StartIndex(startIndex).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NvdAPI.NvdCpeMatchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NvdCpeMatchGet`: RestNvdCpeMatchResponse
	fmt.Fprintf(os.Stdout, "Response from `NvdAPI.NvdCpeMatchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNvdCpeMatchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cveId** | **string** | Only match criteria referenced by this CVE&#39;s configurations, e.g. CVE-2021-44228 | 
 **matchCriteriaId** | **string** | A single match criteria by its NVD UUID | 
 **lastModStartDate** | **string** | Last modified on or after, ISO-8601 (e.g. 2026-01-01T00:00:00.000). Requires lastModEndDate; range at most 120 days | 
 **lastModEndDate** | **string** | Last modified on or before, ISO-8601. Requires lastModStartDate | 
 **resultsPerPage** | **int32** | Page size | [default to 500]
 **startIndex** | **int32** | Zero-based offset of the first result. Must not be combined with a non-empty cursor | [default to 0]
 **cursor** | **string** | VulnCheck extension: opt in to cursor paging. Send it empty for the first page, then the previous response&#39;s nextCursor. Responses carry nextCursor only when this is sent | 

### Return type

[**RestNvdCpeMatchResponse**](RestNvdCpeMatchResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NvdCpesGet

> RestNvdCpesResponse NvdCpesGet(ctx).CpeNameId(cpeNameId).CpeMatchString(cpeMatchString).MatchCriteriaId(matchCriteriaId).LastModStartDate(lastModStartDate).LastModEndDate(lastModEndDate).ResultsPerPage(resultsPerPage).StartIndex(startIndex).Cursor(cursor).Execute()

Search the CPE dictionary (NVD API 2.0 compatible)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vulncheck-oss/sdk-go-v2/v2"
)

func main() {
	cpeNameId := "cpeNameId_example" // string | A single dictionary entry by its NVD UUID (optional)
	cpeMatchString := "cpeMatchString_example" // string | A CPE 2.3 name or prefix, e.g. cpe:2.3:a:apache:log4j. CPE 2.2 URIs are not supported (optional)
	matchCriteriaId := "matchCriteriaId_example" // string | Only the CPE names this match criteria UUID resolves to (optional)
	lastModStartDate := "lastModStartDate_example" // string | Last modified on or after, ISO-8601 (e.g. 2026-01-01T00:00:00.000). Requires lastModEndDate; range at most 120 days (optional)
	lastModEndDate := "lastModEndDate_example" // string | Last modified on or before, ISO-8601. Requires lastModStartDate (optional)
	resultsPerPage := int32(56) // int32 | Page size (optional) (default to 10000)
	startIndex := int32(56) // int32 | Zero-based offset of the first result. Must not be combined with a non-empty cursor (optional) (default to 0)
	cursor := "cursor_example" // string | VulnCheck extension: opt in to cursor paging. Send it empty for the first page, then the previous response's nextCursor. Responses carry nextCursor only when this is sent (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NvdAPI.NvdCpesGet(context.Background()).CpeNameId(cpeNameId).CpeMatchString(cpeMatchString).MatchCriteriaId(matchCriteriaId).LastModStartDate(lastModStartDate).LastModEndDate(lastModEndDate).ResultsPerPage(resultsPerPage).StartIndex(startIndex).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NvdAPI.NvdCpesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NvdCpesGet`: RestNvdCpesResponse
	fmt.Fprintf(os.Stdout, "Response from `NvdAPI.NvdCpesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNvdCpesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cpeNameId** | **string** | A single dictionary entry by its NVD UUID | 
 **cpeMatchString** | **string** | A CPE 2.3 name or prefix, e.g. cpe:2.3:a:apache:log4j. CPE 2.2 URIs are not supported | 
 **matchCriteriaId** | **string** | Only the CPE names this match criteria UUID resolves to | 
 **lastModStartDate** | **string** | Last modified on or after, ISO-8601 (e.g. 2026-01-01T00:00:00.000). Requires lastModEndDate; range at most 120 days | 
 **lastModEndDate** | **string** | Last modified on or before, ISO-8601. Requires lastModStartDate | 
 **resultsPerPage** | **int32** | Page size | [default to 10000]
 **startIndex** | **int32** | Zero-based offset of the first result. Must not be combined with a non-empty cursor | [default to 0]
 **cursor** | **string** | VulnCheck extension: opt in to cursor paging. Send it empty for the first page, then the previous response&#39;s nextCursor. Responses carry nextCursor only when this is sent | 

### Return type

[**RestNvdCpesResponse**](RestNvdCpesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NvdCveHistoryGet

> RestNvdCveHistoryResponse NvdCveHistoryGet(ctx).CveId(cveId).EventName(eventName).ChangeStartDate(changeStartDate).ChangeEndDate(changeEndDate).ResultsPerPage(resultsPerPage).StartIndex(startIndex).Cursor(cursor).Execute()

Search CVE change history (NVD API 2.0 compatible)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vulncheck-oss/sdk-go-v2/v2"
)

func main() {
	cveId := "cveId_example" // string | Only changes to this CVE, e.g. CVE-2021-44228 (optional)
	eventName := "eventName_example" // string | Only changes of this type. One value per request (optional)
	changeStartDate := "changeStartDate_example" // string | Changed on or after, ISO-8601 (e.g. 2026-01-01T00:00:00.000). Requires changeEndDate; range at most 120 days (optional)
	changeEndDate := "changeEndDate_example" // string | Changed on or before, ISO-8601. Requires changeStartDate (optional)
	resultsPerPage := int32(56) // int32 | Page size (optional) (default to 5000)
	startIndex := int32(56) // int32 | Zero-based offset of the first result. Must not be combined with a non-empty cursor (optional) (default to 0)
	cursor := "cursor_example" // string | VulnCheck extension: opt in to cursor paging. Send it empty for the first page, then the previous response's nextCursor. Responses carry nextCursor only when this is sent (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NvdAPI.NvdCveHistoryGet(context.Background()).CveId(cveId).EventName(eventName).ChangeStartDate(changeStartDate).ChangeEndDate(changeEndDate).ResultsPerPage(resultsPerPage).StartIndex(startIndex).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NvdAPI.NvdCveHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NvdCveHistoryGet`: RestNvdCveHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `NvdAPI.NvdCveHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNvdCveHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cveId** | **string** | Only changes to this CVE, e.g. CVE-2021-44228 | 
 **eventName** | **string** | Only changes of this type. One value per request | 
 **changeStartDate** | **string** | Changed on or after, ISO-8601 (e.g. 2026-01-01T00:00:00.000). Requires changeEndDate; range at most 120 days | 
 **changeEndDate** | **string** | Changed on or before, ISO-8601. Requires changeStartDate | 
 **resultsPerPage** | **int32** | Page size | [default to 5000]
 **startIndex** | **int32** | Zero-based offset of the first result. Must not be combined with a non-empty cursor | [default to 0]
 **cursor** | **string** | VulnCheck extension: opt in to cursor paging. Send it empty for the first page, then the previous response&#39;s nextCursor. Responses carry nextCursor only when this is sent | 

### Return type

[**RestNvdCveHistoryResponse**](RestNvdCveHistoryResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NvdCvesGet

> RestNvdCvesResponse NvdCvesGet(ctx).CveId(cveId).CveIds(cveIds).CpeName(cpeName).IsVulnerable(isVulnerable).VirtualMatchString(virtualMatchString).VersionStart(versionStart).VersionStartType(versionStartType).VersionEnd(versionEnd).VersionEndType(versionEndType).VulnStatuses(vulnStatuses).CveTag(cveTag).NoRejected(noRejected).HasKev(hasKev).VcVulnerableCPEs(vcVulnerableCPEs).PubStartDate(pubStartDate).PubEndDate(pubEndDate).LastModStartDate(lastModStartDate).LastModEndDate(lastModEndDate).SourceIdentifier(sourceIdentifier).ResultsPerPage(resultsPerPage).StartIndex(startIndex).Cursor(cursor).Execute()

Search CVEs (NVD API 2.0 compatible)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vulncheck-oss/sdk-go-v2/v2"
)

func main() {
	cveId := "cveId_example" // string | A single CVE ID, e.g. CVE-2021-44228 (optional)
	cveIds := "cveIds_example" // string | Comma-separated CVE IDs, at most 100 (optional)
	cpeName := "cpeName_example" // string | A complete CPE 2.3 name; returns CVEs whose configurations match it. Takes precedence over virtualMatchString (optional)
	isVulnerable := true // bool | Only CVEs where cpeName is marked vulnerable. Requires cpeName (optional)
	virtualMatchString := "virtualMatchString_example" // string | A CPE 2.3 name or prefix (at least part and vendor) matched against configuration criteria (optional)
	versionStart := "versionStart_example" // string | Lower version bound. Requires virtualMatchString and versionStartType (optional)
	versionStartType := "versionStartType_example" // string | Whether versionStart is inclusive. Requires versionStart (optional)
	versionEnd := "versionEnd_example" // string | Upper version bound. Requires virtualMatchString and versionEndType (optional)
	versionEndType := "versionEndType_example" // string | Whether versionEnd is inclusive. Requires versionEnd (optional)
	vulnStatuses := "vulnStatuses_example" // string | Comma-separated vulnStatus values: UndergoingAnalysis, Modified, AwaitingAnalysis, Rejected, Received, Analyzed, Deferred (optional)
	cveTag := "cveTag_example" // string | Only CVEs carrying this tag (optional)
	noRejected := true // bool | Exclude rejected CVEs (optional)
	hasKev := true // bool | Only CVEs in the CISA KEV catalog (optional)
	vcVulnerableCPEs := true // bool | VulnCheck extension: include the flat vcVulnerableCPEs list on each record. Off by default; it can be most of a page's size (optional)
	pubStartDate := "pubStartDate_example" // string | Published on or after, ISO-8601 (e.g. 2026-01-01T00:00:00.000). Requires pubEndDate; range at most 120 days (optional)
	pubEndDate := "pubEndDate_example" // string | Published on or before, ISO-8601. Requires pubStartDate (optional)
	lastModStartDate := "lastModStartDate_example" // string | Last modified on or after, ISO-8601. Requires lastModEndDate; range at most 120 days (optional)
	lastModEndDate := "lastModEndDate_example" // string | Last modified on or before, ISO-8601. Requires lastModStartDate (optional)
	sourceIdentifier := "sourceIdentifier_example" // string | Only CVEs from this source, e.g. cve@mitre.org (optional)
	resultsPerPage := int32(56) // int32 | Page size (optional) (default to 2000)
	startIndex := int32(56) // int32 | Zero-based offset of the first result. Must not be combined with a non-empty cursor (optional) (default to 0)
	cursor := "cursor_example" // string | VulnCheck extension: opt in to cursor paging. Send it empty for the first page, then the previous response's nextCursor. Responses carry nextCursor only when this is sent (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NvdAPI.NvdCvesGet(context.Background()).CveId(cveId).CveIds(cveIds).CpeName(cpeName).IsVulnerable(isVulnerable).VirtualMatchString(virtualMatchString).VersionStart(versionStart).VersionStartType(versionStartType).VersionEnd(versionEnd).VersionEndType(versionEndType).VulnStatuses(vulnStatuses).CveTag(cveTag).NoRejected(noRejected).HasKev(hasKev).VcVulnerableCPEs(vcVulnerableCPEs).PubStartDate(pubStartDate).PubEndDate(pubEndDate).LastModStartDate(lastModStartDate).LastModEndDate(lastModEndDate).SourceIdentifier(sourceIdentifier).ResultsPerPage(resultsPerPage).StartIndex(startIndex).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NvdAPI.NvdCvesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NvdCvesGet`: RestNvdCvesResponse
	fmt.Fprintf(os.Stdout, "Response from `NvdAPI.NvdCvesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNvdCvesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cveId** | **string** | A single CVE ID, e.g. CVE-2021-44228 | 
 **cveIds** | **string** | Comma-separated CVE IDs, at most 100 | 
 **cpeName** | **string** | A complete CPE 2.3 name; returns CVEs whose configurations match it. Takes precedence over virtualMatchString | 
 **isVulnerable** | **bool** | Only CVEs where cpeName is marked vulnerable. Requires cpeName | 
 **virtualMatchString** | **string** | A CPE 2.3 name or prefix (at least part and vendor) matched against configuration criteria | 
 **versionStart** | **string** | Lower version bound. Requires virtualMatchString and versionStartType | 
 **versionStartType** | **string** | Whether versionStart is inclusive. Requires versionStart | 
 **versionEnd** | **string** | Upper version bound. Requires virtualMatchString and versionEndType | 
 **versionEndType** | **string** | Whether versionEnd is inclusive. Requires versionEnd | 
 **vulnStatuses** | **string** | Comma-separated vulnStatus values: UndergoingAnalysis, Modified, AwaitingAnalysis, Rejected, Received, Analyzed, Deferred | 
 **cveTag** | **string** | Only CVEs carrying this tag | 
 **noRejected** | **bool** | Exclude rejected CVEs | 
 **hasKev** | **bool** | Only CVEs in the CISA KEV catalog | 
 **vcVulnerableCPEs** | **bool** | VulnCheck extension: include the flat vcVulnerableCPEs list on each record. Off by default; it can be most of a page&#39;s size | 
 **pubStartDate** | **string** | Published on or after, ISO-8601 (e.g. 2026-01-01T00:00:00.000). Requires pubEndDate; range at most 120 days | 
 **pubEndDate** | **string** | Published on or before, ISO-8601. Requires pubStartDate | 
 **lastModStartDate** | **string** | Last modified on or after, ISO-8601. Requires lastModEndDate; range at most 120 days | 
 **lastModEndDate** | **string** | Last modified on or before, ISO-8601. Requires lastModStartDate | 
 **sourceIdentifier** | **string** | Only CVEs from this source, e.g. cve@mitre.org | 
 **resultsPerPage** | **int32** | Page size | [default to 2000]
 **startIndex** | **int32** | Zero-based offset of the first result. Must not be combined with a non-empty cursor | [default to 0]
 **cursor** | **string** | VulnCheck extension: opt in to cursor paging. Send it empty for the first page, then the previous response&#39;s nextCursor. Responses carry nextCursor only when this is sent | 

### Return type

[**RestNvdCvesResponse**](RestNvdCvesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NvdSourceGet

> RestNvdSourceResponse NvdSourceGet(ctx).SourceIdentifier(sourceIdentifier).LastModStartDate(lastModStartDate).LastModEndDate(lastModEndDate).ResultsPerPage(resultsPerPage).StartIndex(startIndex).Execute()

List data sources (NVD API 2.0 compatible)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vulncheck-oss/sdk-go-v2/v2"
)

func main() {
	sourceIdentifier := "sourceIdentifier_example" // string | Only the source with this identifier, e.g. cve@mitre.org (optional)
	lastModStartDate := "lastModStartDate_example" // string | Last modified on or after, ISO-8601 (e.g. 2026-01-01T00:00:00.000). Requires lastModEndDate; range at most 120 days (optional)
	lastModEndDate := "lastModEndDate_example" // string | Last modified on or before, ISO-8601. Requires lastModStartDate (optional)
	resultsPerPage := int32(56) // int32 | Page size (optional) (default to 1000)
	startIndex := int32(56) // int32 | Zero-based offset of the first result (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NvdAPI.NvdSourceGet(context.Background()).SourceIdentifier(sourceIdentifier).LastModStartDate(lastModStartDate).LastModEndDate(lastModEndDate).ResultsPerPage(resultsPerPage).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NvdAPI.NvdSourceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NvdSourceGet`: RestNvdSourceResponse
	fmt.Fprintf(os.Stdout, "Response from `NvdAPI.NvdSourceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNvdSourceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceIdentifier** | **string** | Only the source with this identifier, e.g. cve@mitre.org | 
 **lastModStartDate** | **string** | Last modified on or after, ISO-8601 (e.g. 2026-01-01T00:00:00.000). Requires lastModEndDate; range at most 120 days | 
 **lastModEndDate** | **string** | Last modified on or before, ISO-8601. Requires lastModStartDate | 
 **resultsPerPage** | **int32** | Page size | [default to 1000]
 **startIndex** | **int32** | Zero-based offset of the first result | [default to 0]

### Return type

[**RestNvdSourceResponse**](RestNvdSourceResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

