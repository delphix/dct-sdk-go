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

// checks if the MSSQLDSourceStagingPushLinkSourceParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MSSQLDSourceStagingPushLinkSourceParameters{}

// MSSQLDSourceStagingPushLinkSourceParameters struct for MSSQLDSourceStagingPushLinkSourceParameters
type MSSQLDSourceStagingPushLinkSourceParameters struct {
	// Name of the dSource to be created.
	Name string `json:"name"`
	// Id of the source to link.
	SourceId *string `json:"source_id,omitempty"`
	// Id of the dataset group where this dSource should belong to.
	GroupId *string `json:"group_id,omitempty"`
	// The notes/description for the dSource.
	Description *string `json:"description,omitempty"`
	// True if LogSync should run for this database.
	LogSyncEnabled *bool `json:"log_sync_enabled,omitempty"`
	// The ID of the SnapSync policy for the dSource.
	SyncPolicyId *string `json:"sync_policy_id,omitempty"`
	// The ID of the Retention policy for the dSource.
	RetentionPolicyId *string `json:"retention_policy_id,omitempty"`
	// Whether the account creating this reporting schedule must be configured as owner of the reporting schedule.
	MakeCurrentAccountOwner *bool `json:"make_current_account_owner,omitempty"`
	// The tags to be created for dSource.
	Tags []Tag `json:"tags,omitempty"`
	// Operations to perform before syncing the created dSource. These operations can quiesce any data prior to syncing.
	OpsPreSync []SourceOperation `json:"ops_pre_sync,omitempty"`
	// Operations to perform after syncing a created dSource.
	OpsPostSync []SourceOperation `json:"ops_post_sync,omitempty"`
	// The ID of the engine to link staging push database on.
	EngineId string `json:"engine_id"`
	// The encryption key to use when restoring encrypted backups.
	EncryptionKey *string `json:"encryption_key,omitempty"`
	// Reference of the SQL instance on the PPT environment that we want to use for pre-provisioning.
	PptRepository string `json:"ppt_repository"`
	// Reference of the host OS user on the PPT host to use for linking.
	PptHostUser string `json:"ppt_host_user"`
	// A user-provided PowerShell script or executable to run prior to restoring from a backup during pre-provisioning.
	StagingPreScript *string `json:"staging_pre_script,omitempty"`
	// A user-provided PowerShell script or executable to run after restoring from a backup during pre-provisioning.
	StagingPostScript *string `json:"staging_post_script,omitempty"`
	// The name of the database to create on the staging environment. This property is mutually exclusive to sync_strategy_managed_type
	StagingDatabaseName string `json:"staging_database_name"`
}

// NewMSSQLDSourceStagingPushLinkSourceParameters instantiates a new MSSQLDSourceStagingPushLinkSourceParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMSSQLDSourceStagingPushLinkSourceParameters(name string, engineId string, pptRepository string, pptHostUser string, stagingDatabaseName string) *MSSQLDSourceStagingPushLinkSourceParameters {
	this := MSSQLDSourceStagingPushLinkSourceParameters{}
	this.Name = name
	var logSyncEnabled bool = false
	this.LogSyncEnabled = &logSyncEnabled
	var makeCurrentAccountOwner bool = true
	this.MakeCurrentAccountOwner = &makeCurrentAccountOwner
	this.EngineId = engineId
	this.PptRepository = pptRepository
	this.PptHostUser = pptHostUser
	this.StagingDatabaseName = stagingDatabaseName
	return &this
}

// NewMSSQLDSourceStagingPushLinkSourceParametersWithDefaults instantiates a new MSSQLDSourceStagingPushLinkSourceParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMSSQLDSourceStagingPushLinkSourceParametersWithDefaults() *MSSQLDSourceStagingPushLinkSourceParameters {
	this := MSSQLDSourceStagingPushLinkSourceParameters{}
	var logSyncEnabled bool = false
	this.LogSyncEnabled = &logSyncEnabled
	var makeCurrentAccountOwner bool = true
	this.MakeCurrentAccountOwner = &makeCurrentAccountOwner
	return &this
}

