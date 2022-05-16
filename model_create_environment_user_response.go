/*
Delphix DCT API

Delphix DCT API

API version: 1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// CreateEnvironmentUserResponse struct for CreateEnvironmentUserResponse
type CreateEnvironmentUserResponse struct {
	// The reference of the created environment user
	UserRef *string `json:"user_ref,omitempty"`
	Job *Job `json:"job,omitempty"`
}

// NewCreateEnvironmentUserResponse instantiates a new CreateEnvironmentUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEnvironmentUserResponse() *CreateEnvironmentUserResponse {
	this := CreateEnvironmentUserResponse{}
	return &this
}

// NewCreateEnvironmentUserResponseWithDefaults instantiates a new CreateEnvironmentUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEnvironmentUserResponseWithDefaults() *CreateEnvironmentUserResponse {
	this := CreateEnvironmentUserResponse{}
	return &this
}

// GetUserRef returns the UserRef field value if set, zero value otherwise.
func (o *CreateEnvironmentUserResponse) GetUserRef() string {
	if o == nil || o.UserRef == nil {
		var ret string
		return ret
	}
	return *o.UserRef
}

// GetUserRefOk returns a tuple with the UserRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentUserResponse) GetUserRefOk() (*string, bool) {
	if o == nil || o.UserRef == nil {
		return nil, false
	}
	return o.UserRef, true
}

// HasUserRef returns a boolean if a field has been set.
func (o *CreateEnvironmentUserResponse) HasUserRef() bool {
	if o != nil && o.UserRef != nil {
		return true
	}

	return false
}

// SetUserRef gets a reference to the given string and assigns it to the UserRef field.
func (o *CreateEnvironmentUserResponse) SetUserRef(v string) {
	o.UserRef = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *CreateEnvironmentUserResponse) GetJob() Job {
	if o == nil || o.Job == nil {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentUserResponse) GetJobOk() (*Job, bool) {
	if o == nil || o.Job == nil {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *CreateEnvironmentUserResponse) HasJob() bool {
	if o != nil && o.Job != nil {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *CreateEnvironmentUserResponse) SetJob(v Job) {
	o.Job = &v
}

func (o CreateEnvironmentUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserRef != nil {
		toSerialize["user_ref"] = o.UserRef
	}
	if o.Job != nil {
		toSerialize["job"] = o.Job
	}
	return json.Marshal(toSerialize)
}

type NullableCreateEnvironmentUserResponse struct {
	value *CreateEnvironmentUserResponse
	isSet bool
}

func (v NullableCreateEnvironmentUserResponse) Get() *CreateEnvironmentUserResponse {
	return v.value
}

func (v *NullableCreateEnvironmentUserResponse) Set(val *CreateEnvironmentUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEnvironmentUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEnvironmentUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEnvironmentUserResponse(val *CreateEnvironmentUserResponse) *NullableCreateEnvironmentUserResponse {
	return &NullableCreateEnvironmentUserResponse{value: val, isSet: true}
}

func (v NullableCreateEnvironmentUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEnvironmentUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


