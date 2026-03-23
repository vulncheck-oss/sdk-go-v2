# \AdvisoryAPI

All URIs are relative to *https://api.vulncheck.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdvisoryGet**](AdvisoryAPI.md#AdvisoryGet) | **Get** /advisory | Query advisories
[**AdvisoryListGet**](AdvisoryAPI.md#AdvisoryListGet) | **Get** /advisory/list | List advisory feeds



## AdvisoryGet

> SearchV4AdvisoryReturnValue AdvisoryGet(ctx).Name(name).CveId(cveId).Vendor(vendor).Product(product).Platform(platform).Version(version).Cpe(cpe).PackageName(packageName).Purl(purl).ReferenceUrl(referenceUrl).ReferenceTag(referenceTag).DescriptionLang(descriptionLang).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Page(page).Limit(limit).StartCursor(startCursor).Cursor(cursor).Execute()

Query advisories



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
	name := "name_example" // string | Filter by advisory feed name (e.g. 'vulncheck') (optional)
	cveId := "cveId_example" // string | Filter by CVE ID (e.g. 'CVE-2024-1234') (optional)
	vendor := "vendor_example" // string | Filter by vendor name (optional)
	product := "product_example" // string | Filter by product name (optional)
	platform := "platform_example" // string | Filter by OS/platform (optional)
	version := "version_example" // string | Filter by product version (semver-aware) (optional)
	cpe := "cpe_example" // string | Filter by CPE (e.g. 'cpe:2.3:a:vendor:product:*') (optional)
	packageName := "packageName_example" // string | Filter by package name (optional)
	purl := "purl_example" // string | Filter by package URL (PURL) (optional)
	referenceUrl := "referenceUrl_example" // string | Filter by reference URL (optional)
	referenceTag := "referenceTag_example" // string | Filter by reference tag (e.g. 'patch', 'advisory') (optional)
	descriptionLang := "descriptionLang_example" // string | Filter by description language (e.g. 'en') (optional)
	updatedAfter := "updatedAfter_example" // string | Return advisories updated after this date (RFC3339 or date-math e.g. 'now-30d') (optional)
	updatedBefore := "updatedBefore_example" // string | Return advisories updated before this date (RFC3339 or date-math) (optional)
	page := int32(56) // int32 | Page number (default: 1) (optional)
	limit := int32(56) // int32 | Results per page, max 100 (default: 10) (optional)
	startCursor := "startCursor_example" // string | Presence activates cursor mode from the first page (value is ignored; cannot be combined with page) (optional)
	cursor := "cursor_example" // string | Cursor from previous response _meta.next_cursor to fetch the next page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvisoryAPI.AdvisoryGet(context.Background()).Name(name).CveId(cveId).Vendor(vendor).Product(product).Platform(platform).Version(version).Cpe(cpe).PackageName(packageName).Purl(purl).ReferenceUrl(referenceUrl).ReferenceTag(referenceTag).DescriptionLang(descriptionLang).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Page(page).Limit(limit).StartCursor(startCursor).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvisoryAPI.AdvisoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdvisoryGet`: SearchV4AdvisoryReturnValue
	fmt.Fprintf(os.Stdout, "Response from `AdvisoryAPI.AdvisoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdvisoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by advisory feed name (e.g. &#39;vulncheck&#39;) | 
 **cveId** | **string** | Filter by CVE ID (e.g. &#39;CVE-2024-1234&#39;) | 
 **vendor** | **string** | Filter by vendor name | 
 **product** | **string** | Filter by product name | 
 **platform** | **string** | Filter by OS/platform | 
 **version** | **string** | Filter by product version (semver-aware) | 
 **cpe** | **string** | Filter by CPE (e.g. &#39;cpe:2.3:a:vendor:product:*&#39;) | 
 **packageName** | **string** | Filter by package name | 
 **purl** | **string** | Filter by package URL (PURL) | 
 **referenceUrl** | **string** | Filter by reference URL | 
 **referenceTag** | **string** | Filter by reference tag (e.g. &#39;patch&#39;, &#39;advisory&#39;) | 
 **descriptionLang** | **string** | Filter by description language (e.g. &#39;en&#39;) | 
 **updatedAfter** | **string** | Return advisories updated after this date (RFC3339 or date-math e.g. &#39;now-30d&#39;) | 
 **updatedBefore** | **string** | Return advisories updated before this date (RFC3339 or date-math) | 
 **page** | **int32** | Page number (default: 1) | 
 **limit** | **int32** | Results per page, max 100 (default: 10) | 
 **startCursor** | **string** | Presence activates cursor mode from the first page (value is ignored; cannot be combined with page) | 
 **cursor** | **string** | Cursor from previous response _meta.next_cursor to fetch the next page | 

### Return type

[**SearchV4AdvisoryReturnValue**](SearchV4AdvisoryReturnValue.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdvisoryListGet

> SearchV4ListFeedReturnValue AdvisoryListGet(ctx).Execute()

List advisory feeds



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvisoryAPI.AdvisoryListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvisoryAPI.AdvisoryListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdvisoryListGet`: SearchV4ListFeedReturnValue
	fmt.Fprintf(os.Stdout, "Response from `AdvisoryAPI.AdvisoryListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdvisoryListGetRequest struct via the builder pattern


### Return type

[**SearchV4ListFeedReturnValue**](SearchV4ListFeedReturnValue.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

