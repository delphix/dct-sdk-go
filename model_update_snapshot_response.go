/*
Delphix DCT API

Delphix DCT API

API version: 3.1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the UpdateSnapshotResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSnapshotResponse{}

// UpdateSnapshotResponse struct for UpdateSnapshotResponse
type UpdateSnapshotResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewUpdateSnapshotResponse instantiates a new UpdateSnapshotResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSnapshotResponse() *UpdateSnapshotResponse {
	this := UpdateSnapshotResponse{}
	return &this
}

// NewUpdateSnapshotResponseWithDefaults instantiates a new UpdateSnapshotResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSnapshotResponseWithDefaults() *UpdateSnapshotResponse {
	this := UpdateSnapshotResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *UpdateSnapshotResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSnapshotResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *UpdateSnapshotResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *UpdateSnapshotResponse) SetJob(v Job) {
	o.Job = &v
}

func (o UpdateSnapshotResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSnapshotResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableUpdateSnapshotResponse struct {
	value *UpdateSnapshotResponse
	isSet bool
}

func (v NullableUpdateSnapshotResponse) Get() *UpdateSnapshotResponse {
	return v.value
}

func (v *NullableUpdateSnapshotResponse) Set(val *UpdateSnapshotResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSnapshotResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSnapshotResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSnapshotResponse(val *UpdateSnapshotResponse) *NullableUpdateSnapshotResponse {
	return &NullableUpdateSnapshotResponse{value: val, isSet: true}
}

func (v NullableUpdateSnapshotResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSnapshotResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

