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

// checks if the StartVCDBResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartVCDBResponse{}

// StartVCDBResponse struct for StartVCDBResponse
type StartVCDBResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewStartVCDBResponse instantiates a new StartVCDBResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartVCDBResponse() *StartVCDBResponse {
	this := StartVCDBResponse{}
	return &this
}

// NewStartVCDBResponseWithDefaults instantiates a new StartVCDBResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartVCDBResponseWithDefaults() *StartVCDBResponse {
	this := StartVCDBResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *StartVCDBResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartVCDBResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *StartVCDBResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *StartVCDBResponse) SetJob(v Job) {
	o.Job = &v
}

func (o StartVCDBResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartVCDBResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableStartVCDBResponse struct {
	value *StartVCDBResponse
	isSet bool
}

func (v NullableStartVCDBResponse) Get() *StartVCDBResponse {
	return v.value
}

func (v *NullableStartVCDBResponse) Set(val *StartVCDBResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStartVCDBResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStartVCDBResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartVCDBResponse(val *StartVCDBResponse) *NullableStartVCDBResponse {
	return &NullableStartVCDBResponse{value: val, isSet: true}
}

func (v NullableStartVCDBResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartVCDBResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


