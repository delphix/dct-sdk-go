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

// checks if the SyncEnginesHyperscaleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncEnginesHyperscaleResponse{}

// SyncEnginesHyperscaleResponse struct for SyncEnginesHyperscaleResponse
type SyncEnginesHyperscaleResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewSyncEnginesHyperscaleResponse instantiates a new SyncEnginesHyperscaleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncEnginesHyperscaleResponse() *SyncEnginesHyperscaleResponse {
	this := SyncEnginesHyperscaleResponse{}
	return &this
}

// NewSyncEnginesHyperscaleResponseWithDefaults instantiates a new SyncEnginesHyperscaleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncEnginesHyperscaleResponseWithDefaults() *SyncEnginesHyperscaleResponse {
	this := SyncEnginesHyperscaleResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *SyncEnginesHyperscaleResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncEnginesHyperscaleResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *SyncEnginesHyperscaleResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *SyncEnginesHyperscaleResponse) SetJob(v Job) {
	o.Job = &v
}

func (o SyncEnginesHyperscaleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncEnginesHyperscaleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableSyncEnginesHyperscaleResponse struct {
	value *SyncEnginesHyperscaleResponse
	isSet bool
}

func (v NullableSyncEnginesHyperscaleResponse) Get() *SyncEnginesHyperscaleResponse {
	return v.value
}

func (v *NullableSyncEnginesHyperscaleResponse) Set(val *SyncEnginesHyperscaleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncEnginesHyperscaleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncEnginesHyperscaleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncEnginesHyperscaleResponse(val *SyncEnginesHyperscaleResponse) *NullableSyncEnginesHyperscaleResponse {
	return &NullableSyncEnginesHyperscaleResponse{value: val, isSet: true}
}

func (v NullableSyncEnginesHyperscaleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncEnginesHyperscaleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