// GetName returns the Name field value
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetName(v string) {
	o.Name = v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetSourceId(v string) {
	o.SourceId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetGroupId(v string) {
	o.GroupId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetDescription(v string) {
	o.Description = &v
}

// GetLogSyncEnabled returns the LogSyncEnabled field value if set, zero value otherwise.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetLogSyncEnabled() bool {
	if o == nil || IsNil(o.LogSyncEnabled) {
		var ret bool
		return ret
	}
	return *o.LogSyncEnabled
}

// GetLogSyncEnabledOk returns a tuple with the LogSyncEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetLogSyncEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LogSyncEnabled) {
		return nil, false
	}
	return o.LogSyncEnabled, true
}

// HasLogSyncEnabled returns a boolean if a field has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) HasLogSyncEnabled() bool {
	if o != nil && !IsNil(o.LogSyncEnabled) {
		return true
	}

	return false
}

// SetLogSyncEnabled gets a reference to the given bool and assigns it to the LogSyncEnabled field.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetLogSyncEnabled(v bool) {
	o.LogSyncEnabled = &v
}

// GetSyncPolicyId returns the SyncPolicyId field value if set, zero value otherwise.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetSyncPolicyId() string {
	if o == nil || IsNil(o.SyncPolicyId) {
		var ret string
		return ret
	}
	return *o.SyncPolicyId
}

// GetSyncPolicyIdOk returns a tuple with the SyncPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetSyncPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.SyncPolicyId) {
		return nil, false
	}
	return o.SyncPolicyId, true
}

// HasSyncPolicyId returns a boolean if a field has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) HasSyncPolicyId() bool {
	if o != nil && !IsNil(o.SyncPolicyId) {
		return true
	}

	return false
}

// SetSyncPolicyId gets a reference to the given string and assigns it to the SyncPolicyId field.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetSyncPolicyId(v string) {
	o.SyncPolicyId = &v
}

// GetRetentionPolicyId returns the RetentionPolicyId field value if set, zero value otherwise.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetRetentionPolicyId() string {
	if o == nil || IsNil(o.RetentionPolicyId) {
		var ret string
		return ret
	}
	return *o.RetentionPolicyId
}

// GetRetentionPolicyIdOk returns a tuple with the RetentionPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetRetentionPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.RetentionPolicyId) {
		return nil, false
	}
	return o.RetentionPolicyId, true
}

// HasRetentionPolicyId returns a boolean if a field has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) HasRetentionPolicyId() bool {
	if o != nil && !IsNil(o.RetentionPolicyId) {
		return true
	}

	return false
}

// SetRetentionPolicyId gets a reference to the given string and assigns it to the RetentionPolicyId field.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetRetentionPolicyId(v string) {
	o.RetentionPolicyId = &v
}

// GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field value if set, zero value otherwise.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetMakeCurrentAccountOwner() bool {
	if o == nil || IsNil(o.MakeCurrentAccountOwner) {
		var ret bool
		return ret
	}
	return *o.MakeCurrentAccountOwner
}

// GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeCurrentAccountOwner) {
		return nil, false
	}
	return o.MakeCurrentAccountOwner, true
}

// HasMakeCurrentAccountOwner returns a boolean if a field has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) HasMakeCurrentAccountOwner() bool {
	if o != nil && !IsNil(o.MakeCurrentAccountOwner) {
		return true
	}

	return false
}

// SetMakeCurrentAccountOwner gets a reference to the given bool and assigns it to the MakeCurrentAccountOwner field.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetMakeCurrentAccountOwner(v bool) {
	o.MakeCurrentAccountOwner = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetTags(v []Tag) {
	o.Tags = v
}

