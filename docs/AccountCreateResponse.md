# AccountCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | The opaque token to use to authenticate for other API calls. The token must be included in all HTTP API calls in a request header named \&quot;Authorization\&quot;, and prefixed with \&quot;apk \&quot; to denote that it is a proprietary API key format. For instance, if the token is \&quot;abc123\&quot;, request must contain the following HTTP request header: \&quot;Authorization: apk abc123\&quot;.  | [optional] 
**Id** | Pointer to **int64** | Numeric ID of the created Account. | [optional] 
**FirstName** | Pointer to **string** | First name of the created Account. | [optional] 
**LastName** | Pointer to **string** | Last name of the created Account. | [optional] 
**Email** | Pointer to **string** | Email of the created Account. | [optional] 
**Username** | Pointer to **string** | Username of the created Account. | [optional] 
**LdapPrincipal** | Pointer to **string** | The LDAP principal of the created Account. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | The tags to be created for this Account. | [optional] 

## Methods

### NewAccountCreateResponse

`func NewAccountCreateResponse() *AccountCreateResponse`

NewAccountCreateResponse instantiates a new AccountCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreateResponseWithDefaults

`func NewAccountCreateResponseWithDefaults() *AccountCreateResponse`

NewAccountCreateResponseWithDefaults instantiates a new AccountCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *AccountCreateResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccountCreateResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccountCreateResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AccountCreateResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetId

`func (o *AccountCreateResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountCreateResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountCreateResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AccountCreateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *AccountCreateResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AccountCreateResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AccountCreateResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AccountCreateResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *AccountCreateResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AccountCreateResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AccountCreateResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AccountCreateResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *AccountCreateResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountCreateResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountCreateResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountCreateResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUsername

`func (o *AccountCreateResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AccountCreateResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AccountCreateResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AccountCreateResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetLdapPrincipal

`func (o *AccountCreateResponse) GetLdapPrincipal() string`

GetLdapPrincipal returns the LdapPrincipal field if non-nil, zero value otherwise.

### GetLdapPrincipalOk

`func (o *AccountCreateResponse) GetLdapPrincipalOk() (*string, bool)`

GetLdapPrincipalOk returns a tuple with the LdapPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPrincipal

`func (o *AccountCreateResponse) SetLdapPrincipal(v string)`

SetLdapPrincipal sets LdapPrincipal field to given value.

### HasLdapPrincipal

`func (o *AccountCreateResponse) HasLdapPrincipal() bool`

HasLdapPrincipal returns a boolean if a field has been set.

### GetTags

`func (o *AccountCreateResponse) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AccountCreateResponse) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AccountCreateResponse) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AccountCreateResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


