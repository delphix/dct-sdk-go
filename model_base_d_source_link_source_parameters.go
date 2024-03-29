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

// checks if the BaseDSourceLinkSourceParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseDSourceLinkSourceParameters{}

// BaseDSourceLinkSourceParameters struct for BaseDSourceLinkSourceParameters
type BaseDSourceLinkSourceParameters struct {
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
}

// NewBaseDSourceLinkSourceParameters instantiates a new BaseDSourceLinkSourceParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseDSourceLinkSourceParameters() *BaseDSourceLinkSourceParameters {
	this := BaseDSourceLinkSourceParameters{}
	var logSyncEnabled bool = false
	this.LogSyncEnabled = &logSyncEnabled
	var makeCurrentAccountOwner bool = true
	this.MakeCurrentAccountOwner = &makeCurrentAccountOwner
	return &this
}

// NewBaseDSourceLinkSourceParametersWithDefaults instantiates a new BaseDSourceLinkSourceParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseDSourceLinkSourceParametersWithDefaults() *BaseDSourceLinkSourceParameters {
	this := BaseDSourceLinkSourceParameters{}
	var logSyncEnabled bool = false
	this.LogSyncEnabled = &logSyncEnabled
	var makeCurrentAccountOwner bool = true
	this.MakeCurrentAccountOwner = &makeCurrentAccountOwner
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BaseDSourceLinkSourceParameters) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDSourceLinkSourceParameters) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseDSourceLinkSourceParameters) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseDSourceLinkSourceParameters) SetName(v string) {
	o.Name = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *BaseDSourceLinkSourceParameters) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDSourceLinkSourceParameters) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *BaseDSourceLinkSourceParameters) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *BaseDSourceLinkSourceParameters) SetSourceId(v string) {
	o.SourceId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *BaseDSourceLinkSourceParameters) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDSourceLinkSourceParameters) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *BaseDSourceLinkSourceParameters) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *BaseDSourceLinkSourceParameters) SetGroupId(v string) {
	o.GroupId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BaseDSourceLinkSourceParameters) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDSourceLinkSourceParameters) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseDSourceLinkSourceParameters) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseDSourceLinkSourceParameters) SetDescription(v string) {
	o.Description = &v
}

// GetLogSyncEnabled returns the LogSyncEnabled field value if set, zero value otherwise.
func (o *BaseDSourceLinkSourceParameters) GetLogSyncEnabled() bool {
	if o == nil || IsNil(o.LogSyncEnabled) {
		var ret bool
		return ret
	}
	return *o.LogSyncEnabled
}

// GetLogSyncEnabledOk returns a tuple with the LogSyncEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDSourceLinkSourceParameters) GetLogSyncEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LogSyncEnabled) {
		return nil, false
	}
	return o.LogSyncEnabled, true
}

// HasLogSyncEnabled returns a boolean if a field has been set.
func (o *BaseDSourceLinkSourceParameters) HasLogSyncEnabled() bool {
	if o != nil && !IsNil(o.LogSyncEnabled) {
		return true
	}

	return false
}

// SetLogSyncEnabled gets a reference to the given bool and assigns it to the LogSyncEnabled field.
func (o *BaseDSourceLinkSourceParameters) SetLogSyncEnabled(v bool) {
	o.LogSyncEnabled = &v
}

// GetSyncPolicyId returns the SyncPolicyId field value if set, zero value otherwise.
func (o *BaseDSourceLinkSourceParameters) GetSyncPolicyId() string {
	if o == nil || IsNil(o.SyncPolicyId) {
		var ret string
		return ret
	}
	return *o.SyncPolicyId
}

// GetSyncPolicyIdOk returns a tuple with the SyncPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDSourceLinkSourceParameters) GetSyncPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.SyncPolicyId) {
		return nil, false
	}
	return o.SyncPolicyId, true
}

// HasSyncPolicyId returns a boolean if a field has been set.
func (o *BaseDSourceLinkSourceParameters) HasSyncPolicyId() bool {
	if o != nil && !IsNil(o.SyncPolicyId) {
		return true
	}

	return false
}

// SetSyncPolicyId gets a reference to the given string and assigns it to the SyncPolicyId field.
func (o *BaseDSourceLinkSourceParameters) SetSyncPolicyId(v string) {
	o.SyncPolicyId = &v
}

// GetRetentionPolicyId returns the RetentionPolicyId field value if set, zero value otherwise.
func (o *BaseDSourceLinkSourceParameters) GetRetentionPolicyId() string {
	if o == nil || IsNil(o.RetentionPolicyId) {
		var ret string
		return ret
	}
	return *o.RetentionPolicyId
}

// GetRetentionPolicyIdOk returns a tuple with the RetentionPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDSourceLinkSourceParameters) GetRetentionPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.RetentionPolicyId) {
		return nil, false
	}
	return o.RetentionPolicyId, true
}

