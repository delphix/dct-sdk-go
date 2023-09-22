# \TestApi

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TimeToUpdateSources**](TestApi.md#TimeToUpdateSources) | **Post** /time-to-update-sources | set sources loop count variable - this is being used during performance testing.



## TimeToUpdateSources

> TimeToUpdateSourcesResponse TimeToUpdateSources(ctx).TimeToUpdateSourcesRequest(timeToUpdateSourcesRequest).Execute()

set sources loop count variable - this is being used during performance testing.

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
    timeToUpdateSourcesRequest := *openapiclient.NewTimeToUpdateSourcesRequest(int32(123), []string{"EnginesList_example"}, int32(123), int32(123)) // TimeToUpdateSourcesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestApi.TimeToUpdateSources(context.Background()).TimeToUpdateSourcesRequest(timeToUpdateSourcesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestApi.TimeToUpdateSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TimeToUpdateSources`: TimeToUpdateSourcesResponse
    fmt.Fprintf(os.Stdout, "Response from `TestApi.TimeToUpdateSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTimeToUpdateSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeToUpdateSourcesRequest** | [**TimeToUpdateSourcesRequest**](TimeToUpdateSourcesRequest.md) |  | 

### Return type

[**TimeToUpdateSourcesResponse**](TimeToUpdateSourcesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