// GetOpsPreSync returns the OpsPreSync field value if set, zero value otherwise.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetOpsPreSync() []SourceOperation {
	if o == nil || IsNil(o.OpsPreSync) {
		var ret []SourceOperation
		return ret
	}
	return o.OpsPreSync
}

// GetOpsPreSyncOk returns a tuple with the OpsPreSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetOpsPreSyncOk() ([]SourceOperation, bool) {
	if o == nil || IsNil(o.OpsPreSync) {
		return nil, false
	}
	return o.OpsPreSync, true
}

// HasOpsPreSync returns a boolean if a field has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) HasOpsPreSync() bool {
	if o != nil && !IsNil(o.OpsPreSync) {
		return true
	}

	return false
}

// SetOpsPreSync gets a reference to the given []SourceOperation and assigns it to the OpsPreSync field.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetOpsPreSync(v []SourceOperation) {
	o.OpsPreSync = v
}

// GetOpsPostSync returns the OpsPostSync field value if set, zero value otherwise.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetOpsPostSync() []SourceOperation {
	if o == nil || IsNil(o.OpsPostSync) {
		var ret []SourceOperation
		return ret
	}
	return o.OpsPostSync
}

// GetOpsPostSyncOk returns a tuple with the OpsPostSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetOpsPostSyncOk() ([]SourceOperation, bool) {
	if o == nil || IsNil(o.OpsPostSync) {
		return nil, false
	}
	return o.OpsPostSync, true
}

// HasOpsPostSync returns a boolean if a field has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) HasOpsPostSync() bool {
	if o != nil && !IsNil(o.OpsPostSync) {
		return true
	}

	return false
}

// SetOpsPostSync gets a reference to the given []SourceOperation and assigns it to the OpsPostSync field.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetOpsPostSync(v []SourceOperation) {
	o.OpsPostSync = v
}

// GetEngineId returns the EngineId field value
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetEngineId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetEngineIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineId, true
}

// SetEngineId sets field value
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetEngineId(v string) {
	o.EngineId = v
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetEncryptionKey() string {
	if o == nil || IsNil(o.EncryptionKey) {
		var ret string
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetEncryptionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionKey) {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) HasEncryptionKey() bool {
	if o != nil && !IsNil(o.EncryptionKey) {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetEncryptionKey(v string) {
	o.EncryptionKey = &v
}

// GetPptRepository returns the PptRepository field value
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetPptRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PptRepository
}

// GetPptRepositoryOk returns a tuple with the PptRepository field value
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetPptRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PptRepository, true
}

// SetPptRepository sets field value
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetPptRepository(v string) {
	o.PptRepository = v
}

// GetPptHostUser returns the PptHostUser field value
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetPptHostUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PptHostUser
}

// GetPptHostUserOk returns a tuple with the PptHostUser field value
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetPptHostUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PptHostUser, true
}

// SetPptHostUser sets field value
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetPptHostUser(v string) {
	o.PptHostUser = v
}

// GetStagingPreScript returns the StagingPreScript field value if set, zero value otherwise.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetStagingPreScript() string {
	if o == nil || IsNil(o.StagingPreScript) {
		var ret string
		return ret
	}
	return *o.StagingPreScript
}

// GetStagingPreScriptOk returns a tuple with the StagingPreScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetStagingPreScriptOk() (*string, bool) {
	if o == nil || IsNil(o.StagingPreScript) {
		return nil, false
	}
	return o.StagingPreScript, true
}

// HasStagingPreScript returns a boolean if a field has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) HasStagingPreScript() bool {
	if o != nil && !IsNil(o.StagingPreScript) {
		return true
	}

	return false
}

// SetStagingPreScript gets a reference to the given string and assigns it to the StagingPreScript field.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetStagingPreScript(v string) {
	o.StagingPreScript = &v
}

// GetStagingPostScript returns the StagingPostScript field value if set, zero value otherwise.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetStagingPostScript() string {
	if o == nil || IsNil(o.StagingPostScript) {
		var ret string
		return ret
	}
	return *o.StagingPostScript
}

