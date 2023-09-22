# OracleDSourceLinkSourceParametersAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalFilePath** | Pointer to **string** | External file path. | [optional] 
**EnvironmentUserId** | Pointer to **string** | Id of the environment user to use for linking. | [optional] 
**BackupLevelEnabled** | Pointer to **bool** | Boolean value indicates whether LEVEL-based incremental backups can be used on the source database. | [optional] 
**RmanChannels** | Pointer to **int32** | Number of parallel channels to use. | [optional] [default to 2]
**FilesPerSet** | Pointer to **int32** | Number of data files to include in each RMAN backup set. | [optional] [default to 5]
**CheckLogical** | Pointer to **bool** | True if extended block checking should be used for this linked database. | [optional] [default to false]
**EncryptedLinkingEnabled** | Pointer to **bool** | True if SnapSync data from the source should be retrieved through an encrypted connection. Enabling this feature can decrease the performance of SnapSync from the source but has no impact on the performance of VDBs created from the retrieved data. | [optional] [default to false]
**CompressedLinkingEnabled** | Pointer to **bool** | True if SnapSync data from the source should be compressed over the network. Enabling this feature will reduce network bandwidth consumption and may significantly improve throughput, especially over slow network. | [optional] [default to true]
**BandwidthLimit** | Pointer to **int32** | Bandwidth limit (MB/s) for SnapSync and LogSync network traffic. A value of 0 means no limit. | [optional] [default to 0]
**NumberOfConnections** | Pointer to **int32** | Total number of transport connections to use during SnapSync. | [optional] [default to 1]
**DiagnoseNoLoggingFaults** | Pointer to **bool** | If true, NOLOGGING operations on this container are treated as faults and cannot be resolved manually. | [optional] [default to true]
**PreProvisioningEnabled** | Pointer to **bool** | If true, pre-provisioning will be performed after every sync. | [optional] [default to false]
**LinkNow** | Pointer to **bool** | True if initial load should be done immediately. | [optional] [default to false]
**ForceFullBackup** | Pointer to **bool** | Whether or not to take another full backup of the source database. | [optional] [default to false]
**DoubleSync** | Pointer to **bool** | True if two SnapSyncs should be performed in immediate succession to reduce the number of logs required to provision the snapshot. This may significantly reduce the time necessary to provision from a snapshot. | [optional] [default to false]
**SkipSpaceCheck** | Pointer to **bool** | Skip check that tests if there is enough space available to store the database in the Delphix Engine. The Delphix Engine estimates how much space a database will occupy after compression and prevents SnapSync if insufficient space is available. This safeguard can be overridden using this option. This may be useful when linking highly compressible databases. | [optional] [default to false]
**DoNotResume** | Pointer to **bool** | Indicates whether a fresh SnapSync must be started regardless if it was possible to resume the current SnapSync. If true, we will not resume but instead ignore previous progress and backup all datafiles even if already completed from previous failed SnapSync. This does not force a full backup, if an incremental was in progress this will start a new incremental snapshot. | [optional] [default to false]
**FilesForFullBackup** | Pointer to **[]int32** | List of datafiles to take a full backup of. This would be useful in situations where certain datafiles could not be backed up during previous SnapSync due to corruption or because they went offline. | [optional] 
**LogSyncMode** | Pointer to **string** | LogSync operation mode for this database. | [optional] [default to "UNDEFINED"]
**LogSyncInterval** | Pointer to **int32** | Interval between LogSync requests, in seconds. | [optional] [default to 5]
**NonSysUsername** | Pointer to **string** | Non-SYS database user to access this database. Only required for username-password auth (Single tenant only). | [optional] 
**NonSysPassword** | Pointer to **string** | Password for non sys user authentication (Single tenant only). | [optional] 
**NonSysVault** | Pointer to **string** | The name or reference of the vault from which to read the database credentials (Single tenant only). | [optional] 
**NonSysHashicorpVaultEngine** | Pointer to **string** | Vault engine name where the credential is stored (Single tenant only). | [optional] 
**NonSysHashicorpVaultSecretPath** | Pointer to **string** | Path in the vault engine where the credential is stored (Single tenant only). | [optional] 
**NonSysHashicorpVaultUsernameKey** | Pointer to **string** | Hashicorp vault key for the username in the key-value store (Single tenant only). | [optional] 
**NonSysHashicorpVaultSecretKey** | Pointer to **string** | Hashicorp vault key for the password in the key-value store (Single tenant only). | [optional] 
**NonSysAzureVaultName** | Pointer to **string** | Azure key vault name (Single tenant only). | [optional] 
**NonSysAzureVaultUsernameKey** | Pointer to **string** | Azure vault key for the username in the key-value store (Single tenant only). | [optional] 
**NonSysAzureVaultSecretKey** | Pointer to **string** | Azure vault key for the password in the key-value store (Single tenant only). | [optional] 
**NonSysCyberarkVaultQueryString** | Pointer to **string** | Query to find a credential in the CyberArk vault (Single tenant only). | [optional] 
**FallbackUsername** | Pointer to **string** | The database fallback username. Optional if bequeath connections are enabled (to be used in case of bequeath connection failures). Only required for username-password auth. | [optional] 
**FallbackPassword** | Pointer to **string** | Password for fallback username. | [optional] 
**FallbackVault** | Pointer to **string** | The name or reference of the vault from which to read the database credentials. | [optional] 
**FallbackHashicorpVaultEngine** | Pointer to **string** | Vault engine name where the credential is stored. | [optional] 
**FallbackHashicorpVaultSecretPath** | Pointer to **string** | Path in the vault engine where the credential is stored. | [optional] 
**FallbackHashicorpVaultUsernameKey** | Pointer to **string** | Hashicorp vault key for the username in the key-value store. | [optional] 
**FallbackHashicorpVaultSecretKey** | Pointer to **string** | Hashicorp vault key for the password in the key-value store. | [optional] 
**FallbackAzureVaultName** | Pointer to **string** | Azure key vault name. | [optional] 
**FallbackAzureVaultUsernameKey** | Pointer to **string** | Azure vault key for the username in the key-value store. | [optional] 
**FallbackAzureVaultSecretKey** | Pointer to **string** | Azure vault key for the password in the key-value store. | [optional] 
**FallbackCyberarkVaultQueryString** | Pointer to **string** | Query to find a credential in the CyberArk vault. | [optional] 
**OpsPreLogSync** | Pointer to [**[]SourceOperation**](SourceOperation.md) | Operations to perform after syncing a created dSource and before running the LogSync. | [optional] 

