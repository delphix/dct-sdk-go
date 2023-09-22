# ScopeTagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]ScopeTag**](ScopeTag.md) | Array of tags with key value pairs along with optional object_type and permissions | [optional] 

## Methods

### NewScopeTagsResponse

`func NewScopeTagsResponse() *ScopeTagsResponse`

NewScopeTagsResponse instantiates a new ScopeTagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeTagsResponseWithDefaults

`func NewScopeTagsResponseWithDefaults() *ScopeTagsResponse`

NewScopeTagsResponseWithDefaults instantiates a new ScopeTagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *ScopeTagsResponse) GetTags() []ScopeTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ScopeTagsResponse) GetTagsOk() (*[]ScopeTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ScopeTagsResponse) SetTags(v []ScopeTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ScopeTagsResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


