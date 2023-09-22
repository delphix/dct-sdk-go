# ProvisionVDBFromBookmarkParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreRefresh** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment before refreshing the VDB. | [optional] 
**PostRefresh** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment after refreshing the VDB. | [optional] 
**PreRollback** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment before rewinding the VDB. | [optional] 
**PostRollback** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment after rewinding the VDB. | [optional] 
**ConfigureClone** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment when the VDB is created or refreshed. | [optional] 
**PreSnapshot** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment before snapshotting a virtual source. These commands can quiesce any data prior to snapshotting. | [optional] 
**PostSnapshot** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment after snapshotting a virtual source. | [optional] 
**PreStart** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment before starting a virtual source. | [optional] 
**PostStart** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment after starting a virtual source. | [optional] 
**PreStop** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment before stopping a virtual source. | [optional] 
**PostStop** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment after stopping a virtual source. | [optional] 
**TargetGroupId** | Pointer to **string** | The ID of the group into which the VDB will be provisioned. If unset, a group is selected randomly on the Engine. | [optional] 
**Name** | Pointer to **string** | The unique name of the provisioned VDB within a group. If unset, a name is randomly generated. | [optional] 
**DatabaseName** | Pointer to **string** | The name of the database on the target environment. Defaults to the value of the name property. | [optional] 
**CdbId** | Pointer to **string** | The ID of the container database (CDB) to provision an Oracle Multitenant database into. This corresponds to a CDB or VCDB API object. When this is not set, a new vCDB will be provisioned. | [optional] 
**ClusterNodeIds** | Pointer to **[]string** | The cluster node ids, name or addresses for this provision operation (Oracle RAC Only). | [optional] 
**TruncateLogOnCheckpoint** | Pointer to **bool** | Whether to truncate log on checkpoint (ASE only). | [optional] 
**OsUsername** | Pointer to **string** | The name of the privileged user to run the provision operation (Oracle Only). | [optional] 
**OsPassword** | Pointer to **string** | The password of the privileged user to run the provision operation (Oracle Only). | [optional] 
**EnvironmentId** | Pointer to **string** | The ID of the target environment where to provision the VDB. If repository_id unambigously identifies a repository, this is unnecessary and ignored. Otherwise, a compatible repository is randomly selected on the environment. | [optional] 
**EnvironmentUserId** | Pointer to **string** | The environment user ID to use to connect to the target environment. | [optional] 
**RepositoryId** | Pointer to **string** | The ID of the target repository where to provision the VDB. A repository typically corresponds to a database installation (Oracle home, database instance, ...). Setting this attribute implicitly determines the environment where to provision the VDB. | [optional] 
**AutoSelectRepository** | Pointer to **bool** | Option to automatically select a compatible environment and repository. Mutually exclusive with repository_id. | [optional] 
**VdbRestart** | Pointer to **bool** | Indicates whether the Engine should automatically restart this virtual source when target host reboot is detected. | [optional] 
**TemplateId** | Pointer to **string** | The ID of the target VDB Template (Oracle Only). | [optional] 
**AuxiliaryTemplateId** | Pointer to **string** | The ID of the configuration template to apply to the auxiliary container database. This is only relevant when provisioning a Multitenant pluggable database into an existing CDB, i.e when the cdb_id property is set.(Oracle Only) | [optional] 
**FileMappingRules** | Pointer to **string** | Target VDB file mapping rules (Oracle Only). Rules must be line separated (\\n or \\r) and each line must have the format \&quot;pattern:replacement\&quot;. Lines are applied in order. | [optional] 
**OracleInstanceName** | Pointer to **string** | Target VDB SID name (Oracle Only). | [optional] 
**UniqueName** | Pointer to **string** | Target VDB db_unique_name (Oracle Only). | [optional] 
**VcdbName** | Pointer to **string** | When provisioning an Oracle Multitenant vCDB (when the cdb_id property is not set), the name of the provisioned vCDB (Oracle Multitenant Only). | [optional] 
**VcdbDatabaseName** | Pointer to **string** | When provisioning an Oracle Multitenant vCDB (when the cdb_id property is not set), the database name of the provisioned vCDB. Defaults to the value of the vcdb_name property. (Oracle Multitenant Only). | [optional] 
**MountPoint** | Pointer to **string** | Mount point for the VDB (Oracle, ASE, AppData). | [optional] 
**OpenResetLogs** | Pointer to **bool** | Whether to open the database after provision (Oracle Only). | [optional] 
**SnapshotPolicyId** | Pointer to **string** | The ID of the snapshot policy for the VDB. | [optional] 
**RetentionPolicyId** | Pointer to **string** | The ID of the retention policy for the VDB. | [optional] 
**RecoveryModel** | Pointer to **string** | Recovery model of the source database (MSSql Only). | [optional] 
**PreScript** | Pointer to **string** | PowerShell script or executable to run prior to provisioning (MSSql Only). | [optional] 
**PostScript** | Pointer to **string** | PowerShell script or executable to run after provisioning (MSSql Only). | [optional] 
**CdcOnProvision** | Pointer to **bool** | Option to enable change data capture (CDC) on both the provisioned VDB and subsequent snapshot-related operations (e.g. refresh, rewind) (MSSql Only). | [optional] 
**OnlineLogSize** | Pointer to **int32** | Online log size in MB (Oracle Only). | [optional] 
**OnlineLogGroups** | Pointer to **int32** | Number of online log groups (Oracle Only). | [optional] 
**ArchiveLog** | Pointer to **bool** | Option to create a VDB in archivelog mode (Oracle Only). | [optional] 
**NewDbid** | Pointer to **bool** | Option to generate a new DB ID for the created VDB (Oracle Only). | [optional] 
**ListenerIds** | Pointer to **[]string** | The listener IDs for this provision operation (Oracle Only). | [optional] 
**CustomEnvVars** | Pointer to **map[string]string** | Environment variable to be set when the engine creates a VDB. See the Engine documentation for the list of allowed/denied environment variables and rules about substitution. | [optional] 
**CustomEnvFiles** | Pointer to **[]string** | Environment files to be sourced when the Engine creates a VDB. This path can be followed by parameters. Paths and parameters are separated by spaces. | [optional] 
**OracleRacCustomEnvFiles** | Pointer to [**[]OracleRacCustomEnvFile**](OracleRacCustomEnvFile.md) | Environment files to be sourced when the Engine creates an Oracle RAC VDB. This path can be followed by parameters. Paths and parameters are separated by spaces. | [optional] 
**OracleRacCustomEnvVars** | Pointer to [**[]OracleRacCustomEnvVar**](OracleRacCustomEnvVar.md) | Environment variable to be set when the engine creates an Oracle RAC VDB. See the Engine documentation for the list of allowed/denied environment variables and rules about substitution. | [optional] 
**ParentTdeKeystorePath** | Pointer to **string** | Path to a copy of the parent&#39;s Oracle transparent data encryption keystore on the target host. Required to provision from snapshots containing encrypted database files. (Oracle Multitenant Only) | [optional] 
**ParentTdeKeystorePassword** | Pointer to **string** | The password of the keystore specified in parentTdeKeystorePath. (Oracle Multitenant Only) | [optional] 
**TdeExportedKeyFileSecret** | Pointer to **string** | Secret to be used while exporting and importing vPDB encryption keys if Transparent Data Encryption is enabled on the vPDB. (Oracle Multitenant Only) | [optional] 
**TdeKeyIdentifier** | Pointer to **string** | ID of the key created by Delphix. (Oracle Multitenant Only) | [optional] 
**TargetVcdbTdeKeystorePath** | Pointer to **string** | Path to the keystore of the target vCDB. (Oracle Multitenant Only) | [optional] 
**CdbTdeKeystorePassword** | Pointer to **string** | The password for the Transparent Data Encryption keystore associated with the CDB. (Oracle Multitenant Only) | [optional] 
**VcdbTdeKeyIdentifier** | Pointer to **string** | ID of the key created by Delphix. (Oracle Multitenant Only) | [optional] 
**AppdataSourceParams** | Pointer to **map[string]interface{}** | The JSON payload conforming to the DraftV4 schema based on the type of application data being manipulated. | [optional] 
**AdditionalMountPoints** | Pointer to [**[]AdditionalMountPoint**](AdditionalMountPoint.md) | Specifies additional locations on which to mount a subdirectory of an AppData container. | [optional] 
**AppdataConfigParams** | Pointer to **map[string]interface{}** | The list of parameters specified by the source config schema in the toolkit | [optional] 
**ConfigParams** | Pointer to **map[string]interface{}** | Database configuration parameter overrides. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | The tags to be created for VDB. | [optional] 
**BookmarkId** | **string** | The ID of the bookmark from which to execute the operation. The bookmark must contain only one VDB. | 
**MakeCurrentAccountOwner** | Pointer to **bool** | Whether the account provisioning this VDB must be configured as owner of the VDB. | [optional] [default to true]

