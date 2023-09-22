# SnapshotCompatibleRepositories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repositories** | Pointer to [**[]EnvironmentRepository**](EnvironmentRepository.md) | Repositories corresponding to the snapshot. A Repository typically corresponds to a database installation. | [optional] 

## Methods

### NewSnapshotCompatibleRepositories

`func NewSnapshotCompatibleRepositories() *SnapshotCompatibleRepositories`

NewSnapshotCompatibleRepositories instantiates a new SnapshotCompatibleRepositories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotCompatibleRepositoriesWithDefaults

`func NewSnapshotCompatibleRepositoriesWithDefaults() *SnapshotCompatibleRepositories`

NewSnapshotCompatibleRepositoriesWithDefaults instantiates a new SnapshotCompatibleRepositories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositories

`func (o *SnapshotCompatibleRepositories) GetRepositories() []EnvironmentRepository`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *SnapshotCompatibleRepositories) GetRepositoriesOk() (*[]EnvironmentRepository, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *SnapshotCompatibleRepositories) SetRepositories(v []EnvironmentRepository)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *SnapshotCompatibleRepositories) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


