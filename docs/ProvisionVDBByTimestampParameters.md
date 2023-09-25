# ProvisionVDBByTimestampParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreRefresh** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment before refreshing the VDB. | [optional] 
**PostRefresh** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment after refreshing the VDB. | [optional] 
**PreSelfRefresh** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment before refreshing the VDB with data from itself. | [optional] 
**PostSelfRefresh** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment after refreshing the VDB with data from itself. | [optional] 
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
**ClusterNodeInstances** | Pointer to [**[]ClusterNodeInstance**](ClusterNodeInstance.md) | The cluster node instances details for this provision operation(Oracle RAC Only).This property is mutually exclusive with cluster_node_ids. | [optional] 
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
**PrivilegedOsUser** | Pointer to **string** | This privileged unix username will be used to create the VDB. Leave this field blank if you do not want to use privilege elevation. The unix privileged username should begin with a letter or an underscore, followed by letters, digits, underscores, or dashes. They can end with a dollar sign (postgres only). | [optional] 
**PostgresPort** | Pointer to **int32** | Port number for Postgres target database (postgres only). | [optional] 
**ConfigSettingsStg** | Pointer to [**[]ConfigSettingsStg**](ConfigSettingsStg.md) | Custom Database-Level config settings (postgres only). | [optional] 
**VcdbRestart** | Pointer to **bool** | Indicates whether the Engine should automatically restart this vCDB when target host reboot is detected. If vdb_restart property value is not explicitly set and vcdb_restart is set as false - the vdb_restart property is defaulted to false. | [optional] 
**MssqlFailoverDriveLetter** | Pointer to **string** | Base drive letter location for mount points. (MSSql Only). | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | The tags to be created for VDB. | [optional] 
**Timestamp** | Pointer to **time.Time** | The point in time from which to execute the operation. Mutually exclusive with timestamp_in_database_timezone. If the timestamp is not set, selects the latest point. | [optional] 
**TimestampInDatabaseTimezone** | Pointer to **string** | The point in time from which to execute the operation, expressed as a date-time in the timezone of the source database. Mutually exclusive with timestamp. | [optional] 
**TimeflowId** | Pointer to **string** | The Timeflow ID. | [optional] 
**EngineId** | Pointer to **string** | The ID of the Engine onto which to provision. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored. | [optional] 
**SourceDataId** | **string** | The ID of the source object (dSource or VDB) to provision from. All other objects referenced by the parameters must live on the same engine as the source. | 
**MakeCurrentAccountOwner** | Pointer to **bool** | Whether the account provisioning this VDB must be configured as owner of the VDB. | [optional] [default to true]

## Methods

### NewProvisionVDBByTimestampParameters

`func NewProvisionVDBByTimestampParameters(sourceDataId string, ) *ProvisionVDBByTimestampParameters`

NewProvisionVDBByTimestampParameters instantiates a new ProvisionVDBByTimestampParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionVDBByTimestampParametersWithDefaults

`func NewProvisionVDBByTimestampParametersWithDefaults() *ProvisionVDBByTimestampParameters`

NewProvisionVDBByTimestampParametersWithDefaults instantiates a new ProvisionVDBByTimestampParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreRefresh

`func (o *ProvisionVDBByTimestampParameters) GetPreRefresh() []Hook`

GetPreRefresh returns the PreRefresh field if non-nil, zero value otherwise.

### GetPreRefreshOk

`func (o *ProvisionVDBByTimestampParameters) GetPreRefreshOk() (*[]Hook, bool)`

GetPreRefreshOk returns a tuple with the PreRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRefresh

`func (o *ProvisionVDBByTimestampParameters) SetPreRefresh(v []Hook)`

SetPreRefresh sets PreRefresh field to given value.

### HasPreRefresh

`func (o *ProvisionVDBByTimestampParameters) HasPreRefresh() bool`

HasPreRefresh returns a boolean if a field has been set.

### GetPostRefresh

`func (o *ProvisionVDBByTimestampParameters) GetPostRefresh() []Hook`

GetPostRefresh returns the PostRefresh field if non-nil, zero value otherwise.

### GetPostRefreshOk

`func (o *ProvisionVDBByTimestampParameters) GetPostRefreshOk() (*[]Hook, bool)`

GetPostRefreshOk returns a tuple with the PostRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostRefresh

`func (o *ProvisionVDBByTimestampParameters) SetPostRefresh(v []Hook)`

SetPostRefresh sets PostRefresh field to given value.

### HasPostRefresh

`func (o *ProvisionVDBByTimestampParameters) HasPostRefresh() bool`

HasPostRefresh returns a boolean if a field has been set.

### GetPreSelfRefresh

`func (o *ProvisionVDBByTimestampParameters) GetPreSelfRefresh() []Hook`

GetPreSelfRefresh returns the PreSelfRefresh field if non-nil, zero value otherwise.

### GetPreSelfRefreshOk

`func (o *ProvisionVDBByTimestampParameters) GetPreSelfRefreshOk() (*[]Hook, bool)`

GetPreSelfRefreshOk returns a tuple with the PreSelfRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSelfRefresh

`func (o *ProvisionVDBByTimestampParameters) SetPreSelfRefresh(v []Hook)`

SetPreSelfRefresh sets PreSelfRefresh field to given value.

### HasPreSelfRefresh

`func (o *ProvisionVDBByTimestampParameters) HasPreSelfRefresh() bool`

HasPreSelfRefresh returns a boolean if a field has been set.

### GetPostSelfRefresh

`func (o *ProvisionVDBByTimestampParameters) GetPostSelfRefresh() []Hook`

GetPostSelfRefresh returns the PostSelfRefresh field if non-nil, zero value otherwise.

### GetPostSelfRefreshOk

`func (o *ProvisionVDBByTimestampParameters) GetPostSelfRefreshOk() (*[]Hook, bool)`

GetPostSelfRefreshOk returns a tuple with the PostSelfRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostSelfRefresh

`func (o *ProvisionVDBByTimestampParameters) SetPostSelfRefresh(v []Hook)`

SetPostSelfRefresh sets PostSelfRefresh field to given value.

### HasPostSelfRefresh

`func (o *ProvisionVDBByTimestampParameters) HasPostSelfRefresh() bool`

HasPostSelfRefresh returns a boolean if a field has been set.

### GetPreRollback

`func (o *ProvisionVDBByTimestampParameters) GetPreRollback() []Hook`

GetPreRollback returns the PreRollback field if non-nil, zero value otherwise.

### GetPreRollbackOk

`func (o *ProvisionVDBByTimestampParameters) GetPreRollbackOk() (*[]Hook, bool)`

GetPreRollbackOk returns a tuple with the PreRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRollback

`func (o *ProvisionVDBByTimestampParameters) SetPreRollback(v []Hook)`

SetPreRollback sets PreRollback field to given value.

### HasPreRollback

`func (o *ProvisionVDBByTimestampParameters) HasPreRollback() bool`

HasPreRollback returns a boolean if a field has been set.

### GetPostRollback

`func (o *ProvisionVDBByTimestampParameters) GetPostRollback() []Hook`

GetPostRollback returns the PostRollback field if non-nil, zero value otherwise.

