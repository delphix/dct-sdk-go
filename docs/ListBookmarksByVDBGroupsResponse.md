# ListBookmarksByVDBGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Bookmark**](Bookmark.md) |  | [optional] 
**ResponseMetadata** | Pointer to [**PaginatedResponseMetadata**](PaginatedResponseMetadata.md) |  | [optional] 

## Methods

### NewListBookmarksByVDBGroupsResponse

`func NewListBookmarksByVDBGroupsResponse() *ListBookmarksByVDBGroupsResponse`

NewListBookmarksByVDBGroupsResponse instantiates a new ListBookmarksByVDBGroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBookmarksByVDBGroupsResponseWithDefaults

`func NewListBookmarksByVDBGroupsResponseWithDefaults() *ListBookmarksByVDBGroupsResponse`

NewListBookmarksByVDBGroupsResponseWithDefaults instantiates a new ListBookmarksByVDBGroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListBookmarksByVDBGroupsResponse) GetItems() []Bookmark`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListBookmarksByVDBGroupsResponse) GetItemsOk() (*[]Bookmark, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListBookmarksByVDBGroupsResponse) SetItems(v []Bookmark)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListBookmarksByVDBGroupsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetResponseMetadata

`func (o *ListBookmarksByVDBGroupsResponse) GetResponseMetadata() PaginatedResponseMetadata`

GetResponseMetadata returns the ResponseMetadata field if non-nil, zero value otherwise.

### GetResponseMetadataOk

`func (o *ListBookmarksByVDBGroupsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool)`

GetResponseMetadataOk returns a tuple with the ResponseMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMetadata

`func (o *ListBookmarksByVDBGroupsResponse) SetResponseMetadata(v PaginatedResponseMetadata)`

SetResponseMetadata sets ResponseMetadata field to given value.

### HasResponseMetadata

`func (o *ListBookmarksByVDBGroupsResponse) HasResponseMetadata() bool`

HasResponseMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


