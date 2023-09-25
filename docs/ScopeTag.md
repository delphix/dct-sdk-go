# ScopeTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Key of the tag | 
**Value** | **string** | Value of the tag | 
**ObjectType** | Pointer to [**ObjectTypeEnum**](ObjectTypeEnum.md) |  | [optional] 
**Permission** | Pointer to [**PermissionEnum**](PermissionEnum.md) |  | [optional] 

## Methods

### NewScopeTag

`func NewScopeTag(key string, value string, ) *ScopeTag`

NewScopeTag instantiates a new ScopeTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeTagWithDefaults

`func NewScopeTagWithDefaults() *ScopeTag`

NewScopeTagWithDefaults instantiates a new ScopeTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ScopeTag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ScopeTag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ScopeTag) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *ScopeTag) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ScopeTag) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ScopeTag) SetValue(v string)`

SetValue sets Value field to given value.


### GetObjectType

`func (o *ScopeTag) GetObjectType() ObjectTypeEnum`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ScopeTag) GetObjectTypeOk() (*ObjectTypeEnum, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ScopeTag) SetObjectType(v ObjectTypeEnum)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ScopeTag) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPermission

`func (o *ScopeTag) GetPermission() PermissionEnum`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ScopeTag) GetPermissionOk() (*PermissionEnum, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ScopeTag) SetPermission(v PermissionEnum)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *ScopeTag) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