### GetPostRollbackOk

`func (o *ProvisionVDBByTimestampParameters) GetPostRollbackOk() (*[]Hook, bool)`

GetPostRollbackOk returns a tuple with the PostRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostRollback

`func (o *ProvisionVDBByTimestampParameters) SetPostRollback(v []Hook)`

SetPostRollback sets PostRollback field to given value.

### HasPostRollback

`func (o *ProvisionVDBByTimestampParameters) HasPostRollback() bool`

HasPostRollback returns a boolean if a field has been set.

### GetConfigureClone

`func (o *ProvisionVDBByTimestampParameters) GetConfigureClone() []Hook`

GetConfigureClone returns the ConfigureClone field if non-nil, zero value otherwise.

### GetConfigureCloneOk

`func (o *ProvisionVDBByTimestampParameters) GetConfigureCloneOk() (*[]Hook, bool)`

GetConfigureCloneOk returns a tuple with the ConfigureClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureClone

`func (o *ProvisionVDBByTimestampParameters) SetConfigureClone(v []Hook)`

SetConfigureClone sets ConfigureClone field to given value.

### HasConfigureClone

`func (o *ProvisionVDBByTimestampParameters) HasConfigureClone() bool`

HasConfigureClone returns a boolean if a field has been set.

### GetPreSnapshot

`func (o *ProvisionVDBByTimestampParameters) GetPreSnapshot() []Hook`

GetPreSnapshot returns the PreSnapshot field if non-nil, zero value otherwise.

### GetPreSnapshotOk

`func (o *ProvisionVDBByTimestampParameters) GetPreSnapshotOk() (*[]Hook, bool)`

GetPreSnapshotOk returns a tuple with the PreSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSnapshot

`func (o *ProvisionVDBByTimestampParameters) SetPreSnapshot(v []Hook)`

SetPreSnapshot sets PreSnapshot field to given value.

### HasPreSnapshot

`func (o *ProvisionVDBByTimestampParameters) HasPreSnapshot() bool`

HasPreSnapshot returns a boolean if a field has been set.

### GetPostSnapshot

`func (o *ProvisionVDBByTimestampParameters) GetPostSnapshot() []Hook`

GetPostSnapshot returns the PostSnapshot field if non-nil, zero value otherwise.

### GetPostSnapshotOk

`func (o *ProvisionVDBByTimestampParameters) GetPostSnapshotOk() (*[]Hook, bool)`

GetPostSnapshotOk returns a tuple with the PostSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostSnapshot

`func (o *ProvisionVDBByTimestampParameters) SetPostSnapshot(v []Hook)`

SetPostSnapshot sets PostSnapshot field to given value.

### HasPostSnapshot

`func (o *ProvisionVDBByTimestampParameters) HasPostSnapshot() bool`

HasPostSnapshot returns a boolean if a field has been set.

### GetPreStart

`func (o *ProvisionVDBByTimestampParameters) GetPreStart() []Hook`

GetPreStart returns the PreStart field if non-nil, zero value otherwise.

### GetPreStartOk

`func (o *ProvisionVDBByTimestampParameters) GetPreStartOk() (*[]Hook, bool)`

GetPreStartOk returns a tuple with the PreStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreStart

`func (o *ProvisionVDBByTimestampParameters) SetPreStart(v []Hook)`

SetPreStart sets PreStart field to given value.

### HasPreStart

`func (o *ProvisionVDBByTimestampParameters) HasPreStart() bool`

HasPreStart returns a boolean if a field has been set.

### GetPostStart

`func (o *ProvisionVDBByTimestampParameters) GetPostStart() []Hook`

GetPostStart returns the PostStart field if non-nil, zero value otherwise.

### GetPostStartOk

`func (o *ProvisionVDBByTimestampParameters) GetPostStartOk() (*[]Hook, bool)`

GetPostStartOk returns a tuple with the PostStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostStart

`func (o *ProvisionVDBByTimestampParameters) SetPostStart(v []Hook)`

SetPostStart sets PostStart field to given value.

### HasPostStart

`func (o *ProvisionVDBByTimestampParameters) HasPostStart() bool`

HasPostStart returns a boolean if a field has been set.

### GetPreStop

`func (o *ProvisionVDBByTimestampParameters) GetPreStop() []Hook`

GetPreStop returns the PreStop field if non-nil, zero value otherwise.

### GetPreStopOk

`func (o *ProvisionVDBByTimestampParameters) GetPreStopOk() (*[]Hook, bool)`

GetPreStopOk returns a tuple with the PreStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreStop

`func (o *ProvisionVDBByTimestampParameters) SetPreStop(v []Hook)`

SetPreStop sets PreStop field to given value.

### HasPreStop

`func (o *ProvisionVDBByTimestampParameters) HasPreStop() bool`

HasPreStop returns a boolean if a field has been set.

### GetPostStop

`func (o *ProvisionVDBByTimestampParameters) GetPostStop() []Hook`

GetPostStop returns the PostStop field if non-nil, zero value otherwise.

### GetPostStopOk

`func (o *ProvisionVDBByTimestampParameters) GetPostStopOk() (*[]Hook, bool)`

GetPostStopOk returns a tuple with the PostStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostStop

`func (o *ProvisionVDBByTimestampParameters) SetPostStop(v []Hook)`

SetPostStop sets PostStop field to given value.

### HasPostStop

`func (o *ProvisionVDBByTimestampParameters) HasPostStop() bool`

HasPostStop returns a boolean if a field has been set.

### GetTargetGroupId

