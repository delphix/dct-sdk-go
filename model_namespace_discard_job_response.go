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

// checks if the NamespaceDiscardJobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamespaceDiscardJobResponse{}

// NamespaceDiscardJobResponse struct for NamespaceDiscardJobResponse
type NamespaceDiscardJobResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewNamespaceDiscardJobResponse instantiates a new NamespaceDiscardJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamespaceDiscardJobResponse() *NamespaceDiscardJobResponse {
	this := NamespaceDiscardJobResponse{}
	return &this
}

// NewNamespaceDiscardJobResponseWithDefaults instantiates a new NamespaceDiscardJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamespaceDiscardJobResponseWithDefaults() *NamespaceDiscardJobResponse {
	this := NamespaceDiscardJobResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *NamespaceDiscardJobResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceDiscardJobResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *NamespaceDiscardJobResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *NamespaceDiscardJobResponse) SetJob(v Job) {
	o.Job = &v
}

func (o NamespaceDiscardJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamespaceDiscardJobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableNamespaceDiscardJobResponse struct {
	value *NamespaceDiscardJobResponse
	isSet bool
}

func (v NullableNamespaceDiscardJobResponse) Get() *NamespaceDiscardJobResponse {
	return v.value
}

func (v *NullableNamespaceDiscardJobResponse) Set(val *NamespaceDiscardJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNamespaceDiscardJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNamespaceDiscardJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamespaceDiscardJobResponse(val *NamespaceDiscardJobResponse) *NullableNamespaceDiscardJobResponse {
	return &NullableNamespaceDiscardJobResponse{value: val, isSet: true}
}

func (v NullableNamespaceDiscardJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamespaceDiscardJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


