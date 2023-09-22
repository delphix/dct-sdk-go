# DeleteScopeObjectTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]ScopeTag**](ScopeTag.md) | List of scope tags to be deleted | [optional] 

## Methods

### NewDeleteScopeObjectTags

`func NewDeleteScopeObjectTags() *DeleteScopeObjectTags`

NewDeleteScopeObjectTags instantiates a new DeleteScopeObjectTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteScopeObjectTagsWithDefaults

`func NewDeleteScopeObjectTagsWithDefaults() *DeleteScopeObjectTags`

NewDeleteScopeObjectTagsWithDefaults instantiates a new DeleteScopeObjectTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *DeleteScopeObjectTags) GetTags() []ScopeTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeleteScopeObjectTags) GetTagsOk() (*[]ScopeTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeleteScopeObjectTags) SetTags(v []ScopeTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeleteScopeObjectTags) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


