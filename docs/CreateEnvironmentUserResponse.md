# CreateEnvironmentUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserRef** | Pointer to **string** | The reference of the created environment user | [optional] 
**Job** | Pointer to [**Job**](Job.md) |  | [optional] 

## Methods

### NewCreateEnvironmentUserResponse

`func NewCreateEnvironmentUserResponse() *CreateEnvironmentUserResponse`

NewCreateEnvironmentUserResponse instantiates a new CreateEnvironmentUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEnvironmentUserResponseWithDefaults

`func NewCreateEnvironmentUserResponseWithDefaults() *CreateEnvironmentUserResponse`

NewCreateEnvironmentUserResponseWithDefaults instantiates a new CreateEnvironmentUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserRef

`func (o *CreateEnvironmentUserResponse) GetUserRef() string`

GetUserRef returns the UserRef field if non-nil, zero value otherwise.

### GetUserRefOk

`func (o *CreateEnvironmentUserResponse) GetUserRefOk() (*string, bool)`

GetUserRefOk returns a tuple with the UserRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRef

`func (o *CreateEnvironmentUserResponse) SetUserRef(v string)`

SetUserRef sets UserRef field to given value.

### HasUserRef

`func (o *CreateEnvironmentUserResponse) HasUserRef() bool`

HasUserRef returns a boolean if a field has been set.

### GetJob

`func (o *CreateEnvironmentUserResponse) GetJob() Job`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *CreateEnvironmentUserResponse) GetJobOk() (*Job, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *CreateEnvironmentUserResponse) SetJob(v Job)`

SetJob sets Job field to given value.

### HasJob

`func (o *CreateEnvironmentUserResponse) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