// HasRetentionPolicyId returns a boolean if a field has been set.
func (o *BaseDSourceLinkSourceParameters) HasRetentionPolicyId() bool {
	if o != nil && !IsNil(o.RetentionPolicyId) {
		return true
	}

	return false
}

// SetRetentionPolicyId gets a reference to the given string and assigns it to the RetentionPolicyId field.
func (o *BaseDSourceLinkSourceParameters) SetRetentionPolicyId(v string) {
	o.RetentionPolicyId = &v
}

// GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field value if set, zero value otherwise.
func (o *BaseDSourceLinkSourceParameters) GetMakeCurrentAccountOwner() bool {
	if o == nil || IsNil(o.MakeCurrentAccountOwner) {
		var ret bool
		return ret
	}
	return *o.MakeCurrentAccountOwner
}

// GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDSourceLinkSourceParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeCurrentAccountOwner) {
		return nil, false
	}
	return o.MakeCurrentAccountOwner, true
}

// HasMakeCurrentAccountOwner returns a boolean if a field has been set.
func (o *BaseDSourceLinkSourceParameters) HasMakeCurrentAccountOwner() bool {
	if o != nil && !IsNil(o.MakeCurrentAccountOwner) {
		return true
	}

	return false
}

// SetMakeCurrentAccountOwner gets a reference to the given bool and assigns it to the MakeCurrentAccountOwner field.
func (o *BaseDSourceLinkSourceParameters) SetMakeCurrentAccountOwner(v bool) {
	o.MakeCurrentAccountOwner = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BaseDSourceLinkSourceParameters) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDSourceLinkSourceParameters) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BaseDSourceLinkSourceParameters) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *BaseDSourceLinkSourceParameters) SetTags(v []Tag) {
	o.Tags = v
}

// GetOpsPreSync returns the OpsPreSync field value if set, zero value otherwise.
func (o *BaseDSourceLinkSourceParameters) GetOpsPreSync() []SourceOperation {
	if o == nil || IsNil(o.OpsPreSync) {
		var ret []SourceOperation
		return ret
	}
	return o.OpsPreSync
}

// GetOpsPreSyncOk returns a tuple with the OpsPreSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDSourceLinkSourceParameters) GetOpsPreSyncOk() ([]SourceOperation, bool) {
	if o == nil || IsNil(o.OpsPreSync) {
		return nil, false
	}
	return o.OpsPreSync, true
}

// HasOpsPreSync returns a boolean if a field has been set.
func (o *BaseDSourceLinkSourceParameters) HasOpsPreSync() bool {
	if o != nil && !IsNil(o.OpsPreSync) {
		return true
	}

	return false
}

// SetOpsPreSync gets a reference to the given []SourceOperation and assigns it to the OpsPreSync field.
func (o *BaseDSourceLinkSourceParameters) SetOpsPreSync(v []SourceOperation) {
	o.OpsPreSync = v
}

// GetOpsPostSync returns the OpsPostSync field value if set, zero value otherwise.
func (o *BaseDSourceLinkSourceParameters) GetOpsPostSync() []SourceOperation {
	if o == nil || IsNil(o.OpsPostSync) {
		var ret []SourceOperation
		return ret
	}
	return o.OpsPostSync
}

// GetOpsPostSyncOk returns a tuple with the OpsPostSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDSourceLinkSourceParameters) GetOpsPostSyncOk() ([]SourceOperation, bool) {
	if o == nil || IsNil(o.OpsPostSync) {
		return nil, false
	}
	return o.OpsPostSync, true
}

// HasOpsPostSync returns a boolean if a field has been set.
func (o *BaseDSourceLinkSourceParameters) HasOpsPostSync() bool {
	if o != nil && !IsNil(o.OpsPostSync) {
		return true
	}

	return false
}

// SetOpsPostSync gets a reference to the given []SourceOperation and assigns it to the OpsPostSync field.
func (o *BaseDSourceLinkSourceParameters) SetOpsPostSync(v []SourceOperation) {
	o.OpsPostSync = v
}

func (o BaseDSourceLinkSourceParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseDSourceLinkSourceParameters) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableBaseDSourceLinkSourceParameters struct {
	value *BaseDSourceLinkSourceParameters
	isSet bool
}

func (v NullableBaseDSourceLinkSourceParameters) Get() *BaseDSourceLinkSourceParameters {
	return v.value
}

func (v *NullableBaseDSourceLinkSourceParameters) Set(val *BaseDSourceLinkSourceParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseDSourceLinkSourceParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseDSourceLinkSourceParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseDSourceLinkSourceParameters(val *BaseDSourceLinkSourceParameters) *NullableBaseDSourceLinkSourceParameters {
	return &NullableBaseDSourceLinkSourceParameters{value: val, isSet: true}
}

func (v NullableBaseDSourceLinkSourceParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseDSourceLinkSourceParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


