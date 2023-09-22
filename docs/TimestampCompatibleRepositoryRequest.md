# TimestampCompatibleRepositoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceDataId** | Pointer to **string** | The ID of the source object (dSource or VDB) to get the compatible repos. All other objects referenced by the parameters must live on the same engine as the source. | [optional] 
**EngineId** | Pointer to **string** | The ID of the Engine from where to get the compatible repos. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored. | [optional] 
**Timestamp** | Pointer to **time.Time** | The point in time from which to execute the operation. If the timestamp is not set, selects the latest point. | [optional] 
**TimeflowId** | Pointer to **string** | ID of the timeflow from which compatible repos need to be fetched, mutually exclusive with source_data_id. | [optional] 
**EnvironmentId** | Pointer to **string** | The ID or name of the target environment. | [optional] 

## Methods

### NewTimestampCompatibleRepositoryRequest

`func NewTimestampCompatibleRepositoryRequest() *TimestampCompatibleRepositoryRequest`

NewTimestampCompatibleRepositoryRequest instantiates a new TimestampCompatibleRepositoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimestampCompatibleRepositoryRequestWithDefaults

`func NewTimestampCompatibleRepositoryRequestWithDefaults() *TimestampCompatibleRepositoryRequest`

NewTimestampCompatibleRepositoryRequestWithDefaults instantiates a new TimestampCompatibleRepositoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceDataId

`func (o *TimestampCompatibleRepositoryRequest) GetSourceDataId() string`

GetSourceDataId returns the SourceDataId field if non-nil, zero value otherwise.

### GetSourceDataIdOk

`func (o *TimestampCompatibleRepositoryRequest) GetSourceDataIdOk() (*string, bool)`

GetSourceDataIdOk returns a tuple with the SourceDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataId

`func (o *TimestampCompatibleRepositoryRequest) SetSourceDataId(v string)`

SetSourceDataId sets SourceDataId field to given value.

### HasSourceDataId

`func (o *TimestampCompatibleRepositoryRequest) HasSourceDataId() bool`

HasSourceDataId returns a boolean if a field has been set.

### GetEngineId

`func (o *TimestampCompatibleRepositoryRequest) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *TimestampCompatibleRepositoryRequest) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *TimestampCompatibleRepositoryRequest) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *TimestampCompatibleRepositoryRequest) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TimestampCompatibleRepositoryRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TimestampCompatibleRepositoryRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TimestampCompatibleRepositoryRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TimestampCompatibleRepositoryRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTimeflowId

`func (o *TimestampCompatibleRepositoryRequest) GetTimeflowId() string`

GetTimeflowId returns the TimeflowId field if non-nil, zero value otherwise.

### GetTimeflowIdOk

`func (o *TimestampCompatibleRepositoryRequest) GetTimeflowIdOk() (*string, bool)`

GetTimeflowIdOk returns a tuple with the TimeflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeflowId

`func (o *TimestampCompatibleRepositoryRequest) SetTimeflowId(v string)`

SetTimeflowId sets TimeflowId field to given value.

### HasTimeflowId

`func (o *TimestampCompatibleRepositoryRequest) HasTimeflowId() bool`

HasTimeflowId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *TimestampCompatibleRepositoryRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *TimestampCompatibleRepositoryRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *TimestampCompatibleRepositoryRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *TimestampCompatibleRepositoryRequest) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


