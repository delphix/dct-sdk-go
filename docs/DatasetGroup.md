# DatasetGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The dataset group ID. | [optional] 
**Name** | Pointer to **string** | The name of this dataset group. | [optional] 
**NamespaceId** | Pointer to **string** | The namespace id of this dataset group. | [optional] 
**NamespaceName** | Pointer to **string** | The namespace name of this dataset group. | [optional] 
**IsReplica** | Pointer to **bool** | Is this a replicated object. | [optional] 
**EngineId** | Pointer to **string** | Id of the Engine that this dataset group belongs to. | [optional] 
**EngineName** | Pointer to **string** | Name of the Engine that this dataset group belongs to. | [optional] 
**Namespace** | Pointer to **string** | The namespace of this dataset group. | [optional] 

## Methods

### NewDatasetGroup

`func NewDatasetGroup() *DatasetGroup`

NewDatasetGroup instantiates a new DatasetGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetGroupWithDefaults

`func NewDatasetGroupWithDefaults() *DatasetGroup`

NewDatasetGroupWithDefaults instantiates a new DatasetGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatasetGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatasetGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatasetGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DatasetGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DatasetGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasetGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasetGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatasetGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespaceId

`func (o *DatasetGroup) GetNamespaceId() string`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *DatasetGroup) GetNamespaceIdOk() (*string, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *DatasetGroup) SetNamespaceId(v string)`

SetNamespaceId sets NamespaceId field to given value.

### HasNamespaceId

`func (o *DatasetGroup) HasNamespaceId() bool`

HasNamespaceId returns a boolean if a field has been set.

### GetNamespaceName

`func (o *DatasetGroup) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *DatasetGroup) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *DatasetGroup) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.

### HasNamespaceName

`func (o *DatasetGroup) HasNamespaceName() bool`

HasNamespaceName returns a boolean if a field has been set.

### GetIsReplica

`func (o *DatasetGroup) GetIsReplica() bool`

GetIsReplica returns the IsReplica field if non-nil, zero value otherwise.

### GetIsReplicaOk

`func (o *DatasetGroup) GetIsReplicaOk() (*bool, bool)`

GetIsReplicaOk returns a tuple with the IsReplica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplica

`func (o *DatasetGroup) SetIsReplica(v bool)`

SetIsReplica sets IsReplica field to given value.

### HasIsReplica

`func (o *DatasetGroup) HasIsReplica() bool`

HasIsReplica returns a boolean if a field has been set.

### GetEngineId

`func (o *DatasetGroup) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *DatasetGroup) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *DatasetGroup) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *DatasetGroup) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetEngineName

`func (o *DatasetGroup) GetEngineName() string`

GetEngineName returns the EngineName field if non-nil, zero value otherwise.

### GetEngineNameOk

`func (o *DatasetGroup) GetEngineNameOk() (*string, bool)`

GetEngineNameOk returns a tuple with the EngineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineName

`func (o *DatasetGroup) SetEngineName(v string)`

SetEngineName sets EngineName field to given value.

### HasEngineName

`func (o *DatasetGroup) HasEngineName() bool`

HasEngineName returns a boolean if a field has been set.

### GetNamespace

`func (o *DatasetGroup) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DatasetGroup) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DatasetGroup) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *DatasetGroup) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


