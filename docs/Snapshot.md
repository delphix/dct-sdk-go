# Snapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Snapshot ID. | [optional] 
**EngineId** | Pointer to **string** | The id of the engine the snapshot belongs to. | [optional] 
**Namespace** | Pointer to **NullableString** | Alternate namespace for this object, for replicated and restored snapshots. | [optional] 
**Name** | Pointer to **string** | The snapshot&#39;s name. | [optional] 
**Consistency** | Pointer to **string** | Indicates what type of recovery strategies must be invoked when provisioning from this snapshot. | [optional] 
**MissingNonLoggedData** | Pointer to **bool** | Indicates if a virtual database provisioned from this snapshot will be missing nologging changes. | [optional] 
**DatasetId** | Pointer to **string** | The ID of the Snapshot&#39;s dSource or VDB. | [optional] 
**CreationTime** | Pointer to **time.Time** | The time when the snapshot was created. | [optional] 
**StartTimestamp** | Pointer to **time.Time** | The timestamp within the parent TimeFlow at which this snapshot was initiated. \\ No recovery earlier than this point needs to be applied in order to provision a database from \\ this snapshot. If start_timestamp equals timestamp, then no recovery needs to be \\ applied in order to provision a database.  | [optional] 
**StartLocation** | Pointer to **string** | The database specific indentifier within the parent TimeFlow at which this snapshot was initiated. \\ No recovery earlier than this point needs to be applied in order to provision a database from \\ this snapshot. If start_location equals location, then no recovery needs to be \\ applied in order to provision a database.  | [optional] 
**Timestamp** | Pointer to **time.Time** | The logical time of the data contained in this Snapshot. | [optional] 
**Location** | Pointer to **string** | Database specific identifier for the data contained in this Snapshot, such as the Log Sequence Number (LSN) for MSsql databases, System Change Number (SCN) for Oracle databases. | [optional] 
**Retention** | Pointer to **int64** | Retention policy, in days. A value of -1 indicates the snapshot should be kept forever. Deprecated in favor of expiration and retain_forever. | [optional] 
**Expiration** | Pointer to **string** | The expiration date of this snapshot. If this is unset and retain_forever is false, the snapshot is subject to the retention policy of its dataset. | [optional] 
**RetainForever** | Pointer to **bool** | Indicates that the snapshot is protected from retention, i.e it will be kept forever. If false, see expiration. | [optional] 
**TimeflowId** | Pointer to **string** | The TimeFlow this snapshot was taken on. | [optional] 
**Timezone** | Pointer to **string** | Time zone of the source database at the time the snapshot was taken. | [optional] 
**Version** | Pointer to **NullableString** | Version of database source repository at the time the snapshot was taken. | [optional] 
**Temporary** | Pointer to **bool** | Indicates that this snapshot is in a transient state and should not be user visible. | [optional] 
**AppdataToolkit** | Pointer to **string** | The toolkit associated with this snapshot. | [optional] 
**AppdataMetadata** | Pointer to **string** | The JSON payload conforming to the DraftV4 schema based on the type of application data being manipulated. | [optional] 
**AseDbEncryptionKey** | Pointer to **string** | Database encryption key present for this snapshot. | [optional] 
**MssqlInternalVersion** | Pointer to **int32** | Internal version of the source database at the time the snapshot was taken. | [optional] 
**MssqlBackupSetUuid** | Pointer to **string** | UUID of the source database backup that was restored for this snapshot. | [optional] 
**MssqlBackupSoftwareType** | Pointer to **string** | Backup software used to restore the source database backup for this snapshot | [optional] 
**MssqlBackupLocationType** | Pointer to **string** | Backup software used to restore the source database backup for this snapshot | [optional] 
**MssqlEmptySnapshot** | Pointer to **bool** | True if the staging push dSource snapshot is empty. | [optional] 
**OracleFromPhysicalStandbyVdb** | Pointer to **bool** | True if this snapshot was taken of a standby database. | [optional] 
**OracleRedoLogSizeInBytes** | Pointer to **int64** | Online redo log size in bytes when this snapshot was taken. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewSnapshot

`func NewSnapshot() *Snapshot`

NewSnapshot instantiates a new Snapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotWithDefaults

`func NewSnapshotWithDefaults() *Snapshot`

NewSnapshotWithDefaults instantiates a new Snapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Snapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Snapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Snapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Snapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEngineId

`func (o *Snapshot) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *Snapshot) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *Snapshot) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *Snapshot) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetNamespace

`func (o *Snapshot) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Snapshot) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Snapshot) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Snapshot) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *Snapshot) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *Snapshot) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetName

`func (o *Snapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Snapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Snapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Snapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConsistency

`func (o *Snapshot) GetConsistency() string`

GetConsistency returns the Consistency field if non-nil, zero value otherwise.

### GetConsistencyOk

`func (o *Snapshot) GetConsistencyOk() (*string, bool)`

GetConsistencyOk returns a tuple with the Consistency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistency

`func (o *Snapshot) SetConsistency(v string)`

SetConsistency sets Consistency field to given value.

### HasConsistency

`func (o *Snapshot) HasConsistency() bool`

HasConsistency returns a boolean if a field has been set.

### GetMissingNonLoggedData

`func (o *Snapshot) GetMissingNonLoggedData() bool`

GetMissingNonLoggedData returns the MissingNonLoggedData field if non-nil, zero value otherwise.

### GetMissingNonLoggedDataOk

`func (o *Snapshot) GetMissingNonLoggedDataOk() (*bool, bool)`

GetMissingNonLoggedDataOk returns a tuple with the MissingNonLoggedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingNonLoggedData

`func (o *Snapshot) SetMissingNonLoggedData(v bool)`

SetMissingNonLoggedData sets MissingNonLoggedData field to given value.

### HasMissingNonLoggedData

`func (o *Snapshot) HasMissingNonLoggedData() bool`

HasMissingNonLoggedData returns a boolean if a field has been set.

### GetDatasetId

`func (o *Snapshot) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Snapshot) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *Snapshot) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.

### HasDatasetId

`func (o *Snapshot) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### GetCreationTime

`func (o *Snapshot) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Snapshot) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Snapshot) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Snapshot) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *Snapshot) GetStartTimestamp() time.Time`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *Snapshot) GetStartTimestampOk() (*time.Time, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *Snapshot) SetStartTimestamp(v time.Time)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *Snapshot) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetStartLocation

`func (o *Snapshot) GetStartLocation() string`

GetStartLocation returns the StartLocation field if non-nil, zero value otherwise.

### GetStartLocationOk

`func (o *Snapshot) GetStartLocationOk() (*string, bool)`

GetStartLocationOk returns a tuple with the StartLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLocation

`func (o *Snapshot) SetStartLocation(v string)`

SetStartLocation sets StartLocation field to given value.

### HasStartLocation

`func (o *Snapshot) HasStartLocation() bool`

HasStartLocation returns a boolean if a field has been set.

### GetTimestamp

`func (o *Snapshot) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Snapshot) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Snapshot) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Snapshot) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetLocation

`func (o *Snapshot) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Snapshot) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Snapshot) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Snapshot) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRetention

`func (o *Snapshot) GetRetention() int64`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *Snapshot) GetRetentionOk() (*int64, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *Snapshot) SetRetention(v int64)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *Snapshot) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetExpiration

`func (o *Snapshot) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *Snapshot) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *Snapshot) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *Snapshot) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetRetainForever

`func (o *Snapshot) GetRetainForever() bool`

GetRetainForever returns the RetainForever field if non-nil, zero value otherwise.

### GetRetainForeverOk

`func (o *Snapshot) GetRetainForeverOk() (*bool, bool)`

GetRetainForeverOk returns a tuple with the RetainForever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainForever

`func (o *Snapshot) SetRetainForever(v bool)`

SetRetainForever sets RetainForever field to given value.

### HasRetainForever

`func (o *Snapshot) HasRetainForever() bool`

HasRetainForever returns a boolean if a field has been set.

### GetTimeflowId

`func (o *Snapshot) GetTimeflowId() string`

GetTimeflowId returns the TimeflowId field if non-nil, zero value otherwise.

### GetTimeflowIdOk

`func (o *Snapshot) GetTimeflowIdOk() (*string, bool)`

GetTimeflowIdOk returns a tuple with the TimeflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeflowId

`func (o *Snapshot) SetTimeflowId(v string)`

SetTimeflowId sets TimeflowId field to given value.

### HasTimeflowId

`func (o *Snapshot) HasTimeflowId() bool`

HasTimeflowId returns a boolean if a field has been set.

