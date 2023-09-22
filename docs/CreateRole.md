# CreateRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The Role name. | 
**Description** | Pointer to **string** | Role description. | [optional] 
**PermissionObjects** | [**[]PermissionObject**](PermissionObject.md) | The list of permissions granted by this role. | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewCreateRole

`func NewCreateRole(name string, permissionObjects []PermissionObject, ) *CreateRole`

NewCreateRole instantiates a new CreateRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleWithDefaults

`func NewCreateRoleWithDefaults() *CreateRole`

NewCreateRoleWithDefaults instantiates a new CreateRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRole) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermissionObjects

`func (o *CreateRole) GetPermissionObjects() []PermissionObject`

GetPermissionObjects returns the PermissionObjects field if non-nil, zero value otherwise.

### GetPermissionObjectsOk

`func (o *CreateRole) GetPermissionObjectsOk() (*[]PermissionObject, bool)`

GetPermissionObjectsOk returns a tuple with the PermissionObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionObjects

`func (o *CreateRole) SetPermissionObjects(v []PermissionObject)`

SetPermissionObjects sets PermissionObjects field to given value.


### GetTags

`func (o *CreateRole) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateRole) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateRole) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateRole) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


