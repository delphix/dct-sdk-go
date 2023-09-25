# OracleDSourceLinkSourceParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the dSource to be created. | [optional] 
**SourceId** | **string** | Id of the source to link. | 
**GroupId** | Pointer to **string** | Id of the dataset group where this dSource should belong to. | [optional] 
**Description** | Pointer to **string** | The notes/description for the dSource. | [optional] 
**LogSyncEnabled** | Pointer to **bool** | True if LogSync should run for this database. | [optional] [default to false]
**MakeCurrentAccountOwner** | Pointer to **bool** | Whether the account creating this reporting schedule must be configured as owner of the reporting schedule. | [optional] [default to true]
**Tags** | Pointer to [**[]Tag**](Tag.md) | The tags to be created for dSource. | [optional] 
**OpsPreSync** | Pointer to [**[]SourceOperation**](SourceOperation.md) | Operations to perform before syncing the created dSource. These operations can quiesce any data prior to syncing. | [optional] 
**OpsPostSync** | Pointer to [**[]SourceOperation**](SourceOperation.md) | Operations to perform after syncing a created dSource. | [optional] 
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

### NewOracleDSourceLinkSourceParameters

`func NewOracleDSourceLinkSourceParameters(sourceId string, ) *OracleDSourceLinkSourceParameters`

NewOracleDSourceLinkSourceParameters instantiates a new OracleDSourceLinkSourceParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleDSourceLinkSourceParametersWithDefaults

`func NewOracleDSourceLinkSourceParametersWithDefaults() *OracleDSourceLinkSourceParameters`

NewOracleDSourceLinkSourceParametersWithDefaults instantiates a new OracleDSourceLinkSourceParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OracleDSourceLinkSourceParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OracleDSourceLinkSourceParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OracleDSourceLinkSourceParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OracleDSourceLinkSourceParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceId

`func (o *OracleDSourceLinkSourceParameters) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *OracleDSourceLinkSourceParameters) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *OracleDSourceLinkSourceParameters) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetGroupId

`func (o *OracleDSourceLinkSourceParameters) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *OracleDSourceLinkSourceParameters) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *OracleDSourceLinkSourceParameters) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *OracleDSourceLinkSourceParameters) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetDescription

`func (o *OracleDSourceLinkSourceParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OracleDSourceLinkSourceParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OracleDSourceLinkSourceParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OracleDSourceLinkSourceParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogSyncEnabled

`func (o *OracleDSourceLinkSourceParameters) GetLogSyncEnabled() bool`

GetLogSyncEnabled returns the LogSyncEnabled field if non-nil, zero value otherwise.

### GetLogSyncEnabledOk

`func (o *OracleDSourceLinkSourceParameters) GetLogSyncEnabledOk() (*bool, bool)`

GetLogSyncEnabledOk returns a tuple with the LogSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSyncEnabled

`func (o *OracleDSourceLinkSourceParameters) SetLogSyncEnabled(v bool)`

SetLogSyncEnabled sets LogSyncEnabled field to given value.

### HasLogSyncEnabled

`func (o *OracleDSourceLinkSourceParameters) HasLogSyncEnabled() bool`

HasLogSyncEnabled returns a boolean if a field has been set.

### GetMakeCurrentAccountOwner

`func (o *OracleDSourceLinkSourceParameters) GetMakeCurrentAccountOwner() bool`

GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field if non-nil, zero value otherwise.

### GetMakeCurrentAccountOwnerOk

`func (o *OracleDSourceLinkSourceParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool)`

GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCurrentAccountOwner

`func (o *OracleDSourceLinkSourceParameters) SetMakeCurrentAccountOwner(v bool)`

SetMakeCurrentAccountOwner sets MakeCurrentAccountOwner field to given value.

### HasMakeCurrentAccountOwner

`func (o *OracleDSourceLinkSourceParameters) HasMakeCurrentAccountOwner() bool`

HasMakeCurrentAccountOwner returns a boolean if a field has been set.

### GetTags

`func (o *OracleDSourceLinkSourceParameters) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OracleDSourceLinkSourceParameters) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OracleDSourceLinkSourceParameters) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OracleDSourceLinkSourceParameters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOpsPreSync

`func (o *OracleDSourceLinkSourceParameters) GetOpsPreSync() []SourceOperation`

GetOpsPreSync returns the OpsPreSync field if non-nil, zero value otherwise.

### GetOpsPreSyncOk

`func (o *OracleDSourceLinkSourceParameters) GetOpsPreSyncOk() (*[]SourceOperation, bool)`

GetOpsPreSyncOk returns a tuple with the OpsPreSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsPreSync

`func (o *OracleDSourceLinkSourceParameters) SetOpsPreSync(v []SourceOperation)`

SetOpsPreSync sets OpsPreSync field to given value.

### HasOpsPreSync

`func (o *OracleDSourceLinkSourceParameters) HasOpsPreSync() bool`

HasOpsPreSync returns a boolean if a field has been set.

### GetOpsPostSync

`func (o *OracleDSourceLinkSourceParameters) GetOpsPostSync() []SourceOperation`

GetOpsPostSync returns the OpsPostSync field if non-nil, zero value otherwise.

### GetOpsPostSyncOk

`func (o *OracleDSourceLinkSourceParameters) GetOpsPostSyncOk() (*[]SourceOperation, bool)`

GetOpsPostSyncOk returns a tuple with the OpsPostSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsPostSync

`func (o *OracleDSourceLinkSourceParameters) SetOpsPostSync(v []SourceOperation)`

