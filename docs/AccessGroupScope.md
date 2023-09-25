# AccessGroupScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Access group scope ID. | [optional] [readonly] 
**Name** | Pointer to **string** | The Access group scope name. | [optional] 
**RoleId** | **string** | The Access group role id. | 
**ScopeType** | Pointer to **string** | Specifies the type of the scope. Scope of type SIMPLE would grant access to all DCT objects. Scope of type SCOPED would grant access to all objects based on objects and object-tags and permissions defined in linked role. Scope of type ADVANCED would grant access to DCT objects based on objects and object-tags and the individual permissions. | [optional] 
**ObjectTags** | Pointer to [**[]ScopeTag**](ScopeTag.md) | The permissions in this access group scope will be granted to all DCT objects tagged with tags matching this property. This is cumulative with objects defined in the &#39;objects&#39; property, and mutually exclusive with scope_type &#39;SIMPLE&#39;. | [optional] 
**Objects** | Pointer to [**[]ScopedObjectItem**](ScopedObjectItem.md) | The permissions in this access group scope will be granted to all DCT objects with matching object id and object type, mutually exclusive with  scope_type &#39;SIMPLE&#39;. | [optional] 
**AlwaysAllowedPermissions** | Pointer to [**[]AlwaysAllowedPermission**](AlwaysAllowedPermission.md) | An array of always allowed permissions which can be used to specify object type and permission. An object with same object type and permission can not be added in &#39;objects&#39; array. | [optional] 

## Methods

### NewAccessGroupScope

`func NewAccessGroupScope(roleId string, ) *AccessGroupScope`

NewAccessGroupScope instantiates a new AccessGroupScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessGroupScopeWithDefaults

`func NewAccessGroupScopeWithDefaults() *AccessGroupScope`

NewAccessGroupScopeWithDefaults instantiates a new AccessGroupScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessGroupScope) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessGroupScope) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessGroupScope) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccessGroupScope) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccessGroupScope) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessGroupScope) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessGroupScope) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessGroupScope) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoleId

`func (o *AccessGroupScope) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *AccessGroupScope) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *AccessGroupScope) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetScopeType

`func (o *AccessGroupScope) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *AccessGroupScope) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *AccessGroupScope) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *AccessGroupScope) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### GetObjectTags

`func (o *AccessGroupScope) GetObjectTags() []ScopeTag`

GetObjectTags returns the ObjectTags field if non-nil, zero value otherwise.

### GetObjectTagsOk

`func (o *AccessGroupScope) GetObjectTagsOk() (*[]ScopeTag, bool)`

GetObjectTagsOk returns a tuple with the ObjectTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTags

`func (o *AccessGroupScope) SetObjectTags(v []ScopeTag)`

SetObjectTags sets ObjectTags field to given value.

### HasObjectTags

`func (o *AccessGroupScope) HasObjectTags() bool`

HasObjectTags returns a boolean if a field has been set.

### GetObjects

`func (o *AccessGroupScope) GetObjects() []ScopedObjectItem`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *AccessGroupScope) GetObjectsOk() (*[]ScopedObjectItem, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *AccessGroupScope) SetObjects(v []ScopedObjectItem)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *AccessGroupScope) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetAlwaysAllowedPermissions

`func (o *AccessGroupScope) GetAlwaysAllowedPermissions() []AlwaysAllowedPermission`

GetAlwaysAllowedPermissions returns the AlwaysAllowedPermissions field if non-nil, zero value otherwise.

### GetAlwaysAllowedPermissionsOk

`func (o *AccessGroupScope) GetAlwaysAllowedPermissionsOk() (*[]AlwaysAllowedPermission, bool)`

GetAlwaysAllowedPermissionsOk returns a tuple with the AlwaysAllowedPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysAllowedPermissions

`func (o *AccessGroupScope) SetAlwaysAllowedPermissions(v []AlwaysAllowedPermission)`

SetAlwaysAllowedPermissions sets AlwaysAllowedPermissions field to given value.

### HasAlwaysAllowedPermissions

`func (o *AccessGroupScope) HasAlwaysAllowedPermissions() bool`

HasAlwaysAllowedPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


