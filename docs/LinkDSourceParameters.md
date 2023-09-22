# LinkDSourceParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the dSource to be created. | [optional] 
**SourceId** | **string** | Id of the source to link. | 
**GroupId** | Pointer to **string** | Id of the dataset group where this dSource should belong to. | [optional] 
**EnvironmentUserId** | Pointer to **string** | Id of the environment user to use for linking. | [optional] 
**Description** | Pointer to **string** | The notes/description for the dSource. | [optional] 
**OracleNonSysUsername** | Pointer to **string** | Non-SYS database user to access this database. Only required for username-password auth (Oracle only). | [optional] 
**OracleNonSysPassword** | Pointer to **string** | Password for non sys user authentication (Oracle only). | [optional] 
**OracleNonSysVault** | Pointer to **string** | The name or reference of the vault from which to read the database credentials (Oracle only). | [optional] 
**OracleNonSysHashicorpVaultEngine** | Pointer to **string** | Vault engine name where the credential is stored (Oracle only). | [optional] 
**OracleNonSysHashicorpVaultSecretPath** | Pointer to **string** | Path in the vault engine where the credential is stored (Oracle only). | [optional] 
**OracleNonSysHashicorpVaultUsernameKey** | Pointer to **string** | Hashicorp vault key for the username in the key-value store (Oracle only). | [optional] 
**OracleNonSysHashicorpVaultSecretKey** | Pointer to **string** | Hashicorp vault key for the password in the key-value store (Oracle only). | [optional] 
**OracleNonSysAzureVaultName** | Pointer to **string** | Azure key vault name (Oracle only). | [optional] 
**OracleNonSysAzureVaultUsernameKey** | Pointer to **string** | Azure vault key for the username in the key-value store (Oracle only). | [optional] 
**OracleNonSysAzureVaultSecretKey** | Pointer to **string** | Azure vault key for the password in the key-value store (Oracle only). | [optional] 
**OracleNonSysCyberarkVaultQueryString** | Pointer to **string** | Query to find a credential in the CyberArk vault (Oracle only). | [optional] 
**OracleFallbackUsername** | Pointer to **string** | The database fallback username. Optional if bequeath connections are enabled (to be used in case of bequeath connection failures). Only required for username-password auth (Oracle only). | [optional] 
**OracleFallbackPassword** | Pointer to **string** | Password for fallback username (Oracle only). | [optional] 
**OracleFallbackVault** | Pointer to **string** | The name or reference of the vault from which to read the database credentials (Oracle only). | [optional] 
**OracleFallbackHashicorpVaultEngine** | Pointer to **string** | Vault engine name where the credential is stored (Oracle only). | [optional] 
**OracleFallbackHashicorpVaultSecretPath** | Pointer to **string** | Path in the vault engine where the credential is stored (Oracle only). | [optional] 
**OracleFallbackHashicorpVaultUsernameKey** | Pointer to **string** | Hashicorp vault key for the username in the key-value store (Oracle only). | [optional] 
**OracleFallbackHashicorpVaultSecretKey** | Pointer to **string** | Hashicorp vault key for the password in the key-value store (Oracle only). | [optional] 
**OracleFallbackAzureVaultName** | Pointer to **string** | Azure key vault name (Oracle only). | [optional] 
**OracleFallbackAzureVaultUsernameKey** | Pointer to **string** | Azure vault key for the username in the key-value store (Oracle only). | [optional] 
**OracleFallbackAzureVaultSecretKey** | Pointer to **string** | Azure vault key for the password in the key-value store (Oracle only). | [optional] 
**OracleFallbackCyberarkVaultQueryString** | Pointer to **string** | Query to find a credential in the CyberArk vault (Oracle only). | [optional] 
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
**LogSyncEnabled** | Pointer to **bool** | True if LogSync should run for this database. | [optional] [default to false]
**ExternalFilePath** | Pointer to **string** | External file path. | [optional] 
**MakeCurrentAccountOwner** | Pointer to **bool** | Whether the account creating this reporting schedule must be configured as owner of the reporting schedule. | [optional] [default to true]
**Tags** | Pointer to [**[]Tag**](Tag.md) | The tags to be created for dSource. | [optional] 
**OpsPreSync** | Pointer to [**[]SourceOperation**](SourceOperation.md) | Operations to perform before syncing the created dSource. These operations can quiesce any data prior to syncing. | [optional] 
**OpsPreLogSync** | Pointer to [**[]SourceOperation**](SourceOperation.md) | Operations to perform after syncing a created dSource and before running the LogSync. | [optional] 
**OpsPostSync** | Pointer to [**[]SourceOperation**](SourceOperation.md) | Operations to perform after syncing a created dSource. | [optional] 

## Methods

### NewLinkDSourceParameters

`func NewLinkDSourceParameters(sourceId string, ) *LinkDSourceParameters`

NewLinkDSourceParameters instantiates a new LinkDSourceParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkDSourceParametersWithDefaults

`func NewLinkDSourceParametersWithDefaults() *LinkDSourceParameters`

NewLinkDSourceParametersWithDefaults instantiates a new LinkDSourceParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LinkDSourceParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinkDSourceParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinkDSourceParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LinkDSourceParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceId

`func (o *LinkDSourceParameters) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *LinkDSourceParameters) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *LinkDSourceParameters) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetGroupId

`func (o *LinkDSourceParameters) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *LinkDSourceParameters) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *LinkDSourceParameters) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *LinkDSourceParameters) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetEnvironmentUserId

`func (o *LinkDSourceParameters) GetEnvironmentUserId() string`

GetEnvironmentUserId returns the EnvironmentUserId field if non-nil, zero value otherwise.

### GetEnvironmentUserIdOk

`func (o *LinkDSourceParameters) GetEnvironmentUserIdOk() (*string, bool)`

GetEnvironmentUserIdOk returns a tuple with the EnvironmentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUserId

`func (o *LinkDSourceParameters) SetEnvironmentUserId(v string)`

SetEnvironmentUserId sets EnvironmentUserId field to given value.

### HasEnvironmentUserId

`func (o *LinkDSourceParameters) HasEnvironmentUserId() bool`

HasEnvironmentUserId returns a boolean if a field has been set.

### GetDescription

`func (o *LinkDSourceParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LinkDSourceParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LinkDSourceParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LinkDSourceParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOracleNonSysUsername

`func (o *LinkDSourceParameters) GetOracleNonSysUsername() string`

GetOracleNonSysUsername returns the OracleNonSysUsername field if non-nil, zero value otherwise.

### GetOracleNonSysUsernameOk

`func (o *LinkDSourceParameters) GetOracleNonSysUsernameOk() (*string, bool)`

GetOracleNonSysUsernameOk returns a tuple with the OracleNonSysUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleNonSysUsername

`func (o *LinkDSourceParameters) SetOracleNonSysUsername(v string)`

SetOracleNonSysUsername sets OracleNonSysUsername field to given value.

### HasOracleNonSysUsername

`func (o *LinkDSourceParameters) HasOracleNonSysUsername() bool`

HasOracleNonSysUsername returns a boolean if a field has been set.

### GetOracleNonSysPassword

`func (o *LinkDSourceParameters) GetOracleNonSysPassword() string`

GetOracleNonSysPassword returns the OracleNonSysPassword field if non-nil, zero value otherwise.

### GetOracleNonSysPasswordOk

`func (o *LinkDSourceParameters) GetOracleNonSysPasswordOk() (*string, bool)`

GetOracleNonSysPasswordOk returns a tuple with the OracleNonSysPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleNonSysPassword

`func (o *LinkDSourceParameters) SetOracleNonSysPassword(v string)`

SetOracleNonSysPassword sets OracleNonSysPassword field to given value.

### HasOracleNonSysPassword

`func (o *LinkDSourceParameters) HasOracleNonSysPassword() bool`

HasOracleNonSysPassword returns a boolean if a field has been set.

### GetOracleNonSysVault

`func (o *LinkDSourceParameters) GetOracleNonSysVault() string`

GetOracleNonSysVault returns the OracleNonSysVault field if non-nil, zero value otherwise.

### GetOracleNonSysVaultOk

`func (o *LinkDSourceParameters) GetOracleNonSysVaultOk() (*string, bool)`

GetOracleNonSysVaultOk returns a tuple with the OracleNonSysVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleNonSysVault

`func (o *LinkDSourceParameters) SetOracleNonSysVault(v string)`

SetOracleNonSysVault sets OracleNonSysVault field to given value.

### HasOracleNonSysVault

`func (o *LinkDSourceParameters) HasOracleNonSysVault() bool`

HasOracleNonSysVault returns a boolean if a field has been set.

### GetOracleNonSysHashicorpVaultEngine

`func (o *LinkDSourceParameters) GetOracleNonSysHashicorpVaultEngine() string`

GetOracleNonSysHashicorpVaultEngine returns the OracleNonSysHashicorpVaultEngine field if non-nil, zero value otherwise.

### GetOracleNonSysHashicorpVaultEngineOk

`func (o *LinkDSourceParameters) GetOracleNonSysHashicorpVaultEngineOk() (*string, bool)`

GetOracleNonSysHashicorpVaultEngineOk returns a tuple with the OracleNonSysHashicorpVaultEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleNonSysHashicorpVaultEngine

`func (o *LinkDSourceParameters) SetOracleNonSysHashicorpVaultEngine(v string)`

SetOracleNonSysHashicorpVaultEngine sets OracleNonSysHashicorpVaultEngine field to given value.

### HasOracleNonSysHashicorpVaultEngine

`func (o *LinkDSourceParameters) HasOracleNonSysHashicorpVaultEngine() bool`

HasOracleNonSysHashicorpVaultEngine returns a boolean if a field has been set.

### GetOracleNonSysHashicorpVaultSecretPath

`func (o *LinkDSourceParameters) GetOracleNonSysHashicorpVaultSecretPath() string`

GetOracleNonSysHashicorpVaultSecretPath returns the OracleNonSysHashicorpVaultSecretPath field if non-nil, zero value otherwise.

### GetOracleNonSysHashicorpVaultSecretPathOk

`func (o *LinkDSourceParameters) GetOracleNonSysHashicorpVaultSecretPathOk() (*string, bool)`

GetOracleNonSysHashicorpVaultSecretPathOk returns a tuple with the OracleNonSysHashicorpVaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleNonSysHashicorpVaultSecretPath

`func (o *LinkDSourceParameters) SetOracleNonSysHashicorpVaultSecretPath(v string)`

SetOracleNonSysHashicorpVaultSecretPath sets OracleNonSysHashicorpVaultSecretPath field to given value.

### HasOracleNonSysHashicorpVaultSecretPath

`func (o *LinkDSourceParameters) HasOracleNonSysHashicorpVaultSecretPath() bool`

HasOracleNonSysHashicorpVaultSecretPath returns a boolean if a field has been set.

### GetOracleNonSysHashicorpVaultUsernameKey

`func (o *LinkDSourceParameters) GetOracleNonSysHashicorpVaultUsernameKey() string`

GetOracleNonSysHashicorpVaultUsernameKey returns the OracleNonSysHashicorpVaultUsernameKey field if non-nil, zero value otherwise.

### GetOracleNonSysHashicorpVaultUsernameKeyOk

`func (o *LinkDSourceParameters) GetOracleNonSysHashicorpVaultUsernameKeyOk() (*string, bool)`

GetOracleNonSysHashicorpVaultUsernameKeyOk returns a tuple with the OracleNonSysHashicorpVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleNonSysHashicorpVaultUsernameKey

`func (o *LinkDSourceParameters) SetOracleNonSysHashicorpVaultUsernameKey(v string)`

SetOracleNonSysHashicorpVaultUsernameKey sets OracleNonSysHashicorpVaultUsernameKey field to given value.

### HasOracleNonSysHashicorpVaultUsernameKey

`func (o *LinkDSourceParameters) HasOracleNonSysHashicorpVaultUsernameKey() bool`

HasOracleNonSysHashicorpVaultUsernameKey returns a boolean if a field has been set.

### GetOracleNonSysHashicorpVaultSecretKey

`func (o *LinkDSourceParameters) GetOracleNonSysHashicorpVaultSecretKey() string`

GetOracleNonSysHashicorpVaultSecretKey returns the OracleNonSysHashicorpVaultSecretKey field if non-nil, zero value otherwise.

### GetOracleNonSysHashicorpVaultSecretKeyOk

`func (o *LinkDSourceParameters) GetOracleNonSysHashicorpVaultSecretKeyOk() (*string, bool)`

GetOracleNonSysHashicorpVaultSecretKeyOk returns a tuple with the OracleNonSysHashicorpVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleNonSysHashicorpVaultSecretKey

`func (o *LinkDSourceParameters) SetOracleNonSysHashicorpVaultSecretKey(v string)`

SetOracleNonSysHashicorpVaultSecretKey sets OracleNonSysHashicorpVaultSecretKey field to given value.

### HasOracleNonSysHashicorpVaultSecretKey

`func (o *LinkDSourceParameters) HasOracleNonSysHashicorpVaultSecretKey() bool`

HasOracleNonSysHashicorpVaultSecretKey returns a boolean if a field has been set.

### GetOracleNonSysAzureVaultName

`func (o *LinkDSourceParameters) GetOracleNonSysAzureVaultName() string`

GetOracleNonSysAzureVaultName returns the OracleNonSysAzureVaultName field if non-nil, zero value otherwise.

### GetOracleNonSysAzureVaultNameOk

`func (o *LinkDSourceParameters) GetOracleNonSysAzureVaultNameOk() (*string, bool)`

GetOracleNonSysAzureVaultNameOk returns a tuple with the OracleNonSysAzureVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleNonSysAzureVaultName

`func (o *LinkDSourceParameters) SetOracleNonSysAzureVaultName(v string)`

SetOracleNonSysAzureVaultName sets OracleNonSysAzureVaultName field to given value.

### HasOracleNonSysAzureVaultName

`func (o *LinkDSourceParameters) HasOracleNonSysAzureVaultName() bool`

HasOracleNonSysAzureVaultName returns a boolean if a field has been set.

### GetOracleNonSysAzureVaultUsernameKey

`func (o *LinkDSourceParameters) GetOracleNonSysAzureVaultUsernameKey() string`

GetOracleNonSysAzureVaultUsernameKey returns the OracleNonSysAzureVaultUsernameKey field if non-nil, zero value otherwise.

### GetOracleNonSysAzureVaultUsernameKeyOk

`func (o *LinkDSourceParameters) GetOracleNonSysAzureVaultUsernameKeyOk() (*string, bool)`

GetOracleNonSysAzureVaultUsernameKeyOk returns a tuple with the OracleNonSysAzureVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleNonSysAzureVaultUsernameKey

`func (o *LinkDSourceParameters) SetOracleNonSysAzureVaultUsernameKey(v string)`

SetOracleNonSysAzureVaultUsernameKey sets OracleNonSysAzureVaultUsernameKey field to given value.

### HasOracleNonSysAzureVaultUsernameKey

`func (o *LinkDSourceParameters) HasOracleNonSysAzureVaultUsernameKey() bool`

HasOracleNonSysAzureVaultUsernameKey returns a boolean if a field has been set.

### GetOracleNonSysAzureVaultSecretKey

`func (o *LinkDSourceParameters) GetOracleNonSysAzureVaultSecretKey() string`

GetOracleNonSysAzureVaultSecretKey returns the OracleNonSysAzureVaultSecretKey field if non-nil, zero value otherwise.

### GetOracleNonSysAzureVaultSecretKeyOk

`func (o *LinkDSourceParameters) GetOracleNonSysAzureVaultSecretKeyOk() (*string, bool)`

GetOracleNonSysAzureVaultSecretKeyOk returns a tuple with the OracleNonSysAzureVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleNonSysAzureVaultSecretKey

`func (o *LinkDSourceParameters) SetOracleNonSysAzureVaultSecretKey(v string)`

SetOracleNonSysAzureVaultSecretKey sets OracleNonSysAzureVaultSecretKey field to given value.

### HasOracleNonSysAzureVaultSecretKey

`func (o *LinkDSourceParameters) HasOracleNonSysAzureVaultSecretKey() bool`

HasOracleNonSysAzureVaultSecretKey returns a boolean if a field has been set.

### GetOracleNonSysCyberarkVaultQueryString

`func (o *LinkDSourceParameters) GetOracleNonSysCyberarkVaultQueryString() string`

GetOracleNonSysCyberarkVaultQueryString returns the OracleNonSysCyberarkVaultQueryString field if non-nil, zero value otherwise.

### GetOracleNonSysCyberarkVaultQueryStringOk

`func (o *LinkDSourceParameters) GetOracleNonSysCyberarkVaultQueryStringOk() (*string, bool)`

GetOracleNonSysCyberarkVaultQueryStringOk returns a tuple with the OracleNonSysCyberarkVaultQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleNonSysCyberarkVaultQueryString

`func (o *LinkDSourceParameters) SetOracleNonSysCyberarkVaultQueryString(v string)`

SetOracleNonSysCyberarkVaultQueryString sets OracleNonSysCyberarkVaultQueryString field to given value.

### HasOracleNonSysCyberarkVaultQueryString

`func (o *LinkDSourceParameters) HasOracleNonSysCyberarkVaultQueryString() bool`

HasOracleNonSysCyberarkVaultQueryString returns a boolean if a field has been set.

### GetOracleFallbackUsername

`func (o *LinkDSourceParameters) GetOracleFallbackUsername() string`

GetOracleFallbackUsername returns the OracleFallbackUsername field if non-nil, zero value otherwise.

### GetOracleFallbackUsernameOk

`func (o *LinkDSourceParameters) GetOracleFallbackUsernameOk() (*string, bool)`

GetOracleFallbackUsernameOk returns a tuple with the OracleFallbackUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleFallbackUsername

`func (o *LinkDSourceParameters) SetOracleFallbackUsername(v string)`

SetOracleFallbackUsername sets OracleFallbackUsername field to given value.

### HasOracleFallbackUsername

`func (o *LinkDSourceParameters) HasOracleFallbackUsername() bool`

HasOracleFallbackUsername returns a boolean if a field has been set.

### GetOracleFallbackPassword

`func (o *LinkDSourceParameters) GetOracleFallbackPassword() string`

GetOracleFallbackPassword returns the OracleFallbackPassword field if non-nil, zero value otherwise.

### GetOracleFallbackPasswordOk

`func (o *LinkDSourceParameters) GetOracleFallbackPasswordOk() (*string, bool)`

GetOracleFallbackPasswordOk returns a tuple with the OracleFallbackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleFallbackPassword

`func (o *LinkDSourceParameters) SetOracleFallbackPassword(v string)`

SetOracleFallbackPassword sets OracleFallbackPassword field to given value.

### HasOracleFallbackPassword

`func (o *LinkDSourceParameters) HasOracleFallbackPassword() bool`

HasOracleFallbackPassword returns a boolean if a field has been set.

### GetOracleFallbackVault

`func (o *LinkDSourceParameters) GetOracleFallbackVault() string`

GetOracleFallbackVault returns the OracleFallbackVault field if non-nil, zero value otherwise.

### GetOracleFallbackVaultOk

`func (o *LinkDSourceParameters) GetOracleFallbackVaultOk() (*string, bool)`

GetOracleFallbackVaultOk returns a tuple with the OracleFallbackVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleFallbackVault

`func (o *LinkDSourceParameters) SetOracleFallbackVault(v string)`

SetOracleFallbackVault sets OracleFallbackVault field to given value.

### HasOracleFallbackVault

`func (o *LinkDSourceParameters) HasOracleFallbackVault() bool`

HasOracleFallbackVault returns a boolean if a field has been set.

### GetOracleFallbackHashicorpVaultEngine

`func (o *LinkDSourceParameters) GetOracleFallbackHashicorpVaultEngine() string`

GetOracleFallbackHashicorpVaultEngine returns the OracleFallbackHashicorpVaultEngine field if non-nil, zero value otherwise.

### GetOracleFallbackHashicorpVaultEngineOk

`func (o *LinkDSourceParameters) GetOracleFallbackHashicorpVaultEngineOk() (*string, bool)`

GetOracleFallbackHashicorpVaultEngineOk returns a tuple with the OracleFallbackHashicorpVaultEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleFallbackHashicorpVaultEngine

`func (o *LinkDSourceParameters) SetOracleFallbackHashicorpVaultEngine(v string)`

SetOracleFallbackHashicorpVaultEngine sets OracleFallbackHashicorpVaultEngine field to given value.

### HasOracleFallbackHashicorpVaultEngine

`func (o *LinkDSourceParameters) HasOracleFallbackHashicorpVaultEngine() bool`

HasOracleFallbackHashicorpVaultEngine returns a boolean if a field has been set.

### GetOracleFallbackHashicorpVaultSecretPath

`func (o *LinkDSourceParameters) GetOracleFallbackHashicorpVaultSecretPath() string`

GetOracleFallbackHashicorpVaultSecretPath returns the OracleFallbackHashicorpVaultSecretPath field if non-nil, zero value otherwise.

### GetOracleFallbackHashicorpVaultSecretPathOk

`func (o *LinkDSourceParameters) GetOracleFallbackHashicorpVaultSecretPathOk() (*string, bool)`

GetOracleFallbackHashicorpVaultSecretPathOk returns a tuple with the OracleFallbackHashicorpVaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleFallbackHashicorpVaultSecretPath

`func (o *LinkDSourceParameters) SetOracleFallbackHashicorpVaultSecretPath(v string)`

SetOracleFallbackHashicorpVaultSecretPath sets OracleFallbackHashicorpVaultSecretPath field to given value.

### HasOracleFallbackHashicorpVaultSecretPath

`func (o *LinkDSourceParameters) HasOracleFallbackHashicorpVaultSecretPath() bool`

HasOracleFallbackHashicorpVaultSecretPath returns a boolean if a field has been set.

### GetOracleFallbackHashicorpVaultUsernameKey

`func (o *LinkDSourceParameters) GetOracleFallbackHashicorpVaultUsernameKey() string`

GetOracleFallbackHashicorpVaultUsernameKey returns the OracleFallbackHashicorpVaultUsernameKey field if non-nil, zero value otherwise.

### GetOracleFallbackHashicorpVaultUsernameKeyOk

`func (o *LinkDSourceParameters) GetOracleFallbackHashicorpVaultUsernameKeyOk() (*string, bool)`

GetOracleFallbackHashicorpVaultUsernameKeyOk returns a tuple with the OracleFallbackHashicorpVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleFallbackHashicorpVaultUsernameKey

`func (o *LinkDSourceParameters) SetOracleFallbackHashicorpVaultUsernameKey(v string)`

SetOracleFallbackHashicorpVaultUsernameKey sets OracleFallbackHashicorpVaultUsernameKey field to given value.

### HasOracleFallbackHashicorpVaultUsernameKey

`func (o *LinkDSourceParameters) HasOracleFallbackHashicorpVaultUsernameKey() bool`

HasOracleFallbackHashicorpVaultUsernameKey returns a boolean if a field has been set.

### GetOracleFallbackHashicorpVaultSecretKey

`func (o *LinkDSourceParameters) GetOracleFallbackHashicorpVaultSecretKey() string`

GetOracleFallbackHashicorpVaultSecretKey returns the OracleFallbackHashicorpVaultSecretKey field if non-nil, zero value otherwise.

### GetOracleFallbackHashicorpVaultSecretKeyOk

`func (o *LinkDSourceParameters) GetOracleFallbackHashicorpVaultSecretKeyOk() (*string, bool)`

GetOracleFallbackHashicorpVaultSecretKeyOk returns a tuple with the OracleFallbackHashicorpVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleFallbackHashicorpVaultSecretKey

`func (o *LinkDSourceParameters) SetOracleFallbackHashicorpVaultSecretKey(v string)`

SetOracleFallbackHashicorpVaultSecretKey sets OracleFallbackHashicorpVaultSecretKey field to given value.

### HasOracleFallbackHashicorpVaultSecretKey

`func (o *LinkDSourceParameters) HasOracleFallbackHashicorpVaultSecretKey() bool`

HasOracleFallbackHashicorpVaultSecretKey returns a boolean if a field has been set.

### GetOracleFallbackAzureVaultName

`func (o *LinkDSourceParameters) GetOracleFallbackAzureVaultName() string`

GetOracleFallbackAzureVaultName returns the OracleFallbackAzureVaultName field if non-nil, zero value otherwise.

### GetOracleFallbackAzureVaultNameOk

`func (o *LinkDSourceParameters) GetOracleFallbackAzureVaultNameOk() (*string, bool)`

GetOracleFallbackAzureVaultNameOk returns a tuple with the OracleFallbackAzureVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleFallbackAzureVaultName

`func (o *LinkDSourceParameters) SetOracleFallbackAzureVaultName(v string)`

SetOracleFallbackAzureVaultName sets OracleFallbackAzureVaultName field to given value.

### HasOracleFallbackAzureVaultName

`func (o *LinkDSourceParameters) HasOracleFallbackAzureVaultName() bool`

HasOracleFallbackAzureVaultName returns a boolean if a field has been set.

### GetOracleFallbackAzureVaultUsernameKey

`func (o *LinkDSourceParameters) GetOracleFallbackAzureVaultUsernameKey() string`

GetOracleFallbackAzureVaultUsernameKey returns the OracleFallbackAzureVaultUsernameKey field if non-nil, zero value otherwise.

### GetOracleFallbackAzureVaultUsernameKeyOk

`func (o *LinkDSourceParameters) GetOracleFallbackAzureVaultUsernameKeyOk() (*string, bool)`

GetOracleFallbackAzureVaultUsernameKeyOk returns a tuple with the OracleFallbackAzureVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleFallbackAzureVaultUsernameKey

`func (o *LinkDSourceParameters) SetOracleFallbackAzureVaultUsernameKey(v string)`

SetOracleFallbackAzureVaultUsernameKey sets OracleFallbackAzureVaultUsernameKey field to given value.

### HasOracleFallbackAzureVaultUsernameKey

`func (o *LinkDSourceParameters) HasOracleFallbackAzureVaultUsernameKey() bool`

HasOracleFallbackAzureVaultUsernameKey returns a boolean if a field has been set.

### GetOracleFallbackAzureVaultSecretKey

`func (o *LinkDSourceParameters) GetOracleFallbackAzureVaultSecretKey() string`

GetOracleFallbackAzureVaultSecretKey returns the OracleFallbackAzureVaultSecretKey field if non-nil, zero value otherwise.

### GetOracleFallbackAzureVaultSecretKeyOk

`func (o *LinkDSourceParameters) GetOracleFallbackAzureVaultSecretKeyOk() (*string, bool)`

GetOracleFallbackAzureVaultSecretKeyOk returns a tuple with the OracleFallbackAzureVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleFallbackAzureVaultSecretKey

`func (o *LinkDSourceParameters) SetOracleFallbackAzureVaultSecretKey(v string)`

SetOracleFallbackAzureVaultSecretKey sets OracleFallbackAzureVaultSecretKey field to given value.

### HasOracleFallbackAzureVaultSecretKey

`func (o *LinkDSourceParameters) HasOracleFallbackAzureVaultSecretKey() bool`

HasOracleFallbackAzureVaultSecretKey returns a boolean if a field has been set.

### GetOracleFallbackCyberarkVaultQueryString

`func (o *LinkDSourceParameters) GetOracleFallbackCyberarkVaultQueryString() string`

GetOracleFallbackCyberarkVaultQueryString returns the OracleFallbackCyberarkVaultQueryString field if non-nil, zero value otherwise.

### GetOracleFallbackCyberarkVaultQueryStringOk

`func (o *LinkDSourceParameters) GetOracleFallbackCyberarkVaultQueryStringOk() (*string, bool)`

GetOracleFallbackCyberarkVaultQueryStringOk returns a tuple with the OracleFallbackCyberarkVaultQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleFallbackCyberarkVaultQueryString

`func (o *LinkDSourceParameters) SetOracleFallbackCyberarkVaultQueryString(v string)`

SetOracleFallbackCyberarkVaultQueryString sets OracleFallbackCyberarkVaultQueryString field to given value.

### HasOracleFallbackCyberarkVaultQueryString

`func (o *LinkDSourceParameters) HasOracleFallbackCyberarkVaultQueryString() bool`

HasOracleFallbackCyberarkVaultQueryString returns a boolean if a field has been set.

### GetBackupLevelEnabled

`func (o *LinkDSourceParameters) GetBackupLevelEnabled() bool`

GetBackupLevelEnabled returns the BackupLevelEnabled field if non-nil, zero value otherwise.

### GetBackupLevelEnabledOk

`func (o *LinkDSourceParameters) GetBackupLevelEnabledOk() (*bool, bool)`

GetBackupLevelEnabledOk returns a tuple with the BackupLevelEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLevelEnabled

`func (o *LinkDSourceParameters) SetBackupLevelEnabled(v bool)`

SetBackupLevelEnabled sets BackupLevelEnabled field to given value.

### HasBackupLevelEnabled

`func (o *LinkDSourceParameters) HasBackupLevelEnabled() bool`

HasBackupLevelEnabled returns a boolean if a field has been set.

### GetRmanChannels

`func (o *LinkDSourceParameters) GetRmanChannels() int32`

GetRmanChannels returns the RmanChannels field if non-nil, zero value otherwise.

### GetRmanChannelsOk

`func (o *LinkDSourceParameters) GetRmanChannelsOk() (*int32, bool)`

GetRmanChannelsOk returns a tuple with the RmanChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmanChannels

`func (o *LinkDSourceParameters) SetRmanChannels(v int32)`

SetRmanChannels sets RmanChannels field to given value.

### HasRmanChannels

`func (o *LinkDSourceParameters) HasRmanChannels() bool`

HasRmanChannels returns a boolean if a field has been set.

### GetFilesPerSet

`func (o *LinkDSourceParameters) GetFilesPerSet() int32`

GetFilesPerSet returns the FilesPerSet field if non-nil, zero value otherwise.

### GetFilesPerSetOk

`func (o *LinkDSourceParameters) GetFilesPerSetOk() (*int32, bool)`

GetFilesPerSetOk returns a tuple with the FilesPerSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesPerSet

`func (o *LinkDSourceParameters) SetFilesPerSet(v int32)`

SetFilesPerSet sets FilesPerSet field to given value.

### HasFilesPerSet

`func (o *LinkDSourceParameters) HasFilesPerSet() bool`

HasFilesPerSet returns a boolean if a field has been set.

### GetCheckLogical

`func (o *LinkDSourceParameters) GetCheckLogical() bool`

GetCheckLogical returns the CheckLogical field if non-nil, zero value otherwise.

### GetCheckLogicalOk

`func (o *LinkDSourceParameters) GetCheckLogicalOk() (*bool, bool)`

GetCheckLogicalOk returns a tuple with the CheckLogical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckLogical

`func (o *LinkDSourceParameters) SetCheckLogical(v bool)`

SetCheckLogical sets CheckLogical field to given value.

### HasCheckLogical

`func (o *LinkDSourceParameters) HasCheckLogical() bool`

HasCheckLogical returns a boolean if a field has been set.

### GetEncryptedLinkingEnabled

`func (o *LinkDSourceParameters) GetEncryptedLinkingEnabled() bool`

GetEncryptedLinkingEnabled returns the EncryptedLinkingEnabled field if non-nil, zero value otherwise.

### GetEncryptedLinkingEnabledOk

`func (o *LinkDSourceParameters) GetEncryptedLinkingEnabledOk() (*bool, bool)`

GetEncryptedLinkingEnabledOk returns a tuple with the EncryptedLinkingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedLinkingEnabled

`func (o *LinkDSourceParameters) SetEncryptedLinkingEnabled(v bool)`

SetEncryptedLinkingEnabled sets EncryptedLinkingEnabled field to given value.

### HasEncryptedLinkingEnabled

`func (o *LinkDSourceParameters) HasEncryptedLinkingEnabled() bool`

HasEncryptedLinkingEnabled returns a boolean if a field has been set.

### GetCompressedLinkingEnabled

`func (o *LinkDSourceParameters) GetCompressedLinkingEnabled() bool`

GetCompressedLinkingEnabled returns the CompressedLinkingEnabled field if non-nil, zero value otherwise.

### GetCompressedLinkingEnabledOk

`func (o *LinkDSourceParameters) GetCompressedLinkingEnabledOk() (*bool, bool)`

GetCompressedLinkingEnabledOk returns a tuple with the CompressedLinkingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedLinkingEnabled

`func (o *LinkDSourceParameters) SetCompressedLinkingEnabled(v bool)`

SetCompressedLinkingEnabled sets CompressedLinkingEnabled field to given value.

### HasCompressedLinkingEnabled

`func (o *LinkDSourceParameters) HasCompressedLinkingEnabled() bool`

HasCompressedLinkingEnabled returns a boolean if a field has been set.

### GetBandwidthLimit

`func (o *LinkDSourceParameters) GetBandwidthLimit() int32`

GetBandwidthLimit returns the BandwidthLimit field if non-nil, zero value otherwise.

### GetBandwidthLimitOk

`func (o *LinkDSourceParameters) GetBandwidthLimitOk() (*int32, bool)`

GetBandwidthLimitOk returns a tuple with the BandwidthLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimit

`func (o *LinkDSourceParameters) SetBandwidthLimit(v int32)`

SetBandwidthLimit sets BandwidthLimit field to given value.

### HasBandwidthLimit

`func (o *LinkDSourceParameters) HasBandwidthLimit() bool`

HasBandwidthLimit returns a boolean if a field has been set.

### GetNumberOfConnections

`func (o *LinkDSourceParameters) GetNumberOfConnections() int32`

GetNumberOfConnections returns the NumberOfConnections field if non-nil, zero value otherwise.

### GetNumberOfConnectionsOk

`func (o *LinkDSourceParameters) GetNumberOfConnectionsOk() (*int32, bool)`

GetNumberOfConnectionsOk returns a tuple with the NumberOfConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfConnections

`func (o *LinkDSourceParameters) SetNumberOfConnections(v int32)`

SetNumberOfConnections sets NumberOfConnections field to given value.

### HasNumberOfConnections

`func (o *LinkDSourceParameters) HasNumberOfConnections() bool`

HasNumberOfConnections returns a boolean if a field has been set.

### GetDiagnoseNoLoggingFaults

`func (o *LinkDSourceParameters) GetDiagnoseNoLoggingFaults() bool`

GetDiagnoseNoLoggingFaults returns the DiagnoseNoLoggingFaults field if non-nil, zero value otherwise.

### GetDiagnoseNoLoggingFaultsOk

`func (o *LinkDSourceParameters) GetDiagnoseNoLoggingFaultsOk() (*bool, bool)`

GetDiagnoseNoLoggingFaultsOk returns a tuple with the DiagnoseNoLoggingFaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnoseNoLoggingFaults

`func (o *LinkDSourceParameters) SetDiagnoseNoLoggingFaults(v bool)`

SetDiagnoseNoLoggingFaults sets DiagnoseNoLoggingFaults field to given value.

### HasDiagnoseNoLoggingFaults

`func (o *LinkDSourceParameters) HasDiagnoseNoLoggingFaults() bool`

HasDiagnoseNoLoggingFaults returns a boolean if a field has been set.

### GetPreProvisioningEnabled

`func (o *LinkDSourceParameters) GetPreProvisioningEnabled() bool`

GetPreProvisioningEnabled returns the PreProvisioningEnabled field if non-nil, zero value otherwise.

### GetPreProvisioningEnabledOk

`func (o *LinkDSourceParameters) GetPreProvisioningEnabledOk() (*bool, bool)`

GetPreProvisioningEnabledOk returns a tuple with the PreProvisioningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreProvisioningEnabled

`func (o *LinkDSourceParameters) SetPreProvisioningEnabled(v bool)`

SetPreProvisioningEnabled sets PreProvisioningEnabled field to given value.

### HasPreProvisioningEnabled

`func (o *LinkDSourceParameters) HasPreProvisioningEnabled() bool`

HasPreProvisioningEnabled returns a boolean if a field has been set.

### GetLinkNow

`func (o *LinkDSourceParameters) GetLinkNow() bool`

GetLinkNow returns the LinkNow field if non-nil, zero value otherwise.

### GetLinkNowOk

`func (o *LinkDSourceParameters) GetLinkNowOk() (*bool, bool)`

GetLinkNowOk returns a tuple with the LinkNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNow

`func (o *LinkDSourceParameters) SetLinkNow(v bool)`

SetLinkNow sets LinkNow field to given value.

### HasLinkNow

`func (o *LinkDSourceParameters) HasLinkNow() bool`

HasLinkNow returns a boolean if a field has been set.

### GetForceFullBackup

`func (o *LinkDSourceParameters) GetForceFullBackup() bool`

GetForceFullBackup returns the ForceFullBackup field if non-nil, zero value otherwise.

### GetForceFullBackupOk

`func (o *LinkDSourceParameters) GetForceFullBackupOk() (*bool, bool)`

GetForceFullBackupOk returns a tuple with the ForceFullBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceFullBackup

`func (o *LinkDSourceParameters) SetForceFullBackup(v bool)`

SetForceFullBackup sets ForceFullBackup field to given value.

### HasForceFullBackup

`func (o *LinkDSourceParameters) HasForceFullBackup() bool`

HasForceFullBackup returns a boolean if a field has been set.

### GetDoubleSync

`func (o *LinkDSourceParameters) GetDoubleSync() bool`

GetDoubleSync returns the DoubleSync field if non-nil, zero value otherwise.

### GetDoubleSyncOk

`func (o *LinkDSourceParameters) GetDoubleSyncOk() (*bool, bool)`

GetDoubleSyncOk returns a tuple with the DoubleSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleSync

`func (o *LinkDSourceParameters) SetDoubleSync(v bool)`

SetDoubleSync sets DoubleSync field to given value.

### HasDoubleSync

`func (o *LinkDSourceParameters) HasDoubleSync() bool`

HasDoubleSync returns a boolean if a field has been set.

### GetSkipSpaceCheck

`func (o *LinkDSourceParameters) GetSkipSpaceCheck() bool`

GetSkipSpaceCheck returns the SkipSpaceCheck field if non-nil, zero value otherwise.

### GetSkipSpaceCheckOk

`func (o *LinkDSourceParameters) GetSkipSpaceCheckOk() (*bool, bool)`

GetSkipSpaceCheckOk returns a tuple with the SkipSpaceCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSpaceCheck

`func (o *LinkDSourceParameters) SetSkipSpaceCheck(v bool)`

SetSkipSpaceCheck sets SkipSpaceCheck field to given value.

### HasSkipSpaceCheck

`func (o *LinkDSourceParameters) HasSkipSpaceCheck() bool`

HasSkipSpaceCheck returns a boolean if a field has been set.

### GetDoNotResume

`func (o *LinkDSourceParameters) GetDoNotResume() bool`

GetDoNotResume returns the DoNotResume field if non-nil, zero value otherwise.

### GetDoNotResumeOk

`func (o *LinkDSourceParameters) GetDoNotResumeOk() (*bool, bool)`

GetDoNotResumeOk returns a tuple with the DoNotResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotResume

`func (o *LinkDSourceParameters) SetDoNotResume(v bool)`

SetDoNotResume sets DoNotResume field to given value.

### HasDoNotResume

`func (o *LinkDSourceParameters) HasDoNotResume() bool`

HasDoNotResume returns a boolean if a field has been set.

### GetFilesForFullBackup

`func (o *LinkDSourceParameters) GetFilesForFullBackup() []int32`

GetFilesForFullBackup returns the FilesForFullBackup field if non-nil, zero value otherwise.

### GetFilesForFullBackupOk

`func (o *LinkDSourceParameters) GetFilesForFullBackupOk() (*[]int32, bool)`

GetFilesForFullBackupOk returns a tuple with the FilesForFullBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesForFullBackup

`func (o *LinkDSourceParameters) SetFilesForFullBackup(v []int32)`

SetFilesForFullBackup sets FilesForFullBackup field to given value.

### HasFilesForFullBackup

`func (o *LinkDSourceParameters) HasFilesForFullBackup() bool`

HasFilesForFullBackup returns a boolean if a field has been set.

### GetLogSyncMode

`func (o *LinkDSourceParameters) GetLogSyncMode() string`

GetLogSyncMode returns the LogSyncMode field if non-nil, zero value otherwise.

### GetLogSyncModeOk

`func (o *LinkDSourceParameters) GetLogSyncModeOk() (*string, bool)`

GetLogSyncModeOk returns a tuple with the LogSyncMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSyncMode

`func (o *LinkDSourceParameters) SetLogSyncMode(v string)`

SetLogSyncMode sets LogSyncMode field to given value.

### HasLogSyncMode

`func (o *LinkDSourceParameters) HasLogSyncMode() bool`

HasLogSyncMode returns a boolean if a field has been set.

### GetLogSyncInterval

`func (o *LinkDSourceParameters) GetLogSyncInterval() int32`

GetLogSyncInterval returns the LogSyncInterval field if non-nil, zero value otherwise.

### GetLogSyncIntervalOk

`func (o *LinkDSourceParameters) GetLogSyncIntervalOk() (*int32, bool)`

GetLogSyncIntervalOk returns a tuple with the LogSyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSyncInterval

`func (o *LinkDSourceParameters) SetLogSyncInterval(v int32)`

SetLogSyncInterval sets LogSyncInterval field to given value.

### HasLogSyncInterval

`func (o *LinkDSourceParameters) HasLogSyncInterval() bool`

HasLogSyncInterval returns a boolean if a field has been set.

### GetLogSyncEnabled

`func (o *LinkDSourceParameters) GetLogSyncEnabled() bool`

GetLogSyncEnabled returns the LogSyncEnabled field if non-nil, zero value otherwise.

### GetLogSyncEnabledOk

`func (o *LinkDSourceParameters) GetLogSyncEnabledOk() (*bool, bool)`

GetLogSyncEnabledOk returns a tuple with the LogSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSyncEnabled

`func (o *LinkDSourceParameters) SetLogSyncEnabled(v bool)`

SetLogSyncEnabled sets LogSyncEnabled field to given value.

### HasLogSyncEnabled

`func (o *LinkDSourceParameters) HasLogSyncEnabled() bool`

HasLogSyncEnabled returns a boolean if a field has been set.

### GetExternalFilePath

`func (o *LinkDSourceParameters) GetExternalFilePath() string`

GetExternalFilePath returns the ExternalFilePath field if non-nil, zero value otherwise.

### GetExternalFilePathOk

`func (o *LinkDSourceParameters) GetExternalFilePathOk() (*string, bool)`

GetExternalFilePathOk returns a tuple with the ExternalFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFilePath

`func (o *LinkDSourceParameters) SetExternalFilePath(v string)`

SetExternalFilePath sets ExternalFilePath field to given value.

### HasExternalFilePath

`func (o *LinkDSourceParameters) HasExternalFilePath() bool`

HasExternalFilePath returns a boolean if a field has been set.

### GetMakeCurrentAccountOwner

`func (o *LinkDSourceParameters) GetMakeCurrentAccountOwner() bool`

GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field if non-nil, zero value otherwise.

### GetMakeCurrentAccountOwnerOk

`func (o *LinkDSourceParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool)`

GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCurrentAccountOwner

`func (o *LinkDSourceParameters) SetMakeCurrentAccountOwner(v bool)`

SetMakeCurrentAccountOwner sets MakeCurrentAccountOwner field to given value.

### HasMakeCurrentAccountOwner

`func (o *LinkDSourceParameters) HasMakeCurrentAccountOwner() bool`

HasMakeCurrentAccountOwner returns a boolean if a field has been set.

### GetTags

`func (o *LinkDSourceParameters) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LinkDSourceParameters) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LinkDSourceParameters) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LinkDSourceParameters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOpsPreSync

