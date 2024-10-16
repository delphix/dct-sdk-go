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

// checks if the StartVDBResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartVDBResponse{}

// StartVDBResponse struct for StartVDBResponse
type StartVDBResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewStartVDBResponse instantiates a new StartVDBResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartVDBResponse() *StartVDBResponse {
	this := StartVDBResponse{}
	return &this
}

// NewStartVDBResponseWithDefaults instantiates a new StartVDBResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartVDBResponseWithDefaults() *StartVDBResponse {
	this := StartVDBResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *StartVDBResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartVDBResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *StartVDBResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *StartVDBResponse) SetJob(v Job) {
	o.Job = &v
}

func (o StartVDBResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartVDBResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableStartVDBResponse struct {
	value *StartVDBResponse
	isSet bool
}

func (v NullableStartVDBResponse) Get() *StartVDBResponse {
	return v.value
}

func (v *NullableStartVDBResponse) Set(val *StartVDBResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStartVDBResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStartVDBResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartVDBResponse(val *StartVDBResponse) *NullableStartVDBResponse {
	return &NullableStartVDBResponse{value: val, isSet: true}
}

func (v NullableStartVDBResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartVDBResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


