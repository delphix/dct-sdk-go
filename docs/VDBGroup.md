# VDBGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the entity. | [readonly] 
**Name** | **string** | A unique name for the entity. | 
**VdbIds** | **[]string** | The list of VDB IDs in this VDBGroup. | 
**IsLocked** | Pointer to **bool** | Indicates whether the VDBGroup is locked. | [optional] 
**LockedBy** | Pointer to **int64** | The Id of the account that locked the VDBGroup. | [optional] 
**LockedByName** | Pointer to **string** | The name of the account that locked the VDBGroup. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewVDBGroup

`func NewVDBGroup(id string, name string, vdbIds []string, ) *VDBGroup`

NewVDBGroup instantiates a new VDBGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVDBGroupWithDefaults

`func NewVDBGroupWithDefaults() *VDBGroup`

NewVDBGroupWithDefaults instantiates a new VDBGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VDBGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VDBGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VDBGroup) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VDBGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VDBGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VDBGroup) SetName(v string)`

SetName sets Name field to given value.


### GetVdbIds

`func (o *VDBGroup) GetVdbIds() []string`

GetVdbIds returns the VdbIds field if non-nil, zero value otherwise.

### GetVdbIdsOk

`func (o *VDBGroup) GetVdbIdsOk() (*[]string, bool)`

GetVdbIdsOk returns a tuple with the VdbIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdbIds

`func (o *VDBGroup) SetVdbIds(v []string)`

SetVdbIds sets VdbIds field to given value.


### GetIsLocked

`func (o *VDBGroup) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *VDBGroup) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *VDBGroup) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *VDBGroup) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetLockedBy

`func (o *VDBGroup) GetLockedBy() int64`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *VDBGroup) GetLockedByOk() (*int64, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *VDBGroup) SetLockedBy(v int64)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *VDBGroup) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.

### GetLockedByName

`func (o *VDBGroup) GetLockedByName() string`

GetLockedByName returns the LockedByName field if non-nil, zero value otherwise.

### GetLockedByNameOk

`func (o *VDBGroup) GetLockedByNameOk() (*string, bool)`

GetLockedByNameOk returns a tuple with the LockedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedByName

`func (o *VDBGroup) SetLockedByName(v string)`

SetLockedByName sets LockedByName field to given value.

### HasLockedByName

`func (o *VDBGroup) HasLockedByName() bool`

HasLockedByName returns a boolean if a field has been set.

### GetTags

`func (o *VDBGroup) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VDBGroup) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VDBGroup) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VDBGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


