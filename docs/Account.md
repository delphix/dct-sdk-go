# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Numeric ID of the Account. | [optional] [readonly] 
**ApiClientId** | Pointer to **string** | The unique ID which is used to identify the identity of an API request. The web server (nginx) configuration must be configured so as to include the external ID as the value of the X_CLIENT_ID HTTP request header when requests are proxied. For OAuth2/JWT based authentication, this typically corresponds to a value extracted from the JWT, uniquely identifying the Account. | [optional] 
**FirstName** | Pointer to **string** | An optional first name for the Account. | [optional] 
**LastName** | Pointer to **string** | An optional last name for the Account. | [optional] 
**Email** | Pointer to **string** | An optional email for the Account. | [optional] 
**Username** | Pointer to **string** | The username for username/password authentication. This can also be used to provide an optional logical name for the Account. | [optional] 
**LdapPrincipal** | Pointer to **string** | This value will be used for linking this account to an LDAP user when authenticated with the same LDAP principal. When accounts authenticate with LDAP, an LDAP principal value is calculated based on the username, msad_domain_name, search_base and username_pattern. | [optional] 
**EffectiveScopes** | Pointer to [**[]EffectiveScope**](EffectiveScope.md) | Access group scopes associated with this account. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | The tags to be created for this Account. | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Account) HasId() bool`

HasId returns a boolean if a field has been set.

### GetApiClientId

`func (o *Account) GetApiClientId() string`

GetApiClientId returns the ApiClientId field if non-nil, zero value otherwise.

### GetApiClientIdOk

`func (o *Account) GetApiClientIdOk() (*string, bool)`

GetApiClientIdOk returns a tuple with the ApiClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientId

`func (o *Account) SetApiClientId(v string)`

SetApiClientId sets ApiClientId field to given value.

### HasApiClientId

`func (o *Account) HasApiClientId() bool`

HasApiClientId returns a boolean if a field has been set.

### GetFirstName

`func (o *Account) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Account) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Account) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Account) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Account) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Account) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Account) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Account) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *Account) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Account) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Account) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Account) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUsername

`func (o *Account) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Account) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Account) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Account) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetLdapPrincipal

`func (o *Account) GetLdapPrincipal() string`

GetLdapPrincipal returns the LdapPrincipal field if non-nil, zero value otherwise.

### GetLdapPrincipalOk

`func (o *Account) GetLdapPrincipalOk() (*string, bool)`

GetLdapPrincipalOk returns a tuple with the LdapPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPrincipal

`func (o *Account) SetLdapPrincipal(v string)`

SetLdapPrincipal sets LdapPrincipal field to given value.

### HasLdapPrincipal

`func (o *Account) HasLdapPrincipal() bool`

HasLdapPrincipal returns a boolean if a field has been set.

### GetEffectiveScopes

`func (o *Account) GetEffectiveScopes() []EffectiveScope`

GetEffectiveScopes returns the EffectiveScopes field if non-nil, zero value otherwise.

### GetEffectiveScopesOk

`func (o *Account) GetEffectiveScopesOk() (*[]EffectiveScope, bool)`

GetEffectiveScopesOk returns a tuple with the EffectiveScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveScopes

`func (o *Account) SetEffectiveScopes(v []EffectiveScope)`

SetEffectiveScopes sets EffectiveScopes field to given value.

### HasEffectiveScopes

`func (o *Account) HasEffectiveScopes() bool`

HasEffectiveScopes returns a boolean if a field has been set.

### GetTags

`func (o *Account) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Account) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Account) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Account) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


