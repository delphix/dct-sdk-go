# UpdateAccessGroupScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The Access group scope name. | [optional] 
**ScopeType** | Pointer to **string** | Specifies the type of the scope. Scope of type SIMPLE would grant access to all DCT objects. Scope of type SCOPED would grant access to all objects based on objects and object-tags and permissions defined in linked role. Scope of type ADVANCED would grant access to DCT objects based on objects and object-tags and the individual permissions. | [optional] 

## Methods

### NewUpdateAccessGroupScope

`func NewUpdateAccessGroupScope() *UpdateAccessGroupScope`

NewUpdateAccessGroupScope instantiates a new UpdateAccessGroupScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccessGroupScopeWithDefaults

`func NewUpdateAccessGroupScopeWithDefaults() *UpdateAccessGroupScope`

NewUpdateAccessGroupScopeWithDefaults instantiates a new UpdateAccessGroupScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAccessGroupScope) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAccessGroupScope) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAccessGroupScope) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAccessGroupScope) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScopeType

`func (o *UpdateAccessGroupScope) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *UpdateAccessGroupScope) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *UpdateAccessGroupScope) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *UpdateAccessGroupScope) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


