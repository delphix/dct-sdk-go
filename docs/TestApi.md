# \TestApi

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnableScaleTesting**](TestApi.md#EnableScaleTesting) | **Post** /enable-scale-testing | This is used for performance testing to enable engine and object duplication.



## EnableScaleTesting

> EnableScaleTesting(ctx).EnableScaleTestingRequest(enableScaleTestingRequest).Execute()

This is used for performance testing to enable engine and object duplication.

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
    enableScaleTestingRequest := *openapiclient.NewEnableScaleTestingRequest(int32(123), []string{"EnginesList_example"}, int32(123), int32(123)) // EnableScaleTestingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestApi.EnableScaleTesting(context.Background()).EnableScaleTestingRequest(enableScaleTestingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestApi.EnableScaleTesting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableScaleTestingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableScaleTestingRequest** | [**EnableScaleTestingRequest**](EnableScaleTestingRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

