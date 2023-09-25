# ASEDSourceLinkSourceParametersAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalFilePath** | Pointer to **string** | External file path. | [optional] 
**MountBase** | Pointer to **string** | The base mount point to use for the NFS mounts. | [optional] 
**LoadBackupPath** | Pointer to **string** | Source database backup location. | [optional] 
**BackupServerName** | Pointer to **string** | Name of the backup server instance. | [optional] 
**BackupHostUser** | Pointer to **string** | OS user for the host where the backup server is located. | [optional] 
**BackupHost** | Pointer to **string** | Host environment where the backup server is located. | [optional] 
**DumpCredentials** | Pointer to **string** | The password credential for the source DB user. | [optional] 
**SourceHostUser** | Pointer to **string** | ID or user reference of the host OS user to use for linking. | [optional] 
**DbUser** | Pointer to **string** | The user name for the source DB user. | [optional] 
**DbPassword** | Pointer to **string** | Password for the database user. | [optional] 
**DbVault** | Pointer to **string** | The name or reference of the vault from which to read the database credentials. | [optional] 
**DbHashicorpVaultEngine** | Pointer to **string** | Vault engine name where the credential is stored. | [optional] 
**DbHashicorpVaultSecretPath** | Pointer to **string** | Path in the vault engine where the credential is stored. | [optional] 
**DbHashicorpVaultUsernameKey** | Pointer to **string** | Hashicorp vault key for the username in the key-value store. | [optional] 
**DbHashicorpVaultSecretKey** | Pointer to **string** | Hashicorp vault key for the password in the key-value store. | [optional] 
**DbAzureVaultName** | Pointer to **string** | Azure key vault name. | [optional] 
**DbAzureVaultUsernameKey** | Pointer to **string** | Azure vault key for the username in the key-value store. | [optional] 
**DbAzureVaultSecretKey** | Pointer to **string** | Azure vault key for the password in the key-value store. | [optional] 
**DbCyberarkVaultQueryString** | Pointer to **string** | Query to find a credential in the CyberArk vault. | [optional] 
**StagingRepository** | Pointer to **string** | The SAP ASE instance on the staging environment that we want to use for validated sync. | [optional] 
**StagingHostUser** | Pointer to **string** | Information about the host OS user on the staging environment to use for linking. | [optional] 
**ValidatedSyncMode** | Pointer to **string** | Information about the host OS user on the staging environment to use for linking. | [optional] [default to "ENABLED"]
**DumpHistoryFileEnabled** | Pointer to **bool** | Specifies if Dump History File is enabled for backup history detection. | [optional] [default to false]
**DropAndRecreateDevices** | Pointer to **bool** | If this parameter is set to true, it will drop the older devices and create new devices during manual sync operations instead of trying to remap the devices. This might increase the space utilization on Delphix Engine. | [optional] [default to false]
**SyncStrategy** | Pointer to **string** | Determines how the Delphix Engine will take a backup: * &#x60;latest_backup&#x60; - Use the most recent backup. * &#x60;new_backup&#x60; - Delphix will take a new backup of your source database. * &#x60;specific_backup&#x60; - Use a specific backup. Using this option requires setting &#x60;ase_backup_files&#x60;. Default is &#x60;new_backup&#x60;.  | [optional] [default to "new_backup"]
**AseBackupFiles** | Pointer to **[]string** | The location of the full backup of the source database to restore from. The backup should be present in the shared backup location for the source database. | [optional] 
**PreValidatedSync** | Pointer to [**[]SourceOperation**](SourceOperation.md) | Operations to perform on the staging source before performing a validated sync. | [optional] 
**PostValidatedSync** | Pointer to [**[]SourceOperation**](SourceOperation.md) | Operations to perform on the staging source after performing a validated sync. | [optional] 

## Methods

### NewASEDSourceLinkSourceParametersAllOf

`func NewASEDSourceLinkSourceParametersAllOf() *ASEDSourceLinkSourceParametersAllOf`

NewASEDSourceLinkSourceParametersAllOf instantiates a new ASEDSourceLinkSourceParametersAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewASEDSourceLinkSourceParametersAllOfWithDefaults

`func NewASEDSourceLinkSourceParametersAllOfWithDefaults() *ASEDSourceLinkSourceParametersAllOf`

