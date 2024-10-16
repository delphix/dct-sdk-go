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

// checks if the CreateVDBGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVDBGroupResponse{}

// CreateVDBGroupResponse struct for CreateVDBGroupResponse
type CreateVDBGroupResponse struct {
	VdbGroup *VDBGroup `json:"vdb_group,omitempty"`
}

// NewCreateVDBGroupResponse instantiates a new CreateVDBGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVDBGroupResponse() *CreateVDBGroupResponse {
	this := CreateVDBGroupResponse{}
	return &this
}

// NewCreateVDBGroupResponseWithDefaults instantiates a new CreateVDBGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVDBGroupResponseWithDefaults() *CreateVDBGroupResponse {
	this := CreateVDBGroupResponse{}
	return &this
}

// GetVdbGroup returns the VdbGroup field value if set, zero value otherwise.
func (o *CreateVDBGroupResponse) GetVdbGroup() VDBGroup {
	if o == nil || IsNil(o.VdbGroup) {
		var ret VDBGroup
		return ret
	}
	return *o.VdbGroup
}

// GetVdbGroupOk returns a tuple with the VdbGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVDBGroupResponse) GetVdbGroupOk() (*VDBGroup, bool) {
	if o == nil || IsNil(o.VdbGroup) {
		return nil, false
	}
	return o.VdbGroup, true
}

// HasVdbGroup returns a boolean if a field has been set.
func (o *CreateVDBGroupResponse) HasVdbGroup() bool {
	if o != nil && !IsNil(o.VdbGroup) {
		return true
	}

	return false
}

// SetVdbGroup gets a reference to the given VDBGroup and assigns it to the VdbGroup field.
func (o *CreateVDBGroupResponse) SetVdbGroup(v VDBGroup) {
	o.VdbGroup = &v
}

func (o CreateVDBGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVDBGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VdbGroup) {
		toSerialize["vdb_group"] = o.VdbGroup
	}
	return toSerialize, nil
}

type NullableCreateVDBGroupResponse struct {
	value *CreateVDBGroupResponse
	isSet bool
}

func (v NullableCreateVDBGroupResponse) Get() *CreateVDBGroupResponse {
	return v.value
}

func (v *NullableCreateVDBGroupResponse) Set(val *CreateVDBGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVDBGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVDBGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVDBGroupResponse(val *CreateVDBGroupResponse) *NullableCreateVDBGroupResponse {
	return &NullableCreateVDBGroupResponse{value: val, isSet: true}
}

func (v NullableCreateVDBGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVDBGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


