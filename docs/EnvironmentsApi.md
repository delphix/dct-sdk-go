# \EnvironmentsApi

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironmentTags**](EnvironmentsApi.md#CreateEnvironmentTags) | **Post** /environments/{environmentId}/tags | Create tags for an Environment.
[**CreateEnvironments**](EnvironmentsApi.md#CreateEnvironments) | **Post** /environments | Create an environment.
[**DeleteEnvironment**](EnvironmentsApi.md#DeleteEnvironment) | **Delete** /environments/{environmentId} | Delete an environment by ID.
[**DisableEnvironment**](EnvironmentsApi.md#DisableEnvironment) | **Put** /environments/{environmentId}/disable | Disable environment.
[**EnableEnvironment**](EnvironmentsApi.md#EnableEnvironment) | **Put** /environments/{environmentId}/enable | Enable a disabled environment.
[**GetEnvironmentById**](EnvironmentsApi.md#GetEnvironmentById) | **Get** /environments/{environmentId} | Returns an environment by ID.
[**GetEnvironments**](EnvironmentsApi.md#GetEnvironments) | **Get** /environments | List all environments.
[**GetTagsEnvironment**](EnvironmentsApi.md#GetTagsEnvironment) | **Get** /environments/{environmentId}/tags | Get tags for an Environment.
[**ListEnvironmentUsers**](EnvironmentsApi.md#ListEnvironmentUsers) | **Get** /environments/{environmentId}/users | List environment users.
[**RefreshEnvironment**](EnvironmentsApi.md#RefreshEnvironment) | **Put** /environments/{environmentId}/refresh | Refresh environment.
[**UpdateEnvironment**](EnvironmentsApi.md#UpdateEnvironment) | **Put** /environments/{environmentId} | Update an environment by ID.



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
    openapiclient "./openapi"
)

func main() {
    environmentId := "environmentId_example" // string | The ID of the environment.
    tagsRequest := *openapiclient.NewTagsRequest() // TagsRequest | Tags information for Environment.

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


## CreateEnvironments

> CreateEnvironmentResponse CreateEnvironments(ctx).EnvironmentCreateParameters(environmentCreateParameters).Execute()

Create an environment.

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
    environmentCreateParameters := *openapiclient.NewEnvironmentCreateParameters(int64(12), "UNIX", "db.host.com") // EnvironmentCreateParameters | The parameters to create an environment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.CreateEnvironments(context.Background()).EnvironmentCreateParameters(environmentCreateParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CreateEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironments`: CreateEnvironmentResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.CreateEnvironments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentsRequest struct via the builder pattern


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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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

> ListEnvironmentsResponse GetEnvironments(ctx).Limit(limit).Cursor(cursor).Execute()

List all environments.

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
    limit := int32(50) // int32 | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. (optional) (default to 100)
    cursor := "RXlhbCBpcyBncmVhdAo=" // string | Cursor to fetch the next or previous page of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.GetEnvironments(context.Background()).Limit(limit).Cursor(cursor).Execute()
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
 **cursor** | **string** | Cursor to fetch the next or previous page of results. | 

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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
)

func main() {
    environmentId := "environmentId_example" // string | List environment users.

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
**environmentId** | **string** | List environment users. | 

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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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