## Methods

### NewProvisionVDBFromBookmarkParameters

`func NewProvisionVDBFromBookmarkParameters(bookmarkId string, ) *ProvisionVDBFromBookmarkParameters`

NewProvisionVDBFromBookmarkParameters instantiates a new ProvisionVDBFromBookmarkParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionVDBFromBookmarkParametersWithDefaults

`func NewProvisionVDBFromBookmarkParametersWithDefaults() *ProvisionVDBFromBookmarkParameters`

NewProvisionVDBFromBookmarkParametersWithDefaults instantiates a new ProvisionVDBFromBookmarkParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreRefresh

`func (o *ProvisionVDBFromBookmarkParameters) GetPreRefresh() []Hook`

GetPreRefresh returns the PreRefresh field if non-nil, zero value otherwise.

### GetPreRefreshOk

`func (o *ProvisionVDBFromBookmarkParameters) GetPreRefreshOk() (*[]Hook, bool)`

GetPreRefreshOk returns a tuple with the PreRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRefresh

`func (o *ProvisionVDBFromBookmarkParameters) SetPreRefresh(v []Hook)`

SetPreRefresh sets PreRefresh field to given value.

### HasPreRefresh

`func (o *ProvisionVDBFromBookmarkParameters) HasPreRefresh() bool`

HasPreRefresh returns a boolean if a field has been set.

### GetPostRefresh

`func (o *ProvisionVDBFromBookmarkParameters) GetPostRefresh() []Hook`

GetPostRefresh returns the PostRefresh field if non-nil, zero value otherwise.

### GetPostRefreshOk

`func (o *ProvisionVDBFromBookmarkParameters) GetPostRefreshOk() (*[]Hook, bool)`

GetPostRefreshOk returns a tuple with the PostRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostRefresh

`func (o *ProvisionVDBFromBookmarkParameters) SetPostRefresh(v []Hook)`

SetPostRefresh sets PostRefresh field to given value.

### HasPostRefresh

`func (o *ProvisionVDBFromBookmarkParameters) HasPostRefresh() bool`

HasPostRefresh returns a boolean if a field has been set.

### GetPreRollback

`func (o *ProvisionVDBFromBookmarkParameters) GetPreRollback() []Hook`

GetPreRollback returns the PreRollback field if non-nil, zero value otherwise.

### GetPreRollbackOk

`func (o *ProvisionVDBFromBookmarkParameters) GetPreRollbackOk() (*[]Hook, bool)`

GetPreRollbackOk returns a tuple with the PreRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRollback

`func (o *ProvisionVDBFromBookmarkParameters) SetPreRollback(v []Hook)`

SetPreRollback sets PreRollback field to given value.

### HasPreRollback

`func (o *ProvisionVDBFromBookmarkParameters) HasPreRollback() bool`

HasPreRollback returns a boolean if a field has been set.

### GetPostRollback

`func (o *ProvisionVDBFromBookmarkParameters) GetPostRollback() []Hook`

GetPostRollback returns the PostRollback field if non-nil, zero value otherwise.

### GetPostRollbackOk

`func (o *ProvisionVDBFromBookmarkParameters) GetPostRollbackOk() (*[]Hook, bool)`

GetPostRollbackOk returns a tuple with the PostRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostRollback

`func (o *ProvisionVDBFromBookmarkParameters) SetPostRollback(v []Hook)`

SetPostRollback sets PostRollback field to given value.

### HasPostRollback

`func (o *ProvisionVDBFromBookmarkParameters) HasPostRollback() bool`

HasPostRollback returns a boolean if a field has been set.

### GetConfigureClone

`func (o *ProvisionVDBFromBookmarkParameters) GetConfigureClone() []Hook`

GetConfigureClone returns the ConfigureClone field if non-nil, zero value otherwise.

### GetConfigureCloneOk

`func (o *ProvisionVDBFromBookmarkParameters) GetConfigureCloneOk() (*[]Hook, bool)`

GetConfigureCloneOk returns a tuple with the ConfigureClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureClone

`func (o *ProvisionVDBFromBookmarkParameters) SetConfigureClone(v []Hook)`

SetConfigureClone sets ConfigureClone field to given value.

### HasConfigureClone

`func (o *ProvisionVDBFromBookmarkParameters) HasConfigureClone() bool`

HasConfigureClone returns a boolean if a field has been set.

### GetPreSnapshot

`func (o *ProvisionVDBFromBookmarkParameters) GetPreSnapshot() []Hook`

GetPreSnapshot returns the PreSnapshot field if non-nil, zero value otherwise.

### GetPreSnapshotOk

`func (o *ProvisionVDBFromBookmarkParameters) GetPreSnapshotOk() (*[]Hook, bool)`

GetPreSnapshotOk returns a tuple with the PreSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSnapshot

`func (o *ProvisionVDBFromBookmarkParameters) SetPreSnapshot(v []Hook)`

SetPreSnapshot sets PreSnapshot field to given value.

### HasPreSnapshot

`func (o *ProvisionVDBFromBookmarkParameters) HasPreSnapshot() bool`

HasPreSnapshot returns a boolean if a field has been set.

