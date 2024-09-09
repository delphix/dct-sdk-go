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

// checks if the UpdateConnectorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateConnectorResponse{}

// UpdateConnectorResponse struct for UpdateConnectorResponse
type UpdateConnectorResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewUpdateConnectorResponse instantiates a new UpdateConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateConnectorResponse() *UpdateConnectorResponse {
	this := UpdateConnectorResponse{}
	return &this
}

// NewUpdateConnectorResponseWithDefaults instantiates a new UpdateConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateConnectorResponseWithDefaults() *UpdateConnectorResponse {
	this := UpdateConnectorResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *UpdateConnectorResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectorResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *UpdateConnectorResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *UpdateConnectorResponse) SetJob(v Job) {
	o.Job = &v
}

func (o UpdateConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateConnectorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableUpdateConnectorResponse struct {
	value *UpdateConnectorResponse
	isSet bool
}

func (v NullableUpdateConnectorResponse) Get() *UpdateConnectorResponse {
	return v.value
}

func (v *NullableUpdateConnectorResponse) Set(val *UpdateConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateConnectorResponse(val *UpdateConnectorResponse) *NullableUpdateConnectorResponse {
	return &NullableUpdateConnectorResponse{value: val, isSet: true}
}

func (v NullableUpdateConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


