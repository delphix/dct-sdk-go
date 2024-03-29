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
	"time"
)

// checks if the MaskingJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaskingJob{}

// MaskingJob A masking job.
type MaskingJob struct {
	// The Masking Job entity ID.
	Id *string `json:"id,omitempty"`
	// The name of this Masking Job.
	Name *string `json:"name,omitempty"`
	Ruleset *MaskingRuleset `json:"ruleset,omitempty"`
	// The type of data being masked by this Job. If the Masking Job is masking a database this is the type of the database (Standard Job only).
	ConnectorType *string `json:"connector_type,omitempty"`
	// Whether this is an on-the-fly masking job (Standard Job only).
	IsOnTheFlyMasking *bool `json:"is_on_the_fly_masking,omitempty"`
	// The date this MaskingJob was created (Standard Job only).
	CreationDate *time.Time `json:"creation_date,omitempty"`
	// The date this MaskingJob was last executed to completion.
	LastCompletedExecutionDate *time.Time `json:"last_completed_execution_date,omitempty"`
	// The status of this MaskingJob's last execution.
	LastExecutionStatus *string `json:"last_execution_status,omitempty"`
	// The ID of this MaskingJob's last execution.
	LastExecutionId *string `json:"last_execution_id,omitempty"`
	// The username of the Connector used by the MaskingJob (Standard Job only). For hyperscale jobs, see the connector of the dataset.
	ConnectorUsername NullableString `json:"connector_username,omitempty"`
	// The password of the Connector used by the MaskingJob (Standard Job only). For hyperscale jobs, see the connector of the dataset.
	ConnectorPassword NullableString `json:"connector_password,omitempty"`
	// The username of the source Connector used by the on-the-fly MaskingJob (Standard Job only).
	OnTheFlySourceConnectorUsername NullableString `json:"on_the_fly_source_connector_username,omitempty"`
	// The password of the source Connector used by the on-the-fly MaskingJob (Standard Job only).
	OnTheFlySourceConnectorPassword NullableString `json:"on_the_fly_source_connector_password,omitempty"`
	// The type of this Job.
	JobType *string `json:"job_type,omitempty"`
	// The ID of the Hyperscale instance of this Job (Hyperscale Job only).
	HyperscaleInstanceId *string `json:"hyperscale_instance_id,omitempty"`
	// Description of the Job (Hyperscale Job only).
	Description *string `json:"description,omitempty"`
	// Dataset of the Hyperscale Job (Hyperscale Job only).
	DatasetId *string `json:"dataset_id,omitempty"`
	// Defines whether execution data will be stored after execution is complete (Hyperscale Job only).
	RetainExecutionData *string `json:"retain_execution_data,omitempty"`
	// Maximum memory to be allocated for each Masking job (Hyperscale Job only).
	MaxMemory *int32 `json:"max_memory,omitempty"`
	// Minimum memory to be allocated for each Masking job (Hyperscale Job only).
	MinMemory *int32 `json:"min_memory,omitempty"`
	// Feedback Size for each Masking job (Hyperscale Job only).
	FeedbackSize *int32 `json:"feedback_size,omitempty"`
	// Stream Row Limit for each Masking job (Hyperscale Job only).
	StreamRowLimit *int32 `json:"stream_row_limit,omitempty"`
	// Number of input streams to be configured for Masking Job (Hyperscale Job only).
	NumInputStreams *int32 `json:"num_input_streams,omitempty"`
	// Maximum number of parallel connection that the Hyperscale instance can have with the source datasource (Hyperscale Job only).
	MaxConcurrentSourceConnections *int32 `json:"max_concurrent_source_connections,omitempty"`
	// Maximum number of parallel connection that the Hyperscale instance can have with the target datasource (Hyperscale Job only).
	MaxConcurrentTargetConnections *int32 `json:"max_concurrent_target_connections,omitempty"`
	// The degree of parallelism (DOP) per Oracle job to recreate the index in the post-load process (Hyperscale Job only).
	ParallelismDegree *int32 `json:"parallelism_degree,omitempty"`
	// The ID of the MaskingJob that was used as the source to create this job (Hyperscale Job only).
	SourceMaskingJobId *string `json:"source_masking_job_id,omitempty"`
	// List of engines that this job can run on (Hyperscale Job only).
	EngineIds []string `json:"engine_ids,omitempty"`
	Tags []Tag `json:"tags,omitempty"`
}