### GetPostSnapshot

`func (o *ProvisionVDBFromBookmarkParameters) GetPostSnapshot() []Hook`

GetPostSnapshot returns the PostSnapshot field if non-nil, zero value otherwise.

### GetPostSnapshotOk

`func (o *ProvisionVDBFromBookmarkParameters) GetPostSnapshotOk() (*[]Hook, bool)`

GetPostSnapshotOk returns a tuple with the PostSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostSnapshot

`func (o *ProvisionVDBFromBookmarkParameters) SetPostSnapshot(v []Hook)`

SetPostSnapshot sets PostSnapshot field to given value.

### HasPostSnapshot

`func (o *ProvisionVDBFromBookmarkParameters) HasPostSnapshot() bool`

HasPostSnapshot returns a boolean if a field has been set.

### GetPreStart

`func (o *ProvisionVDBFromBookmarkParameters) GetPreStart() []Hook`

GetPreStart returns the PreStart field if non-nil, zero value otherwise.

### GetPreStartOk

`func (o *ProvisionVDBFromBookmarkParameters) GetPreStartOk() (*[]Hook, bool)`

GetPreStartOk returns a tuple with the PreStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreStart

`func (o *ProvisionVDBFromBookmarkParameters) SetPreStart(v []Hook)`

SetPreStart sets PreStart field to given value.

### HasPreStart

`func (o *ProvisionVDBFromBookmarkParameters) HasPreStart() bool`

HasPreStart returns a boolean if a field has been set.

### GetPostStart

`func (o *ProvisionVDBFromBookmarkParameters) GetPostStart() []Hook`

GetPostStart returns the PostStart field if non-nil, zero value otherwise.

### GetPostStartOk

`func (o *ProvisionVDBFromBookmarkParameters) GetPostStartOk() (*[]Hook, bool)`

GetPostStartOk returns a tuple with the PostStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostStart

`func (o *ProvisionVDBFromBookmarkParameters) SetPostStart(v []Hook)`

SetPostStart sets PostStart field to given value.

### HasPostStart

`func (o *ProvisionVDBFromBookmarkParameters) HasPostStart() bool`

HasPostStart returns a boolean if a field has been set.

### GetPreStop

`func (o *ProvisionVDBFromBookmarkParameters) GetPreStop() []Hook`

GetPreStop returns the PreStop field if non-nil, zero value otherwise.

### GetPreStopOk

`func (o *ProvisionVDBFromBookmarkParameters) GetPreStopOk() (*[]Hook, bool)`

GetPreStopOk returns a tuple with the PreStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreStop

`func (o *ProvisionVDBFromBookmarkParameters) SetPreStop(v []Hook)`

SetPreStop sets PreStop field to given value.

### HasPreStop

`func (o *ProvisionVDBFromBookmarkParameters) HasPreStop() bool`

HasPreStop returns a boolean if a field has been set.

### GetPostStop

`func (o *ProvisionVDBFromBookmarkParameters) GetPostStop() []Hook`

GetPostStop returns the PostStop field if non-nil, zero value otherwise.

### GetPostStopOk

`func (o *ProvisionVDBFromBookmarkParameters) GetPostStopOk() (*[]Hook, bool)`

GetPostStopOk returns a tuple with the PostStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostStop

`func (o *ProvisionVDBFromBookmarkParameters) SetPostStop(v []Hook)`

SetPostStop sets PostStop field to given value.

### HasPostStop

`func (o *ProvisionVDBFromBookmarkParameters) HasPostStop() bool`

HasPostStop returns a boolean if a field has been set.

### GetTargetGroupId

`func (o *ProvisionVDBFromBookmarkParameters) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *ProvisionVDBFromBookmarkParameters) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *ProvisionVDBFromBookmarkParameters) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *ProvisionVDBFromBookmarkParameters) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### GetName

`func (o *ProvisionVDBFromBookmarkParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisionVDBFromBookmarkParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisionVDBFromBookmarkParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProvisionVDBFromBookmarkParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDatabaseName

`func (o *ProvisionVDBFromBookmarkParameters) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ProvisionVDBFromBookmarkParameters) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ProvisionVDBFromBookmarkParameters) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *ProvisionVDBFromBookmarkParameters) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetCdbId

`func (o *ProvisionVDBFromBookmarkParameters) GetCdbId() string`

GetCdbId returns the CdbId field if non-nil, zero value otherwise.

### GetCdbIdOk

`func (o *ProvisionVDBFromBookmarkParameters) GetCdbIdOk() (*string, bool)`

GetCdbIdOk returns a tuple with the CdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdbId

`func (o *ProvisionVDBFromBookmarkParameters) SetCdbId(v string)`

SetCdbId sets CdbId field to given value.

### HasCdbId

`func (o *ProvisionVDBFromBookmarkParameters) HasCdbId() bool`

HasCdbId returns a boolean if a field has been set.

### GetClusterNodeIds

`func (o *ProvisionVDBFromBookmarkParameters) GetClusterNodeIds() []string`

GetClusterNodeIds returns the ClusterNodeIds field if non-nil, zero value otherwise.

### GetClusterNodeIdsOk

`func (o *ProvisionVDBFromBookmarkParameters) GetClusterNodeIdsOk() (*[]string, bool)`

GetClusterNodeIdsOk returns a tuple with the ClusterNodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeIds

`func (o *ProvisionVDBFromBookmarkParameters) SetClusterNodeIds(v []string)`

SetClusterNodeIds sets ClusterNodeIds field to given value.

### HasClusterNodeIds

`func (o *ProvisionVDBFromBookmarkParameters) HasClusterNodeIds() bool`

HasClusterNodeIds returns a boolean if a field has been set.

### GetTruncateLogOnCheckpoint

`func (o *ProvisionVDBFromBookmarkParameters) GetTruncateLogOnCheckpoint() bool`

GetTruncateLogOnCheckpoint returns the TruncateLogOnCheckpoint field if non-nil, zero value otherwise.

### GetTruncateLogOnCheckpointOk

`func (o *ProvisionVDBFromBookmarkParameters) GetTruncateLogOnCheckpointOk() (*bool, bool)`

GetTruncateLogOnCheckpointOk returns a tuple with the TruncateLogOnCheckpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateLogOnCheckpoint

`func (o *ProvisionVDBFromBookmarkParameters) SetTruncateLogOnCheckpoint(v bool)`

SetTruncateLogOnCheckpoint sets TruncateLogOnCheckpoint field to given value.

### HasTruncateLogOnCheckpoint

`func (o *ProvisionVDBFromBookmarkParameters) HasTruncateLogOnCheckpoint() bool`

HasTruncateLogOnCheckpoint returns a boolean if a field has been set.

### GetOsUsername

`func (o *ProvisionVDBFromBookmarkParameters) GetOsUsername() string`

GetOsUsername returns the OsUsername field if non-nil, zero value otherwise.

### GetOsUsernameOk

`func (o *ProvisionVDBFromBookmarkParameters) GetOsUsernameOk() (*string, bool)`

GetOsUsernameOk returns a tuple with the OsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsUsername

`func (o *ProvisionVDBFromBookmarkParameters) SetOsUsername(v string)`

SetOsUsername sets OsUsername field to given value.

### HasOsUsername

`func (o *ProvisionVDBFromBookmarkParameters) HasOsUsername() bool`

HasOsUsername returns a boolean if a field has been set.

### GetOsPassword

`func (o *ProvisionVDBFromBookmarkParameters) GetOsPassword() string`

GetOsPassword returns the OsPassword field if non-nil, zero value otherwise.

### GetOsPasswordOk

`func (o *ProvisionVDBFromBookmarkParameters) GetOsPasswordOk() (*string, bool)`

GetOsPasswordOk returns a tuple with the OsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPassword

`func (o *ProvisionVDBFromBookmarkParameters) SetOsPassword(v string)`

SetOsPassword sets OsPassword field to given value.

### HasOsPassword

`func (o *ProvisionVDBFromBookmarkParameters) HasOsPassword() bool`

HasOsPassword returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *ProvisionVDBFromBookmarkParameters) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ProvisionVDBFromBookmarkParameters) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ProvisionVDBFromBookmarkParameters) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ProvisionVDBFromBookmarkParameters) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEnvironmentUserId

`func (o *ProvisionVDBFromBookmarkParameters) GetEnvironmentUserId() string`

GetEnvironmentUserId returns the EnvironmentUserId field if non-nil, zero value otherwise.

### GetEnvironmentUserIdOk

`func (o *ProvisionVDBFromBookmarkParameters) GetEnvironmentUserIdOk() (*string, bool)`

GetEnvironmentUserIdOk returns a tuple with the EnvironmentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUserId

`func (o *ProvisionVDBFromBookmarkParameters) SetEnvironmentUserId(v string)`

SetEnvironmentUserId sets EnvironmentUserId field to given value.

### HasEnvironmentUserId

`func (o *ProvisionVDBFromBookmarkParameters) HasEnvironmentUserId() bool`

HasEnvironmentUserId returns a boolean if a field has been set.

### GetRepositoryId

`func (o *ProvisionVDBFromBookmarkParameters) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *ProvisionVDBFromBookmarkParameters) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *ProvisionVDBFromBookmarkParameters) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *ProvisionVDBFromBookmarkParameters) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetAutoSelectRepository

`func (o *ProvisionVDBFromBookmarkParameters) GetAutoSelectRepository() bool`

GetAutoSelectRepository returns the AutoSelectRepository field if non-nil, zero value otherwise.

### GetAutoSelectRepositoryOk

`func (o *ProvisionVDBFromBookmarkParameters) GetAutoSelectRepositoryOk() (*bool, bool)`

GetAutoSelectRepositoryOk returns a tuple with the AutoSelectRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelectRepository

`func (o *ProvisionVDBFromBookmarkParameters) SetAutoSelectRepository(v bool)`

SetAutoSelectRepository sets AutoSelectRepository field to given value.

### HasAutoSelectRepository

`func (o *ProvisionVDBFromBookmarkParameters) HasAutoSelectRepository() bool`

HasAutoSelectRepository returns a boolean if a field has been set.

### GetVdbRestart

`func (o *ProvisionVDBFromBookmarkParameters) GetVdbRestart() bool`

GetVdbRestart returns the VdbRestart field if non-nil, zero value otherwise.

### GetVdbRestartOk

`func (o *ProvisionVDBFromBookmarkParameters) GetVdbRestartOk() (*bool, bool)`

GetVdbRestartOk returns a tuple with the VdbRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdbRestart

`func (o *ProvisionVDBFromBookmarkParameters) SetVdbRestart(v bool)`

SetVdbRestart sets VdbRestart field to given value.

### HasVdbRestart

`func (o *ProvisionVDBFromBookmarkParameters) HasVdbRestart() bool`

HasVdbRestart returns a boolean if a field has been set.

### GetTemplateId

`func (o *ProvisionVDBFromBookmarkParameters) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ProvisionVDBFromBookmarkParameters) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ProvisionVDBFromBookmarkParameters) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ProvisionVDBFromBookmarkParameters) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetAuxiliaryTemplateId

`func (o *ProvisionVDBFromBookmarkParameters) GetAuxiliaryTemplateId() string`

GetAuxiliaryTemplateId returns the AuxiliaryTemplateId field if non-nil, zero value otherwise.

### GetAuxiliaryTemplateIdOk

`func (o *ProvisionVDBFromBookmarkParameters) GetAuxiliaryTemplateIdOk() (*string, bool)`

GetAuxiliaryTemplateIdOk returns a tuple with the AuxiliaryTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryTemplateId

`func (o *ProvisionVDBFromBookmarkParameters) SetAuxiliaryTemplateId(v string)`

SetAuxiliaryTemplateId sets AuxiliaryTemplateId field to given value.

### HasAuxiliaryTemplateId

`func (o *ProvisionVDBFromBookmarkParameters) HasAuxiliaryTemplateId() bool`

HasAuxiliaryTemplateId returns a boolean if a field has been set.

### GetFileMappingRules

`func (o *ProvisionVDBFromBookmarkParameters) GetFileMappingRules() string`

GetFileMappingRules returns the FileMappingRules field if non-nil, zero value otherwise.

### GetFileMappingRulesOk

`func (o *ProvisionVDBFromBookmarkParameters) GetFileMappingRulesOk() (*string, bool)`

GetFileMappingRulesOk returns a tuple with the FileMappingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMappingRules

`func (o *ProvisionVDBFromBookmarkParameters) SetFileMappingRules(v string)`

SetFileMappingRules sets FileMappingRules field to given value.

### HasFileMappingRules

`func (o *ProvisionVDBFromBookmarkParameters) HasFileMappingRules() bool`

HasFileMappingRules returns a boolean if a field has been set.

### GetOracleInstanceName

`func (o *ProvisionVDBFromBookmarkParameters) GetOracleInstanceName() string`

GetOracleInstanceName returns the OracleInstanceName field if non-nil, zero value otherwise.

### GetOracleInstanceNameOk

`func (o *ProvisionVDBFromBookmarkParameters) GetOracleInstanceNameOk() (*string, bool)`

GetOracleInstanceNameOk returns a tuple with the OracleInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleInstanceName

`func (o *ProvisionVDBFromBookmarkParameters) SetOracleInstanceName(v string)`

SetOracleInstanceName sets OracleInstanceName field to given value.

### HasOracleInstanceName

`func (o *ProvisionVDBFromBookmarkParameters) HasOracleInstanceName() bool`

HasOracleInstanceName returns a boolean if a field has been set.

### GetUniqueName

`func (o *ProvisionVDBFromBookmarkParameters) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *ProvisionVDBFromBookmarkParameters) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *ProvisionVDBFromBookmarkParameters) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *ProvisionVDBFromBookmarkParameters) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetVcdbName

`func (o *ProvisionVDBFromBookmarkParameters) GetVcdbName() string`

GetVcdbName returns the VcdbName field if non-nil, zero value otherwise.

