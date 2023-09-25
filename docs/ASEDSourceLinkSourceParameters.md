# ASEDSourceLinkSourceParameters

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
**MountBase** | Pointer to **string** | The base mount point to use for the NFS mounts. | [optional] 
**LoadBackupPath** | **string** | Source database backup location. | 
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

### NewASEDSourceLinkSourceParameters

`func NewASEDSourceLinkSourceParameters(sourceId string, loadBackupPath string, ) *ASEDSourceLinkSourceParameters`

NewASEDSourceLinkSourceParameters instantiates a new ASEDSourceLinkSourceParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewASEDSourceLinkSourceParametersWithDefaults

`func NewASEDSourceLinkSourceParametersWithDefaults() *ASEDSourceLinkSourceParameters`

NewASEDSourceLinkSourceParametersWithDefaults instantiates a new ASEDSourceLinkSourceParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ASEDSourceLinkSourceParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ASEDSourceLinkSourceParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ASEDSourceLinkSourceParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ASEDSourceLinkSourceParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceId

`func (o *ASEDSourceLinkSourceParameters) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ASEDSourceLinkSourceParameters) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ASEDSourceLinkSourceParameters) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetGroupId

`func (o *ASEDSourceLinkSourceParameters) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ASEDSourceLinkSourceParameters) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ASEDSourceLinkSourceParameters) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ASEDSourceLinkSourceParameters) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetDescription

`func (o *ASEDSourceLinkSourceParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ASEDSourceLinkSourceParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ASEDSourceLinkSourceParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ASEDSourceLinkSourceParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogSyncEnabled

`func (o *ASEDSourceLinkSourceParameters) GetLogSyncEnabled() bool`

GetLogSyncEnabled returns the LogSyncEnabled field if non-nil, zero value otherwise.

### GetLogSyncEnabledOk

`func (o *ASEDSourceLinkSourceParameters) GetLogSyncEnabledOk() (*bool, bool)`

GetLogSyncEnabledOk returns a tuple with the LogSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSyncEnabled

`func (o *ASEDSourceLinkSourceParameters) SetLogSyncEnabled(v bool)`

SetLogSyncEnabled sets LogSyncEnabled field to given value.

### HasLogSyncEnabled

`func (o *ASEDSourceLinkSourceParameters) HasLogSyncEnabled() bool`

HasLogSyncEnabled returns a boolean if a field has been set.

### GetMakeCurrentAccountOwner

`func (o *ASEDSourceLinkSourceParameters) GetMakeCurrentAccountOwner() bool`

GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field if non-nil, zero value otherwise.

### GetMakeCurrentAccountOwnerOk

`func (o *ASEDSourceLinkSourceParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool)`

GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCurrentAccountOwner

`func (o *ASEDSourceLinkSourceParameters) SetMakeCurrentAccountOwner(v bool)`

SetMakeCurrentAccountOwner sets MakeCurrentAccountOwner field to given value.

### HasMakeCurrentAccountOwner

`func (o *ASEDSourceLinkSourceParameters) HasMakeCurrentAccountOwner() bool`

HasMakeCurrentAccountOwner returns a boolean if a field has been set.

### GetTags

`func (o *ASEDSourceLinkSourceParameters) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ASEDSourceLinkSourceParameters) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ASEDSourceLinkSourceParameters) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ASEDSourceLinkSourceParameters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOpsPreSync

`func (o *ASEDSourceLinkSourceParameters) GetOpsPreSync() []SourceOperation`

GetOpsPreSync returns the OpsPreSync field if non-nil, zero value otherwise.

### GetOpsPreSyncOk

`func (o *ASEDSourceLinkSourceParameters) GetOpsPreSyncOk() (*[]SourceOperation, bool)`

GetOpsPreSyncOk returns a tuple with the OpsPreSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsPreSync

`func (o *ASEDSourceLinkSourceParameters) SetOpsPreSync(v []SourceOperation)`

SetOpsPreSync sets OpsPreSync field to given value.