## Methods

### NewOracleDSourceLinkSourceParametersAllOf

`func NewOracleDSourceLinkSourceParametersAllOf() *OracleDSourceLinkSourceParametersAllOf`

NewOracleDSourceLinkSourceParametersAllOf instantiates a new OracleDSourceLinkSourceParametersAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleDSourceLinkSourceParametersAllOfWithDefaults

`func NewOracleDSourceLinkSourceParametersAllOfWithDefaults() *OracleDSourceLinkSourceParametersAllOf`

NewOracleDSourceLinkSourceParametersAllOfWithDefaults instantiates a new OracleDSourceLinkSourceParametersAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalFilePath

`func (o *OracleDSourceLinkSourceParametersAllOf) GetExternalFilePath() string`

GetExternalFilePath returns the ExternalFilePath field if non-nil, zero value otherwise.

### GetExternalFilePathOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetExternalFilePathOk() (*string, bool)`

GetExternalFilePathOk returns a tuple with the ExternalFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFilePath

`func (o *OracleDSourceLinkSourceParametersAllOf) SetExternalFilePath(v string)`

SetExternalFilePath sets ExternalFilePath field to given value.

### HasExternalFilePath

`func (o *OracleDSourceLinkSourceParametersAllOf) HasExternalFilePath() bool`

HasExternalFilePath returns a boolean if a field has been set.

### GetEnvironmentUserId

`func (o *OracleDSourceLinkSourceParametersAllOf) GetEnvironmentUserId() string`

GetEnvironmentUserId returns the EnvironmentUserId field if non-nil, zero value otherwise.

### GetEnvironmentUserIdOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetEnvironmentUserIdOk() (*string, bool)`

GetEnvironmentUserIdOk returns a tuple with the EnvironmentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUserId

`func (o *OracleDSourceLinkSourceParametersAllOf) SetEnvironmentUserId(v string)`

SetEnvironmentUserId sets EnvironmentUserId field to given value.

### HasEnvironmentUserId

`func (o *OracleDSourceLinkSourceParametersAllOf) HasEnvironmentUserId() bool`

HasEnvironmentUserId returns a boolean if a field has been set.

### GetBackupLevelEnabled

`func (o *OracleDSourceLinkSourceParametersAllOf) GetBackupLevelEnabled() bool`

GetBackupLevelEnabled returns the BackupLevelEnabled field if non-nil, zero value otherwise.

### GetBackupLevelEnabledOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetBackupLevelEnabledOk() (*bool, bool)`

GetBackupLevelEnabledOk returns a tuple with the BackupLevelEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLevelEnabled

`func (o *OracleDSourceLinkSourceParametersAllOf) SetBackupLevelEnabled(v bool)`

SetBackupLevelEnabled sets BackupLevelEnabled field to given value.

### HasBackupLevelEnabled

`func (o *OracleDSourceLinkSourceParametersAllOf) HasBackupLevelEnabled() bool`

HasBackupLevelEnabled returns a boolean if a field has been set.

### GetRmanChannels

`func (o *OracleDSourceLinkSourceParametersAllOf) GetRmanChannels() int32`

GetRmanChannels returns the RmanChannels field if non-nil, zero value otherwise.

### GetRmanChannelsOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetRmanChannelsOk() (*int32, bool)`

GetRmanChannelsOk returns a tuple with the RmanChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmanChannels

`func (o *OracleDSourceLinkSourceParametersAllOf) SetRmanChannels(v int32)`

SetRmanChannels sets RmanChannels field to given value.

### HasRmanChannels

`func (o *OracleDSourceLinkSourceParametersAllOf) HasRmanChannels() bool`

HasRmanChannels returns a boolean if a field has been set.

### GetFilesPerSet

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFilesPerSet() int32`

GetFilesPerSet returns the FilesPerSet field if non-nil, zero value otherwise.

### GetFilesPerSetOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFilesPerSetOk() (*int32, bool)`

GetFilesPerSetOk returns a tuple with the FilesPerSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesPerSet

`func (o *OracleDSourceLinkSourceParametersAllOf) SetFilesPerSet(v int32)`

SetFilesPerSet sets FilesPerSet field to given value.

