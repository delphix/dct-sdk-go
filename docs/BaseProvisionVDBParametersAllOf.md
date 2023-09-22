# BaseProvisionVDBParametersAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

## Methods

### NewBaseProvisionVDBParametersAllOf

`func NewBaseProvisionVDBParametersAllOf() *BaseProvisionVDBParametersAllOf`

NewBaseProvisionVDBParametersAllOf instantiates a new BaseProvisionVDBParametersAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseProvisionVDBParametersAllOfWithDefaults

`func NewBaseProvisionVDBParametersAllOfWithDefaults() *BaseProvisionVDBParametersAllOf`

NewBaseProvisionVDBParametersAllOfWithDefaults instantiates a new BaseProvisionVDBParametersAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetGroupId

`func (o *BaseProvisionVDBParametersAllOf) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *BaseProvisionVDBParametersAllOf) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *BaseProvisionVDBParametersAllOf) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *BaseProvisionVDBParametersAllOf) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### GetName

`func (o *BaseProvisionVDBParametersAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseProvisionVDBParametersAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseProvisionVDBParametersAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseProvisionVDBParametersAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDatabaseName

`func (o *BaseProvisionVDBParametersAllOf) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *BaseProvisionVDBParametersAllOf) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *BaseProvisionVDBParametersAllOf) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *BaseProvisionVDBParametersAllOf) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetCdbId

`func (o *BaseProvisionVDBParametersAllOf) GetCdbId() string`

GetCdbId returns the CdbId field if non-nil, zero value otherwise.

### GetCdbIdOk

`func (o *BaseProvisionVDBParametersAllOf) GetCdbIdOk() (*string, bool)`

GetCdbIdOk returns a tuple with the CdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdbId

`func (o *BaseProvisionVDBParametersAllOf) SetCdbId(v string)`

SetCdbId sets CdbId field to given value.

### HasCdbId

`func (o *BaseProvisionVDBParametersAllOf) HasCdbId() bool`

HasCdbId returns a boolean if a field has been set.

### GetClusterNodeIds

`func (o *BaseProvisionVDBParametersAllOf) GetClusterNodeIds() []string`

GetClusterNodeIds returns the ClusterNodeIds field if non-nil, zero value otherwise.

### GetClusterNodeIdsOk

`func (o *BaseProvisionVDBParametersAllOf) GetClusterNodeIdsOk() (*[]string, bool)`

GetClusterNodeIdsOk returns a tuple with the ClusterNodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeIds

`func (o *BaseProvisionVDBParametersAllOf) SetClusterNodeIds(v []string)`

SetClusterNodeIds sets ClusterNodeIds field to given value.

### HasClusterNodeIds

`func (o *BaseProvisionVDBParametersAllOf) HasClusterNodeIds() bool`

HasClusterNodeIds returns a boolean if a field has been set.

### GetClusterNodeInstances

`func (o *BaseProvisionVDBParametersAllOf) GetClusterNodeInstances() []ClusterNodeInstance`

GetClusterNodeInstances returns the ClusterNodeInstances field if non-nil, zero value otherwise.

### GetClusterNodeInstancesOk

`func (o *BaseProvisionVDBParametersAllOf) GetClusterNodeInstancesOk() (*[]ClusterNodeInstance, bool)`

GetClusterNodeInstancesOk returns a tuple with the ClusterNodeInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeInstances

`func (o *BaseProvisionVDBParametersAllOf) SetClusterNodeInstances(v []ClusterNodeInstance)`

SetClusterNodeInstances sets ClusterNodeInstances field to given value.

### HasClusterNodeInstances

`func (o *BaseProvisionVDBParametersAllOf) HasClusterNodeInstances() bool`

HasClusterNodeInstances returns a boolean if a field has been set.

### GetTruncateLogOnCheckpoint

`func (o *BaseProvisionVDBParametersAllOf) GetTruncateLogOnCheckpoint() bool`

GetTruncateLogOnCheckpoint returns the TruncateLogOnCheckpoint field if non-nil, zero value otherwise.

### GetTruncateLogOnCheckpointOk

`func (o *BaseProvisionVDBParametersAllOf) GetTruncateLogOnCheckpointOk() (*bool, bool)`

GetTruncateLogOnCheckpointOk returns a tuple with the TruncateLogOnCheckpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateLogOnCheckpoint

`func (o *BaseProvisionVDBParametersAllOf) SetTruncateLogOnCheckpoint(v bool)`

SetTruncateLogOnCheckpoint sets TruncateLogOnCheckpoint field to given value.

### HasTruncateLogOnCheckpoint

`func (o *BaseProvisionVDBParametersAllOf) HasTruncateLogOnCheckpoint() bool`

HasTruncateLogOnCheckpoint returns a boolean if a field has been set.

### GetOsUsername

`func (o *BaseProvisionVDBParametersAllOf) GetOsUsername() string`

GetOsUsername returns the OsUsername field if non-nil, zero value otherwise.

### GetOsUsernameOk

`func (o *BaseProvisionVDBParametersAllOf) GetOsUsernameOk() (*string, bool)`

GetOsUsernameOk returns a tuple with the OsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsUsername

`func (o *BaseProvisionVDBParametersAllOf) SetOsUsername(v string)`

SetOsUsername sets OsUsername field to given value.

### HasOsUsername

`func (o *BaseProvisionVDBParametersAllOf) HasOsUsername() bool`

HasOsUsername returns a boolean if a field has been set.

### GetOsPassword

`func (o *BaseProvisionVDBParametersAllOf) GetOsPassword() string`

GetOsPassword returns the OsPassword field if non-nil, zero value otherwise.

### GetOsPasswordOk

`func (o *BaseProvisionVDBParametersAllOf) GetOsPasswordOk() (*string, bool)`

GetOsPasswordOk returns a tuple with the OsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPassword

`func (o *BaseProvisionVDBParametersAllOf) SetOsPassword(v string)`

SetOsPassword sets OsPassword field to given value.

### HasOsPassword

`func (o *BaseProvisionVDBParametersAllOf) HasOsPassword() bool`

HasOsPassword returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *BaseProvisionVDBParametersAllOf) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *BaseProvisionVDBParametersAllOf) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *BaseProvisionVDBParametersAllOf) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *BaseProvisionVDBParametersAllOf) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEnvironmentUserId

`func (o *BaseProvisionVDBParametersAllOf) GetEnvironmentUserId() string`

GetEnvironmentUserId returns the EnvironmentUserId field if non-nil, zero value otherwise.

### GetEnvironmentUserIdOk

`func (o *BaseProvisionVDBParametersAllOf) GetEnvironmentUserIdOk() (*string, bool)`

GetEnvironmentUserIdOk returns a tuple with the EnvironmentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUserId

`func (o *BaseProvisionVDBParametersAllOf) SetEnvironmentUserId(v string)`

SetEnvironmentUserId sets EnvironmentUserId field to given value.

### HasEnvironmentUserId

`func (o *BaseProvisionVDBParametersAllOf) HasEnvironmentUserId() bool`

HasEnvironmentUserId returns a boolean if a field has been set.

### GetRepositoryId

`func (o *BaseProvisionVDBParametersAllOf) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *BaseProvisionVDBParametersAllOf) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *BaseProvisionVDBParametersAllOf) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *BaseProvisionVDBParametersAllOf) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetAutoSelectRepository

