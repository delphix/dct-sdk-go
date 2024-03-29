/*
Delphix DCT API

Delphix DCT API

API version: 3.9.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the OracleStagingPushDSourceLinkSourceParametersAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OracleStagingPushDSourceLinkSourceParametersAllOf{}

// OracleStagingPushDSourceLinkSourceParametersAllOf struct for OracleStagingPushDSourceLinkSourceParametersAllOf
type OracleStagingPushDSourceLinkSourceParametersAllOf struct {
	// The ID of the engine to link staging push database on.
	EngineId *string `json:"engine_id,omitempty"`
	// The container type of this database.If not provided the request would be considered as a PDB database.
	ContainerType *string `json:"container_type,omitempty"`
	// Id of the environment user to use for linking.
	EnvironmentUserId *string `json:"environment_user_id,omitempty"`
	// The repository reference to link.
	Repository *string `json:"repository,omitempty"`
	// The name of the database.
	DatabaseName *string `json:"database_name,omitempty"`
	// The unique name of the database.
	DatabaseUniqueName *string `json:"database_unique_name,omitempty"`
	// The name (sid) of the instance.
	Sid *string `json:"sid,omitempty"`
	// The base mount point to use for the NFS mounts.
	MountBase *string `json:"mount_base,omitempty"`
	// An array of name value pair of environment variables.
	CustomEnvVariablesPairs []NameValuePair `json:"custom_env_variables_pairs,omitempty"`
	// An array of strings of whitespace-separated parameters to be passed to the source command. The first parameter must be an absolute path to a file that exists on the target environment. Every subsequent parameter will be treated as an argument interpreted by the environment file.
	CustomEnvVariablesPaths []string `json:"custom_env_variables_paths,omitempty"`
	// Boolean value indicates whether this staging database should automatically be restarted when staging host reboot is detected.
	AutoStagingRestart *bool `json:"auto_staging_restart,omitempty"`
	// Boolean value indicates whether this staging database will be configured as a physical standby.
	PhysicalStandby *bool `json:"physical_standby,omitempty"`
	// Boolean value indicates whether this staging database snapshot will be validated by opening it in read-only.
	ValidateSnapshotInReadonly *bool `json:"validate_snapshot_in_readonly,omitempty"`
	// An array of name value pair of Oracle database configuration parameter overrides. This property is deprecated. Use staging_database_config_params instead.
	// Deprecated
	StagingDatabaseTemplates []NameValuePair `json:"staging_database_templates,omitempty"`
	// Oracle database configuration parameter overrides. If both staging_database_templates and staging_database_config_params are specified, staging_database_config_params will be used.
	StagingDatabaseConfigParams map[string]string `json:"staging_database_config_params,omitempty"`
	// Reference of the CDB source config.
	StagingContainerDatabaseReference *string `json:"staging_container_database_reference,omitempty"`
	// Operations to perform after syncing a created dSource and before running the LogSync.
	OpsPreLogSync []SourceOperation `json:"ops_pre_log_sync,omitempty"`
}

// NewOracleStagingPushDSourceLinkSourceParametersAllOf instantiates a new OracleStagingPushDSourceLinkSourceParametersAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOracleStagingPushDSourceLinkSourceParametersAllOf() *OracleStagingPushDSourceLinkSourceParametersAllOf {
	this := OracleStagingPushDSourceLinkSourceParametersAllOf{}
	var autoStagingRestart bool = false
	this.AutoStagingRestart = &autoStagingRestart
	var physicalStandby bool = false
	this.PhysicalStandby = &physicalStandby
	var validateSnapshotInReadonly bool = false
	this.ValidateSnapshotInReadonly = &validateSnapshotInReadonly
	return &this
}

// NewOracleStagingPushDSourceLinkSourceParametersAllOfWithDefaults instantiates a new OracleStagingPushDSourceLinkSourceParametersAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOracleStagingPushDSourceLinkSourceParametersAllOfWithDefaults() *OracleStagingPushDSourceLinkSourceParametersAllOf {
	this := OracleStagingPushDSourceLinkSourceParametersAllOf{}
	var autoStagingRestart bool = false
	this.AutoStagingRestart = &autoStagingRestart
	var physicalStandby bool = false
	this.PhysicalStandby = &physicalStandby
	var validateSnapshotInReadonly bool = false
	this.ValidateSnapshotInReadonly = &validateSnapshotInReadonly
	return &this
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetEngineId() string {
	if o == nil || IsNil(o.EngineId) {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) SetEngineId(v string) {
	o.EngineId = &v
}

// GetContainerType returns the ContainerType field value if set, zero value otherwise.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetContainerType() string {
	if o == nil || IsNil(o.ContainerType) {
		var ret string
		return ret
	}
	return *o.ContainerType
}

// GetContainerTypeOk returns a tuple with the ContainerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetContainerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerType) {
		return nil, false
	}
	return o.ContainerType, true
}

// HasContainerType returns a boolean if a field has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) HasContainerType() bool {
	if o != nil && !IsNil(o.ContainerType) {
		return true
	}

	return false
}

// SetContainerType gets a reference to the given string and assigns it to the ContainerType field.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) SetContainerType(v string) {
	o.ContainerType = &v
}

// GetEnvironmentUserId returns the EnvironmentUserId field value if set, zero value otherwise.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetEnvironmentUserId() string {
	if o == nil || IsNil(o.EnvironmentUserId) {
		var ret string
		return ret
	}
	return *o.EnvironmentUserId
}

// GetEnvironmentUserIdOk returns a tuple with the EnvironmentUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetEnvironmentUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentUserId) {
		return nil, false
	}
	return o.EnvironmentUserId, true
}

// HasEnvironmentUserId returns a boolean if a field has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) HasEnvironmentUserId() bool {
	if o != nil && !IsNil(o.EnvironmentUserId) {
		return true
	}

	return false
}

// SetEnvironmentUserId gets a reference to the given string and assigns it to the EnvironmentUserId field.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) SetEnvironmentUserId(v string) {
	o.EnvironmentUserId = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) SetRepository(v string) {
	o.Repository = &v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetDatabaseName() string {
	if o == nil || IsNil(o.DatabaseName) {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetDatabaseNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseName) {
		return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) HasDatabaseName() bool {
	if o != nil && !IsNil(o.DatabaseName) {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetDatabaseUniqueName returns the DatabaseUniqueName field value if set, zero value otherwise.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetDatabaseUniqueName() string {
	if o == nil || IsNil(o.DatabaseUniqueName) {
		var ret string
		return ret
	}
	return *o.DatabaseUniqueName
}

// GetDatabaseUniqueNameOk returns a tuple with the DatabaseUniqueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetDatabaseUniqueNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseUniqueName) {
		return nil, false
	}
	return o.DatabaseUniqueName, true
}

// HasDatabaseUniqueName returns a boolean if a field has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) HasDatabaseUniqueName() bool {
	if o != nil && !IsNil(o.DatabaseUniqueName) {
		return true
	}

	return false
}

// SetDatabaseUniqueName gets a reference to the given string and assigns it to the DatabaseUniqueName field.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) SetDatabaseUniqueName(v string) {
	o.DatabaseUniqueName = &v
}

// GetSid returns the Sid field value if set, zero value otherwise.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetSid() string {
	if o == nil || IsNil(o.Sid) {
		var ret string
		return ret
	}
	return *o.Sid
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetSidOk() (*string, bool) {
	if o == nil || IsNil(o.Sid) {
		return nil, false
	}
	return o.Sid, true
}

// HasSid returns a boolean if a field has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) HasSid() bool {
	if o != nil && !IsNil(o.Sid) {
		return true
	}

	return false
}

// SetSid gets a reference to the given string and assigns it to the Sid field.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) SetSid(v string) {
	o.Sid = &v
}

// GetMountBase returns the MountBase field value if set, zero value otherwise.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetMountBase() string {
	if o == nil || IsNil(o.MountBase) {
		var ret string
		return ret
	}
	return *o.MountBase
}

// GetMountBaseOk returns a tuple with the MountBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetMountBaseOk() (*string, bool) {
	if o == nil || IsNil(o.MountBase) {
		return nil, false
	}
	return o.MountBase, true
}

// HasMountBase returns a boolean if a field has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) HasMountBase() bool {
	if o != nil && !IsNil(o.MountBase) {
		return true
	}

	return false
}

// SetMountBase gets a reference to the given string and assigns it to the MountBase field.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) SetMountBase(v string) {
	o.MountBase = &v
}

// GetCustomEnvVariablesPairs returns the CustomEnvVariablesPairs field value if set, zero value otherwise.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetCustomEnvVariablesPairs() []NameValuePair {
	if o == nil || IsNil(o.CustomEnvVariablesPairs) {
		var ret []NameValuePair
		return ret
	}
	return o.CustomEnvVariablesPairs
}

// GetCustomEnvVariablesPairsOk returns a tuple with the CustomEnvVariablesPairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetCustomEnvVariablesPairsOk() ([]NameValuePair, bool) {
	if o == nil || IsNil(o.CustomEnvVariablesPairs) {
		return nil, false
	}
	return o.CustomEnvVariablesPairs, true
}

// HasCustomEnvVariablesPairs returns a boolean if a field has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) HasCustomEnvVariablesPairs() bool {
	if o != nil && !IsNil(o.CustomEnvVariablesPairs) {
		return true
	}

	return false
}

// SetCustomEnvVariablesPairs gets a reference to the given []NameValuePair and assigns it to the CustomEnvVariablesPairs field.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) SetCustomEnvVariablesPairs(v []NameValuePair) {
	o.CustomEnvVariablesPairs = v
}

// GetCustomEnvVariablesPaths returns the CustomEnvVariablesPaths field value if set, zero value otherwise.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetCustomEnvVariablesPaths() []string {
	if o == nil || IsNil(o.CustomEnvVariablesPaths) {
		var ret []string
		return ret
	}
	return o.CustomEnvVariablesPaths
}

// GetCustomEnvVariablesPathsOk returns a tuple with the CustomEnvVariablesPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetCustomEnvVariablesPathsOk() ([]string, bool) {
	if o == nil || IsNil(o.CustomEnvVariablesPaths) {
		return nil, false
	}
	return o.CustomEnvVariablesPaths, true
}

// HasCustomEnvVariablesPaths returns a boolean if a field has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) HasCustomEnvVariablesPaths() bool {
	if o != nil && !IsNil(o.CustomEnvVariablesPaths) {
		return true
	}

	return false
}

// SetCustomEnvVariablesPaths gets a reference to the given []string and assigns it to the CustomEnvVariablesPaths field.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) SetCustomEnvVariablesPaths(v []string) {
	o.CustomEnvVariablesPaths = v
}

// GetAutoStagingRestart returns the AutoStagingRestart field value if set, zero value otherwise.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetAutoStagingRestart() bool {
	if o == nil || IsNil(o.AutoStagingRestart) {
		var ret bool
		return ret
	}
	return *o.AutoStagingRestart
}

// GetAutoStagingRestartOk returns a tuple with the AutoStagingRestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetAutoStagingRestartOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoStagingRestart) {
		return nil, false
	}
	return o.AutoStagingRestart, true
}

// HasAutoStagingRestart returns a boolean if a field has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) HasAutoStagingRestart() bool {
	if o != nil && !IsNil(o.AutoStagingRestart) {
		return true
	}

	return false
}

// SetAutoStagingRestart gets a reference to the given bool and assigns it to the AutoStagingRestart field.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) SetAutoStagingRestart(v bool) {
	o.AutoStagingRestart = &v
}

// GetPhysicalStandby returns the PhysicalStandby field value if set, zero value otherwise.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetPhysicalStandby() bool {
	if o == nil || IsNil(o.PhysicalStandby) {
		var ret bool
		return ret
	}
	return *o.PhysicalStandby
}

// GetPhysicalStandbyOk returns a tuple with the PhysicalStandby field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetPhysicalStandbyOk() (*bool, bool) {
	if o == nil || IsNil(o.PhysicalStandby) {
		return nil, false
	}
	return o.PhysicalStandby, true
}

// HasPhysicalStandby returns a boolean if a field has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) HasPhysicalStandby() bool {
	if o != nil && !IsNil(o.PhysicalStandby) {
		return true
	}

	return false
}

// SetPhysicalStandby gets a reference to the given bool and assigns it to the PhysicalStandby field.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) SetPhysicalStandby(v bool) {
	o.PhysicalStandby = &v
}

// GetValidateSnapshotInReadonly returns the ValidateSnapshotInReadonly field value if set, zero value otherwise.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetValidateSnapshotInReadonly() bool {
	if o == nil || IsNil(o.ValidateSnapshotInReadonly) {
		var ret bool
		return ret
	}
	return *o.ValidateSnapshotInReadonly
}

// GetValidateSnapshotInReadonlyOk returns a tuple with the ValidateSnapshotInReadonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetValidateSnapshotInReadonlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ValidateSnapshotInReadonly) {
		return nil, false
	}
	return o.ValidateSnapshotInReadonly, true
}

// HasValidateSnapshotInReadonly returns a boolean if a field has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) HasValidateSnapshotInReadonly() bool {
	if o != nil && !IsNil(o.ValidateSnapshotInReadonly) {
		return true
	}

	return false
}

// SetValidateSnapshotInReadonly gets a reference to the given bool and assigns it to the ValidateSnapshotInReadonly field.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) SetValidateSnapshotInReadonly(v bool) {
	o.ValidateSnapshotInReadonly = &v
}

// GetStagingDatabaseTemplates returns the StagingDatabaseTemplates field value if set, zero value otherwise.
// Deprecated
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetStagingDatabaseTemplates() []NameValuePair {
	if o == nil || IsNil(o.StagingDatabaseTemplates) {
		var ret []NameValuePair
		return ret
	}
	return o.StagingDatabaseTemplates
}

// GetStagingDatabaseTemplatesOk returns a tuple with the StagingDatabaseTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetStagingDatabaseTemplatesOk() ([]NameValuePair, bool) {
	if o == nil || IsNil(o.StagingDatabaseTemplates) {
		return nil, false
	}
	return o.StagingDatabaseTemplates, true
}

// HasStagingDatabaseTemplates returns a boolean if a field has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) HasStagingDatabaseTemplates() bool {
	if o != nil && !IsNil(o.StagingDatabaseTemplates) {
		return true
	}

	return false
}

// SetStagingDatabaseTemplates gets a reference to the given []NameValuePair and assigns it to the StagingDatabaseTemplates field.
// Deprecated
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) SetStagingDatabaseTemplates(v []NameValuePair) {
	o.StagingDatabaseTemplates = v
}

// GetStagingDatabaseConfigParams returns the StagingDatabaseConfigParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetStagingDatabaseConfigParams() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.StagingDatabaseConfigParams
}

// GetStagingDatabaseConfigParamsOk returns a tuple with the StagingDatabaseConfigParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetStagingDatabaseConfigParamsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.StagingDatabaseConfigParams) {
		return nil, false
	}
	return &o.StagingDatabaseConfigParams, true
}

// HasStagingDatabaseConfigParams returns a boolean if a field has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) HasStagingDatabaseConfigParams() bool {
	if o != nil && IsNil(o.StagingDatabaseConfigParams) {
		return true
	}

	return false
}

// SetStagingDatabaseConfigParams gets a reference to the given map[string]string and assigns it to the StagingDatabaseConfigParams field.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) SetStagingDatabaseConfigParams(v map[string]string) {
	o.StagingDatabaseConfigParams = v
}

// GetStagingContainerDatabaseReference returns the StagingContainerDatabaseReference field value if set, zero value otherwise.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetStagingContainerDatabaseReference() string {
	if o == nil || IsNil(o.StagingContainerDatabaseReference) {
		var ret string
		return ret
	}
	return *o.StagingContainerDatabaseReference
}

// GetStagingContainerDatabaseReferenceOk returns a tuple with the StagingContainerDatabaseReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetStagingContainerDatabaseReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.StagingContainerDatabaseReference) {
		return nil, false
	}
	return o.StagingContainerDatabaseReference, true
}

// HasStagingContainerDatabaseReference returns a boolean if a field has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) HasStagingContainerDatabaseReference() bool {
	if o != nil && !IsNil(o.StagingContainerDatabaseReference) {
		return true
	}

	return false
}

// SetStagingContainerDatabaseReference gets a reference to the given string and assigns it to the StagingContainerDatabaseReference field.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) SetStagingContainerDatabaseReference(v string) {
	o.StagingContainerDatabaseReference = &v
}

// GetOpsPreLogSync returns the OpsPreLogSync field value if set, zero value otherwise.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetOpsPreLogSync() []SourceOperation {
	if o == nil || IsNil(o.OpsPreLogSync) {
		var ret []SourceOperation
		return ret
	}
	return o.OpsPreLogSync
}

// GetOpsPreLogSyncOk returns a tuple with the OpsPreLogSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) GetOpsPreLogSyncOk() ([]SourceOperation, bool) {
	if o == nil || IsNil(o.OpsPreLogSync) {
		return nil, false
	}
	return o.OpsPreLogSync, true
}

// HasOpsPreLogSync returns a boolean if a field has been set.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) HasOpsPreLogSync() bool {
	if o != nil && !IsNil(o.OpsPreLogSync) {
		return true
	}

	return false
}

// SetOpsPreLogSync gets a reference to the given []SourceOperation and assigns it to the OpsPreLogSync field.
func (o *OracleStagingPushDSourceLinkSourceParametersAllOf) SetOpsPreLogSync(v []SourceOperation) {
	o.OpsPreLogSync = v
}

func (o OracleStagingPushDSourceLinkSourceParametersAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OracleStagingPushDSourceLinkSourceParametersAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EngineId) {
		toSerialize["engine_id"] = o.EngineId
	}
	if !IsNil(o.ContainerType) {
		toSerialize["container_type"] = o.ContainerType
	}
	if !IsNil(o.EnvironmentUserId) {
		toSerialize["environment_user_id"] = o.EnvironmentUserId
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.DatabaseName) {
		toSerialize["database_name"] = o.DatabaseName
	}
	if !IsNil(o.DatabaseUniqueName) {
		toSerialize["database_unique_name"] = o.DatabaseUniqueName
	}
	if !IsNil(o.Sid) {
		toSerialize["sid"] = o.Sid
	}
	if !IsNil(o.MountBase) {
		toSerialize["mount_base"] = o.MountBase
	}
	if !IsNil(o.CustomEnvVariablesPairs) {
		toSerialize["custom_env_variables_pairs"] = o.CustomEnvVariablesPairs
	}
	if !IsNil(o.CustomEnvVariablesPaths) {
		toSerialize["custom_env_variables_paths"] = o.CustomEnvVariablesPaths
	}
	if !IsNil(o.AutoStagingRestart) {
		toSerialize["auto_staging_restart"] = o.AutoStagingRestart
	}
	if !IsNil(o.PhysicalStandby) {
		toSerialize["physical_standby"] = o.PhysicalStandby
	}
	if !IsNil(o.ValidateSnapshotInReadonly) {
		toSerialize["validate_snapshot_in_readonly"] = o.ValidateSnapshotInReadonly
	}
	if !IsNil(o.StagingDatabaseTemplates) {
		toSerialize["staging_database_templates"] = o.StagingDatabaseTemplates
	}
	if o.StagingDatabaseConfigParams != nil {
		toSerialize["staging_database_config_params"] = o.StagingDatabaseConfigParams
	}
	if !IsNil(o.StagingContainerDatabaseReference) {
		toSerialize["staging_container_database_reference"] = o.StagingContainerDatabaseReference
	}
	if !IsNil(o.OpsPreLogSync) {
		toSerialize["ops_pre_log_sync"] = o.OpsPreLogSync
	}
	return toSerialize, nil
}

type NullableOracleStagingPushDSourceLinkSourceParametersAllOf struct {
	value *OracleStagingPushDSourceLinkSourceParametersAllOf
	isSet bool
}

func (v NullableOracleStagingPushDSourceLinkSourceParametersAllOf) Get() *OracleStagingPushDSourceLinkSourceParametersAllOf {
	return v.value
}

func (v *NullableOracleStagingPushDSourceLinkSourceParametersAllOf) Set(val *OracleStagingPushDSourceLinkSourceParametersAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOracleStagingPushDSourceLinkSourceParametersAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOracleStagingPushDSourceLinkSourceParametersAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOracleStagingPushDSourceLinkSourceParametersAllOf(val *OracleStagingPushDSourceLinkSourceParametersAllOf) *NullableOracleStagingPushDSourceLinkSourceParametersAllOf {
	return &NullableOracleStagingPushDSourceLinkSourceParametersAllOf{value: val, isSet: true}
}

func (v NullableOracleStagingPushDSourceLinkSourceParametersAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOracleStagingPushDSourceLinkSourceParametersAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


