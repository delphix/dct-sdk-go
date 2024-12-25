/*
Delphix DCT API

Delphix DCT API

API version: 3.18.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the VirtualizationPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualizationPolicy{}

// VirtualizationPolicy struct for VirtualizationPolicy
type VirtualizationPolicy struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Whether this virtualization policy is managed by DCT or by an individual Delphix Engine.
	DctManaged *bool `json:"dct_managed,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	// The namespace id of this virtualization policy.
	NamespaceId *string `json:"namespace_id,omitempty"`
	// The namespace name of this virtualization policy.
	NamespaceName *string `json:"namespace_name,omitempty"`
	// Is this a replicated object.
	IsReplica *bool `json:"is_replica,omitempty"`
	EngineId *string `json:"engine_id,omitempty"`
	// The name of the engine the policy belongs to.
	EngineName *string `json:"engine_name,omitempty"`
	PolicyType *string `json:"policy_type,omitempty"`
	TimezoneId *string `json:"timezone_id,omitempty"`
	// True if this is the default policy created when the system is setup.
	DefaultPolicy *bool `json:"default_policy,omitempty"`
	// Whether this policy has been directly applied or inherited. See the effectivePolicies parameter of the list call for details.
	EffectiveType *string `json:"effective_type,omitempty"`
	// Amount of time to keep source data [Retention Policy].
	DataDuration *int32 `json:"data_duration,omitempty"`
	// Time unit for data_duration [Retention Policy].
	DataUnit *string `json:"data_unit,omitempty"`
	// Amount of time to keep log data [Retention Policy].
	LogDuration *int32 `json:"log_duration,omitempty"`
	// Time unit for log_duration [Retention Policy].
	LogUnit *string `json:"log_unit,omitempty"`
	// Number of daily snapshots to keep [Retention Policy].
	NumOfDaily *int32 `json:"num_of_daily,omitempty"`
	// Number of weekly snapshots to keep [Retention Policy].
	NumOfWeekly *int32 `json:"num_of_weekly,omitempty"`
	// Day of week upon which to enforce weekly snapshot retention [Retention Policy].
	DayOfWeek *string `json:"day_of_week,omitempty"`
	// Number of monthly snapshots to keep [Retention Policy].
	NumOfMonthly *int32 `json:"num_of_monthly,omitempty"`
	// Day of month upon which to enforce monthly snapshot retention [Retention Policy].
	DayOfMonth *int32 `json:"day_of_month,omitempty"`
	// Number of yearly snapshots to keep [Retention Policy].
	NumOfYearly *int32 `json:"num_of_yearly,omitempty"`
	// Day of year upon which to enforce yearly snapshot retention, expressed a month / day string (e.g., \"Jan 1\") [Retention Policy].
	DayOfYear *string `json:"day_of_year,omitempty"`
	Schedules []VirtualizationSchedule `json:"schedules,omitempty"`
	// Size of the quota, in bytes. (QUOTA_POLICY only).
	Size NullableInt64 `json:"size,omitempty"`
	// The tags that are applied to this VirtualizationPolicy.
	Tags []Tag `json:"tags,omitempty"`
}

// NewVirtualizationPolicy instantiates a new VirtualizationPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationPolicy() *VirtualizationPolicy {
	this := VirtualizationPolicy{}
	return &this
}

// NewVirtualizationPolicyWithDefaults instantiates a new VirtualizationPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationPolicyWithDefaults() *VirtualizationPolicy {
	this := VirtualizationPolicy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VirtualizationPolicy) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationPolicy) SetName(v string) {
	o.Name = &v
}

// GetDctManaged returns the DctManaged field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetDctManaged() bool {
	if o == nil || IsNil(o.DctManaged) {
		var ret bool
		return ret
	}
	return *o.DctManaged
}

// GetDctManagedOk returns a tuple with the DctManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetDctManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.DctManaged) {
		return nil, false
	}
	return o.DctManaged, true
}

// HasDctManaged returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasDctManaged() bool {
	if o != nil && !IsNil(o.DctManaged) {
		return true
	}

	return false
}

// SetDctManaged gets a reference to the given bool and assigns it to the DctManaged field.
func (o *VirtualizationPolicy) SetDctManaged(v bool) {
	o.DctManaged = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *VirtualizationPolicy) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNamespaceId returns the NamespaceId field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetNamespaceId() string {
	if o == nil || IsNil(o.NamespaceId) {
		var ret string
		return ret
	}
	return *o.NamespaceId
}

// GetNamespaceIdOk returns a tuple with the NamespaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetNamespaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.NamespaceId) {
		return nil, false
	}
	return o.NamespaceId, true
}

// HasNamespaceId returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasNamespaceId() bool {
	if o != nil && !IsNil(o.NamespaceId) {
		return true
	}

	return false
}

// SetNamespaceId gets a reference to the given string and assigns it to the NamespaceId field.
func (o *VirtualizationPolicy) SetNamespaceId(v string) {
	o.NamespaceId = &v
}

// GetNamespaceName returns the NamespaceName field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetNamespaceName() string {
	if o == nil || IsNil(o.NamespaceName) {
		var ret string
		return ret
	}
	return *o.NamespaceName
}

// GetNamespaceNameOk returns a tuple with the NamespaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetNamespaceNameOk() (*string, bool) {
	if o == nil || IsNil(o.NamespaceName) {
		return nil, false
	}
	return o.NamespaceName, true
}

// HasNamespaceName returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasNamespaceName() bool {
	if o != nil && !IsNil(o.NamespaceName) {
		return true
	}

	return false
}

// SetNamespaceName gets a reference to the given string and assigns it to the NamespaceName field.
func (o *VirtualizationPolicy) SetNamespaceName(v string) {
	o.NamespaceName = &v
}

// GetIsReplica returns the IsReplica field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetIsReplica() bool {
	if o == nil || IsNil(o.IsReplica) {
		var ret bool
		return ret
	}
	return *o.IsReplica
}

// GetIsReplicaOk returns a tuple with the IsReplica field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetIsReplicaOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReplica) {
		return nil, false
	}
	return o.IsReplica, true
}

// HasIsReplica returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasIsReplica() bool {
	if o != nil && !IsNil(o.IsReplica) {
		return true
	}

	return false
}

// SetIsReplica gets a reference to the given bool and assigns it to the IsReplica field.
func (o *VirtualizationPolicy) SetIsReplica(v bool) {
	o.IsReplica = &v
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetEngineId() string {
	if o == nil || IsNil(o.EngineId) {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *VirtualizationPolicy) SetEngineId(v string) {
	o.EngineId = &v
}

// GetEngineName returns the EngineName field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetEngineName() string {
	if o == nil || IsNil(o.EngineName) {
		var ret string
		return ret
	}
	return *o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetEngineNameOk() (*string, bool) {
	if o == nil || IsNil(o.EngineName) {
		return nil, false
	}
	return o.EngineName, true
}

// HasEngineName returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasEngineName() bool {
	if o != nil && !IsNil(o.EngineName) {
		return true
	}

	return false
}

// SetEngineName gets a reference to the given string and assigns it to the EngineName field.
func (o *VirtualizationPolicy) SetEngineName(v string) {
	o.EngineName = &v
}

// GetPolicyType returns the PolicyType field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetPolicyType() string {
	if o == nil || IsNil(o.PolicyType) {
		var ret string
		return ret
	}
	return *o.PolicyType
}

// GetPolicyTypeOk returns a tuple with the PolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetPolicyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyType) {
		return nil, false
	}
	return o.PolicyType, true
}

// HasPolicyType returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasPolicyType() bool {
	if o != nil && !IsNil(o.PolicyType) {
		return true
	}

	return false
}

// SetPolicyType gets a reference to the given string and assigns it to the PolicyType field.
func (o *VirtualizationPolicy) SetPolicyType(v string) {
	o.PolicyType = &v
}

// GetTimezoneId returns the TimezoneId field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetTimezoneId() string {
	if o == nil || IsNil(o.TimezoneId) {
		var ret string
		return ret
	}
	return *o.TimezoneId
}

// GetTimezoneIdOk returns a tuple with the TimezoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetTimezoneIdOk() (*string, bool) {
	if o == nil || IsNil(o.TimezoneId) {
		return nil, false
	}
	return o.TimezoneId, true
}

// HasTimezoneId returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasTimezoneId() bool {
	if o != nil && !IsNil(o.TimezoneId) {
		return true
	}

	return false
}

// SetTimezoneId gets a reference to the given string and assigns it to the TimezoneId field.
func (o *VirtualizationPolicy) SetTimezoneId(v string) {
	o.TimezoneId = &v
}

// GetDefaultPolicy returns the DefaultPolicy field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetDefaultPolicy() bool {
	if o == nil || IsNil(o.DefaultPolicy) {
		var ret bool
		return ret
	}
	return *o.DefaultPolicy
}

// GetDefaultPolicyOk returns a tuple with the DefaultPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetDefaultPolicyOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultPolicy) {
		return nil, false
	}
	return o.DefaultPolicy, true
}

// HasDefaultPolicy returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasDefaultPolicy() bool {
	if o != nil && !IsNil(o.DefaultPolicy) {
		return true
	}

	return false
}

// SetDefaultPolicy gets a reference to the given bool and assigns it to the DefaultPolicy field.
func (o *VirtualizationPolicy) SetDefaultPolicy(v bool) {
	o.DefaultPolicy = &v
}

// GetEffectiveType returns the EffectiveType field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetEffectiveType() string {
	if o == nil || IsNil(o.EffectiveType) {
		var ret string
		return ret
	}
	return *o.EffectiveType
}

// GetEffectiveTypeOk returns a tuple with the EffectiveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetEffectiveTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveType) {
		return nil, false
	}
	return o.EffectiveType, true
}

// HasEffectiveType returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasEffectiveType() bool {
	if o != nil && !IsNil(o.EffectiveType) {
		return true
	}

	return false
}

// SetEffectiveType gets a reference to the given string and assigns it to the EffectiveType field.
func (o *VirtualizationPolicy) SetEffectiveType(v string) {
	o.EffectiveType = &v
}

// GetDataDuration returns the DataDuration field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetDataDuration() int32 {
	if o == nil || IsNil(o.DataDuration) {
		var ret int32
		return ret
	}
	return *o.DataDuration
}

// GetDataDurationOk returns a tuple with the DataDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetDataDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.DataDuration) {
		return nil, false
	}
	return o.DataDuration, true
}

// HasDataDuration returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasDataDuration() bool {
	if o != nil && !IsNil(o.DataDuration) {
		return true
	}

	return false
}

// SetDataDuration gets a reference to the given int32 and assigns it to the DataDuration field.
func (o *VirtualizationPolicy) SetDataDuration(v int32) {
	o.DataDuration = &v
}

// GetDataUnit returns the DataUnit field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetDataUnit() string {
	if o == nil || IsNil(o.DataUnit) {
		var ret string
		return ret
	}
	return *o.DataUnit
}

// GetDataUnitOk returns a tuple with the DataUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetDataUnitOk() (*string, bool) {
	if o == nil || IsNil(o.DataUnit) {
		return nil, false
	}
	return o.DataUnit, true
}

// HasDataUnit returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasDataUnit() bool {
	if o != nil && !IsNil(o.DataUnit) {
		return true
	}

	return false
}

// SetDataUnit gets a reference to the given string and assigns it to the DataUnit field.
func (o *VirtualizationPolicy) SetDataUnit(v string) {
	o.DataUnit = &v
}

// GetLogDuration returns the LogDuration field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetLogDuration() int32 {
	if o == nil || IsNil(o.LogDuration) {
		var ret int32
		return ret
	}
	return *o.LogDuration
}

// GetLogDurationOk returns a tuple with the LogDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetLogDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.LogDuration) {
		return nil, false
	}
	return o.LogDuration, true
}

// HasLogDuration returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasLogDuration() bool {
	if o != nil && !IsNil(o.LogDuration) {
		return true
	}

	return false
}

// SetLogDuration gets a reference to the given int32 and assigns it to the LogDuration field.
func (o *VirtualizationPolicy) SetLogDuration(v int32) {
	o.LogDuration = &v
}

// GetLogUnit returns the LogUnit field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetLogUnit() string {
	if o == nil || IsNil(o.LogUnit) {
		var ret string
		return ret
	}
	return *o.LogUnit
}

// GetLogUnitOk returns a tuple with the LogUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetLogUnitOk() (*string, bool) {
	if o == nil || IsNil(o.LogUnit) {
		return nil, false
	}
	return o.LogUnit, true
}

// HasLogUnit returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasLogUnit() bool {
	if o != nil && !IsNil(o.LogUnit) {
		return true
	}

	return false
}

// SetLogUnit gets a reference to the given string and assigns it to the LogUnit field.
func (o *VirtualizationPolicy) SetLogUnit(v string) {
	o.LogUnit = &v
}

// GetNumOfDaily returns the NumOfDaily field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetNumOfDaily() int32 {
	if o == nil || IsNil(o.NumOfDaily) {
		var ret int32
		return ret
	}
	return *o.NumOfDaily
}

// GetNumOfDailyOk returns a tuple with the NumOfDaily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetNumOfDailyOk() (*int32, bool) {
	if o == nil || IsNil(o.NumOfDaily) {
		return nil, false
	}
	return o.NumOfDaily, true
}

// HasNumOfDaily returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasNumOfDaily() bool {
	if o != nil && !IsNil(o.NumOfDaily) {
		return true
	}

	return false
}

// SetNumOfDaily gets a reference to the given int32 and assigns it to the NumOfDaily field.
func (o *VirtualizationPolicy) SetNumOfDaily(v int32) {
	o.NumOfDaily = &v
}

// GetNumOfWeekly returns the NumOfWeekly field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetNumOfWeekly() int32 {
	if o == nil || IsNil(o.NumOfWeekly) {
		var ret int32
		return ret
	}
	return *o.NumOfWeekly
}

// GetNumOfWeeklyOk returns a tuple with the NumOfWeekly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetNumOfWeeklyOk() (*int32, bool) {
	if o == nil || IsNil(o.NumOfWeekly) {
		return nil, false
	}
	return o.NumOfWeekly, true
}

// HasNumOfWeekly returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasNumOfWeekly() bool {
	if o != nil && !IsNil(o.NumOfWeekly) {
		return true
	}

	return false
}

// SetNumOfWeekly gets a reference to the given int32 and assigns it to the NumOfWeekly field.
func (o *VirtualizationPolicy) SetNumOfWeekly(v int32) {
	o.NumOfWeekly = &v
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetDayOfWeek() string {
	if o == nil || IsNil(o.DayOfWeek) {
		var ret string
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetDayOfWeekOk() (*string, bool) {
	if o == nil || IsNil(o.DayOfWeek) {
		return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasDayOfWeek() bool {
	if o != nil && !IsNil(o.DayOfWeek) {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given string and assigns it to the DayOfWeek field.
func (o *VirtualizationPolicy) SetDayOfWeek(v string) {
	o.DayOfWeek = &v
}

// GetNumOfMonthly returns the NumOfMonthly field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetNumOfMonthly() int32 {
	if o == nil || IsNil(o.NumOfMonthly) {
		var ret int32
		return ret
	}
	return *o.NumOfMonthly
}

// GetNumOfMonthlyOk returns a tuple with the NumOfMonthly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetNumOfMonthlyOk() (*int32, bool) {
	if o == nil || IsNil(o.NumOfMonthly) {
		return nil, false
	}
	return o.NumOfMonthly, true
}

// HasNumOfMonthly returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasNumOfMonthly() bool {
	if o != nil && !IsNil(o.NumOfMonthly) {
		return true
	}

	return false
}

// SetNumOfMonthly gets a reference to the given int32 and assigns it to the NumOfMonthly field.
func (o *VirtualizationPolicy) SetNumOfMonthly(v int32) {
	o.NumOfMonthly = &v
}

// GetDayOfMonth returns the DayOfMonth field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetDayOfMonth() int32 {
	if o == nil || IsNil(o.DayOfMonth) {
		var ret int32
		return ret
	}
	return *o.DayOfMonth
}

// GetDayOfMonthOk returns a tuple with the DayOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetDayOfMonthOk() (*int32, bool) {
	if o == nil || IsNil(o.DayOfMonth) {
		return nil, false
	}
	return o.DayOfMonth, true
}

// HasDayOfMonth returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasDayOfMonth() bool {
	if o != nil && !IsNil(o.DayOfMonth) {
		return true
	}

	return false
}

// SetDayOfMonth gets a reference to the given int32 and assigns it to the DayOfMonth field.
func (o *VirtualizationPolicy) SetDayOfMonth(v int32) {
	o.DayOfMonth = &v
}

// GetNumOfYearly returns the NumOfYearly field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetNumOfYearly() int32 {
	if o == nil || IsNil(o.NumOfYearly) {
		var ret int32
		return ret
	}
	return *o.NumOfYearly
}

// GetNumOfYearlyOk returns a tuple with the NumOfYearly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetNumOfYearlyOk() (*int32, bool) {
	if o == nil || IsNil(o.NumOfYearly) {
		return nil, false
	}
	return o.NumOfYearly, true
}

// HasNumOfYearly returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasNumOfYearly() bool {
	if o != nil && !IsNil(o.NumOfYearly) {
		return true
	}

	return false
}

// SetNumOfYearly gets a reference to the given int32 and assigns it to the NumOfYearly field.
func (o *VirtualizationPolicy) SetNumOfYearly(v int32) {
	o.NumOfYearly = &v
}

// GetDayOfYear returns the DayOfYear field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetDayOfYear() string {
	if o == nil || IsNil(o.DayOfYear) {
		var ret string
		return ret
	}
	return *o.DayOfYear
}

// GetDayOfYearOk returns a tuple with the DayOfYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetDayOfYearOk() (*string, bool) {
	if o == nil || IsNil(o.DayOfYear) {
		return nil, false
	}
	return o.DayOfYear, true
}

// HasDayOfYear returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasDayOfYear() bool {
	if o != nil && !IsNil(o.DayOfYear) {
		return true
	}

	return false
}

// SetDayOfYear gets a reference to the given string and assigns it to the DayOfYear field.
func (o *VirtualizationPolicy) SetDayOfYear(v string) {
	o.DayOfYear = &v
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetSchedules() []VirtualizationSchedule {
	if o == nil || IsNil(o.Schedules) {
		var ret []VirtualizationSchedule
		return ret
	}
	return o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetSchedulesOk() ([]VirtualizationSchedule, bool) {
	if o == nil || IsNil(o.Schedules) {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasSchedules() bool {
	if o != nil && !IsNil(o.Schedules) {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given []VirtualizationSchedule and assigns it to the Schedules field.
func (o *VirtualizationPolicy) SetSchedules(v []VirtualizationSchedule) {
	o.Schedules = v
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationPolicy) GetSize() int64 {
	if o == nil || IsNil(o.Size.Get()) {
		var ret int64
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationPolicy) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt64 and assigns it to the Size field.
func (o *VirtualizationPolicy) SetSize(v int64) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *VirtualizationPolicy) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *VirtualizationPolicy) UnsetSize() {
	o.Size.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *VirtualizationPolicy) SetTags(v []Tag) {
	o.Tags = v
}

func (o VirtualizationPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualizationPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DctManaged) {
		toSerialize["dct_managed"] = o.DctManaged
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.NamespaceId) {
		toSerialize["namespace_id"] = o.NamespaceId
	}
	if !IsNil(o.NamespaceName) {
		toSerialize["namespace_name"] = o.NamespaceName
	}
	if !IsNil(o.IsReplica) {
		toSerialize["is_replica"] = o.IsReplica
	}
	if !IsNil(o.EngineId) {
		toSerialize["engine_id"] = o.EngineId
	}
	if !IsNil(o.EngineName) {
		toSerialize["engine_name"] = o.EngineName
	}
	if !IsNil(o.PolicyType) {
		toSerialize["policy_type"] = o.PolicyType
	}
	if !IsNil(o.TimezoneId) {
		toSerialize["timezone_id"] = o.TimezoneId
	}
	if !IsNil(o.DefaultPolicy) {
		toSerialize["default_policy"] = o.DefaultPolicy
	}
	if !IsNil(o.EffectiveType) {
		toSerialize["effective_type"] = o.EffectiveType
	}
	if !IsNil(o.DataDuration) {
		toSerialize["data_duration"] = o.DataDuration
	}
	if !IsNil(o.DataUnit) {
		toSerialize["data_unit"] = o.DataUnit
	}
	if !IsNil(o.LogDuration) {
		toSerialize["log_duration"] = o.LogDuration
	}
	if !IsNil(o.LogUnit) {
		toSerialize["log_unit"] = o.LogUnit
	}
	if !IsNil(o.NumOfDaily) {
		toSerialize["num_of_daily"] = o.NumOfDaily
	}
	if !IsNil(o.NumOfWeekly) {
		toSerialize["num_of_weekly"] = o.NumOfWeekly
	}
	if !IsNil(o.DayOfWeek) {
		toSerialize["day_of_week"] = o.DayOfWeek
	}
	if !IsNil(o.NumOfMonthly) {
		toSerialize["num_of_monthly"] = o.NumOfMonthly
	}
	if !IsNil(o.DayOfMonth) {
		toSerialize["day_of_month"] = o.DayOfMonth
	}
	if !IsNil(o.NumOfYearly) {
		toSerialize["num_of_yearly"] = o.NumOfYearly
	}
	if !IsNil(o.DayOfYear) {
		toSerialize["day_of_year"] = o.DayOfYear
	}
	if !IsNil(o.Schedules) {
		toSerialize["schedules"] = o.Schedules
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableVirtualizationPolicy struct {
	value *VirtualizationPolicy
	isSet bool
}

func (v NullableVirtualizationPolicy) Get() *VirtualizationPolicy {
	return v.value
}

func (v *NullableVirtualizationPolicy) Set(val *VirtualizationPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationPolicy(val *VirtualizationPolicy) *NullableVirtualizationPolicy {
	return &NullableVirtualizationPolicy{value: val, isSet: true}
}

func (v NullableVirtualizationPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


