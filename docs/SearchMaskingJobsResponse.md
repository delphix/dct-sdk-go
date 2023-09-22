# SearchMaskingJobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]MaskingJob**](MaskingJob.md) |  | [optional] 
**ResponseMetadata** | Pointer to [**PaginatedResponseMetadata**](PaginatedResponseMetadata.md) |  | [optional] 

## Methods

### NewSearchMaskingJobsResponse

`func NewSearchMaskingJobsResponse() *SearchMaskingJobsResponse`

NewSearchMaskingJobsResponse instantiates a new SearchMaskingJobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchMaskingJobsResponseWithDefaults

`func NewSearchMaskingJobsResponseWithDefaults() *SearchMaskingJobsResponse`

NewSearchMaskingJobsResponseWithDefaults instantiates a new SearchMaskingJobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SearchMaskingJobsResponse) GetItems() []MaskingJob`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SearchMaskingJobsResponse) GetItemsOk() (*[]MaskingJob, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SearchMaskingJobsResponse) SetItems(v []MaskingJob)`

SetItems sets Items field to given value.

### HasItems

`func (o *SearchMaskingJobsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetResponseMetadata

`func (o *SearchMaskingJobsResponse) GetResponseMetadata() PaginatedResponseMetadata`

GetResponseMetadata returns the ResponseMetadata field if non-nil, zero value otherwise.

### GetResponseMetadataOk

`func (o *SearchMaskingJobsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool)`

GetResponseMetadataOk returns a tuple with the ResponseMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMetadata

`func (o *SearchMaskingJobsResponse) SetResponseMetadata(v PaginatedResponseMetadata)`

SetResponseMetadata sets ResponseMetadata field to given value.

### HasResponseMetadata

`func (o *SearchMaskingJobsResponse) HasResponseMetadata() bool`

HasResponseMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