SetOpsPostSync sets OpsPostSync field to given value.

### HasOpsPostSync

`func (o *OracleDSourceLinkSourceParameters) HasOpsPostSync() bool`

HasOpsPostSync returns a boolean if a field has been set.

### GetExternalFilePath

`func (o *OracleDSourceLinkSourceParameters) GetExternalFilePath() string`

GetExternalFilePath returns the ExternalFilePath field if non-nil, zero value otherwise.

### GetExternalFilePathOk

`func (o *OracleDSourceLinkSourceParameters) GetExternalFilePathOk() (*string, bool)`

GetExternalFilePathOk returns a tuple with the ExternalFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFilePath

`func (o *OracleDSourceLinkSourceParameters) SetExternalFilePath(v string)`

SetExternalFilePath sets ExternalFilePath field to given value.

### HasExternalFilePath

`func (o *OracleDSourceLinkSourceParameters) HasExternalFilePath() bool`

HasExternalFilePath returns a boolean if a field has been set.

### GetEnvironmentUserId

`func (o *OracleDSourceLinkSourceParameters) GetEnvironmentUserId() string`

GetEnvironmentUserId returns the EnvironmentUserId field if non-nil, zero value otherwise.

### GetEnvironmentUserIdOk

`func (o *OracleDSourceLinkSourceParameters) GetEnvironmentUserIdOk() (*string, bool)`

GetEnvironmentUserIdOk returns a tuple with the EnvironmentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUserId

`func (o *OracleDSourceLinkSourceParameters) SetEnvironmentUserId(v string)`

SetEnvironmentUserId sets EnvironmentUserId field to given value.

### HasEnvironmentUserId

`func (o *OracleDSourceLinkSourceParameters) HasEnvironmentUserId() bool`

HasEnvironmentUserId returns a boolean if a field has been set.

### GetBackupLevelEnabled

`func (o *OracleDSourceLinkSourceParameters) GetBackupLevelEnabled() bool`

GetBackupLevelEnabled returns the BackupLevelEnabled field if non-nil, zero value otherwise.

### GetBackupLevelEnabledOk

`func (o *OracleDSourceLinkSourceParameters) GetBackupLevelEnabledOk() (*bool, bool)`

GetBackupLevelEnabledOk returns a tuple with the BackupLevelEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLevelEnabled

`func (o *OracleDSourceLinkSourceParameters) SetBackupLevelEnabled(v bool)`

SetBackupLevelEnabled sets BackupLevelEnabled field to given value.

### HasBackupLevelEnabled

`func (o *OracleDSourceLinkSourceParameters) HasBackupLevelEnabled() bool`

HasBackupLevelEnabled returns a boolean if a field has been set.

### GetRmanChannels

`func (o *OracleDSourceLinkSourceParameters) GetRmanChannels() int32`

GetRmanChannels returns the RmanChannels field if non-nil, zero value otherwise.

### GetRmanChannelsOk

`func (o *OracleDSourceLinkSourceParameters) GetRmanChannelsOk() (*int32, bool)`

GetRmanChannelsOk returns a tuple with the RmanChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmanChannels

`func (o *OracleDSourceLinkSourceParameters) SetRmanChannels(v int32)`

SetRmanChannels sets RmanChannels field to given value.

### HasRmanChannels

`func (o *OracleDSourceLinkSourceParameters) HasRmanChannels() bool`

HasRmanChannels returns a boolean if a field has been set.

### GetFilesPerSet

`func (o *OracleDSourceLinkSourceParameters) GetFilesPerSet() int32`

GetFilesPerSet returns the FilesPerSet field if non-nil, zero value otherwise.

### GetFilesPerSetOk

`func (o *OracleDSourceLinkSourceParameters) GetFilesPerSetOk() (*int32, bool)`

GetFilesPerSetOk returns a tuple with the FilesPerSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesPerSet

`func (o *OracleDSourceLinkSourceParameters) SetFilesPerSet(v int32)`

SetFilesPerSet sets FilesPerSet field to given value.

### HasFilesPerSet

`func (o *OracleDSourceLinkSourceParameters) HasFilesPerSet() bool`

HasFilesPerSet returns a boolean if a field has been set.

### GetCheckLogical

`func (o *OracleDSourceLinkSourceParameters) GetCheckLogical() bool`

GetCheckLogical returns the CheckLogical field if non-nil, zero value otherwise.

### GetCheckLogicalOk

`func (o *OracleDSourceLinkSourceParameters) GetCheckLogicalOk() (*bool, bool)`

GetCheckLogicalOk returns a tuple with the CheckLogical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckLogical

`func (o *OracleDSourceLinkSourceParameters) SetCheckLogical(v bool)`

SetCheckLogical sets CheckLogical field to given value.

### HasCheckLogical

`func (o *OracleDSourceLinkSourceParameters) HasCheckLogical() bool`

HasCheckLogical returns a boolean if a field has been set.

### GetEncryptedLinkingEnabled

`func (o *OracleDSourceLinkSourceParameters) GetEncryptedLinkingEnabled() bool`

GetEncryptedLinkingEnabled returns the EncryptedLinkingEnabled field if non-nil, zero value otherwise.

### GetEncryptedLinkingEnabledOk

`func (o *OracleDSourceLinkSourceParameters) GetEncryptedLinkingEnabledOk() (*bool, bool)`

GetEncryptedLinkingEnabledOk returns a tuple with the EncryptedLinkingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedLinkingEnabled

`func (o *OracleDSourceLinkSourceParameters) SetEncryptedLinkingEnabled(v bool)`

