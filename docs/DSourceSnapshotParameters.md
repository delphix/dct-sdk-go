# DSourceSnapshotParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DropAndRecreateDevices** | Pointer to **bool** | If this parameter is set to true, older devices will be dropped and new devices created instead of trying to remap the devices. This might increase the space utilization on Delphix Engine. (ASE only)  | [optional] 
**SyncStrategy** | Pointer to **string** | Determines how the Delphix Engine will take a backup: * &#x60;latest_backup&#x60; - Use the most recent backup. * &#x60;new_backup&#x60; - Delphix will take a new backup of your source database. * &#x60;specific_backup&#x60; - Use a specific backup. Using this option requires setting   &#x60;ase_backup_files&#x60; for ASE dSources or &#x60;mssql_backup_uuid&#x60; for MSSql dSources. Default is &#x60;new_backup&#x60;. (ASE, MSSql only)  | [optional] 
**AseBackupFiles** | Pointer to **[]string** | When using the &#x60;specific_backup&#x60; sync_strategy, determines the backup files. (ASE Only) | [optional] 
**MssqlBackupUuid** | Pointer to **string** | When using the &#x60;specific_backup&#x60; sync_strategy, determines the Backup Set UUID. (MSSql only) | [optional] 
**CompressionEnabled** | Pointer to **bool** | When using the &#x60;new_backup&#x60; sync_strategy, determines if compression must be enabled. Defaults to the configuration of the ingestion strategy configured on the Delphix Engine for this dSource. (MSSql only) | [optional] 
**AvailabilityGroupBackupPolicy** | Pointer to **string** | When using the &#x60;new_backup&#x60; sync_strategy for an MSSql Availability Group, determines the backup policy: * &#x60;primary&#x60; - Backups only go to the primary node. * &#x60;secondary_only&#x60; - Backups only go to secondary nodes. If secondary nodes are down, backups will fail. * &#x60;prefer_secondary&#x60; - Backups go to secondary nodes, but if secondary nodes are down, backups will go to the primary node. (MSSql only)  | [optional] 
**DoNotResume** | Pointer to **bool** | Indicates whether a fresh SnapSync must be started regardless if it was possible to resume the current SnapSync. If true, we will not resume but instead ignore previous progress and backup all datafiles even if already completed from previous failed SnapSync. This does not force a full backup, if an incremental was in progress this will start a new incremental snapshot. (Oracle only)  | [optional] 
**DoubleSync** | Pointer to **bool** | Indicates whether two SnapSyncs should be performed in immediate succession to reduce the number of logs required to provision the snapshot. This may significantly reduce the time necessary to provision from a snapshot. (Oracle only).  | [optional] 
**ForceFullBackup** | Pointer to **bool** | Whether or not to take another full backup of the source database. (Oracle only) | [optional] 
**SkipSpaceCheck** | Pointer to **bool** | Skip check that tests if there is enough space available to store the database in the Delphix Engine. The Delphix Engine estimates how much space a database will occupy after compression and prevents SnapSync if insufficient space is available. This safeguard can be overridden using this option. This may be useful when linking highly compressible databases. (Oracle only)  | [optional] 
**FilesForPartialFullBackup** | Pointer to **[]int64** | List of datafiles to take a full backup of. This would be useful in situations where certain datafiles could not be backed up during previous SnapSync due to corruption or because they went offline. (Oracle only)  | [optional] 

## Methods

### NewDSourceSnapshotParameters

`func NewDSourceSnapshotParameters() *DSourceSnapshotParameters`

NewDSourceSnapshotParameters instantiates a new DSourceSnapshotParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDSourceSnapshotParametersWithDefaults

`func NewDSourceSnapshotParametersWithDefaults() *DSourceSnapshotParameters`

NewDSourceSnapshotParametersWithDefaults instantiates a new DSourceSnapshotParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDropAndRecreateDevices

`func (o *DSourceSnapshotParameters) GetDropAndRecreateDevices() bool`

GetDropAndRecreateDevices returns the DropAndRecreateDevices field if non-nil, zero value otherwise.

### GetDropAndRecreateDevicesOk

`func (o *DSourceSnapshotParameters) GetDropAndRecreateDevicesOk() (*bool, bool)`

