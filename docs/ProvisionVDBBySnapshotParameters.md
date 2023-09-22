# ProvisionVDBBySnapshotParameters

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
**SnapshotId** | Pointer to **string** | The ID of the snapshot from which to execute the operation. If the snapshot_id is not, selects the latest snapshot. | [optional] 
**EngineId** | Pointer to **string** | The ID of the Engine onto which to provision. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored. | [optional] 
**SourceDataId** | Pointer to **string** | The ID of the source object (dSource or VDB) to provision from. All other objects referenced by the parameters must live on the same engine as the source. If this property is not set, the data_source of the snapshot_id will be used. | [optional] 
**MakeCurrentAccountOwner** | Pointer to **bool** | Whether the account provisioning this VDB must be configured as owner of the VDB. | [optional] [default to true]

## Methods

### NewProvisionVDBBySnapshotParameters

`func NewProvisionVDBBySnapshotParameters() *ProvisionVDBBySnapshotParameters`

NewProvisionVDBBySnapshotParameters instantiates a new ProvisionVDBBySnapshotParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionVDBBySnapshotParametersWithDefaults

`func NewProvisionVDBBySnapshotParametersWithDefaults() *ProvisionVDBBySnapshotParameters`

NewProvisionVDBBySnapshotParametersWithDefaults instantiates a new ProvisionVDBBySnapshotParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreRefresh

`func (o *ProvisionVDBBySnapshotParameters) GetPreRefresh() []Hook`

GetPreRefresh returns the PreRefresh field if non-nil, zero value otherwise.

### GetPreRefreshOk

`func (o *ProvisionVDBBySnapshotParameters) GetPreRefreshOk() (*[]Hook, bool)`

GetPreRefreshOk returns a tuple with the PreRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRefresh

`func (o *ProvisionVDBBySnapshotParameters) SetPreRefresh(v []Hook)`

SetPreRefresh sets PreRefresh field to given value.

### HasPreRefresh

`func (o *ProvisionVDBBySnapshotParameters) HasPreRefresh() bool`

HasPreRefresh returns a boolean if a field has been set.

### GetPostRefresh

`func (o *ProvisionVDBBySnapshotParameters) GetPostRefresh() []Hook`

GetPostRefresh returns the PostRefresh field if non-nil, zero value otherwise.

### GetPostRefreshOk

`func (o *ProvisionVDBBySnapshotParameters) GetPostRefreshOk() (*[]Hook, bool)`

GetPostRefreshOk returns a tuple with the PostRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostRefresh

`func (o *ProvisionVDBBySnapshotParameters) SetPostRefresh(v []Hook)`

SetPostRefresh sets PostRefresh field to given value.

### HasPostRefresh

`func (o *ProvisionVDBBySnapshotParameters) HasPostRefresh() bool`

HasPostRefresh returns a boolean if a field has been set.

### GetPreRollback

`func (o *ProvisionVDBBySnapshotParameters) GetPreRollback() []Hook`

GetPreRollback returns the PreRollback field if non-nil, zero value otherwise.

### GetPreRollbackOk

`func (o *ProvisionVDBBySnapshotParameters) GetPreRollbackOk() (*[]Hook, bool)`

GetPreRollbackOk returns a tuple with the PreRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRollback

`func (o *ProvisionVDBBySnapshotParameters) SetPreRollback(v []Hook)`

SetPreRollback sets PreRollback field to given value.

### HasPreRollback

`func (o *ProvisionVDBBySnapshotParameters) HasPreRollback() bool`

HasPreRollback returns a boolean if a field has been set.

### GetPostRollback

`func (o *ProvisionVDBBySnapshotParameters) GetPostRollback() []Hook`

GetPostRollback returns the PostRollback field if non-nil, zero value otherwise.

### GetPostRollbackOk

`func (o *ProvisionVDBBySnapshotParameters) GetPostRollbackOk() (*[]Hook, bool)`

GetPostRollbackOk returns a tuple with the PostRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostRollback

`func (o *ProvisionVDBBySnapshotParameters) SetPostRollback(v []Hook)`

SetPostRollback sets PostRollback field to given value.

### HasPostRollback

`func (o *ProvisionVDBBySnapshotParameters) HasPostRollback() bool`

HasPostRollback returns a boolean if a field has been set.

### GetConfigureClone

`func (o *ProvisionVDBBySnapshotParameters) GetConfigureClone() []Hook`

GetConfigureClone returns the ConfigureClone field if non-nil, zero value otherwise.

### GetConfigureCloneOk

`func (o *ProvisionVDBBySnapshotParameters) GetConfigureCloneOk() (*[]Hook, bool)`

GetConfigureCloneOk returns a tuple with the ConfigureClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureClone

`func (o *ProvisionVDBBySnapshotParameters) SetConfigureClone(v []Hook)`

SetConfigureClone sets ConfigureClone field to given value.

### HasConfigureClone

`func (o *ProvisionVDBBySnapshotParameters) HasConfigureClone() bool`

HasConfigureClone returns a boolean if a field has been set.

### GetPreSnapshot

`func (o *ProvisionVDBBySnapshotParameters) GetPreSnapshot() []Hook`

GetPreSnapshot returns the PreSnapshot field if non-nil, zero value otherwise.

### GetPreSnapshotOk

`func (o *ProvisionVDBBySnapshotParameters) GetPreSnapshotOk() (*[]Hook, bool)`

GetPreSnapshotOk returns a tuple with the PreSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSnapshot

`func (o *ProvisionVDBBySnapshotParameters) SetPreSnapshot(v []Hook)`

SetPreSnapshot sets PreSnapshot field to given value.

### HasPreSnapshot

`func (o *ProvisionVDBBySnapshotParameters) HasPreSnapshot() bool`

HasPreSnapshot returns a boolean if a field has been set.

### GetPostSnapshot

`func (o *ProvisionVDBBySnapshotParameters) GetPostSnapshot() []Hook`

GetPostSnapshot returns the PostSnapshot field if non-nil, zero value otherwise.

