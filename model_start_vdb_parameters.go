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

// checks if the StartVDBParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartVDBParameters{}

// StartVDBParameters Parameters to start a VDB.
type StartVDBParameters struct {
	// List of specific Oracle Virtual Database Instances to start.
	Instances []int32 `json:"instances,omitempty"`
}

// NewStartVDBParameters instantiates a new StartVDBParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartVDBParameters() *StartVDBParameters {
	this := StartVDBParameters{}
	return &this
}

// NewStartVDBParametersWithDefaults instantiates a new StartVDBParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartVDBParametersWithDefaults() *StartVDBParameters {
	this := StartVDBParameters{}
	return &this
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *StartVDBParameters) GetInstances() []int32 {
	if o == nil || IsNil(o.Instances) {
		var ret []int32
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartVDBParameters) GetInstancesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *StartVDBParameters) HasInstances() bool {
	if o != nil && !IsNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []int32 and assigns it to the Instances field.
func (o *StartVDBParameters) SetInstances(v []int32) {
	o.Instances = v
}

func (o StartVDBParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartVDBParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}
	return toSerialize, nil
}

type NullableStartVDBParameters struct {
	value *StartVDBParameters
	isSet bool
}

func (v NullableStartVDBParameters) Get() *StartVDBParameters {
	return v.value
}

func (v *NullableStartVDBParameters) Set(val *StartVDBParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableStartVDBParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableStartVDBParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartVDBParameters(val *StartVDBParameters) *NullableStartVDBParameters {
	return &NullableStartVDBParameters{value: val, isSet: true}
}

func (v NullableStartVDBParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartVDBParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


