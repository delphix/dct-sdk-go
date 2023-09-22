# Timeflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Timeflow ID. | [optional] 
**EngineId** | Pointer to **string** | The ID of the engine the timeflow belongs to. | [optional] 
**Namespace** | Pointer to **NullableString** | Alternate namespace for this object, for replicated and restored timeflows. | [optional] 
**NamespaceId** | Pointer to **NullableString** | The namespace id of this timeflows. | [optional] 
**NamespaceName** | Pointer to **NullableString** | The namespace name of this timeflows. | [optional] 
**IsReplica** | Pointer to **bool** | Is this a replicated object. | [optional] 
**Name** | Pointer to **string** | The timeflow&#39;s name. | [optional] 
**DatasetId** | Pointer to **string** | The ID of the timeflow&#39;s dSource or VDB. | [optional] 
**CreationType** | Pointer to **string** | The source action that created the timeflow. | [optional] 
**ParentSnapshotId** | Pointer to **string** | The ID of the timeflow&#39;s parent snapshot. | [optional] 
**ParentPointLocation** | Pointer to **string** | The location on the parent timeflow from which this timeflow was provisioned. This will not be present for timeflows derived from linked sources. | [optional] 
**ParentPointTimestamp** | Pointer to **time.Time** | The timestamp on the parent timeflow from which this timeflow was provisioned. This will not be present for timeflows derived from linked sources. | [optional] 
**ParentPointTimeflowId** | Pointer to **string** | A reference to the parent timeflow from which this timeflow was provisioned. This will not be present for timeflows derived from linked sources. | [optional] 
**SourceDataTimestamp** | Pointer to **time.Time** | The timestamp on the root ancestor timeflow from which this timeflow originated. This logical time acts as reference to the origin source data. | [optional] 
**OracleIncarnationId** | Pointer to **string** | Oracle-specific incarnation identifier for this timeflow. | [optional] 
**OracleCdbTimeflowId** | Pointer to **string** | A reference to the mirror CDB timeflow if this is a timeflow for a PDB. | [optional] 
**OracleTdeUuid** | Pointer to **string** | The unique identifier for timeflow-specific TDE objects that reside outside of Delphix storage. | [optional] 
**MssqlDatabaseGuid** | Pointer to **string** | MSSQL-specific recovery branch identifier for this timeflow. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewTimeflow

`func NewTimeflow() *Timeflow`

NewTimeflow instantiates a new Timeflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeflowWithDefaults

`func NewTimeflowWithDefaults() *Timeflow`

NewTimeflowWithDefaults instantiates a new Timeflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Timeflow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Timeflow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Timeflow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Timeflow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEngineId

`func (o *Timeflow) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *Timeflow) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *Timeflow) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *Timeflow) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetNamespace

`func (o *Timeflow) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Timeflow) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Timeflow) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Timeflow) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *Timeflow) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *Timeflow) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetNamespaceId

`func (o *Timeflow) GetNamespaceId() string`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *Timeflow) GetNamespaceIdOk() (*string, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *Timeflow) SetNamespaceId(v string)`

SetNamespaceId sets NamespaceId field to given value.

### HasNamespaceId

`func (o *Timeflow) HasNamespaceId() bool`

HasNamespaceId returns a boolean if a field has been set.

### SetNamespaceIdNil

`func (o *Timeflow) SetNamespaceIdNil(b bool)`

 SetNamespaceIdNil sets the value for NamespaceId to be an explicit nil

### UnsetNamespaceId
`func (o *Timeflow) UnsetNamespaceId()`

UnsetNamespaceId ensures that no value is present for NamespaceId, not even an explicit nil
### GetNamespaceName

`func (o *Timeflow) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *Timeflow) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *Timeflow) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.

### HasNamespaceName

`func (o *Timeflow) HasNamespaceName() bool`

HasNamespaceName returns a boolean if a field has been set.

### SetNamespaceNameNil

`func (o *Timeflow) SetNamespaceNameNil(b bool)`

 SetNamespaceNameNil sets the value for NamespaceName to be an explicit nil

### UnsetNamespaceName
`func (o *Timeflow) UnsetNamespaceName()`

