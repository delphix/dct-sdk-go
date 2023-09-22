# Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Environment object entity ID. | [optional] 
**Name** | Pointer to **string** | The name of this environment. | [optional] 
**Namespace** | Pointer to **NullableString** | The namespace of this environment for replicated and restored objects. | [optional] 
**EngineId** | Pointer to **string** | A reference to the Engine that this Environment connection is associated with. | [optional] 
**Enabled** | Pointer to **bool** | True if this environment is enabled. | [optional] 
**IsCluster** | Pointer to **bool** | True if this environment is a cluster of hosts. | [optional] 
**ClusterHome** | Pointer to **string** | Cluster home for RAC environment. | [optional] 
**IsWindowsTarget** | Pointer to **bool** | True if this windows environment is a target environment. | [optional] 
**Hosts** | Pointer to [**[]Host**](Host.md) | The hosts that are part of this environment. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | The tags to be created for this environment. | [optional] 
**Repositories** | Pointer to [**[]Repository**](Repository.md) | Repositories associated with this environment. A Repository typically corresponds to a database installation. | [optional] 
**Listeners** | Pointer to [**[]OracleListener**](OracleListener.md) | Oracle listeners associated with this environment. | [optional] 

## Methods

### NewEnvironment

`func NewEnvironment() *Environment`

NewEnvironment instantiates a new Environment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentWithDefaults

`func NewEnvironmentWithDefaults() *Environment`

NewEnvironmentWithDefaults instantiates a new Environment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Environment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Environment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Environment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Environment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Environment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Environment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Environment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Environment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *Environment) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Environment) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Environment) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Environment) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *Environment) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *Environment) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetEngineId

`func (o *Environment) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *Environment) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *Environment) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *Environment) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetEnabled

`func (o *Environment) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Environment) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Environment) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Environment) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIsCluster

`func (o *Environment) GetIsCluster() bool`

GetIsCluster returns the IsCluster field if non-nil, zero value otherwise.

### GetIsClusterOk

`func (o *Environment) GetIsClusterOk() (*bool, bool)`

GetIsClusterOk returns a tuple with the IsCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCluster

`func (o *Environment) SetIsCluster(v bool)`

SetIsCluster sets IsCluster field to given value.

### HasIsCluster

`func (o *Environment) HasIsCluster() bool`

HasIsCluster returns a boolean if a field has been set.

### GetClusterHome

`func (o *Environment) GetClusterHome() string`

GetClusterHome returns the ClusterHome field if non-nil, zero value otherwise.

### GetClusterHomeOk

`func (o *Environment) GetClusterHomeOk() (*string, bool)`

GetClusterHomeOk returns a tuple with the ClusterHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterHome

`func (o *Environment) SetClusterHome(v string)`

SetClusterHome sets ClusterHome field to given value.

### HasClusterHome

`func (o *Environment) HasClusterHome() bool`

HasClusterHome returns a boolean if a field has been set.

### GetIsWindowsTarget

`func (o *Environment) GetIsWindowsTarget() bool`

GetIsWindowsTarget returns the IsWindowsTarget field if non-nil, zero value otherwise.

### GetIsWindowsTargetOk

`func (o *Environment) GetIsWindowsTargetOk() (*bool, bool)`

GetIsWindowsTargetOk returns a tuple with the IsWindowsTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWindowsTarget

`func (o *Environment) SetIsWindowsTarget(v bool)`

SetIsWindowsTarget sets IsWindowsTarget field to given value.

### HasIsWindowsTarget

`func (o *Environment) HasIsWindowsTarget() bool`

HasIsWindowsTarget returns a boolean if a field has been set.

### GetHosts

`func (o *Environment) GetHosts() []Host`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *Environment) GetHostsOk() (*[]Host, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *Environment) SetHosts(v []Host)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *Environment) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetTags

`func (o *Environment) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Environment) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Environment) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Environment) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRepositories

`func (o *Environment) GetRepositories() []Repository`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *Environment) GetRepositoriesOk() (*[]Repository, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *Environment) SetRepositories(v []Repository)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *Environment) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetListeners

`func (o *Environment) GetListeners() []OracleListener`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *Environment) GetListenersOk() (*[]OracleListener, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *Environment) SetListeners(v []OracleListener)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *Environment) HasListeners() bool`

HasListeners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


