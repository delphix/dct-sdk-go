# \DSourcesApi

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTagsDsource**](DSourcesApi.md#CreateTagsDsource) | **Post** /dsources/{dsourceId}/tags | Create tags for a dSource.
[**DeleteDsource**](DSourcesApi.md#DeleteDsource) | **Post** /dsources/delete | Delete the specified dSource.
[**DeleteTagsDsource**](DSourcesApi.md#DeleteTagsDsource) | **Post** /dsources/{dsourceId}/tags/delete | Delete tags for a dSource.
[**GetAppdataDsourceLinkingDefaults**](DSourcesApi.md#GetAppdataDsourceLinkingDefaults) | **Post** /dsources/appdata/defaults | Get defaults for an AppData dSource linking.
[**GetAseDsourceLinkingDefaults**](DSourcesApi.md#GetAseDsourceLinkingDefaults) | **Post** /dsources/ase/defaults | Get defaults for an ASE dSource linking.
[**GetDsourceById**](DSourcesApi.md#GetDsourceById) | **Get** /dsources/{dsourceId} | Get a dSource by ID.
[**GetDsourceSnapshots**](DSourcesApi.md#GetDsourceSnapshots) | **Get** /dsources/{dsourceId}/snapshots | List Snapshots for a dSource.
[**GetDsources**](DSourcesApi.md#GetDsources) | **Get** /dsources | List all dSources.
[**GetOracleDsourceLinkingDefaults**](DSourcesApi.md#GetOracleDsourceLinkingDefaults) | **Post** /dsources/oracle/defaults | Get defaults for dSource linking.
[**GetTagsDsource**](DSourcesApi.md#GetTagsDsource) | **Get** /dsources/{dsourceId}/tags | Get tags for a dSource.
[**LinkAppdataDatabase**](DSourcesApi.md#LinkAppdataDatabase) | **Post** /dsources/appdata | Link an AppData database as dSource.
[**LinkAseDatabase**](DSourcesApi.md#LinkAseDatabase) | **Post** /dsources/ase | Link an ASE database as dSource.
[**LinkOracleDatabase**](DSourcesApi.md#LinkOracleDatabase) | **Post** /dsources/oracle | Link Oracle database as dSource.
[**SearchDsources**](DSourcesApi.md#SearchDsources) | **Post** /dsources/search | Search for dSources.
[**SnapshotDsource**](DSourcesApi.md#SnapshotDsource) | **Post** /dsources/{dsourceId}/snapshots | Snapshot a dSource.



## CreateTagsDsource

> TagsResponse CreateTagsDsource(ctx, dsourceId).TagsRequest(tagsRequest).Execute()

Create tags for a dSource.

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
    dsourceId := "dsourceId_example" // string | The ID of the dSource.
    tagsRequest := *openapiclient.NewTagsRequest([]openapiclient.Tag{*openapiclient.NewTag("key-1", "value-1")}) // TagsRequest | Tags information for DSource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DSourcesApi.CreateTagsDsource(context.Background(), dsourceId).TagsRequest(tagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DSourcesApi.CreateTagsDsource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTagsDsource`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `DSourcesApi.CreateTagsDsource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dsourceId** | **string** | The ID of the dSource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagsDsourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagsRequest** | [**TagsRequest**](TagsRequest.md) | Tags information for DSource. | 

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


## DeleteDsource

> Job DeleteDsource(ctx).DeleteDSourceRequest(deleteDSourceRequest).Execute()

Delete the specified dSource.

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
    deleteDSourceRequest := *openapiclient.NewDeleteDSourceRequest("DsourceId_example") // DeleteDSourceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DSourcesApi.DeleteDsource(context.Background()).DeleteDSourceRequest(deleteDSourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DSourcesApi.DeleteDsource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDsource`: Job
    fmt.Fprintf(os.Stdout, "Response from `DSourcesApi.DeleteDsource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDsourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDSourceRequest** | [**DeleteDSourceRequest**](DeleteDSourceRequest.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTagsDsource

> DeleteTagsDsource(ctx, dsourceId).DeleteTag(deleteTag).Execute()

Delete tags for a dSource.

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
    dsourceId := "dsourceId_example" // string | The ID of the dSource.
    deleteTag := *openapiclient.NewDeleteTag() // DeleteTag | The parameters to delete tags (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DSourcesApi.DeleteTagsDsource(context.Background(), dsourceId).DeleteTag(deleteTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DSourcesApi.DeleteTagsDsource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dsourceId** | **string** | The ID of the dSource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagsDsourceRequest struct via the builder pattern


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


## GetAppdataDsourceLinkingDefaults

> AppDataDSourceLinkSourceParameters GetAppdataDsourceLinkingDefaults(ctx).LinkDSourceDefaultRequest(linkDSourceDefaultRequest).Execute()

Get defaults for an AppData dSource linking.

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
    linkDSourceDefaultRequest := *openapiclient.NewLinkDSourceDefaultRequest("SourceId_example") // LinkDSourceDefaultRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DSourcesApi.GetAppdataDsourceLinkingDefaults(context.Background()).LinkDSourceDefaultRequest(linkDSourceDefaultRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DSourcesApi.GetAppdataDsourceLinkingDefaults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppdataDsourceLinkingDefaults`: AppDataDSourceLinkSourceParameters
    fmt.Fprintf(os.Stdout, "Response from `DSourcesApi.GetAppdataDsourceLinkingDefaults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppdataDsourceLinkingDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkDSourceDefaultRequest** | [**LinkDSourceDefaultRequest**](LinkDSourceDefaultRequest.md) |  | 

### Return type

[**AppDataDSourceLinkSourceParameters**](AppDataDSourceLinkSourceParameters.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAseDsourceLinkingDefaults

> ASEDSourceLinkSourceParameters GetAseDsourceLinkingDefaults(ctx).LinkDSourceDefaultRequest(linkDSourceDefaultRequest).Execute()

Get defaults for an ASE dSource linking.

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
    linkDSourceDefaultRequest := *openapiclient.NewLinkDSourceDefaultRequest("SourceId_example") // LinkDSourceDefaultRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DSourcesApi.GetAseDsourceLinkingDefaults(context.Background()).LinkDSourceDefaultRequest(linkDSourceDefaultRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DSourcesApi.GetAseDsourceLinkingDefaults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAseDsourceLinkingDefaults`: ASEDSourceLinkSourceParameters
    fmt.Fprintf(os.Stdout, "Response from `DSourcesApi.GetAseDsourceLinkingDefaults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAseDsourceLinkingDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkDSourceDefaultRequest** | [**LinkDSourceDefaultRequest**](LinkDSourceDefaultRequest.md) |  | 

### Return type

[**ASEDSourceLinkSourceParameters**](ASEDSourceLinkSourceParameters.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDsourceById

> DSource GetDsourceById(ctx, dsourceId).Execute()

Get a dSource by ID.

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
    dsourceId := "dsourceId_example" // string | The ID of the dSource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DSourcesApi.GetDsourceById(context.Background(), dsourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DSourcesApi.GetDsourceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDsourceById`: DSource
    fmt.Fprintf(os.Stdout, "Response from `DSourcesApi.GetDsourceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dsourceId** | **string** | The ID of the dSource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDsourceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DSource**](DSource.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDsourceSnapshots

> ListSnapshotsResponse GetDsourceSnapshots(ctx, dsourceId).Limit(limit).Cursor(cursor).Execute()

List Snapshots for a dSource.

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
    dsourceId := "dsourceId_example" // string | The ID of the dSource.
    limit := int32(50) // int32 | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. (optional) (default to 100)
    cursor := "cursor_example" // string | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the 'prev_cursor' or 'next_cursor' property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DSourcesApi.GetDsourceSnapshots(context.Background(), dsourceId).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DSourcesApi.GetDsourceSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDsourceSnapshots`: ListSnapshotsResponse
    fmt.Fprintf(os.Stdout, "Response from `DSourcesApi.GetDsourceSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dsourceId** | **string** | The ID of the dSource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDsourceSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 

### Return type

[**ListSnapshotsResponse**](ListSnapshotsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDsources

> ListDSourcesResponse GetDsources(ctx).Limit(limit).Cursor(cursor).Sort(sort).Permission(permission).Execute()

List all dSources.

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
    sort := "id" // string | The field to sort results by. A property name with a prepended '-' signifies descending order. (optional)
    permission := openapiclient.PermissionEnum("READ") // PermissionEnum | Restrict the objects, which are allowed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DSourcesApi.GetDsources(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).Permission(permission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DSourcesApi.GetDsources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDsources`: ListDSourcesResponse
    fmt.Fprintf(os.Stdout, "Response from `DSourcesApi.GetDsources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDsourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 
 **permission** | [**PermissionEnum**](PermissionEnum.md) | Restrict the objects, which are allowed. | 

### Return type

[**ListDSourcesResponse**](ListDSourcesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOracleDsourceLinkingDefaults

> OracleDSourceLinkSourceParameters GetOracleDsourceLinkingDefaults(ctx).LinkDSourceDefaultRequest(linkDSourceDefaultRequest).Execute()

Get defaults for dSource linking.

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
    linkDSourceDefaultRequest := *openapiclient.NewLinkDSourceDefaultRequest("SourceId_example") // LinkDSourceDefaultRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DSourcesApi.GetOracleDsourceLinkingDefaults(context.Background()).LinkDSourceDefaultRequest(linkDSourceDefaultRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DSourcesApi.GetOracleDsourceLinkingDefaults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOracleDsourceLinkingDefaults`: OracleDSourceLinkSourceParameters
    fmt.Fprintf(os.Stdout, "Response from `DSourcesApi.GetOracleDsourceLinkingDefaults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOracleDsourceLinkingDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkDSourceDefaultRequest** | [**LinkDSourceDefaultRequest**](LinkDSourceDefaultRequest.md) |  | 

### Return type

[**OracleDSourceLinkSourceParameters**](OracleDSourceLinkSourceParameters.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagsDsource

> TagsResponse GetTagsDsource(ctx, dsourceId).Execute()

Get tags for a dSource.

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
    dsourceId := "dsourceId_example" // string | The ID of the dSource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DSourcesApi.GetTagsDsource(context.Background(), dsourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DSourcesApi.GetTagsDsource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagsDsource`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `DSourcesApi.GetTagsDsource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dsourceId** | **string** | The ID of the dSource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsDsourceRequest struct via the builder pattern


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


## LinkAppdataDatabase

> LinkDSourceResponse LinkAppdataDatabase(ctx).AppDataDSourceLinkSourceParameters(appDataDSourceLinkSourceParameters).Execute()

Link an AppData database as dSource.

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
    appDataDSourceLinkSourceParameters := *openapiclient.NewAppDataDSourceLinkSourceParameters("EnvironmentUser_example", map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}) // AppDataDSourceLinkSourceParameters | The parameters to link an AppData dSource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DSourcesApi.LinkAppdataDatabase(context.Background()).AppDataDSourceLinkSourceParameters(appDataDSourceLinkSourceParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DSourcesApi.LinkAppdataDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkAppdataDatabase`: LinkDSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `DSourcesApi.LinkAppdataDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkAppdataDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appDataDSourceLinkSourceParameters** | [**AppDataDSourceLinkSourceParameters**](AppDataDSourceLinkSourceParameters.md) | The parameters to link an AppData dSource. | 

### Return type

[**LinkDSourceResponse**](LinkDSourceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkAseDatabase

> LinkDSourceResponse LinkAseDatabase(ctx).ASEDSourceLinkSourceParameters(aSEDSourceLinkSourceParameters).Execute()

Link an ASE database as dSource.

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
    aSEDSourceLinkSourceParameters := *openapiclient.NewASEDSourceLinkSourceParameters("SourceId_example", "LoadBackupPath_example") // ASEDSourceLinkSourceParameters | The parameters to link an ASE dSource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DSourcesApi.LinkAseDatabase(context.Background()).ASEDSourceLinkSourceParameters(aSEDSourceLinkSourceParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DSourcesApi.LinkAseDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkAseDatabase`: LinkDSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `DSourcesApi.LinkAseDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkAseDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aSEDSourceLinkSourceParameters** | [**ASEDSourceLinkSourceParameters**](ASEDSourceLinkSourceParameters.md) | The parameters to link an ASE dSource. | 

### Return type

[**LinkDSourceResponse**](LinkDSourceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkOracleDatabase

> LinkDSourceResponse LinkOracleDatabase(ctx).OracleDSourceLinkSourceParameters(oracleDSourceLinkSourceParameters).Execute()

Link Oracle database as dSource.

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
    oracleDSourceLinkSourceParameters := *openapiclient.NewOracleDSourceLinkSourceParameters("SourceId_example") // OracleDSourceLinkSourceParameters | The parameters to link an Oracle dSource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DSourcesApi.LinkOracleDatabase(context.Background()).OracleDSourceLinkSourceParameters(oracleDSourceLinkSourceParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DSourcesApi.LinkOracleDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkOracleDatabase`: LinkDSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `DSourcesApi.LinkOracleDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkOracleDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oracleDSourceLinkSourceParameters** | [**OracleDSourceLinkSourceParameters**](OracleDSourceLinkSourceParameters.md) | The parameters to link an Oracle dSource. | 

### Return type

[**LinkDSourceResponse**](LinkDSourceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDsources

> SearchDSourcesResponse SearchDsources(ctx).Limit(limit).Cursor(cursor).Sort(sort).Permission(permission).SearchBody(searchBody).Execute()

Search for dSources.

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
    sort := "id" // string | The field to sort results by. A property name with a prepended '-' signifies descending order. (optional)
    permission := openapiclient.PermissionEnum("READ") // PermissionEnum | Restrict the objects, which are allowed. (optional)
    searchBody := *openapiclient.NewSearchBody() // SearchBody | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS 'foobar', field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN ['Goku', 'Vegeta'] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically `SEARCH '12'` would match an item with an attribute with an integer value of `123`.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ 'Goku' | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ 'Goku' |  ## Grouping Parenthesis `()` can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS 'foo')  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \"foo\", \"bar\", \"foo bar\", 'foo', 'bar', 'foo bar' | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], ['foo', \"bar\"] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DSourcesApi.SearchDsources(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).Permission(permission).SearchBody(searchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DSourcesApi.SearchDsources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchDsources`: SearchDSourcesResponse
    fmt.Fprintf(os.Stdout, "Response from `DSourcesApi.SearchDsources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDsourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 
 **permission** | [**PermissionEnum**](PermissionEnum.md) | Restrict the objects, which are allowed. | 
 **searchBody** | [**SearchBody**](SearchBody.md) | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS &#39;foobar&#39;, field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN [&#39;Goku&#39;, &#39;Vegeta&#39;] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically &#x60;SEARCH &#39;12&#39;&#x60; would match an item with an attribute with an integer value of &#x60;123&#x60;.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ &#39;Goku&#39; | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ &#39;Goku&#39; |  ## Grouping Parenthesis &#x60;()&#x60; can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS &#39;foo&#39;)  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \&quot;foo\&quot;, \&quot;bar\&quot;, \&quot;foo bar\&quot;, &#39;foo&#39;, &#39;bar&#39;, &#39;foo bar&#39; | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], [&#39;foo&#39;, \&quot;bar\&quot;] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  | 

### Return type

[**SearchDSourcesResponse**](SearchDSourcesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotDsource

> SnapshotDSourceResponse SnapshotDsource(ctx, dsourceId).DSourceSnapshotParameters(dSourceSnapshotParameters).Execute()

Snapshot a dSource.

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
    dsourceId := "dsourceId_example" // string | The ID of the dSource.
    dSourceSnapshotParameters := *openapiclient.NewDSourceSnapshotParameters() // DSourceSnapshotParameters | Optional parameters to snapshot a DSource. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DSourcesApi.SnapshotDsource(context.Background(), dsourceId).DSourceSnapshotParameters(dSourceSnapshotParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DSourcesApi.SnapshotDsource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnapshotDsource`: SnapshotDSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `DSourcesApi.SnapshotDsource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dsourceId** | **string** | The ID of the dSource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotDsourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dSourceSnapshotParameters** | [**DSourceSnapshotParameters**](DSourceSnapshotParameters.md) | Optional parameters to snapshot a DSource. | 

### Return type

[**SnapshotDSourceResponse**](SnapshotDSourceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

