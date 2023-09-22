# ScopedObjectsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]ScopedObjectItem**](ScopedObjectItem.md) | An array of scoped objects | 

## Methods

### NewScopedObjectsRequest

`func NewScopedObjectsRequest(objects []ScopedObjectItem, ) *ScopedObjectsRequest`

NewScopedObjectsRequest instantiates a new ScopedObjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopedObjectsRequestWithDefaults

`func NewScopedObjectsRequestWithDefaults() *ScopedObjectsRequest`

NewScopedObjectsRequestWithDefaults instantiates a new ScopedObjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *ScopedObjectsRequest) GetObjects() []ScopedObjectItem`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ScopedObjectsRequest) GetObjectsOk() (*[]ScopedObjectItem, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ScopedObjectsRequest) SetObjects(v []ScopedObjectItem)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


