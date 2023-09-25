# CreateHostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to [**Job**](Job.md) |  | [optional] 
**ClusterNodeId** | Pointer to **string** | The id of the created cluster node. | [optional] 

## Methods

### NewCreateHostResponse

`func NewCreateHostResponse() *CreateHostResponse`

NewCreateHostResponse instantiates a new CreateHostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHostResponseWithDefaults

`func NewCreateHostResponseWithDefaults() *CreateHostResponse`

NewCreateHostResponseWithDefaults instantiates a new CreateHostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *CreateHostResponse) GetJob() Job`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *CreateHostResponse) GetJobOk() (*Job, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *CreateHostResponse) SetJob(v Job)`

SetJob sets Job field to given value.

### HasJob

`func (o *CreateHostResponse) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetClusterNodeId

`func (o *CreateHostResponse) GetClusterNodeId() string`

GetClusterNodeId returns the ClusterNodeId field if non-nil, zero value otherwise.

### GetClusterNodeIdOk

`func (o *CreateHostResponse) GetClusterNodeIdOk() (*string, bool)`

GetClusterNodeIdOk returns a tuple with the ClusterNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeId

`func (o *CreateHostResponse) SetClusterNodeId(v string)`

SetClusterNodeId sets ClusterNodeId field to given value.

### HasClusterNodeId

`func (o *CreateHostResponse) HasClusterNodeId() bool`

HasClusterNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