### GetVcdbNameOk

`func (o *ProvisionVDBFromBookmarkParameters) GetVcdbNameOk() (*string, bool)`

GetVcdbNameOk returns a tuple with the VcdbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdbName

`func (o *ProvisionVDBFromBookmarkParameters) SetVcdbName(v string)`

SetVcdbName sets VcdbName field to given value.

### HasVcdbName

`func (o *ProvisionVDBFromBookmarkParameters) HasVcdbName() bool`

HasVcdbName returns a boolean if a field has been set.

### GetVcdbDatabaseName

`func (o *ProvisionVDBFromBookmarkParameters) GetVcdbDatabaseName() string`

GetVcdbDatabaseName returns the VcdbDatabaseName field if non-nil, zero value otherwise.

### GetVcdbDatabaseNameOk

`func (o *ProvisionVDBFromBookmarkParameters) GetVcdbDatabaseNameOk() (*string, bool)`

GetVcdbDatabaseNameOk returns a tuple with the VcdbDatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdbDatabaseName

`func (o *ProvisionVDBFromBookmarkParameters) SetVcdbDatabaseName(v string)`

SetVcdbDatabaseName sets VcdbDatabaseName field to given value.

### HasVcdbDatabaseName

`func (o *ProvisionVDBFromBookmarkParameters) HasVcdbDatabaseName() bool`

HasVcdbDatabaseName returns a boolean if a field has been set.

### GetMountPoint

