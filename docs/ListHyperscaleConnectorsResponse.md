# ListHyperscaleConnectorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]HyperscaleConnector**](HyperscaleConnector.md) |  | [optional] 
**ResponseMetadata** | Pointer to [**PaginatedResponseMetadata**](PaginatedResponseMetadata.md) |  | [optional] 

## Methods

### NewListHyperscaleConnectorsResponse

`func NewListHyperscaleConnectorsResponse() *ListHyperscaleConnectorsResponse`

NewListHyperscaleConnectorsResponse instantiates a new ListHyperscaleConnectorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHyperscaleConnectorsResponseWithDefaults

`func NewListHyperscaleConnectorsResponseWithDefaults() *ListHyperscaleConnectorsResponse`

NewListHyperscaleConnectorsResponseWithDefaults instantiates a new ListHyperscaleConnectorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListHyperscaleConnectorsResponse) GetItems() []HyperscaleConnector`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListHyperscaleConnectorsResponse) GetItemsOk() (*[]HyperscaleConnector, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListHyperscaleConnectorsResponse) SetItems(v []HyperscaleConnector)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListHyperscaleConnectorsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetResponseMetadata

`func (o *ListHyperscaleConnectorsResponse) GetResponseMetadata() PaginatedResponseMetadata`

GetResponseMetadata returns the ResponseMetadata field if non-nil, zero value otherwise.

### GetResponseMetadataOk

`func (o *ListHyperscaleConnectorsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool)`

GetResponseMetadataOk returns a tuple with the ResponseMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMetadata

`func (o *ListHyperscaleConnectorsResponse) SetResponseMetadata(v PaginatedResponseMetadata)`

SetResponseMetadata sets ResponseMetadata field to given value.

### HasResponseMetadata

`func (o *ListHyperscaleConnectorsResponse) HasResponseMetadata() bool`

HasResponseMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