### GetPostSnapshotOk

`func (o *ProvisionVDBBySnapshotParameters) GetPostSnapshotOk() (*[]Hook, bool)`

GetPostSnapshotOk returns a tuple with the PostSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostSnapshot

`func (o *ProvisionVDBBySnapshotParameters) SetPostSnapshot(v []Hook)`

SetPostSnapshot sets PostSnapshot field to given value.

### HasPostSnapshot

`func (o *ProvisionVDBBySnapshotParameters) HasPostSnapshot() bool`

HasPostSnapshot returns a boolean if a field has been set.

### GetPreStart

`func (o *ProvisionVDBBySnapshotParameters) GetPreStart() []Hook`

GetPreStart returns the PreStart field if non-nil, zero value otherwise.

### GetPreStartOk

`func (o *ProvisionVDBBySnapshotParameters) GetPreStartOk() (*[]Hook, bool)`

GetPreStartOk returns a tuple with the PreStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreStart

`func (o *ProvisionVDBBySnapshotParameters) SetPreStart(v []Hook)`

SetPreStart sets PreStart field to given value.

### HasPreStart

`func (o *ProvisionVDBBySnapshotParameters) HasPreStart() bool`

HasPreStart returns a boolean if a field has been set.

### GetPostStart

`func (o *ProvisionVDBBySnapshotParameters) GetPostStart() []Hook`

GetPostStart returns the PostStart field if non-nil, zero value otherwise.

### GetPostStartOk

`func (o *ProvisionVDBBySnapshotParameters) GetPostStartOk() (*[]Hook, bool)`

GetPostStartOk returns a tuple with the PostStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostStart

`func (o *ProvisionVDBBySnapshotParameters) SetPostStart(v []Hook)`

SetPostStart sets PostStart field to given value.

### HasPostStart

`func (o *ProvisionVDBBySnapshotParameters) HasPostStart() bool`

HasPostStart returns a boolean if a field has been set.

### GetPreStop

`func (o *ProvisionVDBBySnapshotParameters) GetPreStop() []Hook`

GetPreStop returns the PreStop field if non-nil, zero value otherwise.

### GetPreStopOk

`func (o *ProvisionVDBBySnapshotParameters) GetPreStopOk() (*[]Hook, bool)`

GetPreStopOk returns a tuple with the PreStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreStop

`func (o *ProvisionVDBBySnapshotParameters) SetPreStop(v []Hook)`

SetPreStop sets PreStop field to given value.

### HasPreStop

`func (o *ProvisionVDBBySnapshotParameters) HasPreStop() bool`

HasPreStop returns a boolean if a field has been set.

### GetPostStop

`func (o *ProvisionVDBBySnapshotParameters) GetPostStop() []Hook`

GetPostStop returns the PostStop field if non-nil, zero value otherwise.

### GetPostStopOk

`func (o *ProvisionVDBBySnapshotParameters) GetPostStopOk() (*[]Hook, bool)`

GetPostStopOk returns a tuple with the PostStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostStop

`func (o *ProvisionVDBBySnapshotParameters) SetPostStop(v []Hook)`

SetPostStop sets PostStop field to given value.

### HasPostStop

`func (o *ProvisionVDBBySnapshotParameters) HasPostStop() bool`

HasPostStop returns a boolean if a field has been set.

### GetTargetGroupId

