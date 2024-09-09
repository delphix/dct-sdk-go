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

// checks if the VirtualDatasetHooks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualDatasetHooks{}

// VirtualDatasetHooks VDB operation hooks.
type VirtualDatasetHooks struct {
	// The commands to execute on the target environment before refreshing the VDB.
	PreRefresh []Hook `json:"pre_refresh,omitempty"`
	// The commands to execute on the target environment after refreshing the VDB.
	PostRefresh []Hook `json:"post_refresh,omitempty"`
	// The commands to execute on the target environment before refreshing the VDB with data from itself.
	PreSelfRefresh []Hook `json:"pre_self_refresh,omitempty"`
	// The commands to execute on the target environment after refreshing the VDB with data from itself.
	PostSelfRefresh []Hook `json:"post_self_refresh,omitempty"`
	// The commands to execute on the target environment before rewinding the VDB.
	// Deprecated
	PreRollback []Hook `json:"pre_rollback,omitempty"`
	// The commands to execute on the target environment after rewinding the VDB.
	// Deprecated
	PostRollback []Hook `json:"post_rollback,omitempty"`
	// The commands to execute on the target environment when the VDB is created or refreshed.
	ConfigureClone []Hook `json:"configure_clone,omitempty"`
	// The commands to execute on the target environment before snapshotting a virtual source. These commands can quiesce any data prior to snapshotting.
	PreSnapshot []Hook `json:"pre_snapshot,omitempty"`
	// The commands to execute on the target environment after snapshotting a virtual source.
	PostSnapshot []Hook `json:"post_snapshot,omitempty"`
	// The commands to execute on the target environment before starting a virtual source.
	PreStart []Hook `json:"pre_start,omitempty"`
	// The commands to execute on the target environment after starting a virtual source.
	PostStart []Hook `json:"post_start,omitempty"`
	// The commands to execute on the target environment before stopping a virtual source.
	PreStop []Hook `json:"pre_stop,omitempty"`
	// The commands to execute on the target environment after stopping a virtual source.
	PostStop []Hook `json:"post_stop,omitempty"`
}

// NewVirtualDatasetHooks instantiates a new VirtualDatasetHooks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualDatasetHooks() *VirtualDatasetHooks {
	this := VirtualDatasetHooks{}
	return &this
}

// NewVirtualDatasetHooksWithDefaults instantiates a new VirtualDatasetHooks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualDatasetHooksWithDefaults() *VirtualDatasetHooks {
	this := VirtualDatasetHooks{}
	return &this
}

// GetPreRefresh returns the PreRefresh field value if set, zero value otherwise.
func (o *VirtualDatasetHooks) GetPreRefresh() []Hook {
	if o == nil || IsNil(o.PreRefresh) {
		var ret []Hook
		return ret
	}
	return o.PreRefresh
}

// GetPreRefreshOk returns a tuple with the PreRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDatasetHooks) GetPreRefreshOk() ([]Hook, bool) {
	if o == nil || IsNil(o.PreRefresh) {
		return nil, false
	}
	return o.PreRefresh, true
}

// HasPreRefresh returns a boolean if a field has been set.
func (o *VirtualDatasetHooks) HasPreRefresh() bool {
	if o != nil && !IsNil(o.PreRefresh) {
		return true
	}

	return false
}

// SetPreRefresh gets a reference to the given []Hook and assigns it to the PreRefresh field.
func (o *VirtualDatasetHooks) SetPreRefresh(v []Hook) {
	o.PreRefresh = v
}

// GetPostRefresh returns the PostRefresh field value if set, zero value otherwise.
func (o *VirtualDatasetHooks) GetPostRefresh() []Hook {
	if o == nil || IsNil(o.PostRefresh) {
		var ret []Hook
		return ret
	}
	return o.PostRefresh
}

// GetPostRefreshOk returns a tuple with the PostRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDatasetHooks) GetPostRefreshOk() ([]Hook, bool) {
	if o == nil || IsNil(o.PostRefresh) {
		return nil, false
	}
	return o.PostRefresh, true
}

// HasPostRefresh returns a boolean if a field has been set.
func (o *VirtualDatasetHooks) HasPostRefresh() bool {
	if o != nil && !IsNil(o.PostRefresh) {
		return true
	}

	return false
}

// SetPostRefresh gets a reference to the given []Hook and assigns it to the PostRefresh field.
func (o *VirtualDatasetHooks) SetPostRefresh(v []Hook) {
	o.PostRefresh = v
}

