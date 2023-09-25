# PermissionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionObjects** | [**[]PermissionObject**](PermissionObject.md) | Array of permissions with object type and their permission. | 

## Methods

### NewPermissionsRequest

`func NewPermissionsRequest(permissionObjects []PermissionObject, ) *PermissionsRequest`

NewPermissionsRequest instantiates a new PermissionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsRequestWithDefaults

`func NewPermissionsRequestWithDefaults() *PermissionsRequest`

NewPermissionsRequestWithDefaults instantiates a new PermissionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionObjects

`func (o *PermissionsRequest) GetPermissionObjects() []PermissionObject`

GetPermissionObjects returns the PermissionObjects field if non-nil, zero value otherwise.

### GetPermissionObjectsOk

`func (o *PermissionsRequest) GetPermissionObjectsOk() (*[]PermissionObject, bool)`

GetPermissionObjectsOk returns a tuple with the PermissionObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionObjects

`func (o *PermissionsRequest) SetPermissionObjects(v []PermissionObject)`

SetPermissionObjects sets PermissionObjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


