# \AuthorizationApi

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccessGroupAccountIds**](AuthorizationApi.md#AddAccessGroupAccountIds) | **Post** /access-groups/{accessGroupId}/account-ids | Add account ids to an Access group
[**AddAccessGroupAccountTags**](AuthorizationApi.md#AddAccessGroupAccountTags) | **Post** /access-groups/{accessGroupId}/tags | Add account tags to an Access group
[**AddAccessGroupScopes**](AuthorizationApi.md#AddAccessGroupScopes) | **Post** /access-groups/{accessGroupId}/scopes | Add scopes to an Access group
[**AddAlwaysAllowedPermissions**](AuthorizationApi.md#AddAlwaysAllowedPermissions) | **Post** /access-groups/{accessGroupId}/scopes/{scopeId}/always_allowed_permissions | Add always allowed permissions for given object type.
[**AddObjectsToAccessGroupScope**](AuthorizationApi.md#AddObjectsToAccessGroupScope) | **Post** /access-groups/{accessGroupId}/scopes/{scopeId}/objects | Add objects to the access group scope.
[**AddRolePermissions**](AuthorizationApi.md#AddRolePermissions) | **Post** /roles/{roleId}/permissions | Add permissions to a role.
[**AddTagsToScope**](AuthorizationApi.md#AddTagsToScope) | **Post** /access-groups/{accessGroupId}/scopes/{scopeId}/object-tags | Add object tags to the access group scope.
[**CreateAccessGroup**](AuthorizationApi.md#CreateAccessGroup) | **Post** /access-groups | Create a new access group.
[**CreateRole**](AuthorizationApi.md#CreateRole) | **Post** /roles | Create custom role
[**CreateRoleTags**](AuthorizationApi.md#CreateRoleTags) | **Post** /roles/{roleId}/tags | Create tags for a role.
[**DeleteAccessGroup**](AuthorizationApi.md#DeleteAccessGroup) | **Delete** /access-groups/{accessGroupId} | Delete an Access group.
[**DeleteAccessGroupScopeObjectTags**](AuthorizationApi.md#DeleteAccessGroupScopeObjectTags) | **Post** /access-groups/{accessGroupId}/scopes/{scopeId}/object-tags/delete | Remove tags from the access group scope.
[**DeleteAccessGroupScopeObjects**](AuthorizationApi.md#DeleteAccessGroupScopeObjects) | **Post** /access-groups/{accessGroupId}/scopes/{scopeId}/objects/delete | Remove objects from the access group scope.
[**DeleteRole**](AuthorizationApi.md#DeleteRole) | **Delete** /roles/{roleId} | Delete role by ID.
[**DeleteRoleTag**](AuthorizationApi.md#DeleteRoleTag) | **Post** /roles/{roleId}/tags/delete | Delete tags for a Role.
[**GetAccessGroupById**](AuthorizationApi.md#GetAccessGroupById) | **Get** /access-groups/{accessGroupId} | Returns an Access group by ID.
[**GetAccessGroupScope**](AuthorizationApi.md#GetAccessGroupScope) | **Get** /access-groups/{accessGroupId}/scopes/{scopeId} | Get access group scope.
[**GetAccessGroups**](AuthorizationApi.md#GetAccessGroups) | **Get** /access-groups | List all access groups.
[**GetAllObjectPermissions**](AuthorizationApi.md#GetAllObjectPermissions) | **Get** /auth/object-permissions | Returns all of the possible permissions for all of the objects.
[**GetObjectPermissions**](AuthorizationApi.md#GetObjectPermissions) | **Get** /auth/permissions/objects/{objectType}/{objectId} | Returns permissions for given object.
[**GetRoleById**](AuthorizationApi.md#GetRoleById) | **Get** /roles/{roleId} | Returns role by ID.
[**GetRoleTags**](AuthorizationApi.md#GetRoleTags) | **Get** /roles/{roleId}/tags | Get tags for a Role.
[**GetRoles**](AuthorizationApi.md#GetRoles) | **Get** /roles | List all roles
[**RemoveAccessGroupAccountId**](AuthorizationApi.md#RemoveAccessGroupAccountId) | **Delete** /access-groups/{accessGroupId}/account-ids/{accountId} | Remove the account from the access group.
[**RemoveAccessGroupAccountTags**](AuthorizationApi.md#RemoveAccessGroupAccountTags) | **Post** /access-groups/{accessGroupId}/tags/delete | Remove account tags from an access group.
[**RemoveAccessGroupScope**](AuthorizationApi.md#RemoveAccessGroupScope) | **Delete** /access-groups/{accessGroupId}/scopes/{scopeId} | Remove the scope from the Access group.
[**RemoveAlwaysAllowedPermissions**](AuthorizationApi.md#RemoveAlwaysAllowedPermissions) | **Post** /access-groups/{accessGroupId}/scopes/{scopeId}/always_allowed_permissions/delete | Remove always allowed permissions for given object type.
[**RemoveRolePermissions**](AuthorizationApi.md#RemoveRolePermissions) | **Post** /roles/{roleId}/permissions/delete | Remove permissions from a role.
[**SearchAccessGroups**](AuthorizationApi.md#SearchAccessGroups) | **Post** /access-groups/search | Search for access groups.
[**SearchRoles**](AuthorizationApi.md#SearchRoles) | **Post** /roles/search | Search for roles.
[**UpdateAccessGroup**](AuthorizationApi.md#UpdateAccessGroup) | **Patch** /access-groups/{accessGroupId} | Update an Access group.
[**UpdateAccessGroupScope**](AuthorizationApi.md#UpdateAccessGroupScope) | **Patch** /access-groups/{accessGroupId}/scopes/{scopeId} | Update access group scope.
[**UpdateRole**](AuthorizationApi.md#UpdateRole) | **Patch** /roles/{roleId} | Update a Role.



## AddAccessGroupAccountIds

> AccessGroup AddAccessGroupAccountIds(ctx, accessGroupId).AccessGroupAccountIdsRequest(accessGroupAccountIdsRequest).Execute()

Add account ids to an Access group

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
    accessGroupId := "accessGroupId_example" // string | The ID of the Access group.
    accessGroupAccountIdsRequest := *openapiclient.NewAccessGroupAccountIdsRequest([]int64{int64(123)}) // AccessGroupAccountIdsRequest | Account ids to add to the Access group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.AddAccessGroupAccountIds(context.Background(), accessGroupId).AccessGroupAccountIdsRequest(accessGroupAccountIdsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.AddAccessGroupAccountIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAccessGroupAccountIds`: AccessGroup
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.AddAccessGroupAccountIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessGroupId** | **string** | The ID of the Access group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAccessGroupAccountIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessGroupAccountIdsRequest** | [**AccessGroupAccountIdsRequest**](AccessGroupAccountIdsRequest.md) | Account ids to add to the Access group. | 

### Return type

[**AccessGroup**](AccessGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAccessGroupAccountTags

> AccessGroup AddAccessGroupAccountTags(ctx, accessGroupId).TagsRequest(tagsRequest).Execute()

Add account tags to an Access group

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
    accessGroupId := "accessGroupId_example" // string | The ID of the Access group.
    tagsRequest := *openapiclient.NewTagsRequest([]openapiclient.Tag{*openapiclient.NewTag("key-1", "value-1")}) // TagsRequest | Account Tags to add to the Access group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.AddAccessGroupAccountTags(context.Background(), accessGroupId).TagsRequest(tagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.AddAccessGroupAccountTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAccessGroupAccountTags`: AccessGroup
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.AddAccessGroupAccountTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessGroupId** | **string** | The ID of the Access group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAccessGroupAccountTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagsRequest** | [**TagsRequest**](TagsRequest.md) | Account Tags to add to the Access group. | 

### Return type

[**AccessGroup**](AccessGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAccessGroupScopes

> AccessGroup AddAccessGroupScopes(ctx, accessGroupId).AccessGroupScopesRequest(accessGroupScopesRequest).Execute()

Add scopes to an Access group

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
    accessGroupId := "accessGroupId_example" // string | The ID of the Access group.
    accessGroupScopesRequest := *openapiclient.NewAccessGroupScopesRequest([]openapiclient.AccessGroupScope{*openapiclient.NewAccessGroupScope("RoleId_example")}) // AccessGroupScopesRequest | Scopes to add to the Access group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.AddAccessGroupScopes(context.Background(), accessGroupId).AccessGroupScopesRequest(accessGroupScopesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.AddAccessGroupScopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAccessGroupScopes`: AccessGroup
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.AddAccessGroupScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessGroupId** | **string** | The ID of the Access group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAccessGroupScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessGroupScopesRequest** | [**AccessGroupScopesRequest**](AccessGroupScopesRequest.md) | Scopes to add to the Access group. | 

### Return type

[**AccessGroup**](AccessGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAlwaysAllowedPermissions

> AccessGroupScope AddAlwaysAllowedPermissions(ctx, accessGroupId, scopeId).AlwaysAllowedPermissionRequest(alwaysAllowedPermissionRequest).Execute()

Add always allowed permissions for given object type.

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
    accessGroupId := "accessGroupId_example" // string | The ID of the Access group.
    scopeId := "scopeId_example" // string | The ID of the Access group scope.
    alwaysAllowedPermissionRequest := *openapiclient.NewAlwaysAllowedPermissionRequest([]openapiclient.AlwaysAllowedPermission{*openapiclient.NewAlwaysAllowedPermission(openapiclient.ObjectTypeEnum("ACCESS_GROUP"), openapiclient.PermissionEnum("READ"))}) // AlwaysAllowedPermissionRequest | Add always allowed permissions for given object type.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.AddAlwaysAllowedPermissions(context.Background(), accessGroupId, scopeId).AlwaysAllowedPermissionRequest(alwaysAllowedPermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.AddAlwaysAllowedPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAlwaysAllowedPermissions`: AccessGroupScope
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.AddAlwaysAllowedPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessGroupId** | **string** | The ID of the Access group. | 
**scopeId** | **string** | The ID of the Access group scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAlwaysAllowedPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alwaysAllowedPermissionRequest** | [**AlwaysAllowedPermissionRequest**](AlwaysAllowedPermissionRequest.md) | Add always allowed permissions for given object type. | 

### Return type

[**AccessGroupScope**](AccessGroupScope.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddObjectsToAccessGroupScope

> ScopedObjectItemsResponse AddObjectsToAccessGroupScope(ctx, accessGroupId, scopeId).ScopedObjectsRequest(scopedObjectsRequest).Execute()

Add objects to the access group scope.

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
    accessGroupId := "accessGroupId_example" // string | The ID of the Access group.
    scopeId := "scopeId_example" // string | The ID of the Access group scope.
    scopedObjectsRequest := *openapiclient.NewScopedObjectsRequest([]openapiclient.ScopedObjectItem{*openapiclient.NewScopedObjectItem("1-VDB-OBJECT-ID", openapiclient.ObjectTypeEnum("ACCESS_GROUP"))}) // ScopedObjectsRequest | Add objects to the access group scope.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.AddObjectsToAccessGroupScope(context.Background(), accessGroupId, scopeId).ScopedObjectsRequest(scopedObjectsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.AddObjectsToAccessGroupScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddObjectsToAccessGroupScope`: ScopedObjectItemsResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.AddObjectsToAccessGroupScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessGroupId** | **string** | The ID of the Access group. | 
**scopeId** | **string** | The ID of the Access group scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddObjectsToAccessGroupScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scopedObjectsRequest** | [**ScopedObjectsRequest**](ScopedObjectsRequest.md) | Add objects to the access group scope. | 

### Return type

[**ScopedObjectItemsResponse**](ScopedObjectItemsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRolePermissions

> Role AddRolePermissions(ctx, roleId).PermissionsRequest(permissionsRequest).Execute()

Add permissions to a role.

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
    roleId := "roleId_example" // string | The ID of the role.
    permissionsRequest := *openapiclient.NewPermissionsRequest([]openapiclient.PermissionObject{*openapiclient.NewPermissionObject("VDB", []string{"Permissions_example"})}) // PermissionsRequest | Permissions to add to the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.AddRolePermissions(context.Background(), roleId).PermissionsRequest(permissionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.AddRolePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRolePermissions`: Role
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.AddRolePermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The ID of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRolePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permissionsRequest** | [**PermissionsRequest**](PermissionsRequest.md) | Permissions to add to the role. | 

### Return type

[**Role**](Role.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTagsToScope

> ScopeTagsResponse AddTagsToScope(ctx, accessGroupId, scopeId).ScopeTagsRequest(scopeTagsRequest).Execute()

Add object tags to the access group scope.

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
    accessGroupId := "accessGroupId_example" // string | The ID of the Access group.
    scopeId := "scopeId_example" // string | The ID of the Access group scope.
    scopeTagsRequest := *openapiclient.NewScopeTagsRequest([]openapiclient.ScopeTag{*openapiclient.NewScopeTag("key-1", "value-1")}) // ScopeTagsRequest | Object tags for the access group scope.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.AddTagsToScope(context.Background(), accessGroupId, scopeId).ScopeTagsRequest(scopeTagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.AddTagsToScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTagsToScope`: ScopeTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.AddTagsToScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessGroupId** | **string** | The ID of the Access group. | 
**scopeId** | **string** | The ID of the Access group scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTagsToScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scopeTagsRequest** | [**ScopeTagsRequest**](ScopeTagsRequest.md) | Object tags for the access group scope. | 

### Return type

[**ScopeTagsResponse**](ScopeTagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccessGroup

> AccessGroup CreateAccessGroup(ctx).AccessGroup(accessGroup).Execute()

Create a new access group.

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
    accessGroup := *openapiclient.NewAccessGroup("Name_example") // AccessGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.CreateAccessGroup(context.Background()).AccessGroup(accessGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.CreateAccessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessGroup`: AccessGroup
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.CreateAccessGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessGroup** | [**AccessGroup**](AccessGroup.md) |  | 

### Return type

[**AccessGroup**](AccessGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> Role CreateRole(ctx).CreateRole(createRole).Execute()

Create custom role

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
    createRole := *openapiclient.NewCreateRole("Name_example", []openapiclient.PermissionObject{*openapiclient.NewPermissionObject("VDB", []string{"Permissions_example"})}) // CreateRole |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.CreateRole(context.Background()).CreateRole(createRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.CreateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.CreateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRole** | [**CreateRole**](CreateRole.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoleTags

> TagsResponse CreateRoleTags(ctx, roleId).TagsRequest(tagsRequest).Execute()

Create tags for a role.

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
    roleId := "roleId_example" // string | The ID of the role.
    tagsRequest := *openapiclient.NewTagsRequest([]openapiclient.Tag{*openapiclient.NewTag("key-1", "value-1")}) // TagsRequest | Tags information for Roles.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.CreateRoleTags(context.Background(), roleId).TagsRequest(tagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.CreateRoleTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRoleTags`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.CreateRoleTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The ID of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagsRequest** | [**TagsRequest**](TagsRequest.md) | Tags information for Roles. | 

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


## DeleteAccessGroup

> DeleteAccessGroup(ctx, accessGroupId).Execute()

Delete an Access group.

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
    accessGroupId := "accessGroupId_example" // string | The ID of the Access group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationApi.DeleteAccessGroup(context.Background(), accessGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.DeleteAccessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessGroupId** | **string** | The ID of the Access group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessGroupRequest struct via the builder pattern


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


## DeleteAccessGroupScopeObjectTags

> ScopeTagsResponse DeleteAccessGroupScopeObjectTags(ctx, accessGroupId, scopeId).DeleteScopeObjectTags(deleteScopeObjectTags).Execute()

Remove tags from the access group scope.

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
    accessGroupId := "accessGroupId_example" // string | The ID of the Access group.
    scopeId := "scopeId_example" // string | The ID of the Access group scope.
    deleteScopeObjectTags := *openapiclient.NewDeleteScopeObjectTags() // DeleteScopeObjectTags | The parameters to delete scope objects tags (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.DeleteAccessGroupScopeObjectTags(context.Background(), accessGroupId, scopeId).DeleteScopeObjectTags(deleteScopeObjectTags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.DeleteAccessGroupScopeObjectTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccessGroupScopeObjectTags`: ScopeTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.DeleteAccessGroupScopeObjectTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessGroupId** | **string** | The ID of the Access group. | 
**scopeId** | **string** | The ID of the Access group scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessGroupScopeObjectTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteScopeObjectTags** | [**DeleteScopeObjectTags**](DeleteScopeObjectTags.md) | The parameters to delete scope objects tags | 

### Return type

[**ScopeTagsResponse**](ScopeTagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessGroupScopeObjects

> ScopedObjectItemsResponse DeleteAccessGroupScopeObjects(ctx, accessGroupId, scopeId).DeleteScopedObjectItem(deleteScopedObjectItem).Execute()

Remove objects from the access group scope.

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
    accessGroupId := "accessGroupId_example" // string | The ID of the Access group.
    scopeId := "scopeId_example" // string | The ID of the Access group scope.
    deleteScopedObjectItem := *openapiclient.NewDeleteScopedObjectItem([]openapiclient.ScopedObjectItem{*openapiclient.NewScopedObjectItem("1-VDB-OBJECT-ID", openapiclient.ObjectTypeEnum("ACCESS_GROUP"))}) // DeleteScopedObjectItem | The parameters to delete scope objects (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.DeleteAccessGroupScopeObjects(context.Background(), accessGroupId, scopeId).DeleteScopedObjectItem(deleteScopedObjectItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.DeleteAccessGroupScopeObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccessGroupScopeObjects`: ScopedObjectItemsResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.DeleteAccessGroupScopeObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessGroupId** | **string** | The ID of the Access group. | 
**scopeId** | **string** | The ID of the Access group scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessGroupScopeObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteScopedObjectItem** | [**DeleteScopedObjectItem**](DeleteScopedObjectItem.md) | The parameters to delete scope objects | 

### Return type

[**ScopedObjectItemsResponse**](ScopedObjectItemsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, roleId).Execute()

Delete role by ID.

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
    roleId := "roleId_example" // string | The ID of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationApi.DeleteRole(context.Background(), roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.DeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The ID of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


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


## DeleteRoleTag

> DeleteRoleTag(ctx, roleId).DeleteTag(deleteTag).Execute()

Delete tags for a Role.

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
    roleId := "roleId_example" // string | The ID of the role.
    deleteTag := *openapiclient.NewDeleteTag() // DeleteTag | The parameters to delete tags (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationApi.DeleteRoleTag(context.Background(), roleId).DeleteTag(deleteTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.DeleteRoleTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The ID of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleTagRequest struct via the builder pattern


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


## GetAccessGroupById

> AccessGroup GetAccessGroupById(ctx, accessGroupId).Execute()

Returns an Access group by ID.

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
    accessGroupId := "accessGroupId_example" // string | The ID of the Access group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.GetAccessGroupById(context.Background(), accessGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.GetAccessGroupById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessGroupById`: AccessGroup
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.GetAccessGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessGroupId** | **string** | The ID of the Access group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessGroup**](AccessGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessGroupScope

> AccessGroupScope GetAccessGroupScope(ctx, accessGroupId, scopeId).Execute()

Get access group scope.

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
    accessGroupId := "accessGroupId_example" // string | The ID of the Access group.
    scopeId := "scopeId_example" // string | The ID of the Access group scope.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.GetAccessGroupScope(context.Background(), accessGroupId, scopeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.GetAccessGroupScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessGroupScope`: AccessGroupScope
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.GetAccessGroupScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessGroupId** | **string** | The ID of the Access group. | 
**scopeId** | **string** | The ID of the Access group scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessGroupScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AccessGroupScope**](AccessGroupScope.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessGroups

> ListAccessGroupsResponse GetAccessGroups(ctx).Limit(limit).Cursor(cursor).Sort(sort).Execute()

List all access groups.

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
    resp, r, err := apiClient.AuthorizationApi.GetAccessGroups(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.GetAccessGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessGroups`: ListAccessGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.GetAccessGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 

### Return type

[**ListAccessGroupsResponse**](ListAccessGroupsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllObjectPermissions

> AllObjectPermissionsResponse GetAllObjectPermissions(ctx).Execute()

Returns all of the possible permissions for all of the objects.

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
    resp, r, err := apiClient.AuthorizationApi.GetAllObjectPermissions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.GetAllObjectPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllObjectPermissions`: AllObjectPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.GetAllObjectPermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllObjectPermissionsRequest struct via the builder pattern


### Return type

[**AllObjectPermissionsResponse**](AllObjectPermissionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectPermissions

> ObjectPermissionsResponse GetObjectPermissions(ctx, objectType, objectId).Execute()

Returns permissions for given object.

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
    objectType := "objectType_example" // string | The type of the DCT object.
    objectId := "objectId_example" // string | The ID of the DCT Object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.GetObjectPermissions(context.Background(), objectType, objectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.GetObjectPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectPermissions`: ObjectPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.GetObjectPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of the DCT object. | 
**objectId** | **string** | The ID of the DCT Object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ObjectPermissionsResponse**](ObjectPermissionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleById

> Role GetRoleById(ctx, roleId).Execute()

Returns role by ID.

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
    roleId := "roleId_example" // string | The ID of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.GetRoleById(context.Background(), roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.GetRoleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleById`: Role
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.GetRoleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The ID of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Role**](Role.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleTags

> TagsResponse GetRoleTags(ctx, roleId).Execute()

Get tags for a Role.

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
    roleId := "roleId_example" // string | The ID of the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.GetRoleTags(context.Background(), roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.GetRoleTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleTags`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.GetRoleTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The ID of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleTagsRequest struct via the builder pattern


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


## GetRoles

> ListRolesResponse GetRoles(ctx).Execute()

List all roles

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
    resp, r, err := apiClient.AuthorizationApi.GetRoles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.GetRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoles`: ListRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.GetRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesRequest struct via the builder pattern


### Return type

[**ListRolesResponse**](ListRolesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAccessGroupAccountId

> AccessGroup RemoveAccessGroupAccountId(ctx, accessGroupId, accountId).Execute()

Remove the account from the access group.

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
    accessGroupId := "accessGroupId_example" // string | The ID of the Access group.
    accountId := int64(789) // int64 | The ID of the account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.RemoveAccessGroupAccountId(context.Background(), accessGroupId, accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.RemoveAccessGroupAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveAccessGroupAccountId`: AccessGroup
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.RemoveAccessGroupAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessGroupId** | **string** | The ID of the Access group. | 
**accountId** | **int64** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAccessGroupAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AccessGroup**](AccessGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAccessGroupAccountTags

> AccessGroup RemoveAccessGroupAccountTags(ctx, accessGroupId).DeleteTag(deleteTag).Execute()

Remove account tags from an access group.

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
    accessGroupId := "accessGroupId_example" // string | The ID of the Access group.
    deleteTag := *openapiclient.NewDeleteTag() // DeleteTag | The parameters to delete tags (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.RemoveAccessGroupAccountTags(context.Background(), accessGroupId).DeleteTag(deleteTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.RemoveAccessGroupAccountTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveAccessGroupAccountTags`: AccessGroup
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.RemoveAccessGroupAccountTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessGroupId** | **string** | The ID of the Access group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAccessGroupAccountTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteTag** | [**DeleteTag**](DeleteTag.md) | The parameters to delete tags | 

### Return type

[**AccessGroup**](AccessGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAccessGroupScope

> AccessGroup RemoveAccessGroupScope(ctx, accessGroupId, scopeId).Execute()

Remove the scope from the Access group.

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
    accessGroupId := "accessGroupId_example" // string | The ID of the Access group.
    scopeId := "scopeId_example" // string | The ID of the Access group scope.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.RemoveAccessGroupScope(context.Background(), accessGroupId, scopeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.RemoveAccessGroupScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveAccessGroupScope`: AccessGroup
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.RemoveAccessGroupScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessGroupId** | **string** | The ID of the Access group. | 
**scopeId** | **string** | The ID of the Access group scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAccessGroupScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AccessGroup**](AccessGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAlwaysAllowedPermissions

> AccessGroupScope RemoveAlwaysAllowedPermissions(ctx, accessGroupId, scopeId).AlwaysAllowedPermissionRequest(alwaysAllowedPermissionRequest).Execute()

Remove always allowed permissions for given object type.

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
    accessGroupId := "accessGroupId_example" // string | The ID of the Access group.
    scopeId := "scopeId_example" // string | The ID of the Access group scope.
    alwaysAllowedPermissionRequest := *openapiclient.NewAlwaysAllowedPermissionRequest([]openapiclient.AlwaysAllowedPermission{*openapiclient.NewAlwaysAllowedPermission(openapiclient.ObjectTypeEnum("ACCESS_GROUP"), openapiclient.PermissionEnum("READ"))}) // AlwaysAllowedPermissionRequest | Remove always allowed permissions for given object type.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.RemoveAlwaysAllowedPermissions(context.Background(), accessGroupId, scopeId).AlwaysAllowedPermissionRequest(alwaysAllowedPermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.RemoveAlwaysAllowedPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveAlwaysAllowedPermissions`: AccessGroupScope
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.RemoveAlwaysAllowedPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessGroupId** | **string** | The ID of the Access group. | 
**scopeId** | **string** | The ID of the Access group scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAlwaysAllowedPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alwaysAllowedPermissionRequest** | [**AlwaysAllowedPermissionRequest**](AlwaysAllowedPermissionRequest.md) | Remove always allowed permissions for given object type. | 

### Return type

[**AccessGroupScope**](AccessGroupScope.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRolePermissions

> Role RemoveRolePermissions(ctx, roleId).PermissionsRequest(permissionsRequest).Execute()

Remove permissions from a role.

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
    roleId := "roleId_example" // string | The ID of the role.
    permissionsRequest := *openapiclient.NewPermissionsRequest([]openapiclient.PermissionObject{*openapiclient.NewPermissionObject("VDB", []string{"Permissions_example"})}) // PermissionsRequest | Permissions to remove from the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.RemoveRolePermissions(context.Background(), roleId).PermissionsRequest(permissionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.RemoveRolePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveRolePermissions`: Role
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.RemoveRolePermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The ID of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRolePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permissionsRequest** | [**PermissionsRequest**](PermissionsRequest.md) | Permissions to remove from the role. | 

### Return type

[**Role**](Role.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAccessGroups

> SearchAccessGroupsResponse SearchAccessGroups(ctx).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()

Search for access groups.

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
    resp, r, err := apiClient.AuthorizationApi.SearchAccessGroups(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.SearchAccessGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAccessGroups`: SearchAccessGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.SearchAccessGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAccessGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 
 **searchBody** | [**SearchBody**](SearchBody.md) | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS &#39;foobar&#39;, field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN [&#39;Goku&#39;, &#39;Vegeta&#39;] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically &#x60;SEARCH &#39;12&#39;&#x60; would match an item with an attribute with an integer value of &#x60;123&#x60;.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ &#39;Goku&#39; | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ &#39;Goku&#39; |  ## Grouping Parenthesis &#x60;()&#x60; can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS &#39;foo&#39;)  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \&quot;foo\&quot;, \&quot;bar\&quot;, \&quot;foo bar\&quot;, &#39;foo&#39;, &#39;bar&#39;, &#39;foo bar&#39; | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], [&#39;foo&#39;, \&quot;bar\&quot;] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  | 

### Return type

[**SearchAccessGroupsResponse**](SearchAccessGroupsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRoles

> SearchRolesResponse SearchRoles(ctx).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()

Search for roles.

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
    resp, r, err := apiClient.AuthorizationApi.SearchRoles(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.SearchRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchRoles`: SearchRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.SearchRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 
 **searchBody** | [**SearchBody**](SearchBody.md) | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS &#39;foobar&#39;, field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN [&#39;Goku&#39;, &#39;Vegeta&#39;] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically &#x60;SEARCH &#39;12&#39;&#x60; would match an item with an attribute with an integer value of &#x60;123&#x60;.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ &#39;Goku&#39; | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ &#39;Goku&#39; |  ## Grouping Parenthesis &#x60;()&#x60; can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS &#39;foo&#39;)  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \&quot;foo\&quot;, \&quot;bar\&quot;, \&quot;foo bar\&quot;, &#39;foo&#39;, &#39;bar&#39;, &#39;foo bar&#39; | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], [&#39;foo&#39;, \&quot;bar\&quot;] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  | 

### Return type

[**SearchRolesResponse**](SearchRolesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessGroup

> AccessGroup UpdateAccessGroup(ctx, accessGroupId).AccessGroupUpdateParameters(accessGroupUpdateParameters).Execute()

Update an Access group.

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
    accessGroupId := "accessGroupId_example" // string | The ID of the Access group.
    accessGroupUpdateParameters := *openapiclient.NewAccessGroupUpdateParameters() // AccessGroupUpdateParameters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.UpdateAccessGroup(context.Background(), accessGroupId).AccessGroupUpdateParameters(accessGroupUpdateParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.UpdateAccessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccessGroup`: AccessGroup
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.UpdateAccessGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessGroupId** | **string** | The ID of the Access group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessGroupUpdateParameters** | [**AccessGroupUpdateParameters**](AccessGroupUpdateParameters.md) |  | 

### Return type

[**AccessGroup**](AccessGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessGroupScope

> AccessGroupScope UpdateAccessGroupScope(ctx, accessGroupId, scopeId).UpdateAccessGroupScope(updateAccessGroupScope).Execute()

Update access group scope.

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
    accessGroupId := "accessGroupId_example" // string | The ID of the Access group.
    scopeId := "scopeId_example" // string | The ID of the Access group scope.
    updateAccessGroupScope := *openapiclient.NewUpdateAccessGroupScope() // UpdateAccessGroupScope | Access group scope to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.UpdateAccessGroupScope(context.Background(), accessGroupId, scopeId).UpdateAccessGroupScope(updateAccessGroupScope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.UpdateAccessGroupScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccessGroupScope`: AccessGroupScope
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.UpdateAccessGroupScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessGroupId** | **string** | The ID of the Access group. | 
**scopeId** | **string** | The ID of the Access group scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessGroupScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateAccessGroupScope** | [**UpdateAccessGroupScope**](UpdateAccessGroupScope.md) | Access group scope to update. | 

### Return type

[**AccessGroupScope**](AccessGroupScope.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> Role UpdateRole(ctx, roleId).RoleUpdateParameters(roleUpdateParameters).Execute()

Update a Role.

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
    roleId := "roleId_example" // string | The ID of the role.
    roleUpdateParameters := *openapiclient.NewRoleUpdateParameters() // RoleUpdateParameters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.UpdateRole(context.Background(), roleId).RoleUpdateParameters(roleUpdateParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.UpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.UpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The ID of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleUpdateParameters** | [**RoleUpdateParameters**](RoleUpdateParameters.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