### HasFilesPerSet

`func (o *OracleDSourceLinkSourceParametersAllOf) HasFilesPerSet() bool`

HasFilesPerSet returns a boolean if a field has been set.

### GetCheckLogical

`func (o *OracleDSourceLinkSourceParametersAllOf) GetCheckLogical() bool`

GetCheckLogical returns the CheckLogical field if non-nil, zero value otherwise.

### GetCheckLogicalOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetCheckLogicalOk() (*bool, bool)`

GetCheckLogicalOk returns a tuple with the CheckLogical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckLogical

`func (o *OracleDSourceLinkSourceParametersAllOf) SetCheckLogical(v bool)`

SetCheckLogical sets CheckLogical field to given value.

### HasCheckLogical

`func (o *OracleDSourceLinkSourceParametersAllOf) HasCheckLogical() bool`

HasCheckLogical returns a boolean if a field has been set.

### GetEncryptedLinkingEnabled

`func (o *OracleDSourceLinkSourceParametersAllOf) GetEncryptedLinkingEnabled() bool`

GetEncryptedLinkingEnabled returns the EncryptedLinkingEnabled field if non-nil, zero value otherwise.

### GetEncryptedLinkingEnabledOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetEncryptedLinkingEnabledOk() (*bool, bool)`

GetEncryptedLinkingEnabledOk returns a tuple with the EncryptedLinkingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedLinkingEnabled

`func (o *OracleDSourceLinkSourceParametersAllOf) SetEncryptedLinkingEnabled(v bool)`

SetEncryptedLinkingEnabled sets EncryptedLinkingEnabled field to given value.

### HasEncryptedLinkingEnabled

`func (o *OracleDSourceLinkSourceParametersAllOf) HasEncryptedLinkingEnabled() bool`

HasEncryptedLinkingEnabled returns a boolean if a field has been set.

### GetCompressedLinkingEnabled

`func (o *OracleDSourceLinkSourceParametersAllOf) GetCompressedLinkingEnabled() bool`

GetCompressedLinkingEnabled returns the CompressedLinkingEnabled field if non-nil, zero value otherwise.

### GetCompressedLinkingEnabledOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetCompressedLinkingEnabledOk() (*bool, bool)`

GetCompressedLinkingEnabledOk returns a tuple with the CompressedLinkingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedLinkingEnabled

`func (o *OracleDSourceLinkSourceParametersAllOf) SetCompressedLinkingEnabled(v bool)`

SetCompressedLinkingEnabled sets CompressedLinkingEnabled field to given value.

### HasCompressedLinkingEnabled

`func (o *OracleDSourceLinkSourceParametersAllOf) HasCompressedLinkingEnabled() bool`

HasCompressedLinkingEnabled returns a boolean if a field has been set.

### GetBandwidthLimit

`func (o *OracleDSourceLinkSourceParametersAllOf) GetBandwidthLimit() int32`

GetBandwidthLimit returns the BandwidthLimit field if non-nil, zero value otherwise.

### GetBandwidthLimitOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetBandwidthLimitOk() (*int32, bool)`

GetBandwidthLimitOk returns a tuple with the BandwidthLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimit

`func (o *OracleDSourceLinkSourceParametersAllOf) SetBandwidthLimit(v int32)`

SetBandwidthLimit sets BandwidthLimit field to given value.

### HasBandwidthLimit

`func (o *OracleDSourceLinkSourceParametersAllOf) HasBandwidthLimit() bool`

HasBandwidthLimit returns a boolean if a field has been set.

### GetNumberOfConnections

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNumberOfConnections() int32`

GetNumberOfConnections returns the NumberOfConnections field if non-nil, zero value otherwise.

### GetNumberOfConnectionsOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNumberOfConnectionsOk() (*int32, bool)`

GetNumberOfConnectionsOk returns a tuple with the NumberOfConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfConnections

`func (o *OracleDSourceLinkSourceParametersAllOf) SetNumberOfConnections(v int32)`

SetNumberOfConnections sets NumberOfConnections field to given value.

### HasNumberOfConnections

`func (o *OracleDSourceLinkSourceParametersAllOf) HasNumberOfConnections() bool`

HasNumberOfConnections returns a boolean if a field has been set.

### GetDiagnoseNoLoggingFaults

`func (o *OracleDSourceLinkSourceParametersAllOf) GetDiagnoseNoLoggingFaults() bool`

GetDiagnoseNoLoggingFaults returns the DiagnoseNoLoggingFaults field if non-nil, zero value otherwise.

### GetDiagnoseNoLoggingFaultsOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetDiagnoseNoLoggingFaultsOk() (*bool, bool)`

GetDiagnoseNoLoggingFaultsOk returns a tuple with the DiagnoseNoLoggingFaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnoseNoLoggingFaults

`func (o *OracleDSourceLinkSourceParametersAllOf) SetDiagnoseNoLoggingFaults(v bool)`

SetDiagnoseNoLoggingFaults sets DiagnoseNoLoggingFaults field to given value.

### HasDiagnoseNoLoggingFaults

`func (o *OracleDSourceLinkSourceParametersAllOf) HasDiagnoseNoLoggingFaults() bool`

HasDiagnoseNoLoggingFaults returns a boolean if a field has been set.

### GetPreProvisioningEnabled

`func (o *OracleDSourceLinkSourceParametersAllOf) GetPreProvisioningEnabled() bool`

