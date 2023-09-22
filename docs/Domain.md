# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsadDomainName** | Pointer to **string** | This is used to get full DN for authentication and search. Provide this value only if server is microsoft AD. | [optional] 
**UsernamePattern** | Pointer to **string** | The username_patterns can be used to avoid providing full-dn during login. This will also be used for search of groups,email, first name and last name. | [optional] 
**SearchBase** | Pointer to **string** | Search base used to search for ldap user groups. Leave this field empty if a full username_pattern is provided and server is non microsoft AD. | [optional] 
**GroupAttr** | Pointer to **string** | Group mapped attribute on ldap side used for user group search. | [optional] 
**EmailAttr** | Pointer to **string** | Email mapped attribute on ldap side used for mapping on DCT side account. | [optional] 
**FirstNameAttr** | Pointer to **string** | First name attribute mapped on ldap side used for mapping on DCT side account. | [optional] 
**LastNameAttr** | Pointer to **string** | Last name attribute mapped on ldap side used for mapping on DCT side account. | [optional] 
**ObjectClassAttr** | Pointer to **string** | The name of the objectClass on ldap side under which the user is mapped.This is used to search for the user details. | [optional] 
**SearchAttr** | Pointer to **string** | Search attribute mapped on ldap side using which search on ldap side will be made. | [optional] 

## Methods

### NewDomain

`func NewDomain() *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsadDomainName

`func (o *Domain) GetMsadDomainName() string`

GetMsadDomainName returns the MsadDomainName field if non-nil, zero value otherwise.

### GetMsadDomainNameOk

`func (o *Domain) GetMsadDomainNameOk() (*string, bool)`

GetMsadDomainNameOk returns a tuple with the MsadDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsadDomainName

`func (o *Domain) SetMsadDomainName(v string)`

SetMsadDomainName sets MsadDomainName field to given value.

### HasMsadDomainName

`func (o *Domain) HasMsadDomainName() bool`

HasMsadDomainName returns a boolean if a field has been set.

### GetUsernamePattern

`func (o *Domain) GetUsernamePattern() string`

GetUsernamePattern returns the UsernamePattern field if non-nil, zero value otherwise.

### GetUsernamePatternOk

`func (o *Domain) GetUsernamePatternOk() (*string, bool)`

GetUsernamePatternOk returns a tuple with the UsernamePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernamePattern

`func (o *Domain) SetUsernamePattern(v string)`

SetUsernamePattern sets UsernamePattern field to given value.

### HasUsernamePattern

`func (o *Domain) HasUsernamePattern() bool`

HasUsernamePattern returns a boolean if a field has been set.

### GetSearchBase

`func (o *Domain) GetSearchBase() string`

GetSearchBase returns the SearchBase field if non-nil, zero value otherwise.

### GetSearchBaseOk

`func (o *Domain) GetSearchBaseOk() (*string, bool)`

GetSearchBaseOk returns a tuple with the SearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBase

`func (o *Domain) SetSearchBase(v string)`

SetSearchBase sets SearchBase field to given value.

### HasSearchBase

`func (o *Domain) HasSearchBase() bool`

HasSearchBase returns a boolean if a field has been set.

### GetGroupAttr

`func (o *Domain) GetGroupAttr() string`

GetGroupAttr returns the GroupAttr field if non-nil, zero value otherwise.

### GetGroupAttrOk

`func (o *Domain) GetGroupAttrOk() (*string, bool)`

GetGroupAttrOk returns a tuple with the GroupAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttr

`func (o *Domain) SetGroupAttr(v string)`

SetGroupAttr sets GroupAttr field to given value.

### HasGroupAttr

`func (o *Domain) HasGroupAttr() bool`

HasGroupAttr returns a boolean if a field has been set.

### GetEmailAttr

`func (o *Domain) GetEmailAttr() string`

GetEmailAttr returns the EmailAttr field if non-nil, zero value otherwise.

### GetEmailAttrOk

`func (o *Domain) GetEmailAttrOk() (*string, bool)`

GetEmailAttrOk returns a tuple with the EmailAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAttr

`func (o *Domain) SetEmailAttr(v string)`

SetEmailAttr sets EmailAttr field to given value.

### HasEmailAttr

`func (o *Domain) HasEmailAttr() bool`

HasEmailAttr returns a boolean if a field has been set.

### GetFirstNameAttr

`func (o *Domain) GetFirstNameAttr() string`

GetFirstNameAttr returns the FirstNameAttr field if non-nil, zero value otherwise.

### GetFirstNameAttrOk

`func (o *Domain) GetFirstNameAttrOk() (*string, bool)`

GetFirstNameAttrOk returns a tuple with the FirstNameAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNameAttr

`func (o *Domain) SetFirstNameAttr(v string)`

SetFirstNameAttr sets FirstNameAttr field to given value.

### HasFirstNameAttr

`func (o *Domain) HasFirstNameAttr() bool`

HasFirstNameAttr returns a boolean if a field has been set.

### GetLastNameAttr

`func (o *Domain) GetLastNameAttr() string`

GetLastNameAttr returns the LastNameAttr field if non-nil, zero value otherwise.

### GetLastNameAttrOk

`func (o *Domain) GetLastNameAttrOk() (*string, bool)`

GetLastNameAttrOk returns a tuple with the LastNameAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNameAttr

`func (o *Domain) SetLastNameAttr(v string)`

SetLastNameAttr sets LastNameAttr field to given value.

### HasLastNameAttr

`func (o *Domain) HasLastNameAttr() bool`

HasLastNameAttr returns a boolean if a field has been set.

### GetObjectClassAttr

`func (o *Domain) GetObjectClassAttr() string`

GetObjectClassAttr returns the ObjectClassAttr field if non-nil, zero value otherwise.

### GetObjectClassAttrOk

`func (o *Domain) GetObjectClassAttrOk() (*string, bool)`

GetObjectClassAttrOk returns a tuple with the ObjectClassAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClassAttr

`func (o *Domain) SetObjectClassAttr(v string)`

SetObjectClassAttr sets ObjectClassAttr field to given value.

### HasObjectClassAttr

`func (o *Domain) HasObjectClassAttr() bool`

HasObjectClassAttr returns a boolean if a field has been set.

### GetSearchAttr

`func (o *Domain) GetSearchAttr() string`

GetSearchAttr returns the SearchAttr field if non-nil, zero value otherwise.

### GetSearchAttrOk

`func (o *Domain) GetSearchAttrOk() (*string, bool)`

GetSearchAttrOk returns a tuple with the SearchAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAttr

`func (o *Domain) SetSearchAttr(v string)`

SetSearchAttr sets SearchAttr field to given value.

### HasSearchAttr

`func (o *Domain) HasSearchAttr() bool`

HasSearchAttr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


