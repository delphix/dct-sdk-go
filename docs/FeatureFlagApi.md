# \FeatureFlagApi

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnabledFeaturesFlag**](FeatureFlagApi.md#GetEnabledFeaturesFlag) | **Get** /enabled-features-flag | Get enabled feature flags.



## GetEnabledFeaturesFlag

> []string GetEnabledFeaturesFlag(ctx).Execute()

Get enabled feature flags.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureFlagApi.GetEnabledFeaturesFlag(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagApi.GetEnabledFeaturesFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnabledFeaturesFlag`: []string
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagApi.GetEnabledFeaturesFlag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnabledFeaturesFlagRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

