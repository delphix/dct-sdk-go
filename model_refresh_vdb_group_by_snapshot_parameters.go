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

// checks if the RefreshVDBGroupBySnapshotParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshVDBGroupBySnapshotParameters{}

// RefreshVDBGroupBySnapshotParameters Parameters to refresh a VDB Group by snapshot.
type RefreshVDBGroupBySnapshotParameters struct {
	// List of the pair of VDB and snapshot to refresh from. If this is not set, all VDBs will be refreshed from latest snapshot of their parent.
	VdbSnapshotMappings []VDBGroupRefreshBySnapshot `json:"vdb_snapshot_mappings,omitempty"`
}

// NewRefreshVDBGroupBySnapshotParameters instantiates a new RefreshVDBGroupBySnapshotParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshVDBGroupBySnapshotParameters() *RefreshVDBGroupBySnapshotParameters {
	this := RefreshVDBGroupBySnapshotParameters{}
	return &this
}

// NewRefreshVDBGroupBySnapshotParametersWithDefaults instantiates a new RefreshVDBGroupBySnapshotParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshVDBGroupBySnapshotParametersWithDefaults() *RefreshVDBGroupBySnapshotParameters {
	this := RefreshVDBGroupBySnapshotParameters{}
	return &this
}

// GetVdbSnapshotMappings returns the VdbSnapshotMappings field value if set, zero value otherwise.
func (o *RefreshVDBGroupBySnapshotParameters) GetVdbSnapshotMappings() []VDBGroupRefreshBySnapshot {
	if o == nil || IsNil(o.VdbSnapshotMappings) {
		var ret []VDBGroupRefreshBySnapshot
		return ret
	}
	return o.VdbSnapshotMappings
}

// GetVdbSnapshotMappingsOk returns a tuple with the VdbSnapshotMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshVDBGroupBySnapshotParameters) GetVdbSnapshotMappingsOk() ([]VDBGroupRefreshBySnapshot, bool) {
	if o == nil || IsNil(o.VdbSnapshotMappings) {
		return nil, false
	}
	return o.VdbSnapshotMappings, true
}

// HasVdbSnapshotMappings returns a boolean if a field has been set.
func (o *RefreshVDBGroupBySnapshotParameters) HasVdbSnapshotMappings() bool {
	if o != nil && !IsNil(o.VdbSnapshotMappings) {
		return true
	}

	return false
}

// SetVdbSnapshotMappings gets a reference to the given []VDBGroupRefreshBySnapshot and assigns it to the VdbSnapshotMappings field.
func (o *RefreshVDBGroupBySnapshotParameters) SetVdbSnapshotMappings(v []VDBGroupRefreshBySnapshot) {
	o.VdbSnapshotMappings = v
}

func (o RefreshVDBGroupBySnapshotParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshVDBGroupBySnapshotParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VdbSnapshotMappings) {
		toSerialize["vdb_snapshot_mappings"] = o.VdbSnapshotMappings
	}
	return toSerialize, nil
}

type NullableRefreshVDBGroupBySnapshotParameters struct {
	value *RefreshVDBGroupBySnapshotParameters
	isSet bool
}

func (v NullableRefreshVDBGroupBySnapshotParameters) Get() *RefreshVDBGroupBySnapshotParameters {
	return v.value
}

func (v *NullableRefreshVDBGroupBySnapshotParameters) Set(val *RefreshVDBGroupBySnapshotParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshVDBGroupBySnapshotParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshVDBGroupBySnapshotParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshVDBGroupBySnapshotParameters(val *RefreshVDBGroupBySnapshotParameters) *NullableRefreshVDBGroupBySnapshotParameters {
	return &NullableRefreshVDBGroupBySnapshotParameters{value: val, isSet: true}
}

func (v NullableRefreshVDBGroupBySnapshotParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshVDBGroupBySnapshotParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