// GetStagingPostScriptOk returns a tuple with the StagingPostScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetStagingPostScriptOk() (*string, bool) {
	if o == nil || IsNil(o.StagingPostScript) {
		return nil, false
	}
	return o.StagingPostScript, true
}

// HasStagingPostScript returns a boolean if a field has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) HasStagingPostScript() bool {
	if o != nil && !IsNil(o.StagingPostScript) {
		return true
	}

	return false
}

// SetStagingPostScript gets a reference to the given string and assigns it to the StagingPostScript field.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetStagingPostScript(v string) {
	o.StagingPostScript = &v
}

// GetStagingDatabaseName returns the StagingDatabaseName field value
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetStagingDatabaseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StagingDatabaseName
}

// GetStagingDatabaseNameOk returns a tuple with the StagingDatabaseName field value
// and a boolean to check if the value has been set.
func (o *MSSQLDSourceStagingPushLinkSourceParameters) GetStagingDatabaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StagingDatabaseName, true
}

// SetStagingDatabaseName sets field value
func (o *MSSQLDSourceStagingPushLinkSourceParameters) SetStagingDatabaseName(v string) {
	o.StagingDatabaseName = v
}

func (o MSSQLDSourceStagingPushLinkSourceParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MSSQLDSourceStagingPushLinkSourceParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.SourceId) {
		toSerialize["source_id"] = o.SourceId
	}
	if !IsNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.LogSyncEnabled) {
		toSerialize["log_sync_enabled"] = o.LogSyncEnabled
	}
	if !IsNil(o.SyncPolicyId) {
		toSerialize["sync_policy_id"] = o.SyncPolicyId
	}
	if !IsNil(o.RetentionPolicyId) {
		toSerialize["retention_policy_id"] = o.RetentionPolicyId
	}
	if !IsNil(o.MakeCurrentAccountOwner) {
		toSerialize["make_current_account_owner"] = o.MakeCurrentAccountOwner
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.OpsPreSync) {
		toSerialize["ops_pre_sync"] = o.OpsPreSync
	}
	if !IsNil(o.OpsPostSync) {
		toSerialize["ops_post_sync"] = o.OpsPostSync
	}
	toSerialize["engine_id"] = o.EngineId
	if !IsNil(o.EncryptionKey) {
		toSerialize["encryption_key"] = o.EncryptionKey
	}
	toSerialize["ppt_repository"] = o.PptRepository
	toSerialize["ppt_host_user"] = o.PptHostUser
	if !IsNil(o.StagingPreScript) {
		toSerialize["staging_pre_script"] = o.StagingPreScript
	}
	if !IsNil(o.StagingPostScript) {
		toSerialize["staging_post_script"] = o.StagingPostScript
	}
	toSerialize["staging_database_name"] = o.StagingDatabaseName
	return toSerialize, nil
}

type NullableMSSQLDSourceStagingPushLinkSourceParameters struct {
	value *MSSQLDSourceStagingPushLinkSourceParameters
	isSet bool
}

func (v NullableMSSQLDSourceStagingPushLinkSourceParameters) Get() *MSSQLDSourceStagingPushLinkSourceParameters {
	return v.value
}

func (v *NullableMSSQLDSourceStagingPushLinkSourceParameters) Set(val *MSSQLDSourceStagingPushLinkSourceParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableMSSQLDSourceStagingPushLinkSourceParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableMSSQLDSourceStagingPushLinkSourceParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMSSQLDSourceStagingPushLinkSourceParameters(val *MSSQLDSourceStagingPushLinkSourceParameters) *NullableMSSQLDSourceStagingPushLinkSourceParameters {
	return &NullableMSSQLDSourceStagingPushLinkSourceParameters{value: val, isSet: true}
}

func (v NullableMSSQLDSourceStagingPushLinkSourceParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMSSQLDSourceStagingPushLinkSourceParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

