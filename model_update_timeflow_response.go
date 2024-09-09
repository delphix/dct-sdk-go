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

// checks if the UpdateTimeflowResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTimeflowResponse{}

// UpdateTimeflowResponse struct for UpdateTimeflowResponse
type UpdateTimeflowResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewUpdateTimeflowResponse instantiates a new UpdateTimeflowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTimeflowResponse() *UpdateTimeflowResponse {
	this := UpdateTimeflowResponse{}
	return &this
}

// NewUpdateTimeflowResponseWithDefaults instantiates a new UpdateTimeflowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTimeflowResponseWithDefaults() *UpdateTimeflowResponse {
	this := UpdateTimeflowResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *UpdateTimeflowResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTimeflowResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *UpdateTimeflowResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *UpdateTimeflowResponse) SetJob(v Job) {
	o.Job = &v
}

func (o UpdateTimeflowResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTimeflowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableUpdateTimeflowResponse struct {
	value *UpdateTimeflowResponse
	isSet bool
}

func (v NullableUpdateTimeflowResponse) Get() *UpdateTimeflowResponse {
	return v.value
}

func (v *NullableUpdateTimeflowResponse) Set(val *UpdateTimeflowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTimeflowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTimeflowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTimeflowResponse(val *UpdateTimeflowResponse) *NullableUpdateTimeflowResponse {
	return &NullableUpdateTimeflowResponse{value: val, isSet: true}
}

func (v NullableUpdateTimeflowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTimeflowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


