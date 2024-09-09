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
	"time"
)

// checks if the Timeflow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Timeflow{}

// Timeflow Virtualization Engine Timeflow of a dSource or VDB.
type Timeflow struct {
	// The Timeflow ID.
	Id *string `json:"id,omitempty"`
	// The ID of the engine the timeflow belongs to.
	EngineId *string `json:"engine_id,omitempty"`
	// Alternate namespace for this object, for replicated and restored timeflows.
	Namespace NullableString `json:"namespace,omitempty"`
	// The namespace id of this timeflows.
	NamespaceId NullableString `json:"namespace_id,omitempty"`
	// The namespace name of this timeflows.
	NamespaceName NullableString `json:"namespace_name,omitempty"`
	// Is this a replicated object.
	IsReplica *bool `json:"is_replica,omitempty"`
	// The timeflow's name.
	Name *string `json:"name,omitempty"`
	// The ID of the timeflow's dSource or VDB.
	DatasetId *string `json:"dataset_id,omitempty"`
	// The source action that created the timeflow.
	CreationType *string `json:"creation_type,omitempty"`
	// The ID of the timeflow's parent snapshot.
	ParentSnapshotId *string `json:"parent_snapshot_id,omitempty"`
	// The location on the parent timeflow from which this timeflow was provisioned. This will not be present for timeflows derived from linked sources.
	ParentPointLocation *string `json:"parent_point_location,omitempty"`
	// The timestamp on the parent timeflow from which this timeflow was provisioned. This will not be present for timeflows derived from linked sources.
	ParentPointTimestamp *time.Time `json:"parent_point_timestamp,omitempty"`
	// A reference to the parent timeflow from which this timeflow was provisioned. This will not be present for timeflows derived from linked sources.
	ParentPointTimeflowId *string `json:"parent_point_timeflow_id,omitempty"`
	// The ID of the parent VDB. This is mutually exclusive with parent_dsource_id.
	ParentVdbId *string `json:"parent_vdb_id,omitempty"`
	// The ID of the parent dSource. This is mutually exclusive with parent_vdb_id.
	ParentDsourceId *string `json:"parent_dsource_id,omitempty"`
	// The ID of the source VDB. This is mutually exclusive with source_dsource_id.
	SourceVdbId *string `json:"source_vdb_id,omitempty"`
	// The ID of the source dSource. This is mutually exclusive with source_vdb_id.
	SourceDsourceId *string `json:"source_dsource_id,omitempty"`
	// The timestamp on the root ancestor timeflow from which this timeflow originated. This logical time acts as reference to the origin source data.
	SourceDataTimestamp *time.Time `json:"source_data_timestamp,omitempty"`
	// Oracle-specific incarnation identifier for this timeflow.
	OracleIncarnationId *string `json:"oracle_incarnation_id,omitempty"`
	// A reference to the mirror CDB timeflow if this is a timeflow for a PDB.
	OracleCdbTimeflowId *string `json:"oracle_cdb_timeflow_id,omitempty"`
	// The unique identifier for timeflow-specific TDE objects that reside outside of Delphix storage.
	OracleTdeUuid *string `json:"oracle_tde_uuid,omitempty"`
	// MSSQL-specific recovery branch identifier for this timeflow.
	MssqlDatabaseGuid *string `json:"mssql_database_guid,omitempty"`
	// Whether this timeflow is currently active or not.
	IsActive *bool `json:"is_active,omitempty"`
	// The time when the timeflow was created.
	CreationTimestamp *time.Time `json:"creation_timestamp,omitempty"`
	// The time when this timeflow became active.
	ActivationTimestamp *time.Time `json:"activation_timestamp,omitempty"`
	Tags []Tag `json:"tags,omitempty"`
}

// NewTimeflow instantiates a new Timeflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeflow() *Timeflow {
	this := Timeflow{}
	return &this
}

// NewTimeflowWithDefaults instantiates a new Timeflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeflowWithDefaults() *Timeflow {
	this := Timeflow{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Timeflow) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Timeflow) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Timeflow) SetId(v string) {
	o.Id = &v
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *Timeflow) GetEngineId() string {
	if o == nil || IsNil(o.EngineId) {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *Timeflow) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *Timeflow) SetEngineId(v string) {
	o.EngineId = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Timeflow) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Timeflow) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *Timeflow) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *Timeflow) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *Timeflow) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *Timeflow) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetNamespaceId returns the NamespaceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Timeflow) GetNamespaceId() string {
	if o == nil || IsNil(o.NamespaceId.Get()) {
		var ret string
		return ret
	}
	return *o.NamespaceId.Get()
}

