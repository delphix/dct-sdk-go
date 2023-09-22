# CreateVDBGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**VdbIds** | **[]string** |  | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | The tags to be created for VDB Group. | [optional] 
**MakeCurrentAccountOwner** | Pointer to **bool** | Whether the account creating this VDB group must be configured as owner of the VDB group. | [optional] [default to true]

## Methods

### NewCreateVDBGroupRequest

`func NewCreateVDBGroupRequest(name string, vdbIds []string, ) *CreateVDBGroupRequest`

NewCreateVDBGroupRequest instantiates a new CreateVDBGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVDBGroupRequestWithDefaults

`func NewCreateVDBGroupRequestWithDefaults() *CreateVDBGroupRequest`

NewCreateVDBGroupRequestWithDefaults instantiates a new CreateVDBGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateVDBGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVDBGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVDBGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVdbIds

`func (o *CreateVDBGroupRequest) GetVdbIds() []string`

GetVdbIds returns the VdbIds field if non-nil, zero value otherwise.

### GetVdbIdsOk

`func (o *CreateVDBGroupRequest) GetVdbIdsOk() (*[]string, bool)`

GetVdbIdsOk returns a tuple with the VdbIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdbIds

`func (o *CreateVDBGroupRequest) SetVdbIds(v []string)`

SetVdbIds sets VdbIds field to given value.


### GetTags

`func (o *CreateVDBGroupRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateVDBGroupRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateVDBGroupRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateVDBGroupRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMakeCurrentAccountOwner

`func (o *CreateVDBGroupRequest) GetMakeCurrentAccountOwner() bool`

GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field if non-nil, zero value otherwise.

### GetMakeCurrentAccountOwnerOk

`func (o *CreateVDBGroupRequest) GetMakeCurrentAccountOwnerOk() (*bool, bool)`

GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCurrentAccountOwner

`func (o *CreateVDBGroupRequest) SetMakeCurrentAccountOwner(v bool)`

SetMakeCurrentAccountOwner sets MakeCurrentAccountOwner field to given value.

### HasMakeCurrentAccountOwner

`func (o *CreateVDBGroupRequest) HasMakeCurrentAccountOwner() bool`

HasMakeCurrentAccountOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


