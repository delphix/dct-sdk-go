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

// checks if the RollbackVDBBySnapshotParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RollbackVDBBySnapshotParameters{}

// RollbackVDBBySnapshotParameters struct for RollbackVDBBySnapshotParameters
type RollbackVDBBySnapshotParameters struct {
	// The ID of the snapshot from which to execute the operation. If the snapshot_id is not, selects the latest snapshot.
	SnapshotId *string `json:"snapshot_id,omitempty"`
}

// NewRollbackVDBBySnapshotParameters instantiates a new RollbackVDBBySnapshotParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRollbackVDBBySnapshotParameters() *RollbackVDBBySnapshotParameters {
	this := RollbackVDBBySnapshotParameters{}
	return &this
}

// NewRollbackVDBBySnapshotParametersWithDefaults instantiates a new RollbackVDBBySnapshotParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollbackVDBBySnapshotParametersWithDefaults() *RollbackVDBBySnapshotParameters {
	this := RollbackVDBBySnapshotParameters{}
	return &this
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *RollbackVDBBySnapshotParameters) GetSnapshotId() string {
	if o == nil || IsNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollbackVDBBySnapshotParameters) GetSnapshotIdOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotId) {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *RollbackVDBBySnapshotParameters) HasSnapshotId() bool {
	if o != nil && !IsNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *RollbackVDBBySnapshotParameters) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

func (o RollbackVDBBySnapshotParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RollbackVDBBySnapshotParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SnapshotId) {
		toSerialize["snapshot_id"] = o.SnapshotId
	}
	return toSerialize, nil
}

type NullableRollbackVDBBySnapshotParameters struct {
	value *RollbackVDBBySnapshotParameters
	isSet bool
}

func (v NullableRollbackVDBBySnapshotParameters) Get() *RollbackVDBBySnapshotParameters {
	return v.value
}

func (v *NullableRollbackVDBBySnapshotParameters) Set(val *RollbackVDBBySnapshotParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableRollbackVDBBySnapshotParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableRollbackVDBBySnapshotParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRollbackVDBBySnapshotParameters(val *RollbackVDBBySnapshotParameters) *NullableRollbackVDBBySnapshotParameters {
	return &NullableRollbackVDBBySnapshotParameters{value: val, isSet: true}
}

func (v NullableRollbackVDBBySnapshotParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRollbackVDBBySnapshotParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


