/*
Delphix DCT API

Delphix DCT API

API version: 2.0.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// EnableVDBResponse struct for EnableVDBResponse
type EnableVDBResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewEnableVDBResponse instantiates a new EnableVDBResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableVDBResponse() *EnableVDBResponse {
	this := EnableVDBResponse{}
	return &this
}

// NewEnableVDBResponseWithDefaults instantiates a new EnableVDBResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableVDBResponseWithDefaults() *EnableVDBResponse {
	this := EnableVDBResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *EnableVDBResponse) GetJob() Job {
	if o == nil || o.Job == nil {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableVDBResponse) GetJobOk() (*Job, bool) {
	if o == nil || o.Job == nil {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *EnableVDBResponse) HasJob() bool {
	if o != nil && o.Job != nil {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *EnableVDBResponse) SetJob(v Job) {
	o.Job = &v
}

func (o EnableVDBResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Job != nil {
		toSerialize["job"] = o.Job
	}
	return json.Marshal(toSerialize)
}

type NullableEnableVDBResponse struct {
	value *EnableVDBResponse
	isSet bool
}

func (v NullableEnableVDBResponse) Get() *EnableVDBResponse {
	return v.value
}

func (v *NullableEnableVDBResponse) Set(val *EnableVDBResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableVDBResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableVDBResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableVDBResponse(val *EnableVDBResponse) *NullableEnableVDBResponse {
	return &NullableEnableVDBResponse{value: val, isSet: true}
}

func (v NullableEnableVDBResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableVDBResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