### HasOpsPreSync

`func (o *ASEDSourceLinkSourceParameters) HasOpsPreSync() bool`

HasOpsPreSync returns a boolean if a field has been set.

### GetOpsPostSync

`func (o *ASEDSourceLinkSourceParameters) GetOpsPostSync() []SourceOperation`

GetOpsPostSync returns the OpsPostSync field if non-nil, zero value otherwise.

### GetOpsPostSyncOk

`func (o *ASEDSourceLinkSourceParameters) GetOpsPostSyncOk() (*[]SourceOperation, bool)`

GetOpsPostSyncOk returns a tuple with the OpsPostSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsPostSync

`func (o *ASEDSourceLinkSourceParameters) SetOpsPostSync(v []SourceOperation)`

SetOpsPostSync sets OpsPostSync field to given value.

### HasOpsPostSync

`func (o *ASEDSourceLinkSourceParameters) HasOpsPostSync() bool`

HasOpsPostSync returns a boolean if a field has been set.

### GetExternalFilePath

`func (o *ASEDSourceLinkSourceParameters) GetExternalFilePath() string`

GetExternalFilePath returns the ExternalFilePath field if non-nil, zero value otherwise.

### GetExternalFilePathOk

`func (o *ASEDSourceLinkSourceParameters) GetExternalFilePathOk() (*string, bool)`

GetExternalFilePathOk returns a tuple with the ExternalFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFilePath

`func (o *ASEDSourceLinkSourceParameters) SetExternalFilePath(v string)`

SetExternalFilePath sets ExternalFilePath field to given value.

### HasExternalFilePath

`func (o *ASEDSourceLinkSourceParameters) HasExternalFilePath() bool`

HasExternalFilePath returns a boolean if a field has been set.

### GetMountBase

`func (o *ASEDSourceLinkSourceParameters) GetMountBase() string`

GetMountBase returns the MountBase field if non-nil, zero value otherwise.

### GetMountBaseOk

`func (o *ASEDSourceLinkSourceParameters) GetMountBaseOk() (*string, bool)`

GetMountBaseOk returns a tuple with the MountBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountBase

`func (o *ASEDSourceLinkSourceParameters) SetMountBase(v string)`

SetMountBase sets MountBase field to given value.

### HasMountBase

`func (o *ASEDSourceLinkSourceParameters) HasMountBase() bool`

HasMountBase returns a boolean if a field has been set.

### GetLoadBackupPath

`func (o *ASEDSourceLinkSourceParameters) GetLoadBackupPath() string`

GetLoadBackupPath returns the LoadBackupPath field if non-nil, zero value otherwise.

### GetLoadBackupPathOk

`func (o *ASEDSourceLinkSourceParameters) GetLoadBackupPathOk() (*string, bool)`

GetLoadBackupPathOk returns a tuple with the LoadBackupPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBackupPath

`func (o *ASEDSourceLinkSourceParameters) SetLoadBackupPath(v string)`

SetLoadBackupPath sets LoadBackupPath field to given value.


### GetBackupServerName

`func (o *ASEDSourceLinkSourceParameters) GetBackupServerName() string`

GetBackupServerName returns the BackupServerName field if non-nil, zero value otherwise.

### GetBackupServerNameOk

`func (o *ASEDSourceLinkSourceParameters) GetBackupServerNameOk() (*string, bool)`

GetBackupServerNameOk returns a tuple with the BackupServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupServerName

`func (o *ASEDSourceLinkSourceParameters) SetBackupServerName(v string)`

SetBackupServerName sets BackupServerName field to given value.

### HasBackupServerName

`func (o *ASEDSourceLinkSourceParameters) HasBackupServerName() bool`

HasBackupServerName returns a boolean if a field has been set.

### GetBackupHostUser

`func (o *ASEDSourceLinkSourceParameters) GetBackupHostUser() string`

GetBackupHostUser returns the BackupHostUser field if non-nil, zero value otherwise.

### GetBackupHostUserOk

`func (o *ASEDSourceLinkSourceParameters) GetBackupHostUserOk() (*string, bool)`

