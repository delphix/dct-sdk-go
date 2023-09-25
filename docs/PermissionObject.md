# PermissionObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | Object type. | 
**Permissions** | **[]string** | List of permissions. | 

## Methods

### NewPermissionObject

`func NewPermissionObject(objectType string, permissions []string, ) *PermissionObject`

NewPermissionObject instantiates a new PermissionObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionObjectWithDefaults

`func NewPermissionObjectWithDefaults() *PermissionObject`

NewPermissionObjectWithDefaults instantiates a new PermissionObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *PermissionObject) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PermissionObject) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PermissionObject) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPermissions

`func (o *PermissionObject) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PermissionObject) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PermissionObject) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


