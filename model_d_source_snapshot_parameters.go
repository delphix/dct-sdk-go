/*
Delphix DCT API

Delphix DCT API

API version: 3.16.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the DSourceSnapshotParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DSourceSnapshotParameters{}

// DSourceSnapshotParameters Parameters to snapshot a DSource.
type DSourceSnapshotParameters struct {
	// If this parameter is set to true, older devices will be dropped and new devices created instead of trying to remap the devices. This might increase the space utilization on Delphix Engine. (ASE only) 
	DropAndRecreateDevices *bool `json:"drop_and_recreate_devices,omitempty"`
	// Determines how the Delphix Engine will take a backup: * `latest_backup` - Use the most recent backup. * `new_backup` - Delphix will take a new backup of your source database. * `specific_backup` - Use a specific backup. Using this option requires setting   `ase_backup_files` for ASE dSources or `mssql_backup_uuid` for MSSql dSources. Default is `new_backup`. (ASE, MSSql only) 
	SyncStrategy *string `json:"sync_strategy,omitempty"`
	// When using the `specific_backup` sync_strategy, determines the backup files. (ASE Only)
	AseBackupFiles []string `json:"ase_backup_files,omitempty"`
	// When using the `specific_backup` sync_strategy, determines the Backup Set UUID. (MSSql only)
	MssqlBackupUuid *string `json:"mssql_backup_uuid,omitempty"`
	// When using the `new_backup` sync_strategy, determines if compression must be enabled. Defaults to the configuration of the ingestion strategy configured on the Delphix Engine for this dSource. (MSSql only)
	CompressionEnabled *bool `json:"compression_enabled,omitempty"`
	// When using the `new_backup` sync_strategy for an MSSql Availability Group, determines the backup policy: * `primary` - Backups only go to the primary node. * `secondary_only` - Backups only go to secondary nodes. If secondary nodes are down, backups will fail. * `prefer_secondary` - Backups go to secondary nodes, but if secondary nodes are down, backups will go to the primary node. (MSSql only) 
	AvailabilityGroupBackupPolicy *string `json:"availability_group_backup_policy,omitempty"`
	// Indicates whether a fresh SnapSync must be started regardless if it was possible to resume the current SnapSync. If true, we will not resume but instead ignore previous progress and backup all datafiles even if already completed from previous failed SnapSync. This does not force a full backup, if an incremental was in progress this will start a new incremental snapshot. (Oracle only) 
	DoNotResume *bool `json:"do_not_resume,omitempty"`
	// Indicates whether two SnapSyncs should be performed in immediate succession to reduce the number of logs required to provision the snapshot. This may significantly reduce the time necessary to provision from a snapshot. (Oracle only). 
	DoubleSync *bool `json:"double_sync,omitempty"`
	// Whether or not to take another full backup of the source database. (Oracle only)
	ForceFullBackup *bool `json:"force_full_backup,omitempty"`
	// Skip check that tests if there is enough space available to store the database in the Delphix Engine. The Delphix Engine estimates how much space a database will occupy after compression and prevents SnapSync if insufficient space is available. This safeguard can be overridden using this option. This may be useful when linking highly compressible databases. (Oracle only) 
	SkipSpaceCheck *bool `json:"skip_space_check,omitempty"`
	// List of datafiles to take a full backup of. This would be useful in situations where certain datafiles could not be backed up during previous SnapSync due to corruption or because they went offline. (Oracle only) 
	FilesForPartialFullBackup []int64 `json:"files_for_partial_full_backup,omitempty"`
	// The list of parameters specified by the snapshotParametersDefinition schema in the toolkit (AppData only).
	AppdataParameters map[string]interface{} `json:"appdata_parameters,omitempty"`
}

// NewDSourceSnapshotParameters instantiates a new DSourceSnapshotParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDSourceSnapshotParameters() *DSourceSnapshotParameters {
	this := DSourceSnapshotParameters{}
	return &this
}

// NewDSourceSnapshotParametersWithDefaults instantiates a new DSourceSnapshotParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDSourceSnapshotParametersWithDefaults() *DSourceSnapshotParameters {
	this := DSourceSnapshotParameters{}
	return &this
}

// GetDropAndRecreateDevices returns the DropAndRecreateDevices field value if set, zero value otherwise.
func (o *DSourceSnapshotParameters) GetDropAndRecreateDevices() bool {
	if o == nil || IsNil(o.DropAndRecreateDevices) {
		var ret bool
		return ret
	}
	return *o.DropAndRecreateDevices
}

// GetDropAndRecreateDevicesOk returns a tuple with the DropAndRecreateDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceSnapshotParameters) GetDropAndRecreateDevicesOk() (*bool, bool) {
	if o == nil || IsNil(o.DropAndRecreateDevices) {
		return nil, false
	}
	return o.DropAndRecreateDevices, true
}

// HasDropAndRecreateDevices returns a boolean if a field has been set.
func (o *DSourceSnapshotParameters) HasDropAndRecreateDevices() bool {
	if o != nil && !IsNil(o.DropAndRecreateDevices) {
		return true
	}

	return false
}

// SetDropAndRecreateDevices gets a reference to the given bool and assigns it to the DropAndRecreateDevices field.
func (o *DSourceSnapshotParameters) SetDropAndRecreateDevices(v bool) {
	o.DropAndRecreateDevices = &v
}

// GetSyncStrategy returns the SyncStrategy field value if set, zero value otherwise.
func (o *DSourceSnapshotParameters) GetSyncStrategy() string {
	if o == nil || IsNil(o.SyncStrategy) {
		var ret string
		return ret
	}
	return *o.SyncStrategy
}

// GetSyncStrategyOk returns a tuple with the SyncStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceSnapshotParameters) GetSyncStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.SyncStrategy) {
		return nil, false
	}
	return o.SyncStrategy, true
}

// HasSyncStrategy returns a boolean if a field has been set.
func (o *DSourceSnapshotParameters) HasSyncStrategy() bool {
	if o != nil && !IsNil(o.SyncStrategy) {
		return true
	}

	return false
}

// SetSyncStrategy gets a reference to the given string and assigns it to the SyncStrategy field.
func (o *DSourceSnapshotParameters) SetSyncStrategy(v string) {
	o.SyncStrategy = &v
}

// GetAseBackupFiles returns the AseBackupFiles field value if set, zero value otherwise.
func (o *DSourceSnapshotParameters) GetAseBackupFiles() []string {
	if o == nil || IsNil(o.AseBackupFiles) {
		var ret []string
		return ret
	}
	return o.AseBackupFiles
}

// GetAseBackupFilesOk returns a tuple with the AseBackupFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceSnapshotParameters) GetAseBackupFilesOk() ([]string, bool) {
	if o == nil || IsNil(o.AseBackupFiles) {
		return nil, false
	}
	return o.AseBackupFiles, true
}

// HasAseBackupFiles returns a boolean if a field has been set.
func (o *DSourceSnapshotParameters) HasAseBackupFiles() bool {
	if o != nil && !IsNil(o.AseBackupFiles) {
		return true
	}

	return false
}

// SetAseBackupFiles gets a reference to the given []string and assigns it to the AseBackupFiles field.
func (o *DSourceSnapshotParameters) SetAseBackupFiles(v []string) {
	o.AseBackupFiles = v
}

// GetMssqlBackupUuid returns the MssqlBackupUuid field value if set, zero value otherwise.
func (o *DSourceSnapshotParameters) GetMssqlBackupUuid() string {
	if o == nil || IsNil(o.MssqlBackupUuid) {
		var ret string
		return ret
	}
	return *o.MssqlBackupUuid
}

// GetMssqlBackupUuidOk returns a tuple with the MssqlBackupUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceSnapshotParameters) GetMssqlBackupUuidOk() (*string, bool) {
	if o == nil || IsNil(o.MssqlBackupUuid) {
		return nil, false
	}
	return o.MssqlBackupUuid, true
}

// HasMssqlBackupUuid returns a boolean if a field has been set.
func (o *DSourceSnapshotParameters) HasMssqlBackupUuid() bool {
	if o != nil && !IsNil(o.MssqlBackupUuid) {
		return true
	}

	return false
}

// SetMssqlBackupUuid gets a reference to the given string and assigns it to the MssqlBackupUuid field.
func (o *DSourceSnapshotParameters) SetMssqlBackupUuid(v string) {
	o.MssqlBackupUuid = &v
}

// GetCompressionEnabled returns the CompressionEnabled field value if set, zero value otherwise.
func (o *DSourceSnapshotParameters) GetCompressionEnabled() bool {
	if o == nil || IsNil(o.CompressionEnabled) {
		var ret bool
		return ret
	}
	return *o.CompressionEnabled
}

// GetCompressionEnabledOk returns a tuple with the CompressionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceSnapshotParameters) GetCompressionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CompressionEnabled) {
		return nil, false
	}
	return o.CompressionEnabled, true
}

// HasCompressionEnabled returns a boolean if a field has been set.
func (o *DSourceSnapshotParameters) HasCompressionEnabled() bool {
	if o != nil && !IsNil(o.CompressionEnabled) {
		return true
	}

	return false
}

// SetCompressionEnabled gets a reference to the given bool and assigns it to the CompressionEnabled field.
func (o *DSourceSnapshotParameters) SetCompressionEnabled(v bool) {
	o.CompressionEnabled = &v
}

// GetAvailabilityGroupBackupPolicy returns the AvailabilityGroupBackupPolicy field value if set, zero value otherwise.
func (o *DSourceSnapshotParameters) GetAvailabilityGroupBackupPolicy() string {
	if o == nil || IsNil(o.AvailabilityGroupBackupPolicy) {
		var ret string
		return ret
	}
	return *o.AvailabilityGroupBackupPolicy
}

// GetAvailabilityGroupBackupPolicyOk returns a tuple with the AvailabilityGroupBackupPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceSnapshotParameters) GetAvailabilityGroupBackupPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityGroupBackupPolicy) {
		return nil, false
	}
	return o.AvailabilityGroupBackupPolicy, true
}

// HasAvailabilityGroupBackupPolicy returns a boolean if a field has been set.
func (o *DSourceSnapshotParameters) HasAvailabilityGroupBackupPolicy() bool {
	if o != nil && !IsNil(o.AvailabilityGroupBackupPolicy) {
		return true
	}

	return false
}

// SetAvailabilityGroupBackupPolicy gets a reference to the given string and assigns it to the AvailabilityGroupBackupPolicy field.
func (o *DSourceSnapshotParameters) SetAvailabilityGroupBackupPolicy(v string) {
	o.AvailabilityGroupBackupPolicy = &v
}

// GetDoNotResume returns the DoNotResume field value if set, zero value otherwise.
func (o *DSourceSnapshotParameters) GetDoNotResume() bool {
	if o == nil || IsNil(o.DoNotResume) {
		var ret bool
		return ret
	}
	return *o.DoNotResume
}

// GetDoNotResumeOk returns a tuple with the DoNotResume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceSnapshotParameters) GetDoNotResumeOk() (*bool, bool) {
	if o == nil || IsNil(o.DoNotResume) {
		return nil, false
	}
	return o.DoNotResume, true
}

// HasDoNotResume returns a boolean if a field has been set.
func (o *DSourceSnapshotParameters) HasDoNotResume() bool {
	if o != nil && !IsNil(o.DoNotResume) {
		return true
	}

	return false
}

// SetDoNotResume gets a reference to the given bool and assigns it to the DoNotResume field.
func (o *DSourceSnapshotParameters) SetDoNotResume(v bool) {
	o.DoNotResume = &v
}

// GetDoubleSync returns the DoubleSync field value if set, zero value otherwise.
func (o *DSourceSnapshotParameters) GetDoubleSync() bool {
	if o == nil || IsNil(o.DoubleSync) {
		var ret bool
		return ret
	}
	return *o.DoubleSync
}

// GetDoubleSyncOk returns a tuple with the DoubleSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceSnapshotParameters) GetDoubleSyncOk() (*bool, bool) {
	if o == nil || IsNil(o.DoubleSync) {
		return nil, false
	}
	return o.DoubleSync, true
}

// HasDoubleSync returns a boolean if a field has been set.
func (o *DSourceSnapshotParameters) HasDoubleSync() bool {
	if o != nil && !IsNil(o.DoubleSync) {
		return true
	}

	return false
}

// SetDoubleSync gets a reference to the given bool and assigns it to the DoubleSync field.
func (o *DSourceSnapshotParameters) SetDoubleSync(v bool) {
	o.DoubleSync = &v
}

// GetForceFullBackup returns the ForceFullBackup field value if set, zero value otherwise.
func (o *DSourceSnapshotParameters) GetForceFullBackup() bool {
	if o == nil || IsNil(o.ForceFullBackup) {
		var ret bool
		return ret
	}
	return *o.ForceFullBackup
}

// GetForceFullBackupOk returns a tuple with the ForceFullBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceSnapshotParameters) GetForceFullBackupOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceFullBackup) {
		return nil, false
	}
	return o.ForceFullBackup, true
}

// HasForceFullBackup returns a boolean if a field has been set.
func (o *DSourceSnapshotParameters) HasForceFullBackup() bool {
	if o != nil && !IsNil(o.ForceFullBackup) {
		return true
	}

	return false
}

// SetForceFullBackup gets a reference to the given bool and assigns it to the ForceFullBackup field.
func (o *DSourceSnapshotParameters) SetForceFullBackup(v bool) {
	o.ForceFullBackup = &v
}

// GetSkipSpaceCheck returns the SkipSpaceCheck field value if set, zero value otherwise.
func (o *DSourceSnapshotParameters) GetSkipSpaceCheck() bool {
	if o == nil || IsNil(o.SkipSpaceCheck) {
		var ret bool
		return ret
	}
	return *o.SkipSpaceCheck
}

// GetSkipSpaceCheckOk returns a tuple with the SkipSpaceCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceSnapshotParameters) GetSkipSpaceCheckOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipSpaceCheck) {
		return nil, false
	}
	return o.SkipSpaceCheck, true
}

// HasSkipSpaceCheck returns a boolean if a field has been set.
func (o *DSourceSnapshotParameters) HasSkipSpaceCheck() bool {
	if o != nil && !IsNil(o.SkipSpaceCheck) {
		return true
	}

	return false
}

// SetSkipSpaceCheck gets a reference to the given bool and assigns it to the SkipSpaceCheck field.
func (o *DSourceSnapshotParameters) SetSkipSpaceCheck(v bool) {
	o.SkipSpaceCheck = &v
}

// GetFilesForPartialFullBackup returns the FilesForPartialFullBackup field value if set, zero value otherwise.
func (o *DSourceSnapshotParameters) GetFilesForPartialFullBackup() []int64 {
	if o == nil || IsNil(o.FilesForPartialFullBackup) {
		var ret []int64
		return ret
	}
	return o.FilesForPartialFullBackup
}

// GetFilesForPartialFullBackupOk returns a tuple with the FilesForPartialFullBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceSnapshotParameters) GetFilesForPartialFullBackupOk() ([]int64, bool) {
	if o == nil || IsNil(o.FilesForPartialFullBackup) {
		return nil, false
	}
	return o.FilesForPartialFullBackup, true
}

// HasFilesForPartialFullBackup returns a boolean if a field has been set.
func (o *DSourceSnapshotParameters) HasFilesForPartialFullBackup() bool {
	if o != nil && !IsNil(o.FilesForPartialFullBackup) {
		return true
	}

	return false
}

// SetFilesForPartialFullBackup gets a reference to the given []int64 and assigns it to the FilesForPartialFullBackup field.
func (o *DSourceSnapshotParameters) SetFilesForPartialFullBackup(v []int64) {
	o.FilesForPartialFullBackup = v
}

// GetAppdataParameters returns the AppdataParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DSourceSnapshotParameters) GetAppdataParameters() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AppdataParameters
}

// GetAppdataParametersOk returns a tuple with the AppdataParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DSourceSnapshotParameters) GetAppdataParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AppdataParameters) {
		return map[string]interface{}{}, false
	}
	return o.AppdataParameters, true
}

// HasAppdataParameters returns a boolean if a field has been set.
func (o *DSourceSnapshotParameters) HasAppdataParameters() bool {
	if o != nil && !IsNil(o.AppdataParameters) {
		return true
	}

	return false
}

// SetAppdataParameters gets a reference to the given map[string]interface{} and assigns it to the AppdataParameters field.
func (o *DSourceSnapshotParameters) SetAppdataParameters(v map[string]interface{}) {
	o.AppdataParameters = v
}

func (o DSourceSnapshotParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DSourceSnapshotParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DropAndRecreateDevices) {
		toSerialize["drop_and_recreate_devices"] = o.DropAndRecreateDevices
	}
	if !IsNil(o.SyncStrategy) {
		toSerialize["sync_strategy"] = o.SyncStrategy
	}
	if !IsNil(o.AseBackupFiles) {
		toSerialize["ase_backup_files"] = o.AseBackupFiles
	}
	if !IsNil(o.MssqlBackupUuid) {
		toSerialize["mssql_backup_uuid"] = o.MssqlBackupUuid
	}
	if !IsNil(o.CompressionEnabled) {
		toSerialize["compression_enabled"] = o.CompressionEnabled
	}
	if !IsNil(o.AvailabilityGroupBackupPolicy) {
		toSerialize["availability_group_backup_policy"] = o.AvailabilityGroupBackupPolicy
	}
	if !IsNil(o.DoNotResume) {
		toSerialize["do_not_resume"] = o.DoNotResume
	}
	if !IsNil(o.DoubleSync) {
		toSerialize["double_sync"] = o.DoubleSync
	}
	if !IsNil(o.ForceFullBackup) {
		toSerialize["force_full_backup"] = o.ForceFullBackup
	}
	if !IsNil(o.SkipSpaceCheck) {
		toSerialize["skip_space_check"] = o.SkipSpaceCheck
	}
	if !IsNil(o.FilesForPartialFullBackup) {
		toSerialize["files_for_partial_full_backup"] = o.FilesForPartialFullBackup
	}
	if o.AppdataParameters != nil {
		toSerialize["appdata_parameters"] = o.AppdataParameters
	}
	return toSerialize, nil
}

type NullableDSourceSnapshotParameters struct {
	value *DSourceSnapshotParameters
	isSet bool
}

func (v NullableDSourceSnapshotParameters) Get() *DSourceSnapshotParameters {
	return v.value
}

func (v *NullableDSourceSnapshotParameters) Set(val *DSourceSnapshotParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableDSourceSnapshotParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableDSourceSnapshotParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDSourceSnapshotParameters(val *DSourceSnapshotParameters) *NullableDSourceSnapshotParameters {
	return &NullableDSourceSnapshotParameters{value: val, isSet: true}
}

func (v NullableDSourceSnapshotParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDSourceSnapshotParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