SetEncryptedLinkingEnabled sets EncryptedLinkingEnabled field to given value.

### HasEncryptedLinkingEnabled

`func (o *OracleDSourceLinkSourceParameters) HasEncryptedLinkingEnabled() bool`

HasEncryptedLinkingEnabled returns a boolean if a field has been set.

### GetCompressedLinkingEnabled

`func (o *OracleDSourceLinkSourceParameters) GetCompressedLinkingEnabled() bool`

GetCompressedLinkingEnabled returns the CompressedLinkingEnabled field if non-nil, zero value otherwise.

### GetCompressedLinkingEnabledOk

`func (o *OracleDSourceLinkSourceParameters) GetCompressedLinkingEnabledOk() (*bool, bool)`

GetCompressedLinkingEnabledOk returns a tuple with the CompressedLinkingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedLinkingEnabled

`func (o *OracleDSourceLinkSourceParameters) SetCompressedLinkingEnabled(v bool)`

SetCompressedLinkingEnabled sets CompressedLinkingEnabled field to given value.

### HasCompressedLinkingEnabled

`func (o *OracleDSourceLinkSourceParameters) HasCompressedLinkingEnabled() bool`

HasCompressedLinkingEnabled returns a boolean if a field has been set.

### GetBandwidthLimit

`func (o *OracleDSourceLinkSourceParameters) GetBandwidthLimit() int32`

GetBandwidthLimit returns the BandwidthLimit field if non-nil, zero value otherwise.

### GetBandwidthLimitOk

`func (o *OracleDSourceLinkSourceParameters) GetBandwidthLimitOk() (*int32, bool)`

GetBandwidthLimitOk returns a tuple with the BandwidthLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimit

`func (o *OracleDSourceLinkSourceParameters) SetBandwidthLimit(v int32)`

SetBandwidthLimit sets BandwidthLimit field to given value.

### HasBandwidthLimit

`func (o *OracleDSourceLinkSourceParameters) HasBandwidthLimit() bool`

HasBandwidthLimit returns a boolean if a field has been set.

### GetNumberOfConnections

`func (o *OracleDSourceLinkSourceParameters) GetNumberOfConnections() int32`

GetNumberOfConnections returns the NumberOfConnections field if non-nil, zero value otherwise.

### GetNumberOfConnectionsOk

`func (o *OracleDSourceLinkSourceParameters) GetNumberOfConnectionsOk() (*int32, bool)`

GetNumberOfConnectionsOk returns a tuple with the NumberOfConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfConnections

`func (o *OracleDSourceLinkSourceParameters) SetNumberOfConnections(v int32)`

SetNumberOfConnections sets NumberOfConnections field to given value.

### HasNumberOfConnections

`func (o *OracleDSourceLinkSourceParameters) HasNumberOfConnections() bool`

HasNumberOfConnections returns a boolean if a field has been set.

### GetDiagnoseNoLoggingFaults

`func (o *OracleDSourceLinkSourceParameters) GetDiagnoseNoLoggingFaults() bool`

GetDiagnoseNoLoggingFaults returns the DiagnoseNoLoggingFaults field if non-nil, zero value otherwise.

### GetDiagnoseNoLoggingFaultsOk

`func (o *OracleDSourceLinkSourceParameters) GetDiagnoseNoLoggingFaultsOk() (*bool, bool)`

GetDiagnoseNoLoggingFaultsOk returns a tuple with the DiagnoseNoLoggingFaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnoseNoLoggingFaults

`func (o *OracleDSourceLinkSourceParameters) SetDiagnoseNoLoggingFaults(v bool)`

SetDiagnoseNoLoggingFaults sets DiagnoseNoLoggingFaults field to given value.

### HasDiagnoseNoLoggingFaults

`func (o *OracleDSourceLinkSourceParameters) HasDiagnoseNoLoggingFaults() bool`

HasDiagnoseNoLoggingFaults returns a boolean if a field has been set.

### GetPreProvisioningEnabled

`func (o *OracleDSourceLinkSourceParameters) GetPreProvisioningEnabled() bool`

GetPreProvisioningEnabled returns the PreProvisioningEnabled field if non-nil, zero value otherwise.

### GetPreProvisioningEnabledOk

`func (o *OracleDSourceLinkSourceParameters) GetPreProvisioningEnabledOk() (*bool, bool)`

GetPreProvisioningEnabledOk returns a tuple with the PreProvisioningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreProvisioningEnabled

`func (o *OracleDSourceLinkSourceParameters) SetPreProvisioningEnabled(v bool)`

SetPreProvisioningEnabled sets PreProvisioningEnabled field to given value.

### HasPreProvisioningEnabled

`func (o *OracleDSourceLinkSourceParameters) HasPreProvisioningEnabled() bool`

HasPreProvisioningEnabled returns a boolean if a field has been set.

### GetLinkNow

`func (o *OracleDSourceLinkSourceParameters) GetLinkNow() bool`

GetLinkNow returns the LinkNow field if non-nil, zero value otherwise.

### GetLinkNowOk

`func (o *OracleDSourceLinkSourceParameters) GetLinkNowOk() (*bool, bool)`

GetLinkNowOk returns a tuple with the LinkNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNow

`func (o *OracleDSourceLinkSourceParameters) SetLinkNow(v bool)`

SetLinkNow sets LinkNow field to given value.

### HasLinkNow

`func (o *OracleDSourceLinkSourceParameters) HasLinkNow() bool`

HasLinkNow returns a boolean if a field has been set.

### GetForceFullBackup

`func (o *OracleDSourceLinkSourceParameters) GetForceFullBackup() bool`