// GetPreSelfRefresh returns the PreSelfRefresh field value if set, zero value otherwise.
func (o *VirtualDatasetHooks) GetPreSelfRefresh() []Hook {
	if o == nil || IsNil(o.PreSelfRefresh) {
		var ret []Hook
		return ret
	}
	return o.PreSelfRefresh
}

// GetPreSelfRefreshOk returns a tuple with the PreSelfRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDatasetHooks) GetPreSelfRefreshOk() ([]Hook, bool) {
	if o == nil || IsNil(o.PreSelfRefresh) {
		return nil, false
	}
	return o.PreSelfRefresh, true
}

// HasPreSelfRefresh returns a boolean if a field has been set.
func (o *VirtualDatasetHooks) HasPreSelfRefresh() bool {
	if o != nil && !IsNil(o.PreSelfRefresh) {
		return true
	}

	return false
}

// SetPreSelfRefresh gets a reference to the given []Hook and assigns it to the PreSelfRefresh field.
func (o *VirtualDatasetHooks) SetPreSelfRefresh(v []Hook) {
	o.PreSelfRefresh = v
}

// GetPostSelfRefresh returns the PostSelfRefresh field value if set, zero value otherwise.
func (o *VirtualDatasetHooks) GetPostSelfRefresh() []Hook {
	if o == nil || IsNil(o.PostSelfRefresh) {
		var ret []Hook
		return ret
	}
	return o.PostSelfRefresh
}

// GetPostSelfRefreshOk returns a tuple with the PostSelfRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDatasetHooks) GetPostSelfRefreshOk() ([]Hook, bool) {
	if o == nil || IsNil(o.PostSelfRefresh) {
		return nil, false
	}
	return o.PostSelfRefresh, true
}

// HasPostSelfRefresh returns a boolean if a field has been set.
func (o *VirtualDatasetHooks) HasPostSelfRefresh() bool {
	if o != nil && !IsNil(o.PostSelfRefresh) {
		return true
	}

	return false
}

// SetPostSelfRefresh gets a reference to the given []Hook and assigns it to the PostSelfRefresh field.
func (o *VirtualDatasetHooks) SetPostSelfRefresh(v []Hook) {
	o.PostSelfRefresh = v
}

// GetPreRollback returns the PreRollback field value if set, zero value otherwise.
// Deprecated
func (o *VirtualDatasetHooks) GetPreRollback() []Hook {
	if o == nil || IsNil(o.PreRollback) {
		var ret []Hook
		return ret
	}
	return o.PreRollback
}

// GetPreRollbackOk returns a tuple with the PreRollback field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *VirtualDatasetHooks) GetPreRollbackOk() ([]Hook, bool) {
	if o == nil || IsNil(o.PreRollback) {
		return nil, false
	}
	return o.PreRollback, true
}

// HasPreRollback returns a boolean if a field has been set.
func (o *VirtualDatasetHooks) HasPreRollback() bool {
	if o != nil && !IsNil(o.PreRollback) {
		return true
	}

	return false
}

// SetPreRollback gets a reference to the given []Hook and assigns it to the PreRollback field.
// Deprecated
func (o *VirtualDatasetHooks) SetPreRollback(v []Hook) {
	o.PreRollback = v
}

// GetPostRollback returns the PostRollback field value if set, zero value otherwise.
// Deprecated
func (o *VirtualDatasetHooks) GetPostRollback() []Hook {
	if o == nil || IsNil(o.PostRollback) {
		var ret []Hook
		return ret
	}
	return o.PostRollback
}

// GetPostRollbackOk returns a tuple with the PostRollback field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *VirtualDatasetHooks) GetPostRollbackOk() ([]Hook, bool) {
	if o == nil || IsNil(o.PostRollback) {
		return nil, false
	}
	return o.PostRollback, true
}

// HasPostRollback returns a boolean if a field has been set.
func (o *VirtualDatasetHooks) HasPostRollback() bool {
	if o != nil && !IsNil(o.PostRollback) {
		return true
	}

	return false
}

// SetPostRollback gets a reference to the given []Hook and assigns it to the PostRollback field.
// Deprecated
func (o *VirtualDatasetHooks) SetPostRollback(v []Hook) {
	o.PostRollback = v
}

