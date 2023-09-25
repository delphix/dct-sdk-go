# ObjectPermissionAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Numeric ID of the Account. | [optional] 
**FirstName** | Pointer to **string** | First name of the Account. | [optional] 
**LastName** | Pointer to **string** | Last name of the Account. | [optional] 
**Email** | Pointer to **string** | Email of the Account. | [optional] 
**AccessGroups** | Pointer to [**[]ObjectPermissionAccessGroups**](ObjectPermissionAccessGroups.md) | Access groups of the Account. | [optional] 

## Methods

### NewObjectPermissionAccount

`func NewObjectPermissionAccount() *ObjectPermissionAccount`

NewObjectPermissionAccount instantiates a new ObjectPermissionAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectPermissionAccountWithDefaults

`func NewObjectPermissionAccountWithDefaults() *ObjectPermissionAccount`

NewObjectPermissionAccountWithDefaults instantiates a new ObjectPermissionAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectPermissionAccount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectPermissionAccount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectPermissionAccount) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectPermissionAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *ObjectPermissionAccount) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ObjectPermissionAccount) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ObjectPermissionAccount) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ObjectPermissionAccount) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ObjectPermissionAccount) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ObjectPermissionAccount) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ObjectPermissionAccount) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ObjectPermissionAccount) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *ObjectPermissionAccount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ObjectPermissionAccount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ObjectPermissionAccount) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ObjectPermissionAccount) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAccessGroups

`func (o *ObjectPermissionAccount) GetAccessGroups() []ObjectPermissionAccessGroups`

GetAccessGroups returns the AccessGroups field if non-nil, zero value otherwise.

### GetAccessGroupsOk

`func (o *ObjectPermissionAccount) GetAccessGroupsOk() (*[]ObjectPermissionAccessGroups, bool)`

GetAccessGroupsOk returns a tuple with the AccessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroups

`func (o *ObjectPermissionAccount) SetAccessGroups(v []ObjectPermissionAccessGroups)`

SetAccessGroups sets AccessGroups field to given value.

### HasAccessGroups

`func (o *ObjectPermissionAccount) HasAccessGroups() bool`

HasAccessGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