GetForceFullBackup returns the ForceFullBackup field if non-nil, zero value otherwise.

### GetForceFullBackupOk

`func (o *OracleDSourceLinkSourceParameters) GetForceFullBackupOk() (*bool, bool)`

GetForceFullBackupOk returns a tuple with the ForceFullBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceFullBackup

`func (o *OracleDSourceLinkSourceParameters) SetForceFullBackup(v bool)`

SetForceFullBackup sets ForceFullBackup field to given value.

### HasForceFullBackup

`func (o *OracleDSourceLinkSourceParameters) HasForceFullBackup() bool`

HasForceFullBackup returns a boolean if a field has been set.

### GetDoubleSync

`func (o *OracleDSourceLinkSourceParameters) GetDoubleSync() bool`

GetDoubleSync returns the DoubleSync field if non-nil, zero value otherwise.

### GetDoubleSyncOk

`func (o *OracleDSourceLinkSourceParameters) GetDoubleSyncOk() (*bool, bool)`

GetDoubleSyncOk returns a tuple with the DoubleSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleSync

`func (o *OracleDSourceLinkSourceParameters) SetDoubleSync(v bool)`

SetDoubleSync sets DoubleSync field to given value.

### HasDoubleSync

`func (o *OracleDSourceLinkSourceParameters) HasDoubleSync() bool`

HasDoubleSync returns a boolean if a field has been set.

### GetSkipSpaceCheck

`func (o *OracleDSourceLinkSourceParameters) GetSkipSpaceCheck() bool`

GetSkipSpaceCheck returns the SkipSpaceCheck field if non-nil, zero value otherwise.

### GetSkipSpaceCheckOk

`func (o *OracleDSourceLinkSourceParameters) GetSkipSpaceCheckOk() (*bool, bool)`

GetSkipSpaceCheckOk returns a tuple with the SkipSpaceCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSpaceCheck

`func (o *OracleDSourceLinkSourceParameters) SetSkipSpaceCheck(v bool)`

SetSkipSpaceCheck sets SkipSpaceCheck field to given value.

### HasSkipSpaceCheck

`func (o *OracleDSourceLinkSourceParameters) HasSkipSpaceCheck() bool`

HasSkipSpaceCheck returns a boolean if a field has been set.

### GetDoNotResume

`func (o *OracleDSourceLinkSourceParameters) GetDoNotResume() bool`

GetDoNotResume returns the DoNotResume field if non-nil, zero value otherwise.

### GetDoNotResumeOk

`func (o *OracleDSourceLinkSourceParameters) GetDoNotResumeOk() (*bool, bool)`

GetDoNotResumeOk returns a tuple with the DoNotResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotResume

`func (o *OracleDSourceLinkSourceParameters) SetDoNotResume(v bool)`

SetDoNotResume sets DoNotResume field to given value.

### HasDoNotResume

`func (o *OracleDSourceLinkSourceParameters) HasDoNotResume() bool`

HasDoNotResume returns a boolean if a field has been set.

### GetFilesForFullBackup

`func (o *OracleDSourceLinkSourceParameters) GetFilesForFullBackup() []int32`

GetFilesForFullBackup returns the FilesForFullBackup field if non-nil, zero value otherwise.

### GetFilesForFullBackupOk

`func (o *OracleDSourceLinkSourceParameters) GetFilesForFullBackupOk() (*[]int32, bool)`

GetFilesForFullBackupOk returns a tuple with the FilesForFullBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesForFullBackup

`func (o *OracleDSourceLinkSourceParameters) SetFilesForFullBackup(v []int32)`

SetFilesForFullBackup sets FilesForFullBackup field to given value.

### HasFilesForFullBackup

`func (o *OracleDSourceLinkSourceParameters) HasFilesForFullBackup() bool`

HasFilesForFullBackup returns a boolean if a field has been set.

### GetLogSyncMode

`func (o *OracleDSourceLinkSourceParameters) GetLogSyncMode() string`

GetLogSyncMode returns the LogSyncMode field if non-nil, zero value otherwise.

### GetLogSyncModeOk

`func (o *OracleDSourceLinkSourceParameters) GetLogSyncModeOk() (*string, bool)`

GetLogSyncModeOk returns a tuple with the LogSyncMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSyncMode

`func (o *OracleDSourceLinkSourceParameters) SetLogSyncMode(v string)`

SetLogSyncMode sets LogSyncMode field to given value.

### HasLogSyncMode

`func (o *OracleDSourceLinkSourceParameters) HasLogSyncMode() bool`

HasLogSyncMode returns a boolean if a field has been set.

### GetLogSyncInterval

`func (o *OracleDSourceLinkSourceParameters) GetLogSyncInterval() int32`

GetLogSyncInterval returns the LogSyncInterval field if non-nil, zero value otherwise.

### GetLogSyncIntervalOk

`func (o *OracleDSourceLinkSourceParameters) GetLogSyncIntervalOk() (*int32, bool)`

GetLogSyncIntervalOk returns a tuple with the LogSyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSyncInterval

`func (o *OracleDSourceLinkSourceParameters) SetLogSyncInterval(v int32)`

SetLogSyncInterval sets LogSyncInterval field to given value.

### HasLogSyncInterval

`func (o *OracleDSourceLinkSourceParameters) HasLogSyncInterval() bool`

HasLogSyncInterval returns a boolean if a field has been set.

### GetNonSysUsername

`func (o *OracleDSourceLinkSourceParameters) GetNonSysUsername() string`

GetNonSysUsername returns the NonSysUsername field if non-nil, zero value otherwise.