NewASEDSourceLinkSourceParametersAllOfWithDefaults instantiates a new ASEDSourceLinkSourceParametersAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalFilePath

`func (o *ASEDSourceLinkSourceParametersAllOf) GetExternalFilePath() string`

GetExternalFilePath returns the ExternalFilePath field if non-nil, zero value otherwise.

### GetExternalFilePathOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetExternalFilePathOk() (*string, bool)`

GetExternalFilePathOk returns a tuple with the ExternalFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFilePath

`func (o *ASEDSourceLinkSourceParametersAllOf) SetExternalFilePath(v string)`

SetExternalFilePath sets ExternalFilePath field to given value.

### HasExternalFilePath

`func (o *ASEDSourceLinkSourceParametersAllOf) HasExternalFilePath() bool`

HasExternalFilePath returns a boolean if a field has been set.

### GetMountBase

`func (o *ASEDSourceLinkSourceParametersAllOf) GetMountBase() string`

GetMountBase returns the MountBase field if non-nil, zero value otherwise.

### GetMountBaseOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetMountBaseOk() (*string, bool)`

GetMountBaseOk returns a tuple with the MountBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountBase

`func (o *ASEDSourceLinkSourceParametersAllOf) SetMountBase(v string)`

SetMountBase sets MountBase field to given value.

### HasMountBase

`func (o *ASEDSourceLinkSourceParametersAllOf) HasMountBase() bool`

HasMountBase returns a boolean if a field has been set.

### GetLoadBackupPath

`func (o *ASEDSourceLinkSourceParametersAllOf) GetLoadBackupPath() string`

GetLoadBackupPath returns the LoadBackupPath field if non-nil, zero value otherwise.

### GetLoadBackupPathOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetLoadBackupPathOk() (*string, bool)`

GetLoadBackupPathOk returns a tuple with the LoadBackupPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBackupPath

`func (o *ASEDSourceLinkSourceParametersAllOf) SetLoadBackupPath(v string)`

SetLoadBackupPath sets LoadBackupPath field to given value.

### HasLoadBackupPath

`func (o *ASEDSourceLinkSourceParametersAllOf) HasLoadBackupPath() bool`

HasLoadBackupPath returns a boolean if a field has been set.

### GetBackupServerName

`func (o *ASEDSourceLinkSourceParametersAllOf) GetBackupServerName() string`

GetBackupServerName returns the BackupServerName field if non-nil, zero value otherwise.

### GetBackupServerNameOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetBackupServerNameOk() (*string, bool)`

GetBackupServerNameOk returns a tuple with the BackupServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupServerName

`func (o *ASEDSourceLinkSourceParametersAllOf) SetBackupServerName(v string)`

SetBackupServerName sets BackupServerName field to given value.

### HasBackupServerName

`func (o *ASEDSourceLinkSourceParametersAllOf) HasBackupServerName() bool`

HasBackupServerName returns a boolean if a field has been set.

### GetBackupHostUser

`func (o *ASEDSourceLinkSourceParametersAllOf) GetBackupHostUser() string`

GetBackupHostUser returns the BackupHostUser field if non-nil, zero value otherwise.

### GetBackupHostUserOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetBackupHostUserOk() (*string, bool)`

GetBackupHostUserOk returns a tuple with the BackupHostUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupHostUser

`func (o *ASEDSourceLinkSourceParametersAllOf) SetBackupHostUser(v string)`

SetBackupHostUser sets BackupHostUser field to given value.

### HasBackupHostUser

`func (o *ASEDSourceLinkSourceParametersAllOf) HasBackupHostUser() bool`

HasBackupHostUser returns a boolean if a field has been set.

### GetBackupHost

`func (o *ASEDSourceLinkSourceParametersAllOf) GetBackupHost() string`

GetBackupHost returns the BackupHost field if non-nil, zero value otherwise.

### GetBackupHostOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetBackupHostOk() (*string, bool)`

GetBackupHostOk returns a tuple with the BackupHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupHost

`func (o *ASEDSourceLinkSourceParametersAllOf) SetBackupHost(v string)`

SetBackupHost sets BackupHost field to given value.

### HasBackupHost

`func (o *ASEDSourceLinkSourceParametersAllOf) HasBackupHost() bool`

HasBackupHost returns a boolean if a field has been set.