GetPreProvisioningEnabled returns the PreProvisioningEnabled field if non-nil, zero value otherwise.

### GetPreProvisioningEnabledOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetPreProvisioningEnabledOk() (*bool, bool)`

GetPreProvisioningEnabledOk returns a tuple with the PreProvisioningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreProvisioningEnabled

`func (o *OracleDSourceLinkSourceParametersAllOf) SetPreProvisioningEnabled(v bool)`

SetPreProvisioningEnabled sets PreProvisioningEnabled field to given value.

### HasPreProvisioningEnabled

`func (o *OracleDSourceLinkSourceParametersAllOf) HasPreProvisioningEnabled() bool`

HasPreProvisioningEnabled returns a boolean if a field has been set.

### GetLinkNow

`func (o *OracleDSourceLinkSourceParametersAllOf) GetLinkNow() bool`

GetLinkNow returns the LinkNow field if non-nil, zero value otherwise.

### GetLinkNowOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetLinkNowOk() (*bool, bool)`

GetLinkNowOk returns a tuple with the LinkNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNow

`func (o *OracleDSourceLinkSourceParametersAllOf) SetLinkNow(v bool)`

SetLinkNow sets LinkNow field to given value.

### HasLinkNow

`func (o *OracleDSourceLinkSourceParametersAllOf) HasLinkNow() bool`

HasLinkNow returns a boolean if a field has been set.

### GetForceFullBackup

`func (o *OracleDSourceLinkSourceParametersAllOf) GetForceFullBackup() bool`

GetForceFullBackup returns the ForceFullBackup field if non-nil, zero value otherwise.

### GetForceFullBackupOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetForceFullBackupOk() (*bool, bool)`

GetForceFullBackupOk returns a tuple with the ForceFullBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceFullBackup

`func (o *OracleDSourceLinkSourceParametersAllOf) SetForceFullBackup(v bool)`

SetForceFullBackup sets ForceFullBackup field to given value.

### HasForceFullBackup

`func (o *OracleDSourceLinkSourceParametersAllOf) HasForceFullBackup() bool`

HasForceFullBackup returns a boolean if a field has been set.

### GetDoubleSync

`func (o *OracleDSourceLinkSourceParametersAllOf) GetDoubleSync() bool`

GetDoubleSync returns the DoubleSync field if non-nil, zero value otherwise.

### GetDoubleSyncOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetDoubleSyncOk() (*bool, bool)`

GetDoubleSyncOk returns a tuple with the DoubleSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleSync

`func (o *OracleDSourceLinkSourceParametersAllOf) SetDoubleSync(v bool)`

SetDoubleSync sets DoubleSync field to given value.

### HasDoubleSync

`func (o *OracleDSourceLinkSourceParametersAllOf) HasDoubleSync() bool`

HasDoubleSync returns a boolean if a field has been set.

### GetSkipSpaceCheck

`func (o *OracleDSourceLinkSourceParametersAllOf) GetSkipSpaceCheck() bool`

GetSkipSpaceCheck returns the SkipSpaceCheck field if non-nil, zero value otherwise.

### GetSkipSpaceCheckOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetSkipSpaceCheckOk() (*bool, bool)`

GetSkipSpaceCheckOk returns a tuple with the SkipSpaceCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSpaceCheck

`func (o *OracleDSourceLinkSourceParametersAllOf) SetSkipSpaceCheck(v bool)`

SetSkipSpaceCheck sets SkipSpaceCheck field to given value.

### HasSkipSpaceCheck

`func (o *OracleDSourceLinkSourceParametersAllOf) HasSkipSpaceCheck() bool`

HasSkipSpaceCheck returns a boolean if a field has been set.

### GetDoNotResume

`func (o *OracleDSourceLinkSourceParametersAllOf) GetDoNotResume() bool`

GetDoNotResume returns the DoNotResume field if non-nil, zero value otherwise.

### GetDoNotResumeOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetDoNotResumeOk() (*bool, bool)`

GetDoNotResumeOk returns a tuple with the DoNotResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotResume

`func (o *OracleDSourceLinkSourceParametersAllOf) SetDoNotResume(v bool)`

SetDoNotResume sets DoNotResume field to given value.

### HasDoNotResume

`func (o *OracleDSourceLinkSourceParametersAllOf) HasDoNotResume() bool`

HasDoNotResume returns a boolean if a field has been set.

### GetFilesForFullBackup

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFilesForFullBackup() []int32`

GetFilesForFullBackup returns the FilesForFullBackup field if non-nil, zero value otherwise.

### GetFilesForFullBackupOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFilesForFullBackupOk() (*[]int32, bool)`

GetFilesForFullBackupOk returns a tuple with the FilesForFullBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesForFullBackup

`func (o *OracleDSourceLinkSourceParametersAllOf) SetFilesForFullBackup(v []int32)`

SetFilesForFullBackup sets FilesForFullBackup field to given value.

### HasFilesForFullBackup

`func (o *OracleDSourceLinkSourceParametersAllOf) HasFilesForFullBackup() bool`

HasFilesForFullBackup returns a boolean if a field has been set.

### GetLogSyncMode

`func (o *OracleDSourceLinkSourceParametersAllOf) GetLogSyncMode() string`

GetLogSyncMode returns the LogSyncMode field if non-nil, zero value otherwise.

### GetLogSyncModeOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetLogSyncModeOk() (*string, bool)`

GetLogSyncModeOk returns a tuple with the LogSyncMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSyncMode

`func (o *OracleDSourceLinkSourceParametersAllOf) SetLogSyncMode(v string)`

SetLogSyncMode sets LogSyncMode field to given value.

### HasLogSyncMode

`func (o *OracleDSourceLinkSourceParametersAllOf) HasLogSyncMode() bool`

HasLogSyncMode returns a boolean if a field has been set.

### GetLogSyncInterval

`func (o *OracleDSourceLinkSourceParametersAllOf) GetLogSyncInterval() int32`

GetLogSyncInterval returns the LogSyncInterval field if non-nil, zero value otherwise.

### GetLogSyncIntervalOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetLogSyncIntervalOk() (*int32, bool)`

GetLogSyncIntervalOk returns a tuple with the LogSyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSyncInterval

`func (o *OracleDSourceLinkSourceParametersAllOf) SetLogSyncInterval(v int32)`

SetLogSyncInterval sets LogSyncInterval field to given value.

### HasLogSyncInterval

`func (o *OracleDSourceLinkSourceParametersAllOf) HasLogSyncInterval() bool`

HasLogSyncInterval returns a boolean if a field has been set.

### GetNonSysUsername

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysUsername() string`

GetNonSysUsername returns the NonSysUsername field if non-nil, zero value otherwise.

### GetNonSysUsernameOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysUsernameOk() (*string, bool)`

GetNonSysUsernameOk returns a tuple with the NonSysUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysUsername

`func (o *OracleDSourceLinkSourceParametersAllOf) SetNonSysUsername(v string)`

SetNonSysUsername sets NonSysUsername field to given value.

### HasNonSysUsername

`func (o *OracleDSourceLinkSourceParametersAllOf) HasNonSysUsername() bool`

HasNonSysUsername returns a boolean if a field has been set.

### GetNonSysPassword

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysPassword() string`

GetNonSysPassword returns the NonSysPassword field if non-nil, zero value otherwise.

### GetNonSysPasswordOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysPasswordOk() (*string, bool)`

GetNonSysPasswordOk returns a tuple with the NonSysPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysPassword

`func (o *OracleDSourceLinkSourceParametersAllOf) SetNonSysPassword(v string)`

SetNonSysPassword sets NonSysPassword field to given value.

### HasNonSysPassword

`func (o *OracleDSourceLinkSourceParametersAllOf) HasNonSysPassword() bool`

HasNonSysPassword returns a boolean if a field has been set.

### GetNonSysVault

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysVault() string`

GetNonSysVault returns the NonSysVault field if non-nil, zero value otherwise.

### GetNonSysVaultOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysVaultOk() (*string, bool)`

GetNonSysVaultOk returns a tuple with the NonSysVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysVault

`func (o *OracleDSourceLinkSourceParametersAllOf) SetNonSysVault(v string)`

SetNonSysVault sets NonSysVault field to given value.

### HasNonSysVault

`func (o *OracleDSourceLinkSourceParametersAllOf) HasNonSysVault() bool`

HasNonSysVault returns a boolean if a field has been set.

### GetNonSysHashicorpVaultEngine

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysHashicorpVaultEngine() string`

GetNonSysHashicorpVaultEngine returns the NonSysHashicorpVaultEngine field if non-nil, zero value otherwise.

### GetNonSysHashicorpVaultEngineOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysHashicorpVaultEngineOk() (*string, bool)`

GetNonSysHashicorpVaultEngineOk returns a tuple with the NonSysHashicorpVaultEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysHashicorpVaultEngine

`func (o *OracleDSourceLinkSourceParametersAllOf) SetNonSysHashicorpVaultEngine(v string)`

SetNonSysHashicorpVaultEngine sets NonSysHashicorpVaultEngine field to given value.

### HasNonSysHashicorpVaultEngine

`func (o *OracleDSourceLinkSourceParametersAllOf) HasNonSysHashicorpVaultEngine() bool`

HasNonSysHashicorpVaultEngine returns a boolean if a field has been set.

### GetNonSysHashicorpVaultSecretPath

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysHashicorpVaultSecretPath() string`

GetNonSysHashicorpVaultSecretPath returns the NonSysHashicorpVaultSecretPath field if non-nil, zero value otherwise.

### GetNonSysHashicorpVaultSecretPathOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysHashicorpVaultSecretPathOk() (*string, bool)`

GetNonSysHashicorpVaultSecretPathOk returns a tuple with the NonSysHashicorpVaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysHashicorpVaultSecretPath

`func (o *OracleDSourceLinkSourceParametersAllOf) SetNonSysHashicorpVaultSecretPath(v string)`

SetNonSysHashicorpVaultSecretPath sets NonSysHashicorpVaultSecretPath field to given value.

### HasNonSysHashicorpVaultSecretPath

`func (o *OracleDSourceLinkSourceParametersAllOf) HasNonSysHashicorpVaultSecretPath() bool`

HasNonSysHashicorpVaultSecretPath returns a boolean if a field has been set.

### GetNonSysHashicorpVaultUsernameKey

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysHashicorpVaultUsernameKey() string`

GetNonSysHashicorpVaultUsernameKey returns the NonSysHashicorpVaultUsernameKey field if non-nil, zero value otherwise.

### GetNonSysHashicorpVaultUsernameKeyOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysHashicorpVaultUsernameKeyOk() (*string, bool)`

