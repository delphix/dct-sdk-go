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

// checks if the DeleteCDBParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteCDBParameters{}

// DeleteCDBParameters Parameters to delete a CDB.
type DeleteCDBParameters struct {
	// Whether to continue the operation upon failures.
	Force *bool `json:"force,omitempty"`
}

// NewDeleteCDBParameters instantiates a new DeleteCDBParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteCDBParameters() *DeleteCDBParameters {
	this := DeleteCDBParameters{}
	var force bool = false
	this.Force = &force
	return &this
}

// NewDeleteCDBParametersWithDefaults instantiates a new DeleteCDBParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteCDBParametersWithDefaults() *DeleteCDBParameters {
	this := DeleteCDBParameters{}
	var force bool = false
	this.Force = &force
	return &this
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *DeleteCDBParameters) GetForce() bool {
	if o == nil || IsNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCDBParameters) GetForceOk() (*bool, bool) {
	if o == nil || IsNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *DeleteCDBParameters) HasForce() bool {
	if o != nil && !IsNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *DeleteCDBParameters) SetForce(v bool) {
	o.Force = &v
}

func (o DeleteCDBParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteCDBParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	return toSerialize, nil
}

type NullableDeleteCDBParameters struct {
	value *DeleteCDBParameters
	isSet bool
}

func (v NullableDeleteCDBParameters) Get() *DeleteCDBParameters {
	return v.value
}

func (v *NullableDeleteCDBParameters) Set(val *DeleteCDBParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteCDBParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteCDBParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteCDBParameters(val *DeleteCDBParameters) *NullableDeleteCDBParameters {
	return &NullableDeleteCDBParameters{value: val, isSet: true}
}

func (v NullableDeleteCDBParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteCDBParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