// GetConfigureClone returns the ConfigureClone field value if set, zero value otherwise.
func (o *VirtualDatasetHooks) GetConfigureClone() []Hook {
	if o == nil || IsNil(o.ConfigureClone) {
		var ret []Hook
		return ret
	}
	return o.ConfigureClone
}

// GetConfigureCloneOk returns a tuple with the ConfigureClone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDatasetHooks) GetConfigureCloneOk() ([]Hook, bool) {
	if o == nil || IsNil(o.ConfigureClone) {
		return nil, false
	}
	return o.ConfigureClone, true
}

// HasConfigureClone returns a boolean if a field has been set.
func (o *VirtualDatasetHooks) HasConfigureClone() bool {
	if o != nil && !IsNil(o.ConfigureClone) {
		return true
	}

	return false
}

// SetConfigureClone gets a reference to the given []Hook and assigns it to the ConfigureClone field.
func (o *VirtualDatasetHooks) SetConfigureClone(v []Hook) {
	o.ConfigureClone = v
}

// GetPreSnapshot returns the PreSnapshot field value if set, zero value otherwise.
func (o *VirtualDatasetHooks) GetPreSnapshot() []Hook {
	if o == nil || IsNil(o.PreSnapshot) {
		var ret []Hook
		return ret
	}
	return o.PreSnapshot
}

// GetPreSnapshotOk returns a tuple with the PreSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDatasetHooks) GetPreSnapshotOk() ([]Hook, bool) {
	if o == nil || IsNil(o.PreSnapshot) {
		return nil, false
	}
	return o.PreSnapshot, true
}

// HasPreSnapshot returns a boolean if a field has been set.
func (o *VirtualDatasetHooks) HasPreSnapshot() bool {
	if o != nil && !IsNil(o.PreSnapshot) {
		return true
	}

	return false
}

// SetPreSnapshot gets a reference to the given []Hook and assigns it to the PreSnapshot field.
func (o *VirtualDatasetHooks) SetPreSnapshot(v []Hook) {
	o.PreSnapshot = v
}

// GetPostSnapshot returns the PostSnapshot field value if set, zero value otherwise.
func (o *VirtualDatasetHooks) GetPostSnapshot() []Hook {
	if o == nil || IsNil(o.PostSnapshot) {
		var ret []Hook
		return ret
	}
	return o.PostSnapshot
}

// GetPostSnapshotOk returns a tuple with the PostSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDatasetHooks) GetPostSnapshotOk() ([]Hook, bool) {
	if o == nil || IsNil(o.PostSnapshot) {
		return nil, false
	}
	return o.PostSnapshot, true
}

// HasPostSnapshot returns a boolean if a field has been set.
func (o *VirtualDatasetHooks) HasPostSnapshot() bool {
	if o != nil && !IsNil(o.PostSnapshot) {
		return true
	}

	return false
}

// SetPostSnapshot gets a reference to the given []Hook and assigns it to the PostSnapshot field.
func (o *VirtualDatasetHooks) SetPostSnapshot(v []Hook) {
	o.PostSnapshot = v
}

// GetPreStart returns the PreStart field value if set, zero value otherwise.
func (o *VirtualDatasetHooks) GetPreStart() []Hook {
	if o == nil || IsNil(o.PreStart) {
		var ret []Hook
		return ret
	}
	return o.PreStart
}

// GetPreStartOk returns a tuple with the PreStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDatasetHooks) GetPreStartOk() ([]Hook, bool) {
	if o == nil || IsNil(o.PreStart) {
		return nil, false
	}
	return o.PreStart, true
}

// HasPreStart returns a boolean if a field has been set.
func (o *VirtualDatasetHooks) HasPreStart() bool {
	if o != nil && !IsNil(o.PreStart) {
		return true
	}

	return false
}

// SetPreStart gets a reference to the given []Hook and assigns it to the PreStart field.
func (o *VirtualDatasetHooks) SetPreStart(v []Hook) {
	o.PreStart = v
}

// GetPostStart returns the PostStart field value if set, zero value otherwise.
func (o *VirtualDatasetHooks) GetPostStart() []Hook {
	if o == nil || IsNil(o.PostStart) {
		var ret []Hook
		return ret
	}
	return o.PostStart
}

// GetPostStartOk returns a tuple with the PostStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDatasetHooks) GetPostStartOk() ([]Hook, bool) {
	if o == nil || IsNil(o.PostStart) {
		return nil, false
	}
	return o.PostStart, true
}

