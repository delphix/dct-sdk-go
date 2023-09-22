# ProvisionVDBBySnapshotDefaultsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotId** | Pointer to **string** | The ID of the snapshot from which to execute the operation. | [optional] 
**EngineId** | Pointer to **string** | The ID of the Engine onto which to provision. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored. | [optional] 
**SourceDataId** | Pointer to **string** | The ID of the source object (dSource or VDB) to provision from. If this property is not set, the data_source of the snapshot_id will be used. | [optional] 

## Methods

### NewProvisionVDBBySnapshotDefaultsRequest

`func NewProvisionVDBBySnapshotDefaultsRequest() *ProvisionVDBBySnapshotDefaultsRequest`

NewProvisionVDBBySnapshotDefaultsRequest instantiates a new ProvisionVDBBySnapshotDefaultsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionVDBBySnapshotDefaultsRequestWithDefaults

`func NewProvisionVDBBySnapshotDefaultsRequestWithDefaults() *ProvisionVDBBySnapshotDefaultsRequest`

NewProvisionVDBBySnapshotDefaultsRequestWithDefaults instantiates a new ProvisionVDBBySnapshotDefaultsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotId

`func (o *ProvisionVDBBySnapshotDefaultsRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ProvisionVDBBySnapshotDefaultsRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ProvisionVDBBySnapshotDefaultsRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ProvisionVDBBySnapshotDefaultsRequest) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetEngineId

`func (o *ProvisionVDBBySnapshotDefaultsRequest) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *ProvisionVDBBySnapshotDefaultsRequest) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *ProvisionVDBBySnapshotDefaultsRequest) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *ProvisionVDBBySnapshotDefaultsRequest) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetSourceDataId

`func (o *ProvisionVDBBySnapshotDefaultsRequest) GetSourceDataId() string`

GetSourceDataId returns the SourceDataId field if non-nil, zero value otherwise.

### GetSourceDataIdOk

`func (o *ProvisionVDBBySnapshotDefaultsRequest) GetSourceDataIdOk() (*string, bool)`

GetSourceDataIdOk returns a tuple with the SourceDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataId

`func (o *ProvisionVDBBySnapshotDefaultsRequest) SetSourceDataId(v string)`

SetSourceDataId sets SourceDataId field to given value.

### HasSourceDataId

`func (o *ProvisionVDBBySnapshotDefaultsRequest) HasSourceDataId() bool`

HasSourceDataId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