`func (o *BaseProvisionVDBParametersAllOf) GetAutoSelectRepository() bool`

GetAutoSelectRepository returns the AutoSelectRepository field if non-nil, zero value otherwise.

### GetAutoSelectRepositoryOk

`func (o *BaseProvisionVDBParametersAllOf) GetAutoSelectRepositoryOk() (*bool, bool)`

GetAutoSelectRepositoryOk returns a tuple with the AutoSelectRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelectRepository

`func (o *BaseProvisionVDBParametersAllOf) SetAutoSelectRepository(v bool)`

SetAutoSelectRepository sets AutoSelectRepository field to given value.

### HasAutoSelectRepository

`func (o *BaseProvisionVDBParametersAllOf) HasAutoSelectRepository() bool`

HasAutoSelectRepository returns a boolean if a field has been set.

### GetVdbRestart

`func (o *BaseProvisionVDBParametersAllOf) GetVdbRestart() bool`

GetVdbRestart returns the VdbRestart field if non-nil, zero value otherwise.

### GetVdbRestartOk

`func (o *BaseProvisionVDBParametersAllOf) GetVdbRestartOk() (*bool, bool)`

GetVdbRestartOk returns a tuple with the VdbRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdbRestart

`func (o *BaseProvisionVDBParametersAllOf) SetVdbRestart(v bool)`

SetVdbRestart sets VdbRestart field to given value.

### HasVdbRestart

`func (o *BaseProvisionVDBParametersAllOf) HasVdbRestart() bool`

HasVdbRestart returns a boolean if a field has been set.

### GetTemplateId

`func (o *BaseProvisionVDBParametersAllOf) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *BaseProvisionVDBParametersAllOf) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *BaseProvisionVDBParametersAllOf) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *BaseProvisionVDBParametersAllOf) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetAuxiliaryTemplateId

`func (o *BaseProvisionVDBParametersAllOf) GetAuxiliaryTemplateId() string`

GetAuxiliaryTemplateId returns the AuxiliaryTemplateId field if non-nil, zero value otherwise.

### GetAuxiliaryTemplateIdOk

`func (o *BaseProvisionVDBParametersAllOf) GetAuxiliaryTemplateIdOk() (*string, bool)`

GetAuxiliaryTemplateIdOk returns a tuple with the AuxiliaryTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryTemplateId

`func (o *BaseProvisionVDBParametersAllOf) SetAuxiliaryTemplateId(v string)`

SetAuxiliaryTemplateId sets AuxiliaryTemplateId field to given value.

### HasAuxiliaryTemplateId

`func (o *BaseProvisionVDBParametersAllOf) HasAuxiliaryTemplateId() bool`

HasAuxiliaryTemplateId returns a boolean if a field has been set.

### GetFileMappingRules

`func (o *BaseProvisionVDBParametersAllOf) GetFileMappingRules() string`

GetFileMappingRules returns the FileMappingRules field if non-nil, zero value otherwise.

### GetFileMappingRulesOk

`func (o *BaseProvisionVDBParametersAllOf) GetFileMappingRulesOk() (*string, bool)`

GetFileMappingRulesOk returns a tuple with the FileMappingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMappingRules

`func (o *BaseProvisionVDBParametersAllOf) SetFileMappingRules(v string)`

SetFileMappingRules sets FileMappingRules field to given value.

### HasFileMappingRules

`func (o *BaseProvisionVDBParametersAllOf) HasFileMappingRules() bool`

HasFileMappingRules returns a boolean if a field has been set.

### GetOracleInstanceName

`func (o *BaseProvisionVDBParametersAllOf) GetOracleInstanceName() string`

GetOracleInstanceName returns the OracleInstanceName field if non-nil, zero value otherwise.

### GetOracleInstanceNameOk

`func (o *BaseProvisionVDBParametersAllOf) GetOracleInstanceNameOk() (*string, bool)`

GetOracleInstanceNameOk returns a tuple with the OracleInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleInstanceName

`func (o *BaseProvisionVDBParametersAllOf) SetOracleInstanceName(v string)`

SetOracleInstanceName sets OracleInstanceName field to given value.

### HasOracleInstanceName

`func (o *BaseProvisionVDBParametersAllOf) HasOracleInstanceName() bool`

HasOracleInstanceName returns a boolean if a field has been set.

### GetUniqueName

`func (o *BaseProvisionVDBParametersAllOf) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *BaseProvisionVDBParametersAllOf) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *BaseProvisionVDBParametersAllOf) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *BaseProvisionVDBParametersAllOf) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetVcdbName

`func (o *BaseProvisionVDBParametersAllOf) GetVcdbName() string`

GetVcdbName returns the VcdbName field if non-nil, zero value otherwise.

### GetVcdbNameOk

`func (o *BaseProvisionVDBParametersAllOf) GetVcdbNameOk() (*string, bool)`

GetVcdbNameOk returns a tuple with the VcdbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdbName

`func (o *BaseProvisionVDBParametersAllOf) SetVcdbName(v string)`

SetVcdbName sets VcdbName field to given value.

### HasVcdbName

`func (o *BaseProvisionVDBParametersAllOf) HasVcdbName() bool`

HasVcdbName returns a boolean if a field has been set.

### GetVcdbDatabaseName

`func (o *BaseProvisionVDBParametersAllOf) GetVcdbDatabaseName() string`

GetVcdbDatabaseName returns the VcdbDatabaseName field if non-nil, zero value otherwise.

### GetVcdbDatabaseNameOk

`func (o *BaseProvisionVDBParametersAllOf) GetVcdbDatabaseNameOk() (*string, bool)`

GetVcdbDatabaseNameOk returns a tuple with the VcdbDatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdbDatabaseName

`func (o *BaseProvisionVDBParametersAllOf) SetVcdbDatabaseName(v string)`

SetVcdbDatabaseName sets VcdbDatabaseName field to given value.

### HasVcdbDatabaseName

`func (o *BaseProvisionVDBParametersAllOf) HasVcdbDatabaseName() bool`

HasVcdbDatabaseName returns a boolean if a field has been set.

### GetMountPoint

