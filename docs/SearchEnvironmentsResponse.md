# SearchEnvironmentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Environment**](Environment.md) |  | [optional] 
**ResponseMetadata** | Pointer to [**PaginatedResponseMetadata**](PaginatedResponseMetadata.md) |  | [optional] 

## Methods

### NewSearchEnvironmentsResponse

`func NewSearchEnvironmentsResponse() *SearchEnvironmentsResponse`

NewSearchEnvironmentsResponse instantiates a new SearchEnvironmentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEnvironmentsResponseWithDefaults

`func NewSearchEnvironmentsResponseWithDefaults() *SearchEnvironmentsResponse`

NewSearchEnvironmentsResponseWithDefaults instantiates a new SearchEnvironmentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SearchEnvironmentsResponse) GetItems() []Environment`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SearchEnvironmentsResponse) GetItemsOk() (*[]Environment, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SearchEnvironmentsResponse) SetItems(v []Environment)`

SetItems sets Items field to given value.

### HasItems

`func (o *SearchEnvironmentsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetResponseMetadata

`func (o *SearchEnvironmentsResponse) GetResponseMetadata() PaginatedResponseMetadata`

GetResponseMetadata returns the ResponseMetadata field if non-nil, zero value otherwise.

### GetResponseMetadataOk

`func (o *SearchEnvironmentsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool)`

GetResponseMetadataOk returns a tuple with the ResponseMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMetadata

`func (o *SearchEnvironmentsResponse) SetResponseMetadata(v PaginatedResponseMetadata)`

SetResponseMetadata sets ResponseMetadata field to given value.

### HasResponseMetadata

`func (o *SearchEnvironmentsResponse) HasResponseMetadata() bool`

HasResponseMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


