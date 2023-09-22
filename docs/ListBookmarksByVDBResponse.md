# ListBookmarksByVDBResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Bookmark**](Bookmark.md) |  | [optional] 
**ResponseMetadata** | Pointer to [**PaginatedResponseMetadata**](PaginatedResponseMetadata.md) |  | [optional] 

## Methods

### NewListBookmarksByVDBResponse

`func NewListBookmarksByVDBResponse() *ListBookmarksByVDBResponse`

NewListBookmarksByVDBResponse instantiates a new ListBookmarksByVDBResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBookmarksByVDBResponseWithDefaults

`func NewListBookmarksByVDBResponseWithDefaults() *ListBookmarksByVDBResponse`

NewListBookmarksByVDBResponseWithDefaults instantiates a new ListBookmarksByVDBResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListBookmarksByVDBResponse) GetItems() []Bookmark`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListBookmarksByVDBResponse) GetItemsOk() (*[]Bookmark, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListBookmarksByVDBResponse) SetItems(v []Bookmark)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListBookmarksByVDBResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetResponseMetadata

`func (o *ListBookmarksByVDBResponse) GetResponseMetadata() PaginatedResponseMetadata`

GetResponseMetadata returns the ResponseMetadata field if non-nil, zero value otherwise.

### GetResponseMetadataOk

`func (o *ListBookmarksByVDBResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool)`

GetResponseMetadataOk returns a tuple with the ResponseMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMetadata

`func (o *ListBookmarksByVDBResponse) SetResponseMetadata(v PaginatedResponseMetadata)`

SetResponseMetadata sets ResponseMetadata field to given value.

### HasResponseMetadata

`func (o *ListBookmarksByVDBResponse) HasResponseMetadata() bool`

HasResponseMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