`func (o *ProvisionVDBByTimestampParameters) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *ProvisionVDBByTimestampParameters) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *ProvisionVDBByTimestampParameters) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *ProvisionVDBByTimestampParameters) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### GetName

`func (o *ProvisionVDBByTimestampParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisionVDBByTimestampParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisionVDBByTimestampParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProvisionVDBByTimestampParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDatabaseName

`func (o *ProvisionVDBByTimestampParameters) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ProvisionVDBByTimestampParameters) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ProvisionVDBByTimestampParameters) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *ProvisionVDBByTimestampParameters) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetCdbId

`func (o *ProvisionVDBByTimestampParameters) GetCdbId() string`

GetCdbId returns the CdbId field if non-nil, zero value otherwise.

### GetCdbIdOk

`func (o *ProvisionVDBByTimestampParameters) GetCdbIdOk() (*string, bool)`

GetCdbIdOk returns a tuple with the CdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdbId

`func (o *ProvisionVDBByTimestampParameters) SetCdbId(v string)`

SetCdbId sets CdbId field to given value.

### HasCdbId

`func (o *ProvisionVDBByTimestampParameters) HasCdbId() bool`

HasCdbId returns a boolean if a field has been set.

### GetClusterNodeIds

`func (o *ProvisionVDBByTimestampParameters) GetClusterNodeIds() []string`

GetClusterNodeIds returns the ClusterNodeIds field if non-nil, zero value otherwise.

### GetClusterNodeIdsOk

`func (o *ProvisionVDBByTimestampParameters) GetClusterNodeIdsOk() (*[]string, bool)`

GetClusterNodeIdsOk returns a tuple with the ClusterNodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeIds

`func (o *ProvisionVDBByTimestampParameters) SetClusterNodeIds(v []string)`

SetClusterNodeIds sets ClusterNodeIds field to given value.

### HasClusterNodeIds

`func (o *ProvisionVDBByTimestampParameters) HasClusterNodeIds() bool`

HasClusterNodeIds returns a boolean if a field has been set.

### GetClusterNodeInstances

`func (o *ProvisionVDBByTimestampParameters) GetClusterNodeInstances() []ClusterNodeInstance`

GetClusterNodeInstances returns the ClusterNodeInstances field if non-nil, zero value otherwise.

### GetClusterNodeInstancesOk

`func (o *ProvisionVDBByTimestampParameters) GetClusterNodeInstancesOk() (*[]ClusterNodeInstance, bool)`

GetClusterNodeInstancesOk returns a tuple with the ClusterNodeInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeInstances

`func (o *ProvisionVDBByTimestampParameters) SetClusterNodeInstances(v []ClusterNodeInstance)`

SetClusterNodeInstances sets ClusterNodeInstances field to given value.

### HasClusterNodeInstances

`func (o *ProvisionVDBByTimestampParameters) HasClusterNodeInstances() bool`

HasClusterNodeInstances returns a boolean if a field has been set.

### GetTruncateLogOnCheckpoint

`func (o *ProvisionVDBByTimestampParameters) GetTruncateLogOnCheckpoint() bool`

GetTruncateLogOnCheckpoint returns the TruncateLogOnCheckpoint field if non-nil, zero value otherwise.

### GetTruncateLogOnCheckpointOk

`func (o *ProvisionVDBByTimestampParameters) GetTruncateLogOnCheckpointOk() (*bool, bool)`

GetTruncateLogOnCheckpointOk returns a tuple with the TruncateLogOnCheckpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateLogOnCheckpoint

`func (o *ProvisionVDBByTimestampParameters) SetTruncateLogOnCheckpoint(v bool)`

SetTruncateLogOnCheckpoint sets TruncateLogOnCheckpoint field to given value.

### HasTruncateLogOnCheckpoint

`func (o *ProvisionVDBByTimestampParameters) HasTruncateLogOnCheckpoint() bool`

HasTruncateLogOnCheckpoint returns a boolean if a field has been set.

### GetOsUsername

`func (o *ProvisionVDBByTimestampParameters) GetOsUsername() string`

GetOsUsername returns the OsUsername field if non-nil, zero value otherwise.

### GetOsUsernameOk

`func (o *ProvisionVDBByTimestampParameters) GetOsUsernameOk() (*string, bool)`

GetOsUsernameOk returns a tuple with the OsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsUsername

`func (o *ProvisionVDBByTimestampParameters) SetOsUsername(v string)`

SetOsUsername sets OsUsername field to given value.

### HasOsUsername

`func (o *ProvisionVDBByTimestampParameters) HasOsUsername() bool`

HasOsUsername returns a boolean if a field has been set.

### GetOsPassword

`func (o *ProvisionVDBByTimestampParameters) GetOsPassword() string`

GetOsPassword returns the OsPassword field if non-nil, zero value otherwise.

### GetOsPasswordOk

`func (o *ProvisionVDBByTimestampParameters) GetOsPasswordOk() (*string, bool)`

GetOsPasswordOk returns a tuple with the OsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPassword

`func (o *ProvisionVDBByTimestampParameters) SetOsPassword(v string)`

SetOsPassword sets OsPassword field to given value.

### HasOsPassword

`func (o *ProvisionVDBByTimestampParameters) HasOsPassword() bool`

HasOsPassword returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *ProvisionVDBByTimestampParameters) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ProvisionVDBByTimestampParameters) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ProvisionVDBByTimestampParameters) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ProvisionVDBByTimestampParameters) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEnvironmentUserId

`func (o *ProvisionVDBByTimestampParameters) GetEnvironmentUserId() string`

GetEnvironmentUserId returns the EnvironmentUserId field if non-nil, zero value otherwise.

### GetEnvironmentUserIdOk

`func (o *ProvisionVDBByTimestampParameters) GetEnvironmentUserIdOk() (*string, bool)`

GetEnvironmentUserIdOk returns a tuple with the EnvironmentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUserId

`func (o *ProvisionVDBByTimestampParameters) SetEnvironmentUserId(v string)`

SetEnvironmentUserId sets EnvironmentUserId field to given value.

### HasEnvironmentUserId

`func (o *ProvisionVDBByTimestampParameters) HasEnvironmentUserId() bool`

HasEnvironmentUserId returns a boolean if a field has been set.

### GetRepositoryId

`func (o *ProvisionVDBByTimestampParameters) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *ProvisionVDBByTimestampParameters) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *ProvisionVDBByTimestampParameters) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *ProvisionVDBByTimestampParameters) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetAutoSelectRepository

`func (o *ProvisionVDBByTimestampParameters) GetAutoSelectRepository() bool`

GetAutoSelectRepository returns the AutoSelectRepository field if non-nil, zero value otherwise.

### GetAutoSelectRepositoryOk

`func (o *ProvisionVDBByTimestampParameters) GetAutoSelectRepositoryOk() (*bool, bool)`

GetAutoSelectRepositoryOk returns a tuple with the AutoSelectRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelectRepository

`func (o *ProvisionVDBByTimestampParameters) SetAutoSelectRepository(v bool)`

SetAutoSelectRepository sets AutoSelectRepository field to given value.

### HasAutoSelectRepository

`func (o *ProvisionVDBByTimestampParameters) HasAutoSelectRepository() bool`

HasAutoSelectRepository returns a boolean if a field has been set.

### GetVdbRestart

`func (o *ProvisionVDBByTimestampParameters) GetVdbRestart() bool`

GetVdbRestart returns the VdbRestart field if non-nil, zero value otherwise.

### GetVdbRestartOk

`func (o *ProvisionVDBByTimestampParameters) GetVdbRestartOk() (*bool, bool)`

GetVdbRestartOk returns a tuple with the VdbRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdbRestart

`func (o *ProvisionVDBByTimestampParameters) SetVdbRestart(v bool)`

SetVdbRestart sets VdbRestart field to given value.

### HasVdbRestart

`func (o *ProvisionVDBByTimestampParameters) HasVdbRestart() bool`

HasVdbRestart returns a boolean if a field has been set.

### GetTemplateId

