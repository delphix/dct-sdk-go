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

// checks if the UpdateEngineAutoTaggingConfigResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEngineAutoTaggingConfigResponse{}

// UpdateEngineAutoTaggingConfigResponse struct for UpdateEngineAutoTaggingConfigResponse
type UpdateEngineAutoTaggingConfigResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewUpdateEngineAutoTaggingConfigResponse instantiates a new UpdateEngineAutoTaggingConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEngineAutoTaggingConfigResponse() *UpdateEngineAutoTaggingConfigResponse {
	this := UpdateEngineAutoTaggingConfigResponse{}
	return &this
}

// NewUpdateEngineAutoTaggingConfigResponseWithDefaults instantiates a new UpdateEngineAutoTaggingConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEngineAutoTaggingConfigResponseWithDefaults() *UpdateEngineAutoTaggingConfigResponse {
	this := UpdateEngineAutoTaggingConfigResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *UpdateEngineAutoTaggingConfigResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEngineAutoTaggingConfigResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *UpdateEngineAutoTaggingConfigResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *UpdateEngineAutoTaggingConfigResponse) SetJob(v Job) {
	o.Job = &v
}

func (o UpdateEngineAutoTaggingConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEngineAutoTaggingConfigResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableUpdateEngineAutoTaggingConfigResponse struct {
	value *UpdateEngineAutoTaggingConfigResponse
	isSet bool
}

func (v NullableUpdateEngineAutoTaggingConfigResponse) Get() *UpdateEngineAutoTaggingConfigResponse {
	return v.value
}

func (v *NullableUpdateEngineAutoTaggingConfigResponse) Set(val *UpdateEngineAutoTaggingConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEngineAutoTaggingConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEngineAutoTaggingConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEngineAutoTaggingConfigResponse(val *UpdateEngineAutoTaggingConfigResponse) *NullableUpdateEngineAutoTaggingConfigResponse {
	return &NullableUpdateEngineAutoTaggingConfigResponse{value: val, isSet: true}
}

func (v NullableUpdateEngineAutoTaggingConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEngineAutoTaggingConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


