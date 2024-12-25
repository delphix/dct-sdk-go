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
	"time"
)

// checks if the VDBOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VDBOrder{}

// VDBOrder struct for VDBOrder
type VDBOrder struct {
	// Vdb id
	VdbId *string `json:"vdb_id,omitempty"`
	// Dictates order of operations on VDBs. Operations can be performed in parallel <br> for all VDBs or sequentially. Below are possible valid and invalid orderings given an example <br> VDB group with 3 vdbs (A, B, and C).<br> Valid:<br> {\"vdb_id\":\"vdb-1\", \"order\":\"1\"} {\"vdb_id\":\"vdb-2\", order:\"1\"} {vdb_id:\"vdb-3\", order:\"1\"} (parallel)<br> {vdb_id:\"vdb-1\", order:\"1\"} {vdb_id:\"vdb-2\", order:\"2\"} {vdb_id:\"vdb-3\", order:\"3\"} (sequential)<br> Invalid:<br> {vdb_id:\"vdb-1\", order:\"A\"} {vdb_id:\"vdb-2\", order:\"B\"} {vdb_id:\"vdb-3\", order:\"C\"} (sequential)<br><br> In the sequential case the vdbs with priority 1 is the first to be started and the last to<br> be stopped. This value is set on creation of VDB groups.
	Order *int32 `json:"order,omitempty"`
	VdbName *string `json:"vdb_name,omitempty"`
	// The last time the VDB was successfully refreshed as a part of VDB Group refresh operation in UTC timezone.
	LastRefreshTimeWithGroupRefresh *time.Time `json:"last_refresh_time_with_group_refresh,omitempty"`
	// Indicates if the VDB is in sync with the VDB Group or not. If this VDB is was last refreshed as part of the VDB Group then this value will be true.
	InSync *bool `json:"in_sync,omitempty"`
}

// NewVDBOrder instantiates a new VDBOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVDBOrder() *VDBOrder {
	this := VDBOrder{}
	var order int32 = 0
	this.Order = &order
	return &this
}

// NewVDBOrderWithDefaults instantiates a new VDBOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVDBOrderWithDefaults() *VDBOrder {
	this := VDBOrder{}
	var order int32 = 0
	this.Order = &order
	return &this
}

// GetVdbId returns the VdbId field value if set, zero value otherwise.
func (o *VDBOrder) GetVdbId() string {
	if o == nil || IsNil(o.VdbId) {
		var ret string
		return ret
	}
	return *o.VdbId
}

// GetVdbIdOk returns a tuple with the VdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBOrder) GetVdbIdOk() (*string, bool) {
	if o == nil || IsNil(o.VdbId) {
		return nil, false
	}
	return o.VdbId, true
}

// HasVdbId returns a boolean if a field has been set.
func (o *VDBOrder) HasVdbId() bool {
	if o != nil && !IsNil(o.VdbId) {
		return true
	}

	return false
}

// SetVdbId gets a reference to the given string and assigns it to the VdbId field.
func (o *VDBOrder) SetVdbId(v string) {
	o.VdbId = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *VDBOrder) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBOrder) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *VDBOrder) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *VDBOrder) SetOrder(v int32) {
	o.Order = &v
}

// GetVdbName returns the VdbName field value if set, zero value otherwise.
func (o *VDBOrder) GetVdbName() string {
	if o == nil || IsNil(o.VdbName) {
		var ret string
		return ret
	}
	return *o.VdbName
}

// GetVdbNameOk returns a tuple with the VdbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBOrder) GetVdbNameOk() (*string, bool) {
	if o == nil || IsNil(o.VdbName) {
		return nil, false
	}
	return o.VdbName, true
}

// HasVdbName returns a boolean if a field has been set.
func (o *VDBOrder) HasVdbName() bool {
	if o != nil && !IsNil(o.VdbName) {
		return true
	}

	return false
}

// SetVdbName gets a reference to the given string and assigns it to the VdbName field.
func (o *VDBOrder) SetVdbName(v string) {
	o.VdbName = &v
}

// GetLastRefreshTimeWithGroupRefresh returns the LastRefreshTimeWithGroupRefresh field value if set, zero value otherwise.
func (o *VDBOrder) GetLastRefreshTimeWithGroupRefresh() time.Time {
	if o == nil || IsNil(o.LastRefreshTimeWithGroupRefresh) {
		var ret time.Time
		return ret
	}
	return *o.LastRefreshTimeWithGroupRefresh
}

// GetLastRefreshTimeWithGroupRefreshOk returns a tuple with the LastRefreshTimeWithGroupRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBOrder) GetLastRefreshTimeWithGroupRefreshOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastRefreshTimeWithGroupRefresh) {
		return nil, false
	}
	return o.LastRefreshTimeWithGroupRefresh, true
}

// HasLastRefreshTimeWithGroupRefresh returns a boolean if a field has been set.
func (o *VDBOrder) HasLastRefreshTimeWithGroupRefresh() bool {
	if o != nil && !IsNil(o.LastRefreshTimeWithGroupRefresh) {
		return true
	}

	return false
}

// SetLastRefreshTimeWithGroupRefresh gets a reference to the given time.Time and assigns it to the LastRefreshTimeWithGroupRefresh field.
func (o *VDBOrder) SetLastRefreshTimeWithGroupRefresh(v time.Time) {
	o.LastRefreshTimeWithGroupRefresh = &v
}

// GetInSync returns the InSync field value if set, zero value otherwise.
func (o *VDBOrder) GetInSync() bool {
	if o == nil || IsNil(o.InSync) {
		var ret bool
		return ret
	}
	return *o.InSync
}

// GetInSyncOk returns a tuple with the InSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBOrder) GetInSyncOk() (*bool, bool) {
	if o == nil || IsNil(o.InSync) {
		return nil, false
	}
	return o.InSync, true
}

// HasInSync returns a boolean if a field has been set.
func (o *VDBOrder) HasInSync() bool {
	if o != nil && !IsNil(o.InSync) {
		return true
	}

	return false
}

// SetInSync gets a reference to the given bool and assigns it to the InSync field.
func (o *VDBOrder) SetInSync(v bool) {
	o.InSync = &v
}

func (o VDBOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VDBOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VdbId) {
		toSerialize["vdb_id"] = o.VdbId
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.VdbName) {
		toSerialize["vdb_name"] = o.VdbName
	}
	if !IsNil(o.LastRefreshTimeWithGroupRefresh) {
		toSerialize["last_refresh_time_with_group_refresh"] = o.LastRefreshTimeWithGroupRefresh
	}
	if !IsNil(o.InSync) {
		toSerialize["in_sync"] = o.InSync
	}
	return toSerialize, nil
}

type NullableVDBOrder struct {
	value *VDBOrder
	isSet bool
}

func (v NullableVDBOrder) Get() *VDBOrder {
	return v.value
}

func (v *NullableVDBOrder) Set(val *VDBOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableVDBOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableVDBOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVDBOrder(val *VDBOrder) *NullableVDBOrder {
	return &NullableVDBOrder{value: val, isSet: true}
}

func (v NullableVDBOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVDBOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


