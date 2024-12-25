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

// checks if the UpgradeDsourceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpgradeDsourceResponse{}

// UpgradeDsourceResponse struct for UpgradeDsourceResponse
type UpgradeDsourceResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewUpgradeDsourceResponse instantiates a new UpgradeDsourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradeDsourceResponse() *UpgradeDsourceResponse {
	this := UpgradeDsourceResponse{}
	return &this
}

// NewUpgradeDsourceResponseWithDefaults instantiates a new UpgradeDsourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradeDsourceResponseWithDefaults() *UpgradeDsourceResponse {
	this := UpgradeDsourceResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *UpgradeDsourceResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeDsourceResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *UpgradeDsourceResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *UpgradeDsourceResponse) SetJob(v Job) {
	o.Job = &v
}

func (o UpgradeDsourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpgradeDsourceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableUpgradeDsourceResponse struct {
	value *UpgradeDsourceResponse
	isSet bool
}

func (v NullableUpgradeDsourceResponse) Get() *UpgradeDsourceResponse {
	return v.value
}

func (v *NullableUpgradeDsourceResponse) Set(val *UpgradeDsourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradeDsourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradeDsourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradeDsourceResponse(val *UpgradeDsourceResponse) *NullableUpgradeDsourceResponse {
	return &NullableUpgradeDsourceResponse{value: val, isSet: true}
}

func (v NullableUpgradeDsourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradeDsourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