// NewMaskingJob instantiates a new MaskingJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaskingJob() *MaskingJob {
	this := MaskingJob{}
	return &this
}

// NewMaskingJobWithDefaults instantiates a new MaskingJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaskingJobWithDefaults() *MaskingJob {
	this := MaskingJob{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MaskingJob) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MaskingJob) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MaskingJob) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MaskingJob) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MaskingJob) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MaskingJob) SetName(v string) {
	o.Name = &v
}

// GetRuleset returns the Ruleset field value if set, zero value otherwise.
func (o *MaskingJob) GetRuleset() MaskingRuleset {
	if o == nil || IsNil(o.Ruleset) {
		var ret MaskingRuleset
		return ret
	}
	return *o.Ruleset
}

// GetRulesetOk returns a tuple with the Ruleset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetRulesetOk() (*MaskingRuleset, bool) {
	if o == nil || IsNil(o.Ruleset) {
		return nil, false
	}
	return o.Ruleset, true
}

// HasRuleset returns a boolean if a field has been set.
func (o *MaskingJob) HasRuleset() bool {
	if o != nil && !IsNil(o.Ruleset) {
		return true
	}

	return false
}

// SetRuleset gets a reference to the given MaskingRuleset and assigns it to the Ruleset field.
func (o *MaskingJob) SetRuleset(v MaskingRuleset) {
	o.Ruleset = &v
}

// GetConnectorType returns the ConnectorType field value if set, zero value otherwise.
func (o *MaskingJob) GetConnectorType() string {
	if o == nil || IsNil(o.ConnectorType) {
		var ret string
		return ret
	}
	return *o.ConnectorType
}

// GetConnectorTypeOk returns a tuple with the ConnectorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetConnectorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectorType) {
		return nil, false
	}
	return o.ConnectorType, true
}

// HasConnectorType returns a boolean if a field has been set.
func (o *MaskingJob) HasConnectorType() bool {
	if o != nil && !IsNil(o.ConnectorType) {
		return true
	}

	return false
}

// SetConnectorType gets a reference to the given string and assigns it to the ConnectorType field.
func (o *MaskingJob) SetConnectorType(v string) {
	o.ConnectorType = &v
}

// GetIsOnTheFlyMasking returns the IsOnTheFlyMasking field value if set, zero value otherwise.
func (o *MaskingJob) GetIsOnTheFlyMasking() bool {
	if o == nil || IsNil(o.IsOnTheFlyMasking) {
		var ret bool
		return ret
	}
	return *o.IsOnTheFlyMasking
}

// GetIsOnTheFlyMaskingOk returns a tuple with the IsOnTheFlyMasking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetIsOnTheFlyMaskingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOnTheFlyMasking) {
		return nil, false
	}
	return o.IsOnTheFlyMasking, true
}

// HasIsOnTheFlyMasking returns a boolean if a field has been set.
func (o *MaskingJob) HasIsOnTheFlyMasking() bool {
	if o != nil && !IsNil(o.IsOnTheFlyMasking) {
		return true
	}

	return false
}

// SetIsOnTheFlyMasking gets a reference to the given bool and assigns it to the IsOnTheFlyMasking field.
func (o *MaskingJob) SetIsOnTheFlyMasking(v bool) {
	o.IsOnTheFlyMasking = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *MaskingJob) GetCreationDate() time.Time {
	if o == nil || IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *MaskingJob) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *MaskingJob) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetLastCompletedExecutionDate returns the LastCompletedExecutionDate field value if set, zero value otherwise.
func (o *MaskingJob) GetLastCompletedExecutionDate() time.Time {
	if o == nil || IsNil(o.LastCompletedExecutionDate) {
		var ret time.Time
		return ret
	}
	return *o.LastCompletedExecutionDate
}

// GetLastCompletedExecutionDateOk returns a tuple with the LastCompletedExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetLastCompletedExecutionDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastCompletedExecutionDate) {
		return nil, false
	}
	return o.LastCompletedExecutionDate, true
}

