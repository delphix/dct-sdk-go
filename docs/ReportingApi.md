# \ReportingApi

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVirtualizationStorageSummaryReport**](ReportingApi.md#GetVirtualizationStorageSummaryReport) | **Get** /reporting/virtualization-storage-summary-report | Gets the storage summary report for virtualization engines.



## GetVirtualizationStorageSummaryReport

> VirtualizationStorageSummaryReportResponse GetVirtualizationStorageSummaryReport(ctx).Execute()

Gets the storage summary report for virtualization engines.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.GetVirtualizationStorageSummaryReport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.GetVirtualizationStorageSummaryReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualizationStorageSummaryReport`: VirtualizationStorageSummaryReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.GetVirtualizationStorageSummaryReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualizationStorageSummaryReportRequest struct via the builder pattern


### Return type

[**VirtualizationStorageSummaryReportResponse**](VirtualizationStorageSummaryReportResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

