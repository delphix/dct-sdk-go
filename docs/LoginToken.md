# LoginToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Opaque token that validates the successful account login and is used to authenticate subsequent api calls. This token needs to be sent as part of &#39;Authorization&#39; header for all api calls prefixed with value contained in &#39;token_type&#39; property. For example, if the &#39;access_token&#39; value is \&quot;abc123\&quot; and &#39;token_type&#39; is \&quot;Bearer\&quot; then HTTP requests should contain following header : \&quot;Authorization: Bearer abc123\&quot;  | 
**TokenType** | **string** | Type of the token returned in &#39;access_token&#39; property. | 
**ExpiresIn** | **int64** | Seconds duration after which the token will expire. | 

## Methods

### NewLoginToken

`func NewLoginToken(accessToken string, tokenType string, expiresIn int64, ) *LoginToken`

NewLoginToken instantiates a new LoginToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginTokenWithDefaults

`func NewLoginTokenWithDefaults() *LoginToken`

NewLoginTokenWithDefaults instantiates a new LoginToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *LoginToken) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *LoginToken) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *LoginToken) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetTokenType

`func (o *LoginToken) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *LoginToken) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *LoginToken) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetExpiresIn

`func (o *LoginToken) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *LoginToken) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *LoginToken) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


