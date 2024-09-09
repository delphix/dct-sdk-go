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

// checks if the DeleteTimeflowResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteTimeflowResponse{}

// DeleteTimeflowResponse struct for DeleteTimeflowResponse
type DeleteTimeflowResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewDeleteTimeflowResponse instantiates a new DeleteTimeflowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTimeflowResponse() *DeleteTimeflowResponse {
	this := DeleteTimeflowResponse{}
	return &this
}

// NewDeleteTimeflowResponseWithDefaults instantiates a new DeleteTimeflowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTimeflowResponseWithDefaults() *DeleteTimeflowResponse {
	this := DeleteTimeflowResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *DeleteTimeflowResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTimeflowResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *DeleteTimeflowResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *DeleteTimeflowResponse) SetJob(v Job) {
	o.Job = &v
}

func (o DeleteTimeflowResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteTimeflowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableDeleteTimeflowResponse struct {
	value *DeleteTimeflowResponse
	isSet bool
}

func (v NullableDeleteTimeflowResponse) Get() *DeleteTimeflowResponse {
	return v.value
}

func (v *NullableDeleteTimeflowResponse) Set(val *DeleteTimeflowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTimeflowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTimeflowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTimeflowResponse(val *DeleteTimeflowResponse) *NullableDeleteTimeflowResponse {
	return &NullableDeleteTimeflowResponse{value: val, isSet: true}
}

func (v NullableDeleteTimeflowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTimeflowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


