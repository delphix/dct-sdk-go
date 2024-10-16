/*
Delphix DCT API

Delphix DCT API

API version: 3.17.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the UpdateVDBParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVDBParameters{}

// UpdateVDBParameters Parameters to update a VDB.
type UpdateVDBParameters struct {
	// The unique name of the VDB within a group.
	Name *string `json:"name,omitempty"`
	// The username of the database user (Oracle, ASE Only).
	DbUsername *string `json:"db_username,omitempty"`
	// The password of the database user (Oracle, ASE Only).
	DbPassword *string `json:"db_password,omitempty"`
	// Whether db_username and db_password must be validated, if present, against the VDB. This must be set to false when credentials validation is not possible, for instance if the VDB is known to be disabled.
	ValidateDbCredentials *bool `json:"validate_db_credentials,omitempty"`
	// Whether to enable VDB restart.
	AutoRestart *bool `json:"auto_restart,omitempty"`
	// The environment user ID to use to connect to the target environment.
	EnvironmentUserId *string `json:"environment_user_id,omitempty"`
	// The ID of the target VDB Template (Oracle and MSSql Only).
	TemplateId *string `json:"template_id,omitempty"`
	// The listener IDs for this provision operation (Oracle Only).
	ListenerIds []string `json:"listener_ids,omitempty"`
	// Whether to enable new DBID for Oracle
	NewDbid *bool `json:"new_dbid,omitempty"`
	// Whether to enable CDC on provision for MSSql
	CdcOnProvision *bool `json:"cdc_on_provision,omitempty"`
	// Pre script for MSSql.
	PreScript *string `json:"pre_script,omitempty"`
	// Post script for MSSql.
	PostScript *string `json:"post_script,omitempty"`
	Hooks *VirtualDatasetHooks `json:"hooks,omitempty"`
	// Environment variable to be set when the engine administers a VDB. See the Engine documentation for the list of allowed/denied environment variables and rules about substitution. Custom environment variables can only be updated while the VDB is disabled.
	CustomEnvVars *map[string]string `json:"custom_env_vars,omitempty"`
	// Environment files to be sourced when the Engine administers a VDB. This path can be followed by parameters. Paths and parameters are separated by spaces. Custom environment variables can only be updated while the VDB is disabled.
	CustomEnvFiles []string `json:"custom_env_files,omitempty"`
	// Environment files to be sourced when the Engine administers an Oracle RAC VDB. This path can be followed by parameters. Paths and parameters are separated by spaces. Custom environment variables can only be updated while the VDB is disabled.
	OracleRacCustomEnvFiles []OracleRacCustomEnvFile `json:"oracle_rac_custom_env_files,omitempty"`
	// Environment variable to be set when the engine administers an Oracle RAC VDB. See the Engine documentation for the list of allowed/denied environment variables and rules about substitution. Custom environment variables can only be updated while the VDB is disabled.
	OracleRacCustomEnvVars []OracleRacCustomEnvVar `json:"oracle_rac_custom_env_vars,omitempty"`
	// Path to a copy of the parent's Oracle transparent data encryption keystore on the target host. Required to provision from snapshots containing encrypted database files. (Oracle Multitenant Only)
	ParentTdeKeystorePath *string `json:"parent_tde_keystore_path,omitempty"`
	// The password of the keystore specified in parentTdeKeystorePath. (Oracle Multitenant Only)
	ParentTdeKeystorePassword *string `json:"parent_tde_keystore_password,omitempty"`
	// ID of the key created by Delphix. (Oracle Multitenant Only)
	TdeKeyIdentifier *string `json:"tde_key_identifier,omitempty"`
	// Path to the keystore of the target vCDB. (Oracle Multitenant Only)
	TargetVcdbTdeKeystorePath *string `json:"target_vcdb_tde_keystore_path,omitempty"`
	// The password for the Transparent Data Encryption keystore associated with the CDB. (Oracle Multitenant Only)
	CdbTdeKeystorePassword *string `json:"cdb_tde_keystore_password,omitempty"`
	// The JSON payload conforming to the DraftV4 schema based on the type of application data being manipulated.
	AppdataSourceParams map[string]interface{} `json:"appdata_source_params,omitempty"`
	// Specifies additional locations on which to mount a subdirectory of an AppData container. Can only be updated while the VDB is disabled.
	AdditionalMountPoints []AdditionalMountPoint `json:"additional_mount_points,omitempty"`
	// The parameters specified by the source config schema in the toolkit
	AppdataConfigParams map[string]interface{} `json:"appdata_config_params,omitempty"`
	// Database configuration parameter overrides.
	ConfigParams map[string]interface{} `json:"config_params,omitempty"`
	// Mount point for the VDB (AppData only), can only be updated while the VDB is disabled.
	MountPoint *string `json:"mount_point,omitempty"`
	// List of jdbc connection strings which are used to connect with the database.
	OracleServices []string `json:"oracle_services,omitempty"`
}

// NewUpdateVDBParameters instantiates a new UpdateVDBParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVDBParameters() *UpdateVDBParameters {
	this := UpdateVDBParameters{}
	var validateDbCredentials bool = true
	this.ValidateDbCredentials = &validateDbCredentials
	return &this
}

// NewUpdateVDBParametersWithDefaults instantiates a new UpdateVDBParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVDBParametersWithDefaults() *UpdateVDBParameters {
	this := UpdateVDBParameters{}
	var validateDbCredentials bool = true
	this.ValidateDbCredentials = &validateDbCredentials
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateVDBParameters) SetName(v string) {
	o.Name = &v
}

// GetDbUsername returns the DbUsername field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetDbUsername() string {
	if o == nil || IsNil(o.DbUsername) {
		var ret string
		return ret
	}
	return *o.DbUsername
}

// GetDbUsernameOk returns a tuple with the DbUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetDbUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.DbUsername) {
		return nil, false
	}
	return o.DbUsername, true
}

// HasDbUsername returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasDbUsername() bool {
	if o != nil && !IsNil(o.DbUsername) {
		return true
	}

	return false
}

// SetDbUsername gets a reference to the given string and assigns it to the DbUsername field.
func (o *UpdateVDBParameters) SetDbUsername(v string) {
	o.DbUsername = &v
}

// GetDbPassword returns the DbPassword field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetDbPassword() string {
	if o == nil || IsNil(o.DbPassword) {
		var ret string
		return ret
	}
	return *o.DbPassword
}

// GetDbPasswordOk returns a tuple with the DbPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetDbPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.DbPassword) {
		return nil, false
	}
	return o.DbPassword, true
}

// HasDbPassword returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasDbPassword() bool {
	if o != nil && !IsNil(o.DbPassword) {
		return true
	}

	return false
}

// SetDbPassword gets a reference to the given string and assigns it to the DbPassword field.
func (o *UpdateVDBParameters) SetDbPassword(v string) {
	o.DbPassword = &v
}

// GetValidateDbCredentials returns the ValidateDbCredentials field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetValidateDbCredentials() bool {
	if o == nil || IsNil(o.ValidateDbCredentials) {
		var ret bool
		return ret
	}
	return *o.ValidateDbCredentials
}

// GetValidateDbCredentialsOk returns a tuple with the ValidateDbCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetValidateDbCredentialsOk() (*bool, bool) {
	if o == nil || IsNil(o.ValidateDbCredentials) {
		return nil, false
	}
	return o.ValidateDbCredentials, true
}

// HasValidateDbCredentials returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasValidateDbCredentials() bool {
	if o != nil && !IsNil(o.ValidateDbCredentials) {
		return true
	}

	return false
}

// SetValidateDbCredentials gets a reference to the given bool and assigns it to the ValidateDbCredentials field.
func (o *UpdateVDBParameters) SetValidateDbCredentials(v bool) {
	o.ValidateDbCredentials = &v
}

// GetAutoRestart returns the AutoRestart field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetAutoRestart() bool {
	if o == nil || IsNil(o.AutoRestart) {
		var ret bool
		return ret
	}
	return *o.AutoRestart
}

// GetAutoRestartOk returns a tuple with the AutoRestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetAutoRestartOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoRestart) {
		return nil, false
	}
	return o.AutoRestart, true
}

// HasAutoRestart returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasAutoRestart() bool {
	if o != nil && !IsNil(o.AutoRestart) {
		return true
	}

	return false
}

// SetAutoRestart gets a reference to the given bool and assigns it to the AutoRestart field.
func (o *UpdateVDBParameters) SetAutoRestart(v bool) {
	o.AutoRestart = &v
}

// GetEnvironmentUserId returns the EnvironmentUserId field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetEnvironmentUserId() string {
	if o == nil || IsNil(o.EnvironmentUserId) {
		var ret string
		return ret
	}
	return *o.EnvironmentUserId
}

// GetEnvironmentUserIdOk returns a tuple with the EnvironmentUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetEnvironmentUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentUserId) {
		return nil, false
	}
	return o.EnvironmentUserId, true
}

// HasEnvironmentUserId returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasEnvironmentUserId() bool {
	if o != nil && !IsNil(o.EnvironmentUserId) {
		return true
	}

	return false
}

// SetEnvironmentUserId gets a reference to the given string and assigns it to the EnvironmentUserId field.
func (o *UpdateVDBParameters) SetEnvironmentUserId(v string) {
	o.EnvironmentUserId = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *UpdateVDBParameters) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetListenerIds returns the ListenerIds field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetListenerIds() []string {
	if o == nil || IsNil(o.ListenerIds) {
		var ret []string
		return ret
	}
	return o.ListenerIds
}

// GetListenerIdsOk returns a tuple with the ListenerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetListenerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ListenerIds) {
		return nil, false
	}
	return o.ListenerIds, true
}

// HasListenerIds returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasListenerIds() bool {
	if o != nil && !IsNil(o.ListenerIds) {
		return true
	}

	return false
}

// SetListenerIds gets a reference to the given []string and assigns it to the ListenerIds field.
func (o *UpdateVDBParameters) SetListenerIds(v []string) {
	o.ListenerIds = v
}

// GetNewDbid returns the NewDbid field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetNewDbid() bool {
	if o == nil || IsNil(o.NewDbid) {
		var ret bool
		return ret
	}
	return *o.NewDbid
}

// GetNewDbidOk returns a tuple with the NewDbid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetNewDbidOk() (*bool, bool) {
	if o == nil || IsNil(o.NewDbid) {
		return nil, false
	}
	return o.NewDbid, true
}

// HasNewDbid returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasNewDbid() bool {
	if o != nil && !IsNil(o.NewDbid) {
		return true
	}

	return false
}

// SetNewDbid gets a reference to the given bool and assigns it to the NewDbid field.
func (o *UpdateVDBParameters) SetNewDbid(v bool) {
	o.NewDbid = &v
}

// GetCdcOnProvision returns the CdcOnProvision field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetCdcOnProvision() bool {
	if o == nil || IsNil(o.CdcOnProvision) {
		var ret bool
		return ret
	}
	return *o.CdcOnProvision
}

// GetCdcOnProvisionOk returns a tuple with the CdcOnProvision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetCdcOnProvisionOk() (*bool, bool) {
	if o == nil || IsNil(o.CdcOnProvision) {
		return nil, false
	}
	return o.CdcOnProvision, true
}

// HasCdcOnProvision returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasCdcOnProvision() bool {
	if o != nil && !IsNil(o.CdcOnProvision) {
		return true
	}

	return false
}

// SetCdcOnProvision gets a reference to the given bool and assigns it to the CdcOnProvision field.
func (o *UpdateVDBParameters) SetCdcOnProvision(v bool) {
	o.CdcOnProvision = &v
}

// GetPreScript returns the PreScript field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetPreScript() string {
	if o == nil || IsNil(o.PreScript) {
		var ret string
		return ret
	}
	return *o.PreScript
}

// GetPreScriptOk returns a tuple with the PreScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetPreScriptOk() (*string, bool) {
	if o == nil || IsNil(o.PreScript) {
		return nil, false
	}
	return o.PreScript, true
}

// HasPreScript returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasPreScript() bool {
	if o != nil && !IsNil(o.PreScript) {
		return true
	}

	return false
}

// SetPreScript gets a reference to the given string and assigns it to the PreScript field.
func (o *UpdateVDBParameters) SetPreScript(v string) {
	o.PreScript = &v
}

// GetPostScript returns the PostScript field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetPostScript() string {
	if o == nil || IsNil(o.PostScript) {
		var ret string
		return ret
	}
	return *o.PostScript
}

// GetPostScriptOk returns a tuple with the PostScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetPostScriptOk() (*string, bool) {
	if o == nil || IsNil(o.PostScript) {
		return nil, false
	}
	return o.PostScript, true
}

// HasPostScript returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasPostScript() bool {
	if o != nil && !IsNil(o.PostScript) {
		return true
	}

	return false
}

// SetPostScript gets a reference to the given string and assigns it to the PostScript field.
func (o *UpdateVDBParameters) SetPostScript(v string) {
	o.PostScript = &v
}

// GetHooks returns the Hooks field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetHooks() VirtualDatasetHooks {
	if o == nil || IsNil(o.Hooks) {
		var ret VirtualDatasetHooks
		return ret
	}
	return *o.Hooks
}

// GetHooksOk returns a tuple with the Hooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetHooksOk() (*VirtualDatasetHooks, bool) {
	if o == nil || IsNil(o.Hooks) {
		return nil, false
	}
	return o.Hooks, true
}

// HasHooks returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasHooks() bool {
	if o != nil && !IsNil(o.Hooks) {
		return true
	}

	return false
}

// SetHooks gets a reference to the given VirtualDatasetHooks and assigns it to the Hooks field.
func (o *UpdateVDBParameters) SetHooks(v VirtualDatasetHooks) {
	o.Hooks = &v
}

// GetCustomEnvVars returns the CustomEnvVars field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetCustomEnvVars() map[string]string {
	if o == nil || IsNil(o.CustomEnvVars) {
		var ret map[string]string
		return ret
	}
	return *o.CustomEnvVars
}

// GetCustomEnvVarsOk returns a tuple with the CustomEnvVars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetCustomEnvVarsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomEnvVars) {
		return nil, false
	}
	return o.CustomEnvVars, true
}

// HasCustomEnvVars returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasCustomEnvVars() bool {
	if o != nil && !IsNil(o.CustomEnvVars) {
		return true
	}

	return false
}

// SetCustomEnvVars gets a reference to the given map[string]string and assigns it to the CustomEnvVars field.
func (o *UpdateVDBParameters) SetCustomEnvVars(v map[string]string) {
	o.CustomEnvVars = &v
}

// GetCustomEnvFiles returns the CustomEnvFiles field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetCustomEnvFiles() []string {
	if o == nil || IsNil(o.CustomEnvFiles) {
		var ret []string
		return ret
	}
	return o.CustomEnvFiles
}

// GetCustomEnvFilesOk returns a tuple with the CustomEnvFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetCustomEnvFilesOk() ([]string, bool) {
	if o == nil || IsNil(o.CustomEnvFiles) {
		return nil, false
	}
	return o.CustomEnvFiles, true
}

// HasCustomEnvFiles returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasCustomEnvFiles() bool {
	if o != nil && !IsNil(o.CustomEnvFiles) {
		return true
	}

	return false
}

// SetCustomEnvFiles gets a reference to the given []string and assigns it to the CustomEnvFiles field.
func (o *UpdateVDBParameters) SetCustomEnvFiles(v []string) {
	o.CustomEnvFiles = v
}

// GetOracleRacCustomEnvFiles returns the OracleRacCustomEnvFiles field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetOracleRacCustomEnvFiles() []OracleRacCustomEnvFile {
	if o == nil || IsNil(o.OracleRacCustomEnvFiles) {
		var ret []OracleRacCustomEnvFile
		return ret
	}
	return o.OracleRacCustomEnvFiles
}

// GetOracleRacCustomEnvFilesOk returns a tuple with the OracleRacCustomEnvFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetOracleRacCustomEnvFilesOk() ([]OracleRacCustomEnvFile, bool) {
	if o == nil || IsNil(o.OracleRacCustomEnvFiles) {
		return nil, false
	}
	return o.OracleRacCustomEnvFiles, true
}

// HasOracleRacCustomEnvFiles returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasOracleRacCustomEnvFiles() bool {
	if o != nil && !IsNil(o.OracleRacCustomEnvFiles) {
		return true
	}

	return false
}

// SetOracleRacCustomEnvFiles gets a reference to the given []OracleRacCustomEnvFile and assigns it to the OracleRacCustomEnvFiles field.
func (o *UpdateVDBParameters) SetOracleRacCustomEnvFiles(v []OracleRacCustomEnvFile) {
	o.OracleRacCustomEnvFiles = v
}

// GetOracleRacCustomEnvVars returns the OracleRacCustomEnvVars field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetOracleRacCustomEnvVars() []OracleRacCustomEnvVar {
	if o == nil || IsNil(o.OracleRacCustomEnvVars) {
		var ret []OracleRacCustomEnvVar
		return ret
	}
	return o.OracleRacCustomEnvVars
}

// GetOracleRacCustomEnvVarsOk returns a tuple with the OracleRacCustomEnvVars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetOracleRacCustomEnvVarsOk() ([]OracleRacCustomEnvVar, bool) {
	if o == nil || IsNil(o.OracleRacCustomEnvVars) {
		return nil, false
	}
	return o.OracleRacCustomEnvVars, true
}

// HasOracleRacCustomEnvVars returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasOracleRacCustomEnvVars() bool {
	if o != nil && !IsNil(o.OracleRacCustomEnvVars) {
		return true
	}

	return false
}

// SetOracleRacCustomEnvVars gets a reference to the given []OracleRacCustomEnvVar and assigns it to the OracleRacCustomEnvVars field.
func (o *UpdateVDBParameters) SetOracleRacCustomEnvVars(v []OracleRacCustomEnvVar) {
	o.OracleRacCustomEnvVars = v
}

// GetParentTdeKeystorePath returns the ParentTdeKeystorePath field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetParentTdeKeystorePath() string {
	if o == nil || IsNil(o.ParentTdeKeystorePath) {
		var ret string
		return ret
	}
	return *o.ParentTdeKeystorePath
}

// GetParentTdeKeystorePathOk returns a tuple with the ParentTdeKeystorePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetParentTdeKeystorePathOk() (*string, bool) {
	if o == nil || IsNil(o.ParentTdeKeystorePath) {
		return nil, false
	}
	return o.ParentTdeKeystorePath, true
}

// HasParentTdeKeystorePath returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasParentTdeKeystorePath() bool {
	if o != nil && !IsNil(o.ParentTdeKeystorePath) {
		return true
	}

	return false
}

// SetParentTdeKeystorePath gets a reference to the given string and assigns it to the ParentTdeKeystorePath field.
func (o *UpdateVDBParameters) SetParentTdeKeystorePath(v string) {
	o.ParentTdeKeystorePath = &v
}

// GetParentTdeKeystorePassword returns the ParentTdeKeystorePassword field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetParentTdeKeystorePassword() string {
	if o == nil || IsNil(o.ParentTdeKeystorePassword) {
		var ret string
		return ret
	}
	return *o.ParentTdeKeystorePassword
}

// GetParentTdeKeystorePasswordOk returns a tuple with the ParentTdeKeystorePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetParentTdeKeystorePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.ParentTdeKeystorePassword) {
		return nil, false
	}
	return o.ParentTdeKeystorePassword, true
}

// HasParentTdeKeystorePassword returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasParentTdeKeystorePassword() bool {
	if o != nil && !IsNil(o.ParentTdeKeystorePassword) {
		return true
	}

	return false
}

// SetParentTdeKeystorePassword gets a reference to the given string and assigns it to the ParentTdeKeystorePassword field.
func (o *UpdateVDBParameters) SetParentTdeKeystorePassword(v string) {
	o.ParentTdeKeystorePassword = &v
}

// GetTdeKeyIdentifier returns the TdeKeyIdentifier field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetTdeKeyIdentifier() string {
	if o == nil || IsNil(o.TdeKeyIdentifier) {
		var ret string
		return ret
	}
	return *o.TdeKeyIdentifier
}

// GetTdeKeyIdentifierOk returns a tuple with the TdeKeyIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetTdeKeyIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.TdeKeyIdentifier) {
		return nil, false
	}
	return o.TdeKeyIdentifier, true
}

// HasTdeKeyIdentifier returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasTdeKeyIdentifier() bool {
	if o != nil && !IsNil(o.TdeKeyIdentifier) {
		return true
	}

	return false
}

// SetTdeKeyIdentifier gets a reference to the given string and assigns it to the TdeKeyIdentifier field.
func (o *UpdateVDBParameters) SetTdeKeyIdentifier(v string) {
	o.TdeKeyIdentifier = &v
}

// GetTargetVcdbTdeKeystorePath returns the TargetVcdbTdeKeystorePath field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetTargetVcdbTdeKeystorePath() string {
	if o == nil || IsNil(o.TargetVcdbTdeKeystorePath) {
		var ret string
		return ret
	}
	return *o.TargetVcdbTdeKeystorePath
}

// GetTargetVcdbTdeKeystorePathOk returns a tuple with the TargetVcdbTdeKeystorePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetTargetVcdbTdeKeystorePathOk() (*string, bool) {
	if o == nil || IsNil(o.TargetVcdbTdeKeystorePath) {
		return nil, false
	}
	return o.TargetVcdbTdeKeystorePath, true
}

// HasTargetVcdbTdeKeystorePath returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasTargetVcdbTdeKeystorePath() bool {
	if o != nil && !IsNil(o.TargetVcdbTdeKeystorePath) {
		return true
	}

	return false
}

// SetTargetVcdbTdeKeystorePath gets a reference to the given string and assigns it to the TargetVcdbTdeKeystorePath field.
func (o *UpdateVDBParameters) SetTargetVcdbTdeKeystorePath(v string) {
	o.TargetVcdbTdeKeystorePath = &v
}

// GetCdbTdeKeystorePassword returns the CdbTdeKeystorePassword field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetCdbTdeKeystorePassword() string {
	if o == nil || IsNil(o.CdbTdeKeystorePassword) {
		var ret string
		return ret
	}
	return *o.CdbTdeKeystorePassword
}

// GetCdbTdeKeystorePasswordOk returns a tuple with the CdbTdeKeystorePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetCdbTdeKeystorePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.CdbTdeKeystorePassword) {
		return nil, false
	}
	return o.CdbTdeKeystorePassword, true
}

// HasCdbTdeKeystorePassword returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasCdbTdeKeystorePassword() bool {
	if o != nil && !IsNil(o.CdbTdeKeystorePassword) {
		return true
	}

	return false
}

// SetCdbTdeKeystorePassword gets a reference to the given string and assigns it to the CdbTdeKeystorePassword field.
func (o *UpdateVDBParameters) SetCdbTdeKeystorePassword(v string) {
	o.CdbTdeKeystorePassword = &v
}

// GetAppdataSourceParams returns the AppdataSourceParams field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetAppdataSourceParams() map[string]interface{} {
	if o == nil || IsNil(o.AppdataSourceParams) {
		var ret map[string]interface{}
		return ret
	}
	return o.AppdataSourceParams
}

// GetAppdataSourceParamsOk returns a tuple with the AppdataSourceParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetAppdataSourceParamsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AppdataSourceParams) {
		return map[string]interface{}{}, false
	}
	return o.AppdataSourceParams, true
}

// HasAppdataSourceParams returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasAppdataSourceParams() bool {
	if o != nil && !IsNil(o.AppdataSourceParams) {
		return true
	}

	return false
}

// SetAppdataSourceParams gets a reference to the given map[string]interface{} and assigns it to the AppdataSourceParams field.
func (o *UpdateVDBParameters) SetAppdataSourceParams(v map[string]interface{}) {
	o.AppdataSourceParams = v
}

// GetAdditionalMountPoints returns the AdditionalMountPoints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateVDBParameters) GetAdditionalMountPoints() []AdditionalMountPoint {
	if o == nil {
		var ret []AdditionalMountPoint
		return ret
	}
	return o.AdditionalMountPoints
}

// GetAdditionalMountPointsOk returns a tuple with the AdditionalMountPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateVDBParameters) GetAdditionalMountPointsOk() ([]AdditionalMountPoint, bool) {
	if o == nil || IsNil(o.AdditionalMountPoints) {
		return nil, false
	}
	return o.AdditionalMountPoints, true
}

// HasAdditionalMountPoints returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasAdditionalMountPoints() bool {
	if o != nil && !IsNil(o.AdditionalMountPoints) {
		return true
	}

	return false
}

// SetAdditionalMountPoints gets a reference to the given []AdditionalMountPoint and assigns it to the AdditionalMountPoints field.
func (o *UpdateVDBParameters) SetAdditionalMountPoints(v []AdditionalMountPoint) {
	o.AdditionalMountPoints = v
}

// GetAppdataConfigParams returns the AppdataConfigParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateVDBParameters) GetAppdataConfigParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AppdataConfigParams
}

// GetAppdataConfigParamsOk returns a tuple with the AppdataConfigParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateVDBParameters) GetAppdataConfigParamsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AppdataConfigParams) {
		return map[string]interface{}{}, false
	}
	return o.AppdataConfigParams, true
}

// HasAppdataConfigParams returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasAppdataConfigParams() bool {
	if o != nil && !IsNil(o.AppdataConfigParams) {
		return true
	}

	return false
}

// SetAppdataConfigParams gets a reference to the given map[string]interface{} and assigns it to the AppdataConfigParams field.
func (o *UpdateVDBParameters) SetAppdataConfigParams(v map[string]interface{}) {
	o.AppdataConfigParams = v
}

// GetConfigParams returns the ConfigParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateVDBParameters) GetConfigParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ConfigParams
}

// GetConfigParamsOk returns a tuple with the ConfigParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateVDBParameters) GetConfigParamsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ConfigParams) {
		return map[string]interface{}{}, false
	}
	return o.ConfigParams, true
}

// HasConfigParams returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasConfigParams() bool {
	if o != nil && !IsNil(o.ConfigParams) {
		return true
	}

	return false
}

// SetConfigParams gets a reference to the given map[string]interface{} and assigns it to the ConfigParams field.
func (o *UpdateVDBParameters) SetConfigParams(v map[string]interface{}) {
	o.ConfigParams = v
}

// GetMountPoint returns the MountPoint field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetMountPoint() string {
	if o == nil || IsNil(o.MountPoint) {
		var ret string
		return ret
	}
	return *o.MountPoint
}

// GetMountPointOk returns a tuple with the MountPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetMountPointOk() (*string, bool) {
	if o == nil || IsNil(o.MountPoint) {
		return nil, false
	}
	return o.MountPoint, true
}

// HasMountPoint returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasMountPoint() bool {
	if o != nil && !IsNil(o.MountPoint) {
		return true
	}

	return false
}

// SetMountPoint gets a reference to the given string and assigns it to the MountPoint field.
func (o *UpdateVDBParameters) SetMountPoint(v string) {
	o.MountPoint = &v
}

// GetOracleServices returns the OracleServices field value if set, zero value otherwise.
func (o *UpdateVDBParameters) GetOracleServices() []string {
	if o == nil || IsNil(o.OracleServices) {
		var ret []string
		return ret
	}
	return o.OracleServices
}

// GetOracleServicesOk returns a tuple with the OracleServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBParameters) GetOracleServicesOk() ([]string, bool) {
	if o == nil || IsNil(o.OracleServices) {
		return nil, false
	}
	return o.OracleServices, true
}

// HasOracleServices returns a boolean if a field has been set.
func (o *UpdateVDBParameters) HasOracleServices() bool {
	if o != nil && !IsNil(o.OracleServices) {
		return true
	}

	return false
}

// SetOracleServices gets a reference to the given []string and assigns it to the OracleServices field.
func (o *UpdateVDBParameters) SetOracleServices(v []string) {
	o.OracleServices = v
}

func (o UpdateVDBParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVDBParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DbUsername) {
		toSerialize["db_username"] = o.DbUsername
	}
	if !IsNil(o.DbPassword) {
		toSerialize["db_password"] = o.DbPassword
	}
	if !IsNil(o.ValidateDbCredentials) {
		toSerialize["validate_db_credentials"] = o.ValidateDbCredentials
	}
	if !IsNil(o.AutoRestart) {
		toSerialize["auto_restart"] = o.AutoRestart
	}
	if !IsNil(o.EnvironmentUserId) {
		toSerialize["environment_user_id"] = o.EnvironmentUserId
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.ListenerIds) {
		toSerialize["listener_ids"] = o.ListenerIds
	}
	if !IsNil(o.NewDbid) {
		toSerialize["new_dbid"] = o.NewDbid
	}
	if !IsNil(o.CdcOnProvision) {
		toSerialize["cdc_on_provision"] = o.CdcOnProvision
	}
	if !IsNil(o.PreScript) {
		toSerialize["pre_script"] = o.PreScript
	}
	if !IsNil(o.PostScript) {
		toSerialize["post_script"] = o.PostScript
	}
	if !IsNil(o.Hooks) {
		toSerialize["hooks"] = o.Hooks
	}
	if !IsNil(o.CustomEnvVars) {
		toSerialize["custom_env_vars"] = o.CustomEnvVars
	}
	if !IsNil(o.CustomEnvFiles) {
		toSerialize["custom_env_files"] = o.CustomEnvFiles
	}
	if !IsNil(o.OracleRacCustomEnvFiles) {
		toSerialize["oracle_rac_custom_env_files"] = o.OracleRacCustomEnvFiles
	}
	if !IsNil(o.OracleRacCustomEnvVars) {
		toSerialize["oracle_rac_custom_env_vars"] = o.OracleRacCustomEnvVars
	}
	if !IsNil(o.ParentTdeKeystorePath) {
		toSerialize["parent_tde_keystore_path"] = o.ParentTdeKeystorePath
	}
	if !IsNil(o.ParentTdeKeystorePassword) {
		toSerialize["parent_tde_keystore_password"] = o.ParentTdeKeystorePassword
	}
	if !IsNil(o.TdeKeyIdentifier) {
		toSerialize["tde_key_identifier"] = o.TdeKeyIdentifier
	}
	if !IsNil(o.TargetVcdbTdeKeystorePath) {
		toSerialize["target_vcdb_tde_keystore_path"] = o.TargetVcdbTdeKeystorePath
	}
	if !IsNil(o.CdbTdeKeystorePassword) {
		toSerialize["cdb_tde_keystore_password"] = o.CdbTdeKeystorePassword
	}
	if !IsNil(o.AppdataSourceParams) {
		toSerialize["appdata_source_params"] = o.AppdataSourceParams
	}
	if o.AdditionalMountPoints != nil {
		toSerialize["additional_mount_points"] = o.AdditionalMountPoints
	}
	if o.AppdataConfigParams != nil {
		toSerialize["appdata_config_params"] = o.AppdataConfigParams
	}
	if o.ConfigParams != nil {
		toSerialize["config_params"] = o.ConfigParams
	}
	if !IsNil(o.MountPoint) {
		toSerialize["mount_point"] = o.MountPoint
	}
	if !IsNil(o.OracleServices) {
		toSerialize["oracle_services"] = o.OracleServices
	}
	return toSerialize, nil
}

type NullableUpdateVDBParameters struct {
	value *UpdateVDBParameters
	isSet bool
}

func (v NullableUpdateVDBParameters) Get() *UpdateVDBParameters {
	return v.value
}

func (v *NullableUpdateVDBParameters) Set(val *UpdateVDBParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVDBParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVDBParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVDBParameters(val *UpdateVDBParameters) *NullableUpdateVDBParameters {
	return &NullableUpdateVDBParameters{value: val, isSet: true}
}

func (v NullableUpdateVDBParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVDBParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