`func (o *ProvisionVDBBySnapshotParameters) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *ProvisionVDBBySnapshotParameters) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *ProvisionVDBBySnapshotParameters) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *ProvisionVDBBySnapshotParameters) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### GetName

`func (o *ProvisionVDBBySnapshotParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisionVDBBySnapshotParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisionVDBBySnapshotParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProvisionVDBBySnapshotParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDatabaseName

`func (o *ProvisionVDBBySnapshotParameters) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ProvisionVDBBySnapshotParameters) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ProvisionVDBBySnapshotParameters) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *ProvisionVDBBySnapshotParameters) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetCdbId

`func (o *ProvisionVDBBySnapshotParameters) GetCdbId() string`

GetCdbId returns the CdbId field if non-nil, zero value otherwise.

### GetCdbIdOk

`func (o *ProvisionVDBBySnapshotParameters) GetCdbIdOk() (*string, bool)`

GetCdbIdOk returns a tuple with the CdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdbId

`func (o *ProvisionVDBBySnapshotParameters) SetCdbId(v string)`

SetCdbId sets CdbId field to given value.

### HasCdbId

`func (o *ProvisionVDBBySnapshotParameters) HasCdbId() bool`

HasCdbId returns a boolean if a field has been set.

### GetClusterNodeIds

`func (o *ProvisionVDBBySnapshotParameters) GetClusterNodeIds() []string`

GetClusterNodeIds returns the ClusterNodeIds field if non-nil, zero value otherwise.

### GetClusterNodeIdsOk

`func (o *ProvisionVDBBySnapshotParameters) GetClusterNodeIdsOk() (*[]string, bool)`

GetClusterNodeIdsOk returns a tuple with the ClusterNodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeIds

`func (o *ProvisionVDBBySnapshotParameters) SetClusterNodeIds(v []string)`

SetClusterNodeIds sets ClusterNodeIds field to given value.

### HasClusterNodeIds

`func (o *ProvisionVDBBySnapshotParameters) HasClusterNodeIds() bool`

HasClusterNodeIds returns a boolean if a field has been set.

### GetTruncateLogOnCheckpoint

`func (o *ProvisionVDBBySnapshotParameters) GetTruncateLogOnCheckpoint() bool`

GetTruncateLogOnCheckpoint returns the TruncateLogOnCheckpoint field if non-nil, zero value otherwise.

### GetTruncateLogOnCheckpointOk

`func (o *ProvisionVDBBySnapshotParameters) GetTruncateLogOnCheckpointOk() (*bool, bool)`

GetTruncateLogOnCheckpointOk returns a tuple with the TruncateLogOnCheckpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateLogOnCheckpoint

`func (o *ProvisionVDBBySnapshotParameters) SetTruncateLogOnCheckpoint(v bool)`

SetTruncateLogOnCheckpoint sets TruncateLogOnCheckpoint field to given value.

### HasTruncateLogOnCheckpoint

`func (o *ProvisionVDBBySnapshotParameters) HasTruncateLogOnCheckpoint() bool`

HasTruncateLogOnCheckpoint returns a boolean if a field has been set.

### GetOsUsername

`func (o *ProvisionVDBBySnapshotParameters) GetOsUsername() string`

GetOsUsername returns the OsUsername field if non-nil, zero value otherwise.

### GetOsUsernameOk

`func (o *ProvisionVDBBySnapshotParameters) GetOsUsernameOk() (*string, bool)`

GetOsUsernameOk returns a tuple with the OsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsUsername

`func (o *ProvisionVDBBySnapshotParameters) SetOsUsername(v string)`

SetOsUsername sets OsUsername field to given value.

### HasOsUsername

`func (o *ProvisionVDBBySnapshotParameters) HasOsUsername() bool`

HasOsUsername returns a boolean if a field has been set.

### GetOsPassword

`func (o *ProvisionVDBBySnapshotParameters) GetOsPassword() string`

GetOsPassword returns the OsPassword field if non-nil, zero value otherwise.

### GetOsPasswordOk

`func (o *ProvisionVDBBySnapshotParameters) GetOsPasswordOk() (*string, bool)`

GetOsPasswordOk returns a tuple with the OsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPassword

`func (o *ProvisionVDBBySnapshotParameters) SetOsPassword(v string)`

SetOsPassword sets OsPassword field to given value.

### HasOsPassword

`func (o *ProvisionVDBBySnapshotParameters) HasOsPassword() bool`

HasOsPassword returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *ProvisionVDBBySnapshotParameters) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ProvisionVDBBySnapshotParameters) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ProvisionVDBBySnapshotParameters) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ProvisionVDBBySnapshotParameters) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEnvironmentUserId

`func (o *ProvisionVDBBySnapshotParameters) GetEnvironmentUserId() string`

GetEnvironmentUserId returns the EnvironmentUserId field if non-nil, zero value otherwise.

### GetEnvironmentUserIdOk

`func (o *ProvisionVDBBySnapshotParameters) GetEnvironmentUserIdOk() (*string, bool)`

GetEnvironmentUserIdOk returns a tuple with the EnvironmentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUserId

`func (o *ProvisionVDBBySnapshotParameters) SetEnvironmentUserId(v string)`

SetEnvironmentUserId sets EnvironmentUserId field to given value.

### HasEnvironmentUserId

`func (o *ProvisionVDBBySnapshotParameters) HasEnvironmentUserId() bool`

HasEnvironmentUserId returns a boolean if a field has been set.

### GetRepositoryId

`func (o *ProvisionVDBBySnapshotParameters) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *ProvisionVDBBySnapshotParameters) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *ProvisionVDBBySnapshotParameters) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *ProvisionVDBBySnapshotParameters) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetAutoSelectRepository

`func (o *ProvisionVDBBySnapshotParameters) GetAutoSelectRepository() bool`

GetAutoSelectRepository returns the AutoSelectRepository field if non-nil, zero value otherwise.

### GetAutoSelectRepositoryOk

`func (o *ProvisionVDBBySnapshotParameters) GetAutoSelectRepositoryOk() (*bool, bool)`

GetAutoSelectRepositoryOk returns a tuple with the AutoSelectRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelectRepository

`func (o *ProvisionVDBBySnapshotParameters) SetAutoSelectRepository(v bool)`

SetAutoSelectRepository sets AutoSelectRepository field to given value.

### HasAutoSelectRepository

`func (o *ProvisionVDBBySnapshotParameters) HasAutoSelectRepository() bool`

HasAutoSelectRepository returns a boolean if a field has been set.

### GetVdbRestart

`func (o *ProvisionVDBBySnapshotParameters) GetVdbRestart() bool`

GetVdbRestart returns the VdbRestart field if non-nil, zero value otherwise.

### GetVdbRestartOk

`func (o *ProvisionVDBBySnapshotParameters) GetVdbRestartOk() (*bool, bool)`

GetVdbRestartOk returns a tuple with the VdbRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdbRestart

`func (o *ProvisionVDBBySnapshotParameters) SetVdbRestart(v bool)`

SetVdbRestart sets VdbRestart field to given value.

### HasVdbRestart

`func (o *ProvisionVDBBySnapshotParameters) HasVdbRestart() bool`

HasVdbRestart returns a boolean if a field has been set.

### GetTemplateId

`func (o *ProvisionVDBBySnapshotParameters) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ProvisionVDBBySnapshotParameters) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ProvisionVDBBySnapshotParameters) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ProvisionVDBBySnapshotParameters) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetAuxiliaryTemplateId

`func (o *ProvisionVDBBySnapshotParameters) GetAuxiliaryTemplateId() string`

GetAuxiliaryTemplateId returns the AuxiliaryTemplateId field if non-nil, zero value otherwise.

### GetAuxiliaryTemplateIdOk

`func (o *ProvisionVDBBySnapshotParameters) GetAuxiliaryTemplateIdOk() (*string, bool)`

GetAuxiliaryTemplateIdOk returns a tuple with the AuxiliaryTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryTemplateId

`func (o *ProvisionVDBBySnapshotParameters) SetAuxiliaryTemplateId(v string)`

SetAuxiliaryTemplateId sets AuxiliaryTemplateId field to given value.

### HasAuxiliaryTemplateId

`func (o *ProvisionVDBBySnapshotParameters) HasAuxiliaryTemplateId() bool`

HasAuxiliaryTemplateId returns a boolean if a field has been set.

### GetFileMappingRules

`func (o *ProvisionVDBBySnapshotParameters) GetFileMappingRules() string`

GetFileMappingRules returns the FileMappingRules field if non-nil, zero value otherwise.

### GetFileMappingRulesOk

`func (o *ProvisionVDBBySnapshotParameters) GetFileMappingRulesOk() (*string, bool)`

GetFileMappingRulesOk returns a tuple with the FileMappingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMappingRules

`func (o *ProvisionVDBBySnapshotParameters) SetFileMappingRules(v string)`

SetFileMappingRules sets FileMappingRules field to given value.

### HasFileMappingRules

`func (o *ProvisionVDBBySnapshotParameters) HasFileMappingRules() bool`

HasFileMappingRules returns a boolean if a field has been set.

### GetOracleInstanceName

`func (o *ProvisionVDBBySnapshotParameters) GetOracleInstanceName() string`

GetOracleInstanceName returns the OracleInstanceName field if non-nil, zero value otherwise.

### GetOracleInstanceNameOk

`func (o *ProvisionVDBBySnapshotParameters) GetOracleInstanceNameOk() (*string, bool)`

GetOracleInstanceNameOk returns a tuple with the OracleInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleInstanceName

`func (o *ProvisionVDBBySnapshotParameters) SetOracleInstanceName(v string)`

SetOracleInstanceName sets OracleInstanceName field to given value.

### HasOracleInstanceName

`func (o *ProvisionVDBBySnapshotParameters) HasOracleInstanceName() bool`

HasOracleInstanceName returns a boolean if a field has been set.

### GetUniqueName

`func (o *ProvisionVDBBySnapshotParameters) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *ProvisionVDBBySnapshotParameters) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *ProvisionVDBBySnapshotParameters) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *ProvisionVDBBySnapshotParameters) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetVcdbName

`func (o *ProvisionVDBBySnapshotParameters) GetVcdbName() string`

GetVcdbName returns the VcdbName field if non-nil, zero value otherwise.

### GetVcdbNameOk

`func (o *ProvisionVDBBySnapshotParameters) GetVcdbNameOk() (*string, bool)`

GetVcdbNameOk returns a tuple with the VcdbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdbName

`func (o *ProvisionVDBBySnapshotParameters) SetVcdbName(v string)`

SetVcdbName sets VcdbName field to given value.

### HasVcdbName

`func (o *ProvisionVDBBySnapshotParameters) HasVcdbName() bool`

HasVcdbName returns a boolean if a field has been set.

### GetVcdbDatabaseName

`func (o *ProvisionVDBBySnapshotParameters) GetVcdbDatabaseName() string`

GetVcdbDatabaseName returns the VcdbDatabaseName field if non-nil, zero value otherwise.

### GetVcdbDatabaseNameOk

`func (o *ProvisionVDBBySnapshotParameters) GetVcdbDatabaseNameOk() (*string, bool)`

GetVcdbDatabaseNameOk returns a tuple with the VcdbDatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdbDatabaseName

`func (o *ProvisionVDBBySnapshotParameters) SetVcdbDatabaseName(v string)`

SetVcdbDatabaseName sets VcdbDatabaseName field to given value.

### HasVcdbDatabaseName

`func (o *ProvisionVDBBySnapshotParameters) HasVcdbDatabaseName() bool`

HasVcdbDatabaseName returns a boolean if a field has been set.

### GetMountPoint

