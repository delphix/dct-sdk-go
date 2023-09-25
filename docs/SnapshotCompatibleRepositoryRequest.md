# SnapshotCompatibleRepositoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceDataId** | Pointer to **string** | The ID of the source object (dSource or VDB) to get the compatible repos. All other objects referenced by the parameters must live on the same engine as the source. | [optional] 
**EngineId** | Pointer to **string** | The ID of the Engine from where to get the compatible repos. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored. | [optional] 
**SnapshotId** | Pointer to **string** | The ID of the snapshot from which to execute the operation. | [optional] 
**EnvironmentId** | Pointer to **string** | The ID or name of the target environment. | [optional] 

## Methods

### NewSnapshotCompatibleRepositoryRequest

`func NewSnapshotCompatibleRepositoryRequest() *SnapshotCompatibleRepositoryRequest`

NewSnapshotCompatibleRepositoryRequest instantiates a new SnapshotCompatibleRepositoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotCompatibleRepositoryRequestWithDefaults

`func NewSnapshotCompatibleRepositoryRequestWithDefaults() *SnapshotCompatibleRepositoryRequest`

NewSnapshotCompatibleRepositoryRequestWithDefaults instantiates a new SnapshotCompatibleRepositoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceDataId

`func (o *SnapshotCompatibleRepositoryRequest) GetSourceDataId() string`

GetSourceDataId returns the SourceDataId field if non-nil, zero value otherwise.

### GetSourceDataIdOk

`func (o *SnapshotCompatibleRepositoryRequest) GetSourceDataIdOk() (*string, bool)`

GetSourceDataIdOk returns a tuple with the SourceDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataId

`func (o *SnapshotCompatibleRepositoryRequest) SetSourceDataId(v string)`

SetSourceDataId sets SourceDataId field to given value.

### HasSourceDataId

`func (o *SnapshotCompatibleRepositoryRequest) HasSourceDataId() bool`

HasSourceDataId returns a boolean if a field has been set.

### GetEngineId

`func (o *SnapshotCompatibleRepositoryRequest) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *SnapshotCompatibleRepositoryRequest) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *SnapshotCompatibleRepositoryRequest) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *SnapshotCompatibleRepositoryRequest) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetSnapshotId

`func (o *SnapshotCompatibleRepositoryRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *SnapshotCompatibleRepositoryRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *SnapshotCompatibleRepositoryRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *SnapshotCompatibleRepositoryRequest) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *SnapshotCompatibleRepositoryRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *SnapshotCompatibleRepositoryRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *SnapshotCompatibleRepositoryRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *SnapshotCompatibleRepositoryRequest) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


