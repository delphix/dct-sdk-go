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

// checks if the UpdateVDBGroupParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVDBGroupParameters{}

// UpdateVDBGroupParameters Parameters to update a VDB group.
type UpdateVDBGroupParameters struct {
	// The name of the VDB group.
	Name *string `json:"name,omitempty"`
	VdbIds []string `json:"vdb_ids,omitempty"`
	// Dictates order of operations on VDBs. Operations can be performed in parallel <br> for all VDBs or sequentially. Below are possible valid and invalid orderings given an example <br> VDB group with 3 vdbs (A, B, and C).<br> Valid:<br> {\"vdb_id\":\"vdb-1\", \"order\":\"1\"} {\"vdb_id\":\"vdb-2\", order:\"1\"} {vdb_id:\"vdb-3\", order:\"1\"} (parallel)<br> {vdb_id:\"vdb-1\", order:\"1\"} {vdb_id:\"vdb-2\", order:\"2\"} {vdb_id:\"vdb-3\", order:\"3\"} (sequential)<br> Invalid:<br> {vdb_id:\"vdb-1\", order:\"A\"} {vdb_id:\"vdb-2\", order:\"B\"} {vdb_id:\"vdb-3\", order:\"C\"} (sequential)<br><br> In the sequential case the vdbs with priority 1 is the first to be started and the last to<br> be stopped. This value is set on creation of VDB groups.
	Vdbs []CreateVDBGroupOrder `json:"vdbs,omitempty"`
}

// NewUpdateVDBGroupParameters instantiates a new UpdateVDBGroupParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVDBGroupParameters() *UpdateVDBGroupParameters {
	this := UpdateVDBGroupParameters{}
	return &this
}

// NewUpdateVDBGroupParametersWithDefaults instantiates a new UpdateVDBGroupParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVDBGroupParametersWithDefaults() *UpdateVDBGroupParameters {
	this := UpdateVDBGroupParameters{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateVDBGroupParameters) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBGroupParameters) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateVDBGroupParameters) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateVDBGroupParameters) SetName(v string) {
	o.Name = &v
}

// GetVdbIds returns the VdbIds field value if set, zero value otherwise.
func (o *UpdateVDBGroupParameters) GetVdbIds() []string {
	if o == nil || IsNil(o.VdbIds) {
		var ret []string
		return ret
	}
	return o.VdbIds
}

// GetVdbIdsOk returns a tuple with the VdbIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBGroupParameters) GetVdbIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.VdbIds) {
		return nil, false
	}
	return o.VdbIds, true
}

// HasVdbIds returns a boolean if a field has been set.
func (o *UpdateVDBGroupParameters) HasVdbIds() bool {
	if o != nil && !IsNil(o.VdbIds) {
		return true
	}

	return false
}

// SetVdbIds gets a reference to the given []string and assigns it to the VdbIds field.
func (o *UpdateVDBGroupParameters) SetVdbIds(v []string) {
	o.VdbIds = v
}

// GetVdbs returns the Vdbs field value if set, zero value otherwise.
func (o *UpdateVDBGroupParameters) GetVdbs() []CreateVDBGroupOrder {
	if o == nil || IsNil(o.Vdbs) {
		var ret []CreateVDBGroupOrder
		return ret
	}
	return o.Vdbs
}

// GetVdbsOk returns a tuple with the Vdbs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVDBGroupParameters) GetVdbsOk() ([]CreateVDBGroupOrder, bool) {
	if o == nil || IsNil(o.Vdbs) {
		return nil, false
	}
	return o.Vdbs, true
}

// HasVdbs returns a boolean if a field has been set.
func (o *UpdateVDBGroupParameters) HasVdbs() bool {
	if o != nil && !IsNil(o.Vdbs) {
		return true
	}

	return false
}

// SetVdbs gets a reference to the given []CreateVDBGroupOrder and assigns it to the Vdbs field.
func (o *UpdateVDBGroupParameters) SetVdbs(v []CreateVDBGroupOrder) {
	o.Vdbs = v
}

func (o UpdateVDBGroupParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVDBGroupParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.VdbIds) {
		toSerialize["vdb_ids"] = o.VdbIds
	}
	if !IsNil(o.Vdbs) {
		toSerialize["vdbs"] = o.Vdbs
	}
	return toSerialize, nil
}

type NullableUpdateVDBGroupParameters struct {
	value *UpdateVDBGroupParameters
	isSet bool
}

func (v NullableUpdateVDBGroupParameters) Get() *UpdateVDBGroupParameters {
	return v.value
}

func (v *NullableUpdateVDBGroupParameters) Set(val *UpdateVDBGroupParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVDBGroupParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVDBGroupParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVDBGroupParameters(val *UpdateVDBGroupParameters) *NullableUpdateVDBGroupParameters {
	return &NullableUpdateVDBGroupParameters{value: val, isSet: true}
}

func (v NullableUpdateVDBGroupParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVDBGroupParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


