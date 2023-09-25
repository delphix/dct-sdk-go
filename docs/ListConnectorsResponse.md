# ListConnectorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Connector**](Connector.md) |  | [optional] 
**ResponseMetadata** | Pointer to [**PaginatedResponseMetadata**](PaginatedResponseMetadata.md) |  | [optional] 

## Methods

### NewListConnectorsResponse

`func NewListConnectorsResponse() *ListConnectorsResponse`

NewListConnectorsResponse instantiates a new ListConnectorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectorsResponseWithDefaults

`func NewListConnectorsResponseWithDefaults() *ListConnectorsResponse`

NewListConnectorsResponseWithDefaults instantiates a new ListConnectorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListConnectorsResponse) GetItems() []Connector`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListConnectorsResponse) GetItemsOk() (*[]Connector, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListConnectorsResponse) SetItems(v []Connector)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListConnectorsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetResponseMetadata

`func (o *ListConnectorsResponse) GetResponseMetadata() PaginatedResponseMetadata`

GetResponseMetadata returns the ResponseMetadata field if non-nil, zero value otherwise.

### GetResponseMetadataOk

`func (o *ListConnectorsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool)`

GetResponseMetadataOk returns a tuple with the ResponseMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMetadata

`func (o *ListConnectorsResponse) SetResponseMetadata(v PaginatedResponseMetadata)`

SetResponseMetadata sets ResponseMetadata field to given value.

### HasResponseMetadata

`func (o *ListConnectorsResponse) HasResponseMetadata() bool`

HasResponseMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