### GetDumpCredentials

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDumpCredentials() string`

GetDumpCredentials returns the DumpCredentials field if non-nil, zero value otherwise.

### GetDumpCredentialsOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDumpCredentialsOk() (*string, bool)`

GetDumpCredentialsOk returns a tuple with the DumpCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDumpCredentials

`func (o *ASEDSourceLinkSourceParametersAllOf) SetDumpCredentials(v string)`

SetDumpCredentials sets DumpCredentials field to given value.

### HasDumpCredentials

`func (o *ASEDSourceLinkSourceParametersAllOf) HasDumpCredentials() bool`

HasDumpCredentials returns a boolean if a field has been set.

### GetSourceHostUser

`func (o *ASEDSourceLinkSourceParametersAllOf) GetSourceHostUser() string`

GetSourceHostUser returns the SourceHostUser field if non-nil, zero value otherwise.

### GetSourceHostUserOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetSourceHostUserOk() (*string, bool)`

GetSourceHostUserOk returns a tuple with the SourceHostUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHostUser

`func (o *ASEDSourceLinkSourceParametersAllOf) SetSourceHostUser(v string)`

SetSourceHostUser sets SourceHostUser field to given value.

### HasSourceHostUser

`func (o *ASEDSourceLinkSourceParametersAllOf) HasSourceHostUser() bool`

HasSourceHostUser returns a boolean if a field has been set.

### GetDbUser

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbUser() string`

GetDbUser returns the DbUser field if non-nil, zero value otherwise.

### GetDbUserOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbUserOk() (*string, bool)`

GetDbUserOk returns a tuple with the DbUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUser

`func (o *ASEDSourceLinkSourceParametersAllOf) SetDbUser(v string)`

SetDbUser sets DbUser field to given value.

### HasDbUser

`func (o *ASEDSourceLinkSourceParametersAllOf) HasDbUser() bool`

HasDbUser returns a boolean if a field has been set.

### GetDbPassword

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbPassword() string`

GetDbPassword returns the DbPassword field if non-nil, zero value otherwise.

### GetDbPasswordOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbPasswordOk() (*string, bool)`

GetDbPasswordOk returns a tuple with the DbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPassword

`func (o *ASEDSourceLinkSourceParametersAllOf) SetDbPassword(v string)`

SetDbPassword sets DbPassword field to given value.

### HasDbPassword

`func (o *ASEDSourceLinkSourceParametersAllOf) HasDbPassword() bool`

HasDbPassword returns a boolean if a field has been set.

### GetDbVault

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbVault() string`

GetDbVault returns the DbVault field if non-nil, zero value otherwise.

### GetDbVaultOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbVaultOk() (*string, bool)`

GetDbVaultOk returns a tuple with the DbVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbVault

`func (o *ASEDSourceLinkSourceParametersAllOf) SetDbVault(v string)`

SetDbVault sets DbVault field to given value.

### HasDbVault

`func (o *ASEDSourceLinkSourceParametersAllOf) HasDbVault() bool`

HasDbVault returns a boolean if a field has been set.

### GetDbHashicorpVaultEngine

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbHashicorpVaultEngine() string`

GetDbHashicorpVaultEngine returns the DbHashicorpVaultEngine field if non-nil, zero value otherwise.

### GetDbHashicorpVaultEngineOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbHashicorpVaultEngineOk() (*string, bool)`

GetDbHashicorpVaultEngineOk returns a tuple with the DbHashicorpVaultEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHashicorpVaultEngine

`func (o *ASEDSourceLinkSourceParametersAllOf) SetDbHashicorpVaultEngine(v string)`

SetDbHashicorpVaultEngine sets DbHashicorpVaultEngine field to given value.

### HasDbHashicorpVaultEngine

`func (o *ASEDSourceLinkSourceParametersAllOf) HasDbHashicorpVaultEngine() bool`

HasDbHashicorpVaultEngine returns a boolean if a field has been set.

### GetDbHashicorpVaultSecretPath

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbHashicorpVaultSecretPath() string`

GetDbHashicorpVaultSecretPath returns the DbHashicorpVaultSecretPath field if non-nil, zero value otherwise.

### GetDbHashicorpVaultSecretPathOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbHashicorpVaultSecretPathOk() (*string, bool)`

GetDbHashicorpVaultSecretPathOk returns a tuple with the DbHashicorpVaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHashicorpVaultSecretPath

`func (o *ASEDSourceLinkSourceParametersAllOf) SetDbHashicorpVaultSecretPath(v string)`

SetDbHashicorpVaultSecretPath sets DbHashicorpVaultSecretPath field to given value.

### HasDbHashicorpVaultSecretPath

`func (o *ASEDSourceLinkSourceParametersAllOf) HasDbHashicorpVaultSecretPath() bool`

HasDbHashicorpVaultSecretPath returns a boolean if a field has been set.

### GetDbHashicorpVaultUsernameKey

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbHashicorpVaultUsernameKey() string`

GetDbHashicorpVaultUsernameKey returns the DbHashicorpVaultUsernameKey field if non-nil, zero value otherwise.

### GetDbHashicorpVaultUsernameKeyOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbHashicorpVaultUsernameKeyOk() (*string, bool)`

GetDbHashicorpVaultUsernameKeyOk returns a tuple with the DbHashicorpVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHashicorpVaultUsernameKey

`func (o *ASEDSourceLinkSourceParametersAllOf) SetDbHashicorpVaultUsernameKey(v string)`

SetDbHashicorpVaultUsernameKey sets DbHashicorpVaultUsernameKey field to given value.

### HasDbHashicorpVaultUsernameKey

`func (o *ASEDSourceLinkSourceParametersAllOf) HasDbHashicorpVaultUsernameKey() bool`

HasDbHashicorpVaultUsernameKey returns a boolean if a field has been set.

### GetDbHashicorpVaultSecretKey

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbHashicorpVaultSecretKey() string`

GetDbHashicorpVaultSecretKey returns the DbHashicorpVaultSecretKey field if non-nil, zero value otherwise.

### GetDbHashicorpVaultSecretKeyOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbHashicorpVaultSecretKeyOk() (*string, bool)`

GetDbHashicorpVaultSecretKeyOk returns a tuple with the DbHashicorpVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHashicorpVaultSecretKey

`func (o *ASEDSourceLinkSourceParametersAllOf) SetDbHashicorpVaultSecretKey(v string)`

SetDbHashicorpVaultSecretKey sets DbHashicorpVaultSecretKey field to given value.

### HasDbHashicorpVaultSecretKey

`func (o *ASEDSourceLinkSourceParametersAllOf) HasDbHashicorpVaultSecretKey() bool`

HasDbHashicorpVaultSecretKey returns a boolean if a field has been set.

### GetDbAzureVaultName

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbAzureVaultName() string`

GetDbAzureVaultName returns the DbAzureVaultName field if non-nil, zero value otherwise.

### GetDbAzureVaultNameOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbAzureVaultNameOk() (*string, bool)`

GetDbAzureVaultNameOk returns a tuple with the DbAzureVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbAzureVaultName

`func (o *ASEDSourceLinkSourceParametersAllOf) SetDbAzureVaultName(v string)`

SetDbAzureVaultName sets DbAzureVaultName field to given value.

### HasDbAzureVaultName

`func (o *ASEDSourceLinkSourceParametersAllOf) HasDbAzureVaultName() bool`

HasDbAzureVaultName returns a boolean if a field has been set.

### GetDbAzureVaultUsernameKey

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbAzureVaultUsernameKey() string`

GetDbAzureVaultUsernameKey returns the DbAzureVaultUsernameKey field if non-nil, zero value otherwise.

### GetDbAzureVaultUsernameKeyOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbAzureVaultUsernameKeyOk() (*string, bool)`

GetDbAzureVaultUsernameKeyOk returns a tuple with the DbAzureVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbAzureVaultUsernameKey

`func (o *ASEDSourceLinkSourceParametersAllOf) SetDbAzureVaultUsernameKey(v string)`

SetDbAzureVaultUsernameKey sets DbAzureVaultUsernameKey field to given value.

### HasDbAzureVaultUsernameKey

`func (o *ASEDSourceLinkSourceParametersAllOf) HasDbAzureVaultUsernameKey() bool`

HasDbAzureVaultUsernameKey returns a boolean if a field has been set.

### GetDbAzureVaultSecretKey

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbAzureVaultSecretKey() string`

GetDbAzureVaultSecretKey returns the DbAzureVaultSecretKey field if non-nil, zero value otherwise.

### GetDbAzureVaultSecretKeyOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbAzureVaultSecretKeyOk() (*string, bool)`