// HasLastCompletedExecutionDate returns a boolean if a field has been set.
func (o *MaskingJob) HasLastCompletedExecutionDate() bool {
	if o != nil && !IsNil(o.LastCompletedExecutionDate) {
		return true
	}

	return false
}

// SetLastCompletedExecutionDate gets a reference to the given time.Time and assigns it to the LastCompletedExecutionDate field.
func (o *MaskingJob) SetLastCompletedExecutionDate(v time.Time) {
	o.LastCompletedExecutionDate = &v
}

// GetLastExecutionStatus returns the LastExecutionStatus field value if set, zero value otherwise.
func (o *MaskingJob) GetLastExecutionStatus() string {
	if o == nil || IsNil(o.LastExecutionStatus) {
		var ret string
		return ret
	}
	return *o.LastExecutionStatus
}

// GetLastExecutionStatusOk returns a tuple with the LastExecutionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetLastExecutionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LastExecutionStatus) {
		return nil, false
	}
	return o.LastExecutionStatus, true
}

// HasLastExecutionStatus returns a boolean if a field has been set.
func (o *MaskingJob) HasLastExecutionStatus() bool {
	if o != nil && !IsNil(o.LastExecutionStatus) {
		return true
	}

	return false
}

// SetLastExecutionStatus gets a reference to the given string and assigns it to the LastExecutionStatus field.
func (o *MaskingJob) SetLastExecutionStatus(v string) {
	o.LastExecutionStatus = &v
}

// GetLastExecutionId returns the LastExecutionId field value if set, zero value otherwise.
func (o *MaskingJob) GetLastExecutionId() string {
	if o == nil || IsNil(o.LastExecutionId) {
		var ret string
		return ret
	}
	return *o.LastExecutionId
}

// GetLastExecutionIdOk returns a tuple with the LastExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetLastExecutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastExecutionId) {
		return nil, false
	}
	return o.LastExecutionId, true
}

// HasLastExecutionId returns a boolean if a field has been set.
func (o *MaskingJob) HasLastExecutionId() bool {
	if o != nil && !IsNil(o.LastExecutionId) {
		return true
	}

	return false
}

// SetLastExecutionId gets a reference to the given string and assigns it to the LastExecutionId field.
func (o *MaskingJob) SetLastExecutionId(v string) {
	o.LastExecutionId = &v
}

// GetConnectorUsername returns the ConnectorUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MaskingJob) GetConnectorUsername() string {
	if o == nil || IsNil(o.ConnectorUsername.Get()) {
		var ret string
		return ret
	}
	return *o.ConnectorUsername.Get()
}

// GetConnectorUsernameOk returns a tuple with the ConnectorUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MaskingJob) GetConnectorUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectorUsername.Get(), o.ConnectorUsername.IsSet()
}

// HasConnectorUsername returns a boolean if a field has been set.
func (o *MaskingJob) HasConnectorUsername() bool {
	if o != nil && o.ConnectorUsername.IsSet() {
		return true
	}

	return false
}

// SetConnectorUsername gets a reference to the given NullableString and assigns it to the ConnectorUsername field.
func (o *MaskingJob) SetConnectorUsername(v string) {
	o.ConnectorUsername.Set(&v)
}
// SetConnectorUsernameNil sets the value for ConnectorUsername to be an explicit nil
func (o *MaskingJob) SetConnectorUsernameNil() {
	o.ConnectorUsername.Set(nil)
}

// UnsetConnectorUsername ensures that no value is present for ConnectorUsername, not even an explicit nil
func (o *MaskingJob) UnsetConnectorUsername() {
	o.ConnectorUsername.Unset()
}

// GetConnectorPassword returns the ConnectorPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MaskingJob) GetConnectorPassword() string {
	if o == nil || IsNil(o.ConnectorPassword.Get()) {
		var ret string
		return ret
	}
	return *o.ConnectorPassword.Get()
}

// GetConnectorPasswordOk returns a tuple with the ConnectorPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MaskingJob) GetConnectorPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectorPassword.Get(), o.ConnectorPassword.IsSet()
}

// HasConnectorPassword returns a boolean if a field has been set.
func (o *MaskingJob) HasConnectorPassword() bool {
	if o != nil && o.ConnectorPassword.IsSet() {
		return true
	}

	return false
}

