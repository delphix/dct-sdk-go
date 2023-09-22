# ObjectPermissionAccessGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the access group. | [optional] 
**Name** | Pointer to **string** | Name of the access group. | [optional] 
**Permissions** | Pointer to **[]string** | Permissions for the object in this access group. | [optional] 

## Methods

### NewObjectPermissionAccessGroups

`func NewObjectPermissionAccessGroups() *ObjectPermissionAccessGroups`

NewObjectPermissionAccessGroups instantiates a new ObjectPermissionAccessGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectPermissionAccessGroupsWithDefaults

`func NewObjectPermissionAccessGroupsWithDefaults() *ObjectPermissionAccessGroups`

NewObjectPermissionAccessGroupsWithDefaults instantiates a new ObjectPermissionAccessGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectPermissionAccessGroups) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectPermissionAccessGroups) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectPermissionAccessGroups) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectPermissionAccessGroups) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ObjectPermissionAccessGroups) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectPermissionAccessGroups) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectPermissionAccessGroups) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectPermissionAccessGroups) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *ObjectPermissionAccessGroups) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ObjectPermissionAccessGroups) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ObjectPermissionAccessGroups) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ObjectPermissionAccessGroups) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