GetBackupHostUserOk returns a tuple with the BackupHostUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupHostUser

`func (o *ASEDSourceLinkSourceParameters) SetBackupHostUser(v string)`

SetBackupHostUser sets BackupHostUser field to given value.

### HasBackupHostUser

`func (o *ASEDSourceLinkSourceParameters) HasBackupHostUser() bool`

HasBackupHostUser returns a boolean if a field has been set.

### GetBackupHost

`func (o *ASEDSourceLinkSourceParameters) GetBackupHost() string`

GetBackupHost returns the BackupHost field if non-nil, zero value otherwise.

### GetBackupHostOk

`func (o *ASEDSourceLinkSourceParameters) GetBackupHostOk() (*string, bool)`

GetBackupHostOk returns a tuple with the BackupHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupHost

`func (o *ASEDSourceLinkSourceParameters) SetBackupHost(v string)`

SetBackupHost sets BackupHost field to given value.

### HasBackupHost

`func (o *ASEDSourceLinkSourceParameters) HasBackupHost() bool`

HasBackupHost returns a boolean if a field has been set.

### GetDumpCredentials

`func (o *ASEDSourceLinkSourceParameters) GetDumpCredentials() string`

GetDumpCredentials returns the DumpCredentials field if non-nil, zero value otherwise.

### GetDumpCredentialsOk

`func (o *ASEDSourceLinkSourceParameters) GetDumpCredentialsOk() (*string, bool)`

GetDumpCredentialsOk returns a tuple with the DumpCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDumpCredentials

`func (o *ASEDSourceLinkSourceParameters) SetDumpCredentials(v string)`

SetDumpCredentials sets DumpCredentials field to given value.

### HasDumpCredentials

`func (o *ASEDSourceLinkSourceParameters) HasDumpCredentials() bool`

HasDumpCredentials returns a boolean if a field has been set.

### GetSourceHostUser

`func (o *ASEDSourceLinkSourceParameters) GetSourceHostUser() string`

GetSourceHostUser returns the SourceHostUser field if non-nil, zero value otherwise.

### GetSourceHostUserOk

`func (o *ASEDSourceLinkSourceParameters) GetSourceHostUserOk() (*string, bool)`

GetSourceHostUserOk returns a tuple with the SourceHostUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHostUser

`func (o *ASEDSourceLinkSourceParameters) SetSourceHostUser(v string)`

SetSourceHostUser sets SourceHostUser field to given value.

### HasSourceHostUser

`func (o *ASEDSourceLinkSourceParameters) HasSourceHostUser() bool`

HasSourceHostUser returns a boolean if a field has been set.

### GetDbUser

`func (o *ASEDSourceLinkSourceParameters) GetDbUser() string`

GetDbUser returns the DbUser field if non-nil, zero value otherwise.

### GetDbUserOk

`func (o *ASEDSourceLinkSourceParameters) GetDbUserOk() (*string, bool)`

GetDbUserOk returns a tuple with the DbUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUser

`func (o *ASEDSourceLinkSourceParameters) SetDbUser(v string)`

SetDbUser sets DbUser field to given value.

### HasDbUser

`func (o *ASEDSourceLinkSourceParameters) HasDbUser() bool`

HasDbUser returns a boolean if a field has been set.

### GetDbPassword

`func (o *ASEDSourceLinkSourceParameters) GetDbPassword() string`

GetDbPassword returns the DbPassword field if non-nil, zero value otherwise.

### GetDbPasswordOk

`func (o *ASEDSourceLinkSourceParameters) GetDbPasswordOk() (*string, bool)`

GetDbPasswordOk returns a tuple with the DbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPassword

`func (o *ASEDSourceLinkSourceParameters) SetDbPassword(v string)`

SetDbPassword sets DbPassword field to given value.

### HasDbPassword

`func (o *ASEDSourceLinkSourceParameters) HasDbPassword() bool`

HasDbPassword returns a boolean if a field has been set.

### GetDbVault

