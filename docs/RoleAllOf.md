# RoleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Role ID. | [optional] 
**SystemRole** | Pointer to **bool** | System role are pre defined roles. System roles cannot be modified. | [optional] 

## Methods

### NewRoleAllOf

`func NewRoleAllOf() *RoleAllOf`

NewRoleAllOf instantiates a new RoleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAllOfWithDefaults

`func NewRoleAllOfWithDefaults() *RoleAllOf`

NewRoleAllOfWithDefaults instantiates a new RoleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoleAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoleAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSystemRole

`func (o *RoleAllOf) GetSystemRole() bool`

GetSystemRole returns the SystemRole field if non-nil, zero value otherwise.

### GetSystemRoleOk

`func (o *RoleAllOf) GetSystemRoleOk() (*bool, bool)`

GetSystemRoleOk returns a tuple with the SystemRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRole

`func (o *RoleAllOf) SetSystemRole(v bool)`

SetSystemRole sets SystemRole field to given value.

### HasSystemRole

`func (o *RoleAllOf) HasSystemRole() bool`

HasSystemRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


