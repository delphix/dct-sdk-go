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

// checks if the CreateReplicationProfileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateReplicationProfileResponse{}

// CreateReplicationProfileResponse struct for CreateReplicationProfileResponse
type CreateReplicationProfileResponse struct {
	Job *Job `json:"job,omitempty"`
	// The id of the created ReplicationProfile.
	ReplicationProfileId *string `json:"replication_profile_id,omitempty"`
}

// NewCreateReplicationProfileResponse instantiates a new CreateReplicationProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReplicationProfileResponse() *CreateReplicationProfileResponse {
	this := CreateReplicationProfileResponse{}
	return &this
}

// NewCreateReplicationProfileResponseWithDefaults instantiates a new CreateReplicationProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReplicationProfileResponseWithDefaults() *CreateReplicationProfileResponse {
	this := CreateReplicationProfileResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *CreateReplicationProfileResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateReplicationProfileResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *CreateReplicationProfileResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *CreateReplicationProfileResponse) SetJob(v Job) {
	o.Job = &v
}

// GetReplicationProfileId returns the ReplicationProfileId field value if set, zero value otherwise.
func (o *CreateReplicationProfileResponse) GetReplicationProfileId() string {
	if o == nil || IsNil(o.ReplicationProfileId) {
		var ret string
		return ret
	}
	return *o.ReplicationProfileId
}

// GetReplicationProfileIdOk returns a tuple with the ReplicationProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateReplicationProfileResponse) GetReplicationProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicationProfileId) {
		return nil, false
	}
	return o.ReplicationProfileId, true
}

// HasReplicationProfileId returns a boolean if a field has been set.
func (o *CreateReplicationProfileResponse) HasReplicationProfileId() bool {
	if o != nil && !IsNil(o.ReplicationProfileId) {
		return true
	}

	return false
}

// SetReplicationProfileId gets a reference to the given string and assigns it to the ReplicationProfileId field.
func (o *CreateReplicationProfileResponse) SetReplicationProfileId(v string) {
	o.ReplicationProfileId = &v
}

func (o CreateReplicationProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateReplicationProfileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.ReplicationProfileId) {
		toSerialize["replication_profile_id"] = o.ReplicationProfileId
	}
	return toSerialize, nil
}

type NullableCreateReplicationProfileResponse struct {
	value *CreateReplicationProfileResponse
	isSet bool
}

func (v NullableCreateReplicationProfileResponse) Get() *CreateReplicationProfileResponse {
	return v.value
}

func (v *NullableCreateReplicationProfileResponse) Set(val *CreateReplicationProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReplicationProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReplicationProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReplicationProfileResponse(val *CreateReplicationProfileResponse) *NullableCreateReplicationProfileResponse {
	return &NullableCreateReplicationProfileResponse{value: val, isSet: true}
}

func (v NullableCreateReplicationProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReplicationProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


