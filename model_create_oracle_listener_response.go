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

// checks if the CreateOracleListenerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOracleListenerResponse{}

// CreateOracleListenerResponse struct for CreateOracleListenerResponse
type CreateOracleListenerResponse struct {
	// The reference of the created Oracle listener
	ListenerRef *string `json:"listener_ref,omitempty"`
	Job *Job `json:"job,omitempty"`
}

// NewCreateOracleListenerResponse instantiates a new CreateOracleListenerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOracleListenerResponse() *CreateOracleListenerResponse {
	this := CreateOracleListenerResponse{}
	return &this
}

// NewCreateOracleListenerResponseWithDefaults instantiates a new CreateOracleListenerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOracleListenerResponseWithDefaults() *CreateOracleListenerResponse {
	this := CreateOracleListenerResponse{}
	return &this
}

// GetListenerRef returns the ListenerRef field value if set, zero value otherwise.
func (o *CreateOracleListenerResponse) GetListenerRef() string {
	if o == nil || IsNil(o.ListenerRef) {
		var ret string
		return ret
	}
	return *o.ListenerRef
}

// GetListenerRefOk returns a tuple with the ListenerRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOracleListenerResponse) GetListenerRefOk() (*string, bool) {
	if o == nil || IsNil(o.ListenerRef) {
		return nil, false
	}
	return o.ListenerRef, true
}

// HasListenerRef returns a boolean if a field has been set.
func (o *CreateOracleListenerResponse) HasListenerRef() bool {
	if o != nil && !IsNil(o.ListenerRef) {
		return true
	}

	return false
}

// SetListenerRef gets a reference to the given string and assigns it to the ListenerRef field.
func (o *CreateOracleListenerResponse) SetListenerRef(v string) {
	o.ListenerRef = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *CreateOracleListenerResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOracleListenerResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *CreateOracleListenerResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *CreateOracleListenerResponse) SetJob(v Job) {
	o.Job = &v
}

func (o CreateOracleListenerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOracleListenerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ListenerRef) {
		toSerialize["listener_ref"] = o.ListenerRef
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableCreateOracleListenerResponse struct {
	value *CreateOracleListenerResponse
	isSet bool
}

func (v NullableCreateOracleListenerResponse) Get() *CreateOracleListenerResponse {
	return v.value
}

func (v *NullableCreateOracleListenerResponse) Set(val *CreateOracleListenerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOracleListenerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOracleListenerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOracleListenerResponse(val *CreateOracleListenerResponse) *NullableCreateOracleListenerResponse {
	return &NullableCreateOracleListenerResponse{value: val, isSet: true}
}

func (v NullableCreateOracleListenerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOracleListenerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