`func (o *ProvisionVDBBySnapshotParameters) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *ProvisionVDBBySnapshotParameters) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *ProvisionVDBBySnapshotParameters) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *ProvisionVDBBySnapshotParameters) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### GetOpenResetLogs

`func (o *ProvisionVDBBySnapshotParameters) GetOpenResetLogs() bool`

GetOpenResetLogs returns the OpenResetLogs field if non-nil, zero value otherwise.

### GetOpenResetLogsOk

`func (o *ProvisionVDBBySnapshotParameters) GetOpenResetLogsOk() (*bool, bool)`

GetOpenResetLogsOk returns a tuple with the OpenResetLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenResetLogs

`func (o *ProvisionVDBBySnapshotParameters) SetOpenResetLogs(v bool)`

SetOpenResetLogs sets OpenResetLogs field to given value.

### HasOpenResetLogs

`func (o *ProvisionVDBBySnapshotParameters) HasOpenResetLogs() bool`

HasOpenResetLogs returns a boolean if a field has been set.

### GetSnapshotPolicyId

`func (o *ProvisionVDBBySnapshotParameters) GetSnapshotPolicyId() string`

GetSnapshotPolicyId returns the SnapshotPolicyId field if non-nil, zero value otherwise.

### GetSnapshotPolicyIdOk

`func (o *ProvisionVDBBySnapshotParameters) GetSnapshotPolicyIdOk() (*string, bool)`

GetSnapshotPolicyIdOk returns a tuple with the SnapshotPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotPolicyId

`func (o *ProvisionVDBBySnapshotParameters) SetSnapshotPolicyId(v string)`

SetSnapshotPolicyId sets SnapshotPolicyId field to given value.

### HasSnapshotPolicyId

`func (o *ProvisionVDBBySnapshotParameters) HasSnapshotPolicyId() bool`

HasSnapshotPolicyId returns a boolean if a field has been set.

### GetRetentionPolicyId

`func (o *ProvisionVDBBySnapshotParameters) GetRetentionPolicyId() string`

GetRetentionPolicyId returns the RetentionPolicyId field if non-nil, zero value otherwise.

### GetRetentionPolicyIdOk

`func (o *ProvisionVDBBySnapshotParameters) GetRetentionPolicyIdOk() (*string, bool)`

GetRetentionPolicyIdOk returns a tuple with the RetentionPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicyId

`func (o *ProvisionVDBBySnapshotParameters) SetRetentionPolicyId(v string)`

SetRetentionPolicyId sets RetentionPolicyId field to given value.

### HasRetentionPolicyId

`func (o *ProvisionVDBBySnapshotParameters) HasRetentionPolicyId() bool`

HasRetentionPolicyId returns a boolean if a field has been set.

### GetRecoveryModel

`func (o *ProvisionVDBBySnapshotParameters) GetRecoveryModel() string`

GetRecoveryModel returns the RecoveryModel field if non-nil, zero value otherwise.

### GetRecoveryModelOk

`func (o *ProvisionVDBBySnapshotParameters) GetRecoveryModelOk() (*string, bool)`

GetRecoveryModelOk returns a tuple with the RecoveryModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryModel

`func (o *ProvisionVDBBySnapshotParameters) SetRecoveryModel(v string)`

SetRecoveryModel sets RecoveryModel field to given value.

### HasRecoveryModel

`func (o *ProvisionVDBBySnapshotParameters) HasRecoveryModel() bool`

HasRecoveryModel returns a boolean if a field has been set.

### GetPreScript

`func (o *ProvisionVDBBySnapshotParameters) GetPreScript() string`

GetPreScript returns the PreScript field if non-nil, zero value otherwise.

### GetPreScriptOk

`func (o *ProvisionVDBBySnapshotParameters) GetPreScriptOk() (*string, bool)`

GetPreScriptOk returns a tuple with the PreScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreScript

`func (o *ProvisionVDBBySnapshotParameters) SetPreScript(v string)`

SetPreScript sets PreScript field to given value.

### HasPreScript

`func (o *ProvisionVDBBySnapshotParameters) HasPreScript() bool`

HasPreScript returns a boolean if a field has been set.

### GetPostScript

`func (o *ProvisionVDBBySnapshotParameters) GetPostScript() string`

GetPostScript returns the PostScript field if non-nil, zero value otherwise.

### GetPostScriptOk

`func (o *ProvisionVDBBySnapshotParameters) GetPostScriptOk() (*string, bool)`

GetPostScriptOk returns a tuple with the PostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScript

`func (o *ProvisionVDBBySnapshotParameters) SetPostScript(v string)`

SetPostScript sets PostScript field to given value.

### HasPostScript

`func (o *ProvisionVDBBySnapshotParameters) HasPostScript() bool`

HasPostScript returns a boolean if a field has been set.

### GetCdcOnProvision

`func (o *ProvisionVDBBySnapshotParameters) GetCdcOnProvision() bool`

GetCdcOnProvision returns the CdcOnProvision field if non-nil, zero value otherwise.

### GetCdcOnProvisionOk

`func (o *ProvisionVDBBySnapshotParameters) GetCdcOnProvisionOk() (*bool, bool)`

GetCdcOnProvisionOk returns a tuple with the CdcOnProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdcOnProvision

`func (o *ProvisionVDBBySnapshotParameters) SetCdcOnProvision(v bool)`

SetCdcOnProvision sets CdcOnProvision field to given value.

### HasCdcOnProvision

`func (o *ProvisionVDBBySnapshotParameters) HasCdcOnProvision() bool`

HasCdcOnProvision returns a boolean if a field has been set.

### GetOnlineLogSize

`func (o *ProvisionVDBBySnapshotParameters) GetOnlineLogSize() int32`

GetOnlineLogSize returns the OnlineLogSize field if non-nil, zero value otherwise.

### GetOnlineLogSizeOk

`func (o *ProvisionVDBBySnapshotParameters) GetOnlineLogSizeOk() (*int32, bool)`

GetOnlineLogSizeOk returns a tuple with the OnlineLogSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineLogSize

`func (o *ProvisionVDBBySnapshotParameters) SetOnlineLogSize(v int32)`

SetOnlineLogSize sets OnlineLogSize field to given value.

### HasOnlineLogSize

`func (o *ProvisionVDBBySnapshotParameters) HasOnlineLogSize() bool`

HasOnlineLogSize returns a boolean if a field has been set.

### GetOnlineLogGroups

`func (o *ProvisionVDBBySnapshotParameters) GetOnlineLogGroups() int32`

GetOnlineLogGroups returns the OnlineLogGroups field if non-nil, zero value otherwise.

### GetOnlineLogGroupsOk

`func (o *ProvisionVDBBySnapshotParameters) GetOnlineLogGroupsOk() (*int32, bool)`

GetOnlineLogGroupsOk returns a tuple with the OnlineLogGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineLogGroups

`func (o *ProvisionVDBBySnapshotParameters) SetOnlineLogGroups(v int32)`

SetOnlineLogGroups sets OnlineLogGroups field to given value.

### HasOnlineLogGroups

`func (o *ProvisionVDBBySnapshotParameters) HasOnlineLogGroups() bool`

HasOnlineLogGroups returns a boolean if a field has been set.

### GetArchiveLog

`func (o *ProvisionVDBBySnapshotParameters) GetArchiveLog() bool`

GetArchiveLog returns the ArchiveLog field if non-nil, zero value otherwise.

### GetArchiveLogOk

`func (o *ProvisionVDBBySnapshotParameters) GetArchiveLogOk() (*bool, bool)`

GetArchiveLogOk returns a tuple with the ArchiveLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLog

`func (o *ProvisionVDBBySnapshotParameters) SetArchiveLog(v bool)`

SetArchiveLog sets ArchiveLog field to given value.

### HasArchiveLog

`func (o *ProvisionVDBBySnapshotParameters) HasArchiveLog() bool`

HasArchiveLog returns a boolean if a field has been set.

### GetNewDbid

`func (o *ProvisionVDBBySnapshotParameters) GetNewDbid() bool`

GetNewDbid returns the NewDbid field if non-nil, zero value otherwise.

### GetNewDbidOk

`func (o *ProvisionVDBBySnapshotParameters) GetNewDbidOk() (*bool, bool)`

GetNewDbidOk returns a tuple with the NewDbid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDbid

`func (o *ProvisionVDBBySnapshotParameters) SetNewDbid(v bool)`

SetNewDbid sets NewDbid field to given value.

### HasNewDbid

`func (o *ProvisionVDBBySnapshotParameters) HasNewDbid() bool`

HasNewDbid returns a boolean if a field has been set.

### GetListenerIds

`func (o *ProvisionVDBBySnapshotParameters) GetListenerIds() []string`

GetListenerIds returns the ListenerIds field if non-nil, zero value otherwise.

### GetListenerIdsOk

`func (o *ProvisionVDBBySnapshotParameters) GetListenerIdsOk() (*[]string, bool)`

GetListenerIdsOk returns a tuple with the ListenerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerIds

`func (o *ProvisionVDBBySnapshotParameters) SetListenerIds(v []string)`

SetListenerIds sets ListenerIds field to given value.

### HasListenerIds

`func (o *ProvisionVDBBySnapshotParameters) HasListenerIds() bool`

HasListenerIds returns a boolean if a field has been set.

### GetCustomEnvVars

`func (o *ProvisionVDBBySnapshotParameters) GetCustomEnvVars() map[string]string`

GetCustomEnvVars returns the CustomEnvVars field if non-nil, zero value otherwise.

### GetCustomEnvVarsOk

`func (o *ProvisionVDBBySnapshotParameters) GetCustomEnvVarsOk() (*map[string]string, bool)`

GetCustomEnvVarsOk returns a tuple with the CustomEnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvVars

`func (o *ProvisionVDBBySnapshotParameters) SetCustomEnvVars(v map[string]string)`

SetCustomEnvVars sets CustomEnvVars field to given value.

### HasCustomEnvVars

`func (o *ProvisionVDBBySnapshotParameters) HasCustomEnvVars() bool`

HasCustomEnvVars returns a boolean if a field has been set.

### GetCustomEnvFiles

`func (o *ProvisionVDBBySnapshotParameters) GetCustomEnvFiles() []string`

GetCustomEnvFiles returns the CustomEnvFiles field if non-nil, zero value otherwise.

### GetCustomEnvFilesOk

`func (o *ProvisionVDBBySnapshotParameters) GetCustomEnvFilesOk() (*[]string, bool)`

GetCustomEnvFilesOk returns a tuple with the CustomEnvFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvFiles

`func (o *ProvisionVDBBySnapshotParameters) SetCustomEnvFiles(v []string)`

SetCustomEnvFiles sets CustomEnvFiles field to given value.

### HasCustomEnvFiles

`func (o *ProvisionVDBBySnapshotParameters) HasCustomEnvFiles() bool`

HasCustomEnvFiles returns a boolean if a field has been set.

### GetOracleRacCustomEnvFiles

`func (o *ProvisionVDBBySnapshotParameters) GetOracleRacCustomEnvFiles() []OracleRacCustomEnvFile`

GetOracleRacCustomEnvFiles returns the OracleRacCustomEnvFiles field if non-nil, zero value otherwise.

### GetOracleRacCustomEnvFilesOk

`func (o *ProvisionVDBBySnapshotParameters) GetOracleRacCustomEnvFilesOk() (*[]OracleRacCustomEnvFile, bool)`

GetOracleRacCustomEnvFilesOk returns a tuple with the OracleRacCustomEnvFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleRacCustomEnvFiles

`func (o *ProvisionVDBBySnapshotParameters) SetOracleRacCustomEnvFiles(v []OracleRacCustomEnvFile)`

SetOracleRacCustomEnvFiles sets OracleRacCustomEnvFiles field to given value.

### HasOracleRacCustomEnvFiles

`func (o *ProvisionVDBBySnapshotParameters) HasOracleRacCustomEnvFiles() bool`

HasOracleRacCustomEnvFiles returns a boolean if a field has been set.

### GetOracleRacCustomEnvVars

`func (o *ProvisionVDBBySnapshotParameters) GetOracleRacCustomEnvVars() []OracleRacCustomEnvVar`

GetOracleRacCustomEnvVars returns the OracleRacCustomEnvVars field if non-nil, zero value otherwise.

### GetOracleRacCustomEnvVarsOk

`func (o *ProvisionVDBBySnapshotParameters) GetOracleRacCustomEnvVarsOk() (*[]OracleRacCustomEnvVar, bool)`

GetOracleRacCustomEnvVarsOk returns a tuple with the OracleRacCustomEnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleRacCustomEnvVars

`func (o *ProvisionVDBBySnapshotParameters) SetOracleRacCustomEnvVars(v []OracleRacCustomEnvVar)`

SetOracleRacCustomEnvVars sets OracleRacCustomEnvVars field to given value.

### HasOracleRacCustomEnvVars

`func (o *ProvisionVDBBySnapshotParameters) HasOracleRacCustomEnvVars() bool`

HasOracleRacCustomEnvVars returns a boolean if a field has been set.

### GetParentTdeKeystorePath

`func (o *ProvisionVDBBySnapshotParameters) GetParentTdeKeystorePath() string`

GetParentTdeKeystorePath returns the ParentTdeKeystorePath field if non-nil, zero value otherwise.

### GetParentTdeKeystorePathOk

`func (o *ProvisionVDBBySnapshotParameters) GetParentTdeKeystorePathOk() (*string, bool)`

GetParentTdeKeystorePathOk returns a tuple with the ParentTdeKeystorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTdeKeystorePath

`func (o *ProvisionVDBBySnapshotParameters) SetParentTdeKeystorePath(v string)`

SetParentTdeKeystorePath sets ParentTdeKeystorePath field to given value.

### HasParentTdeKeystorePath

`func (o *ProvisionVDBBySnapshotParameters) HasParentTdeKeystorePath() bool`

HasParentTdeKeystorePath returns a boolean if a field has been set.

### GetParentTdeKeystorePassword

`func (o *ProvisionVDBBySnapshotParameters) GetParentTdeKeystorePassword() string`

GetParentTdeKeystorePassword returns the ParentTdeKeystorePassword field if non-nil, zero value otherwise.

### GetParentTdeKeystorePasswordOk

`func (o *ProvisionVDBBySnapshotParameters) GetParentTdeKeystorePasswordOk() (*string, bool)`

GetParentTdeKeystorePasswordOk returns a tuple with the ParentTdeKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTdeKeystorePassword

`func (o *ProvisionVDBBySnapshotParameters) SetParentTdeKeystorePassword(v string)`

SetParentTdeKeystorePassword sets ParentTdeKeystorePassword field to given value.

### HasParentTdeKeystorePassword

`func (o *ProvisionVDBBySnapshotParameters) HasParentTdeKeystorePassword() bool`

HasParentTdeKeystorePassword returns a boolean if a field has been set.

### GetTdeExportedKeyFileSecret

`func (o *ProvisionVDBBySnapshotParameters) GetTdeExportedKeyFileSecret() string`

GetTdeExportedKeyFileSecret returns the TdeExportedKeyFileSecret field if non-nil, zero value otherwise.

### GetTdeExportedKeyFileSecretOk

`func (o *ProvisionVDBBySnapshotParameters) GetTdeExportedKeyFileSecretOk() (*string, bool)`

GetTdeExportedKeyFileSecretOk returns a tuple with the TdeExportedKeyFileSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdeExportedKeyFileSecret

`func (o *ProvisionVDBBySnapshotParameters) SetTdeExportedKeyFileSecret(v string)`

SetTdeExportedKeyFileSecret sets TdeExportedKeyFileSecret field to given value.

### HasTdeExportedKeyFileSecret

`func (o *ProvisionVDBBySnapshotParameters) HasTdeExportedKeyFileSecret() bool`

HasTdeExportedKeyFileSecret returns a boolean if a field has been set.

### GetTdeKeyIdentifier

`func (o *ProvisionVDBBySnapshotParameters) GetTdeKeyIdentifier() string`

GetTdeKeyIdentifier returns the TdeKeyIdentifier field if non-nil, zero value otherwise.

### GetTdeKeyIdentifierOk

`func (o *ProvisionVDBBySnapshotParameters) GetTdeKeyIdentifierOk() (*string, bool)`

GetTdeKeyIdentifierOk returns a tuple with the TdeKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdeKeyIdentifier

`func (o *ProvisionVDBBySnapshotParameters) SetTdeKeyIdentifier(v string)`

SetTdeKeyIdentifier sets TdeKeyIdentifier field to given value.

### HasTdeKeyIdentifier

`func (o *ProvisionVDBBySnapshotParameters) HasTdeKeyIdentifier() bool`

HasTdeKeyIdentifier returns a boolean if a field has been set.

### GetTargetVcdbTdeKeystorePath

`func (o *ProvisionVDBBySnapshotParameters) GetTargetVcdbTdeKeystorePath() string`

GetTargetVcdbTdeKeystorePath returns the TargetVcdbTdeKeystorePath field if non-nil, zero value otherwise.

### GetTargetVcdbTdeKeystorePathOk

`func (o *ProvisionVDBBySnapshotParameters) GetTargetVcdbTdeKeystorePathOk() (*string, bool)`

GetTargetVcdbTdeKeystorePathOk returns a tuple with the TargetVcdbTdeKeystorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVcdbTdeKeystorePath

`func (o *ProvisionVDBBySnapshotParameters) SetTargetVcdbTdeKeystorePath(v string)`

SetTargetVcdbTdeKeystorePath sets TargetVcdbTdeKeystorePath field to given value.

### HasTargetVcdbTdeKeystorePath

`func (o *ProvisionVDBBySnapshotParameters) HasTargetVcdbTdeKeystorePath() bool`

HasTargetVcdbTdeKeystorePath returns a boolean if a field has been set.

### GetCdbTdeKeystorePassword

`func (o *ProvisionVDBBySnapshotParameters) GetCdbTdeKeystorePassword() string`

GetCdbTdeKeystorePassword returns the CdbTdeKeystorePassword field if non-nil, zero value otherwise.

### GetCdbTdeKeystorePasswordOk

`func (o *ProvisionVDBBySnapshotParameters) GetCdbTdeKeystorePasswordOk() (*string, bool)`

GetCdbTdeKeystorePasswordOk returns a tuple with the CdbTdeKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdbTdeKeystorePassword

`func (o *ProvisionVDBBySnapshotParameters) SetCdbTdeKeystorePassword(v string)`

SetCdbTdeKeystorePassword sets CdbTdeKeystorePassword field to given value.

### HasCdbTdeKeystorePassword

`func (o *ProvisionVDBBySnapshotParameters) HasCdbTdeKeystorePassword() bool`

HasCdbTdeKeystorePassword returns a boolean if a field has been set.

### GetVcdbTdeKeyIdentifier

`func (o *ProvisionVDBBySnapshotParameters) GetVcdbTdeKeyIdentifier() string`

GetVcdbTdeKeyIdentifier returns the VcdbTdeKeyIdentifier field if non-nil, zero value otherwise.

### GetVcdbTdeKeyIdentifierOk

`func (o *ProvisionVDBBySnapshotParameters) GetVcdbTdeKeyIdentifierOk() (*string, bool)`

GetVcdbTdeKeyIdentifierOk returns a tuple with the VcdbTdeKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdbTdeKeyIdentifier

`func (o *ProvisionVDBBySnapshotParameters) SetVcdbTdeKeyIdentifier(v string)`

SetVcdbTdeKeyIdentifier sets VcdbTdeKeyIdentifier field to given value.

### HasVcdbTdeKeyIdentifier

`func (o *ProvisionVDBBySnapshotParameters) HasVcdbTdeKeyIdentifier() bool`

HasVcdbTdeKeyIdentifier returns a boolean if a field has been set.

### GetAppdataSourceParams

`func (o *ProvisionVDBBySnapshotParameters) GetAppdataSourceParams() map[string]interface{}`

GetAppdataSourceParams returns the AppdataSourceParams field if non-nil, zero value otherwise.

### GetAppdataSourceParamsOk

`func (o *ProvisionVDBBySnapshotParameters) GetAppdataSourceParamsOk() (*map[string]interface{}, bool)`

GetAppdataSourceParamsOk returns a tuple with the AppdataSourceParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppdataSourceParams

`func (o *ProvisionVDBBySnapshotParameters) SetAppdataSourceParams(v map[string]interface{})`

SetAppdataSourceParams sets AppdataSourceParams field to given value.

### HasAppdataSourceParams

`func (o *ProvisionVDBBySnapshotParameters) HasAppdataSourceParams() bool`

HasAppdataSourceParams returns a boolean if a field has been set.

### GetAdditionalMountPoints

`func (o *ProvisionVDBBySnapshotParameters) GetAdditionalMountPoints() []AdditionalMountPoint`

GetAdditionalMountPoints returns the AdditionalMountPoints field if non-nil, zero value otherwise.

### GetAdditionalMountPointsOk

`func (o *ProvisionVDBBySnapshotParameters) GetAdditionalMountPointsOk() (*[]AdditionalMountPoint, bool)`

GetAdditionalMountPointsOk returns a tuple with the AdditionalMountPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMountPoints

`func (o *ProvisionVDBBySnapshotParameters) SetAdditionalMountPoints(v []AdditionalMountPoint)`

SetAdditionalMountPoints sets AdditionalMountPoints field to given value.

### HasAdditionalMountPoints

`func (o *ProvisionVDBBySnapshotParameters) HasAdditionalMountPoints() bool`

HasAdditionalMountPoints returns a boolean if a field has been set.

### SetAdditionalMountPointsNil

`func (o *ProvisionVDBBySnapshotParameters) SetAdditionalMountPointsNil(b bool)`

 SetAdditionalMountPointsNil sets the value for AdditionalMountPoints to be an explicit nil

### UnsetAdditionalMountPoints
`func (o *ProvisionVDBBySnapshotParameters) UnsetAdditionalMountPoints()`

UnsetAdditionalMountPoints ensures that no value is present for AdditionalMountPoints, not even an explicit nil
### GetAppdataConfigParams

`func (o *ProvisionVDBBySnapshotParameters) GetAppdataConfigParams() map[string]interface{}`

GetAppdataConfigParams returns the AppdataConfigParams field if non-nil, zero value otherwise.

### GetAppdataConfigParamsOk

`func (o *ProvisionVDBBySnapshotParameters) GetAppdataConfigParamsOk() (*map[string]interface{}, bool)`

GetAppdataConfigParamsOk returns a tuple with the AppdataConfigParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppdataConfigParams

`func (o *ProvisionVDBBySnapshotParameters) SetAppdataConfigParams(v map[string]interface{})`

SetAppdataConfigParams sets AppdataConfigParams field to given value.

### HasAppdataConfigParams

`func (o *ProvisionVDBBySnapshotParameters) HasAppdataConfigParams() bool`

HasAppdataConfigParams returns a boolean if a field has been set.

### SetAppdataConfigParamsNil

`func (o *ProvisionVDBBySnapshotParameters) SetAppdataConfigParamsNil(b bool)`

 SetAppdataConfigParamsNil sets the value for AppdataConfigParams to be an explicit nil

### UnsetAppdataConfigParams
`func (o *ProvisionVDBBySnapshotParameters) UnsetAppdataConfigParams()`

UnsetAppdataConfigParams ensures that no value is present for AppdataConfigParams, not even an explicit nil
### GetConfigParams

`func (o *ProvisionVDBBySnapshotParameters) GetConfigParams() map[string]interface{}`

GetConfigParams returns the ConfigParams field if non-nil, zero value otherwise.

### GetConfigParamsOk

`func (o *ProvisionVDBBySnapshotParameters) GetConfigParamsOk() (*map[string]interface{}, bool)`

GetConfigParamsOk returns a tuple with the ConfigParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParams

`func (o *ProvisionVDBBySnapshotParameters) SetConfigParams(v map[string]interface{})`

SetConfigParams sets ConfigParams field to given value.

### HasConfigParams

`func (o *ProvisionVDBBySnapshotParameters) HasConfigParams() bool`

HasConfigParams returns a boolean if a field has been set.

### SetConfigParamsNil

`func (o *ProvisionVDBBySnapshotParameters) SetConfigParamsNil(b bool)`

 SetConfigParamsNil sets the value for ConfigParams to be an explicit nil

### UnsetConfigParams
`func (o *ProvisionVDBBySnapshotParameters) UnsetConfigParams()`

UnsetConfigParams ensures that no value is present for ConfigParams, not even an explicit nil
### GetTags

`func (o *ProvisionVDBBySnapshotParameters) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProvisionVDBBySnapshotParameters) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProvisionVDBBySnapshotParameters) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProvisionVDBBySnapshotParameters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSnapshotId

`func (o *ProvisionVDBBySnapshotParameters) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ProvisionVDBBySnapshotParameters) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ProvisionVDBBySnapshotParameters) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ProvisionVDBBySnapshotParameters) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetEngineId

`func (o *ProvisionVDBBySnapshotParameters) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *ProvisionVDBBySnapshotParameters) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *ProvisionVDBBySnapshotParameters) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *ProvisionVDBBySnapshotParameters) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetSourceDataId

`func (o *ProvisionVDBBySnapshotParameters) GetSourceDataId() string`

GetSourceDataId returns the SourceDataId field if non-nil, zero value otherwise.

### GetSourceDataIdOk

`func (o *ProvisionVDBBySnapshotParameters) GetSourceDataIdOk() (*string, bool)`

GetSourceDataIdOk returns a tuple with the SourceDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataId

`func (o *ProvisionVDBBySnapshotParameters) SetSourceDataId(v string)`

SetSourceDataId sets SourceDataId field to given value.

### HasSourceDataId

`func (o *ProvisionVDBBySnapshotParameters) HasSourceDataId() bool`

HasSourceDataId returns a boolean if a field has been set.

### GetMakeCurrentAccountOwner

`func (o *ProvisionVDBBySnapshotParameters) GetMakeCurrentAccountOwner() bool`

GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field if non-nil, zero value otherwise.

### GetMakeCurrentAccountOwnerOk

`func (o *ProvisionVDBBySnapshotParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool)`

GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCurrentAccountOwner

`func (o *ProvisionVDBBySnapshotParameters) SetMakeCurrentAccountOwner(v bool)`

SetMakeCurrentAccountOwner sets MakeCurrentAccountOwner field to given value.

### HasMakeCurrentAccountOwner

`func (o *ProvisionVDBBySnapshotParameters) HasMakeCurrentAccountOwner() bool`

HasMakeCurrentAccountOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


