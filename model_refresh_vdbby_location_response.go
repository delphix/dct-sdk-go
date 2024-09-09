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

// checks if the RefreshVDBByLocationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshVDBByLocationResponse{}

// RefreshVDBByLocationResponse struct for RefreshVDBByLocationResponse
type RefreshVDBByLocationResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewRefreshVDBByLocationResponse instantiates a new RefreshVDBByLocationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshVDBByLocationResponse() *RefreshVDBByLocationResponse {
	this := RefreshVDBByLocationResponse{}
	return &this
}

// NewRefreshVDBByLocationResponseWithDefaults instantiates a new RefreshVDBByLocationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshVDBByLocationResponseWithDefaults() *RefreshVDBByLocationResponse {
	this := RefreshVDBByLocationResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *RefreshVDBByLocationResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshVDBByLocationResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *RefreshVDBByLocationResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *RefreshVDBByLocationResponse) SetJob(v Job) {
	o.Job = &v
}

func (o RefreshVDBByLocationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshVDBByLocationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableRefreshVDBByLocationResponse struct {
	value *RefreshVDBByLocationResponse
	isSet bool
}

func (v NullableRefreshVDBByLocationResponse) Get() *RefreshVDBByLocationResponse {
	return v.value
}

func (v *NullableRefreshVDBByLocationResponse) Set(val *RefreshVDBByLocationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshVDBByLocationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshVDBByLocationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshVDBByLocationResponse(val *RefreshVDBByLocationResponse) *NullableRefreshVDBByLocationResponse {
	return &NullableRefreshVDBByLocationResponse{value: val, isSet: true}
}

func (v NullableRefreshVDBByLocationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshVDBByLocationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


