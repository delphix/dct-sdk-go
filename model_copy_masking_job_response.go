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

// checks if the CopyMaskingJobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CopyMaskingJobResponse{}

// CopyMaskingJobResponse struct for CopyMaskingJobResponse
type CopyMaskingJobResponse struct {
	Job *Job `json:"job,omitempty"`
	MaskingJobId *string `json:"masking_job_id,omitempty"`
}

// NewCopyMaskingJobResponse instantiates a new CopyMaskingJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCopyMaskingJobResponse() *CopyMaskingJobResponse {
	this := CopyMaskingJobResponse{}
	return &this
}

// NewCopyMaskingJobResponseWithDefaults instantiates a new CopyMaskingJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCopyMaskingJobResponseWithDefaults() *CopyMaskingJobResponse {
	this := CopyMaskingJobResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *CopyMaskingJobResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMaskingJobResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *CopyMaskingJobResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *CopyMaskingJobResponse) SetJob(v Job) {
	o.Job = &v
}

// GetMaskingJobId returns the MaskingJobId field value if set, zero value otherwise.
func (o *CopyMaskingJobResponse) GetMaskingJobId() string {
	if o == nil || IsNil(o.MaskingJobId) {
		var ret string
		return ret
	}
	return *o.MaskingJobId
}

// GetMaskingJobIdOk returns a tuple with the MaskingJobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyMaskingJobResponse) GetMaskingJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.MaskingJobId) {
		return nil, false
	}
	return o.MaskingJobId, true
}

// HasMaskingJobId returns a boolean if a field has been set.
func (o *CopyMaskingJobResponse) HasMaskingJobId() bool {
	if o != nil && !IsNil(o.MaskingJobId) {
		return true
	}

	return false
}

// SetMaskingJobId gets a reference to the given string and assigns it to the MaskingJobId field.
func (o *CopyMaskingJobResponse) SetMaskingJobId(v string) {
	o.MaskingJobId = &v
}

func (o CopyMaskingJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CopyMaskingJobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.MaskingJobId) {
		toSerialize["masking_job_id"] = o.MaskingJobId
	}
	return toSerialize, nil
}

type NullableCopyMaskingJobResponse struct {
	value *CopyMaskingJobResponse
	isSet bool
}

func (v NullableCopyMaskingJobResponse) Get() *CopyMaskingJobResponse {
	return v.value
}

func (v *NullableCopyMaskingJobResponse) Set(val *CopyMaskingJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCopyMaskingJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCopyMaskingJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCopyMaskingJobResponse(val *CopyMaskingJobResponse) *NullableCopyMaskingJobResponse {
	return &NullableCopyMaskingJobResponse{value: val, isSet: true}
}

func (v NullableCopyMaskingJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCopyMaskingJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