### GetNonSysUsernameOk

`func (o *OracleDSourceLinkSourceParameters) GetNonSysUsernameOk() (*string, bool)`

GetNonSysUsernameOk returns a tuple with the NonSysUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysUsername

`func (o *OracleDSourceLinkSourceParameters) SetNonSysUsername(v string)`

SetNonSysUsername sets NonSysUsername field to given value.

### HasNonSysUsername

`func (o *OracleDSourceLinkSourceParameters) HasNonSysUsername() bool`

HasNonSysUsername returns a boolean if a field has been set.

### GetNonSysPassword

`func (o *OracleDSourceLinkSourceParameters) GetNonSysPassword() string`

GetNonSysPassword returns the NonSysPassword field if non-nil, zero value otherwise.

### GetNonSysPasswordOk

`func (o *OracleDSourceLinkSourceParameters) GetNonSysPasswordOk() (*string, bool)`

GetNonSysPasswordOk returns a tuple with the NonSysPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysPassword

`func (o *OracleDSourceLinkSourceParameters) SetNonSysPassword(v string)`

SetNonSysPassword sets NonSysPassword field to given value.

### HasNonSysPassword

`func (o *OracleDSourceLinkSourceParameters) HasNonSysPassword() bool`

HasNonSysPassword returns a boolean if a field has been set.

### GetNonSysVault

`func (o *OracleDSourceLinkSourceParameters) GetNonSysVault() string`

GetNonSysVault returns the NonSysVault field if non-nil, zero value otherwise.

### GetNonSysVaultOk

`func (o *OracleDSourceLinkSourceParameters) GetNonSysVaultOk() (*string, bool)`

GetNonSysVaultOk returns a tuple with the NonSysVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysVault

`func (o *OracleDSourceLinkSourceParameters) SetNonSysVault(v string)`

SetNonSysVault sets NonSysVault field to given value.

### HasNonSysVault

`func (o *OracleDSourceLinkSourceParameters) HasNonSysVault() bool`

HasNonSysVault returns a boolean if a field has been set.

### GetNonSysHashicorpVaultEngine

`func (o *OracleDSourceLinkSourceParameters) GetNonSysHashicorpVaultEngine() string`

GetNonSysHashicorpVaultEngine returns the NonSysHashicorpVaultEngine field if non-nil, zero value otherwise.

### GetNonSysHashicorpVaultEngineOk

`func (o *OracleDSourceLinkSourceParameters) GetNonSysHashicorpVaultEngineOk() (*string, bool)`

GetNonSysHashicorpVaultEngineOk returns a tuple with the NonSysHashicorpVaultEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysHashicorpVaultEngine

`func (o *OracleDSourceLinkSourceParameters) SetNonSysHashicorpVaultEngine(v string)`

SetNonSysHashicorpVaultEngine sets NonSysHashicorpVaultEngine field to given value.

### HasNonSysHashicorpVaultEngine

`func (o *OracleDSourceLinkSourceParameters) HasNonSysHashicorpVaultEngine() bool`

HasNonSysHashicorpVaultEngine returns a boolean if a field has been set.

### GetNonSysHashicorpVaultSecretPath

`func (o *OracleDSourceLinkSourceParameters) GetNonSysHashicorpVaultSecretPath() string`

GetNonSysHashicorpVaultSecretPath returns the NonSysHashicorpVaultSecretPath field if non-nil, zero value otherwise.

### GetNonSysHashicorpVaultSecretPathOk

`func (o *OracleDSourceLinkSourceParameters) GetNonSysHashicorpVaultSecretPathOk() (*string, bool)`

GetNonSysHashicorpVaultSecretPathOk returns a tuple with the NonSysHashicorpVaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysHashicorpVaultSecretPath

`func (o *OracleDSourceLinkSourceParameters) SetNonSysHashicorpVaultSecretPath(v string)`

SetNonSysHashicorpVaultSecretPath sets NonSysHashicorpVaultSecretPath field to given value.

### HasNonSysHashicorpVaultSecretPath

`func (o *OracleDSourceLinkSourceParameters) HasNonSysHashicorpVaultSecretPath() bool`

HasNonSysHashicorpVaultSecretPath returns a boolean if a field has been set.

### GetNonSysHashicorpVaultUsernameKey

`func (o *OracleDSourceLinkSourceParameters) GetNonSysHashicorpVaultUsernameKey() string`

GetNonSysHashicorpVaultUsernameKey returns the NonSysHashicorpVaultUsernameKey field if non-nil, zero value otherwise.

### GetNonSysHashicorpVaultUsernameKeyOk

`func (o *OracleDSourceLinkSourceParameters) GetNonSysHashicorpVaultUsernameKeyOk() (*string, bool)`

GetNonSysHashicorpVaultUsernameKeyOk returns a tuple with the NonSysHashicorpVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysHashicorpVaultUsernameKey

`func (o *OracleDSourceLinkSourceParameters) SetNonSysHashicorpVaultUsernameKey(v string)`

SetNonSysHashicorpVaultUsernameKey sets NonSysHashicorpVaultUsernameKey field to given value.

### HasNonSysHashicorpVaultUsernameKey

`func (o *OracleDSourceLinkSourceParameters) HasNonSysHashicorpVaultUsernameKey() bool`

HasNonSysHashicorpVaultUsernameKey returns a boolean if a field has been set.

### GetNonSysHashicorpVaultSecretKey

`func (o *OracleDSourceLinkSourceParameters) GetNonSysHashicorpVaultSecretKey() string`

