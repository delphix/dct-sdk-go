# ScopedObjectItemsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]ScopedObjectItem**](ScopedObjectItem.md) | Array of access group scope objects with object id and object type | [optional] 

## Methods

### NewScopedObjectItemsResponse

`func NewScopedObjectItemsResponse() *ScopedObjectItemsResponse`

NewScopedObjectItemsResponse instantiates a new ScopedObjectItemsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopedObjectItemsResponseWithDefaults

`func NewScopedObjectItemsResponseWithDefaults() *ScopedObjectItemsResponse`

NewScopedObjectItemsResponseWithDefaults instantiates a new ScopedObjectItemsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *ScopedObjectItemsResponse) GetObjects() []ScopedObjectItem`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ScopedObjectItemsResponse) GetObjectsOk() (*[]ScopedObjectItem, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ScopedObjectItemsResponse) SetObjects(v []ScopedObjectItem)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ScopedObjectItemsResponse) HasObjects() bool`

HasObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