// GetNamespaceIdOk returns a tuple with the NamespaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Timeflow) GetNamespaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NamespaceId.Get(), o.NamespaceId.IsSet()
}

// HasNamespaceId returns a boolean if a field has been set.
func (o *Timeflow) HasNamespaceId() bool {
	if o != nil && o.NamespaceId.IsSet() {
		return true
	}

	return false
}

// SetNamespaceId gets a reference to the given NullableString and assigns it to the NamespaceId field.
func (o *Timeflow) SetNamespaceId(v string) {
	o.NamespaceId.Set(&v)
}
// SetNamespaceIdNil sets the value for NamespaceId to be an explicit nil
func (o *Timeflow) SetNamespaceIdNil() {
	o.NamespaceId.Set(nil)
}

// UnsetNamespaceId ensures that no value is present for NamespaceId, not even an explicit nil
func (o *Timeflow) UnsetNamespaceId() {
	o.NamespaceId.Unset()
}

// GetNamespaceName returns the NamespaceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Timeflow) GetNamespaceName() string {
	if o == nil || IsNil(o.NamespaceName.Get()) {
		var ret string
		return ret
	}
	return *o.NamespaceName.Get()
}

// GetNamespaceNameOk returns a tuple with the NamespaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Timeflow) GetNamespaceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NamespaceName.Get(), o.NamespaceName.IsSet()
}

// HasNamespaceName returns a boolean if a field has been set.
func (o *Timeflow) HasNamespaceName() bool {
	if o != nil && o.NamespaceName.IsSet() {
		return true
	}

	return false
}

// SetNamespaceName gets a reference to the given NullableString and assigns it to the NamespaceName field.
func (o *Timeflow) SetNamespaceName(v string) {
	o.NamespaceName.Set(&v)
}
// SetNamespaceNameNil sets the value for NamespaceName to be an explicit nil
func (o *Timeflow) SetNamespaceNameNil() {
	o.NamespaceName.Set(nil)
}

// UnsetNamespaceName ensures that no value is present for NamespaceName, not even an explicit nil
func (o *Timeflow) UnsetNamespaceName() {
	o.NamespaceName.Unset()
}

// GetIsReplica returns the IsReplica field value if set, zero value otherwise.
func (o *Timeflow) GetIsReplica() bool {
	if o == nil || IsNil(o.IsReplica) {
		var ret bool
		return ret
	}
	return *o.IsReplica
}

// GetIsReplicaOk returns a tuple with the IsReplica field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetIsReplicaOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReplica) {
		return nil, false
	}
	return o.IsReplica, true
}

// HasIsReplica returns a boolean if a field has been set.
func (o *Timeflow) HasIsReplica() bool {
	if o != nil && !IsNil(o.IsReplica) {
		return true
	}

	return false
}

// SetIsReplica gets a reference to the given bool and assigns it to the IsReplica field.
func (o *Timeflow) SetIsReplica(v bool) {
	o.IsReplica = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Timeflow) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Timeflow) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Timeflow) SetName(v string) {
	o.Name = &v
}

// GetDatasetId returns the DatasetId field value if set, zero value otherwise.
func (o *Timeflow) GetDatasetId() string {
	if o == nil || IsNil(o.DatasetId) {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetDatasetIdOk() (*string, bool) {
	if o == nil || IsNil(o.DatasetId) {
		return nil, false
	}
	return o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *Timeflow) HasDatasetId() bool {
	if o != nil && !IsNil(o.DatasetId) {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *Timeflow) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetCreationType returns the CreationType field value if set, zero value otherwise.
func (o *Timeflow) GetCreationType() string {
	if o == nil || IsNil(o.CreationType) {
		var ret string
		return ret
	}
	return *o.CreationType
}

// GetCreationTypeOk returns a tuple with the CreationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetCreationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CreationType) {
		return nil, false
	}
	return o.CreationType, true
}

// HasCreationType returns a boolean if a field has been set.
func (o *Timeflow) HasCreationType() bool {
	if o != nil && !IsNil(o.CreationType) {
		return true
	}

	return false
}

// SetCreationType gets a reference to the given string and assigns it to the CreationType field.
func (o *Timeflow) SetCreationType(v string) {
	o.CreationType = &v
}

// GetParentSnapshotId returns the ParentSnapshotId field value if set, zero value otherwise.
func (o *Timeflow) GetParentSnapshotId() string {
	if o == nil || IsNil(o.ParentSnapshotId) {
		var ret string
		return ret
	}
	return *o.ParentSnapshotId
}

// GetParentSnapshotIdOk returns a tuple with the ParentSnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetParentSnapshotIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentSnapshotId) {
		return nil, false
	}
	return o.ParentSnapshotId, true
}

