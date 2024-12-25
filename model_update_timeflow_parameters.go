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

// checks if the UpdateTimeflowParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTimeflowParameters{}

// UpdateTimeflowParameters Parameters to update a Timeflow.
type UpdateTimeflowParameters struct {
	// The name of the timeflow.
	Name *string `json:"name,omitempty"`
}

// NewUpdateTimeflowParameters instantiates a new UpdateTimeflowParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTimeflowParameters() *UpdateTimeflowParameters {
	this := UpdateTimeflowParameters{}
	return &this
}

// NewUpdateTimeflowParametersWithDefaults instantiates a new UpdateTimeflowParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTimeflowParametersWithDefaults() *UpdateTimeflowParameters {
	this := UpdateTimeflowParameters{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateTimeflowParameters) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTimeflowParameters) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateTimeflowParameters) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateTimeflowParameters) SetName(v string) {
	o.Name = &v
}

func (o UpdateTimeflowParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTimeflowParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUpdateTimeflowParameters struct {
	value *UpdateTimeflowParameters
	isSet bool
}

func (v NullableUpdateTimeflowParameters) Get() *UpdateTimeflowParameters {
	return v.value
}

func (v *NullableUpdateTimeflowParameters) Set(val *UpdateTimeflowParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTimeflowParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTimeflowParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTimeflowParameters(val *UpdateTimeflowParameters) *NullableUpdateTimeflowParameters {
	return &NullableUpdateTimeflowParameters{value: val, isSet: true}
}

func (v NullableUpdateTimeflowParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTimeflowParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


