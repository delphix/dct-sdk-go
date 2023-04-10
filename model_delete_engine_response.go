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

// checks if the DeleteEngineResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteEngineResponse{}

// DeleteEngineResponse struct for DeleteEngineResponse
type DeleteEngineResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewDeleteEngineResponse instantiates a new DeleteEngineResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteEngineResponse() *DeleteEngineResponse {
	this := DeleteEngineResponse{}
	return &this
}

// NewDeleteEngineResponseWithDefaults instantiates a new DeleteEngineResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteEngineResponseWithDefaults() *DeleteEngineResponse {
	this := DeleteEngineResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *DeleteEngineResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteEngineResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *DeleteEngineResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *DeleteEngineResponse) SetJob(v Job) {
	o.Job = &v
}

func (o DeleteEngineResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteEngineResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableDeleteEngineResponse struct {
	value *DeleteEngineResponse
	isSet bool
}

func (v NullableDeleteEngineResponse) Get() *DeleteEngineResponse {
	return v.value
}

func (v *NullableDeleteEngineResponse) Set(val *DeleteEngineResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteEngineResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteEngineResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteEngineResponse(val *DeleteEngineResponse) *NullableDeleteEngineResponse {
	return &NullableDeleteEngineResponse{value: val, isSet: true}
}

func (v NullableDeleteEngineResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteEngineResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