`func (o *ASEDSourceLinkSourceParameters) GetDbVault() string`

GetDbVault returns the DbVault field if non-nil, zero value otherwise.

### GetDbVaultOk

`func (o *ASEDSourceLinkSourceParameters) GetDbVaultOk() (*string, bool)`

GetDbVaultOk returns a tuple with the DbVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbVault

`func (o *ASEDSourceLinkSourceParameters) SetDbVault(v string)`

SetDbVault sets DbVault field to given value.

### HasDbVault

`func (o *ASEDSourceLinkSourceParameters) HasDbVault() bool`

HasDbVault returns a boolean if a field has been set.

### GetDbHashicorpVaultEngine

`func (o *ASEDSourceLinkSourceParameters) GetDbHashicorpVaultEngine() string`

GetDbHashicorpVaultEngine returns the DbHashicorpVaultEngine field if non-nil, zero value otherwise.

### GetDbHashicorpVaultEngineOk

`func (o *ASEDSourceLinkSourceParameters) GetDbHashicorpVaultEngineOk() (*string, bool)`

GetDbHashicorpVaultEngineOk returns a tuple with the DbHashicorpVaultEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHashicorpVaultEngine

`func (o *ASEDSourceLinkSourceParameters) SetDbHashicorpVaultEngine(v string)`

SetDbHashicorpVaultEngine sets DbHashicorpVaultEngine field to given value.

### HasDbHashicorpVaultEngine

`func (o *ASEDSourceLinkSourceParameters) HasDbHashicorpVaultEngine() bool`

HasDbHashicorpVaultEngine returns a boolean if a field has been set.

### GetDbHashicorpVaultSecretPath

`func (o *ASEDSourceLinkSourceParameters) GetDbHashicorpVaultSecretPath() string`

GetDbHashicorpVaultSecretPath returns the DbHashicorpVaultSecretPath field if non-nil, zero value otherwise.

### GetDbHashicorpVaultSecretPathOk

`func (o *ASEDSourceLinkSourceParameters) GetDbHashicorpVaultSecretPathOk() (*string, bool)`

GetDbHashicorpVaultSecretPathOk returns a tuple with the DbHashicorpVaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHashicorpVaultSecretPath

`func (o *ASEDSourceLinkSourceParameters) SetDbHashicorpVaultSecretPath(v string)`

SetDbHashicorpVaultSecretPath sets DbHashicorpVaultSecretPath field to given value.

### HasDbHashicorpVaultSecretPath

`func (o *ASEDSourceLinkSourceParameters) HasDbHashicorpVaultSecretPath() bool`

HasDbHashicorpVaultSecretPath returns a boolean if a field has been set.

### GetDbHashicorpVaultUsernameKey

`func (o *ASEDSourceLinkSourceParameters) GetDbHashicorpVaultUsernameKey() string`

GetDbHashicorpVaultUsernameKey returns the DbHashicorpVaultUsernameKey field if non-nil, zero value otherwise.

### GetDbHashicorpVaultUsernameKeyOk

`func (o *ASEDSourceLinkSourceParameters) GetDbHashicorpVaultUsernameKeyOk() (*string, bool)`

GetDbHashicorpVaultUsernameKeyOk returns a tuple with the DbHashicorpVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHashicorpVaultUsernameKey

`func (o *ASEDSourceLinkSourceParameters) SetDbHashicorpVaultUsernameKey(v string)`

SetDbHashicorpVaultUsernameKey sets DbHashicorpVaultUsernameKey field to given value.

### HasDbHashicorpVaultUsernameKey

`func (o *ASEDSourceLinkSourceParameters) HasDbHashicorpVaultUsernameKey() bool`

HasDbHashicorpVaultUsernameKey returns a boolean if a field has been set.

### GetDbHashicorpVaultSecretKey

`func (o *ASEDSourceLinkSourceParameters) GetDbHashicorpVaultSecretKey() string`

GetDbHashicorpVaultSecretKey returns the DbHashicorpVaultSecretKey field if non-nil, zero value otherwise.

### GetDbHashicorpVaultSecretKeyOk

