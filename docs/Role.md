# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The Role name. | 
**Description** | Pointer to **string** | Role description. | [optional] 
**PermissionObjects** | [**[]PermissionObject**](PermissionObject.md) | The list of permissions granted by this role. | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Id** | Pointer to **string** | The Role ID. | [optional] 
**SystemRole** | Pointer to **bool** | System role are pre defined roles. System roles cannot be modified. | [optional] 

## Methods

### NewRole

`func NewRole(name string, permissionObjects []PermissionObject, ) *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Role) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Role) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Role) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Role) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Role) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Role) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Role) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermissionObjects

`func (o *Role) GetPermissionObjects() []PermissionObject`

GetPermissionObjects returns the PermissionObjects field if non-nil, zero value otherwise.

### GetPermissionObjectsOk

`func (o *Role) GetPermissionObjectsOk() (*[]PermissionObject, bool)`

GetPermissionObjectsOk returns a tuple with the PermissionObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionObjects

`func (o *Role) SetPermissionObjects(v []PermissionObject)`

SetPermissionObjects sets PermissionObjects field to given value.


### GetTags

`func (o *Role) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Role) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Role) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Role) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetId

`func (o *Role) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Role) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Role) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Role) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSystemRole

`func (o *Role) GetSystemRole() bool`

GetSystemRole returns the SystemRole field if non-nil, zero value otherwise.

### GetSystemRoleOk

`func (o *Role) GetSystemRoleOk() (*bool, bool)`

GetSystemRoleOk returns a tuple with the SystemRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRole

`func (o *Role) SetSystemRole(v bool)`

SetSystemRole sets SystemRole field to given value.

### HasSystemRole

`func (o *Role) HasSystemRole() bool`

HasSystemRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


