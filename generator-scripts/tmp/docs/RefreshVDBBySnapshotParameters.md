# RefreshVDBBySnapshotParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotId** | Pointer to **string** | The ID of the snapshot from which to execute the operation. If the snapshot_id is not, selects the latest snapshot. | [optional] 

## Methods

### NewRefreshVDBBySnapshotParameters

`func NewRefreshVDBBySnapshotParameters() *RefreshVDBBySnapshotParameters`

NewRefreshVDBBySnapshotParameters instantiates a new RefreshVDBBySnapshotParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshVDBBySnapshotParametersWithDefaults

`func NewRefreshVDBBySnapshotParametersWithDefaults() *RefreshVDBBySnapshotParameters`

NewRefreshVDBBySnapshotParametersWithDefaults instantiates a new RefreshVDBBySnapshotParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotId

`func (o *RefreshVDBBySnapshotParameters) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *RefreshVDBBySnapshotParameters) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *RefreshVDBBySnapshotParameters) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *RefreshVDBBySnapshotParameters) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


