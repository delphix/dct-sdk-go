/*
Delphix DCT API

Delphix DCT API

API version: 3.5.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the AppDataDSourceLinkSourceParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppDataDSourceLinkSourceParameters{}

// AppDataDSourceLinkSourceParameters struct for AppDataDSourceLinkSourceParameters
type AppDataDSourceLinkSourceParameters struct {
	// Name of the dSource to be created.
	Name *string `json:"name,omitempty"`
	// Id of the source to link.
	SourceId *string `json:"source_id,omitempty"`
	// Id of the dataset group where this dSource should belong to.
	GroupId *string `json:"group_id,omitempty"`
	// The notes/description for the dSource.
	Description *string `json:"description,omitempty"`
	// True if LogSync should run for this database.
	LogSyncEnabled *bool `json:"log_sync_enabled,omitempty"`
	// Whether the account creating this reporting schedule must be configured as owner of the reporting schedule.
	MakeCurrentAccountOwner *bool `json:"make_current_account_owner,omitempty"`
	// The tags to be created for dSource.
	Tags []Tag `json:"tags,omitempty"`
	// Operations to perform before syncing the created dSource. These operations can quiesce any data prior to syncing.
	OpsPreSync []SourceOperation `json:"ops_pre_sync,omitempty"`
	// Operations to perform after syncing a created dSource.
	OpsPostSync []SourceOperation `json:"ops_post_sync,omitempty"`
	// The type of link to create. Default is AppDataDirect.  * `AppDataDirect` - Represents the AppData specific parameters of a link request for a source directly replicated into the Delphix Engine. * `AppDataStaged` - Represents the AppData specific parameters of a link request for a source with a staging source. 
	LinkType *string `json:"link_type,omitempty"`
	// The base mount point for the NFS mount on the staging environment [AppDataStaged only].
	StagingMountBase *string `json:"staging_mount_base,omitempty"`
	// The environment used as an intermediate stage to pull data into Delphix [AppDataStaged only].
	StagingEnvironment *string `json:"staging_environment,omitempty"`
	// The environment user used to access the staging environment [AppDataStaged only].
	StagingEnvironmentUser *string `json:"staging_environment_user,omitempty"`
	// The OS user to use for linking.
	EnvironmentUser string `json:"environment_user"`
	// List of subdirectories in the source to exclude when syncing data. These paths are relative to the root of the source directory. [AppDataDirect only]
	Excludes []string `json:"excludes,omitempty"`
	// List of symlinks in the source to follow when syncing data. These paths are relative to the root of the source directory. All other symlinks are preserved. [AppDataDirect only]
	FollowSymlinks []string `json:"follow_symlinks,omitempty"`
	// The JSON payload conforming to the DraftV4 schema based on the type of application data being manipulated.
	Parameters map[string]interface{} `json:"parameters"`
	// The JSON payload conforming to the snapshot parameters definition in a LUA toolkit or platform plugin.
	SyncParameters map[string]interface{} `json:"sync_parameters"`
}

// NewAppDataDSourceLinkSourceParameters instantiates a new AppDataDSourceLinkSourceParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDataDSourceLinkSourceParameters(environmentUser string, parameters map[string]interface{}, syncParameters map[string]interface{}) *AppDataDSourceLinkSourceParameters {
	this := AppDataDSourceLinkSourceParameters{}
	var logSyncEnabled bool = false
	this.LogSyncEnabled = &logSyncEnabled
	var makeCurrentAccountOwner bool = true
	this.MakeCurrentAccountOwner = &makeCurrentAccountOwner
	var linkType string = "AppDataDirect"
	this.LinkType = &linkType
	this.EnvironmentUser = environmentUser
	this.Parameters = parameters
	this.SyncParameters = syncParameters
	return &this
}

// NewAppDataDSourceLinkSourceParametersWithDefaults instantiates a new AppDataDSourceLinkSourceParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDataDSourceLinkSourceParametersWithDefaults() *AppDataDSourceLinkSourceParameters {
	this := AppDataDSourceLinkSourceParameters{}
	var logSyncEnabled bool = false
	this.LogSyncEnabled = &logSyncEnabled
	var makeCurrentAccountOwner bool = true
	this.MakeCurrentAccountOwner = &makeCurrentAccountOwner
	var linkType string = "AppDataDirect"
	this.LinkType = &linkType
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppDataDSourceLinkSourceParameters) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppDataDSourceLinkSourceParameters) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppDataDSourceLinkSourceParameters) SetName(v string) {
	o.Name = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *AppDataDSourceLinkSourceParameters) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *AppDataDSourceLinkSourceParameters) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *AppDataDSourceLinkSourceParameters) SetSourceId(v string) {
	o.SourceId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *AppDataDSourceLinkSourceParameters) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *AppDataDSourceLinkSourceParameters) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *AppDataDSourceLinkSourceParameters) SetGroupId(v string) {
	o.GroupId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AppDataDSourceLinkSourceParameters) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AppDataDSourceLinkSourceParameters) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AppDataDSourceLinkSourceParameters) SetDescription(v string) {
	o.Description = &v
}

// GetLogSyncEnabled returns the LogSyncEnabled field value if set, zero value otherwise.
func (o *AppDataDSourceLinkSourceParameters) GetLogSyncEnabled() bool {
	if o == nil || IsNil(o.LogSyncEnabled) {
		var ret bool
		return ret
	}
	return *o.LogSyncEnabled
}

// GetLogSyncEnabledOk returns a tuple with the LogSyncEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetLogSyncEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LogSyncEnabled) {
		return nil, false
	}
	return o.LogSyncEnabled, true
}

// HasLogSyncEnabled returns a boolean if a field has been set.
func (o *AppDataDSourceLinkSourceParameters) HasLogSyncEnabled() bool {
	if o != nil && !IsNil(o.LogSyncEnabled) {
		return true
	}

	return false
}

// SetLogSyncEnabled gets a reference to the given bool and assigns it to the LogSyncEnabled field.
func (o *AppDataDSourceLinkSourceParameters) SetLogSyncEnabled(v bool) {
	o.LogSyncEnabled = &v
}

// GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field value if set, zero value otherwise.
func (o *AppDataDSourceLinkSourceParameters) GetMakeCurrentAccountOwner() bool {
	if o == nil || IsNil(o.MakeCurrentAccountOwner) {
		var ret bool
		return ret
	}
	return *o.MakeCurrentAccountOwner
}

// GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeCurrentAccountOwner) {
		return nil, false
	}
	return o.MakeCurrentAccountOwner, true
}

// HasMakeCurrentAccountOwner returns a boolean if a field has been set.
func (o *AppDataDSourceLinkSourceParameters) HasMakeCurrentAccountOwner() bool {
	if o != nil && !IsNil(o.MakeCurrentAccountOwner) {
		return true
	}

	return false
}

// SetMakeCurrentAccountOwner gets a reference to the given bool and assigns it to the MakeCurrentAccountOwner field.
func (o *AppDataDSourceLinkSourceParameters) SetMakeCurrentAccountOwner(v bool) {
	o.MakeCurrentAccountOwner = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AppDataDSourceLinkSourceParameters) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AppDataDSourceLinkSourceParameters) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *AppDataDSourceLinkSourceParameters) SetTags(v []Tag) {
	o.Tags = v
}

// GetOpsPreSync returns the OpsPreSync field value if set, zero value otherwise.
func (o *AppDataDSourceLinkSourceParameters) GetOpsPreSync() []SourceOperation {
	if o == nil || IsNil(o.OpsPreSync) {
		var ret []SourceOperation
		return ret
	}
	return o.OpsPreSync
}

// GetOpsPreSyncOk returns a tuple with the OpsPreSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetOpsPreSyncOk() ([]SourceOperation, bool) {
	if o == nil || IsNil(o.OpsPreSync) {
		return nil, false
	}
	return o.OpsPreSync, true
}

// HasOpsPreSync returns a boolean if a field has been set.
func (o *AppDataDSourceLinkSourceParameters) HasOpsPreSync() bool {
	if o != nil && !IsNil(o.OpsPreSync) {
		return true
	}

	return false
}

// SetOpsPreSync gets a reference to the given []SourceOperation and assigns it to the OpsPreSync field.
func (o *AppDataDSourceLinkSourceParameters) SetOpsPreSync(v []SourceOperation) {
	o.OpsPreSync = v
}

// GetOpsPostSync returns the OpsPostSync field value if set, zero value otherwise.
func (o *AppDataDSourceLinkSourceParameters) GetOpsPostSync() []SourceOperation {
	if o == nil || IsNil(o.OpsPostSync) {
		var ret []SourceOperation
		return ret
	}
	return o.OpsPostSync
}

// GetOpsPostSyncOk returns a tuple with the OpsPostSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetOpsPostSyncOk() ([]SourceOperation, bool) {
	if o == nil || IsNil(o.OpsPostSync) {
		return nil, false
	}
	return o.OpsPostSync, true
}

// HasOpsPostSync returns a boolean if a field has been set.
func (o *AppDataDSourceLinkSourceParameters) HasOpsPostSync() bool {
	if o != nil && !IsNil(o.OpsPostSync) {
		return true
	}

	return false
}

// SetOpsPostSync gets a reference to the given []SourceOperation and assigns it to the OpsPostSync field.
func (o *AppDataDSourceLinkSourceParameters) SetOpsPostSync(v []SourceOperation) {
	o.OpsPostSync = v
}

// GetLinkType returns the LinkType field value if set, zero value otherwise.
func (o *AppDataDSourceLinkSourceParameters) GetLinkType() string {
	if o == nil || IsNil(o.LinkType) {
		var ret string
		return ret
	}
	return *o.LinkType
}

// GetLinkTypeOk returns a tuple with the LinkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetLinkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LinkType) {
		return nil, false
	}
	return o.LinkType, true
}

// HasLinkType returns a boolean if a field has been set.
func (o *AppDataDSourceLinkSourceParameters) HasLinkType() bool {
	if o != nil && !IsNil(o.LinkType) {
		return true
	}

	return false
}

// SetLinkType gets a reference to the given string and assigns it to the LinkType field.
func (o *AppDataDSourceLinkSourceParameters) SetLinkType(v string) {
	o.LinkType = &v
}

// GetStagingMountBase returns the StagingMountBase field value if set, zero value otherwise.
func (o *AppDataDSourceLinkSourceParameters) GetStagingMountBase() string {
	if o == nil || IsNil(o.StagingMountBase) {
		var ret string
		return ret
	}
	return *o.StagingMountBase
}

// GetStagingMountBaseOk returns a tuple with the StagingMountBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetStagingMountBaseOk() (*string, bool) {
	if o == nil || IsNil(o.StagingMountBase) {
		return nil, false
	}
	return o.StagingMountBase, true
}

// HasStagingMountBase returns a boolean if a field has been set.
func (o *AppDataDSourceLinkSourceParameters) HasStagingMountBase() bool {
	if o != nil && !IsNil(o.StagingMountBase) {
		return true
	}

	return false
}

// SetStagingMountBase gets a reference to the given string and assigns it to the StagingMountBase field.
func (o *AppDataDSourceLinkSourceParameters) SetStagingMountBase(v string) {
	o.StagingMountBase = &v
}

// GetStagingEnvironment returns the StagingEnvironment field value if set, zero value otherwise.
func (o *AppDataDSourceLinkSourceParameters) GetStagingEnvironment() string {
	if o == nil || IsNil(o.StagingEnvironment) {
		var ret string
		return ret
	}
	return *o.StagingEnvironment
}

// GetStagingEnvironmentOk returns a tuple with the StagingEnvironment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetStagingEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.StagingEnvironment) {
		return nil, false
	}
	return o.StagingEnvironment, true
}

// HasStagingEnvironment returns a boolean if a field has been set.
func (o *AppDataDSourceLinkSourceParameters) HasStagingEnvironment() bool {
	if o != nil && !IsNil(o.StagingEnvironment) {
		return true
	}

	return false
}

// SetStagingEnvironment gets a reference to the given string and assigns it to the StagingEnvironment field.
func (o *AppDataDSourceLinkSourceParameters) SetStagingEnvironment(v string) {
	o.StagingEnvironment = &v
}

// GetStagingEnvironmentUser returns the StagingEnvironmentUser field value if set, zero value otherwise.
func (o *AppDataDSourceLinkSourceParameters) GetStagingEnvironmentUser() string {
	if o == nil || IsNil(o.StagingEnvironmentUser) {
		var ret string
		return ret
	}
	return *o.StagingEnvironmentUser
}

// GetStagingEnvironmentUserOk returns a tuple with the StagingEnvironmentUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetStagingEnvironmentUserOk() (*string, bool) {
	if o == nil || IsNil(o.StagingEnvironmentUser) {
		return nil, false
	}
	return o.StagingEnvironmentUser, true
}

// HasStagingEnvironmentUser returns a boolean if a field has been set.
func (o *AppDataDSourceLinkSourceParameters) HasStagingEnvironmentUser() bool {
	if o != nil && !IsNil(o.StagingEnvironmentUser) {
		return true
	}

	return false
}

// SetStagingEnvironmentUser gets a reference to the given string and assigns it to the StagingEnvironmentUser field.
func (o *AppDataDSourceLinkSourceParameters) SetStagingEnvironmentUser(v string) {
	o.StagingEnvironmentUser = &v
}

// GetEnvironmentUser returns the EnvironmentUser field value
func (o *AppDataDSourceLinkSourceParameters) GetEnvironmentUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentUser
}

// GetEnvironmentUserOk returns a tuple with the EnvironmentUser field value
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetEnvironmentUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentUser, true
}

// SetEnvironmentUser sets field value
func (o *AppDataDSourceLinkSourceParameters) SetEnvironmentUser(v string) {
	o.EnvironmentUser = v
}

// GetExcludes returns the Excludes field value if set, zero value otherwise.
func (o *AppDataDSourceLinkSourceParameters) GetExcludes() []string {
	if o == nil || IsNil(o.Excludes) {
		var ret []string
		return ret
	}
	return o.Excludes
}

// GetExcludesOk returns a tuple with the Excludes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetExcludesOk() ([]string, bool) {
	if o == nil || IsNil(o.Excludes) {
		return nil, false
	}
	return o.Excludes, true
}

// HasExcludes returns a boolean if a field has been set.
func (o *AppDataDSourceLinkSourceParameters) HasExcludes() bool {
	if o != nil && !IsNil(o.Excludes) {
		return true
	}

	return false
}

// SetExcludes gets a reference to the given []string and assigns it to the Excludes field.
func (o *AppDataDSourceLinkSourceParameters) SetExcludes(v []string) {
	o.Excludes = v
}

// GetFollowSymlinks returns the FollowSymlinks field value if set, zero value otherwise.
func (o *AppDataDSourceLinkSourceParameters) GetFollowSymlinks() []string {
	if o == nil || IsNil(o.FollowSymlinks) {
		var ret []string
		return ret
	}
	return o.FollowSymlinks
}

// GetFollowSymlinksOk returns a tuple with the FollowSymlinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetFollowSymlinksOk() ([]string, bool) {
	if o == nil || IsNil(o.FollowSymlinks) {
		return nil, false
	}
	return o.FollowSymlinks, true
}

// HasFollowSymlinks returns a boolean if a field has been set.
func (o *AppDataDSourceLinkSourceParameters) HasFollowSymlinks() bool {
	if o != nil && !IsNil(o.FollowSymlinks) {
		return true
	}

	return false
}

// SetFollowSymlinks gets a reference to the given []string and assigns it to the FollowSymlinks field.
func (o *AppDataDSourceLinkSourceParameters) SetFollowSymlinks(v []string) {
	o.FollowSymlinks = v
}

// GetParameters returns the Parameters field value
func (o *AppDataDSourceLinkSourceParameters) GetParameters() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// SetParameters sets field value
func (o *AppDataDSourceLinkSourceParameters) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

// GetSyncParameters returns the SyncParameters field value
func (o *AppDataDSourceLinkSourceParameters) GetSyncParameters() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.SyncParameters
}

// GetSyncParametersOk returns a tuple with the SyncParameters field value
// and a boolean to check if the value has been set.
func (o *AppDataDSourceLinkSourceParameters) GetSyncParametersOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.SyncParameters, true
}

// SetSyncParameters sets field value
func (o *AppDataDSourceLinkSourceParameters) SetSyncParameters(v map[string]interface{}) {
	o.SyncParameters = v
}

func (o AppDataDSourceLinkSourceParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppDataDSourceLinkSourceParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
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
	if !IsNil(o.LinkType) {
		toSerialize["link_type"] = o.LinkType
	}
	if !IsNil(o.StagingMountBase) {
		toSerialize["staging_mount_base"] = o.StagingMountBase
	}
	if !IsNil(o.StagingEnvironment) {
		toSerialize["staging_environment"] = o.StagingEnvironment
	}
	if !IsNil(o.StagingEnvironmentUser) {
		toSerialize["staging_environment_user"] = o.StagingEnvironmentUser
	}
	toSerialize["environment_user"] = o.EnvironmentUser
	if !IsNil(o.Excludes) {
		toSerialize["excludes"] = o.Excludes
	}
	if !IsNil(o.FollowSymlinks) {
		toSerialize["follow_symlinks"] = o.FollowSymlinks
	}
	toSerialize["parameters"] = o.Parameters
	toSerialize["sync_parameters"] = o.SyncParameters
	return toSerialize, nil
}

type NullableAppDataDSourceLinkSourceParameters struct {
	value *AppDataDSourceLinkSourceParameters
	isSet bool
}

func (v NullableAppDataDSourceLinkSourceParameters) Get() *AppDataDSourceLinkSourceParameters {
	return v.value
}

func (v *NullableAppDataDSourceLinkSourceParameters) Set(val *AppDataDSourceLinkSourceParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDataDSourceLinkSourceParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDataDSourceLinkSourceParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDataDSourceLinkSourceParameters(val *AppDataDSourceLinkSourceParameters) *NullableAppDataDSourceLinkSourceParameters {
	return &NullableAppDataDSourceLinkSourceParameters{value: val, isSet: true}
}

func (v NullableAppDataDSourceLinkSourceParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDataDSourceLinkSourceParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

