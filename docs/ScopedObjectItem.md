# ScopedObjectItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | **string** | ID of the object | 
**ObjectType** | [**ObjectTypeEnum**](ObjectTypeEnum.md) |  | 
**Permission** | Pointer to [**PermissionEnum**](PermissionEnum.md) |  | [optional] 

## Methods

### NewScopedObjectItem

`func NewScopedObjectItem(objectId string, objectType ObjectTypeEnum, ) *ScopedObjectItem`

NewScopedObjectItem instantiates a new ScopedObjectItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopedObjectItemWithDefaults

`func NewScopedObjectItemWithDefaults() *ScopedObjectItem`

NewScopedObjectItemWithDefaults instantiates a new ScopedObjectItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectId

`func (o *ScopedObjectItem) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ScopedObjectItem) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ScopedObjectItem) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetObjectType

`func (o *ScopedObjectItem) GetObjectType() ObjectTypeEnum`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ScopedObjectItem) GetObjectTypeOk() (*ObjectTypeEnum, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ScopedObjectItem) SetObjectType(v ObjectTypeEnum)`

SetObjectType sets ObjectType field to given value.


### GetPermission

`func (o *ScopedObjectItem) GetPermission() PermissionEnum`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ScopedObjectItem) GetPermissionOk() (*PermissionEnum, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ScopedObjectItem) SetPermission(v PermissionEnum)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *ScopedObjectItem) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


