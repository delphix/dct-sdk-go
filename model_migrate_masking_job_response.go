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
)

// checks if the MigrateMaskingJobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MigrateMaskingJobResponse{}

// MigrateMaskingJobResponse struct for MigrateMaskingJobResponse
type MigrateMaskingJobResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewMigrateMaskingJobResponse instantiates a new MigrateMaskingJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrateMaskingJobResponse() *MigrateMaskingJobResponse {
	this := MigrateMaskingJobResponse{}
	return &this
}

// NewMigrateMaskingJobResponseWithDefaults instantiates a new MigrateMaskingJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrateMaskingJobResponseWithDefaults() *MigrateMaskingJobResponse {
	this := MigrateMaskingJobResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *MigrateMaskingJobResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateMaskingJobResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *MigrateMaskingJobResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *MigrateMaskingJobResponse) SetJob(v Job) {
	o.Job = &v
}

func (o MigrateMaskingJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MigrateMaskingJobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableMigrateMaskingJobResponse struct {
	value *MigrateMaskingJobResponse
	isSet bool
}

func (v NullableMigrateMaskingJobResponse) Get() *MigrateMaskingJobResponse {
	return v.value
}

func (v *NullableMigrateMaskingJobResponse) Set(val *MigrateMaskingJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrateMaskingJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrateMaskingJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrateMaskingJobResponse(val *MigrateMaskingJobResponse) *NullableMigrateMaskingJobResponse {
	return &NullableMigrateMaskingJobResponse{value: val, isSet: true}
}

func (v NullableMigrateMaskingJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrateMaskingJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


