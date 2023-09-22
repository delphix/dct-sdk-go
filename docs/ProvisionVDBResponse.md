# ProvisionVDBResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to [**Job**](Job.md) |  | [optional] 
**VdbId** | Pointer to **string** | The ID of the provisioned vdb. | [optional] 

## Methods

### NewProvisionVDBResponse

`func NewProvisionVDBResponse() *ProvisionVDBResponse`

NewProvisionVDBResponse instantiates a new ProvisionVDBResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionVDBResponseWithDefaults

`func NewProvisionVDBResponseWithDefaults() *ProvisionVDBResponse`

NewProvisionVDBResponseWithDefaults instantiates a new ProvisionVDBResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *ProvisionVDBResponse) GetJob() Job`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *ProvisionVDBResponse) GetJobOk() (*Job, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *ProvisionVDBResponse) SetJob(v Job)`

SetJob sets Job field to given value.

### HasJob

`func (o *ProvisionVDBResponse) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetVdbId

`func (o *ProvisionVDBResponse) GetVdbId() string`

GetVdbId returns the VdbId field if non-nil, zero value otherwise.

### GetVdbIdOk

`func (o *ProvisionVDBResponse) GetVdbIdOk() (*string, bool)`

GetVdbIdOk returns a tuple with the VdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdbId

`func (o *ProvisionVDBResponse) SetVdbId(v string)`

SetVdbId sets VdbId field to given value.

### HasVdbId

`func (o *ProvisionVDBResponse) HasVdbId() bool`

HasVdbId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


