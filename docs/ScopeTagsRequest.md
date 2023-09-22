# ScopeTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | [**[]ScopeTag**](ScopeTag.md) | Array of tags with key value pairs along with optional object_type and permissions | 

## Methods

### NewScopeTagsRequest

`func NewScopeTagsRequest(tags []ScopeTag, ) *ScopeTagsRequest`

NewScopeTagsRequest instantiates a new ScopeTagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeTagsRequestWithDefaults

`func NewScopeTagsRequestWithDefaults() *ScopeTagsRequest`

NewScopeTagsRequestWithDefaults instantiates a new ScopeTagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *ScopeTagsRequest) GetTags() []ScopeTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ScopeTagsRequest) GetTagsOk() (*[]ScopeTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ScopeTagsRequest) SetTags(v []ScopeTag)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


