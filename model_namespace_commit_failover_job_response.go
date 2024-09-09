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

// checks if the NamespaceCommitFailoverJobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamespaceCommitFailoverJobResponse{}

// NamespaceCommitFailoverJobResponse struct for NamespaceCommitFailoverJobResponse
type NamespaceCommitFailoverJobResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewNamespaceCommitFailoverJobResponse instantiates a new NamespaceCommitFailoverJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamespaceCommitFailoverJobResponse() *NamespaceCommitFailoverJobResponse {
	this := NamespaceCommitFailoverJobResponse{}
	return &this
}

// NewNamespaceCommitFailoverJobResponseWithDefaults instantiates a new NamespaceCommitFailoverJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamespaceCommitFailoverJobResponseWithDefaults() *NamespaceCommitFailoverJobResponse {
	this := NamespaceCommitFailoverJobResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *NamespaceCommitFailoverJobResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceCommitFailoverJobResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *NamespaceCommitFailoverJobResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *NamespaceCommitFailoverJobResponse) SetJob(v Job) {
	o.Job = &v
}

func (o NamespaceCommitFailoverJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamespaceCommitFailoverJobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableNamespaceCommitFailoverJobResponse struct {
	value *NamespaceCommitFailoverJobResponse
	isSet bool
}

func (v NullableNamespaceCommitFailoverJobResponse) Get() *NamespaceCommitFailoverJobResponse {
	return v.value
}

func (v *NullableNamespaceCommitFailoverJobResponse) Set(val *NamespaceCommitFailoverJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNamespaceCommitFailoverJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNamespaceCommitFailoverJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamespaceCommitFailoverJobResponse(val *NamespaceCommitFailoverJobResponse) *NullableNamespaceCommitFailoverJobResponse {
	return &NullableNamespaceCommitFailoverJobResponse{value: val, isSet: true}
}

func (v NullableNamespaceCommitFailoverJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamespaceCommitFailoverJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