GetNonSysHashicorpVaultSecretKey returns the NonSysHashicorpVaultSecretKey field if non-nil, zero value otherwise.

### GetNonSysHashicorpVaultSecretKeyOk

`func (o *OracleDSourceLinkSourceParameters) GetNonSysHashicorpVaultSecretKeyOk() (*string, bool)`

GetNonSysHashicorpVaultSecretKeyOk returns a tuple with the NonSysHashicorpVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysHashicorpVaultSecretKey

`func (o *OracleDSourceLinkSourceParameters) SetNonSysHashicorpVaultSecretKey(v string)`

SetNonSysHashicorpVaultSecretKey sets NonSysHashicorpVaultSecretKey field to given value.

### HasNonSysHashicorpVaultSecretKey

`func (o *OracleDSourceLinkSourceParameters) HasNonSysHashicorpVaultSecretKey() bool`

HasNonSysHashicorpVaultSecretKey returns a boolean if a field has been set.

### GetNonSysAzureVaultName

`func (o *OracleDSourceLinkSourceParameters) GetNonSysAzureVaultName() string`

GetNonSysAzureVaultName returns the NonSysAzureVaultName field if non-nil, zero value otherwise.

### GetNonSysAzureVaultNameOk

`func (o *OracleDSourceLinkSourceParameters) GetNonSysAzureVaultNameOk() (*string, bool)`

GetNonSysAzureVaultNameOk returns a tuple with the NonSysAzureVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysAzureVaultName

`func (o *OracleDSourceLinkSourceParameters) SetNonSysAzureVaultName(v string)`

SetNonSysAzureVaultName sets NonSysAzureVaultName field to given value.

### HasNonSysAzureVaultName

`func (o *OracleDSourceLinkSourceParameters) HasNonSysAzureVaultName() bool`

HasNonSysAzureVaultName returns a boolean if a field has been set.

### GetNonSysAzureVaultUsernameKey

`func (o *OracleDSourceLinkSourceParameters) GetNonSysAzureVaultUsernameKey() string`

GetNonSysAzureVaultUsernameKey returns the NonSysAzureVaultUsernameKey field if non-nil, zero value otherwise.

### GetNonSysAzureVaultUsernameKeyOk

`func (o *OracleDSourceLinkSourceParameters) GetNonSysAzureVaultUsernameKeyOk() (*string, bool)`

GetNonSysAzureVaultUsernameKeyOk returns a tuple with the NonSysAzureVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysAzureVaultUsernameKey

`func (o *OracleDSourceLinkSourceParameters) SetNonSysAzureVaultUsernameKey(v string)`

SetNonSysAzureVaultUsernameKey sets NonSysAzureVaultUsernameKey field to given value.

### HasNonSysAzureVaultUsernameKey

`func (o *OracleDSourceLinkSourceParameters) HasNonSysAzureVaultUsernameKey() bool`

HasNonSysAzureVaultUsernameKey returns a boolean if a field has been set.

### GetNonSysAzureVaultSecretKey

`func (o *OracleDSourceLinkSourceParameters) GetNonSysAzureVaultSecretKey() string`

GetNonSysAzureVaultSecretKey returns the NonSysAzureVaultSecretKey field if non-nil, zero value otherwise.

### GetNonSysAzureVaultSecretKeyOk

`func (o *OracleDSourceLinkSourceParameters) GetNonSysAzureVaultSecretKeyOk() (*string, bool)`

GetNonSysAzureVaultSecretKeyOk returns a tuple with the NonSysAzureVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysAzureVaultSecretKey

`func (o *OracleDSourceLinkSourceParameters) SetNonSysAzureVaultSecretKey(v string)`

SetNonSysAzureVaultSecretKey sets NonSysAzureVaultSecretKey field to given value.

### HasNonSysAzureVaultSecretKey

`func (o *OracleDSourceLinkSourceParameters) HasNonSysAzureVaultSecretKey() bool`

HasNonSysAzureVaultSecretKey returns a boolean if a field has been set.

### GetNonSysCyberarkVaultQueryString

`func (o *OracleDSourceLinkSourceParameters) GetNonSysCyberarkVaultQueryString() string`

GetNonSysCyberarkVaultQueryString returns the NonSysCyberarkVaultQueryString field if non-nil, zero value otherwise.

### GetNonSysCyberarkVaultQueryStringOk

`func (o *OracleDSourceLinkSourceParameters) GetNonSysCyberarkVaultQueryStringOk() (*string, bool)`

GetNonSysCyberarkVaultQueryStringOk returns a tuple with the NonSysCyberarkVaultQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSysCyberarkVaultQueryString

`func (o *OracleDSourceLinkSourceParameters) SetNonSysCyberarkVaultQueryString(v string)`

SetNonSysCyberarkVaultQueryString sets NonSysCyberarkVaultQueryString field to given value.

### HasNonSysCyberarkVaultQueryString

`func (o *OracleDSourceLinkSourceParameters) HasNonSysCyberarkVaultQueryString() bool`

HasNonSysCyberarkVaultQueryString returns a boolean if a field has been set.

### GetFallbackUsername

`func (o *OracleDSourceLinkSourceParameters) GetFallbackUsername() string`

GetFallbackUsername returns the FallbackUsername field if non-nil, zero value otherwise.

### GetFallbackUsernameOk

`func (o *OracleDSourceLinkSourceParameters) GetFallbackUsernameOk() (*string, bool)`

GetFallbackUsernameOk returns a tuple with the FallbackUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUsername

`func (o *OracleDSourceLinkSourceParameters) SetFallbackUsername(v string)`