GetDbAzureVaultSecretKeyOk returns a tuple with the DbAzureVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbAzureVaultSecretKey

`func (o *ASEDSourceLinkSourceParametersAllOf) SetDbAzureVaultSecretKey(v string)`

SetDbAzureVaultSecretKey sets DbAzureVaultSecretKey field to given value.

### HasDbAzureVaultSecretKey

`func (o *ASEDSourceLinkSourceParametersAllOf) HasDbAzureVaultSecretKey() bool`

HasDbAzureVaultSecretKey returns a boolean if a field has been set.

### GetDbCyberarkVaultQueryString

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbCyberarkVaultQueryString() string`

GetDbCyberarkVaultQueryString returns the DbCyberarkVaultQueryString field if non-nil, zero value otherwise.

### GetDbCyberarkVaultQueryStringOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDbCyberarkVaultQueryStringOk() (*string, bool)`

GetDbCyberarkVaultQueryStringOk returns a tuple with the DbCyberarkVaultQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCyberarkVaultQueryString

`func (o *ASEDSourceLinkSourceParametersAllOf) SetDbCyberarkVaultQueryString(v string)`

SetDbCyberarkVaultQueryString sets DbCyberarkVaultQueryString field to given value.

### HasDbCyberarkVaultQueryString

`func (o *ASEDSourceLinkSourceParametersAllOf) HasDbCyberarkVaultQueryString() bool`

HasDbCyberarkVaultQueryString returns a boolean if a field has been set.

### GetStagingRepository

`func (o *ASEDSourceLinkSourceParametersAllOf) GetStagingRepository() string`

GetStagingRepository returns the StagingRepository field if non-nil, zero value otherwise.

### GetStagingRepositoryOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetStagingRepositoryOk() (*string, bool)`

GetStagingRepositoryOk returns a tuple with the StagingRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingRepository

`func (o *ASEDSourceLinkSourceParametersAllOf) SetStagingRepository(v string)`

SetStagingRepository sets StagingRepository field to given value.

### HasStagingRepository

`func (o *ASEDSourceLinkSourceParametersAllOf) HasStagingRepository() bool`

HasStagingRepository returns a boolean if a field has been set.

### GetStagingHostUser

`func (o *ASEDSourceLinkSourceParametersAllOf) GetStagingHostUser() string`

GetStagingHostUser returns the StagingHostUser field if non-nil, zero value otherwise.

### GetStagingHostUserOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetStagingHostUserOk() (*string, bool)`

GetStagingHostUserOk returns a tuple with the StagingHostUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingHostUser

`func (o *ASEDSourceLinkSourceParametersAllOf) SetStagingHostUser(v string)`

SetStagingHostUser sets StagingHostUser field to given value.

### HasStagingHostUser

`func (o *ASEDSourceLinkSourceParametersAllOf) HasStagingHostUser() bool`

HasStagingHostUser returns a boolean if a field has been set.

### GetValidatedSyncMode

`func (o *ASEDSourceLinkSourceParametersAllOf) GetValidatedSyncMode() string`

GetValidatedSyncMode returns the ValidatedSyncMode field if non-nil, zero value otherwise.

### GetValidatedSyncModeOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetValidatedSyncModeOk() (*string, bool)`

GetValidatedSyncModeOk returns a tuple with the ValidatedSyncMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedSyncMode

`func (o *ASEDSourceLinkSourceParametersAllOf) SetValidatedSyncMode(v string)`

SetValidatedSyncMode sets ValidatedSyncMode field to given value.

### HasValidatedSyncMode

`func (o *ASEDSourceLinkSourceParametersAllOf) HasValidatedSyncMode() bool`

HasValidatedSyncMode returns a boolean if a field has been set.

### GetDumpHistoryFileEnabled

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDumpHistoryFileEnabled() bool`

GetDumpHistoryFileEnabled returns the DumpHistoryFileEnabled field if non-nil, zero value otherwise.

### GetDumpHistoryFileEnabledOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDumpHistoryFileEnabledOk() (*bool, bool)`

GetDumpHistoryFileEnabledOk returns a tuple with the DumpHistoryFileEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDumpHistoryFileEnabled

`func (o *ASEDSourceLinkSourceParametersAllOf) SetDumpHistoryFileEnabled(v bool)`