GetNonSysHashicorpVaultUsernameKeyOk returns a tuple with the NonSysHashicorpVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysHashicorpVaultUsernameKey

`func (o *OracleDSourceLinkSourceParametersAllOf) SetNonSysHashicorpVaultUsernameKey(v string)`

SetNonSysHashicorpVaultUsernameKey sets NonSysHashicorpVaultUsernameKey field to given value.

### HasNonSysHashicorpVaultUsernameKey

`func (o *OracleDSourceLinkSourceParametersAllOf) HasNonSysHashicorpVaultUsernameKey() bool`

HasNonSysHashicorpVaultUsernameKey returns a boolean if a field has been set.

### GetNonSysHashicorpVaultSecretKey

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysHashicorpVaultSecretKey() string`

GetNonSysHashicorpVaultSecretKey returns the NonSysHashicorpVaultSecretKey field if non-nil, zero value otherwise.

### GetNonSysHashicorpVaultSecretKeyOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysHashicorpVaultSecretKeyOk() (*string, bool)`

GetNonSysHashicorpVaultSecretKeyOk returns a tuple with the NonSysHashicorpVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysHashicorpVaultSecretKey

`func (o *OracleDSourceLinkSourceParametersAllOf) SetNonSysHashicorpVaultSecretKey(v string)`

SetNonSysHashicorpVaultSecretKey sets NonSysHashicorpVaultSecretKey field to given value.

### HasNonSysHashicorpVaultSecretKey

`func (o *OracleDSourceLinkSourceParametersAllOf) HasNonSysHashicorpVaultSecretKey() bool`

HasNonSysHashicorpVaultSecretKey returns a boolean if a field has been set.

### GetNonSysAzureVaultName

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysAzureVaultName() string`

GetNonSysAzureVaultName returns the NonSysAzureVaultName field if non-nil, zero value otherwise.

### GetNonSysAzureVaultNameOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysAzureVaultNameOk() (*string, bool)`

GetNonSysAzureVaultNameOk returns a tuple with the NonSysAzureVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysAzureVaultName

`func (o *OracleDSourceLinkSourceParametersAllOf) SetNonSysAzureVaultName(v string)`

SetNonSysAzureVaultName sets NonSysAzureVaultName field to given value.

### HasNonSysAzureVaultName

`func (o *OracleDSourceLinkSourceParametersAllOf) HasNonSysAzureVaultName() bool`

HasNonSysAzureVaultName returns a boolean if a field has been set.

### GetNonSysAzureVaultUsernameKey

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysAzureVaultUsernameKey() string`

GetNonSysAzureVaultUsernameKey returns the NonSysAzureVaultUsernameKey field if non-nil, zero value otherwise.

### GetNonSysAzureVaultUsernameKeyOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysAzureVaultUsernameKeyOk() (*string, bool)`

GetNonSysAzureVaultUsernameKeyOk returns a tuple with the NonSysAzureVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysAzureVaultUsernameKey

`func (o *OracleDSourceLinkSourceParametersAllOf) SetNonSysAzureVaultUsernameKey(v string)`

SetNonSysAzureVaultUsernameKey sets NonSysAzureVaultUsernameKey field to given value.

### HasNonSysAzureVaultUsernameKey

`func (o *OracleDSourceLinkSourceParametersAllOf) HasNonSysAzureVaultUsernameKey() bool`

HasNonSysAzureVaultUsernameKey returns a boolean if a field has been set.

### GetNonSysAzureVaultSecretKey

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysAzureVaultSecretKey() string`

GetNonSysAzureVaultSecretKey returns the NonSysAzureVaultSecretKey field if non-nil, zero value otherwise.

### GetNonSysAzureVaultSecretKeyOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysAzureVaultSecretKeyOk() (*string, bool)`

GetNonSysAzureVaultSecretKeyOk returns a tuple with the NonSysAzureVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysAzureVaultSecretKey

`func (o *OracleDSourceLinkSourceParametersAllOf) SetNonSysAzureVaultSecretKey(v string)`

SetNonSysAzureVaultSecretKey sets NonSysAzureVaultSecretKey field to given value.

### HasNonSysAzureVaultSecretKey

`func (o *OracleDSourceLinkSourceParametersAllOf) HasNonSysAzureVaultSecretKey() bool`

HasNonSysAzureVaultSecretKey returns a boolean if a field has been set.

### GetNonSysCyberarkVaultQueryString

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysCyberarkVaultQueryString() string`

GetNonSysCyberarkVaultQueryString returns the NonSysCyberarkVaultQueryString field if non-nil, zero value otherwise.

### GetNonSysCyberarkVaultQueryStringOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetNonSysCyberarkVaultQueryStringOk() (*string, bool)`

GetNonSysCyberarkVaultQueryStringOk returns a tuple with the NonSysCyberarkVaultQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysCyberarkVaultQueryString

`func (o *OracleDSourceLinkSourceParametersAllOf) SetNonSysCyberarkVaultQueryString(v string)`

SetNonSysCyberarkVaultQueryString sets NonSysCyberarkVaultQueryString field to given value.

### HasNonSysCyberarkVaultQueryString