`func (o *ProvisionVDBFromBookmarkParameters) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *ProvisionVDBFromBookmarkParameters) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *ProvisionVDBFromBookmarkParameters) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *ProvisionVDBFromBookmarkParameters) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### GetOpenResetLogs

`func (o *ProvisionVDBFromBookmarkParameters) GetOpenResetLogs() bool`

GetOpenResetLogs returns the OpenResetLogs field if non-nil, zero value otherwise.

### GetOpenResetLogsOk

`func (o *ProvisionVDBFromBookmarkParameters) GetOpenResetLogsOk() (*bool, bool)`

GetOpenResetLogsOk returns a tuple with the OpenResetLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenResetLogs

`func (o *ProvisionVDBFromBookmarkParameters) SetOpenResetLogs(v bool)`

SetOpenResetLogs sets OpenResetLogs field to given value.

### HasOpenResetLogs

`func (o *ProvisionVDBFromBookmarkParameters) HasOpenResetLogs() bool`

HasOpenResetLogs returns a boolean if a field has been set.

### GetSnapshotPolicyId

`func (o *ProvisionVDBFromBookmarkParameters) GetSnapshotPolicyId() string`

GetSnapshotPolicyId returns the SnapshotPolicyId field if non-nil, zero value otherwise.

### GetSnapshotPolicyIdOk

`func (o *ProvisionVDBFromBookmarkParameters) GetSnapshotPolicyIdOk() (*string, bool)`

GetSnapshotPolicyIdOk returns a tuple with the SnapshotPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotPolicyId

`func (o *ProvisionVDBFromBookmarkParameters) SetSnapshotPolicyId(v string)`

SetSnapshotPolicyId sets SnapshotPolicyId field to given value.

### HasSnapshotPolicyId

`func (o *ProvisionVDBFromBookmarkParameters) HasSnapshotPolicyId() bool`

HasSnapshotPolicyId returns a boolean if a field has been set.

### GetRetentionPolicyId

`func (o *ProvisionVDBFromBookmarkParameters) GetRetentionPolicyId() string`

GetRetentionPolicyId returns the RetentionPolicyId field if non-nil, zero value otherwise.

### GetRetentionPolicyIdOk

`func (o *ProvisionVDBFromBookmarkParameters) GetRetentionPolicyIdOk() (*string, bool)`

GetRetentionPolicyIdOk returns a tuple with the RetentionPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicyId

`func (o *ProvisionVDBFromBookmarkParameters) SetRetentionPolicyId(v string)`

SetRetentionPolicyId sets RetentionPolicyId field to given value.

### HasRetentionPolicyId

`func (o *ProvisionVDBFromBookmarkParameters) HasRetentionPolicyId() bool`

HasRetentionPolicyId returns a boolean if a field has been set.

### GetRecoveryModel

`func (o *ProvisionVDBFromBookmarkParameters) GetRecoveryModel() string`

GetRecoveryModel returns the RecoveryModel field if non-nil, zero value otherwise.

### GetRecoveryModelOk

`func (o *ProvisionVDBFromBookmarkParameters) GetRecoveryModelOk() (*string, bool)`

GetRecoveryModelOk returns a tuple with the RecoveryModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryModel

`func (o *ProvisionVDBFromBookmarkParameters) SetRecoveryModel(v string)`

SetRecoveryModel sets RecoveryModel field to given value.

### HasRecoveryModel

`func (o *ProvisionVDBFromBookmarkParameters) HasRecoveryModel() bool`

HasRecoveryModel returns a boolean if a field has been set.

### GetPreScript

`func (o *ProvisionVDBFromBookmarkParameters) GetPreScript() string`

GetPreScript returns the PreScript field if non-nil, zero value otherwise.

### GetPreScriptOk

`func (o *ProvisionVDBFromBookmarkParameters) GetPreScriptOk() (*string, bool)`

GetPreScriptOk returns a tuple with the PreScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreScript

`func (o *ProvisionVDBFromBookmarkParameters) SetPreScript(v string)`

SetPreScript sets PreScript field to given value.

### HasPreScript

`func (o *ProvisionVDBFromBookmarkParameters) HasPreScript() bool`

HasPreScript returns a boolean if a field has been set.

### GetPostScript

`func (o *ProvisionVDBFromBookmarkParameters) GetPostScript() string`

GetPostScript returns the PostScript field if non-nil, zero value otherwise.

### GetPostScriptOk

`func (o *ProvisionVDBFromBookmarkParameters) GetPostScriptOk() (*string, bool)`

GetPostScriptOk returns a tuple with the PostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScript

`func (o *ProvisionVDBFromBookmarkParameters) SetPostScript(v string)`

SetPostScript sets PostScript field to given value.

### HasPostScript

`func (o *ProvisionVDBFromBookmarkParameters) HasPostScript() bool`

HasPostScript returns a boolean if a field has been set.

### GetCdcOnProvision

`func (o *ProvisionVDBFromBookmarkParameters) GetCdcOnProvision() bool`

GetCdcOnProvision returns the CdcOnProvision field if non-nil, zero value otherwise.

### GetCdcOnProvisionOk

`func (o *ProvisionVDBFromBookmarkParameters) GetCdcOnProvisionOk() (*bool, bool)`

GetCdcOnProvisionOk returns a tuple with the CdcOnProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdcOnProvision

`func (o *ProvisionVDBFromBookmarkParameters) SetCdcOnProvision(v bool)`

SetCdcOnProvision sets CdcOnProvision field to given value.

### HasCdcOnProvision

`func (o *ProvisionVDBFromBookmarkParameters) HasCdcOnProvision() bool`

HasCdcOnProvision returns a boolean if a field has been set.

### GetOnlineLogSize

`func (o *ProvisionVDBFromBookmarkParameters) GetOnlineLogSize() int32`

GetOnlineLogSize returns the OnlineLogSize field if non-nil, zero value otherwise.

### GetOnlineLogSizeOk

`func (o *ProvisionVDBFromBookmarkParameters) GetOnlineLogSizeOk() (*int32, bool)`

GetOnlineLogSizeOk returns a tuple with the OnlineLogSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineLogSize

`func (o *ProvisionVDBFromBookmarkParameters) SetOnlineLogSize(v int32)`

SetOnlineLogSize sets OnlineLogSize field to given value.

### HasOnlineLogSize

`func (o *ProvisionVDBFromBookmarkParameters) HasOnlineLogSize() bool`

HasOnlineLogSize returns a boolean if a field has been set.

### GetOnlineLogGroups

`func (o *ProvisionVDBFromBookmarkParameters) GetOnlineLogGroups() int32`

GetOnlineLogGroups returns the OnlineLogGroups field if non-nil, zero value otherwise.

### GetOnlineLogGroupsOk

`func (o *ProvisionVDBFromBookmarkParameters) GetOnlineLogGroupsOk() (*int32, bool)`

GetOnlineLogGroupsOk returns a tuple with the OnlineLogGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineLogGroups

`func (o *ProvisionVDBFromBookmarkParameters) SetOnlineLogGroups(v int32)`

SetOnlineLogGroups sets OnlineLogGroups field to given value.

### HasOnlineLogGroups

`func (o *ProvisionVDBFromBookmarkParameters) HasOnlineLogGroups() bool`

HasOnlineLogGroups returns a boolean if a field has been set.

### GetArchiveLog

`func (o *ProvisionVDBFromBookmarkParameters) GetArchiveLog() bool`

GetArchiveLog returns the ArchiveLog field if non-nil, zero value otherwise.

### GetArchiveLogOk

`func (o *ProvisionVDBFromBookmarkParameters) GetArchiveLogOk() (*bool, bool)`

GetArchiveLogOk returns a tuple with the ArchiveLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLog

`func (o *ProvisionVDBFromBookmarkParameters) SetArchiveLog(v bool)`

SetArchiveLog sets ArchiveLog field to given value.

### HasArchiveLog

`func (o *ProvisionVDBFromBookmarkParameters) HasArchiveLog() bool`

HasArchiveLog returns a boolean if a field has been set.

### GetNewDbid

`func (o *ProvisionVDBFromBookmarkParameters) GetNewDbid() bool`

GetNewDbid returns the NewDbid field if non-nil, zero value otherwise.

### GetNewDbidOk

`func (o *ProvisionVDBFromBookmarkParameters) GetNewDbidOk() (*bool, bool)`

GetNewDbidOk returns a tuple with the NewDbid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDbid

`func (o *ProvisionVDBFromBookmarkParameters) SetNewDbid(v bool)`

SetNewDbid sets NewDbid field to given value.

### HasNewDbid

`func (o *ProvisionVDBFromBookmarkParameters) HasNewDbid() bool`

HasNewDbid returns a boolean if a field has been set.

### GetListenerIds

`func (o *ProvisionVDBFromBookmarkParameters) GetListenerIds() []string`

GetListenerIds returns the ListenerIds field if non-nil, zero value otherwise.

### GetListenerIdsOk

`func (o *ProvisionVDBFromBookmarkParameters) GetListenerIdsOk() (*[]string, bool)`

GetListenerIdsOk returns a tuple with the ListenerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerIds

`func (o *ProvisionVDBFromBookmarkParameters) SetListenerIds(v []string)`

SetListenerIds sets ListenerIds field to given value.

### HasListenerIds

`func (o *ProvisionVDBFromBookmarkParameters) HasListenerIds() bool`

HasListenerIds returns a boolean if a field has been set.

### GetCustomEnvVars

`func (o *ProvisionVDBFromBookmarkParameters) GetCustomEnvVars() map[string]string`

GetCustomEnvVars returns the CustomEnvVars field if non-nil, zero value otherwise.

### GetCustomEnvVarsOk

`func (o *ProvisionVDBFromBookmarkParameters) GetCustomEnvVarsOk() (*map[string]string, bool)`

GetCustomEnvVarsOk returns a tuple with the CustomEnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvVars

`func (o *ProvisionVDBFromBookmarkParameters) SetCustomEnvVars(v map[string]string)`

SetCustomEnvVars sets CustomEnvVars field to given value.

### HasCustomEnvVars

`func (o *ProvisionVDBFromBookmarkParameters) HasCustomEnvVars() bool`

HasCustomEnvVars returns a boolean if a field has been set.

### GetCustomEnvFiles

`func (o *ProvisionVDBFromBookmarkParameters) GetCustomEnvFiles() []string`

GetCustomEnvFiles returns the CustomEnvFiles field if non-nil, zero value otherwise.

### GetCustomEnvFilesOk

`func (o *ProvisionVDBFromBookmarkParameters) GetCustomEnvFilesOk() (*[]string, bool)`

GetCustomEnvFilesOk returns a tuple with the CustomEnvFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvFiles

`func (o *ProvisionVDBFromBookmarkParameters) SetCustomEnvFiles(v []string)`

SetCustomEnvFiles sets CustomEnvFiles field to given value.

### HasCustomEnvFiles

`func (o *ProvisionVDBFromBookmarkParameters) HasCustomEnvFiles() bool`

HasCustomEnvFiles returns a boolean if a field has been set.

### GetOracleRacCustomEnvFiles

`func (o *ProvisionVDBFromBookmarkParameters) GetOracleRacCustomEnvFiles() []OracleRacCustomEnvFile`

GetOracleRacCustomEnvFiles returns the OracleRacCustomEnvFiles field if non-nil, zero value otherwise.

### GetOracleRacCustomEnvFilesOk

`func (o *ProvisionVDBFromBookmarkParameters) GetOracleRacCustomEnvFilesOk() (*[]OracleRacCustomEnvFile, bool)`

GetOracleRacCustomEnvFilesOk returns a tuple with the OracleRacCustomEnvFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleRacCustomEnvFiles

`func (o *ProvisionVDBFromBookmarkParameters) SetOracleRacCustomEnvFiles(v []OracleRacCustomEnvFile)`

SetOracleRacCustomEnvFiles sets OracleRacCustomEnvFiles field to given value.

### HasOracleRacCustomEnvFiles

`func (o *ProvisionVDBFromBookmarkParameters) HasOracleRacCustomEnvFiles() bool`

HasOracleRacCustomEnvFiles returns a boolean if a field has been set.

### GetOracleRacCustomEnvVars

`func (o *ProvisionVDBFromBookmarkParameters) GetOracleRacCustomEnvVars() []OracleRacCustomEnvVar`

GetOracleRacCustomEnvVars returns the OracleRacCustomEnvVars field if non-nil, zero value otherwise.

### GetOracleRacCustomEnvVarsOk

`func (o *ProvisionVDBFromBookmarkParameters) GetOracleRacCustomEnvVarsOk() (*[]OracleRacCustomEnvVar, bool)`

GetOracleRacCustomEnvVarsOk returns a tuple with the OracleRacCustomEnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleRacCustomEnvVars

`func (o *ProvisionVDBFromBookmarkParameters) SetOracleRacCustomEnvVars(v []OracleRacCustomEnvVar)`

SetOracleRacCustomEnvVars sets OracleRacCustomEnvVars field to given value.

### HasOracleRacCustomEnvVars

`func (o *ProvisionVDBFromBookmarkParameters) HasOracleRacCustomEnvVars() bool`

HasOracleRacCustomEnvVars returns a boolean if a field has been set.

### GetParentTdeKeystorePath

`func (o *ProvisionVDBFromBookmarkParameters) GetParentTdeKeystorePath() string`

GetParentTdeKeystorePath returns the ParentTdeKeystorePath field if non-nil, zero value otherwise.

### GetParentTdeKeystorePathOk

`func (o *ProvisionVDBFromBookmarkParameters) GetParentTdeKeystorePathOk() (*string, bool)`

GetParentTdeKeystorePathOk returns a tuple with the ParentTdeKeystorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTdeKeystorePath

`func (o *ProvisionVDBFromBookmarkParameters) SetParentTdeKeystorePath(v string)`

SetParentTdeKeystorePath sets ParentTdeKeystorePath field to given value.

### HasParentTdeKeystorePath

`func (o *ProvisionVDBFromBookmarkParameters) HasParentTdeKeystorePath() bool`

HasParentTdeKeystorePath returns a boolean if a field has been set.

### GetParentTdeKeystorePassword

`func (o *ProvisionVDBFromBookmarkParameters) GetParentTdeKeystorePassword() string`

GetParentTdeKeystorePassword returns the ParentTdeKeystorePassword field if non-nil, zero value otherwise.

### GetParentTdeKeystorePasswordOk

`func (o *ProvisionVDBFromBookmarkParameters) GetParentTdeKeystorePasswordOk() (*string, bool)`

GetParentTdeKeystorePasswordOk returns a tuple with the ParentTdeKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTdeKeystorePassword

`func (o *ProvisionVDBFromBookmarkParameters) SetParentTdeKeystorePassword(v string)`

SetParentTdeKeystorePassword sets ParentTdeKeystorePassword field to given value.

### HasParentTdeKeystorePassword

`func (o *ProvisionVDBFromBookmarkParameters) HasParentTdeKeystorePassword() bool`

HasParentTdeKeystorePassword returns a boolean if a field has been set.

### GetTdeExportedKeyFileSecret

`func (o *ProvisionVDBFromBookmarkParameters) GetTdeExportedKeyFileSecret() string`

GetTdeExportedKeyFileSecret returns the TdeExportedKeyFileSecret field if non-nil, zero value otherwise.

### GetTdeExportedKeyFileSecretOk

`func (o *ProvisionVDBFromBookmarkParameters) GetTdeExportedKeyFileSecretOk() (*string, bool)`

GetTdeExportedKeyFileSecretOk returns a tuple with the TdeExportedKeyFileSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdeExportedKeyFileSecret

`func (o *ProvisionVDBFromBookmarkParameters) SetTdeExportedKeyFileSecret(v string)`

SetTdeExportedKeyFileSecret sets TdeExportedKeyFileSecret field to given value.

### HasTdeExportedKeyFileSecret

`func (o *ProvisionVDBFromBookmarkParameters) HasTdeExportedKeyFileSecret() bool`

HasTdeExportedKeyFileSecret returns a boolean if a field has been set.

### GetTdeKeyIdentifier

`func (o *ProvisionVDBFromBookmarkParameters) GetTdeKeyIdentifier() string`

GetTdeKeyIdentifier returns the TdeKeyIdentifier field if non-nil, zero value otherwise.

### GetTdeKeyIdentifierOk

`func (o *ProvisionVDBFromBookmarkParameters) GetTdeKeyIdentifierOk() (*string, bool)`

GetTdeKeyIdentifierOk returns a tuple with the TdeKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdeKeyIdentifier

`func (o *ProvisionVDBFromBookmarkParameters) SetTdeKeyIdentifier(v string)`

SetTdeKeyIdentifier sets TdeKeyIdentifier field to given value.

### HasTdeKeyIdentifier

`func (o *ProvisionVDBFromBookmarkParameters) HasTdeKeyIdentifier() bool`

HasTdeKeyIdentifier returns a boolean if a field has been set.

### GetTargetVcdbTdeKeystorePath

`func (o *ProvisionVDBFromBookmarkParameters) GetTargetVcdbTdeKeystorePath() string`

GetTargetVcdbTdeKeystorePath returns the TargetVcdbTdeKeystorePath field if non-nil, zero value otherwise.

### GetTargetVcdbTdeKeystorePathOk

`func (o *ProvisionVDBFromBookmarkParameters) GetTargetVcdbTdeKeystorePathOk() (*string, bool)`

GetTargetVcdbTdeKeystorePathOk returns a tuple with the TargetVcdbTdeKeystorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVcdbTdeKeystorePath

`func (o *ProvisionVDBFromBookmarkParameters) SetTargetVcdbTdeKeystorePath(v string)`

SetTargetVcdbTdeKeystorePath sets TargetVcdbTdeKeystorePath field to given value.

### HasTargetVcdbTdeKeystorePath

`func (o *ProvisionVDBFromBookmarkParameters) HasTargetVcdbTdeKeystorePath() bool`

HasTargetVcdbTdeKeystorePath returns a boolean if a field has been set.

### GetCdbTdeKeystorePassword

`func (o *ProvisionVDBFromBookmarkParameters) GetCdbTdeKeystorePassword() string`

GetCdbTdeKeystorePassword returns the CdbTdeKeystorePassword field if non-nil, zero value otherwise.

### GetCdbTdeKeystorePasswordOk

`func (o *ProvisionVDBFromBookmarkParameters) GetCdbTdeKeystorePasswordOk() (*string, bool)`

GetCdbTdeKeystorePasswordOk returns a tuple with the CdbTdeKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdbTdeKeystorePassword

`func (o *ProvisionVDBFromBookmarkParameters) SetCdbTdeKeystorePassword(v string)`

SetCdbTdeKeystorePassword sets CdbTdeKeystorePassword field to given value.

### HasCdbTdeKeystorePassword

`func (o *ProvisionVDBFromBookmarkParameters) HasCdbTdeKeystorePassword() bool`

HasCdbTdeKeystorePassword returns a boolean if a field has been set.

### GetVcdbTdeKeyIdentifier

`func (o *ProvisionVDBFromBookmarkParameters) GetVcdbTdeKeyIdentifier() string`

GetVcdbTdeKeyIdentifier returns the VcdbTdeKeyIdentifier field if non-nil, zero value otherwise.

### GetVcdbTdeKeyIdentifierOk

`func (o *ProvisionVDBFromBookmarkParameters) GetVcdbTdeKeyIdentifierOk() (*string, bool)`

GetVcdbTdeKeyIdentifierOk returns a tuple with the VcdbTdeKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdbTdeKeyIdentifier

`func (o *ProvisionVDBFromBookmarkParameters) SetVcdbTdeKeyIdentifier(v string)`

SetVcdbTdeKeyIdentifier sets VcdbTdeKeyIdentifier field to given value.

### HasVcdbTdeKeyIdentifier

`func (o *ProvisionVDBFromBookmarkParameters) HasVcdbTdeKeyIdentifier() bool`

HasVcdbTdeKeyIdentifier returns a boolean if a field has been set.

### GetAppdataSourceParams

`func (o *ProvisionVDBFromBookmarkParameters) GetAppdataSourceParams() map[string]interface{}`

GetAppdataSourceParams returns the AppdataSourceParams field if non-nil, zero value otherwise.

### GetAppdataSourceParamsOk

`func (o *ProvisionVDBFromBookmarkParameters) GetAppdataSourceParamsOk() (*map[string]interface{}, bool)`

GetAppdataSourceParamsOk returns a tuple with the AppdataSourceParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppdataSourceParams

`func (o *ProvisionVDBFromBookmarkParameters) SetAppdataSourceParams(v map[string]interface{})`

SetAppdataSourceParams sets AppdataSourceParams field to given value.

### HasAppdataSourceParams

`func (o *ProvisionVDBFromBookmarkParameters) HasAppdataSourceParams() bool`

HasAppdataSourceParams returns a boolean if a field has been set.

### GetAdditionalMountPoints

`func (o *ProvisionVDBFromBookmarkParameters) GetAdditionalMountPoints() []AdditionalMountPoint`

GetAdditionalMountPoints returns the AdditionalMountPoints field if non-nil, zero value otherwise.

### GetAdditionalMountPointsOk

`func (o *ProvisionVDBFromBookmarkParameters) GetAdditionalMountPointsOk() (*[]AdditionalMountPoint, bool)`

GetAdditionalMountPointsOk returns a tuple with the AdditionalMountPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMountPoints

`func (o *ProvisionVDBFromBookmarkParameters) SetAdditionalMountPoints(v []AdditionalMountPoint)`

SetAdditionalMountPoints sets AdditionalMountPoints field to given value.

### HasAdditionalMountPoints

`func (o *ProvisionVDBFromBookmarkParameters) HasAdditionalMountPoints() bool`

HasAdditionalMountPoints returns a boolean if a field has been set.

### SetAdditionalMountPointsNil

`func (o *ProvisionVDBFromBookmarkParameters) SetAdditionalMountPointsNil(b bool)`

 SetAdditionalMountPointsNil sets the value for AdditionalMountPoints to be an explicit nil

### UnsetAdditionalMountPoints
`func (o *ProvisionVDBFromBookmarkParameters) UnsetAdditionalMountPoints()`

UnsetAdditionalMountPoints ensures that no value is present for AdditionalMountPoints, not even an explicit nil
### GetAppdataConfigParams

`func (o *ProvisionVDBFromBookmarkParameters) GetAppdataConfigParams() map[string]interface{}`

GetAppdataConfigParams returns the AppdataConfigParams field if non-nil, zero value otherwise.

### GetAppdataConfigParamsOk

`func (o *ProvisionVDBFromBookmarkParameters) GetAppdataConfigParamsOk() (*map[string]interface{}, bool)`

GetAppdataConfigParamsOk returns a tuple with the AppdataConfigParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppdataConfigParams

`func (o *ProvisionVDBFromBookmarkParameters) SetAppdataConfigParams(v map[string]interface{})`

SetAppdataConfigParams sets AppdataConfigParams field to given value.

### HasAppdataConfigParams

`func (o *ProvisionVDBFromBookmarkParameters) HasAppdataConfigParams() bool`

HasAppdataConfigParams returns a boolean if a field has been set.

### SetAppdataConfigParamsNil

`func (o *ProvisionVDBFromBookmarkParameters) SetAppdataConfigParamsNil(b bool)`

 SetAppdataConfigParamsNil sets the value for AppdataConfigParams to be an explicit nil

### UnsetAppdataConfigParams
`func (o *ProvisionVDBFromBookmarkParameters) UnsetAppdataConfigParams()`

UnsetAppdataConfigParams ensures that no value is present for AppdataConfigParams, not even an explicit nil
### GetConfigParams

`func (o *ProvisionVDBFromBookmarkParameters) GetConfigParams() map[string]interface{}`

GetConfigParams returns the ConfigParams field if non-nil, zero value otherwise.

### GetConfigParamsOk

`func (o *ProvisionVDBFromBookmarkParameters) GetConfigParamsOk() (*map[string]interface{}, bool)`

GetConfigParamsOk returns a tuple with the ConfigParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParams

`func (o *ProvisionVDBFromBookmarkParameters) SetConfigParams(v map[string]interface{})`

SetConfigParams sets ConfigParams field to given value.

### HasConfigParams

`func (o *ProvisionVDBFromBookmarkParameters) HasConfigParams() bool`

HasConfigParams returns a boolean if a field has been set.

### SetConfigParamsNil

`func (o *ProvisionVDBFromBookmarkParameters) SetConfigParamsNil(b bool)`

 SetConfigParamsNil sets the value for ConfigParams to be an explicit nil

### UnsetConfigParams
`func (o *ProvisionVDBFromBookmarkParameters) UnsetConfigParams()`

UnsetConfigParams ensures that no value is present for ConfigParams, not even an explicit nil
### GetTags

`func (o *ProvisionVDBFromBookmarkParameters) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProvisionVDBFromBookmarkParameters) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProvisionVDBFromBookmarkParameters) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProvisionVDBFromBookmarkParameters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetBookmarkId

