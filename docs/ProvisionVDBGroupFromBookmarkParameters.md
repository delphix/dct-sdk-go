# ProvisionVDBGroupFromBookmarkParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the created VDB group name. | 
**BookmarkId** | **string** | ID of a bookmark to provision this VDB Group from. | 
**ProvisionParameters** | [**map[string]BaseProvisionVDBParameters**](BaseProvisionVDBParameters.md) | Provision parameters for each of the VDBs which will need to be provisioned. The key must be the vdb_id of the corresponding entry from the bookmark, and the value the provision parameters for the VDB which will be cloned from the bookmark. | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | The tags to be created for VDB Group. | [optional] 
**MakeCurrentAccountOwner** | Pointer to **bool** | Whether the account provisioning this VDB group must be configured as owner of the VDB group. | [optional] [default to true]

## Methods

### NewProvisionVDBGroupFromBookmarkParameters

`func NewProvisionVDBGroupFromBookmarkParameters(name string, bookmarkId string, provisionParameters map[string]BaseProvisionVDBParameters, ) *ProvisionVDBGroupFromBookmarkParameters`

NewProvisionVDBGroupFromBookmarkParameters instantiates a new ProvisionVDBGroupFromBookmarkParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionVDBGroupFromBookmarkParametersWithDefaults

`func NewProvisionVDBGroupFromBookmarkParametersWithDefaults() *ProvisionVDBGroupFromBookmarkParameters`

NewProvisionVDBGroupFromBookmarkParametersWithDefaults instantiates a new ProvisionVDBGroupFromBookmarkParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProvisionVDBGroupFromBookmarkParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisionVDBGroupFromBookmarkParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisionVDBGroupFromBookmarkParameters) SetName(v string)`

SetName sets Name field to given value.


### GetBookmarkId

`func (o *ProvisionVDBGroupFromBookmarkParameters) GetBookmarkId() string`

GetBookmarkId returns the BookmarkId field if non-nil, zero value otherwise.

### GetBookmarkIdOk

`func (o *ProvisionVDBGroupFromBookmarkParameters) GetBookmarkIdOk() (*string, bool)`

GetBookmarkIdOk returns a tuple with the BookmarkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarkId

`func (o *ProvisionVDBGroupFromBookmarkParameters) SetBookmarkId(v string)`

SetBookmarkId sets BookmarkId field to given value.


### GetProvisionParameters

`func (o *ProvisionVDBGroupFromBookmarkParameters) GetProvisionParameters() map[string]BaseProvisionVDBParameters`

GetProvisionParameters returns the ProvisionParameters field if non-nil, zero value otherwise.

### GetProvisionParametersOk

`func (o *ProvisionVDBGroupFromBookmarkParameters) GetProvisionParametersOk() (*map[string]BaseProvisionVDBParameters, bool)`

GetProvisionParametersOk returns a tuple with the ProvisionParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionParameters

`func (o *ProvisionVDBGroupFromBookmarkParameters) SetProvisionParameters(v map[string]BaseProvisionVDBParameters)`

SetProvisionParameters sets ProvisionParameters field to given value.


### GetTags

`func (o *ProvisionVDBGroupFromBookmarkParameters) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProvisionVDBGroupFromBookmarkParameters) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProvisionVDBGroupFromBookmarkParameters) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProvisionVDBGroupFromBookmarkParameters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMakeCurrentAccountOwner

`func (o *ProvisionVDBGroupFromBookmarkParameters) GetMakeCurrentAccountOwner() bool`

GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field if non-nil, zero value otherwise.

### GetMakeCurrentAccountOwnerOk

`func (o *ProvisionVDBGroupFromBookmarkParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool)`

GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCurrentAccountOwner

`func (o *ProvisionVDBGroupFromBookmarkParameters) SetMakeCurrentAccountOwner(v bool)`

SetMakeCurrentAccountOwner sets MakeCurrentAccountOwner field to given value.

### HasMakeCurrentAccountOwner

`func (o *ProvisionVDBGroupFromBookmarkParameters) HasMakeCurrentAccountOwner() bool`

HasMakeCurrentAccountOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


