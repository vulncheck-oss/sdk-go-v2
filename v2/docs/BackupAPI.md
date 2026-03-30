# \BackupAPI

All URIs are relative to *https://api.vulncheck.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V4GetBackupByName**](BackupAPI.md#V4GetBackupByName) | **Get** /v4/backup/{index} | Get backup by feed name
[**V4ListBackups**](BackupAPI.md#V4ListBackups) | **Get** /v4/backup | List available backups



## V4GetBackupByName

> BackupBackupResponse V4GetBackupByName(ctx, index).Execute()

Get backup by feed name



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
	index := "index_example" // string | Feed name (e.g. 'vulncheck')

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.V4GetBackupByName(context.Background(), index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.V4GetBackupByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V4GetBackupByName`: BackupBackupResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.V4GetBackupByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Feed name (e.g. &#39;vulncheck&#39;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV4GetBackupByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BackupBackupResponse**](BackupBackupResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4ListBackups

> BackupListBackupsResponse V4ListBackups(ctx).Execute()

List available backups



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
	resp, r, err := apiClient.BackupAPI.V4ListBackups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.V4ListBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V4ListBackups`: BackupListBackupsResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.V4ListBackups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV4ListBackupsRequest struct via the builder pattern


### Return type

[**BackupListBackupsResponse**](BackupListBackupsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

