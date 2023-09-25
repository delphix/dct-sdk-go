# AccountCreateParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAdmin** | Pointer to **bool** | Whether the created account must be granted to admin role. | [optional] [default to false]
**GenerateApiKey** | Pointer to **bool** | Whether an API key must be generated for this Account. This must be set if the Account will be used for API key based authentication, and unset otherwise. | [optional] [default to false]
**ApiClientId** | Pointer to **string** | The unique ID which is used to identify the identity of an API request. The web server (nginx) configuration must be configured so as to include the external ID as the value of the X_CLIENT_ID HTTP request header when requests are proxied. If this value isn&#39;t set, the application will automatically generate one. For OAuth2/JWT based authentication, this typically corresponds to a value extracted from the JWT, uniquely identifying the Account. | [optional] 
**FirstName** | Pointer to **string** | An optional first name for the Account. | [optional] 
**LastName** | Pointer to **string** | An optional last name for the Account. | [optional] 
**Email** | Pointer to **string** | An optional email for the Account. | [optional] 
**Username** | Pointer to **string** | The username for username/password authentication. This can also be used to provide an optional logical name for the Account. | [optional] 
**Password** | Pointer to **string** | The password for username/password authentication. | [optional] 
**LdapPrincipal** | Pointer to **string** | This value will be used for linking this account to an LDAP user when authenticated with the same LDAP principal. When accounts authenticate with LDAP, an LDAP principal value is calculated based on the username, msad_domain_name, search_base and username_pattern. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | The tags to be created for this Account. | [optional] 

## Methods

### NewAccountCreateParameter

`func NewAccountCreateParameter() *AccountCreateParameter`

NewAccountCreateParameter instantiates a new AccountCreateParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreateParameterWithDefaults

`func NewAccountCreateParameterWithDefaults() *AccountCreateParameter`

NewAccountCreateParameterWithDefaults instantiates a new AccountCreateParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAdmin

`func (o *AccountCreateParameter) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *AccountCreateParameter) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *AccountCreateParameter) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *AccountCreateParameter) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetGenerateApiKey

`func (o *AccountCreateParameter) GetGenerateApiKey() bool`

GetGenerateApiKey returns the GenerateApiKey field if non-nil, zero value otherwise.

### GetGenerateApiKeyOk

`func (o *AccountCreateParameter) GetGenerateApiKeyOk() (*bool, bool)`

GetGenerateApiKeyOk returns a tuple with the GenerateApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateApiKey

`func (o *AccountCreateParameter) SetGenerateApiKey(v bool)`

SetGenerateApiKey sets GenerateApiKey field to given value.

### HasGenerateApiKey

`func (o *AccountCreateParameter) HasGenerateApiKey() bool`

HasGenerateApiKey returns a boolean if a field has been set.

### GetApiClientId

`func (o *AccountCreateParameter) GetApiClientId() string`

GetApiClientId returns the ApiClientId field if non-nil, zero value otherwise.

### GetApiClientIdOk

`func (o *AccountCreateParameter) GetApiClientIdOk() (*string, bool)`

GetApiClientIdOk returns a tuple with the ApiClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientId

`func (o *AccountCreateParameter) SetApiClientId(v string)`

SetApiClientId sets ApiClientId field to given value.

### HasApiClientId

`func (o *AccountCreateParameter) HasApiClientId() bool`

HasApiClientId returns a boolean if a field has been set.

### GetFirstName

`func (o *AccountCreateParameter) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AccountCreateParameter) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AccountCreateParameter) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AccountCreateParameter) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *AccountCreateParameter) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AccountCreateParameter) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AccountCreateParameter) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AccountCreateParameter) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *AccountCreateParameter) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountCreateParameter) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountCreateParameter) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountCreateParameter) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUsername

`func (o *AccountCreateParameter) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AccountCreateParameter) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AccountCreateParameter) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AccountCreateParameter) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *AccountCreateParameter) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AccountCreateParameter) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AccountCreateParameter) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AccountCreateParameter) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetLdapPrincipal

`func (o *AccountCreateParameter) GetLdapPrincipal() string`

GetLdapPrincipal returns the LdapPrincipal field if non-nil, zero value otherwise.

### GetLdapPrincipalOk

`func (o *AccountCreateParameter) GetLdapPrincipalOk() (*string, bool)`

GetLdapPrincipalOk returns a tuple with the LdapPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPrincipal

`func (o *AccountCreateParameter) SetLdapPrincipal(v string)`

SetLdapPrincipal sets LdapPrincipal field to given value.

### HasLdapPrincipal

`func (o *AccountCreateParameter) HasLdapPrincipal() bool`

HasLdapPrincipal returns a boolean if a field has been set.

### GetTags

`func (o *AccountCreateParameter) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AccountCreateParameter) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AccountCreateParameter) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AccountCreateParameter) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


