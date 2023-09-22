# TokenInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Flag to identify if the token is active. | [optional] 
**TokenType** | Pointer to **string** | Type of the token. | [optional] 
**AccountId** | Pointer to **int64** | Numeric ID of the account. | [optional] 
**FirstName** | Pointer to **string** | First name for the Account. | [optional] 
**LastName** | Pointer to **string** | Last name for the Account. | [optional] 
**Email** | Pointer to **string** | Email for the Account. | [optional] 
**Username** | Pointer to **string** | The username or logical name for the Account. | [optional] 
**LdapPrincipal** | Pointer to **string** | The LDAP Principal for the Account. | [optional] 
**Exp** | Pointer to **int64** | Seconds duration after which the token will expire. | [optional] 

## Methods

### NewTokenInfoResponse

`func NewTokenInfoResponse() *TokenInfoResponse`

NewTokenInfoResponse instantiates a new TokenInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenInfoResponseWithDefaults

`func NewTokenInfoResponseWithDefaults() *TokenInfoResponse`

NewTokenInfoResponseWithDefaults instantiates a new TokenInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *TokenInfoResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TokenInfoResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TokenInfoResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *TokenInfoResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetTokenType

`func (o *TokenInfoResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TokenInfoResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TokenInfoResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *TokenInfoResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetAccountId

`func (o *TokenInfoResponse) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TokenInfoResponse) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TokenInfoResponse) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TokenInfoResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetFirstName

`func (o *TokenInfoResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *TokenInfoResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *TokenInfoResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *TokenInfoResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *TokenInfoResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *TokenInfoResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *TokenInfoResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *TokenInfoResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *TokenInfoResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TokenInfoResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TokenInfoResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TokenInfoResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUsername

`func (o *TokenInfoResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TokenInfoResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TokenInfoResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TokenInfoResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetLdapPrincipal

`func (o *TokenInfoResponse) GetLdapPrincipal() string`

GetLdapPrincipal returns the LdapPrincipal field if non-nil, zero value otherwise.

### GetLdapPrincipalOk

`func (o *TokenInfoResponse) GetLdapPrincipalOk() (*string, bool)`

GetLdapPrincipalOk returns a tuple with the LdapPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPrincipal

`func (o *TokenInfoResponse) SetLdapPrincipal(v string)`

SetLdapPrincipal sets LdapPrincipal field to given value.

### HasLdapPrincipal

`func (o *TokenInfoResponse) HasLdapPrincipal() bool`

HasLdapPrincipal returns a boolean if a field has been set.

### GetExp

`func (o *TokenInfoResponse) GetExp() int64`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *TokenInfoResponse) GetExpOk() (*int64, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *TokenInfoResponse) SetExp(v int64)`

SetExp sets Exp field to given value.

### HasExp

`func (o *TokenInfoResponse) HasExp() bool`

HasExp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