GetDropAndRecreateDevicesOk returns a tuple with the DropAndRecreateDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropAndRecreateDevices

`func (o *DSourceSnapshotParameters) SetDropAndRecreateDevices(v bool)`

SetDropAndRecreateDevices sets DropAndRecreateDevices field to given value.

### HasDropAndRecreateDevices

`func (o *DSourceSnapshotParameters) HasDropAndRecreateDevices() bool`

HasDropAndRecreateDevices returns a boolean if a field has been set.

### GetSyncStrategy

`func (o *DSourceSnapshotParameters) GetSyncStrategy() string`

GetSyncStrategy returns the SyncStrategy field if non-nil, zero value otherwise.

### GetSyncStrategyOk

`func (o *DSourceSnapshotParameters) GetSyncStrategyOk() (*string, bool)`

GetSyncStrategyOk returns a tuple with the SyncStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStrategy

`func (o *DSourceSnapshotParameters) SetSyncStrategy(v string)`

SetSyncStrategy sets SyncStrategy field to given value.

### HasSyncStrategy

`func (o *DSourceSnapshotParameters) HasSyncStrategy() bool`

HasSyncStrategy returns a boolean if a field has been set.

### GetAseBackupFiles

`func (o *DSourceSnapshotParameters) GetAseBackupFiles() []string`

GetAseBackupFiles returns the AseBackupFiles field if non-nil, zero value otherwise.

### GetAseBackupFilesOk

`func (o *DSourceSnapshotParameters) GetAseBackupFilesOk() (*[]string, bool)`

GetAseBackupFilesOk returns a tuple with the AseBackupFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseBackupFiles

`func (o *DSourceSnapshotParameters) SetAseBackupFiles(v []string)`

SetAseBackupFiles sets AseBackupFiles field to given value.

### HasAseBackupFiles

`func (o *DSourceSnapshotParameters) HasAseBackupFiles() bool`

HasAseBackupFiles returns a boolean if a field has been set.

### GetMssqlBackupUuid

`func (o *DSourceSnapshotParameters) GetMssqlBackupUuid() string`

GetMssqlBackupUuid returns the MssqlBackupUuid field if non-nil, zero value otherwise.

### GetMssqlBackupUuidOk

`func (o *DSourceSnapshotParameters) GetMssqlBackupUuidOk() (*string, bool)`

GetMssqlBackupUuidOk returns a tuple with the MssqlBackupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlBackupUuid

`func (o *DSourceSnapshotParameters) SetMssqlBackupUuid(v string)`

SetMssqlBackupUuid sets MssqlBackupUuid field to given value.

### HasMssqlBackupUuid

`func (o *DSourceSnapshotParameters) HasMssqlBackupUuid() bool`

HasMssqlBackupUuid returns a boolean if a field has been set.

### GetCompressionEnabled

`func (o *DSourceSnapshotParameters) GetCompressionEnabled() bool`

GetCompressionEnabled returns the CompressionEnabled field if non-nil, zero value otherwise.

### GetCompressionEnabledOk

`func (o *DSourceSnapshotParameters) GetCompressionEnabledOk() (*bool, bool)`

GetCompressionEnabledOk returns a tuple with the CompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionEnabled

`func (o *DSourceSnapshotParameters) SetCompressionEnabled(v bool)`

SetCompressionEnabled sets CompressionEnabled field to given value.

### HasCompressionEnabled

`func (o *DSourceSnapshotParameters) HasCompressionEnabled() bool`

HasCompressionEnabled returns a boolean if a field has been set.

### GetAvailabilityGroupBackupPolicy

`func (o *DSourceSnapshotParameters) GetAvailabilityGroupBackupPolicy() string`

GetAvailabilityGroupBackupPolicy returns the AvailabilityGroupBackupPolicy field if non-nil, zero value otherwise.

### GetAvailabilityGroupBackupPolicyOk

`func (o *DSourceSnapshotParameters) GetAvailabilityGroupBackupPolicyOk() (*string, bool)`

GetAvailabilityGroupBackupPolicyOk returns a tuple with the AvailabilityGroupBackupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityGroupBackupPolicy

`func (o *DSourceSnapshotParameters) SetAvailabilityGroupBackupPolicy(v string)`

