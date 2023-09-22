# \ReportingApi

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReportingSchedule**](ReportingApi.md#CreateReportingSchedule) | **Post** /reporting/schedule | Create a new report schedule.
[**CreateReportingScheduleTags**](ReportingApi.md#CreateReportingScheduleTags) | **Post** /reporting/schedule/{reportId}/tags | Create tags for a report schedule.
[**DeleteReportingSchedule**](ReportingApi.md#DeleteReportingSchedule) | **Delete** /reporting/schedule/{reportId} | Delete report schedule by ID.
[**DeleteReportingScheduleTag**](ReportingApi.md#DeleteReportingScheduleTag) | **Post** /reporting/schedule/{reportId}/tags/delete | Delete tags for a report schedule.
[**GetApiUsageReport**](ReportingApi.md#GetApiUsageReport) | **Get** /reporting/api-usage-report | Gets the report of API usage metrics over a given time period. This API returns at the most 10000 results in the response to protect against the server running out of memory. Users might not hit this limit with the default report without any &#39;apiUsageReportGroupByParam&#39; param but can hit this limit if the groupBy is by too granular like by &#39;client_name&#39; or &#39;user_agent&#39; only. Hence it is advisable to use startDate and endDate to &#39;limit&#39; the scope of the report.
[**GetDsourceConsumptionReport**](ReportingApi.md#GetDsourceConsumptionReport) | **Get** /reporting/dsource-consumption-report | Gets the consumption report for virtualization engine dSources.
[**GetDsourceUsageReport**](ReportingApi.md#GetDsourceUsageReport) | **Get** /reporting/dsource-usage-report | Gets the usage report for virtualization engine dSources.
[**GetProductInfo**](ReportingApi.md#GetProductInfo) | **Get** /reporting/product_info | Returns the DCT Product Information.
[**GetReportingScheduleById**](ReportingApi.md#GetReportingScheduleById) | **Get** /reporting/schedule/{reportId} | Returns a report schedule by ID.
[**GetReportingScheduleTags**](ReportingApi.md#GetReportingScheduleTags) | **Get** /reporting/schedule/{reportId}/tags | Get tags for a report schedule.
[**GetReportingSchedules**](ReportingApi.md#GetReportingSchedules) | **Get** /reporting/schedule | List all report schedules.
[**GetVdbInventoryReport**](ReportingApi.md#GetVdbInventoryReport) | **Get** /reporting/vdb-inventory-report | Gets the inventory report for virtualization engine VDBs.
[**GetVirtualizationStorageSummaryReport**](ReportingApi.md#GetVirtualizationStorageSummaryReport) | **Get** /reporting/virtualization-storage-summary-report | Gets the storage summary report for virtualization engines.
[**SearchDsourceConsumptionReport**](ReportingApi.md#SearchDsourceConsumptionReport) | **Post** /reporting/dsource-consumption-report/search | Search the consumption report for virtualization engine DSources.
[**SearchDsourceUsageReport**](ReportingApi.md#SearchDsourceUsageReport) | **Post** /reporting/dsource-usage-report/search | Search the usage report for virtualization engine dSources.
[**SearchVdbInventoryReport**](ReportingApi.md#SearchVdbInventoryReport) | **Post** /reporting/vdb-inventory-report/search | Search the inventory report for virtualization engine VDBs.
[**SearchVirtualizationStorageSummaryReport**](ReportingApi.md#SearchVirtualizationStorageSummaryReport) | **Post** /reporting/virtualization-storage-summary-report/search | Search the storage summary report for virtualization engines.
[**UpdateReportingSchedule**](ReportingApi.md#UpdateReportingSchedule) | **Put** /reporting/schedule/{reportId} | Update a reporting schedule by ID.



## CreateReportingSchedule

> ReportingSchedule CreateReportingSchedule(ctx).ReportingScheduleCreateParameters(reportingScheduleCreateParameters).Execute()

Create a new report schedule.

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
    reportingScheduleCreateParameters := *openapiclient.NewReportingScheduleCreateParameters("VIRTUALIZATION_STORAGE_SUMMARY", "0 0 2 ? * Mon-Fri", "CSV", true, []string{"user@example.com"}) // ReportingScheduleCreateParameters | The parameters to create a reporting schedule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.CreateReportingSchedule(context.Background()).ReportingScheduleCreateParameters(reportingScheduleCreateParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.CreateReportingSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReportingSchedule`: ReportingSchedule
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.CreateReportingSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReportingScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportingScheduleCreateParameters** | [**ReportingScheduleCreateParameters**](ReportingScheduleCreateParameters.md) | The parameters to create a reporting schedule. | 

### Return type

[**ReportingSchedule**](ReportingSchedule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReportingScheduleTags

> TagsResponse CreateReportingScheduleTags(ctx, reportId).TagsRequest(tagsRequest).Execute()

Create tags for a report schedule.

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
    reportId := int32(56) // int32 | The ID of the report schedule.
    tagsRequest := *openapiclient.NewTagsRequest([]openapiclient.Tag{*openapiclient.NewTag("key-1", "value-1")}) // TagsRequest | Tags information for report schedule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.CreateReportingScheduleTags(context.Background(), reportId).TagsRequest(tagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.CreateReportingScheduleTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReportingScheduleTags`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.CreateReportingScheduleTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **int32** | The ID of the report schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReportingScheduleTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagsRequest** | [**TagsRequest**](TagsRequest.md) | Tags information for report schedule. | 

### Return type

[**TagsResponse**](TagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReportingSchedule

> DeleteReportingSchedule(ctx, reportId).Execute()

Delete report schedule by ID.

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
    reportId := int32(56) // int32 | The ID of the report schedule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReportingApi.DeleteReportingSchedule(context.Background(), reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.DeleteReportingSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **int32** | The ID of the report schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReportingScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReportingScheduleTag

> DeleteReportingScheduleTag(ctx, reportId).DeleteTag(deleteTag).Execute()

Delete tags for a report schedule.

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
    reportId := int32(56) // int32 | The ID of the report schedule.
    deleteTag := *openapiclient.NewDeleteTag() // DeleteTag | The parameters to delete tags (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReportingApi.DeleteReportingScheduleTag(context.Background(), reportId).DeleteTag(deleteTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.DeleteReportingScheduleTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **int32** | The ID of the report schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReportingScheduleTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteTag** | [**DeleteTag**](DeleteTag.md) | The parameters to delete tags | 

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


## GetApiUsageReport

> ApiUsageReportResponse GetApiUsageReport(ctx).StartDate(startDate).EndDate(endDate).ApiMetricKind(apiMetricKind).GroupBy(groupBy).ClientName(clientName).UserAgent(userAgent).DctVersion(dctVersion).Execute()

Gets the report of API usage metrics over a given time period. This API returns at the most 10000 results in the response to protect against the server running out of memory. Users might not hit this limit with the default report without any 'apiUsageReportGroupByParam' param but can hit this limit if the groupBy is by too granular like by 'client_name' or 'user_agent' only. Hence it is advisable to use startDate and endDate to 'limit' the scope of the report.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/delphix/dct-sdk-go"
)

func main() {
    startDate := time.Now() // time.Time | Report start date/time. Defaults to first API request. (optional)
    endDate := time.Now() // time.Time | Report end date/time. Defaults to current time. (optional)
    apiMetricKind := "automation" // string | Restrict the list to API usage metric of the given kind (optional)
    groupBy := []string{"GroupBy_example"} // []string | The field to group results by. (optional)
    clientName := []string{"Inner_example"} // []string | The Client names to be included in the report. (optional)
    userAgent := []string{"Inner_example"} // []string | The UserAgent names to be included in the report. (optional)
    dctVersion := []string{"Inner_example"} // []string | The DCT versions to be included in the report. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.GetApiUsageReport(context.Background()).StartDate(startDate).EndDate(endDate).ApiMetricKind(apiMetricKind).GroupBy(groupBy).ClientName(clientName).UserAgent(userAgent).DctVersion(dctVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.GetApiUsageReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiUsageReport`: ApiUsageReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.GetApiUsageReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiUsageReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | Report start date/time. Defaults to first API request. | 
 **endDate** | **time.Time** | Report end date/time. Defaults to current time. | 
 **apiMetricKind** | **string** | Restrict the list to API usage metric of the given kind | 
 **groupBy** | **[]string** | The field to group results by. | 
 **clientName** | **[]string** | The Client names to be included in the report. | 
 **userAgent** | **[]string** | The UserAgent names to be included in the report. | 
 **dctVersion** | **[]string** | The DCT versions to be included in the report. | 

### Return type

[**ApiUsageReportResponse**](ApiUsageReportResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDsourceConsumptionReport

> DSourceConsumptionReportResponse GetDsourceConsumptionReport(ctx).Limit(limit).Cursor(cursor).Sort(sort).Execute()

Gets the consumption report for virtualization engine dSources.

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
    limit := int32(100) // int32 | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 10000. (optional) (default to 10000)
    cursor := "cursor_example" // string | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the 'prev_cursor' or 'next_cursor' property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. (optional)
    sort := "name" // string | The field to sort results by. A property name with a prepended '-' signifies descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.GetDsourceConsumptionReport(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.GetDsourceConsumptionReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDsourceConsumptionReport`: DSourceConsumptionReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.GetDsourceConsumptionReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDsourceConsumptionReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 10000. | [default to 10000]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 

### Return type

[**DSourceConsumptionReportResponse**](DSourceConsumptionReportResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDsourceUsageReport

> DSourceUsageReportResponse GetDsourceUsageReport(ctx).Limit(limit).Cursor(cursor).Sort(sort).Execute()

Gets the usage report for virtualization engine dSources.

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
    limit := int32(100) // int32 | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 10000. (optional) (default to 10000)
    cursor := "cursor_example" // string | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the 'prev_cursor' or 'next_cursor' property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. (optional)
    sort := "name" // string | The field to sort results by. A property name with a prepended '-' signifies descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.GetDsourceUsageReport(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.GetDsourceUsageReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDsourceUsageReport`: DSourceUsageReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.GetDsourceUsageReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDsourceUsageReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 10000. | [default to 10000]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 

### Return type

[**DSourceUsageReportResponse**](DSourceUsageReportResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductInfo

> ProductInfo GetProductInfo(ctx).Execute()

Returns the DCT Product Information.

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
    resp, r, err := apiClient.ReportingApi.GetProductInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.GetProductInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductInfo`: ProductInfo
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.GetProductInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductInfoRequest struct via the builder pattern


### Return type

[**ProductInfo**](ProductInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportingScheduleById

> ReportingSchedule GetReportingScheduleById(ctx, reportId).Execute()

Returns a report schedule by ID.

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
    reportId := int32(56) // int32 | The ID of the report schedule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.GetReportingScheduleById(context.Background(), reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.GetReportingScheduleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportingScheduleById`: ReportingSchedule
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.GetReportingScheduleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **int32** | The ID of the report schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportingScheduleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportingSchedule**](ReportingSchedule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportingScheduleTags

> TagsResponse GetReportingScheduleTags(ctx, reportId).Execute()

Get tags for a report schedule.

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
    reportId := int32(56) // int32 | The ID of the report schedule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.GetReportingScheduleTags(context.Background(), reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.GetReportingScheduleTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportingScheduleTags`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.GetReportingScheduleTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **int32** | The ID of the report schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportingScheduleTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagsResponse**](TagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportingSchedules

> ListReportingScheduleResponse GetReportingSchedules(ctx).Limit(limit).Cursor(cursor).Sort(sort).Execute()

List all report schedules.

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
    limit := int32(50) // int32 | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. (optional) (default to 100)
    cursor := "cursor_example" // string | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the 'prev_cursor' or 'next_cursor' property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. (optional)
    sort := "report_id" // string | The field to sort results by. A property name with a prepended '-' signifies descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.GetReportingSchedules(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.GetReportingSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportingSchedules`: ListReportingScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.GetReportingSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportingSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 

### Return type

[**ListReportingScheduleResponse**](ListReportingScheduleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVdbInventoryReport

> VDBInventoryReportResponse GetVdbInventoryReport(ctx).Limit(limit).Cursor(cursor).Sort(sort).Execute()

Gets the inventory report for virtualization engine VDBs.

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
    limit := int32(100) // int32 | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 10000. (optional) (default to 10000)
    cursor := "cursor_example" // string | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the 'prev_cursor' or 'next_cursor' property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. (optional)
    sort := "name" // string | The field to sort results by. A property name with a prepended '-' signifies descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.GetVdbInventoryReport(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.GetVdbInventoryReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVdbInventoryReport`: VDBInventoryReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.GetVdbInventoryReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVdbInventoryReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 10000. | [default to 10000]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 

### Return type

[**VDBInventoryReportResponse**](VDBInventoryReportResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualizationStorageSummaryReport

> VirtualizationStorageSummaryReportResponse GetVirtualizationStorageSummaryReport(ctx).Limit(limit).Cursor(cursor).Sort(sort).Execute()

Gets the storage summary report for virtualization engines.

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
    limit := int32(100) // int32 | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 10000. (optional) (default to 10000)
    cursor := "cursor_example" // string | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the 'prev_cursor' or 'next_cursor' property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. (optional)
    sort := "engine_id" // string | The field to sort results by. A property name with a prepended '-' signifies descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.GetVirtualizationStorageSummaryReport(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.GetVirtualizationStorageSummaryReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualizationStorageSummaryReport`: VirtualizationStorageSummaryReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.GetVirtualizationStorageSummaryReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualizationStorageSummaryReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 10000. | [default to 10000]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 

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


## SearchDsourceConsumptionReport

> DSourceConsumptionReportResponse SearchDsourceConsumptionReport(ctx).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()

Search the consumption report for virtualization engine DSources.

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
    limit := int32(100) // int32 | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 10000. (optional) (default to 10000)
    cursor := "cursor_example" // string | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the 'prev_cursor' or 'next_cursor' property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. (optional)
    sort := "name" // string | The field to sort results by. A property name with a prepended '-' signifies descending order. (optional)
    searchBody := *openapiclient.NewSearchBody() // SearchBody | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS 'foobar', field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN ['Goku', 'Vegeta'] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically `SEARCH '12'` would match an item with an attribute with an integer value of `123`.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ 'Goku' | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ 'Goku' |  ## Grouping Parenthesis `()` can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS 'foo')  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \"foo\", \"bar\", \"foo bar\", 'foo', 'bar', 'foo bar' | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], ['foo', \"bar\"] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.SearchDsourceConsumptionReport(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.SearchDsourceConsumptionReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchDsourceConsumptionReport`: DSourceConsumptionReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.SearchDsourceConsumptionReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDsourceConsumptionReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 10000. | [default to 10000]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 
 **searchBody** | [**SearchBody**](SearchBody.md) | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS &#39;foobar&#39;, field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN [&#39;Goku&#39;, &#39;Vegeta&#39;] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically &#x60;SEARCH &#39;12&#39;&#x60; would match an item with an attribute with an integer value of &#x60;123&#x60;.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ &#39;Goku&#39; | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ &#39;Goku&#39; |  ## Grouping Parenthesis &#x60;()&#x60; can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS &#39;foo&#39;)  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \&quot;foo\&quot;, \&quot;bar\&quot;, \&quot;foo bar\&quot;, &#39;foo&#39;, &#39;bar&#39;, &#39;foo bar&#39; | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], [&#39;foo&#39;, \&quot;bar\&quot;] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  | 

### Return type

[**DSourceConsumptionReportResponse**](DSourceConsumptionReportResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDsourceUsageReport

> DSourceUsageReportResponse SearchDsourceUsageReport(ctx).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()

Search the usage report for virtualization engine dSources.

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
    limit := int32(100) // int32 | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 10000. (optional) (default to 10000)
    cursor := "cursor_example" // string | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the 'prev_cursor' or 'next_cursor' property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. (optional)
    sort := "name" // string | The field to sort results by. A property name with a prepended '-' signifies descending order. (optional)
    searchBody := *openapiclient.NewSearchBody() // SearchBody | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS 'foobar', field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN ['Goku', 'Vegeta'] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically `SEARCH '12'` would match an item with an attribute with an integer value of `123`.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ 'Goku' | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ 'Goku' |  ## Grouping Parenthesis `()` can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS 'foo')  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \"foo\", \"bar\", \"foo bar\", 'foo', 'bar', 'foo bar' | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], ['foo', \"bar\"] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.SearchDsourceUsageReport(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.SearchDsourceUsageReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchDsourceUsageReport`: DSourceUsageReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.SearchDsourceUsageReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDsourceUsageReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 10000. | [default to 10000]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 
 **searchBody** | [**SearchBody**](SearchBody.md) | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS &#39;foobar&#39;, field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN [&#39;Goku&#39;, &#39;Vegeta&#39;] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically &#x60;SEARCH &#39;12&#39;&#x60; would match an item with an attribute with an integer value of &#x60;123&#x60;.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ &#39;Goku&#39; | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ &#39;Goku&#39; |  ## Grouping Parenthesis &#x60;()&#x60; can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS &#39;foo&#39;)  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \&quot;foo\&quot;, \&quot;bar\&quot;, \&quot;foo bar\&quot;, &#39;foo&#39;, &#39;bar&#39;, &#39;foo bar&#39; | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], [&#39;foo&#39;, \&quot;bar\&quot;] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  | 

### Return type

[**DSourceUsageReportResponse**](DSourceUsageReportResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchVdbInventoryReport

> VDBInventoryReportResponse SearchVdbInventoryReport(ctx).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()

Search the inventory report for virtualization engine VDBs.

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
    limit := int32(100) // int32 | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 10000. (optional) (default to 10000)
    cursor := "cursor_example" // string | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the 'prev_cursor' or 'next_cursor' property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. (optional)
    sort := "name" // string | The field to sort results by. A property name with a prepended '-' signifies descending order. (optional)
    searchBody := *openapiclient.NewSearchBody() // SearchBody | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS 'foobar', field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN ['Goku', 'Vegeta'] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically `SEARCH '12'` would match an item with an attribute with an integer value of `123`.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ 'Goku' | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ 'Goku' |  ## Grouping Parenthesis `()` can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS 'foo')  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \"foo\", \"bar\", \"foo bar\", 'foo', 'bar', 'foo bar' | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], ['foo', \"bar\"] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.SearchVdbInventoryReport(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.SearchVdbInventoryReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchVdbInventoryReport`: VDBInventoryReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.SearchVdbInventoryReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchVdbInventoryReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 10000. | [default to 10000]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 
 **searchBody** | [**SearchBody**](SearchBody.md) | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS &#39;foobar&#39;, field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN [&#39;Goku&#39;, &#39;Vegeta&#39;] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically &#x60;SEARCH &#39;12&#39;&#x60; would match an item with an attribute with an integer value of &#x60;123&#x60;.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ &#39;Goku&#39; | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ &#39;Goku&#39; |  ## Grouping Parenthesis &#x60;()&#x60; can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS &#39;foo&#39;)  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \&quot;foo\&quot;, \&quot;bar\&quot;, \&quot;foo bar\&quot;, &#39;foo&#39;, &#39;bar&#39;, &#39;foo bar&#39; | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], [&#39;foo&#39;, \&quot;bar\&quot;] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  | 

### Return type

[**VDBInventoryReportResponse**](VDBInventoryReportResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchVirtualizationStorageSummaryReport

> VirtualizationStorageSummaryReportResponse SearchVirtualizationStorageSummaryReport(ctx).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()

Search the storage summary report for virtualization engines.

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
    limit := int32(100) // int32 | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 10000. (optional) (default to 10000)
    cursor := "cursor_example" // string | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the 'prev_cursor' or 'next_cursor' property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. (optional)
    sort := "engine_id" // string | The field to sort results by. A property name with a prepended '-' signifies descending order. (optional)
    searchBody := *openapiclient.NewSearchBody() // SearchBody | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS 'foobar', field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN ['Goku', 'Vegeta'] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically `SEARCH '12'` would match an item with an attribute with an integer value of `123`.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ 'Goku' | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ 'Goku' |  ## Grouping Parenthesis `()` can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS 'foo')  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \"foo\", \"bar\", \"foo bar\", 'foo', 'bar', 'foo bar' | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], ['foo', \"bar\"] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.SearchVirtualizationStorageSummaryReport(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.SearchVirtualizationStorageSummaryReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchVirtualizationStorageSummaryReport`: VirtualizationStorageSummaryReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.SearchVirtualizationStorageSummaryReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchVirtualizationStorageSummaryReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 10000. | [default to 10000]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 
 **searchBody** | [**SearchBody**](SearchBody.md) | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS &#39;foobar&#39;, field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN [&#39;Goku&#39;, &#39;Vegeta&#39;] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically &#x60;SEARCH &#39;12&#39;&#x60; would match an item with an attribute with an integer value of &#x60;123&#x60;.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ &#39;Goku&#39; | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ &#39;Goku&#39; |  ## Grouping Parenthesis &#x60;()&#x60; can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS &#39;foo&#39;)  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \&quot;foo\&quot;, \&quot;bar\&quot;, \&quot;foo bar\&quot;, &#39;foo&#39;, &#39;bar&#39;, &#39;foo bar&#39; | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], [&#39;foo&#39;, \&quot;bar\&quot;] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  | 

### Return type

[**VirtualizationStorageSummaryReportResponse**](VirtualizationStorageSummaryReportResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReportingSchedule

> ReportingSchedule UpdateReportingSchedule(ctx, reportId).ReportingSchedule(reportingSchedule).Execute()

Update a reporting schedule by ID.

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
    reportId := int32(56) // int32 | The ID of the report schedule.
    reportingSchedule := *openapiclient.NewReportingSchedule("VIRTUALIZATION_STORAGE_SUMMARY", "0 0 2 ? * Mon-Fri", "CSV", true, []string{"user@example.com"}) // ReportingSchedule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportingApi.UpdateReportingSchedule(context.Background(), reportId).ReportingSchedule(reportingSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.UpdateReportingSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReportingSchedule`: ReportingSchedule
    fmt.Fprintf(os.Stdout, "Response from `ReportingApi.UpdateReportingSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **int32** | The ID of the report schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReportingScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportingSchedule** | [**ReportingSchedule**](ReportingSchedule.md) |  | 

### Return type

[**ReportingSchedule**](ReportingSchedule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