`func (o *ASEDSourceLinkSourceParameters) GetDbHashicorpVaultSecretKeyOk() (*string, bool)`

GetDbHashicorpVaultSecretKeyOk returns a tuple with the DbHashicorpVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHashicorpVaultSecretKey

`func (o *ASEDSourceLinkSourceParameters) SetDbHashicorpVaultSecretKey(v string)`

SetDbHashicorpVaultSecretKey sets DbHashicorpVaultSecretKey field to given value.

### HasDbHashicorpVaultSecretKey

`func (o *ASEDSourceLinkSourceParameters) HasDbHashicorpVaultSecretKey() bool`

HasDbHashicorpVaultSecretKey returns a boolean if a field has been set.

### GetDbAzureVaultName

`func (o *ASEDSourceLinkSourceParameters) GetDbAzureVaultName() string`

GetDbAzureVaultName returns the DbAzureVaultName field if non-nil, zero value otherwise.

### GetDbAzureVaultNameOk

`func (o *ASEDSourceLinkSourceParameters) GetDbAzureVaultNameOk() (*string, bool)`

GetDbAzureVaultNameOk returns a tuple with the DbAzureVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbAzureVaultName

`func (o *ASEDSourceLinkSourceParameters) SetDbAzureVaultName(v string)`

SetDbAzureVaultName sets DbAzureVaultName field to given value.

### HasDbAzureVaultName

`func (o *ASEDSourceLinkSourceParameters) HasDbAzureVaultName() bool`

HasDbAzureVaultName returns a boolean if a field has been set.

### GetDbAzureVaultUsernameKey

`func (o *ASEDSourceLinkSourceParameters) GetDbAzureVaultUsernameKey() string`

GetDbAzureVaultUsernameKey returns the DbAzureVaultUsernameKey field if non-nil, zero value otherwise.

### GetDbAzureVaultUsernameKeyOk

`func (o *ASEDSourceLinkSourceParameters) GetDbAzureVaultUsernameKeyOk() (*string, bool)`

GetDbAzureVaultUsernameKeyOk returns a tuple with the DbAzureVaultUsernameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbAzureVaultUsernameKey

`func (o *ASEDSourceLinkSourceParameters) SetDbAzureVaultUsernameKey(v string)`

SetDbAzureVaultUsernameKey sets DbAzureVaultUsernameKey field to given value.

### HasDbAzureVaultUsernameKey

`func (o *ASEDSourceLinkSourceParameters) HasDbAzureVaultUsernameKey() bool`

HasDbAzureVaultUsernameKey returns a boolean if a field has been set.

### GetDbAzureVaultSecretKey

`func (o *ASEDSourceLinkSourceParameters) GetDbAzureVaultSecretKey() string`

GetDbAzureVaultSecretKey returns the DbAzureVaultSecretKey field if non-nil, zero value otherwise.

### GetDbAzureVaultSecretKeyOk

`func (o *ASEDSourceLinkSourceParameters) GetDbAzureVaultSecretKeyOk() (*string, bool)`

GetDbAzureVaultSecretKeyOk returns a tuple with the DbAzureVaultSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbAzureVaultSecretKey

`func (o *ASEDSourceLinkSourceParameters) SetDbAzureVaultSecretKey(v string)`

SetDbAzureVaultSecretKey sets DbAzureVaultSecretKey field to given value.

### HasDbAzureVaultSecretKey

`func (o *ASEDSourceLinkSourceParameters) HasDbAzureVaultSecretKey() bool`

HasDbAzureVaultSecretKey returns a boolean if a field has been set.

### GetDbCyberarkVaultQueryString

`func (o *ASEDSourceLinkSourceParameters) GetDbCyberarkVaultQueryString() string`

GetDbCyberarkVaultQueryString returns the DbCyberarkVaultQueryString field if non-nil, zero value otherwise.

### GetDbCyberarkVaultQueryStringOk

`func (o *ASEDSourceLinkSourceParameters) GetDbCyberarkVaultQueryStringOk() (*string, bool)`

