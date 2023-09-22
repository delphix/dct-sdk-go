# SearchAccessGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]AccessGroup**](AccessGroup.md) |  | [optional] 
**ResponseMetadata** | Pointer to [**PaginatedResponseMetadata**](PaginatedResponseMetadata.md) |  | [optional] 

## Methods

### NewSearchAccessGroupsResponse

`func NewSearchAccessGroupsResponse() *SearchAccessGroupsResponse`

NewSearchAccessGroupsResponse instantiates a new SearchAccessGroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchAccessGroupsResponseWithDefaults

`func NewSearchAccessGroupsResponseWithDefaults() *SearchAccessGroupsResponse`

NewSearchAccessGroupsResponseWithDefaults instantiates a new SearchAccessGroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SearchAccessGroupsResponse) GetItems() []AccessGroup`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SearchAccessGroupsResponse) GetItemsOk() (*[]AccessGroup, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SearchAccessGroupsResponse) SetItems(v []AccessGroup)`

SetItems sets Items field to given value.

### HasItems

`func (o *SearchAccessGroupsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetResponseMetadata

`func (o *SearchAccessGroupsResponse) GetResponseMetadata() PaginatedResponseMetadata`

GetResponseMetadata returns the ResponseMetadata field if non-nil, zero value otherwise.

### GetResponseMetadataOk

`func (o *SearchAccessGroupsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool)`

GetResponseMetadataOk returns a tuple with the ResponseMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMetadata

`func (o *SearchAccessGroupsResponse) SetResponseMetadata(v PaginatedResponseMetadata)`

SetResponseMetadata sets ResponseMetadata field to given value.

### HasResponseMetadata

`func (o *SearchAccessGroupsResponse) HasResponseMetadata() bool`

HasResponseMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