SetFallbackUsername sets FallbackUsername field to given value.

### HasFallbackUsername

`func (o *OracleDSourceLinkSourceParameters) HasFallbackUsername() bool`

HasFallbackUsername returns a boolean if a field has been set.

### GetFallbackPassword

`func (o *OracleDSourceLinkSourceParameters) GetFallbackPassword() string`

GetFallbackPassword returns the FallbackPassword field if non-nil, zero value otherwise.

### GetFallbackPasswordOk

`func (o *OracleDSourceLinkSourceParameters) GetFallbackPasswordOk() (*string, bool)`

GetFallbackPasswordOk returns a tuple with the FallbackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackPassword

`func (o *OracleDSourceLinkSourceParameters) SetFallbackPassword(v string)`

SetFallbackPassword sets FallbackPassword field to given value.

### HasFallbackPassword

`func (o *OracleDSourceLinkSourceParameters) HasFallbackPassword() bool`

HasFallbackPassword returns a boolean if a field has been set.

### GetFallbackVault

`func (o *OracleDSourceLinkSourceParameters) GetFallbackVault() string`

GetFallbackVault returns the FallbackVault field if non-nil, zero value otherwise.

### GetFallbackVaultOk

`func (o *OracleDSourceLinkSourceParameters) GetFallbackVaultOk() (*string, bool)`

GetFallbackVaultOk returns a tuple with the FallbackVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackVault

`func (o *OracleDSourceLinkSourceParameters) SetFallbackVault(v string)`

SetFallbackVault sets FallbackVault field to given value.

### HasFallbackVault

`func (o *OracleDSourceLinkSourceParameters) HasFallbackVault() bool`

HasFallbackVault returns a boolean if a field has been set.

### GetFallbackHashicorpVaultEngine

`func (o *OracleDSourceLinkSourceParameters) GetFallbackHashicorpVaultEngine() string`

GetFallbackHashicorpVaultEngine returns the FallbackHashicorpVaultEngine field if non-nil, zero value otherwise.

### GetFallbackHashicorpVaultEngineOk

`func (o *OracleDSourceLinkSourceParameters) GetFallbackHashicorpVaultEngineOk() (*string, bool)`

GetFallbackHashicorpVaultEngineOk returns a tuple with the FallbackHashicorpVaultEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackHashicorpVaultEngine

`func (o *OracleDSourceLinkSourceParameters) SetFallbackHashicorpVaultEngine(v string)`

SetFallbackHashicorpVaultEngine sets FallbackHashicorpVaultEngine field to given value.

### HasFallbackHashicorpVaultEngine

`func (o *OracleDSourceLinkSourceParameters) HasFallbackHashicorpVaultEngine() bool`

HasFallbackHashicorpVaultEngine returns a boolean if a field has been set.

### GetFallbackHashicorpVaultSecretPath

`func (o *OracleDSourceLinkSourceParameters) GetFallbackHashicorpVaultSecretPath() string`

GetFallbackHashicorpVaultSecretPath returns the FallbackHashicorpVaultSecretPath field if non-nil, zero value otherwise.

### GetFallbackHashicorpVaultSecretPathOk

`func (o *OracleDSourceLinkSourceParameters) GetFallbackHashicorpVaultSecretPathOk() (*string, bool)`

GetFallbackHashicorpVaultSecretPathOk returns a tuple with the FallbackHashicorpVaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackHashicorpVaultSecretPath

`func (o *OracleDSourceLinkSourceParameters) SetFallbackHashicorpVaultSecretPath(v string)`

SetFallbackHashicorpVaultSecretPath sets FallbackHashicorpVaultSecretPath field to given value.

### HasFallbackHashicorpVaultSecretPath

`func (o *OracleDSourceLinkSourceParameters) HasFallbackHashicorpVaultSecretPath() bool`

HasFallbackHashicorpVaultSecretPath returns a boolean if a field has been set.

### GetFallbackHashicorpVaultUsernameKey

`func (o *OracleDSourceLinkSourceParameters) GetFallbackHashicorpVaultUsernameKey() string`

GetFallbackHashicorpVaultUsernameKey returns the FallbackHashicorpVaultUsernameKey field if non-nil, zero value otherwise.

### GetFallbackHashicorpVaultUsernameKeyOk

`func (o *OracleDSourceLinkSourceParameters) GetFallbackHashicorpVaultUsernameKeyOk() (*string, bool)`

GetFallbackHashicorpVaultUsernameKeyOk returns a tuple with the FallbackHashicorpVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackHashicorpVaultUsernameKey

`func (o *OracleDSourceLinkSourceParameters) SetFallbackHashicorpVaultUsernameKey(v string)`

SetFallbackHashicorpVaultUsernameKey sets FallbackHashicorpVaultUsernameKey field to given value.

### HasFallbackHashicorpVaultUsernameKey

`func (o *OracleDSourceLinkSourceParameters) HasFallbackHashicorpVaultUsernameKey() bool`

HasFallbackHashicorpVaultUsernameKey returns a boolean if a field has been set.

### GetFallbackHashicorpVaultSecretKey

`func (o *OracleDSourceLinkSourceParameters) GetFallbackHashicorpVaultSecretKey() string`

GetFallbackHashicorpVaultSecretKey returns the FallbackHashicorpVaultSecretKey field if non-nil, zero value otherwise.

### GetFallbackHashicorpVaultSecretKeyOk

`func (o *OracleDSourceLinkSourceParameters) GetFallbackHashicorpVaultSecretKeyOk() (*string, bool)`