`func (o *BaseProvisionVDBParametersAllOf) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *BaseProvisionVDBParametersAllOf) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *BaseProvisionVDBParametersAllOf) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *BaseProvisionVDBParametersAllOf) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### GetOpenResetLogs

`func (o *BaseProvisionVDBParametersAllOf) GetOpenResetLogs() bool`

GetOpenResetLogs returns the OpenResetLogs field if non-nil, zero value otherwise.

### GetOpenResetLogsOk

`func (o *BaseProvisionVDBParametersAllOf) GetOpenResetLogsOk() (*bool, bool)`

GetOpenResetLogsOk returns a tuple with the OpenResetLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenResetLogs

`func (o *BaseProvisionVDBParametersAllOf) SetOpenResetLogs(v bool)`

SetOpenResetLogs sets OpenResetLogs field to given value.

### HasOpenResetLogs

`func (o *BaseProvisionVDBParametersAllOf) HasOpenResetLogs() bool`

HasOpenResetLogs returns a boolean if a field has been set.

### GetSnapshotPolicyId

`func (o *BaseProvisionVDBParametersAllOf) GetSnapshotPolicyId() string`

GetSnapshotPolicyId returns the SnapshotPolicyId field if non-nil, zero value otherwise.

### GetSnapshotPolicyIdOk

`func (o *BaseProvisionVDBParametersAllOf) GetSnapshotPolicyIdOk() (*string, bool)`

GetSnapshotPolicyIdOk returns a tuple with the SnapshotPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotPolicyId

`func (o *BaseProvisionVDBParametersAllOf) SetSnapshotPolicyId(v string)`

SetSnapshotPolicyId sets SnapshotPolicyId field to given value.

### HasSnapshotPolicyId

`func (o *BaseProvisionVDBParametersAllOf) HasSnapshotPolicyId() bool`

HasSnapshotPolicyId returns a boolean if a field has been set.

### GetRetentionPolicyId

`func (o *BaseProvisionVDBParametersAllOf) GetRetentionPolicyId() string`

GetRetentionPolicyId returns the RetentionPolicyId field if non-nil, zero value otherwise.

### GetRetentionPolicyIdOk

`func (o *BaseProvisionVDBParametersAllOf) GetRetentionPolicyIdOk() (*string, bool)`

GetRetentionPolicyIdOk returns a tuple with the RetentionPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicyId

`func (o *BaseProvisionVDBParametersAllOf) SetRetentionPolicyId(v string)`

SetRetentionPolicyId sets RetentionPolicyId field to given value.

### HasRetentionPolicyId

`func (o *BaseProvisionVDBParametersAllOf) HasRetentionPolicyId() bool`

HasRetentionPolicyId returns a boolean if a field has been set.

### GetRecoveryModel

`func (o *BaseProvisionVDBParametersAllOf) GetRecoveryModel() string`

GetRecoveryModel returns the RecoveryModel field if non-nil, zero value otherwise.

### GetRecoveryModelOk

`func (o *BaseProvisionVDBParametersAllOf) GetRecoveryModelOk() (*string, bool)`

GetRecoveryModelOk returns a tuple with the RecoveryModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryModel

`func (o *BaseProvisionVDBParametersAllOf) SetRecoveryModel(v string)`

SetRecoveryModel sets RecoveryModel field to given value.

### HasRecoveryModel

`func (o *BaseProvisionVDBParametersAllOf) HasRecoveryModel() bool`

HasRecoveryModel returns a boolean if a field has been set.

### GetPreScript

`func (o *BaseProvisionVDBParametersAllOf) GetPreScript() string`

GetPreScript returns the PreScript field if non-nil, zero value otherwise.

### GetPreScriptOk

`func (o *BaseProvisionVDBParametersAllOf) GetPreScriptOk() (*string, bool)`

GetPreScriptOk returns a tuple with the PreScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreScript

`func (o *BaseProvisionVDBParametersAllOf) SetPreScript(v string)`

SetPreScript sets PreScript field to given value.

### HasPreScript

`func (o *BaseProvisionVDBParametersAllOf) HasPreScript() bool`

HasPreScript returns a boolean if a field has been set.

### GetPostScript

`func (o *BaseProvisionVDBParametersAllOf) GetPostScript() string`

GetPostScript returns the PostScript field if non-nil, zero value otherwise.

### GetPostScriptOk

`func (o *BaseProvisionVDBParametersAllOf) GetPostScriptOk() (*string, bool)`

GetPostScriptOk returns a tuple with the PostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScript

`func (o *BaseProvisionVDBParametersAllOf) SetPostScript(v string)`

SetPostScript sets PostScript field to given value.

### HasPostScript

`func (o *BaseProvisionVDBParametersAllOf) HasPostScript() bool`

HasPostScript returns a boolean if a field has been set.

### GetCdcOnProvision

`func (o *BaseProvisionVDBParametersAllOf) GetCdcOnProvision() bool`

GetCdcOnProvision returns the CdcOnProvision field if non-nil, zero value otherwise.

### GetCdcOnProvisionOk

`func (o *BaseProvisionVDBParametersAllOf) GetCdcOnProvisionOk() (*bool, bool)`

GetCdcOnProvisionOk returns a tuple with the CdcOnProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdcOnProvision

`func (o *BaseProvisionVDBParametersAllOf) SetCdcOnProvision(v bool)`

SetCdcOnProvision sets CdcOnProvision field to given value.

### HasCdcOnProvision

`func (o *BaseProvisionVDBParametersAllOf) HasCdcOnProvision() bool`

HasCdcOnProvision returns a boolean if a field has been set.

### GetOnlineLogSize

`func (o *BaseProvisionVDBParametersAllOf) GetOnlineLogSize() int32`

GetOnlineLogSize returns the OnlineLogSize field if non-nil, zero value otherwise.

### GetOnlineLogSizeOk

`func (o *BaseProvisionVDBParametersAllOf) GetOnlineLogSizeOk() (*int32, bool)`

GetOnlineLogSizeOk returns a tuple with the OnlineLogSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineLogSize

`func (o *BaseProvisionVDBParametersAllOf) SetOnlineLogSize(v int32)`

SetOnlineLogSize sets OnlineLogSize field to given value.

### HasOnlineLogSize

`func (o *BaseProvisionVDBParametersAllOf) HasOnlineLogSize() bool`

HasOnlineLogSize returns a boolean if a field has been set.

### GetOnlineLogGroups

`func (o *BaseProvisionVDBParametersAllOf) GetOnlineLogGroups() int32`

GetOnlineLogGroups returns the OnlineLogGroups field if non-nil, zero value otherwise.

### GetOnlineLogGroupsOk

`func (o *BaseProvisionVDBParametersAllOf) GetOnlineLogGroupsOk() (*int32, bool)`

GetOnlineLogGroupsOk returns a tuple with the OnlineLogGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineLogGroups

`func (o *BaseProvisionVDBParametersAllOf) SetOnlineLogGroups(v int32)`

SetOnlineLogGroups sets OnlineLogGroups field to given value.

### HasOnlineLogGroups

`func (o *BaseProvisionVDBParametersAllOf) HasOnlineLogGroups() bool`

HasOnlineLogGroups returns a boolean if a field has been set.

### GetArchiveLog

`func (o *BaseProvisionVDBParametersAllOf) GetArchiveLog() bool`

GetArchiveLog returns the ArchiveLog field if non-nil, zero value otherwise.

### GetArchiveLogOk

`func (o *BaseProvisionVDBParametersAllOf) GetArchiveLogOk() (*bool, bool)`

GetArchiveLogOk returns a tuple with the ArchiveLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLog

`func (o *BaseProvisionVDBParametersAllOf) SetArchiveLog(v bool)`

SetArchiveLog sets ArchiveLog field to given value.

### HasArchiveLog

`func (o *BaseProvisionVDBParametersAllOf) HasArchiveLog() bool`

HasArchiveLog returns a boolean if a field has been set.

### GetNewDbid

`func (o *BaseProvisionVDBParametersAllOf) GetNewDbid() bool`

GetNewDbid returns the NewDbid field if non-nil, zero value otherwise.

### GetNewDbidOk

`func (o *BaseProvisionVDBParametersAllOf) GetNewDbidOk() (*bool, bool)`

GetNewDbidOk returns a tuple with the NewDbid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDbid

`func (o *BaseProvisionVDBParametersAllOf) SetNewDbid(v bool)`

SetNewDbid sets NewDbid field to given value.

### HasNewDbid

`func (o *BaseProvisionVDBParametersAllOf) HasNewDbid() bool`

HasNewDbid returns a boolean if a field has been set.

### GetListenerIds

`func (o *BaseProvisionVDBParametersAllOf) GetListenerIds() []string`

GetListenerIds returns the ListenerIds field if non-nil, zero value otherwise.

### GetListenerIdsOk

`func (o *BaseProvisionVDBParametersAllOf) GetListenerIdsOk() (*[]string, bool)`

GetListenerIdsOk returns a tuple with the ListenerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerIds

`func (o *BaseProvisionVDBParametersAllOf) SetListenerIds(v []string)`

SetListenerIds sets ListenerIds field to given value.

### HasListenerIds

`func (o *BaseProvisionVDBParametersAllOf) HasListenerIds() bool`

HasListenerIds returns a boolean if a field has been set.

### GetCustomEnvVars

`func (o *BaseProvisionVDBParametersAllOf) GetCustomEnvVars() map[string]string`

GetCustomEnvVars returns the CustomEnvVars field if non-nil, zero value otherwise.

### GetCustomEnvVarsOk

`func (o *BaseProvisionVDBParametersAllOf) GetCustomEnvVarsOk() (*map[string]string, bool)`

GetCustomEnvVarsOk returns a tuple with the CustomEnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvVars

`func (o *BaseProvisionVDBParametersAllOf) SetCustomEnvVars(v map[string]string)`

SetCustomEnvVars sets CustomEnvVars field to given value.

### HasCustomEnvVars

`func (o *BaseProvisionVDBParametersAllOf) HasCustomEnvVars() bool`

HasCustomEnvVars returns a boolean if a field has been set.

### GetCustomEnvFiles

`func (o *BaseProvisionVDBParametersAllOf) GetCustomEnvFiles() []string`

GetCustomEnvFiles returns the CustomEnvFiles field if non-nil, zero value otherwise.

### GetCustomEnvFilesOk

`func (o *BaseProvisionVDBParametersAllOf) GetCustomEnvFilesOk() (*[]string, bool)`

GetCustomEnvFilesOk returns a tuple with the CustomEnvFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvFiles

`func (o *BaseProvisionVDBParametersAllOf) SetCustomEnvFiles(v []string)`

SetCustomEnvFiles sets CustomEnvFiles field to given value.

### HasCustomEnvFiles

`func (o *BaseProvisionVDBParametersAllOf) HasCustomEnvFiles() bool`

HasCustomEnvFiles returns a boolean if a field has been set.

### GetOracleRacCustomEnvFiles

`func (o *BaseProvisionVDBParametersAllOf) GetOracleRacCustomEnvFiles() []OracleRacCustomEnvFile`

GetOracleRacCustomEnvFiles returns the OracleRacCustomEnvFiles field if non-nil, zero value otherwise.

### GetOracleRacCustomEnvFilesOk

`func (o *BaseProvisionVDBParametersAllOf) GetOracleRacCustomEnvFilesOk() (*[]OracleRacCustomEnvFile, bool)`

GetOracleRacCustomEnvFilesOk returns a tuple with the OracleRacCustomEnvFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleRacCustomEnvFiles

`func (o *BaseProvisionVDBParametersAllOf) SetOracleRacCustomEnvFiles(v []OracleRacCustomEnvFile)`

SetOracleRacCustomEnvFiles sets OracleRacCustomEnvFiles field to given value.

### HasOracleRacCustomEnvFiles

`func (o *BaseProvisionVDBParametersAllOf) HasOracleRacCustomEnvFiles() bool`

HasOracleRacCustomEnvFiles returns a boolean if a field has been set.

### GetOracleRacCustomEnvVars

`func (o *BaseProvisionVDBParametersAllOf) GetOracleRacCustomEnvVars() []OracleRacCustomEnvVar`

GetOracleRacCustomEnvVars returns the OracleRacCustomEnvVars field if non-nil, zero value otherwise.

### GetOracleRacCustomEnvVarsOk

`func (o *BaseProvisionVDBParametersAllOf) GetOracleRacCustomEnvVarsOk() (*[]OracleRacCustomEnvVar, bool)`

GetOracleRacCustomEnvVarsOk returns a tuple with the OracleRacCustomEnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleRacCustomEnvVars

`func (o *BaseProvisionVDBParametersAllOf) SetOracleRacCustomEnvVars(v []OracleRacCustomEnvVar)`

SetOracleRacCustomEnvVars sets OracleRacCustomEnvVars field to given value.

### HasOracleRacCustomEnvVars

`func (o *BaseProvisionVDBParametersAllOf) HasOracleRacCustomEnvVars() bool`

HasOracleRacCustomEnvVars returns a boolean if a field has been set.

### GetParentTdeKeystorePath

`func (o *BaseProvisionVDBParametersAllOf) GetParentTdeKeystorePath() string`

GetParentTdeKeystorePath returns the ParentTdeKeystorePath field if non-nil, zero value otherwise.

### GetParentTdeKeystorePathOk

`func (o *BaseProvisionVDBParametersAllOf) GetParentTdeKeystorePathOk() (*string, bool)`

GetParentTdeKeystorePathOk returns a tuple with the ParentTdeKeystorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTdeKeystorePath

`func (o *BaseProvisionVDBParametersAllOf) SetParentTdeKeystorePath(v string)`

SetParentTdeKeystorePath sets ParentTdeKeystorePath field to given value.

### HasParentTdeKeystorePath

`func (o *BaseProvisionVDBParametersAllOf) HasParentTdeKeystorePath() bool`

HasParentTdeKeystorePath returns a boolean if a field has been set.

### GetParentTdeKeystorePassword

`func (o *BaseProvisionVDBParametersAllOf) GetParentTdeKeystorePassword() string`

GetParentTdeKeystorePassword returns the ParentTdeKeystorePassword field if non-nil, zero value otherwise.

### GetParentTdeKeystorePasswordOk

`func (o *BaseProvisionVDBParametersAllOf) GetParentTdeKeystorePasswordOk() (*string, bool)`

GetParentTdeKeystorePasswordOk returns a tuple with the ParentTdeKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTdeKeystorePassword

`func (o *BaseProvisionVDBParametersAllOf) SetParentTdeKeystorePassword(v string)`

SetParentTdeKeystorePassword sets ParentTdeKeystorePassword field to given value.

### HasParentTdeKeystorePassword

`func (o *BaseProvisionVDBParametersAllOf) HasParentTdeKeystorePassword() bool`

HasParentTdeKeystorePassword returns a boolean if a field has been set.

### GetTdeExportedKeyFileSecret

`func (o *BaseProvisionVDBParametersAllOf) GetTdeExportedKeyFileSecret() string`

GetTdeExportedKeyFileSecret returns the TdeExportedKeyFileSecret field if non-nil, zero value otherwise.

### GetTdeExportedKeyFileSecretOk

`func (o *BaseProvisionVDBParametersAllOf) GetTdeExportedKeyFileSecretOk() (*string, bool)`

GetTdeExportedKeyFileSecretOk returns a tuple with the TdeExportedKeyFileSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdeExportedKeyFileSecret

`func (o *BaseProvisionVDBParametersAllOf) SetTdeExportedKeyFileSecret(v string)`

SetTdeExportedKeyFileSecret sets TdeExportedKeyFileSecret field to given value.

### HasTdeExportedKeyFileSecret

`func (o *BaseProvisionVDBParametersAllOf) HasTdeExportedKeyFileSecret() bool`

HasTdeExportedKeyFileSecret returns a boolean if a field has been set.

### GetTdeKeyIdentifier

`func (o *BaseProvisionVDBParametersAllOf) GetTdeKeyIdentifier() string`

GetTdeKeyIdentifier returns the TdeKeyIdentifier field if non-nil, zero value otherwise.

### GetTdeKeyIdentifierOk

`func (o *BaseProvisionVDBParametersAllOf) GetTdeKeyIdentifierOk() (*string, bool)`

GetTdeKeyIdentifierOk returns a tuple with the TdeKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdeKeyIdentifier

`func (o *BaseProvisionVDBParametersAllOf) SetTdeKeyIdentifier(v string)`

SetTdeKeyIdentifier sets TdeKeyIdentifier field to given value.

### HasTdeKeyIdentifier

`func (o *BaseProvisionVDBParametersAllOf) HasTdeKeyIdentifier() bool`

HasTdeKeyIdentifier returns a boolean if a field has been set.

### GetTargetVcdbTdeKeystorePath

`func (o *BaseProvisionVDBParametersAllOf) GetTargetVcdbTdeKeystorePath() string`

GetTargetVcdbTdeKeystorePath returns the TargetVcdbTdeKeystorePath field if non-nil, zero value otherwise.

### GetTargetVcdbTdeKeystorePathOk

`func (o *BaseProvisionVDBParametersAllOf) GetTargetVcdbTdeKeystorePathOk() (*string, bool)`

GetTargetVcdbTdeKeystorePathOk returns a tuple with the TargetVcdbTdeKeystorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVcdbTdeKeystorePath

`func (o *BaseProvisionVDBParametersAllOf) SetTargetVcdbTdeKeystorePath(v string)`

SetTargetVcdbTdeKeystorePath sets TargetVcdbTdeKeystorePath field to given value.

### HasTargetVcdbTdeKeystorePath

`func (o *BaseProvisionVDBParametersAllOf) HasTargetVcdbTdeKeystorePath() bool`

HasTargetVcdbTdeKeystorePath returns a boolean if a field has been set.

### GetCdbTdeKeystorePassword

`func (o *BaseProvisionVDBParametersAllOf) GetCdbTdeKeystorePassword() string`

GetCdbTdeKeystorePassword returns the CdbTdeKeystorePassword field if non-nil, zero value otherwise.

### GetCdbTdeKeystorePasswordOk

`func (o *BaseProvisionVDBParametersAllOf) GetCdbTdeKeystorePasswordOk() (*string, bool)`

GetCdbTdeKeystorePasswordOk returns a tuple with the CdbTdeKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdbTdeKeystorePassword

`func (o *BaseProvisionVDBParametersAllOf) SetCdbTdeKeystorePassword(v string)`

SetCdbTdeKeystorePassword sets CdbTdeKeystorePassword field to given value.

### HasCdbTdeKeystorePassword

`func (o *BaseProvisionVDBParametersAllOf) HasCdbTdeKeystorePassword() bool`

HasCdbTdeKeystorePassword returns a boolean if a field has been set.

### GetVcdbTdeKeyIdentifier

`func (o *BaseProvisionVDBParametersAllOf) GetVcdbTdeKeyIdentifier() string`

GetVcdbTdeKeyIdentifier returns the VcdbTdeKeyIdentifier field if non-nil, zero value otherwise.

### GetVcdbTdeKeyIdentifierOk

`func (o *BaseProvisionVDBParametersAllOf) GetVcdbTdeKeyIdentifierOk() (*string, bool)`

GetVcdbTdeKeyIdentifierOk returns a tuple with the VcdbTdeKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdbTdeKeyIdentifier

`func (o *BaseProvisionVDBParametersAllOf) SetVcdbTdeKeyIdentifier(v string)`

SetVcdbTdeKeyIdentifier sets VcdbTdeKeyIdentifier field to given value.

### HasVcdbTdeKeyIdentifier

`func (o *BaseProvisionVDBParametersAllOf) HasVcdbTdeKeyIdentifier() bool`

HasVcdbTdeKeyIdentifier returns a boolean if a field has been set.

### GetAppdataSourceParams

`func (o *BaseProvisionVDBParametersAllOf) GetAppdataSourceParams() map[string]interface{}`

GetAppdataSourceParams returns the AppdataSourceParams field if non-nil, zero value otherwise.

### GetAppdataSourceParamsOk

`func (o *BaseProvisionVDBParametersAllOf) GetAppdataSourceParamsOk() (*map[string]interface{}, bool)`

GetAppdataSourceParamsOk returns a tuple with the AppdataSourceParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppdataSourceParams

`func (o *BaseProvisionVDBParametersAllOf) SetAppdataSourceParams(v map[string]interface{})`

SetAppdataSourceParams sets AppdataSourceParams field to given value.

### HasAppdataSourceParams

`func (o *BaseProvisionVDBParametersAllOf) HasAppdataSourceParams() bool`

HasAppdataSourceParams returns a boolean if a field has been set.

### GetAdditionalMountPoints

`func (o *BaseProvisionVDBParametersAllOf) GetAdditionalMountPoints() []AdditionalMountPoint`

GetAdditionalMountPoints returns the AdditionalMountPoints field if non-nil, zero value otherwise.

### GetAdditionalMountPointsOk

`func (o *BaseProvisionVDBParametersAllOf) GetAdditionalMountPointsOk() (*[]AdditionalMountPoint, bool)`

GetAdditionalMountPointsOk returns a tuple with the AdditionalMountPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMountPoints

`func (o *BaseProvisionVDBParametersAllOf) SetAdditionalMountPoints(v []AdditionalMountPoint)`

SetAdditionalMountPoints sets AdditionalMountPoints field to given value.

### HasAdditionalMountPoints

`func (o *BaseProvisionVDBParametersAllOf) HasAdditionalMountPoints() bool`

HasAdditionalMountPoints returns a boolean if a field has been set.

### SetAdditionalMountPointsNil

`func (o *BaseProvisionVDBParametersAllOf) SetAdditionalMountPointsNil(b bool)`

 SetAdditionalMountPointsNil sets the value for AdditionalMountPoints to be an explicit nil

### UnsetAdditionalMountPoints
`func (o *BaseProvisionVDBParametersAllOf) UnsetAdditionalMountPoints()`

UnsetAdditionalMountPoints ensures that no value is present for AdditionalMountPoints, not even an explicit nil
### GetAppdataConfigParams

`func (o *BaseProvisionVDBParametersAllOf) GetAppdataConfigParams() map[string]interface{}`

GetAppdataConfigParams returns the AppdataConfigParams field if non-nil, zero value otherwise.

### GetAppdataConfigParamsOk

`func (o *BaseProvisionVDBParametersAllOf) GetAppdataConfigParamsOk() (*map[string]interface{}, bool)`

GetAppdataConfigParamsOk returns a tuple with the AppdataConfigParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppdataConfigParams

`func (o *BaseProvisionVDBParametersAllOf) SetAppdataConfigParams(v map[string]interface{})`

SetAppdataConfigParams sets AppdataConfigParams field to given value.

### HasAppdataConfigParams

`func (o *BaseProvisionVDBParametersAllOf) HasAppdataConfigParams() bool`

HasAppdataConfigParams returns a boolean if a field has been set.

### SetAppdataConfigParamsNil

`func (o *BaseProvisionVDBParametersAllOf) SetAppdataConfigParamsNil(b bool)`

 SetAppdataConfigParamsNil sets the value for AppdataConfigParams to be an explicit nil

### UnsetAppdataConfigParams
`func (o *BaseProvisionVDBParametersAllOf) UnsetAppdataConfigParams()`

UnsetAppdataConfigParams ensures that no value is present for AppdataConfigParams, not even an explicit nil
### GetConfigParams

`func (o *BaseProvisionVDBParametersAllOf) GetConfigParams() map[string]interface{}`

GetConfigParams returns the ConfigParams field if non-nil, zero value otherwise.

### GetConfigParamsOk

`func (o *BaseProvisionVDBParametersAllOf) GetConfigParamsOk() (*map[string]interface{}, bool)`

GetConfigParamsOk returns a tuple with the ConfigParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParams

`func (o *BaseProvisionVDBParametersAllOf) SetConfigParams(v map[string]interface{})`

SetConfigParams sets ConfigParams field to given value.

### HasConfigParams

`func (o *BaseProvisionVDBParametersAllOf) HasConfigParams() bool`

HasConfigParams returns a boolean if a field has been set.

### SetConfigParamsNil

`func (o *BaseProvisionVDBParametersAllOf) SetConfigParamsNil(b bool)`

 SetConfigParamsNil sets the value for ConfigParams to be an explicit nil

### UnsetConfigParams
`func (o *BaseProvisionVDBParametersAllOf) UnsetConfigParams()`

UnsetConfigParams ensures that no value is present for ConfigParams, not even an explicit nil
### GetPrivilegedOsUser

`func (o *BaseProvisionVDBParametersAllOf) GetPrivilegedOsUser() string`

GetPrivilegedOsUser returns the PrivilegedOsUser field if non-nil, zero value otherwise.

### GetPrivilegedOsUserOk

`func (o *BaseProvisionVDBParametersAllOf) GetPrivilegedOsUserOk() (*string, bool)`

GetPrivilegedOsUserOk returns a tuple with the PrivilegedOsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegedOsUser

`func (o *BaseProvisionVDBParametersAllOf) SetPrivilegedOsUser(v string)`

SetPrivilegedOsUser sets PrivilegedOsUser field to given value.

### HasPrivilegedOsUser

`func (o *BaseProvisionVDBParametersAllOf) HasPrivilegedOsUser() bool`

HasPrivilegedOsUser returns a boolean if a field has been set.

### GetPostgresPort

`func (o *BaseProvisionVDBParametersAllOf) GetPostgresPort() int32`

GetPostgresPort returns the PostgresPort field if non-nil, zero value otherwise.

### GetPostgresPortOk

`func (o *BaseProvisionVDBParametersAllOf) GetPostgresPortOk() (*int32, bool)`

GetPostgresPortOk returns a tuple with the PostgresPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresPort

`func (o *BaseProvisionVDBParametersAllOf) SetPostgresPort(v int32)`

SetPostgresPort sets PostgresPort field to given value.

### HasPostgresPort

`func (o *BaseProvisionVDBParametersAllOf) HasPostgresPort() bool`

HasPostgresPort returns a boolean if a field has been set.

### GetConfigSettingsStg

`func (o *BaseProvisionVDBParametersAllOf) GetConfigSettingsStg() []ConfigSettingsStg`

GetConfigSettingsStg returns the ConfigSettingsStg field if non-nil, zero value otherwise.

### GetConfigSettingsStgOk

`func (o *BaseProvisionVDBParametersAllOf) GetConfigSettingsStgOk() (*[]ConfigSettingsStg, bool)`

GetConfigSettingsStgOk returns a tuple with the ConfigSettingsStg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSettingsStg

`func (o *BaseProvisionVDBParametersAllOf) SetConfigSettingsStg(v []ConfigSettingsStg)`

SetConfigSettingsStg sets ConfigSettingsStg field to given value.

### HasConfigSettingsStg

`func (o *BaseProvisionVDBParametersAllOf) HasConfigSettingsStg() bool`

HasConfigSettingsStg returns a boolean if a field has been set.

### GetVcdbRestart

`func (o *BaseProvisionVDBParametersAllOf) GetVcdbRestart() bool`

GetVcdbRestart returns the VcdbRestart field if non-nil, zero value otherwise.

### GetVcdbRestartOk

`func (o *BaseProvisionVDBParametersAllOf) GetVcdbRestartOk() (*bool, bool)`

GetVcdbRestartOk returns a tuple with the VcdbRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdbRestart

`func (o *BaseProvisionVDBParametersAllOf) SetVcdbRestart(v bool)`

SetVcdbRestart sets VcdbRestart field to given value.

### HasVcdbRestart

`func (o *BaseProvisionVDBParametersAllOf) HasVcdbRestart() bool`

HasVcdbRestart returns a boolean if a field has been set.

### GetMssqlFailoverDriveLetter

`func (o *BaseProvisionVDBParametersAllOf) GetMssqlFailoverDriveLetter() string`

GetMssqlFailoverDriveLetter returns the MssqlFailoverDriveLetter field if non-nil, zero value otherwise.

### GetMssqlFailoverDriveLetterOk

`func (o *BaseProvisionVDBParametersAllOf) GetMssqlFailoverDriveLetterOk() (*string, bool)`

GetMssqlFailoverDriveLetterOk returns a tuple with the MssqlFailoverDriveLetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlFailoverDriveLetter

`func (o *BaseProvisionVDBParametersAllOf) SetMssqlFailoverDriveLetter(v string)`

SetMssqlFailoverDriveLetter sets MssqlFailoverDriveLetter field to given value.

### HasMssqlFailoverDriveLetter

`func (o *BaseProvisionVDBParametersAllOf) HasMssqlFailoverDriveLetter() bool`

HasMssqlFailoverDriveLetter returns a boolean if a field has been set.

### GetTags

`func (o *BaseProvisionVDBParametersAllOf) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BaseProvisionVDBParametersAllOf) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BaseProvisionVDBParametersAllOf) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BaseProvisionVDBParametersAllOf) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


