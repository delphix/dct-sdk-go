# ProvisionVDBFromBookmarkDefaultsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BookmarkId** | **string** | The ID of the bookmark from which to execute the operation. The bookmark must contain only one VDB. | 

## Methods

### NewProvisionVDBFromBookmarkDefaultsRequest

`func NewProvisionVDBFromBookmarkDefaultsRequest(bookmarkId string, ) *ProvisionVDBFromBookmarkDefaultsRequest`

NewProvisionVDBFromBookmarkDefaultsRequest instantiates a new ProvisionVDBFromBookmarkDefaultsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionVDBFromBookmarkDefaultsRequestWithDefaults

`func NewProvisionVDBFromBookmarkDefaultsRequestWithDefaults() *ProvisionVDBFromBookmarkDefaultsRequest`

NewProvisionVDBFromBookmarkDefaultsRequestWithDefaults instantiates a new ProvisionVDBFromBookmarkDefaultsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookmarkId

`func (o *ProvisionVDBFromBookmarkDefaultsRequest) GetBookmarkId() string`

GetBookmarkId returns the BookmarkId field if non-nil, zero value otherwise.

### GetBookmarkIdOk

`func (o *ProvisionVDBFromBookmarkDefaultsRequest) GetBookmarkIdOk() (*string, bool)`

GetBookmarkIdOk returns a tuple with the BookmarkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarkId

`func (o *ProvisionVDBFromBookmarkDefaultsRequest) SetBookmarkId(v string)`

SetBookmarkId sets BookmarkId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