`func (o *ProvisionVDBByTimestampParameters) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ProvisionVDBByTimestampParameters) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ProvisionVDBByTimestampParameters) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ProvisionVDBByTimestampParameters) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetAuxiliaryTemplateId

`func (o *ProvisionVDBByTimestampParameters) GetAuxiliaryTemplateId() string`

GetAuxiliaryTemplateId returns the AuxiliaryTemplateId field if non-nil, zero value otherwise.

### GetAuxiliaryTemplateIdOk

`func (o *ProvisionVDBByTimestampParameters) GetAuxiliaryTemplateIdOk() (*string, bool)`

GetAuxiliaryTemplateIdOk returns a tuple with the AuxiliaryTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryTemplateId

`func (o *ProvisionVDBByTimestampParameters) SetAuxiliaryTemplateId(v string)`

SetAuxiliaryTemplateId sets AuxiliaryTemplateId field to given value.

### HasAuxiliaryTemplateId

`func (o *ProvisionVDBByTimestampParameters) HasAuxiliaryTemplateId() bool`

HasAuxiliaryTemplateId returns a boolean if a field has been set.

### GetFileMappingRules

`func (o *ProvisionVDBByTimestampParameters) GetFileMappingRules() string`

GetFileMappingRules returns the FileMappingRules field if non-nil, zero value otherwise.

### GetFileMappingRulesOk

`func (o *ProvisionVDBByTimestampParameters) GetFileMappingRulesOk() (*string, bool)`

GetFileMappingRulesOk returns a tuple with the FileMappingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMappingRules

`func (o *ProvisionVDBByTimestampParameters) SetFileMappingRules(v string)`

SetFileMappingRules sets FileMappingRules field to given value.

### HasFileMappingRules

`func (o *ProvisionVDBByTimestampParameters) HasFileMappingRules() bool`

HasFileMappingRules returns a boolean if a field has been set.

### GetOracleInstanceName

`func (o *ProvisionVDBByTimestampParameters) GetOracleInstanceName() string`

GetOracleInstanceName returns the OracleInstanceName field if non-nil, zero value otherwise.

### GetOracleInstanceNameOk

`func (o *ProvisionVDBByTimestampParameters) GetOracleInstanceNameOk() (*string, bool)`

GetOracleInstanceNameOk returns a tuple with the OracleInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleInstanceName

`func (o *ProvisionVDBByTimestampParameters) SetOracleInstanceName(v string)`

SetOracleInstanceName sets OracleInstanceName field to given value.

### HasOracleInstanceName

`func (o *ProvisionVDBByTimestampParameters) HasOracleInstanceName() bool`

HasOracleInstanceName returns a boolean if a field has been set.

### GetUniqueName

`func (o *ProvisionVDBByTimestampParameters) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *ProvisionVDBByTimestampParameters) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *ProvisionVDBByTimestampParameters) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *ProvisionVDBByTimestampParameters) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetVcdbName

`func (o *ProvisionVDBByTimestampParameters) GetVcdbName() string`

GetVcdbName returns the VcdbName field if non-nil, zero value otherwise.

### GetVcdbNameOk

`func (o *ProvisionVDBByTimestampParameters) GetVcdbNameOk() (*string, bool)`

GetVcdbNameOk returns a tuple with the VcdbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdbName

`func (o *ProvisionVDBByTimestampParameters) SetVcdbName(v string)`

SetVcdbName sets VcdbName field to given value.

### HasVcdbName

`func (o *ProvisionVDBByTimestampParameters) HasVcdbName() bool`

HasVcdbName returns a boolean if a field has been set.

### GetVcdbDatabaseName

`func (o *ProvisionVDBByTimestampParameters) GetVcdbDatabaseName() string`

GetVcdbDatabaseName returns the VcdbDatabaseName field if non-nil, zero value otherwise.

### GetVcdbDatabaseNameOk

`func (o *ProvisionVDBByTimestampParameters) GetVcdbDatabaseNameOk() (*string, bool)`

GetVcdbDatabaseNameOk returns a tuple with the VcdbDatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdbDatabaseName

`func (o *ProvisionVDBByTimestampParameters) SetVcdbDatabaseName(v string)`

SetVcdbDatabaseName sets VcdbDatabaseName field to given value.

### HasVcdbDatabaseName

`func (o *ProvisionVDBByTimestampParameters) HasVcdbDatabaseName() bool`

HasVcdbDatabaseName returns a boolean if a field has been set.

### GetMountPoint