// SetConnectorPassword gets a reference to the given NullableString and assigns it to the ConnectorPassword field.
func (o *MaskingJob) SetConnectorPassword(v string) {
	o.ConnectorPassword.Set(&v)
}
// SetConnectorPasswordNil sets the value for ConnectorPassword to be an explicit nil
func (o *MaskingJob) SetConnectorPasswordNil() {
	o.ConnectorPassword.Set(nil)
}

// UnsetConnectorPassword ensures that no value is present for ConnectorPassword, not even an explicit nil
func (o *MaskingJob) UnsetConnectorPassword() {
	o.ConnectorPassword.Unset()
}

// GetOnTheFlySourceConnectorUsername returns the OnTheFlySourceConnectorUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MaskingJob) GetOnTheFlySourceConnectorUsername() string {
	if o == nil || IsNil(o.OnTheFlySourceConnectorUsername.Get()) {
		var ret string
		return ret
	}
	return *o.OnTheFlySourceConnectorUsername.Get()
}

// GetOnTheFlySourceConnectorUsernameOk returns a tuple with the OnTheFlySourceConnectorUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MaskingJob) GetOnTheFlySourceConnectorUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnTheFlySourceConnectorUsername.Get(), o.OnTheFlySourceConnectorUsername.IsSet()
}

// HasOnTheFlySourceConnectorUsername returns a boolean if a field has been set.
func (o *MaskingJob) HasOnTheFlySourceConnectorUsername() bool {
	if o != nil && o.OnTheFlySourceConnectorUsername.IsSet() {
		return true
	}

	return false
}

// SetOnTheFlySourceConnectorUsername gets a reference to the given NullableString and assigns it to the OnTheFlySourceConnectorUsername field.
func (o *MaskingJob) SetOnTheFlySourceConnectorUsername(v string) {
	o.OnTheFlySourceConnectorUsername.Set(&v)
}
// SetOnTheFlySourceConnectorUsernameNil sets the value for OnTheFlySourceConnectorUsername to be an explicit nil
func (o *MaskingJob) SetOnTheFlySourceConnectorUsernameNil() {
	o.OnTheFlySourceConnectorUsername.Set(nil)
}

// UnsetOnTheFlySourceConnectorUsername ensures that no value is present for OnTheFlySourceConnectorUsername, not even an explicit nil
func (o *MaskingJob) UnsetOnTheFlySourceConnectorUsername() {
	o.OnTheFlySourceConnectorUsername.Unset()
}

// GetOnTheFlySourceConnectorPassword returns the OnTheFlySourceConnectorPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MaskingJob) GetOnTheFlySourceConnectorPassword() string {
	if o == nil || IsNil(o.OnTheFlySourceConnectorPassword.Get()) {
		var ret string
		return ret
	}
	return *o.OnTheFlySourceConnectorPassword.Get()
}

// GetOnTheFlySourceConnectorPasswordOk returns a tuple with the OnTheFlySourceConnectorPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MaskingJob) GetOnTheFlySourceConnectorPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnTheFlySourceConnectorPassword.Get(), o.OnTheFlySourceConnectorPassword.IsSet()
}

// HasOnTheFlySourceConnectorPassword returns a boolean if a field has been set.
func (o *MaskingJob) HasOnTheFlySourceConnectorPassword() bool {
	if o != nil && o.OnTheFlySourceConnectorPassword.IsSet() {
		return true
	}

	return false
}

// SetOnTheFlySourceConnectorPassword gets a reference to the given NullableString and assigns it to the OnTheFlySourceConnectorPassword field.
func (o *MaskingJob) SetOnTheFlySourceConnectorPassword(v string) {
	o.OnTheFlySourceConnectorPassword.Set(&v)
}
// SetOnTheFlySourceConnectorPasswordNil sets the value for OnTheFlySourceConnectorPassword to be an explicit nil
func (o *MaskingJob) SetOnTheFlySourceConnectorPasswordNil() {
	o.OnTheFlySourceConnectorPassword.Set(nil)
}

// UnsetOnTheFlySourceConnectorPassword ensures that no value is present for OnTheFlySourceConnectorPassword, not even an explicit nil
func (o *MaskingJob) UnsetOnTheFlySourceConnectorPassword() {
	o.OnTheFlySourceConnectorPassword.Unset()
}

