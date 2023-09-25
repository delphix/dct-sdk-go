# TagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | [**[]Tag**](Tag.md) | Array of tags with key value pairs | 

## Methods

### NewTagsRequest

`func NewTagsRequest(tags []Tag, ) *TagsRequest`

NewTagsRequest instantiates a new TagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagsRequestWithDefaults

`func NewTagsRequestWithDefaults() *TagsRequest`

NewTagsRequestWithDefaults instantiates a new TagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *TagsRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TagsRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TagsRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