// HasParentSnapshotId returns a boolean if a field has been set.
func (o *Timeflow) HasParentSnapshotId() bool {
	if o != nil && !IsNil(o.ParentSnapshotId) {
		return true
	}

	return false
}

// SetParentSnapshotId gets a reference to the given string and assigns it to the ParentSnapshotId field.
func (o *Timeflow) SetParentSnapshotId(v string) {
	o.ParentSnapshotId = &v
}

// GetParentPointLocation returns the ParentPointLocation field value if set, zero value otherwise.
func (o *Timeflow) GetParentPointLocation() string {
	if o == nil || IsNil(o.ParentPointLocation) {
		var ret string
		return ret
	}
	return *o.ParentPointLocation
}

// GetParentPointLocationOk returns a tuple with the ParentPointLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetParentPointLocationOk() (*string, bool) {
	if o == nil || IsNil(o.ParentPointLocation) {
		return nil, false
	}
	return o.ParentPointLocation, true
}

// HasParentPointLocation returns a boolean if a field has been set.
func (o *Timeflow) HasParentPointLocation() bool {
	if o != nil && !IsNil(o.ParentPointLocation) {
		return true
	}

	return false
}

// SetParentPointLocation gets a reference to the given string and assigns it to the ParentPointLocation field.
func (o *Timeflow) SetParentPointLocation(v string) {
	o.ParentPointLocation = &v
}

// GetParentPointTimestamp returns the ParentPointTimestamp field value if set, zero value otherwise.
func (o *Timeflow) GetParentPointTimestamp() time.Time {
	if o == nil || IsNil(o.ParentPointTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.ParentPointTimestamp
}

// GetParentPointTimestampOk returns a tuple with the ParentPointTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetParentPointTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ParentPointTimestamp) {
		return nil, false
	}
	return o.ParentPointTimestamp, true
}

// HasParentPointTimestamp returns a boolean if a field has been set.
func (o *Timeflow) HasParentPointTimestamp() bool {
	if o != nil && !IsNil(o.ParentPointTimestamp) {
		return true
	}

	return false
}

// SetParentPointTimestamp gets a reference to the given time.Time and assigns it to the ParentPointTimestamp field.
func (o *Timeflow) SetParentPointTimestamp(v time.Time) {
	o.ParentPointTimestamp = &v
}

// GetParentPointTimeflowId returns the ParentPointTimeflowId field value if set, zero value otherwise.
func (o *Timeflow) GetParentPointTimeflowId() string {
	if o == nil || IsNil(o.ParentPointTimeflowId) {
		var ret string
		return ret
	}
	return *o.ParentPointTimeflowId
}

// GetParentPointTimeflowIdOk returns a tuple with the ParentPointTimeflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetParentPointTimeflowIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentPointTimeflowId) {
		return nil, false
	}
	return o.ParentPointTimeflowId, true
}

// HasParentPointTimeflowId returns a boolean if a field has been set.
func (o *Timeflow) HasParentPointTimeflowId() bool {
	if o != nil && !IsNil(o.ParentPointTimeflowId) {
		return true
	}

	return false
}

// SetParentPointTimeflowId gets a reference to the given string and assigns it to the ParentPointTimeflowId field.
func (o *Timeflow) SetParentPointTimeflowId(v string) {
	o.ParentPointTimeflowId = &v
}

// GetParentVdbId returns the ParentVdbId field value if set, zero value otherwise.
func (o *Timeflow) GetParentVdbId() string {
	if o == nil || IsNil(o.ParentVdbId) {
		var ret string
		return ret
	}
	return *o.ParentVdbId
}

// GetParentVdbIdOk returns a tuple with the ParentVdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetParentVdbIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentVdbId) {
		return nil, false
	}
	return o.ParentVdbId, true
}

