/*
Delphix DCT API

Delphix DCT API

API version: 3.18.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the RefreshEnvironmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshEnvironmentResponse{}

// RefreshEnvironmentResponse struct for RefreshEnvironmentResponse
type RefreshEnvironmentResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewRefreshEnvironmentResponse instantiates a new RefreshEnvironmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshEnvironmentResponse() *RefreshEnvironmentResponse {
	this := RefreshEnvironmentResponse{}
	return &this
}

// NewRefreshEnvironmentResponseWithDefaults instantiates a new RefreshEnvironmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshEnvironmentResponseWithDefaults() *RefreshEnvironmentResponse {
	this := RefreshEnvironmentResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *RefreshEnvironmentResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshEnvironmentResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *RefreshEnvironmentResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *RefreshEnvironmentResponse) SetJob(v Job) {
	o.Job = &v
}

func (o RefreshEnvironmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshEnvironmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableRefreshEnvironmentResponse struct {
	value *RefreshEnvironmentResponse
	isSet bool
}

func (v NullableRefreshEnvironmentResponse) Get() *RefreshEnvironmentResponse {
	return v.value
}

func (v *NullableRefreshEnvironmentResponse) Set(val *RefreshEnvironmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshEnvironmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshEnvironmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshEnvironmentResponse(val *RefreshEnvironmentResponse) *NullableRefreshEnvironmentResponse {
	return &NullableRefreshEnvironmentResponse{value: val, isSet: true}
}

func (v NullableRefreshEnvironmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshEnvironmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