`func (o *ProvisionVDBByTimestampParameters) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *ProvisionVDBByTimestampParameters) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *ProvisionVDBByTimestampParameters) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *ProvisionVDBByTimestampParameters) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### GetOpenResetLogs

`func (o *ProvisionVDBByTimestampParameters) GetOpenResetLogs() bool`

GetOpenResetLogs returns the OpenResetLogs field if non-nil, zero value otherwise.

### GetOpenResetLogsOk

`func (o *ProvisionVDBByTimestampParameters) GetOpenResetLogsOk() (*bool, bool)`

GetOpenResetLogsOk returns a tuple with the OpenResetLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenResetLogs

`func (o *ProvisionVDBByTimestampParameters) SetOpenResetLogs(v bool)`

SetOpenResetLogs sets OpenResetLogs field to given value.

### HasOpenResetLogs

`func (o *ProvisionVDBByTimestampParameters) HasOpenResetLogs() bool`

HasOpenResetLogs returns a boolean if a field has been set.

### GetSnapshotPolicyId

`func (o *ProvisionVDBByTimestampParameters) GetSnapshotPolicyId() string`

GetSnapshotPolicyId returns the SnapshotPolicyId field if non-nil, zero value otherwise.

### GetSnapshotPolicyIdOk

`func (o *ProvisionVDBByTimestampParameters) GetSnapshotPolicyIdOk() (*string, bool)`

GetSnapshotPolicyIdOk returns a tuple with the SnapshotPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotPolicyId

`func (o *ProvisionVDBByTimestampParameters) SetSnapshotPolicyId(v string)`

SetSnapshotPolicyId sets SnapshotPolicyId field to given value.

### HasSnapshotPolicyId

`func (o *ProvisionVDBByTimestampParameters) HasSnapshotPolicyId() bool`

HasSnapshotPolicyId returns a boolean if a field has been set.

### GetRetentionPolicyId

`func (o *ProvisionVDBByTimestampParameters) GetRetentionPolicyId() string`

GetRetentionPolicyId returns the RetentionPolicyId field if non-nil, zero value otherwise.

### GetRetentionPolicyIdOk

`func (o *ProvisionVDBByTimestampParameters) GetRetentionPolicyIdOk() (*string, bool)`

GetRetentionPolicyIdOk returns a tuple with the RetentionPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicyId

`func (o *ProvisionVDBByTimestampParameters) SetRetentionPolicyId(v string)`

SetRetentionPolicyId sets RetentionPolicyId field to given value.

### HasRetentionPolicyId

`func (o *ProvisionVDBByTimestampParameters) HasRetentionPolicyId() bool`

HasRetentionPolicyId returns a boolean if a field has been set.

### GetRecoveryModel

`func (o *ProvisionVDBByTimestampParameters) GetRecoveryModel() string`

GetRecoveryModel returns the RecoveryModel field if non-nil, zero value otherwise.

### GetRecoveryModelOk

`func (o *ProvisionVDBByTimestampParameters) GetRecoveryModelOk() (*string, bool)`

GetRecoveryModelOk returns a tuple with the RecoveryModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryModel

`func (o *ProvisionVDBByTimestampParameters) SetRecoveryModel(v string)`

SetRecoveryModel sets RecoveryModel field to given value.

### HasRecoveryModel

`func (o *ProvisionVDBByTimestampParameters) HasRecoveryModel() bool`

HasRecoveryModel returns a boolean if a field has been set.

### GetPreScript

`func (o *ProvisionVDBByTimestampParameters) GetPreScript() string`

GetPreScript returns the PreScript field if non-nil, zero value otherwise.

### GetPreScriptOk

`func (o *ProvisionVDBByTimestampParameters) GetPreScriptOk() (*string, bool)`

GetPreScriptOk returns a tuple with the PreScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreScript

`func (o *ProvisionVDBByTimestampParameters) SetPreScript(v string)`

SetPreScript sets PreScript field to given value.

### HasPreScript

`func (o *ProvisionVDBByTimestampParameters) HasPreScript() bool`

HasPreScript returns a boolean if a field has been set.

### GetPostScript

`func (o *ProvisionVDBByTimestampParameters) GetPostScript() string`

GetPostScript returns the PostScript field if non-nil, zero value otherwise.

### GetPostScriptOk

`func (o *ProvisionVDBByTimestampParameters) GetPostScriptOk() (*string, bool)`

GetPostScriptOk returns a tuple with the PostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScript

`func (o *ProvisionVDBByTimestampParameters) SetPostScript(v string)`

SetPostScript sets PostScript field to given value.

### HasPostScript

`func (o *ProvisionVDBByTimestampParameters) HasPostScript() bool`

HasPostScript returns a boolean if a field has been set.

### GetCdcOnProvision

`func (o *ProvisionVDBByTimestampParameters) GetCdcOnProvision() bool`

GetCdcOnProvision returns the CdcOnProvision field if non-nil, zero value otherwise.

### GetCdcOnProvisionOk

`func (o *ProvisionVDBByTimestampParameters) GetCdcOnProvisionOk() (*bool, bool)`

GetCdcOnProvisionOk returns a tuple with the CdcOnProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdcOnProvision

`func (o *ProvisionVDBByTimestampParameters) SetCdcOnProvision(v bool)`

SetCdcOnProvision sets CdcOnProvision field to given value.

### HasCdcOnProvision

`func (o *ProvisionVDBByTimestampParameters) HasCdcOnProvision() bool`

HasCdcOnProvision returns a boolean if a field has been set.

### GetOnlineLogSize

`func (o *ProvisionVDBByTimestampParameters) GetOnlineLogSize() int32`

GetOnlineLogSize returns the OnlineLogSize field if non-nil, zero value otherwise.

### GetOnlineLogSizeOk

`func (o *ProvisionVDBByTimestampParameters) GetOnlineLogSizeOk() (*int32, bool)`

GetOnlineLogSizeOk returns a tuple with the OnlineLogSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineLogSize

`func (o *ProvisionVDBByTimestampParameters) SetOnlineLogSize(v int32)`

SetOnlineLogSize sets OnlineLogSize field to given value.

### HasOnlineLogSize

`func (o *ProvisionVDBByTimestampParameters) HasOnlineLogSize() bool`

HasOnlineLogSize returns a boolean if a field has been set.

### GetOnlineLogGroups

`func (o *ProvisionVDBByTimestampParameters) GetOnlineLogGroups() int32`

GetOnlineLogGroups returns the OnlineLogGroups field if non-nil, zero value otherwise.

### GetOnlineLogGroupsOk

`func (o *ProvisionVDBByTimestampParameters) GetOnlineLogGroupsOk() (*int32, bool)`

GetOnlineLogGroupsOk returns a tuple with the OnlineLogGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineLogGroups

`func (o *ProvisionVDBByTimestampParameters) SetOnlineLogGroups(v int32)`

SetOnlineLogGroups sets OnlineLogGroups field to given value.

### HasOnlineLogGroups

`func (o *ProvisionVDBByTimestampParameters) HasOnlineLogGroups() bool`

HasOnlineLogGroups returns a boolean if a field has been set.

### GetArchiveLog

`func (o *ProvisionVDBByTimestampParameters) GetArchiveLog() bool`

GetArchiveLog returns the ArchiveLog field if non-nil, zero value otherwise.

### GetArchiveLogOk

`func (o *ProvisionVDBByTimestampParameters) GetArchiveLogOk() (*bool, bool)`

GetArchiveLogOk returns a tuple with the ArchiveLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLog

`func (o *ProvisionVDBByTimestampParameters) SetArchiveLog(v bool)`

SetArchiveLog sets ArchiveLog field to given value.

### HasArchiveLog

`func (o *ProvisionVDBByTimestampParameters) HasArchiveLog() bool`

HasArchiveLog returns a boolean if a field has been set.

### GetNewDbid

`func (o *ProvisionVDBByTimestampParameters) GetNewDbid() bool`

GetNewDbid returns the NewDbid field if non-nil, zero value otherwise.

### GetNewDbidOk

`func (o *ProvisionVDBByTimestampParameters) GetNewDbidOk() (*bool, bool)`

GetNewDbidOk returns a tuple with the NewDbid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDbid

`func (o *ProvisionVDBByTimestampParameters) SetNewDbid(v bool)`

SetNewDbid sets NewDbid field to given value.

### HasNewDbid

`func (o *ProvisionVDBByTimestampParameters) HasNewDbid() bool`

HasNewDbid returns a boolean if a field has been set.

### GetListenerIds

`func (o *ProvisionVDBByTimestampParameters) GetListenerIds() []string`

GetListenerIds returns the ListenerIds field if non-nil, zero value otherwise.

### GetListenerIdsOk

`func (o *ProvisionVDBByTimestampParameters) GetListenerIdsOk() (*[]string, bool)`

GetListenerIdsOk returns a tuple with the ListenerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerIds

`func (o *ProvisionVDBByTimestampParameters) SetListenerIds(v []string)`

SetListenerIds sets ListenerIds field to given value.

### HasListenerIds

`func (o *ProvisionVDBByTimestampParameters) HasListenerIds() bool`

HasListenerIds returns a boolean if a field has been set.

### GetCustomEnvVars

`func (o *ProvisionVDBByTimestampParameters) GetCustomEnvVars() map[string]string`

GetCustomEnvVars returns the CustomEnvVars field if non-nil, zero value otherwise.

### GetCustomEnvVarsOk

`func (o *ProvisionVDBByTimestampParameters) GetCustomEnvVarsOk() (*map[string]string, bool)`

GetCustomEnvVarsOk returns a tuple with the CustomEnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvVars

`func (o *ProvisionVDBByTimestampParameters) SetCustomEnvVars(v map[string]string)`

SetCustomEnvVars sets CustomEnvVars field to given value.

### HasCustomEnvVars

`func (o *ProvisionVDBByTimestampParameters) HasCustomEnvVars() bool`

HasCustomEnvVars returns a boolean if a field has been set.

### GetCustomEnvFiles

`func (o *ProvisionVDBByTimestampParameters) GetCustomEnvFiles() []string`

GetCustomEnvFiles returns the CustomEnvFiles field if non-nil, zero value otherwise.

### GetCustomEnvFilesOk

`func (o *ProvisionVDBByTimestampParameters) GetCustomEnvFilesOk() (*[]string, bool)`

GetCustomEnvFilesOk returns a tuple with the CustomEnvFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvFiles

`func (o *ProvisionVDBByTimestampParameters) SetCustomEnvFiles(v []string)`

SetCustomEnvFiles sets CustomEnvFiles field to given value.

### HasCustomEnvFiles

`func (o *ProvisionVDBByTimestampParameters) HasCustomEnvFiles() bool`

HasCustomEnvFiles returns a boolean if a field has been set.

### GetOracleRacCustomEnvFiles

`func (o *ProvisionVDBByTimestampParameters) GetOracleRacCustomEnvFiles() []OracleRacCustomEnvFile`

GetOracleRacCustomEnvFiles returns the OracleRacCustomEnvFiles field if non-nil, zero value otherwise.

### GetOracleRacCustomEnvFilesOk

`func (o *ProvisionVDBByTimestampParameters) GetOracleRacCustomEnvFilesOk() (*[]OracleRacCustomEnvFile, bool)`

GetOracleRacCustomEnvFilesOk returns a tuple with the OracleRacCustomEnvFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleRacCustomEnvFiles

`func (o *ProvisionVDBByTimestampParameters) SetOracleRacCustomEnvFiles(v []OracleRacCustomEnvFile)`

SetOracleRacCustomEnvFiles sets OracleRacCustomEnvFiles field to given value.

### HasOracleRacCustomEnvFiles

`func (o *ProvisionVDBByTimestampParameters) HasOracleRacCustomEnvFiles() bool`

HasOracleRacCustomEnvFiles returns a boolean if a field has been set.

### GetOracleRacCustomEnvVars

`func (o *ProvisionVDBByTimestampParameters) GetOracleRacCustomEnvVars() []OracleRacCustomEnvVar`

GetOracleRacCustomEnvVars returns the OracleRacCustomEnvVars field if non-nil, zero value otherwise.

### GetOracleRacCustomEnvVarsOk

`func (o *ProvisionVDBByTimestampParameters) GetOracleRacCustomEnvVarsOk() (*[]OracleRacCustomEnvVar, bool)`

GetOracleRacCustomEnvVarsOk returns a tuple with the OracleRacCustomEnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleRacCustomEnvVars

`func (o *ProvisionVDBByTimestampParameters) SetOracleRacCustomEnvVars(v []OracleRacCustomEnvVar)`

SetOracleRacCustomEnvVars sets OracleRacCustomEnvVars field to given value.

### HasOracleRacCustomEnvVars

`func (o *ProvisionVDBByTimestampParameters) HasOracleRacCustomEnvVars() bool`

HasOracleRacCustomEnvVars returns a boolean if a field has been set.

### GetParentTdeKeystorePath

`func (o *ProvisionVDBByTimestampParameters) GetParentTdeKeystorePath() string`

GetParentTdeKeystorePath returns the ParentTdeKeystorePath field if non-nil, zero value otherwise.

### GetParentTdeKeystorePathOk

`func (o *ProvisionVDBByTimestampParameters) GetParentTdeKeystorePathOk() (*string, bool)`

GetParentTdeKeystorePathOk returns a tuple with the ParentTdeKeystorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTdeKeystorePath

`func (o *ProvisionVDBByTimestampParameters) SetParentTdeKeystorePath(v string)`

SetParentTdeKeystorePath sets ParentTdeKeystorePath field to given value.

### HasParentTdeKeystorePath

`func (o *ProvisionVDBByTimestampParameters) HasParentTdeKeystorePath() bool`

HasParentTdeKeystorePath returns a boolean if a field has been set.

### GetParentTdeKeystorePassword

`func (o *ProvisionVDBByTimestampParameters) GetParentTdeKeystorePassword() string`

GetParentTdeKeystorePassword returns the ParentTdeKeystorePassword field if non-nil, zero value otherwise.

### GetParentTdeKeystorePasswordOk

`func (o *ProvisionVDBByTimestampParameters) GetParentTdeKeystorePasswordOk() (*string, bool)`

GetParentTdeKeystorePasswordOk returns a tuple with the ParentTdeKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTdeKeystorePassword

`func (o *ProvisionVDBByTimestampParameters) SetParentTdeKeystorePassword(v string)`

SetParentTdeKeystorePassword sets ParentTdeKeystorePassword field to given value.

### HasParentTdeKeystorePassword

`func (o *ProvisionVDBByTimestampParameters) HasParentTdeKeystorePassword() bool`

HasParentTdeKeystorePassword returns a boolean if a field has been set.

### GetTdeExportedKeyFileSecret

`func (o *ProvisionVDBByTimestampParameters) GetTdeExportedKeyFileSecret() string`

GetTdeExportedKeyFileSecret returns the TdeExportedKeyFileSecret field if non-nil, zero value otherwise.

### GetTdeExportedKeyFileSecretOk

`func (o *ProvisionVDBByTimestampParameters) GetTdeExportedKeyFileSecretOk() (*string, bool)`

GetTdeExportedKeyFileSecretOk returns a tuple with the TdeExportedKeyFileSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdeExportedKeyFileSecret

`func (o *ProvisionVDBByTimestampParameters) SetTdeExportedKeyFileSecret(v string)`

SetTdeExportedKeyFileSecret sets TdeExportedKeyFileSecret field to given value.

### HasTdeExportedKeyFileSecret

`func (o *ProvisionVDBByTimestampParameters) HasTdeExportedKeyFileSecret() bool`

HasTdeExportedKeyFileSecret returns a boolean if a field has been set.

### GetTdeKeyIdentifier

`func (o *ProvisionVDBByTimestampParameters) GetTdeKeyIdentifier() string`

GetTdeKeyIdentifier returns the TdeKeyIdentifier field if non-nil, zero value otherwise.

### GetTdeKeyIdentifierOk

`func (o *ProvisionVDBByTimestampParameters) GetTdeKeyIdentifierOk() (*string, bool)`

GetTdeKeyIdentifierOk returns a tuple with the TdeKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdeKeyIdentifier

`func (o *ProvisionVDBByTimestampParameters) SetTdeKeyIdentifier(v string)`

SetTdeKeyIdentifier sets TdeKeyIdentifier field to given value.

### HasTdeKeyIdentifier

`func (o *ProvisionVDBByTimestampParameters) HasTdeKeyIdentifier() bool`

HasTdeKeyIdentifier returns a boolean if a field has been set.

### GetTargetVcdbTdeKeystorePath

`func (o *ProvisionVDBByTimestampParameters) GetTargetVcdbTdeKeystorePath() string`

GetTargetVcdbTdeKeystorePath returns the TargetVcdbTdeKeystorePath field if non-nil, zero value otherwise.

### GetTargetVcdbTdeKeystorePathOk

`func (o *ProvisionVDBByTimestampParameters) GetTargetVcdbTdeKeystorePathOk() (*string, bool)`

GetTargetVcdbTdeKeystorePathOk returns a tuple with the TargetVcdbTdeKeystorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVcdbTdeKeystorePath

`func (o *ProvisionVDBByTimestampParameters) SetTargetVcdbTdeKeystorePath(v string)`

SetTargetVcdbTdeKeystorePath sets TargetVcdbTdeKeystorePath field to given value.

### HasTargetVcdbTdeKeystorePath

`func (o *ProvisionVDBByTimestampParameters) HasTargetVcdbTdeKeystorePath() bool`

HasTargetVcdbTdeKeystorePath returns a boolean if a field has been set.

### GetCdbTdeKeystorePassword

`func (o *ProvisionVDBByTimestampParameters) GetCdbTdeKeystorePassword() string`

GetCdbTdeKeystorePassword returns the CdbTdeKeystorePassword field if non-nil, zero value otherwise.

### GetCdbTdeKeystorePasswordOk

`func (o *ProvisionVDBByTimestampParameters) GetCdbTdeKeystorePasswordOk() (*string, bool)`

GetCdbTdeKeystorePasswordOk returns a tuple with the CdbTdeKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdbTdeKeystorePassword

`func (o *ProvisionVDBByTimestampParameters) SetCdbTdeKeystorePassword(v string)`

SetCdbTdeKeystorePassword sets CdbTdeKeystorePassword field to given value.

### HasCdbTdeKeystorePassword

`func (o *ProvisionVDBByTimestampParameters) HasCdbTdeKeystorePassword() bool`

HasCdbTdeKeystorePassword returns a boolean if a field has been set.

### GetVcdbTdeKeyIdentifier

`func (o *ProvisionVDBByTimestampParameters) GetVcdbTdeKeyIdentifier() string`

GetVcdbTdeKeyIdentifier returns the VcdbTdeKeyIdentifier field if non-nil, zero value otherwise.

### GetVcdbTdeKeyIdentifierOk

`func (o *ProvisionVDBByTimestampParameters) GetVcdbTdeKeyIdentifierOk() (*string, bool)`

GetVcdbTdeKeyIdentifierOk returns a tuple with the VcdbTdeKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdbTdeKeyIdentifier

`func (o *ProvisionVDBByTimestampParameters) SetVcdbTdeKeyIdentifier(v string)`

SetVcdbTdeKeyIdentifier sets VcdbTdeKeyIdentifier field to given value.

### HasVcdbTdeKeyIdentifier

`func (o *ProvisionVDBByTimestampParameters) HasVcdbTdeKeyIdentifier() bool`

HasVcdbTdeKeyIdentifier returns a boolean if a field has been set.

### GetAppdataSourceParams

`func (o *ProvisionVDBByTimestampParameters) GetAppdataSourceParams() map[string]interface{}`

GetAppdataSourceParams returns the AppdataSourceParams field if non-nil, zero value otherwise.

### GetAppdataSourceParamsOk

`func (o *ProvisionVDBByTimestampParameters) GetAppdataSourceParamsOk() (*map[string]interface{}, bool)`

GetAppdataSourceParamsOk returns a tuple with the AppdataSourceParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppdataSourceParams

`func (o *ProvisionVDBByTimestampParameters) SetAppdataSourceParams(v map[string]interface{})`

SetAppdataSourceParams sets AppdataSourceParams field to given value.

### HasAppdataSourceParams

`func (o *ProvisionVDBByTimestampParameters) HasAppdataSourceParams() bool`

HasAppdataSourceParams returns a boolean if a field has been set.

### GetAdditionalMountPoints

`func (o *ProvisionVDBByTimestampParameters) GetAdditionalMountPoints() []AdditionalMountPoint`

GetAdditionalMountPoints returns the AdditionalMountPoints field if non-nil, zero value otherwise.

### GetAdditionalMountPointsOk

`func (o *ProvisionVDBByTimestampParameters) GetAdditionalMountPointsOk() (*[]AdditionalMountPoint, bool)`

GetAdditionalMountPointsOk returns a tuple with the AdditionalMountPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMountPoints

`func (o *ProvisionVDBByTimestampParameters) SetAdditionalMountPoints(v []AdditionalMountPoint)`

SetAdditionalMountPoints sets AdditionalMountPoints field to given value.

### HasAdditionalMountPoints

`func (o *ProvisionVDBByTimestampParameters) HasAdditionalMountPoints() bool`

HasAdditionalMountPoints returns a boolean if a field has been set.

### SetAdditionalMountPointsNil

`func (o *ProvisionVDBByTimestampParameters) SetAdditionalMountPointsNil(b bool)`

 SetAdditionalMountPointsNil sets the value for AdditionalMountPoints to be an explicit nil

### UnsetAdditionalMountPoints
`func (o *ProvisionVDBByTimestampParameters) UnsetAdditionalMountPoints()`

UnsetAdditionalMountPoints ensures that no value is present for AdditionalMountPoints, not even an explicit nil
### GetAppdataConfigParams

`func (o *ProvisionVDBByTimestampParameters) GetAppdataConfigParams() map[string]interface{}`

GetAppdataConfigParams returns the AppdataConfigParams field if non-nil, zero value otherwise.

### GetAppdataConfigParamsOk

`func (o *ProvisionVDBByTimestampParameters) GetAppdataConfigParamsOk() (*map[string]interface{}, bool)`

GetAppdataConfigParamsOk returns a tuple with the AppdataConfigParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppdataConfigParams

`func (o *ProvisionVDBByTimestampParameters) SetAppdataConfigParams(v map[string]interface{})`

SetAppdataConfigParams sets AppdataConfigParams field to given value.

### HasAppdataConfigParams

`func (o *ProvisionVDBByTimestampParameters) HasAppdataConfigParams() bool`

HasAppdataConfigParams returns a boolean if a field has been set.

### SetAppdataConfigParamsNil

`func (o *ProvisionVDBByTimestampParameters) SetAppdataConfigParamsNil(b bool)`

 SetAppdataConfigParamsNil sets the value for AppdataConfigParams to be an explicit nil

### UnsetAppdataConfigParams
`func (o *ProvisionVDBByTimestampParameters) UnsetAppdataConfigParams()`

UnsetAppdataConfigParams ensures that no value is present for AppdataConfigParams, not even an explicit nil
### GetConfigParams

`func (o *ProvisionVDBByTimestampParameters) GetConfigParams() map[string]interface{}`

GetConfigParams returns the ConfigParams field if non-nil, zero value otherwise.

### GetConfigParamsOk

`func (o *ProvisionVDBByTimestampParameters) GetConfigParamsOk() (*map[string]interface{}, bool)`

GetConfigParamsOk returns a tuple with the ConfigParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParams

`func (o *ProvisionVDBByTimestampParameters) SetConfigParams(v map[string]interface{})`

SetConfigParams sets ConfigParams field to given value.

### HasConfigParams

`func (o *ProvisionVDBByTimestampParameters) HasConfigParams() bool`

HasConfigParams returns a boolean if a field has been set.

### SetConfigParamsNil

`func (o *ProvisionVDBByTimestampParameters) SetConfigParamsNil(b bool)`

 SetConfigParamsNil sets the value for ConfigParams to be an explicit nil

### UnsetConfigParams
`func (o *ProvisionVDBByTimestampParameters) UnsetConfigParams()`

UnsetConfigParams ensures that no value is present for ConfigParams, not even an explicit nil
### GetPrivilegedOsUser

`func (o *ProvisionVDBByTimestampParameters) GetPrivilegedOsUser() string`

GetPrivilegedOsUser returns the PrivilegedOsUser field if non-nil, zero value otherwise.

### GetPrivilegedOsUserOk

`func (o *ProvisionVDBByTimestampParameters) GetPrivilegedOsUserOk() (*string, bool)`

GetPrivilegedOsUserOk returns a tuple with the PrivilegedOsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegedOsUser

`func (o *ProvisionVDBByTimestampParameters) SetPrivilegedOsUser(v string)`

SetPrivilegedOsUser sets PrivilegedOsUser field to given value.

### HasPrivilegedOsUser

`func (o *ProvisionVDBByTimestampParameters) HasPrivilegedOsUser() bool`

HasPrivilegedOsUser returns a boolean if a field has been set.

### GetPostgresPort

`func (o *ProvisionVDBByTimestampParameters) GetPostgresPort() int32`

GetPostgresPort returns the PostgresPort field if non-nil, zero value otherwise.

### GetPostgresPortOk

`func (o *ProvisionVDBByTimestampParameters) GetPostgresPortOk() (*int32, bool)`

GetPostgresPortOk returns a tuple with the PostgresPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresPort

`func (o *ProvisionVDBByTimestampParameters) SetPostgresPort(v int32)`

SetPostgresPort sets PostgresPort field to given value.

### HasPostgresPort

`func (o *ProvisionVDBByTimestampParameters) HasPostgresPort() bool`

HasPostgresPort returns a boolean if a field has been set.

### GetConfigSettingsStg

`func (o *ProvisionVDBByTimestampParameters) GetConfigSettingsStg() []ConfigSettingsStg`

GetConfigSettingsStg returns the ConfigSettingsStg field if non-nil, zero value otherwise.

### GetConfigSettingsStgOk

`func (o *ProvisionVDBByTimestampParameters) GetConfigSettingsStgOk() (*[]ConfigSettingsStg, bool)`

GetConfigSettingsStgOk returns a tuple with the ConfigSettingsStg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSettingsStg

`func (o *ProvisionVDBByTimestampParameters) SetConfigSettingsStg(v []ConfigSettingsStg)`

SetConfigSettingsStg sets ConfigSettingsStg field to given value.

### HasConfigSettingsStg

`func (o *ProvisionVDBByTimestampParameters) HasConfigSettingsStg() bool`

HasConfigSettingsStg returns a boolean if a field has been set.

### GetVcdbRestart

`func (o *ProvisionVDBByTimestampParameters) GetVcdbRestart() bool`

GetVcdbRestart returns the VcdbRestart field if non-nil, zero value otherwise.

### GetVcdbRestartOk

`func (o *ProvisionVDBByTimestampParameters) GetVcdbRestartOk() (*bool, bool)`

GetVcdbRestartOk returns a tuple with the VcdbRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdbRestart

`func (o *ProvisionVDBByTimestampParameters) SetVcdbRestart(v bool)`

SetVcdbRestart sets VcdbRestart field to given value.

### HasVcdbRestart

`func (o *ProvisionVDBByTimestampParameters) HasVcdbRestart() bool`

HasVcdbRestart returns a boolean if a field has been set.

### GetMssqlFailoverDriveLetter

`func (o *ProvisionVDBByTimestampParameters) GetMssqlFailoverDriveLetter() string`

GetMssqlFailoverDriveLetter returns the MssqlFailoverDriveLetter field if non-nil, zero value otherwise.

### GetMssqlFailoverDriveLetterOk

`func (o *ProvisionVDBByTimestampParameters) GetMssqlFailoverDriveLetterOk() (*string, bool)`

GetMssqlFailoverDriveLetterOk returns a tuple with the MssqlFailoverDriveLetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlFailoverDriveLetter

`func (o *ProvisionVDBByTimestampParameters) SetMssqlFailoverDriveLetter(v string)`

SetMssqlFailoverDriveLetter sets MssqlFailoverDriveLetter field to given value.

### HasMssqlFailoverDriveLetter

`func (o *ProvisionVDBByTimestampParameters) HasMssqlFailoverDriveLetter() bool`

HasMssqlFailoverDriveLetter returns a boolean if a field has been set.

### GetTags

`func (o *ProvisionVDBByTimestampParameters) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProvisionVDBByTimestampParameters) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProvisionVDBByTimestampParameters) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProvisionVDBByTimestampParameters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *ProvisionVDBByTimestampParameters) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProvisionVDBByTimestampParameters) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProvisionVDBByTimestampParameters) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProvisionVDBByTimestampParameters) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTimestampInDatabaseTimezone

`func (o *ProvisionVDBByTimestampParameters) GetTimestampInDatabaseTimezone() string`

GetTimestampInDatabaseTimezone returns the TimestampInDatabaseTimezone field if non-nil, zero value otherwise.

### GetTimestampInDatabaseTimezoneOk

`func (o *ProvisionVDBByTimestampParameters) GetTimestampInDatabaseTimezoneOk() (*string, bool)`

GetTimestampInDatabaseTimezoneOk returns a tuple with the TimestampInDatabaseTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampInDatabaseTimezone

`func (o *ProvisionVDBByTimestampParameters) SetTimestampInDatabaseTimezone(v string)`

SetTimestampInDatabaseTimezone sets TimestampInDatabaseTimezone field to given value.

### HasTimestampInDatabaseTimezone

`func (o *ProvisionVDBByTimestampParameters) HasTimestampInDatabaseTimezone() bool`

HasTimestampInDatabaseTimezone returns a boolean if a field has been set.

### GetTimeflowId

`func (o *ProvisionVDBByTimestampParameters) GetTimeflowId() string`

GetTimeflowId returns the TimeflowId field if non-nil, zero value otherwise.

### GetTimeflowIdOk

`func (o *ProvisionVDBByTimestampParameters) GetTimeflowIdOk() (*string, bool)`