// HasParentVdbId returns a boolean if a field has been set.
func (o *Timeflow) HasParentVdbId() bool {
	if o != nil && !IsNil(o.ParentVdbId) {
		return true
	}

	return false
}

// SetParentVdbId gets a reference to the given string and assigns it to the ParentVdbId field.
func (o *Timeflow) SetParentVdbId(v string) {
	o.ParentVdbId = &v
}

// GetParentDsourceId returns the ParentDsourceId field value if set, zero value otherwise.
func (o *Timeflow) GetParentDsourceId() string {
	if o == nil || IsNil(o.ParentDsourceId) {
		var ret string
		return ret
	}
	return *o.ParentDsourceId
}

// GetParentDsourceIdOk returns a tuple with the ParentDsourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetParentDsourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentDsourceId) {
		return nil, false
	}
	return o.ParentDsourceId, true
}

// HasParentDsourceId returns a boolean if a field has been set.
func (o *Timeflow) HasParentDsourceId() bool {
	if o != nil && !IsNil(o.ParentDsourceId) {
		return true
	}

	return false
}

// SetParentDsourceId gets a reference to the given string and assigns it to the ParentDsourceId field.
func (o *Timeflow) SetParentDsourceId(v string) {
	o.ParentDsourceId = &v
}

// GetSourceVdbId returns the SourceVdbId field value if set, zero value otherwise.
func (o *Timeflow) GetSourceVdbId() string {
	if o == nil || IsNil(o.SourceVdbId) {
		var ret string
		return ret
	}
	return *o.SourceVdbId
}

// GetSourceVdbIdOk returns a tuple with the SourceVdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetSourceVdbIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceVdbId) {
		return nil, false
	}
	return o.SourceVdbId, true
}

// HasSourceVdbId returns a boolean if a field has been set.
func (o *Timeflow) HasSourceVdbId() bool {
	if o != nil && !IsNil(o.SourceVdbId) {
		return true
	}

	return false
}

// SetSourceVdbId gets a reference to the given string and assigns it to the SourceVdbId field.
func (o *Timeflow) SetSourceVdbId(v string) {
	o.SourceVdbId = &v
}

// GetSourceDsourceId returns the SourceDsourceId field value if set, zero value otherwise.
func (o *Timeflow) GetSourceDsourceId() string {
	if o == nil || IsNil(o.SourceDsourceId) {
		var ret string
		return ret
	}
	return *o.SourceDsourceId
}

// GetSourceDsourceIdOk returns a tuple with the SourceDsourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetSourceDsourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceDsourceId) {
		return nil, false
	}
	return o.SourceDsourceId, true
}

// HasSourceDsourceId returns a boolean if a field has been set.
func (o *Timeflow) HasSourceDsourceId() bool {
	if o != nil && !IsNil(o.SourceDsourceId) {
		return true
	}

	return false
}

// SetSourceDsourceId gets a reference to the given string and assigns it to the SourceDsourceId field.
func (o *Timeflow) SetSourceDsourceId(v string) {
	o.SourceDsourceId = &v
}

// GetSourceDataTimestamp returns the SourceDataTimestamp field value if set, zero value otherwise.
func (o *Timeflow) GetSourceDataTimestamp() time.Time {
	if o == nil || IsNil(o.SourceDataTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.SourceDataTimestamp
}

// GetSourceDataTimestampOk returns a tuple with the SourceDataTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetSourceDataTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SourceDataTimestamp) {
		return nil, false
	}
	return o.SourceDataTimestamp, true
}

// HasSourceDataTimestamp returns a boolean if a field has been set.
func (o *Timeflow) HasSourceDataTimestamp() bool {
	if o != nil && !IsNil(o.SourceDataTimestamp) {
		return true
	}

	return false
}

// SetSourceDataTimestamp gets a reference to the given time.Time and assigns it to the SourceDataTimestamp field.
func (o *Timeflow) SetSourceDataTimestamp(v time.Time) {
	o.SourceDataTimestamp = &v
}

// GetOracleIncarnationId returns the OracleIncarnationId field value if set, zero value otherwise.
func (o *Timeflow) GetOracleIncarnationId() string {
	if o == nil || IsNil(o.OracleIncarnationId) {
		var ret string
		return ret
	}
	return *o.OracleIncarnationId
}