`func (o *ProvisionVDBFromBookmarkParameters) GetBookmarkId() string`

GetBookmarkId returns the BookmarkId field if non-nil, zero value otherwise.

### GetBookmarkIdOk

`func (o *ProvisionVDBFromBookmarkParameters) GetBookmarkIdOk() (*string, bool)`

GetBookmarkIdOk returns a tuple with the BookmarkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarkId

`func (o *ProvisionVDBFromBookmarkParameters) SetBookmarkId(v string)`

SetBookmarkId sets BookmarkId field to given value.


### GetMakeCurrentAccountOwner

`func (o *ProvisionVDBFromBookmarkParameters) GetMakeCurrentAccountOwner() bool`

GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field if non-nil, zero value otherwise.

### GetMakeCurrentAccountOwnerOk

`func (o *ProvisionVDBFromBookmarkParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool)`

GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCurrentAccountOwner

`func (o *ProvisionVDBFromBookmarkParameters) SetMakeCurrentAccountOwner(v bool)`

SetMakeCurrentAccountOwner sets MakeCurrentAccountOwner field to given value.

### HasMakeCurrentAccountOwner

`func (o *ProvisionVDBFromBookmarkParameters) HasMakeCurrentAccountOwner() bool`

HasMakeCurrentAccountOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