// GetJobType returns the JobType field value if set, zero value otherwise.
func (o *MaskingJob) GetJobType() string {
	if o == nil || IsNil(o.JobType) {
		var ret string
		return ret
	}
	return *o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetJobTypeOk() (*string, bool) {
	if o == nil || IsNil(o.JobType) {
		return nil, false
	}
	return o.JobType, true
}

// HasJobType returns a boolean if a field has been set.
func (o *MaskingJob) HasJobType() bool {
	if o != nil && !IsNil(o.JobType) {
		return true
	}

	return false
}

// SetJobType gets a reference to the given string and assigns it to the JobType field.
func (o *MaskingJob) SetJobType(v string) {
	o.JobType = &v
}

// GetHyperscaleInstanceId returns the HyperscaleInstanceId field value if set, zero value otherwise.
func (o *MaskingJob) GetHyperscaleInstanceId() string {
	if o == nil || IsNil(o.HyperscaleInstanceId) {
		var ret string
		return ret
	}
	return *o.HyperscaleInstanceId
}

// GetHyperscaleInstanceIdOk returns a tuple with the HyperscaleInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetHyperscaleInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.HyperscaleInstanceId) {
		return nil, false
	}
	return o.HyperscaleInstanceId, true
}

// HasHyperscaleInstanceId returns a boolean if a field has been set.
func (o *MaskingJob) HasHyperscaleInstanceId() bool {
	if o != nil && !IsNil(o.HyperscaleInstanceId) {
		return true
	}

	return false
}

// SetHyperscaleInstanceId gets a reference to the given string and assigns it to the HyperscaleInstanceId field.
func (o *MaskingJob) SetHyperscaleInstanceId(v string) {
	o.HyperscaleInstanceId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MaskingJob) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MaskingJob) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MaskingJob) SetDescription(v string) {
	o.Description = &v
}

// GetDatasetId returns the DatasetId field value if set, zero value otherwise.
func (o *MaskingJob) GetDatasetId() string {
	if o == nil || IsNil(o.DatasetId) {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetDatasetIdOk() (*string, bool) {
	if o == nil || IsNil(o.DatasetId) {
		return nil, false
	}
	return o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *MaskingJob) HasDatasetId() bool {
	if o != nil && !IsNil(o.DatasetId) {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *MaskingJob) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetRetainExecutionData returns the RetainExecutionData field value if set, zero value otherwise.
func (o *MaskingJob) GetRetainExecutionData() string {
	if o == nil || IsNil(o.RetainExecutionData) {
		var ret string
		return ret
	}
	return *o.RetainExecutionData
}

// GetRetainExecutionDataOk returns a tuple with the RetainExecutionData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetRetainExecutionDataOk() (*string, bool) {
	if o == nil || IsNil(o.RetainExecutionData) {
		return nil, false
	}
	return o.RetainExecutionData, true
}

// HasRetainExecutionData returns a boolean if a field has been set.
func (o *MaskingJob) HasRetainExecutionData() bool {
	if o != nil && !IsNil(o.RetainExecutionData) {
		return true
	}

	return false
}

// SetRetainExecutionData gets a reference to the given string and assigns it to the RetainExecutionData field.
func (o *MaskingJob) SetRetainExecutionData(v string) {
	o.RetainExecutionData = &v
}

// GetMaxMemory returns the MaxMemory field value if set, zero value otherwise.
func (o *MaskingJob) GetMaxMemory() int32 {
	if o == nil || IsNil(o.MaxMemory) {
		var ret int32
		return ret
	}
	return *o.MaxMemory
}

// GetMaxMemoryOk returns a tuple with the MaxMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetMaxMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxMemory) {
		return nil, false
	}
	return o.MaxMemory, true
}

// HasMaxMemory returns a boolean if a field has been set.
func (o *MaskingJob) HasMaxMemory() bool {
	if o != nil && !IsNil(o.MaxMemory) {
		return true
	}

	return false
}

// SetMaxMemory gets a reference to the given int32 and assigns it to the MaxMemory field.
func (o *MaskingJob) SetMaxMemory(v int32) {
	o.MaxMemory = &v
}