// HasPostStart returns a boolean if a field has been set.
func (o *VirtualDatasetHooks) HasPostStart() bool {
	if o != nil && !IsNil(o.PostStart) {
		return true
	}

	return false
}

// SetPostStart gets a reference to the given []Hook and assigns it to the PostStart field.
func (o *VirtualDatasetHooks) SetPostStart(v []Hook) {
	o.PostStart = v
}

// GetPreStop returns the PreStop field value if set, zero value otherwise.
func (o *VirtualDatasetHooks) GetPreStop() []Hook {
	if o == nil || IsNil(o.PreStop) {
		var ret []Hook
		return ret
	}
	return o.PreStop
}

// GetPreStopOk returns a tuple with the PreStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDatasetHooks) GetPreStopOk() ([]Hook, bool) {
	if o == nil || IsNil(o.PreStop) {
		return nil, false
	}
	return o.PreStop, true
}

// HasPreStop returns a boolean if a field has been set.
func (o *VirtualDatasetHooks) HasPreStop() bool {
	if o != nil && !IsNil(o.PreStop) {
		return true
	}

	return false
}

// SetPreStop gets a reference to the given []Hook and assigns it to the PreStop field.
func (o *VirtualDatasetHooks) SetPreStop(v []Hook) {
	o.PreStop = v
}

// GetPostStop returns the PostStop field value if set, zero value otherwise.
func (o *VirtualDatasetHooks) GetPostStop() []Hook {
	if o == nil || IsNil(o.PostStop) {
		var ret []Hook
		return ret
	}
	return o.PostStop
}

// GetPostStopOk returns a tuple with the PostStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDatasetHooks) GetPostStopOk() ([]Hook, bool) {
	if o == nil || IsNil(o.PostStop) {
		return nil, false
	}
	return o.PostStop, true
}

// HasPostStop returns a boolean if a field has been set.
func (o *VirtualDatasetHooks) HasPostStop() bool {
	if o != nil && !IsNil(o.PostStop) {
		return true
	}

	return false
}

// SetPostStop gets a reference to the given []Hook and assigns it to the PostStop field.
func (o *VirtualDatasetHooks) SetPostStop(v []Hook) {
	o.PostStop = v
}

func (o VirtualDatasetHooks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualDatasetHooks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreRefresh) {
		toSerialize["pre_refresh"] = o.PreRefresh
	}
	if !IsNil(o.PostRefresh) {
		toSerialize["post_refresh"] = o.PostRefresh
	}
	if !IsNil(o.PreSelfRefresh) {
		toSerialize["pre_self_refresh"] = o.PreSelfRefresh
	}
	if !IsNil(o.PostSelfRefresh) {
		toSerialize["post_self_refresh"] = o.PostSelfRefresh
	}
	if !IsNil(o.PreRollback) {
		toSerialize["pre_rollback"] = o.PreRollback
	}
	if !IsNil(o.PostRollback) {
		toSerialize["post_rollback"] = o.PostRollback
	}
	if !IsNil(o.ConfigureClone) {
		toSerialize["configure_clone"] = o.ConfigureClone
	}
	if !IsNil(o.PreSnapshot) {
		toSerialize["pre_snapshot"] = o.PreSnapshot
	}
	if !IsNil(o.PostSnapshot) {
		toSerialize["post_snapshot"] = o.PostSnapshot
	}
	if !IsNil(o.PreStart) {
		toSerialize["pre_start"] = o.PreStart
	}
	if !IsNil(o.PostStart) {
		toSerialize["post_start"] = o.PostStart
	}
	if !IsNil(o.PreStop) {
		toSerialize["pre_stop"] = o.PreStop
	}
	if !IsNil(o.PostStop) {
		toSerialize["post_stop"] = o.PostStop
	}
	return toSerialize, nil
}

type NullableVirtualDatasetHooks struct {
	value *VirtualDatasetHooks
	isSet bool
}

func (v NullableVirtualDatasetHooks) Get() *VirtualDatasetHooks {
	return v.value
}

func (v *NullableVirtualDatasetHooks) Set(val *VirtualDatasetHooks) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualDatasetHooks) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualDatasetHooks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualDatasetHooks(val *VirtualDatasetHooks) *NullableVirtualDatasetHooks {
	return &NullableVirtualDatasetHooks{value: val, isSet: true}
}

func (v NullableVirtualDatasetHooks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualDatasetHooks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


