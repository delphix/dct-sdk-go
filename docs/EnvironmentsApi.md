# \EnvironmentsApi

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompatibleRepositoriesByLocation**](EnvironmentsApi.md#CompatibleRepositoriesByLocation) | **Post** /environments/compatible_repositories_by_location | Get compatible repositories corresponding to the location.
[**CompatibleRepositoriesBySnapshot**](EnvironmentsApi.md#CompatibleRepositoriesBySnapshot) | **Post** /environments/compatible_repositories_by_snapshot | Get compatible repositories corresponding to the snapshot.
[**CompatibleRepositoriesByTimestamp**](EnvironmentsApi.md#CompatibleRepositoriesByTimestamp) | **Post** /environments/compatible_repositories_by_timestamp | Get compatible repositories corresponding to the timestamp.
[**CompatibleRepositoriesFromBookmark**](EnvironmentsApi.md#CompatibleRepositoriesFromBookmark) | **Post** /environments/compatible_repositories_from_bookmark | Get compatible repositories corresponding to the bookmark.
[**CreateEnvironment**](EnvironmentsApi.md#CreateEnvironment) | **Post** /environments | Create an environment.
[**CreateEnvironmentTags**](EnvironmentsApi.md#CreateEnvironmentTags) | **Post** /environments/{environmentId}/tags | Create tags for an Environment.
[**CreateEnvironmentUser**](EnvironmentsApi.md#CreateEnvironmentUser) | **Post** /environments/{environmentId}/users | Create environment user.
[**CreateHost**](EnvironmentsApi.md#CreateHost) | **Post** /environments/{environmentId}/hosts | Create a new Host.
[**DeleteEnvironment**](EnvironmentsApi.md#DeleteEnvironment) | **Delete** /environments/{environmentId} | Delete an environment by ID.
[**DeleteEnvironmentTags**](EnvironmentsApi.md#DeleteEnvironmentTags) | **Post** /environments/{environmentId}/tags/delete | Delete tags for an Environment.
[**DeleteEnvironmentUser**](EnvironmentsApi.md#DeleteEnvironmentUser) | **Delete** /environments/{environmentId}/users/{userRef} | Delete environment user.
[**DeleteHost**](EnvironmentsApi.md#DeleteHost) | **Delete** /environments/{environmentId}/hosts/{hostId} | Delete a Host.
[**DisableEnvironment**](EnvironmentsApi.md#DisableEnvironment) | **Post** /environments/{environmentId}/disable | Disable environment.
[**EnableEnvironment**](EnvironmentsApi.md#EnableEnvironment) | **Post** /environments/{environmentId}/enable | Enable a disabled environment.
[**GetEnvironmentById**](EnvironmentsApi.md#GetEnvironmentById) | **Get** /environments/{environmentId} | Returns an environment by ID.
[**GetEnvironments**](EnvironmentsApi.md#GetEnvironments) | **Get** /environments | List all environments.
[**GetTagsEnvironment**](EnvironmentsApi.md#GetTagsEnvironment) | **Get** /environments/{environmentId}/tags | Get tags for an Environment.
[**ListEnvironmentUsers**](EnvironmentsApi.md#ListEnvironmentUsers) | **Get** /environments/{environmentId}/users | List environment users.
[**PrimaryEnvironmentUser**](EnvironmentsApi.md#PrimaryEnvironmentUser) | **Post** /environments/{environmentId}/users/{userRef}/primary | Set primary environment user.
[**RefreshEnvironment**](EnvironmentsApi.md#RefreshEnvironment) | **Post** /environments/{environmentId}/refresh | Refresh environment.
[**SearchEnvironments**](EnvironmentsApi.md#SearchEnvironments) | **Post** /environments/search | Search for environments.
[**UpdateEnvironment**](EnvironmentsApi.md#UpdateEnvironment) | **Patch** /environments/{environmentId} | Update an environment by ID.
[**UpdateEnvironmentUser**](EnvironmentsApi.md#UpdateEnvironmentUser) | **Put** /environments/{environmentId}/users/{userRef} | Update environment user.
[**UpdateHost**](EnvironmentsApi.md#UpdateHost) | **Patch** /environments/{environmentId}/hosts/{hostId} | Update a Host.
[**UpdateRepository**](EnvironmentsApi.md#UpdateRepository) | **Patch** /environments/{environmentId}/repository/{repositoryId} | Update a Repository.



## CompatibleRepositoriesByLocation

> LocationCompatibleEnvironmentsResponse CompatibleRepositoriesByLocation(ctx).LocationCompatibleRepositoryRequest(locationCompatibleRepositoryRequest).Execute()

Get compatible repositories corresponding to the location.

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
    locationCompatibleRepositoryRequest := *openapiclient.NewLocationCompatibleRepositoryRequest() // LocationCompatibleRepositoryRequest | The request to get compatible repositories for provisioning a new VDB by location.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.CompatibleRepositoriesByLocation(context.Background()).LocationCompatibleRepositoryRequest(locationCompatibleRepositoryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CompatibleRepositoriesByLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompatibleRepositoriesByLocation`: LocationCompatibleEnvironmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.CompatibleRepositoriesByLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompatibleRepositoriesByLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locationCompatibleRepositoryRequest** | [**LocationCompatibleRepositoryRequest**](LocationCompatibleRepositoryRequest.md) | The request to get compatible repositories for provisioning a new VDB by location. | 

### Return type

[**LocationCompatibleEnvironmentsResponse**](LocationCompatibleEnvironmentsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompatibleRepositoriesBySnapshot

> SnapshotCompatibleEnvironmentsResponse CompatibleRepositoriesBySnapshot(ctx).SnapshotCompatibleRepositoryRequest(snapshotCompatibleRepositoryRequest).Execute()

Get compatible repositories corresponding to the snapshot.

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
    snapshotCompatibleRepositoryRequest := *openapiclient.NewSnapshotCompatibleRepositoryRequest() // SnapshotCompatibleRepositoryRequest | The request to get compatible repositories for provisioning a new VDB by snapshot.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.CompatibleRepositoriesBySnapshot(context.Background()).SnapshotCompatibleRepositoryRequest(snapshotCompatibleRepositoryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CompatibleRepositoriesBySnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompatibleRepositoriesBySnapshot`: SnapshotCompatibleEnvironmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.CompatibleRepositoriesBySnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompatibleRepositoriesBySnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **snapshotCompatibleRepositoryRequest** | [**SnapshotCompatibleRepositoryRequest**](SnapshotCompatibleRepositoryRequest.md) | The request to get compatible repositories for provisioning a new VDB by snapshot. | 

### Return type

[**SnapshotCompatibleEnvironmentsResponse**](SnapshotCompatibleEnvironmentsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompatibleRepositoriesByTimestamp

> TimestampCompatibleEnvironmentsResponse CompatibleRepositoriesByTimestamp(ctx).TimestampCompatibleRepositoryRequest(timestampCompatibleRepositoryRequest).Execute()

Get compatible repositories corresponding to the timestamp.

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
    timestampCompatibleRepositoryRequest := *openapiclient.NewTimestampCompatibleRepositoryRequest() // TimestampCompatibleRepositoryRequest | The request to get compatible repositories for provisioning a new VDB by timestamp.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.CompatibleRepositoriesByTimestamp(context.Background()).TimestampCompatibleRepositoryRequest(timestampCompatibleRepositoryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CompatibleRepositoriesByTimestamp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompatibleRepositoriesByTimestamp`: TimestampCompatibleEnvironmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.CompatibleRepositoriesByTimestamp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompatibleRepositoriesByTimestampRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestampCompatibleRepositoryRequest** | [**TimestampCompatibleRepositoryRequest**](TimestampCompatibleRepositoryRequest.md) | The request to get compatible repositories for provisioning a new VDB by timestamp. | 

### Return type

[**TimestampCompatibleEnvironmentsResponse**](TimestampCompatibleEnvironmentsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompatibleRepositoriesFromBookmark

> BookmarkCompatibleEnvironmentsResponse CompatibleRepositoriesFromBookmark(ctx).BookmarkCompatibleRepositoryRequest(bookmarkCompatibleRepositoryRequest).Execute()

Get compatible repositories corresponding to the bookmark.

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
    bookmarkCompatibleRepositoryRequest := *openapiclient.NewBookmarkCompatibleRepositoryRequest("BookmarkId_example") // BookmarkCompatibleRepositoryRequest | The request to get compatible repositories for provisioning a new VDB by bookmark.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.CompatibleRepositoriesFromBookmark(context.Background()).BookmarkCompatibleRepositoryRequest(bookmarkCompatibleRepositoryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CompatibleRepositoriesFromBookmark``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompatibleRepositoriesFromBookmark`: BookmarkCompatibleEnvironmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.CompatibleRepositoriesFromBookmark`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompatibleRepositoriesFromBookmarkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookmarkCompatibleRepositoryRequest** | [**BookmarkCompatibleRepositoryRequest**](BookmarkCompatibleRepositoryRequest.md) | The request to get compatible repositories for provisioning a new VDB by bookmark. | 

### Return type

[**BookmarkCompatibleEnvironmentsResponse**](BookmarkCompatibleEnvironmentsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnvironment

> CreateEnvironmentResponse CreateEnvironment(ctx).EnvironmentCreateParameters(environmentCreateParameters).Execute()

Create an environment.

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
    environmentCreateParameters := *openapiclient.NewEnvironmentCreateParameters("12", "UNIX", "db.host.com") // EnvironmentCreateParameters | The parameters to create an environment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.CreateEnvironment(context.Background()).EnvironmentCreateParameters(environmentCreateParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CreateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironment`: CreateEnvironmentResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.CreateEnvironment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentCreateParameters** | [**EnvironmentCreateParameters**](EnvironmentCreateParameters.md) | The parameters to create an environment. | 

### Return type

[**CreateEnvironmentResponse**](CreateEnvironmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnvironmentTags

> TagsResponse CreateEnvironmentTags(ctx, environmentId).TagsRequest(tagsRequest).Execute()

Create tags for an Environment.

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
    environmentId := "environmentId_example" // string | The ID of the environment.
    tagsRequest := *openapiclient.NewTagsRequest([]openapiclient.Tag{*openapiclient.NewTag("key-1", "value-1")}) // TagsRequest | Tags information for Environment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.CreateEnvironmentTags(context.Background(), environmentId).TagsRequest(tagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CreateEnvironmentTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironmentTags`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.CreateEnvironmentTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagsRequest** | [**TagsRequest**](TagsRequest.md) | Tags information for Environment. | 

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


## CreateEnvironmentUser

> CreateEnvironmentUserResponse CreateEnvironmentUser(ctx, environmentId).EnvironmentUserParams(environmentUserParams).Execute()

Create environment user.

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
    environmentId := "environmentId_example" // string | The ID of the environment.
    environmentUserParams := *openapiclient.NewEnvironmentUserParams() // EnvironmentUserParams | The parameters to create an environment user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.CreateEnvironmentUser(context.Background(), environmentId).EnvironmentUserParams(environmentUserParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CreateEnvironmentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironmentUser`: CreateEnvironmentUserResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.CreateEnvironmentUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentUserParams** | [**EnvironmentUserParams**](EnvironmentUserParams.md) | The parameters to create an environment user. | 

### Return type

[**CreateEnvironmentUserResponse**](CreateEnvironmentUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHost

> CreateHostResponse CreateHost(ctx, environmentId).HostCreateParameters(hostCreateParameters).Execute()

Create a new Host.

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
    environmentId := "environmentId_example" // string | The ID of the environment.
    hostCreateParameters := *openapiclient.NewHostCreateParameters() // HostCreateParameters | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.CreateHost(context.Background(), environmentId).HostCreateParameters(hostCreateParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CreateHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHost`: CreateHostResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.CreateHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostCreateParameters** | [**HostCreateParameters**](HostCreateParameters.md) |  | 

### Return type

[**CreateHostResponse**](CreateHostResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironment

> DeleteEnvironmentResponse DeleteEnvironment(ctx, environmentId).Execute()

Delete an environment by ID.

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
    environmentId := "environmentId_example" // string | The ID of the environment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.DeleteEnvironment(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.DeleteEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEnvironment`: DeleteEnvironmentResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.DeleteEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteEnvironmentResponse**](DeleteEnvironmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironmentTags

> DeleteEnvironmentTags(ctx, environmentId).DeleteTag(deleteTag).Execute()

Delete tags for an Environment.

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
    environmentId := "environmentId_example" // string | The ID of the environment.
    deleteTag := *openapiclient.NewDeleteTag() // DeleteTag | The parameters to delete tags (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EnvironmentsApi.DeleteEnvironmentTags(context.Background(), environmentId).DeleteTag(deleteTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.DeleteEnvironmentTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentTagsRequest struct via the builder pattern


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


## DeleteEnvironmentUser

> DeleteEnvironmentUserResponse DeleteEnvironmentUser(ctx, environmentId, userRef).Execute()

Delete environment user.

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
    environmentId := "environmentId_example" // string | The ID of the environment.
    userRef := "userRef_example" // string | Environment user reference.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.DeleteEnvironmentUser(context.Background(), environmentId, userRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.DeleteEnvironmentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEnvironmentUser`: DeleteEnvironmentUserResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.DeleteEnvironmentUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 
**userRef** | **string** | Environment user reference. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteEnvironmentUserResponse**](DeleteEnvironmentUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHost

> DeleteHostResponse DeleteHost(ctx, environmentId, hostId).Execute()

Delete a Host.

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
    environmentId := "environmentId_example" // string | The ID of the environment.
    hostId := "hostId_example" // string | The ID of the host.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.DeleteHost(context.Background(), environmentId, hostId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.DeleteHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteHost`: DeleteHostResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.DeleteHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 
**hostId** | **string** | The ID of the host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteHostResponse**](DeleteHostResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableEnvironment

> DisableEnvironmentResponse DisableEnvironment(ctx, environmentId).Execute()

Disable environment.

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
    environmentId := "environmentId_example" // string | The ID of the environment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.DisableEnvironment(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.DisableEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableEnvironment`: DisableEnvironmentResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.DisableEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DisableEnvironmentResponse**](DisableEnvironmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableEnvironment

> EnableEnvironmentResponse EnableEnvironment(ctx, environmentId).Execute()

Enable a disabled environment.

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
    environmentId := "environmentId_example" // string | The ID of the environment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnableEnvironment(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnableEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableEnvironment`: EnableEnvironmentResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnableEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnableEnvironmentResponse**](EnableEnvironmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentById

> Environment GetEnvironmentById(ctx, environmentId).Execute()

Returns an environment by ID.

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
    environmentId := "environmentId_example" // string | The ID of the environment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.GetEnvironmentById(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.GetEnvironmentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentById`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.GetEnvironmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Environment**](Environment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironments

> ListEnvironmentsResponse GetEnvironments(ctx).Limit(limit).Cursor(cursor).Sort(sort).Execute()

List all environments.

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
    resp, r, err := apiClient.EnvironmentsApi.GetEnvironments(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.GetEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironments`: ListEnvironmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.GetEnvironments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 

### Return type

[**ListEnvironmentsResponse**](ListEnvironmentsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagsEnvironment

> TagsResponse GetTagsEnvironment(ctx, environmentId).Execute()

Get tags for an Environment.

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
    environmentId := "environmentId_example" // string | The ID of the environment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.GetTagsEnvironment(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.GetTagsEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagsEnvironment`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.GetTagsEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsEnvironmentRequest struct via the builder pattern


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


## ListEnvironmentUsers

> ListEnvironmentUsers ListEnvironmentUsers(ctx, environmentId).Execute()

List environment users.

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
    environmentId := "environmentId_example" // string | The ID of the environment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.ListEnvironmentUsers(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.ListEnvironmentUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEnvironmentUsers`: ListEnvironmentUsers
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.ListEnvironmentUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListEnvironmentUsers**](ListEnvironmentUsers.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrimaryEnvironmentUser

> PrimaryEnvironmentUserResponse PrimaryEnvironmentUser(ctx, environmentId, userRef).Execute()

Set primary environment user.

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
    environmentId := "environmentId_example" // string | The ID of the environment.
    userRef := "userRef_example" // string | Environment user reference.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.PrimaryEnvironmentUser(context.Background(), environmentId, userRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.PrimaryEnvironmentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrimaryEnvironmentUser`: PrimaryEnvironmentUserResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.PrimaryEnvironmentUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 
**userRef** | **string** | Environment user reference. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrimaryEnvironmentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PrimaryEnvironmentUserResponse**](PrimaryEnvironmentUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshEnvironment

> RefreshEnvironmentResponse RefreshEnvironment(ctx, environmentId).Execute()

Refresh environment.

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
    environmentId := "environmentId_example" // string | The ID of the environment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.RefreshEnvironment(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.RefreshEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshEnvironment`: RefreshEnvironmentResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.RefreshEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RefreshEnvironmentResponse**](RefreshEnvironmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchEnvironments

> SearchEnvironmentsResponse SearchEnvironments(ctx).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()

Search for environments.

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
    resp, r, err := apiClient.EnvironmentsApi.SearchEnvironments(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.SearchEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchEnvironments`: SearchEnvironmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.SearchEnvironments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 
 **searchBody** | [**SearchBody**](SearchBody.md) | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS &#39;foobar&#39;, field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN [&#39;Goku&#39;, &#39;Vegeta&#39;] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically &#x60;SEARCH &#39;12&#39;&#x60; would match an item with an attribute with an integer value of &#x60;123&#x60;.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ &#39;Goku&#39; | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ &#39;Goku&#39; |  ## Grouping Parenthesis &#x60;()&#x60; can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS &#39;foo&#39;)  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \&quot;foo\&quot;, \&quot;bar\&quot;, \&quot;foo bar\&quot;, &#39;foo&#39;, &#39;bar&#39;, &#39;foo bar&#39; | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], [&#39;foo&#39;, \&quot;bar\&quot;] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  | 

### Return type

[**SearchEnvironmentsResponse**](SearchEnvironmentsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnvironment

> UpdateEnvironmentResponse UpdateEnvironment(ctx, environmentId).EnvironmentUpdateParameters(environmentUpdateParameters).Execute()

Update an environment by ID.

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
    environmentId := "environmentId_example" // string | The ID of the environment.
    environmentUpdateParameters := *openapiclient.NewEnvironmentUpdateParameters() // EnvironmentUpdateParameters | the parameters to update an environment (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.UpdateEnvironment(context.Background(), environmentId).EnvironmentUpdateParameters(environmentUpdateParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.UpdateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEnvironment`: UpdateEnvironmentResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.UpdateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentUpdateParameters** | [**EnvironmentUpdateParameters**](EnvironmentUpdateParameters.md) | the parameters to update an environment | 

### Return type

[**UpdateEnvironmentResponse**](UpdateEnvironmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnvironmentUser

> UpdateEnvironmentUserResponse UpdateEnvironmentUser(ctx, environmentId, userRef).EnvironmentUserParams(environmentUserParams).Execute()

Update environment user.

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
    environmentId := "environmentId_example" // string | The ID of the environment.
    userRef := "userRef_example" // string | Environment user reference.
    environmentUserParams := *openapiclient.NewEnvironmentUserParams() // EnvironmentUserParams | The parameters to create an environment user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.UpdateEnvironmentUser(context.Background(), environmentId, userRef).EnvironmentUserParams(environmentUserParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.UpdateEnvironmentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEnvironmentUser`: UpdateEnvironmentUserResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.UpdateEnvironmentUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 
**userRef** | **string** | Environment user reference. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **environmentUserParams** | [**EnvironmentUserParams**](EnvironmentUserParams.md) | The parameters to create an environment user. | 

### Return type

[**UpdateEnvironmentUserResponse**](UpdateEnvironmentUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHost

> UpdateHostResponse UpdateHost(ctx, environmentId, hostId).HostUpdateParameters(hostUpdateParameters).Execute()

Update a Host.

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
    environmentId := "environmentId_example" // string | The ID of the environment.
    hostId := "hostId_example" // string | The ID of the host.
    hostUpdateParameters := *openapiclient.NewHostUpdateParameters() // HostUpdateParameters | the parameters to update a host.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.UpdateHost(context.Background(), environmentId, hostId).HostUpdateParameters(hostUpdateParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.UpdateHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHost`: UpdateHostResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.UpdateHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 
**hostId** | **string** | The ID of the host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hostUpdateParameters** | [**HostUpdateParameters**](HostUpdateParameters.md) | the parameters to update a host. | 

### Return type

[**UpdateHostResponse**](UpdateHostResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRepository

> UpdateRepositoryResponse UpdateRepository(ctx, environmentId, repositoryId).UpdateRepositoryParameters(updateRepositoryParameters).Execute()

Update a Repository.

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
    environmentId := "environmentId_example" // string | The ID of the environment.
    repositoryId := "repositoryId_example" // string | The ID of the repository.
    updateRepositoryParameters := *openapiclient.NewUpdateRepositoryParameters("DatabaseType_example") // UpdateRepositoryParameters | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.UpdateRepository(context.Background(), environmentId, repositoryId).UpdateRepositoryParameters(updateRepositoryParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.UpdateRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRepository`: UpdateRepositoryResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.UpdateRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment. | 
**repositoryId** | **string** | The ID of the repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRepositoryParameters** | [**UpdateRepositoryParameters**](UpdateRepositoryParameters.md) |  | 

### Return type

[**UpdateRepositoryResponse**](UpdateRepositoryResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

