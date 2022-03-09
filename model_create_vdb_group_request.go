/*
Delphix API Gateway

Delphix API Gateway API

API version: 1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CreateVDBGroupRequest struct for CreateVDBGroupRequest
type CreateVDBGroupRequest struct {
	Name string `json:"name"`
	VdbIds []string `json:"vdb_ids"`
}

// NewCreateVDBGroupRequest instantiates a new CreateVDBGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVDBGroupRequest(name string, vdbIds []string) *CreateVDBGroupRequest {
	this := CreateVDBGroupRequest{}
	this.Name = name
	this.VdbIds = vdbIds
	return &this
}

// NewCreateVDBGroupRequestWithDefaults instantiates a new CreateVDBGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVDBGroupRequestWithDefaults() *CreateVDBGroupRequest {
	this := CreateVDBGroupRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateVDBGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateVDBGroupRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateVDBGroupRequest) SetName(v string) {
	o.Name = v
}

// GetVdbIds returns the VdbIds field value
func (o *CreateVDBGroupRequest) GetVdbIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.VdbIds
}

// GetVdbIdsOk returns a tuple with the VdbIds field value
// and a boolean to check if the value has been set.
func (o *CreateVDBGroupRequest) GetVdbIdsOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VdbIds, true
}

// SetVdbIds sets field value
func (o *CreateVDBGroupRequest) SetVdbIds(v []string) {
	o.VdbIds = v
}

func (o CreateVDBGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["vdb_ids"] = o.VdbIds
	}
	return json.Marshal(toSerialize)
}

type NullableCreateVDBGroupRequest struct {
	value *CreateVDBGroupRequest
	isSet bool
}

func (v NullableCreateVDBGroupRequest) Get() *CreateVDBGroupRequest {
	return v.value
}

func (v *NullableCreateVDBGroupRequest) Set(val *CreateVDBGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVDBGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVDBGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVDBGroupRequest(val *CreateVDBGroupRequest) *NullableCreateVDBGroupRequest {
	return &NullableCreateVDBGroupRequest{value: val, isSet: true}
}

func (v NullableCreateVDBGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVDBGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


