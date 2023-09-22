# DataPointFromBookmarkParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BookmarkId** | **string** | The ID of the bookmark from which to execute the operation. The bookmark must contain only one VDB. | 

## Methods

### NewDataPointFromBookmarkParameters

`func NewDataPointFromBookmarkParameters(bookmarkId string, ) *DataPointFromBookmarkParameters`

NewDataPointFromBookmarkParameters instantiates a new DataPointFromBookmarkParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataPointFromBookmarkParametersWithDefaults

`func NewDataPointFromBookmarkParametersWithDefaults() *DataPointFromBookmarkParameters`

NewDataPointFromBookmarkParametersWithDefaults instantiates a new DataPointFromBookmarkParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookmarkId

`func (o *DataPointFromBookmarkParameters) GetBookmarkId() string`

GetBookmarkId returns the BookmarkId field if non-nil, zero value otherwise.

### GetBookmarkIdOk

`func (o *DataPointFromBookmarkParameters) GetBookmarkIdOk() (*string, bool)`

GetBookmarkIdOk returns a tuple with the BookmarkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarkId

`func (o *DataPointFromBookmarkParameters) SetBookmarkId(v string)`

SetBookmarkId sets BookmarkId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