GetTimeflowIdOk returns a tuple with the TimeflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeflowId

`func (o *ProvisionVDBByTimestampParameters) SetTimeflowId(v string)`

SetTimeflowId sets TimeflowId field to given value.

### HasTimeflowId

`func (o *ProvisionVDBByTimestampParameters) HasTimeflowId() bool`

HasTimeflowId returns a boolean if a field has been set.

### GetEngineId

`func (o *ProvisionVDBByTimestampParameters) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *ProvisionVDBByTimestampParameters) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *ProvisionVDBByTimestampParameters) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *ProvisionVDBByTimestampParameters) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetSourceDataId

`func (o *ProvisionVDBByTimestampParameters) GetSourceDataId() string`

GetSourceDataId returns the SourceDataId field if non-nil, zero value otherwise.

### GetSourceDataIdOk

`func (o *ProvisionVDBByTimestampParameters) GetSourceDataIdOk() (*string, bool)`

GetSourceDataIdOk returns a tuple with the SourceDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataId

`func (o *ProvisionVDBByTimestampParameters) SetSourceDataId(v string)`

SetSourceDataId sets SourceDataId field to given value.


### GetMakeCurrentAccountOwner

`func (o *ProvisionVDBByTimestampParameters) GetMakeCurrentAccountOwner() bool`

GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field if non-nil, zero value otherwise.

### GetMakeCurrentAccountOwnerOk

`func (o *ProvisionVDBByTimestampParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool)`

GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCurrentAccountOwner

`func (o *ProvisionVDBByTimestampParameters) SetMakeCurrentAccountOwner(v bool)`

SetMakeCurrentAccountOwner sets MakeCurrentAccountOwner field to given value.

### HasMakeCurrentAccountOwner

`func (o *ProvisionVDBByTimestampParameters) HasMakeCurrentAccountOwner() bool`

HasMakeCurrentAccountOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