GetDbCyberarkVaultQueryStringOk returns a tuple with the DbCyberarkVaultQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCyberarkVaultQueryString

`func (o *ASEDSourceLinkSourceParameters) SetDbCyberarkVaultQueryString(v string)`

SetDbCyberarkVaultQueryString sets DbCyberarkVaultQueryString field to given value.

### HasDbCyberarkVaultQueryString

`func (o *ASEDSourceLinkSourceParameters) HasDbCyberarkVaultQueryString() bool`

HasDbCyberarkVaultQueryString returns a boolean if a field has been set.

### GetStagingRepository

`func (o *ASEDSourceLinkSourceParameters) GetStagingRepository() string`

GetStagingRepository returns the StagingRepository field if non-nil, zero value otherwise.

### GetStagingRepositoryOk

`func (o *ASEDSourceLinkSourceParameters) GetStagingRepositoryOk() (*string, bool)`

GetStagingRepositoryOk returns a tuple with the StagingRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingRepository

`func (o *ASEDSourceLinkSourceParameters) SetStagingRepository(v string)`

SetStagingRepository sets StagingRepository field to given value.

### HasStagingRepository

`func (o *ASEDSourceLinkSourceParameters) HasStagingRepository() bool`

HasStagingRepository returns a boolean if a field has been set.

### GetStagingHostUser

`func (o *ASEDSourceLinkSourceParameters) GetStagingHostUser() string`

GetStagingHostUser returns the StagingHostUser field if non-nil, zero value otherwise.

### GetStagingHostUserOk

`func (o *ASEDSourceLinkSourceParameters) GetStagingHostUserOk() (*string, bool)`

GetStagingHostUserOk returns a tuple with the StagingHostUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingHostUser

`func (o *ASEDSourceLinkSourceParameters) SetStagingHostUser(v string)`

SetStagingHostUser sets StagingHostUser field to given value.

### HasStagingHostUser

`func (o *ASEDSourceLinkSourceParameters) HasStagingHostUser() bool`

HasStagingHostUser returns a boolean if a field has been set.

### GetValidatedSyncMode

`func (o *ASEDSourceLinkSourceParameters) GetValidatedSyncMode() string`

GetValidatedSyncMode returns the ValidatedSyncMode field if non-nil, zero value otherwise.

### GetValidatedSyncModeOk

`func (o *ASEDSourceLinkSourceParameters) GetValidatedSyncModeOk() (*string, bool)`

GetValidatedSyncModeOk returns a tuple with the ValidatedSyncMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedSyncMode

`func (o *ASEDSourceLinkSourceParameters) SetValidatedSyncMode(v string)`

SetValidatedSyncMode sets ValidatedSyncMode field to given value.

### HasValidatedSyncMode

`func (o *ASEDSourceLinkSourceParameters) HasValidatedSyncMode() bool`

HasValidatedSyncMode returns a boolean if a field has been set.

### GetDumpHistoryFileEnabled

`func (o *ASEDSourceLinkSourceParameters) GetDumpHistoryFileEnabled() bool`

GetDumpHistoryFileEnabled returns the DumpHistoryFileEnabled field if non-nil, zero value otherwise.

### GetDumpHistoryFileEnabledOk

`func (o *ASEDSourceLinkSourceParameters) GetDumpHistoryFileEnabledOk() (*bool, bool)`

GetDumpHistoryFileEnabledOk returns a tuple with the DumpHistoryFileEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDumpHistoryFileEnabled

`func (o *ASEDSourceLinkSourceParameters) SetDumpHistoryFileEnabled(v bool)`

SetDumpHistoryFileEnabled sets DumpHistoryFileEnabled field to given value.

### HasDumpHistoryFileEnabled

`func (o *ASEDSourceLinkSourceParameters) HasDumpHistoryFileEnabled() bool`

HasDumpHistoryFileEnabled returns a boolean if a field has been set.

### GetDropAndRecreateDevices

`func (o *ASEDSourceLinkSourceParameters) GetDropAndRecreateDevices() bool`

