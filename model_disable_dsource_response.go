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

// checks if the DisableDsourceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisableDsourceResponse{}

// DisableDsourceResponse struct for DisableDsourceResponse
type DisableDsourceResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewDisableDsourceResponse instantiates a new DisableDsourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisableDsourceResponse() *DisableDsourceResponse {
	this := DisableDsourceResponse{}
	return &this
}

// NewDisableDsourceResponseWithDefaults instantiates a new DisableDsourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisableDsourceResponseWithDefaults() *DisableDsourceResponse {
	this := DisableDsourceResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *DisableDsourceResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisableDsourceResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *DisableDsourceResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *DisableDsourceResponse) SetJob(v Job) {
	o.Job = &v
}

func (o DisableDsourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisableDsourceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableDisableDsourceResponse struct {
	value *DisableDsourceResponse
	isSet bool
}

func (v NullableDisableDsourceResponse) Get() *DisableDsourceResponse {
	return v.value
}

func (v *NullableDisableDsourceResponse) Set(val *DisableDsourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDisableDsourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDisableDsourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisableDsourceResponse(val *DisableDsourceResponse) *NullableDisableDsourceResponse {
	return &NullableDisableDsourceResponse{value: val, isSet: true}
}

func (v NullableDisableDsourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisableDsourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


