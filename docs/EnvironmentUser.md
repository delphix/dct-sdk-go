# EnvironmentUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserRef** | Pointer to **string** | Environment user reference | [optional] 
**Username** | Pointer to **string** | Username of environment user | [optional] 
**PrimaryUser** | Pointer to **bool** | This indicates if this user is primary or not | [optional] 

## Methods

### NewEnvironmentUser

`func NewEnvironmentUser() *EnvironmentUser`

NewEnvironmentUser instantiates a new EnvironmentUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentUserWithDefaults

`func NewEnvironmentUserWithDefaults() *EnvironmentUser`

NewEnvironmentUserWithDefaults instantiates a new EnvironmentUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserRef

`func (o *EnvironmentUser) GetUserRef() string`

GetUserRef returns the UserRef field if non-nil, zero value otherwise.

### GetUserRefOk

`func (o *EnvironmentUser) GetUserRefOk() (*string, bool)`

GetUserRefOk returns a tuple with the UserRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRef

`func (o *EnvironmentUser) SetUserRef(v string)`

SetUserRef sets UserRef field to given value.

### HasUserRef

`func (o *EnvironmentUser) HasUserRef() bool`

HasUserRef returns a boolean if a field has been set.

### GetUsername

`func (o *EnvironmentUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EnvironmentUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EnvironmentUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EnvironmentUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPrimaryUser

`func (o *EnvironmentUser) GetPrimaryUser() bool`

GetPrimaryUser returns the PrimaryUser field if non-nil, zero value otherwise.

### GetPrimaryUserOk

`func (o *EnvironmentUser) GetPrimaryUserOk() (*bool, bool)`

GetPrimaryUserOk returns a tuple with the PrimaryUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryUser

`func (o *EnvironmentUser) SetPrimaryUser(v bool)`

SetPrimaryUser sets PrimaryUser field to given value.

### HasPrimaryUser

`func (o *EnvironmentUser) HasPrimaryUser() bool`

HasPrimaryUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


