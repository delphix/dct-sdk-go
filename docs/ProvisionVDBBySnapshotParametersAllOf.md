# ProvisionVDBBySnapshotParametersAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineId** | Pointer to **string** | The ID of the Engine onto which to provision. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored. | [optional] 
**SourceDataId** | Pointer to **string** | The ID of the source object (dSource or VDB) to provision from. All other objects referenced by the parameters must live on the same engine as the source. If this property is not set, the data_source of the snapshot_id will be used. | [optional] 
**MakeCurrentAccountOwner** | Pointer to **bool** | Whether the account provisioning this VDB must be configured as owner of the VDB. | [optional] [default to true]

## Methods

### NewProvisionVDBBySnapshotParametersAllOf

`func NewProvisionVDBBySnapshotParametersAllOf() *ProvisionVDBBySnapshotParametersAllOf`

NewProvisionVDBBySnapshotParametersAllOf instantiates a new ProvisionVDBBySnapshotParametersAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionVDBBySnapshotParametersAllOfWithDefaults

`func NewProvisionVDBBySnapshotParametersAllOfWithDefaults() *ProvisionVDBBySnapshotParametersAllOf`

NewProvisionVDBBySnapshotParametersAllOfWithDefaults instantiates a new ProvisionVDBBySnapshotParametersAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineId

`func (o *ProvisionVDBBySnapshotParametersAllOf) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *ProvisionVDBBySnapshotParametersAllOf) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *ProvisionVDBBySnapshotParametersAllOf) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *ProvisionVDBBySnapshotParametersAllOf) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetSourceDataId

`func (o *ProvisionVDBBySnapshotParametersAllOf) GetSourceDataId() string`

GetSourceDataId returns the SourceDataId field if non-nil, zero value otherwise.

### GetSourceDataIdOk

`func (o *ProvisionVDBBySnapshotParametersAllOf) GetSourceDataIdOk() (*string, bool)`

GetSourceDataIdOk returns a tuple with the SourceDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataId

`func (o *ProvisionVDBBySnapshotParametersAllOf) SetSourceDataId(v string)`

SetSourceDataId sets SourceDataId field to given value.

### HasSourceDataId

`func (o *ProvisionVDBBySnapshotParametersAllOf) HasSourceDataId() bool`

HasSourceDataId returns a boolean if a field has been set.

### GetMakeCurrentAccountOwner

`func (o *ProvisionVDBBySnapshotParametersAllOf) GetMakeCurrentAccountOwner() bool`

GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field if non-nil, zero value otherwise.

### GetMakeCurrentAccountOwnerOk

`func (o *ProvisionVDBBySnapshotParametersAllOf) GetMakeCurrentAccountOwnerOk() (*bool, bool)`

GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCurrentAccountOwner

`func (o *ProvisionVDBBySnapshotParametersAllOf) SetMakeCurrentAccountOwner(v bool)`

SetMakeCurrentAccountOwner sets MakeCurrentAccountOwner field to given value.

### HasMakeCurrentAccountOwner

`func (o *ProvisionVDBBySnapshotParametersAllOf) HasMakeCurrentAccountOwner() bool`

HasMakeCurrentAccountOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


