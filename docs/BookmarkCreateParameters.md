# BookmarkCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The user-defined name of this bookmark. | 
**VdbIds** | Pointer to **[]string** | The IDs of the VDBs to create the Bookmark on. This parameter is mutually exclusive with snapshot_ids. | [optional] 
**SnapshotIds** | Pointer to **[]string** | The IDs of the snapshots that will be part of the Bookmark. This parameter is mutually exclusive with vdb_ids.  | [optional] 
**Retention** | Pointer to **int64** | The retention policy for this bookmark, in days. A value of -1 indicates the bookmark should be kept forever. Deprecated in favor of expiration and retain_forever. | [optional] 
**Expiration** | Pointer to **string** | The expiration for this bookmark. Mutually exclusive with retention and retain_forever. | [optional] 
**RetainForever** | Pointer to **bool** | Indicates that the bookmark should be retained forever. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | The tags to be created for this Bookmark. | [optional] 
**MakeCurrentAccountOwner** | Pointer to **bool** | Whether the account creating this bookmark must be configured as owner of the bookmark. | [optional] [default to true]

## Methods

### NewBookmarkCreateParameters

`func NewBookmarkCreateParameters(name string, ) *BookmarkCreateParameters`

NewBookmarkCreateParameters instantiates a new BookmarkCreateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookmarkCreateParametersWithDefaults

`func NewBookmarkCreateParametersWithDefaults() *BookmarkCreateParameters`

NewBookmarkCreateParametersWithDefaults instantiates a new BookmarkCreateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BookmarkCreateParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BookmarkCreateParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BookmarkCreateParameters) SetName(v string)`

SetName sets Name field to given value.


### GetVdbIds

`func (o *BookmarkCreateParameters) GetVdbIds() []string`

GetVdbIds returns the VdbIds field if non-nil, zero value otherwise.

### GetVdbIdsOk

`func (o *BookmarkCreateParameters) GetVdbIdsOk() (*[]string, bool)`

GetVdbIdsOk returns a tuple with the VdbIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdbIds

`func (o *BookmarkCreateParameters) SetVdbIds(v []string)`

SetVdbIds sets VdbIds field to given value.

### HasVdbIds

`func (o *BookmarkCreateParameters) HasVdbIds() bool`

HasVdbIds returns a boolean if a field has been set.

### GetSnapshotIds

`func (o *BookmarkCreateParameters) GetSnapshotIds() []string`

GetSnapshotIds returns the SnapshotIds field if non-nil, zero value otherwise.

### GetSnapshotIdsOk

`func (o *BookmarkCreateParameters) GetSnapshotIdsOk() (*[]string, bool)`

GetSnapshotIdsOk returns a tuple with the SnapshotIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotIds

`func (o *BookmarkCreateParameters) SetSnapshotIds(v []string)`

SetSnapshotIds sets SnapshotIds field to given value.

### HasSnapshotIds

`func (o *BookmarkCreateParameters) HasSnapshotIds() bool`

HasSnapshotIds returns a boolean if a field has been set.

### GetRetention

`func (o *BookmarkCreateParameters) GetRetention() int64`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *BookmarkCreateParameters) GetRetentionOk() (*int64, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *BookmarkCreateParameters) SetRetention(v int64)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *BookmarkCreateParameters) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetExpiration

`func (o *BookmarkCreateParameters) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *BookmarkCreateParameters) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *BookmarkCreateParameters) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *BookmarkCreateParameters) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetRetainForever

`func (o *BookmarkCreateParameters) GetRetainForever() bool`

GetRetainForever returns the RetainForever field if non-nil, zero value otherwise.

### GetRetainForeverOk

`func (o *BookmarkCreateParameters) GetRetainForeverOk() (*bool, bool)`

GetRetainForeverOk returns a tuple with the RetainForever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainForever

`func (o *BookmarkCreateParameters) SetRetainForever(v bool)`

SetRetainForever sets RetainForever field to given value.

### HasRetainForever

`func (o *BookmarkCreateParameters) HasRetainForever() bool`

HasRetainForever returns a boolean if a field has been set.

### GetTags

`func (o *BookmarkCreateParameters) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BookmarkCreateParameters) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BookmarkCreateParameters) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BookmarkCreateParameters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMakeCurrentAccountOwner

`func (o *BookmarkCreateParameters) GetMakeCurrentAccountOwner() bool`

GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field if non-nil, zero value otherwise.

### GetMakeCurrentAccountOwnerOk

`func (o *BookmarkCreateParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool)`

GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCurrentAccountOwner

`func (o *BookmarkCreateParameters) SetMakeCurrentAccountOwner(v bool)`

SetMakeCurrentAccountOwner sets MakeCurrentAccountOwner field to given value.

### HasMakeCurrentAccountOwner

`func (o *BookmarkCreateParameters) HasMakeCurrentAccountOwner() bool`

HasMakeCurrentAccountOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


