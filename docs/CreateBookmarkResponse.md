# CreateBookmarkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bookmark** | Pointer to [**Bookmark**](Bookmark.md) |  | [optional] 
**Job** | Pointer to [**Job**](Job.md) |  | [optional] 

## Methods

### NewCreateBookmarkResponse

`func NewCreateBookmarkResponse() *CreateBookmarkResponse`

NewCreateBookmarkResponse instantiates a new CreateBookmarkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBookmarkResponseWithDefaults

`func NewCreateBookmarkResponseWithDefaults() *CreateBookmarkResponse`

NewCreateBookmarkResponseWithDefaults instantiates a new CreateBookmarkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookmark

`func (o *CreateBookmarkResponse) GetBookmark() Bookmark`

GetBookmark returns the Bookmark field if non-nil, zero value otherwise.

### GetBookmarkOk

`func (o *CreateBookmarkResponse) GetBookmarkOk() (*Bookmark, bool)`

GetBookmarkOk returns a tuple with the Bookmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmark

`func (o *CreateBookmarkResponse) SetBookmark(v Bookmark)`

SetBookmark sets Bookmark field to given value.

### HasBookmark

`func (o *CreateBookmarkResponse) HasBookmark() bool`

HasBookmark returns a boolean if a field has been set.

### GetJob

`func (o *CreateBookmarkResponse) GetJob() Job`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *CreateBookmarkResponse) GetJobOk() (*Job, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *CreateBookmarkResponse) SetJob(v Job)`

SetJob sets Job field to given value.

### HasJob

`func (o *CreateBookmarkResponse) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


