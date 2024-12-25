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

// checks if the RollbackVDBBySnapshotResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RollbackVDBBySnapshotResponse{}

// RollbackVDBBySnapshotResponse struct for RollbackVDBBySnapshotResponse
type RollbackVDBBySnapshotResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewRollbackVDBBySnapshotResponse instantiates a new RollbackVDBBySnapshotResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRollbackVDBBySnapshotResponse() *RollbackVDBBySnapshotResponse {
	this := RollbackVDBBySnapshotResponse{}
	return &this
}

// NewRollbackVDBBySnapshotResponseWithDefaults instantiates a new RollbackVDBBySnapshotResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollbackVDBBySnapshotResponseWithDefaults() *RollbackVDBBySnapshotResponse {
	this := RollbackVDBBySnapshotResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *RollbackVDBBySnapshotResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollbackVDBBySnapshotResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *RollbackVDBBySnapshotResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *RollbackVDBBySnapshotResponse) SetJob(v Job) {
	o.Job = &v
}

func (o RollbackVDBBySnapshotResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RollbackVDBBySnapshotResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableRollbackVDBBySnapshotResponse struct {
	value *RollbackVDBBySnapshotResponse
	isSet bool
}

func (v NullableRollbackVDBBySnapshotResponse) Get() *RollbackVDBBySnapshotResponse {
	return v.value
}

func (v *NullableRollbackVDBBySnapshotResponse) Set(val *RollbackVDBBySnapshotResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRollbackVDBBySnapshotResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRollbackVDBBySnapshotResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRollbackVDBBySnapshotResponse(val *RollbackVDBBySnapshotResponse) *NullableRollbackVDBBySnapshotResponse {
	return &NullableRollbackVDBBySnapshotResponse{value: val, isSet: true}
}

func (v NullableRollbackVDBBySnapshotResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRollbackVDBBySnapshotResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


