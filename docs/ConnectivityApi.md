# \ConnectivityApi

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectivityCheck**](ConnectivityApi.md#ConnectivityCheck) | **Post** /connectivity/check | Checks connectivity between an engine and a remote host machine on a given port.
[**DatabaseConnectivityCheck**](ConnectivityApi.md#DatabaseConnectivityCheck) | **Post** /database/connectivity/check | Tests the validity of the supplied database credentials, returning an error if unable to connect to the database.



## ConnectivityCheck

> ConnectivityCheckResponse ConnectivityCheck(ctx).ConnectivityCheckParameters(connectivityCheckParameters).Execute()

Checks connectivity between an engine and a remote host machine on a given port.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/delphix/dct-sdk-go"
)

func main() {
    connectivityCheckParameters := *openapiclient.NewConnectivityCheckParameters("engine-123", "test.host.com", NullableInt32(22)) // ConnectivityCheckParameters | The api to check connectivity of engine and a remote host on given port.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectivityApi.ConnectivityCheck(context.Background()).ConnectivityCheckParameters(connectivityCheckParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectivityApi.ConnectivityCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectivityCheck`: ConnectivityCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectivityApi.ConnectivityCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectivityCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectivityCheckParameters** | [**ConnectivityCheckParameters**](ConnectivityCheckParameters.md) | The api to check connectivity of engine and a remote host on given port. | 

### Return type

[**ConnectivityCheckResponse**](ConnectivityCheckResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseConnectivityCheck

> ConnectivityCheckResponse DatabaseConnectivityCheck(ctx).DatabaseConnectivityCheckParameters(databaseConnectivityCheckParameters).Execute()

Tests the validity of the supplied database credentials, returning an error if unable to connect to the database.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/delphix/dct-sdk-go"
)

func main() {
    databaseConnectivityCheckParameters := *openapiclient.NewDatabaseConnectivityCheckParameters("CredentialsType_example", "1-SOURCE-CONFIG-1") // DatabaseConnectivityCheckParameters | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectivityApi.DatabaseConnectivityCheck(context.Background()).DatabaseConnectivityCheckParameters(databaseConnectivityCheckParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectivityApi.DatabaseConnectivityCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseConnectivityCheck`: ConnectivityCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectivityApi.DatabaseConnectivityCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseConnectivityCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseConnectivityCheckParameters** | [**DatabaseConnectivityCheckParameters**](DatabaseConnectivityCheckParameters.md) |  | 

### Return type

[**ConnectivityCheckResponse**](ConnectivityCheckResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

