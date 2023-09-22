# AccountUpdateParameter

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

## Methods

### NewAccountUpdateParameter

`func NewAccountUpdateParameter() *AccountUpdateParameter`

NewAccountUpdateParameter instantiates a new AccountUpdateParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUpdateParameterWithDefaults

`func NewAccountUpdateParameterWithDefaults() *AccountUpdateParameter`

NewAccountUpdateParameterWithDefaults instantiates a new AccountUpdateParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountUpdateParameter) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountUpdateParameter) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountUpdateParameter) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AccountUpdateParameter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetApiClientId

`func (o *AccountUpdateParameter) GetApiClientId() string`

GetApiClientId returns the ApiClientId field if non-nil, zero value otherwise.

### GetApiClientIdOk

`func (o *AccountUpdateParameter) GetApiClientIdOk() (*string, bool)`

GetApiClientIdOk returns a tuple with the ApiClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientId

`func (o *AccountUpdateParameter) SetApiClientId(v string)`

SetApiClientId sets ApiClientId field to given value.

### HasApiClientId

`func (o *AccountUpdateParameter) HasApiClientId() bool`

HasApiClientId returns a boolean if a field has been set.

### GetFirstName

`func (o *AccountUpdateParameter) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AccountUpdateParameter) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AccountUpdateParameter) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AccountUpdateParameter) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *AccountUpdateParameter) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AccountUpdateParameter) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AccountUpdateParameter) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AccountUpdateParameter) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *AccountUpdateParameter) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountUpdateParameter) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountUpdateParameter) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountUpdateParameter) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUsername

`func (o *AccountUpdateParameter) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AccountUpdateParameter) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AccountUpdateParameter) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AccountUpdateParameter) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetLdapPrincipal

`func (o *AccountUpdateParameter) GetLdapPrincipal() string`

GetLdapPrincipal returns the LdapPrincipal field if non-nil, zero value otherwise.

### GetLdapPrincipalOk

`func (o *AccountUpdateParameter) GetLdapPrincipalOk() (*string, bool)`

GetLdapPrincipalOk returns a tuple with the LdapPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPrincipal

`func (o *AccountUpdateParameter) SetLdapPrincipal(v string)`

SetLdapPrincipal sets LdapPrincipal field to given value.

### HasLdapPrincipal

`func (o *AccountUpdateParameter) HasLdapPrincipal() bool`

HasLdapPrincipal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


