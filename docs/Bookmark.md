# Bookmark

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Bookmark object entity ID. | [optional] [readonly] 
**Name** | Pointer to **string** | The user-defined name of this bookmark. | [optional] 
**CreationDate** | Pointer to **time.Time** | The date and time that this bookmark was created. | [optional] [readonly] 
**VdbIds** | Pointer to **[]string** | The list of VDB IDs associated with this bookmark. | [optional] 
**Retention** | Pointer to **int64** | The retention policy for this bookmark, in days. A value of -1 indicates the bookmark should be kept forever. Deprecated in favor of expiration. | [optional] 
**Expiration** | Pointer to **string** | The expiration for this bookmark. When unset, indicates the bookmark is kept forever. | [optional] 
**Status** | Pointer to **NullableString** | A message with details about operation progress or state of this bookmark. | [optional] [readonly] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | The tags to be created for this Bookmark. | [optional] 

## Methods

### NewBookmark

`func NewBookmark() *Bookmark`

NewBookmark instantiates a new Bookmark object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookmarkWithDefaults

`func NewBookmarkWithDefaults() *Bookmark`

NewBookmarkWithDefaults instantiates a new Bookmark object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bookmark) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bookmark) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bookmark) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Bookmark) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Bookmark) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bookmark) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bookmark) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Bookmark) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreationDate

`func (o *Bookmark) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Bookmark) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Bookmark) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Bookmark) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetVdbIds

`func (o *Bookmark) GetVdbIds() []string`

GetVdbIds returns the VdbIds field if non-nil, zero value otherwise.

### GetVdbIdsOk

`func (o *Bookmark) GetVdbIdsOk() (*[]string, bool)`

GetVdbIdsOk returns a tuple with the VdbIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdbIds

`func (o *Bookmark) SetVdbIds(v []string)`

SetVdbIds sets VdbIds field to given value.

### HasVdbIds

`func (o *Bookmark) HasVdbIds() bool`

HasVdbIds returns a boolean if a field has been set.

### GetRetention

`func (o *Bookmark) GetRetention() int64`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *Bookmark) GetRetentionOk() (*int64, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *Bookmark) SetRetention(v int64)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *Bookmark) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetExpiration

`func (o *Bookmark) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *Bookmark) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *Bookmark) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *Bookmark) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetStatus

`func (o *Bookmark) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Bookmark) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Bookmark) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Bookmark) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Bookmark) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Bookmark) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTags

`func (o *Bookmark) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Bookmark) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Bookmark) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Bookmark) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