SetDumpHistoryFileEnabled sets DumpHistoryFileEnabled field to given value.

### HasDumpHistoryFileEnabled

`func (o *ASEDSourceLinkSourceParametersAllOf) HasDumpHistoryFileEnabled() bool`

HasDumpHistoryFileEnabled returns a boolean if a field has been set.

### GetDropAndRecreateDevices

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDropAndRecreateDevices() bool`

GetDropAndRecreateDevices returns the DropAndRecreateDevices field if non-nil, zero value otherwise.

### GetDropAndRecreateDevicesOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetDropAndRecreateDevicesOk() (*bool, bool)`

GetDropAndRecreateDevicesOk returns a tuple with the DropAndRecreateDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropAndRecreateDevices

`func (o *ASEDSourceLinkSourceParametersAllOf) SetDropAndRecreateDevices(v bool)`

SetDropAndRecreateDevices sets DropAndRecreateDevices field to given value.

### HasDropAndRecreateDevices

`func (o *ASEDSourceLinkSourceParametersAllOf) HasDropAndRecreateDevices() bool`

HasDropAndRecreateDevices returns a boolean if a field has been set.

### GetSyncStrategy

`func (o *ASEDSourceLinkSourceParametersAllOf) GetSyncStrategy() string`

GetSyncStrategy returns the SyncStrategy field if non-nil, zero value otherwise.

### GetSyncStrategyOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetSyncStrategyOk() (*string, bool)`

GetSyncStrategyOk returns a tuple with the SyncStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStrategy

`func (o *ASEDSourceLinkSourceParametersAllOf) SetSyncStrategy(v string)`

SetSyncStrategy sets SyncStrategy field to given value.

### HasSyncStrategy

`func (o *ASEDSourceLinkSourceParametersAllOf) HasSyncStrategy() bool`

HasSyncStrategy returns a boolean if a field has been set.

### GetAseBackupFiles

`func (o *ASEDSourceLinkSourceParametersAllOf) GetAseBackupFiles() []string`

GetAseBackupFiles returns the AseBackupFiles field if non-nil, zero value otherwise.

### GetAseBackupFilesOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetAseBackupFilesOk() (*[]string, bool)`

GetAseBackupFilesOk returns a tuple with the AseBackupFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseBackupFiles

`func (o *ASEDSourceLinkSourceParametersAllOf) SetAseBackupFiles(v []string)`

SetAseBackupFiles sets AseBackupFiles field to given value.

### HasAseBackupFiles

`func (o *ASEDSourceLinkSourceParametersAllOf) HasAseBackupFiles() bool`

HasAseBackupFiles returns a boolean if a field has been set.

### GetPreValidatedSync

`func (o *ASEDSourceLinkSourceParametersAllOf) GetPreValidatedSync() []SourceOperation`

GetPreValidatedSync returns the PreValidatedSync field if non-nil, zero value otherwise.

### GetPreValidatedSyncOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetPreValidatedSyncOk() (*[]SourceOperation, bool)`

GetPreValidatedSyncOk returns a tuple with the PreValidatedSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreValidatedSync

`func (o *ASEDSourceLinkSourceParametersAllOf) SetPreValidatedSync(v []SourceOperation)`

SetPreValidatedSync sets PreValidatedSync field to given value.

### HasPreValidatedSync

`func (o *ASEDSourceLinkSourceParametersAllOf) HasPreValidatedSync() bool`

HasPreValidatedSync returns a boolean if a field has been set.

### GetPostValidatedSync

`func (o *ASEDSourceLinkSourceParametersAllOf) GetPostValidatedSync() []SourceOperation`

GetPostValidatedSync returns the PostValidatedSync field if non-nil, zero value otherwise.

### GetPostValidatedSyncOk

`func (o *ASEDSourceLinkSourceParametersAllOf) GetPostValidatedSyncOk() (*[]SourceOperation, bool)`

GetPostValidatedSyncOk returns a tuple with the PostValidatedSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostValidatedSync

`func (o *ASEDSourceLinkSourceParametersAllOf) SetPostValidatedSync(v []SourceOperation)`

SetPostValidatedSync sets PostValidatedSync field to given value.

### HasPostValidatedSync

`func (o *ASEDSourceLinkSourceParametersAllOf) HasPostValidatedSync() bool`

HasPostValidatedSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


