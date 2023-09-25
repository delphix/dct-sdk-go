# \LoginApi

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountLogin**](LoginApi.md#AccountLogin) | **Post** /login | Login to Account with Username and Password 
[**AccountLogout**](LoginApi.md#AccountLogout) | **Post** /logout | Invalidates username/password and SSO login session.
[**TokenInfo**](LoginApi.md#TokenInfo) | **Post** /token-info | Get Token Information 



## AccountLogin

> LoginToken AccountLogin(ctx).AccountLoginParameter(accountLoginParameter).Execute()

Login to Account with Username and Password 

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
    accountLoginParameter := *openapiclient.NewAccountLoginParameter("Username_example", "Password_example") // AccountLoginParameter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginApi.AccountLogin(context.Background()).AccountLoginParameter(accountLoginParameter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.AccountLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountLogin`: LoginToken
    fmt.Fprintf(os.Stdout, "Response from `LoginApi.AccountLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountLoginParameter** | [**AccountLoginParameter**](AccountLoginParameter.md) |  | 

### Return type

[**LoginToken**](LoginToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountLogout

> AccountLogout(ctx).Authorization(authorization).Execute()

Invalidates username/password and SSO login session.

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
    authorization := "authorization_example" // string | Access token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LoginApi.AccountLogout(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.AccountLogout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountLogoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Access token | 

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


## TokenInfo

> TokenInfoResponse TokenInfo(ctx).TokenInfoRequest(tokenInfoRequest).Execute()

Get Token Information 

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
    tokenInfoRequest := *openapiclient.NewTokenInfoRequest("Bearer eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIyIiwiaXNzIjoiYXBpZ3ctc2VydmljZXMtYXBwIiwiaWF0IjoxNjYwMDI1MTczLCJlbWFpbCI6ImppdGVuZHJhLm1hdGh1ckBkZWxwaGl4LmNvbSIsInVzZXJuYW1lIjoidXNlci00MjciLCJleHAiOjE2NjAxMTE1NzN9._9LgnIlkKUr2KVqjeFYqru3GxJr2-ztSmP0XO3vBkRo") // TokenInfoRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginApi.TokenInfo(context.Background()).TokenInfoRequest(tokenInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.TokenInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokenInfo`: TokenInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `LoginApi.TokenInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenInfoRequest** | [**TokenInfoRequest**](TokenInfoRequest.md) |  | 

### Return type

[**TokenInfoResponse**](TokenInfoResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