// GetMinMemory returns the MinMemory field value if set, zero value otherwise.
func (o *MaskingJob) GetMinMemory() int32 {
	if o == nil || IsNil(o.MinMemory) {
		var ret int32
		return ret
	}
	return *o.MinMemory
}

// GetMinMemoryOk returns a tuple with the MinMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetMinMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.MinMemory) {
		return nil, false
	}
	return o.MinMemory, true
}

// HasMinMemory returns a boolean if a field has been set.
func (o *MaskingJob) HasMinMemory() bool {
	if o != nil && !IsNil(o.MinMemory) {
		return true
	}

	return false
}

// SetMinMemory gets a reference to the given int32 and assigns it to the MinMemory field.
func (o *MaskingJob) SetMinMemory(v int32) {
	o.MinMemory = &v
}

// GetFeedbackSize returns the FeedbackSize field value if set, zero value otherwise.
func (o *MaskingJob) GetFeedbackSize() int32 {
	if o == nil || IsNil(o.FeedbackSize) {
		var ret int32
		return ret
	}
	return *o.FeedbackSize
}

// GetFeedbackSizeOk returns a tuple with the FeedbackSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetFeedbackSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FeedbackSize) {
		return nil, false
	}
	return o.FeedbackSize, true
}

// HasFeedbackSize returns a boolean if a field has been set.
func (o *MaskingJob) HasFeedbackSize() bool {
	if o != nil && !IsNil(o.FeedbackSize) {
		return true
	}

	return false
}

// SetFeedbackSize gets a reference to the given int32 and assigns it to the FeedbackSize field.
func (o *MaskingJob) SetFeedbackSize(v int32) {
	o.FeedbackSize = &v
}

// GetStreamRowLimit returns the StreamRowLimit field value if set, zero value otherwise.
func (o *MaskingJob) GetStreamRowLimit() int32 {
	if o == nil || IsNil(o.StreamRowLimit) {
		var ret int32
		return ret
	}
	return *o.StreamRowLimit
}

// GetStreamRowLimitOk returns a tuple with the StreamRowLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetStreamRowLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.StreamRowLimit) {
		return nil, false
	}
	return o.StreamRowLimit, true
}

// HasStreamRowLimit returns a boolean if a field has been set.
func (o *MaskingJob) HasStreamRowLimit() bool {
	if o != nil && !IsNil(o.StreamRowLimit) {
		return true
	}

	return false
}

// SetStreamRowLimit gets a reference to the given int32 and assigns it to the StreamRowLimit field.
func (o *MaskingJob) SetStreamRowLimit(v int32) {
	o.StreamRowLimit = &v
}

// GetNumInputStreams returns the NumInputStreams field value if set, zero value otherwise.
func (o *MaskingJob) GetNumInputStreams() int32 {
	if o == nil || IsNil(o.NumInputStreams) {
		var ret int32
		return ret
	}
	return *o.NumInputStreams
}

// GetNumInputStreamsOk returns a tuple with the NumInputStreams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetNumInputStreamsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumInputStreams) {
		return nil, false
	}
	return o.NumInputStreams, true
}

// HasNumInputStreams returns a boolean if a field has been set.
func (o *MaskingJob) HasNumInputStreams() bool {
	if o != nil && !IsNil(o.NumInputStreams) {
		return true
	}

	return false
}

// SetNumInputStreams gets a reference to the given int32 and assigns it to the NumInputStreams field.
func (o *MaskingJob) SetNumInputStreams(v int32) {
	o.NumInputStreams = &v
}

// GetMaxConcurrentSourceConnections returns the MaxConcurrentSourceConnections field value if set, zero value otherwise.
func (o *MaskingJob) GetMaxConcurrentSourceConnections() int32 {
	if o == nil || IsNil(o.MaxConcurrentSourceConnections) {
		var ret int32
		return ret
	}
	return *o.MaxConcurrentSourceConnections
}

// GetMaxConcurrentSourceConnectionsOk returns a tuple with the MaxConcurrentSourceConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetMaxConcurrentSourceConnectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxConcurrentSourceConnections) {
		return nil, false
	}
	return o.MaxConcurrentSourceConnections, true
}

