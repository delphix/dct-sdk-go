/*
Delphix DCT API

Delphix DCT API

API version: 3.9.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the CreateMaskingJobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMaskingJobResponse{}

// CreateMaskingJobResponse struct for CreateMaskingJobResponse
type CreateMaskingJobResponse struct {
	// The ID of the created masking job.
	Id *string `json:"id,omitempty"`
	Job *Job `json:"job,omitempty"`
}

// NewCreateMaskingJobResponse instantiates a new CreateMaskingJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMaskingJobResponse() *CreateMaskingJobResponse {
	this := CreateMaskingJobResponse{}
	return &this
}

// NewCreateMaskingJobResponseWithDefaults instantiates a new CreateMaskingJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMaskingJobResponseWithDefaults() *CreateMaskingJobResponse {
	this := CreateMaskingJobResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateMaskingJobResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMaskingJobResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateMaskingJobResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateMaskingJobResponse) SetId(v string) {
	o.Id = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *CreateMaskingJobResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMaskingJobResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *CreateMaskingJobResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *CreateMaskingJobResponse) SetJob(v Job) {
	o.Job = &v
}

func (o CreateMaskingJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMaskingJobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableCreateMaskingJobResponse struct {
	value *CreateMaskingJobResponse
	isSet bool
}

func (v NullableCreateMaskingJobResponse) Get() *CreateMaskingJobResponse {
	return v.value
}

func (v *NullableCreateMaskingJobResponse) Set(val *CreateMaskingJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMaskingJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMaskingJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMaskingJobResponse(val *CreateMaskingJobResponse) *NullableCreateMaskingJobResponse {
	return &NullableCreateMaskingJobResponse{value: val, isSet: true}
}

func (v NullableCreateMaskingJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMaskingJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