`func (o *LinkDSourceParameters) GetOpsPreSync() []SourceOperation`

GetOpsPreSync returns the OpsPreSync field if non-nil, zero value otherwise.

### GetOpsPreSyncOk

`func (o *LinkDSourceParameters) GetOpsPreSyncOk() (*[]SourceOperation, bool)`

GetOpsPreSyncOk returns a tuple with the OpsPreSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsPreSync

`func (o *LinkDSourceParameters) SetOpsPreSync(v []SourceOperation)`

SetOpsPreSync sets OpsPreSync field to given value.

### HasOpsPreSync

`func (o *LinkDSourceParameters) HasOpsPreSync() bool`

HasOpsPreSync returns a boolean if a field has been set.

### GetOpsPreLogSync

`func (o *LinkDSourceParameters) GetOpsPreLogSync() []SourceOperation`

GetOpsPreLogSync returns the OpsPreLogSync field if non-nil, zero value otherwise.

### GetOpsPreLogSyncOk

`func (o *LinkDSourceParameters) GetOpsPreLogSyncOk() (*[]SourceOperation, bool)`

GetOpsPreLogSyncOk returns a tuple with the OpsPreLogSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsPreLogSync

`func (o *LinkDSourceParameters) SetOpsPreLogSync(v []SourceOperation)`

SetOpsPreLogSync sets OpsPreLogSync field to given value.

### HasOpsPreLogSync

`func (o *LinkDSourceParameters) HasOpsPreLogSync() bool`

HasOpsPreLogSync returns a boolean if a field has been set.

### GetOpsPostSync

`func (o *LinkDSourceParameters) GetOpsPostSync() []SourceOperation`

GetOpsPostSync returns the OpsPostSync field if non-nil, zero value otherwise.

### GetOpsPostSyncOk

`func (o *LinkDSourceParameters) GetOpsPostSyncOk() (*[]SourceOperation, bool)`

GetOpsPostSyncOk returns a tuple with the OpsPostSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsPostSync

`func (o *LinkDSourceParameters) SetOpsPostSync(v []SourceOperation)`

SetOpsPostSync sets OpsPostSync field to given value.

### HasOpsPostSync

`func (o *LinkDSourceParameters) HasOpsPostSync() bool`

HasOpsPostSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