GetDropAndRecreateDevices returns the DropAndRecreateDevices field if non-nil, zero value otherwise.

### GetDropAndRecreateDevicesOk

`func (o *ASEDSourceLinkSourceParameters) GetDropAndRecreateDevicesOk() (*bool, bool)`

GetDropAndRecreateDevicesOk returns a tuple with the DropAndRecreateDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropAndRecreateDevices

`func (o *ASEDSourceLinkSourceParameters) SetDropAndRecreateDevices(v bool)`

SetDropAndRecreateDevices sets DropAndRecreateDevices field to given value.

### HasDropAndRecreateDevices

`func (o *ASEDSourceLinkSourceParameters) HasDropAndRecreateDevices() bool`

HasDropAndRecreateDevices returns a boolean if a field has been set.

### GetSyncStrategy

`func (o *ASEDSourceLinkSourceParameters) GetSyncStrategy() string`

GetSyncStrategy returns the SyncStrategy field if non-nil, zero value otherwise.

### GetSyncStrategyOk

`func (o *ASEDSourceLinkSourceParameters) GetSyncStrategyOk() (*string, bool)`

GetSyncStrategyOk returns a tuple with the SyncStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStrategy

`func (o *ASEDSourceLinkSourceParameters) SetSyncStrategy(v string)`

SetSyncStrategy sets SyncStrategy field to given value.

### HasSyncStrategy

`func (o *ASEDSourceLinkSourceParameters) HasSyncStrategy() bool`

HasSyncStrategy returns a boolean if a field has been set.

### GetAseBackupFiles

`func (o *ASEDSourceLinkSourceParameters) GetAseBackupFiles() []string`

GetAseBackupFiles returns the AseBackupFiles field if non-nil, zero value otherwise.

### GetAseBackupFilesOk

`func (o *ASEDSourceLinkSourceParameters) GetAseBackupFilesOk() (*[]string, bool)`

GetAseBackupFilesOk returns a tuple with the AseBackupFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseBackupFiles

`func (o *ASEDSourceLinkSourceParameters) SetAseBackupFiles(v []string)`

SetAseBackupFiles sets AseBackupFiles field to given value.

### HasAseBackupFiles

`func (o *ASEDSourceLinkSourceParameters) HasAseBackupFiles() bool`

HasAseBackupFiles returns a boolean if a field has been set.

### GetPreValidatedSync

`func (o *ASEDSourceLinkSourceParameters) GetPreValidatedSync() []SourceOperation`

GetPreValidatedSync returns the PreValidatedSync field if non-nil, zero value otherwise.

### GetPreValidatedSyncOk

`func (o *ASEDSourceLinkSourceParameters) GetPreValidatedSyncOk() (*[]SourceOperation, bool)`

GetPreValidatedSyncOk returns a tuple with the PreValidatedSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreValidatedSync

`func (o *ASEDSourceLinkSourceParameters) SetPreValidatedSync(v []SourceOperation)`

SetPreValidatedSync sets PreValidatedSync field to given value.

### HasPreValidatedSync

`func (o *ASEDSourceLinkSourceParameters) HasPreValidatedSync() bool`

HasPreValidatedSync returns a boolean if a field has been set.

### GetPostValidatedSync

`func (o *ASEDSourceLinkSourceParameters) GetPostValidatedSync() []SourceOperation`

GetPostValidatedSync returns the PostValidatedSync field if non-nil, zero value otherwise.

### GetPostValidatedSyncOk

`func (o *ASEDSourceLinkSourceParameters) GetPostValidatedSyncOk() (*[]SourceOperation, bool)`

GetPostValidatedSyncOk returns a tuple with the PostValidatedSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostValidatedSync

`func (o *ASEDSourceLinkSourceParameters) SetPostValidatedSync(v []SourceOperation)`

SetPostValidatedSync sets PostValidatedSync field to given value.

### HasPostValidatedSync

`func (o *ASEDSourceLinkSourceParameters) HasPostValidatedSync() bool`

HasPostValidatedSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