// HasMaxConcurrentSourceConnections returns a boolean if a field has been set.
func (o *MaskingJob) HasMaxConcurrentSourceConnections() bool {
	if o != nil && !IsNil(o.MaxConcurrentSourceConnections) {
		return true
	}

	return false
}

// SetMaxConcurrentSourceConnections gets a reference to the given int32 and assigns it to the MaxConcurrentSourceConnections field.
func (o *MaskingJob) SetMaxConcurrentSourceConnections(v int32) {
	o.MaxConcurrentSourceConnections = &v
}

// GetMaxConcurrentTargetConnections returns the MaxConcurrentTargetConnections field value if set, zero value otherwise.
func (o *MaskingJob) GetMaxConcurrentTargetConnections() int32 {
	if o == nil || IsNil(o.MaxConcurrentTargetConnections) {
		var ret int32
		return ret
	}
	return *o.MaxConcurrentTargetConnections
}

// GetMaxConcurrentTargetConnectionsOk returns a tuple with the MaxConcurrentTargetConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetMaxConcurrentTargetConnectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxConcurrentTargetConnections) {
		return nil, false
	}
	return o.MaxConcurrentTargetConnections, true
}

// HasMaxConcurrentTargetConnections returns a boolean if a field has been set.
func (o *MaskingJob) HasMaxConcurrentTargetConnections() bool {
	if o != nil && !IsNil(o.MaxConcurrentTargetConnections) {
		return true
	}

	return false
}

// SetMaxConcurrentTargetConnections gets a reference to the given int32 and assigns it to the MaxConcurrentTargetConnections field.
func (o *MaskingJob) SetMaxConcurrentTargetConnections(v int32) {
	o.MaxConcurrentTargetConnections = &v
}

// GetParallelismDegree returns the ParallelismDegree field value if set, zero value otherwise.
func (o *MaskingJob) GetParallelismDegree() int32 {
	if o == nil || IsNil(o.ParallelismDegree) {
		var ret int32
		return ret
	}
	return *o.ParallelismDegree
}

// GetParallelismDegreeOk returns a tuple with the ParallelismDegree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetParallelismDegreeOk() (*int32, bool) {
	if o == nil || IsNil(o.ParallelismDegree) {
		return nil, false
	}
	return o.ParallelismDegree, true
}

// HasParallelismDegree returns a boolean if a field has been set.
func (o *MaskingJob) HasParallelismDegree() bool {
	if o != nil && !IsNil(o.ParallelismDegree) {
		return true
	}

	return false
}

// SetParallelismDegree gets a reference to the given int32 and assigns it to the ParallelismDegree field.
func (o *MaskingJob) SetParallelismDegree(v int32) {
	o.ParallelismDegree = &v
}

// GetSourceMaskingJobId returns the SourceMaskingJobId field value if set, zero value otherwise.
func (o *MaskingJob) GetSourceMaskingJobId() string {
	if o == nil || IsNil(o.SourceMaskingJobId) {
		var ret string
		return ret
	}
	return *o.SourceMaskingJobId
}

// GetSourceMaskingJobIdOk returns a tuple with the SourceMaskingJobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetSourceMaskingJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceMaskingJobId) {
		return nil, false
	}
	return o.SourceMaskingJobId, true
}

// HasSourceMaskingJobId returns a boolean if a field has been set.
func (o *MaskingJob) HasSourceMaskingJobId() bool {
	if o != nil && !IsNil(o.SourceMaskingJobId) {
		return true
	}

	return false
}

// SetSourceMaskingJobId gets a reference to the given string and assigns it to the SourceMaskingJobId field.
func (o *MaskingJob) SetSourceMaskingJobId(v string) {
	o.SourceMaskingJobId = &v
}

// GetEngineIds returns the EngineIds field value if set, zero value otherwise.
func (o *MaskingJob) GetEngineIds() []string {
	if o == nil || IsNil(o.EngineIds) {
		var ret []string
		return ret
	}
	return o.EngineIds
}

// GetEngineIdsOk returns a tuple with the EngineIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetEngineIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.EngineIds) {
		return nil, false
	}
	return o.EngineIds, true
}

// HasEngineIds returns a boolean if a field has been set.
func (o *MaskingJob) HasEngineIds() bool {
	if o != nil && !IsNil(o.EngineIds) {
		return true
	}

	return false
}