SetAvailabilityGroupBackupPolicy sets AvailabilityGroupBackupPolicy field to given value.

### HasAvailabilityGroupBackupPolicy

`func (o *DSourceSnapshotParameters) HasAvailabilityGroupBackupPolicy() bool`

HasAvailabilityGroupBackupPolicy returns a boolean if a field has been set.

### GetDoNotResume

`func (o *DSourceSnapshotParameters) GetDoNotResume() bool`

GetDoNotResume returns the DoNotResume field if non-nil, zero value otherwise.

### GetDoNotResumeOk

`func (o *DSourceSnapshotParameters) GetDoNotResumeOk() (*bool, bool)`

GetDoNotResumeOk returns a tuple with the DoNotResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotResume

`func (o *DSourceSnapshotParameters) SetDoNotResume(v bool)`

SetDoNotResume sets DoNotResume field to given value.

### HasDoNotResume

`func (o *DSourceSnapshotParameters) HasDoNotResume() bool`

HasDoNotResume returns a boolean if a field has been set.

### GetDoubleSync

`func (o *DSourceSnapshotParameters) GetDoubleSync() bool`

GetDoubleSync returns the DoubleSync field if non-nil, zero value otherwise.

### GetDoubleSyncOk

`func (o *DSourceSnapshotParameters) GetDoubleSyncOk() (*bool, bool)`

GetDoubleSyncOk returns a tuple with the DoubleSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleSync

`func (o *DSourceSnapshotParameters) SetDoubleSync(v bool)`

SetDoubleSync sets DoubleSync field to given value.

### HasDoubleSync

`func (o *DSourceSnapshotParameters) HasDoubleSync() bool`

HasDoubleSync returns a boolean if a field has been set.

### GetForceFullBackup

`func (o *DSourceSnapshotParameters) GetForceFullBackup() bool`

GetForceFullBackup returns the ForceFullBackup field if non-nil, zero value otherwise.

### GetForceFullBackupOk

`func (o *DSourceSnapshotParameters) GetForceFullBackupOk() (*bool, bool)`

GetForceFullBackupOk returns a tuple with the ForceFullBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceFullBackup

`func (o *DSourceSnapshotParameters) SetForceFullBackup(v bool)`

SetForceFullBackup sets ForceFullBackup field to given value.

### HasForceFullBackup

`func (o *DSourceSnapshotParameters) HasForceFullBackup() bool`

HasForceFullBackup returns a boolean if a field has been set.

### GetSkipSpaceCheck

`func (o *DSourceSnapshotParameters) GetSkipSpaceCheck() bool`

GetSkipSpaceCheck returns the SkipSpaceCheck field if non-nil, zero value otherwise.

### GetSkipSpaceCheckOk

`func (o *DSourceSnapshotParameters) GetSkipSpaceCheckOk() (*bool, bool)`

GetSkipSpaceCheckOk returns a tuple with the SkipSpaceCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSpaceCheck

`func (o *DSourceSnapshotParameters) SetSkipSpaceCheck(v bool)`

SetSkipSpaceCheck sets SkipSpaceCheck field to given value.

### HasSkipSpaceCheck

`func (o *DSourceSnapshotParameters) HasSkipSpaceCheck() bool`

HasSkipSpaceCheck returns a boolean if a field has been set.

### GetFilesForPartialFullBackup

`func (o *DSourceSnapshotParameters) GetFilesForPartialFullBackup() []int64`

GetFilesForPartialFullBackup returns the FilesForPartialFullBackup field if non-nil, zero value otherwise.

### GetFilesForPartialFullBackupOk

`func (o *DSourceSnapshotParameters) GetFilesForPartialFullBackupOk() (*[]int64, bool)`

GetFilesForPartialFullBackupOk returns a tuple with the FilesForPartialFullBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesForPartialFullBackup

`func (o *DSourceSnapshotParameters) SetFilesForPartialFullBackup(v []int64)`

SetFilesForPartialFullBackup sets FilesForPartialFullBackup field to given value.

### HasFilesForPartialFullBackup

`func (o *DSourceSnapshotParameters) HasFilesForPartialFullBackup() bool`

HasFilesForPartialFullBackup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