GetFallbackHashicorpVaultSecretKeyOk returns a tuple with the FallbackHashicorpVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackHashicorpVaultSecretKey

`func (o *OracleDSourceLinkSourceParameters) SetFallbackHashicorpVaultSecretKey(v string)`

SetFallbackHashicorpVaultSecretKey sets FallbackHashicorpVaultSecretKey field to given value.

### HasFallbackHashicorpVaultSecretKey

`func (o *OracleDSourceLinkSourceParameters) HasFallbackHashicorpVaultSecretKey() bool`

HasFallbackHashicorpVaultSecretKey returns a boolean if a field has been set.

### GetFallbackAzureVaultName

`func (o *OracleDSourceLinkSourceParameters) GetFallbackAzureVaultName() string`

GetFallbackAzureVaultName returns the FallbackAzureVaultName field if non-nil, zero value otherwise.

### GetFallbackAzureVaultNameOk

`func (o *OracleDSourceLinkSourceParameters) GetFallbackAzureVaultNameOk() (*string, bool)`

GetFallbackAzureVaultNameOk returns a tuple with the FallbackAzureVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackAzureVaultName

`func (o *OracleDSourceLinkSourceParameters) SetFallbackAzureVaultName(v string)`

SetFallbackAzureVaultName sets FallbackAzureVaultName field to given value.

### HasFallbackAzureVaultName

`func (o *OracleDSourceLinkSourceParameters) HasFallbackAzureVaultName() bool`

HasFallbackAzureVaultName returns a boolean if a field has been set.

### GetFallbackAzureVaultUsernameKey

`func (o *OracleDSourceLinkSourceParameters) GetFallbackAzureVaultUsernameKey() string`

GetFallbackAzureVaultUsernameKey returns the FallbackAzureVaultUsernameKey field if non-nil, zero value otherwise.

### GetFallbackAzureVaultUsernameKeyOk

`func (o *OracleDSourceLinkSourceParameters) GetFallbackAzureVaultUsernameKeyOk() (*string, bool)`

GetFallbackAzureVaultUsernameKeyOk returns a tuple with the FallbackAzureVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackAzureVaultUsernameKey

`func (o *OracleDSourceLinkSourceParameters) SetFallbackAzureVaultUsernameKey(v string)`

SetFallbackAzureVaultUsernameKey sets FallbackAzureVaultUsernameKey field to given value.

### HasFallbackAzureVaultUsernameKey

`func (o *OracleDSourceLinkSourceParameters) HasFallbackAzureVaultUsernameKey() bool`

HasFallbackAzureVaultUsernameKey returns a boolean if a field has been set.

### GetFallbackAzureVaultSecretKey

`func (o *OracleDSourceLinkSourceParameters) GetFallbackAzureVaultSecretKey() string`

GetFallbackAzureVaultSecretKey returns the FallbackAzureVaultSecretKey field if non-nil, zero value otherwise.

### GetFallbackAzureVaultSecretKeyOk

`func (o *OracleDSourceLinkSourceParameters) GetFallbackAzureVaultSecretKeyOk() (*string, bool)`

GetFallbackAzureVaultSecretKeyOk returns a tuple with the FallbackAzureVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackAzureVaultSecretKey

`func (o *OracleDSourceLinkSourceParameters) SetFallbackAzureVaultSecretKey(v string)`

SetFallbackAzureVaultSecretKey sets FallbackAzureVaultSecretKey field to given value.

### HasFallbackAzureVaultSecretKey

`func (o *OracleDSourceLinkSourceParameters) HasFallbackAzureVaultSecretKey() bool`

HasFallbackAzureVaultSecretKey returns a boolean if a field has been set.

### GetFallbackCyberarkVaultQueryString

`func (o *OracleDSourceLinkSourceParameters) GetFallbackCyberarkVaultQueryString() string`

GetFallbackCyberarkVaultQueryString returns the FallbackCyberarkVaultQueryString field if non-nil, zero value otherwise.

### GetFallbackCyberarkVaultQueryStringOk

`func (o *OracleDSourceLinkSourceParameters) GetFallbackCyberarkVaultQueryStringOk() (*string, bool)`

GetFallbackCyberarkVaultQueryStringOk returns a tuple with the FallbackCyberarkVaultQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackCyberarkVaultQueryString

`func (o *OracleDSourceLinkSourceParameters) SetFallbackCyberarkVaultQueryString(v string)`

SetFallbackCyberarkVaultQueryString sets FallbackCyberarkVaultQueryString field to given value.

### HasFallbackCyberarkVaultQueryString

`func (o *OracleDSourceLinkSourceParameters) HasFallbackCyberarkVaultQueryString() bool`

HasFallbackCyberarkVaultQueryString returns a boolean if a field has been set.

### GetOpsPreLogSync

`func (o *OracleDSourceLinkSourceParameters) GetOpsPreLogSync() []SourceOperation`

GetOpsPreLogSync returns the OpsPreLogSync field if non-nil, zero value otherwise.

### GetOpsPreLogSyncOk

`func (o *OracleDSourceLinkSourceParameters) GetOpsPreLogSyncOk() (*[]SourceOperation, bool)`

GetOpsPreLogSyncOk returns a tuple with the OpsPreLogSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsPreLogSync

`func (o *OracleDSourceLinkSourceParameters) SetOpsPreLogSync(v []SourceOperation)`

SetOpsPreLogSync sets OpsPreLogSync field to given value.

### HasOpsPreLogSync

`func (o *OracleDSourceLinkSourceParameters) HasOpsPreLogSync() bool`

HasOpsPreLogSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


