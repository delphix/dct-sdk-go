# BookmarkCompatibleRepositoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BookmarkId** | **string** | The ID of the bookmark from which to execute the operation. The bookmark must contain only one VDB. | 
**EnvironmentId** | Pointer to **string** | The ID or name of the target environment. | [optional] 

## Methods

### NewBookmarkCompatibleRepositoryRequest

`func NewBookmarkCompatibleRepositoryRequest(bookmarkId string, ) *BookmarkCompatibleRepositoryRequest`

NewBookmarkCompatibleRepositoryRequest instantiates a new BookmarkCompatibleRepositoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookmarkCompatibleRepositoryRequestWithDefaults

`func NewBookmarkCompatibleRepositoryRequestWithDefaults() *BookmarkCompatibleRepositoryRequest`

NewBookmarkCompatibleRepositoryRequestWithDefaults instantiates a new BookmarkCompatibleRepositoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookmarkId

`func (o *BookmarkCompatibleRepositoryRequest) GetBookmarkId() string`

GetBookmarkId returns the BookmarkId field if non-nil, zero value otherwise.

### GetBookmarkIdOk

`func (o *BookmarkCompatibleRepositoryRequest) GetBookmarkIdOk() (*string, bool)`

GetBookmarkIdOk returns a tuple with the BookmarkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarkId

`func (o *BookmarkCompatibleRepositoryRequest) SetBookmarkId(v string)`

SetBookmarkId sets BookmarkId field to given value.


### GetEnvironmentId

`func (o *BookmarkCompatibleRepositoryRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *BookmarkCompatibleRepositoryRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *BookmarkCompatibleRepositoryRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *BookmarkCompatibleRepositoryRequest) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


