# \ManagementApi

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEngineTags**](ManagementApi.md#CreateEngineTags) | **Post** /management/engines/{engineId}/tags | Create tags for an Engine.
[**CreateHashicorpVault**](ManagementApi.md#CreateHashicorpVault) | **Post** /management/vaults/hashicorp | Configure a new Hashicorp Vault
[**CreateHashicorpVaultTags**](ManagementApi.md#CreateHashicorpVaultTags) | **Post** /management/vaults/hashicorp/{vaultId}/tags | Create tags for a Hashicorp vault.
[**DeleteEngineTags**](ManagementApi.md#DeleteEngineTags) | **Post** /management/engines/{engineId}/tags/delete | Delete tags for an Engine.
[**DeleteHashicorpVault**](ManagementApi.md#DeleteHashicorpVault) | **Delete** /management/vaults/hashicorp/{vaultId} | Delete a Hashicorp vault by id
[**DeleteHashicorpVaultTag**](ManagementApi.md#DeleteHashicorpVaultTag) | **Post** /management/vaults/hashicorp/{vaultId}/tags/delete | Delete tags for a Hashicorp vault.
[**GetApiClassification**](ManagementApi.md#GetApiClassification) | **Get** /management/api-classification | Get api classification.
[**GetEngineTags**](ManagementApi.md#GetEngineTags) | **Get** /management/engines/{engineId}/tags | Get tags for a Engine.
[**GetHashicorpVault**](ManagementApi.md#GetHashicorpVault) | **Get** /management/vaults/hashicorp/{vaultId} | Get a Hashicorp vault by id
[**GetHashicorpVaultTags**](ManagementApi.md#GetHashicorpVaultTags) | **Get** /management/vaults/hashicorp/{vaultId}/tags | Get tags for a Hashicorp vault.
[**GetHashicorpVaults**](ManagementApi.md#GetHashicorpVaults) | **Get** /management/vaults/hashicorp | Returns a list of configured Hashicorp vaults.
[**GetLdapConfig**](ManagementApi.md#GetLdapConfig) | **Get** /management/ldap-config | Returns the LDAP configuration
[**GetMetadataDatabase**](ManagementApi.md#GetMetadataDatabase) | **Get** /management/metadata-database | Returns configuration information about the metadata database which stores the product data.
[**GetProxyConfiguration**](ManagementApi.md#GetProxyConfiguration) | **Get** /management/proxy | Returns the current Proxy configuration.
[**GetRegisteredEngine**](ManagementApi.md#GetRegisteredEngine) | **Get** /management/engines/{engineId} | Returns a registered engine by ID.
[**GetRegisteredEngines**](ManagementApi.md#GetRegisteredEngines) | **Get** /management/engines | Returns a list of registered engines.
[**GetSamlConfig**](ManagementApi.md#GetSamlConfig) | **Get** /management/saml-config | Returns the SAML configuration
[**GetSmtpConfig**](ManagementApi.md#GetSmtpConfig) | **Get** /management/smtp | Returns the SMTP configuration
[**ListProperties**](ManagementApi.md#ListProperties) | **Get** /management/properties | Get global properties.
[**RegisterEngine**](ManagementApi.md#RegisterEngine) | **Post** /management/engines | Register an engine.
[**SearchEngines**](ManagementApi.md#SearchEngines) | **Post** /management/engines/search | Search for engines.
[**SearchHashicorpVaults**](ManagementApi.md#SearchHashicorpVaults) | **Post** /management/vaults/hashicorp/search | Search for configured Hashicorp vaults.
[**UnregisterEngine**](ManagementApi.md#UnregisterEngine) | **Delete** /management/engines/{engineId} | Unregister an engine.
[**UpdateApiClassification**](ManagementApi.md#UpdateApiClassification) | **Put** /management/api-classification | Update the api classification to new version.
[**UpdateLdapConfig**](ManagementApi.md#UpdateLdapConfig) | **Put** /management/ldap-config | Update LDAP Config.
[**UpdateProperties**](ManagementApi.md#UpdateProperties) | **Patch** /management/properties | Update value of predefined properties.
[**UpdateProxyConfiguration**](ManagementApi.md#UpdateProxyConfiguration) | **Put** /management/proxy | Update Proxy configuration.
[**UpdateRegisteredEngine**](ManagementApi.md#UpdateRegisteredEngine) | **Put** /management/engines/{engineId} | Update a registered engine.
[**UpdateSamlConfig**](ManagementApi.md#UpdateSamlConfig) | **Put** /management/saml-config | Update SAML Config.
[**UpdateSmtpConfig**](ManagementApi.md#UpdateSmtpConfig) | **Put** /management/smtp | Update SMTP Config.
[**ValidateJavaPath**](ManagementApi.md#ValidateJavaPath) | **Post** /management/engines/{engineId}/validate/java-path | Validate java path for the remote host machine.
[**ValidateLdapConfig**](ManagementApi.md#ValidateLdapConfig) | **Post** /management/ldap-config/validate | Validate LDAP Config. Without username/password, DCT performs an anonymous bind against the LDAP server. If credentials are provided DCT validates that authentication and mapping of optional properties are actually working with provided credentials. LDAP search is only validated if search attributes are set.
[**ValidateSmtpConfig**](ManagementApi.md#ValidateSmtpConfig) | **Post** /management/smtp/validate | Validate SMTP Config.



## CreateEngineTags

> TagsResponse CreateEngineTags(ctx, engineId).TagsRequest(tagsRequest).Execute()

Create tags for an Engine.

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
    engineId := "engineId_example" // string | The ID of the registered engine.
    tagsRequest := *openapiclient.NewTagsRequest([]openapiclient.Tag{*openapiclient.NewTag("key-1", "value-1")}) // TagsRequest | Tags information for Engine.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.CreateEngineTags(context.Background(), engineId).TagsRequest(tagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.CreateEngineTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEngineTags`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.CreateEngineTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**engineId** | **string** | The ID of the registered engine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEngineTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagsRequest** | [**TagsRequest**](TagsRequest.md) | Tags information for Engine. | 

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


## CreateHashicorpVault

> HashicorpVault CreateHashicorpVault(ctx).HashicorpVault(hashicorpVault).Execute()

Configure a new Hashicorp Vault

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
    hashicorpVault := *openapiclient.NewHashicorpVault() // HashicorpVault | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.CreateHashicorpVault(context.Background()).HashicorpVault(hashicorpVault).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.CreateHashicorpVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHashicorpVault`: HashicorpVault
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.CreateHashicorpVault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHashicorpVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hashicorpVault** | [**HashicorpVault**](HashicorpVault.md) |  | 

### Return type

[**HashicorpVault**](HashicorpVault.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHashicorpVaultTags

> TagsResponse CreateHashicorpVaultTags(ctx, vaultId).TagsRequest(tagsRequest).Execute()

Create tags for a Hashicorp vault.

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
    vaultId := int64(789) // int64 | Numeric ID of the Hashicorp vault
    tagsRequest := *openapiclient.NewTagsRequest([]openapiclient.Tag{*openapiclient.NewTag("key-1", "value-1")}) // TagsRequest | Tags information for Hashicorp vault.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.CreateHashicorpVaultTags(context.Background(), vaultId).TagsRequest(tagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.CreateHashicorpVaultTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHashicorpVaultTags`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.CreateHashicorpVaultTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **int64** | Numeric ID of the Hashicorp vault | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHashicorpVaultTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagsRequest** | [**TagsRequest**](TagsRequest.md) | Tags information for Hashicorp vault. | 

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


## DeleteEngineTags

> DeleteEngineTags(ctx, engineId).DeleteTag(deleteTag).Execute()

Delete tags for an Engine.

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
    engineId := "engineId_example" // string | The ID of the registered engine.
    deleteTag := *openapiclient.NewDeleteTag() // DeleteTag | The parameters to delete tags (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagementApi.DeleteEngineTags(context.Background(), engineId).DeleteTag(deleteTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.DeleteEngineTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**engineId** | **string** | The ID of the registered engine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEngineTagsRequest struct via the builder pattern


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


## DeleteHashicorpVault

> DeleteHashicorpVault(ctx, vaultId).Execute()

Delete a Hashicorp vault by id

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
    vaultId := int64(789) // int64 | Numeric ID of the Hashicorp vault

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagementApi.DeleteHashicorpVault(context.Background(), vaultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.DeleteHashicorpVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **int64** | Numeric ID of the Hashicorp vault | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHashicorpVaultRequest struct via the builder pattern


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


## DeleteHashicorpVaultTag

> DeleteHashicorpVaultTag(ctx, vaultId).DeleteTag(deleteTag).Execute()

Delete tags for a Hashicorp vault.

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
    vaultId := int64(789) // int64 | Numeric ID of the Hashicorp vault
    deleteTag := *openapiclient.NewDeleteTag() // DeleteTag | The parameters to delete tags (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagementApi.DeleteHashicorpVaultTag(context.Background(), vaultId).DeleteTag(deleteTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.DeleteHashicorpVaultTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **int64** | Numeric ID of the Hashicorp vault | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHashicorpVaultTagRequest struct via the builder pattern


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


## GetApiClassification

> APIClassificationConfig GetApiClassification(ctx).Execute()

Get api classification.

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
    resp, r, err := apiClient.ManagementApi.GetApiClassification(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.GetApiClassification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiClassification`: APIClassificationConfig
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.GetApiClassification`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiClassificationRequest struct via the builder pattern


### Return type

[**APIClassificationConfig**](APIClassificationConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEngineTags

> TagsResponse GetEngineTags(ctx, engineId).Execute()

Get tags for a Engine.

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
    engineId := "engineId_example" // string | The ID of the registered engine.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.GetEngineTags(context.Background(), engineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.GetEngineTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEngineTags`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.GetEngineTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**engineId** | **string** | The ID of the registered engine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEngineTagsRequest struct via the builder pattern


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


## GetHashicorpVault

> HashicorpVault GetHashicorpVault(ctx, vaultId).Execute()

Get a Hashicorp vault by id

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
    vaultId := int64(789) // int64 | Numeric ID of the Hashicorp vault

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.GetHashicorpVault(context.Background(), vaultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.GetHashicorpVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHashicorpVault`: HashicorpVault
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.GetHashicorpVault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **int64** | Numeric ID of the Hashicorp vault | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHashicorpVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HashicorpVault**](HashicorpVault.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHashicorpVaultTags

> TagsResponse GetHashicorpVaultTags(ctx, vaultId).Execute()

Get tags for a Hashicorp vault.

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
    vaultId := int64(789) // int64 | Numeric ID of the Hashicorp vault

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.GetHashicorpVaultTags(context.Background(), vaultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.GetHashicorpVaultTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHashicorpVaultTags`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.GetHashicorpVaultTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **int64** | Numeric ID of the Hashicorp vault | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHashicorpVaultTagsRequest struct via the builder pattern


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


## GetHashicorpVaults

> ListHashicorpVaultsResponse GetHashicorpVaults(ctx).Limit(limit).Cursor(cursor).Sort(sort).Execute()

Returns a list of configured Hashicorp vaults.

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
    resp, r, err := apiClient.ManagementApi.GetHashicorpVaults(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.GetHashicorpVaults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHashicorpVaults`: ListHashicorpVaultsResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.GetHashicorpVaults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHashicorpVaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 

### Return type

[**ListHashicorpVaultsResponse**](ListHashicorpVaultsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLdapConfig

> LDAPConfigParams GetLdapConfig(ctx).Execute()

Returns the LDAP configuration

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
    resp, r, err := apiClient.ManagementApi.GetLdapConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.GetLdapConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLdapConfig`: LDAPConfigParams
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.GetLdapConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapConfigRequest struct via the builder pattern


### Return type

[**LDAPConfigParams**](LDAPConfigParams.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadataDatabase

> MetadataDbInfo GetMetadataDatabase(ctx).Execute()

Returns configuration information about the metadata database which stores the product data.

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
    resp, r, err := apiClient.ManagementApi.GetMetadataDatabase(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.GetMetadataDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetadataDatabase`: MetadataDbInfo
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.GetMetadataDatabase`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataDatabaseRequest struct via the builder pattern


### Return type

[**MetadataDbInfo**](MetadataDbInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProxyConfiguration

> ProxyConfiguration GetProxyConfiguration(ctx).Execute()

Returns the current Proxy configuration.

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
    resp, r, err := apiClient.ManagementApi.GetProxyConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.GetProxyConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProxyConfiguration`: ProxyConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.GetProxyConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProxyConfigurationRequest struct via the builder pattern


### Return type

[**ProxyConfiguration**](ProxyConfiguration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegisteredEngine

> RegisteredEngine GetRegisteredEngine(ctx, engineId).Execute()

Returns a registered engine by ID.

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
    engineId := "engineId_example" // string | The ID of the registered engine.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.GetRegisteredEngine(context.Background(), engineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.GetRegisteredEngine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegisteredEngine`: RegisteredEngine
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.GetRegisteredEngine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**engineId** | **string** | The ID of the registered engine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegisteredEngine**](RegisteredEngine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegisteredEngines

> ListRegisteredEnginesResponse GetRegisteredEngines(ctx).Limit(limit).Cursor(cursor).Sort(sort).Execute()

Returns a list of registered engines.

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
    resp, r, err := apiClient.ManagementApi.GetRegisteredEngines(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.GetRegisteredEngines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegisteredEngines`: ListRegisteredEnginesResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.GetRegisteredEngines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredEnginesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 

### Return type

[**ListRegisteredEnginesResponse**](ListRegisteredEnginesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSamlConfig

> SAMLConfigParams GetSamlConfig(ctx).Execute()

Returns the SAML configuration

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
    resp, r, err := apiClient.ManagementApi.GetSamlConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.GetSamlConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSamlConfig`: SAMLConfigParams
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.GetSamlConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSamlConfigRequest struct via the builder pattern


### Return type

[**SAMLConfigParams**](SAMLConfigParams.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmtpConfig

> SMTPConfigParams GetSmtpConfig(ctx).Execute()

Returns the SMTP configuration

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
    resp, r, err := apiClient.ManagementApi.GetSmtpConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.GetSmtpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmtpConfig`: SMTPConfigParams
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.GetSmtpConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmtpConfigRequest struct via the builder pattern


### Return type

[**SMTPConfigParams**](SMTPConfigParams.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProperties

> GlobalProperties ListProperties(ctx).Execute()

Get global properties.

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
    resp, r, err := apiClient.ManagementApi.ListProperties(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.ListProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProperties`: GlobalProperties
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.ListProperties`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPropertiesRequest struct via the builder pattern


### Return type

[**GlobalProperties**](GlobalProperties.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterEngine

> RegisteredEngine RegisterEngine(ctx).EngineRegistrationParameter(engineRegistrationParameter).Execute()

Register an engine.

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
    engineRegistrationParameter := *openapiclient.NewEngineRegistrationParameter("Name_example", "Hostname_example") // EngineRegistrationParameter | The parameters to register an engine.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.RegisterEngine(context.Background()).EngineRegistrationParameter(engineRegistrationParameter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.RegisterEngine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterEngine`: RegisteredEngine
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.RegisterEngine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engineRegistrationParameter** | [**EngineRegistrationParameter**](EngineRegistrationParameter.md) | The parameters to register an engine. | 

### Return type

[**RegisteredEngine**](RegisteredEngine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchEngines

> SearchEnginesResponse SearchEngines(ctx).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()

Search for engines.

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
    resp, r, err := apiClient.ManagementApi.SearchEngines(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.SearchEngines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchEngines`: SearchEnginesResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.SearchEngines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchEnginesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 
 **searchBody** | [**SearchBody**](SearchBody.md) | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS &#39;foobar&#39;, field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN [&#39;Goku&#39;, &#39;Vegeta&#39;] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically &#x60;SEARCH &#39;12&#39;&#x60; would match an item with an attribute with an integer value of &#x60;123&#x60;.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ &#39;Goku&#39; | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ &#39;Goku&#39; |  ## Grouping Parenthesis &#x60;()&#x60; can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS &#39;foo&#39;)  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \&quot;foo\&quot;, \&quot;bar\&quot;, \&quot;foo bar\&quot;, &#39;foo&#39;, &#39;bar&#39;, &#39;foo bar&#39; | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], [&#39;foo&#39;, \&quot;bar\&quot;] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  | 

### Return type

[**SearchEnginesResponse**](SearchEnginesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchHashicorpVaults

> SearchHashicorpVaultsResponse SearchHashicorpVaults(ctx).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()

Search for configured Hashicorp vaults.

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
    resp, r, err := apiClient.ManagementApi.SearchHashicorpVaults(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.SearchHashicorpVaults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchHashicorpVaults`: SearchHashicorpVaultsResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.SearchHashicorpVaults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchHashicorpVaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 
 **searchBody** | [**SearchBody**](SearchBody.md) | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS &#39;foobar&#39;, field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN [&#39;Goku&#39;, &#39;Vegeta&#39;] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically &#x60;SEARCH &#39;12&#39;&#x60; would match an item with an attribute with an integer value of &#x60;123&#x60;.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ &#39;Goku&#39; | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ &#39;Goku&#39; |  ## Grouping Parenthesis &#x60;()&#x60; can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS &#39;foo&#39;)  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \&quot;foo\&quot;, \&quot;bar\&quot;, \&quot;foo bar\&quot;, &#39;foo&#39;, &#39;bar&#39;, &#39;foo bar&#39; | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], [&#39;foo&#39;, \&quot;bar\&quot;] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  | 

### Return type

[**SearchHashicorpVaultsResponse**](SearchHashicorpVaultsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterEngine

> DeleteEngineResponse UnregisterEngine(ctx, engineId).Execute()

Unregister an engine.

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
    engineId := "engineId_example" // string | The ID of the registered engine.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.UnregisterEngine(context.Background(), engineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.UnregisterEngine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnregisterEngine`: DeleteEngineResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.UnregisterEngine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**engineId** | **string** | The ID of the registered engine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteEngineResponse**](DeleteEngineResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiClassification

> APIClassificationConfig UpdateApiClassification(ctx).APIClassificationConfig(aPIClassificationConfig).Execute()

Update the api classification to new version.

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
    aPIClassificationConfig := *openapiclient.NewAPIClassificationConfig() // APIClassificationConfig | Request to update api classification config.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.UpdateApiClassification(context.Background()).APIClassificationConfig(aPIClassificationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.UpdateApiClassification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiClassification`: APIClassificationConfig
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.UpdateApiClassification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiClassificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aPIClassificationConfig** | [**APIClassificationConfig**](APIClassificationConfig.md) | Request to update api classification config. | 

### Return type

[**APIClassificationConfig**](APIClassificationConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLdapConfig

> LDAPConfigParams UpdateLdapConfig(ctx).LDAPConfigParams(lDAPConfigParams).Execute()

Update LDAP Config.

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
    lDAPConfigParams := *openapiclient.NewLDAPConfigParams() // LDAPConfigParams | The parameters to update the LDAP config.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.UpdateLdapConfig(context.Background()).LDAPConfigParams(lDAPConfigParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.UpdateLdapConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLdapConfig`: LDAPConfigParams
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.UpdateLdapConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLdapConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPConfigParams** | [**LDAPConfigParams**](LDAPConfigParams.md) | The parameters to update the LDAP config. | 

### Return type

[**LDAPConfigParams**](LDAPConfigParams.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProperties

> GlobalProperties UpdateProperties(ctx).GlobalProperties(globalProperties).Execute()

Update value of predefined properties.

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
    globalProperties := *openapiclient.NewGlobalProperties() // GlobalProperties | The parameters to update property value.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.UpdateProperties(context.Background()).GlobalProperties(globalProperties).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.UpdateProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProperties`: GlobalProperties
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.UpdateProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalProperties** | [**GlobalProperties**](GlobalProperties.md) | The parameters to update property value. | 

### Return type

[**GlobalProperties**](GlobalProperties.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProxyConfiguration

> ProxyConfiguration UpdateProxyConfiguration(ctx).ProxyConfiguration(proxyConfiguration).Execute()

Update Proxy configuration.

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
    proxyConfiguration := *openapiclient.NewProxyConfiguration("proxy.server.com", int32(3128), true) // ProxyConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.UpdateProxyConfiguration(context.Background()).ProxyConfiguration(proxyConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.UpdateProxyConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProxyConfiguration`: ProxyConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.UpdateProxyConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProxyConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proxyConfiguration** | [**ProxyConfiguration**](ProxyConfiguration.md) |  | 

### Return type

[**ProxyConfiguration**](ProxyConfiguration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRegisteredEngine

> RegisteredEngine UpdateRegisteredEngine(ctx, engineId).RegisteredEngine(registeredEngine).Execute()

Update a registered engine.

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
    engineId := "engineId_example" // string | The ID of the registered engine.
    registeredEngine := *openapiclient.NewRegisteredEngine("Name_example", "Hostname_example") // RegisteredEngine | The updated registration engine information.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.UpdateRegisteredEngine(context.Background(), engineId).RegisteredEngine(registeredEngine).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.UpdateRegisteredEngine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRegisteredEngine`: RegisteredEngine
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.UpdateRegisteredEngine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**engineId** | **string** | The ID of the registered engine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRegisteredEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registeredEngine** | [**RegisteredEngine**](RegisteredEngine.md) | The updated registration engine information. | 

### Return type

[**RegisteredEngine**](RegisteredEngine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSamlConfig

> SAMLConfigParams UpdateSamlConfig(ctx).SAMLConfigParams(sAMLConfigParams).Execute()

Update SAML Config.

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
    sAMLConfigParams := *openapiclient.NewSAMLConfigParams() // SAMLConfigParams | The parameters to update the SAML config.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.UpdateSamlConfig(context.Background()).SAMLConfigParams(sAMLConfigParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.UpdateSamlConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSamlConfig`: SAMLConfigParams
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.UpdateSamlConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSamlConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sAMLConfigParams** | [**SAMLConfigParams**](SAMLConfigParams.md) | The parameters to update the SAML config. | 

### Return type

[**SAMLConfigParams**](SAMLConfigParams.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSmtpConfig

> SMTPConfigParams UpdateSmtpConfig(ctx).SMTPConfigParams(sMTPConfigParams).Execute()

Update SMTP Config.

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
    sMTPConfigParams := *openapiclient.NewSMTPConfigParams() // SMTPConfigParams | The parameters to update the SMTP config.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.UpdateSmtpConfig(context.Background()).SMTPConfigParams(sMTPConfigParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.UpdateSmtpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSmtpConfig`: SMTPConfigParams
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.UpdateSmtpConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSmtpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sMTPConfigParams** | [**SMTPConfigParams**](SMTPConfigParams.md) | The parameters to update the SMTP config. | 

### Return type

[**SMTPConfigParams**](SMTPConfigParams.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateJavaPath

> ValidateJavaPath(ctx, engineId).ValidateJavaParameters(validateJavaParameters).Execute()

Validate java path for the remote host machine.

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
    engineId := "engineId_example" // string | The ID of the registered engine.
    validateJavaParameters := *openapiclient.NewValidateJavaParameters("home/jdk/", NullableInt32(22), "test.host.com") // ValidateJavaParameters | The api to check connectivity of engine and a remote host on given port.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagementApi.ValidateJavaPath(context.Background(), engineId).ValidateJavaParameters(validateJavaParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.ValidateJavaPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**engineId** | **string** | The ID of the registered engine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateJavaPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validateJavaParameters** | [**ValidateJavaParameters**](ValidateJavaParameters.md) | The api to check connectivity of engine and a remote host on given port. | 

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


## ValidateLdapConfig

> LdapValidateResponse ValidateLdapConfig(ctx).LdapConfigValidateParameter(ldapConfigValidateParameter).Execute()

Validate LDAP Config. Without username/password, DCT performs an anonymous bind against the LDAP server. If credentials are provided DCT validates that authentication and mapping of optional properties are actually working with provided credentials. LDAP search is only validated if search attributes are set.

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
    ldapConfigValidateParameter := *openapiclient.NewLdapConfigValidateParameter() // LdapConfigValidateParameter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.ValidateLdapConfig(context.Background()).LdapConfigValidateParameter(ldapConfigValidateParameter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.ValidateLdapConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateLdapConfig`: LdapValidateResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.ValidateLdapConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateLdapConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapConfigValidateParameter** | [**LdapConfigValidateParameter**](LdapConfigValidateParameter.md) |  | 

### Return type

[**LdapValidateResponse**](LdapValidateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateSmtpConfig

> ValidateSmtpConfig(ctx).SMTPConfigValidate(sMTPConfigValidate).Execute()

Validate SMTP Config.

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
    sMTPConfigValidate := *openapiclient.NewSMTPConfigValidate("ToAddress_example") // SMTPConfigValidate | The parameters to validate the SMTP config.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagementApi.ValidateSmtpConfig(context.Background()).SMTPConfigValidate(sMTPConfigValidate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.ValidateSmtpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateSmtpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sMTPConfigValidate** | [**SMTPConfigValidate**](SMTPConfigValidate.md) | The parameters to validate the SMTP config. | 

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