// SetEngineIds gets a reference to the given []string and assigns it to the EngineIds field.
func (o *MaskingJob) SetEngineIds(v []string) {
	o.EngineIds = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MaskingJob) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJob) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MaskingJob) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *MaskingJob) SetTags(v []Tag) {
	o.Tags = v
}

func (o MaskingJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaskingJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Ruleset) {
		toSerialize["ruleset"] = o.Ruleset
	}
	if !IsNil(o.ConnectorType) {
		toSerialize["connector_type"] = o.ConnectorType
	}
	if !IsNil(o.IsOnTheFlyMasking) {
		toSerialize["is_on_the_fly_masking"] = o.IsOnTheFlyMasking
	}
	if !IsNil(o.CreationDate) {
		toSerialize["creation_date"] = o.CreationDate
	}
	if !IsNil(o.LastCompletedExecutionDate) {
		toSerialize["last_completed_execution_date"] = o.LastCompletedExecutionDate
	}
	if !IsNil(o.LastExecutionStatus) {
		toSerialize["last_execution_status"] = o.LastExecutionStatus
	}
	if !IsNil(o.LastExecutionId) {
		toSerialize["last_execution_id"] = o.LastExecutionId
	}
	if o.ConnectorUsername.IsSet() {
		toSerialize["connector_username"] = o.ConnectorUsername.Get()
	}
	if o.ConnectorPassword.IsSet() {
		toSerialize["connector_password"] = o.ConnectorPassword.Get()
	}
	if o.OnTheFlySourceConnectorUsername.IsSet() {
		toSerialize["on_the_fly_source_connector_username"] = o.OnTheFlySourceConnectorUsername.Get()
	}
	if o.OnTheFlySourceConnectorPassword.IsSet() {
		toSerialize["on_the_fly_source_connector_password"] = o.OnTheFlySourceConnectorPassword.Get()
	}
	if !IsNil(o.JobType) {
		toSerialize["job_type"] = o.JobType
	}
	if !IsNil(o.HyperscaleInstanceId) {
		toSerialize["hyperscale_instance_id"] = o.HyperscaleInstanceId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DatasetId) {
		toSerialize["dataset_id"] = o.DatasetId
	}
	if !IsNil(o.RetainExecutionData) {
		toSerialize["retain_execution_data"] = o.RetainExecutionData
	}
	if !IsNil(o.MaxMemory) {
		toSerialize["max_memory"] = o.MaxMemory
	}
	if !IsNil(o.MinMemory) {
		toSerialize["min_memory"] = o.MinMemory
	}
	if !IsNil(o.FeedbackSize) {
		toSerialize["feedback_size"] = o.FeedbackSize
	}
	if !IsNil(o.StreamRowLimit) {
		toSerialize["stream_row_limit"] = o.StreamRowLimit
	}
	if !IsNil(o.NumInputStreams) {
		toSerialize["num_input_streams"] = o.NumInputStreams
	}
	if !IsNil(o.MaxConcurrentSourceConnections) {
		toSerialize["max_concurrent_source_connections"] = o.MaxConcurrentSourceConnections
	}
	if !IsNil(o.MaxConcurrentTargetConnections) {
		toSerialize["max_concurrent_target_connections"] = o.MaxConcurrentTargetConnections
	}
	if !IsNil(o.ParallelismDegree) {
		toSerialize["parallelism_degree"] = o.ParallelismDegree
	}
	if !IsNil(o.SourceMaskingJobId) {
		toSerialize["source_masking_job_id"] = o.SourceMaskingJobId
	}
	if !IsNil(o.EngineIds) {
		toSerialize["engine_ids"] = o.EngineIds
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableMaskingJob struct {
	value *MaskingJob
	isSet bool
}

func (v NullableMaskingJob) Get() *MaskingJob {
	return v.value
}

func (v *NullableMaskingJob) Set(val *MaskingJob) {
	v.value = val
	v.isSet = true
}

func (v NullableMaskingJob) IsSet() bool {
	return v.isSet
}

func (v *NullableMaskingJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaskingJob(val *MaskingJob) *NullableMaskingJob {
	return &NullableMaskingJob{value: val, isSet: true}
}

func (v NullableMaskingJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaskingJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


