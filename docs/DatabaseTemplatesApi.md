# \DatabaseTemplatesApi

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatabaseTemplate**](DatabaseTemplatesApi.md#CreateDatabaseTemplate) | **Post** /database-templates | Create a database template.
[**CreateDatabaseTemplateTags**](DatabaseTemplatesApi.md#CreateDatabaseTemplateTags) | **Post** /database-templates/{databaseTemplateId}/tags | Create tags for a DatabaseTemplate.
[**DeleteDatabaseTemplate**](DatabaseTemplatesApi.md#DeleteDatabaseTemplate) | **Delete** /database-templates/{databaseTemplateId} | Delete a DatabaseTemplate by ID.
[**DeleteDatabaseTemplateTag**](DatabaseTemplatesApi.md#DeleteDatabaseTemplateTag) | **Post** /database-templates/{databaseTemplateId}/tags/delete | Delete tags for a DatabaseTemplate.
[**GetDatabaseTemplateById**](DatabaseTemplatesApi.md#GetDatabaseTemplateById) | **Get** /database-templates/{databaseTemplateId} | Retrieve a DatabaseTemplate by ID.
[**GetDatabaseTemplateTags**](DatabaseTemplatesApi.md#GetDatabaseTemplateTags) | **Get** /database-templates/{databaseTemplateId}/tags | Get tags for a DatabaseTemplate.
[**GetDatabaseTemplates**](DatabaseTemplatesApi.md#GetDatabaseTemplates) | **Get** /database-templates | Retrieve the list of database templates.
[**ImportDatabaseTemplates**](DatabaseTemplatesApi.md#ImportDatabaseTemplates) | **Post** /database-templates/import | Imports the database templates from an an engine.
[**SearchDatabaseTemplates**](DatabaseTemplatesApi.md#SearchDatabaseTemplates) | **Post** /database-templates/search | Search DatabaseTemplates.
[**UndoImportDatabaseTemplates**](DatabaseTemplatesApi.md#UndoImportDatabaseTemplates) | **Post** /database-templates/undo-import | Undo an import of DatabaseTemplates on an Engine.
[**UpdateDatabaseTemplate**](DatabaseTemplatesApi.md#UpdateDatabaseTemplate) | **Patch** /database-templates/{databaseTemplateId} | Updates a DatabaseTemplate by ID



## CreateDatabaseTemplate

> CreateDatabaseTemplateResponse CreateDatabaseTemplate(ctx).DatabaseTemplateCreateParameters(databaseTemplateCreateParameters).Execute()

Create a database template.

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
    databaseTemplateCreateParameters := *openapiclient.NewDatabaseTemplateCreateParameters("template-name", "OracleVirtualSource") // DatabaseTemplateCreateParameters | The parameters to create a database template.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseTemplatesApi.CreateDatabaseTemplate(context.Background()).DatabaseTemplateCreateParameters(databaseTemplateCreateParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseTemplatesApi.CreateDatabaseTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabaseTemplate`: CreateDatabaseTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `DatabaseTemplatesApi.CreateDatabaseTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseTemplateCreateParameters** | [**DatabaseTemplateCreateParameters**](DatabaseTemplateCreateParameters.md) | The parameters to create a database template. | 

### Return type

[**CreateDatabaseTemplateResponse**](CreateDatabaseTemplateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDatabaseTemplateTags

> TagsResponse CreateDatabaseTemplateTags(ctx, databaseTemplateId).TagsRequest(tagsRequest).Execute()

Create tags for a DatabaseTemplate.

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
    databaseTemplateId := "databaseTemplateId_example" // string | The ID of the Database Template.
    tagsRequest := *openapiclient.NewTagsRequest([]openapiclient.Tag{*openapiclient.NewTag("key-1", "value-1")}) // TagsRequest | Tags information for a DatabaseTemplate.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseTemplatesApi.CreateDatabaseTemplateTags(context.Background(), databaseTemplateId).TagsRequest(tagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseTemplatesApi.CreateDatabaseTemplateTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabaseTemplateTags`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `DatabaseTemplatesApi.CreateDatabaseTemplateTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseTemplateId** | **string** | The ID of the Database Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseTemplateTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagsRequest** | [**TagsRequest**](TagsRequest.md) | Tags information for a DatabaseTemplate. | 

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


## DeleteDatabaseTemplate

> DeleteDatabaseTemplateResponse DeleteDatabaseTemplate(ctx, databaseTemplateId).Execute()

Delete a DatabaseTemplate by ID.

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
    databaseTemplateId := "databaseTemplateId_example" // string | The ID of the Database Template.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseTemplatesApi.DeleteDatabaseTemplate(context.Background(), databaseTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseTemplatesApi.DeleteDatabaseTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDatabaseTemplate`: DeleteDatabaseTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `DatabaseTemplatesApi.DeleteDatabaseTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseTemplateId** | **string** | The ID of the Database Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteDatabaseTemplateResponse**](DeleteDatabaseTemplateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabaseTemplateTag

> DeleteDatabaseTemplateTag(ctx, databaseTemplateId).DeleteTag(deleteTag).Execute()

Delete tags for a DatabaseTemplate.

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
    databaseTemplateId := "databaseTemplateId_example" // string | The ID of the Database Template.
    deleteTag := *openapiclient.NewDeleteTag() // DeleteTag | The parameters to delete tags (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DatabaseTemplatesApi.DeleteDatabaseTemplateTag(context.Background(), databaseTemplateId).DeleteTag(deleteTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseTemplatesApi.DeleteDatabaseTemplateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseTemplateId** | **string** | The ID of the Database Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseTemplateTagRequest struct via the builder pattern


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


## GetDatabaseTemplateById

> DatabaseTemplate GetDatabaseTemplateById(ctx, databaseTemplateId).Execute()

Retrieve a DatabaseTemplate by ID.

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
    databaseTemplateId := "databaseTemplateId_example" // string | The ID of the Database Template.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseTemplatesApi.GetDatabaseTemplateById(context.Background(), databaseTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseTemplatesApi.GetDatabaseTemplateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseTemplateById`: DatabaseTemplate
    fmt.Fprintf(os.Stdout, "Response from `DatabaseTemplatesApi.GetDatabaseTemplateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseTemplateId** | **string** | The ID of the Database Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseTemplateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatabaseTemplate**](DatabaseTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseTemplateTags

> TagsResponse GetDatabaseTemplateTags(ctx, databaseTemplateId).Execute()

Get tags for a DatabaseTemplate.

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
    databaseTemplateId := "databaseTemplateId_example" // string | The ID of the Database Template.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseTemplatesApi.GetDatabaseTemplateTags(context.Background(), databaseTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseTemplatesApi.GetDatabaseTemplateTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseTemplateTags`: TagsResponse
    fmt.Fprintf(os.Stdout, "Response from `DatabaseTemplatesApi.GetDatabaseTemplateTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseTemplateId** | **string** | The ID of the Database Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseTemplateTagsRequest struct via the builder pattern


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


## GetDatabaseTemplates

> ListDatabaseTemplatesResponse GetDatabaseTemplates(ctx).Limit(limit).Cursor(cursor).Sort(sort).Execute()

Retrieve the list of database templates.

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
    resp, r, err := apiClient.DatabaseTemplatesApi.GetDatabaseTemplates(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseTemplatesApi.GetDatabaseTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseTemplates`: ListDatabaseTemplatesResponse
    fmt.Fprintf(os.Stdout, "Response from `DatabaseTemplatesApi.GetDatabaseTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 

### Return type

[**ListDatabaseTemplatesResponse**](ListDatabaseTemplatesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportDatabaseTemplates

> ImportDatabaseTemplates(ctx).EngineIdBody(engineIdBody).Execute()

Imports the database templates from an an engine.

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
    engineIdBody := *openapiclient.NewEngineIdBody() // EngineIdBody | Body containing the ID of the registered engine.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DatabaseTemplatesApi.ImportDatabaseTemplates(context.Background()).EngineIdBody(engineIdBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseTemplatesApi.ImportDatabaseTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportDatabaseTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engineIdBody** | [**EngineIdBody**](EngineIdBody.md) | Body containing the ID of the registered engine. | 

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


## SearchDatabaseTemplates

> SearchDatabaseTemplatesResponse SearchDatabaseTemplates(ctx).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()

Search DatabaseTemplates.

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
    resp, r, err := apiClient.DatabaseTemplatesApi.SearchDatabaseTemplates(context.Background()).Limit(limit).Cursor(cursor).Sort(sort).SearchBody(searchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseTemplatesApi.SearchDatabaseTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchDatabaseTemplates`: SearchDatabaseTemplatesResponse
    fmt.Fprintf(os.Stdout, "Response from `DatabaseTemplatesApi.SearchDatabaseTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDatabaseTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. The value of this property must be extracted from the &#39;prev_cursor&#39; or &#39;next_cursor&#39; property of a PaginatedResponseMetadata which is contained in the response of list and search API endpoints. | 
 **sort** | **string** | The field to sort results by. A property name with a prepended &#39;-&#39; signifies descending order. | 
 **searchBody** | [**SearchBody**](SearchBody.md) | A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS &#39;foobar&#39;, field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN [&#39;Goku&#39;, &#39;Vegeta&#39;] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically &#x60;SEARCH &#39;12&#39;&#x60; would match an item with an attribute with an integer value of &#x60;123&#x60;.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ &#39;Goku&#39; | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ &#39;Goku&#39; |  ## Grouping Parenthesis &#x60;()&#x60; can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS &#39;foo&#39;)  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \&quot;foo\&quot;, \&quot;bar\&quot;, \&quot;foo bar\&quot;, &#39;foo&#39;, &#39;bar&#39;, &#39;foo bar&#39; | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], [&#39;foo&#39;, \&quot;bar\&quot;] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression.  | 

### Return type

[**SearchDatabaseTemplatesResponse**](SearchDatabaseTemplatesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UndoImportDatabaseTemplates

> UndoImportDatabaseTemplates(ctx).EngineIdBody(engineIdBody).Execute()

Undo an import of DatabaseTemplates on an Engine.

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
    engineIdBody := *openapiclient.NewEngineIdBody() // EngineIdBody | Body containing the ID of the registered engine.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DatabaseTemplatesApi.UndoImportDatabaseTemplates(context.Background()).EngineIdBody(engineIdBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseTemplatesApi.UndoImportDatabaseTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUndoImportDatabaseTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engineIdBody** | [**EngineIdBody**](EngineIdBody.md) | Body containing the ID of the registered engine. | 

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


## UpdateDatabaseTemplate

> UpdateDatabaseTemplateResponse UpdateDatabaseTemplate(ctx, databaseTemplateId).UpdateDatabaseTemplateParameters(updateDatabaseTemplateParameters).Execute()

Updates a DatabaseTemplate by ID

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
    databaseTemplateId := "databaseTemplateId_example" // string | The ID of the Database Template.
    updateDatabaseTemplateParameters := *openapiclient.NewUpdateDatabaseTemplateParameters() // UpdateDatabaseTemplateParameters | The new data to update a VDB. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseTemplatesApi.UpdateDatabaseTemplate(context.Background(), databaseTemplateId).UpdateDatabaseTemplateParameters(updateDatabaseTemplateParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseTemplatesApi.UpdateDatabaseTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDatabaseTemplate`: UpdateDatabaseTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `DatabaseTemplatesApi.UpdateDatabaseTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseTemplateId** | **string** | The ID of the Database Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatabaseTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDatabaseTemplateParameters** | [**UpdateDatabaseTemplateParameters**](UpdateDatabaseTemplateParameters.md) | The new data to update a VDB. | 

### Return type

[**UpdateDatabaseTemplateResponse**](UpdateDatabaseTemplateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

