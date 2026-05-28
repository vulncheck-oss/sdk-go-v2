# \EndpointsAPI

All URIs are relative to *https://api.vulncheck.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupGet**](EndpointsAPI.md#BackupGet) | **Get** /v3/backup | Return a list of indexes with backup and endpoint links
[**BackupIndexGet**](EndpointsAPI.md#BackupIndexGet) | **Get** /v3/backup/{index} | Retrieve a list of backups by index
[**CpeGet**](EndpointsAPI.md#CpeGet) | **Get** /v3/cpe | Return CVE &#39;s associated with a specific NIST CPE
[**EntitlementsGet**](EndpointsAPI.md#EntitlementsGet) | **Get** /v3/entitlements | Retrieve user entitlements
[**IndexGet**](EndpointsAPI.md#IndexGet) | **Get** /v3/index | Return a list of available indexes with endpoint links
[**OpenapiGet**](EndpointsAPI.md#OpenapiGet) | **Get** /v3/openapi | Return OpenAPI specification
[**PdnsVulncheckC2Get**](EndpointsAPI.md#PdnsVulncheckC2Get) | **Get** /v3/pdns/vulncheck-c2 | Retrieve a list of C2 Hostnames
[**PurlGet**](EndpointsAPI.md#PurlGet) | **Get** /v3/purl | Request vulnerabilities related to a PURL
[**PurlsPost**](EndpointsAPI.md#PurlsPost) | **Post** /v3/purls | Request vulnerabilities related to a list of PURLs
[**RulesInitialAccessTypeGet**](EndpointsAPI.md#RulesInitialAccessTypeGet) | **Get** /v3/rules/initial-access/{type} | Retrieve set of initial-access detection rules
[**SearchCpeGet**](EndpointsAPI.md#SearchCpeGet) | **Get** /v3/search/cpe | Return CPEs and associated CPEs by searching CPE fields
[**SearchCveGet**](EndpointsAPI.md#SearchCveGet) | **Get** /v3/search/cve | Search all indices for a CVE
[**TagsVulncheckC2Get**](EndpointsAPI.md#TagsVulncheckC2Get) | **Get** /v3/tags/vulncheck-c2 | Retrieve a list of C2 IP addresses



## BackupGet

> RenderResponseArrayParamsIndexBackupList BackupGet(ctx).Execute()

Return a list of indexes with backup and endpoint links



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
	resp, r, err := apiClient.EndpointsAPI.BackupGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.BackupGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupGet`: RenderResponseArrayParamsIndexBackupList
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.BackupGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBackupGetRequest struct via the builder pattern


### Return type

[**RenderResponseArrayParamsIndexBackupList**](RenderResponseArrayParamsIndexBackupList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupIndexGet

> RenderResponseWithMetadataV3controllersBackupResponseDataV3controllersBackupResponseMetadata BackupIndexGet(ctx, index).Execute()

Retrieve a list of backups by index



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
	index := "index_example" // string | Name of an exploit, vulnerability, or advisory index

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.BackupIndexGet(context.Background(), index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.BackupIndexGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupIndexGet`: RenderResponseWithMetadataV3controllersBackupResponseDataV3controllersBackupResponseMetadata
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.BackupIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Name of an exploit, vulnerability, or advisory index | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RenderResponseWithMetadataV3controllersBackupResponseDataV3controllersBackupResponseMetadata**](RenderResponseWithMetadataV3controllersBackupResponseDataV3controllersBackupResponseMetadata.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CpeGet

> RenderResponseWithMetadataArrayStringV3controllersResponseMetadata CpeGet(ctx).Cpe(cpe).IsVulnerable(isVulnerable).Execute()

Return CVE 's associated with a specific NIST CPE



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
	cpe := "cpe_example" // string | CPE designation to lookup
	isVulnerable := "isVulnerable_example" // string | Filter by vulnerability status (true/false). Defaults to false if not provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.CpeGet(context.Background()).Cpe(cpe).IsVulnerable(isVulnerable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.CpeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CpeGet`: RenderResponseWithMetadataArrayStringV3controllersResponseMetadata
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.CpeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCpeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cpe** | **string** | CPE designation to lookup | 
 **isVulnerable** | **string** | Filter by vulnerability status (true/false). Defaults to false if not provided. | 

### Return type

[**RenderResponseWithMetadataArrayStringV3controllersResponseMetadata**](RenderResponseWithMetadataArrayStringV3controllersResponseMetadata.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsGet

> ModelsEntitlements EntitlementsGet(ctx).Execute()

Retrieve user entitlements



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
	resp, r, err := apiClient.EndpointsAPI.EntitlementsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.EntitlementsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitlementsGet`: ModelsEntitlements
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.EntitlementsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsGetRequest struct via the builder pattern


### Return type

[**ModelsEntitlements**](ModelsEntitlements.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexGet

> RenderResponseArrayParamsIndexList IndexGet(ctx).Execute()

Return a list of available indexes with endpoint links



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
	resp, r, err := apiClient.EndpointsAPI.IndexGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.IndexGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IndexGet`: RenderResponseArrayParamsIndexList
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.IndexGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIndexGetRequest struct via the builder pattern


### Return type

[**RenderResponseArrayParamsIndexList**](RenderResponseArrayParamsIndexList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenapiGet

> map[string]interface{} OpenapiGet(ctx).Execute()

Return OpenAPI specification



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
	resp, r, err := apiClient.EndpointsAPI.OpenapiGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.OpenapiGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenapiGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.OpenapiGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenapiGetRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PdnsVulncheckC2Get

> string PdnsVulncheckC2Get(ctx).Format(format).Execute()

Retrieve a list of C2 Hostnames



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
	format := "format_example" // string | Format of the Hostnames in the response (Defaults To: text) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.PdnsVulncheckC2Get(context.Background()).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.PdnsVulncheckC2Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PdnsVulncheckC2Get`: string
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.PdnsVulncheckC2Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPdnsVulncheckC2GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Format of the Hostnames in the response (Defaults To: text) | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurlGet

> RenderResponseWithMetadataV3controllersPurlResponseDataV3controllersPurlResponseMetadata PurlGet(ctx).Purl(purl).Execute()

Request vulnerabilities related to a PURL



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
	purl := "purl_example" // string | URL string used to identify and locate a software package

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.PurlGet(context.Background()).Purl(purl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.PurlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurlGet`: RenderResponseWithMetadataV3controllersPurlResponseDataV3controllersPurlResponseMetadata
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.PurlGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurlGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purl** | **string** | URL string used to identify and locate a software package | 

### Return type

[**RenderResponseWithMetadataV3controllersPurlResponseDataV3controllersPurlResponseMetadata**](RenderResponseWithMetadataV3controllersPurlResponseDataV3controllersPurlResponseMetadata.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurlsPost

> RenderResponseWithMetadataV3controllersPurlsResponseDataV3controllersPurlsResponseMetadata PurlsPost(ctx).RequestBody(requestBody).Execute()

Request vulnerabilities related to a list of PURLs



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
	requestBody := []string{"Property_example"} // []string | PURL strings used to identify and locate software packages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.PurlsPost(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.PurlsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurlsPost`: RenderResponseWithMetadataV3controllersPurlsResponseDataV3controllersPurlsResponseMetadata
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.PurlsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurlsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | PURL strings used to identify and locate software packages | 

### Return type

[**RenderResponseWithMetadataV3controllersPurlsResponseDataV3controllersPurlsResponseMetadata**](RenderResponseWithMetadataV3controllersPurlsResponseDataV3controllersPurlsResponseMetadata.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesInitialAccessTypeGet

> string RulesInitialAccessTypeGet(ctx, type_).Execute()

Retrieve set of initial-access detection rules



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
	type_ := "type__example" // string | Type of ruleset to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.RulesInitialAccessTypeGet(context.Background(), type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.RulesInitialAccessTypeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesInitialAccessTypeGet`: string
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.RulesInitialAccessTypeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Type of ruleset to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesInitialAccessTypeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCpeGet

> RenderResponseWithMetadataSearchResponsesSearchResponseMetadata SearchCpeGet(ctx).Part(part).Vendor(vendor).Product(product).Version(version).IsVulnerable(isVulnerable).Execute()

Return CPEs and associated CPEs by searching CPE fields



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
	part := "[a]pplication, [o]perating system, [h]ardware" // string | CPE part to lookup (optional)
	vendor := "vendor_example" // string | CPE vendor to lookup (optional)
	product := "product_example" // string | CPE product to lookup (optional)
	version := "version_example" // string | CPE version to lookup (optional)
	isVulnerable := "isVulnerable_example" // string | Filter by vulnerability status (true/false). Defaults to false if not provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.SearchCpeGet(context.Background()).Part(part).Vendor(vendor).Product(product).Version(version).IsVulnerable(isVulnerable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.SearchCpeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCpeGet`: RenderResponseWithMetadataSearchResponsesSearchResponseMetadata
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.SearchCpeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCpeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **part** | **string** | CPE part to lookup | 
 **vendor** | **string** | CPE vendor to lookup | 
 **product** | **string** | CPE product to lookup | 
 **version** | **string** | CPE version to lookup | 
 **isVulnerable** | **string** | Filter by vulnerability status (true/false). Defaults to false if not provided. | 

### Return type

[**RenderResponseWithMetadataSearchResponsesSearchResponseMetadata**](RenderResponseWithMetadataSearchResponsesSearchResponseMetadata.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCveGet

> RenderResponseWithMetadataArrayIndexCveSearchHitIndexCveSearchMeta SearchCveGet(ctx).Cve(cve).Page(page).Limit(limit).Cursor(cursor).StartCursor(startCursor).Execute()

Search all indices for a CVE



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
	cve := "cve_example" // string | CVE ID to search for (e.g. CVE-2024-1234)
	page := int32(56) // int32 | Page number (default: 1, page mode only) (optional)
	limit := int32(56) // int32 | Maximum number of results per page (default: 500, max: 1000) (optional)
	cursor := "cursor_example" // string | Continue cursor paging, or use an empty value to start cursor paging (optional)
	startCursor := "startCursor_example" // string | Start cursor paging (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.SearchCveGet(context.Background()).Cve(cve).Page(page).Limit(limit).Cursor(cursor).StartCursor(startCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.SearchCveGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCveGet`: RenderResponseWithMetadataArrayIndexCveSearchHitIndexCveSearchMeta
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.SearchCveGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCveGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cve** | **string** | CVE ID to search for (e.g. CVE-2024-1234) | 
 **page** | **int32** | Page number (default: 1, page mode only) | 
 **limit** | **int32** | Maximum number of results per page (default: 500, max: 1000) | 
 **cursor** | **string** | Continue cursor paging, or use an empty value to start cursor paging | 
 **startCursor** | **string** | Start cursor paging | 

### Return type

[**RenderResponseWithMetadataArrayIndexCveSearchHitIndexCveSearchMeta**](RenderResponseWithMetadataArrayIndexCveSearchHitIndexCveSearchMeta.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TagsVulncheckC2Get

> string TagsVulncheckC2Get(ctx).Format(format).Execute()

Retrieve a list of C2 IP addresses



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
	format := "format_example" // string | Format of the IP Addresses in the response (Defaults To: text) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.TagsVulncheckC2Get(context.Background()).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.TagsVulncheckC2Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TagsVulncheckC2Get`: string
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.TagsVulncheckC2Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTagsVulncheckC2GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Format of the IP Addresses in the response (Defaults To: text) | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