// GetOracleIncarnationIdOk returns a tuple with the OracleIncarnationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetOracleIncarnationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OracleIncarnationId) {
		return nil, false
	}
	return o.OracleIncarnationId, true
}

// HasOracleIncarnationId returns a boolean if a field has been set.
func (o *Timeflow) HasOracleIncarnationId() bool {
	if o != nil && !IsNil(o.OracleIncarnationId) {
		return true
	}

	return false
}

// SetOracleIncarnationId gets a reference to the given string and assigns it to the OracleIncarnationId field.
func (o *Timeflow) SetOracleIncarnationId(v string) {
	o.OracleIncarnationId = &v
}

// GetOracleCdbTimeflowId returns the OracleCdbTimeflowId field value if set, zero value otherwise.
func (o *Timeflow) GetOracleCdbTimeflowId() string {
	if o == nil || IsNil(o.OracleCdbTimeflowId) {
		var ret string
		return ret
	}
	return *o.OracleCdbTimeflowId
}

// GetOracleCdbTimeflowIdOk returns a tuple with the OracleCdbTimeflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetOracleCdbTimeflowIdOk() (*string, bool) {
	if o == nil || IsNil(o.OracleCdbTimeflowId) {
		return nil, false
	}
	return o.OracleCdbTimeflowId, true
}

// HasOracleCdbTimeflowId returns a boolean if a field has been set.
func (o *Timeflow) HasOracleCdbTimeflowId() bool {
	if o != nil && !IsNil(o.OracleCdbTimeflowId) {
		return true
	}

	return false
}

// SetOracleCdbTimeflowId gets a reference to the given string and assigns it to the OracleCdbTimeflowId field.
func (o *Timeflow) SetOracleCdbTimeflowId(v string) {
	o.OracleCdbTimeflowId = &v
}

// GetOracleTdeUuid returns the OracleTdeUuid field value if set, zero value otherwise.
func (o *Timeflow) GetOracleTdeUuid() string {
	if o == nil || IsNil(o.OracleTdeUuid) {
		var ret string
		return ret
	}
	return *o.OracleTdeUuid
}

// GetOracleTdeUuidOk returns a tuple with the OracleTdeUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetOracleTdeUuidOk() (*string, bool) {
	if o == nil || IsNil(o.OracleTdeUuid) {
		return nil, false
	}
	return o.OracleTdeUuid, true
}

// HasOracleTdeUuid returns a boolean if a field has been set.
func (o *Timeflow) HasOracleTdeUuid() bool {
	if o != nil && !IsNil(o.OracleTdeUuid) {
		return true
	}

	return false
}

// SetOracleTdeUuid gets a reference to the given string and assigns it to the OracleTdeUuid field.
func (o *Timeflow) SetOracleTdeUuid(v string) {
	o.OracleTdeUuid = &v
}

// GetMssqlDatabaseGuid returns the MssqlDatabaseGuid field value if set, zero value otherwise.
func (o *Timeflow) GetMssqlDatabaseGuid() string {
	if o == nil || IsNil(o.MssqlDatabaseGuid) {
		var ret string
		return ret
	}
	return *o.MssqlDatabaseGuid
}

// GetMssqlDatabaseGuidOk returns a tuple with the MssqlDatabaseGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetMssqlDatabaseGuidOk() (*string, bool) {
	if o == nil || IsNil(o.MssqlDatabaseGuid) {
		return nil, false
	}
	return o.MssqlDatabaseGuid, true
}

// HasMssqlDatabaseGuid returns a boolean if a field has been set.
func (o *Timeflow) HasMssqlDatabaseGuid() bool {
	if o != nil && !IsNil(o.MssqlDatabaseGuid) {
		return true
	}

	return false
}

// SetMssqlDatabaseGuid gets a reference to the given string and assigns it to the MssqlDatabaseGuid field.
func (o *Timeflow) SetMssqlDatabaseGuid(v string) {
	o.MssqlDatabaseGuid = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *Timeflow) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *Timeflow) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *Timeflow) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetCreationTimestamp returns the CreationTimestamp field value if set, zero value otherwise.
func (o *Timeflow) GetCreationTimestamp() time.Time {
	if o == nil || IsNil(o.CreationTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetCreationTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTimestamp) {
		return nil, false
	}
	return o.CreationTimestamp, true
}

// HasCreationTimestamp returns a boolean if a field has been set.
func (o *Timeflow) HasCreationTimestamp() bool {
	if o != nil && !IsNil(o.CreationTimestamp) {
		return true
	}

	return false
}

