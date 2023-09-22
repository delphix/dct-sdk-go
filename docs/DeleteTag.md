# DeleteTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key of the tag | [optional] 
**Value** | Pointer to **string** | Value of the tag | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of tags to be deleted | [optional] 

## Methods

### NewDeleteTag

`func NewDeleteTag() *DeleteTag`

NewDeleteTag instantiates a new DeleteTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteTagWithDefaults

`func NewDeleteTagWithDefaults() *DeleteTag`

NewDeleteTagWithDefaults instantiates a new DeleteTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DeleteTag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DeleteTag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DeleteTag) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DeleteTag) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *DeleteTag) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DeleteTag) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DeleteTag) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DeleteTag) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTags

`func (o *DeleteTag) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeleteTag) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeleteTag) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeleteTag) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