`func (o *OracleDSourceLinkSourceParametersAllOf) HasNonSysCyberarkVaultQueryString() bool`

HasNonSysCyberarkVaultQueryString returns a boolean if a field has been set.

### GetFallbackUsername

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackUsername() string`

GetFallbackUsername returns the FallbackUsername field if non-nil, zero value otherwise.

### GetFallbackUsernameOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackUsernameOk() (*string, bool)`

GetFallbackUsernameOk returns a tuple with the FallbackUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUsername

`func (o *OracleDSourceLinkSourceParametersAllOf) SetFallbackUsername(v string)`

SetFallbackUsername sets FallbackUsername field to given value.

### HasFallbackUsername

`func (o *OracleDSourceLinkSourceParametersAllOf) HasFallbackUsername() bool`

HasFallbackUsername returns a boolean if a field has been set.

### GetFallbackPassword

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackPassword() string`

GetFallbackPassword returns the FallbackPassword field if non-nil, zero value otherwise.

### GetFallbackPasswordOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackPasswordOk() (*string, bool)`

GetFallbackPasswordOk returns a tuple with the FallbackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackPassword

`func (o *OracleDSourceLinkSourceParametersAllOf) SetFallbackPassword(v string)`

SetFallbackPassword sets FallbackPassword field to given value.

### HasFallbackPassword

`func (o *OracleDSourceLinkSourceParametersAllOf) HasFallbackPassword() bool`

HasFallbackPassword returns a boolean if a field has been set.

### GetFallbackVault

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackVault() string`

GetFallbackVault returns the FallbackVault field if non-nil, zero value otherwise.

### GetFallbackVaultOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackVaultOk() (*string, bool)`

GetFallbackVaultOk returns a tuple with the FallbackVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackVault

`func (o *OracleDSourceLinkSourceParametersAllOf) SetFallbackVault(v string)`

SetFallbackVault sets FallbackVault field to given value.

### HasFallbackVault

`func (o *OracleDSourceLinkSourceParametersAllOf) HasFallbackVault() bool`

HasFallbackVault returns a boolean if a field has been set.

### GetFallbackHashicorpVaultEngine

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackHashicorpVaultEngine() string`

GetFallbackHashicorpVaultEngine returns the FallbackHashicorpVaultEngine field if non-nil, zero value otherwise.

### GetFallbackHashicorpVaultEngineOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackHashicorpVaultEngineOk() (*string, bool)`

GetFallbackHashicorpVaultEngineOk returns a tuple with the FallbackHashicorpVaultEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackHashicorpVaultEngine

`func (o *OracleDSourceLinkSourceParametersAllOf) SetFallbackHashicorpVaultEngine(v string)`

SetFallbackHashicorpVaultEngine sets FallbackHashicorpVaultEngine field to given value.

### HasFallbackHashicorpVaultEngine

`func (o *OracleDSourceLinkSourceParametersAllOf) HasFallbackHashicorpVaultEngine() bool`

HasFallbackHashicorpVaultEngine returns a boolean if a field has been set.

### GetFallbackHashicorpVaultSecretPath

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackHashicorpVaultSecretPath() string`

GetFallbackHashicorpVaultSecretPath returns the FallbackHashicorpVaultSecretPath field if non-nil, zero value otherwise.

### GetFallbackHashicorpVaultSecretPathOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackHashicorpVaultSecretPathOk() (*string, bool)`

GetFallbackHashicorpVaultSecretPathOk returns a tuple with the FallbackHashicorpVaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackHashicorpVaultSecretPath

`func (o *OracleDSourceLinkSourceParametersAllOf) SetFallbackHashicorpVaultSecretPath(v string)`

SetFallbackHashicorpVaultSecretPath sets FallbackHashicorpVaultSecretPath field to given value.

### HasFallbackHashicorpVaultSecretPath

`func (o *OracleDSourceLinkSourceParametersAllOf) HasFallbackHashicorpVaultSecretPath() bool`

HasFallbackHashicorpVaultSecretPath returns a boolean if a field has been set.

### GetFallbackHashicorpVaultUsernameKey

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackHashicorpVaultUsernameKey() string`

GetFallbackHashicorpVaultUsernameKey returns the FallbackHashicorpVaultUsernameKey field if non-nil, zero value otherwise.

### GetFallbackHashicorpVaultUsernameKeyOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackHashicorpVaultUsernameKeyOk() (*string, bool)`

GetFallbackHashicorpVaultUsernameKeyOk returns a tuple with the FallbackHashicorpVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackHashicorpVaultUsernameKey

`func (o *OracleDSourceLinkSourceParametersAllOf) SetFallbackHashicorpVaultUsernameKey(v string)`

SetFallbackHashicorpVaultUsernameKey sets FallbackHashicorpVaultUsernameKey field to given value.

### HasFallbackHashicorpVaultUsernameKey

`func (o *OracleDSourceLinkSourceParametersAllOf) HasFallbackHashicorpVaultUsernameKey() bool`

HasFallbackHashicorpVaultUsernameKey returns a boolean if a field has been set.

### GetFallbackHashicorpVaultSecretKey

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackHashicorpVaultSecretKey() string`

GetFallbackHashicorpVaultSecretKey returns the FallbackHashicorpVaultSecretKey field if non-nil, zero value otherwise.

### GetFallbackHashicorpVaultSecretKeyOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackHashicorpVaultSecretKeyOk() (*string, bool)`

