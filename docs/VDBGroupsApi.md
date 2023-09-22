# \VDBGroupsApi

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVdbGroup**](VDBGroupsApi.md#CreateVdbGroup) | **Post** /vdb-groups | Create a new VDBGroup.
[**CreateVdbGroupsTags**](VDBGroupsApi.md#CreateVdbGroupsTags) | **Post** /vdb-groups/{vdbGroupId}/tags | Create tags for a VDB Group.
[**DeleteVdbGroup**](VDBGroupsApi.md#DeleteVdbGroup) | **Delete** /vdb-groups/{vdbGroupId} | Delete a VDBGoup.
[**DeleteVdbGroupTags**](VDBGroupsApi.md#DeleteVdbGroupTags) | **Post** /vdb-groups/{vdbGroupId}/tags/delete | Delete tags for a VDB Group.
[**GetBookmarksByVdbGroup**](VDBGroupsApi.md#GetBookmarksByVdbGroup) | **Get** /vdb-groups/{vdbGroupId}/bookmarks | List bookmarks compatible with this VDB Group.
[**GetVdbGroup**](VDBGroupsApi.md#GetVdbGroup) | **Get** /vdb-groups/{vdbGroupId} | Get a VDBGroup by name.
[**GetVdbGroupTags**](VDBGroupsApi.md#GetVdbGroupTags) | **Get** /vdb-groups/{vdbGroupId}/tags | Get tags for a VDB Group.
[**GetVdbGroups**](VDBGroupsApi.md#GetVdbGroups) | **Get** /vdb-groups | List all VDBGroups.
[**LockVdbGroup**](VDBGroupsApi.md#LockVdbGroup) | **Post** /vdb-groups/{vdbGroupId}/lock | Lock a VDB Group.
[**ProvisionVdbGroupFromBookmark**](VDBGroupsApi.md#ProvisionVdbGroupFromBookmark) | **Post** /vdb-groups/provision_from_bookmark | Provision a new VDB Group from a Bookmark.
[**RefreshVdbGroup**](VDBGroupsApi.md#RefreshVdbGroup) | **Post** /vdb-groups/{vdbGroupId}/refresh | Refresh a VDBGroup.
[**RollbackVdbGroup**](VDBGroupsApi.md#RollbackVdbGroup) | **Post** /vdb-groups/{vdbGroupId}/rollback | Rollback a VDBGroup.
[**SearchBookmarksByVdbGroup**](VDBGroupsApi.md#SearchBookmarksByVdbGroup) | **Post** /vdb-groups/{vdbGroupId}/bookmarks/search | Search for bookmarks compatible with this VDB Group.
[**SearchVdbGroups**](VDBGroupsApi.md#SearchVdbGroups) | **Post** /vdb-groups/search | Search for VDB Groups.
[**UnlockVdbGroup**](VDBGroupsApi.md#UnlockVdbGroup) | **Post** /vdb-groups/{vdbGroupId}/unlock | Unlock a VDB Group.
[**UpdateVdbGroupById**](VDBGroupsApi.md#UpdateVdbGroupById) | **Patch** /vdb-groups/{vdbGroupId} | Update values of a VDB group.



## CreateVdbGroup

> CreateVDBGroupResponse CreateVdbGroup(ctx).CreateVDBGroupRequest(createVDBGroupRequest).Execute()

Create a new VDBGroup.

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
    createVDBGroupRequest := *openapiclient.NewCreateVDBGroupRequest("Name_example", []string{"VdbIds_example"}) // CreateVDBGroupRequest | The parameters to create a VDBGroup.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.CreateVdbGroup(context.Background()).CreateVDBGroupRequest(createVDBGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.CreateVdbGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVdbGroup`: CreateVDBGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.CreateVdbGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVdbGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVDBGroupRequest** | [**CreateVDBGroupRequest**](CreateVDBGroupRequest.md) | The parameters to create a VDBGroup. | 

### Return type

[**CreateVDBGroupResponse**](CreateVDBGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVdbGroupsTags

> TagsResponse CreateVdbGroupsTags(ctx, vdbGroupId).TagsRequest(tagsRequest).Execute()

Create tags for a VDB Group.

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
    vdbGroupId := "vdbGroupId_example" // string | The ID or name of the VDBGroup.
    tagsRequest := *openapiclient.NewTagsRequest([]openapiclient.Tag{*openapiclient.NewTag("key-1", "value-1")}) // TagsRequest | Tags information for VDB Group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.CreateVdbGroupsTags(context.Background(), vdbGroupId).TagsRequest(tagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.CreateVdbGroupsTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVdbGroupsTags`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.CreateVdbGroupsTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbGroupId** | **string** | The ID or name of the VDBGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVdbGroupsTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagsRequest** | [**TagsRequest**](TagsRequest.md) | Tags information for VDB Group. | 

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


## DeleteVdbGroup

> DeleteVdbGroup(ctx, vdbGroupId).Execute()

Delete a VDBGoup.

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
    vdbGroupId := "vdbGroupId_example" // string | The ID or name of the VDBGroup.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VDBGroupsApi.DeleteVdbGroup(context.Background(), vdbGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.DeleteVdbGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbGroupId** | **string** | The ID or name of the VDBGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVdbGroupRequest struct via the builder pattern


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


## DeleteVdbGroupTags

> DeleteVdbGroupTags(ctx, vdbGroupId).DeleteTag(deleteTag).Execute()

Delete tags for a VDB Group.

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
    vdbGroupId := "vdbGroupId_example" // string | The ID or name of the VDBGroup.
    deleteTag := *openapiclient.NewDeleteTag() // DeleteTag | The parameters to delete tags (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VDBGroupsApi.DeleteVdbGroupTags(context.Background(), vdbGroupId).DeleteTag(deleteTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.DeleteVdbGroupTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbGroupId** | **string** | The ID or name of the VDBGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVdbGroupTagsRequest struct via the builder pattern


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


## GetBookmarksByVdbGroup

> ListBookmarksByVDBGroupsResponse GetBookmarksByVdbGroup(ctx, vdbGroupId).Limit(limit).Cursor(cursor).Sort(sort).Execute()

List bookmarks compatible with this VDB Group.

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
    vdbGroupId := "vdbGroupId_example" // string | The ID or name of the VDBGroup.
    limit := int32(50) // int32 | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. (optional) (default to 100)
    cursor := "cursor_example" // string | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the 'prev_cursor' or 'next_cursor' property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. (optional)
    sort := "id" // string | The field to sort results by. A property name with a prepended '-' signifies descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.GetBookmarksByVdbGroup(context.Background(), vdbGroupId).Limit(limit).Cursor(cursor).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.GetBookmarksByVdbGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBookmarksByVdbGroup`: ListBookmarksByVDBGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.GetBookmarksByVdbGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbGroupId** | **string** | The ID or name of the VDBGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookmarksByVdbGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 

### Return type

[**ListBookmarksByVDBGroupsResponse**](ListBookmarksByVDBGroupsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVdbGroup

> VDBGroup GetVdbGroup(ctx, vdbGroupId).Execute()

Get a VDBGroup by name.

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
    vdbGroupId := "vdbGroupId_example" // string | The ID or name of the VDBGroup.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.GetVdbGroup(context.Background(), vdbGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.GetVdbGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVdbGroup`: VDBGroup
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.GetVdbGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbGroupId** | **string** | The ID or name of the VDBGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVdbGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VDBGroup**](VDBGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVdbGroupTags

> TagsResponse GetVdbGroupTags(ctx, vdbGroupId).Execute()

Get tags for a VDB Group.

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
    vdbGroupId := "vdbGroupId_example" // string | The ID or name of the VDBGroup.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.GetVdbGroupTags(context.Background(), vdbGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.GetVdbGroupTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVdbGroupTags`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.GetVdbGroupTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbGroupId** | **string** | The ID or name of the VDBGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVdbGroupTagsRequest struct via the builder pattern


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


## GetVdbGroups

> ListVDBGroupsResponse GetVdbGroups(ctx).Limit(limit).Cursor(cursor).Sort(sort).Execute()

List all VDBGroups.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.GetVdbGroups(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.GetVdbGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVdbGroups`: ListVDBGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.GetVdbGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVdbGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 

### Return type

[**ListVDBGroupsResponse**](ListVDBGroupsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockVdbGroup

> VDBGroup LockVdbGroup(ctx, vdbGroupId).LockVDBGroupParameters(lockVDBGroupParameters).Execute()

Lock a VDB Group.

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
    vdbGroupId := "vdbGroupId_example" // string | The ID or name of the VDBGroup.
    lockVDBGroupParameters := *openapiclient.NewLockVDBGroupParameters() // LockVDBGroupParameters | The parameters to lock a VDB Group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.LockVdbGroup(context.Background(), vdbGroupId).LockVDBGroupParameters(lockVDBGroupParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.LockVdbGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LockVdbGroup`: VDBGroup
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.LockVdbGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbGroupId** | **string** | The ID or name of the VDBGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockVdbGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lockVDBGroupParameters** | [**LockVDBGroupParameters**](LockVDBGroupParameters.md) | The parameters to lock a VDB Group. | 

### Return type

[**VDBGroup**](VDBGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisionVdbGroupFromBookmark

> ProvisionVDBGroupFromBookmarkResponse ProvisionVdbGroupFromBookmark(ctx).ProvisionVDBGroupFromBookmarkParameters(provisionVDBGroupFromBookmarkParameters).Execute()

Provision a new VDB Group from a Bookmark.

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
    provisionVDBGroupFromBookmarkParameters := *openapiclient.NewProvisionVDBGroupFromBookmarkParameters("Name_example", "BookmarkId_example", map[string]BaseProvisionVDBParameters{"key": *openapiclient.NewBaseProvisionVDBParameters()}) // ProvisionVDBGroupFromBookmarkParameters | The parameters to provision a VDB group from a Bookmark. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.ProvisionVdbGroupFromBookmark(context.Background()).ProvisionVDBGroupFromBookmarkParameters(provisionVDBGroupFromBookmarkParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.ProvisionVdbGroupFromBookmark``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionVdbGroupFromBookmark`: ProvisionVDBGroupFromBookmarkResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.ProvisionVdbGroupFromBookmark`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisionVdbGroupFromBookmarkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisionVDBGroupFromBookmarkParameters** | [**ProvisionVDBGroupFromBookmarkParameters**](ProvisionVDBGroupFromBookmarkParameters.md) | The parameters to provision a VDB group from a Bookmark. | 

### Return type

[**ProvisionVDBGroupFromBookmarkResponse**](ProvisionVDBGroupFromBookmarkResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshVdbGroup

> RefreshVDBGroupResponse RefreshVdbGroup(ctx, vdbGroupId).RefreshVDBGroupParameters(refreshVDBGroupParameters).Execute()

Refresh a VDBGroup.

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
    vdbGroupId := "vdbGroupId_example" // string | The ID or name of the VDBGroup.
    refreshVDBGroupParameters := *openapiclient.NewRefreshVDBGroupParameters("BookmarkId_example") // RefreshVDBGroupParameters | The parameters to refresh a VDBGroup. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.RefreshVdbGroup(context.Background(), vdbGroupId).RefreshVDBGroupParameters(refreshVDBGroupParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.RefreshVdbGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshVdbGroup`: RefreshVDBGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.RefreshVdbGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbGroupId** | **string** | The ID or name of the VDBGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshVdbGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refreshVDBGroupParameters** | [**RefreshVDBGroupParameters**](RefreshVDBGroupParameters.md) | The parameters to refresh a VDBGroup. | 

### Return type

[**RefreshVDBGroupResponse**](RefreshVDBGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackVdbGroup

> RollbackVDBGroupResponse RollbackVdbGroup(ctx, vdbGroupId).RollbackVDBGroupParameters(rollbackVDBGroupParameters).Execute()

Rollback a VDBGroup.

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
    vdbGroupId := "vdbGroupId_example" // string | The ID or name of the VDBGroup.
    rollbackVDBGroupParameters := *openapiclient.NewRollbackVDBGroupParameters("BookmarkId_example") // RollbackVDBGroupParameters | The parameters to rollback a VDBGroup. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.RollbackVdbGroup(context.Background(), vdbGroupId).RollbackVDBGroupParameters(rollbackVDBGroupParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.RollbackVdbGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RollbackVdbGroup`: RollbackVDBGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.RollbackVdbGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbGroupId** | **string** | The ID or name of the VDBGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackVdbGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rollbackVDBGroupParameters** | [**RollbackVDBGroupParameters**](RollbackVDBGroupParameters.md) | The parameters to rollback a VDBGroup. | 

### Return type

[**RollbackVDBGroupResponse**](RollbackVDBGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchBookmarksByVdbGroup

> SearchBookmarksByVDBGroupsResponse SearchBookmarksByVdbGroup(ctx, vdbGroupId).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()

Search for bookmarks compatible with this VDB Group.

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
    vdbGroupId := "vdbGroupId_example" // string | The ID or name of the VDBGroup.
    limit := int32(50) // int32 | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. (optional) (default to 100)
    cursor := "cursor_example" // string | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the 'prev_cursor' or 'next_cursor' property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. (optional)
    sort := "id" // string | The field to sort results by. A property name with a prepended '-' signifies descending order. (optional)
    searchBody := *openapiclient.NewSearchBody() // SearchBody | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS 'foobar', field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN ['Goku', 'Vegeta'] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically `SEARCH '12'` would match an item with an attribute with an integer value of `123`.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ 'Goku' | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ 'Goku' |  ## Grouping Parenthesis `()` can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS 'foo')  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \"foo\", \"bar\", \"foo bar\", 'foo', 'bar', 'foo bar' | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], ['foo', \"bar\"] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.SearchBookmarksByVdbGroup(context.Background(), vdbGroupId).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.SearchBookmarksByVdbGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchBookmarksByVdbGroup`: SearchBookmarksByVDBGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.SearchBookmarksByVdbGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbGroupId** | **string** | The ID or name of the VDBGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchBookmarksByVdbGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 
 **searchBody** | [**SearchBody**](SearchBody.md) | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS &#39;foobar&#39;, field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN [&#39;Goku&#39;, &#39;Vegeta&#39;] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically &#x60;SEARCH &#39;12&#39;&#x60; would match an item with an attribute with an integer value of &#x60;123&#x60;.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ &#39;Goku&#39; | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ &#39;Goku&#39; |  ## Grouping Parenthesis &#x60;()&#x60; can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS &#39;foo&#39;)  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \&quot;foo\&quot;, \&quot;bar\&quot;, \&quot;foo bar\&quot;, &#39;foo&#39;, &#39;bar&#39;, &#39;foo bar&#39; | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], [&#39;foo&#39;, \&quot;bar\&quot;] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  | 

### Return type

[**SearchBookmarksByVDBGroupsResponse**](SearchBookmarksByVDBGroupsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchVdbGroups

> SearchVDBGroupResponse SearchVdbGroups(ctx).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()

Search for VDB Groups.

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
    searchBody := *openapiclient.NewSearchBody() // SearchBody | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS 'foobar', field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN ['Goku', 'Vegeta'] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically `SEARCH '12'` would match an item with an attribute with an integer value of `123`.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ 'Goku' | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ 'Goku' |  ## Grouping Parenthesis `()` can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS 'foo')  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \"foo\", \"bar\", \"foo bar\", 'foo', 'bar', 'foo bar' | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], ['foo', \"bar\"] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.SearchVdbGroups(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.SearchVdbGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchVdbGroups`: SearchVDBGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.SearchVdbGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchVdbGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 
 **searchBody** | [**SearchBody**](SearchBody.md) | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS &#39;foobar&#39;, field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN [&#39;Goku&#39;, &#39;Vegeta&#39;] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically &#x60;SEARCH &#39;12&#39;&#x60; would match an item with an attribute with an integer value of &#x60;123&#x60;.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ &#39;Goku&#39; | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ &#39;Goku&#39; |  ## Grouping Parenthesis &#x60;()&#x60; can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS &#39;foo&#39;)  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \&quot;foo\&quot;, \&quot;bar\&quot;, \&quot;foo bar\&quot;, &#39;foo&#39;, &#39;bar&#39;, &#39;foo bar&#39; | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], [&#39;foo&#39;, \&quot;bar\&quot;] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  | 

### Return type

[**SearchVDBGroupResponse**](SearchVDBGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlockVdbGroup

> VDBGroup UnlockVdbGroup(ctx, vdbGroupId).Execute()

Unlock a VDB Group.

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
    vdbGroupId := "vdbGroupId_example" // string | The ID or name of the VDBGroup.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.UnlockVdbGroup(context.Background(), vdbGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.UnlockVdbGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlockVdbGroup`: VDBGroup
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.UnlockVdbGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbGroupId** | **string** | The ID or name of the VDBGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockVdbGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VDBGroup**](VDBGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVdbGroupById

> VDBGroup UpdateVdbGroupById(ctx, vdbGroupId).UpdateVDBGroupParameters(updateVDBGroupParameters).Execute()

Update values of a VDB group.

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
    vdbGroupId := "vdbGroupId_example" // string | The ID or name of the VDBGroup.
    updateVDBGroupParameters := *openapiclient.NewUpdateVDBGroupParameters() // UpdateVDBGroupParameters | The new data to update a VDB group. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.UpdateVdbGroupById(context.Background(), vdbGroupId).UpdateVDBGroupParameters(updateVDBGroupParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.UpdateVdbGroupById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVdbGroupById`: VDBGroup
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.UpdateVdbGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbGroupId** | **string** | The ID or name of the VDBGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVdbGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateVDBGroupParameters** | [**UpdateVDBGroupParameters**](UpdateVDBGroupParameters.md) | The new data to update a VDB group. | 

### Return type

[**VDBGroup**](VDBGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