### GetTimezone

`func (o *Snapshot) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Snapshot) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Snapshot) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Snapshot) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetVersion

`func (o *Snapshot) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Snapshot) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Snapshot) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Snapshot) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Snapshot) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Snapshot) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetTemporary

`func (o *Snapshot) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *Snapshot) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *Snapshot) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *Snapshot) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.

### GetAppdataToolkit

`func (o *Snapshot) GetAppdataToolkit() string`

GetAppdataToolkit returns the AppdataToolkit field if non-nil, zero value otherwise.

### GetAppdataToolkitOk

`func (o *Snapshot) GetAppdataToolkitOk() (*string, bool)`

GetAppdataToolkitOk returns a tuple with the AppdataToolkit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppdataToolkit

`func (o *Snapshot) SetAppdataToolkit(v string)`

SetAppdataToolkit sets AppdataToolkit field to given value.

### HasAppdataToolkit

`func (o *Snapshot) HasAppdataToolkit() bool`

HasAppdataToolkit returns a boolean if a field has been set.

### GetAppdataMetadata

`func (o *Snapshot) GetAppdataMetadata() string`

GetAppdataMetadata returns the AppdataMetadata field if non-nil, zero value otherwise.

### GetAppdataMetadataOk

`func (o *Snapshot) GetAppdataMetadataOk() (*string, bool)`

GetAppdataMetadataOk returns a tuple with the AppdataMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppdataMetadata

`func (o *Snapshot) SetAppdataMetadata(v string)`

SetAppdataMetadata sets AppdataMetadata field to given value.

### HasAppdataMetadata

`func (o *Snapshot) HasAppdataMetadata() bool`

HasAppdataMetadata returns a boolean if a field has been set.

### GetAseDbEncryptionKey

`func (o *Snapshot) GetAseDbEncryptionKey() string`

GetAseDbEncryptionKey returns the AseDbEncryptionKey field if non-nil, zero value otherwise.

### GetAseDbEncryptionKeyOk

`func (o *Snapshot) GetAseDbEncryptionKeyOk() (*string, bool)`

GetAseDbEncryptionKeyOk returns a tuple with the AseDbEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbEncryptionKey

`func (o *Snapshot) SetAseDbEncryptionKey(v string)`

SetAseDbEncryptionKey sets AseDbEncryptionKey field to given value.

### HasAseDbEncryptionKey

`func (o *Snapshot) HasAseDbEncryptionKey() bool`

HasAseDbEncryptionKey returns a boolean if a field has been set.

### GetMssqlInternalVersion

`func (o *Snapshot) GetMssqlInternalVersion() int32`

GetMssqlInternalVersion returns the MssqlInternalVersion field if non-nil, zero value otherwise.

### GetMssqlInternalVersionOk

`func (o *Snapshot) GetMssqlInternalVersionOk() (*int32, bool)`

GetMssqlInternalVersionOk returns a tuple with the MssqlInternalVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlInternalVersion

`func (o *Snapshot) SetMssqlInternalVersion(v int32)`

SetMssqlInternalVersion sets MssqlInternalVersion field to given value.

### HasMssqlInternalVersion

`func (o *Snapshot) HasMssqlInternalVersion() bool`

HasMssqlInternalVersion returns a boolean if a field has been set.

### GetMssqlBackupSetUuid

`func (o *Snapshot) GetMssqlBackupSetUuid() string`

GetMssqlBackupSetUuid returns the MssqlBackupSetUuid field if non-nil, zero value otherwise.

### GetMssqlBackupSetUuidOk

`func (o *Snapshot) GetMssqlBackupSetUuidOk() (*string, bool)`

GetMssqlBackupSetUuidOk returns a tuple with the MssqlBackupSetUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlBackupSetUuid

`func (o *Snapshot) SetMssqlBackupSetUuid(v string)`

SetMssqlBackupSetUuid sets MssqlBackupSetUuid field to given value.

### HasMssqlBackupSetUuid

`func (o *Snapshot) HasMssqlBackupSetUuid() bool`

HasMssqlBackupSetUuid returns a boolean if a field has been set.

### GetMssqlBackupSoftwareType

`func (o *Snapshot) GetMssqlBackupSoftwareType() string`

GetMssqlBackupSoftwareType returns the MssqlBackupSoftwareType field if non-nil, zero value otherwise.

### GetMssqlBackupSoftwareTypeOk

`func (o *Snapshot) GetMssqlBackupSoftwareTypeOk() (*string, bool)`

GetMssqlBackupSoftwareTypeOk returns a tuple with the MssqlBackupSoftwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlBackupSoftwareType

`func (o *Snapshot) SetMssqlBackupSoftwareType(v string)`

SetMssqlBackupSoftwareType sets MssqlBackupSoftwareType field to given value.

### HasMssqlBackupSoftwareType

`func (o *Snapshot) HasMssqlBackupSoftwareType() bool`

HasMssqlBackupSoftwareType returns a boolean if a field has been set.

### GetMssqlBackupLocationType

`func (o *Snapshot) GetMssqlBackupLocationType() string`

GetMssqlBackupLocationType returns the MssqlBackupLocationType field if non-nil, zero value otherwise.

### GetMssqlBackupLocationTypeOk

`func (o *Snapshot) GetMssqlBackupLocationTypeOk() (*string, bool)`

GetMssqlBackupLocationTypeOk returns a tuple with the MssqlBackupLocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlBackupLocationType

`func (o *Snapshot) SetMssqlBackupLocationType(v string)`

SetMssqlBackupLocationType sets MssqlBackupLocationType field to given value.

### HasMssqlBackupLocationType

`func (o *Snapshot) HasMssqlBackupLocationType() bool`

HasMssqlBackupLocationType returns a boolean if a field has been set.

### GetMssqlEmptySnapshot

`func (o *Snapshot) GetMssqlEmptySnapshot() bool`

GetMssqlEmptySnapshot returns the MssqlEmptySnapshot field if non-nil, zero value otherwise.

### GetMssqlEmptySnapshotOk

`func (o *Snapshot) GetMssqlEmptySnapshotOk() (*bool, bool)`

GetMssqlEmptySnapshotOk returns a tuple with the MssqlEmptySnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlEmptySnapshot

`func (o *Snapshot) SetMssqlEmptySnapshot(v bool)`

SetMssqlEmptySnapshot sets MssqlEmptySnapshot field to given value.

### HasMssqlEmptySnapshot

`func (o *Snapshot) HasMssqlEmptySnapshot() bool`

HasMssqlEmptySnapshot returns a boolean if a field has been set.

### GetOracleFromPhysicalStandbyVdb

`func (o *Snapshot) GetOracleFromPhysicalStandbyVdb() bool`

GetOracleFromPhysicalStandbyVdb returns the OracleFromPhysicalStandbyVdb field if non-nil, zero value otherwise.

### GetOracleFromPhysicalStandbyVdbOk

`func (o *Snapshot) GetOracleFromPhysicalStandbyVdbOk() (*bool, bool)`

GetOracleFromPhysicalStandbyVdbOk returns a tuple with the OracleFromPhysicalStandbyVdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleFromPhysicalStandbyVdb

`func (o *Snapshot) SetOracleFromPhysicalStandbyVdb(v bool)`

SetOracleFromPhysicalStandbyVdb sets OracleFromPhysicalStandbyVdb field to given value.

### HasOracleFromPhysicalStandbyVdb

`func (o *Snapshot) HasOracleFromPhysicalStandbyVdb() bool`

HasOracleFromPhysicalStandbyVdb returns a boolean if a field has been set.

### GetOracleRedoLogSizeInBytes

`func (o *Snapshot) GetOracleRedoLogSizeInBytes() int64`

GetOracleRedoLogSizeInBytes returns the OracleRedoLogSizeInBytes field if non-nil, zero value otherwise.

### GetOracleRedoLogSizeInBytesOk

`func (o *Snapshot) GetOracleRedoLogSizeInBytesOk() (*int64, bool)`

GetOracleRedoLogSizeInBytesOk returns a tuple with the OracleRedoLogSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleRedoLogSizeInBytes

`func (o *Snapshot) SetOracleRedoLogSizeInBytes(v int64)`

SetOracleRedoLogSizeInBytes sets OracleRedoLogSizeInBytes field to given value.

### HasOracleRedoLogSizeInBytes

`func (o *Snapshot) HasOracleRedoLogSizeInBytes() bool`

HasOracleRedoLogSizeInBytes returns a boolean if a field has been set.

### GetTags

`func (o *Snapshot) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Snapshot) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Snapshot) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Snapshot) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