GetFallbackHashicorpVaultSecretKeyOk returns a tuple with the FallbackHashicorpVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackHashicorpVaultSecretKey

`func (o *OracleDSourceLinkSourceParametersAllOf) SetFallbackHashicorpVaultSecretKey(v string)`

SetFallbackHashicorpVaultSecretKey sets FallbackHashicorpVaultSecretKey field to given value.

### HasFallbackHashicorpVaultSecretKey

`func (o *OracleDSourceLinkSourceParametersAllOf) HasFallbackHashicorpVaultSecretKey() bool`

HasFallbackHashicorpVaultSecretKey returns a boolean if a field has been set.

### GetFallbackAzureVaultName

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackAzureVaultName() string`

GetFallbackAzureVaultName returns the FallbackAzureVaultName field if non-nil, zero value otherwise.

### GetFallbackAzureVaultNameOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackAzureVaultNameOk() (*string, bool)`

GetFallbackAzureVaultNameOk returns a tuple with the FallbackAzureVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackAzureVaultName

`func (o *OracleDSourceLinkSourceParametersAllOf) SetFallbackAzureVaultName(v string)`

SetFallbackAzureVaultName sets FallbackAzureVaultName field to given value.

### HasFallbackAzureVaultName

`func (o *OracleDSourceLinkSourceParametersAllOf) HasFallbackAzureVaultName() bool`

HasFallbackAzureVaultName returns a boolean if a field has been set.

### GetFallbackAzureVaultUsernameKey

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackAzureVaultUsernameKey() string`

GetFallbackAzureVaultUsernameKey returns the FallbackAzureVaultUsernameKey field if non-nil, zero value otherwise.

### GetFallbackAzureVaultUsernameKeyOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackAzureVaultUsernameKeyOk() (*string, bool)`

GetFallbackAzureVaultUsernameKeyOk returns a tuple with the FallbackAzureVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackAzureVaultUsernameKey

`func (o *OracleDSourceLinkSourceParametersAllOf) SetFallbackAzureVaultUsernameKey(v string)`

SetFallbackAzureVaultUsernameKey sets FallbackAzureVaultUsernameKey field to given value.

### HasFallbackAzureVaultUsernameKey

`func (o *OracleDSourceLinkSourceParametersAllOf) HasFallbackAzureVaultUsernameKey() bool`

HasFallbackAzureVaultUsernameKey returns a boolean if a field has been set.

### GetFallbackAzureVaultSecretKey

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackAzureVaultSecretKey() string`

GetFallbackAzureVaultSecretKey returns the FallbackAzureVaultSecretKey field if non-nil, zero value otherwise.

### GetFallbackAzureVaultSecretKeyOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackAzureVaultSecretKeyOk() (*string, bool)`

GetFallbackAzureVaultSecretKeyOk returns a tuple with the FallbackAzureVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackAzureVaultSecretKey

`func (o *OracleDSourceLinkSourceParametersAllOf) SetFallbackAzureVaultSecretKey(v string)`

SetFallbackAzureVaultSecretKey sets FallbackAzureVaultSecretKey field to given value.

### HasFallbackAzureVaultSecretKey

`func (o *OracleDSourceLinkSourceParametersAllOf) HasFallbackAzureVaultSecretKey() bool`

HasFallbackAzureVaultSecretKey returns a boolean if a field has been set.

### GetFallbackCyberarkVaultQueryString

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackCyberarkVaultQueryString() string`

GetFallbackCyberarkVaultQueryString returns the FallbackCyberarkVaultQueryString field if non-nil, zero value otherwise.

### GetFallbackCyberarkVaultQueryStringOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetFallbackCyberarkVaultQueryStringOk() (*string, bool)`

GetFallbackCyberarkVaultQueryStringOk returns a tuple with the FallbackCyberarkVaultQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackCyberarkVaultQueryString

`func (o *OracleDSourceLinkSourceParametersAllOf) SetFallbackCyberarkVaultQueryString(v string)`

SetFallbackCyberarkVaultQueryString sets FallbackCyberarkVaultQueryString field to given value.

### HasFallbackCyberarkVaultQueryString

`func (o *OracleDSourceLinkSourceParametersAllOf) HasFallbackCyberarkVaultQueryString() bool`

HasFallbackCyberarkVaultQueryString returns a boolean if a field has been set.

### GetOpsPreLogSync

`func (o *OracleDSourceLinkSourceParametersAllOf) GetOpsPreLogSync() []SourceOperation`

GetOpsPreLogSync returns the OpsPreLogSync field if non-nil, zero value otherwise.

### GetOpsPreLogSyncOk

`func (o *OracleDSourceLinkSourceParametersAllOf) GetOpsPreLogSyncOk() (*[]SourceOperation, bool)`

GetOpsPreLogSyncOk returns a tuple with the OpsPreLogSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsPreLogSync

`func (o *OracleDSourceLinkSourceParametersAllOf) SetOpsPreLogSync(v []SourceOperation)`

SetOpsPreLogSync sets OpsPreLogSync field to given value.

### HasOpsPreLogSync

`func (o *OracleDSourceLinkSourceParametersAllOf) HasOpsPreLogSync() bool`

HasOpsPreLogSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


