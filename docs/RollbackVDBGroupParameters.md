# RollbackVDBGroupParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BookmarkId** | Pointer to **string** | ID of a bookmark to rollback this VDB Group to. | [optional] 

## Methods

### NewRollbackVDBGroupParameters

`func NewRollbackVDBGroupParameters() *RollbackVDBGroupParameters`

NewRollbackVDBGroupParameters instantiates a new RollbackVDBGroupParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollbackVDBGroupParametersWithDefaults

`func NewRollbackVDBGroupParametersWithDefaults() *RollbackVDBGroupParameters`

NewRollbackVDBGroupParametersWithDefaults instantiates a new RollbackVDBGroupParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookmarkId

`func (o *RollbackVDBGroupParameters) GetBookmarkId() string`

GetBookmarkId returns the BookmarkId field if non-nil, zero value otherwise.

### GetBookmarkIdOk

`func (o *RollbackVDBGroupParameters) GetBookmarkIdOk() (*string, bool)`

GetBookmarkIdOk returns a tuple with the BookmarkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarkId

`func (o *RollbackVDBGroupParameters) SetBookmarkId(v string)`

SetBookmarkId sets BookmarkId field to given value.

### HasBookmarkId

`func (o *RollbackVDBGroupParameters) HasBookmarkId() bool`

HasBookmarkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