UnsetNamespaceName ensures that no value is present for NamespaceName, not even an explicit nil
### GetIsReplica

`func (o *Timeflow) GetIsReplica() bool`

GetIsReplica returns the IsReplica field if non-nil, zero value otherwise.

### GetIsReplicaOk

`func (o *Timeflow) GetIsReplicaOk() (*bool, bool)`

GetIsReplicaOk returns a tuple with the IsReplica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplica

`func (o *Timeflow) SetIsReplica(v bool)`

SetIsReplica sets IsReplica field to given value.

### HasIsReplica

`func (o *Timeflow) HasIsReplica() bool`

HasIsReplica returns a boolean if a field has been set.

### GetName

`func (o *Timeflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Timeflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Timeflow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Timeflow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDatasetId

`func (o *Timeflow) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Timeflow) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *Timeflow) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.

### HasDatasetId

`func (o *Timeflow) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### GetCreationType

`func (o *Timeflow) GetCreationType() string`

GetCreationType returns the CreationType field if non-nil, zero value otherwise.

### GetCreationTypeOk

`func (o *Timeflow) GetCreationTypeOk() (*string, bool)`

GetCreationTypeOk returns a tuple with the CreationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationType

`func (o *Timeflow) SetCreationType(v string)`

SetCreationType sets CreationType field to given value.

### HasCreationType

`func (o *Timeflow) HasCreationType() bool`

HasCreationType returns a boolean if a field has been set.

### GetParentSnapshotId

`func (o *Timeflow) GetParentSnapshotId() string`

GetParentSnapshotId returns the ParentSnapshotId field if non-nil, zero value otherwise.

### GetParentSnapshotIdOk

`func (o *Timeflow) GetParentSnapshotIdOk() (*string, bool)`

GetParentSnapshotIdOk returns a tuple with the ParentSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSnapshotId

`func (o *Timeflow) SetParentSnapshotId(v string)`

SetParentSnapshotId sets ParentSnapshotId field to given value.

### HasParentSnapshotId

`func (o *Timeflow) HasParentSnapshotId() bool`

HasParentSnapshotId returns a boolean if a field has been set.

### GetParentPointLocation

`func (o *Timeflow) GetParentPointLocation() string`

GetParentPointLocation returns the ParentPointLocation field if non-nil, zero value otherwise.

### GetParentPointLocationOk

`func (o *Timeflow) GetParentPointLocationOk() (*string, bool)`

GetParentPointLocationOk returns a tuple with the ParentPointLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPointLocation

`func (o *Timeflow) SetParentPointLocation(v string)`

SetParentPointLocation sets ParentPointLocation field to given value.

### HasParentPointLocation

`func (o *Timeflow) HasParentPointLocation() bool`

HasParentPointLocation returns a boolean if a field has been set.

### GetParentPointTimestamp

`func (o *Timeflow) GetParentPointTimestamp() time.Time`

GetParentPointTimestamp returns the ParentPointTimestamp field if non-nil, zero value otherwise.

### GetParentPointTimestampOk

`func (o *Timeflow) GetParentPointTimestampOk() (*time.Time, bool)`

GetParentPointTimestampOk returns a tuple with the ParentPointTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPointTimestamp

`func (o *Timeflow) SetParentPointTimestamp(v time.Time)`

SetParentPointTimestamp sets ParentPointTimestamp field to given value.

### HasParentPointTimestamp

`func (o *Timeflow) HasParentPointTimestamp() bool`

HasParentPointTimestamp returns a boolean if a field has been set.

### GetParentPointTimeflowId

`func (o *Timeflow) GetParentPointTimeflowId() string`

GetParentPointTimeflowId returns the ParentPointTimeflowId field if non-nil, zero value otherwise.

### GetParentPointTimeflowIdOk

`func (o *Timeflow) GetParentPointTimeflowIdOk() (*string, bool)`

GetParentPointTimeflowIdOk returns a tuple with the ParentPointTimeflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPointTimeflowId

`func (o *Timeflow) SetParentPointTimeflowId(v string)`

SetParentPointTimeflowId sets ParentPointTimeflowId field to given value.

### HasParentPointTimeflowId

`func (o *Timeflow) HasParentPointTimeflowId() bool`

HasParentPointTimeflowId returns a boolean if a field has been set.

### GetSourceDataTimestamp

`func (o *Timeflow) GetSourceDataTimestamp() time.Time`

GetSourceDataTimestamp returns the SourceDataTimestamp field if non-nil, zero value otherwise.

### GetSourceDataTimestampOk

`func (o *Timeflow) GetSourceDataTimestampOk() (*time.Time, bool)`

GetSourceDataTimestampOk returns a tuple with the SourceDataTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataTimestamp

`func (o *Timeflow) SetSourceDataTimestamp(v time.Time)`

SetSourceDataTimestamp sets SourceDataTimestamp field to given value.

### HasSourceDataTimestamp

`func (o *Timeflow) HasSourceDataTimestamp() bool`

HasSourceDataTimestamp returns a boolean if a field has been set.

### GetOracleIncarnationId

`func (o *Timeflow) GetOracleIncarnationId() string`

GetOracleIncarnationId returns the OracleIncarnationId field if non-nil, zero value otherwise.

### GetOracleIncarnationIdOk

`func (o *Timeflow) GetOracleIncarnationIdOk() (*string, bool)`

GetOracleIncarnationIdOk returns a tuple with the OracleIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleIncarnationId

`func (o *Timeflow) SetOracleIncarnationId(v string)`

SetOracleIncarnationId sets OracleIncarnationId field to given value.

### HasOracleIncarnationId

`func (o *Timeflow) HasOracleIncarnationId() bool`

HasOracleIncarnationId returns a boolean if a field has been set.

### GetOracleCdbTimeflowId

`func (o *Timeflow) GetOracleCdbTimeflowId() string`

GetOracleCdbTimeflowId returns the OracleCdbTimeflowId field if non-nil, zero value otherwise.

### GetOracleCdbTimeflowIdOk

`func (o *Timeflow) GetOracleCdbTimeflowIdOk() (*string, bool)`

GetOracleCdbTimeflowIdOk returns a tuple with the OracleCdbTimeflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleCdbTimeflowId

`func (o *Timeflow) SetOracleCdbTimeflowId(v string)`

SetOracleCdbTimeflowId sets OracleCdbTimeflowId field to given value.

### HasOracleCdbTimeflowId

`func (o *Timeflow) HasOracleCdbTimeflowId() bool`

HasOracleCdbTimeflowId returns a boolean if a field has been set.

### GetOracleTdeUuid

`func (o *Timeflow) GetOracleTdeUuid() string`

GetOracleTdeUuid returns the OracleTdeUuid field if non-nil, zero value otherwise.

### GetOracleTdeUuidOk

`func (o *Timeflow) GetOracleTdeUuidOk() (*string, bool)`

GetOracleTdeUuidOk returns a tuple with the OracleTdeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleTdeUuid

`func (o *Timeflow) SetOracleTdeUuid(v string)`

SetOracleTdeUuid sets OracleTdeUuid field to given value.

### HasOracleTdeUuid

`func (o *Timeflow) HasOracleTdeUuid() bool`

HasOracleTdeUuid returns a boolean if a field has been set.

### GetMssqlDatabaseGuid

`func (o *Timeflow) GetMssqlDatabaseGuid() string`

GetMssqlDatabaseGuid returns the MssqlDatabaseGuid field if non-nil, zero value otherwise.

### GetMssqlDatabaseGuidOk

`func (o *Timeflow) GetMssqlDatabaseGuidOk() (*string, bool)`

GetMssqlDatabaseGuidOk returns a tuple with the MssqlDatabaseGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlDatabaseGuid

`func (o *Timeflow) SetMssqlDatabaseGuid(v string)`

SetMssqlDatabaseGuid sets MssqlDatabaseGuid field to given value.

### HasMssqlDatabaseGuid

`func (o *Timeflow) HasMssqlDatabaseGuid() bool`

HasMssqlDatabaseGuid returns a boolean if a field has been set.

### GetTags

`func (o *Timeflow) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Timeflow) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Timeflow) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Timeflow) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


