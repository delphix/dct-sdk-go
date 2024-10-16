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

// checks if the ProvisionVDBResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisionVDBResponse{}

// ProvisionVDBResponse struct for ProvisionVDBResponse
type ProvisionVDBResponse struct {
	Job *Job `json:"job,omitempty"`
	// The ID of the provisioned vdb.
	VdbId *string `json:"vdb_id,omitempty"`
}

// NewProvisionVDBResponse instantiates a new ProvisionVDBResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionVDBResponse() *ProvisionVDBResponse {
	this := ProvisionVDBResponse{}
	return &this
}

// NewProvisionVDBResponseWithDefaults instantiates a new ProvisionVDBResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionVDBResponseWithDefaults() *ProvisionVDBResponse {
	this := ProvisionVDBResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *ProvisionVDBResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionVDBResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *ProvisionVDBResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *ProvisionVDBResponse) SetJob(v Job) {
	o.Job = &v
}

// GetVdbId returns the VdbId field value if set, zero value otherwise.
func (o *ProvisionVDBResponse) GetVdbId() string {
	if o == nil || IsNil(o.VdbId) {
		var ret string
		return ret
	}
	return *o.VdbId
}

// GetVdbIdOk returns a tuple with the VdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionVDBResponse) GetVdbIdOk() (*string, bool) {
	if o == nil || IsNil(o.VdbId) {
		return nil, false
	}
	return o.VdbId, true
}

// HasVdbId returns a boolean if a field has been set.
func (o *ProvisionVDBResponse) HasVdbId() bool {
	if o != nil && !IsNil(o.VdbId) {
		return true
	}

	return false
}

// SetVdbId gets a reference to the given string and assigns it to the VdbId field.
func (o *ProvisionVDBResponse) SetVdbId(v string) {
	o.VdbId = &v
}

func (o ProvisionVDBResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisionVDBResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.VdbId) {
		toSerialize["vdb_id"] = o.VdbId
	}
	return toSerialize, nil
}

type NullableProvisionVDBResponse struct {
	value *ProvisionVDBResponse
	isSet bool
}

func (v NullableProvisionVDBResponse) Get() *ProvisionVDBResponse {
	return v.value
}

func (v *NullableProvisionVDBResponse) Set(val *ProvisionVDBResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionVDBResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionVDBResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionVDBResponse(val *ProvisionVDBResponse) *NullableProvisionVDBResponse {
	return &NullableProvisionVDBResponse{value: val, isSet: true}
}

func (v NullableProvisionVDBResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionVDBResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


