# \EnginesApi

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectivityCheck**](EnginesApi.md#ConnectivityCheck) | **Post** /engines/connectivity/check | Checks connectivity between an engine and a remote host on a given port.



## ConnectivityCheck

> EngineConnectivityCheckResponse ConnectivityCheck(ctx).EngineConnectivityCheckRequest(engineConnectivityCheckRequest).Execute()

Checks connectivity between an engine and a remote host on a given port.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    engineConnectivityCheckRequest := *openapiclient.NewEngineConnectivityCheckRequest("EngineId_example", "Host_example", NullableInt32(123)) // EngineConnectivityCheckRequest | The api to check connectivity of engine and a remote host on given port.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnginesApi.ConnectivityCheck(context.Background()).EngineConnectivityCheckRequest(engineConnectivityCheckRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnginesApi.ConnectivityCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectivityCheck`: EngineConnectivityCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `EnginesApi.ConnectivityCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectivityCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engineConnectivityCheckRequest** | [**EngineConnectivityCheckRequest**](EngineConnectivityCheckRequest.md) | The api to check connectivity of engine and a remote host on given port. | 

### Return type

[**EngineConnectivityCheckResponse**](EngineConnectivityCheckResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