// SetCreationTimestamp gets a reference to the given time.Time and assigns it to the CreationTimestamp field.
func (o *Timeflow) SetCreationTimestamp(v time.Time) {
	o.CreationTimestamp = &v
}

// GetActivationTimestamp returns the ActivationTimestamp field value if set, zero value otherwise.
func (o *Timeflow) GetActivationTimestamp() time.Time {
	if o == nil || IsNil(o.ActivationTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.ActivationTimestamp
}

// GetActivationTimestampOk returns a tuple with the ActivationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetActivationTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ActivationTimestamp) {
		return nil, false
	}
	return o.ActivationTimestamp, true
}

// HasActivationTimestamp returns a boolean if a field has been set.
func (o *Timeflow) HasActivationTimestamp() bool {
	if o != nil && !IsNil(o.ActivationTimestamp) {
		return true
	}

	return false
}

// SetActivationTimestamp gets a reference to the given time.Time and assigns it to the ActivationTimestamp field.
func (o *Timeflow) SetActivationTimestamp(v time.Time) {
	o.ActivationTimestamp = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Timeflow) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeflow) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Timeflow) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *Timeflow) SetTags(v []Tag) {
	o.Tags = v
}

func (o Timeflow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Timeflow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.EngineId) {
		toSerialize["engine_id"] = o.EngineId
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.NamespaceId.IsSet() {
		toSerialize["namespace_id"] = o.NamespaceId.Get()
	}
	if o.NamespaceName.IsSet() {
		toSerialize["namespace_name"] = o.NamespaceName.Get()
	}
	if !IsNil(o.IsReplica) {
		toSerialize["is_replica"] = o.IsReplica
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DatasetId) {
		toSerialize["dataset_id"] = o.DatasetId
	}
	if !IsNil(o.CreationType) {
		toSerialize["creation_type"] = o.CreationType
	}
	if !IsNil(o.ParentSnapshotId) {
		toSerialize["parent_snapshot_id"] = o.ParentSnapshotId
	}
	if !IsNil(o.ParentPointLocation) {
		toSerialize["parent_point_location"] = o.ParentPointLocation
	}
	if !IsNil(o.ParentPointTimestamp) {
		toSerialize["parent_point_timestamp"] = o.ParentPointTimestamp
	}
	if !IsNil(o.ParentPointTimeflowId) {
		toSerialize["parent_point_timeflow_id"] = o.ParentPointTimeflowId
	}
	if !IsNil(o.ParentVdbId) {
		toSerialize["parent_vdb_id"] = o.ParentVdbId
	}
	if !IsNil(o.ParentDsourceId) {
		toSerialize["parent_dsource_id"] = o.ParentDsourceId
	}
	if !IsNil(o.SourceVdbId) {
		toSerialize["source_vdb_id"] = o.SourceVdbId
	}
	if !IsNil(o.SourceDsourceId) {
		toSerialize["source_dsource_id"] = o.SourceDsourceId
	}
	if !IsNil(o.SourceDataTimestamp) {
		toSerialize["source_data_timestamp"] = o.SourceDataTimestamp
	}
	if !IsNil(o.OracleIncarnationId) {
		toSerialize["oracle_incarnation_id"] = o.OracleIncarnationId
	}
	if !IsNil(o.OracleCdbTimeflowId) {
		toSerialize["oracle_cdb_timeflow_id"] = o.OracleCdbTimeflowId
	}
	if !IsNil(o.OracleTdeUuid) {
		toSerialize["oracle_tde_uuid"] = o.OracleTdeUuid
	}
	if !IsNil(o.MssqlDatabaseGuid) {
		toSerialize["mssql_database_guid"] = o.MssqlDatabaseGuid
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.CreationTimestamp) {
		toSerialize["creation_timestamp"] = o.CreationTimestamp
	}
	if !IsNil(o.ActivationTimestamp) {
		toSerialize["activation_timestamp"] = o.ActivationTimestamp
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableTimeflow struct {
	value *Timeflow
	isSet bool
}

func (v NullableTimeflow) Get() *Timeflow {
	return v.value
}

func (v *NullableTimeflow) Set(val *Timeflow) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeflow) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeflow(val *Timeflow) *NullableTimeflow {
	return &NullableTimeflow{value: val, isSet: true}
}

func (v NullableTimeflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


