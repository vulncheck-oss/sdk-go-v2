# \EndpointsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupGet**](EndpointsAPI.md#BackupGet) | **Get** /backup | Return a list of indexes with backup and endpoint links
[**BackupIndexGet**](EndpointsAPI.md#BackupIndexGet) | **Get** /backup/{index} | Retrieve a list of backups by index
[**CpeGet**](EndpointsAPI.md#CpeGet) | **Get** /cpe | Return CVE &#39;s associated with a specific NIST CPE
[**EntitlementsGet**](EndpointsAPI.md#EntitlementsGet) | **Get** /entitlements | Retrieve user entitlements
[**IndexGet**](EndpointsAPI.md#IndexGet) | **Get** /index | Return a list of available indexes with endpoint links
[**OpenapiGet**](EndpointsAPI.md#OpenapiGet) | **Get** /openapi | Return OpenAPI specification
[**PdnsVulncheckC2Get**](EndpointsAPI.md#PdnsVulncheckC2Get) | **Get** /pdns/vulncheck-c2 | Retrieve a list of C2 Hostnames
[**PurlGet**](EndpointsAPI.md#PurlGet) | **Get** /purl | Request vulnerabilities related to a PURL
[**RulesInitialAccessTypeGet**](EndpointsAPI.md#RulesInitialAccessTypeGet) | **Get** /rules/initial-access/{type} | Retrieve set of initial-access detection rules
[**TagsVulncheckC2Get**](EndpointsAPI.md#TagsVulncheckC2Get) | **Get** /tags/vulncheck-c2 | Retrieve a list of C2 IP addresses



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
	openapiclient "github.com/vulncheck-oss/sdk-go-v2/vulncheck"
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
	openapiclient "github.com/vulncheck-oss/sdk-go-v2/vulncheck"
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

> RenderResponseWithMetadataArrayStringV3controllersResponseMetadata CpeGet(ctx).Cpe(cpe).Execute()

Return CVE 's associated with a specific NIST CPE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vulncheck-oss/sdk-go-v2/vulncheck"
)

func main() {
	cpe := "cpe_example" // string | CPE designation to lookup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.CpeGet(context.Background()).Cpe(cpe).Execute()
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
	openapiclient "github.com/vulncheck-oss/sdk-go-v2/vulncheck"
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
	openapiclient "github.com/vulncheck-oss/sdk-go-v2/vulncheck"
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
	openapiclient "github.com/vulncheck-oss/sdk-go-v2/vulncheck"
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
	openapiclient "github.com/vulncheck-oss/sdk-go-v2/vulncheck"
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
	openapiclient "github.com/vulncheck-oss/sdk-go-v2/vulncheck"
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
	openapiclient "github.com/vulncheck-oss/sdk-go-v2/vulncheck"
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
	openapiclient "github.com/vulncheck-oss/sdk-go-v2/vulncheck"
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

