# \HyperscaleInstanceApi

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHyperscaleInstanceTags**](HyperscaleInstanceApi.md#CreateHyperscaleInstanceTags) | **Post** /hyperscale-instances/{hyperscaleInstanceId}/tags | Create tags for a Hyperscale Instance.
[**DeleteHyperscaleInstanceTags**](HyperscaleInstanceApi.md#DeleteHyperscaleInstanceTags) | **Post** /hyperscale-instances/{hyperscaleInstanceId}/tags/delete | Delete tags for a Hyperscale Instance.
[**GetHyperscaleInstanceById**](HyperscaleInstanceApi.md#GetHyperscaleInstanceById) | **Get** /hyperscale-instances/{hyperscaleInstanceId} | Returns a Hyperscale Instance by ID.
[**GetHyperscaleInstanceTags**](HyperscaleInstanceApi.md#GetHyperscaleInstanceTags) | **Get** /hyperscale-instances/{hyperscaleInstanceId}/tags | Get tags for a Hyperscale Instance.
[**GetHyperscaleInstances**](HyperscaleInstanceApi.md#GetHyperscaleInstances) | **Get** /hyperscale-instances | Returns a list of Hyperscale instances.
[**RegisterHyperscaleInstance**](HyperscaleInstanceApi.md#RegisterHyperscaleInstance) | **Post** /hyperscale-instances | Register a Hyperscale instance
[**SearchHyperscaleInstances**](HyperscaleInstanceApi.md#SearchHyperscaleInstances) | **Post** /hyperscale-instances/search | Search for Hyperscale instances.
[**UnregisterHyperscaleInstance**](HyperscaleInstanceApi.md#UnregisterHyperscaleInstance) | **Delete** /hyperscale-instances/{hyperscaleInstanceId} | Unregister a Hyperscale Instance.
[**UpdateHyperscaleInstance**](HyperscaleInstanceApi.md#UpdateHyperscaleInstance) | **Patch** /hyperscale-instances/{hyperscaleInstanceId} | Update an Hyperscale Instance



## CreateHyperscaleInstanceTags

> TagsResponse CreateHyperscaleInstanceTags(ctx, hyperscaleInstanceId).TagsRequest(tagsRequest).Execute()

Create tags for a Hyperscale Instance.

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
    hyperscaleInstanceId := "hyperscaleInstanceId_example" // string | The ID of hyperscale instance.
    tagsRequest := *openapiclient.NewTagsRequest([]openapiclient.Tag{*openapiclient.NewTag("key-1", "value-1")}) // TagsRequest | Tags information for Hyperscale Instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HyperscaleInstanceApi.CreateHyperscaleInstanceTags(context.Background(), hyperscaleInstanceId).TagsRequest(tagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HyperscaleInstanceApi.CreateHyperscaleInstanceTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHyperscaleInstanceTags`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `HyperscaleInstanceApi.CreateHyperscaleInstanceTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hyperscaleInstanceId** | **string** | The ID of hyperscale instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHyperscaleInstanceTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagsRequest** | [**TagsRequest**](TagsRequest.md) | Tags information for Hyperscale Instance. | 

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


## DeleteHyperscaleInstanceTags

> DeleteHyperscaleInstanceTags(ctx, hyperscaleInstanceId).DeleteTag(deleteTag).Execute()

Delete tags for a Hyperscale Instance.

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
    hyperscaleInstanceId := "hyperscaleInstanceId_example" // string | The ID of hyperscale instance.
    deleteTag := *openapiclient.NewDeleteTag() // DeleteTag | The parameters to delete tags (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HyperscaleInstanceApi.DeleteHyperscaleInstanceTags(context.Background(), hyperscaleInstanceId).DeleteTag(deleteTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HyperscaleInstanceApi.DeleteHyperscaleInstanceTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hyperscaleInstanceId** | **string** | The ID of hyperscale instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHyperscaleInstanceTagsRequest struct via the builder pattern


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


## GetHyperscaleInstanceById

> HyperscaleInstance GetHyperscaleInstanceById(ctx, hyperscaleInstanceId).Execute()

Returns a Hyperscale Instance by ID.

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
    hyperscaleInstanceId := "hyperscaleInstanceId_example" // string | The ID of hyperscale instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HyperscaleInstanceApi.GetHyperscaleInstanceById(context.Background(), hyperscaleInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HyperscaleInstanceApi.GetHyperscaleInstanceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHyperscaleInstanceById`: HyperscaleInstance
    fmt.Fprintf(os.Stdout, "Response from `HyperscaleInstanceApi.GetHyperscaleInstanceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hyperscaleInstanceId** | **string** | The ID of hyperscale instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHyperscaleInstanceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HyperscaleInstance**](HyperscaleInstance.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHyperscaleInstanceTags

> TagsResponse GetHyperscaleInstanceTags(ctx, hyperscaleInstanceId).Execute()

Get tags for a Hyperscale Instance.

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
    hyperscaleInstanceId := "hyperscaleInstanceId_example" // string | The ID of hyperscale instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HyperscaleInstanceApi.GetHyperscaleInstanceTags(context.Background(), hyperscaleInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HyperscaleInstanceApi.GetHyperscaleInstanceTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHyperscaleInstanceTags`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `HyperscaleInstanceApi.GetHyperscaleInstanceTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hyperscaleInstanceId** | **string** | The ID of hyperscale instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHyperscaleInstanceTagsRequest struct via the builder pattern


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


## GetHyperscaleInstances

> ListHyperscaleInstancesResponse GetHyperscaleInstances(ctx).Limit(limit).Cursor(cursor).Sort(sort).Execute()

Returns a list of Hyperscale instances.

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
    resp, r, err := apiClient.HyperscaleInstanceApi.GetHyperscaleInstances(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HyperscaleInstanceApi.GetHyperscaleInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHyperscaleInstances`: ListHyperscaleInstancesResponse
    fmt.Fprintf(os.Stdout, "Response from `HyperscaleInstanceApi.GetHyperscaleInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHyperscaleInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 

### Return type

[**ListHyperscaleInstancesResponse**](ListHyperscaleInstancesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterHyperscaleInstance

> HyperscaleInstance RegisterHyperscaleInstance(ctx).HyperscaleInstanceRegistrationParameter(hyperscaleInstanceRegistrationParameter).Execute()

Register a Hyperscale instance

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
    hyperscaleInstanceRegistrationParameter := *openapiclient.NewHyperscaleInstanceRegistrationParameter("Name_example", "Hostname_example", "ApiKey_example") // HyperscaleInstanceRegistrationParameter | The parameters to register a Hyperscale instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HyperscaleInstanceApi.RegisterHyperscaleInstance(context.Background()).HyperscaleInstanceRegistrationParameter(hyperscaleInstanceRegistrationParameter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HyperscaleInstanceApi.RegisterHyperscaleInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterHyperscaleInstance`: HyperscaleInstance
    fmt.Fprintf(os.Stdout, "Response from `HyperscaleInstanceApi.RegisterHyperscaleInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterHyperscaleInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hyperscaleInstanceRegistrationParameter** | [**HyperscaleInstanceRegistrationParameter**](HyperscaleInstanceRegistrationParameter.md) | The parameters to register a Hyperscale instance. | 

### Return type

[**HyperscaleInstance**](HyperscaleInstance.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchHyperscaleInstances

> SearchHyperscaleInstancesResponse SearchHyperscaleInstances(ctx).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()

Search for Hyperscale instances.

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
    resp, r, err := apiClient.HyperscaleInstanceApi.SearchHyperscaleInstances(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HyperscaleInstanceApi.SearchHyperscaleInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchHyperscaleInstances`: SearchHyperscaleInstancesResponse
    fmt.Fprintf(os.Stdout, "Response from `HyperscaleInstanceApi.SearchHyperscaleInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchHyperscaleInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 
 **searchBody** | [**SearchBody**](SearchBody.md) | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS &#39;foobar&#39;, field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN [&#39;Goku&#39;, &#39;Vegeta&#39;] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically &#x60;SEARCH &#39;12&#39;&#x60; would match an item with an attribute with an integer value of &#x60;123&#x60;.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ &#39;Goku&#39; | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ &#39;Goku&#39; |  ## Grouping Parenthesis &#x60;()&#x60; can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS &#39;foo&#39;)  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \&quot;foo\&quot;, \&quot;bar\&quot;, \&quot;foo bar\&quot;, &#39;foo&#39;, &#39;bar&#39;, &#39;foo bar&#39; | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], [&#39;foo&#39;, \&quot;bar\&quot;] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  | 

### Return type

[**SearchHyperscaleInstancesResponse**](SearchHyperscaleInstancesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterHyperscaleInstance

> UnregisterHyperscaleInstanceResponse UnregisterHyperscaleInstance(ctx, hyperscaleInstanceId).Execute()

Unregister a Hyperscale Instance.

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
    hyperscaleInstanceId := "hyperscaleInstanceId_example" // string | The ID of hyperscale instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HyperscaleInstanceApi.UnregisterHyperscaleInstance(context.Background(), hyperscaleInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HyperscaleInstanceApi.UnregisterHyperscaleInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnregisterHyperscaleInstance`: UnregisterHyperscaleInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `HyperscaleInstanceApi.UnregisterHyperscaleInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hyperscaleInstanceId** | **string** | The ID of hyperscale instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterHyperscaleInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UnregisterHyperscaleInstanceResponse**](UnregisterHyperscaleInstanceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHyperscaleInstance

> HyperscaleInstance UpdateHyperscaleInstance(ctx, hyperscaleInstanceId).HyperscaleInstanceUpdateParams(hyperscaleInstanceUpdateParams).Execute()

Update an Hyperscale Instance

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
    hyperscaleInstanceId := "hyperscaleInstanceId_example" // string | The ID of hyperscale instance.
    hyperscaleInstanceUpdateParams := *openapiclient.NewHyperscaleInstanceUpdateParams() // HyperscaleInstanceUpdateParams | Update parameters for a hyperscale instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HyperscaleInstanceApi.UpdateHyperscaleInstance(context.Background(), hyperscaleInstanceId).HyperscaleInstanceUpdateParams(hyperscaleInstanceUpdateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HyperscaleInstanceApi.UpdateHyperscaleInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHyperscaleInstance`: HyperscaleInstance
    fmt.Fprintf(os.Stdout, "Response from `HyperscaleInstanceApi.UpdateHyperscaleInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hyperscaleInstanceId** | **string** | The ID of hyperscale instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHyperscaleInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hyperscaleInstanceUpdateParams** | [**HyperscaleInstanceUpdateParams**](HyperscaleInstanceUpdateParams.md) | Update parameters for a hyperscale instance. | 

### Return type

[**HyperscaleInstance**](HyperscaleInstance.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

